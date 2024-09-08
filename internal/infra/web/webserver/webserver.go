package webserver

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc // Map de métodos para handlers
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

// AddHandler agora aceita o método HTTP como parâmetro
func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	if s.Handlers[path] == nil {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[path][method] = handler
}

// loop through the handlers and add them to the router
// register middleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, methodHandlers := range s.Handlers {
		for method, handler := range methodHandlers {
			if path[0] != '/' {
				panic("chi: routing pattern must begin with '/' in '" + path + "'")
			}

			switch method {
			case "GET":
				s.Router.Get(path, handler)
			case "POST":
				s.Router.Post(path, handler)
			case "PUT":
				s.Router.Put(path, handler)
			case "DELETE":
				s.Router.Delete(path, handler)
			// Adicione outros métodos conforme necessário
			default:
				s.Router.Handle(path, handler)
			}
		}
	}

	// Adicione um prefixo ':' à porta
	addr := ":" + s.WebServerPort

	// Verifique se há erros ao iniciar o servidor
	log.Printf("Starting web server on port %s", s.WebServerPort)
	if err := http.ListenAndServe(addr, s.Router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
