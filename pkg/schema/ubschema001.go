package schema

import (
	"emperror.dev/errors"
	"encoding/json"
	"regexp"
	"time"
)

type DataField struct {
	Tag      string   `json:"tag"`
	Ind1     string   `json:"ind1"`
	Ind2     string   `json:"ind2"`
	Subfield []string `json:"subfield"`
}

type Classification struct {
	Ddc   *Ddc     `json:"ddc,omitempty"`
	Local *Local   `json:"local,omitempty"`
	Rvk   []string `json:"rvk,omitempty"`
	Sdnb  []string `json:"sdnb,omitempty"`
	Udc   []string `json:"udc,omitempty"`
}

type Ddc struct {
	All  []string `json:"all,omitempty"`
	Ed23 []string `json:"ed23,omitempty"`
	//General []string
	Sdnb []string `json:"sdnb,omitempty"`
}

type Local struct {
	All []string `json:"all,omitempty"`
}

type Location struct {
	Digital    []*Url        `json:"digital,omitempty"`
	Electronic []*Electronic `json:"electronic,omitempty"`
	Holding    []Holding     `json:"holding,omitempty"`
}

type Holding struct {
	CallNumber  string  `json:"callNumber,omitempty"`
	Item        []*Item `json:"item,omitempty"`
	Library     string  `json:"library,omitempty"`
	Location    string  `json:"location,omitempty"`
	Note        string  `json:"note,omitempty"`
	Summary     string  `json:"summary,omitempty"`
	SummaryNote string  `json:"summaryNote,omitempty"`
}

type Item struct {
	ItemCallNumber string `json:"itemCallNumber,omitempty"`
	Note           string `json:"notel,omitempty"`
}

type Electronic struct {
	Library       string `json:"library,omitempty"`
	Collection    string `json:"collection,omitempty"`
	Url           string `json:"url,omitempty"`
	AlmaAccessUrl string `json:"almaAccessUrl,omitempty"`
	Interface     string `json:"interface,omitempty"`
	Coverage      string `json:"coverage,omitempty"`
	Note          string `json:"note,omitempty"`
}

type Subject struct {
	Genre      map[string][]string     `json:"genre,omitempty"`
	Geographic []map[string]Geographic `json:"geographic,omitempty"`
	Local      map[string][]string     `json:"local,omitempty"`
	Music      []string                `json:"music,omitempty"`
	Name       *Name                   `json:"name,omitempty"`
	Temporal   map[string][]string     `json:"temporal,omitempty"`
	TitleInfo  []map[string]Work       `json:"titleInfo,omitempty"`
	Topic      []map[string]Topic      `json:"topic,omitempty"`
	Undefined  []string                `json:"undefined,omitempty"`
}

type OriginInfo struct {
	CopyrightDate []string                `json:"copyrightDate,omitempty"`
	Distribution  []*PublicationNote      `json:"distribution,omitempty"`
	Edition       []string                `json:"edition,omitempty"`
	EditionAltRep []string                `json:"editionAltRep,omitempty"`
	Geographic    []map[string]Geographic `json:"geographic,omitempty"`
	Manufacture   []*PublicationNote      `json:"manufacture,omitempty"`
	Production    []*PublicationNote      `json:"production,omitempty"`
	Publication   []*PublicationNote      `json:"publication,omitempty"`
}

type PublicationNote struct {
	Date        string   `json:"date,omitempty"`
	LinkedField string   `json:"linkedField,omitempty"`
	Place       []string `json:"place,omitempty"`
	Publisher   []string `json:"publisher,omitempty"`
}

type NoteWithURL struct {
	Add     string   `json:"add,omitempty"`
	Licence string   `json:"licence,omitempty"`
	Main    string   `json:"main,omitempty"`
	Url     []string `json:"url,omitempty"`
}

type Cartographics struct {
	Coded   []*Coded   `json:"coded,omitempty"`
	Display []*Display `json:"display,omitempty"`
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
	Doi                      []string    `json:"doi,omitempty"`
	Isbn                     []*IsbnIsmn `json:"isbn,omitempty"`
	IsbnInvalid              []string    `json:"isbnInvalid,omitempty"`
	Ismn                     []*IsbnIsmn `json:"ismn,omitempty"`
	Issn                     []string    `json:"issn,omitempty"`
	IssnInvalid              []string    `json:"issnInvalid,omitempty"`
	Issnl                    []string    `json:"issnl,omitempty"`
	IssueNumber              []string    `json:"issueNumber,omitempty"`
	MatrixNumber             []string    `json:"matrixNumber,omitempty"`
	MusicPlate               []string    `json:"musicPlate,omitempty"`
	MusicPublisher           []string    `json:"musicPublisher,omitempty"`
	Urn                      []string    `json:"urn,omitempty"`
	VideoRecordingIdentifier []string    `json:"videoRecordingIdentifier,omitempty"`
}

