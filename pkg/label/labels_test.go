package label

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_labelSetToFingerprint(t *testing.T) {
	type args struct {
		ls LabelSet
	}
	tests := []struct {
		name string
		args args
		want Fingerprint
	}{
		{
			name: "test1",
			args: args{
				ls: LabelSet{
					"test1": "test1",
					"test2": "test2",
				},
			},
			want: Fingerprint(16292663137245379196),
		},
		{
			name: "test1",
			args: args{
				ls: LabelSet{
					"test2": "test2",
					"test1": "test1",
				},
			},
			want: Fingerprint(16292663137245379196),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := labelSetToFingerprint(tt.args.ls)
			if got != tt.want {
				t.Errorf("labelSetToFingerprint() = %v, want %v", got, tt.want)
			}
			var got2 Fingerprint
			assert.NoError(t, got2.Parse(got.String()))
			assert.Equal(t, tt.want, got2)
		})
	}
}
