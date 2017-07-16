package factors

import (
	"fmt"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
	gsm "github.com/bradleypeabody/gorilla-sessions-memcache"

	"github.com/gorilla/sessions"
)

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
	memcacheClient := memcache.New("localhost:11211")
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
		mensagem = fmt.Sprintf("Quantidade de acessos em %s: %d", urlSession[2], contador)
		return http.StatusOK, mensagem
	}
	return http.StatusNotFound, "Favor inserir /factor6/tipo (fs para filesystem ou mc para memcached)"
}
