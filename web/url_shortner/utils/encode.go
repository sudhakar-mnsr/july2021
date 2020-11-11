package base62

import (
	"math"
	"strings"
)

const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const b = 62

// Function encodes the given database ID to a base62 string
func ToBase62(num int) string {
	r := num % b
	res := string(base[r])
	div := num / b
	q := int(math.Floor(float64(div)))
