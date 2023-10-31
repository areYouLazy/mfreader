package mifare

import (
	"encoding/hex"

	"github.com/areYouLazy/mfreader/utils"
)

type Sector struct {
	Number      int
	Blocks      []*Block
	AccessBytes []byte
	Data        []byte
}

func (s Sector) DataAsHexString() string {
	return hex.EncodeToString(s.Data)
}

func (s Sector) AccessBytesAsBool() []bool {
	return utils.ByteSliceToBoolSlice(s.AccessBytes)
}

func (s Sector) AccessBytesAsBits() []bool {
	return utils.ByteSliceToBoolSlice(s.AccessBytes)
}

func (s Sector) AccessBytesAsHexString() string {
	return hex.EncodeToString(s.AccessBytes)
}
