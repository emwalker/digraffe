// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type ActivityLineItem struct {
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

type ActivityLineItemConnection struct {
	Edges    []*ActivityLineItemEdge `json:"edges"`
	PageInfo PageInfo                `json:"pageInfo"`
}

type ActivityLineItemEdge struct {
	Cursor string           `json:"cursor"`
	Node   ActivityLineItem `json:"node"`
}

type Alert struct {
	Text string    `json:"text"`
	Type AlertType `json:"type"`
	ID   string    `json:"id"`
}

type Alertable interface {
	IsAlertable()
}

type DeleteLinkInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	LinkID           string  `json:"linkId"`
}

type DeleteLinkPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	DeletedLinkID    string  `json:"deletedLinkId"`
}

type DeleteTopicInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	TopicID          string  `json:"topicId"`
}

type DeleteTopicPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	DeletedTopicID   string  `json:"deletedTopicId"`
}

type LinkConnection struct {
	Edges    []*LinkEdge `json:"edges"`
	PageInfo PageInfo    `json:"pageInfo"`
}

type LinkEdge struct {
	Cursor string    `json:"cursor"`
	Node   LinkValue `json:"node"`
}

type Namespaceable interface {
	IsNamespaceable()
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type RepositoryConnection struct {
	Edges []*RepositoryEdge `json:"edges"`
}

type RepositoryEdge struct {
	Cursor     string     `json:"cursor"`
	Node       Repository `json:"node"`
	IsSelected bool       `json:"isSelected"`
}

type ResourceIdentifiable interface {
	IsResourceIdentifiable()
}

type SearchResultItem interface {
	IsSearchResultItem()
}

type SearchResultItemConnection struct {
	Edges []*SearchResultItemEdge `json:"edges"`
}

type SearchResultItemEdge struct {
	Node SearchResultItem `json:"node"`
}

type SelectRepositoryInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	RepositoryID     *string `json:"repositoryId"`
}

type SelectRepositoryPayload struct {
	Repository *Repository `json:"repository"`
	Viewer     User        `json:"viewer"`
}

type TopicConnection struct {
	Edges    []*TopicEdge `json:"edges"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type TopicEdge struct {
	Cursor string     `json:"cursor"`
	Node   TopicValue `json:"node"`
}

type UpdateLinkTopicsInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	LinkID           string   `json:"linkId"`
	ParentTopicIds   []string `json:"parentTopicIds"`
}

type UpdateLinkTopicsPayload struct {
	Link LinkValue `json:"link"`
}

type UpdateTopicInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	Description      *string  `json:"description"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	TopicIds         []string `json:"topicIds"`
}

type UpdateTopicParentTopicsInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	TopicID          string   `json:"topicId"`
	ParentTopicIds   []string `json:"parentTopicIds"`
}

type UpdateTopicParentTopicsPayload struct {
	Alerts []Alert    `json:"alerts"`
	Topic  TopicValue `json:"topic"`
}

func (UpdateTopicParentTopicsPayload) IsAlertable() {}

type UpdateTopicPayload struct {
	Alerts []Alert    `json:"alerts"`
	Topic  TopicValue `json:"topic"`
}

type UpsertLinkInput struct {
	AddParentTopicIds []string `json:"addParentTopicIds"`
	ClientMutationID  *string  `json:"clientMutationId"`
	OrganizationLogin string   `json:"organizationLogin"`
	RepositoryName    string   `json:"repositoryName"`
	Title             *string  `json:"title"`
	URL               string   `json:"url"`
}

type UpsertLinkPayload struct {
	Alerts   []Alert   `json:"alerts"`
	LinkEdge *LinkEdge `json:"linkEdge"`
}

func (UpsertLinkPayload) IsAlertable() {}

type UpsertTopicInput struct {
	ClientMutationID  *string  `json:"clientMutationId"`
	Description       *string  `json:"description"`
	Name              string   `json:"name"`
	OrganizationLogin string   `json:"organizationLogin"`
	RepositoryName    string   `json:"repositoryName"`
	TopicIds          []string `json:"topicIds"`
}

type UpsertTopicPayload struct {
	Alerts    []Alert    `json:"alerts"`
	TopicEdge *TopicEdge `json:"topicEdge"`
}

func (UpsertTopicPayload) IsAlertable() {}

type AlertType string

const (
	AlertTypeSuccess AlertType = "SUCCESS"
	AlertTypeWarn    AlertType = "WARN"
	AlertTypeError   AlertType = "ERROR"
)

func (e AlertType) IsValid() bool {
	switch e {
	case AlertTypeSuccess, AlertTypeWarn, AlertTypeError:
		return true
	}
	return false
}

func (e AlertType) String() string {
	return string(e)
}

func (e *AlertType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlertType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlertType", str)
	}
	return nil
}

func (e AlertType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
