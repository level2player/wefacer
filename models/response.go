package models

import (
	"encoding/xml"
	"time"
)

type IResponse interface {
	EncodeResponse(reqhead RequestHead, responsecontent string) (data []byte, err error)
}

type ResponseHead struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
}

type TextResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}
type ImageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MediaId      string
}
type VoiceResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MediaId      string
}
type ErrorResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

func (textResponse TextResponse) EncodeResponse(reqhead RequestHead, responsecontent string) (data []byte, err error) {
	textResponse.MsgType = "text"
	textResponse.CreateTime = time.Second
	textResponse.FromUserName = reqhead.ToUserName
	textResponse.ToUserName = reqhead.FromUserName
	if len(responsecontent) == 0 {
		textResponse.Content = "照片处理失败"
	} else {
		textResponse.Content = responsecontent
	}

	data, err = xml.Marshal(textResponse)
	return
}
func (imageResponse ImageResponse) EncodeResponse(reqhead RequestHead, responsecontent string) (data []byte, err error) {

	data, err = xml.Marshal(&imageResponse)
	return
}
func (voiceResponse VoiceResponse) EncodeResponse(reqhead RequestHead, responsecontent string) (data []byte, err error) {

	data, err = xml.Marshal(&voiceResponse)
	return
}
func (errorResponse ErrorResponse) EncodeResponse(reqhead RequestHead, responsecontent string) (data []byte, err error) {
	errorResponse.MsgType = "text"
	errorResponse.CreateTime = time.Second
	errorResponse.FromUserName = reqhead.ToUserName
	errorResponse.ToUserName = reqhead.FromUserName
	errorResponse.Content = responsecontent
	data, err = xml.Marshal(&errorResponse)
	return
}
