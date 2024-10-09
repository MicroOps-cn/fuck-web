package ldapservice

import "testing"

func Test_getPhoneNumberFilter(t *testing.T) {
	type args struct {
		name    string
		phoneNo string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "+86", args: args{name: "phoneNumber", phoneNo: "+8612345678"}, want: "(|(phoneNumber=+8612345678)(phoneNumber=12345678))"},
		{name: "+86 ", args: args{name: "phoneNumber", phoneNo: "+86 12345678"}, want: "(|(phoneNumber=+86 12345678)(phoneNumber=12345678))"},
		{name: "+86-", args: args{name: "phoneNumber", phoneNo: "+86-12345678"}, want: "(|(phoneNumber=+86-12345678)(phoneNumber=12345678))"},
		{name: "no", args: args{name: "phoneNumber", phoneNo: "12345678"}, want: "(phoneNumber=12345678)"},
		{name: "+86+86-+86", args: args{name: "phoneNumber", phoneNo: "+86+86-+8612345678"}, want: "(|(phoneNumber=+86+86-+8612345678)(phoneNumber=+86-+8612345678))"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPhoneNumberFilter(tt.args.name, tt.args.phoneNo); got != tt.want {
				t.Errorf("getPhoneNumberFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
