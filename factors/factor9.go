package factors

import "net/http"

// ServerHealth contains the Server Health for checks
type ServerHealth struct {
	Rc  int
	Msg string
}

// Factor9 functions enables you to disable the server health, and also to check serverhealth
func Factor9(serverhealth **ServerHealth, urlSession []string) (rc int, msg string) {
	var sh *ServerHealth
	sh = *serverhealth
	if len(urlSession) > 2 {
		switch urlSession[2] {
		case "desabilita":
			sh.Rc = 500
			sh.Msg = "Serviço desabilitado"
		case "habilita":
			sh.Rc = 500
			sh.Msg = "Serviço habilitado"
		default:
			return http.StatusNotFound, "Favor inserir /status, /habilita ou /desabilita"
		}
	}
	return sh.Rc, sh.Msg
}
