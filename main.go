package main

import (
	"encoding/json"
	"fmt"

	"github.com/n8os/nhlgo/stats"
)

var baseUrl = "https://statsapi.web.nhl.com/api/v1/"

// todo, renewable http.Client defined in this file

// type Client struct {
// 	BaseUrl string
// }

// func NewClient(baseUrl *string) *Client {
// 	var url string

// 	if baseUrl == nil {
// 		url = "https://statsapi.web.nhl.com/api/v1"
// 	} else {
// 		url = *baseUrl
// 	}

// 	return &Client{
// 		BaseUrl: url,
// 	}
// }

func main() {
	// todo, router (i.e. /teams) aka "/teams" on discord
	if true {
		// TeamStruct := teams.GetTeams(base_url)
		// for _, team := range TeamStruct.Teams {
		// 	fmt.Println(team.Name)
		// }

		// fmt.Println(prettyPrint(stats.GetTeams(base_url)))
		// fmt.Println(prettyPrint(stats.GetTeam(base_url, "15")))

		fmt.Println(prettyPrint(stats.GetPerson(baseUrl, "8474157")))

	}
}

// PrettyPrint to print struct in a readable way
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
