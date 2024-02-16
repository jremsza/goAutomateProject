package main

import "testing"

func TestLinearRegression(t *testing.T) {
	for _, tc := range testCases {
		calculatedSlope, calculatedIntercept, calculatedRSquared := CalculateLinearRegression(x, y)
		if calculatedSlope != tc.ExpectedSlope || calculatedIntercept != tc.ExpectedIntercept || calculatedRSquared != tc.ExpectedRSquared {
			t.Errorf("%s failed: expected slope=%v, intercept=%v, R^2=%v; got slope=%v, intercept=%v, R^2=%v", tc.Name, tc.ExpectedSlope, tc.ExpectedIntercept, tc.ExpectedRSquared, calculatedSlope, calculatedIntercept, calculatedRSquared)
		}
	}
}
