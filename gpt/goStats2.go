package main

import (
	"fmt"

	"goAutomated/data"

	"github.com/montanaflynn/stats"
)

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
			fmt.Printf("Error in linear regression for %s: %v\n", setName, err)
			continue
		}

		// Extract coefficients
		intercept, slope := regression[0].X, regression[1].Y

		// Calculate R-squared
		rSquared := stats.RSquared(series, regression)

		// Calculate residuals
		residuals := make([]float64, len(dataSet.Y))
		for i, coord := range series {
			residuals[i] = coord.Y - (intercept + slope*coord.X)
		}

		// Calculate standard error
		stdErr := stats.StdDevP(residuals)

		// Output results
		fmt.Printf("%s - Linear Regression: y = %.2fx + %.2f\n", setName, slope, intercept)
		fmt.Printf("R-squared: %.2f\n", rSquared)
		fmt.Printf("Standard Error: %.2f\n", stdErr)
	}
}
