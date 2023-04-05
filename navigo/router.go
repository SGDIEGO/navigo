package navigo

import (
	"fmt"
	"net/http"

	"github.com/SGDIEGO/Navigo/navigo/util"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func (rt *Route) MatchRoute(r *http.Request) bool {
	return (rt.Method == r.Method) && (rt.Path == r.URL.Path)
}

type navigo struct {
	path    string
	routes  []Route
	child   *navigo
	brother *navigo
}

func NewMux() *navigo {
	return &navigo{
		path:    "",
		routes:  []Route{},
		child:   nil,
		brother: nil,
	}
}

func (rou *navigo) Search(r *http.Request) *Route {
	for _, route := range rou.routes {
		if route.MatchRoute(r) {
			return &route
		}
	}
	return nil
}

func (rou *navigo) WalkTree(r *http.Request) *Route {
	group := rou

	fmt.Printf("group.path: %v\n", group.path)
	fmt.Printf("r.URL.Path: %v\n", r.URL.Path)

	route := group.Search(r)
	if route != nil {
		return route
	}

	if group.child != nil {
		return group.child.WalkTree(r)
	}

	if group.brother != nil {
		return group.brother.WalkTree(r)
	}

	return nil
}

func (rou *navigo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	route := rou.WalkTree(r)
	if route != nil {
		// Serve handler
		util.LogRoute(route.Path, route.Method)
		route.HandlerFunc.ServeHTTP(w, r)
		return
	}

	util.LogRoute("", "")
	http.NotFound(w, r)
}

func (r *navigo) addGroup(newgroup *navigo) {
	// If doesnt have any router child, so add it as first
	if r.child == nil {
		r.child = newgroup
		return
	}

	// Else add as a brother
	pos := r.child
	for pos.brother != nil {
		pos = pos.brother
	}
	pos.brother = newgroup
}

func (r *navigo) Group(Path string) *navigo {

	newNav := &navigo{
		path:    r.path + Path,
		routes:  []Route{},
		child:   nil,
		brother: nil,
	}

	r.addGroup(newNav)
	return newNav
}

func (r *navigo) GET(Path string, HandlerFunc http.HandlerFunc) {
	r.addRoute("GET", Path, HandlerFunc)
}

func (r *navigo) POST(Path string, HandlerFunc http.HandlerFunc) {
	r.addRoute("Post", Path, HandlerFunc)
}

func (r *navigo) PUT(Path string, HandlerFunc http.HandlerFunc) {
	r.addRoute("PUT", Path, HandlerFunc)
}

func (r *navigo) DELETE(Path string, HandlerFunc http.HandlerFunc) {
	r.addRoute("DELETE", Path, HandlerFunc)
}

func (r *navigo) addRoute(Method string, Path string, HandlerFunc http.HandlerFunc) {
	route := Route{
		Path:        r.path + Path,
		Method:      Method,
		HandlerFunc: HandlerFunc,
	}

	util.AddRoute(route.Path, route.Method)

	r.routes = append(r.routes, route)
}
