package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Anscombe dataset
	xData := [][]float64{
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	}
	yData := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Convert data to matrices
	x := mat.NewDense(len(xData[0]), len(xData), nil)
	for i := range xData {
		for j := range xData[i] {
			x.Set(j, i, xData[i][j])
		}
	}
	y := mat.NewDense(len(yData), 1, yData)

	// Calculate linear regression coefficients
	var regression mat.VecDense
	err := regression.Solve(x.T(), y)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate predicted values
	var predicted mat.VecDense
	predicted.Mul(x.T(), &regression)

	// Calculate residuals
	var residuals mat.VecDense
	residuals.Sub(y, &predicted)

	// Calculate R-squared
	rSquared := rSquared(&predicted, y)

	// Output results
	fmt.Println("Linear Regression Coefficients:")
	fmt.Printf("Intercept: %.2f\n", regression.AtVec(0))
	fmt.Printf("Slope: %.2f\n", regression.AtVec(1))

	fmt.Println("\nStatistical Summary:")
	fmt.Printf("R-squared: %.4f\n", rSquared)
}

// Calculate R-squared
func rSquared(predicted, actual *mat.VecDense) float64 {
	var rss, tss float64

	// Calculate Residual Sum of Squares (RSS)
	var residuals mat.VecDense
	residuals.Sub(actual, predicted)
	rss = mat.Dot(&residuals, &residuals)

	// Calculate Total Sum of Squares (TSS)
	mean := mat.Sum(actual) / float64(actual.Len())
	var deviations mat.VecDense
	deviations.SubVec(actual, mat.NewVecDense(actual.Len(), []float64{mean}))
	tss = mat.Dot(&deviations, &deviations)

	// Calculate R-squared
	return 1 - rss/tss
}
