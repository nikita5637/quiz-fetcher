package sixty_seconds

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// Result ...
type Result struct {
	Title   string
	Total   float64
	Tours   []float64
	Pregame float64
}

// Response ...
type Response struct {
	Results string
}

// GetGameResult ...
func (f *Fetcher) GetGameResult(ctx context.Context, externalID int32) (model.GameResult, error) {
	url := fmt.Sprintf(f.url+f.gameResultPath, externalID)
	resp, err := f.client.Get(url)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("getting HTTP response error: %w", err)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("decoding response error: %w", err)
	}
	defer resp.Body.Close()

	var results []Result
	err = json.Unmarshal([]byte(response.Results), &results)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("unmarshaling results error: %w", err)
	}

	place := 0
	tmp := make([]string, 0, 5)
	builder := strings.Builder{}
	builder.WriteString("{")

	for i, result := range results {
		if result.Title != f.team {
			continue
		}

		place = i + 1
		tmp = append(tmp, fmt.Sprintf("\"Сумма\":%d", int(result.Total)))

		for round, points := range result.Tours {
			tmp = append(tmp, fmt.Sprintf("\"%d\":%d", round+1, int(points)))
		}

		tmp = append(tmp, fmt.Sprintf("\"Матрица\":%d", int(result.Pregame)))

		break
	}

	builder.WriteString(strings.Join(tmp, ","))
	builder.WriteString("}")

	return model.GameResult{
		ResultPlace: uint32(place),
		RoundPoints: maybe.Just(builder.String()),
	}, nil
}
