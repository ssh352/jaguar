package struc

import (
	"fmt"
)

type Logon struct {
	SenderCompID     []byte
	TargetCompID     []byte
	HeartBtInt       int32
	Password         []byte
	DefaultApplVerID []byte
}

func NewLogon(logon map[string]string) *Logon {
	return &Logon{
		SenderCompID:     []byte(logon["SenderCompID"]),
		TargetCompID:     []byte(logon["TargetCompID"]),
		HeartBtInt:       256,
		Password:         []byte(logon["Password"]),
		DefaultApplVerID: []byte(logon["DefaultApplVerID"]),
	}
}

func Test() {
	fmt.Println("test!")
}

func (logon *Logon) Marshal() []byte {
	msg := make([]byte, 92)
	copy(msg, logon.SenderCompID)
	copy(msg[20:], logon.TargetCompID)
	x := uint32(logon.HeartBtInt)
	data := make([]byte, 4)
	data[0] = (byte)(x >> 24)
	data[1] = (byte)(x >> 16)
	data[2] = (byte)(x >> 8)
	data[3] = (byte)(x)
	copy(msg[40:], data)
	copy(msg[44:], logon.Password)
	copy(msg[60:], logon.DefaultApplVerID)
	return msg
}
