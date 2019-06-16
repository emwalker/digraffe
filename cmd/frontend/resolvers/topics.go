package resolvers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	"github.com/emwalker/digraph/cmd/frontend/loaders"
	"github.com/emwalker/digraph/cmd/frontend/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

type topicResolver struct{ *Resolver }

// ByName provides a way to sort a topic by name.
type ByName []*models.Topic

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func getTopicLoader(ctx context.Context) *loaders.TopicLoader {
	return ctx.Value(loaders.TopicLoaderKey).(*loaders.TopicLoader)
}

func fetchTopic(
	ctx context.Context, exec boil.ContextExecutor, topicID string, actor *models.User,
) (*models.Topic, error) {
	topic, err := models.Topics(
		qm.InnerJoin("organization_members om on topics.organization_id = om.organization_id"),
		qm.Where("topics.id = ? and om.user_id = ?", topicID, actor.ID),
	).One(ctx, exec)
	return topic, err
}

func topicConnection(view *models.View, rows []*models.Topic, err error) (models.TopicConnection, error) {
	if err != nil {
		return models.TopicConnection{}, err
	}

	sort.Sort(ByName(rows))

	edges := make([]*models.TopicEdge, len(rows))
	for i, topic := range rows {
		edges[i] = &models.TopicEdge{Node: models.TopicValue{topic, false, view}}
	}

	return models.TopicConnection{Edges: edges}, nil
}

func topicRepository(ctx context.Context, topic *models.TopicValue) (*models.Repository, error) {
	if topic.R != nil && topic.R.Repository != nil {
		return topic.R.Repository, nil
	}
	repo, err := fetchRepository(ctx, topic.RepositoryID)
	return &repo, err
}

func topicOrganization(ctx context.Context, topic *models.TopicValue) (*models.Organization, error) {
	if topic.R != nil && topic.R.Organization != nil {
		return topic.R.Organization, nil
	}
	org, err := fetchOrganization(ctx, topic.OrganizationID)
	return &org, err
}

func availableTopics(
	ctx context.Context, exec boil.ContextExecutor, view *models.View, searchString *string, first *int,
	excludeTopicIds []string,
) (models.TopicConnection, error) {
	mods := []qm.QueryMod{
		qm.InnerJoin("organizations o on o.id = topics.organization_id"),
		qm.InnerJoin("organization_members om on om.organization_id = o.id"),
		qm.InnerJoin("users u on om.user_id = u.id"),
		qm.Where("u.id = ?", view.ViewerID),
		qm.OrderBy("topics.name"),
		qm.Limit(limitFrom(first)),
	}

	if searchString != nil {
		mods = append(mods,
			qm.Where("to_tsvector('synonymsdict', topics.synonyms) @@ to_tsquery(?)", wildcardStringArray(*searchString)),
		)
	}

	if len(excludeTopicIds) > 0 {
		mods = append(mods, qm.Where("topics.id != all(?)", types.Array(excludeTopicIds)))
	}

	topics, err := models.Topics(mods...).All(ctx, exec)
	if err != nil {
		return models.TopicConnection{}, err
	}

	return topicConnection(view, topics, err)
}

// AvailableParentTopics returns the topics that can be added to the link.
func (r *topicResolver) AvailableParentTopics(
	ctx context.Context, topic *models.TopicValue, searchString *string, first *int, after *string,
	last *int, before *string,
) (models.TopicConnection, error) {
	return availableTopics(ctx, r.DB, topic.View, searchString, first, []string{topic.ID})
}

// ChildTopics returns a set of topics.
func (r *topicResolver) ChildTopics(
	ctx context.Context, topic *models.TopicValue, searchString *string, first *int, after *string,
	last *int, before *string,
) (models.TopicConnection, error) {
	log.Printf("Fetching child topics for topic %s", topic.ID)

	mods := topic.View.Filter([]qm.QueryMod{
		qm.Load("ParentTopics"),
		qm.InnerJoin("repositories r on topics.repository_id = r.id"),
		qm.OrderBy("topics.name"),
	})

	if searchString != nil && *searchString != "" {
		mods = append(
			mods,
			qm.Where("to_tsvector('synonymsdict', topics.synonyms) @@ to_tsquery(?)", wildcardStringQuery(*searchString)),
		)
	}

	topics, err := topic.ChildTopics(mods...).All(ctx, r.DB)
	return topicConnection(topic.View, topics, err)
}

// CreatedAt returns the time of the topic's creation.
func (r *topicResolver) CreatedAt(_ context.Context, topic *models.TopicValue) (string, error) {
	return topic.CreatedAt.Format(time.RFC3339), nil
}

// Description returns a description of the topic.
func (r *topicResolver) Description(_ context.Context, topic *models.TopicValue) (*string, error) {
	return topic.Description.Ptr(), nil
}

