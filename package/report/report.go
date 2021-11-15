package report

import (
	"fmt"
	"math"
	"time"

	"github.com/signintech/gopdf"
)

func setPageNumb(pdf *gopdf.GoPdf, currentPageNumber int) {
	//show soma logo on the left side
	/*xp := leftMargin / 2
	yp := pageHeight - bottomMargin
	getSomaIcon(pdf, xp, yp)*/

	//show page number on the right side
	//pdf.SetX(pageWidth - rightMargin - leftMargin - 27)
	pdf.SetX(rightMargin)
	pdf.SetY(pageHeight - bottomMargin)

	//set timestamp
	timeSignature := time.Now().Format("02/01/2006 15:04")
	pdf.SetTextColor(0, 0, 0)
	setFont(pdf, 8)
	pdf.Text("This report was generated on " + timeSignature)

	pdf.SetX(pageWidth - rightMargin - leftMargin - 27)
	pdf.SetTextColor(0, 122, 204)
	setFontBold(pdf, 8)
	pdf.Text("SOMA /")
	pdf.SetTextColor(0, 0, 0)

	setFont(pdf, 8)
	pdf.SetX(pageWidth - rightMargin - leftMargin)
	//pdf.SetY(pageHeight - bottomMargin)
	pdf.Text(fmt.Sprintf("/ PAGE %d/%d", currentPageNumber, totalPageNumber))
}

func calculatePages(rows float64, rowHeight float64) int {
	firstPageSpace := pageHeight - headingMargin - footerMargin
	pageSpace := pageHeight - topMargin - footerMargin
	totalRequiredSpace := rows * rowHeight
	totalPages := math.Ceil((totalRequiredSpace-firstPageSpace)/pageSpace) + 1
	return int(totalPages)
}
