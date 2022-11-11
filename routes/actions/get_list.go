package actions

import (
	"storefront/pkg/response"
	"storefront/template"
)

type GetListData struct {
	Title string
}

func GetList() (response.Response, error) {
	data := GetListData{Title: "Product list"}
	body, err := template.Load("list", data)

	if err != nil {
		return response.InternalServerError(), err
	}

	return response.HttpResponse(body, 200), nil
}
