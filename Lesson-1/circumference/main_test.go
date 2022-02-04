package main

import (
	"testing"
	"github.com/GoLevel-1/calculate"
	"github.com/GoLevel-1/errApp"
)

// func TestInputUser(t *testing.T) {
// 	testCases := []struct{
// 		name string
// 		area float64
// 		want float64
// 		wantErr error
// 	}{
// 		{
// 			name: "Test-1 failure, cannot be a negative number",
// 			area: -20,
// 			wantErr: errApp.ErrAreaNegativNumber,
// 		},
// 		{
// 			name: "Test-2 failure, area cannot be zero",
// 			area: 0,
// 			wantErr: errApp.ErrAreaZero,
// 		},
// 		{
// 			name: "Test-3 successful",
// 			area: 6358.5,
// 			want: 6358.5,
// 			wantErr: nil,
// 		},

// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			result, err := userInput()
// 			if err != nil {
// 				if err != tc.wantErr {
// 				t.Errorf("want error %s, got diametr %s", tc.wantErr, err)
// 				}
// 			}
// 			if result != tc.want {
// 				t.Errorf("want diametr %.2f, got diametr %.2f", tc.want, result)
// 			}
// 		})
// 	}
// }

func TestDiametr(t *testing.T) {
	testCases := []struct{
		name string
		area float64
		want float64
	}{
		{
			name: "Test successful",
			area: 113.04,
			want: 12.00,
		},
		{
			name: "Test successful",
			area: 6358.5,
			want: 90.00,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculate.Diametr(tc.area)
			if result != tc.want {
				t.Errorf("want diametr %.2f, got diametr %.2f", tc.want, result)
			}
		})
	}
}

func TestCircumference(t *testing.T) {
	testCases := []struct{
		name string
		area float64
		want float64
	}{
		{
			name: "Test-1 successful",
			area: 113.04,
			want: 37.68,
		},
		{
			name: "Test-2 successful",
			area: 6358.5,
			want: 282.60,
		},
		{
			name: "Test-3 successful",
			area: 28.26,
			want: 18.84,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calculate.Circumference(tc.area)
			if result != tc.want {
				t.Errorf("want diametr %.2f, got diametr %.2f", tc.want, result)
			}
		})
	}
}