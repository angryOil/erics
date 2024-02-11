package usecase

import (
	"context"
	"erics/application/war_cafe/core/ports/out/repository"
	"erics/application/war_cafe/core/request"
	"erics/application/war_cafe/core/response"
	"erics/internal/war_cafe/domain/factory"
	"erics/internal/war_cafe/domain/model"
	"erics/internal/war_cafe/domain/service"
	"erics/internal/war_cafe/domain/service/warCafe"
	"reflect"
	"testing"
)

var _ repository.Repository = (*mockRepo)(nil)

type mockRepo struct {
}

func (m mockRepo) Save(ctx context.Context, model model.Model) {}

func (m mockRepo) Load(ctx context.Context, ID int64) model.Model {
	return model.Model{
		ID:       ID,
		Type:     string(model.INVEST),
		TargetID: 99,
		Number:   10,
	}
}

var mock = mockRepo{}

func TestImplement_Invested(t *testing.T) {
	investFactory := factory.InvestWarCafeFactory{}
	investService := warCafe.NewInvestService(investFactory)

	type fields struct {
		repo    repository.Repository
		service service.WarCafeService
	}
	type args struct {
		request request.InvestedRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    response.InvestedResponse
		wantErr bool
	}{
		{
			fields: fields{
				repo:    mock,
				service: service.NewWarCafeService(investService),
			},
			args: args{
				request: request.InvestedRequest{
					Ctx:                         context.TODO(),
					ID:                          10,
					WarCafeID:                   99,
					CurrentNumberOfParticipants: 10,
					AddNumberOfParticipants:     3,
				},
			},
			want:    response.InvestedResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Implement{
				repo:    tt.fields.repo,
				service: tt.fields.service,
			}
			got, err := i.Invested(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Invested() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invested() got = %v, want %v", got, tt.want)
			}
		})
	}
}
