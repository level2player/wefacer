package models

import (
	"encoding/xml"

	"reflect"
)

type RequestHead struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MsgId        int
}

func GetMsgHead(content []byte) (reqhead RequestHead, err error) {
	err = xml.Unmarshal(content, &reqhead)
	return reqhead, err
}

type IRequest interface {
	UnmarshalRequest(content []byte, head RequestHead) IRequest
	MakeRequest(rchan chan string)
}

type TextRequest struct {
	RequestHead RequestHead
	Content     string
}

func (textRequest TextRequest) UnmarshalRequest(content []byte, head RequestHead) IRequest {
	xml.Unmarshal(content, &textRequest)
	textRequest.RequestHead = head
	return textRequest
}
func (textRequest TextRequest) MakeRequest(rchan chan string) {

	rchan <- "Now Recive Msgtext: " + textRequest.Content
}

type ImageRequest struct {
	RequestHead RequestHead
	PicUrl      string
	MediaId     string
}

func (imageRequest ImageRequest) UnmarshalRequest(content []byte, head RequestHead) IRequest {
	xml.Unmarshal(content, &imageRequest)
	imageRequest.RequestHead = head
	return imageRequest
}
func (imageRequest ImageRequest) MakeRequest(rchan chan string) {
	rchan <- "Now down load wechat image url: " + imageRequest.PicUrl
}

func HandleRequest(content []byte, head RequestHead) IRequest {
	v := reflect.New(reflect.ValueOf(regStruct[head.MsgType]).Type()).Elem()
	request := v.Interface().(IRequest).UnmarshalRequest(content, head)
	return request
	//request.MakeRequest(rchan)
}

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

func init() {
	regStruct = make(map[string]interface{})
	regStruct["text"] = TextRequest{}
	regStruct["image"] = ImageRequest{}
}
