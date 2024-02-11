package warCafe

import (
	"erics/internal/war_cafe/domain/factory"
	"erics/internal/war_cafe/domain/model"
	"erics/internal/war_cafe/domain/specification/warcafe_spec/invest_spec"
)

type InvestService interface {
	Invest(
		model model.Model,
		targetID int64,
		currentNumber int,
		addNumber int,
	) (model.Model, error)
}

var _ InvestService = (*warCafeInvestService)(nil)

func NewInvestService(cafeFactory factory.WarCafeFactory) InvestService {
	return &warCafeInvestService{
		investWarCafeFactory: cafeFactory,
	}
}

type warCafeInvestService struct {
	investWarCafeFactory factory.WarCafeFactory
}

func (w *warCafeInvestService) Invest(
	model model.Model,
	targetID int64,
	currentNumber int,
	addNumber int,
) (model.Model, error) {
	investModel := w.investWarCafeFactory.Create(
		model.ID,
		targetID,
		currentNumber+addNumber,
	)

	_spec := invest_spec.NewAndSpecification(
		invest_spec.MaxSum{
			CurrentCount: currentNumber,
			AddCount:     addNumber,
			MaxSumNumber: 20,
		},
		invest_spec.MinAddNumber{
			AddCount:  addNumber,
			MinNumber: 1,
		},
		invest_spec.MaxAddNumber{
			AddCount:  addNumber,
			MaxNumber: 5,
		},
	)

	return investModel.Apply(_spec)
}
