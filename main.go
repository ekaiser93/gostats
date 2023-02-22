package main

import (
	"fmt"

	"example.com/gostats/calculations"
)

func main() {
	playerOne := calculations.Batter{
		PlayerID:             "troutmi01",
		YearID:               2014,
		Stint:                1,
		TeamID:               "LAA",
		LeagueID:             "AL",
		Games:                157,
		AtBats:               602,
		Runs:                 115,
		Hits:                 173,
		Doubles:              39,
		Triples:              9,
		Homeruns:             36,
		RunsBattedIn:         111,
		StolenBases:          16,
		CaughtStealing:       2,
		Walks:                83,
		Strikeouts:           184,
		IntentionalWalks:     6,
		HitByPitch:           10,
		SacrificeBunt:        0,
		SacrificeFly:         10,
		GroundIntoDoublePlay: 6,
	}

	playerTwo := calculations.Batter{
		PlayerID:             "cabremi01",
		YearID:               2013,
		Stint:                1,
		TeamID:               "DET",
		LeagueID:             "AL",
		Games:                148,
		AtBats:               555,
		Runs:                 103,
		Hits:                 193,
		Doubles:              26,
		Triples:              1,
		Homeruns:             44,
		RunsBattedIn:         137,
		StolenBases:          3,
		CaughtStealing:       0,
		Walks:                90,
		Strikeouts:           94,
		IntentionalWalks:     19,
		HitByPitch:           5,
		SacrificeBunt:        0,
		SacrificeFly:         2,
		GroundIntoDoublePlay: 19,
	}

	playerThree := calculations.Batter{
		PlayerID:             "hamiljo03",
		YearID:               2010,
		Stint:                1,
		TeamID:               "TEX",
		LeagueID:             "AL",
		Games:                133,
		AtBats:               518,
		Runs:                 95,
		Hits:                 186,
		Doubles:              40,
		Triples:              3,
		Homeruns:             32,
		RunsBattedIn:         100,
		StolenBases:          8,
		CaughtStealing:       1,
		Walks:                43,
		Strikeouts:           95,
		IntentionalWalks:     5,
		HitByPitch:           5,
		SacrificeBunt:        1,
		SacrificeFly:         4,
		GroundIntoDoublePlay: 11,
	}

	playerFour := calculations.Batter{
		PlayerID:             "mauerjo01",
		YearID:               2009,
		Stint:                1,
		TeamID:               "MIN",
		LeagueID:             "AL",
		Games:                138,
		AtBats:               523,
		Runs:                 94,
		Hits:                 191,
		Doubles:              30,
		Triples:              1,
		Homeruns:             28,
		RunsBattedIn:         96,
		StolenBases:          4,
		CaughtStealing:       1,
		Walks:                76,
		Strikeouts:           63,
		IntentionalWalks:     14,
		HitByPitch:           2,
		SacrificeBunt:        0,
		SacrificeFly:         5,
		GroundIntoDoublePlay: 13,
	}

	fmt.Println("Play Ball!")

	// create a slice containing our players
	batters := []calculations.Batter{playerOne, playerTwo, playerThree, playerFour}

	// loop through the batters and get the averages for each and print
	for _, batter := range batters {
		calculations.GetAverages(&batter)
		fmt.Println(batter)
	}
}
