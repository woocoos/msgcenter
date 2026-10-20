package kosdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/api/fs"
	"github.com/woocoos/knockout-go/api/fs/alioss"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
)

func NewSDK(cfg *conf.Configuration, db *ent.Client) (*api.SDK, error) {
	kosdk, err := api.NewSDK(cfg)
	if err != nil {
		return nil, err
	}
	// 获取所有tenantID
	tenantIDs, err := db.MsgTemplate.Query().Where(msgtemplate.TenantIDNotNil()).GroupBy(msgtemplate.FieldTenantID).Ints(context.Background())
	if err != nil {
		return nil, err
	}
	ret, _, err := kosdk.Fs().FileIdentityAPI.GetFileIdentities(context.Background(), &fs.GetFileIdentitiesRequest{TenantIDs: tenantIDs, IsDefault: true})
	if err != nil {
		return nil, err
	}
	fs.RegisterS3Provider(fs.KindAliOSS, alioss.BuildProvider)
	for _, fi := range ret {
		err = kosdk.Fs().RegistryProvider(fs.ToProviderConfig(fi), fi.TenantID.String())
		if err != nil {
			return nil, err
		}
	}
	return kosdk, nil
}

// DefaultFilePath 存储默认文件路径
func DefaultFilePath(tenantID int, rawURL, baseDir, dataDir string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	ext := filepath.Ext(u.Path)
	fileName := MD5String([]byte(rawURL)) + ext
	localPath := filepath.Join(baseDir, strconv.Itoa(tenantID), dataDir, fileName)
	return localPath, nil
}

// ResolveMountPath resolves a remote storage URL to a local mount path.
// If the URL matches a registered provider's BucketUrl, it looks up the bucket
// in mountPaths and returns the local path as {mountPath}/{objectKey}.
// Returns empty string and false if no matching mount path is found.
func ResolveMountPath(sdk *api.SDK, mountPaths map[string]string, tenantID, rawURL string) (string, bool) {
	if len(mountPaths) == 0 || sdk == nil {
		return "", false
	}
	provider, err := sdk.Fs().GetProviderByBizKey(tenantID)
	if err != nil {
		return "", false
	}
	pcfg := provider.ProviderConfig()
	bucketURL := pcfg.BucketUrl
	if bucketURL == "" {
		return "", false
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", false
	}
	bu, err := url.Parse(bucketURL)
	if err != nil {
		return "", false
	}
	// Match scheme and host to determine if the URL belongs to this bucket.
	if !strings.EqualFold(u.Scheme, bu.Scheme) || !strings.EqualFold(u.Host, bu.Host) {
		return "", false
	}
	// Look up mount path by bucket name.
	mountPath, ok := mountPaths[pcfg.Bucket]
	if !ok || mountPath == "" {
		return "", false
	}
	objectKey, err := provider.ParseUrlKey(rawURL)
	if err != nil || objectKey == "" {
		return "", false
	}
	return filepath.Join(mountPath, objectKey), true
}

// MD5String 计算md5
func MD5String(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
