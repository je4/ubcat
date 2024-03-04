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

type Subject struct {
	Genre      map[string][]string     `json:"genre,omitempty"`
	Geographic []map[string]Geographic `json:"geographic,omitempty"`
	Local      map[string][]string     `json:"local,omitempty"`
	Music      []string                `json:"music,omitempty"`
	Name       Name                    `json:"name,omitempty"`
	Temporal   map[string][]string     `json:"temporal,omitempty"`
	Topic      []map[string]Topic      `json:"topic,omitempty"`
	Undefined  []string                `json:"undefined,omitempty"`
}

type OriginInfo struct {
	CopyrightDate []string                `json:"copyrightDate,omitempty"`
	Distribution  []PublicationNote       `json:"distribution,omitempty"`
	Edition       []string                `json:"edition,omitempty"`
	EditionAltRep []string                `json:"editionAltRep,omitempty"`
	Geographic    []map[string]Geographic `json:"geographic,omitempty"`
	Manufacture   []PublicationNote       `json:"manufacture,omitempty"`
	Production    []PublicationNote       `json:"production,omitempty"`
	Publication   []PublicationNote       `json:"publication,omitempty"`
}

type PublicationNote struct {
	Date        string   `json:"date,omitempty"`
	LinkedField string   `json:"linkedField,omitempty"`
	Place       []string `json:"place,omitempty"`
	Publisher   []string `json:"publisher,omitempty"`
}

type NoteWithURL struct {
	Add     string `json:"add,omitempty"`
	Licence string `json:"licence,omitempty"`
	Main    string `json:"main,omitempty"`
	Url     string `json:"url,omitempty"`
	//Url     []string `json:"url,omitempty"`
}

type Cartographics struct {
	Coded   []Coded   `json:"coded,omitempty"`
	Display []Display `json:"display,omitempty"`
}

type Coded struct {
	Coordinates []string `json:"coordinates,omitempty"`
	Scale       string   `json:"scale,omitempty"`
}

type Display struct {
	Coordinates string `json:"coordinates,omitempty"`
	Other       string `json:"other,omitempty"`
	Projection  string `json:"projection,omitempty"`
	Scale       string `json:"scale,omitempty"`
}

type Extension struct {
	LocalCodeIZ        []string `json:"localCodeIZ,omitempty"`
	LocalCodeNZ        []string `json:"localCodeNZ,omitempty"`
	SubjectCategoryUbs []string `json:"subjectCategoryUbs,omitempty"`
}

type Identifier struct {
	Doi                      []string   `json:"doi,omitempty"`
	Isbn                     []IsbnIsmn `json:"isbn,omitempty"`
	IsbnInvalid              []string   `json:"isbnInvalid,omitempty"`
	Ismn                     []IsbnIsmn `json:"ismn,omitempty"`
	Issn                     []string   `json:"issn,omitempty"`
	IssnInvalid              []string   `json:"issnInvalid,omitempty"`
	Issnl                    []string   `json:"issnl,omitempty"`
	IssueNumber              []string   `json:"issueNumber,omitempty"`
	MatrixNumber             []string   `json:"matrixNumber,omitempty"`
	MusicPlate               []string   `json:"musicPlate,omitempty"`
	MusicPublisher           []string   `json:"musicPublisher,omitempty"`
	Urn                      []string   `json:"urn,omitempty"`
	VideoRecordingIdentifier []string   `json:"videoRecordingIdentifier,omitempty"`
}

type IsbnIsmn struct {
	Id        []string `json:"identifier,omitempty"`
	Qualifier []string `json:"qualifier,omitempty"`
}

type Note struct {
	Acquisition                     []string      `json:"acquisition,omitempty"`
	Action                          []string      `json:"action,omitempty"`
	AdditionalPhysicalForm          []NoteWithURL `json:"additionalPhysicalForm,omitempty"`
	Bibliography                    []string      `json:"bibliography,omitempty"`
	Binding                         []NoteWithURL `json:"binding,omitempty"`
	Bibliographical                 []NoteWithURL `json:"bibliographical,omitempty"`
	Citation                        []NoteWithURL `json:"citation,omitempty"`
	Credits                         []string      `json:"credits,omitempty"`
	Exhibitions                     []string      `json:"exhibitions,omitempty"`
	Funding                         []string      `json:"funding,omitempty"`
	General                         []string      `json:"general,omitempty"`
	Han                             []string      `json:"han,omitempty"`
	Language                        []string      `json:"language,omitempty"`
	MediumOfPerformance             []string      `json:"mediumOfPerformance,omitempty"`
	Other                           []string      `json:"other,omitempty"`
	Ownership                       []NoteWithURL `json:"ownership,omitempty"`
	Performers                      []string      `json:"performers,omitempty"`
	PreferredCitation               []string      `json:"preferredCitation,omitempty"`
	Publications                    []string      `json:"publications,omitempty"`
	Reproduction                    []string      `json:"reproduction,omitempty"`
	Restriction                     []NoteWithURL `json:"restriction,omitempty"`
	StatementOfResponsibility       []string      `json:"statementOfResponsibility,omitempty"`
	StatementOfResponsibilityAltRep []string      `json:"statementOfResponsibilityAltRep,omitempty"`
	SystemDetails                   []string      `json:"systemDetails,omitempty"`
	Thesis                          []string      `json:"thesis,omitempty"`
	Venue                           []string      `json:"venue,omitempty"`
	VersionIdentification           []string      `json:"versionIdentification,omitempty"`
}

