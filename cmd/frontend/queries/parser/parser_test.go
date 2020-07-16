package parser

import (
	"reflect"
	"testing"
)

func ptr(str string) *string {
	return &str
}

func TestParsing(t *testing.T) {
	testData := []struct {
		name         string
		input        string
		stringTokens []string
		topics       []TopicSpec
	}{
		{
			name:         "Two tokens",
			input:        "York New",
			stringTokens: []string{"York", "New"},
			topics:       []TopicSpec(nil),
		},
		{
			name:         "An empty string",
			input:        "",
			stringTokens: []string(nil),
			topics:       []TopicSpec(nil),
		},
		{
			name:         "A topic",
			input:        "in:/wiki/topics/1",
			stringTokens: []string(nil),
			topics:       []TopicSpec{{resourcePath: "/wiki/topics/1"}},
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			s := Parse(&td.input)

			if !reflect.DeepEqual(s.stringTokens, td.stringTokens) {
				t.Fatalf("Expected %#v, got %#v", td.stringTokens, s.stringTokens)
			}

			if !reflect.DeepEqual(s.Topics, td.topics) {
				t.Fatalf("Expected %#v, got %#v", td.topics, s.Topics)
			}
		})
	}
}