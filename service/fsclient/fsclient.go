package fsclient

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/api/fs"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/fileidentity"
	urlx "net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Client struct {
	db        *ent.Client
	kosdk     *api.SDK
	providers map[int]S3Provider
}

func NewClient(db *ent.Client, kosdk *api.SDK) (*Client, error) {
	client := &Client{
		db:    db,
		kosdk: kosdk,
	}
	providers, err := client.loadProviders()
	if err != nil {
		return nil, err
	}
	client.providers = providers
	return client, nil
}

func (c *Client) loadProviders() (map[int]S3Provider, error) {
	fis, err := c.db.FileIdentity.Query().Where(fileidentity.IsDefault(true)).WithSource().All(context.Background())
	if err != nil {
		return nil, err
	}
	// 获取租户的s3实例
	providers := make(map[int]S3Provider)
	for _, fi := range fis {
		config := c.toProviderConfig(fi)
		provider, err := c.kosdk.Fs().GetProvider(config)
		if err != nil {
			return nil, err
		}
		providers[fi.TenantID] = S3Provider{
			S3Provider:     provider,
			ProviderConfig: config,
			TenantID:       fi.TenantID,
		}
	}
	return providers, nil
}

func (c *Client) toProviderConfig(fi *ent.FileIdentity) *fs.ProviderConfig {
	return &fs.ProviderConfig{
		Kind:              fs.Kind(fi.Edges.Source.Kind.String()),
		Bucket:            fi.Edges.Source.Bucket,
		BucketUrl:         fi.Edges.Source.BucketURL,
		Endpoint:          fi.Edges.Source.Endpoint,
		EndpointImmutable: fi.Edges.Source.EndpointImmutable,
		StsEndpoint:       fi.Edges.Source.StsEndpoint,
		AccessKeyID:       fi.AccessKeyID,
		AccessKeySecret:   fi.AccessKeySecret,
		Policy:            fi.Policy,
		Region:            fi.Edges.Source.Region,
		RoleArn:           fi.RoleArn,
		DurationSeconds:   fi.DurationSeconds,
	}
}

// GetProvider 获取租户的s3实例
func (c *Client) GetProvider(tenantID int) S3Provider {
	return c.providers[tenantID]
}

type S3Provider struct {
	fs.S3Provider
	ProviderConfig *fs.ProviderConfig
	TenantID       int
}

// DefaultDownloadObject 下载对象到本地，消息中心默认处理方式
func (p *S3Provider) DefaultDownloadObject(url, baseDir, dataDir string) (string, error) {
	key, err := p.ParseUrlKey(url)
	if err != nil {
		return "", err
	}
	localPath, err := DefaultFilePath(p.TenantID, url, baseDir, dataDir)
	if err != nil {
		return "", err
	}
	// 文件存在则不下载
	if FileExists(localPath) {
		return "", nil
	}
	return localPath, p.downloadObject(key, localPath)
}

func (p *S3Provider) downloadObject(key string, localPath string) error {
	// 获取对象
	getObjOutput, err := p.S3Client().GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(p.ProviderConfig.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer getObjOutput.Body.Close()
	// 创建本地文件夹（如果不存在）
	err = os.MkdirAll(filepath.Dir(localPath), os.ModePerm)
	if err != nil {
		return err
	}
	// 保存到本地文件
	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.ReadFrom(getObjOutput.Body)
	if err != nil {
		return err
	}
	return nil
}

// ParseUrlKey 从url解析出key
func (p *S3Provider) ParseUrlKey(url string) (key string, err error) {
	u, err := urlx.Parse(url)
	if err != nil {
		return
	}
	if p.ProviderConfig.Kind == fs.KindMinio {
		key = strings.TrimPrefix(u.Path, "/"+p.ProviderConfig.Bucket)
	} else {
		key = strings.TrimPrefix(u.Path, "/")
	}
	return key, nil
}

// DefaultFileName 存储默认文件名
func DefaultFileName(url string) (string, error) {
	ext, err := ParseUrlExt(url)
	if err != nil {
		return "", err
	}
	return MD5String([]byte(url)) + ext, nil
}

// DefaultFilePath 存储默认文件路径
func DefaultFilePath(tenantID int, url, baseDir, dataDir string) (string, error) {
	fileName, err := DefaultFileName(url)
	if err != nil {
		return "", err
	}
	localPath := filepath.Join(baseDir, strconv.Itoa(tenantID), dataDir, fileName)
	return localPath, nil
}

// ParseUrlExt 从url解析出文件扩展名
func ParseUrlExt(url string) (ext string, err error) {
	u, err := urlx.Parse(url)
	if err != nil {
		return
	}
	return filepath.Ext(u.Path), nil
}

// MD5String 计算md5
func MD5String(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// FileExists 检查文件是否存在
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
