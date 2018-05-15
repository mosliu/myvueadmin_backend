package libs

import (
    "crypto/md5"
    "encoding/hex"
)

func Md5(buf []byte) string {
    h := md5.New()
    h.Write(buf)
    rs := hex.EncodeToString(h.Sum(nil))
    return rs
}

//create md5 string
func Str2Md5(s string) string {
    return Md5([]byte(s))
}

//password hash function
func Pwdhash(str string) string {
    return Str2Md5(str)
}