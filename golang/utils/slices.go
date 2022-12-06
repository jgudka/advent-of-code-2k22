package utils

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func CreateSetFromList(items []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func Reverse[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func Pop[T any](s []T) (T, []T) {
	lastElement, s := s[len(s)-1], s[:len(s)-1]
	return lastElement, s
}

func PopMultiple[T any](s []T, quantity int) ([]T, []T) {
	lastElements, s := s[len(s)-quantity:], s[:len(s)-quantity]
	return lastElements, s
}
