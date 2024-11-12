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
			Files: []*schema.Files{
				{
					Media: &schema.Media{
						Medias: []*schema.Medias{
							{
								Height:       0,
								License_abbr: "",
								Mimetype:     "",
								Path:         "",
								Pronom:       "",
								Type:         "metadata",
								Uri:          "",
								Width:        0,
							},
							{
								Height:       0,
								License_abbr: "",
								Mimetype:     "",
								Path:         "",
								Pronom:       "",
								Type:         "image",
								Uri:          "",
								Width:        0,
							},
							{
								Height:       0,
								License_abbr: "",
								Mimetype:     "",
								Path:         "",
								Pronom:       "",
								Type:         "video",
								Uri:          "",
								Width:        0,
							},
							{
								Height:       0,
								License_abbr: "",
								Mimetype:     "",
								Path:         "",
								Pronom:       "",
								Type:         "video",
								Uri:          "",
								Width:        0,
							},
						},
						Poster: &schema.Medias{
							Height:       5476,
							License_abbr: "PDM 1.0 Deed",
							Mimetype:     "image/tiff",
							Path:         "vfs://ub-reprofiler/mets-container/bau_1/2011b/BAU_1_002632345_20120919T230052_master_ver1.zip/002632345/image/3955426.tif",
							Pronom:       "fmt/353",
							Type:         "image",
							Uri:          "mediaserver:bau_1/002632345_003_1_image_3955426.tif",
							Width:        8945,
						},
					},
					Structure: nil,
				},
			},
			Fulltext:   nil,
			Identifier: nil,
			Language:   nil,
			Location: &schema.Location{
				Digital: []*schema.Url{
					{
						Content: "",
						Format:  "",
						Url:     "https://www.e-manuscripta.ch/bau/content/titleinfo/3802578",
						Note:    "Link auf e-manuscripta",
					},
					{
						Content: "Porträt Vorschau",
						Format:  "bild",
						Url:     "http://www.ub.unibas.ch/digi/a100/portraet/bs/m/IBB_1_004261027.jpg",
					},
				},
				Electronic: []*schema.Electronic{
					{
						Library:       "AFREE",
						Collection:    "Project MUSE Open Access Books",
						Url:           "https://eu03.alma.exlibrisgroup.com/view/uresolver/41SLSP_UBS/openurl?u.ignore_date_coverage=true&amp;portfolio_pid=53375898680005504&amp;Force_direct=true",
						AlmaAccessUrl: "https://slsp-ubs.primo.exlibrisgroup.com/discovery/openurl?institution=41SLSP_UBS&amp;vid=41SLSP_UBS:live&amp;?u.ignore_date_coverage=true&amp;rft.mms_id=9972497358305504",
						Interface:     "Project Muse",
						Coverage:      "",
						Note:          "",
					},
					{
						Library:       "A145",
						Collection:    "ProQuest One Literature",
						Url:           "https://eu03.alma.exlibrisgroup.com/view/uresolver/41SLSP_UBS/openurl?u.ignore_date_coverage=true&amp;portfolio_pid=53409066940005504&amp;Force_direct=true",
						AlmaAccessUrl: "",
						Interface:     "",
						Coverage:      "Available from 01/01/1616 until 31/01/1640.",
						Note:          "",
					},
				},
				Holding: []*schema.Holding{
					{
						CallNumber:  "UBH Test",
						Item:        nil,
						Library:     "A100",
						Location:    "",
						Note:        "",
						Summary:     "",
						SummaryNote: "",
					},
					{
						CallNumber:  "UBH Test 2",
						Item:        nil,
						Library:     "A125",
						Location:    "",
						Note:        "",
						Summary:     "",
						SummaryNote: "",
					},
				},
			},
			Name: nil,
			Note: &schema.Note{
				Acquisition:            nil,
				Action:                 nil,
				AdditionalPhysicalForm: nil,
				Bibliography:           nil,
				Binding:                nil,
				Bibliographical:        nil,
				Citation: []*schema.NoteWithURL{
					{
						Add:     "1407625X",
						License: "",
						Main:    "VD18",
						Url:     nil,
					},
				},
				Credits:             nil,
				Exhibitions:         nil,
				Funding:             nil,
				General:             []string{"note 1", "note 2"},
				Han:                 nil,
				Language:            nil,
				MediumOfPerformance: nil,
				Other:               nil,
				Ownership: []*schema.NoteWithURL{
					{
						Add:     "",
						License: "",
						Main:    "Gehörte früher mir",
						Url:     []string{"http://test.ch"},
					},
				},
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
				"991085081549705507",
				"(EXLNZ-41SLSP_NETWORK)991085081549705501",
				"9919502480105504", "991170524020205501"},
			RelatedItem: nil,
			Subject:     nil,
			TableOfContents: []string{"Theil 1",
				"Geschichte der deutschen Malerei bis 1450",
				"Theil 2",
				"<<Die>> flandrische Malerei des fünfzehnten Jahrhunderts, Liefg. 1"},
			TargetAudience: nil,
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
							Identifier: []string{"(DE-588)118560093"},
							Label:      "Karl V., Heiliges Römisches Reich, Kaiser (1500-1558)",
							Role:       []string{"scr"},
						},
						{
							Identifier: []string{"(DE-588)118718444"},
							Label:      "Granvelle, Antoine Perrenot de (1517-1586)",
							Role:       []string{"scr", "aut"},
						},
						{
							Identifier: []string{"(noid)Seld, V."},
							Label:      "Seld, V.",
							Role:       []string{"scr", "aut"},
						},
					},
				},
				{
					Name: "author",
					Agent: []*schema.Agent{
						{
							Identifier: []string{"(DE-588)118718444"},
							Label:      "Granvelle, Antoine Perrenot de (1517-1586)",
							Role:       []string{"scr", "aut"},
						},
						{
							Identifier: []string{"(noid)Seld, V."},
							Label:      "Seld, V.",
							Role:       []string{"scr", "aut"},
						},
					},
				},
			},
			Concepts: []*schema.ConceptFacets{
				{
					Name: "publicationPlace",
					Concept: []*schema.Concept{
						{
							Identifier: []string{"(DE-588)118718444"},
							Label:      "Basel",
						},
					},
				},
			},
			Daterange: nil,
			Strings: []*schema.StringFacets{
				{
					Name:   "resourceType",
					String: []string{"book", "archivalMaterial"},
				},
				{
					Name:   "geigyNummer",
					String: []string{"100"},
				},
			},
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