type IsbnIsmn struct {
	Id        string   `json:"identifier,omitempty"`
	Qualifier []string `json:"qualifier,omitempty"`
}

type Note struct {
	Acquisition                     []string       `json:"acquisition,omitempty"`
	Action                          []string       `json:"action,omitempty"`
	AdditionalPhysicalForm          []*NoteWithURL `json:"additionalPhysicalForm,omitempty"`
	Bibliography                    []string       `json:"bibliography,omitempty"`
	Binding                         []*NoteWithURL `json:"binding,omitempty"`
	Bibliographical                 []*NoteWithURL `json:"bibliographical,omitempty"`
	Citation                        []*NoteWithURL `json:"citation,omitempty"`
	Credits                         []string       `json:"credits,omitempty"`
	Exhibitions                     []string       `json:"exhibitions,omitempty"`
	Funding                         []string       `json:"funding,omitempty"`
	General                         []string       `json:"general,omitempty"`
	Han                             []string       `json:"han,omitempty"`
	Language                        []string       `json:"language,omitempty"`
	MediumOfPerformance             []string       `json:"mediumOfPerformance,omitempty"`
	Other                           []string       `json:"other,omitempty"`
	Ownership                       []*NoteWithURL `json:"ownership,omitempty"`
	Performers                      []string       `json:"performers,omitempty"`
	PreferredCitation               []string       `json:"preferredCitation,omitempty"`
	Publications                    []string       `json:"publications,omitempty"`
	Reproduction                    []string       `json:"reproduction,omitempty"`
	Restriction                     []*NoteWithURL `json:"restriction,omitempty"`
	StatementOfResponsibility       []string       `json:"statementOfResponsibility,omitempty"`
	StatementOfResponsibilityAltRep []string       `json:"statementOfResponsibilityAltRep,omitempty"`
	SystemDetails                   []string       `json:"systemDetails,omitempty"`
	Thesis                          []string       `json:"thesis,omitempty"`
	Venue                           []string       `json:"venue,omitempty"`
	VersionIdentification           []string       `json:"versionIdentification,omitempty"`
}

type Name struct {
	Conference []map[string]Conference `json:"conference,omitempty"`
	Corporate  []map[string]Corporate  `json:"corporate,omitempty"`
	Family     []map[string]Family     `json:"family,omitempty"`
	Personal   []map[string]Personal   `json:"personal,omitempty"`
}

type Date string

var dateFormats = []string{
	"2006",
	"2006-01",
	"2006-01-02",
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
}

func (d Date) GetTime() (time.Time, error) {
	for _, format := range dateFormats {
		t, err := time.Parse(format, string(d))
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("cannot parse date " + string(d))
}

type Conference struct {
	Date        string   `json:"date,omitempty"`
	Description []string `json:"description,omitempty"`
	// EntityType      []string `json:"entityType,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	// Level           string   `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	Role            []string `json:"role,omitempty"`
	// UseFor          []string `json:"json,omitempty"`
	Variant []string `json:"variant,omitempty"`
}

type Corporate struct {
	Description []string `json:"description,omitempty"`
	// EntityType      []string `json:"entityType,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	// Level           string   `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	PlaceOfBusiness []string `json:"placeOfBusiness,omitempty"`
	Related         []string `json:"related,omitempty"`
	Role            []string `json:"role,omitempty"`
	// UseFor          []string `json:"json,omitempty"`
	Variant []string `json:"variant,omitempty"`
}

type Family struct {
	Date string `json:"date,omitempty"`
	// EntityType      []string `json:"entityType,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	// Level           string   `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	Role            []string `json:"role,omitempty"`
	// UseFor          []string `json:"json,omitempty"`
	Variant []string `json:"variant,omitempty"`
}

type PersonalDate struct {
	Original string `json:"original,omitempty"`
	Birth    Date   `json:"birth,omitempty"`
	Death    Date   `json:"death,omitempty"`
}

func (p *PersonalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Original)
}

