package stats

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Standings struct {
	Copyright string `json:"copyright"`
	Records   []struct {
		StandingsType string `json:"standingsType"`
		League        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"league"`
		Division struct {
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
		TeamRecords []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"leagueRecord"`
			RegulationWins     int    `json:"regulationWins"`
			GoalsAgainst       int    `json:"goalsAgainst"`
			GoalsScored        int    `json:"goalsScored"`
			Points             int    `json:"points"`
			DivisionRank       string `json:"divisionRank"`
			DivisionL10Rank    string `json:"divisionL10Rank"`
			DivisionRoadRank   string `json:"divisionRoadRank"`
			DivisionHomeRank   string `json:"divisionHomeRank"`
			ConferenceRank     string `json:"conferenceRank"`
			ConferenceL10Rank  string `json:"conferenceL10Rank"`
			ConferenceRoadRank string `json:"conferenceRoadRank"`
			ConferenceHomeRank string `json:"conferenceHomeRank"`
			LeagueRank         string `json:"leagueRank"`
			LeagueL10Rank      string `json:"leagueL10Rank"`
			LeagueRoadRank     string `json:"leagueRoadRank"`
			LeagueHomeRank     string `json:"leagueHomeRank"`
			WildCardRank       string `json:"wildCardRank"`
			Row                int    `json:"row"`
			GamesPlayed        int    `json:"gamesPlayed"`
			Streak             struct {
				StreakType   string `json:"streakType"`
				StreakNumber int    `json:"streakNumber"`
				StreakCode   string `json:"streakCode"`
			} `json:"streak"`
			ClinchIndicator  string    `json:"clinchIndicator,omitempty"`
			PointsPercentage float64   `json:"pointsPercentage"`
			PpDivisionRank   string    `json:"ppDivisionRank"`
			PpConferenceRank string    `json:"ppConferenceRank"`
			PpLeagueRank     string    `json:"ppLeagueRank"`
			LastUpdated      time.Time `json:"lastUpdated"`
		} `json:"teamRecords"`
	} `json:"records"`
}

var standingsEndpoint string = "standings"

func GetStandings(base string) Standings {
	url := base + "/" + standingsEndpoint

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		log.Fatal(err)
	}

	var result Standings
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		log.Fatal(err)
	}

	return result
}
