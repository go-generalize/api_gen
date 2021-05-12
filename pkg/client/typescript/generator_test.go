package clientts

import (
	"os"
	"testing"

	"github.com/go-generalize/api_gen/pkg/parser"
	"github.com/google/go-cmp/cmp"
)

func TestGenerate(t *testing.T) {
	type args struct {
		gr *parser.Group
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
			wantPath: "./testdata/api_client.ts",
			args: args{
				gr: group,
			},
		},
	}
	for _, tt := range tests {
		tt := tt // escape: Using the variable on range scope `tt` in function literal
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGenerator(tt.args.gr, "2.0").GenerateClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantBytes, err := os.ReadFile(tt.wantPath)

			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(string(wantBytes), got); diff != "" {
				t.Errorf("Generate() = %v, diff = %v", got, diff)
			}
		})
	}
}
