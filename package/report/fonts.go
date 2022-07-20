package report

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

var (
	//go:embed fonts/arial.ttf
	arial []byte

	//go:embed fonts/arial-bold.ttf
	arialBold []byte

	//go:embed fonts/arial-italic.ttf
	arialItalic []byte
)

func addFonts(pdf *gopdf.GoPdf) {
	//add fonts

	if err := pdf.AddTTFFontByReader("arial", bytes.NewBuffer(arial)); err != nil {
		fmt.Println("error font")
		log.Print(err.Error())
		return
	}
	if err := pdf.AddTTFFontByReader("arial-italic", bytes.NewBuffer(arialItalic)); err != nil {
		fmt.Println("error font")
		log.Print(err.Error())
		return
	}

	if err := pdf.AddTTFFontByReader("arial-bold", bytes.NewBuffer(arialBold)); err != nil {
		fmt.Println("error font")
		log.Print(err.Error())
		return
	}
	/*if err := pdf.AddTTFFont("arial-bold-italic", "fonts/arial/arial-bold-italic.ttf"); err != nil {
		fmt.Println("error font")
		log.Print(err.Error())
		return
	}*/
}

func setFont(pdf *gopdf.GoPdf, size int) {
	if err := pdf.SetFont("arial", "", size); err != nil {
		log.Print(err.Error())
		return
	}
}
func setFontItalic(pdf *gopdf.GoPdf, size int) {
	if err := pdf.SetFont("arial-italic", "", size); err != nil {
		log.Print(err.Error())
		return
	}
}

func setFontBold(pdf *gopdf.GoPdf, size int) {
	if err := pdf.SetFont("arial-bold", "", size); err != nil {
		log.Print(err.Error())
		return
	}
}
