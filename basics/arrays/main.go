package main

import (
	"fmt"
)

func main() {
	// array()
	// sliceFunc()
	game()
}

func array() {
	var numbers [5]int
	fmt.Println("the numbers are:", numbers)

	numbers[0] = 43
	numbers[4] = 24

	fmt.Println("the numbers array is now:", numbers)

	names := [5]string{"Taiwo", "Joshua", "Olumide", "Adisa", "Funmi"}

	fmt.Println("the names of me and my fam:", names)

	names[1] = "Kenny"
	names[2] = "Idowu"
	names[3] = "Jesse"

	fmt.Println("the updated list is:", names)

	if len(names) == len(numbers) {
		fmt.Println("The length is the same:", len(numbers))
	}

	for i, v := range names {
		fmt.Println("the name at index I :", i, "with value of:", v)
	}

	matrix := [3][3]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("the array is", matrix)
}

func sliceFunc() {
	var numbers = make([]int, 5)
	fmt.Println("the numbers are:", numbers)

	for i := range numbers {
		numbers[i] = 2*i + 2
	}

	fmt.Println("the numbers slices is now:", numbers)

	names := []string{"Taiwo", "Joshua", "Olumide", "Adisa", "Funmi"}

	fmt.Println("the names of me and my fam:", names)

	names[1] = "Kenny"
	names[2] = "Idowu"
	names[3] = "Jesse"
	names = append(names, "Ileri")

	fmt.Println("the updated list is:", names)

	if len(names) == len(numbers) {
		fmt.Println("The length is the same:", len(numbers))
	}

	for i, v := range names {
		fmt.Println("the name at index I :", i, "with value of:", v)
	}

	matrix := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("the slices is", matrix)
}

func game() {
	var party [4]string

	party[0] = "Taiwo"
	party[1] = "Collections"
	party[2] = "Claude"
	party[3] = "Go"

	fmt.Println("Your party:", party)
	fmt.Println("Party size:", len(party))

	var lootbag []string

	lootbag = append(lootbag, "FireSword")
	lootbag = append(lootbag, "IceShield")
	lootbag = append(lootbag, "PoisonArrow")

	fmt.Println("Loot bag:", lootbag)
	fmt.Println("Items carried:", len(lootbag))
	fmt.Println("Bag capacity:", cap(lootbag))

	lootbag = append(lootbag[:0], lootbag[1:]...)
	fmt.Println("After using item:", lootbag)

	weakness := map[string]string{
		"Dragon":   "IceShield",
		"Skeleton": "FireSword",
		"Troll":    "PoisonArrow",
		"Ghost":    "FireSword",
	}

	fmt.Println("Troll fears:", weakness["Troll"])

	item, ok := weakness["Dragon"]
	if !ok {
		fmt.Println("Unknown monster - no weakness data")
	} else {
		fmt.Println("Weakness found", item)
	}

	delete(weakness, "Skeleton")

	fmt.Println("Monsters remaining:", len(weakness))

	for i, hero := range party {
		fmt.Printf("Here %d: %s reporting for duty\n", i, hero)
	}

	monster := "dddd"
	requiredItem, ok := weakness[monster]

	if !ok {
		fmt.Println("Monster not registered — no weakness data.")
		return
	}
	itemFound := false

	for _, item := range lootbag {
		if item == requiredItem {
			itemFound = true
			break
		}
	}

	if itemFound {
		fmt.Printf("Used %s against %s - monster defeated!\n", requiredItem, monster)
		for i, item := range lootbag {
			if item == requiredItem {
				lootbag = append(lootbag[:i], lootbag[i+1:]...)
				break
			}
		}
	} else {
		fmt.Println("You don't have the right item - monster escapes!")
	}

	fmt.Println("\n---- End of dungeon run ---")
	fmt.Printf("Heroes: %d\n", len(party))
	fmt.Printf("Remaining loot (%d items):\n", len(lootbag))

	for _, item := range lootbag {
		fmt.Println("This weapon is not used -", item)
	}

}
