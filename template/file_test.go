package template

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/woocoos/knockout-go/pkg/identity"
	"testing"
)

func TestTemplate_ValidateFilePath(t *testing.T) {
	t.Run("relative", func(t *testing.T) {
		tpl := &Template{
			Config: Config{
				BaseDir:       "/base",
				DataDir:       "data",
				TmpDir:        "/tmp",
				AttachmentDir: "/attachment",
			},
		}
		ctx := identity.WithTenantID(context.Background(), 1)
		err := tpl.ValidateFilePath(ctx, "/data/1/1.txt", TplPathKindData)
		assert.NoError(t, err)
	})
}
