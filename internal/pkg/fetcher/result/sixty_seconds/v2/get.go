package sixty_seconds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// Main ...
type Main struct {
	Name       string    `json:"name,omitempty"`
	Summ       float64   `json:"summ,omitempty"`
	Answers    [][][]any `json:"answers,omitempty"`
	MatrixSumm float64   `json:"matrix_summ,omitempty"`
}

// Response ...
type Response struct {
	Mains []Main `json:"main,omitempty"`
}

// GetGameResult ...
func (f *Fetcher) GetGameResult(ctx context.Context, externalID int32) (model.GameResult, error) {
	req, err := http.NewRequest(http.MethodPost, f.url+f.gameResultPath, strings.NewReader(fmt.Sprintf("game_id=%d", externalID)))
	if err != nil {
		return model.GameResult{}, fmt.Errorf("creating HTTP request error: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := f.client.Do(req)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("getting HTTP response error: %w", err)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("decoding response error: %w", err)
	}
	defer resp.Body.Close()

	place := 0
	tmp := make([]string, 0, 5)
	builder := strings.Builder{}
	builder.WriteString("{")
	for i, main := range response.Mains {
		if main.Name != f.team {
			continue
		}

		place = i + 1

		tmp = append(tmp, fmt.Sprintf("\"Сумма\":%d", int(main.Summ)))

		for round, answers := range main.Answers {
			roundSumm := 0.0
			for _, answer := range answers {
				roundSumm += answer[0].(float64)
			}
			tmp = append(tmp, fmt.Sprintf("\"%d\":%d", round+1, int(roundSumm)))
		}

		tmp = append(tmp, fmt.Sprintf("\"Матрица\":%d", int(main.MatrixSumm)))

		break
	}
	builder.WriteString(strings.Join(tmp, ","))
	builder.WriteString("}")

	return model.GameResult{
		ResultPlace: uint32(place),
		RoundPoints: maybe.Just(builder.String()),
	}, nil
}
