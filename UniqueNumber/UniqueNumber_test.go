package UniqueNumber

import (
	"testing"
	"time"
)

func TestUniqueNumberGenerator_Next(t *testing.T) {
	type fields struct {
		number uint64
		now    func() time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "CheckNumber",
			fields: fields{
				number: 0,
				now: func() time.Time {
					dt, _ := time.Parse("20060102150405", "20060102150405")
					return dt
				},
			},
			want: "20060102150405.0000000001",
		},
		{
			name: "CheckNumber",
			fields: fields{
				number: 100,
				now: func() time.Time {
					dt, _ := time.Parse("20060102150405", "20190101000000")
					return dt
				},
			},
			want: "20190101000000.0000000101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &Generator{
				number: tt.fields.number,
				now:    tt.fields.now,
			}
			if got := self.Next(); got != tt.want {
				t.Errorf("UniqueNumberGenerator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
