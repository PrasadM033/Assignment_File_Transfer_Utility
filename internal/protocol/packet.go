package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
)


const (
	MSG_METADATA = 1
	MSG_CHUNK    = 2
	MSG_END      = 3
)


type Packet struct {
	Type   uint8
	Length uint32
	Data   []byte
}
func TryDecode(buffer []byte) (Packet, []byte, error) {
	// Minimum header size = 1 byte type + 4 bytes length
	if len(buffer) < 5 {
		return Packet{}, buffer, errors.New("not enough data")
	}

	p := Packet{}
	p.Type = buffer[0]
	p.Length = binary.BigEndian.Uint32(buffer[1:5])

	// Check if full packet is available
	if len(buffer) < int(5+p.Length) {
		return Packet{}, buffer, errors.New("incomplete packet")
	}

	p.Data = buffer[5 : 5+p.Length]

	remaining := buffer[5+p.Length:]

	return p, remaining, nil
}

func Encode(p Packet) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, p.Type); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, uint32(len(p.Data))); err != nil {
		return nil, err
	}
	buf.Write(p.Data)
	return buf.Bytes(), nil
}

func Decode(reader *bytes.Reader) (Packet, error) {
	var p Packet
	binary.Read(reader, binary.BigEndian, &p.Type)
	binary.Read(reader, binary.BigEndian, &p.Length)

	p.Data = make([]byte, p.Length)
	reader.Read(p.Data)

	return p, nil
}