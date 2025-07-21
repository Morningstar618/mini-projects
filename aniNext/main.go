package main

import (
	"aninext/lib/player"
	"aninext/lib/seasons"
	"fmt"
)

func main() {
	aniList := seasons.GetSeasonsNow()

	fmt.Print("\nPlease press Cntrl+C to exit or select an anime from the above list to watch it: ")

	var userChoice int
	fmt.Scanln(&userChoice)
	anime := aniList[userChoice-1]

	player.LaunchPlayerWindow(anime)
}
