package search_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/emwalker/digraph/cmd/frontend/models"
	"github.com/emwalker/digraph/cmd/frontend/queries"
	in "github.com/emwalker/digraph/test/integration"
)

func TestSearchInTopic(t *testing.T) {
	ctx := context.Background()
	mutator := in.NewMutator(in.MutatorOptions{})

	childTopic1 := mutator.MakeTopic(in.TopicOptions{
		ParentTopicIds: []string{in.Everything.ID},
		Name:           "Child topic 1",
	})

	childTopic2 := mutator.MakeTopic(in.TopicOptions{
		ParentTopicIds: []string{in.Everything.ID},
		Name:           "Child topic 2",
	})

	grandchildLink := mutator.MakeLink(in.LinkOptions{
		ParentTopicIds: []string{childTopic1.ID, childTopic2.ID},
		Title:          "New York Timely",
		URL:            "http://nytimely.com",
	})

	mutator.MakeLink(in.LinkOptions{
		ParentTopicIds: []string{childTopic1.ID},
		Title:          "Link with two parents",
		URL:            "http://link-with-two-parents.com",
	})

	grandchildTopic1 := mutator.MakeTopic(in.TopicOptions{
		Name:           "Grandchild topic 1abc",
		ParentTopicIds: []string{childTopic1.ID},
	})

	grandchildTopic2 := mutator.MakeTopic(in.TopicOptions{
		Name:           "Grandchild topic 2abc",
		ParentTopicIds: []string{childTopic2.ID},
	})

	greatGrandchildLink1 := mutator.MakeLink(in.LinkOptions{
		ParentTopicIds: []string{grandchildTopic1.ID},
		Title:          "Great granchild link 1",
		URL:            "http://great-grandchild-1.org",
	})

	greatGrandchildLink2 := mutator.MakeLink(in.LinkOptions{
		ParentTopicIds: []string{grandchildTopic1.ID, grandchildTopic2.ID},
		Title:          "Great granchild link 2",
		URL:            "http://great-grandchild-2.org",
	})

	testCases := []struct {
		name         string
		searchString string
		parentTopic  *models.TopicValue
		topics       in.Condition
		links        in.Condition
	}{
		{
			name:         "When an empty string is provided",
			searchString: "",
			parentTopic:  in.Everything,
			topics:       in.Condition{in.Exactly, 100},
			links:        in.Condition{in.Exactly, 100},
		},
		{
			name:         "When a link matches",
			searchString: "New York Timely",
			parentTopic:  childTopic1,
			topics:       in.Condition{in.Exactly, 0},
			links:        in.Condition{in.Exactly, 1},
		},
		{
			name:         "When a child topic matches",
			searchString: "Child topic",
			parentTopic:  in.Everything,
			topics:       in.Condition{in.Exactly, 3}, // FIXME: should be 2
			links:        in.Condition{in.Anything, 0},
		},
		{
			name:         "When a descendant topic matches",
			searchString: "Grandchild topic 1abc",
			parentTopic:  in.Everything,
			topics:       in.Condition{in.Exactly, 1},
			links:        in.Condition{in.Exactly, 0},
		},
		{
			name:         "When the search is for a URL and the parent topic is not the root",
			searchString: grandchildLink.URL,
			parentTopic:  childTopic1,
			topics:       in.Condition{in.Exactly, 0},
			links:        in.Condition{in.Exactly, 1},
		},
		{
			name:         "When the child topic is inlined in the query",
			searchString: fmt.Sprintf("in:/wiki/topics/%s %s", childTopic1.ID, grandchildLink.URL),
			parentTopic:  in.Everything,
			topics:       in.Condition{in.Exactly, 0},
			links:        in.Condition{in.Exactly, 1},
		},
		{
			name:         "When there is no such url",
			searchString: fmt.Sprintf("in:/wiki/topics/%s http://no-such-url", childTopic1.ID),
			parentTopic:  in.Everything,
			topics:       in.Condition{in.Exactly, 0},
			links:        in.Condition{in.Exactly, 0},
		},
		{
			name:         "When the url is a descendant url",
			searchString: fmt.Sprintf("in:/wiki/topics/%s %s", childTopic1.ID, greatGrandchildLink1.URL),
			parentTopic:  in.Everything,
			topics:       in.Condition{in.Exactly, 0},
			links:        in.Condition{in.Exactly, 1},
		},
		{
			name:         "When the child topic is inlined and a link title matches",
			searchString: fmt.Sprintf("in:/wiki/topics/%s New York Timely", childTopic1.ID),
			parentTopic:  in.Everything,
			topics:       in.Condition{in.Exactly, 0},
			links:        in.Condition{in.Exactly, 1},
		},
		{
			name: "When two topics are inlined",
			searchString: fmt.Sprintf(
				"in:/wiki/topics/%s in:/wiki/topics/%s",
				childTopic1.ID,
				childTopic2.ID,
			),
			parentTopic: in.Everything,
			topics:      in.Condition{in.Exactly, 0},
			// http://nytimely.com and http://great-grandchild-2.org
			links: in.Condition{in.Exactly, 1}, // FIXME: should be 2
		},
		{
			name: "Descendant links from the intersection of the transitive closures are included",
			searchString: fmt.Sprintf(
				"in:/wiki/topics/%s in:/wiki/topics/%s %s",
				childTopic1.ID,
				childTopic2.ID,
				greatGrandchildLink2.URL,
			),
			parentTopic: in.Everything,
			topics:      in.Condition{in.Exactly, 0},
			// http://great-grandchild-2.org
			links: in.Condition{in.Exactly, 0}, // FIXME: should be 1
		},
	}

	for _, td := range testCases {
		t.Run(td.name, func(t *testing.T) {
			query, err := queries.NewSearch(td.parentTopic, &td.searchString).Exec(ctx, mutator.DB)
			in.Must(err)

			topics, err := query.DescendantTopics(ctx, mutator.DB, 100)
			in.Must(err)

			if !td.topics.Evaluate(len(topics)) {
				t.Fatalf(td.topics.Describe(len(topics)))
			}

			links, err := query.DescendantLinks(ctx, mutator.DB, 100)
			in.Must(err)

			if !td.links.Evaluate(len(links)) {
				t.Fatalf(td.links.Describe(len(links)))
			}
		})
	}
}