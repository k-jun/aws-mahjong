package deck

import (
	"aws-mahjong/tile"
	"errors"
	"math/rand"
	"time"
)

type Deck interface {
	Draw() (*tile.Tile, error)
	Count() int
}

type DeckImpl struct {
	tiles []tile.Tile
}

var (
	TimeNowUnix = time.Now().Unix()
)

var (
	RunOutOfTileErr = errors.New("no tile exist on deck")
)

func NewDeck() Deck {

	tiles := []tile.Tile{}
	for _, t := range tile.All {
		tiles = append(tiles, t)
	}

	deck := DeckImpl{tiles: tiles}
	deck.shuffle()
	return &deck
}

func (d *DeckImpl) Count() int {
	return len(d.tiles)
}

func (d *DeckImpl) Draw() (*tile.Tile, error) {
	if len(d.tiles) > 0 {
		tile := d.tiles[0]
		d.tiles = d.tiles[1:]
		return &tile, nil
	}
	return nil, RunOutOfTileErr
}

func (d *DeckImpl) shuffle() {
	rand.Seed(TimeNowUnix)
	rand.Shuffle(len(d.tiles), func(i, j int) { d.tiles[i], d.tiles[j] = d.tiles[j], d.tiles[i] })
}
