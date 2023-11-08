package games

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := New(Config{})
		assert.NotNil(t, s)
	})
}

func TestSyncer_Enabled(t *testing.T) {
	type fields struct {
		disabledGamesFetchers []string
		enabled               bool
		fetchers              []Fetcher
		gameServiceClient     GameServiceClient
		name                  string
		period                time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "enabled",
			fields: fields{
				enabled: true,
			},
			want: true,
		},
		{
			name: "disabled",
			fields: fields{
				enabled: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syncer{
				disabledGamesFetchers: tt.fields.disabledGamesFetchers,
				enabled:               tt.fields.enabled,
				fetchers:              tt.fields.fetchers,
				gameServiceClient:     tt.fields.gameServiceClient,
				name:                  tt.fields.name,
				period:                tt.fields.period,
			}
			if got := s.Enabled(); got != tt.want {
				t.Errorf("Syncer.Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSyncer_GetPeriod(t *testing.T) {
	type fields struct {
		disabledGamesFetchers []string
		enabled               bool
		fetchers              []Fetcher
		gameServiceClient     GameServiceClient
		name                  string
		period                time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "tc1",
			fields: fields{
				period: 60 * time.Second,
			},
			want: 60 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syncer{
				disabledGamesFetchers: tt.fields.disabledGamesFetchers,
				enabled:               tt.fields.enabled,
				fetchers:              tt.fields.fetchers,
				gameServiceClient:     tt.fields.gameServiceClient,
				name:                  tt.fields.name,
				period:                tt.fields.period,
			}
			if got := s.GetPeriod(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Syncer.GetPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSyncer_GetName(t *testing.T) {
	type fields struct {
		disabledGamesFetchers []string
		enabled               bool
		fetchers              []Fetcher
		gameServiceClient     GameServiceClient
		name                  string
		period                time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "tc1",
			fields: fields{
				name: "name",
			},
			want: "name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Syncer{
				disabledGamesFetchers: tt.fields.disabledGamesFetchers,
				enabled:               tt.fields.enabled,
				fetchers:              tt.fields.fetchers,
				gameServiceClient:     tt.fields.gameServiceClient,
				name:                  tt.fields.name,
				period:                tt.fields.period,
			}
			if got := s.GetName(); got != tt.want {
				t.Errorf("Syncer.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
