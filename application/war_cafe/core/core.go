package core

import (
	"erics/application/war_cafe/core/request"
	"erics/application/war_cafe/core/response"
)

type UseCase interface {
	Invested(request request.InvestedRequest) (response.InvestedResponse, error)
}
