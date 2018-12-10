package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"net/url"

	pl "github.com/PuerkitoBio/purell"
	"github.com/emwalker/digraph/common"
	"github.com/emwalker/digraph/models"
	"github.com/emwalker/digraph/services/pageinfo"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UpsertLinkResult struct {
	Alerts      []models.Alert
	Link        *models.Link
	LinkCreated bool
}

type URL struct {
	CanonicalURL string
	Input        string
	Sha1         string
}

const normalizationFlags = pl.FlagRemoveDefaultPort |
	pl.FlagDecodeDWORDHost |
	pl.FlagDecodeOctalHost |
	pl.FlagDecodeHexHost |
	pl.FlagRemoveUnnecessaryHostDots |
	pl.FlagRemoveDotSegments |
	pl.FlagRemoveDuplicateSlashes |
	pl.FlagUppercaseEscapes |
	pl.FlagDecodeUnnecessaryEscapes |
	pl.FlagEncodeNecessaryEscapes |
	pl.FlagSortQuery

var Fetcher pageinfo.Fetcher = &pageinfo.HtmlFetcher{}

func NormalizeUrl(url string) (*URL, error) {
	canonical, err := pl.NormalizeURLString(url, normalizationFlags)
	if err != nil {
		return nil, err
	}

	sha1 := fmt.Sprintf("%x", sha1.Sum([]byte(canonical)))
	return &URL{canonical, url, sha1}, nil
}

func providedOrFetchedTitle(url string, providedTitle *string) (string, error) {
	if providedTitle != nil && *providedTitle != "" {
		return *providedTitle, nil
	}

	log.Print("Fetching title of ", url)
	pageInfo, err := Fetcher.FetchPage(url)
	if err != nil {
		return "", err
	}

	if pageInfo.Title != nil {
		return *pageInfo.Title, nil
	}

	return "", nil
}

func isURL(name string) bool {
	_, err := url.ParseRequestURI(name)
	if err != nil {
		return false
	}
	return true
}

func (c Connection) addParentTopicsToLink(
	ctx context.Context, link models.Link, parentTopicIds []string,
) error {
	if len(parentTopicIds) < 1 {
		return nil
	}

	var topicIds []interface{}
	for _, topicID := range parentTopicIds {
		topicIds = append(topicIds, topicID)
	}

	overlappingTopics, err := link.ParentTopics(
		qm.Select("id"),
		qm.WhereIn("id in ?", topicIds...),
	).All(ctx, c.Exec)

	if err != nil {
		return err
	}

	seen := make(map[string]bool)
	for _, topic := range overlappingTopics {
		seen[topic.ID] = true
	}

	var insertIds []string
	for _, topicID := range parentTopicIds {
		if _, ok := seen[topicID]; !ok {
			insertIds = append(insertIds, topicID)
		}
	}

	if len(insertIds) < 1 {
		return nil
	}

	topics := common.TopicsFromIds(insertIds)
	return link.AddParentTopics(ctx, c.Exec, false, topics...)
}

func (c Connection) UpsertLink(
	ctx context.Context, repo *models.Repository, providedUrl string, providedTitle *string,
	parentTopicIds []string,
) (*UpsertLinkResult, error) {
	var alerts []models.Alert

	url, err := NormalizeUrl(providedUrl)
	if err != nil {
		return nil, err
	}

	title, err := providedOrFetchedTitle(url.CanonicalURL, providedTitle)
	if err != nil {
		alerts = append(alerts,
			*models.NewAlert(models.AlertTypeWarn, fmt.Sprintf("Not a valid link: %s", providedUrl)),
		)
		return &UpsertLinkResult{Alerts: alerts}, nil
	}

	link := models.Link{
		OrganizationID: repo.OrganizationID,
		RepositoryID:   repo.ID,
		Sha1:           url.Sha1,
		Title:          title,
		URL:            url.CanonicalURL,
	}

	existing, err := repo.Links(qm.Where("sha1 like ?", url.Sha1)).Count(ctx, c.Exec)
	if err != nil {
		return nil, err
	}

	err = link.Upsert(
		ctx,
		c.Exec,
		true,
		[]string{"repository_id", "sha1"},
		boil.Whitelist("url", "title"),
		boil.Infer(),
	)

	if err != nil {
		return nil, err
	}

	err = c.addParentTopicsToLink(ctx, link, parentTopicIds)
	if err != nil {
		return nil, err
	}

	if existing > 0 {
		alerts = []models.Alert{
			*models.NewAlert(models.AlertTypeSuccess, fmt.Sprintf("An existing link %s was found", providedUrl)),
		}
	}

	return &UpsertLinkResult{
		Alerts:      alerts,
		Link:        &link,
		LinkCreated: existing < 1,
	}, nil
}