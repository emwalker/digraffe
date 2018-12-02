package resolvers_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/emwalker/digraph/models"
	"github.com/emwalker/digraph/resolvers"
	"github.com/emwalker/digraph/resolvers/pageinfo"
)


const orgId = "45dc89a6-e6f0-11e8-8bc1-6f4d565e3ddb"

var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB = newTestDb()
	defer testDB.Close()
	os.Exit(m.Run())
}

func newTestDb() *sql.DB {
	var err error
	if testDB, err = sql.Open("postgres", "dbname=digraph_dev user=postgres sslmode=disable"); err != nil {
		log.Fatal("Unable to connect to the database", err)
	}
	return testDB
}

type testFetcher struct{}

func startMutationTest(t *testing.T, db *sql.DB) (models.MutationResolver, context.Context) {
	resolver := &resolvers.MutationResolver{
		&resolvers.Resolver{DB: db},
		&testFetcher{},
	}
	return resolver, context.Background()
}

func (f *testFetcher) FetchPage(url string) (*pageinfo.PageInfo, error) {
	title := "Gnusto's blog"
	return &pageinfo.PageInfo{
		URL:   url,
		Title: &title,
	}, nil
}