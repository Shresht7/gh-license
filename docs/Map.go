package main

// Map applies the given function to each item in the given slice and returns a new slice with the results.
func Map[T, R any](items []T, fn func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}
