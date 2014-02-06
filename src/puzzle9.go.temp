package main

import (
	"Matrix"
	"Search"
	"fmt"
	"math/cmplx"
)

type Puzzle9 struct {
	M *Matrix.Matrix
}

func (this Puzzle9) Hp(Goal Search.Astar) float64 {
	result := 0.0
	if this.M.GetMRows() == Goal.(Puzzle9).M.GetMRows() && this.M.GetNColumns() == Goal.(Puzzle9).M.GetNColumns() {
		r := this.M.GetMRows()
		c := this.M.GetNColumns()

		for i := 1; i < r; i++ {
			for j := 1; j < c; j++ {
				if this.M.GetValue(i, j) != Goal.(Puzzle9).M.GetValue(i, j) {
					result = result + 1
				}
			}
		}
	}
	return result
}

func (this Puzzle9) H(x Search.Astar) float64 {
	return 1.0
}

func (this Puzzle9) Equal(OtherNode Search.Astar) bool {
	return Matrix.Equal(this.M, OtherNode.(Puzzle9).M)
}

func (this Puzzle9) Neighbors() []Search.Astar {
	Nmax := this.M.GetMRows() * this.M.GetNColumns() // pivote element
	//Search element
	Px := 0
	Py := 0
	flag := false
	for i := 1; i <= this.M.GetMRows(); i++ {
		for j := 1; j <= this.M.GetNColumns(); j++ {
			if cmplx.Abs(this.M.GetValue(i, j)) == float64(Nmax) {
				Px = i
				Py = j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	neighbors := []Puzzle9{}

	if Px-1 >= 1 {

		n := this.M.Copy()
		temp := n.GetValue(Px-1, Py)
		n.SetValue(Px-1, Py, n.GetValue(Px, Py))
		n.SetValue(Px, Py, temp)
		var neighbor Puzzle9

		neighbor.M = n
		neighbors = append(neighbors, neighbor)

	}
	if Px+1 <= this.M.GetNColumns() {
		n := this.M.Copy()
		temp := n.GetValue(Px+1, Py)
		n.SetValue(Px+1, Py, n.GetValue(Px, Py))
		n.SetValue(Px, Py, temp)
		var neighbor Puzzle9

		neighbor.M = n
		neighbors = append(neighbors, neighbor)
	}

	if Py-1 >= 1 {
		n := this.M.Copy()
		temp := n.GetValue(Px, Py-1)
		n.SetValue(Px, Py-1, n.GetValue(Px, Py))
		n.SetValue(Px, Py, temp)
		var neighbor Puzzle9

		neighbor.M = n
		neighbors = append(neighbors, neighbor)
	}
	if Py+1 <= this.M.GetMRows() {
		n := this.M.Copy()
		temp := n.GetValue(Px, Py+1)
		n.SetValue(Px, Py+1, n.GetValue(Px, Py))
		n.SetValue(Px, Py, temp)
		var neighbor Puzzle9

		neighbor.M = n
		neighbors = append(neighbors, neighbor)
	}

	s := make([]Search.Astar, len(neighbors))
	for i, v := range neighbors {
		s[i] = Search.Astar(v)
	}

	return s
}

func main() {

	M1, E1 := Matrix.FromFile("Search/test/position2.txt")
	Mg, E2 := Matrix.FromFile("Search/test/goal.txt")

	if E1 == nil && E2 == nil {
		var p1 Puzzle9
		p1.M = M1
		var goal Puzzle9
		goal.M = Mg

		path, er := Search.AStar(p1, goal, 1000)
		fmt.Println(er, " :/")
		if er == nil {
			fmt.Println(len(path))
			for i := 0; i < len(path); i++ {
				fmt.Println("--> ", path[i].(Puzzle9).M.ToString())
			}
		}

	}
}
