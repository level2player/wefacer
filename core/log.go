package core

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Logger *log.Logger
)

func init() {
	creat_dir(WefacerConfig.ConfigMap["log_dir"])
}

func Print_log(fotmat_logstr string, v ...interface{}) {
	pathstr := WefacerConfig.ConfigMap["log_dir"] + "/" + time.Now().Format("20060102") + ".log"
	logfile, err := path_exists_or_creat(pathstr)
	defer logfile.Close()
	if err != nil {
		log.Println(err)
		return
	}
	Logger = log.New(logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Printf("log:%s", fmt.Sprintf(fotmat_logstr, v...))
}

func path_exists_or_creat(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if err != nil {
		return os.Create(path)
	} else {
		return os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	}
}

func creat_dir(path string) {
	if len(path) == 0 {
		path = "./logs"
	}
	if !is_dir_exist(path) {
		mik_dir(path)
	}
}

//判断路径是否存在
func is_dir_exist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}

}

//创建路径
func mik_dir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("add dir error:", err)
		return err
	} else {
		return nil
	}
}
