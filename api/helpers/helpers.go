package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func ContainsDomainError(url string) bool {
	return !strings.Contains(url, os.Getenv("DOMAIN"))
}
