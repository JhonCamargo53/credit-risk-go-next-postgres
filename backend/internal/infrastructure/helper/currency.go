package helper

import "github.com/leekchan/accounting"

var ac = accounting.Accounting{
	Symbol:    "$",
	Precision: 0,
	Thousand:  ".",
	Decimal:   ",",
}

func FormatCOP(value float64) string {
	return ac.FormatMoney(value)
}
