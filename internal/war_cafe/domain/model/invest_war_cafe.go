package model

import (
	"erics/internal/war_cafe/domain/specification/warcafe_spec/invest_spec"
	"errors"
)

type WarCafe[T InvestWarCafe] interface {
	Apply(
		spec invest_spec.Spec,
	) (Model, error)
}

var _ WarCafe[InvestWarCafe] = (*InvestWarCafe)(nil)

type InvestWarCafe struct {
	ID                   int64
	WarCafeID            int64
	NumberOfParticipants int

	Type TypeWarCafe
}

var (
	NotSatisfiedSpecificationErr = errors.New("not satisfied specification")
)

func (i InvestWarCafe) Apply(
	spec invest_spec.Spec,
) (Model, error) {
	if !spec.IsSatisfy() {
		return Model{}, NotSatisfiedSpecificationErr
	}

	return Model{
		ID:       i.ID,
		Type:     string(i.Type),
		TargetID: i.WarCafeID,
		Number:   i.NumberOfParticipants,
	}, nil
}