var personalDateRegexp = regexp.MustCompile(`^(.[^-]+)?-([^-]+)?$`)

func (p *PersonalDate) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return errors.Wrap(err, "cannot unmarshal personal date")
	}
	p.Original = str
	matches := personalDateRegexp.FindStringSubmatch(str)
	if len(matches) == 3 {
		p.Birth = Date(matches[1])
		p.Death = Date(matches[2])
	}
	return nil
}

type Personal struct {
	Date *PersonalDate `json:"date,omitempty"`
	// EntityType      []string      `json:"entityType,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	// Level           string        `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	PlaceOfActivity []string `json:"placeOfActivity,omitempty"`
	PlaceOfBirth    []string `json:"placeOfBirth,omitempty"`
	Profession      []string `json:"profession,omitempty"`
	Related         []string `json:"related,omitempty"`
	Role            []string `json:"role,omitempty"`
	// UseFor          []string      `json:"json,omitempty"`
	Variant []string `json:"variant,omitempty"`
}

type Geographic struct {
	Coordinates []*Coordinates `json:"coordinates,omitempty"`
	Description []string       `json:"description,omitempty"`
	// EntityType      []string       `json:"entityType,omitempty"`
	GeoNamesId []string `json:"geoNamesId,omitempty"`
	Identifier string   `json:"identifier,omitempty"`
	// Level           string         `json:"level,omitempty"`
	NamePart        string   `json:"namePart,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	Related         []string `json:"related,omitempty"`
	Role            []string `json:"role,omitempty"`
	// UseFor          []string       `json:"json,omitempty"`
	Variant []string `json:"variant,omitempty"`
}

type Work struct {
	// EntityType      []string `json:"entityType,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	// Level           string   `json:"level,omitempty"`
	Name            string   `json:"namePart,omitempty"`
	Title           string   `json:"title,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	// UseFor          []string `json:"json,omitempty"`
	Variant []string `json:"variant,omitempty"`
}

type Coordinates struct {
	Lat string `json:"lat,omitempty"`
	Lon string `json:"lon,omitempty"`
}

type Topic struct {
	// EntityType      []string `json:"entityType,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Label      string `json:"label,omitempty"`
	// Level           string   `json:"level,omitempty"`
	Mapped          []string `json:"mapped,omitempty"`
	OtherIdentifier []string `json:"otherIdentifier,omitempty"`
	// UseFor          []string `json:"json,omitempty"`
	Variant []string `json:"variant,omitempty"`
}

type PhysicalDescription struct {
	Arrangement               []string `json:"arrangement,omitempty"`
	DateSequentialDesignation []string `json:"dateSequentialDesignation,omitempty"`
	Extent                    []Extent `json:"extent,omitempty"`
	Frequency                 []string `json:"frequency,omitempty"`
	Medium                    []string `json:"medium,omitempty"`
	NotatedMusic              []string `json:"notatedMusic,omitempty"`
}

type Extent struct {
	Extent string `json:"extent,omitempty"`
}

type RelatedItem struct {
	Constituent  []*Related        `json:"constituent,omitempty"`
	Host         []*Host           `json:"host,omitempty"`
	IssuedWith   []*Related        `json:"issuedWith,omitempty"`
	Location     []*Url            `json:"location,omitempty"`
	Original     []string          `json:"original,omitempty"`
	OriginalNote []string          `json:"originalNote,omitempty"`
	Other        []*Related        `json:"other,omitempty"`
	OtherFormat  []*Related        `json:"otherFormat,omitempty"`
	OtherVersion []*Related        `json:"otherVersion,omitempty"`
	Preceding    []*Related        `json:"preceding,omitempty"`
	Series       []*Series         `json:"series,omitempty"`
	Succeeding   []*Related        `json:"succeeding,omitempty"`
	Work         []map[string]Work `json:"work,omitempty"`
}

type Series struct {
	Conference         *Conference `json:"conference,omitempty"`
	Corporate          *Corporate  `json:"corporate,omitempty"`
	InternalIdentifier string      `json:"internalIdentifier,omitempty"`
	LinkedField        string      `json:"linkedField,omitempty"`
	PartName           []string    `json:"partName,omitempty"`
	PartNumber         []string    `json:"partNumber,omitempty"`
	Personal           *Personal   `json:"personal,omitempty"`
	Title              string      `json:"title,omitempty"`
	VolumeDesignation  string      `json:"volumeDesignation,omitempty"`
}

