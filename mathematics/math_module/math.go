package mathematics

func Add(adds ...int) int {
	var result int = adds[0]
	for a := 1; a < len(adds); a++ {
		result += adds[a]
	}
	return result
}

func Sub(subs ...int) int {
	var result int = subs[0]
	for a := 1; a < len(subs); a++ {
		result -= subs[a]
	}
	return result
}

func Mul(muls ...int) int {
	var result = muls[0]
	for a := 1; a < len(muls); a++ {
		result *= muls[a]
	}
	return result
}

func Div(divs ...int) float32 {
	var result float32 = float32(divs[0])
	for a := 1; a < len(divs); a++ {
		result /= float32(divs[a])
	}
	return result
}