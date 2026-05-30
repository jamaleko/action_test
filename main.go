package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type MatchResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`

	HomeTeam struct {
		Name string `json:"name"`
	} `json:"homeTeam"`

	AwayTeam struct {
		Name string `json:"name"`
	} `json:"awayTeam"`

	Score struct {
		FullTime struct {
			Home *int `json:"home"`
			Away *int `json:"away"`
		} `json:"fullTime"`
	} `json:"score"`
}

func main() {

	req, err := http.NewRequest(
		"GET",
		"https://api.football-data.org/v4/matches/552096",
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set(
		"X-Auth-Token",
		os.Getenv("TOKEN"),
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var match MatchResponse

	err = json.NewDecoder(resp.Body).Decode(&match)
	if err != nil {
		panic(err)
	}

	fmt.Println("STATUS:", match.Status)

	home := "-"
	away := "-"

	if match.Score.FullTime.Home != nil {
		home = fmt.Sprintf("%d", *match.Score.FullTime.Home)
	}

	if match.Score.FullTime.Away != nil {
		away = fmt.Sprintf("%d", *match.Score.FullTime.Away)
	}

	fmt.Println("SCORE:", home+"-"+away)

	fmt.Println(
		match.HomeTeam.Name,
		"vs",
		match.AwayTeam.Name,
	)
}
