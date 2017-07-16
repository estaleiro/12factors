package factors

import (
	"net/http"
	"os"
)

// Factor3 gets the system Environment Variable 'MENSAGEM' and prints it into http response.
func Factor3() (int, string) {
	mensagem := os.Getenv("MENSAGEM")
	if mensagem == "" {
		return http.StatusInternalServerError, "Erro - Variável MENSAGEM não definida"
	}
	return http.StatusOK, mensagem
}
