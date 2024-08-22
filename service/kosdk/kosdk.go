package kosdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/api/fs"
	"github.com/woocoos/msgcenter/ent"
	urlx "net/url"
	"path/filepath"
	"strconv"
)

func NewSDK(cfg *conf.Configuration, db *ent.Client) (*api.SDK, error) {
	kosdk, err := api.NewSDK(cfg)
	if err != nil {
		return nil, err
	}
	ret, _, err := kosdk.Fs().FileIdentityAPI.GetFileIdentities(context.Background(), fs.GetFileIdentitiesRequest{IsDefault: aws.Bool(true)})
	if err != nil {
		return nil, err
	}
	for _, fi := range ret {
		err = kosdk.Fs().RegistryProvider(fs.ToProviderConfig(fi), fi.TenantID.String())
		if err != nil {
			return nil, err
		}
	}
	return kosdk, nil
}

// DefaultFilePath 存储默认文件路径
func DefaultFilePath(tenantID int, url, baseDir, dataDir string) (string, error) {
	u, err := urlx.Parse(url)
	if err != nil {
		return "", err
	}
	ext := filepath.Ext(u.Path)
	if err != nil {
		return "", err
	}
	fileName := MD5String([]byte(url)) + ext
	localPath := filepath.Join(baseDir, strconv.Itoa(tenantID), dataDir, fileName)
	return localPath, nil
}

// MD5String 计算md5
func MD5String(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
