package route

/**
 * Author: jsongo<jsongo@qq.com>
 */

import (
	"strings"

	restful "github.com/emicklei/go-restful"
)

// HandleOption is the option for a router item
type HandleOption struct {
	Handler restful.RouteFunction
	RspType string // mime type
	ReqType string // mime type
}

// RouterOption is just an array of HandleOption
type RouterOption map[string]HandleOption

// Router is the self-define router
type Router struct {
	pathMap map[string]*restful.WebService
	WS      *restful.Container
}

// NewRouter creates a new router
func NewRouter() *Router {
	router := new(Router)
	router.WS = restful.NewContainer()
	return router
}

// AddRouter add a router
func (router *Router) AddRouter(prefix string,
	data RouterOption) *Router {
	var ws *restful.WebService

	if router.pathMap[prefix] == nil {
		ws = new(restful.WebService)
		// ws.Consumes(restful.MIME_JSON)
		// ws.Produces(restful.MIME_JSON)
		ws.Path(prefix)
		// fmt.Println(prefix)
	} else {
		ws = router.pathMap[prefix]
	}
	for pathMethod, option := range data {
		var methodFunc func(string) *restful.RouteBuilder
		arr := strings.Split(pathMethod, "<=")
		path := strings.TrimSpace(arr[0])
		var method string
		if len(arr) < 2 {
			// 默认get方法
			method = "Get"
		} else {
			method = strings.ToUpper(strings.TrimSpace(arr[1]))
		}
		// fmt.Println(method)
		switch method {
		case "POST":
			methodFunc = ws.POST
		case "GET":
			methodFunc = ws.GET
		case "PUT":
			methodFunc = ws.PUT
		default:
			methodFunc = ws.GET
		}
		// ws.Route(ws.HEAD(path).To(item.Handler))
		builder := methodFunc(path).To(option.Handler)
		if len(option.RspType) != 0 {
			// fmt.Println(option.RspType)
			builder.Produces(option.RspType)
		} else {
			builder.Produces(restful.MIME_JSON)
		}
		if len(option.ReqType) != 0 {
			// fmt.Println(option.ReqType)
			builder.Consumes(option.ReqType)
		} else {
			builder.Consumes(restful.MIME_JSON)
		}
		ws.Route(builder)
	}
	// ws.Route(ws.GET("/{name}").To(say.Hello))
	router.WS.Add(ws)
	return router
}
