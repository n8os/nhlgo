package main

import (
	"encoding/json"
	"fmt"

	"github.com/n8os/nhlgo/teams"
)

var base_url = "https://statsapi.web.nhl.com/api/v1/"

// todo, renewable http.Client defined in this file

func main() {
	// todo, router (i.e. /teams) aka "/teams" on discord
	if true {
		TeamStruct := teams.GetTeams(base_url)
		// for _, team := range TeamStruct.Teams {
		// 	fmt.Println(team.Name)
		// }
		fmt.Println(PrettyPrint(TeamStruct))
	}
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
