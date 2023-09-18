package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
)

// RandomStr generate random string,exclude 0,i,l
func RandomStr(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	var result string
	for i := 0; i < n; i++ {
		randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		randomChar := charset[randomInt.Int64()]
		result += string(randomChar)
	}
	return result
}

func SHA256(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	return hex.EncodeToString(sha.Sum(nil))
}

// RemoveSliceElement 移除数组指定元素
func RemoveSliceElement[T int | int64 | string | float32 | float64](a []T, el T) []T {
	j := 0
	for _, v := range a {
		if v != el {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// UpdateSliceElement 更新数组指定元素
func UpdateSliceElement[T int | int64 | string | float32 | float64](a []T, newEl T, oldEl T) []T {
	for i, v := range a {
		if v == oldEl {
			a[i] = newEl
		}
	}
	return a
}

// DiffArrays 查找两数组新增及删除的元素: a:新数组  b:旧数组
func DiffArrays[T int | int64 | string | float32 | float64](a []T, b []T) ([]T, []T) {
	var added, removed []T
	hash := make(map[T]bool)
	for _, num := range a {
		hash[num] = true
	}
	// 删除的元素
	for _, num := range b {
		if _, ok := hash[num]; !ok {
			removed = append(removed, num)
		} else {
			// 如果元素存在，则删除该元素
			delete(hash, num)
		}
	}
	// 新增的元素
	for num := range hash {
		added = append(added, num)
	}
	return added, removed
}

// RemoveDuplicateElement 去重
func RemoveDuplicateElement[T int | int64 | string | float32 | float64](arr []T) []T {
	if arr == nil {
		return nil
	}
	temp := make(map[string]bool)
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		key := fmt.Sprint(v)
		if _, ok := temp[key]; !ok {
			temp[key] = true
			result = append(result, v)
		}
	}
	return result
}

// MaskEmail 邮箱脱敏处理
func MaskEmail(email string) string {
	// 使用正则表达式匹配邮箱地址的用户名部分
	re := regexp.MustCompile(`([^@]+)@`)
	matches := re.FindStringSubmatch(email)
	if len(matches) != 2 {
		return email
	}
	// 获取用户名部分
	un := matches[0]
	// 保留前三个字符
	mun := un[:3] + "****"
	// 替换原始邮箱地址中的用户名部分为脱敏后的用户名
	maskedEmail := re.ReplaceAllString(email, mun+"@")
	return maskedEmail
}
