package struc

type Header struct {
	MsgType    uint32
	BodyLength uint32
}

func NewHeader(msgType, bodyLength uint32) *Header {
	return &Header{
		MsgType:    msgType,
		BodyLength: bodyLength,
	}
}

func (header *Header) Marshal() []byte {
	msg := make([]byte, 8)
	x := uint(header.MsgType)
	msg[0] = (byte)((x & 0xFF000000) >> 24)
	msg[1] = (byte)((x & 0x00FF0000) >> 16)
	msg[2] = (byte)((x & 0x0000FF00) >> 8)
	msg[3] = (byte)(x & 0x000000FF)
	x = uint(header.BodyLength)
	msg[4] = (byte)((x & 0xFF000000) >> 24)
	msg[5] = (byte)((x & 0x00FF0000) >> 16)
	msg[6] = (byte)((x & 0x0000FF00) >> 8)
	msg[7] = (byte)(x & 0x000000FF)
	return msg
}
