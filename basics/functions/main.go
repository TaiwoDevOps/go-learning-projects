package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var dungeonName string
var totalRooms int
var roomTraps [5]bool
var enteredRoom int
var currentPlayer string

func init() {
	dungeonName = "The Cursed Depths"
	totalRooms = 5

	for i := range roomTraps {
		roomTraps[i] = rand.Intn(2) == 1
	}

	fmt.Println("Dungeons initialising.... traps armed")
	fmt.Printf("Dungeon: %s | Rooms: %d\n\n", dungeonName, totalRooms)
}

func validateEntry(playerLevel int, minLevel int) {
	if playerLevel == 0 {
		fmt.Println("Critical: player not initialised. Shutting down.")
		os.Exit(1)
	}
	if playerLevel < minLevel {
		fmt.Printf("Player level too low. Dungeon requires level {%d}. Exiting.\n", minLevel)
		os.Exit(1)
	}
	fmt.Printf("Player level {%d} accepted. Entering dungeon\n", playerLevel)

}

func welcomePlayer(name string, rooms int) string {
	return fmt.Sprintf("Welcome %s. Survive %d rooms in The Cursed Depths.", name, rooms)
}

func collectLoot(player string, items ...string) int {
	for _, item := range items {
		fmt.Printf("%s found: %v\n", player, item)
	}
	return len(items)
}

func totalScore(scores ...int) int {
	sum := 0
	for _, s := range scores {
		sum += s
	}

	return sum
}

func enterRoom(number int, name string) {
	enteredRoom = number
	currentPlayer = name
	defer fmt.Printf("Player %s exited room %d safely\n", name, number)
	fmt.Printf("Player %s just entered the room %d\n", name, enteredRoom)
	fmt.Printf("Player %s is fighting monster in room %d...\n", name, enteredRoom)
}

func trapRoom(number int, isTrapped bool) {
	defer fmt.Printf("Leaving trapped room %d\n", number)
	if isTrapped {
		panic(fmt.Sprintf("Room %d: floor collapsed!", number))
	}
	fmt.Printf("Room %d: no trap found, safe passage\n", number)
}

func enterTrapRoom(number int, isTrapped bool, lootBag *[]string, player string) bool {
	recovered := false

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Lucky charm activated! Recovered: {%s}\n", r)
			recovered = true
		}
	}()

	trapRoom(number, isTrapped)
	*lootBag = append(*lootBag, "TrapKey")
	fmt.Printf(" %s looted a TrapKey from room %d\n", player, number)
	return recovered
}

func quitDungeon(reason string) {
	defer fmt.Println("Dungeon quit handler firing")
	fmt.Printf("Player quit \"%s\"}\n", reason)
	fmt.Println("Thank you for playing Dungeon Depth. Goodbye.")
	os.Exit(0)
}

func printInventory(lootBag []string, weaponPower map[string]int) {
	fmt.Println("\n ------ Inventory ----- ")
	if len(lootBag) == 0 {
		fmt.Println("Loot bag is empty")
		return
	}

	score := 0
	for _, item := range lootBag {
		power, ok := weaponPower[item]
		if ok {
			fmt.Printf(" %s (power: %d)\n", item, power)
			score += power
		} else {
			fmt.Printf(" %s (power: unknown)\n", item)
		}
	}

	fmt.Printf(" Total power: %d\n", score)
}

func reportRoom(number int, cleared bool) {
	if cleared {
		fmt.Printf("Room %d cleared.\n", number)
	} else {
		fmt.Printf("Room %d failed.\n", number)
	}
}

func openChest(number int) {
	defer fmt.Printf("Closing chest in room %d\n", number)
	defer fmt.Printf("Logging chest access for room %d\n", number)
	fmt.Printf("Opening chest in room %d\n", number)
	if strings.TrimSpace(currentPlayer) != "" {
		fmt.Printf("%s takes the treasure!\n", currentPlayer)
		return
	}
	fmt.Println("Plauyer takes the treasure!")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	validateEntry(5, 3)

	fmt.Println(welcomePlayer("Taiwo", totalRooms))
	fmt.Println()

	lootBag := []string{}

	weaponPower := map[string]int{
		"FireSword":   40,
		"IceShield":   30,
		"PoisonArrow": 25,
		"TrapKey":     15,
	}
	roomLoot := [5]string{"FireSword", "IceShield", "PoisonArrow", "FireSword", "IceShield"}

	roomsCleared := 0
	roomScores := []int{}

	for i := 0; i < totalRooms; i++ {
		roomNumber := i + 1
		fmt.Printf("=============== ROOM %d ================\n", roomNumber)
		fmt.Println("Options: [enter] advance | [q] quit")
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "q" {
			quitDungeon("player chose to flee😒")
		}

		isTrapped := roomTraps[i]

		if isTrapped {
			fmt.Printf("Room %d is trapped\n", roomNumber)
			enterTrapRoom(roomNumber, true, &lootBag, "Taiwo")
		} else {
			enterRoom(roomNumber, "Taiwo")

			drop := roomLoot[i]
			fmt.Printf("\nRoom %d loot drop:\n", roomNumber)
			count := collectLoot("Taiwo", drop)
			if count > 0 {
				lootBag = append(lootBag, drop)
			}

			roomScores = append(roomScores, 20)
			roomsCleared++
		}

		printInventory(lootBag, weaponPower)
		fmt.Println()
	}

	fmt.Println("============ DUNGEON COMPLETE ==============")
	fmt.Printf("Rooms cleared: %d / %d\n", roomsCleared, totalRooms)

	fmt.Printf("Final score: %d\n", totalScore(roomScores...))

	fmt.Println("\nFinal loot bag:")
	for i, item := range lootBag {
		fmt.Printf(". [%d] %s\n", i, item)
	}

	fmt.Println("\nWeapon registry (all known weapons):")
	for weapon, power := range weaponPower {
		fmt.Printf(" %s -> power %d\n", weapon, power)
	}

}
