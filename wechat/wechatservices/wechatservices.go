package wechatservices

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"wefacer/core"
	"wefacer/recognition/recognitionservices"
)

func ReceiveRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Host=" + r.Host)
	log.Println("RemoteAddr=" + r.RemoteAddr)
	log.Println("RequestURI=" + r.RequestURI)
	if checkWeChatSignature(w, r) {
		log.Println("receive msg....begin handle msg")
		recognitionservices.HandleMsg(r, func(resbuffer []byte) {
			fmt.Println(string(resbuffer))
			fmt.Fprintf(w, string(resbuffer))
		})
	}

}

func checkWeChatSignature(w http.ResponseWriter, r *http.Request) bool {
	if r != nil {
		r.ParseForm()
		token := "lsy_token"
		signature := strings.Join(r.Form["signature"], "")
		timestamp := strings.Join(r.Form["timestamp"], "")
		nonce := strings.Join(r.Form["nonce"], "")
		echostr := strings.Join(r.Form["echostr"], "")
		tmps := []string{token, timestamp, nonce}
		sort.Strings(tmps)
		tmpStr := tmps[0] + tmps[1] + tmps[2]
		tmp := core.Str2sha1(tmpStr)
		if tmp == signature {
			fmt.Fprintf(w, echostr)
			return true
		} else {
			log.Println("wechat token validation error")
			return false
		}
	}
	log.Println("wechat token validation error")
	return false
}
