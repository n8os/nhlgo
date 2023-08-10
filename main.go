package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/n8os/nhlgo/stats"
)

var baseUrl = "https://statsapi.web.nhl.com/api/v1"

// TODO: renewable http.Client defined in this file
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

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func main() {
	// TODO: router (i.e. /teams) aka "/teams" on discord
	var output, entityId, datestr string
	var bail bool
	// slice holdering arguments
	options := []string{"team", "person", "schedule", "standings", "teams", "exit"}

	// init message
	fmt.Println("Ready:")

	// commandline args
	// optional
	// TODO: move
	// fmt.Println(len(os.Args), os.Args)

	// infinite loop, awaiting input via scanner
	// interactive shell against api
	// TODO this is temporary
	for {
		// get input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		// team id second argument
		if len(input) > 1 {
			entityId = input[1]
		}
		// switch based on input
		switch input[0] {
		case "teams":
			// output = prettyPrint(stats.GetTeams(baseUrl))

			// get json from static file instead of query
			staticTeams := stats.GetTeamsFile()
			output += "ID TEAM\n-------\n"
			for i := 0; i < len(staticTeams.Teams); i++ {
				output += fmt.Sprintf("%v %v\n", staticTeams.Teams[i].ID, staticTeams.Teams[i].Name)
			}
		case "team":
			// optional from command line args
			// if Contains(os.Args, "-team") {
			// 	teamid = "16"
			// }
			teamStruct := stats.GetTeam(baseUrl, entityId)
			output = teamStruct.Teams[0].Name
		case "person":
			// entityId = "8474157"
			output = prettyPrint(stats.GetPerson(baseUrl, entityId))
		case "schedule":
			// date
			if len(input) > 2 {
				datestr = input[2]
			}
			fmt.Println(datestr)
			output = prettyPrint(stats.GetSchedule(baseUrl, entityId, datestr))
		case "standings":
			output = prettyPrint(stats.GetStandings(baseUrl))
		case "roster":
			roster := stats.GetRoster(baseUrl, entityId)
			output = roster.Teams[0].Name + "\n"
			// output = prettyPrint(roster)
			players := roster.Teams[0].Roster.Roster
			for i := range players {
				output += fmt.Sprintf("\n%v [%v]\n", players[i].Person.FullName, players[i].Person.ID)
			}
			// output = roster.Teams[0].Roster.Roster[0]
		case "exit":
			// exit by breaking loop
			bail = true
		default:
			// persistent note
			output += "\n\n----------GO NHL---------\nAvailable Options:\n"
			for _, element := range options {
				output += "    " + element
			}
		}

		fmt.Println(output)

		// break the forloop
		if bail {
			break
		}
	}
}

// PrettyPrint to print struct in a readable way
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
