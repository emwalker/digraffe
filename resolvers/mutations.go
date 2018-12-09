package resolvers

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/emwalker/digraph/common"
	"github.com/emwalker/digraph/models"
	"github.com/emwalker/digraph/services"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func init() {
	log.SetOutput(os.Stdout)
}

// https://stackoverflow.com/a/23502629/61048
func transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return
}

// UpsertTopic creates a new topic.
func (r *MutationResolver) UpsertTopic(
	ctx context.Context, input models.UpsertTopicInput,
) (*models.UpsertTopicPayload, error) {
	var result *services.UpsertTopicResult
	var err error

	err = transact(r.DB, func(tx *sql.Tx) error {
		result, err = services.UpsertTopic(
			ctx,
			tx,
			input.OrganizationID,
			input.Name,
			input.Description,
			input.TopicIds,
		)

		return err
	})

	if err != nil {
		return nil, err
	}

	if result.Topic == nil {
		return &models.UpsertTopicPayload{Alerts: result.Alerts}, nil
	}

	return &models.UpsertTopicPayload{
		Alerts:    result.Alerts,
		TopicEdge: &models.TopicEdge{Node: *result.Topic},
	}, nil
}

// SelectTopic updates the currently selected topic.
func (r *MutationResolver) SelectTopic(
	ctx context.Context, input models.SelectTopicInput,
) (*models.SelectTopicPayload, error) {
	panic("not implemented")
}

// UpdateTopic updates the fields on a topic.
func (r *MutationResolver) UpdateTopic(
	ctx context.Context, input models.UpdateTopicInput,
) (*models.UpdateTopicPayload, error) {
	topic := models.Topic{
		OrganizationID: input.OrganizationID,
		Name:           input.Name,
		Description:    null.StringFromPtr(input.Description),
		ID:             input.ID,
	}

	_, err := topic.Update(ctx, r.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &models.UpdateTopicPayload{Topic: topic}, nil
}

// UpsertLink adds a new link to the database.
func (r *MutationResolver) UpsertLink(
	ctx context.Context, input models.UpsertLinkInput,
) (*models.UpsertLinkPayload, error) {
	var result *services.UpsertLinkResult
	var err error

	err = transact(r.DB, func(tx *sql.Tx) error {
		result, err = services.UpsertLink(
			ctx,
			tx,
			input.OrganizationID,
			input.URL,
			input.Title,
			input.AddParentTopicIds,
		)

		return err
	})

	if err != nil {
		return nil, err
	}

	return &models.UpsertLinkPayload{
		Alerts:   result.Alerts,
		LinkEdge: &models.LinkEdge{Node: result.Link},
	}, nil
}

// UpdateLinkTopics sets the parent topics on a link.
func (r *MutationResolver) UpdateLinkTopics(
	ctx context.Context, input models.UpdateLinkTopicsInput,
) (*models.UpdateLinkTopicsPayload, error) {
	link, err := models.FindLink(ctx, r.DB, input.LinkID)
	if err != nil {
		return nil, err
	}

	topics := common.TopicsFromIds(input.ParentTopicIds)
	if err = link.SetParentTopics(ctx, r.DB, false, topics...); err != nil {
		return nil, err
	}

	if err = link.Reload(ctx, r.DB); err != nil {
		return nil, err
	}

	return &models.UpdateLinkTopicsPayload{
		Link: *link,
	}, nil
}

// UpdateTopicParentTopics sets the parent topics on a topic.
func (r *MutationResolver) UpdateTopicParentTopics(
	ctx context.Context, input models.UpdateTopicParentTopicsInput,
) (*models.UpdateTopicParentTopicsPayload, error) {
	var result *services.UpdateTopicParentTopicsResult
	var topic *models.Topic
	var err error

	err = transact(r.DB, func(tx *sql.Tx) error {
		if topic, err = models.FindTopic(ctx, tx, input.TopicID); err != nil {
			return err
		}

		result, err = services.UpdateTopicParentTopics(ctx, tx, topic, input.ParentTopicIds)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &models.UpdateTopicParentTopicsPayload{
		Alerts: result.Alerts,
		Topic:  *result.Topic,
	}, nil
}
