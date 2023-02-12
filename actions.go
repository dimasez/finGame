package main

import (
	"fmt"
	"math/rand"
)

//Меню с выбором действий
func actionsOffer() {
	fmt.Println()
	fmt.Println("--------------------------------------------------------")
	fmt.Println()
	fmt.Println("ИДЕТ", GameTurn, "ДЕНЬ  ДЕЯТЕЛЬНОСТИ ВАШЕГО ПРЕДПРИЯТИЯ")
	fmt.Println()
	fmt.Println("ВАМ ДОСТУПНЫ СЛЕДУЮЩИЕ ОПЕРАЦИИ:")
	rndAmount1 := 5 + rand.Intn(10)
	rndPrice1 := 7 + rand.Intn(2)
	fmt.Println("1. Продать", rndAmount1, "единиц товара по цене", rndPrice1, "рублей за штуку")
	rndAmount2 := 10 + rand.Intn(20)
	rndPrice2 := 5 + rand.Intn(2)
	fmt.Println("2. Продать", rndAmount2, "единиц товара по цене", rndPrice2, "рублей за штуку")
	fmt.Println()
	rndAmount3 := 7 + rand.Intn(10)
	fmt.Println("3. Купить", rndAmount3, "единиц сырья по цене", 2, "рублей за штуку")
	rndAmount4 := 2 + rand.Intn(8)
	fmt.Println("4. Купить", rndAmount4, "единиц сырья по цене", 2, "рублей за штуку")
	rndAmount5 := 300 + rand.Intn(300)
	fmt.Println()
	fmt.Println("5. Купить 1 станок по цене", rndAmount5, "рублей")
	fmt.Println("6. Погасить займ на сумму", 50, "рублей")
	fmt.Println("7. Купить лотерейный билет за", 5, "рублей. Есть шанс выиграть 100 рублей!!!")
	fmt.Println()
	fmt.Print("Укажите номер операции (0 - ничего не делать): ")

	var answer int
	_, err := fmt.Scanln(&answer)
	fmt.Println()
	if err != nil {
		fmt.Println("Вы отказались от сделки")
		fmt.Println()
	}

	//1-2 Продажа продукта. ОПТ и РОЗНИЦА
	if answer == 1 {
		productSell(rndAmount1, rndPrice1)
	} else if answer == 2 {
		productSell(rndAmount2, rndPrice2)
	} else if answer == 3 {
		buyRaw(rndAmount3)
	} else if answer == 4 {
		buyRaw(rndAmount4)
	} else if answer == 5 {
		machineBuy(rndAmount5)
	} else if answer == 6 {
		loanRepayment(50)
	} else if answer == 7 {
		lottery()
	} else {
		fmt.Println("Вы отказались от сделки")
		fmt.Println()
	}
}

