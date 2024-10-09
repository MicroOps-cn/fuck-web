package models

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToken_To(t *testing.T) {
	type data struct {
		Name string `json:"name"`
	}
	type args interface{}
	tests := []struct {
		name    string
		token   Token
		args    args
		wantErr bool
		want    interface{}
	}{
		{name: "test struct to struct", token: Token{Type: TokenTypeCode, Data: sql.RawBytes(`{"name":"lion"}`)}, args: &data{}, want: &data{Name: "lion"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.token.To(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("To() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.Equal(t, tt.args, tt.want)
		})
	}
}
