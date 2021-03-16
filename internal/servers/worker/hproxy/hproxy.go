package hproxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/hashicorp/boundary/internal/servers/controller/common"
)

type Services struct {
	repoAuthMethodFn common.AuthTokenRepoFactory
}

func NewService(atRepoFn common.AuthTokenRepoFactory) (Services, error) {
	return Services{
		repoAuthMethodFn: atRepoFn,
	}, nil
}

func (s *Services) HttpProxyHandlerV1() {

	repoAuthMethodFn, _ := s.repoAuthMethodFn()

	director := func(req *http.Request) {
	}

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Printf("%s has requested the url %s from ip %s", r.Header.Get("Boundary-At-Token"), r.Header.Get("Boundary-Forward-For"), r.Header.Get("X-Forwarded-For"))

		pass, err := repoAuthMethodFn.GetAuthMethodPasswordByAuthTokenId(ctx, r.Header.Get("Boundary-At-Token"))
		if err != nil {
			log.Fatal(err)
			return
		}

		r.Header.Add("Authorization", pass)
		origin, _ := url.Parse(r.Header.Get("Boundary-Forward-For"))

		r.Header.Add("X-Forwarded-Host", r.Host)
		r.Header.Add("X-Origin-Host", origin.Host)
		r.URL.Scheme = "http"
		r.URL.Host = origin.Host

		proxy.ServeHTTP(w, r)
	})

	log.Print("Service started sucessfully.")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
