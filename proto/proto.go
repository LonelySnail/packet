package proto

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"bytes"
	"encoding/binary"
)

/*
	magicNumber|version|serializationType|msgId|msgLength | msgData
		1         1          	2			4	   4

	magicNumber:   协议Header的起始标识
	version:       协议版本号
	serializationType: 序列化类型(json,bson,protobuf,msgpack)
	msgId:    消息id

*/

type Header struct{
	MagicNumber     byte
	Version 	    byte
	SeriaType 		uint16
	MsgId 			uint32
	MsgLength 		uint32
}

type Message struct{
	*Header
	payLoad []byte
}

func newMessage() *Message {
	msg := new(Message)
	return msg
}

func (this *Header)setMagicNum()  {
	this.MagicNumber= 0x12
}

func (this *Header)setVersion(version byte)  {
	this.Version = version
}

func (this *Header)setSeriaType(Type uint16)  {
	this.SeriaType = Type
}

func (this *Header)setMsgId(msgId uint32)  {
	this.MsgId = msgId
}

func (this *Header)setMsgLength(length uint32)  {
	this.MsgLength = length
}

func (this *Header)GetMagicNum(magic byte) bool  {
	return magic == 0x12
}


func (this *Message)setHeader(header *Header)  {
	this.Header = header
}

func (this *Message)setPayLoad(payLoad []byte)  {
	this.payLoad = payLoad
}

func Packet()  []byte{
	header := new(Header)
	header.setMagicNum()
	header.setVersion(byte(1))
	header.setSeriaType(0)
	header.setMsgId(1)
	a := make(map[string]interface{})
	a["1"] =1
	a["2"] =2
	a["3"] = true
	a["4"] = "5"
	payLoad,_ := json.Marshal(a)
	header.setMsgLength(uint32(len(payLoad)))
	buf := new(bytes.Buffer)
	binary.Write(buf,binary.BigEndian,header)
	buf.Write(payLoad)

	return  buf.Bytes()
}

func UnPacket(r io.Reader)  error {
	b := make([]byte,12)
	n,err :=io.ReadFull(r,b)
	if err != nil || n <12 {
		return err
	}

	header := new(Header)
	if !header.GetMagicNum(b[0]) {
		return  errors.New("message is wrong")
	}
	buf := bytes.NewBuffer(b)
	binary.Read(buf,binary.BigEndian,header)
	//header.MagicNumber = b[0]
	//header.Version = b[1]
	//header.SeriaType = binary.BigEndian.Uint16(b[2:4])
	//header.MsgId = binary.BigEndian.Uint32(b[4:8])
	//header.MsgLength = binary.BigEndian.Uint32(b[8:])
	b = make([]byte,header.MsgLength)
	n,err = io.ReadFull(r,b)
	if err != nil || n < int(header.MsgLength) {
		return err
	}
	msg := newMessage()
	msg.setHeader(header)
	msg.setPayLoad(b)

	fmt.Println(string(msg.payLoad),"3333333333333")
	return  nil
}



func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

