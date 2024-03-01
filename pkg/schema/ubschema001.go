package schema

import (
	"encoding/json"
	"time"
)

type DataField struct {
	Tag      string   `json:"tag"`
	Ind1     string   `json:"ind1"`
	Ind2     string   `json:"ind2"`
	Subfield []string `json:"subfield"`
}

type Classification struct {
	All     []string `json:"all"`
	General []string `json:"general"`
}

type HoldingLocation struct {
	Library    string `json:"library,omitempty"`
	Location   string `json:"location,omitempty"`
	CallNumber string `json:"callNumber,omitempty"`
}

type ItemLocation struct {
	Library           string `json:"library,omitempty"`
	Location          string `json:"location,omitempty"`
	HoldingCallNumber string `json:"holdingCallNumber,omitempty"`
}

type Location struct {
	Holding []HoldingLocation `json:"holding,omitempty"`
	Item    []ItemLocation    `json:"item,omitempty"`
}

type SubjectName struct {
	Personal []json.RawMessage `json:"personal,omitempty"`
}

type Subject struct {
	Name SubjectName `json:"name,omitempty"`
}

type OriginInfo struct {
	Place     []string `json:"place,omitempty"`
	Publisher []string `json:"publisher,omitempty"`
	Date      string   `json:"date,omitempty"`
}

type AccessCondition struct {
	Add     []string `json:"add,omitempty"`
	Licence []string `json:"licence,omitempty"`
	Main    []string `json:"main,omitempty"`
	Url     []string `json:"url,omitempty"`
}

type Cartographics struct {
	Coded   []Coded   `json:"coded,omitempty"`
	Display []Display `json:"display,omitempty"`
}

type Coded struct {
	Coordinates []string `json:"coordinates,omitempty"`
	Scale       []string `json:"scale,omitempty"`
}

type Display struct {
	Coordinates []string `json:"coordinates,omitempty"`
	Other       []string `json:"other,omitempty"`
	Projection  []string `json:"projection,omitempty"`
	Scale       []string `json:"scale,omitempty"`
}

type Mapping001 struct {
	Abstract        []string                  `json:"abstract,omitempty"`
	AccessCondition []AccessCondition         `json:"accessCondition,omitempty"`
	Cartographics   Cartographics             `json:"cartographics,omitempty"`
	Classification  map[string]Classification `json:"classification,omitempty"`
	Date            []DateRange               `json:"date,omitempty"`
	//	DateRange           []DateRange                `json:"daterange,omitempty"`
	Extension           map[string]json.RawMessage `json:"extension,omitempty"`
	Files               map[string]json.RawMessage `json:"files,omitempty"`
	Fulltext            []string                   `json:"fulltext,omitempty"`
	Identifier          map[string]json.RawMessage `json:"identifier,omitempty"`
	Language            []string                   `json:"language,omitempty"`
	Location            Location                   `json:"location,omitempty"`
	Name                map[string]json.RawMessage `json:"name,omitempty"`
	Note                map[string]json.RawMessage `json:"note,omitempty"`
	OriginInfo          OriginInfo                 `json:"originInfo,omitempty"`
	PhysicalDescription map[string]json.RawMessage `json:"physicalDescription,omitempty"`
	RecordIdentifier    []string                   `json:"recordIdentifier,omitempty"`
	RelatedItem         map[string]json.RawMessage `json:"relatedItem,omitempty"`
	Subject             Subject                    `json:"subject,omitempty"`
	TableOfContents     map[string]json.RawMessage `json:"tableOfContents,omitempty"`
	TargetAudience      map[string]json.RawMessage `json:"targetAudience,omitempty"`
	TitleInfo           map[string]json.RawMessage `json:"titleInfo,omitempty"`
}

type Facets struct {
	Name   string   `json:"name"`
	String []string `json:"string"`
}

type ACL struct {
	Content []string `json:"content,omitempty"`
	Meta    []string `json:"meta,omitempty"`
	Preview []string `json:"preview,omitempty"`
}

type UBSchema001 struct {
	Timestamp    time.Time         `json:"timestamp"`
	OaiId        []string          `json:"oai_id"`
	Ldr          map[string]string `json:"LDR,omitempty"`
	Controlfield map[string]string `json:"controlfield,omitempty"`
	// Datafield    []DataField         `json:"datafield,omitempty"`
	// FieldLists map[string][]string `json:"fieldlists,omitempty"`
	Mapping Mapping001 `json:"mapping,omitempty"`
	Facets  []Facets   `json:"facets,omitempty"`
	Sets    []string   `json:"sets,omitempty"`
	Flags   []string   `json:"flags,omitempty"`
	ACL     ACL        `json:"acl,omitempty"`
}
