package vcards_test

import (
	"reflect"
	"testing"

	"github.com/lucasepe/cirql/internal/vcards"
)

func TestParamsMethods(t *testing.T) {
	tests := []struct {
		name       string
		setup      func() vcards.Params
		testAction func(p vcards.Params) any
		want       any
	}{
		{
			name: "Get returns first value",
			setup: func() vcards.Params {
				return vcards.Params{"TYPE": {"HOME", "WORK"}}
			},
			testAction: func(p vcards.Params) any {
				return p.Get("TYPE")
			},
			want: "HOME",
		},
		{
			name: "Get returns empty string if key not present",
			setup: func() vcards.Params {
				return vcards.Params{}
			},
			testAction: func(p vcards.Params) any {
				return p.Get("NOTFOUND")
			},
			want: "",
		},
		{
			name: "Add appends a value",
			setup: func() vcards.Params {
				p := vcards.Params{}
				p.Add("TYPE", "HOME")
				p.Add("TYPE", "WORK")
				return p
			},
			testAction: func(p vcards.Params) any {
				return p["TYPE"]
			},
			want: []string{"HOME", "WORK"},
		},
		{
			name: "Set replaces existing values",
			setup: func() vcards.Params {
				p := vcards.Params{}
				p.Add("TYPE", "HOME")
				p.Set("TYPE", "WORK")
				return p
			},
			testAction: func(p vcards.Params) any {
				return p["TYPE"]
			},
			want: []string{"WORK"},
		},
		{
			name: "Types returns lowercased type values",
			setup: func() vcards.Params {
				return vcards.Params{"TYPE": {"HOME", "WoRk"}}
			},
			testAction: func(p vcards.Params) any {
				return p.Types()
			},
			want: []string{"home", "work"},
		},
		{
			name: "HasType true when type exists (case-insensitive)",
			setup: func() vcards.Params {
				return vcards.Params{"TYPE": {"HOME", "WoRk"}}
			},
			testAction: func(p vcards.Params) any {
				return p.HasType("work")
			},
			want: true,
		},
		{
			name: "HasType false when type doesn't exist",
			setup: func() vcards.Params {
				return vcards.Params{"TYPE": {"HOME"}}
			},
			testAction: func(p vcards.Params) any {
				return p.HasType("work")
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := tt.setup()
			got := tt.testAction(params)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
