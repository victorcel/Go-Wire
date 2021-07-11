package di

import (
	"github.com/victorcel/pruebasGolan/models"
	"reflect"
	"testing"
)

func Test_providerMessage(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Message
		wantErr bool
	}{
		{
			name: "test_success",
			args: args{
				phrase: "prueba",
			},
			want:    models.NewMessage("prueba"),
			wantErr: false,
		},
		{
			name: "test_error",
			args: args{
				phrase: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := providerMessage(tt.args.phrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("providerUserRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("providerMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
