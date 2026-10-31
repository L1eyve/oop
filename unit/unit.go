package unit

import (
	"fmt"
	"oop/attributes"
)

type Unit interface {
	GetName() string
	GetHealth() *attributes.Health
	GetArmor() *attributes.Armor
	GetAttack() *attributes.Attack
	GetPosition() attributes.Point
	GetSymbol() string
	Move(x, y int) error
	Clone() Unit
	AttackTarget(target Unit)
}

type BaseUnit struct {
	name   string
	symbol string
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

func (b *BaseUnit) GetName() string { return b.name }

func (b *BaseUnit) GetHealth() *attributes.Health { return b.health }

func (b *BaseUnit) GetArmor() *attributes.Armor { return b.armor }

func (b *BaseUnit) GetAttack() *attributes.Attack { return b.attack }

func (b *BaseUnit) GetPosition() attributes.Point { return b.pos }

func (b *BaseUnit) GetSymbol() string { return b.symbol }

func (b *BaseUnit) Move(x, y int) error {
	b.pos = attributes.Point{
		X: x,
		Y: y,
	}
	fmt.Printf("\n>>> %s перемещается на (%d, %d) <<<", b.name, x, y)
	return nil
}

func (b *BaseUnit) AttackTarget(target Unit) {
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

func (s *Swordsman) Clone() Unit {
	return &Swordsman{Infantry{BaseUnit{
		name:   "Мечник",
		symbol: "⚔️ ",
		health: attributes.NewHealth(s.health.GetValue()),
		armor:  attributes.NewArmor(s.armor.GetValue()),
		attack: attributes.NewAttack(s.attack.GetValue()),
		pos:    s.pos,
	}}}
}

func (s *Spearman) Clone() Unit {
	return &Spearman{Infantry{BaseUnit{
		name:   "Копейщик",
		symbol: "🔱",
		health: attributes.NewHealth(s.health.GetValue()),
		armor:  attributes.NewArmor(s.armor.GetValue()),
		attack: attributes.NewAttack(s.attack.GetValue()),
		pos:    s.pos,
	}}}
}

func (l *LightArcher) Clone() Unit {
	return &LightArcher{Archer{
		BaseUnit: BaseUnit{
			name:   "Лёгкий лучник",
			symbol: "🏹",
			health: attributes.NewHealth(l.health.GetValue()),
			armor:  attributes.NewArmor(l.armor.GetValue()),
			attack: attributes.NewAttack(l.attack.GetValue()),
			pos:    l.pos,
		},
		Range: l.Range,
	}}
}

func (h *HeavyArcher) Clone() Unit {
	return &HeavyArcher{Archer{
		BaseUnit: BaseUnit{
			name:   "Тяжёлый лучник",
			symbol: "🎯",
			health: attributes.NewHealth(h.health.GetValue()),
			armor:  attributes.NewArmor(h.armor.GetValue()),
			attack: attributes.NewAttack(h.attack.GetValue()),
			pos:    h.pos,
		},
		Range: h.Range,
	}}
}

func (c *Cataphract) Clone() Unit {
	return &Cataphract{Cavalry{
		BaseUnit: BaseUnit{
			name:   "Катафракт",
			symbol: "🐎",
			health: attributes.NewHealth(c.health.GetValue()),
			armor:  attributes.NewArmor(c.armor.GetValue()),
			attack: attributes.NewAttack(c.attack.GetValue()),
			pos:    c.pos,
		},
		Speed: c.Speed,
	}}
}

func (h *HorseArcher) Clone() Unit {
	return &HorseArcher{Cavalry{
		BaseUnit: BaseUnit{
			name:   "Конный лучник",
			symbol: "🏇",
			health: attributes.NewHealth(h.health.GetValue()),
			armor:  attributes.NewArmor(h.armor.GetValue()),
			attack: attributes.NewAttack(h.attack.GetValue()),
			pos:    h.pos,
		},
		Speed: h.Speed,
	}}
}

func NewBaseUnit(newName, newSymbol string, hp, armor, atk int) *BaseUnit {
	return &BaseUnit{
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
