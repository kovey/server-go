package router

import "github.com/kovey/server-go/http/context"

type Routers struct {
	get  map[string]*Router
	put  map[string]*Router
	post map[string]*Router
	del  map[string]*Router
}

func NewRouters() *Routers {
	return &Routers{get: make(map[string]*Router), put: make(map[string]*Router), post: make(map[string]*Router), del: make(map[string]*Router)}
}

func (r *Routers) AddGet(router *Router) {
	r.get[router.Path()] = router
}

func (r *Routers) AddPost(router *Router) {
	r.post[router.Path()] = router
}

func (r *Routers) AddPut(router *Router) {
	r.put[router.Path()] = router
}

func (r *Routers) AddDel(router *Router) {
	r.del[router.Path()] = router
}

func (r *Routers) Get(path string) *Router {
	router, ok := r.get[path]
	if !ok {
		return nil
	}

	return router
}

func (r *Routers) Post(path string) *Router {
	router, ok := r.post[path]
	if !ok {
		return nil
	}

	return router
}

func (r *Routers) Put(path string) *Router {
	router, ok := r.put[path]
	if !ok {
		return nil
	}

	return router
}

func (r *Routers) Del(path string) *Router {
	router, ok := r.del[path]
	if !ok {
		return nil
	}

	return router
}

func (r *Routers) Router(path string, intput *context.Input) *Router {
	if intput.Is("post") {
		return r.Post(path)
	}

	if intput.Is("get") {
		return r.Get(path)
	}

	if intput.Is("put") {
		return r.Put(path)
	}

	return r.Del(path)
}
