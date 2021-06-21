package fakeout

import (
	"example.com/hello/core"
)

type Fakeout struct {
	core.Game
}

func (g *Fakeout) Name() string {
	return "fakeout"
}

func Init() core.Gamelike {
	buildQuestions()
	return &Fakeout{
		Game: *core.NewGame(NewFakeoutHub, NewFakeoutClient),
	}
}
