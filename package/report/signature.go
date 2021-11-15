package report

import (
	"github.com/signintech/gopdf"
)

func addModuleSignature(pdf *gopdf.GoPdf) {

	spaceLength := 5.0
	nameLength := 40.0
	signatureLength := 30.0
	dateLenth := 20.0
	vspace := 5.0

	//normalise width
	width := pageWidth - leftMargin - rightMargin
	totalLength := spaceLength*2 + nameLength + signatureLength + dateLenth
	spaceLength = spaceLength / totalLength * width
	nameLength = nameLength / totalLength * width
	signatureLength = signatureLength / totalLength * width
	dateLenth = dateLenth / totalLength * width

	setFont(pdf, 8)

	signHeight := 5.0

	//examiner name
	addHrWithLen(pdf, rightMargin, pageHeight-signHeight*bottomMargin, nameLength, 1)

	addTextBlock(pdf, rightMargin, pageHeight-signHeight*bottomMargin+vspace, nameLength, vspace, "Examiner's Name", true)
	//examiner signature
	addHrWithLen(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin, signatureLength, 1)
	addTextBlock(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin+vspace, signatureLength, vspace, "Signature", true)
	//examiner date
	addHrWithLen(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin, dateLenth, 1)
	addTextBlock(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin+vspace, dateLenth, vspace, "Date", true)

	signHeight = 3

	//
	//examiner name
	addHrWithLen(pdf, rightMargin, pageHeight-signHeight*bottomMargin, nameLength, 1)

	addTextBlock(pdf, rightMargin, pageHeight-signHeight*bottomMargin+vspace, nameLength, vspace, "Head of Department's Name", true)
	//examiner signature
	addHrWithLen(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin, signatureLength, 1)
	addTextBlock(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin+vspace, signatureLength, vspace, "Signature", true)
	//examiner date
	addHrWithLen(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin, dateLenth, 1)
	addTextBlock(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin+vspace, dateLenth, vspace, "Date", true)

}

func addSemesterSignature(pdf *gopdf.GoPdf) {

	spaceLength := 5.0
	nameLength := 40.0
	signatureLength := 30.0
	dateLenth := 20.0
	vspace := 5.0

	//normalise width
	width := pageWidth - leftMargin - rightMargin
	totalLength := spaceLength*2 + nameLength + signatureLength + dateLenth
	spaceLength = spaceLength / totalLength * width
	nameLength = nameLength / totalLength * width
	signatureLength = signatureLength / totalLength * width
	dateLenth = dateLenth / totalLength * width

	setFont(pdf, 8)

	signHeight := 5.0

	//examiner name
	addHrWithLen(pdf, rightMargin, pageHeight-signHeight*bottomMargin, nameLength, 1)

	addTextBlock(pdf, rightMargin, pageHeight-signHeight*bottomMargin+vspace, nameLength, vspace, "Head of Department's Name", true)
	//examiner signature
	addHrWithLen(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin, signatureLength, 1)
	addTextBlock(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin+vspace, signatureLength, vspace, "Signature", true)
	//examiner date
	addHrWithLen(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin, dateLenth, 1)
	addTextBlock(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin+vspace, dateLenth, vspace, "Date", true)

	signHeight = 3

	//
	//examiner name
	addHrWithLen(pdf, rightMargin, pageHeight-signHeight*bottomMargin, nameLength, 1)

	addTextBlock(pdf, rightMargin, pageHeight-signHeight*bottomMargin+vspace, nameLength, vspace, "Examination Officer's Name", true)
	//examiner signature
	addHrWithLen(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin, signatureLength, 1)
	addTextBlock(pdf, rightMargin+spaceLength+nameLength, pageHeight-signHeight*bottomMargin+vspace, signatureLength, vspace, "Signature", true)
	//examiner date
	addHrWithLen(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin, dateLenth, 1)
	addTextBlock(pdf, rightMargin+2*spaceLength+nameLength+signatureLength, pageHeight-signHeight*bottomMargin+vspace, dateLenth, vspace, "Date", true)

}
