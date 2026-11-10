package damageable

import "oop/attributes"

type Damageable interface {
	GetId() string
	GetName() string
	GetHealth() *attributes.Health
	GetArmor() *attributes.Armor
	GetPosition() attributes.Point
	GetSymbol() int
}
