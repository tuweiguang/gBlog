package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gBlog/common"
	"strconv"
)

// ToInt64 类型转换，获得int64
func ToInt64(v interface{}) (re int64, err error) {
	switch v.(type) {
	case string:
		re, err = strconv.ParseInt(v.(string), 10, 64)
	case float64:
		re = int64(v.(float64))
	case float32:
		re = int64(v.(float32))
	case int64:
		re = v.(int64)
	case int32:
		re = v.(int64)
	default:
		err = errors.New("不能转换")
	}
	return
}

// 加盐加密
// 加密(password+盐)
func PasswordMD5(password, username string) string {
	h := md5.New()
	h.Write([]byte(password + username + common.PasswordSalt))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}