// DisplayName returns the name of the topic.  The name is obtained by finding the first synonym
// in the current locale.  If there is no synonym in the current locale, the first English synonym
// is returned.  If there is no English synonym, the first synonym is returned.
func (r *topicResolver) DisplayName(_ context.Context, topic *models.TopicValue) (string, error) {
	synonyms, err := topic.SynonymList()
	if err != nil {
		return "<name missing>", err
	}

	name, ok := synonyms.NameForLocale(models.LocaleIdentifierEn)
	if !ok {
		return "<name missing>", errors.New("name not found")
	}

	return name, nil
}

func (r *topicResolver) ID(_ context.Context, topic *models.TopicValue) (string, error) {
	return topic.ID, nil
}

// Links returns a set of links.
func (r *topicResolver) Links(
	ctx context.Context, topic *models.TopicValue, searchString *string, first *int, after *string,
	last *int, before *string,
) (models.LinkConnection, error) {
	log.Printf("Fetching links for topic %s", topic.Summary())

	mods := topic.View.Filter([]qm.QueryMod{
		qm.Load("ParentTopics"),
		qm.OrderBy("links.created_at desc"),
		qm.InnerJoin("repositories r on links.repository_id = r.id"),
	})

	if !r.Actor.IsGuest() {
		mods = append(
			mods,
			qm.Load(
				models.LinkRels.UserLinkReviews, qm.Where("user_link_reviews.user_id = ?", r.Actor.ID), qm.Limit(1),
			),
		)
	}

	if searchString != nil && *searchString != "" {
		mods = append(mods, qm.Where("links.title ~~* all(?)", wildcardStringArray(*searchString)))
	}

	rows, err := topic.ChildLinks(mods...).All(ctx, r.DB)
	return linkConnection(topic.View, rows, len(rows), err)
}

// Loading is true if the topic is being loaded.  Only used on the client.
func (r *topicResolver) Loading(_ context.Context, topic *models.TopicValue) (bool, error) {
	return false, nil
}

func (r *topicResolver) Name(ctx context.Context, topic *models.TopicValue) (string, error) {
	return topic.Name, nil
}

// NewlyAdded returns true if the topic was just added.
func (r *topicResolver) NewlyAdded(_ context.Context, topic *models.TopicValue) (bool, error) {
	return topic.NewlyAdded, nil
}

// Organization returns an organization.
func (r *topicResolver) Organization(
	ctx context.Context, topic *models.TopicValue,
) (models.Organization, error) {
	org, err := topicOrganization(ctx, topic)
	return *org, err
}

// ParentTopics returns a set of topics.
func (r *topicResolver) ParentTopics(
	ctx context.Context, topic *models.TopicValue, first *int, after *string, last *int, before *string,
) (models.TopicConnection, error) {
	if topic.R != nil && topic.R.ParentTopics != nil {
		return topicConnection(topic.View, topic.R.ParentTopics, nil)
	}

	log.Printf("Fetching parent topics for topic %s", topic.ID)
	mods := topic.View.Filter([]qm.QueryMod{
		qm.InnerJoin("repositories r on topics.repository_id = r.id"),
		qm.OrderBy("topics.name"),
	})

	topics, err := topic.ParentTopics(mods...).All(ctx, r.DB)
	return topicConnection(topic.View, topics, err)
}

// Repository returns the repostory of the topic.
func (r *topicResolver) Repository(
	ctx context.Context, topic *models.TopicValue,
) (models.Repository, error) {
	repo, err := topicRepository(ctx, topic)
	return *repo, err
}

// ResourcePath returns a path to the item.
func (r *topicResolver) ResourcePath(ctx context.Context, topic *models.TopicValue) (string, error) {
	repo, err := topicRepository(ctx, topic)
	if err != nil {
		return "", err
	}

	org, err := topicOrganization(ctx, topic)
	if err != nil {
		return "", err
	}

	if repo.System {
		return fmt.Sprintf("/%s/topics/%s", org.Login, topic.ID), nil
	}
	return fmt.Sprintf("/%s/%s/topics/%s", org.Login, repo.Name, topic.ID), nil
}

func (r *topicResolver) matchingDescendantTopics(
	ctx context.Context, topic *models.TopicValue, searchString string, limit int,
) ([]*models.TopicValue, error) {
	var rows []struct {
		ID string
	}

	query := wildcardStringQuery(searchString)
	log.Printf("Looking for descendent topics with query: %s", query)

	err := queries.Raw(`
	with recursive child_topics as (
	  select parent_id, parent_id as child_id
	  from topic_topics where parent_id = $1
	union
	  select pt.child_id, ct.child_id
	  from topic_topics ct
	  inner join child_topics pt on pt.child_id = ct.parent_id
	)
	select distinct t.id
	from topics t
	inner join child_topics ct on ct.child_id = t.id
	where (
		case $2
		when '' then true
		else to_tsvector('synonymsdict', t.synonyms) @@ to_tsquery('synonymsdict', $2)
		end
	)
	limit $3
	`, topic.ID, query, limit).Bind(ctx, r.DB, &rows)

	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, nil
	}

	var ids []interface{}
	for _, row := range rows {
		ids = append(ids, row.ID)
	}

	mods := topic.View.Filter([]qm.QueryMod{
		qm.Load("ParentTopics"),
		qm.WhereIn("topics.id in ?", ids...),
		qm.InnerJoin("repositories r on topics.repository_id = r.id"),
	})

	topics, err := models.Topics(mods...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	var topicValues []*models.TopicValue
	for _, t := range topics {
		topicValues = append(topicValues, &models.TopicValue{t, false, topic.View})
	}

	return topicValues, nil
}

