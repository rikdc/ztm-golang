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

import "fmt"

type Player struct {
	health, energy       int
	maxHealth, maxEnergy int

	name string
}

func (player *Player) addHealth(amount int) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "add", amount, "health ->", player.health)
}

func (player *Player) damage(amount int) {
	player.health -= amount
	if player.health < 0 {
		player.health = 0
	}
	fmt.Println(player.name, "add", amount, "health ->", player.health)
}

func (player *Player) spendEnergy(energy int) {
	player.energy += energy
	if player.energy < 0 {
		player.energy = 0
	}
	fmt.Println(player.name, "add", energy, "energy ->", player.energy)

}

func (player *Player) addEnergy(energy int) {
	player.energy += energy
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println(player.name, "add", energy, "energy ->", player.energy)
}

func (player *Player) showStats() {
	fmt.Println("---")
	fmt.Println(player.name)
	fmt.Println("Health: ", player.health)
	fmt.Println("Energy: ", player.energy)
}

func main() {
	player := Player{
		name:      "Knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.damage(50)
	player.addEnergy(100)
	player.spendEnergy(50)
	player.addHealth(5)

	player.showStats()
}
