package toHex

const alpha = "0123456789abcdef"

func toHex(num int) string {
	var n uint32
	var res []byte

	if num == 0 {
		return "0"
	}

	if num < 0 {
		n = uint32(^(-num) + 1)
	} else {
		n = uint32(num)
	}

	for n > 0 {
		res = append(res, alpha[n%16])
		n = n / 16
	}

	// reverse
	for i := len(res)/2 - 1; i >= 0; i-- {
		opp := len(res) - 1 - i
		res[i], res[opp] = res[opp], res[i]
	}

	return string(res)
}
