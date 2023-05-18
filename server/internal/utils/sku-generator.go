package utils

import (
	"server/internal/models"
	"strconv"
)

func hash(str string) string {
	x, y := 0, 0
	for _, c := range str {
		x = x ^ int(c)
		y = y * int(c) % 256
	}
	return strconv.Itoa(x) + strconv.Itoa(y)
}

// 3 first letters of name + hash of name + "-" + brand + "-" + color + "-" + size
// hash of name in case of same prefix of name
func SkuGenerator(p *models.Products) string {
	name := p.Name
	brand := p.Brand
	color := p.Color
	size := p.Size

	n := len(name)
	if n > 3 {
		n = 3
	}

	for name[n-1] == ' ' {
		n--
	}

	sku := name[:n] + hash(name) + "-" + brand + "-" + color + "-" + size

	return sku
}
