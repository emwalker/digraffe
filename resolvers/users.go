package resolvers

import (
	"context"
	"time"

	"github.com/emwalker/digraph/models"
)

type userResolver struct {
	*Resolver
}

// AvatarURL provides a link to a picture of the user.
func (r *userResolver) AvatarURL(_ context.Context, user *models.User) (string, error) {
	url := user.GithubAvatarURL.Ptr()
	if url != nil {
		return *url, nil
	}
	return "", nil
}

// CreatedAt returns of the creation of the user account.
func (r *userResolver) CreatedAt(_ context.Context, user *models.User) (string, error) {
	return user.CreatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) DefaultRepository(
	ctx context.Context, user *models.User,
) (*models.Repository, error) {
	return user.DefaultRepo(ctx, r.DB)
}

// Email returns the email of a user.
func (r *userResolver) PrimaryEmail(_ context.Context, user *models.User) (string, error) {
	return user.PrimaryEmail, nil
}

// Repositories returns the repositories to which the user has access
func (r *userResolver) Repositories(
	ctx context.Context, user *models.User, first *int, after *string, last *int, before *string,
) (models.RepositoryConnection, error) {
	var edges []*models.RepositoryEdge
	var err error
	var repos []*models.Repository

	selectedID := ""
	if id := user.SelectedRepositoryID.Ptr(); id != nil {
		selectedID = *id
	}

	if repos, err = user.OwnerRepositories().All(ctx, r.DB); err != nil {
		return models.RepositoryConnection{}, err
	}

	for _, repo := range repos {
		edges = append(edges, &models.RepositoryEdge{
			Node:       *repo,
			IsSelected: repo.ID == selectedID,
		})
	}

	return models.RepositoryConnection{Edges: edges}, nil
}

func (r *userResolver) SelectedRepository(
	ctx context.Context, user *models.User,
) (*models.Repository, error) {
	return user.SelectedRepository().One(ctx, r.DB)
}

// UpdatedAt returns the time of the most recent update.
func (r *userResolver) UpdatedAt(_ context.Context, user *models.User) (string, error) {
	return user.UpdatedAt.Format(time.RFC3339), nil
}
