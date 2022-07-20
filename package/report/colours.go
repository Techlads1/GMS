package report

import "github.com/signintech/gopdf"

func setLightGrayBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(247, 248, 250)
	//pdf.SetFillColor(234, 236, 240)
}
func setGreyBackgroundColor(pdf *gopdf.GoPdf) {
	//https://www.color-hex.com/color-palette/15606
	//#a7a7a7 	(167,167,167)
	//#b4b4b4 	(180,180,180)
	//#c0c0c0 	(192,192,192)
	//#cdcdcd 	(205,205,205)
	//#dadada 	(218,218,218)
	//pdf.SetFillColor(169, 169, 169)
	//pdf.SetFillColor(205, 205, 205)
	pdf.SetFillColor(218, 218, 218)
}

func setDarkGrayBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(223, 223, 223)
}

func setDefaultBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(255, 255, 255)
}

func setCriticalBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(192, 0, 0)
}

func setHighBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(255, 0, 0)
}

func setMediumBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(255, 192, 0)
}
func setLowBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(255, 255, 0)
}
func setNegligibleBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(0, 112, 192)
}

func setOrangeBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(255, 165, 0)
}

func setDarkGreenBackgroundColor(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(76, 175, 81)
}
