package main

import (
	"log"
	"net/http"
	"wefacer/core"
	"wefacer/wechat/wechatservices"
)

func main() {
	http.HandleFunc("/", wechatservices.ReceiveRequest)                       //设置访问的路由
	err := http.ListenAndServe(":"+core.WefacerConfig.ConfigMap["port"], nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
