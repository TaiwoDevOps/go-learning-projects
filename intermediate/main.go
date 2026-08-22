package main

import (
	"fmt"
	"strings"
)

// Generics
type Bag[T any] struct {
	items []T
}

type HasHP interface {
	GetHP() int
}

func (b *Bag[T]) Add(item T) {
	b.items = append(b.items, item)
}

func (b *Bag[T]) Remove(index int) (T, bool) {

	var t T
	if index < 0 || index > len(b.items)-1 {
		return t, false
	}

	t = b.items[index]
	temp := b.items[len(b.items)-1]
	b.items[index] = temp

	var empty T
	b.items[len(b.items)-1] = empty

	b.items = b.items[:len(b.items)-1]

	return t, true

}

func (b *Bag[T]) List() []T {
	return b.items
}

// Generic with constraints
func (e *Entity) GetHP() int {
	return e.HP
}

func Strongest[T HasHP](fighters []T) T {
	var highestHP T
	if len(fighters) <= 0 {
		return highestHP
	}

	highestHP = fighters[0]
	for _, v := range fighters {
		if v.GetHP() > highestHP.GetHP() {
			highestHP = v
		}
	}

	return highestHP
}

type Entity struct {
	Name string
	HP   int
}

type Combatant interface {
	TakeDamage(amount int)
	IsAlive() bool
}

type Interactable interface {
	Enter() string
	Describe() string
}

// composition
type Room struct {
	Name        string
	Description string
	Connections []*Room
	HasEnemy    bool
	HasTreasure bool
}

type PuzzleRoom struct {
	Name       string
	CipherText string
}

type TrapRoom struct {
	Name   string
	Damage int
}

type TreasureRoom struct {
	Name string
	Gold int
}

type RuneRoom struct {
	Name      string
	CiherText string
	Shift     int
	Solved    bool
}

// struct embedding
type Player struct {
	Entity
	MaxHp int
	Inventory
}

type Enemy struct {
	Entity
}

type Inventory struct {
	Gold int
	Item []string
}

func (p *Entity) TakeDamage(amount int) {
	p.HP = p.HP - amount
}
func (p *Entity) IsAlive() bool {
	return p.HP > 0
}

func (e *Entity) Describe() string {
	return fmt.Sprintf("%s (%d HP)", e.Name, e.HP)
}

func (p *Player) Heal(amount int) {
	if p.HP < p.MaxHp && (p.HP+amount) <= p.MaxHp {
		p.HP = p.HP + amount
	} else {
		p.HP = p.MaxHp
	}
}

func (p *Player) Describe() string {
	return fmt.Sprintf("%s health power is %d", p.Name, p.HP)
}

func Fight(p, e Combatant) {
	p.TakeDamage(15)
	if p.IsAlive() {
		fmt.Println("Kael is still alive and just fought back")
		e.TakeDamage(10)
	}
}

// interface
func (p PuzzleRoom) Enter() string {
	return fmt.Sprintf("You just entered into: %s \n", p.Describe())
}

func (p PuzzleRoom) Describe() string {
	return fmt.Sprintf("%s", p.Name)
}

func (p TrapRoom) Enter() string {
	return fmt.Sprintf("You just entered into: %s ", p.Describe())
}

func (p TrapRoom) Describe() string {
	return fmt.Sprintf("%s is a trap room and you might suffer %d%% damage", p.Name, p.Damage)
}

func (p TreasureRoom) Enter() string {
	return fmt.Sprintf("You just entered into: %s ", p.Describe())
}

func (p TreasureRoom) Describe() string {
	return fmt.Sprintf("Welcome to %s, you can walk away with %d gold bars", p.Name, p.Gold)
}

func EnterRoom(r Interactable) {
	fmt.Println(r.Describe())
	fmt.Println(r.Enter())
}

type Spell func(target string) int

// closures
func makeSpell(name string, baseDamage int) Spell {

	timesCast := 0

	return func(target string) int {
		timesCast++
		damage := baseDamage + (timesCast * 2)
		fmt.Printf("Kael casts %s on %s for %d damage! (cast #%d)\n", name, target, damage, timesCast)
		return damage
	}
}

func buildFloor() *Room {
	armoury := &Room{
		Name:        "Armoury",
		HasTreasure: true,
	}

	guardRoom := &Room{
		Name:     "Guard Room",
		HasEnemy: true,
		Connections: []*Room{
			armoury,
		},
	}

	trapRoom := &Room{
		Name: "Trap Room",
	}

	bossChamber := &Room{
		Name:     "Boss Chamber",
		HasEnemy: true,
	}

	shadowHall := &Room{
		Name: "Shadow Hall",
		Connections: []*Room{
			trapRoom, bossChamber,
		},
	}

	entrance := &Room{
		Name: "Entrance",
		Connections: []*Room{
			guardRoom,
			shadowHall,
		},
	}

	return entrance
}

