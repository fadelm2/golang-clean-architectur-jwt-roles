package config

import (
	"golang-clean-architecture/internal/delivery/http"
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/delivery/http/route"
	"golang-clean-architecture/internal/gateway/messaging"
	"golang-clean-architecture/internal/repository"
	"golang-clean-architecture/internal/usecase"
	"golang-clean-architecture/internal/util"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *gorm.DB
	App       *fiber.App
	Log       *logrus.Logger
	Validate  *validator.Validate
	Config    *viper.Viper
	Producer  *kafka.Producer
	SecretKey string
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	contactRepository := repository.NewContactRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)
	tokenSecretKey := config.SecretKey

	//setup producer
	userProducer := messaging.NewUserProducer(config.Producer, config.Log)
	contactProducer := messaging.NewContactProducer(config.Producer, config.Log)
	addressProducer := messaging.NewAddressProducer(config.Producer, config.Log)

	tokenUtil := util.NewTokenUtil(tokenSecretKey)

	//setup use cases
	userUseCase := usecase.NewUserCase(config.DB, config.Log, config.Validate, userRepository, userProducer, tokenUtil)
	ContactUseCase := usecase.NewContactUseCase(config.DB, config.Log, config.Validate, contactRepository, contactProducer)
	addressUseCase := usecase.NewAddressUseCase(config.DB, config.Log, config.Validate, contactRepository, addressRepository, addressProducer)

	//setup controller
	userController := http.NewUserController(config.Log, userUseCase)
	ContactController := http.NewContactController(ContactUseCase, config.Log)
	addressController := http.NewAddressController(addressUseCase, config.Log)

	requestLogMiddleware := middleware.NewRequestLogger(userUseCase)
	authMiddlewareAdmin := middleware.NewAuthAdmin(userUseCase, tokenUtil)
	authMiddlewareCustomer := middleware.NewAuthCustomer(userUseCase, tokenUtil)
	authMiddlewareSuperAdmin := middleware.NewAuthSuperAdmin(userUseCase, tokenUtil)
	authMiddlewareDriver := middleware.NewAuthDriver(userUseCase, tokenUtil)

	routeConfig := route.RouteConfig{
		App:                      config.App,
		UserController:           userController,
		ContactController:        ContactController,
		AddressController:        addressController,
		AuthAdminMiddleware:      authMiddlewareAdmin,
		AuthCustomerMiddleware:   authMiddlewareCustomer,
		AuthSuperAdminMiddleware: authMiddlewareSuperAdmin,
		AuthDriverMiddleware:     authMiddlewareDriver,
		RequestLoggerMiddleware:  requestLogMiddleware,
	}
	routeConfig.Setup()
}
