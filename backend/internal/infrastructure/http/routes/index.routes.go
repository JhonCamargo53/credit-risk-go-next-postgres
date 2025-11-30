package routes

import (
	"github.com/gorilla/mux"
)

func RegisterAllRoutes(router *mux.Router) {
	RegisterAssetRoutes(router)
	RegisterAuthRoutes(router)
	RegisterCreditRequestRoutes(router)
	RegisterCreditStatusRoutes(router)
	RegisterCustomerRoutes(router)
	RegisterDocumentTypeRoutes(router)
	RegisterUserRoutes(router)
	RegisterHealthRoutes(router)
	RegisterCustomerAssetRoutes(router)
	RegisterMetricRoutes(router)
}
