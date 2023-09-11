package main

type PersonalNumber int

type Number interface {
	~int | ~float64
}

func main() {
	m := map[string]int{"Filipe": 1000, "Marina": 2000}
	m2 := map[string]float64{"Filipe": 100.0, "Marina": 200.20}
	m3 := map[string]PersonalNumber{"Filipe": 100, "Marina": 200}
	println(sum(m))
	println(sum(m2))
	println(sum(m3))
	println(compare(10, 10.20))
}

func sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}
