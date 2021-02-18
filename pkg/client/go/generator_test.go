package clientgo

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/go-generalize/api_gen/pkg/parser"
)

func TestGenerate(t *testing.T) {
	type args struct {
		gr          *parser.Group
		packageName string
	}
	group, err := parser.Parse("../../../server_generator/sample")

	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		outputFile string
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "server_generator/sample",
			outputFile: "./sample/client.go",
			args: args{
				gr:          group,
				packageName: "client",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.gr, tt.args.packageName, "1.0")
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantWriter {
				fmt.Println(got)

				ioutil.WriteFile(tt.outputFile, []byte(got), 0644)

				t.Errorf("Generate() = %v, want %v", got, tt.wantWriter)
			}
		})
	}
}
