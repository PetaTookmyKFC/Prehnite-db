package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	routes routes
}
type (
	routes map[string]route
	route  map[string]http.HandlerFunc
)

func StartServer(port string) *Server {
	ser := Server{
		port:   port,
		routes: make(map[string]route),
	}
	fmt.Println("Starting Server")
	// Start Server on port
	go http.ListenAndServe(port, nil)

	fmt.Println("Started Listener")

	return &ser
}

func (rl routes) addRoute(path string, method string, handler http.HandlerFunc) {
	r := route{}
	if _, ok := rl[path]; ok {
		r = rl[path]
	}

	r[method] = handler
	rl[path] = r
}

func (ser *Server) Get(path string, handler http.HandlerFunc) {
	ser.routes.addRoute(path, "GET", handler)
	http.HandleFunc(path, ser.TriggerRoute)
}

func (ser *Server) Post(path string, handler http.HandlerFunc) {
	ser.routes.addRoute(path, "POST", handler)
	http.HandleFunc(path, ser.TriggerRoute)
}

func (ser *Server) TriggerRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get R -Path: %s", r.URL.Path)
	fmt.Printf("Get R -Method: %s", r.Method)
	ser.routes[r.URL.Path][r.Method](w, r)
}
