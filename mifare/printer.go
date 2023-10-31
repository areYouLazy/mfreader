package mifare

import (
	"encoding/json"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// pretty print

func prettyPrintMIFAREManufacturerInfos(mifare *MIFARE) {
	tw := table.NewWriter()
	// tw.SetTitle("Manufacturer Data")
	tw.Style().Title.Align = text.AlignCenter

	headerRow := table.Row{"UID", "BCC", "SAK", "ATQA"}

	tw.AppendHeader(headerRow)
	dataRow := table.Row{mifare.UID, mifare.BCC, mifare.SAK, mifare.ATQA}

	tw.AppendRow(dataRow)

	fmt.Println(tw.Render())
}

func prettyPrintMIFAREInfos(mifare *MIFARE) {
	// generate a new table writer
	tw := table.NewWriter()
	// tw.SetTitle("MIFARE Card Dump")
	tw.Style().Title.Align = text.AlignCenter

	// setup table headers
	mainHeaderRow := table.Row{"Sector", "Block", "Data", "Access Bits"}
	emptyHeaderRow := table.Row{"", "", "", ""}

	dataLable := fmt.Sprintf("%sKey A%s %sAccess Bits%s %sKey B%s", Red, Reset, Green, Reset, Blue, Reset)
	subtitleHeaderRow := table.Row{"", "", dataLable, ""}

	// append headers to table writer
	tw.AppendHeader(mainHeaderRow)
	tw.AppendHeader(emptyHeaderRow)
	tw.AppendHeader(subtitleHeaderRow)

	// setup some style
	tw.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:        "Sector",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Name:        "Block",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Name:        "Data",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Name:        "Access Bits",
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
	})

	// iterate sectors and format rows
	for sectorIDX, sector := range mifare.Sectors {
		// iterate blocks in every sector
		for _, block := range sector.Blocks {
			// get data as hex string
			dataString := block.DataAsHexString()

			// if is trailer sector, put some fency colors
			if block.IsTrailerSector {
				keyA := fmt.Sprintf("%s%s%s",
					Red,
					dataString[0:12],
					Reset,
				)
				accBits := fmt.Sprintf("%s%s%s",
					Green,
					dataString[12:20],
					Reset,
				)
				keyB := fmt.Sprintf("%s%s%s",
					Blue,
					dataString[20:],
					Reset,
				)
				dataString = fmt.Sprintf("%s%s%s", keyA, accBits, keyB)
			}

			// if rights has not been extracted, paint with yellow otherwise green
			var rights string
			if block.RightsAsString() == "ERR" {
				rights = fmt.Sprintf("%s%s%s", Yellow, block.RightsAsString(), Reset)
			} else {
				rights = fmt.Sprintf("%s%s%s", Green, block.RightsAsString(), Reset)
			}

			// generate a new table row
			tw.AppendRow(table.Row{
				sectorIDX,
				block.Number,
				dataString,
				rights,
			})
		}

		// separate sectors
		tw.AppendSeparator()

	}

	// renter table
	fmt.Println(tw.Render())
}

func jsonPrintMIFAREManufacturerInfos(mifare *MIFARE) {

	tmp := struct {
		UID  string
		BCC  string
		SAK  string
		ATQA string
	}{
		UID:  mifare.UID,
		BCC:  mifare.BCC,
		SAK:  mifare.SAK,
		ATQA: mifare.ATQA,
	}

	buf, _ := json.MarshalIndent(tmp, "", "  ")

	fmt.Println(string(buf))
}

func jsonPrintMIFAREInfos(mifare *MIFARE) {
	type stringBlock struct {
		Number          int
		Position        int
		Rights          string
		Data            string
		IsTrailerSector bool
	}

	type stringSector struct {
		Number      int
		Blocks      []*stringBlock
		AccessBytes string
		Data        string
	}

	type stringMIFARE struct {
		Sectors []*stringSector
		Size    int
		UID     string
		BCC     string
		SAK     string
		ATQA    string
	}

	mf := new(stringMIFARE)

	mf.Size = mifare.Size
	mf.UID = mifare.UID
	mf.BCC = mifare.BCC
	mf.SAK = mifare.SAK
	mf.ATQA = mifare.ATQA

	for _, sector := range mifare.Sectors {
		s := new(stringSector)
		s.Number = sector.Number
		s.AccessBytes = sector.AccessBytesAsHexString()
		s.Data = sector.DataAsHexString()

		for _, block := range sector.Blocks {
			b := new(stringBlock)
			b.Number = block.Number
			b.Position = block.Position
			b.Rights = block.RightsAsString()
			b.Data = block.DataAsHexString()
			b.IsTrailerSector = block.IsTrailerSector

			s.Blocks = append(s.Blocks, b)
		}

		mf.Sectors = append(mf.Sectors, s)
	}

	buf, _ := json.MarshalIndent(mf, "", "  ")
	fmt.Println(string(buf))
}
