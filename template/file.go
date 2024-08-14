package template

import (
	"github.com/woocoos/msgcenter/service/fsclient"
	"os"
)

// EnableTplFile 启用模板文件
// tpl 模板地址
func (t *Template) EnableTplFile(tpl string, tenantID int) error {
	if tpl == "" {
		return nil
	}
	provider := t.FSClient.GetProvider(tenantID)
	tplPath, err := provider.DefaultDownloadObject(tpl, t.BaseDir, t.DataDir)
	if err != nil {
		return err
	}
	// 加载模板
	_, err = t.ParseFiles(tplPath)
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
	tplPath, err := fsclient.DefaultFilePath(tenantID, tpl, t.BaseDir, t.DataDir)
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
	provider := t.FSClient.GetProvider(tenantID)
	for _, att := range attachments {
		if att == "" {
			continue
		}
		_, err := provider.DefaultDownloadObject(att, t.BaseDir, t.AttachmentDir)
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
		dataPath, err := fsclient.DefaultFilePath(tenantID, att, t.BaseDir, t.AttachmentDir)
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
