package main

import (
	pb "example"
	"fmt"
)

const (
	MsgID_Test1 = iota
	MsgID_Test2
)

func MessageHandler_Test1(msgid uint16, msg interface{}) {
	p := msg.(*pb.MsgTest1)
	fmt.Println("message handler msgid:", msgid, " body:", p)
}

func MessageHandler_Test2(msgid uint16, msg interface{}) {
	p := msg.(*pb.MsgTest2)
	fmt.Println("message handler msgid:", msgid, " body:", p)
}

func RegistMsg() {
	//register
	RegisterMessage(MsgID_Test1, &pb.MsgTest1{}, MessageHandler_Test1)
	RegisterMessage(MsgID_Test2, &pb.MsgTest2{}, MessageHandler_Test2)
}
