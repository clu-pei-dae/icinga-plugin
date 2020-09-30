package icinga

import "testing"

func TestCreateOutputData(t *testing.T) {
	type args struct {
		message  string
		perfdata string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "simple test", args: args{
			message:  "Hello",
			perfdata: "World",
		}, want: "Hello | World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateOutputData(tt.args.message, tt.args.perfdata); got != tt.want {
				t.Errorf("CreateOutputData() = %v, want %v", got, tt.want)
			}
		})
	}
}
