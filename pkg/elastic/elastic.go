package elastic

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/je4/ubcat/v2/pkg/schema"
	"github.com/je4/utils/v2/pkg/zLogger"
	"slices"
)

func NewElastic(elastic *elasticsearch.TypedClient, logger zLogger.ZLogger, index string) *Elastic {
	return &Elastic{
		elastic: elastic,
		logger:  logger,
		index:   index,
	}
}

type Elastic struct {
	elastic *elasticsearch.TypedClient
	logger  zLogger.ZLogger
	index   string
}

func (e *Elastic) SimpleSearch(query string, embedding []float32, embeddingName string) ([]*schema.UBSchema, error) {
	res := []*schema.UBSchema{}
	esMust := []types.Query{}
	if query != "" {
		esMust = append(esMust, types.Query{
			SimpleQueryString: &types.SimpleQueryStringQuery{
				Query: query,
			},
		})
	} else {
		esMust = append(esMust, types.Query{
			MatchAll: &types.MatchAllQuery{},
		})
	}
	if len(embedding) > 0 {
		if !slices.Contains([]string{"marc", "prose", "json"}, embeddingName) {
			return nil, errors.Errorf("unknown embedding name '%s'", embeddingName)
		}
		vectorBytes, err := json.Marshal(embedding)
		if err != nil {
			return nil, errors.Wrap(err, "cannot marshal params")
		}
		field := "embedding_" + embeddingName
		esMust = append(esMust, types.Query{
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
		})
	}
	searchRequest := &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Filter: []types.Query{},
				Must:   esMust,
			},
		},
	}
	resp, err := e.elastic.Search().Index(e.index).SourceExcludes_("embedding_marc", "embedding_prose", "embedding_json").Request(searchRequest).Do(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "cannot search")
	}
	for _, hit := range resp.Hits.Hits {
		entry := &schema.UBSchema{}
		if err := json.Unmarshal(hit.Source_, entry); err != nil {
			return nil, errors.Wrap(err, "cannot unmarshal json")
		}
		res = append(res, entry)
	}
	return res, nil
}
