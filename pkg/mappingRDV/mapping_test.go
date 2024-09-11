package mappingRDV

import (
	"encoding/json"
	"github.com/je4/ubcat/v2/pkg/schema"
	"testing"
	"time"
)

func TestMappingRDV(t *testing.T) {
	content := &schema.UBSchema001{
		Timestamp:    time.Time{},
		OaiId:        nil,
		Ldr:          nil,
		Controlfield: nil,
		Mapping: &schema.Mapping001{
			Abstract:        nil,
			AccessCondition: nil,
			Cartographics:   nil,
			Classification:  nil,
			Date:            nil,
			Extension:       nil,
			Files:           nil,
			Fulltext:        nil,
			Identifier:      nil,
			Language:        nil,
			Location:        nil,
			Name:            nil,
			Note:            nil,
			OriginInfo: &schema.OriginInfo{
				CopyrightDate: nil,
				Distribution:  nil,
				Edition:       nil,
				EditionAltRep: nil,
				Geographic:    nil,
				Manufacture:   nil,
				Production: []*schema.PublicationNote{
					{
						Date:        "",
						LinkedField: "",
						Place:       []string{"Place 1", "Place 2"},
						Publisher:   []string{"Publisher 1", "Publisher 2"},
					},
				},
				Publication: nil,
			},
			PhysicalDescription: nil,
			RecordIdentifier:    nil,
			RelatedItem:         nil,
			Subject:             nil,
			TableOfContents:     nil,
			TargetAudience:      nil,
			TitleInfo: &schema.TitleInfo{
				Abbreviated: nil,
				Alternative: nil,
				Main: []*schema.Title{
					{
						Title:       "Title 1",
						SubTitle:    "",
						PartName:    nil,
						PartNumber:  nil,
						LinkedField: "",
					},
					{
						Title:       "Title 2",
						SubTitle:    "",
						PartName:    nil,
						PartNumber:  nil,
						LinkedField: "",
					},
				},
				Translated: nil,
				Uniform:    nil,
			},
		},
		Facets: []schema.Facets{
			{
				Agents: []schema.AgentFacets{
					{
						Name: "scribe",
						Agent: []schema.Agent{
							{
								Identifer: []string{"4711"},
								Label:     "AgentLabel 1",
								Role:      nil,
							},
						},
					},
				},
				Concepts:  nil,
				Daterange: nil,
				Strings:   nil,
			},
		},
		Sets:           nil,
		Flags:          nil,
		ACL:            schema.ACL{},
		EmbeddingProse: nil,
		EmbeddingMarc:  nil,
		EmbeddingJson:  nil,
	}

	result := (*MappingRDV)(content).Map()

	//	t.Logf("result: %v", result)
	/*
		if fmt.Sprintf("%v", result) != "map[mainTitle:[{Title 1  } {Title 2  }] originInfo:[{Place 1, Place 2  }]]" {
			t.Errorf("unexpected result: %v", result)
		}
	*/
	if titles, ok := result["mainTitle"]; !ok {
		t.Errorf("mainTitle not found")
	} else {
		if len(titles) != 2 {
			t.Errorf("expected 2 titles, got %d", len(titles))
		}
	}
	if originInfos, ok := result["originInfo"]; !ok {
		t.Errorf("originInfo not found")
	} else {
		if len(originInfos) != 1 {
			t.Errorf("expected 1 originInfo, got %d", len(originInfos))
		}
	}
	if facetAutographScribe, ok := result["facetAutographScribe"]; !ok {
		t.Errorf("facetAutographScribe not found")
	} else {
		if len(facetAutographScribe) != 1 {
			t.Errorf("expected 1 facetAutographScribe, got %d", len(facetAutographScribe))
		}
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Errorf("cannot marshal result: %v", err)
	}
	t.Logf("result: \n%s", data)
}
