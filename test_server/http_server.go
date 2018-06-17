package http_test

import (
	"net/http"
)

func testhandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", testhandler)

}
