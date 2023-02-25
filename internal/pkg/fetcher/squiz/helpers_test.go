package squiz

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/squiz/mocks"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx          context.Context
	gamesFetcher *GamesFetcher

	gameTypeMatchStorage *mocks.GameTypeMatchStorage
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gameTypeMatchStorage: mocks.NewGameTypeMatchStorage(t),
	}

	fx.gamesFetcher = &GamesFetcher{
		gameTypeMatchStorage: fx.gameTypeMatchStorage,
	}

	t.Cleanup(func() {})

	return fx
}

func Test_convertGameToModelGame(t *testing.T) {
	t.Run("invalid href", func(t *testing.T) {
		fx := tearUp(t)

		g := game{
			Href: "#href",
		}

		got, err := fx.gamesFetcher.convertGameToModelGame(fx.ctx, g)
		assert.Equal(t, model.Game{}, got)
		assert.Error(t, err)
	})

	t.Run("error while get game type", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "invalid description").Return(0, errors.New("some error"))

		g := game{
			Href:        "#123",
			Description: "invalid description",
		}

		got, err := fx.gamesFetcher.convertGameToModelGame(fx.ctx, g)
		assert.Equal(t, model.Game{}, got)
		assert.Error(t, err)
	})

	t.Run("game type is 0", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "description").Return(0, nil)

		g := game{
			Href:        "#123",
			Description: "description",
		}

		got, err := fx.gamesFetcher.convertGameToModelGame(fx.ctx, g)
		assert.Equal(t, model.Game{}, got)
		assert.Error(t, err)
	})

	t.Run("invalid game number", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "description").Return(1, nil)

		g := game{
			Href:        "#123",
			Description: "description",
			Number:      "",
		}

		got, err := fx.gamesFetcher.convertGameToModelGame(fx.ctx, g)
		assert.Equal(t, model.Game{}, got)
		assert.Error(t, err)
	})

	t.Run("invalid date time", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "description").Return(1, nil)

		g := game{
			Href:        "#123",
			Description: "description",
			Number:      "1",
			DateTime:    "invalid date time",
		}

		got, err := fx.gamesFetcher.convertGameToModelGame(fx.ctx, g)
		assert.Equal(t, model.Game{}, got)
		assert.Error(t, err)
	})

	t.Run("invalid payment info", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "description").Return(1, nil)

		g := game{
			Href:        "#123",
			Description: "description",
			Number:      "1",
			DateTime:    "21 января 2023 19:30",
			PaymentInfo: "invalid payment info",
		}

		got, err := fx.gamesFetcher.convertGameToModelGame(fx.ctx, g)
		assert.Equal(t, model.Game{}, got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "description").Return(1, nil)

		g := game{
			Href:        "#123",
			Description: "description",
			Name:        "name",
			Number:      "1",
			DateTime:    "21 января 2023 19:30",
			PaymentInfo: "400 рублей с человека, оплата только наличными",
		}

		got, err := fx.gamesFetcher.convertGameToModelGame(fx.ctx, g)
		assert.Equal(t, model.Game{
			ExternalID:  123,
			LeagueID:    leagueID,
			Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
			Number:      "1",
			Name:        "name",
			PlaceID:     0,
			DateTime:    convertTime("2023-01-21 16:30"),
			Price:       400,
			PaymentType: "cash",
			MaxPlayers:  8,
		}, got)
		assert.NoError(t, err)
	})
}

