package web

import (
	"net"
	"net/http"
)

func StartServer(addr *string) (err error) {
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		return err
	}
	if err = http.Serve(l, Rout()); err != nil {
		return err
	}
	return nil
}

func Rout() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.Handle("/", page())
	mux.HandleFunc("/create", handlerCreateFile)
	mux.HandleFunc("/getfile", handlerGetFile)
	return mux
}
