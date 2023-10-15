package squiz

import (
	"context"
	"testing"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game/squiz/v2/mocks"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
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

func convertTime(str string) time.Time {
	return time_utils.ConvertTime(str)
}

func Test_getGameDate(t *testing.T) {
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
				text: "",
			},
			want: "",
		},
		{
			name: "tc2",
			args: args{
				text: "Описание: \"ТВ-шоу\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра с раундами по мотивам легендарных ТВ-передач.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-19\"<br /><br />",
			},
			want: "2023-08-19",
		},
		{
			name: "tc3",
			args: args{
				text: "Описание: \"18+\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">Parkking</a>\", \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">Александровский парк, 4, корп. 3</a>\", \"59.955873,30.315782\"<br />Спецпроект: \"\"<br />Текст попапа: \"Вопросы для взрослых, черный юмор, без цензуры.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-10\"<br /><br />",
			},
			want: "2023-08-10",
		},
		{
			name: "tc4",
			args: args{
				text: "Описание: \"\"<br />Дополнение описания: \"повтор от 11 августа\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра на общие темы. Самый популярный вариант.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-12\"<br /><br />",
			},
			want: "2023-08-12",
		},
		{
			name: "tc5",
			args: args{
				text: "Описание: \"Гарри Поттер\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра для любителей книг и фильмов о «Мальчике, который выжил». Вас ждут волшебные раунды на внимательность, логику и знания.\"<br />SMS: \"Запись в резерв\"<br />Дата: \"2023-08-05\"<br /><br />",
			},
			want: "2023-08-05",
		},
		{
			name: "tc6",
			args: args{
				text: "Описание: \"Битва городов\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра на общие темы. Один пакет вопросов играем одновременно в разных городах.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-11\"<br /><br />",
			},
			want: "2023-08-11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameDate(tt.args.text); got != tt.want {
				t.Errorf("getGameDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGameDescription(t *testing.T) {
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
				text: "",
			},
			want: "",
		},
		{
			name: "tc2",
			args: args{
				text: "Описание: \"ТВ-шоу\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра с раундами по мотивам легендарных ТВ-передач.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-19\"<br /><br />",
			},
			want: "Игра с раундами по мотивам легендарных ТВ-передач.",
		},
		{
			name: "tc3",
			args: args{
				text: "Описание: \"18+\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">Parkking</a>\", \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">Александровский парк, 4, корп. 3</a>\", \"59.955873,30.315782\"<br />Спецпроект: \"\"<br />Текст попапа: \"Вопросы для взрослых, черный юмор, без цензуры.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-10\"<br /><br />",
			},
			want: "Вопросы для взрослых, черный юмор, без цензуры.",
		},
		{
			name: "tc4",
			args: args{
				text: "Описание: \"\"<br />Дополнение описания: \"повтор от 11 августа\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра на общие темы. Самый популярный вариант.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-12\"<br /><br />",
			},
			want: "Игра на общие темы. Самый популярный вариант.",
		},
		{
			name: "tc5",
			args: args{
				text: "Описание: \"Гарри Поттер\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра для любителей книг и фильмов о «Мальчике, который выжил». Вас ждут волшебные раунды на внимательность, логику и знания.\"<br />SMS: \"Запись в резерв\"<br />Дата: \"2023-08-05\"<br /><br />",
			},
			want: "Игра для любителей книг и фильмов о «Мальчике, который выжил». Вас ждут волшебные раунды на внимательность, логику и знания.",
		},
		{
			name: "tc6",
			args: args{
				text: "Описание: \"Битва городов\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра на общие темы. Один пакет вопросов играем одновременно в разных городах.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-11\"<br /><br />",
			},
			want: "Игра на общие темы. Один пакет вопросов играем одновременно в разных городах.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameDescription(tt.args.text); got != tt.want {
				t.Errorf("getGameDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_getGameFormat(t *testing.T) {
	type args struct {
		characteristics []Characteristic
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				characteristics: []Characteristic{
					{
						Title: "Формат игры",
						Value: "Онлайн",
					},
				},
			},
			want: "Онлайн",
		},
		{
			name: "tc2",
			args: args{
				characteristics: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameFormat(tt.args.characteristics); got != tt.want {
				t.Errorf("getGameFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGameName(t *testing.T) {
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
				text: "",
			},
			want: "",
		},
		{
			name: "tc2",
			args: args{
				text: "Описание: \"ТВ-шоу\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра с раундами по мотивам легендарных ТВ-передач.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-19\"<br /><br />",
			},
			want: "ТВ-шоу",
		},
		{
			name: "tc3",
			args: args{
				text: "Описание: \"\"<br />Дополнение описания: \"повтор от 17 августа\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра на общие темы. Самый популярный вариант.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-18\"<br /><br />",
			},
			want: "",
		},
		{
			name: "tc4",
			args: args{
				text: "Описание: \"Битва городов\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра на общие темы. Один пакет вопросов играем одновременно в разных городах.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-11\"<br /><br />",
			},
			want: "Битва городов",
		},
		{
			name: "tc5",
			args: args{
				text: "Описание: \"18+\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">Parkking</a>\", \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">Александровский парк, 4, корп. 3</a>\", \"59.955873,30.315782\"<br />Спецпроект: \"\"<br />Текст попапа: \"Вопросы для взрослых, черный юмор, без цензуры.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-10\"<br /><br />",
			},
			want: "18+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameName(tt.args.text); got != tt.want {
				t.Errorf("getGameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGamePlaceAddress(t *testing.T) {
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
				text: "",
			},
			want: "",
		},
		{
			name: "tc2",
			args: args{
				text: "Описание: \"ТВ-шоу\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра с раундами по мотивам легендарных ТВ-передач.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-19\"<br /><br />",
			},
			want: "ул. Ломоносова, 16",
		},
		{
			name: "tc3",
			args: args{
				text: "Описание: \"18+\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">Parkking</a>\", \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">Александровский парк, 4, корп. 3</a>\", \"59.955873,30.315782\"<br />Спецпроект: \"\"<br />Текст попапа: \"Вопросы для взрослых, черный юмор, без цензуры.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-10\"<br /><br />",
			},
			want: "Александровский парк, 4, корп. 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGamePlaceAddress(tt.args.text); got != tt.want {
				t.Errorf("getGamePlaceAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGamePlaceName(t *testing.T) {
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
				text: "",
			},
			want: "",
		},
		{
			name: "tc2",
			args: args{
				text: "Описание: \"ТВ-шоу\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \" <a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">BarBQ Night</a>\", \"<a href=\"https://yandex.ru/maps/org/barbq_night/133678003397/?ll=30.338599%2C59.927607&amp;z=15\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">ул. Ломоносова, 16</a>\", \"59.927621, 30.338757\"<br />Спецпроект: \"\"<br />Текст попапа: \"Игра с раундами по мотивам легендарных ТВ-передач.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-19\"<br /><br />",
			},
			want: "BarBQ Night",
		},
		{
			name: "tc3",
			args: args{
				text: "Описание: \"18+\"<br />Дополнение описания: \"\"<br />Особенности: \"\"<br />Локация: \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(255, 255, 255);\">Parkking</a>\", \"<a href=\"https://yandex.ru/maps/org/parkking/1281972478/?ll=30.315706%2C59.956148&amp;z=14\" target=\"_blank\" rel=\"noreferrer noopener\" style=\"color: rgb(177, 177, 177);\">Александровский парк, 4, корп. 3</a>\", \"59.955873,30.315782\"<br />Спецпроект: \"\"<br />Текст попапа: \"Вопросы для взрослых, черный юмор, без цензуры.\"<br />SMS: \"Запись на игру\"<br />Дата: \"2023-08-10\"<br /><br />",
			},
			want: "Parkking",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGamePlaceName(tt.args.text); got != tt.want {
				t.Errorf("getGamePlaceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGameTime(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				title: "5 августа, сб :: 16:00",
			},
			want: "16:00",
		},
		{
			name: "tc2",
			args: args{
				title: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameTime(tt.args.title); got != tt.want {
				t.Errorf("getGameTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
