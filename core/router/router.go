package router

import (
	"fmt"
	"net/http"
)

type router struct {
	routeMethodMap map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func New() *router {
	newRouter := new(router)
	newRouter.routeMethodMap = make(map[string]map[string]func(http.ResponseWriter, *http.Request))
	return newRouter
}

func (r *router) Get(route string, newControllerMethod func(http.ResponseWriter, *http.Request)) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.routeMethodMap[route]["GET"] = newControllerMethod
}

func (r *router) Post(route string, newControllerMethod func(http.ResponseWriter, *http.Request)) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.routeMethodMap[route]["POST"] = newControllerMethod
}

func (r *router) Put(route string, newControllerMethod func(http.ResponseWriter, *http.Request)) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.routeMethodMap[route]["PUT"] = newControllerMethod
}

func (r *router) Delete(route string, newControllerMethod func(http.ResponseWriter, *http.Request)) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.routeMethodMap[route]["DELETE"] = newControllerMethod
}

func (r *router) Head(route string, newControllerMethod func(http.ResponseWriter, *http.Request)) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.routeMethodMap[route]["HEAD"] = newControllerMethod
}

func (r *router) Options(route string, newControllerMethod func(http.ResponseWriter, *http.Request)) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.routeMethodMap[route]["OPTIONS"] = newControllerMethod
}

func (r *router) trace(route string, newControllerMethod func(http.ResponseWriter, *http.Request)) {
	if r.routeMethodMap[route] == nil {
		r.routeMethodMap[route] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	r.routeMethodMap[route]["TRACE"] = newControllerMethod
}

func (r *router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	method := req.Method
	requestURI := req.RequestURI
	fmt.Println("method = " + method + ",req url = " + requestURI)
	r.routeMethodMap[requestURI][method](writer, req)
}
