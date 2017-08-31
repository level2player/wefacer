package core

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//对字符串进行哈希加密
func Str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//下载http 图片转换成base64str返回
func GetImageUrlBase64(httpurl string) (base64str string, err error) {
	resp, err1 := http.Get(httpurl)
	defer resp.Body.Close()
	if err1 != nil {
		err = err1
		return
	}
	content, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
		return
	}
	base64str = GetImageBase64(content)
	return
}

//buffer转换成base64格式
func GetImageBase64(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}
