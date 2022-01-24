package di

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/db"
	"github.com/Levor/birthday/internal/db/repositories"
	"github.com/Levor/birthday/internal/http/handlers"
	"github.com/Levor/birthday/internal/http/routes"
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
	}

	providers = append(providers, db.Providers()...)
	providers = append(providers, handlers.HandlerProviders()...)

	for _, provider := range providers {
		err := Container.Provide(provider)
		if err != nil {
			log.Fatal(err)
		}
	}

}
