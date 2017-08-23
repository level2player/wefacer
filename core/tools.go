package core
import(
	"io"
	"crypto/sha1"
     "fmt"
     "encoding/base64"     
     "io/ioutil"
)
//对字符串进行哈希加密
func Str2sha1(data string)string{
    t:=sha1.New()
    io.WriteString(t,data)
    return fmt.Sprintf("%x",t.Sum(nil))
}

//将图片转换成base64格式
func GetImageUrlBase64(path string)string{
     picbyte,_:=ioutil.ReadFile(path)
     return base64.StdEncoding.EncodeToString(picbyte)
}

func GetImageBase64(content []byte)string{
     return base64.StdEncoding.EncodeToString(content)
}

