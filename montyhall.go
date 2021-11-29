package main

import (
	"fmt"
	"math/rand"
	"time"
)

type doors []bool

func (d *doors) init() {
	rand.Seed(time.Now().UnixNano())
	*d = make(doors, 3)
	(*d)[rand.Intn(len(*d))] = true
}

func (d *doors) pick() int {
	return rand.Intn(len(*d))
}

func (d *doors) eliminate(picked int) int {
	notPicked := []int{}
	for i := 0; i < len(*d); i++ {
		if i == picked {
			continue
		}
		notPicked = append(notPicked, i)
	}

	if (*d)[notPicked[0]] {
		//fmt.Println("Hostmaster must choose:")
		return notPicked[1]
	} else if (*d)[notPicked[1]] {
		//fmt.Println("Hostmaster must choose:")
		return notPicked[0]
	} else {
		//fmt.Println("Hostmaster random choice: ", notPicked)
		rand.Seed(time.Now().UnixNano())
		return notPicked[rand.Intn(len(notPicked))]
	}
}

func play(changeDoors bool) bool {
	var d doors

	d.init()
	fmt.Println("Doors:", d)

	picked := d.pick()
	fmt.Println("Player picked:", picked)

	eliminated := d.eliminate(picked)
	fmt.Println("Showmaster eliminated:", eliminated)

	if changeDoors {
		newPicked := -1
		for i := 0; i < len(d); i++ {
			if i == picked || i == eliminated {
				continue
			}
			newPicked = i
			break
		}
		fmt.Println("Player switches doors from:", picked, "; to: ", newPicked)
		picked = newPicked
	} else {
		fmt.Println("Player sticks to their choice")
	}

	hasWon := d[picked]
	if hasWon {
		fmt.Println("Player won")
	} else {
		fmt.Println("Player lost")
	}
	fmt.Println("")
	return hasWon
}

func main() {
	numGames := 100000
	changeDoors := true
	won := 0
	for i := 0; i < numGames; i++ {
		if play(changeDoors) {
			won++
		}
	}
	percentage := float64(won) / float64(numGames)
	fmt.Println("Won: ", won, "; Total: ", numGames, "; %: ", percentage)
}
