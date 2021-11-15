package report

import (
	"strings"

	"github.com/signintech/gopdf"
)

func mainHeader(pdf *gopdf.GoPdf, mainTitle, title, qrs, doi string) {
	//get positions
	xp := pdf.GetX()
	yp := pdf.GetY()

	//get logo
	getLogo(pdf, xp, yp)

	//wiret header
	xp = leftMargin + logoSize
	yp = topMargin
	//yp += logoSize / 3.0
	//pdf.SetX(xp)
	//pdf.SetY(yp)
	setFont(pdf, 12)

	//title = strings.ToUpper(title)
	//fmt.Printf("Title: %v", title)

	//lines, _ := pdf.SplitText(title, 200)
	titleWidth := availablePageWidth - logoSize - qrSize // - 50
	x, y := addMultiLineBlock(pdf, xp, yp, titleWidth, 30, strings.ToUpper(mainTitle), true)
	addMultiLineBlock(pdf, x, y, titleWidth, 15.0, title, true)
	//lines := wrapTextLines(pdf, title, titleWidth)
	//fmt.Printf("lines: %v\n",lines)
	//addMultiLines(pdf, xp, 15.0, lines)
	//pdf.MultiCell(&rect, title)
	//pdf.Text(title)
	//pdf.RectFromUpperLeftWithStyle(50, 100, 400, 600, "FD")
	//pdf.SetFillColor(0, 0, 0)
	//pdf.MultiCell(nil, title)
	xp = pageWidth - rightMargin - qrSize
	yp = topMargin
	//fmt.Printf("Margin right  = %f, margin left = %f", pdf.MarginRight(), pdf.MarginLeft())
	getQR(pdf, xp, yp, qrs, doi)

	addHr(pdf, topMargin+logoSize+15)
	//addHrGreyH(pdf, topMargin+logoSize+10, 1.0)
}
