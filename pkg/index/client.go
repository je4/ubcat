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

type Result struct {
	Docs  map[string]*schema.UBSchema
	Total int64
	From  int64
	Num   int64
}

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

func (c *Client) Search(ctx context.Context, querystring string, filter map[string]string, vectorMarc, vectorJson, vectorProse []float32, from, num int64) (*Result, error) {
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

	wildcardQuery := map[string]types.WildcardQuery{}
	for k, v := range filter {
		val := v + "*"
		wildcardQuery[k] = types.WildcardQuery{Value: &val}
	}
	if len(wildcardQuery) > 0 {
		esMust = append(esMust, types.Query{
			Wildcard: wildcardQuery,
		})
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
		From(int(from)).
		Size(int(num))

	elasticResponse, err := elasticQuery.Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot search")
	}

	var docs = map[string]*schema.UBSchema{}
	for _, hit := range elasticResponse.Hits.Hits {
		var s = &schema.UBSchema{}
		if err := json.Unmarshal(hit.Source_, s); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal document %v", hit.Source_)
		}
		s.Score_ = float64(hit.Score_)
		s.Id_ = hit.Id_
		docs[hit.Id_] = s
	}
	result := &Result{
		Docs: docs,
		From: from,
		Num:  num,
	}
	if elasticResponse.Hits.Total != nil {
		result.Total = elasticResponse.Hits.Total.Value
	}
	return result, nil
}
func (c *Client) SearchKNN(ctx context.Context, filter map[string]string, vector []float32, vectorField string, k int64, numCandidates int64) (*Result, error) {
	esMust := []types.Query{}
	wildcardQuery := map[string]types.WildcardQuery{}
	for k, v := range filter {
		val := v + "*"
		wildcardQuery[k] = types.WildcardQuery{Value: &val}
	}
	if len(wildcardQuery) > 0 {
		esMust = append(esMust, types.Query{
			Wildcard: wildcardQuery,
		})
	}
	knnQuery := types.KnnQuery{
		Field:         vectorField,
		QueryVector:   vector,
		K:             k,
		NumCandidates: numCandidates,
	}
	if len(esMust) > 0 {
		knnQuery.Filter = esMust
	}

	searchRequest := &search.Request{}
	searchRequest.Knn = []types.KnnQuery{
		knnQuery,
	}
	elasticQuery := c.client.Search().
		Index(c.index).
		// SourceExcludes_("embedding_marc", "embedding_json", "embedding_prose").
		Request(searchRequest)

	elasticResponse, err := elasticQuery.Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cannot search")
	}

	var docs = map[string]*schema.UBSchema{}
	for _, hit := range elasticResponse.Hits.Hits {
		var s = &schema.UBSchema{}
		if err := json.Unmarshal(hit.Source_, s); err != nil {
			return nil, errors.Wrapf(err, "cannot unmarshal document %v", hit.Source_)
		}
		s.Score_ = float64(hit.Score_)
		s.Id_ = hit.Id_
		docs[hit.Id_] = s
	}
	result := &Result{
		Docs: docs,
		From: 0,
		Num:  k,
	}
	if elasticResponse.Hits.Total != nil {
		result.Total = elasticResponse.Hits.Total.Value
	}
	return result, nil
}
