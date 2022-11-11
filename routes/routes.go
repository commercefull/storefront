package routes

import (
	"log"
	"storefront/pkg/request"
	"storefront/pkg/response"
	"storefront/pkg/router"
	"storefront/routes/actions"
)

func Handle(r request.Request) response.Response {

	routes := map[string]interface{}{
		"GET /":               actions.GetIndex,
		"GET /list/:slug":     actions.GetCategory,
		"GET /categories":     actions.GetCategory,
		"GET /category/:slug": actions.GetCategory,
	}

	path, _ := r.RequestPath()
	res, err := router.Router(path, routes)

	if err != nil {
		log.Panic(err)
	}

	return res
}
