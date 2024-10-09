package opts

import "github.com/MicroOps-cn/fuck-web/pkg/errors"

type WithGetUserOptions func(o *GetUserOptions)

type GetUserOptions struct {
	Id          string
	Username    string
	Email       string
	PhoneNumber string
	Ext         bool
	Err         error
}

func WithPhoneNumber(no string) WithGetUserOptions {
	return func(o *GetUserOptions) {
		o.PhoneNumber = no
		if len(o.PhoneNumber) == 0 && o.Err != nil {
			o.Err = errors.LackParameterError("phoneNumber")
		}
	}
}

func WithUsername(username string) WithGetUserOptions {
	return func(o *GetUserOptions) {
		o.Username = username
		if len(o.Username) == 0 && o.Err != nil {
			o.Err = errors.LackParameterError("username")
		}
	}
}

func WithUserExt(o *GetUserOptions) {
	o.Ext = true
}

func WithoutUserExt(o *GetUserOptions) {
	o.Ext = false
}

func WithUserId(id string) WithGetUserOptions {
	return func(o *GetUserOptions) {
		o.Id = id
		if len(o.Id) == 0 && o.Err != nil {
			o.Err = errors.LackParameterError("id")
		}
	}
}

func WithEmail(email string) WithGetUserOptions {
	return func(o *GetUserOptions) {
		o.Email = email
		if len(o.Email) == 0 && o.Err != nil {
			o.Err = errors.LackParameterError("email")
		}
	}
}

func NewGetUserOptions(opts ...WithGetUserOptions) *GetUserOptions {
	o := &GetUserOptions{}
	for _, option := range opts {
		option(o)
	}
	return o
}
