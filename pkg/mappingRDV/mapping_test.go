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
			/*Files: []*schema.Files{
				{
					Media: &schema.Media{
						Medias: []*schema.Medias{
							{
								Height:       0,
								License_abbr: "InC",
								Mimetype:     "image/tiff",
								Path:         "",
								Pronom:       "fmt/153",
								Type:         "metadata",
								Uri:          "",
								Width:        0,
							},
							{
								Height:       1520,
								License_abbr: "InC",
								Mimetype:     "image/tiff",
								Path:         "",
								Pronom:       "fmt/153",
								Type:         "image",
								Uri:          "",
								Width:        3480,
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
					Structure: &schema.Structure{
						DigitalObject: &schema.DigitalObject{
							DigitalContainer: nil,
							FromRecordHeader: nil,
							Header:           nil,
							Id:               "md345678",
							RecordIdentifier: nil,
							Label:            "",
							Type:             "",
						},
					},
				},
			},*/
			/*Files: []*schema.Files{
				{
					Media: &schema.Media{
						Medias: []*schema.Medias{
							{
								License_abbr: "PDM 1.0 Deed",
								Mimetype:     "application/xml",
								Path:         "vfs://ub-reprofiler/mets-container/bau_1/2018/BAU_1_002632345_20190421T145549_gen3_ver1.zip/002632345/export_mets.xml",
								Pronom:       "UNKNOWN",
								Type:         "metadata",
								Uri:          "mediaserver:bau_1/002632345_003_1_export_mets.xml",
							},
							{
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
					Structure: &schema.Structure{
						DigitalObject: &schema.DigitalObject{
							DigitalContainer: nil,
							FromRecordHeader: nil,
							Header:           nil,
							Id:               "md3955424",
							RecordIdentifier: nil,
							Label:            "",
							Type:             "document",
						},
					},
				},
			},*/
			Files:      nil,
			Fulltext:   nil,
			Identifier: nil,
			Language:   []string{"ger", "eng", "ger"},
			Location: &schema.Location{
				Digital: []*schema.Url{
					{
						Content: "",
						Format:  "",
						Url:     "https://www.e-manuscripta.ch/bau/content/titleinfo/3802578",
						Note:    "Link auf e-manuscripta",
					},
					/*{
						Content: "Porträt Vorschau",
						Format:  "bild",
						Url:     "http://www.ub.unibas.ch/digi/a100/portraet/bs/m/IBB_1_004261027.jpg",
					},*/
					{
						Content: "Porträt Vorschau",
						Format:  "bild",
						Url:     "http://www.ub.unibas.ch/digi/a100/portraet/2011-04/BAU_1_005322299.jpg",
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
			Name: &schema.Name{
				Conference: nil,
				Corporate:  nil,
				Family:     nil,
				Personal: []map[string]schema.Personal{
					{
						"gnd": schema.Personal{
							Date: &schema.PersonalDate{
								Original: "1964-",
								Birth:    "",
								Death:    "",
							},
							Gender:          "male",
							Identifier:      "(DE-588)138393389",
							NamePart:        "Brown, Andrew",
							OtherIdentifier: []string{"(orcid)0000-0001-8463-7145"},
							PlaceOfActivity: nil,
							PlaceOfBirth:    nil,
							Profession:      nil,
							Related:         nil,
							Role:            []string{"aut"},
							Variant:         nil,
						},
						"idref": schema.Personal{
							Date: &schema.PersonalDate{
								Original: "1964-....",
								Birth:    "",
								Death:    "",
							},
							Gender:          "",
							Identifier:      "(IDREF)035764236",
							NamePart:        "Brown, Andrew D.",
							OtherIdentifier: nil,
							PlaceOfActivity: nil,
							PlaceOfBirth:    nil,
							Profession:      nil,
							Related:         nil,
							Role:            []string{"cre"},
							Variant:         nil,
						},
						"unknown": schema.Personal{
							Date:            nil,
							Gender:          "",
							Identifier:      "",
							NamePart:        "Brown, A. D.",
							OtherIdentifier: nil,
							PlaceOfActivity: nil,
							PlaceOfBirth:    nil,
							Profession:      nil,
							Related:         nil,
							Role:            nil,
							Variant:         nil,
						},
					},
					{
						"gnd": schema.Personal{
							Date: &schema.PersonalDate{
								Original: "1964-",
								Birth:    "",
								Death:    "",
							},
							Gender:          "male",
							Identifier:      "(DE-588)13839338977",
							NamePart:        "Green, Anna",
							OtherIdentifier: []string{"(orcid)0000-0001-8463-7145"},
							PlaceOfActivity: nil,
							PlaceOfBirth:    nil,
							Profession:      nil,
							Related:         nil,
							Role:            []string{"cmp"},
							Variant:         nil,
						},
						"idref": schema.Personal{
							Date: &schema.PersonalDate{
								Original: "1964-....",
								Birth:    "",
								Death:    "",
							},
							Gender:          "",
							Identifier:      "(IDREF)03576423699",
							NamePart:        "Green, Anna D.",
							OtherIdentifier: nil,
							PlaceOfActivity: nil,
							PlaceOfBirth:    nil,
							Profession:      nil,
							Related:         nil,
							Role:            []string{"cre"},
							Variant:         nil,
						},
						"unknown": schema.Personal{
							Date:            nil,
							Gender:          "",
							Identifier:      "",
							NamePart:        "Green, A. D.",
							OtherIdentifier: nil,
							PlaceOfActivity: nil,
							PlaceOfBirth:    nil,
							Profession:      nil,
							Related:         nil,
							Role:            nil,
							Variant:         nil,
						},
					},
				},
			},
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
				Language:            []string{"Epos in ursprünglicher italienischer Volkssprache, Einleitung und Kommentar auf Englisch"},
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
			/*RecordIdentifier: []string{"(swissbib)219446946-41slsp_network",
			"(NEBIS)002014884EBI01",
			"(IDSBB)001950248DSV01",
			"991085081549705507",
			"(EXLNZ-41SLSP_NETWORK)991085081549705501",
			"9919502480105504", "991170524020205501"},*/
			RecordIdentifier: []string{"9920530890105504"},
			RelatedItem: &schema.RelatedItem{
				Constituent: []*schema.Related{
					{
						DisplayConstant:           "",
						Identifier:                nil,
						InternalIdentifier:        "(swissbib)176683909-41slsp_network",
						VolumeDesignation:         "",
						StatementOfResponsibility: "",
						LinkedField:               "",
						Title:                     "Fortpflanzung, Zuchthygiene und Haustierbesamung",
					},
					{
						DisplayConstant:           "",
						Identifier:                nil,
						InternalIdentifier:        "(IDSBB)176683907DSV01",
						VolumeDesignation:         "",
						StatementOfResponsibility: "",
						LinkedField:               "",
						Title:                     "Fortpflanzung, Zuchthygiene und Haustierbesamung",
					},
					{
						DisplayConstant:           "",
						Identifier:                nil,
						InternalIdentifier:        "(swissbib)222331992-41slsp_network",
						VolumeDesignation:         "",
						StatementOfResponsibility: "",
						LinkedField:               "",
						Title:                     "Mitteilungen der Tierärztlichen Gesellschaft zur Bekämpfung des Kurpfuschertums",
					},
					{
						DisplayConstant:           "nur 1980",
						Identifier:                nil,
						InternalIdentifier:        "",
						VolumeDesignation:         "",
						StatementOfResponsibility: "",
						LinkedField:               "",
						Title:                     "Jährliche Kurpfuscherliste",
					},
				},
				Host: []*schema.Host{
					{
						Title:              "Les cahiers Ciba",
						Publisher:          "",
						PartYear:           "",
						PartNumber:         "",
						Part:               "Vol. 4, No. 38 (1951), p. 1282-1307",
						PartSort:           "4/38/1951/1282-1307",
						Identifier:         nil,
						InternalIdentifier: "(IDSBB)003787128DSV01",
					},
				},
				IssuedWith:   nil,
				Location:     nil,
				Original:     nil,
				OriginalNote: nil,
				Other:        nil,
				OtherFormat:  nil,
				OtherVersion: nil,
				Preceding:    nil,
				Series: []*schema.Series{
					{
						Conference:         nil,
						Corporate:          nil,
						InternalIdentifier: "",
						LinkedField:        "",
						PartName:           nil,
						PartNumber:         nil,
						Personal:           nil,
						Title:              "Berner Schriften zur Kunst",
						VolumeDesignation:  "Bd. 8",
					},
					{
						Conference:         nil,
						Corporate:          nil,
						InternalIdentifier: "(NEBIS)000043484EBI01",
						LinkedField:        "",
						PartName:           nil,
						PartNumber:         nil,
						Personal:           nil,
						Title:              "Berner Schriften zur Kunst",
						VolumeDesignation:  "8",
					},
					{
						Conference:         nil,
						Corporate:          nil,
						InternalIdentifier: "(RERO)0243741-41slsp",
						LinkedField:        "",
						PartName:           nil,
						PartNumber:         nil,
						Personal:           nil,
						Title:              "Berner Schriften zur Kunst",
						VolumeDesignation:  "8",
					},
				},
				Succeeding: nil,
				Work:       nil,
			},
			Subject: &schema.Subject{
				Genre: map[string][]string{
					"idsbb": {"Test"},
					"other": {"TestR", "TestR2"},
				},
				Geographic: nil,
				Local:      nil,
				Music:      []string{"Ouvertüre", "19.-20. Jh."},
				Name: &schema.Name{
					Conference: nil,
					Corporate:  nil,
					Family: []map[string]schema.Family{
						{
							"gnd": schema.Family{
								Date:            "",
								Identifier:      "",
								NamePart:        "Gebrüder Grimm",
								OtherIdentifier: []string{"(orcid)0000-0001-8463-7145"},
								Role:            []string{"aut"},
								Variant:         nil,
							},
						},
					},
					Personal: []map[string]schema.Personal{
						{
							"gnd": schema.Personal{
								Date: &schema.PersonalDate{
									Original: "1964-",
									Birth:    "",
									Death:    "",
								},
								Gender:          "male",
								Identifier:      "(DE-588)138393389",
								NamePart:        "Brown, Andrew",
								OtherIdentifier: []string{"(orcid)0000-0001-8463-7145"},
								PlaceOfActivity: nil,
								PlaceOfBirth:    nil,
								Profession:      nil,
								Related:         nil,
								Role:            []string{"aut"},
								Variant:         nil,
							},
						},
					},
				},
				Temporal: nil,
				TitleInfo: []map[string]schema.Work{
					{
						"gnd": schema.Work{
							Identifier:      "(DE-588)1067523391",
							Name:            "Schoeck, Othmar (1886-1957)",
							Title:           "Gaselen ; arrangiert",
							OtherIdentifier: nil,
							Variant:         nil,
						},
					},
				},
				Topic: []map[string]schema.Topic{
					{
						"gnd": schema.Topic{
							Identifier:      "(DE-588)123456",
							Label:           "Brot",
							Mapped:          nil,
							OtherIdentifier: nil,
							Variant:         nil,
						},
					},
				},
				Undefined: nil,
			},
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
				{
					Name:   "digitized",
					String: []string{"yes"},
				},
			},
		},

		Sets:  nil,
		Flags: []string{"all", "portraets", "e-rara"},
		ACL: &schema.ACL{
			Content: []string{"global/guest"},
			Meta:    []string{"global/guest"},
			Preview: []string{"global/guest"},
		},
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
