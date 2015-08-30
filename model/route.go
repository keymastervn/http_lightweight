package model

type Router struct {
	Url     string
	Request func()
}

func (router *Router) Add(url string, req func()) {
	route.Url = url
	route.TypeRequest = req
}

func (routers *Routers) Add(url string, req func()) {
	var router Router
	router.Add(url, req)
	routers = Append(routers, router)
}
