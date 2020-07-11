package tile

import (
	"log"
	"strconv"
)

type TileKind string

var (
	suhai TileKind = "suhai"
	pinzu TileKind = "pinzu"
	souzu TileKind = "souzu"
	manzu TileKind = "manzu"
	one   TileKind = "1"
	two   TileKind = "2"
	three TileKind = "3"
	four  TileKind = "4"
	five  TileKind = "5"
	aka   TileKind = "aka"
	six   TileKind = "6"
	seven TileKind = "7"
	eight TileKind = "8"
	nine  TileKind = "9"

	zihai   TileKind = "zihai"
	kaze    TileKind = "kaze"
	yakuhai TileKind = "yakuhai"
)

type Tile struct {
	kind []TileKind
	name string
}

func (t *Tile) Name() string {
	return t.name
}

func (t *Tile) IsSame(a *Tile) bool {
	if t.Name() == a.Name() {
		return true
	}

	akaMapper := map[string]string{
		Souzu5.Name():    Souzu5Aka.Name(),
		Manzu5.Name():    Manzu5Aka.Name(),
		Pinzu5.Name():    Pinzu5Aka.Name(),
		Souzu5Aka.Name(): Souzu5.Name(),
		Manzu5Aka.Name(): Manzu5.Name(),
		Pinzu5Aka.Name(): Pinzu5.Name(),
	}

	if t.Number() == 5 && akaMapper[t.Name()] == a.Name() {
		return true
	}

	return false

}

func (t *Tile) Number() int {

	for _, k := range t.kind {
		for _, knum := range []TileKind{one, two, three, four, five, six, seven, eight, nine} {
			if knum == k {
				num, err := strconv.Atoi(string(knum))
				if err != nil {
					log.Println(err)
					return 0
				}
				return num
			}
		}
	}
	return 0
}

func (t *Tile) Kinds() []TileKind {
	return t.kind
}

var (
	// suits
	Pinzu1 = Tile{
		kind: []TileKind{suhai, one, pinzu},
		name: "pinzu1",
	}
	Pinzu2 = Tile{
		kind: []TileKind{suhai, two, pinzu},
		name: "pinzu2",
	}
	Pinzu3 = Tile{
		kind: []TileKind{suhai, three, pinzu},
		name: "pinzu3",
	}
	Pinzu4 = Tile{
		kind: []TileKind{suhai, four, pinzu},
		name: "pinzu4",
	}
	Pinzu5 = Tile{
		kind: []TileKind{suhai, five, pinzu},
		name: "pinzu5",
	}
	Pinzu5Aka = Tile{
		kind: []TileKind{suhai, five, pinzu, aka},
		name: "pinzu5aka",
	}
	Pinzu6 = Tile{
		kind: []TileKind{suhai, six, pinzu},
		name: "pinzu6",
	}
	Pinzu7 = Tile{
		kind: []TileKind{suhai, seven, pinzu},
		name: "pinzu7",
	}
	Pinzu8 = Tile{
		kind: []TileKind{suhai, eight, pinzu},
		name: "pinzu8",
	}
	Pinzu9 = Tile{
		kind: []TileKind{suhai, nine, pinzu},
		name: "pinzu9",
	}
	Souzu1 = Tile{
		kind: []TileKind{suhai, one, pinzu},
		name: "souzu1",
	}
	Souzu2 = Tile{
		kind: []TileKind{suhai, two, souzu},
		name: "souzu2",
	}
	Souzu3 = Tile{
		kind: []TileKind{suhai, three, souzu},
		name: "souzu3",
	}
	Souzu4 = Tile{
		kind: []TileKind{suhai, four, souzu},
		name: "souzu4",
	}
	Souzu5 = Tile{
		kind: []TileKind{suhai, five, souzu},
		name: "souzu5",
	}
	Souzu5Aka = Tile{
		kind: []TileKind{suhai, five, souzu, aka},
		name: "souzu5aka",
	}
	Souzu6 = Tile{
		kind: []TileKind{suhai, six, souzu},
		name: "souzu6",
	}
	Souzu7 = Tile{
		kind: []TileKind{suhai, seven, souzu},
		name: "souzu7",
	}
	Souzu8 = Tile{
		kind: []TileKind{suhai, eight, souzu},
		name: "souzu8",
	}
	Souzu9 = Tile{
		kind: []TileKind{suhai, nine, souzu},
		name: "souzu9",
	}
	Manzu1 = Tile{
		kind: []TileKind{suhai, one, manzu},
		name: "manzu1",
	}
	Manzu2 = Tile{
		kind: []TileKind{suhai, two, manzu},
		name: "manzu2",
	}
	Manzu3 = Tile{
		kind: []TileKind{suhai, three, manzu},
		name: "manzu3",
	}
	Manzu4 = Tile{
		kind: []TileKind{suhai, four, manzu},
		name: "manzu4",
	}
	Manzu5 = Tile{
		kind: []TileKind{suhai, five, manzu},
		name: "manzu5",
	}
	Manzu5Aka = Tile{
		kind: []TileKind{suhai, five, souzu, aka},
		name: "manzu5aka",
	}
	Manzu6 = Tile{
		kind: []TileKind{suhai, six, manzu},
		name: "manzu6",
	}
	Manzu7 = Tile{
		kind: []TileKind{suhai, seven, manzu},
		name: "manzu7",
	}
	Manzu8 = Tile{
		kind: []TileKind{suhai, eight, manzu},
		name: "manzu8",
	}
	Manzu9 = Tile{
		kind: []TileKind{suhai, nine, manzu},
		name: "manzu9",
	}

	// honors
	Chun = Tile{
		kind: []TileKind{zihai, yakuhai},
		name: "red",
	}
	Hatu = Tile{
		kind: []TileKind{zihai, yakuhai},
		name: "green",
	}
	Haku = Tile{
		kind: []TileKind{zihai, yakuhai},
		name: "white",
	}
	North = Tile{
		kind: []TileKind{zihai, kaze},
		name: "north",
	}
	East = Tile{
		kind: []TileKind{zihai, kaze},
		name: "east",
	}
	West = Tile{
		kind: []TileKind{zihai, kaze},
		name: "west",
	}
	South = Tile{
		kind: []TileKind{zihai, kaze},
		name: "south",
	}
)

