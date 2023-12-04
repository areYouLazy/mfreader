package mifare

// def accbits_to_permission_sector(accbits):
//     permissions = {
//         '000': "- A | A   - | A A [read B]",
//         '010': "- - | A   - | A - [read B]",
//         '100': "- B | A/B - | - B",
//         '110': "- - | A/B - | - -",
//         '001': "- A | A   A | A A [transport]",
//         '011': "- B | A/B B | - B",
//         '101': "- - | A/B B | - -",
//         '111': "- - | A/B - | - -",
//     }
//     if isinstance(accbits, BitArray):
//         return permissions.get(accbits.bin, "unknown")
//     else:
//         return ""

// def accbits_to_permission_data(accbits):
//     permissions = {
//         '000': "A/B | A/B   | A/B | A/B [transport]",
//         '010': "A/B |  -    |  -  |  -  [r/w]",
//         '100': "A/B |   B   |  -  |  -  [r/w]",
//         '110': "A/B |   B   |   B | A/B [value]",
//         '001': "A/B |  -    |  -  | A/B [value]",
//         '011': "  B |   B   |  -  |  -  [r/w]",
//         '101': "  B |  -    |  -  |  -  [r/w]",
//         '111': " -  |  -    |  -  |  -  [r/w]",
//     }
//     if isinstance(accbits, BitArray):
//         return permissions.get(accbits.bin, "unknown")
//     else:
//         return ""

var (
	AccessBitsToSectorPermissions = map[string]string{
		"000": "- A | A   - | A A [read B]",
		"010": "- - | A   - | A - [read B]",
		"100": "- B | A/B - | - B",
		"110": "- - | A/B - | - -",
		"001": "- A | A   A | A A [transport]",
		"011": "- B | A/B B | - B",
		"101": "- - | A/B B | - -",
		"111": "- - | A/B - | - -",
	}

	AccessBitsToDataPermissions = map[string]string{
		"000": "A/B | A/B   | A/B | A/B [transport]",
		"010": "A/B |  -    |  -  |  -  [r/w]",
		"100": "A/B |   B   |  -  |  -  [r/w]",
		"110": "A/B |   B   |   B | A/B [value]",
		"001": "A/B |  -    |  -  | A/B [value]",
		"011": "  B |   B   |  -  |  -  [r/w]",
		"101": "  B |  -    |  -  |  -  [r/w]",
		"111": " -  |  -    |  -  |  -  [r/w]",
	}
)
