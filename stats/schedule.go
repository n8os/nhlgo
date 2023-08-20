package stats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/n8os/nhlgo/client"
)

const scheduleEndpoint string = "/schedule"

type Schedule struct {
	Copyright    string `json:"copyright"`
	TotalItems   int    `json:"totalItems"`
	TotalEvents  int    `json:"totalEvents"`
	TotalGames   int    `json:"totalGames"`
	TotalMatches int    `json:"totalMatches"`
	MetaData     struct {
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	Wait  int `json:"wait"`
	Dates []struct {
		Date         string `json:"date"`
		TotalItems   int    `json:"totalItems"`
		TotalEvents  int    `json:"totalEvents"`
		TotalGames   int    `json:"totalGames"`
		TotalMatches int    `json:"totalMatches"`
		Games        []struct {
			GamePk   int       `json:"gamePk"`
			Link     string    `json:"link"`
			GameType string    `json:"gameType"`
			Season   string    `json:"season"`
			GameDate time.Time `json:"gameDate"`
			Status   struct {
				AbstractGameState string `json:"abstractGameState"`
				CodedGameState    string `json:"codedGameState"`
				DetailedState     string `json:"detailedState"`
				StatusCode        string `json:"statusCode"`
				StartTimeTBD      bool   `json:"startTimeTBD"`
			} `json:"status"`
			Teams struct {
				Away struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Ot     int    `json:"ot"`
						Type   string `json:"type"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
				} `json:"away"`
				Home struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Ot     int    `json:"ot"`
						Type   string `json:"type"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Link string `json:"link"`
					} `json:"team"`
				} `json:"home"`
			} `json:"teams"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"venue"`
			Content struct {
				Link string `json:"link"`
			} `json:"content"`
		} `json:"games"`
		Events  []any `json:"events"`
		Matches []any `json:"matches"`
	} `json:"dates"`
}

// GET
// ?date=2022-01-02&teamId=14
func GetSchedule(client *client.Client, teamID int, dateStr string) (Schedule, error) {
	response, err := client.GetRequest(fmt.Sprintf("%v?teamId=%v&date=%v", scheduleEndpoint, teamID, dateStr))
	if err != nil {
		log.Fatal(err)
	}
	var result Schedule
	if err := json.Unmarshal(response, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}
