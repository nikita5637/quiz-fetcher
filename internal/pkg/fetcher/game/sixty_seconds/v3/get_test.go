package sixty_seconds

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mocks"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	"github.com/stretchr/testify/suite"
)

const (
	game25678 = "/quizgames/game/25678/"
	game25679 = "/quizgames/game/25679/"
	game25680 = "/quizgames/game/25680/"
	game25681 = "/quizgames/game/25681/"
	game25682 = "/quizgames/game/25682/"
	game25683 = "/quizgames/game/25683/"
	game25684 = "/quizgames/game/25684/"
	game25685 = "/quizgames/game/25685/"
	game25686 = "/quizgames/game/25686/"
	game25687 = "/quizgames/game/25687/"
	game25688 = "/quizgames/game/25688/"
	game25781 = "/quizgames/game/25781/"
	game25782 = "/quizgames/game/25782/"
	game25783 = "/quizgames/game/25783/"
	game25784 = "/quizgames/game/25784/"
	game25785 = "/quizgames/game/25785/"
	game25786 = "/quizgames/game/25786/"
	game25787 = "/quizgames/game/25787/"
	game25788 = "/quizgames/game/25788/"
	game25789 = "/quizgames/game/25789/"
	game25790 = "/quizgames/game/25790/"
	game25791 = "/quizgames/game/25791/"
	game28702 = "/quizgames/game/28702/"
	game28703 = "/quizgames/game/28703/"
	game30313 = "/quizgames/game/30313/"
	game30325 = "/quizgames/game/30325/"
	game30314 = "/quizgames/game/30314/"
	game30326 = "/quizgames/game/30326/"
	game30315 = "/quizgames/game/30315/"
	game30327 = "/quizgames/game/30327/"
	game30316 = "/quizgames/game/30316/"
	game30328 = "/quizgames/game/30328/"
	game30317 = "/quizgames/game/30317/"
	game30329 = "/quizgames/game/30329/"
	game30318 = "/quizgames/game/30318/"
	game30330 = "/quizgames/game/30330/"
	game30319 = "/quizgames/game/30319/"
	game30331 = "/quizgames/game/30331/"
	game30320 = "/quizgames/game/30320/"
	game30332 = "/quizgames/game/30332/"
	game30321 = "/quizgames/game/30321/"
	game30333 = "/quizgames/game/30333/"
)

// GetSuite ...
type GetSuite struct {
	suite.Suite

	ctx context.Context

	placeStorage *mocks.PlaceStorage

	fetcher *Fetcher

	svr *httptest.Server
}

// SetupSuite ...
func (s *GetSuite) SetupSuite() {
	s.svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var r io.Reader
		switch req.URL.Path {
		case "/html1":
			r = strings.NewReader(html1)
		case "/html2":
			r = strings.NewReader(html2)
		case "/html3":
			r = strings.NewReader(html3)
		case "/html4":
			r = strings.NewReader(html4)
		case game25678:
			r = strings.NewReader(html25678)
		case game25679:
			r = strings.NewReader(html25679)
		case game25680:
			r = strings.NewReader(html25680)
		case game25681:
			r = strings.NewReader(html25681)
		case game25682:
			r = strings.NewReader(html25682)
		case game25683:
			r = strings.NewReader(html25683)
		case game25684:
			r = strings.NewReader(html25684)
		case game25685:
			r = strings.NewReader(html25685)
		case game25686:
			r = strings.NewReader(html25686)
		case game25687:
			r = strings.NewReader(html25687)
		case game25688:
			r = strings.NewReader(html25688)
		case game25781:
			r = strings.NewReader(html25781)
		case game25782:
			r = strings.NewReader(html25782)
		case game25783:
			r = strings.NewReader(html25783)
		case game25784:
			r = strings.NewReader(html25784)
		case game25785:
			r = strings.NewReader(html25785)
		case game25786:
			r = strings.NewReader(html25786)
		case game25787:
			r = strings.NewReader(html25787)
		case game25788:
			r = strings.NewReader(html25788)
		case game25789:
			r = strings.NewReader(html25789)
		case game25790:
			r = strings.NewReader(html25790)
		case game25791:
			r = strings.NewReader(html25791)
		case game28702:
			r = strings.NewReader(html28702)
		case game28703:
			r = strings.NewReader(html28703)
		case game30313:
			r = strings.NewReader(html30313)
		case game30325:
			r = strings.NewReader(html30325)
		case game30314:
			r = strings.NewReader(html30314)
		case game30326:
			r = strings.NewReader(html30326)
		case game30315:
			r = strings.NewReader(html30315)
		case game30327:
			r = strings.NewReader(html30327)
		case game30316:
			r = strings.NewReader(html30316)
		case game30328:
			r = strings.NewReader(html30328)
		case game30317:
			r = strings.NewReader(html30317)
		case game30329:
			r = strings.NewReader(html30329)
		case game30318:
			r = strings.NewReader(html30318)
		case game30330:
			r = strings.NewReader(html30330)
		case game30319:
			r = strings.NewReader(html30319)
		case game30331:
			r = strings.NewReader(html30331)
		case game30320:
			r = strings.NewReader(html30320)
		case game30332:
			r = strings.NewReader(html30332)
		case game30321:
			r = strings.NewReader(html30321)
		case game30333:
			r = strings.NewReader(html30333)
		}
		_, err := io.Copy(w, r)
		s.NoError(err)
	}))
}

