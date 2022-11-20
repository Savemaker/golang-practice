package main

const (
	ErrNoResult = GameError("could not resolve winner")
)

type GameError string

func (g GameError) Error() string {
	return string(g)
}

type Game interface {
	Play() (string, error)
}

func ResolveScore(homeTeam string, awayTeam string, homeTeamScore int, awayTeamScore int) (string, error) {
	switch homeTeamScore > awayTeamScore {
	case true:
		return homeTeam + " win!", nil
	case false:
		return awayTeam + " win!", nil
	default:
		return "", ErrNoResult
	}
}
