package service

import (
	"erics/internal/war_cafe/domain/model"
	"erics/internal/war_cafe/domain/service/warCafe"
	"erics/internal/war_cafe/domain/specification/warcafe_spec"
	"errors"
)

type WarCafeService interface {
	Invest( // 현재 투자자 수와 추가될 인원 수를 받아, 투자한다
		model model.Model,
		targetID int64,
		currentNumber int,
		addNumber int,
	) (model.Model, error)
}

var _ WarCafeService = (*warCafeService)(nil)

type warCafeService struct {
	investService warCafe.InvestService
}

func NewWarCafeService(
	investService warCafe.InvestService,
) WarCafeService {
	return &warCafeService{
		investService: investService,
	}
}

func (w *warCafeService) Invest(
	m model.Model,
	targetID int64,
	currentNumber int,
	addNumber int,
) (model.Model, error) {

	spec := warcafe_spec.NewSpec(m, currentNumber)
	if !spec.IsValid() {
		return m, errors.New("current num is not matched bye")
	}

	// 일단 모델을 받아서 invest 를 해
	m, err := w.investService.Invest(m, targetID, currentNumber, addNumber)

	return m, err
}
