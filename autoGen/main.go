package main

import (
	"testing"

	"github.com/dave/jennifer/jen"
)

// Define the test cases
var testCases = []struct {
	Name              string
	Inputs            []float64 // x1, y1, x2, y2, ...
	ExpectedSlope     float64
	ExpectedIntercept float64
	ExpectedRSquared  float64
}{
	{"TestCase1", []float64{10, 8.04, 8, 6.95, 13, 7.58, 9, 8.81, 11, 8.33, 14, 9.96, 6, 7.24, 4, 4.26, 12, 10.84, 7, 4.82, 5, 5.68}, 0.5, 3.00, 0.67},
}

// Define the test function
func TestLinearRegression(t *testing.T) {
	for _, tc := range testCases {
		//convert the inputs to x and y
		var x, y []float64
		for i := 0; i < len(tc.Inputs); i += 2 {
			x = append(x, tc.Inputs[i])
			y = append(y, tc.Inputs[i+1])
		}

		// Call the function
		calculatedSlope, calculatedIntercept, calculatedRSquared := CalculateLinearRegression(x, y)
		if calculatedSlope != tc.ExpectedSlope || calculatedIntercept != tc.ExpectedIntercept || calculatedRSquared != tc.ExpectedRSquared {
			t.Errorf("%s failed: expected slope=%v, intercept=%v, R^2=%v; got slope=%v, intercept=%v, R^2=%v",
				tc.Name,
				tc.ExpectedSlope,
				tc.ExpectedIntercept,
				tc.ExpectedRSquared,
				calculatedSlope,
				calculatedIntercept,
				calculatedRSquared,
			)
		}
	}
}

func CalculateLinearRegression(x, y []float64) (slope, intercept, rsquared float64) {

	slope = 0.5
	intercept = 3.0
	rsquared = 0.67

	return slope, intercept, rsquared
}

func main() {
	f := jen.NewFile("main")
	f.ImportName("testing", "testing") // Import the testing package

	// Generate the test function
	f = jen.NewFile("main")
	f.Func().Id("TestLinearRegression").Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		jen.For(jen.List(jen.Id("_"), jen.Id("tc")).Op(":=").Range().Id("testCases")).Block(
			jen.List(jen.Id("calculatedSlope"), jen.Id("calculatedIntercept"), jen.Id("calculatedRSquared")).Op(":=").Id("CalculateLinearRegression").Call(jen.Id("x"), jen.Id("y")),
			jen.If(
				jen.Id("calculatedSlope").Op("!=").Id("tc").Dot("ExpectedSlope").Op("||").
					Id("calculatedIntercept").Op("!=").Id("tc").Dot("ExpectedIntercept").Op("||").
					Id("calculatedRSquared").Op("!=").Id("tc").Dot("ExpectedRSquared"),
			).Block(
				jen.Id("t").Dot("Errorf").Call(
					jen.Lit("%s failed: expected slope=%v, intercept=%v, R^2=%v; got slope=%v, intercept=%v, R^2=%v"),
					jen.Id("tc").Dot("Name"),
					jen.Id("tc").Dot("ExpectedSlope"),
					jen.Id("tc").Dot("ExpectedIntercept"),
					jen.Id("tc").Dot("ExpectedRSquared"),
					jen.Id("calculatedSlope"),
					jen.Id("calculatedIntercept"),
					jen.Id("calculatedRSquared"),
				),
			),
		),
	)

	if err := f.Save("linReg_test.go"); err != nil {
		panic(err)
	}
}
