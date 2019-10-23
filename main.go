package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	// 触发条件
	QualifyingFactorQtyPrefix    = "QFQ"
	QualifyingFactorAmountPrefix = "QF$"
)

func main() {
	fmt.Println("hello Horus")
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now[:len("2006-01-02")])
	fmt.Println(now[len("2006-01-02 "):])

	var intf interface{}

	intf = "horus test string"

	switch a := intf.(type) {
	case string:
		fmt.Println(a)
	default:
		fmt.Println("others")
	}

	qtyExp := regexp.MustCompile(`QFQ\d+`)
	amountExp := regexp.MustCompile(`QF\$\d+`)

	result1 := qtyExp.FindString("QF1Q1")

	result2 := amountExp.FindString("QF$111")

	result3 := qtyExp.FindStringIndex("QFQ1")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	fmt.Println(strings.LastIndex("QFQ5", QualifyingFactorQtyPrefix))

	fmt.Println(strings.LastIndex("QF1", QualifyingFactorQtyPrefix))

	qty := "QFQ5"[len(QualifyingFactorQtyPrefix):]
	fmt.Println(qty)

}
