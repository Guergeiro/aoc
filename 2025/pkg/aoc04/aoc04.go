package aoc04

import (
	"slices"

	"github.com/guergeiro/aoc/2025/internal/grid"
)

func Part1(arr []string) int {
	papers := createPapers(arr)

	sum := 0
	for _, paper := range papers.GetData() {
		if paper.Value() == '@' {
			if canRemovePaper(paper) {
				sum += 1
			}
		}
	}
	return sum
}

func Part2(arr []string) int {
	papers := createPapers(arr)
	removedPapers := removePapers(papers)
	return len(removedPapers)
}

func createPapers(arr []string) grid.Grid[rune] {
	input := [][]rune{}
	for _, str := range arr {
		input = append(input, []rune(str))
	}
	return grid.NewGrid(input)
}

func removePapers(g grid.Grid[rune]) []*grid.Coordinate[rune] {
	papersToBeRemoved := []*grid.Coordinate[rune]{}
	for _, p := range g.GetData() {
		if p.Value() == '@' && canRemovePaper(p) {
			papersToBeRemoved = append(papersToBeRemoved, p)
		}
	}

	if len(papersToBeRemoved) == 0 {
		return papersToBeRemoved
	}

	for _, p := range papersToBeRemoved {
		removePaper(p)
	}

	return slices.Concat(papersToBeRemoved, removePapers(g))
}

func removePaper(p *grid.Coordinate[rune]) {
	p.SetValue('.')
}

func canRemovePaper(p *grid.Coordinate[rune]) bool {
	adjacentPapers := getAdjacentPapers(p)
	return len(adjacentPapers) < 4
}

func getAdjacentPapers(p *grid.Coordinate[rune]) []*grid.Coordinate[rune] {
	papers := []*grid.Coordinate[rune]{}
	if p.Top() != nil {
		if p.Top().Value() == '@' {
			papers = append(papers, p.Top())
		}
	}
	if p.TopRight() != nil {
		if p.TopRight().Value() == '@' {
			papers = append(papers, p.TopRight())
		}
	}
	if p.Right() != nil {
		if p.Right().Value() == '@' {
			papers = append(papers, p.Right())
		}
	}
	if p.BottomRight() != nil {
		if p.BottomRight().Value() == '@' {
			papers = append(papers, p.BottomRight())
		}
	}
	if p.Bottom() != nil {
		if p.Bottom().Value() == '@' {
			papers = append(papers, p.Bottom())
		}
	}
	if p.BottomLeft() != nil {
		if p.BottomLeft().Value() == '@' {
			papers = append(papers, p.BottomLeft())
		}
	}
	if p.Left() != nil {
		if p.Left().Value() == '@' {
			papers = append(papers, p.Left())
		}
	}
	if p.TopLeft() != nil {
		if p.TopLeft().Value() == '@' {
			papers = append(papers, p.TopLeft())
		}
	}
	return papers
}
