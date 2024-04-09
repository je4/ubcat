package index

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/je4/ubcat/v2/pkg/schema"
)

func NewClient(index string, client *elasticsearch.TypedClient) *Client {
	return &Client{
		index:  index,
		client: client,
	}
}

type Client struct {
	index  string
	client *elasticsearch.TypedClient
}

func (c *Client) Index() string {
	return c.index
}

func (c *Client) GetDocuments(ctx context.Context, identifiers ...string) (map[string]*schema.UBSchema, error) {
	mgetResponse, err := c.client.Mget().Index(c.index).Ids(identifiers...).Do(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get documents %v from index %s", c.index, identifiers)
	}

	var result = map[string]*schema.UBSchema{}
	for _, docInt := range mgetResponse.Docs {
		doc, ok := docInt.(*types.GetResult)
		if !ok {
			return nil, errors.Errorf("cannot convert document %v to GetResult", docInt)
		}
		var s = &schema.UBSchema{}
		if err := json.Unmarshal(doc.Source_, s); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal document %v", doc.Source_)
		}
		result[doc.Id_] = s
	}
	return result, nil
}

func vectorQuery(vector []float32, field string) (*types.Query, error) {
	vectorBytes, err := json.Marshal(vector)
	if err != nil {
		return nil, errors.Wrap(err, "cannot marshal params")
	}
	return &types.Query{
		ScriptScore: &types.ScriptScoreQuery{
			Query: &types.Query{
				Exists: &types.ExistsQuery{
					Field: field,
				},
			},
			Script: &types.InlineScript{
				Source: fmt.Sprintf("cosineSimilarity(params.queryVector, '%s') + 1.0", field),
				Params: map[string]json.RawMessage{
					"queryVector": vectorBytes,
				},
			},
		},
	}, nil
}

func (c *Client) Search(ctx context.Context, querystring string, vectorMarc, vectorJson, vectorProse []float32, from, num int) (map[string]*schema.UBSchema, error) {
	esMust := []types.Query{}
	if len(vectorMarc) > 0 {
		vQuery, err := vectorQuery(vectorMarc, "embedding_marc")
		if err != nil {
			return nil, errors.Wrap(err, "cannot create vector query")
		}
		esMust = append(esMust, *vQuery)
	}
	if len(vectorJson) > 0 {
		vQuery, err := vectorQuery(vectorJson, "embedding_json")
		if err != nil {
			return nil, errors.Wrap(err, "cannot create vector query")
		}
		esMust = append(esMust, *vQuery)
	}
	if len(vectorProse) > 0 {
		vQuery, err := vectorQuery(vectorProse, "embedding_prose")
		if err != nil {
			return nil, errors.Wrap(err, "cannot create vector query")
		}
		esMust = append(esMust, *vQuery)
	}
	if querystring != "" && len(esMust) == 0 {
		esMust = append(esMust, types.Query{
			SimpleQueryString: &types.SimpleQueryStringQuery{
				Query: querystring,
			},
		})
	}
	searchRequest := &search.Request{}
	searchRequest.Query = &types.Query{
		Bool: &types.BoolQuery{
			Filter: []types.Query{},
			Must:   esMust,
		},
	}

	elasticQuery := c.client.Search().
		Index(c.index).
		// SourceExcludes_("embedding_marc", "embedding_json", "embedding_prose").
		Request(searchRequest).
		From(from).
		Size(num)

	elasticResponse, err := elasticQuery.Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot search")
	}

	var result = map[string]*schema.UBSchema{}
	for _, hit := range elasticResponse.Hits.Hits {
		var s = &schema.UBSchema{}
		if err := json.Unmarshal(hit.Source_, s); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal document %v", hit.Source_)
		}
		result[hit.Id_] = s
	}
	return result, nil
}
