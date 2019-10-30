package util

import (
	"base/log"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"usercmd"
)

const CmdHeaderSize int = 2

type Message interface {
	Marshal() (data []byte, err error)
	MarshalTo(data []byte) (n int, err error)
	Size() (n int)
	Unmarshal(data []byte) error
}

// 获取指令号
func GetCmdType(buf []byte) uint16 {
	if len(buf) < CmdHeaderSize {
		return 0
	}
	return uint16(buf[1]) | uint16(buf[0])<<8
}

// 生成二进制数据,返回数据和标识
func EncodeCmd(cmd usercmd.UserCmd, msg proto.Message) ([]byte, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error.Println("EncodeCmd fail cmd:", cmd)
		return nil, err
	}
	p := make([]byte, len(data)+CmdHeaderSize)
	binary.BigEndian.PutUint16(p[0:], uint16(cmd))
	copy(p[2:], data)
	return p, nil
}

func DecodeCmd(buf []byte, pb proto.Message) proto.Message {
	if len(buf) < CmdHeaderSize {
		return nil
	}
	var tmpBuff []byte
	tmpBuff = buf[CmdHeaderSize:]
	err := proto.Unmarshal(tmpBuff, pb)
	if err != nil {
		return nil
	}
	return pb
}