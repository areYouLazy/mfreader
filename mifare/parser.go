package mifare

import (
	"fmt"

	"github.com/areYouLazy/mfreader/utils"
)

// ANSI escape codes for color formatting
const (
	Reset  = "\x1b[0m"
	Red    = "\x1b[31m"
	Green  = "\x1b[32m"
	Yellow = "\x1b[33m"
	Blue   = "\x1b[34m"
)

var (
	allowedSizes = []int{
		320,
		1024,
		4096,
	}
)

func extractRightsBitsForGivenBlockNumber(accessBits []bool, blockPosition int) []bool {
	bits := make([]bool, 3)
	inverted := make([]bool, 3)

	// Block access bits based on block number
	switch blockPosition {
	case 0:
		// Block 0 access bits
		bits[0] = accessBits[11]
		bits[1] = accessBits[23]
		bits[2] = accessBits[19]
		inverted[0] = accessBits[7]
		inverted[1] = accessBits[3]
		inverted[2] = accessBits[15]
	case 1:
		// Block 1 access bits
		bits[0] = accessBits[10]
		bits[1] = accessBits[22]
		bits[2] = accessBits[18]
		inverted[0] = accessBits[6]
		inverted[1] = accessBits[2]
		inverted[2] = accessBits[14]
	case 2:
		// Block 2 access bits
		bits[0] = accessBits[9]
		bits[1] = accessBits[21]
		bits[2] = accessBits[17]
		inverted[0] = accessBits[5]
		inverted[1] = accessBits[1]
		inverted[2] = accessBits[13]
	default:
		// Sector trailer / Block 3 access bits
		bits[0] = accessBits[8]
		bits[1] = accessBits[20]
		bits[2] = accessBits[16]
		inverted[0] = accessBits[4]
		inverted[1] = accessBits[0]
		inverted[2] = accessBits[12]
	}

	// Invert the inverted array
	for i := range inverted {
		inverted[i] = !inverted[i]
	}

	// Check if bits and inverted match
	if utils.EqualBits(bits, inverted) {
		return bits
	}

	// return nil
	return bits
}

func isValidSize(dataSize int) bool {
	for _, size := range allowedSizes {
		if dataSize == size {
			return true
		}
	}

	return false
}

func splitByteSliceIntoChunks(data []byte, chunkSize int) ([][]byte, error) {
	// Check if the chunkSize is valid
	if chunkSize <= 0 {
		return nil, fmt.Errorf("chunk size must be greater than 0")
	}

	// Initialize an empty slice to store the chunks
	var chunks [][]byte

	// Split the byte slice into equal parts of chunkSize
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

// parseMIFAREBinaryFile generates a MIFARE object from a binary file dump
func parseMIFAREBinaryFile(data []byte) (*MIFARE, error) {
	absoluteBlockNumber := 0

	mifare := &MIFARE{
		Sectors: make([]*Sector, 0),
		Size:    len(data),
		UID:     "",
		BCC:     "",
		SAK:     "",
		ATQA:    "",
	}

	// ensure data is of valid length
	dataSize := len(data)
	if !isValidSize(dataSize) {
		return nil, fmt.Errorf("wrong file size: %d bytes. Only 320, 1024, or 4096 bytes allowed", dataSize)
	}

	// split data in chunks (sectors)
	rawSectors, _ := splitByteSliceIntoChunks(data, 64)

	// iterate sectors
	for idx, rawSector := range rawSectors {
		// generate Sector object
		s := &Sector{
			Number:      idx,
			Blocks:      make([]*Block, 0),
			AccessBytes: make([]byte, 0),
			Data:        rawSector,
		}

		// split sector in chunks (blocks)
		rawBlocks, _ := splitByteSliceIntoChunks(s.Data, 16)

		// iterate rawblocks and generate Block objects
		for idx, rawBlock := range rawBlocks {
			b := &Block{
				Number:   absoluteBlockNumber,
				Position: idx,
				Rights:   make([]bool, 0),
				Data:     rawBlock,
			}

			// append new block
			s.Blocks = append(s.Blocks, b)

			// before incrementing absoluteBlockNumber
			// check if this is the first block
			// and extract manufacter data
			if b.Number == 0 {
				hexBuffer := b.DataAsHexString()

				mifare.UID = hexBuffer[0:8]
				mifare.BCC = hexBuffer[8:10]
				mifare.SAK = hexBuffer[10:12]
				mifare.ATQA = hexBuffer[12:14]
			}

			// increment block counter
			absoluteBlockNumber++
		}

		// mark last block in the sector as TrailingBlock
		s.Blocks[len(s.Blocks)-1].IsTrailerSector = true

		// extract access bytes from trailer sector block
		for _, block := range s.Blocks {
			if block.IsTrailerSector {
				s.AccessBytes = block.Data[6:9]
			}
		}

		// get access bytes for current sector as bool
		accessBits := s.AccessBytesAsBits()

		// iterate blocks and calculate rights
		for _, block := range s.Blocks {
			block.Rights = extractRightsBitsForGivenBlockNumber(accessBits, block.Position)
		}

		// append new sector
		mifare.Sectors = append(mifare.Sectors, s)
	}

	return mifare, nil
}
