package deck

import (
	"aws-mahjong/tile"
	"math/rand"
	"time"
)

type Deck struct {
	tiles []tile.Tile
}

func NewDeck() *Deck {

	tiles := []tile.Tile{}
	for _, t := range tile.All {
		for i := 0; i < tile.Count; i++ {
			tiles = append(tiles, t)
		}
	}

	deck := Deck{tiles: tiles}
	deck.shuffle()
	return &deck
}

func (d *Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.tiles), func(i, j int) { d.tiles[i], d.tiles[j] = d.tiles[j], d.tiles[i] })
}
