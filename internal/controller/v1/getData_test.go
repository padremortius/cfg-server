package v1

import "testing"

func Test_getProfileAndName(t *testing.T) {
	type args struct {
		aName string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{name: "Test 1", args: args{aName: "dev-test-profile"}, want: "dev-test", want1: "profile"},
		{name: "Test 2", args: args{aName: "testapp"}, want: "testapp", want1: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getProfileAndName(tt.args.aName)
			if got != tt.want {
				t.Errorf("getProfileAndName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getProfileAndName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
