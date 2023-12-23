package decimal

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// DecimalUtils 浮点数计算
type DecimalUtils struct {
}

// AddDecimal 接受两个表示十进制数的字符串参数，返回它们的和。
// 返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) AddDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	result := decA.Add(decB)
	return result.StringFixed(2), nil
}

// SubtractDecimal 接受两个表示十进制数的字符串参数，返回第一个参数减去第二个参数的结果。
// 返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) SubtractDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	result := decA.Sub(decB)
	return result.StringFixed(2), nil
}

// MultiplyDecimal 接受两个表示十进制数的字符串参数，返回它们的乘积。
// 返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) MultiplyDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	result := decA.Mul(decB)
	return result.StringFixed(2), nil
}

// DivideDecimal 接受两个表示十进制数的字符串参数，返回第一个参数除以第二个参数的结果。
// 如果第二个参数为零，则返回错误。返回结果为保留两位小数的字符串，如果有错误则返回错误。
func (c DecimalUtils) DivideDecimal(a, b string) (string, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	if decB.IsZero() {
		return "", fmt.Errorf("除数不能为零")
	}
	result := decA.Div(decB)
	return result.StringFixed(2), nil
}

// CompareDecimal 比较两个表示十进制数的字符串参数。
// 如果 a > b 返回 1，如果 a < b 返回 -1，如果 a == b 返回 0。
func (c DecimalUtils) CompareDecimal(a, b string) (int, error) {
	decA, err := decimal.NewFromString(a)
	if err != nil {
		return 0, err
	}
	decB, err := decimal.NewFromString(b)
	if err != nil {
		return 0, err
	}
	return decA.Cmp(decB), nil
}

// SubtractPrices 接受两个字符串类型的价格，执行减法操作，并返回结果字符串
//func SubtractPrices(originalPriceStr, promotionPriceStr string) (string, error) {
//	originalPrice, _, err := big.NewFloat(0).Parse(originalPriceStr, 10)
//	if err != nil {
//		return "", fmt.Errorf("error parsing original price: %v", err)
//	}
//
//	promotionPrice, _, err := big.NewFloat(0).Parse(promotionPriceStr, 10)
//	if err != nil {
//		return "", fmt.Errorf("error parsing promotion price: %v", err)
//	}
//
//	finalPrice := new(big.Float).Sub(originalPrice, promotionPrice)
//	return finalPrice.Text('f', 2), nil // 返回结果，保留两位小数
//}
