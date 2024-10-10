package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getData())
}

func getPeriod(period int) string {
	suffix := ""
	if period == 1 {
		suffix = "ST"
	} else if period == 2 {
		suffix = "ND"
	} else if period == 3 {
		suffix = "RD"
	} else {
		suffix = "TH"
	}

	return (fmt.Sprintf("%x", period) + suffix)
}

func parseHourMinuteDuration(s string) (time.Duration, error) {
	// Split the string into hours and minutes
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format: %s", s)
	}

	// Parse hours
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hours: %s", parts[0])
	}

	// Parse minutes
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %s", parts[1])
	}

	// Convert to duration
	duration := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute

	return duration, nil
}
func parseTimezoneOffset(offsetString string) (int, error) {
	offset, err := parseHourMinuteDuration(offsetString)
	if err != nil {
		return 0, err
	}
	return int(offset.Seconds()), nil
}

func formatTime(timestamp time.Time, timezoneOffset int) string {
	adjustedTime := timestamp.Add(time.Duration(timezoneOffset) * time.Second)
	return adjustedTime.Format("3:04 PM")
}
func getTodaysDate() string {
	return time.Now().Format("2006-01-02")
}
func getData() string {
	res, err := http.Get("https://api-web.nhle.com/v1/score/" + getTodaysDate())
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data ResponseRoot
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	var games []GamesResponse

	for _, game := range data.Games {
		var time, period string
		if game.GameState == "LIVE" || game.GameState == "CRIT" {
			period = getPeriod(game.PeriodDescriptor.Number)
			time = game.Clock.TimeRemaining
		} else if game.GameState == "PRE" || game.GameState == "FUT" {
			period = game.GameState
			offset, err := parseTimezoneOffset(game.EasternUTCOffset)
			if err != nil {
				panic(err)
			}
			time = formatTime(game.StartTimeUTC, offset)
		} else {
			period = game.GameState
			time = "FINAL"
		}

		homeLogo := game.HomeTeam.Abbrev[0:1]
		awayLogo := game.AwayTeam.Abbrev[0:1]
		coercedGame := &GamesResponse{
			Period:    period,
			Time:      time,
			HomeTeam:  game.HomeTeam.Name.Default,
			AwayTeam:  game.AwayTeam.Name.Default,
			HomeLogo:  homeLogo,
			AwayLogo:  awayLogo,
			HomeScore: game.HomeTeam.Score,
			AwayScore: game.AwayTeam.Score,
			HomeSog:   game.HomeTeam.Sog,
			AwaySog:   game.AwayTeam.Sog,
		}
		games = append(games, *coercedGame)
	}

	gamesBytes, err := json.Marshal(games)
	if err != nil {
		panic(err)
	}

	return string(gamesBytes)
}

func main() {
	go getData()
	http.HandleFunc("/", greet)
	http.ListenAndServe(":1234", nil)
}
