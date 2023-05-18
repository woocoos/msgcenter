package label

import "testing"

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
			if got := labelSetToFingerprint(tt.args.ls); got != tt.want {
				t.Errorf("labelSetToFingerprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
