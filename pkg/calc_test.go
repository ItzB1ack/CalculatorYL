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
			name:       "Простое сложение",
			expression: "2+2",
			want:      4,
			wantErr:   false,
		},
		{
			name:       "Простое вычитание",
			expression: "5-3",
			want:      2,
			wantErr:   false,
		},
		{
			name:       "Простое умножение",
			expression: "4*3",
			want:      12,
			wantErr:   false,
		},
		{
			name:       "Простое деление",
			expression: "8/2",
			want:      4,
			wantErr:   false,
		},
		{
			name:       "Сложное выражение со скобками",
			expression: "(2+3)*4",
			want:      20,
			wantErr:   false,
		},
		{
			name:       "Сложное выражение с несколькими операторами",
			expression: "2+3*4",
			want:      14,
			wantErr:   false,
		},
		{
			name:       "Выражение с вложенными скобками",
			expression: "((2+3)*2)+1",
			want:      11,
			wantErr:   false,
		},
		{
			name:        "Деление на ноль",
			expression:  "1/0",
			wantErr:     true,
			expectedErr: DivideByZero,
		},
		{
			name:        "Некорректные скобки",
			expression:  "((1+2)*3",
			wantErr:     true,
			expectedErr: ErrorInBrackets,
		},
		{
			name:        "Некорректное выражение - двойные операторы",
			expression:  "1++2",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Некорректное выражение - начинается с оператора",
			expression:  "+1+2",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Некорректное выражение - заканчивается оператором",
			expression:  "1+2+",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Пустое выражение",
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
			name:       "Валидное выражение",
			expression: "1+2*3",
			wantErr:   false,
		},
		{
			name:        "Пустое выражение",
			expression:  "",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Начинается с оператора",
			expression:  "+1+2",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Заканчивается оператором",
			expression:  "1+2+",
			wantErr:     true,
			expectedErr: ErrorInExpression,
		},
		{
			name:        "Двойные операторы",
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
			name:       "Валидные скобки",
			expression: "(1+2)*3",
			wantErr:   false,
		},
		{
			name:       "Вложенные скобки",
			expression: "((1+2)*3)",
			wantErr:   false,
		},
		{
			name:        "Незакрытая скобка",
			expression:  "(1+2",
			wantErr:     true,
			expectedErr: ErrorInBrackets,
		},
		{
			name:        "Избыточная закрывающая скобка",
			expression:  "(1+2))",
			wantErr:     true,
			expectedErr: ErrorInBrackets,
		},
		{
			name:        "Несовпадающие скобки",
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
