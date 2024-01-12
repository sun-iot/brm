package util

import (
	"testing"
)

func TestGetGitLsRemote(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Official Source",
			args:    args{url: "https://github.com/Homebrew/brew.git"},
			wantErr: false,
		},
		{
			name:    "Aliyun Source",
			args:    args{url: "https://mirrors.aliyun.com/homebrew/brew.git"},
			wantErr: false,
		},
		{
			name:    "Tsinghua Source",
			args:    args{url: "https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git"},
			wantErr: false,
		},
		{
			name:    "USTC Source",
			args:    args{url: "https://mirrors.ustc.edu.cn/brew.git"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGitLsRemote(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGitLsRemote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Log(got)
		})
	}
}