func exploreFloor(room *Room, depth int) {
	if room == nil {
		return
	}

	indent := strings.Repeat(" ", depth)
	fmt.Printf("%s[ 🚪 %s]\n", indent, room.Name)

	if room.HasEnemy {
		fmt.Printf("%s x enemy lurks here\n", indent)
	}

	if room.HasTreasure {
		fmt.Printf("%s + treasure found\n", indent)
	}

	for _, next := range room.Connections {
		exploreFloor(next, depth+1)
	}
}

// recurssion
func countRooms(room *Room) int {
	if room == nil {
		return 0
	}

	total := 1
	for _, next := range room.Connections {
		total += countRooms(next)
	}

	return total
}

func countEnemies(room *Room) int {
	if room == nil {
		return 0
	}

	totalEnemies := 0

	if room.HasEnemy {
		totalEnemies++
	}
	for _, next := range room.Connections {
		totalEnemies += countEnemies(next)
	}
	return totalEnemies
}

func findPath(current *Room, target *Room, visited map[*Room]bool, path []string) ([]string, bool) {

	if current == nil || target == nil {
		return path, false
	}

	if visited[current] {
		return path, false
	}

	visited[current] = true
	path = append(path, current.Name)

	if current == target {
		return path, true
	}

	for _, next := range current.Connections {
		if result, found := findPath(next, target, visited, path); found {
			return result, true
		}
	}

	visited[current] = false
	path = path[:len(path)-1]

	return path, false

}

func decodePuzzle(cipher string, shift int) string {
	convertCipher := []rune(cipher)
	for i, r := range convertCipher {
		if r >= 'a' && r <= 'z' {
			convertCipher[i] = r - rune(shift)
		} else if r >= 'A' && r <= 'Z' {
			convertCipher[i] = r - rune(shift)
		} else {
			convertCipher[i] = r
		}
	}
	return string(convertCipher)
}

func formatRoomName(name string) string {
	runes := []rune(name)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - 32
	}
	return string(runes)
}

func parseCommand(input string) (verb, arg string) {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

func main() {

	floor := buildFloor()
	fmt.Println("==== Exploring Floor 1 ====")
	exploreFloor(floor, 0)
	fmt.Println("==== Dungeon Census ====")
	fmt.Printf("Total rooms: %d\n", countRooms(floor))
	fmt.Printf("Total enemies: %d\n", countEnemies(floor))
	result, found := findPath(floor, floor.Connections[1], make(map[*Room]bool), []string{})
	fmt.Printf("Targeted path to Boss Chamber: %v and found: %v\n", result, found)

	kael := &Player{
		Entity: Entity{
			Name: "Kael",
			HP:   60,
		},
		MaxHp: 1000,
		Inventory: Inventory{
			Gold: 1,
			Item: []string{"Spell Book"},
		},
	}

	kael.TakeDamage(10)

	fmt.Println(kael.Describe())

	kael.Heal(1000)

	fmt.Println(kael.Describe())

	puzzle := RuneRoom{
		Name:      "shadow hall",
		CiherText: "Wkh Ervv Dzdlwv Ehklqg Wklv Grru",
		Shift:     3,
	}

	fmt.Println(formatRoomName(puzzle.Name))

	verb, args := parseCommand("decode 3")
	fmt.Println(verb, args)

	if verb == "decode" {
		decoded := decodePuzzle(puzzle.CiherText, puzzle.Shift)
		fmt.Println(decoded)
		puzzle.Solved = true
	}

	fmt.Printf("The puzzle %v\n", puzzle)

	player1 := Player{
		Entity: Entity{
			Name: "Jemmy",
			HP:   100,
		},
		MaxHp: 1000,
		Inventory: Inventory{
			Gold: 1,
			Item: []string{"Spell Book"},
		},
	}

	fmt.Printf("%s inventory item %v \n", player1.Name, player1.Inventory.Item)

	someone := Player{}

	fmt.Printf("Health Power: %d and Gold count: %d \n", someone.HP, someone.Inventory.Gold)

	Fight(kael, &Enemy{
		Entity: Entity{
			Name: "Goblin",
			HP:   100,
		},
	})
	EnterRoom(PuzzleRoom{
		Name:       "shadow hall",
		CipherText: "Wkh Ervv Dzdlwv Ehklqg Wklv Grru",
	})

	EnterRoom(TrapRoom{
		Name:   "Guards quarters",
		Damage: 10,
	})
	EnterRoom(TreasureRoom{
		Name: "Treasury",
		Gold: 10,
	})

	lootBag := Bag[string]{
		items: []string{"weapon 1", "weapon 2", "weapon 3"},
	}
	fmt.Println(lootBag.List())
	fmt.Println(lootBag.Remove(3))
	lootBag.Add("Weapon 4")
	fmt.Println(lootBag.List())
	fmt.Println(lootBag.Remove(1))
	fmt.Println(lootBag.List())

	enemies := []*Enemy{
		{
			Entity: Entity{
				Name: "Draculla",
				HP:   10,
			},
		},
		{
			Entity: Entity{
				Name: "Peter Obi",
				HP:   60,
			},
		},
		{
			Entity: Entity{
				Name: "Tinubu",
				HP:   100,
			},
		},
	}

	res := Strongest(enemies)

	fmt.Println(res.GetHP())

}
