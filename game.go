package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Game struct {
	Width  int
	Height int
	Cells  []*Cell
	Gen    int
}

// Assumes:
// 0 1 2 3 4
// 1
// 2
// 3
// 4
func (g *Game) Init(width, height int) {
	g.Width = width
	g.Height = height
	g.Gen = 1
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			rand.Seed(time.Now().UnixNano())
			living := rand.Intn(2) == 0

			// 5/5 blinker
			// living := false
			// if (i == 2 && j == 1) || (i == 2 && j == 2) || (i == 2 && j == 3) {
			// 	living = true
			// }

			cell := &Cell{X: i, Y: j, Living: living}
			g.Cells = append(g.Cells, cell)
		}
	}

	g.PrintCells()
	for g.HasLivingCells() {
		time.Sleep(1000 * time.Millisecond)
		g.Generation()
		g.PrintCells()
	}
}

func (g *Game) HasLivingCells() bool {
	result := false
	for i := 0; i < len(g.Cells); i++ {
		if g.Cells[i].IsLiving() {
			result = true
			break
		}
	}

	return result
}

func (g *Game) PrintCells() {
	fmt.Printf("Generation %v\n", g.Gen)
	for i := 0; i < g.Width; i++ {
		for j := 0; j < g.Height; j++ {
			cell := g.FindCell(i, j)
			if cell.IsLiving() {
				fmt.Printf("O ")
			} else {
				fmt.Printf("X ")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

func (g *Game) Generation() {
	g.Gen++
	var wg sync.WaitGroup

	for i := 0; i < len(g.Cells); i++ {
		wg.Add(1)

		go func(j int, wg *sync.WaitGroup) {
			defer wg.Done()

			cell := g.Cells[j]
			cell.SetNextGenLiving(g)
		}(i, &wg)
	}

	wg.Wait()

	for i := 0; i < len(g.Cells); i++ {
		wg.Add(1)

		go func(j int, wg *sync.WaitGroup) {
			defer wg.Done()

			cell := g.Cells[j]
			cell.SetLiving()
		}(i, &wg)
	}

	wg.Wait()
}

func (g *Game) FindCell(x, y int) *Cell {
	cell := new(Cell)

	for i := 0; i < len(g.Cells); i++ {
		c := g.Cells[i]
		if x == c.X && y == c.Y {
			cell = c
			break
		}
	}

	return cell
}
