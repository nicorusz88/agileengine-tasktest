package dependencies

import (
	"github.com/agileengine-tasktest/src/api/repository"
	"github.com/agileengine-tasktest/src/api/service"
)

// Container container for dependency injection
type Container struct {
	AccountsHandler    service.AccountsHandler
	AccountsRepository repository.AccountsPersistenceHandler
}

// GetContainer returns a container with the environment data and dependencies to inject in the components
func GetContainer() *Container {

	var container Container

	container.AccountsHandler = service.NewAccountsHandler()
	container.AccountsRepository = repository.NewAccountsPersistenceHandler()

	return &container
}
