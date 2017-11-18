package main

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/j0holo/web-framework/parser"
)

func main() {
	config := parser.Config("./config.json")

	srv := newServer()
	defer srv.Close()

	http.HandleFunc("/", indexHandler)

	go srv.ListenAndServeTLS(config.Main.TLSCert, config.Main.TLSKey)
	srv.ListenAndServe()
}

func newServer() http.Server {
	return http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			// Select only save ciphers.
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		},
		TLSNextProto:      make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      25 * time.Second,
		IdleTimeout:       20 * time.Second,
		// Use DefaultMaxHeaderBytes (1 MB).
		MaxHeaderBytes: 0,
		ConnState:      nil,
		ErrorLog:       nil,
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to the index page."))
}
