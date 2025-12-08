package day6

var ops = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

var opNeutral = map[string]int{
	"+": 0,
	"*": 1,
}

func OpGetter(opsLine []string) ([]int, []func(a, b int) int) {
	currentVals, opperations := []int{}, []func(a, b int) int{}
	for _, op := range opsLine {
		f, initial := ops[op], opNeutral[op]

		opperations = append(opperations, f)
		currentVals = append(currentVals, initial)
	}

	return currentVals, opperations
}
