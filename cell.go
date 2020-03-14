package main

type Cell struct {
	Living        bool
	NextGenLiving bool
	X             int
	Y             int
}

func (c *Cell) IsLiving() bool {
	return c.Living == true
}

func (c *Cell) SetLiving() {
	c.Living = c.NextGenLiving
}

func (c *Cell) SetNextGenLiving(game *Game) {
	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	// Any live cell with two or three live neighbours lives on to the next generation.
	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

	count := c.LivingNeighborsCount(game)
	if c.IsLiving() {
		if count < 2 {
			c.NextGenLiving = false
		} else if count <= 3 {
			c.NextGenLiving = true
		} else {
			c.NextGenLiving = false
		}
	} else {
		if count == 3 {
			c.NextGenLiving = true
		} else {
			c.NextGenLiving = false
		}
	}
}

func (c *Cell) Neighbors(game *Game) []Cell {
	// X and Y adjustments
	adjustments := [][]int{
		[]int{-1, -1}, []int{0, -1}, []int{1, -1}, []int{-1, 0}, []int{1, 0}, []int{-1, 1}, []int{0, 1}, []int{1, 1},
	}
	var cells []Cell
	for i := 0; i < len(adjustments); i++ {
		adjustment := adjustments[i]
		x := c.X + adjustment[0]
		y := c.Y + adjustment[1]
		cell := game.FindCell(x, y)

		if cell != nil {
			cells = append(cells, *cell)
		}
	}

	return cells
}

func (c *Cell) LivingNeighborsCount(game *Game) int {
	var count int
	neighbors := c.Neighbors(game)

	for i := 0; i < len(neighbors); i++ {
		if neighbors[i].IsLiving() {
			count++
		}
	}

	return count
}
