package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//
// Encode MsgHeader method
//
func (rcv *MsgHeader) Encode() ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})

	if err := binary.Write(buff, binary.BigEndian, rcv); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

//
// Decode MsgHeader method
//
func (rcv *MsgHeader) Decode(data []byte) error {
	buff := bytes.NewBuffer(data)

	if err := binary.Read(buff, binary.BigEndian, rcv); err != nil {
		return err
	}

	return nil
}

//
// String MsgHeader method
//
func (rcv *MsgHeader) String() string {
	return fmt.Sprintf(
		`ProtVer: %d
Length: %d
SeqNum: %d
MID: %d
TOut: %d
Session: %d
Flags: %d`,
		rcv.ProtVer,
		rcv.Length,
		rcv.SeqNum,
		rcv.MID,
		rcv.TOut,
		rcv.Session,
		rcv.Flags,
	)
}

//
// Encode RROComIni method
//
func (rcv *RROComIni) Encode() ([]byte, error) {
	bytes, err := rcv.Header.Encode()
	if err != nil {
		return nil, err
	}

	b := make([]byte, 6)
	binary.BigEndian.PutUint32(b[:], rcv.IDdev)
	binary.BigEndian.PutUint16(b[4:], rcv.KSize)

	buff := bytes.NewBuffer([]byte{})

	if err := binary.Read(buff, binary.BigEndian, rcv.G); err != nil {
		return nil, err
	}
}

func generateIndexMap(base, size uint16) *map[string]uint16 {
	m := map[string]uint16{
		"G": base,
		"P": base + (size+7)/8,
		"A": base + 2*(size+7)/8,
	}
}

//
// String RROComIni method
//
func (rcv *RROComIni) String() string {
	return fmt.Sprintf(
		`%s
IDdev: %d
KSize: %d
G: %v
P: %v
A: %v
MSize: %d
MAC: %v
MACKey: %v
ZPad: %v
CRC: %d`,
		rcv.Header, // rcv.Header.String()
		rcv.IDdev,
		rcv.KSize,
		rcv.G,
		rcv.P,
		rcv.A,
		rcv.MSize,
		rcv.MAC,
		rcv.MACKey,
		rcv.ZPad,
		rcv.CRC,
	)
}
