package factory

import (
	"fmt"
	"oop/unit"
)

type UnitFactory interface {
	CreateUnit(unitType string) (unit.Unit, error)
}

type InfantryFactory struct{}
type ArcherFactory struct{}
type CavalryFactory struct{}

func (i *InfantryFactory) CreateUnit(unitType string) (unit.Unit, error) {
	switch unitType {
	case "swordsman":
		return &unit.Swordsman{Infantry: unit.Infantry{
			BaseUnit: *unit.NewBaseUnit("Мечник", "⚔️ ", 100, 20, 30)}}, nil
	case "spearman":
		return &unit.Spearman{Infantry: unit.Infantry{
			BaseUnit: *unit.NewBaseUnit("Копейщик", "🔱", 90, 15, 40)}}, nil
	}
	return nil, fmt.Errorf("неизвестный тип пехоты: %s", unitType)

}

func (i *ArcherFactory) CreateUnit(unitType string) (unit.Unit, error) {
	switch unitType {
	case "light":
		return &unit.LightArcher{Archer: unit.Archer{
			BaseUnit: *unit.NewBaseUnit("Лёгкий лучник", "🏹", 70, 5, 35), Range: 5}}, nil
	case "heavy":
		return &unit.HeavyArcher{Archer: unit.Archer{
			BaseUnit: *unit.NewBaseUnit("Тяжёлый лучник", "🎯", 85, 15, 45), Range: 15}}, nil
	}
	return nil, fmt.Errorf("неизвестный тип лучника: %s", unitType)

}

func (i *CavalryFactory) CreateUnit(unitType string) (unit.Unit, error) {
	switch unitType {
	case "cataphract":
		return &unit.Cataphract{Cavalry: unit.Cavalry{
			BaseUnit: *unit.NewBaseUnit("Катафракт", "🐎", 120, 40, 50), Speed: 4}}, nil
	case "horse_archer":
		return &unit.HorseArcher{Cavalry: unit.Cavalry{
			BaseUnit: *unit.NewBaseUnit("Конный лучник", "🏇", 95, 15, 35), Speed: 6}}, nil
	}
	return nil, fmt.Errorf("неизвестный тип конницы: %s", unitType)

}
