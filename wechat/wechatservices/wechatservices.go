package wechatservices

import(
	  "net/http"
	  "wefacer/core"
	  "wefacer/recognition/recognitionservices"
	  "fmt"
	  "log"
	  "strings"
	  "sort"
)

func ReceiveRequest(w http.ResponseWriter,r *http.Request){
	  if r!=nil{
		if checkSignature(w,r){
			defer r.Body.Close()
			//消息处理
			recognitionservices.HandleMsg(r,func(resbuffer []byte){
				log.Println(string(resbuffer))
				fmt.Fprintf(w,string(resbuffer))
			})
		}
	  }
}



func checkSignature(w http.ResponseWriter,r *http.Request)bool{
    r.ParseForm()
    token :="lsy_token"
    signature:=strings.Join(r.Form["signature"],"")
     timestamp :=strings.Join(r.Form["timestamp"],"")
    var nonce string=strings.Join(r.Form["nonce"],"")
    var echostr string=strings.Join(r.Form["echostr"],"")
    tmps:=[]string{token,timestamp,nonce}
    sort.Strings(tmps)
    tmpStr:=tmps[0]+tmps[1]+tmps[2]
    tmp:=core.Str2sha1(tmpStr)
    if tmp==signature{
		fmt.Fprintf(w,echostr)
		return true
	}else{
		fmt.Println("token error")
	}
	return false
}


		

