package recognitionservices
import(
	"log"
	"net/http"
	"io/ioutil"
	"wefacer/models"
	"time"
	"strconv"
	"wefacer/core"
)



func HandleMsg(req *http.Request,action func(resbuffer []byte)){
	if req!=nil{
		t1 := time.Now()
		content,err:= ioutil.ReadAll(req.Body)
		if err!=nil{
			log.Println(err.Error())
		}
		msghead,err:=models.GetMsgHead(content)
		if err!=nil{
			log.Println(err.Error())
		}
		if msghead.MsgType!="image"{
			action(MakeResponse(msghead,models.FaceAuto{},true))
			return
		}
		request:=models.HandleRequest(content,msghead)
		var dentifyFace models.IdentifyFace
		dentifyFace=models.BaiduDentifyFace{}
		faceAutochan := make(chan models.FaceAuto)
		faceAutoerrchan := make(chan bool)
		timeout := make (chan bool, 1)
	
		go dentifyFace.DentifyFace(request,msghead,faceAutochan,faceAutoerrchan)
		Timing(timeout)
		select{
		case value:=<-faceAutochan:
				action(MakeResponse(msghead,value,false))
				log.Println("执行时间:", time.Since(t1),"\n")
				return
		case <-faceAutoerrchan:
				action(MakeResponse(msghead,models.FaceAuto{},true))
				log.Println("执行时间:", time.Since(t1),"\n")
				return
		case <- timeout:
				action([]byte(MakeResponse(msghead,models.FaceAuto{},true)))
				return
			}
	}
}


 func MakeResponse(requestHead models.RequestHead,faceAuto models.FaceAuto,isError bool)[]byte{
	var response models.IResponse
	if isError{
	  response=models.ErrorResponse{}
	}else{
	  response=models.TextResponse{}
	}
	rescontent,err:= response.EncodeResponse(requestHead,faceAuto,MakeResponseContent(faceAuto))
	if err!=nil{
		log.Printf(err.Error())
	}
	return rescontent
 }


func Timing(timeout chan bool){
	go func(){
				time.Sleep(4800000000) 
				timeout <- true
			}()
}

func MakeResponseContent(faceAuto models.FaceAuto)string{
	var  contentstr string
	for index,result:= range faceAuto.Result{
		index++
		contentstr+="第"+strconv.Itoa(index)+"人\n"
		contentstr+="年龄:"+strconv.Itoa(int(result.Age))+"\n"
		contentstr+="人种:"+result.Race+"\n"
		contentstr+="性别:"+core.ConvertGender(result.Gender)+"\n"
		contentstr+="表情:"+core.Convertexpression(result.Expression)+"\n"
		contentstr+="眼镜:"+core.Convertglasses(result.Glasses)+"\n"
	}
	return contentstr
}



