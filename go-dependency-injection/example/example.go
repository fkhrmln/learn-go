package example

import "errors"

type ExampleRespository struct {
	Error bool
}

func NewExampleRepository(isError bool) *ExampleRespository {
	return &ExampleRespository{Error: isError}
}

type ExampleService struct {
	ExampleRespository *ExampleRespository
	Error              bool
}

func NewExampleService(exampleRepository *ExampleRespository, isError bool) (*ExampleService, error) {
	if exampleRepository.Error {
		return nil, errors.New("Failed Create Repository")
	}

	return &ExampleService{ExampleRespository: exampleRepository, Error: isError}, nil
}

type ExampleController struct {
	ExampleService *ExampleService
}

func NewExampleController(exampleService *ExampleService) (*ExampleController, error) {
	if exampleService.Error {
		return nil, errors.New("Failed Create Service")
	}

	return &ExampleController{ExampleService: exampleService}, nil
}
