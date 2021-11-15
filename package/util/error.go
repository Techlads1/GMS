package util

func CheckError(err error) bool {
	if err != nil {
		//log.Error(err)
		return true
	} else {
		return false
	}
}
