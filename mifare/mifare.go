package mifare

type Product int

const (
	MIFAREDESFire Product = iota
	MIFAREPlus
	MIFAREUltralight
	NTAG
	RFU
	NTAGI2C
	MIFAREDESFireLight
)

type MIFARE struct {
	Sectors []*Sector
	Size    int
	UID     string
	BCC     string
	SAK     string
	ATQA    string
	// product Product
}
