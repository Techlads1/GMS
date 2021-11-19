package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func GenerateLifetimeNumber(firstName, middleName, lastName, gender string, dob time.Time) string {
	firstName = strings.TrimSpace(firstName)
	middleName = strings.TrimSpace(middleName)
	lastName = strings.TrimSpace(lastName)
	gender = strings.TrimSpace(gender)

	dobStr := dob.Format("02/01/2006") //  strings.TrimSpace(dob) //dob should be in this format: dd/mm/yyyy
	details := fmt.Sprintf("%s|%s|%s|%s|%s", firstName, middleName, lastName, gender, dobStr)

	hash := sha256.Sum256([]byte(details))

	hexHash := hex.EncodeToString(hash[:])
	fmt.Println(hexHash)
	l := len(hexHash)
	//lluid := hexHash[0:4] + hexHash[l-4:l]
	//lluid := hexHash[0:2] + " " + hexHash[2:4] + " " + hexHash[l-4:l-2] + " " + hexHash[l-2:l]
	lluid := hexHash[0:4] + " " + hexHash[l-4:l]
	return strings.ToUpper(lluid)
}

func GenerateLifetimeNumbers(firstName, middleName, lastName, gender string) string {
	firstName = strings.TrimSpace(firstName)
	middleName = strings.TrimSpace(middleName)
	lastName = strings.TrimSpace(lastName)
	gender = strings.TrimSpace(gender)

	//dobStr := dob.Format("02/01/2006") //  strings.TrimSpace(dob) //dob should be in this format: dd/mm/yyyy
	details := fmt.Sprintf("%s|%s|%s|%s|", firstName, middleName, lastName, gender)

	hash := sha256.Sum256([]byte(details))

	hexHash := hex.EncodeToString(hash[:])
	fmt.Println(hexHash)
	l := len(hexHash)
	//lluid := hexHash[0:4] + hexHash[l-4:l]
	//lluid := hexHash[0:2] + " " + hexHash[2:4] + " " + hexHash[l-4:l-2] + " " + hexHash[l-2:l]
	lluid := hexHash[0:4] + " " + hexHash[l-4:l]
	return strings.ToUpper(lluid)
}
