package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

type Settings struct {
	Host string `envconfig:"HOST" default:"0.0.0.0"`
	Port string `envconfig:"PORT" default:"5555"`
	// Domain string `envconfig:"DOMAIN" required:"true"`
	// GlobalUsers means that user@ part is globally unique across all domains
	// WARNING: if you toggle this existing users won't work anymore for safety reasons!
	// GlobalUsers   bool   `envconfig:"GLOBAL_USERS" required:"false" default:false`
	// Secret        string `envconfig:"SECRET" required:"true"`
	// SiteOwnerName string `envconfig:"SITE_OWNER_NAME" required:"true"`
	// SiteOwnerURL  string `envconfig:"SITE_OWNER_URL" required:"true"`
	// SiteName      string `envconfig:"SITE_NAME" required:"true"`

	// ForceMigrate bool   `envconfig:"FORCE_MIGRATE" required:"false" default:false`
	// TorProxyURL  string `envconfig:"TOR_PROXY_URL"`
}

var (
	s      Settings
	router = mux.NewRouter()
	log    = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr})
)

//go:embed index.html
var indexHTML string

//go:embed static
var static embed.FS

func main() {
	// err := envconfig.Process("", &s)
	envconfig.Process("", &s)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("couldn't process envconfig.")
	// }

	router.Path("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, indexHTML)
		},
	)

	router.PathPrefix("/static/").Handler(http.FileServer(http.FS(static)))

	server := &http.Server{
		Handler:      cors.Default().Handler(router),
		Addr:         s.Host + ":" + s.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Debug().Str("addr", server.Addr).Msg("listening")
	server.ListenAndServe()
}
