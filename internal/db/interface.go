package db

import (
	"context"
	"plagChecker/internal/model"
)

type DB interface {
	CreateMetadata(ctx context.Context, metadata *model.Metadata) error
	SelectStudentLabs(ctx context.Context, name string) ([]model.StudentCheckResult, error)
	SelectLabMetadata(ctx context.Context, labID string) ([]model.Metadata, error)
	SelectVariantMetadata(ctx context.Context, labID, variant string) ([]model.Metadata, error)
}
