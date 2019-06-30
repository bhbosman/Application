package StringDirectory

import (
	"reflect"
	"testing"
)

func TestNewStringRep(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *StringRep
	}{
		{
			name:,
			args: args{},
			want: nil,
		}
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringRep(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringRep() = %v, want %v", got, tt.want)
			}
		})
	}
}
