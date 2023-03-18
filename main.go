package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	triangle := readFile("./data/hard.json")
	fmt.Printf("Triangle path count: %d\n", calTrianglePaths(len(triangle)))
	fmt.Printf("Maximum path sum: %d\n", maxPathSum(triangle))
}

func readFile(filename string) (triangle [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &triangle)
	if err != nil {
		panic(err)
	}
	return triangle
}

func calTrianglePaths(n int) int {
	return int(math.Pow(2, float64(n-1)))
}

func maxPathSum(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := range triangle[i] {
			if triangle[i+1][j] > triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
	}

	return triangle[0][0]
}
