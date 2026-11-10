package unit

import (
	"fmt"
	"oop/attributes"
	"oop/damageable"
)

type Unit interface {
	GetId() string
	GetIdBase() string
	GetName() string
	GetHealth() *attributes.Health
	GetArmor() *attributes.Armor
	GetAttack() *attributes.Attack
	GetPosition() attributes.Point
	GetSymbol() int
	Move(x, y int) error
	Clone() Unit
	AttackTarget(target damageable.Damageable)
}

type BaseUnit struct {
	id     string
	idBase string
	name   string
	symbol int
	health *attributes.Health
	armor  *attributes.Armor
	attack *attributes.Attack
	pos    attributes.Point
}

type Infantry struct {
	BaseUnit
}

type Swordsman struct {
	Infantry
}

type Spearman struct {
	Infantry
}

type Archer struct {
	BaseUnit
	Range int
}

type LightArcher struct {
	Archer
}

type HeavyArcher struct {
	Archer
}

type Cavalry struct {
	BaseUnit
	Speed int
}

type Cataphract struct {
	Cavalry
}

type HorseArcher struct {
	Cavalry
}

func (b *BaseUnit) GetId() string { return b.id }

func (b *BaseUnit) GetIdBase() string { return b.idBase }

func (b *BaseUnit) GetName() string { return b.name }

func (b *BaseUnit) GetHealth() *attributes.Health { return b.health }

func (b *BaseUnit) GetArmor() *attributes.Armor { return b.armor }

func (b *BaseUnit) GetAttack() *attributes.Attack { return b.attack }

func (b *BaseUnit) GetPosition() attributes.Point { return b.pos }

func (b *BaseUnit) GetSymbol() int { return b.symbol }

func (b *BaseUnit) Move(x, y int) error {
	b.pos = attributes.Point{
		X: x,
		Y: y,
	}
	fmt.Printf("\n>>> %s перемещается на (%d, %d) <<<", b.name, x, y)
	return nil
}

func (b *BaseUnit) AttackTarget(target damageable.Damageable) {
	if target.GetHealth().GetValue() == 0 {
		fmt.Printf("\n>>> %s уже мертв, атаковать нельзя <<<", target.GetName())
		return
	}
	damage := b.GetAttack().CalculateDamage(target.GetArmor())
	fmt.Printf("\n>>> %s атакует %s <<<", b.name, target.GetName())
	if dead := target.GetHealth().Decrease(damage); dead {
		fmt.Printf("\n>>> %s умирает <<<", target.GetName())
	}

}

func (b *BaseUnit) cloneBase() BaseUnit {
	return BaseUnit{
		id:     b.id,
		idBase: b.idBase,
		name:   b.name,
		symbol: b.symbol,
		health: attributes.NewHealth(b.health.GetValue()),
		armor:  attributes.NewArmor(b.armor.GetValue()),
		attack: attributes.NewAttack(b.attack.GetValue()),
		pos:    b.pos,
	}
}

func (s *Swordsman) Clone() Unit {
	return &Swordsman{Infantry{BaseUnit: s.BaseUnit.cloneBase()}}
}

func (s *Spearman) Clone() Unit {
	return &Spearman{Infantry{BaseUnit: s.BaseUnit.cloneBase()}}
}

func (l *LightArcher) Clone() Unit {
	return &LightArcher{Archer{
		BaseUnit: l.BaseUnit.cloneBase(),
		Range:    l.Range,
	}}
}

func (h *HeavyArcher) Clone() Unit {
	return &HeavyArcher{Archer{
		BaseUnit: h.BaseUnit.cloneBase(),
		Range:    h.Range,
	}}
}

func (c *Cataphract) Clone() Unit {
	return &Cataphract{Cavalry{
		BaseUnit: c.BaseUnit.cloneBase(),
		Speed:    c.Speed,
	}}
}

func (h *HorseArcher) Clone() Unit {
	return &HorseArcher{Cavalry{
		BaseUnit: h.BaseUnit.cloneBase(),
		Speed:    h.Speed,
	}}
}

func NewBaseUnit(id, idBase, newName string, newSymbol, hp, armor, atk int) *BaseUnit {
	return &BaseUnit{
		id:     id,
		idBase: idBase,
		name:   newName,
		symbol: newSymbol,
		health: attributes.NewHealth(hp),
		armor:  attributes.NewArmor(armor),
		attack: attributes.NewAttack(atk),
		pos: attributes.Point{
			X: 0,
			Y: 0,
		},
	}
}
