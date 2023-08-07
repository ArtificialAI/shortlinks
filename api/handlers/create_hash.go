package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"shortlinks/db"
	"shortlinks/utils"
)

type body struct {
	Link string
}

func generateHash() string {
	// choose the minimum number of characters depending
	// on the number of links stored in the database
	// 4 => ~7M
	// 5 => ~380M
	// 6 => ~19B
	return utils.RandomString(5)
}

func CreateHash(r *http.Request) (int, string) {
	decoder := json.NewDecoder(r.Body)
	var data body
	err := decoder.Decode(&data)
	if err != nil {
		return http.StatusBadRequest, ""
	}

	link := data.Link
	if link == "" {
		return http.StatusBadRequest, ""
	}

	isValid, message := utils.IsValidUrl(link)
	if !isValid {
		return http.StatusForbidden, `{"message":"` + message + `" }`
	}

	// we should not make new hashes if the URL is already existing
	if existedHash, _ := db.ExistenceCheck(link); existedHash != "" {
		log.Printf("An existed entry found: '%v' with hash: '%v'", link, existedHash)
		return http.StatusCreated, `{"hash":"` + existedHash + `"}`
	}

	hash := generateHash()

	log.Printf("Saving link '%s' with hash '%s'", link, hash)
	if err := db.InsertLink(hash, link); err != nil {
		log.Printf("Error executing insert-link.sql: %v\n", err)
		return http.StatusInternalServerError, ""
	}

	return http.StatusCreated, `{"hash":"` + hash + `"}`
}
