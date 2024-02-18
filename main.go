package main

import "go-concurrency-patterns/samples"

func main() {
	size := 1500

	samples.GaussianElimination(size, true)
	samples.GaussianElimination(size, false)
}
