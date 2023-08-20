package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/n8os/nhlgo/client"
	"github.com/n8os/nhlgo/stats"
)

func main() {
	var output string
	// flags package doesn't parse after argument so just use flags for now
	// TODO: something about that?
	// if len(os.Args) < 2 {
	// 	log.Fatalf("Usage: %s ID\n\nExample: %[1]s 15 (Washington Captials),", os.Args[0])
	// }
	// initArg := os.Args[1]
	// fmt.Println(initArg)

	// flags
	initArg := flag.String("cmd", "teams", "Type of query to run")
	entityID := flag.Int("id", 15, "Entity ID to query")
	dateStr := flag.String("date", "", "Date string in format YYYY-MM-DD")
	flag.Parse()

	client := client.NewClient()

	switch *initArg {
	case "teams":
		// get json from static file instead of query
		staticTeams, err := stats.GetTeamsFile()
		if err != nil {
			fmt.Println(err)
		}
		output += "ID TEAM\n-------\n"
		for i := 0; i < len(staticTeams.Teams); i++ {
			output += fmt.Sprintf("%v %v\n", staticTeams.Teams[i].ID, staticTeams.Teams[i].Name)
		}
	case "team":
		teamStruct, err := stats.GetTeam(client, *entityID)
		if err != nil {
			fmt.Println(err)
		}
		output = teamStruct.Teams[0].Name
		// output = prettyPrint(teamStruct)
	case "person":
		if *entityID == 15 {
			*entityID = 8471214 // OV as default
		}
		// entityID = "8474157"
		personStruct, err := stats.GetPerson(client, *entityID)
		if err != nil {
			fmt.Println(err)
		}
		output = prettyPrint(personStruct)
	case "schedule":
		schedStruct, err := stats.GetSchedule(client, *entityID, *dateStr)
		if err != nil {
			fmt.Println(err)
		}
		output = prettyPrint(schedStruct)
	case "standings":
		standStruct, err := stats.GetStandings(client)
		if err != nil {
			fmt.Println(err)
		}
		output = prettyPrint(standStruct)
	case "roster":
		roster, err := stats.GetRoster(client, *entityID)
		if err != nil {
			fmt.Println(err)
		}
		output = roster.Teams[0].Name + "\n"
		players := roster.Teams[0].Roster.Roster
		for i := range players {
			output += fmt.Sprintf("\n%v [%v]\n", players[i].Person.FullName, players[i].Person.ID)
		}
	default:
		output = "Command not found, try --help"
	}

	fmt.Println(output)
}

// PrettyPrint to print struct in a readable way
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
