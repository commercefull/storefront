package actions

import "storefront/pkg/response"

func GetNoFound() (response.Response, error) {
	return response.HttpResponse("non found", 404), nil
}
