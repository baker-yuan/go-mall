package decimal

import (
	"testing"
)

func TestAddDecimalStrings(t *testing.T) {
	sum, err := DecimalUtils{}.AddDecimal("10.00", "5.50")
	if err != nil {
		t.Errorf("AddDecimalStrings 返回错误: %v", err)
	}
	if sum != "15.50" {
		t.Errorf("AddDecimalStrings 期望得到 %s, 实际得到 %s", "15.50", sum)
	}
}

func TestSubtractDecimalStrings(t *testing.T) {
	diff, err := DecimalUtils{}.SubtractDecimal("10.00", "5.50")
	if err != nil {
		t.Errorf("SubtractDecimalStrings 返回错误: %v", err)
	}
	if diff != "4.50" {
		t.Errorf("SubtractDecimalStrings 期望得到 %s, 实际得到 %s", "4.50", diff)
	}
}

func TestMultiplyDecimalStrings(t *testing.T) {
	product, err := DecimalUtils{}.MultiplyDecimal("10.00", "5.50")
	if err != nil {
		t.Errorf("MultiplyDecimalStrings 返回错误: %v", err)
	}
	if product != "55.00" {
		t.Errorf("MultiplyDecimalStrings 期望得到 %s, 实际得到 %s", "55.00", product)
	}
}

func TestDivideDecimalStrings(t *testing.T) {
	quotient, err := DecimalUtils{}.DivideDecimal("10.00", "5.00")
	if err != nil {
		t.Errorf("DivideDecimalStrings 返回错误: %v", err)
	}
	if quotient != "2.00" {
		t.Errorf("DivideDecimalStrings 期望得到 %s, 实际得到 %s", "2.00", quotient)
	}

	_, err = DecimalUtils{}.DivideDecimal("10.00", "0.00")
	if err == nil {
		t.Errorf("DivideDecimalStrings 期望得到除以零的错误, 但实际没有错误")
	}
}
