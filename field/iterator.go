package field

import "oop/unit"

type GameFieldIterator struct {
	Field *GameField
	Index int
}

func (it *GameFieldIterator) Next() bool {
	it.Index++
	return it.Index < len(it.Field.units)
}

func (it *GameFieldIterator) Value() unit.Unit {
	if it.Index >= 0 && it.Index < len(it.Field.units) {
		return it.Field.units[it.Index]
	}

	return nil
}
