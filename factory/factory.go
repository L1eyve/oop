package factory

import (
	"fmt"
	"oop/id"
	"oop/unit"
)

type UnitFactory interface {
	CreateUnit(unitType, idBase string) (unit.Unit, error)
}

type InfantryFactory struct{}
type ArcherFactory struct{}
type CavalryFactory struct{}

func (i *InfantryFactory) CreateUnit(unitType, idBase string) (unit.Unit, error) {
	id, err := id.GenerateId()
	if err != nil {
		return nil, fmt.Errorf("ошибка при генерации id новой пехоты: %w", err)
	}
	switch unitType {
	case "swordsman":
		return &unit.Swordsman{Infantry: unit.Infantry{
			BaseUnit: *unit.NewBaseUnit(id, idBase, "Мечник", 1, 100, 20, 30)}}, nil
	case "spearman":
		return &unit.Spearman{Infantry: unit.Infantry{
			BaseUnit: *unit.NewBaseUnit(id, idBase, "Копейщик", 2, 90, 15, 40)}}, nil
	}
	return nil, fmt.Errorf("неизвестный тип пехоты: %s", unitType)

}

func (i *ArcherFactory) CreateUnit(unitType, idBase string) (unit.Unit, error) {
	id, err := id.GenerateId()
	if err != nil {
		return nil, fmt.Errorf("ошибка при генерации id нового лучника: %w", err)
	}

	switch unitType {
	case "light":
		return &unit.LightArcher{Archer: unit.Archer{
			BaseUnit: *unit.NewBaseUnit(id, idBase, "Лёгкий лучник", 3, 70, 5, 35), Range: 5}}, nil
	case "heavy":
		return &unit.HeavyArcher{Archer: unit.Archer{
			BaseUnit: *unit.NewBaseUnit(id, idBase, "Тяжёлый лучник", 4, 85, 15, 45), Range: 15}}, nil
	}
	return nil, fmt.Errorf("неизвестный тип лучника: %s", unitType)

}

func (i *CavalryFactory) CreateUnit(unitType, idBase string) (unit.Unit, error) {
	id, err := id.GenerateId()
	if err != nil {
		return nil, fmt.Errorf("ошибка при генерации id новой кавалерии: %w", err)
	}

	switch unitType {
	case "cataphract":
		return &unit.Cataphract{Cavalry: unit.Cavalry{
			BaseUnit: *unit.NewBaseUnit(id, idBase, "Катафракт", 5, 120, 40, 50), Speed: 4}}, nil
	case "horse_archer":
		return &unit.HorseArcher{Cavalry: unit.Cavalry{
			BaseUnit: *unit.NewBaseUnit(id, idBase, "Конный лучник", 6, 95, 15, 35), Speed: 6}}, nil
	}
	return nil, fmt.Errorf("неизвестный тип конницы: %s", unitType)

}
