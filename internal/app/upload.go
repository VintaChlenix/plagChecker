package app

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"plagChecker/internal/dto"
	"plagChecker/internal/model"
	"plagChecker/pkg/checker"
	"strings"
)

func (a *App) UploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tmpl, err := template.ParseFiles("internal/templates/upload.html")
	if err != nil {
		a.log.Errorf("failed to parse template file: %w", err)
		return
	}
	switch r.Method {
	case "GET":
		tmpl.Execute(w, nil)
	case "POST":
		request := dto.UploadRequest{UploadInfo: model.UploadInfo{
			URL:     r.FormValue("url"),
			Name:    r.FormValue("name"),
			LabID:   r.FormValue("lab_id"),
			Variant: r.FormValue("variant"),
			Ext:     r.FormValue("ext"),
		}}

		fileName, err := a.uploadFile(request)
		if err != nil {
			a.log.Errorf("failed to upload file: %w", err)
			tmpl.Execute(w, "failed to upload file")
			return
		}

		metadata, err := a.countMetadata(fileName, request)
		if err != nil {
			a.log.Errorf("failed to count metadata: %w", err)
			tmpl.Execute(w, "failed to count metadata")
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			return
		}

		checkResult, err := a.checkMetadata(ctx, metadata)
		if err != nil {
			a.log.Errorf("failed to check metadata: %w", err)
			tmpl.Execute(w, "failed to check metadata")
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			return
		}

		if err := a.storeSending(ctx, metadata, checkResult); err != nil {
			a.log.Errorf("failed to store sending: %w", err)
			tmpl.Execute(w, "failed to store your sending")
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			return
		}

		switch checkResult.Result {
		case model.CheckResultType0:
			a.log.Infof("student: %s | checkResult: %s explanation: %s | matchPercentage: %v | original: %s", request.Name, checkResult.Result, checkResult.Explanation, checkResult.MatchPercentage, checkResult.Original)
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			tmpl.Execute(w, "Plagiarism!!! Type0")
			return
		case model.CheckResultType1:
			a.log.Infof("student: %s | checkResult: %s explanation: %s | matchPercentage: %v | original: %s", request.Name, checkResult.Result, checkResult.Explanation, checkResult.MatchPercentage, checkResult.Original)
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			tmpl.Execute(w, "Plagiarism!!! Type1")
			return
		case model.CheckResultType2:
			a.log.Infof("student: %s | checkResult: %s explanation: %s | matchPercentage: %v | original: %s", request.Name, checkResult.Result, checkResult.Explanation, checkResult.MatchPercentage, checkResult.Original)
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			tmpl.Execute(w, "Plagiarism!!! Type2")
			return
		case model.CheckResultType3:
			a.log.Infof("student: %s | checkResult: %s explanation: %s | matchPercentage: %v | original: %s", request.Name, checkResult.Result, checkResult.Explanation, checkResult.MatchPercentage, checkResult.Original)
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			tmpl.Execute(w, "Plagiarism!!! Type3")
			return
		}

		a.log.Infof("student: %s | checkResult: %s explanation: %s | matchPercentage: %v | most similar: %s", request.Name, checkResult.Result, checkResult.Explanation, checkResult.MatchPercentage, checkResult.Original)

		if err := a.storeMetadata(ctx, metadata); err != nil {
			a.log.Errorf("failed to store metadata")
			tmpl.Execute(w, "failed to upload file")
			if err := os.Remove(fileName); err != nil {
				a.log.Errorf("failed to remove file %s: %w", fileName, err)
			}
			return
		}

		if err := os.Remove(fileName); err != nil {
			a.log.Errorf("failed to remove file %s: %w", fileName, err)
		}
		tmpl.Execute(w, "Very good!")
	}
}

func (a *App) storeSending(ctx context.Context, metadata *dto.CountMetadataResponse, checkResult *dto.CheckMetadataResponse) error {
	sending := &model.Sending{
		Name:    metadata.Metadata.Name,
		LabID:   metadata.Metadata.LabID,
		Variant: metadata.Metadata.Variant,
		Results: checkResult.MatchPercentage,
	}
	if err := a.db.CreateSending(ctx, sending); err != nil {
		return fmt.Errorf("failed to create sending: %w", err)
	}
	return nil
}

