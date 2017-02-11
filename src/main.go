package main

import (
	pb "example"
	"github.com/golang/protobuf/proto"
)

func main() {
	RegistMsg()

	t1 := pb.MsgTest1{
		Id:    11111,
		Name:  "name1111",
		Email: "1111@example.com",
	}
	out, _ := proto.Marshal(&t1)
	HandleRawData(MsgID_Test1, out)

	t2 := pb.MsgTest2{
		Id:    22222,
		Name:  "name222",
		Email: "2222@example.com",
	}
	out, _ = proto.Marshal(&t2)
	HandleRawData(MsgID_Test2, out)
}
