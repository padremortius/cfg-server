package common

import (
	"reflect"
	"testing"
)

func TestStructToJSONBytes(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "Test 1", args: args{v: 1}, want: []byte("1"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToJSONBytes(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToJSONBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToJSONBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToYamlBytes(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "Test 1", args: args{v: 1}, want: []byte("1\n"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToYamlBytes(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToYamlBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToYamlBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToTomlBytes(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "Test 1", args: args{v: 1}, want: []byte("1"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToTomlBytes(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToTomlBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToTomlBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
