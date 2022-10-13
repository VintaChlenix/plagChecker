package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"os"
	"plagChecker/internal/dto"
	"plagChecker/internal/model"
	"plagChecker/pkg/checker"
)

func (a *App) UploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	url1 := "https://drive.google.com/uc?id=1RX1xgNcBdDBmRIWolQ3fRSklb7XIhOgn&export=download"
	//url2 := "https://drive.google.com/u/0/uc?id=1DyK-8dTA9K1Di10zB5xI8_AyWp7vLzna&export=download"

	name := chi.URLParam(r, "name")
	labID := chi.URLParam(r, "labID")
	variant := chi.URLParam(r, "variant")

	fileName, err := a.uploadFile(url1, name, labID, variant)
	if err != nil {
		a.log.Errorf("failed to upload file: %w")
		w.Write([]byte("failed to upload file"))
		return
	}

	metadata, err := a.countMetadata(fileName, name, labID, variant)
	if err != nil {
		a.log.Errorf("failed to count metadata: %w", err)
		w.Write([]byte("failed to count metadata"))
		if err := os.Remove(fileName); err != nil {
			a.log.Errorf("failed to remove file %s: %w", fileName, err)
		}
		return
	}

	checkResult, err := a.checkMetadata(ctx, metadata)
	if err != nil {
		a.log.Errorf("failed to check metadata: %w", err)
		w.Write([]byte("failed to check metadata"))
		if err := os.Remove(fileName); err != nil {
			a.log.Errorf("failed to remove file %s: %w", fileName, err)
		}
		return
	}
	switch checkResult.Result {
	case model.CheckResultType1:
		a.log.Infof("student: %s | checkResult: %s explanation: %s", name, checkResult.Result, checkResult.Explanation)
		if err := os.Remove(fileName); err != nil {
			a.log.Errorf("failed to remove file %s: %w", fileName, err)
		}
		w.Write([]byte("Plagiarism!!! Type1"))
		return
	case model.CheckResultType2:
		a.log.Infof("student: %s | checkResult: %s explanation: %s", name, checkResult.Result, checkResult.Explanation)
		if err := os.Remove(fileName); err != nil {
			a.log.Errorf("failed to remove file %s: %w", fileName, err)
		}
		w.Write([]byte("Plagiarism!!! Type2"))
		return
	case model.CheckResultType3:
		a.log.Infof("student: %s | checkResult: %s explanation: %s", name, checkResult.Result, checkResult.Explanation)
		if err := os.Remove(fileName); err != nil {
			a.log.Errorf("failed to remove file %s: %w", fileName, err)
		}
		w.Write([]byte("OK."))
		return
	}

	if err := a.storeMetadata(ctx, metadata); err != nil {
		a.log.Errorf("failed to store metadata")
		w.Write([]byte("failed to upload file"))
		if err := os.Remove(fileName); err != nil {
			a.log.Errorf("failed to remove file %s: %w", fileName, err)
		}
		return
	}

	if err := os.Remove(fileName); err != nil {
		a.log.Errorf("failed to remove file %s: %w", fileName, err)
	}
	w.Write([]byte("Very good!!!"))
}

func (a *App) uploadFile(url, name, labID, variant string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}
	defer response.Body.Close()

	output, err := os.Create(fmt.Sprintf("%s_%s_%s", name, labID, variant))
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}
	return fmt.Sprintf("%s_%s_%s", name, labID, variant), nil
}

func (a *App) countMetadata(fileName, name, labID, variant string) (*dto.CountMetadataResponse, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	norm, err := checker.Normalize(f)
	if err != nil {
		return nil, fmt.Errorf("failed to normalize file: %w", err)
	}

	sum, err := checker.GetSum(f)
	if err != nil {
		return nil, fmt.Errorf("failed to get sum of the file: %w", err)
	}

	tokens, err := checker.GetTokens(f)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokens of the file: %w", err)
	}
	return &dto.CountMetadataResponse{
		Metadata: model.Metadata{
			Name:     name,
			LabID:    labID,
			Variant:  variant,
			NormCode: norm,
			Sum:      sum,
			Tokens:   tokens,
		},
	}, nil
}

func (a *App) checkMetadata(ctx context.Context, metadata *dto.CountMetadataResponse) (*dto.CheckMetadataResponse, error) {
	otherMetadata, err := a.db.SelectVariantMetadata(ctx, metadata.Metadata.LabID, metadata.Metadata.Variant)
	if err != nil {
		return nil, fmt.Errorf("failed to select metadata by variant: %w", err)
	}
	for i := range otherMetadata {
		if checker.SumCheck(metadata.Metadata.Sum, otherMetadata[i].Sum) {
			return &dto.CheckMetadataResponse{
				Result:      model.CheckResultType1,
				Explanation: "Sums are identical, this file was sent before!",
			}, nil
		}

		if checker.DiffCheck(metadata.Metadata.NormCode, otherMetadata[i].NormCode) > float64(60) {
			return &dto.CheckMetadataResponse{
				Result:      model.CheckResultType1,
				Explanation: "Diff check failed. Code plagiarized with a little changes.",
			}, nil
		}

		if checker.TokensCheck(metadata.Metadata.Tokens, otherMetadata[i].Tokens) > float64(60) {
			return &dto.CheckMetadataResponse{
				Result:      model.CheckResultType2,
				Explanation: "TokensCheck failed. Code plagiarized with renamings, but structures are similar.",
			}, nil
		}
	}
	return &dto.CheckMetadataResponse{
		Result:      model.CheckResultType4,
		Explanation: "That's good!",
	}, nil
}

func (a *App) storeMetadata(ctx context.Context, metadata *dto.CountMetadataResponse) error {
	if err := a.db.CreateMetadata(ctx, &metadata.Metadata); err != nil {
		return fmt.Errorf("failed to create metadata: %w", err)
	}
	return nil
}
