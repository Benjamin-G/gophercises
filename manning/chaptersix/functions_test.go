package chaptersix

import (
	"io"
	"strings"
	"testing"
)

func TestReadFull(t *testing.T) {
	type args struct {
		r   io.Reader
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:  "Test 3",
			args:  args{r: strings.NewReader("Hell0"), buf: []byte{0xE6, 0xB1, 0x89}},
			wantN: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := ReadFull(tt.args.r, tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ReadFull() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_countEmptyLines(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "empty lines",
			args: args{reader: strings.NewReader(
				`foo
			bar

			baz
			`)},
			want: 0,
		},
		{
			name: "empty lines",
			args: args{reader: strings.NewReader(
				`foo
			bar

			baz
			`)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countEmptyLines(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("countEmptyLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("countEmptyLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}
