package container

import (
	"github.com/MAAARKIN/unico/config"
	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/repository"
	"github.com/MAAARKIN/unico/service"
)

type services struct {
	FeiraService service.FeiraService
}

type Dependency struct {
	Services services
}

func Injector(cfg config.Config) Dependency {
	conn := db.StartDatabase(cfg)

	feiraStore := repository.NewFeiraStorePostgres(conn)
	feiraService := service.NewFeiraService(feiraStore)

	//new repositories/services here and propagate to services cdi

	services := services{
		feiraService,
	}

	return Dependency{
		Services: services,
	}
}
