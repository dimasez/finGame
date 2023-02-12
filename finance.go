package main

import (
	"fmt"
	"math/rand"
)

var CurrentAssets = Assets{
	Machine: MeasureTypes{
		1000,
		1,
	},
	Raw: MeasureTypes{
		100,
		50,
	},
	Product: MeasureTypes{
		100,
		25,
	},
	Money: 6,
}

var CurrentLiability = Liability{
	Equity: 500,
	Debt:   706,
}

//Суммирование/подсчет активов
func assetsCalc() int {
	return CurrentAssets.Machine.Money + CurrentAssets.Raw.Money + CurrentAssets.Product.Money + CurrentAssets.Money
}

//Подсчет ПАССИВА
func equityCalc() int {
	return CurrentLiability.Equity + CurrentLiability.Debt
}

//Продажа товара - все 5 операций/функции
func productSell(amount, price int) {
	//Проверяем - хватает ли товара для продажи
	if CurrentAssets.Product.Physical >= amount {
		//Увеличиваем деньги на выручку
		CurrentAssets.Money = CurrentAssets.Money + amount*price
		//Уменьшаем товар финансово - на себестоимость проданного товара
		productCost := (CurrentAssets.Product.Money / CurrentAssets.Product.Physical) * amount //Себестоимость проданной партии товара
		CurrentAssets.Product.Money = CurrentAssets.Product.Money - productCost
		//Уменьшаем товар физически
		CurrentAssets.Product.Physical = CurrentAssets.Product.Physical - amount
		//Увеличиваем Equity на размер прибыли
		profit := amount*price - productCost
		CurrentLiability.Equity = CurrentLiability.Equity + profit
		//Выводим результат в терминал
		fmt.Println("Вы продали", amount, "единиц товара по цене", price, "рублей")
		fmt.Println("Ваша прибыль составила", profit, "рублей")
		fmt.Println()
	} else {
		fmt.Println("У вас недостаточно товара для продажи")
		fmt.Println("На складе", CurrentAssets.Product.Physical, "единиц товара")
		fmt.Println("Вы хотите продать", amount, "единиц товара")
		fmt.Println()
	}

}

//Купить сырье
func buyRaw(amount int) {
	//Себестоимость всегда 2
	//Проверка достаточно денег на счете для покупки
	if CurrentAssets.Money >= amount*2 {
		//Увеличиваем сырье
		CurrentAssets.Raw.Physical = CurrentAssets.Raw.Physical + amount
		CurrentAssets.Raw.Money = CurrentAssets.Raw.Money + amount*2

		//Уменьшаем деньги
		CurrentAssets.Money = CurrentAssets.Money - amount*2

		fmt.Println("Вы купили", amount, "единиц сырья")
		fmt.Println()
	} else {
		fmt.Println("У вас недостаточно средств для покупки", amount, "единиц сырья")
		fmt.Println()
	}
}

//Производство. Amount - кол-во товара, производимого за ход одним станком
func productionTurn(amount int) {
	if amount*2*CurrentAssets.Machine.Physical <= CurrentAssets.Raw.Physical {
		rawCost := CurrentAssets.Raw.Money / CurrentAssets.Raw.Physical //Расчет текущей себестоимости сырья
		//Увеличиваем товар
		CurrentAssets.Product.Physical = CurrentAssets.Product.Physical + amount*CurrentAssets.Machine.Physical
		CurrentAssets.Product.Money = CurrentAssets.Product.Money + rawCost*(amount*2)*CurrentAssets.Machine.Physical

		//Уменьшаем сырье
		CurrentAssets.Raw.Physical = CurrentAssets.Raw.Physical - (amount*2)*CurrentAssets.Machine.Physical
		CurrentAssets.Raw.Money = CurrentAssets.Raw.Money - rawCost*(amount*2)*CurrentAssets.Machine.Physical

		fmt.Println("Произведено", amount*CurrentAssets.Machine.Physical, "единиц продукции")
		fmt.Println()
	} else {
		fmt.Println("У вас недостаточно сырья для производства продукции")
		fmt.Println("Нужно минимум", amount*2*CurrentAssets.Machine.Physical, "единиц сырья")
		fmt.Println()
	}
}

//Погашение займа
func loanRepayment (payment int) {
	//Проверяем - хватает ли денег на погашение
	if payment <= CurrentAssets.Money {
		//Проверяем - превышает ли займ сумму погашения
		if payment <= CurrentLiability.Debt {
			CurrentLiability.Debt = CurrentLiability.Debt - payment
			CurrentAssets.Money = CurrentAssets.Money - payment
			fmt.Println("Вы погасили часть займа на сумму", payment, "рублей")
			fmt.Println()
		} else {
			//Уменьшаем деньги на финальный остаток займа
			CurrentAssets.Money = CurrentAssets.Money - CurrentLiability.Debt
			//Обнуляем займ
			CurrentLiability.Debt = 0
			fmt.Println("Поздравляем - ВЫ ПОЛНОСТЬЮ ПОГАСИЛИ ЗАЙМ")
			fmt.Println()
		}
	} else {
		fmt.Println("У вас недостаточно средств на счете для погашения займа в размере", payment, "рублей")
		fmt.Println()
	}
}

//Покупка станка
func machineBuy(price int) {
	if CurrentAssets.Money >= price {
		CurrentAssets.Machine.Physical++
		CurrentAssets.Machine.Money = CurrentAssets.Machine.Money + price
		CurrentAssets.Money = CurrentAssets.Money - price
		fmt.Println("Поздравляем - Вы купили новый станок!")
		fmt.Println("Ваши производственные мощности увеличились!")
		fmt.Println("Теперь вы будете выпускать больше единиц продукции за один ход!")
		fmt.Println()

	} else {
		fmt.Println("У вас недостаточно средств для покупки станка")
		fmt.Println()
	}
}

//Лоттерея
func lottery () {
	if CurrentAssets.Money >= 5 {
		luckyNumber := rand.Intn(10)
		if luckyNumber == 5 {
			CurrentAssets.Money = CurrentAssets.Money + 95
			CurrentLiability.Equity = CurrentLiability.Equity + 95
			fmt.Println("УРА! Вы выиграли в лотерею 100 рублей!!!")
			fmt.Println()
		} else {
			CurrentAssets.Money = CurrentAssets.Money - 5
			CurrentLiability.Equity = CurrentLiability.Equity - 5
			fmt.Println("К сожалению - вы ничего не выиграли :(")
			fmt.Println("Повезет в следующий раз :)")
			fmt.Println()
		}
	} else {
		fmt.Println("У вас недостаточно средств, чтобы купить лотерейный билет :(")
		fmt.Println()
	}
}

func ratePayment() {
	//Если хватает денег - списываем проценты
	if CurrentLiability.Debt/DebtRate <= CurrentAssets.Money {
		CurrentAssets.Money = CurrentAssets.Money - CurrentLiability.Debt/DebtRate
		CurrentLiability.Equity = CurrentLiability.Equity - CurrentLiability.Debt/DebtRate
		fmt.Println("С вас списано", CurrentLiability.Debt/DebtRate, "рублей процентов по кредиту")
	} else {
		//Если денег не хватает - увеличиваем долг - уменьшаем Equity
		CurrentLiability.Debt = CurrentLiability.Debt + CurrentLiability.Debt/DebtRate
		CurrentLiability.Equity = CurrentLiability.Equity - CurrentLiability.Debt/DebtRate
		fmt.Println("С вас списано", CurrentLiability.Debt/DebtRate, "рублей процентов по кредиту")
	}
	DebtTurn = 0
}