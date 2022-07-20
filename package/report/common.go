package report

import (
	"github.com/signintech/gopdf"
)

const (
	//1 millimeter = 3.7795275591 pixel
	//logo size
	logoSize      = 70.00
	qrSize        = 70.00
	somaSize      = 15.00
	leftMargin    = 30.00
	rightMargin   = 30.00
	topMargin     = 30.00
	bottomMargin  = 20.00
	headingMargin = topMargin + logoSize + 20.00
	footerMargin  = bottomMargin + 25.00
	//topMagin    = 75.590551181 // 20 mm = 0.79 inc
	//bottomMagin = 75.590551181 // 20 mm = 0.79 inc
	//leftMargin  = 96.0         // 25.4 mm = 1 inc
	//rightMargin = 56.692913386 // 15 mm = 0.59 inc
	// //595.28, 841.89 = A4

	padding = 10

	//heightPage          = 842 - topMargin
	brSize = 20

	rowSize = 15

	moduleCodeSpan = 3
	//cellHeight      = 20 //25

	dateFormat    = "02/01/2006 15:04"
	instituteName = "Dar es Salaam Institute of Technology (DIT)"
)

var (
	//for potrait
	pageWidth  = 595.28
	pageHeight = 841.89

	//for landscape
	//pageWidth  = 841.89
	//pageHeight = 595.28

	availablePageWidth  = pageWidth - leftMargin - rightMargin
	availablePageHeight = pageHeight - topMargin - bottomMargin

	cellOptionCentre = gopdf.CellOption{
		Align:  gopdf.Middle | gopdf.Center,
		Border: 0,
		Float:  gopdf.Right,
	}

	cellOptionLeft = gopdf.CellOption{
		Align:  gopdf.Middle | gopdf.Left,
		Border: 0,
		Float:  gopdf.Right,
	}
	cellPadding = 10.0
)

var (
	//page
	currentPageNumber int
	totalPageNumber   int
	spaceLen          float64 // length New Line

	//table
	totalRows  float64
	cellWidth  float64
	cellHeight float64
)

func initPDF(landscape bool) *gopdf.GoPdf {

	if landscape && (pageHeight > pageWidth) {
		//change some grobal variables
		temp := pageHeight
		pageHeight = pageWidth
		pageWidth = temp

		availablePageWidth = pageWidth - leftMargin - rightMargin
		availablePageHeight = pageHeight - topMargin - bottomMargin

	}
	pdf := &gopdf.GoPdf{}
	rect := gopdf.Rect{W: pageWidth, H: pageHeight}

	pdf.Start(
		gopdf.Config{
			PageSize: rect,
		})
	pdf.SetMargins(leftMargin, topMargin, rightMargin, bottomMargin)
	return pdf
}

func checkEndOfPageWithoutBr(pdf *gopdf.GoPdf, deltaY float64) {
	checkEndOfPage(pdf, deltaY, false)
}

func checkEndOfPageWithBr(pdf *gopdf.GoPdf, deltaY float64) {
	checkEndOfPage(pdf, deltaY, true)
}

func checkTableEndOfPage(pdf *gopdf.GoPdf, y float64) bool {
	if y > pageHeight-footerMargin {
		currentPageNumber++
		pdf.AddPage()
		setPageNumb(pdf, currentPageNumber)
		return true
	}
	return false
}

func checkEndOfPage(pdf *gopdf.GoPdf, deltaY float64, needBr bool) bool {
	if (pdf.GetY() + deltaY) > pageHeight {
		currentPageNumber++
		pdf.AddPage()
		setPageNumb(pdf, currentPageNumber)
		return true

	} else {
		if needBr {
			pdf.Br(brSize)
		}
	}
	return false
}

func addHrGreyH(pdf *gopdf.GoPdf, yLeft, h float64) {
	pdf.SetStrokeColor(236, 239, 241)
	pdf.SetLineWidth(h)
	pdf.Line(leftMargin, yLeft, leftMargin+availablePageWidth, yLeft)
}
func addVrGreyH(pdf *gopdf.GoPdf, x, yLeft, rHeight, h float64) {
	pdf.SetStrokeColor(236, 239, 241)
	pdf.SetLineWidth(h)
	pdf.Line(x, yLeft, x, yLeft+rHeight)
}

func addHrGrey(pdf *gopdf.GoPdf, yLeft float64) {
	addHrGreyH(pdf, yLeft, 2)
}

func addHr(pdf *gopdf.GoPdf, yLeft float64) {
	pdf.SetStrokeColor(34, 166, 242)
	pdf.SetLineWidth(0.5)
	pdf.Line(leftMargin, yLeft, leftMargin+availablePageWidth, yLeft)
}
func addHrWithLen(pdf *gopdf.GoPdf, x, y, l, w float64) {
	pdf.SetStrokeColor(34, 166, 242)
	pdf.SetLineWidth(w)
	pdf.Line(x, y, x+l, y)
}
func addHrWithLenGrey(pdf *gopdf.GoPdf, x, y, l, w float64) {
	pdf.SetStrokeColor(236, 239, 241)
	pdf.SetLineWidth(w)
	pdf.Line(x, y, x+l, y)
}

func showLine(pdf *gopdf.GoPdf) {
	pdf.SetX(200)
	pdf.Line(pdf.MarginLeft(), pdf.GetY(), 575.0, pdf.GetY())
}

func space(pdf *gopdf.GoPdf) {
	pdf.Br(spaceLen)
}
