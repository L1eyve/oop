package main

import (
	"fmt"
	"oop/factory"
	"oop/field"
)

func main() {
	gameField := field.NewGameField(8, 8, 10)

	infFactory := &factory.InfantryFactory{}
	archFactory := &factory.ArcherFactory{}
	cavFactory := &factory.CavalryFactory{}

	swordsman, _ := infFactory.CreateUnit("swordsman")
	spearman, _ := infFactory.CreateUnit("spearman")
	lightArcher, _ := archFactory.CreateUnit("light")
	cataphract, _ := cavFactory.CreateUnit("cataphract")

	gameField.AddUnit(swordsman, 1, 3)
	gameField.AddUnit(spearman, 2, 1)
	gameField.AddUnit(lightArcher, 4, 3)
	gameField.AddUnit(cataphract, 6, 5)

	gameField.PrintGameField()
	gameField.PrintUnitsInfo()

	swordsman.AttackTarget(cataphract)

	if err := gameField.MoveUnit(cataphract, 7, 7); err != nil {
		fmt.Println("\nОшибка перемещения: ", err)
	}

	cataphract.AttackTarget(lightArcher)
	cataphract.AttackTarget(lightArcher)

	gameField.PrintGameField()
	gameField.PrintUnitsInfo()

	cataphract.AttackTarget(lightArcher)

}
