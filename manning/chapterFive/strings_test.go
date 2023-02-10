package chapterFive

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

var global string

func BenchmarkConcatV1(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat1(s)
	}
	global = local
}

func BenchmarkConcatV2(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat2(s)
	}
	global = local
}

func BenchmarkConcatV3(b *testing.B) {
	var local string
	s := getInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		local = concat3(s)
	}
	global = local
}

func getInput() []string {
	n := 1_000
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = string(make([]byte, 1_000))
	}
	return s
}

func Test_getBytes1(t *testing.T) {
	type args struct {
		reader io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Hell0 world",
			args: args{reader: strings.NewReader("Hell0 world    \n")},
			want: []byte("Hell0 world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBytes1(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBytes1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBytes1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBytes2(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Hell0 world",
			args: args{reader: strings.NewReader("Hell0 world    \n")},
			want: []byte("Hell0 world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBytes2(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBytes2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBytes2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
