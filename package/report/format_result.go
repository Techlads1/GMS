package report

import "fmt"

func formatScore(score int32) string {
	if score == 0 {
		return "-"
	} else {
		return fmt.Sprintf("%d", score)
	}
}
