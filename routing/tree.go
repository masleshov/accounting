package routing

import "accounting/accounting/service"

// RouteTree represents pair with key - string route and value - service method
type RouteTree map[string]interface{}

// NewRouteTree returns tree with defined routes and service methods
func NewRouteTree() RouteTree {
	tree := make(RouteTree)
	tree.defineRoutes()
	return tree
}

func (tree *RouteTree) defineRoutes() {
	(*tree)["/get_users"] = service.GetUsers
	(*tree)["/create_user"] = service.CreateUser
	(*tree)["/update_user"] = service.UpdateUser
	(*tree)["/delete_user"] = service.DeleteUser
}
