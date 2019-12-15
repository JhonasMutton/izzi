package usecase

import "github.com/google/wire"

var UseCaseSet = wire.NewSet(NewInsuranceUseCase, wire.Bind(new(IInsuranceUseCase), new(InsuranceUseCase)),
	NewValidationsUseCase, wire.Bind(new(IValidationsUseCase), new(ValidationsUseCase)))
