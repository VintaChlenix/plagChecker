package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"plagChecker/internal/db"
	"plagChecker/internal/model"
)

type Client struct {
	db *pgx.Conn
}

var _ db.DB = Client{}

func NewClient(connectionString string) (*Client, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	return &Client{db: conn}, nil
}

func (c Client) CreateMetadata(ctx context.Context, metadata *model.Metadata) error {
	q := `
		INSERT INTO
		  metadata(name, lab_id, variant, norm_code, sum, tokens)
		VALUES
		  ($1, $2, $3, $4, $5, $6)
	`
	if _, err := c.db.Exec(
		ctx,
		q,
		metadata.Name,
		metadata.LabID,
		metadata.Variant,
		metadata.NormCode,
		metadata.Sum,
		metadata.Tokens,
	); err != nil {
		return fmt.Errorf("failed to insert student metadata: %w", err)
	}
	return nil
}

func (c Client) SelectStudentLabs(ctx context.Context, name string) ([]model.StudentCheckResult, error) {
	q := `
		SELECT
		  lab_id, variant
		FROM
		  metadata
		WHERE
		  name = $1
	`
	rows, err := c.db.Query(ctx, q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	studentLabs := make([]model.StudentCheckResult, 0)
	for rows.Next() {
		var studentLab model.StudentCheckResult
		if err := rows.Scan(&studentLab.LabID, &studentLab.Variant); err != nil {
			return nil, fmt.Errorf("failed to parse student lab: %w", err)
		}
		studentLabs = append(studentLabs, studentLab)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse student labs: %w", err)
	}
	return studentLabs, nil
}

func (c Client) SelectVariantMetadata(ctx context.Context, labID, variant string) ([]model.Metadata, error) {
	q := `
		SELECT
		  name, lab_id, variant, norm_code, sum, tokens
		FROM
		  metadata
		WHERE
		  lab_id = $1 AND variant = $2
	`
	rows, err := c.db.Query(ctx, q, labID, variant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	studentsMetadata := make([]model.Metadata, 0)
	for rows.Next() {
		var studentMetadata model.Metadata
		if err := rows.Scan(&studentMetadata.Name, &studentMetadata.LabID, &studentMetadata.Variant, &studentMetadata.NormCode, &studentMetadata.Sum, &studentMetadata.Tokens); err != nil {
			return nil, fmt.Errorf("failed to parse student metadata: %w", err)
		}
		studentsMetadata = append(studentsMetadata, studentMetadata)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse students metadata: %w", err)
	}
	return studentsMetadata, nil
}
