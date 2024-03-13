package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	randomSource *rand.Rand
}

func (g *Game) CheckNumber(number int) bool {
	return number == g.randomSource.Int()
}

func main() {
	g := Game{
		randomSource: rand.New(rand.NewSource(95)),
	}
	g.CheckNumber(21)
	g.CheckNumber(12)

	gCopy := Game{
		randomSource: rand.New(rand.NewSource(95)),
	}

	for i := 0; i < 2; i++ {
		gCopy.randomSource.Int()
	}

	winNumber := gCopy.randomSource.Int()

	fmt.Println("Следующим будет число:", winNumber)
	if g.CheckNumber(winNumber) {
		fmt.Println("Вы победили!")
	} else {
		fmt.Println("Вы проиграли :(")
	}
}
