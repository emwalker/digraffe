package resolvers_test

import (
	"testing"

	"github.com/emwalker/digraph/cmd/frontend/models"
	"github.com/emwalker/digraph/cmd/frontend/resolvers"
	"github.com/emwalker/digraph/cmd/frontend/services"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func TestCreateSession(t *testing.T) {
	ctx := testContext()
	m := newMutator(t, testViewer)

	email := "gnusto@frotz.net"
	name := "Gnusto Frotz"
	login := "gfrotz"

	count, err := models.Users(qm.Where("primary_email = ?", email)).Count(ctx, testDB)
	if err != nil {
		t.Fatal(err)
	}

	if count > 0 {
		t.Fatalf("Expected there to be no users with the email %s", email)
	}

	input := models.CreateSessionInput{
		GithubAvatarURL: "https://some/url",
		GithubUsername:  login,
		Name:            name,
		PrimaryEmail:    email,
	}

	// Doesn't work if we do not have an admin session
	payload, err := m.resolver.CreateSession(ctx, input)
	if err != resolvers.ErrUnauthorized {
		t.Fatal(err)
	}

	// Works if we have an admin session
	rc := resolvers.GetRequestContext(ctx)
	rc.SetIsAdminSession(true)

	payload, err = m.resolver.CreateSession(ctx, input)
	if err != nil {
		t.Fatal(err)
	}

	user := payload.UserEdge.Node
	defer func() error {
		if _, err := user.Delete(m.ctx, m.db); err != nil {
			t.Fatal(err)
		}

		_, err = models.Organizations(qm.Where("login = ?", login)).DeleteAll(m.ctx, m.db)
		if err != nil {
			t.Fatal(err)
		}

		return err
	}()

	if payload.UserEdge == nil || payload.UserEdge.Node == nil {
		t.Fatal("There should be a user edge and a node")
	}

	if user.Name != name {
		t.Fatalf("Expected name to be %s, was %s instead", name, user.Name)
	}

	count, err = models.Users(qm.Where("primary_email = ?", email)).Count(m.ctx, m.db)
	if err != nil {
		t.Fatal(err)
	}

	if count < 1 {
		t.Fatalf("Expected there to be a new user with the email %s", email)
	}

	if payload.SessionEdge == nil {
		t.Fatal("A session should have been created")
	}

	resolvers.ClearRequestSession(ctx)
}

func TestDestroySession(t *testing.T) {
	ctx := testContext()

	username := testViewer.GithubUsername.Ptr()
	avatarURL := testViewer.GithubAvatarURL.Ptr()

	c := services.New(testDB, testViewer, nil)
	result, err := c.CreateSession(
		ctx, testViewer.Name, testViewer.PrimaryEmail, *username, *avatarURL,
	)

	if err != nil {
		t.Fatal(err)
	}
	defer result.Cleanup()

	session := result.Session

	m := newMutator(t, testViewer)

	input := models.DeleteSessionInput{SessionID: session.ID}
	payload, err := m.resolver.DeleteSession(m.ctx, input)

	if err != nil {
		t.Fatal(err)
	}

	if payload.DeletedSessionID != session.ID {
		t.Fatal("Expected a session id in the response")
	}
}
