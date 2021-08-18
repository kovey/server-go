package server

import (
	"bytes"
	"encoding/binary"
)

type LenType int32

const (
	U_INT_8  LenType = 1
	INT_8    LenType = 1
	U_INT_16 LenType = 2
	INT_16   LenType = 2
	U_INT_32 LenType = 4
	INT_32   LenType = 4
	U_INT_64 LenType = 8
	INT_64   LenType = 8
)

type Config struct {
	OpenCheck      bool
	BodyLenOffset  int32
	PackageLenType LenType
	PackageMax     int32
	HeaderLen      int32
}

func NewConfig(openCheck bool, bodyLenOffset int32, headerLen int32, packageLenType LenType) *Config {
	return &Config{OpenCheck: openCheck, BodyLenOffset: bodyLenOffset, HeaderLen: headerLen, PackageLenType: packageLenType, PackageMax: 2097152}
}

func BytesToInt8(buf []byte) int8 {
	bBuf := bytes.NewBuffer(buf)
	var data int8
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}

func BytesToInt16(buf []byte) int16 {
	bBuf := bytes.NewBuffer(buf)
	var data int16
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}

func BytesToInt32(buf []byte) int32 {
	bBuf := bytes.NewBuffer(buf)
	var data int32
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}

func BytesToInt64(buf []byte) int64 {
	bBuf := bytes.NewBuffer(buf)
	var data int64
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}

func BytesToUInt8(buf []byte) uint8 {
	bBuf := bytes.NewBuffer(buf)
	var data uint8
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}

func BytesToUInt16(buf []byte) uint16 {
	bBuf := bytes.NewBuffer(buf)
	var data uint16
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}

func BytesToUInt32(buf []byte) uint32 {
	bBuf := bytes.NewBuffer(buf)
	var data uint32
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}

func BytesToUInt64(buf []byte) uint64 {
	bBuf := bytes.NewBuffer(buf)
	var data uint64
	binary.Read(bBuf, binary.BigEndian, &data)
	return data
}
