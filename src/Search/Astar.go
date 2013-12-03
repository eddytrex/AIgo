package Search

import (
	"errors"
)

type Astar interface {
	Hp(Goal Astar) float64

	H(x Astar) float64

	Equal(OtherNode Astar) bool

	Neighbors() []Astar
}

type NewNode struct {
	Node Astar

	Come_From *NewNode

	G float64
	F float64
}

func (this NewNode) Neighbors() []NewNode {
	result := []NewNode{}

	for _, n := range this.Node.Neighbors() {

		var a NewNode
		a.Node = n
		a.G = 0.0
		a.F = 0.0

		result = append(result, a)
	}
	return result
}

func AStar(Start, Goal Astar, n int) ([]Astar, error) {

	Open := make([]NewNode, 1)
	Open[0].Node = Start
	Open[0].G = 0
	Open[0].F = Open[0].G + Open[0].Node.Hp(Goal)

	var NGoal NewNode
	NGoal.Node = Goal

	Close := []NewNode{}

	i := 0

	for len(Open) > 0 {
		current, index := Minimun_F(Open, Goal)
		i++
		if current.Node.Equal(Goal) {
			return current.Recostruct_path(), nil
		}

		if i > n {
			break
		}

		Open = append(Open[:index], Open[index+1:]...)
		Close = append(Close, current)
		for _, n := range current.Neighbors() {
			tentative_g_score := current.G + current.Node.H(n.Node)
			tentative_f_score := tentative_g_score + n.Node.Hp(Goal)

			_, Is := n.In(Close)
			if Is && tentative_f_score >= n.F {
				continue

			}
			if !Is || tentative_f_score < n.F {

				n.Come_From = &current
				n.G = tentative_g_score
				n.F = tentative_f_score

				_, b := n.In(Open)
				if !b {
					Open = append(Open, n)
				}
			}

		}
	}

	return nil, errors.New(" the search has fail")
}

func (n NewNode) In(Set []NewNode) (*NewNode, bool) {
	for i := 1; i < len(Set); i++ {
		if Set[i].Node.Equal(n.Node) {
			return &Set[i], true
		}
	}
	return nil, false
}

func (current NewNode) Recostruct_path() (path []Astar) {
	result := make([]Astar, 1)

	if current.Come_From != nil {
		p := (*current.Come_From).Recostruct_path()
		return append(p, current.Node)
	}

	result[0] = current.Node
	return result
}

func Minimun_F(Open []NewNode, Goal Astar) (NewNode, int) {
	minimum := Open[0].Node.Hp(Goal) + Open[0].G
	minimum_Node := Open[0]
	index := 0

	for i := 1; i < len(Open); i++ {

		compare := Open[i].Node.Hp(Goal) + Open[i].G

		if compare < minimum {
			minimum = compare
			minimum_Node = Open[i]
			index = i
		}
	}
	return minimum_Node, index
}
