package route

import (
	"msss_slackapi/model"
	"msss_slackapi/request"
)

func init() {
	var SlackRouters model.[]*Router
	make(SlackRouters, 5)
}

func Link() model.[]*Router {
	SlackRouters.Add("/hello", request.HelloWorld())
	SlackRouters.Add("/now", request.Timenow())
	SlackRouters.Add("/checkdj", request.CheckDelayedJobs())

	return SlackRouters
}
