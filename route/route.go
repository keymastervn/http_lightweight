package route

import (
	"mapi/model"
	"mapi/request"
)

func init() {
	var Routers model.[]*Router
	make(Routers, 5)
}

func Link() model.[]*Router {
	Routers.Add("/hello", request.HelloWorld())
	Routers.Add("/now", request.Timenow())
	Routers.Add("/checkdj", request.CheckDelayedJobs())

	return Routers
}
