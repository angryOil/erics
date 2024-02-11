package application

import (
	"erics/application/war_cafe/core"
	"erics/internal/war_cafe/domain/factory"
	"erics/internal/war_cafe/domain/service"
	"erics/internal/war_cafe/domain/service/warCafe"
	"erics/internal/war_cafe/usecase"
)

// container 띄울 준비를 하고
// 논리상 조합한다

// di 에 관한 모든 요소를 조합받아서 조립하고
// application 을 만든다

type App struct {
	useCase core.UseCase
}

func MustProvide() App {
	return App{
		useCase: usecase.NewImplement(
			nil,
			service.NewWarCafeService(
				warCafe.NewInvestService(
					factory.NewInvestWarCafeFactory(),
				),
			),
		),
	}
}
