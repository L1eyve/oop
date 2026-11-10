package field

import (
	"errors"
	"fmt"
	"oop/attributes"
	"oop/base"
	"oop/factory"
	"oop/render"
	"oop/unit"
)

type GameField struct {
	Width, Height, MaxUnits int
	units                   []unit.Unit
	bases                   map[string]*base.Base
	cells                   [][]int
}

func NewGameField(width, height, maxUnits int) *GameField {
	cells := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
		for j := range cells[i] {
			cells[i][j] = -1
		}
	}

	return &GameField{
		Width:    width,
		Height:   height,
		MaxUnits: maxUnits,
		units:    []unit.Unit{},
		bases:    map[string]*base.Base{},
		cells:    cells,
	}
}

func (f *GameField) IfCellFree(x, y int) (bool, error) {
	if x < 0 || x >= f.Width || y >= f.Height || y < 0 {
		return false, errors.New("выход за границы")
	}

	if f.cells[y][x] != -1 {
		return false, fmt.Errorf("клетка (%d, %d) занята", x, y)
	} else {
		return true, nil
	}
}

func (f *GameField) AddUnit(u unit.Unit, x, y int) error {
	if len(f.units) >= f.MaxUnits {
		return errors.New("уже максимальное количество юнитов")
	}

	if ok, err := f.IfCellFree(x, y); !ok {
		return fmt.Errorf("Ошибка при проверке клетки: %w", err)
	}

	u.Move(x, y)
	f.cells[y][x] = u.GetSymbol()
	f.units = append(f.units, u)
	return nil
}

func (f *GameField) AddBase(b *base.Base, x, y int) error {
	if ok, err := f.IfCellFree(x, y); !ok {
		return fmt.Errorf("Ошибка при проверке клетки: %w", err)
	}

	f.cells[y][x] = b.GetSymbol()

	f.bases[b.GetId()] = b
	return nil
}

func (f *GameField) RemoveUnit(index int) error {
	if len(f.units) < index || 0 > index {
		return errors.New("неверный индекс")
	}

	f.units = append(f.units[:index], f.units[index+1:]...)
	return nil
}

func (f *GameField) GetFreeCell(idBase string) (int, int) {
	b := f.bases[idBase]
	pos := b.GetPosition()
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if y == 0 && x == 0 {
				continue
			}

			newX, newY := pos.X+x, pos.Y+y
			if ok, _ := f.IfCellFree(newX, newY); ok {
				return newX, newY
			}

		}
	}
	return -1, -1
}

func (f *GameField) BaseAddUnit(idBase, unitType string) (unit.Unit, error) {
	var unitFactory factory.UnitFactory
	b, ok := f.bases[idBase]
	if !ok {
		return nil, errors.New("база не найдена")
	}

	if b.GetMaxUnits() <= len(b.GetUnits()) {
		return nil, errors.New("количесто юнитов базы уже максимально")
	}

	cost, ok := attributes.Costs[unitType]
	if !ok {
		return nil, errors.New("неизвестный тип юнита")
	}

	if b.GetIron().GetValue() < cost.Iron && b.GetWood().GetValue() < cost.Wood {
		return nil, errors.New("недостаточно ресурсов")
	}

	x, y := f.GetFreeCell(idBase)
	if x < 0 || y < 0 {
		return nil, errors.New("нет свободных клеток возле базы")
	}

	switch unitType {
	case "swordsman", "spearman":
		unitFactory = &factory.InfantryFactory{}
	case "light", "heavy":
		unitFactory = &factory.ArcherFactory{}
	case "cataphract", "horse_archer":
		unitFactory = &factory.CavalryFactory{}
	}

	u, err := unitFactory.CreateUnit(unitType, idBase)
	if err != nil {
		return nil, fmt.Errorf("ошибка при добавлении юнита базой: %w", err)
	}

	b.GetWood().Decrease(cost.Wood)
	b.GetIron().Decrease(cost.Iron)
	f.AddUnit(u, x, y)
	b.AddUnit(u)
	return u, nil
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
	if ok, err := f.IfCellFree(x, y); !ok {
		return fmt.Errorf("Ошибка при проверке клетки: %w", err)
	}

	pos := u.GetPosition()
	f.cells[pos.Y][pos.X] = -1
	f.cells[y][x] = u.GetSymbol()

	return u.Move(x, y)
}

func (f *GameField) PrintGameField() error {
	f.DeleteDeadUnits()
	fmt.Print("\n\n   ")
	for x := range f.Width {
		fmt.Printf("%d ", x%10)
	}
	fmt.Println()

	for y, c1 := range f.cells {
		fmt.Printf("%d  ", y%10)
		for _, c2 := range c1 {
			fmt.Print(render.IdentifySymbol(c2))
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
			render.IdentifySymbol(u.GetSymbol()), u.GetName(), u.GetHealth().GetValue(),
			u.GetArmor().GetValue(), u.GetAttack().GetValue())
	}
	fmt.Println()
	return nil
}

func (f *GameField) DeleteDeadUnits() error {
	aliveUnits := make([]unit.Unit, 0, f.MaxUnits)
	for _, u := range f.units {
		if u.GetHealth().GetValue() > 0 {
			aliveUnits = append(aliveUnits, u)
		} else {
			pos := u.GetPosition()
			f.cells[pos.Y][pos.X] = -1
			baseId := u.GetIdBase()
			b := f.bases[baseId]
			b.RemoveUnit(u.GetId())
		}
	}

	f.units = aliveUnits
	return nil
}

func (f *GameField) Iterator() *GameFieldIterator {
	return &GameFieldIterator{
		Field: f,
		Index: -1,
	}
}
