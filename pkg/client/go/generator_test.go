package clientgo

import (
	"os"
	"testing"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/google/go-cmp/cmp"
)

func TestGenerate(t *testing.T) {
	type args struct {
		gr                *parser.Group
		baseDirImportPath string
		packageName       string
	}
	group, err := parser.Parse("../../../server_generator/sample")

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		args     args
		wantPath string
		wantErr  bool
	}{
		{
			name:     "server_generator/sample",
			wantPath: "./sample/client.go",
			args: args{
				gr:                group,
				baseDirImportPath: "root",
				packageName:       "client",
			},
		},
	}
	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in function literal
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGenerator(tt.args.gr, tt.args.baseDirImportPath, tt.args.packageName, "1.0").GenerateClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantBytes, err := os.ReadFile(tt.wantPath)

			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, string(wantBytes)); diff != "" {
				t.Errorf("Generate() diff: %v", diff)
			}
		})
	}
}
