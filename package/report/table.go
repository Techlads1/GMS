package report

import (
	"github.com/signintech/gopdf"
)

func addTable(pdf *gopdf.GoPdf, x, y, rowHeight, padding float64, colWidth []float64, tableHeader []string, data [][]string, fontSize int) {
	//calculate number of pages

	totalRows := len(data)
	totalPageNumber += calculatePages(float64(totalRows), rowHeight)
	//fmt.Printf("Estmated number rows = %d =  pages  = %d\n", totalRows, totalPageNumber)

	setPageNumb(pdf, currentPageNumber) //set page number on the first page
	setFontBold(pdf, fontSize)

	addHrGreyH(pdf, y, 0.5)
	_, y = addFillRow(pdf, x, y, rowHeight, padding, colWidth, tableHeader, true)
	addHrGreyH(pdf, y, 0.5)

	setDefaultBackgroundColor(pdf)
	pdf.SetTextColor(0, 0, 0)
	setFont(pdf, fontSize)

	for _, d := range data {
		_, y = addRow(pdf, x, y, rowHeight, padding, colWidth, d, true)
		addHrGreyH(pdf, y, 0.5)
		if checkTableEndOfPage(pdf, y) {
			setFontBold(pdf, fontSize)
			_, y = addFillRow(pdf, leftMargin, topMargin, rowHeight, padding, colWidth, tableHeader, true)
			setFont(pdf, fontSize)
		}
	}
}

// addRow adds table row
func addRow(pdf *gopdf.GoPdf, x, y, h, p float64, w []float64, data []string, alignCentre bool) (float64, float64) {
	for c, d := range data {

		x, y = addCol(pdf, x, y, w[c], h, p, d, false)
		if c == len(w)-1 {
			continue
		}
		addVrGreyH(pdf, x, y, h, 0.5)
	}
	return x, y + h
}

// addFillRow adds table row
func addFillRow(pdf *gopdf.GoPdf, x, y, h, p float64, w []float64, data []string, alignCentre bool) (float64, float64) {
	setLightGrayBackgroundColor(pdf)
	for c, d := range data {
		x, y = addFillCol(pdf, x, y, w[c], h, p, d, false)
		if c == len(w)-1 {
			continue
		}
		addVrGreyH(pdf, x, y, h, 0.5)
	}
	pdf.SetTextColor(0, 0, 0)
	return x, y + h
}

// addCol adds table column
func addCol(pdf *gopdf.GoPdf, x, y, w, h, p float64, text string, alignCentre bool) (float64, float64) {
	rect := gopdf.Rect{
		W: w - 2*p,
		H: h,
	}
	pdf.SetX(x + p)
	pdf.SetY(y)
	if alignCentre {
		pdf.CellWithOption(&rect, text, cellOptionCentre)

	} else {
		pdf.CellWithOption(&rect, text, cellOptionLeft)
	}
	return x + w, y
}

// addFillCol adds table column
func addFillCol(pdf *gopdf.GoPdf, x, y, w, h, p float64, text string, alignCentre bool) (float64, float64) {
	pdf.SetTextColor(0, 0, 0)
	setLightGrayBackgroundColor(pdf)
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "F")
	rect := gopdf.Rect{
		W: w - 2*p,
		H: h,
	}
	pdf.SetX(x + p)
	pdf.SetY(y)
	if alignCentre {
		pdf.CellWithOption(&rect, text, cellOptionCentre)

	} else {
		pdf.CellWithOption(&rect, text, cellOptionLeft)
	}
	setDefaultBackgroundColor(pdf)
	return x + w, y
}

func addBlock(pdf *gopdf.GoPdf, x, y, w, h float64, text string, alignCentre bool) (float64, float64) {
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "D")

	if alignCentre {
		rect := gopdf.Rect{
			W: w - 2*cellPadding,
			H: h,
		}
		pdf.SetX(x)
		pdf.SetY(y)
		pdf.CellWithOption(&rect, text, cellOptionCentre)

	} else {
		rect := gopdf.Rect{
			W: w - 2*cellPadding,
			H: h,
		}
		pdf.SetX(x + cellPadding) //fix as it does not respect left padding
		pdf.SetY(y)
		pdf.CellWithOption(&rect, text, cellOptionLeft)
	}
	return x + w, y
}

func addSheddedBlock(pdf *gopdf.GoPdf, x, y, w, h float64, text string, alignCentre bool) (float64, float64) {
	setGreyBackgroundColor(pdf)
	//setLightGrayBackgroundColor(pdf)
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")

	if alignCentre {
		rect := gopdf.Rect{
			W: w - 2*cellPadding,
			H: h,
		}
		pdf.SetX(x)
		pdf.SetY(y)
		pdf.CellWithOption(&rect, text, cellOptionCentre)

	} else {
		rect := gopdf.Rect{
			W: w - 2*cellPadding,
			H: h,
		}
		pdf.SetX(x + cellPadding) //fix as it does not respect left padding
		pdf.SetY(y)
		pdf.CellWithOption(&rect, text, cellOptionLeft)
	}
	setDefaultBackgroundColor(pdf)
	return x + w, y
}

func addTextBlock(pdf *gopdf.GoPdf, x, y, w, h float64, text string, alignCentre bool) (float64, float64) {
	rect := gopdf.Rect{
		W: w - 2*cellPadding,
		H: h,
	}
	pdf.SetX(x)
	pdf.SetY(y)
	if alignCentre {
		pdf.CellWithOption(&rect, text, cellOptionCentre)

	} else {
		pdf.CellWithOption(&rect, text, cellOptionLeft)
	}
	return x, y + h
}

func normaliseTableWidth(data []float64, width float64) []float64 {
	totalWidth := 0.0
	for _, w := range data {
		totalWidth += w
	}
	for c, w := range data {
		data[c] = w / totalWidth * width
	}
	return data
}
