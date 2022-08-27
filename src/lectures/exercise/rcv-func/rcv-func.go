//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type HYDPlayer struct {
	name      string
	health    int
	maxHealth int
	energy    int
	maxEnergy int
}

// method-2 receving func
func changeHealth(hyd *HYDPlayer) {

	hyd.health = rand.Intn(5008)
}

// method-2 receving func
func changeMaxHealth(hyd *HYDPlayer) {

	hyd.maxHealth = rand.Intn(7008)
}

// method-1 receving func
func (hyd *HYDPlayer) changeEnergy() {

	hyd.energy = rand.Intn(8800)
}

// method-1 receving func
func (hyd *HYDPlayer) changeMaxEnergy() {

	hyd.maxEnergy = rand.Intn(3800)
}

func main() {
	var hydp HYDPlayer
	hydp.name = "Krish Health Stats"
	rand.Seed(time.Now().UnixNano())
	changeHealth(&hydp)
	changeMaxHealth(&hydp)
	hydp.changeEnergy()
	hydp.changeMaxEnergy()
	fmt.Println("HYD Player", hydp)
}
