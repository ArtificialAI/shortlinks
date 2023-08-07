package utils

import (
	"log"
	"net/url"
	"os"
)

func IsValidUrl(str string) (bool, string) {
	log.Printf("Validating '%s'", str)

	u, err := url.Parse(str)

	if err != nil {
		return false, "Invalid url format"
	}

	if (u.Scheme == "") && (os.Getenv("IS_URL_SCHEME_REQUIRED") == "true") {
		return false, "Missing scheme"
	}

	if u.Host == "" {
		return false, "Missing host"
	}

	return true, ""
}
