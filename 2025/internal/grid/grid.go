package grid

func NewGrid[T any](input [][]T) Grid[T] {
	numRows := len(input)
	numCols := len(input[0])

	grid := Grid[T]{
		data:    []*Coordinate[T]{},
		numCols: numCols,
	}

	// Create all points in 1D array
	for row := 0; row < numRows; row += 1 {
		for col := 0; col < numCols; col += 1 {
			point := &Coordinate[T]{
				value: input[row][col],
			}
			grid.data = append(grid.data, point)
		}
	}

	// Connect points via pointers using 1D indexing
	for row := 0; row < numRows; row += 1 {
		for col := 0; col < numCols; col += 1 {
			idx := grid.idx(row, col)
			if row > 0 {
				grid.data[idx].top = grid.data[grid.idx(row-1, col)]
			}
			if row < numRows-1 {
				grid.data[idx].bottom = grid.data[grid.idx(row+1, col)]
			}
			if col > 0 {
				grid.data[idx].left = grid.data[grid.idx(row, col-1)]
			}
			if col < numCols-1 {
				grid.data[idx].right = grid.data[grid.idx(row, col+1)]
			}
		}
	}

	return grid
}

type Grid[T any] struct {
	data    []*Coordinate[T]
	numCols int
}

func (g Grid[T]) idx(row, col int) int {
	return row*g.numCols + col
}

func (g Grid[T]) GetData() []*Coordinate[T] {
	return g.data
}

func (g Grid[T]) GetMap() [][]*Coordinate[T] {
	numRows := len(g.data) / g.numCols
	result := [][]*Coordinate[T]{}

	for row := 0; row < numRows; row += 1 {
		rowData := []*Coordinate[T]{}
		for col := 0; col < g.numCols; col += 1 {
			idx := row*g.numCols + col
			rowData = append(rowData, g.data[idx])
		}
		result = append(result, rowData)
	}

	return result
}

type Coordinate[T any] struct {
	value  T
	top    *Coordinate[T]
	right  *Coordinate[T]
	bottom *Coordinate[T]
	left   *Coordinate[T]
}

func (p *Coordinate[T]) Value() T {
	return p.value
}

func (p *Coordinate[T]) SetValue(value T) *Coordinate[T] {
	p.value = value
	return p
}

func (p *Coordinate[T]) Top() *Coordinate[T] {
	return p.top
}

func (p *Coordinate[T]) Bottom() *Coordinate[T] {
	return p.bottom
}

func (p *Coordinate[T]) Right() *Coordinate[T] {
	return p.right
}

func (p *Coordinate[T]) Left() *Coordinate[T] {
	return p.left
}

func (p *Coordinate[T]) TopRight() *Coordinate[T] {
	if p.top != nil {
		return p.top.right
	}
	return nil
}

func (p *Coordinate[T]) BottomRight() *Coordinate[T] {
	if p.bottom != nil {
		return p.bottom.right
	}
	return nil
}

func (p *Coordinate[T]) BottomLeft() *Coordinate[T] {
	if p.bottom != nil {
		return p.bottom.left
	}
	return nil
}

func (p *Coordinate[T]) TopLeft() *Coordinate[T] {
	if p.top != nil {
		return p.top.left
	}
	return nil
}
