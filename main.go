package main

import (
	"fmt"
	"os"

	//unicommon "github.com/unidoc/unidoc/common"
	"github.com/unidoc/unidoc/pdf/creator"
	pdf "github.com/unidoc/unidoc/pdf/model"
	"github.com/unidoc/unidoc/pdf/model/fonts"
)

type pdfText struct {
	x     float64
	y     float64
	label string
}

func main() {

	// Test Data
	texts := []pdfText{
		{95, 112.4, "Club Name"},
		{380, 112.4, "Club City"},
		{518, 112.4, "Club State"},

		{105, 128, "League Name"},

		{100, 280, "Player Name"},
		{359, 280, "Player DOB"},
		{476, 280, "X"},
		{522, 280, "X"},

		{100, 295, "Player Address"},
		{394, 295, "Player City"},

		{66, 310, "Player State"},
		{157, 310, "Player Zip"},
		{265, 310, "Player Email"},

		{95, 336, "Parent Name"},
		{326, 336, "Parent Home"},
		{466, 336, "Parent Work"},

		{95, 351, "Parent Email"},
		{319, 351, "Parent Cell"},
		{485.5, 351, "X"},
		{516, 351, "X"},

		{95, 366, "Parent Name"},
		{326, 366, "Parent Home"},
		{466, 366, "Parent Work"},

		{95, 381, "Parent Email"},
		{319, 381, "Parent Cell"},
		{485.5, 381, "X"},
		{516, 381, "X"},

		{65, 420, "Emergency Contact Name"},
		{309, 420, "Contact Phone 1"},
		{454, 420, "Contact Phone 2"},

		{65, 435, "Emergency Contact Name"},
		{309, 435, "Contact Phone 1"},
		{454, 435, "Contact Phone 2"},

		{141, 462, "Allergies"},

		{173, 476, "Medical Conditions"},

		{81, 502, "Physician Name"},
		{313, 502, "Physician Phone 1"},
		{457, 502, "Physician Phone 2"},

		{179, 517, "Insurance Company"},
		{448, 517, "Insurance Phone"},

		{121, 532, "Insured Name"},
		{477, 532, "Policy Number"},
	}

	c, err := loadPDF("US_Club_Form.pdf")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, text := range texts {
		addTextToPDF(text.x, text.y, text.label, c)
	}

	if err := savePDF("output.pdf", c); err != nil {
		fmt.Printf("Error Saving PDF: %s\n", err)
		os.Exit(1)
	}

	// TODO: Verify that the user specified a CSV file on the command line
	// TODO: Open CSV
	// TODO: For each line in the CSV, generate a new PDF

}

func loadPDF(name string) (*creator.Creator, error) {

	// Read the input pdf file.
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return nil, err
	}

	c := creator.New()

	page, err := pdfReader.GetPage(1)
	if err != nil {
		return nil, err
	}

	err = c.AddPage(page)
	if err != nil {
		return nil, err
	}

	return c, nil

}

func addTextToPDF(xPos float64, yPos float64, text string, c *creator.Creator) error {
	p := creator.NewParagraph(text)

	// Change to times bold font (default is helvetica).
	p.SetFont(fonts.NewFontTimesBold())
	p.SetPos(xPos, yPos)

	return c.Draw(p)
}

func savePDF(name string, c *creator.Creator) error {
	return c.WriteToFile(name)
}
