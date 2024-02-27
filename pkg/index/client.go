package index

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
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
