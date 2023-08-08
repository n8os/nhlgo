package stats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"log"
)

type Person struct {
	Copyright string `json:"copyright"`
	People    []struct {
		ID                 int    `json:"id"`
		FullName           string `json:"fullName"`
		Link               string `json:"link"`
		FirstName          string `json:"firstName"`
		LastName           string `json:"lastName"`
		PrimaryNumber      string `json:"primaryNumber"`
		BirthDate          string `json:"birthDate"`
		CurrentAge         int    `json:"currentAge"`
		BirthCity          string `json:"birthCity"`
		BirthStateProvince string `json:"birthStateProvince"`
		BirthCountry       string `json:"birthCountry"`
		Nationality        string `json:"nationality"`
		Height             string `json:"height"`
		Weight             int    `json:"weight"`
		Active             bool   `json:"active"`
		AlternateCaptain   bool   `json:"alternateCaptain"`
		Captain            bool   `json:"captain"`
		Rookie             bool   `json:"rookie"`
		ShootsCatches      string `json:"shootsCatches"`
		RosterStatus       string `json:"rosterStatus"`
		CurrentTeam        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"currentTeam"`
		PrimaryPosition struct {
			Code         string `json:"code"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			Abbreviation string `json:"abbreviation"`
		} `json:"primaryPosition"`
	} `json:"people"`
}

var endpoint string = "people"

func GetPerson(base string, personid string) Person {
	url := fmt.Sprintf("%v/%v/%v", base, endpoint, personid)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var result Person
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		log.Fatal(err)
	}

	return result
}
