package stats

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/n8os/nhlgo/client"
)

const teamsEndpoint string = "/teams"

type Teams struct {
	Copyright string `json:"copyright"`
	Teams     []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Link  string `json:"link"`
		Venue struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Link     string `json:"link"`
			City     string `json:"city"`
			TimeZone struct {
				ID     string `json:"id"`
				Offset int    `json:"offset"`
				Tz     string `json:"tz"`
			} `json:"timeZone"`
		} `json:"venue"`
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
		Roster struct {
			Roster []struct {
				Person struct {
					ID       int    `json:"id"`
					FullName string `json:"fullName"`
					Link     string `json:"link"`
				} `json:"person"`
				JerseyNumber string `json:"jerseyNumber"`
				Position     struct {
					Code         string `json:"code"`
					Name         string `json:"name"`
					Type         string `json:"type"`
					Abbreviation string `json:"abbreviation"`
				} `json:"position"`
			} `json:"roster"`
			Link string `json:"link"`
		} `json:"roster"`
		ShortName       string `json:"shortName"`
		OfficialSiteURL string `json:"officialSiteUrl"`
		FranchiseID     int    `json:"franchiseId"`
		Active          bool   `json:"active"`
	} `json:"teams"`
}

type Roster struct {
	Copyright string `json:"copyright"`
	Roster    []struct {
		Person struct {
			ID       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"person"`
		JerseyNumber string `json:"jerseyNumber"`
		Position     struct {
			Code         string `json:"code"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			Abbreviation string `json:"abbreviation"`
		} `json:"position"`
	} `json:"roster"`
	Link string `json:"link"`
}

// GET https://statsapi.web.nhl.com/api/v1/teams/
func GetTeams(client *client.Client) (Teams, error) {
	response, err := client.GetRequest(teamsEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	var result Teams
	if err := json.Unmarshal(response, &result); err != nil { // Parse []byte to the go struct pointer
		log.Fatal(err)
	}
	return result, nil
}

// local /static/teams.json
func GetTeamsFile() (Teams, error) {
	fileBytes, err := os.ReadFile("./static/teams.json")
	if err != nil {
		log.Fatal(err)
	}
	var sTeams Teams
	if err := json.Unmarshal(fileBytes, &sTeams); err != nil {
		log.Fatal(err)
	}
	return sTeams, nil
}

// GET https://statsapi.web.nhl.com/api/v1/teams/ID
func GetTeam(client *client.Client, teamID int) (Teams, error) {
	response, err := client.GetRequest(fmt.Sprintf("%v/%v", teamsEndpoint, teamID))
	if err != nil {
		log.Fatal(err)
	}
	var result Teams
	if err := json.Unmarshal(response, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}

// GET https://statsapi.web.nhl.com/api/v1/teams/ID?expand=team.roster
func GetRoster(client *client.Client, teamID int) (Teams, error) {
	response, err := client.GetRequest(fmt.Sprintf("%v/%v?expand=team.roster", teamsEndpoint, teamID))
	if err != nil {
		log.Fatal(err)
	}
	var result Teams
	if err := json.Unmarshal(response, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}
