package game

import (
	"fmt"
	"strings"
)

var towers = [3]rune{'a', 'b', 'c'}

var Movements = [6][2]rune{{'a', 'b'}, {'a', 'c'}, {'b', 'a'}, {'b', 'c'}, {'c', 'a'}, {'c', 'b'}}

type State struct {
	a Tower
	b Tower
	c Tower
}

func BlankState() *State {
	return NewState(Tower{}, Tower{}, Tower{})
}

func NewState(a Tower, b Tower, c Tower) *State {
	return &State{a, b, c}
}

func ValidTowerName(name rune) bool {
	return name == 'a' || name == 'b' || name == 'c'
}

func RemainingTower(from, to rune) (rune, error) {
	var remaining rune
	if !ValidTowerName(from) || !ValidTowerName(to) {
		return remaining, fmt.Errorf("from '%c' or to '%c' is not a valid tower", from, to)
	}
	for _, v := range towers {
		if !(v == from || v == to) {
			remaining = v
		}
	}
	return remaining, nil
}

func (state *State) Eq(other State) bool {
	return state.a.Eq(other.a) && state.b.Eq(other.b) && state.c.Eq(other.c)
}

func (state *State) IsValid() bool {
	return state.a.IsValid() && state.b.IsValid() && state.c.IsValid()
}

func (state *State) getTower(tower rune) (*Tower, error) {
	switch tower {
	case 'a':
		return &state.a, nil
	case 'b':
		return &state.b, nil
	case 'c':
		return &state.c, nil
	}
	return nil, fmt.Errorf("unknown tower")
}

func (state *State) Move(fromName, toName rune) (*State, error) {
	from, err := state.getTower(fromName)
	if err != nil {
		return nil, err
	}
	to, err := state.getTower(toName)
	if err != nil {
		return nil, err
	}
	remaining, err := RemainingTower(fromName, toName)
	if err != nil {
		return nil, err
	}

	fromClone := from.Clone()
	disk, ok := fromClone.Pop()
	if !ok {
		return nil, fmt.Errorf("empty start tower")
	}
	toClone := to.Clone()
	toClone.Push(disk)
	newState := &State{}
	if err := newState.SetTower(fromName, *fromClone); err != nil {
		return nil, err
	}
	if err := newState.SetTower(toName, *toClone); err != nil {
		return nil, err
	}
	if remainingTower, err := state.getTower(remaining); err != nil {
		return nil, err
	} else if err := newState.SetTower(remaining, *remainingTower); err != nil {
		return nil, err
	}
	if !newState.IsValid() {
		return nil, fmt.Errorf("invalid movent")
	}
	return newState, nil
}

func (state *State) SetTower(name rune, value Tower) error {
	switch name {
	case 'a':
		state.a = value
		return nil
	case 'b':
		state.b = value
		return nil
	case 'c':
		state.c = value
		return nil
	}
	return fmt.Errorf("%c is a invalid tower", name)
}

func (state State) String() string {
	return strings.ReplaceAll(fmt.Sprintf("(%v, %v, %v)", state.a, state.b, state.c), "&", "")
}

func (state State) ApplyRules() []*State {
	nextStates := []*State{}
	for _, movement := range Movements {
		if _state, err := state.Move(movement[0], movement[1]); err == nil {
			nextStates = append(nextStates, _state)
		}
	}
	return nextStates
}

func ExpectedDirection(disk, numberDisks int) int {
	if disk%2 == 0 && numberDisks%2 == 0 {
		return -1 // left
	} else if disk%2 != 0 && numberDisks%2 == 0 {
		return 1 // right
	} else if disk%2 == 0 && numberDisks%2 != 0 {
		return 1 // right
	} else if disk%2 != 0 && numberDisks%2 != 0 {
		return -1 // left
	}
	return 0
}

func Movement(start, result State) (int, [2]rune, int, error) { // disco, movimiento, peso, error
	for _, movement := range Movements {
		if state, err := start.Move(movement[0], movement[1]); err == nil {
			if state.Eq(result) {
				tower, _ := start.getTower(movement[0])
				disk := tower.GetDisc()
				weight, _ := MovementWeight(movement)
				return disk, movement, weight, nil
			}
		}
	}
	return 0, [2]rune{}, 0, fmt.Errorf("no movement was found that of that result")
}

func MovementWeight(movement [2]rune) (int, error) {
	s, err := towerPos(movement[0])
	if err != nil {
		return 0, err
	}
	e, err := towerPos(movement[1])
	if err != nil {
		return 0, err
	}
	w := e - s
	if w == 2 {
		w = -1
	} else if w == -2 {
		w = 1
	}
	return w, nil
}

func towerPos(pos rune) (int, error) {
	switch pos {
	case 'a':
		return 0, nil
	case 'b':
		return 1, nil
	case 'c':
		return 2, nil
	default:
		return 0, fmt.Errorf("invalid movemnt")
	}
}
