package resolvers_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/emwalker/digraph/models"
	"github.com/emwalker/digraph/resolvers"
	"github.com/emwalker/digraph/services"
	"github.com/emwalker/digraph/services/pageinfo"
	helpers "github.com/emwalker/digraph/testing"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const orgId = "45dc89a6-e6f0-11e8-8bc1-6f4d565e3ddb"

var (
	testDB    *sql.DB
	testActor *models.User
)

type testFetcher struct{}

type mutator struct {
	actor    *models.User
	ctx      context.Context
	db       *sql.DB
	resolver models.MutationResolver
	t        *testing.T
}

func (m mutator) defaultRepo() *models.Repository {
	repo, err := m.actor.OwnerRepositories(
		qm.InnerJoin("organizations o on o.id = repositories.organization_id"),
		qm.Where("repositories.system and o.login = ?", m.actor.Login),
	).One(m.ctx, testDB)
	if err != nil {
		panic(err)
	}
	return repo
}

func TestMain(m *testing.M) {
	services.Fetcher = &testFetcher{}
	testDB = newTestDb()
	defer testDB.Close()

	var err error

	testActor, err = models.Users().One(context.Background(), testDB)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func newTestDb() *sql.DB {
	var err error
	if testDB, err = sql.Open("postgres", "dbname=digraph_dev user=postgres sslmode=disable"); err != nil {
		log.Fatal("Unable to connect to the database", err)
	}
	return testDB
}

func newMutator(t *testing.T, actor *models.User) mutator {
	resolver := &resolvers.MutationResolver{
		&resolvers.Resolver{DB: testDB, Actor: actor},
	}

	return mutator{
		actor:    actor,
		ctx:      context.Background(),
		db:       testDB,
		resolver: resolver,
		t:        t,
	}
}

func (f *testFetcher) FetchPage(url string) (*pageinfo.PageInfo, error) {
	title := "Gnusto's blog"
	return &pageinfo.PageInfo{
		URL:   url,
		Title: &title,
	}, nil
}

func (m mutator) addParentTopicToTopic(child, parent *models.Topic) {
	everything, err := models.Topics(qm.Where("name like 'Everything'")).One(context.Background(), testDB)
	if err != nil {
		m.t.Fatal(err)
	}

	input := models.UpdateTopicParentTopicsInput{
		TopicID:        child.ID,
		ParentTopicIds: []string{everything.ID, parent.ID},
	}

	if _, err := m.resolver.UpdateTopicParentTopics(m.ctx, input); err != nil {
		m.t.Fatal(err)
	}
}

func (m mutator) addParentTopicToLink(link *models.Link, topic *models.Topic) {
	input := models.UpdateLinkTopicsInput{
		LinkID:         link.ID,
		ParentTopicIds: []string{topic.ID},
	}

	if _, err := m.resolver.UpdateLinkTopics(m.ctx, input); err != nil {
		m.t.Fatal(err)
	}
}

func (m mutator) deleteTopic(topic models.Topic) {
	count, err := topic.Delete(m.ctx, m.db)
	if err != nil {
		m.t.Fatal(err)
	}

	if count != int64(1) {
		m.t.Fatal("Expected a single row to be deleted")
	}
}

func (m mutator) createTopic(name string) (*models.Topic, helpers.CleanupFunc) {
	parentTopic, err := models.Topics(qm.Where("name like 'Everything'")).One(m.ctx, m.db)
	if err != nil {
		m.t.Fatal(err)
	}

	input := models.UpsertTopicInput{
		Name:              name,
		OrganizationLogin: m.actor.Login,
		RepositoryName:    m.defaultRepo().Name,
		TopicIds:          []string{parentTopic.ID},
	}

	payload, err := m.resolver.UpsertTopic(m.ctx, input)
	if err != nil {
		m.t.Fatal(err)
	}

	topic := payload.TopicEdge.Node

	cleanup := func() error {
		m.deleteTopic(topic)
		return nil
	}

	return &topic, cleanup
}

func (m mutator) createLink(title, url string) (*models.Link, helpers.CleanupFunc) {
	payload1, err := m.resolver.UpsertLink(m.ctx, models.UpsertLinkInput{
		AddParentTopicIds: []string{},
		OrganizationLogin: testActor.Login,
		RepositoryName:    m.defaultRepo().Name,
		Title:             &title,
		URL:               url,
	})
	if err != nil {
		m.t.Fatal(err)
	}

	link := payload1.LinkEdge.Node

	cleanup := func() error {
		count, err := link.Delete(m.ctx, testDB)
		if err != nil {
			m.t.Fatal(err)
		}

		if count != int64(1) {
			m.t.Fatal("Expected at least one updated record")
		}
		return nil
	}

	return &link, cleanup
}
