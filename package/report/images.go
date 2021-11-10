package report

import (
	"bytes"
	_ "embed"
	"fmt"
	"gateway/package/log"

	"github.com/signintech/gopdf"
	"github.com/yeqown/go-qrcode"
)

//go:embed images/dit-logo.png
var ditLogo []byte

//go:embed images/soma-icon.png
var somaLogo []byte

func getLogo(pdf *gopdf.GoPdf, posX, posY float64) {
	logoImageHolder, err := gopdf.ImageHolderByBytes(ditLogo)
	if err != nil {
		log.Errorf("error crating a holder: %v")
	}

	if err := pdf.ImageByHolder(logoImageHolder, posX, posY-10, &gopdf.Rect{W: logoSize + 20, H: logoSize + 20}); err != nil {
		log.Errorf("could not place logo: %v", err)
	}
}

func getSomaIcon(pdf *gopdf.GoPdf, posX, posY float64) {
	somaIconHolder, err := gopdf.ImageHolderByBytes(somaLogo)
	if err != nil {
		log.Errorf("error creative a holder: %v", err)
	}

	if err := pdf.ImageByHolder(somaIconHolder, posX, posY, &gopdf.Rect{W: somaSize, H: somaSize}); err != nil {
		log.Errorf("could not place logo: %v", err)
	}
}

func getQR(pdf *gopdf.GoPdf, posX, posY float64, qrs, doi string) {
	qrc, err := qrcode.New(qrs)

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	var qrBytes bytes.Buffer
	err = qrc.SaveTo(&qrBytes)
	if err != nil {
		fmt.Printf("error creating qr byte: %v", err)
	}

	qrImageHolder, err := gopdf.ImageHolderByBytes(qrBytes.Bytes())
	if err != nil {
		log.Errorf("error creating qr byte: %v", err)
	}

	if err := pdf.ImageByHolder(qrImageHolder, posX, posY, &gopdf.Rect{W: qrSize, H: qrSize}); err != nil {
		log.Errorf("could not create qr image: %v", err)
	}

	pdf.SetY(posY + qrSize + 10)
	pdf.SetX(posX + 1)
	setFont(pdf, 9)
	pdf.Text("DOI: " + doi)
}

//addChart
func addChart(pdf *gopdf.GoPdf, posX, posY, chartWidth, chartHeight float64, chartData []byte) {
	chartBytes, err := gopdf.ImageHolderByBytes(chartData)
	if err != nil {
		log.Errorf("error crating a holder: %v")
	}

	if err := pdf.ImageByHolder(chartBytes, posX, posY, &gopdf.Rect{W: chartWidth, H: chartHeight}); err != nil {
		log.Errorf("could not place logo: %v", err)
	}
}
