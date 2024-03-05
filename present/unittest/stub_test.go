package main

import (
	"errors"
	"testing"
)

type foorwareStub struct {
	err error
}

func (f foorwareStub) do() error {
	return f.err
}

func Test_personStub_do(t *testing.T) {
	type fields struct {
		fw footware
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Without error",
			fields: fields{
				fw: foorwareStub{err: nil},
			},
			wantErr: false,
		},
		{
			name: "With error",
			fields: fields{
				fw: foorwareStub{err: errors.New("error occur")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := personStub{
				fw: tt.fields.fw,
			}
			if err := p.do(); (err != nil) != tt.wantErr {
				t.Errorf("personStub.do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