var AllTailKind = []Tile{
	Pinzu1, Pinzu2, Pinzu3, Pinzu4, Pinzu5, Pinzu5Aka, Pinzu6, Pinzu7, Pinzu8, Pinzu9,
	Souzu1, Souzu2, Souzu3, Souzu4, Souzu5, Souzu5Aka, Souzu6, Souzu7, Souzu8, Souzu9,
	Manzu1, Manzu2, Manzu3, Manzu4, Manzu5, Manzu5Aka, Manzu6, Manzu7, Manzu8, Manzu9,
	North, East, West, South,
	Chun, Hatu, Haku,
}

var All = []Tile{
	// pinzu
	Pinzu1, Pinzu2, Pinzu3, Pinzu4, Pinzu5, Pinzu6, Pinzu7, Pinzu8, Pinzu9,
	Pinzu1, Pinzu2, Pinzu3, Pinzu4, Pinzu5, Pinzu6, Pinzu7, Pinzu8, Pinzu9,
	Pinzu1, Pinzu2, Pinzu3, Pinzu4, Pinzu5, Pinzu6, Pinzu7, Pinzu8, Pinzu9,
	Pinzu1, Pinzu2, Pinzu3, Pinzu4, Pinzu5Aka, Pinzu6, Pinzu7, Pinzu8, Pinzu9,

	// souzu
	Souzu1, Souzu2, Souzu3, Souzu4, Souzu5, Souzu6, Souzu7, Souzu8, Souzu9,
	Souzu1, Souzu2, Souzu3, Souzu4, Souzu5, Souzu6, Souzu7, Souzu8, Souzu9,
	Souzu1, Souzu2, Souzu3, Souzu4, Souzu5, Souzu6, Souzu7, Souzu8, Souzu9,
	Souzu1, Souzu2, Souzu3, Souzu4, Souzu5Aka, Souzu6, Souzu7, Souzu8, Souzu9,

	// manzu
	Manzu1, Manzu2, Manzu3, Manzu4, Manzu5, Manzu6, Manzu7, Manzu8, Manzu9,
	Manzu1, Manzu2, Manzu3, Manzu4, Manzu5, Manzu6, Manzu7, Manzu8, Manzu9,
	Manzu1, Manzu2, Manzu3, Manzu4, Manzu5, Manzu6, Manzu7, Manzu8, Manzu9,
	Manzu1, Manzu2, Manzu3, Manzu4, Manzu5Aka, Manzu6, Manzu7, Manzu8, Manzu9,

	// kaze
	North, East, West, South,
	North, East, West, South,
	North, East, West, South,
	North, East, West, South,

	// yakuhai
	Chun, Hatu, Haku,
	Chun, Hatu, Haku,
	Chun, Hatu, Haku,
	Chun, Hatu, Haku,
}
