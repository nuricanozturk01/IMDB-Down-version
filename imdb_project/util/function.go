package util

import "log"

const PUBLIC_PREFIX = "/api/v1/public"
const PRIVATE_PREFIX = "/api/v1"

func CheckError(err error, message string) {
	if err != nil {
		log.Panic(message, err)
	}
}
