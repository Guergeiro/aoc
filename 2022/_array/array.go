package array

func Reduce[T, M any](slice []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, value := range slice {
		acc = f(acc, value)
	}
	return acc
}

func Contains[T comparable](slice []T, neddle T) bool {
	for _, value := range slice {
		if value == neddle {
			return true
		}
	}
	return false
}

func ContainsInAll[T comparable](slices [][]T, neddle T) bool {
	for _, slice := range slices {
		if Contains(slice, neddle) == false {
			return false
		}
	}

	return true
}
