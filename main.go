package main

import (
	"fmt"
)

var GameTurn = 0
var DebtTurn = 0
var DebtRate = 50       //Ставка по займу. На нее делится тело займа - 25 - это 4%
var IsGameBegin = false //Для первого хода не нужен вывод действий

func main() {

	//Бесконечный цикл
	for {
		fmt.Println("Ваш ход обсчитывается")

		//1. Спиннер
		showSpinner()

		//2. Главная игровая функция. Функция продажи товара - пересчет состояния
		if IsGameBegin == true {
			//Меню для выбора действий игроку
			actionsOffer()
			//Производство
			productionTurn(4)
		}

		//3. Отрисовка таблицы
		drawTable()

		//Конец хода. Увеличиваем счетчики
		IsGameBegin = true
		moneySound()
		GameTurn++
		DebtTurn++
		fmt.Println()

		//Каждые 10 ходов списываются проценты (4%) по
		if DebtTurn == 10 {
			ratePayment()
		}
	}
}