type Host struct {
	Title              string   `json:"title,omitempty"`
	Publisher          string   `json:"publisher,omitempty"`
	PartYear           string   `json:"partYear,omitempty"`
	PartNumber         string   `json:"partNumber,omitempty"`
	Part               string   `json:"part,omitempty"`
	PartSort           string   `json:"partSort,omitempty"`
	Identifier         []string `json:"identifier,omitempty"`
	InternalIdentifier string   `json:"internalIdentifer,omitempty"`
}

type Url struct {
	Content string `json:"content,omitempty"`
	Format  string `json:"format,omitempty"`
	Url     string `json:"url,omitempty"`
	Note    string `json:"note,omitempty"`
}

type Related struct {
	DisplayConstant           string   `json:"displayConstant,omitempty"`
	Identifier                []string `json:"identifier,omitempty"`
	InternalIdentifier        string   `json:"internalIdentifier,omitempty"`
	VolumeDesignation         string   `json:"volumeDesignation,omitempty"`
	StatementOfResponsibility string   `json:"statementOfResponsibility,omitempty"`
	LinkedField               string   `json:"linkedField,omitempty"`
	Title                     string   `json:"title,omitempty"`
}

type TitleInfo struct {
	Abbreviated []string `json:"abbreviated,omitempty"`
	Alternative []*Title `json:"alternative,omitempty"`
	Main        []*Title `json:"main,omitempty"`
	Translated  []*Title `json:"translated,omitempty"`
	Uniform     []*Title `json:"uniform,omitempty"`
}

type Title struct {
	Title       string   `json:"title,omitempty"`
	SubTitle    string   `json:"subTitle,omitempty"`
	PartName    []string `json:"partName,omitempty"`
	PartNumber  []string `json:"partNumber,omitempty"`
	LinkedField string   `json:"linkedField,omitempty"`
}

type Files struct {
	Media     *Media                     `json:"media,omitempty"`
	Structure map[string]json.RawMessage `json:"structure,omitempty"`
}

type Media struct {
	Medias []*Medias `json:"medias,omitempty"`
	Poster *Medias   `json:"poster,omitempty"`
}

type Medias struct {
	Height   int64  `json:"height,omitempty"`
	Licence  string `json:"licence,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
	Path     string `json:"path,omitempty"`
	Pronom   string `json:"pronom,omitempty"`
	Type     string `json:"type,omitempty"`
	Uri      string `json:"uri,omitempty"`
	Width    int64  `json:"width,omitempty"`
}

type Mapping001 struct {
	Abstract        []string        `json:"abstract,omitempty"`
	AccessCondition []*NoteWithURL  `json:"accessCondition,omitempty"`
	Cartographics   *Cartographics  `json:"cartographics,omitempty"`
	Classification  *Classification `json:"classification,omitempty"`
	Date            []*DateRange    `json:"date,omitempty"`
	//	DateRange           []DateRange                `json:"daterange,omitempty"`
	Extension           *Extension           `json:"extension,omitempty"`
	Files               []*Files             `json:"files,omitempty"`
	Fulltext            []string             `json:"fulltext,omitempty"`
	Identifier          *Identifier          `json:"identifier,omitempty"`
	Language            []string             `json:"language,omitempty"`
	Location            *Location            `json:"location,omitempty"`
	Name                *Name                `json:"name,omitempty"`
	Note                *Note                `json:"note,omitempty"`
	OriginInfo          *OriginInfo          `json:"originInfo,omitempty"`
	PhysicalDescription *PhysicalDescription `json:"physicalDescription,omitempty"`
	RecordIdentifier    []string             `json:"recordIdentifier,omitempty"`
	RelatedItem         *RelatedItem         `json:"relatedItem,omitempty"`
	Subject             *Subject             `json:"subject,omitempty"`
	TableOfContents     []string             `json:"tableOfContents,omitempty"`
	TargetAudience      []string             `json:"targetAudience,omitempty"`
	TitleInfo           *TitleInfo           `json:"titleInfo,omitempty"`
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
	Mapping        *Mapping001   `json:"mapping,omitempty"`
	Facets         []Facets      `json:"facets,omitempty"`
	Sets           []string      `json:"sets,omitempty"`
	Flags          []string      `json:"flags,omitempty"`
	ACL            ACL           `json:"acl,omitempty"`
	EmbeddingProse [1536]float32 `json:"embedding_prose,omitempty"`
	EmbeddingMarc  [1536]float32 `json:"embedding_marc,omitempty"`
}
