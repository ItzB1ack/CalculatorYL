package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		want        float64
		wantErr     bool
		expectedErr string
	}{
		{
			name:       "Simple addition",
			expression: "2+2",
			want:      4,
			wantErr:   false,
		},
		{
			name:       "Simple subtraction",
			expression: "5-3",
			want:      2,
			wantErr:   false,
		},
		{
			name:       "Simple multiplication",
			expression: "4*3",
			want:      12,
			wantErr:   false,
		},
		{
			name:       "Simple division",
			expression: "8/2",
			want:      4,
			wantErr:   false,
		},
		{
			name:       "Complex expression with brackets",
			expression: "(2+3)*4",
			want:      20,
			wantErr:   false,
		},
		{
			name:       "Complex expression with multiple operators",
			expression: "2+3*4",
			want:      14,
			wantErr:   false,
		},
		{
			name:       "Expression with nested brackets",
			expression: "((2+3)*2)+1",
			want:      11,
			wantErr:   false,
		},
		{
			name:        "Division by zero",
			expression:  "1/0",
			wantErr:     true,
			expectedErr: DivideByZero,
		},
		{
			name:        "Invalid brackets",
			expression:  "((1+2)*3",
			wantErr:     true,
			expectedErr: ErrorInBrackets,
		},
		{
			name:        "Invalid expression - double operators",
			expression:  "1++2",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Invalid expression - starts with operator",
			expression:  "+1+2",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Invalid expression - ends with operator",
			expression:  "1+2+",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Empty expression",
			expression:  "",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calc(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("Calc() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateExpression(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		wantErr     bool
		expectedErr string
	}{
		{
			name:       "Valid expression",
			expression: "1+2*3",
			wantErr:   false,
		},
		{
			name:        "Empty expression",
			expression:  "",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Starts with operator",
			expression:  "+1+2",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Ends with operator",
			expression:  "1+2+",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Double operators",
			expression:  "1++2",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateExpression(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("validateExpression() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidateBrackets(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		wantErr     bool
		expectedErr string
	}{
		{
			name:       "Valid brackets",
			expression: "(1+2)*3",
			wantErr:   false,
		},
		{
			name:       "Nested brackets",
			expression: "((1+2)*3)",
			wantErr:   false,
		},
		{
			name:        "Unclosed bracket",
			expression:  "(1+2",
			wantErr:     true,
			expectedErr: ErrorInBrackets,
		},
		{
			name:        "Extra closing bracket",
			expression:  "(1+2))",
			wantErr:     true,
			expectedErr: ErrorInBrackets,
		},
		{
			name:        "Mismatched brackets",
			expression:  ")(1+2)",
			wantErr:     true,
			expectedErr: ErrorInBrackets,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBrackets(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateBrackets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("validateBrackets() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
