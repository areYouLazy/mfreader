package mifare

func ParseAndPrettyPrintMIFAREBinaryDump(data []byte) {
	mifare, err := parseMIFAREBinaryFile(data)
	if err != nil {
		panic(err.Error())
	}

	prettyPrintMIFAREInfos(mifare)
}

func ParseAndPrettyPrintMIFAREManufacturerBinaryDump(data []byte) {
	mifare, err := parseMIFAREBinaryFile(data)
	if err != nil {
		panic(err.Error())
	}

	prettyPrintMIFAREManufacturerInfos(mifare)
}

func ParseAnsJSONPrintMIFAREManufacturerBinaryDump(data []byte) {
	mifare, err := parseMIFAREBinaryFile(data)
	if err != nil {
		panic(err.Error())
	}

	jsonPrintMIFAREManufacturerInfos(mifare)
}

func ParseAnsJSONPrintMIFAREBinaryDump(data []byte) {
	mifare, err := parseMIFAREBinaryFile(data)
	if err != nil {
		panic(err.Error())
	}

	jsonPrintMIFAREInfos(mifare)
}
