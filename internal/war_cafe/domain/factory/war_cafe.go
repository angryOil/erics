package factory

import (
	"erics/internal/war_cafe/domain/model"
)

// concrete factory

type WarCafeFactory interface {
	Create(
		ID int64,
		targetID int64,
		numberOfParticipants int,
	) model.InvestWarCafe
}

var _ WarCafeFactory = (*InvestWarCafeFactory)(nil)

func NewInvestWarCafeFactory() WarCafeFactory {
	return &InvestWarCafeFactory{}
}

type InvestWarCafeFactory struct{}

func (w InvestWarCafeFactory) Create(
	ID int64,
	targetID int64,
	numberOfParticipants int,
) model.InvestWarCafe {

	return model.InvestWarCafe{
		ID:                   ID,
		WarCafeID:            targetID,
		NumberOfParticipants: numberOfParticipants,
		Type:                 model.INVEST,
	}
}
