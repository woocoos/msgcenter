package template

import (
	"github.com/woocoos/knockout-go/api/fs"
	"github.com/woocoos/msgcenter/service/kosdk"
	"os"
	"strconv"
)

// EnableTplFile 启用模板文件
// tpl 模板地址
func (t *Template) EnableTplFile(tpl string, tenantID int) error {
	if tpl == "" {
		return nil
	}
	localFile, err := kosdk.DefaultFilePath(tenantID, tpl, t.BaseDir, t.DataDir)
	if err != nil {
		return err
	}
	err = t.KOSdk.Fs().DownloadObjectByKey(strconv.Itoa(tenantID), tpl, localFile, fs.WithOverwrittenFile(false))
	if err != nil {
		return err
	}
	// 加载模板
	_, err = t.ParseFiles(localFile)
	if err != nil {
		return err
	}
	return nil
}

// RemoveTplFile 移除模板文件
// tpl 模板地址
func (t *Template) RemoveTplFile(tpl string, tenantID int) error {
	if tpl == "" {
		return nil
	}
	tplPath, err := kosdk.DefaultFilePath(tenantID, tpl, t.BaseDir, t.DataDir)
	if err != nil {
		return err
	}
	// 移除已删除的模板
	err = t.RemoveTemplates(tplPath)
	if err != nil {
		return err
	}
	// 将文件删除
	_, err = os.Stat(tplPath)
	if err == nil {
		err = os.Remove(tplPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// EnableAttachFile 启用模板文件
// tpl 模板地址
func (t *Template) EnableAttachFile(attachments []string, tenantID int) error {
	if len(attachments) == 0 {
		return nil
	}
	for _, att := range attachments {
		if att == "" {
			continue
		}
		localFile, err := kosdk.DefaultFilePath(tenantID, att, t.BaseDir, t.AttachmentDir)
		if err != nil {
			return err
		}
		err = t.KOSdk.Fs().DownloadObjectByKey(strconv.Itoa(tenantID), att, localFile, fs.WithOverwrittenFile(false))
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveAttachFile 移除模板文件
// tpl 模板地址
func (t *Template) RemoveAttachFile(attachments []string, tenantID int) error {
	if len(attachments) == 0 {
		return nil
	}
	// 删除附件
	for _, att := range attachments {
		if att == "" {
			continue
		}
		dataPath, err := kosdk.DefaultFilePath(tenantID, att, t.BaseDir, t.AttachmentDir)
		if err != nil {
			return err
		}
		_, err = os.Stat(dataPath)
		if err == nil {
			err = os.Remove(dataPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
