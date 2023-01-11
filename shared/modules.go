package shared

import (
	"emailQue/utils"
	"reflect"
	"time"

	"github.com/CloudyKit/jet"
	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
)

var templates = jet.NewHTMLSet("./templates")
var moneyFormat = accounting.Accounting{Symbol: "₦ ", Precision: 2}

var logoByte, _ = utils.OpenandReadFile("templates/logo.svg")
var logoEncoded = utils.EncodeImageBase64(logoByte)

func fmtMoney(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("fmtMoney", 1, 1)
	val, ok := args.Get(0).Interface().(decimal.Decimal)
	if !ok {
		zero := "0.0"
		return reflect.ValueOf(zero)
	}

	return reflect.ValueOf(moneyFormat.FormatMoneyDecimal(val))
}

func formatTime(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("fmtTime", 1, 1)
	val, ok := args.Get(0).Interface().(time.Time)
	if !ok {
		zero := "0001-01-01 00:00:00 +0000"
		return reflect.ValueOf(zero)
	}

	return reflect.ValueOf(val.Format("3:04PM"))
}

func formatDate(args jet.Arguments) reflect.Value {
	args.RequireNumOfArguments("fmtDate", 1, 1)
	val, ok := args.Get(0).Interface().(time.Time)
	if !ok {
		zero := "0001-01-01 00:00:00 +0000"
		return reflect.ValueOf(zero)
	}

	return reflect.ValueOf(val.Format("YYYY-MM-DD"))
}
