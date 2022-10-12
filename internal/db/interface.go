package db

import (
	"context"
	"plagChecker/internal/model"
)

type DB interface {
	CreateMetadata(ctx context.Context, metadata *model.Metadata) error
	SelectStudentMetadata(ctx context.Context, name string, labID int, variant int) (*model.Metadata, error)
	SelectVariantMetadata(ctx context.Context, labID, variant int) ([]model.Metadata, error)
}
