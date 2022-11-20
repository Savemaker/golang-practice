package main

type NHL struct {
	homeTeam      string
	homeTeamScore int
	awayTeam      string
	awayTeamScore int
}

func (n NHL) Play() (string, error) {
	return ResolveScore(n.homeTeam, n.awayTeam, n.homeTeamScore, n.awayTeamScore)
}
