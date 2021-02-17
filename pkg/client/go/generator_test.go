package clientgo

import (
	"bytes"
	"testing"

	"github.com/go-generalize/api_gen/pkg/parser"
)

func TestGenerate(t *testing.T) {
	type args struct {
		gr *parser.Group
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := Generate(writer, tt.args.gr); (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Generate() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
