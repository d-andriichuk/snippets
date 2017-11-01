package main

import (
	"fmt"
	"log"
	"unknown/snippets/obturator/common"
)

func main() {
	msgHeader := &common.MsgHeader{
		ProtVer: 1,
		Length:  32,
		SeqNum:  16,
		MID:     0x0001,
		TOut:    5,
		Session: 4,
		Flags:   0x0001,
	}

	fmt.Printf("--------Base msgHeader: \n%s\n", msgHeader.String())

	bytes, err := msgHeader.Encode()
	if err != nil {
		log.Fatalf("Failed to encode msgHeader: %s\n", err.Error())
	}

	fmt.Printf("--------MsgHeader encoded bytes %v\n", bytes)

	msgHeader = &common.MsgHeader{}

	fmt.Printf("--------Empty msgHeader: \n%s\n", msgHeader.String())

	err = msgHeader.Decode(bytes)
	if err != nil {
		log.Fatalf("Failed to decode msgHeader: %s\n", err.Error())
	}

	fmt.Printf("--------Decoded msgHeader: \n%s\n", msgHeader.String())

	rroComIni := &common.RROComIni{
		Header: msgHeader,
		IDdev:  143,
		KSize:  1,
		G:      []byte{2},
		P:      []byte{3},
		A:      []byte{4},
		MSize:  4,
		MAC:    []byte{5},
		MACKey: []byte{6},
		ZPad:   []byte{0},
		CRC:    1234,
	}

	fmt.Printf("--------Base RROComIni: \n%s\n", rroComIni.String())
}
