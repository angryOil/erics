package repository

import (
	"context"
	"erics/internal/war_cafe/domain/model"
)

type Repository interface {
	Save(ctx context.Context, model model.Model)
	Load(ctx context.Context, ID int64) model.Model
}
