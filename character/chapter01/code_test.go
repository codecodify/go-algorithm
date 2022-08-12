package chapter01

import "testing"

func Test_reformat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{
			name:    "测试数字字符数大于非数字字符数",
			args:    struct{ s string }{s: "abc1234"},
			wantRes: "1a2b3c4",
		},
		{
			name:    "测试数字字符数等于非数字字符数",
			args:    struct{ s string }{s: "abc123"},
			wantRes: "1a2b3c",
		},
		{
			name:    "测试数字字符数小于非数字字符数",
			args:    struct{ s string }{s: "abcd123"},
			wantRes: "a1b2c3d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := reformat(tt.args.s); gotRes != tt.wantRes {
				t.Errorf("reformat() = %v\n, want %v", gotRes, tt.wantRes)
			} else {
				t.Logf("reformat() = %v\n, want %v, test pass", gotRes, tt.wantRes)
			}
		})
	}
}
