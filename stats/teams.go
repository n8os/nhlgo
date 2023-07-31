package stats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"log"
)

type Teams struct {
	Copyright string `json:"copyright"`
	Teams     []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Link  string `json:"link"`
		Venue struct {
			Name     string `json:"name"`
			Link     string `json:"link"`
			City     string `json:"city"`
			TimeZone struct {
				ID     string `json:"id"`
				Offset int    `json:"offset"`
				Tz     string `json:"tz"`
			} `json:"timeZone"`
		} `json:"venue,omitempty"`
		Abbreviation    string `json:"abbreviation"`
		TeamName        string `json:"teamName"`
		LocationName    string `json:"locationName"`
		FirstYearOfPlay string `json:"firstYearOfPlay"`
		Division        struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			NameShort    string `json:"nameShort"`
			Link         string `json:"link"`
			Abbreviation string `json:"abbreviation"`
		} `json:"division"`
		Conference struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"conference"`
		Franchise struct {
			FranchiseID int    `json:"franchiseId"`
			TeamName    string `json:"teamName"`
			Link        string `json:"link"`
		} `json:"franchise"`
		ShortName       string `json:"shortName"`
		OfficialSiteURL string `json:"officialSiteUrl"`
		FranchiseID     int    `json:"franchiseId"`
		Active          bool   `json:"active"`
	} `json:"teams"`
}

var teamsEndpoint string = "teams"

func GetTeams(base string) Teams {
	url := base + teamsEndpoint

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var result Teams
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		log.Fatal(err)
	}

	return result
}

func GetTeam(base string, teamID string) Teams {
	url := fmt.Sprintf("%v%v/%v", base, teamsEndpoint, teamID)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var result Teams
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		log.Fatal(err)
	}

	return result
}
