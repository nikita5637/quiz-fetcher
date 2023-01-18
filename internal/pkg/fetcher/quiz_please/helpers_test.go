package quiz_please

import (
	"testing"
	"time"

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
				ExternalID:  1,
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        "name",
				PlaceID:     0,
				DateTime:    dt,
				Price:       400,
				PaymentType: "",
				MaxPlayers:  9,
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
				ExternalID:  1,
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        "name",
				PlaceID:     0,
				DateTime:    dt,
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  9,
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
				ExternalID:  1,
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        "name",
				PlaceID:     0,
				DateTime:    dt,
				Price:       400,
				PaymentType: "card",
				MaxPlayers:  9,
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
				ExternalID:  1,
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        "name",
				PlaceID:     0,
				DateTime:    dt,
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
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
				ExternalID:  1,
				LeagueID:    leagueID,
				Type:        1,
				Number:      "1",
				Name:        "name",
				PlaceID:     0,
				DateTime:    dt,
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
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

func Test_getPlaceKey(t *testing.T) {
	type args struct {
		name    string
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				address: "address",
				name:    "name",
			},
			want: "name:address",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPlaceKey(tt.args.name, tt.args.address); got != tt.want {
				t.Errorf("getPlaceKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