func (r *topicResolver) matchingDescendantLinks(
	ctx context.Context, topic *models.TopicValue, searchString string, limit int,
) ([]*models.LinkValue, error) {
	var rows []struct {
		ID string
	}

	query := wildcardStringQuery(searchString)
	log.Printf("Searching for links that match query: %s", query)

	err := queries.Raw(`
	with recursive child_topics as (
	  select parent_id, parent_id as child_id
	  from topic_topics where parent_id = $1
	union
	  select pt.child_id, ct.child_id
	  from topic_topics ct
	  inner join child_topics pt on pt.child_id = ct.parent_id
	)
	select l.id from links l
	inner join link_topics lt on l.id = lt.child_id
	inner join child_topics ct on ct.child_id = lt.parent_id
	where (
		case $2
		when '' then true
		else to_tsvector('linksdict', l.title) @@ to_tsquery('linksdict', $2)
		end
	)
	limit $3
	`, topic.ID, query, limit).Bind(ctx, r.DB, &rows)

	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, nil
	}

	var ids []interface{}
	for _, row := range rows {
		ids = append(ids, row.ID)
	}

	mods := topic.View.Filter([]qm.QueryMod{
		qm.Load("ParentTopics"),
		qm.WhereIn("links.id in ?", ids...),
		qm.InnerJoin("repositories r on links.repository_id = r.id"),
	})

	links, err := models.Links(mods...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	var linkValues []*models.LinkValue
	for _, l := range links {
		linkValues = append(linkValues, &models.LinkValue{l, false, topic.View})
	}

	return linkValues, nil
}

func (r *topicResolver) Search(
	ctx context.Context, topic *models.TopicValue, searchString string, first *int, after *string,
	last *int, before *string,
) (models.SearchResultItemConnection, error) {
	log.Printf("Searching topic %s for '%s'", topic.ID, searchString)

	var (
		err    error
		topics []*models.TopicValue
		links  []*models.LinkValue
		edges  []*models.SearchResultItemEdge
	)

	var limit = math.MaxInt32
	if first != nil {
		limit = *first
	}

	if topics, err = r.matchingDescendantTopics(ctx, topic, searchString, limit); err != nil {
		return models.SearchResultItemConnection{}, err
	}

	limit -= len(topics)
	if links, err = r.matchingDescendantLinks(ctx, topic, searchString, limit); err != nil {
		return models.SearchResultItemConnection{}, err
	}

	for _, topic := range topics {
		edges = append(edges, &models.SearchResultItemEdge{Node: *topic})
	}

	for _, link := range links {
		edges = append(edges, &models.SearchResultItemEdge{Node: *link})
	}

	return models.SearchResultItemConnection{Edges: edges}, nil
}

// Synonyms return the synonyms for this topic.
func (r *topicResolver) Synonyms(ctx context.Context, topic *models.TopicValue) ([]models.Synonym, error) {
	var out []models.Synonym

	synonyms, err := topic.SynonymList()
	if err != nil {
		return out, nil
	}

	for _, synonym := range synonyms.Values {
		out = append(out, models.Synonym{Locale: synonym.Locale, Name: synonym.Name})
	}

	return out, nil
}

// UpdatedAt returns the time of the most recent update.
func (r *topicResolver) UpdatedAt(_ context.Context, topic *models.TopicValue) (string, error) {
	return topic.UpdatedAt.Format(time.RFC3339), nil
}

// ViewerCanAddSynonym returns true if the viewer can add a synonym.
func (r *topicResolver) ViewerCanUpdate(ctx context.Context, topic *models.TopicValue) (bool, error) {
	log.Printf("Fetching value for ViewerCanAddSynonym for %s", topic.ID)
	mods := topic.View.Filter([]qm.QueryMod{
		qm.InnerJoin("repositories r on topics.repository_id = r.id"),
		qm.Where("topics.id = ?", topic.ID),
	})

	count, err := models.Topics(mods...).Count(ctx, r.DB)
	return count > 0, err
}

// ViewerCanAddSynonym returns true if the viewer can delete a synonym.
func (r *topicResolver) ViewerCanDeleteSynonyms(ctx context.Context, topic *models.TopicValue) (bool, error) {
	synonyms, err := topic.SynonymList()
	if err != nil {
		return false, err
	}

	if len(synonyms.Values) < 2 {
		return false, nil
	}

	return r.ViewerCanUpdate(ctx, topic)
}
