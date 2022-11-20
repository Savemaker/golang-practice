package main

type Venue struct {
	tonight Game
}

func (v Venue) start() (string, error) {
	return v.tonight.Play()
}
