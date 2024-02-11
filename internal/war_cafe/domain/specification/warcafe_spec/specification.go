package warcafe_spec

import "erics/internal/war_cafe/domain/model"

type Spec interface {
	IsValid() bool
}

type spec struct {
	m          model.Model
	currentNum int
}

var _ Spec = (*spec)(nil)

func (s spec) IsValid() bool {
	return s.currentNum == s.m.Number
}

func NewSpec(m model.Model, currentNum int) Spec {
	return &spec{
		m:          m,
		currentNum: currentNum,
	}
}
