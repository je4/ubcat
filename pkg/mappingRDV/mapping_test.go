package mappingRDV

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/je4/ubcat/v2/pkg/schema"
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
			Location: &schema.Location{
				Digital: []*schema.Url{
					{
						Content: "",
						Format:  "",
						Url:     "https://www.e-manuscripta.ch/bau/content/titleinfo/3802578",
						Note:    "Link auf e-manuscripta",
					},
				},
				Electronic: nil,
				Holding: []*schema.Holding{
					{
						CallNumber:  "UBH Test",
						Item:        nil,
						Library:     "",
						Location:    "",
						Note:        "",
						Summary:     "",
						SummaryNote: "",
					},
					{
						CallNumber:  "UBH Test 2",
						Item:        nil,
						Library:     "",
						Location:    "",
						Note:        "",
						Summary:     "",
						SummaryNote: "",
					},
				},
			},
			Name: nil,
			Note: &schema.Note{
				Acquisition:                     nil,
				Action:                          nil,
				AdditionalPhysicalForm:          nil,
				Bibliography:                    nil,
				Binding:                         nil,
				Bibliographical:                 nil,
				Citation:                        nil,
				Credits:                         nil,
				Exhibitions:                     nil,
				Funding:                         nil,
				General:                         []string{"note 1", "note 2"},
				Han:                             nil,
				Language:                        nil,
				MediumOfPerformance:             nil,
				Other:                           nil,
				Ownership:                       nil,
				Performers:                      nil,
				PreferredCitation:               nil,
				Publications:                    nil,
				Reproduction:                    nil,
				Restriction:                     nil,
				StatementOfResponsibility:       nil,
				StatementOfResponsibilityAltRep: nil,
				SystemDetails:                   nil,
				Thesis:                          nil,
				Venue:                           nil,
				VersionIdentification:           nil,
			},
			OriginInfo: &schema.OriginInfo{
				CopyrightDate: nil,
				Distribution:  nil,
				Edition:       nil,
				EditionAltRep: nil,
				Geographic:    nil,
				Manufacture:   nil,
				Production: []*schema.PublicationNote{
					{
						Date:        "13. März 1958",
						LinkedField: "",
						Place:       []string{"[S.l.]", "Basel"},
						Publisher:   nil, //[]string{"Verlag 1", "Verlag 2"},
					},
				},
				Publication: nil,
			},
			PhysicalDescription: &schema.PhysicalDescription{
				Arrangement:               nil,
				DateSequentialDesignation: nil,
				Extent: []*schema.Extent{
					{
						AccompanyingMaterial: "",
						Dimensions:           "23 cm",
						Extent:               "VIII, 230 S.",
						PhysicalDetails:      "Ill.",
					},
				},
				Frequency:    nil,
				Medium:       []string{"Papier", "15,9 x 14,3 cm"},
				NotatedMusic: nil,
			},
			RecordIdentifier: []string{"(swissbib)219446946-41slsp_network",
				"(NEBIS)002014884EBI01",
				"(IDSBB)001950248DSV01",
				"991085081549705501",
				"(EXLNZ-41SLSP_NETWORK)991085081549705501",
				"9919502480105504"},
			RelatedItem:     nil,
			Subject:         nil,
			TableOfContents: nil,
			TargetAudience:  nil,
			TitleInfo: &schema.TitleInfo{
				Abbreviated: nil,
				Alternative: nil,
				Main: []*schema.Title{
					{
						Title:       "<<Der>> Brief an Ernsten Grarteri zu Solmbs Herrn zu Mintzenberg",
						SubTitle:    "Untertitel",
						PartName:    nil,
						PartNumber:  nil,
						LinkedField: "",
					},
				},
				Translated: nil,
				Uniform:    nil,
			},
		},
		Facets: &schema.Facets{

			Agents: []*schema.AgentFacets{
				{
					Name: "scribe",
					Agent: []*schema.Agent{
						{
							Identifer: []string{"(DE-588)118560093"},
							Label:     "Karl V., Heiliges Römisches Reich, Kaiser (1500-1558)",
							Role:      []string{"scr"},
						},
						{
							Identifer: []string{"(DE-588)118718444"},
							Label:     "Granvelle, Antoine Perrenot de (1517-1586)",
							Role:      []string{"scr", "aut"},
						},
						{
							Identifer: []string{"(noid)Seld, V."},
							Label:     "Seld, V.",
							Role:      []string{"scr", "aut"},
						},
					},
				},
				{
					Name: "author",
					Agent: []*schema.Agent{
						{
							Identifer: []string{"(DE-588)118718444"},
							Label:     "Granvelle, Antoine Perrenot de (1517-1586)",
							Role:      []string{"scr", "aut"},
						},
						{
							Identifer: []string{"(noid)Seld, V."},
							Label:     "Seld, V.",
							Role:      []string{"scr", "aut"},
						},
					},
				},
			},
			Concepts:  nil,
			Daterange: nil,
			Strings:   nil,
		},

		Sets:           nil,
		Flags:          []string{"all", "portraets"},
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
	if titles, ok := result["titleInfoMainTitle"]; !ok {
		t.Errorf("mainTitle not found")
	} else {
		if len(titles) != 1 {
			t.Errorf("expected 1 titles, got %d", len(titles))
		}
	}
	if originInfos, ok := result["originInfoProduction"]; !ok {
		t.Errorf("originInfo not found")
	} else {
		if len(originInfos) != 1 {
			t.Errorf("expected 1 originInfo, got %d", len(originInfos))
		}
	}
	if facetAutographScribe, ok := result["facetAutographScribe"]; !ok {
		t.Errorf("facetAutographScribe not found")
	} else {
		if len(facetAutographScribe) != 3 {
			t.Errorf("expected 3 facetAutographScribe, got %d", len(facetAutographScribe))
		}
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Errorf("cannot marshal result: %v", err)
	}
	t.Logf("result: \n%s", data)
}
