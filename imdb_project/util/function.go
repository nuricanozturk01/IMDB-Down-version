package util

import "log"

func CheckError(err error, message string) {
	if err != nil {
		log.Panic(message, err)
	}
}
