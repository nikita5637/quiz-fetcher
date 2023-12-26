package sixty_seconds

import (
	"reflect"
	"testing"
	"time"

	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

func Test_convertDateTime(t *testing.T) {
	type args struct {
		dateTime string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				dateTime: "                            9 Jan 2024, 7:30 p.m.",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				dateTime: "9 Jan 2024, 7:30 p.m.",
			},
			want:    time_utils.ConvertTime("2024-01-09 16:30"),
			wantErr: false,
		},
		{
			name: "ok",
			args: args{
				dateTime: "9 января 2024 г. 19:30",
			},
			want:    time_utils.ConvertTime("2024-01-09 16:30"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertDateTime(tt.args.dateTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertDateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
