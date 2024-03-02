package main

import (
	"calc"
	"fmt"
	"github.com/bojanz/currency"
)

func main() {
	// считаем баланс в рублях
	sum := calc.AddInts(1, 56, 78, -33)

	// вывод будет с копейками, поэтому умножаем на 100
	total, err := currency.NewAmountFromInt64(int64(sum*100), "RUB")
	if err != nil {
		panic(err)
	}
	// справа от значения будет добавляться знак рубля
	locale := currency.NewLocale("ru")
	formatter := currency.NewFormatter(locale)

	fmt.Println("Баланс:", formatter.Format(total))
}
