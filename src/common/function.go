package common

import (
    "crypto/md5"
    "encoding/hex"
    "math/rand"
    "strconv"
)
func IsNull(args []string) bool{
    for _, v := range args {
        if v == "" {
            return true
        }
    }
    return false
}

func IsNil(err error) {
    if err == nil {
        panic("error")
    }
}

//返回一个加密后的md5字符串
func STMd5(s string) string{
    h := md5.New()
    h.Write([]byte(s)) // 需要加密的字符串为 123456
    cipherStr := h.Sum(nil)
    return hex.EncodeToString(cipherStr)
}

//返回一个随机数 string
func RandNum() string{
    return strconv.Itoa(rand.Intn(999999))
}
