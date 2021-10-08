package Dailyguoguo

import (
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {

	var res string
	// 对负数进行判断,并将除数和被除数转化为绝对值
	if (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0) {
		res += "-"
	}
	numerator, denominator = int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator)))
	res += strconv.Itoa(numerator / denominator)
	numerator %= denominator
	// 判断是否整除
	if numerator == 0 {
		return res
	}
	res += "."
	// map k 代表除数,v 代表此次除数在 res 的长度
	data := make(map[int]int)
	data[numerator] = len(res)
	for numerator != 0 {
		// 注意每次除完,numerator需要*10
		numerator *= 10
		res += strconv.Itoa(numerator / denominator)
		numerator %= denominator
		if index, ok := data[numerator]; ok {
			res = res[:index] + "(" + res[index:] + ")"
			break
		} else {
			data[numerator] = len(res)
		}
	}
	return res
}