// TearDownSuite ...
func (s *GetSuite) TearDownSuite() {
	s.svr.Close()
}

// SetupTest ...
func (s *GetSuite) SetupTest() {
	s.ctx = context.Background()

	s.placeStorage = mocks.NewPlaceStorage(s.T())
}

// TestGetGamesList ...
func (s *GetSuite) TestGetGamesList() {
	s.Run("ok", func() {
		s.fetcher = New(Config{
			PlaceStorage: s.placeStorage,

			NeedToFetchOpenLeague:  true,
			NeedToFetchFirstLeague: true,
			SchedulePath:           "/html1",
			URL:                    s.svr.URL,
		})

		s.placeStorage.EXPECT().GetPlaceByNameAndAddress(s.ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(mysql.Place{
			ExternalID: 1,
		}, nil).Times(21)

		s.placeStorage.EXPECT().GetPlaceByNameAndAddress(s.ctx, "Фрегат \"Благодать\"", "Петровская наб., 2А").Return(mysql.Place{
			ExternalID: 2,
		}, nil).Times(1)

		got, err := s.fetcher.GetGamesList(s.ctx)
		expected := []model.Game{
			{
				ExternalID:  maybe.Just(int32(25678)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#3",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-06-17 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25781)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#2",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-06-18 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25679)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#4",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-06-24 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25782)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#3",
				Name:        maybe.Just(firstLeague),
				PlaceID:     2,
				DateTime:    time_utils.ConvertTime("2024-06-25 16:30"),
				Price:       2400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25680)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#5",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-01 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25783)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#4",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-02 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25681)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#6",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-08 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25784)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#5",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-09 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25682)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#7",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-15 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25785)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#6",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-16 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25683)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#8",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-22 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25786)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#7",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-23 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25684)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#9",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-29 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25787)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#8",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-30 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25685)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#10",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-05 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25788)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#9",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-06 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25686)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#11",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-12 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25789)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#10",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-13 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25687)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#12",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-19 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25790)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#11",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-20 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25688)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      final,
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-26 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25791)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      final,
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-27 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}

		s.Equal(expected, got)
		s.NoError(err)
	})

	s.Run("ok", func() {
		s.fetcher = New(Config{
			PlaceStorage: s.placeStorage,

			NeedToFetchOpenLeague:  false,
			NeedToFetchFirstLeague: true,
			SchedulePath:           "/html2",
			URL:                    s.svr.URL,
		})

		s.placeStorage.EXPECT().GetPlaceByNameAndAddress(s.ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(mysql.Place{
			ExternalID: 1,
		}, nil).Times(6)

		got, err := s.fetcher.GetGamesList(s.ctx)
		s.Equal([]model.Game{
			{
				ExternalID:  maybe.Just(int32(25786)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#7",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-23 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25787)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#8",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-07-30 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25788)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#9",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-06 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25789)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#10",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-13 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25790)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#11",
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-20 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(25791)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      final,
				Name:        maybe.Just(firstLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-08-27 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		s.NoError(err)
	})

	s.Run("ok", func() {
		s.fetcher = New(Config{
			PlaceStorage: s.placeStorage,

			NeedToFetchOpenLeague:  true,
			NeedToFetchFirstLeague: false,
			SchedulePath:           "/html3",
			URL:                    s.svr.URL,
		})

		s.placeStorage.EXPECT().GetPlaceByNameAndAddress(s.ctx, "Rossi's Club", "ул. Зодчего Росси, 1-3").Return(mysql.Place{
			ExternalID: 8,
		}, nil).Twice()

		got, err := s.fetcher.GetGamesList(s.ctx)
		s.Equal([]model.Game{
			{
				ExternalID:  maybe.Just(int32(28702)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#2",
				Name:        maybe.Just(openLeague),
				PlaceID:     8,
				DateTime:    time_utils.ConvertTime("2024-12-09 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		s.NoError(err)
	})

	s.Run("ok", func() {
		s.fetcher = New(Config{
			PlaceStorage: s.placeStorage,

			NeedToFetchOpenLeague:  true,
			NeedToFetchFirstLeague: true,
			SchedulePath:           "/html4",
			URL:                    s.svr.URL,
		})

		s.placeStorage.EXPECT().GetPlaceByNameAndAddress(s.ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(mysql.Place{
			ExternalID: 1,
		}, nil).Times(16)

		got, err := s.fetcher.GetGamesList(s.ctx)
		expected := []model.Game{
			{
				ExternalID:  maybe.Just(int32(30313)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#4",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-03-31 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30325)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#4",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-01 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30314)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#5",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-07 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30326)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#5",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-08 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30315)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#6",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-14 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30327)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#6",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-15 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30316)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#7",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-21 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30328)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#7",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-22 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30317)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#8",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-28 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30329)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#8",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-04-29 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30318)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#9",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-05-05 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30330)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#9",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-05-06 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30319)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#10",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-05-12 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30332)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "#11",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-05-20 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30321)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "Финал",
				Name:        maybe.Just(openLeague),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-05-26 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(30333)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "Финал",
				Name:        maybe.Just(string(firstLeague)),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2025-05-27 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}
		s.Equal(expected, got)
		s.NoError(err)
	})
}

// Test_getExternalID ...
func Test_getExternalID(t *testing.T) {
	t.Parallel()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				path: "/quizgames/game/",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				path: "/quizgames/game/25781/",
			},
			want:    25781,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getExternalID(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getExternalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getExternalID() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_getName ...
func Test_getName(t *testing.T) {
	t.Parallel()

	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				text: "Первая лига | Игра #1",
			},
			want: firstLeague,
		},
		{
			name: "tc2",
			args: args{
				text: "Первая лига Игра #1",
			},
			want: "",
		},
		{
			name: "tc3",
			args: args{
				text: "\n               Первая лига | Игра #1",
			},
			want: firstLeague,
		},
		{
			name: "tc4",
			args: args{
				text: openLeagueFinal,
			},
			want: openLeague,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getName(tt.args.text); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_getNumber ...
func Test_getNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				text: "Первая лига | Игра #1",
			},
			want: "#1",
		},
		{
			name: "tc2",
			args: args{
				text: "Первая лига Игра #1",
			},
			want: "",
		},
		{
			name: "tc3",
			args: args{
				text: openLeagueFinal,
			},
			want: final,
		},
		{
			name: "tc4",
			args: args{
				text: "\n                            Первая лига | Игра #1\n                        ",
			},
			want: "#1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumber(tt.args.text); got != tt.want {
				t.Errorf("getNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_getPrice ...
func Test_getPrice(t *testing.T) {
	t.Parallel()

	type args struct {
		price string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				price: "					     1500 руб. с команды",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				price: "1500 руб. с команды",
			},
			want:    1500,
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				price: "400 руб. с человека",
			},
			want:    400,
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				price: "1500 руб. с команды + депозит",
			},
			want:    1500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPrice(tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
