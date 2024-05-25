package util

import "log"

func ForEachIndexedString(slice []string, f func(int, string)) {
	for i, v := range slice {
		f(i, v)
	}
}

func ForEach[T any](slice []T, f func(T)) {
	for _, v := range slice {
		f(v)
	}
}

func CheckError(err error, message string) {
	if err != nil {
		log.Panic(message, err)
	}
}
