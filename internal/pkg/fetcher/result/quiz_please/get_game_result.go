package quiz_please

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

const (
	placeColumnName = "Место"
	teamColumnName  = "Название команды"
)

var (
	permittedColumnNames = map[string]struct{}{
		placeColumnName: {},
		teamColumnName:  {},
		"1 раунд":       {},
		"2 раунд":       {},
		"3 раунд":       {},
		"4 раунд":       {},
		"5 раунд":       {},
		"6 раунд":       {},
		"7 раунд":       {},
		"8 раунд":       {},
		"итого":         {},
	}
)

// GetGameResult ...
func (f *Fetcher) GetGameResult(ctx context.Context, externalID int32) (model.GameResult, error) {
	url := fmt.Sprintf(f.url+f.gameResultPath, externalID)
	resp, err := f.client.Get(url)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("getting HTTP response error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return model.GameResult{}, fmt.Errorf("status code is %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("reading document from HTTP response error: %w", err)
	}

	gameTable := doc.Find(".game-table")
	if gameTable.Length() != 1 {
		return model.GameResult{}, fmt.Errorf("number of game tables on page not equal 1")
	}

	tHead := gameTable.Find("thead")
	if tHead.Length() != 1 {
		return model.GameResult{}, fmt.Errorf("number of tHead not equal 1")
	}

	tHeadTRs := tHead.Find("tr")
	if tHeadTRs.Length() != 1 {
		return model.GameResult{}, fmt.Errorf("number of <tr> tags in tHead not equal 1")
	}

	columnIndexes := make([]int, 0, 9)
	columnNames := make([]string, 0, 9)
	columnPlaceIndex := -1
	columnTeamIndex := -1
	tHeadTRs.Find("td").Each(func(i int, s *goquery.Selection) {
		columnName := strings.TrimSpace(s.Text())
		if _, ok := permittedColumnNames[columnName]; ok {
			switch columnName {
			case placeColumnName:
				columnPlaceIndex = i
			case teamColumnName:
				columnTeamIndex = i
			default:
				columnIndexes = append(columnIndexes, i)
				columnNames = append(columnNames, columnName)
			}
		}
	})

	if columnPlaceIndex == -1 {
		return model.GameResult{}, fmt.Errorf("column with place not found")
	}

	if columnTeamIndex == -1 {
		return model.GameResult{}, fmt.Errorf("column with team not found")
	}

	tBody := gameTable.Find("tbody")
	if tBody.Length() != 1 {
		return model.GameResult{}, fmt.Errorf("number of tBody not equal 1")
	}

	strPlace := ""
	tmp := make([]string, 0, len(columnIndexes))
	builder := strings.Builder{}
	builder.WriteString("{")
	tBody.Find("tr").EachWithBreak(func(_ int, tr *goquery.Selection) bool {
		tdValues := tr.Find("td").Map(func(_ int, td *goquery.Selection) string {
			return strings.TrimSpace(td.Text())
		})

		if tdValues[columnTeamIndex] != f.team {
			return true
		}

		strPlace = tdValues[columnPlaceIndex]

		for i, columnIndex := range columnIndexes {
			tmp = append(tmp, fmt.Sprintf("\"%s\":%s", columnNames[i], tdValues[columnIndex]))
		}

		return false
	})
	builder.WriteString(strings.Join(tmp, ","))
	builder.WriteString("}")

	place, err := strconv.ParseUint(strPlace, 10, 64)
	if err != nil {
		return model.GameResult{}, fmt.Errorf("parsing uint error: %w", err)
	}

	return model.GameResult{
		ResultPlace: uint32(place),
		RoundPoints: maybe.Just(builder.String()),
	}, nil
}
