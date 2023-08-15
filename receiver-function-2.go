package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, MaxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "Add", amount, "health -> ", player.health)
}

// Function that applies
func (player *Player) applyDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "Damage", amount, "->", player.health)
}

//Add energy
func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.MaxEnergy {
		player.energy = player.MaxEnergy
	}
	fmt.Println(player.name, "Add", amount, "energy -> ", player.energy)
}

// Consume Energy
func (player *Player) consumeEnergy(amount uint) {
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Consume", amount, "energy ->", player.energy)
}

//Implement receiver functions to create stat modification for a video game character
func main() {
	player := Player{
		name:      "knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		MaxEnergy: 500,
	}

	player.applyDamage(99)
	player.addHealth(10)
	player.consumeEnergy(99)
	player.addEnergy(10)

	player.consumeEnergy(9999)
}
