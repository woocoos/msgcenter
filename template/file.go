package template

import (
	"context"
	"fmt"
	"github.com/woocoos/knockout-go/pkg/identity"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type TplPathKind int

const (
	TplPathKindData TplPathKind = iota
	TplPathKindTmp
	TplPathKindAttachment
)

// ValidateFilePath Verify compliance with file service conventions.
// path rule: {PathKind}/{tenantId}/{xxx}/{filename}
func (t *Template) ValidateFilePath(ctx context.Context, path string, kind TplPathKind) error {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}

	path = strings.TrimPrefix(path, "/")
	var rp string
	switch kind {
	case TplPathKindData:
		rp = t.DataDir
	case TplPathKindTmp:
		rp = t.TmpDir
	case TplPathKindAttachment:
		rp = t.AttachmentDir
	}
	prefixPath := filepath.Join(rp, strconv.Itoa(tid)) + "/"
	if !strings.HasPrefix(path, prefixPath) {
		return fmt.Errorf("invalid path: %s,must be like:%s/xxx", path, prefixPath)
	}
	return nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// GetTplDataPath 将tpl临时文件路径转为正式存储路径
func (t *Template) GetTplDataPath(tempPath string) string {
	return filepath.Join(
		t.BaseDir,
		t.DataDir,
		strings.TrimPrefix(
			strings.TrimPrefix(tempPath, "/"),
			strings.TrimPrefix(t.TmpDir, "/"),
		),
	)
}

// GetTplTempPath 获取tpl正式文件路径
func (t *Template) GetTplTempPath(tempPath string) string {
	return filepath.Join(t.BaseDir, tempPath)
}

// EnableTplDataFile 启用模板文件
// tplPath 为temp文件路径
func (t *Template) EnableTplDataFile(tplPath string) error {
	if tplPath == "" {
		return nil
	}
	// 将temp文件复制到data目录下
	distName := t.GetTplDataPath(tplPath)
	_, err := CopyFile(distName, t.GetTplTempPath(tplPath))
	if err != nil {
		return err
	}
	// 加载模板
	_, err = t.ParseFiles(distName)
	if err != nil {
		return err
	}
	return nil
}

// RemoveTplDataFile 移除data目录模板
// tplPath 为temp文件路径
func (t *Template) RemoveTplDataFile(tplPath string) error {
	if tplPath == "" {
		return nil
	}
	// 将文件从data目录下删除
	dataPath := t.GetTplDataPath(tplPath)
	_, err := os.Stat(dataPath)
	if err == nil {
		err = os.Remove(dataPath)
		if err != nil {
			return err
		}
	}
	return nil
}
