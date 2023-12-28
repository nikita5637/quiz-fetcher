package shaker

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
				dateTime: "",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				dateTime: "2023-12-26T19:30:00",
			},
			want:    time_utils.ConvertTime("2023-12-26 16:30"),
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
