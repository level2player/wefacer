package recognitionservices

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"wefacer/core"
	"wefacer/models"
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
	var dentifyFace models.IdentifyFace = models.BaiduDentifyFace{}
	faceAutochan := make(chan models.DefaultFaceAuto)
	faceAutoerrchan := make(chan bool)
	timeout := make(chan bool, 1)
	Timing(timeout)
	defer log.Println("execute time:", time.Since(t_now), "\n")
	go dentifyFace.DentifyFace(request, msghead, faceAutochan, faceAutoerrchan)
	select {
	case value := <-faceAutochan:
		SendResponse(MakeResponse(msghead, value))
		break
	case <-faceAutoerrchan:
		SendResponse(MakeErrorResponse(msghead))
		break
	case <-timeout:
		SendResponse([]byte(MakeErrorResponse(msghead)))
		break
	}
}

func MakeResponse(requestHead models.RequestHead, faceAuto models.DefaultFaceAuto) []byte {
	var response models.IResponse = models.TextResponse{}
	rescontent, err := response.EncodeResponse(requestHead, MakeResponseContent(faceAuto))
	if err != nil {
		log.Printf(err.Error())
	}
	return rescontent
}

func MakeResponseContent(faceAuto models.DefaultFaceAuto) string {
	var contentstr string
	for index, result := range faceAuto.Face {
		index++
		contentstr += "第" + strconv.Itoa(index) + "人\n"
		contentstr += "年龄:" + strconv.Itoa(int(result.Age)) + "\n"
		contentstr += "人种:" + result.Race + "\n"
		contentstr += "性别:" + core.ConvertGender(result.Gender) + "\n"
		contentstr += "表情:" + core.Convertexpression(result.Expression) + "\n"
		contentstr += "眼镜:" + core.Convertglasses(result.Glasses) + "\n"
	}
	return contentstr
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