func (a *App) uploadFile(request dto.UploadRequest) (string, error) {
	downloadURL := a.parseURL(request.URL)

	response, err := http.Get(downloadURL)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}
	defer response.Body.Close()

	output, err := os.Create(fmt.Sprintf("%s_%s_%s.%s", request.Name, request.LabID, request.Variant, request.Ext))
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}
	return fmt.Sprintf("%s_%s_%s.%s", request.Name, request.LabID, request.Variant, request.Ext), nil
}

func (a *App) parseURL(url string) string {
	slice := strings.Split(url, "/")
	return fmt.Sprintf("https://www.googleapis.com/drive/v3/files/%s?alt=media&key=%s", slice[5], a.config.APIKey)
}

func (a *App) countMetadata(fileName string, request dto.UploadRequest) (*dto.CountMetadataResponse, error) {
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
			Name:     request.Name,
			LabID:    request.LabID,
			Variant:  request.Variant,
			NormCode: norm,
			Sum:      sum,
			Tokens:   tokens,
		},
	}, nil
}

func (a *App) checkMetadata(ctx context.Context, metadata *dto.CountMetadataResponse) (*dto.CheckMetadataResponse, error) {
	otherMetadata, err := a.db.SelectLabMetadata(ctx, metadata.Metadata.LabID)
	if err != nil {
		return nil, fmt.Errorf("failed to select metadata by variant: %w", err)
	}

	original := ""
	matchPercentage := make([]float64, 4)

	for i := range otherMetadata {
		if checker.SumCheck(metadata.Metadata.Sum, otherMetadata[i].Sum) {
			return &dto.CheckMetadataResponse{
				Result:          model.CheckResultType0,
				Explanation:     "Sums are identical, this file was sent before!",
				MatchPercentage: []float64{1, 100, 100, 100},
				Original:        otherMetadata[i].Name,
			}, nil
		}
		diff := checker.DiffCheck(metadata.Metadata.NormCode, otherMetadata[i].NormCode)
		token := checker.TokensCheck(metadata.Metadata.Tokens, otherMetadata[i].Tokens)
		metric := checker.MetricsCheck(metadata.Metadata.Tokens, otherMetadata[i].Tokens)

		if (matchPercentage[1]+matchPercentage[2]+matchPercentage[3])/3 < (diff+token+metric)/3 {
			original = otherMetadata[i].Name
			matchPercentage[1] = diff
			matchPercentage[2] = token
			matchPercentage[3] = metric
		}
	}

	if matchPercentage[1] > a.config.ReferenceValues.DiffValue {
		return &dto.CheckMetadataResponse{
			Result:          model.CheckResultType1,
			Explanation:     "Diff check failed. Code plagiarized with a little changes.",
			MatchPercentage: matchPercentage,
			Original:        original,
		}, nil
	}
	if matchPercentage[2] > a.config.ReferenceValues.TokensValue {
		return &dto.CheckMetadataResponse{
			Result:          model.CheckResultType2,
			Explanation:     "Tokens Check failed. Code plagiarized with cosmetic changes, but structures are similar.",
			MatchPercentage: matchPercentage,
			Original:        original,
		}, nil
	}
	if matchPercentage[3] > a.config.ReferenceValues.MetricsValue {
		return &dto.CheckMetadataResponse{
			Result:          model.CheckResultType3,
			Explanation:     "Metrics Check failed. The program is rewritten in some way with the general preservation of the logic of work and functionality. However, syntactically it may be completely different from the original",
			MatchPercentage: matchPercentage,
			Original:        original,
		}, nil
	}
	return &dto.CheckMetadataResponse{
		Result:          model.CheckResultType4,
		Explanation:     "That's good!",
		MatchPercentage: matchPercentage,
		Original:        original,
	}, nil
}

func (a *App) storeMetadata(ctx context.Context, metadata *dto.CountMetadataResponse) error {
	if err := a.db.CreateMetadata(ctx, &metadata.Metadata); err != nil {
		return fmt.Errorf("failed to create metadata: %w", err)
	}
	return nil
}
