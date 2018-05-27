package main

import (
	"github.com/gorilla/mux"
	"github.com/beessheet/beessheet/api/api_sheet"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"net"
	"fmt"
	"github.com/beessheet/beessheet/pages/page_home"
)

func main() {

	r := mux.NewRouter()
	apiV1Route := r.PathPrefix("/api/v1/").Subrouter()
	apiBook := apiV1Route.PathPrefix("/books/{bookId}").Subrouter()
	apiBook.HandleFunc("/sheet", api_sheet.HandleList)

	r.HandleFunc("/", page_home.HandleHome)

	http.Handle("/", r)
	serve("127.0.0.1:0", r)
}


func serve(addr string, handler http.Handler) error {
	srv := &http.Server{Addr: addr, Handler: handler}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	port := getPort(ln)
	fmt.Println(port)
	listen := make(chan bool)
	go func() {
		fmt.Println("http://127.0.0.1:" + port + "/")
		<-listen
		open.Run("http://127.0.0.1:" + port + "/")
	}()
	listen <- true
	return srv.Serve(ln.(*net.TCPListener))

}


func getPort(listen net.Listener) string {
	addr := listen.Addr().String()
	_, port, err := net.SplitHostPort(addr)

	if err != nil {
		panic(err)
	}
	return port
}

