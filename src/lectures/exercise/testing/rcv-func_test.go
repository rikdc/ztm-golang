//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "Knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}
}

func TestHealthRanges(t *testing.T) {

	player := newPlayer()

	player.addHealth(100)

	if player.health > player.maxHealth {
		t.Fatalf("health should not exceed %v", player.maxHealth)
	}

	player.damage(player.maxHealth + 1)
	if player.health < 0 {
		t.Fatalf("health cannot be below zero, currently %v", player.health)
	}

	if player.health > player.maxHealth {
		t.Fatalf("health should not exceed %v", player.maxHealth)
	}
}

func TestEnergyRanges(t *testing.T) {

	player := newPlayer()

	player.addEnergy(100)

	if player.energy > player.maxEnergy {
		t.Fatalf("energy should not exceed %v", player.maxEnergy)
	}

	player.damage(player.maxEnergy + 1)
	if player.energy < 0 {
		t.Fatalf("energy cannot be below zero, currently %v", player.energy)
	}

	if player.energy > player.maxEnergy {
		t.Fatalf("energy should not exceed %v", player.maxEnergy)
	}
}
