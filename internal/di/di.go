package di

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/db"
	"github.com/Levor/birthday/internal/db/repositories"
	"github.com/Levor/birthday/internal/http/handlers"
	"github.com/Levor/birthday/internal/http/routes"
	"github.com/Levor/birthday/internal/services"
	"github.com/Levor/birthday/internal/utils"
	"go.uber.org/dig"
	"log"
)

var Container *dig.Container

func init() {
	Container = dig.New()
	providers := []interface{}{
		config.Read,
		routes.API,
		repositories.NewWorkersRepository,
		repositories.NewUserRepository,
		utils.NewJWT,
	}

	providers = append(providers, db.Providers()...)
	providers = append(providers, handlers.HandlerProviders()...)
	providers = append(providers, services.ServiceProviders()...)

	for _, provider := range providers {
		err := Container.Provide(provider)
		if err != nil {
			log.Fatal(err)
		}
	}

}
