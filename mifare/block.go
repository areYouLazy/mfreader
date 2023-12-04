package mifare

import (
	"encoding/hex"
	"fmt"

	"github.com/areYouLazy/mfreader/utils"
)

// Block represents a 16 bytes block
// A sector is composed by 4 blocks
type Block struct {
	//Number is the absolute block number
	Number int

	//Position is the block position inside a sector (0, 1, 2 or 3)
	Position int

	//Rights extracted from Sector's Access Bits
	Rights []bool

	//Data is a 16 byte slice, this is the actual block data
	Data []byte

	//True if this block is a trailer sector (the last block in the sector)
	IsTrailerSector bool
}

func (b Block) RightsAsString() string {
	s := ""

	if len(b.Rights) != 3 {
		return "ERR"
	}

	for _, b := range b.Rights {
		if b {
			s = fmt.Sprintf("%s%s", s, "1")
		} else {
			s = fmt.Sprintf("%s%s", s, "0")
		}
	}

	return s
}

func (b Block) RightsAsHumanString() string {
	// get permissions for humans
	if b.IsTrailerSector {
		return AccessBitsToSectorPermissions[b.RightsAsString()]
	} else {
		return AccessBitsToDataPermissions[b.RightsAsString()]
	}
}

func (b Block) DataAsHexString() string {
	return hex.EncodeToString(b.Data)
}

func (b Block) DataAsString() string {
	return string(b.Data)
}

func (b Block) DataAsBytes() []byte {
	return b.Data
}

func (b Block) DataAsBool() []bool {
	return utils.ByteSliceToBoolSlice(b.Data)
}
