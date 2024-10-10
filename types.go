package main

import "time"

type ResponseRoot struct {
	PrevDate    string `json:"prevDate"`
	CurrentDate string `json:"currentDate"`
	NextDate    string `json:"nextDate"`
	GameWeek    []struct {
		Date          string `json:"date"`
		DayAbbrev     string `json:"dayAbbrev"`
		NumberOfGames int    `json:"numberOfGames"`
	} `json:"gameWeek"`
	OddsPartners []struct {
		PartnerID   int    `json:"partnerId"`
		Country     string `json:"country"`
		Name        string `json:"name"`
		ImageURL    string `json:"imageUrl"`
		SiteURL     string `json:"siteUrl,omitempty"`
		BgColor     string `json:"bgColor"`
		TextColor   string `json:"textColor"`
		AccentColor string `json:"accentColor"`
	} `json:"oddsPartners"`
	Games []struct {
		ID       int    `json:"id"`
		Season   int    `json:"season"`
		GameType int    `json:"gameType"`
		GameDate string `json:"gameDate"`
		Venue    struct {
			Default string `json:"default"`
		} `json:"venue"`
		StartTimeUTC     time.Time `json:"startTimeUTC"`
		EasternUTCOffset string    `json:"easternUTCOffset"`
		VenueUTCOffset   string    `json:"venueUTCOffset"`
		TvBroadcasts     []struct {
			ID             int    `json:"id"`
			Market         string `json:"market"`
			CountryCode    string `json:"countryCode"`
			Network        string `json:"network"`
			SequenceNumber int    `json:"sequenceNumber"`
		} `json:"tvBroadcasts"`
		GameState         string `json:"gameState"`
		GameScheduleState string `json:"gameScheduleState"`
		AwayTeam          struct {
			ID   int `json:"id"`
			Name struct {
				Default string `json:"default"`
			} `json:"name"`
			Abbrev string `json:"abbrev"`
			Score  int    `json:"score"`
			Sog    int    `json:"sog"`
			Logo   string `json:"logo"`
		} `json:"awayTeam"`
		HomeTeam struct {
			ID   int `json:"id"`
			Name struct {
				Default string `json:"default"`
			} `json:"name"`
			Abbrev string `json:"abbrev"`
			Score  int    `json:"score"`
			Sog    int    `json:"sog"`
			Logo   string `json:"logo"`
		} `json:"homeTeam"`
		GameCenterLink  string `json:"gameCenterLink"`
		ThreeMinRecap   string `json:"threeMinRecap"`
		ThreeMinRecapFr string `json:"threeMinRecapFr"`
		CondensedGame   string `json:"condensedGame"`
		CondensedGameFr string `json:"condensedGameFr"`
		Clock           struct {
			TimeRemaining    string `json:"timeRemaining"`
			SecondsRemaining int    `json:"secondsRemaining"`
			Running          bool   `json:"running"`
			InIntermission   bool   `json:"inIntermission"`
		} `json:"clock"`
		NeutralSite      bool   `json:"neutralSite"`
		VenueTimezone    string `json:"venueTimezone"`
		Period           int    `json:"period"`
		PeriodDescriptor struct {
			Number               int    `json:"number"`
			PeriodType           string `json:"periodType"`
			MaxRegulationPeriods int    `json:"maxRegulationPeriods"`
		} `json:"periodDescriptor"`
		GameOutcome struct {
			LastPeriodType string `json:"lastPeriodType"`
		} `json:"gameOutcome"`
		Goals []struct {
			Period           int `json:"period"`
			PeriodDescriptor struct {
				Number               int    `json:"number"`
				PeriodType           string `json:"periodType"`
				MaxRegulationPeriods int    `json:"maxRegulationPeriods"`
			} `json:"periodDescriptor"`
			TimeInPeriod string `json:"timeInPeriod"`
			PlayerID     int    `json:"playerId"`
			Name         struct {
				Default string `json:"default"`
			} `json:"name"`
			FirstName struct {
				Default string `json:"default"`
			} `json:"firstName"`
			LastName struct {
				Default string `json:"default"`
			} `json:"lastName"`
			GoalModifier string `json:"goalModifier"`
			Assists      []struct {
				PlayerID int `json:"playerId"`
				Name     struct {
					Default string `json:"default"`
				} `json:"name"`
				AssistsToDate int `json:"assistsToDate"`
			} `json:"assists"`
			Mugshot                   string `json:"mugshot"`
			TeamAbbrev                string `json:"teamAbbrev"`
			GoalsToDate               int    `json:"goalsToDate"`
			AwayScore                 int    `json:"awayScore"`
			HomeScore                 int    `json:"homeScore"`
			Strength                  string `json:"strength"`
			HighlightClipSharingURL   string `json:"highlightClipSharingUrl"`
			HighlightClipSharingURLFr string `json:"highlightClipSharingUrlFr"`
			HighlightClip             int64  `json:"highlightClip"`
			HighlightClipFr           int64  `json:"highlightClipFr"`
			DiscreteClip              int64  `json:"discreteClip"`
			DiscreteClipFr            int64  `json:"discreteClipFr"`
		} `json:"goals"`
	} `json:"games"`
}

type GamesResponse struct {
	Period    string `json:"period"`
	Time      string `json:"time"`
	HomeTeam  string `json:"home_team"`
	AwayTeam  string `json:"away_team"`
	HomeLogo  string `json:"home_logo"`
	AwayLogo  string `json:"away_logo"`
	HomeScore int    `json:"home_score"`
	AwayScore int    `json:"away_score"`
	HomeSog   int    `json:"home_sog"`
	AwaySog   int    `json:"away_sog"`
}
