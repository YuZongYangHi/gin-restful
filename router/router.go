package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ViewsBehavior interface {
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	ExtendsAction() map[string]gin.HandlerFunc
}

type RouterController struct {
	Views ViewsBehavior
	ctx   *gin.Context
}

func (router *RouterController) PostAndListMathEq(cm string, queryMath bool, ctx *gin.Context) bool {
	if queryMath {
		return cm == ctx.Request.Method && !router.IsQuery(ctx) && router.RowEqReqURI(ctx) && !router.FormIsEmpty(ctx)
	}
	return cm == ctx.Request.Method && router.RowEqReqURI(ctx)

}

func (router *RouterController) RouterPostVerify(method string, ctx *gin.Context) bool {
	if router.PostAndListMathEq(http.MethodPost, true, ctx) {
		return true
	}
	return false
}

func (router *RouterController) RouterListAndRetrieveVerify(method string, ctx *gin.Context) {
	if method == "" {
		method = http.MethodGet
	}

	if router.RouterListVerify(method, ctx) {
		router.Views.List(ctx)
	} else if router.RouterRetrieveVerify(method, ctx) {
		router.Views.Retrieve(ctx)
	}
}

func (router *RouterController) RouterDeleteVerify(method string, ctx *gin.Context) bool {
	if method == ctx.Request.Method && router.RegexResourceId(ctx) && !router.IsQuery(ctx) {
		return true
	}
	return false
}

func (router *RouterController) RouterListVerify(method string, ctx *gin.Context) bool {
	if router.PostAndListMathEq(http.MethodGet, false, ctx) {
		return true
	}
	return false
}

func (router *RouterController) RouterRetrieveVerify(method string, ctx *gin.Context) bool {
	if ctx.Request.Method == method && router.RegexResourceId(ctx) {
		return true
	}
	return false
}

func (router *RouterController) SourceURI(ctx *gin.Context) string {
	return ctx.FullPath()
}

func (router *RouterController) ManyRouterAlias(ctx *gin.Context) []string {
	return strings.Split(router.SourceURI(ctx), "*")
}

func (router *RouterController) URISplit(ctx *gin.Context) []string {
	uri := router.SourceURI(ctx)
	return strings.Split(uri, "*")
}

func (router *RouterController) URIRowQuery(ctx *gin.Context) string {
	return ctx.Request.URL.RawQuery
}

func (router *RouterController) IsQuery(ctx *gin.Context) bool {
	if router.URIRowQuery(ctx) != "" {
		return true
	}
	return false
}

func (router *RouterController) RowEqReqURI(ctx *gin.Context) bool {
	return router.URISplit(ctx)[0] == strings.Split(ctx.Request.RequestURI, "?")[0]
}

func (router *RouterController) GetResource(ctx *gin.Context) string {
	path := router.ManyRouterAlias(ctx)
	return ctx.Param(path[1])
}

func (router *RouterController) IsEmptyResource(ctx *gin.Context) bool {
	source := router.GetResource(ctx)
	if source == "/" {
		return true
	}
	return false
}

func (router *RouterController) AccessResourceFilter(source []string) bool {
	if len(source) == 2 {
		return true
	} else if len(source) == 3 {
		if source[2] == "" {
			return true
		}
		return false
	}
	return false
}

func (router *RouterController) GetResourceId(ctx *gin.Context) string {
	source := router.GetResource(ctx)
	splitArray := strings.Split(source, "/")

	if len(splitArray) >= 2 && router.AccessResourceFilter(splitArray) {
		return splitArray[1]
	}
	return ""
}

func (router *RouterController) RegexResourceId(ctx *gin.Context) bool {

	if !router.IsEmptyResource(ctx) && router.GetResourceId(ctx) != "" {
		return true
	}
	return false
}

func (router *RouterController) FormIsEmpty(ctx *gin.Context) bool {
	buf := make([]byte, 1024)
	n, _ := ctx.Request.Body.Read(buf)
	if string(buf[0:n]) == "" {
		return true
	}
	return false
}

func (router *RouterController) RouterUpdateVerify(method string, ctx *gin.Context) bool {
	if ctx.Request.Method == method && router.RegexResourceId(ctx) && !router.FormIsEmpty(ctx) {
		return true
	}
	return false
}

func (router *RouterController) Dispath(ctx *gin.Context) {
	router.ctx = ctx
	switch {
	case router.RouterPostVerify(ctx.Request.Method, ctx):
		router.Views.Create(ctx)
	case router.RouterDeleteVerify(http.MethodDelete, ctx):
		router.Views.Delete(ctx)
	case ctx.Request.Method == http.MethodGet:
		if router.RouterListVerify(http.MethodGet, ctx) {
			router.Views.List(ctx)
		} else if router.RouterRetrieveVerify(http.MethodGet, ctx) {
			router.Views.Retrieve(ctx)
		} else {
			router.Transition(ctx)
		}
	case router.RouterUpdateVerify(http.MethodPut, ctx):
		router.Views.Update(ctx)
	default:
		router.Transition(ctx)
	}
}

func (router *RouterController) Transition(ctx *gin.Context) {

	sets := router.Views.ExtendsAction()
	relatively := router.URISplit(router.ctx)[0]
	sg := false

	for path, f := range sets {
		if router.ctx.Request.RequestURI == relatively+path {
			f(router.ctx)
			sg = true
		}
	}

	if !sg {
		ctx.JSON(404, "url not found")
	}
}

func (router *RouterController) Register(v ViewsBehavior) func(ctx *gin.Context) {
	router.Views = v
	return router.Dispath
}