type Name struct {
	Conference []map[string]Conference `json:"conference,omitempty"`
	Corporate  []map[string]Corporate  `json:"corporate,omitempty"`
	Family     []map[string]Family     `json:"family,omitempty"`
	Personal   []map[string]Personal   `json:"personal,omitempty"`
}

type Conference struct {
	Date            string   `json:"date,omitempty"`
	Description     []string `json:"description,omitempty"`
	EntityType      []string `json:"entityType,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	Level           string   `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	Role            []string `json:"role,omitempty"`
	UseFor          []string `json:"json,omitempty"`
	Variant         []string `json:"variant,omitempty"`
}

type Corporate struct {
	Description     []string `json:"description,omitempty"`
	EntityType      []string `json:"entityType,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	Level           string   `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	PlaceOfBusiness []string `json:"placeOfBusiness,omitempty"`
	Related         []string `json:"related,omitempty"`
	Role            []string `json:"role,omitempty"`
	UseFor          []string `json:"json,omitempty"`
	Variant         []string `json:"variant,omitempty"`
}

type Family struct {
	Date            string   `json:"date,omitempty"`
	EntityType      []string `json:"entityType,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	Level           string   `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	Role            []string `json:"role,omitempty"`
	UseFor          []string `json:"json,omitempty"`
	Variant         []string `json:"variant,omitempty"`
}

type Personal struct {
	Date            string   `json:"date,omitempty"`
	EntityType      []string `json:"entityType,omitempty"`
	Gender          string   `json:"gender,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	Level           string   `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	PlaceOfActivity []string `json:"placeOfActivity,omitempty"`
	PlaceOfBirth    []string `json:"placeOfBirth,omitempty"`
	Profession      []string `json:"profession,omitempty"`
	Related         []string `json:"related,omitempty"`
	Role            []string `json:"role,omitempty"`
	UseFor          []string `json:"json,omitempty"`
	Variant         []string `json:"variant,omitempty"`
}

type Geographic struct {
	Coordinates     []Coordinates `json:"coordinates,omitempty"`
	Description     []string      `json:"description,omitempty"`
	EntityType      []string      `json:"entityType,omitempty"`
	GeoNamesId      []string      `json:"geoNamesId,omitempty"`
	Identifier      string        `json:"identifier,omitempty"`
	Level           string        `json:"level,omitempty"`
	NamePart        string        `json:"namePart,omitempty"`
	OtherIdentifier []string      `json:"otherIdentifier,omitempty"`
	Related         []string      `json:"related,omitempty"`
	Role            []string      `json:"role,omitempty"`
	UseFor          []string      `json:"json,omitempty"`
	Variant         []string      `json:"variant,omitempty"`
}

type Coordinates struct {
	Lat string `json:"lat,omitempty"`
	Lon string `json:"lon,omitempty"`
}

type Topic struct {
	EntityType      []string `json:"entityType,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	Label           string   `json:"label,omitempty"`
	Level           string   `json:"level,omitempty"`
	Mapped          []string `json:"mapped,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	UseFor          []string `json:"json,omitempty"`
	Variant         []string `json:"variant,omitempty"`
}

type Mapping001 struct {
	Abstract        []string      `json:"abstract,omitempty"`
	AccessCondition []NoteWithURL `json:"accessCondition,omitempty"`
	Cartographics   Cartographics `json:"cartographics,omitempty"`
	//Classification  map[string]Classification `json:"classification,omitempty"`
	Date []DateRange `json:"date,omitempty"`
	//	DateRange           []DateRange                `json:"daterange,omitempty"`
	Extension           Extension                  `json:"extension,omitempty"`
	Files               map[string]json.RawMessage `json:"files,omitempty"`
	Fulltext            []string                   `json:"fulltext,omitempty"`
	Identifier          Identifier                 `json:"identifier,omitempty"`
	Language            []string                   `json:"language,omitempty"`
	Location            Location                   `json:"location,omitempty"`
	Name                Name                       `json:"name,omitempty"`
	Note                Note                       `json:"note,omitempty"`
	OriginInfo          OriginInfo                 `json:"originInfo,omitempty"`
	PhysicalDescription map[string]json.RawMessage `json:"physicalDescription,omitempty"`
	RecordIdentifier    []string                   `json:"recordIdentifier,omitempty"`
	RelatedItem         map[string]json.RawMessage `json:"relatedItem,omitempty"`
	Subject             Subject                    `json:"subject,omitempty"`
	TableOfContents     []string                   `json:"tableOfContents,omitempty"`
	TargetAudience      []string                   `json:"targetAudience,omitempty"`
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
