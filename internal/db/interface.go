package db

import (
	"context"
	"plagChecker/internal/model"
)

type DB interface {
	CreateMetadata(ctx context.Context, metadata *model.Metadata) error
	SelectStudentMetadata(ctx context.Context, name string, labID string, variant string) (*model.Metadata, error)
	SelectVariantMetadata(ctx context.Context, labID, variant string) ([]model.Metadata, error)
}
