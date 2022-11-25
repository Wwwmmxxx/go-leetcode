package book01

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		// {
		// 	name: "success2",
		// 	args: args{
		// 		s: "a.",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "fail1",
		// 	args: args{
		// 		s: ".,",
		// 	},
		// 	want: false,
		// },
		{
			name: "fail2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
