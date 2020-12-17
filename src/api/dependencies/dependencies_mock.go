package dependencies

import (
	"github.com/agileengine-tasktest/src/api/repository"
	"github.com/agileengine-tasktest/src/api/service"
)

// NewMockContainer returns a container with mock dependencies
func NewMockContainer() *Container {

	var container Container

	container.AccountsHandler = service.NewAccountsHandler()
	container.AccountsRepository = repository.NewAccountsPersistenceHandler()

	return &container
}
