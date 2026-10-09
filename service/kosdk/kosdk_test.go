package kosdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveMountPath_NilSDK(t *testing.T) {
	mountPaths := map[string]string{"mybucket": "/mnt/oss/mybucket"}
	path, ok := ResolveMountPath(nil, mountPaths, "1", "https://mybucket.oss-cn-hangzhou.aliyuncs.com/file.pdf")
	assert.False(t, ok)
	assert.Empty(t, path)
}

func TestResolveMountPath_EmptyMountPaths(t *testing.T) {
	path, ok := ResolveMountPath(nil, nil, "1", "https://mybucket.oss-cn-hangzhou.aliyuncs.com/file.pdf")
	assert.False(t, ok)
	assert.Empty(t, path)
}
