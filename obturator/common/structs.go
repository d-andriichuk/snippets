package common

//
// MsgHeader type
//
type MsgHeader struct {
	ProtVer uint32
	Length  uint32
	SeqNum  byte
	MID     uint16
	TOut    uint16
	Session uint16
	Flags   uint16
}

//
// RROComIni type
//
type RROComIni struct {
	Header *MsgHeader
	IDdev  uint32
	KSize  uint16
	G      []byte
	P      []byte
	A      []byte
	MSize  uint16
	MAC    []byte
	MACKey []byte
	ZPad   []byte
	CRC    uint32
}
