package factors

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"

	"github.com/gorilla/sessions"
)

/*
Factor 6 program
Please, install Memcached in some machine, and export an environment variable called MEMCACHE_HOST
with the value HOST:PORT before starting the program.
*/
var err error

func incrementaContador(sessao *sessions.Session, w http.ResponseWriter, r *http.Request) int {
	qtdeAcesso := sessao.Values["acessos"]
	if qtdeAcesso == nil {
		sessao.Values["acessos"] = 0
	} else {
		sessao.Values["acessos"] = qtdeAcesso.(int) + 1
	}
	sessao.Save(r, w)
	return sessao.Values["acessos"].(int)
}

func sessaoFS(r *http.Request) (sessao *sessions.Session, err error) {
	storeFS := sessions.NewFilesystemStore("", []byte("senha"))
	sessao, err = storeFS.Get(r, "sessaoFS")
	if err != nil {
		sessao, _ = storeFS.New(r, "sessaoFS")
	}
	return sessao, nil
}

func sessaoMC(r *http.Request) (sessao *sessions.Session, err error) {
	var memcacheHost string
	if memcacheHost = os.Getenv("MEMCACHE_HOST"); memcacheHost == "" {
		memcacheHost = "localhost:11211"
	}

	memcacheClient := memcache.New(memcacheHost)
	storeMC := gsm.NewMemcacheStore(memcacheClient, "session_prefix_", []byte("secret-key-goes-here"))
	sessao, _ = storeMC.Get(r, "sessaoMC")
	return sessao, err
}

//Factor6 exports the
func Factor6(urlSession []string, w http.ResponseWriter, r *http.Request) (rc int, msg string) {
	var mensagem string
	var sessao *sessions.Session
	if len(urlSession) > 2 {
		switch urlSession[2] {
		case "fs":
			sessao, err = sessaoFS(r)
		case "mc":
			sessao, err = sessaoMC(r)
		default:
			return http.StatusNotFound, "Favor inserir /factor6/tipo (fs para filesystem ou mc para memcached)"
		}
		if err != nil {
			return http.StatusInternalServerError, fmt.Sprintf("Erro ao criar sessão %s", urlSession[2])
		}
		contador := incrementaContador(sessao, w, r)
		hostname, _ := os.Hostname()
		mensagem = fmt.Sprintf("Olá. Você foi atendido pelo servidor %s \n", hostname)
		mensagem = mensagem + fmt.Sprintf("Seu SESSIONID é %s\n", sessao.ID)
		mensagem = mensagem + fmt.Sprintf("Quantidade de acessos em %s: %d\n", urlSession[2], contador)
		return http.StatusOK, mensagem
	}
	return http.StatusNotFound, "Favor inserir /factor6/tipo (fs para filesystem ou mc para memcached)"
}
