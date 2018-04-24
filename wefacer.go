package main

import (
	"log"
	"net/http"
	"wefacer/core"
	"wefacer/wechat/wechatservices"
	"github.com/spf13/cobra"
	"os"
)
var RootCmd = &cobra.Command{
    Use: "-v",
    Run: func(cmd *cobra.Command, args []string) {
        println("wefacer version is 0.0.1")
	},
}
var UpCmd= &cobra.Command{
	Use: "up",
    Run: func(cmd *cobra.Command, args []string) {
        LoadWefacer()
	},
}
func init(){
    RootCmd.AddCommand(UpCmd)
}

func LoadWefacer(){
	core.Print_log("start wefacer ...")
	http.HandleFunc("/", wechatservices.ReceiveRequest)
	err := http.ListenAndServe(":"+core.WefacerConfig.ConfigMap["port"], nil)
	if err != nil {
		core.Print_log(err.Error())
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        log.Println(err)
        os.Exit(1)
    }
}
