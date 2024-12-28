package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	calc "github.com/ItzB1ack/CalculatorYL/pkg"
)

func TestConfigFromEnv(t *testing.T) {
	tests := []struct {
		name     string
		envPort  string
		expected string
	}{
		{
			name:     "Порт по умолчанию при пустом env",
			envPort:  "",
			expected: "8080",
		},
		{
			name:     "Пользовательский порт из env",
			envPort:  "3000",
			expected: "3000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envPort != "" {
				os.Setenv("PORT", tt.envPort)
				defer os.Unsetenv("PORT")
			} else {
				os.Unsetenv("PORT")
			}

			config := ConfigFromEnv()
			if config.Address != tt.expected {
				t.Errorf("ConfigFromEnv() = %v, want %v", config.Address, tt.expected)
			}
		})
	}
}

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		requestBody    interface{}
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "Корректное выражение",
			method: "POST",
			requestBody: Request{
				Expression: "2+2",
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"result": float64(4),
			},
		},
		{
			name:   "Некорректное выражение - деление на ноль",
			method: "POST",
			requestBody: Request{
				Expression: "1/0",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"error": calc.DivideByZero,
			},
		},
		{
			name:   "Некорректное выражение - скобки",
			method: "POST",
			requestBody: Request{
				Expression: "((1+2)",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"error": calc.ErrorInBrackets,
			},
		},
		{
			name:   "Некорректный JSON запрос",
			method: "POST",
			requestBody: map[string]interface{}{
				"invalid_field": "value",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Запрос невалиден",
			},
		},
		{
			name:           "Некорректный HTTP метод",
			method:         "GET",
			requestBody:    nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody: map[string]interface{}{
				"error": "Данный метод не поддерживается",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			var err error

			if tt.requestBody != nil {
				reqBody, err = json.Marshal(tt.requestBody)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req := httptest.NewRequest(tt.method, "/", bytes.NewBuffer(reqBody))
			w := httptest.NewRecorder()

			CalcHandler(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("CalcHandler() status = %v, want %v", w.Code, tt.expectedStatus)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for key, expectedValue := range tt.expectedBody {
				if actualValue, ok := response[key]; !ok {
					t.Errorf("Response missing expected key %q", key)
				} else if actualValue != expectedValue {
					t.Errorf("Response[%q] = %v, want %v", key, actualValue, expectedValue)
				}
			}
		})
	}
}

func TestApplication_RunServer(t *testing.T) {
	os.Setenv("PORT", "0")
	defer os.Unsetenv("PORT")

	app := New()
	
	go func() {
		err := app.RunServer()
		if err != nil && err.Error() != "http: Server closed" {
			t.Errorf("RunServer() error = %v", err)
		}
	}()
}
