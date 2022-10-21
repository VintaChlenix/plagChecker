package db

import (
	"context"
	"plagChecker/internal/model"
)

type DB interface {
	CreateMetadata(ctx context.Context, metadata *model.Metadata) error
	SelectStudentLabs(ctx context.Context, name string) ([]model.StudentCheckResult, error)
	SelectVariantMetadata(ctx context.Context, labID, variant string) ([]model.Metadata, error)
}
