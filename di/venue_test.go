package main

import (
	"strings"
	"testing"
)

func TestVenue(t *testing.T) {

	nhl := NHL{
		homeTeam:      "LA Kings",
		homeTeamScore: 3,
		awayTeam:      "Toronto",
		awayTeamScore: 2,
	}

	apl := APL{
		homeTeam:      "Manchester City",
		homeTeamScore: 3,
		awayTeam:      "Liverpool",
		awayTeamScore: 4,
	}

	t.Run("test NHL Kings Toronto game is happening and Kings wins", func(t *testing.T) {
		venue := Venue{
			tonight: &nhl,
		}
		gameResult, _ := venue.start()
		if !strings.Contains(gameResult, "LA Kings") {
			t.Fatalf("Kings did not win. Game result %q", gameResult)
		}
	})

	t.Run("test APL game is happening and Liverpool wins", func(t *testing.T) {
		venue := Venue{
			tonight: &apl,
		}
		gameResult, _ := venue.start()
		if !strings.Contains(gameResult, "Liverpool") {
			t.Fatalf("Liverpool did not win. Game result %q", gameResult)
		}
	})
}
