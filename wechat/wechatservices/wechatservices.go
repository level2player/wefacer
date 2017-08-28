package wechatservices

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"wefacer/core"
)

func ReceiveRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Receive Msg.....")
	defer r.Body.Close()
	checkWeChatSignature(w, r)
	// if checkWeChatSignature(w, r) {
	// 	recognitionservices.HandleMsg(r, func(resbuffer []byte) {
	// 		fmt.Fprintf(w, string(resbuffer))
	// 	})
	// }

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
