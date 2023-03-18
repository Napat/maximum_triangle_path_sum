package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	type TestCase struct {
		testName        string
		inputFilename   string
		expectPathCount int
		expectMps       int
	}

	testCases := []TestCase{
		{testName: "TestMain ./data/easy1.json", inputFilename: "./data/easy1.json", expectPathCount: 8, expectMps: 237},
		{testName: "TestMain ./data/easy2.json", inputFilename: "./data/easy2.json", expectPathCount: 8, expectMps: 11},
		{testName: "TestMain ./data/easy3.json", inputFilename: "./data/easy3.json", expectPathCount: 8, expectMps: 321},
		{testName: "TestMain ./data/medium1.json", inputFilename: "./data/medium1.json", expectPathCount: 512, expectMps: 691},
		{testName: "TestMain ./data/medium2.json", inputFilename: "./data/medium2.json", expectPathCount: 16384, expectMps: 1074},
		{testName: "TestMain ./data/hard.json", inputFilename: "./data/hard.json", expectPathCount: 9223372036854775807, expectMps: 7273},
	}

	for _, testCase := range testCases {
		triangle := readFile(testCase.inputFilename)
		pathCount := calTrianglePaths(len(triangle))
		mps := maxPathSum(triangle)
		assert.Equal(t, testCase.expectPathCount, pathCount)
		assert.Equal(t, testCase.expectMps, mps)
	}
}
