package report

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/signintech/gopdf"
)

func splitString(pdf *gopdf.GoPdf, input *string, width float64) []string {
	re := regexp.MustCompile(`[[:cntrl:]]|[\x{FFFD}]`)
	pureString := re.ReplaceAllString(*input, "")
	lines, err := pdf.SplitText(pureString, width)
	if err != nil {
		return []string{pureString}
	}
	var emptyString int
	for i, v := range lines {
		if v == "" {
			copy(lines[i:], lines[i+1:])
			lines = lines[:len(lines)-1]
			emptyString++
		}
	}
	if emptyString > 0 {
		lines = lines[:len(lines)-emptyString]
	}
	return lines
}

func addMultiLines(pdf *gopdf.GoPdf, x, deltaY float64, lines []string) {
	for _, line := range lines {
		pdf.SetX(x)
		pdf.Cell(nil, line)
		pdf.SetY(pdf.GetY() + deltaY)
	}
}

func addMultiLineBlock(pdf *gopdf.GoPdf, x, y, lineWidth, lineHeight float64, text string, alignCentre bool) (float64, float64) {
	//lines, _ := pdf.SplitText(title, 200)
	//titleWidth := availablePageWidth - logoSize - qrSize - 50
	lines := wrapTextLines(pdf, text, lineWidth-2*padding)
	for _, line := range lines {
		x, y = addTextBlock(pdf, x, y, lineWidth, lineHeight, line, alignCentre)
	}
	return x, y
}

// wrapTextLines splits a string into multiple lines so that the text
// fits in the specified width. The text is wrapped on word boundaries.
// Newline characters ("\r" and "\n") also cause text to be split.
// You can find out the number of lines needed to wrap some
// text by checking the length of the returned array.
func wrapTextLines(pdf *gopdf.GoPdf, text string, width float64) (ret []string) {
	// isWhiteSpace returns true if all the chars. in 's' are white-spaces
	isWhiteSpace := func(s string) bool {
		for _, r := range s {
			if !unicode.IsSpace(r) {
				return false
			}
		}
		return len(s) > 0
	}

	// splitLines splits 's' into several lines using line breaks in 's'
	splitLines := func(s string) []string {
		split := func(lines []string, sep string) (ret []string) {
			for _, line := range lines {
				if strings.Contains(line, sep) {
					ret = append(ret, strings.Split(line, sep)...)
					continue
				}
				ret = append(ret, line)
			}
			return ret
		}
		return split(split(split([]string{s}, "\r\n"), "\r"), "\n")
	} //

	fit := func(s string, step, n int, width float64) int {
		for max := len(s); n > 0 && n <= max; {
			w, _ := pdf.MeasureTextWidth(s[:n]) //.GetTextWidth(s[:n])
			switch step {
			case 1, 3: //       keep halving (or - 1) until n chars fit in width
				if w <= width {
					return n
				}
				n--
				if step == 1 {
					n /= 2
				}
			case 2: //               increase n until n chars won't fit in width
				if w > width {
					return n
				}
				n = 1 + int((float64(n) * 1.1)) //    increase n by 1 + 20% of n
			}
		}
		return 0
	}
	// split text into lines. then break lines based on text width
	//te := splitLines(text)
	//fmt.Printf("%v\n", te)
	for _, line := range splitLines(text) {
		w, err := pdf.MeasureTextWidth(line)
		if err != nil {
			fmt.Printf("error measuring text widith: %v\n", err)
		}
		for w > width {
			n := len(line) //    reduce, increase, then reduce n to get best fit
			for i := 1; i <= 3; i++ {
				n = fit(line, i, n, width)
			}
			// move to the last word (if white-space is found)
			found, max := false, n
			for n > 0 {
				if isWhiteSpace(line[n-1 : n]) {
					found = true
					break
				}
				n--
			}
			if !found {
				n = max
			}
			if n <= 0 {
				break
			}
			ret = append(ret, line[:n])
			line = line[n:]
		}
		ret = append(ret, line)
	}
	return ret
}
