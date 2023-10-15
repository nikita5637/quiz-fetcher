package quiz_please

import (
	"testing"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	"github.com/stretchr/testify/assert"
)

func Test_convertGameToModelGame(t *testing.T) {
	loc, err := time.LoadLocation(time_utils.TimeZoneMoscow)
	assert.NoError(t, err)

	dt, err := time.ParseInLocation(timeFormatString, "02.01.22 19:30", loc)
	assert.NoError(t, err)

	type args struct {
		game game
	}
	tests := []struct {
		name    string
		args    args
		want    model.Game
		wantErr bool
	}{
		{
			name: "invalid game date time",
			args: args{
				game: game{
					DateTime: "invalid",
				},
			},
			want:    model.Game{},
			wantErr: true,
		},
		{
			name: "invalid game price",
			args: args{
				game: game{
					DateTime: "02.01.22 19:30",
					Price:    "invalid",
				},
			},
			want:    model.Game{},
			wantErr: true,
		},
		{
			name: "empty payment type",
			args: args{
				game: game{
					ExternalID:   1,
					GameTypeID:   1,
					Number:       "1",
					Name:         "name",
					PlaceName:    "place",
					PlaceAddress: "address",
					DateTime:     "02.01.22 19:30",
					Price:        "400₽",
					Text:         "",
					MaxPlayers:   9,
				},
			},
			want: model.Game{
				ExternalID:  maybe.Just(int32(1)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        maybe.Just("name"),
				PlaceID:     0,
				DateTime:    dt.UTC(),
				Price:       400,
				PaymentType: maybe.Nothing[string](),
				MaxPlayers:  9,
				IsInMaster:  true,
			},
			wantErr: false,
		},
		{
			name: "payment type \"cash\"",
			args: args{
				game: game{
					ExternalID:   1,
					GameTypeID:   1,
					Number:       "1",
					Name:         "name",
					PlaceName:    "place",
					PlaceAddress: "address",
					DateTime:     "02.01.22 19:30",
					Price:        "400₽",
					Text:         "с человека, наличные",
					MaxPlayers:   9,
				},
			},
			want: model.Game{
				ExternalID:  maybe.Just(int32(1)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        maybe.Just("name"),
				PlaceID:     0,
				DateTime:    dt.UTC(),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  9,
				IsInMaster:  true,
			},
			wantErr: false,
		},
		{
			name: "payment type \"card\"",
			args: args{
				game: game{
					ExternalID:   1,
					GameTypeID:   1,
					Number:       "1",
					Name:         "name",
					PlaceName:    "place",
					PlaceAddress: "address",
					DateTime:     "02.01.22 19:30",
					Price:        "400₽",
					Text:         "с человека, карта",
					MaxPlayers:   9,
				},
			},
			want: model.Game{
				ExternalID:  maybe.Just(int32(1)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        maybe.Just("name"),
				PlaceID:     0,
				DateTime:    dt.UTC(),
				Price:       400,
				PaymentType: maybe.Just("card"),
				MaxPlayers:  9,
				IsInMaster:  true,
			},
			wantErr: false,
		},
		{
			name: "payment type \"cash,card\"",
			args: args{
				game: game{
					ExternalID:   1,
					GameTypeID:   1,
					Number:       "1",
					Name:         "name",
					PlaceName:    "place",
					PlaceAddress: "address",
					DateTime:     "02.01.22 19:30",
					Price:        "400₽",
					Text:         "с человека, наличные или карта",
					MaxPlayers:   9,
				},
			},
			want: model.Game{
				ExternalID:  maybe.Just(int32(1)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        maybe.Just("name"),
				PlaceID:     0,
				DateTime:    dt.UTC(),
				Price:       400,
				PaymentType: maybe.Just("cash,card"),
				MaxPlayers:  9,
				IsInMaster:  true,
			},
			wantErr: false,
		},
		{
			name: "game name without SPB",
			args: args{
				game: game{
					ExternalID:   1,
					GameTypeID:   1,
					Number:       "1",
					Name:         "name SPB",
					PlaceName:    "place",
					PlaceAddress: "address",
					DateTime:     "02.01.22 19:30",
					Price:        "400₽",
					Text:         "с человека, наличные или карта",
					MaxPlayers:   9,
				},
			},
			want: model.Game{
				ExternalID:  maybe.Just(int32(1)),
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        maybe.Just("name"),
				PlaceID:     0,
				DateTime:    dt.UTC(),
				Price:       400,
				PaymentType: maybe.Just("cash,card"),
				MaxPlayers:  9,
				IsInMaster:  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertGameToModelGame(tt.args.game)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertGameToModelGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
