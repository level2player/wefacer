package recognitionservices

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
	"wefacer/core"
	"wefacer/models"
)

func HandleMsg(req *http.Request, SendResponse func(resbuffer []byte)) {
	t_now := time.Now()
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		core.Print_log("Msg Handle Error" + err.Error())
	}
	msghead, err := models.GetMsgHead(content)
	if err != nil {
		core.Print_log("Msg Handle Error" + err.Error())
	}
	request := models.HandleRequest(content, msghead)
	if request!=nil{
		v := reflect.New(reflect.ValueOf(models.FaceAutoStruct[core.WefacerConfig.ConfigMap["faceauto_type"]]).Type()).Elem()
		var dentifyFace models.IdentifyFace = v.Interface().(models.IdentifyFace)
		faceAutochan := make(chan string)
		faceAutoerrchan := make(chan bool)
		timeout := make(chan bool, 1)
		Timing(timeout)
		defer core.Print_log("execute time:%s", time.Since(t_now))
		go dentifyFace.DentifyFace(request, msghead, faceAutochan, faceAutoerrchan)
		select {
		case value := <-faceAutochan:
			SendResponse(MakeResponse(msghead, value))
		case <-faceAutoerrchan:
			SendResponse(MakeErrorResponse(msghead))
		case <-timeout:
			SendResponse([]byte(MakeErrorResponse(msghead)))
		}
	}
}

func MakeResponse(requestHead models.RequestHead, responseContent string) []byte {
	var response models.IResponse = models.TextResponse{}
	rescontent, err := response.EncodeResponse(requestHead, responseContent)
	if err != nil {
		core.Print_log(err.Error())
	}
	return rescontent
}

func MakeErrorResponse(requestHead models.RequestHead) []byte {
	var response models.IResponse = models.ErrorResponse{}
	rescontent, err := response.EncodeResponse(requestHead, "图片识别错误,请发高清自拍无码大图,帅斌你不要搞事情!!!")
	if err != nil {
		core.Print_log(err.Error())
	}
	return rescontent
}

func Timing(timeout chan bool) {
	go func() {
		time.Sleep(4800000000)
		timeout <- true
	}()
}
