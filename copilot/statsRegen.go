package main

import (
	"fmt"
	"log"

	"goAutomated/data"

	"github.com/montanaflynn/stats"
)

type Data struct {
	X []float64
	Y []float64
}

func calculateSum(dataSet Data, meanX, meanY float64) (float64, float64) {
	var sumNumerator, sumDenominator float64
	for i := range dataSet.X {
		sumNumerator += (dataSet.X[i] - meanX) * (dataSet.Y[i] - meanY)
		sumDenominator += (dataSet.X[i] - meanX) * (dataSet.X[i] - meanX)
	}
	return sumNumerator, sumDenominator
}

func main() {
	// Retrieve the dataset
	anscombeData := data.AnscombeData()

	// Iterate over the dataset
	for setName, dataSet := range anscombeData {
		var series stats.Series
		for i := range dataSet.X {
			series = append(series, stats.Coordinate{X: dataSet.X[i], Y: dataSet.Y[i]})
		}

		// Calculate linear regression
		regression, err := stats.LinearRegression(series)
		if err != nil {
			log.Fatalf("Error in linear regression for %s: %v\n", setName, err)
		}

		// intercept and slope from the Coordinates
		intercept, slope := regression[0].X, regression[1].Y

		// Manual calculation of linear regression
		meanX, _ := stats.Mean(dataSet.X)
		meanY, _ := stats.Mean(dataSet.Y)
		sumNumerator, sumDenominator := calculateSum(dataSet, meanX, meanY)
		manualSlope := sumNumerator / sumDenominator
		manualIntercept := meanY - manualSlope*meanX

		// Calculate correlation and R squared
		correlation, err := stats.Correlation(dataSet.X, dataSet.Y)
		if err != nil {
			log.Fatalf("Error in calculating correlation for %s: %v\n", setName, err)
		}
		rSquared := correlation * correlation

		fmt.Printf("%s - Linear Regression from package: y = %.2fx + %.2f, R squared: %.2f\n", setName, slope, intercept, rSquared)
		fmt.Printf("%s - Manual Linear Regression: y = %.2fx + %.2f\n", setName, manualSlope, manualIntercept)
	}
}
