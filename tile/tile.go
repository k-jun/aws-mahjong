package tile

import (
	"log"
	"strconv"
)

type TileKind string

const Count = 4

var (
	suit      TileKind = "suit"      // suuzi
	dot       TileKind = "dot"       // pinzu
	banboo    TileKind = "banboo"    // sozu
	character TileKind = "character" // manzu
	one       TileKind = "1"
	two       TileKind = "2"
	three     TileKind = "3"
	four      TileKind = "4"
	five      TileKind = "5"
	six       TileKind = "6"
	seven     TileKind = "7"
	eight     TileKind = "8"
	nine      TileKind = "9"

	honor  TileKind = "honor"  // zihai
	wind   TileKind = "wind"   // kaze
	dragon TileKind = "dragon" // yakuhai
)

type Tile struct {
	kind []TileKind
	name string
}

func (t *Tile) Name() string {
	return t.name
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
	Dot1 = Tile{
		kind: []TileKind{suit, one, dot},
		name: "dot1",
	}
	Dot2 = Tile{
		kind: []TileKind{suit, two, dot},
		name: "dot2",
	}
	Dot3 = Tile{
		kind: []TileKind{suit, three, dot},
		name: "dot3",
	}
	Dot4 = Tile{
		kind: []TileKind{suit, four, dot},
		name: "dot4",
	}
	Dot5 = Tile{
		kind: []TileKind{suit, four, dot},
		name: "dot5",
	}
	Dot6 = Tile{
		kind: []TileKind{suit, six, dot},
		name: "dot6",
	}
	Dot7 = Tile{
		kind: []TileKind{suit, seven, dot},
		name: "dot7",
	}
	Dot8 = Tile{
		kind: []TileKind{suit, eight, dot},
		name: "dot8",
	}
	Dot9 = Tile{
		kind: []TileKind{suit, nine, dot},
		name: "dot9",
	}
	Banboo1 = Tile{
		kind: []TileKind{suit, one, dot},
		name: "banboo1",
	}
	Banboo2 = Tile{
		kind: []TileKind{suit, two, banboo},
		name: "banboo2",
	}
	Banboo3 = Tile{
		kind: []TileKind{suit, three, banboo},
		name: "banboo3",
	}
	Banboo5 = Tile{
		kind: []TileKind{suit, four, banboo},
		name: "banboo5",
	}
	Banboo4 = Tile{
		kind: []TileKind{suit, four, banboo},
		name: "banboo4",
	}
	Banboo6 = Tile{
		kind: []TileKind{suit, six, banboo},
		name: "banboo6",
	}
	Banboo7 = Tile{
		kind: []TileKind{suit, seven, banboo},
		name: "banboo7",
	}
	Banboo8 = Tile{
		kind: []TileKind{suit, eight, banboo},
		name: "banboo8",
	}
	Banboo9 = Tile{
		kind: []TileKind{suit, nine, banboo},
		name: "banboo9",
	}
	Character1 = Tile{
		kind: []TileKind{suit, one, character},
		name: "character1",
	}
	Character2 = Tile{
		kind: []TileKind{suit, two, character},
		name: "character2",
	}
	Character3 = Tile{
		kind: []TileKind{suit, three, character},
		name: "character3",
	}
	Character4 = Tile{
		kind: []TileKind{suit, four, character},
		name: "character4",
	}
	Character5 = Tile{
		kind: []TileKind{suit, four, character},
		name: "character5",
	}
	Character6 = Tile{
		kind: []TileKind{suit, six, character},
		name: "character6",
	}
	Character7 = Tile{
		kind: []TileKind{suit, seven, character},
		name: "character7",
	}
	Character8 = Tile{
		kind: []TileKind{suit, eight, character},
		name: "character8",
	}
	Character9 = Tile{
		kind: []TileKind{suit, nine, character},
		name: "character9",
	}

	// honors
	Red = Tile{
		kind: []TileKind{honor, dragon},
		name: "red",
	}
	Green = Tile{
		kind: []TileKind{honor, dragon},
		name: "green",
	}
	White = Tile{
		kind: []TileKind{honor, dragon},
		name: "white",
	}
	North = Tile{
		kind: []TileKind{honor, wind},
		name: "north",
	}
	East = Tile{
		kind: []TileKind{honor, wind},
		name: "east",
	}
	West = Tile{
		kind: []TileKind{honor, wind},
		name: "west",
	}
	South = Tile{
		kind: []TileKind{honor, wind},
		name: "south",
	}
)

var All = []Tile{
	Dot1, Dot2, Dot3, Dot4, Dot5, Dot6, Dot7, Dot8, Dot9,
	Banboo1, Banboo2, Banboo3, Banboo4, Banboo5, Banboo6, Banboo7, Banboo8, Banboo9,
	Character1, Character2, Character3, Character4, Character5, Character6, Character7, Character8, Character9,
	North, East, West, South,
	Red, Green, White,
}
