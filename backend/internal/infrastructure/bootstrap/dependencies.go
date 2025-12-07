package bootstrap

import (
	_ "github.com/JhonCamargo53/prueba-tecnica/docs"
	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/asset"
	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/auth"
	creditRequest "github.com/JhonCamargo53/prueba-tecnica/internal/application/services/credit-request"
	creditStatus "github.com/JhonCamargo53/prueba-tecnica/internal/application/services/credit-status"
	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/customer"
	customerAsset "github.com/JhonCamargo53/prueba-tecnica/internal/application/services/customer-asset"
	documentType "github.com/JhonCamargo53/prueba-tecnica/internal/application/services/document-type"
	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/role"
	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/user"
	"github.com/JhonCamargo53/prueba-tecnica/internal/config"
	adapters "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/ai/credit-risk/adapter/gorm"
	repositories "github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/database/gorm/adapters"
	"github.com/JhonCamargo53/prueba-tecnica/internal/infrastructure/http/handlers"
	"gorm.io/gorm"
)

func InitializeDependencies(db *gorm.DB, cfg *config.Config) {

	/* Assets */
	assetRepo := repositories.NewAssetGormRepository(db)
	assetService := asset.NewAssetService(assetRepo)
	handlers.InitAssetHandler(assetService)

	/* Roles */
	roleRepo := repositories.NewRoleGormRepository(db)
	role.NewRoleService(roleRepo)

	/* Users */
	userRepo := repositories.NewUserGormRepository(db)
	userService := user.NewUserService(userRepo, roleRepo)
	handlers.InitUserHandler(userService)

	/* Auth */
	authService := auth.NewAuthService(userRepo, []byte(cfg.JWTSecretKey))
	handlers.InitAuthHandler(authService)

	/* DocumentTypes */
	documentTypeRepo := repositories.NewDocumentTypeGormRepository(db)
	documentTypeService := documentType.NewDocumentTypeService(documentTypeRepo)
	handlers.InitDocumentTypeHandler(documentTypeService)

	/* CreditStatus */
	creditStatusRepo := repositories.NewCreditStatusGormRepository(db)
	creditStatusService := creditStatus.NewCreditStatusService(creditStatusRepo)
	handlers.InitCreditStatusHandler(creditStatusService)

	/* Customers & CreditRequests repos */
	customerRepo := repositories.NewCustomerGormRepository(db)
	creditRequestRepo := repositories.NewCreditRequestGormRepository(db)

	/* Risk */
	riskEvaluator := adapters.NewRiskEvaluatorAdapter()

	/* CustomerAsset */
	customerAssetRepo := repositories.NewCustomerAssetGormRepository(db)
	customerAssetService := customerAsset.NewCustomerAssetService(
		customerAssetRepo,
		customerRepo,
		assetRepo,
		creditRequestRepo,
		riskEvaluator,
	)
	handlers.InitCustomerAssetHandler(customerAssetService)

	/* CreditRequest */
	creditRequestService := creditRequest.NewCreditRequestService(
		creditRequestRepo,
		customerRepo,
		creditStatusRepo,
		customerAssetRepo,
		riskEvaluator,
	)
	handlers.InitCreditRequestHandler(creditRequestService)

	/* Customers */
	customerService := customer.NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)
	handlers.InitCustomerHandler(customerService)

}
