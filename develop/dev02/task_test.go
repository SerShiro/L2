package dev

import "testing"

func Test_unpackString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{s: "a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{s: "abcd"},
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "test3",
			args:    args{s: "45"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "test4",
			args:    args{s: ""},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpackString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unpackString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
