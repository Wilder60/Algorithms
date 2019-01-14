package DynamicProgramming

func FibonacciNumber(N int) int {
	var Map map[int]int = map[int]int{1: 1, 0: 0}
	return dynamicFib(N, Map)
}

func dynamicFib(N int, prevNum map[int]int) int {
	Value, ok := prevNum[N]
	if ok != true {
		prevNum[N] = dynamicFib(N-1, prevNum) + dynamicFib(N-2, prevNum)
		return prevNum[N]
	}
	return Value
}
