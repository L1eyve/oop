package field

import (
	"errors"
	"fmt"
	"oop/unit"
)

type GameField struct {
	Width, Height, MaxUnits int
	units                   []unit.Unit
}

func NewGameField(width, height, maxUnits int) *GameField {
	return &GameField{
		Width:    width,
		Height:   height,
		MaxUnits: maxUnits,
		units:    []unit.Unit{},
	}
}

func (f *GameField) AddUnit(u unit.Unit, x, y int) error {
	if len(f.units) >= f.MaxUnits {
		return errors.New("уже максимальное количество юнитов")
	}

	if x <= 0 || x >= f.Height || y >= f.Width || 0 >= y {
		return errors.New("выход за размер поля")
	}

	if err := f.MoveUnit(u, x, y); err != nil {
		return fmt.Errorf("ошибка при добавлении юнита: %w", err)
	}

	f.units = append(f.units, u)
	return nil
}

func (f *GameField) RemoveUnit(index int) error {
	if len(f.units) < index || 0 > index {
		return errors.New("неверный индекс")
	}

	f.units = append(f.units[:index], f.units[index+1:]...)
	return nil
}

func (f *GameField) Clone() *GameField {
	newGameField := NewGameField(f.Width, f.Height, f.MaxUnits)
	it := f.Iterator()
	for it.Next() {
		u := it.Value()
		newGameField.units = append(newGameField.units, u.Clone())
	}

	return newGameField
}

func (f *GameField) MoveUnit(u unit.Unit, x, y int) error {
	if 0 > x && x > f.Width && 0 > y && y > f.Height {
		return errors.New("выход за границы")
	}

	it := f.Iterator()
	for it.Next() {
		unit := it.Value()
		if unit.GetPosition() == u.GetPosition() {
			return fmt.Errorf("клетка (%d, %d) занята", x, y)
		}
	}

	return u.Move(x, y)
}

func (f *GameField) PrintGameField() error {
	f.DeleteDeadUnits()
	fmt.Print("\n\n   ")
	for x := range f.Width {
		fmt.Printf("%d  ", x%10)
	}
	fmt.Println()

	for y := range f.Height {
		fmt.Printf("%d  ", y%10)

		for x := range f.Width {
			symbol := "⬛"
			it := f.Iterator()
			for it.Next() {
				u := it.Value()
				pos := u.GetPosition()
				if pos.X == x && pos.Y == y {
					symbol = u.GetSymbol()
					break
				}
			}
			fmt.Printf("%s ", symbol)
		}
		fmt.Println()
	}
	return nil
}

func (f *GameField) PrintUnitsInfo() error {
	it := f.Iterator()
	for it.Next() {
		u := it.Value()
		fmt.Printf("\n[%s] %-15s | Health: %3d Armor: %2d Attack: %2d ",
			u.GetSymbol(), u.GetName(), u.GetHealth().GetValue(),
			u.GetArmor().GetValue(), u.GetAttack().GetValue())
	}
	return nil
}

func (f *GameField) DeleteDeadUnits() error {
	it := f.Iterator()
	for it.Next() {
		u := it.Value()
		if u.GetHealth().GetValue() == 0 {
			f.units = append(f.units[:it.Index], f.units[it.Index+1:]...)
		}
	}
	return nil
}

type GameFieldIterator struct {
	Field *GameField
	Index int
}

func (f *GameField) Iterator() *GameFieldIterator {
	return &GameFieldIterator{
		Field: f,
		Index: -1,
	}
}

func (it *GameFieldIterator) Next() bool {
	it.Index++
	return it.Index < len(it.Field.units)
}

func (it *GameFieldIterator) Value() unit.Unit {
	if it.Index >= 0 && it.Index <= len(it.Field.units) {
		return it.Field.units[it.Index]
	}

	return nil
}
