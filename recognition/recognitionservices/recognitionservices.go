package recognitionservices

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"wefacer/models"
	"reflect"
	"wefacer/core"
)

func HandleMsg(req *http.Request, SendResponse func(resbuffer []byte)) {
	t_now := time.Now()
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Msg Handle Error" + err.Error())
	}
	msghead, err := models.GetMsgHead(content)
	if err != nil {
		log.Println("Msg Handle Error" + err.Error())
	}
	request := models.HandleRequest(content, msghead)
	v := reflect.New(reflect.ValueOf(models.FaceAutoStruct[core.WefacerConfig.ConfigMap["faceauto_type"]]).Type()).Elem()
	var dentifyFace models.IdentifyFace = v.Interface().(models.IdentifyFace)
	faceAutochan := make(chan string)
	faceAutoerrchan := make(chan bool)
	timeout := make(chan bool, 1)
	Timing(timeout)
	defer log.Println("execute time:", time.Since(t_now), "\n")
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

func MakeResponse(requestHead models.RequestHead, responseContent string) []byte {
	var response models.IResponse = models.TextResponse{}
	rescontent, err := response.EncodeResponse(requestHead, responseContent)
	if err != nil {
		log.Printf(err.Error())
	}
	return rescontent
}

func MakeErrorResponse(requestHead models.RequestHead) []byte {
	var response models.IResponse = models.ErrorResponse{}
	rescontent, err := response.EncodeResponse(requestHead, "Message Handle Error Call 13575468007")
	if err != nil {
		log.Printf(err.Error())
	}
	return rescontent
}

func Timing(timeout chan bool) {
	go func() {
		time.Sleep(4800000000)
		timeout <- true
	}()
}
