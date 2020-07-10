package deck

import (
	"aws-mahjong/tile"
	"errors"
	"math/rand"
	"time"
)

type Deck struct {
	tiles []tile.Tile
}

var (
	RunOutOfTileErr = errors.New("no tile exist on deck")
)

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

func (d *Deck) Draw() (*tile.Tile, error) {
	if len(d.tiles) > 0 {
		tile := d.tiles[0]
		d.tiles = d.tiles[1:]
		return &tile, nil
	}
	return nil, RunOutOfTileErr

}

func (d *Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.tiles), func(i, j int) { d.tiles[i], d.tiles[j] = d.tiles[j], d.tiles[i] })
}
