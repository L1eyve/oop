package base

import (
	"fmt"
	"oop/attributes"
	"oop/id"
	"oop/unit"
)

type Base struct {
	id       string
	name     string
	symbol   int
	health   *attributes.Health
	armor    *attributes.Armor
	wood     *attributes.Resource
	iron     *attributes.Resource
	maxUnits int
	pos      attributes.Point
	units    []unit.Unit
}

func NewBase(health, armor, wood, iron, maxUnits, x, y int) *Base {
	id, err := id.GenerateId()
	if err != nil {
		fmt.Println("ошибка при генерации id новой базы: ", err)
		return nil
	}
	return &Base{
		id:       id,
		name:     "База",
		symbol:   0,
		health:   attributes.NewHealth(health),
		armor:    attributes.NewArmor(armor),
		wood:     attributes.NewResource(wood),
		iron:     attributes.NewResource(iron),
		maxUnits: maxUnits,
		units:    []unit.Unit{},
		pos: attributes.Point{
			X: x,
			Y: y,
		},
	}
}

func (b *Base) GetName() string { return b.name }

func (b *Base) AddUnit(u unit.Unit) { b.units = append(b.units, u) }

func (b *Base) GetId() string { return b.id }

func (b *Base) GetHealth() *attributes.Health { return b.health }

func (b *Base) GetWood() *attributes.Resource { return b.wood }

func (b *Base) GetIron() *attributes.Resource { return b.iron }

func (b *Base) GetArmor() *attributes.Armor { return b.armor }

func (b *Base) GetPosition() attributes.Point { return b.pos }

func (b *Base) GetSymbol() int { return b.symbol }

func (b *Base) GetUnits() []unit.Unit { return b.units }

func (b *Base) GetMaxUnits() int { return b.maxUnits }

func (b *Base) RemoveUnit(id string) {
	for i, u := range b.units {
		if u.GetId() == id {
			b.units = append(b.units[:i], b.units[i+1:]...)
		}
	}
}

func (b *Base) Clone() *Base {
	return &Base{
		id:       b.id,
		symbol:   0,
		health:   attributes.NewHealth(b.health.GetValue()),
		armor:    attributes.NewArmor(b.armor.GetValue()),
		wood:     attributes.NewResource(b.wood.GetValue()),
		iron:     attributes.NewResource(b.iron.GetValue()),
		maxUnits: b.maxUnits,
		pos:      b.pos,
		units:    b.units,
	}
}
