//go:build wireinject
// +build wireinject

package example

import "github.com/google/wire"

func InitializeExampleController(isError bool) (*ExampleController, error) {
	wire.Build(NewExampleController, NewExampleService, NewExampleRepository)

	return nil, nil
}
