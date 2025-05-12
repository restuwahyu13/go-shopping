package api

import (
	"compress/zlib"
	"os"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	cinf "restuwahyu13/shopping-cart/internal/domain/interface/common"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	popt "restuwahyu13/shopping-cart/internal/domain/output/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/handler"
	"restuwahyu13/shopping-cart/internal/infrastructure/route"
	"restuwahyu13/shopping-cart/internal/infrastructure/service"
	"restuwahyu13/shopping-cart/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/oxequa/grace"
	"github.com/unrolled/secure"
	"github.com/uptrace/bun"
)

type Api struct {
	ENV    *cdto.Environtment
	ROUTER *chi.Mux
	DB     *bun.DB
	RDS    pinf.IRedis
}

func NewApi(options Api) cinf.IApi {
	return Api{
		ENV:    options.ENV,
		ROUTER: options.ROUTER,
		DB:     options.DB,
		RDS:    options.RDS,
	}
}

func (i Api) Middleware() {
	if i.ENV.APP.ENV != cons.PROD {
		i.ROUTER.Use(middleware.Logger)
	}

	i.ROUTER.Use(middleware.Recoverer)
	i.ROUTER.Use(middleware.RealIP)
	i.ROUTER.Use(middleware.CleanPath)
	i.ROUTER.Use(middleware.NoCache)
	i.ROUTER.Use(middleware.GetHead)
	i.ROUTER.Use(middleware.Compress(zlib.BestCompression))
	i.ROUTER.Use(middleware.AllowContentType("application/json"))
	i.ROUTER.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		MaxAge:             900,
	}))
	i.ROUTER.Use(secure.New(secure.Options{
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		STSIncludeSubdomains: true,
		STSPreload:           true,
		STSSeconds:           900,
	}).Handler)
}

func (i Api) Router() {
	serviceAuth := service.NewAuthService(service.AuthService{ENV: i.ENV, DB: i.DB, RDS: i.RDS})
	usecaseAuth := usecase.NewAuthUsecase(usecase.AuthUsecase{SERVICE: serviceAuth})
	handlerAuth := handler.NewAuthHandler(handler.AuthHandler{USECASE: usecaseAuth})
	route.NewAuthRoute(route.AuthRoute{ROUTER: i.ROUTER, HANDLER: handlerAuth})

	serviceShopping := service.NewShoppingService(service.ShoppingService{ENV: i.ENV, DB: i.DB, RDS: i.RDS})
	usecaseShopping := usecase.NewShoppingUsecase(usecase.ShoppingUsecase{SERVICE: serviceShopping})
	handlerShopping := handler.NewShoppingHandler(handler.ShoppingHandler{USECASE: usecaseShopping})
	route.NewShoppingRoute(route.ShoppingRoute{ROUTER: i.ROUTER, HANDLER: handlerShopping, ENV: i.ENV, REDIS: i.RDS})

	servicePayment := service.NewPaymentService(service.PaymentService{ENV: i.ENV, DB: i.DB, RDS: i.RDS})
	usecasePayment := usecase.NewPaymentUsecase(usecase.PaymentUsecase{SERVICE: servicePayment})
	handlerPayment := handler.NewPaymentHandler(handler.PaymentHandler{USECASE: usecasePayment})
	route.NewPaymentRoute(route.PaymentRoute{ROUTER: i.ROUTER, HANDLER: handlerPayment, ENV: i.ENV, REDIS: i.RDS})
}

func (i Api) Listener() {
	err := pkg.Graceful(func() *popt.GracefulConfig {
		return &popt.GracefulConfig{HANDLER: i.ROUTER, ENV: i.ENV}
	})

	recover := grace.Recover(&err)
	recover.Stack()

	if err != nil {
		pkg.Logrus("fatal", err)
		os.Exit(1)
	}
}
