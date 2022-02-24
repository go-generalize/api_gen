package clientgo

import (
	"os"
	"testing"

	"github.com/go-generalize/api_gen/v2/pkg/parser"
	"github.com/google/go-cmp/cmp"
)

func TestGenerate(t *testing.T) {
	type args struct {
		gr                *parser.Group
		baseDirImportPath string
		packageName       string
	}
	group, err := parser.Parse("../../../samples/standard/api")

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
			name:     "samples/standard",
			wantPath: "./sample/client.go.want",
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
			got, err := NewGenerator(tt.args.gr, tt.args.baseDirImportPath, tt.args.packageName, "devel").GenerateClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantBytes, err := os.ReadFile(tt.wantPath)

			if err != nil {
				// nolint:errcheck
				os.WriteFile(tt.wantPath, []byte(got), 0774)
				t.Fatal(err)
			}

			if diff := cmp.Diff(string(wantBytes), got); diff != "" {
				t.Errorf("Generate() diff: %v", diff)
			}
		})
	}
}
