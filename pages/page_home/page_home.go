package page_home

import "net/http"

func HandleHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello"))

}
