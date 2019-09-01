package queries_test

import (
	"testing"

	"github.com/emwalker/digraph/cmd/frontend/queries"
	"github.com/volatiletech/sqlboiler/types"
)

func TestWildcardStringArray(t *testing.T) {
	testData := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "A simple case",
			input:  "York New",
			output: `{"%York%","%New%"}`,
		},
		{
			name:   "When there is a comma",
			input:  "York New,",
			output: `{"%York%","%New,%"}`,
		},
		{
			name:   "When there is a {",
			input:  "{York New",
			output: `{"%{York%","%New%"}`,
		},
		{
			name:   "When there is a }",
			input:  "York} New",
			output: `{"%York}%","%New%"}`,
		},
		{
			name:   "When there is a %",
			input:  "York% New",
			output: `{"%York%%","%New%"}`,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			q := queries.NewSearchQuery(td.input)

			actual, ok := q.WildcardStringArray().(*types.StringArray)
			if !ok {
				t.Fatalf("Expected a StringArray, got: %#v", actual)
			}

			value, err := actual.Value()
			if err != nil {
				t.Fatal(err)
			}

			if value != td.output {
				t.Fatalf("Expected %#v, got %#v", td.output, value)
			}
		})
	}
}