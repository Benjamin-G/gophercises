package chapterSeven

import (
	"errors"
	"reflect"
	"testing"
)

func Test_runError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "Test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runError()
			if _, isErr := err.(barError); isErr {
				t.Errorf("runError() error = %[1]t %[1]v, wantErr %[2]t", err, barError{})
			}
		})
	}
}

func Test_runError2(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "Test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := runError2()

			if errors.Is(err, transientError{}) {
				t.Errorf("runError() error = %[1]t %[1]v, wantErr %[2]t", err, transientError{})
			}
		})
	}
}

func TestGetRoute1(t *testing.T) {
	type args struct {
		srcLat float32
		srcLng float32
		dstLat float32
		dstLng float32
	}
	tests := []struct {
		name    string
		args    args
		want    Route
		wantErr bool
	}{
		{
			name: "TestGetRoute1",
			args: args{srcLat: float32(100),
				srcLng: 100,
				dstLat: 100,
				dstLng: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoute1(tt.args.srcLat, tt.args.srcLng, tt.args.dstLat, tt.args.dstLng)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoute1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoute1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRoute3(t *testing.T) {
	type args struct {
		srcLat float32
		srcLng float32
		dstLat float32
		dstLng float32
	}
	tests := []struct {
		name    string
		args    args
		want    Route
		wantErr bool
	}{
		{
			name: "TestGetRoute1",
			args: args{srcLat: float32(100),
				srcLng: 100,
				dstLat: 100,
				dstLng: 100},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoute3(tt.args.srcLat, tt.args.srcLng, tt.args.dstLat, tt.args.dstLng)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRoute3() error = %[1]v ,\n wantErr %[2]v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoute3() got = %v, want %v", got, tt.want)
			}
		})
	}
}
