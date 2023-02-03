package https

import "net/http"

type St struct {
	addr   string
	server *http.Server
}