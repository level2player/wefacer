package wechatservices

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"wefacer/core"
	"wefacer/recognition/recognitionservices"
)

func ReceiveRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	core.Print_log("receive msg....,Hoset=%s,RemoteAddr=%s,RequestURI=%s", r.Host, r.RemoteAddr, r.RequestURI)
	if check_wechat_signature(w, r) {
		recognitionservices.HandleMsg(r, func(resbuffer []byte) {
			core.Print_log("ask response,reulst=%s", string(resbuffer))
			fmt.Fprintf(w, string(resbuffer))
		})
	}

}

func check_wechat_signature(w http.ResponseWriter, r *http.Request) bool {
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
			core.Print_log("check_wechat_signature success")
			fmt.Fprintf(w, echostr)
			return true
		} else {
			core.Print_log("wechat token validation error")
			return false
		}
	}
	core.Print_log("wechat token validation error")
	return false
}