func Test_dirtyHooksForDateTime(t *testing.T) {
	type args struct {
		dateTime string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "dirty hook 1",
			args: args{
				dateTime: "21января 2023 16:00",
			},
			want: "21 января 2023 16:00",
		},
		{
			name: "ok",
			args: args{
				dateTime: "21 января 2023 16:00",
			},
			want: "21 января 2023 16:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dirtyHooksForDateTime(tt.args.dateTime); got != tt.want {
				t.Errorf("dirtyHooksForDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInfoFromHTML(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{

		{
			name: "test case 1",
			args: args{
				text: "16:00, BarBQ Night, ул. Ломоносова, 16",
			},
			want:    "BarBQ Night",
			want1:   "ул. Ломоносова, 16",
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				text: "19:30, Parkking, Александровский парк, 4, корп. 3",
			},
			want:    "Parkking",
			want1:   "Александровский парк, 4, корп. 3",
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				text: "19:30, BarBQ Night, ул. Ломоносова, 16",
			},
			want:    "BarBQ Night",
			want1:   "ул. Ломоносова, 16",
			wantErr: false,
		},
		{
			name: "test case 4",
			args: args{
				text: "19:30, BarBQ Night",
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getInfoFromHTML(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGamePlaceAndAddressFromText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getGamePlaceAndAddressFromText() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getGamePlaceAndAddressFromText() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getInfoFromPopup(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name            string
		args            args
		dateTime        string
		gameDescription string
		paymentInfo     string
		wantErr         bool
	}{
		{
			name: "test case 1",
			args: args{
				html: ` 17 декабря 2022<br/><br/> 16:00<br/>  18:00<br/>Игра, предназначенная для команд без большого опыта в интеллектуальных играх.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "17 декабря 2022 16:00",
			gameDescription: "Игра, предназначенная для команд без большого опыта в интеллектуальных играх.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 2",
			args: args{
				html: ` 23 декабря 2022<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  1000 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "23 декабря 2022 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "1000 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 3",
			args: args{
				html: ` 24 декабря 2022<br/><br/> 16:00<br/>  18:00<br/> Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "24 декабря 2022 16:00",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 4",
			args: args{
				html: ` 28 декабря 2022<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "28 декабря 2022 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 5",
			args: args{
				html: ` 29 декабря 2022<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "29 декабря 2022 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 6",
			args: args{
				html: ` 30 декабря 2022<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "30 декабря 2022 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 7",
			args: args{
				html: ` 5 января 2023<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "5 января 2023 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 8",
			args: args{
				html: ` 6 января 2023<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "6 января 2023 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 9",
			args: args{
				html: ` 7 января 2023<br/><br/> 16:00<br/>  18:00<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "7 января 2023 16:00",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 10",
			args: args{
				html: ` 12 января 2023<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "12 января 2023 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 11",
			args: args{
				html: ` 13 января 2023<br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "13 января 2023 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 12",
			args: args{
				html: ` 13 января 2023<br/><br/><br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "",
			gameDescription: "",
			paymentInfo:     "",
			wantErr:         true,
		},
		{
			name: "test case 13",
			args: args{
				html: ` 13 января 2023      <br/><br/> 19:30<br/>  21:30<br/>Игра на общие темы. Самый популярный и массовый вариант.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "13 января 2023 19:30",
			gameDescription: "Игра на общие темы. Самый популярный и массовый вариант.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 14",
			args: args{
				html: ` 13 января 2023      <br/><br/> 19:30<br/>  21:30<br/>Игра со странным описанием.<br/>  400 рублей с человека, оплата только наличными<br/>`,
			},
			dateTime:        "13 января 2023 19:30",
			gameDescription: "Игра со странным описанием.",
			paymentInfo:     "400 рублей с человека, оплата только наличными",
			wantErr:         false,
		},
		{
			name: "test case 15",
			args: args{
				html: ``,
			},
			dateTime:        "",
			gameDescription: "",
			paymentInfo:     "",
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dateTime, gameDescription, gamePaymentInfo, err := getInfoFromPopup(context.Background(), tt.args.html)
			assert.Equal(t, tt.dateTime, dateTime)
			assert.Equal(t, tt.gameDescription, gameDescription)
			assert.Equal(t, tt.paymentInfo, gamePaymentInfo)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_getInfoFromStrong(t *testing.T) {
	type args struct {
		strong string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "test case 1",
			args: args{
				strong: "OnlineИгра 244",
			},
			want:  "Online",
			want1: "244",
		},
		{
			name: "test case 2",
			args: args{
				strong: "Игра 319.1",
			},
			want:  "",
			want1: "319.1",
		},
		{
			name: "test case 3",
			args: args{
				strong: "Online18+Игра 9",
			},
			want:  "Online18+",
			want1: "9",
		},
		{
			name: "test case 4",
			args: args{
				strong: "ТВ-шоуИгра 7",
			},
			want:  "ТВ-шоу",
			want1: "7",
		},
		{
			name: "test case 5",
			args: args{
				strong: "OnlineИгра 245",
			},
			want:  "Online",
			want1: "245",
		},
		{
			name: "test case 6",
			args: args{
				strong: "Сериалы. Кино. МузыкаИгра 35",
			},
			want:  "Сериалы. Кино. Музыка",
			want1: "35",
		},
		{
			name: "test case 7",
			args: args{
				strong: "OnlineЛига новичковИгра 13",
			},
			want:  "OnlineЛига новичков",
			want1: "13",
		},
		{
			name: "test case 8",
			args: args{
				strong: "Лига новичковИгра 3",
			},
			want:  "Лига новичков",
			want1: "3",
		},
		{
			name: "test case 9",
			args: args{
				strong: "OnlineСовместная игра с Pikabu",
			},
			want:  "",
			want1: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getInfoFromStrong(tt.args.strong)
			if got != tt.want {
				t.Errorf("getInfoFromStrong() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getInfoFromStrong() got1 = %v, want %v", got1, tt.want1)
			}
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

func convertTime(str string) time.Time {
	timeFormat := "2006-01-02 15:04"

	t, err := time.Parse(timeFormat, str)
	if err != nil {
		return time.Time{}
	}

	return t
}
