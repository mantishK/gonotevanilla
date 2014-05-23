package router

import (
	"github.com/mantishK/gonotevanilla/filters"
	"net/http"
)

type router struct {
	routeMethodMap  map[string]map[string]func(http.ResponseWriter, *http.Request)
	routerFilterMap map[string]map[string][]filters.Filterable
}

func New() *router {
	newRouter := new(router)
	newRouter.routeMethodMap = make(map[string]map[string]func(http.ResponseWriter, *http.Request))
	newRouter.routerFilterMap = make(map[string]map[string][]filters.Filterable)
	return newRouter
}

func (r *router) setMethodMap(httpMethod, route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice []filters.Filterable) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	if len(filterSlice) != 0 {
		if r.routerFilterMap[route] == nil {
			r.routerFilterMap[route] = make(map[string][]filters.Filterable)
			if r.routerFilterMap[route][httpMethod] == nil {
				r.routerFilterMap[route][httpMethod] = make([]filters.Filterable, 0, 5)
			}
		}
		for i := range filterSlice {
			r.routerFilterMap[route][httpMethod] = append(r.routerFilterMap[route][httpMethod], filterSlice[i])
		}
	}
	r.routeMethodMap[route][httpMethod] = newControllerMethod
}

func (r *router) Get(route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice ...filters.Filterable) {
	r.setMethodMap("GET", route, newControllerMethod, filterSlice)
}

func (r *router) Post(route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice ...filters.Filterable) {
	r.setMethodMap("POST", route, newControllerMethod, filterSlice)
}

func (r *router) Put(route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice ...filters.Filterable) {
	r.setMethodMap("PUT", route, newControllerMethod, filterSlice)
}

func (r *router) Delete(route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice ...filters.Filterable) {
	r.setMethodMap("DELETE", route, newControllerMethod, filterSlice)
}

func (r *router) Head(route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice ...filters.Filterable) {
	r.setMethodMap("HEAD", route, newControllerMethod, filterSlice)
}

func (r *router) Options(route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice ...filters.Filterable) {
	r.setMethodMap("OPTIONS", route, newControllerMethod, filterSlice)
}

func (r *router) Trace(route string, newControllerMethod func(http.ResponseWriter, *http.Request),
	filterSlice ...filters.Filterable) {
	r.setMethodMap("TRACE", route, newControllerMethod, filterSlice)
}

func (r *router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	method := req.Method
	requestURI := strings.Split(req.RequestURI, "?")[0]
	filterReturnVal := r.executeFilters(method, requestURI, writer, req)
	if filterReturnVal == true {
		r.routeMethodMap[requestURI][method](writer, req)
	}
}

func (r *router) executeFilters(method, requestURI string, writer http.ResponseWriter, req *http.Request) bool {
	filtersSlice := r.routerFilterMap[requestURI][method]
	for i := range filtersSlice {
		returnVal := r.routerFilterMap[requestURI][method][i].Filter(writer, req)
		if returnVal == false {
			return false
		}
	}
	return true
}
