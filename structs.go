package main

//Актив
type Assets struct {
	Machine MeasureTypes
	Raw     MeasureTypes
	Product MeasureTypes
	Money   int
}

//Пассив
type Liability struct {
	Equity int
	Debt   int
}

//Единицы измерения = в деньгах и в штуках
type MeasureTypes struct {
	Money    int
	Physical int
}
