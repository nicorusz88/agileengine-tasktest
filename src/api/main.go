package main

import (
	"github.com/agileengine-tasktest/src/api/builder"
	"github.com/agileengine-tasktest/src/api/dependencies"
)

func main() {
	environmentBuilder := builder.Build()

	container := dependencies.GetContainer()
	environmentBuilder.ConfigureRouter(container)
	environmentBuilder.Run()
}

