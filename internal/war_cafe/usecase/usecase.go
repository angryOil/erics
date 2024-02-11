package usecase

import (
	"erics/application/war_cafe/core"
	"erics/application/war_cafe/core/ports/out/repository"
	"erics/application/war_cafe/core/request"
	"erics/application/war_cafe/core/response"
	"erics/internal/war_cafe/domain/service"
)

var _ core.UseCase = (*Implement)(nil)

type Implement struct {
	repo    repository.Repository
	service service.WarCafeService
}

func NewImplement(

	repo repository.Repository,
	service service.WarCafeService,
) core.UseCase {
	return &Implement{
		repo:    repo,
		service: service,
	}
}

func (i *Implement) Invested(request request.InvestedRequest) (response.InvestedResponse, error) {
	ID, warCafeID, current, add := request.ID, request.WarCafeID, request.CurrentNumberOfParticipants, request.AddNumberOfParticipants

	warCafe := i.repo.Load(request.Ctx, ID) // 이미 만들어진 것을 가져온다

	// 도메인 이벤트(투자)를 호출한다
	updatedWarCafe, err := i.service.Invest(warCafe, warCafeID, current, add)
	if err != nil {
		return response.InvestedResponse{}, err
	}

	// 투자된 전쟁 카페를 저장소에 반영한다
	i.repo.Save(request.Ctx, updatedWarCafe) // 순수한 도메인의 이벤트를 저장한다

	return response.InvestedResponse{}, nil
}
