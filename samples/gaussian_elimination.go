package samples

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func GaussianElimination(size int, concurrent bool) {
	var durationInMilliseconds int64
	var msg string

	if concurrent {
		msg = "concurrent algorithm"
		durationInMilliseconds = concurrentGaussianElimination(size)
	} else {
		msg = "sequential algorithm"
		durationInMilliseconds = sequentialGaussianElimination(size)
	}

	fmt.Println(msg, durationInMilliseconds, "ms")
}

func sequentialGaussianElimination(size int) int64 {
	augmentedMatrix := randomMatrix(size)

	start := time.Now()
	for i := 0; i < len(augmentedMatrix[0]); i++ {
		for j := i + 1; j < len(augmentedMatrix); j++ {
			factor := GetFactor(augmentedMatrix[j][i], augmentedMatrix[i][i])
			scaledRow := scaleRow(augmentedMatrix[i], factor)
			subtractedRow := subtractRows(augmentedMatrix[j], scaledRow)
			swapRows(&subtractedRow, &augmentedMatrix[j])
		}
	}
	end := time.Now()
	return end.Sub(start).Milliseconds()
}

func concurrentGaussianElimination(size int) int64 {
	augmentedMatrix := randomMatrix(size)

	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < len(augmentedMatrix[0]); i++ {
		wg.Wait()
		for j := i + 1; j < len(augmentedMatrix); j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()

				factor := GetFactor(augmentedMatrix[j][i], augmentedMatrix[i][i])
				scaledRow := scaleRow(augmentedMatrix[i], factor)
				subtractedRow := subtractRows(augmentedMatrix[j], scaledRow)
				swapRows(&subtractedRow, &augmentedMatrix[j])
			}(i, j)
		}
	}
	end := time.Now()
	return end.Sub(start).Milliseconds()
}

func scaleRow(row []float64, factor float64) []float64 {
	newRow := make([]float64, len(row))

	for i := 0; i < len(row); i++ {
		newRow[i] = row[i] * factor
	}

	return newRow
}

func subtractRows(rowA, rowB []float64) []float64 {
	newRow := make([]float64, len(rowA))

	for i := 0; i < len(rowA); i++ {
		newRow[i] = rowA[i] - rowB[i]
	}

	return newRow
}

func swapRows(rowA, rowB *[]float64) {
	temp := *rowA
	*rowA = *rowB
	*rowB = temp
}

func GetFactor(value, pivot float64) float64 {
	return value / pivot
}

func printMatrix(m [][]float64) {
	for _, row := range m {
		fmt.Println(row)
	}
}

func randomMatrix(size int) [][]float64 {
	m := make([][]float64, size)
	//v := make([]float64, size)

	for i := 0; i < size; i++ {
		m[i] = make([]float64, size+1)
		for j := 0; j < size+1; j++ {
			m[i][j] = roundFloat64(rand.Float64()*10, 2)
		}
	}

	return m
}

func roundFloat64(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
