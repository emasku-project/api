package failure

import "net/http"

type Failure struct {
	Message string `json:"message" validate:"required"`
}

type App struct {
	Code    int
	Message string
	Err     error
}

func New(code int, message string, err error) *App {
	return &App{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewUnauthorized() *App {
	return &App{
		Code:    http.StatusUnauthorized,
		Message: "Anda tidak terautentikasi!",
	}
}

func NewForbidden() *App {
	return &App{
		Code:    http.StatusForbidden,
		Message: "Anda tidak memiliki akses!",
	}
}

func NewInternal(err error) *App {
	return &App{
		Code:    http.StatusInternalServerError,
		Message: "Terjadi kesalahan tak terduga pada server.",
		Err:     err,
	}
}
