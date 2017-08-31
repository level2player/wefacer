package main

import (
	"log"
	"net/http"
	"wefacer/core"
	"wefacer/wechat/wechatservices"
)

func main() {
	core.Print_log("start wefacer ...")
	http.HandleFunc("/", wechatservices.ReceiveRequest)
	err := http.ListenAndServe(":"+core.WefacerConfig.ConfigMap["port"], nil)
	if err != nil {
		core.Print_log(err.Error())
		log.Fatal("ListenAndServe: ", err)
	}
}
