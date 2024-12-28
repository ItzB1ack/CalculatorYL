package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	calc "github.com/ItzB1ack/CalculatorYL/pkg"
)

type Config struct {
	Address string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Address = os.Getenv("PORT")

	if config.Address == "" {
		config.Address = "8080"
	}

	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		request := new(Request)
		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(request)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "Запрос невалиден"}`)
			return
		}

		result, err := calc.Calc(request.Expression)

		if err != nil {
			switch err {
			case calc.ErrBrackets, calc.ErrValues, calc.ErrDivisionByZero, calc.ErrAllowed:
				w.WriteHeader(http.StatusUnprocessableEntity)
				responce := Response{Error: err.Error()}

				json.NewEncoder(w).Encode(responce)
			default:
				w.WriteHeader(http.StatusInternalServerError)
				responce := Response{Error: "Внутренняя ошибка сервера"}

				json.NewEncoder(w).Encode(responce)
			}

		} else {
			responce := Response{Result: result}
			json.NewEncoder(w).Encode(responce)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"error": "Данный метод не поддерживается"}`)
	}

}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Address, nil)
}
