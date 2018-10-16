package token

import (
	"crypto/rsa"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/agungdwiprasetyo/agungkiki-backend/config"
	model "github.com/agungdwiprasetyo/agungkiki-backend/src/model"
)

func TestNew(t *testing.T) {
	conf := config.New()
	_, filename, _, _ := runtime.Caller(0)
	paths := strings.Split(filename, "/")
	path := strings.Join(paths[:len(paths)-2], "/")
	os.Setenv("APP_PATH", path)

	type args struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
		age        time.Duration
	}
	tests := []struct {
		name string
		args args
		want *Token
	}{
		{
			name: "Testcase #1: Positive",
			args: args{
				privateKey: conf.LoadPrivateKey(),
				publicKey:  conf.LoadPublicKey(),
				age:        12 * time.Hour,
			},
			want: &Token{
				privateKey: conf.LoadPrivateKey(),
				publicKey:  conf.LoadPublicKey(),
				age:        12 * time.Hour,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.privateKey, tt.args.publicKey, tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToken_Generate(t *testing.T) {
	conf := config.New()
	_, filename, _, _ := runtime.Caller(0)
	paths := strings.Split(filename, "/")
	path := strings.Join(paths[:len(paths)-2], "/")
	os.Setenv("APP_PATH", path)

	type fields struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
		age        time.Duration
	}
	type args struct {
		cl *Claim
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Testcase #1: Positive",
			fields: fields{
				privateKey: conf.LoadPrivateKey(),
				publicKey:  conf.LoadPublicKey(),
				age:        12 * time.Hour,
			},
			args: args{
				cl: &Claim{
					Audience: "user",
					User: &model.User{
						Name: "agung",
						Role: &model.Role{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &Token{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
				age:        tt.fields.age,
			}
			_, err := tok.Generate(tt.args.cl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Token.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestToken_Extract(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	paths := strings.Split(filename, "/")
	path := strings.Join(paths[:len(paths)-2], "/")
	os.Setenv("APP_PATH", path)
	conf := config.New()

	type fields struct {
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
		age        time.Duration
	}

	type args struct {
		tokenString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  *Claim
	}{
		{
			name: "Testcase #1: Positive",
			fields: fields{
				privateKey: conf.LoadPrivateKey(),
				publicKey:  conf.LoadPublicKey(),
				age:        12 * time.Hour,
			},
			args: args{
				tokenString: "asdadadsa",
			},
			want:  false,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &Token{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
				age:        tt.fields.age,
			}
			got, got1 := tok.Extract(tt.args.tokenString)
			if got != tt.want {
				t.Errorf("Token.Extract() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Token.Extract() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
