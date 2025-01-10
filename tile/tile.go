package tile

import (
	"dice-game/entities"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	// 1. 기본 타일 2. 이벤트 타일 3. 상점 타일 4. 중간 보스 타일 5. 최종 보스 타일 6. 저주 타일
	Type  int
	Image *ebiten.Image
	X     float64
	Y     float64
}

func shuffle(arr []Tile) []Tile {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

var (
	DefaultTile    Tile
	EventTile      Tile
	StoreTile      Tile
	MiddleBossTile Tile
	LastBossTile   Tile
	CursedTile     Tile
)

func init() {
	TileHash := loadTileImg()
	DefaultTile = Tile{
		Type:  1,
		Image: TileHash["DefaultTile"],
	}
	EventTile = Tile{
		Type:  2,
		Image: TileHash["EventTile"],
	}
	StoreTile = Tile{
		Type:  3,
		Image: TileHash["StoreTile"],
	}
	MiddleBossTile = Tile{
		Type:  4,
		Image: TileHash["MiddleBossTile"],
	}
	LastBossTile = Tile{
		Type:  5,
		Image: TileHash["LastBossTile"],
	}
	CursedTile = Tile{
		Type:  6,
		Image: TileHash["CursedTile"],
	}
}

func loadTileImg() map[string]*ebiten.Image {
	tiles := make(map[string]*ebiten.Image)

	err := filepath.Walk("assets/tiles", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() && (filepath.Ext(path) == ".png") {
			tileName := info.Name()[:len(info.Name())-4]
			img, _, err := ebitenutil.NewImageFromFile(path)
			if err != nil {
				log.Fatal(err)
			}

			tiles[tileName] = img
		}

		return nil
	})
	if err != nil {
		log.Fatal()
	}
	return tiles
}

func NewTileMap() [10][10]Tile {
	Tiles := [10][10]Tile{}

	for idx := range Tiles {
		Tiles[idx][0] = DefaultTile
		Tiles[idx][8] = DefaultTile
		Tiles[idx][9] = LastBossTile

		arr := shuffle([]Tile{
			DefaultTile,
			DefaultTile,
			DefaultTile,
			EventTile,
			EventTile,
			StoreTile,
			MiddleBossTile,
		})
		for i := 1; i < 8; i++ {
			Tiles[idx][i] = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		}
	}
	return Tiles
}

func NowTile(tileMap [10][10]Tile, player *entities.Player) Tile {
	return tileMap[player.NowIndex/10][player.NowIndex%10]
}
