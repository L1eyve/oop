package main

import (
	"fmt"
	"oop/base"
	"oop/field"
)

func main() {
	gameField := field.NewGameField(8, 8, 10)

	base := base.NewBase(300, 30, 5, 0, 10, 1, 5)
	gameField.AddBase(base, 1, 5)

	sw1, err := gameField.BaseAddUnit(base.GetId(), "swordsman")
	if err != nil {
		fmt.Println("ошибка при добавлении юнита базой: ", err)
	}
	sw2, err := gameField.BaseAddUnit(base.GetId(), "swordsman")
	if err != nil {
		fmt.Println("ошибка при добавлении юнита базой: ", err)
	}

	gameField.PrintGameField()
	gameField.PrintUnitsInfo()

	sw1.AttackTarget(sw2)

	if err := gameField.MoveUnit(sw1, 7, 7); err != nil {
		fmt.Println("\nОшибка перемещения: ", err)
	}

	gameField.PrintGameField()
	gameField.PrintUnitsInfo()

	gameField.PrintGameField()
	gameField.PrintUnitsInfo()
}
