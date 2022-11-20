package main

type APL struct {
	homeTeam      string
	homeTeamScore int
	awayTeam      string
	awayTeamScore int
}

func (a APL) Play() (string, error) {
	return ResolveScore(a.homeTeam, a.awayTeam, a.homeTeamScore, a.awayTeamScore)
}
