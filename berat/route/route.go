package route

import (
	"net/http"

	"github.com/ctopher7/sirclo/v2/berat/logic"
	"github.com/ctopher7/sirclo/v2/berat/resource"
)

type routeConfig struct {
	path    string
	method  string
	handler func(http.ResponseWriter, *http.Request)
}

func New(res *resource.Resource) {
	routes := routeBuilder{make(map[string]func(http.ResponseWriter, *http.Request)), make(map[string]map[string]func(http.ResponseWriter, *http.Request))}

	handlers := logic.NewHandler(res)

	routes.register(routeConfig{
		path:    "/berat",
		method:  "GET",
		handler: handlers.GetBerat,
	})
	routes.register(routeConfig{
		path:    "/berat",
		method:  "POST",
		handler: handlers.InsertBerat,
	})

	routes.serve()
}

type routeBuilder struct {
	routes   map[string]func(http.ResponseWriter, *http.Request)
	internal map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func (self *routeBuilder) register(routeCfg routeConfig) {
	if _, ok := self.internal[routeCfg.path]; !ok {
		self.internal[routeCfg.path] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	self.internal[routeCfg.path][routeCfg.method] = routeCfg.handler
}

func (self *routeBuilder) serve() {
	for path, v := range self.internal {
		tempV := v
		self.routes[path] = func(w http.ResponseWriter, r *http.Request) {
			for method, fn := range tempV {
				if method == r.Method {
					fn(w, r)
					return
				}
			}
			http.Error(w, "", http.StatusNotFound)
		}
	}
	for k, v := range self.routes {
		http.HandleFunc(k, v)
	}
}
