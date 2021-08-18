package server

type Header struct {
	PackageLen int32
	BodyLen    int32
	HBuf       []byte
}

func NewHeader() *Header {
	return &Header{PackageLen: 0, BodyLen: 0}
}

func (h *Header) Parse(buf []byte, c *Config) {
	length := int32(c.PackageLenType)
	h.HBuf = buf[c.BodyLenOffset : c.BodyLenOffset+length]
	if c.PackageLenType == U_INT_8 {
		h.BodyLen = int32(BytesToUInt8(h.HBuf))
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	if c.PackageLenType == INT_8 {
		h.BodyLen = int32(BytesToInt8(h.HBuf))
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	if c.PackageLenType == U_INT_16 {
		h.BodyLen = int32(BytesToUInt16(h.HBuf))
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	if c.PackageLenType == INT_16 {
		h.BodyLen = int32(BytesToInt16(h.HBuf))
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	if c.PackageLenType == U_INT_32 {
		h.BodyLen = int32(BytesToUInt32(h.HBuf))
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	if c.PackageLenType == INT_32 {
		h.BodyLen = BytesToInt32(h.HBuf)
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	if c.PackageLenType == U_INT_64 {
		h.BodyLen = int32(BytesToUInt64(h.HBuf))
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	if c.PackageLenType == INT_64 {
		h.BodyLen = int32(BytesToInt64(h.HBuf))
		h.PackageLen = c.HeaderLen + h.BodyLen
		return
	}

	h.BodyLen = BytesToInt32(h.HBuf)
	h.PackageLen = c.HeaderLen + h.BodyLen
}
