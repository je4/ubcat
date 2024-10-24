package mappingRDV

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/je4/ubcat/v2/pkg/schema"
)

type MappingRDV schema.UBSchema001

type Element struct {
	Text     string                     `json:"text,omitempty"`
	Link     string                     `json:"link,omitempty"`
	Extended map[string]json.RawMessage `json:"extended,omitempty"`
}

func (e Element) String() string {
	if e.Link != "" {
		return fmt.Sprintf("'%s:[%s]'", e.Text, e.Link)
	}
	return fmt.Sprintf("'%s'", e.Text)
}

func appendText(e *Element, text, separator string) {
	if text != "" {
		if e.Text != "" {
			e.Text += separator
		}
		e.Text += text
	}
}

func (m *MappingRDV) GetAbstract() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if len(m.Mapping.Abstract) == 0 {
		return
	}
	result = []Element{}
	key = "abstract"
	ok = true
	for _, v := range m.Mapping.Abstract {
		if v == "" {
			continue
		}
		e := Element{
			Text: v,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetNoteGeneral() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Note == nil {
		return
	}
	if len(m.Mapping.Note.General) == 0 {
		return
	}
	result = []Element{}
	key = "noteGeneral"
	ok = true
	for _, v := range m.Mapping.Note.General {
		if v == "" {
			continue
		}
		e := Element{
			Text: v,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoDistribution() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Distribution) == 0 {
		return
	}
	result = []Element{}
	key = "originInfoDistribution"
	ok = true
	for _, v := range m.Mapping.OriginInfo.Distribution {
		if v == nil {
			continue
		}

		e := Element{}

		appendText(&e, strings.Join(v.Place, ", "), "")
		appendText(&e, strings.Join(v.Publisher, ", "), " : ")
		appendText(&e, v.Date, ", ")

		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoDistributionPlace() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Distribution) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Distribution {
		if v == nil {
			continue
		}
		if len(v.Place) == 0 {
			continue
		}
		key = "originInfoDistributionPlace"
		ok = true
		e := Element{
			Text: strings.Join(v.Place, ", "),
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoDistributionDate() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Distribution) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Distribution {
		if v == nil {
			continue
		}
		if len(v.Date) == 0 {
			continue
		}
		key = "originInfoDistributionDate"
		ok = true
		e := Element{
			Text: v.Date,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoManufacture() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Manufacture) == 0 {
		return
	}
	result = []Element{}
	key = "originInfoManufacture"
	ok = true
	for _, v := range m.Mapping.OriginInfo.Manufacture {
		if v == nil {
			continue
		}

		e := Element{}

		appendText(&e, strings.Join(v.Place, ", "), "")
		appendText(&e, strings.Join(v.Publisher, ", "), " : ")
		appendText(&e, v.Date, ", ")

		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoManufacturePlace() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Manufacture) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Manufacture {
		if v == nil {
			continue
		}
		if len(v.Place) == 0 {
			continue
		}
		key = "originInfoManufacturePlace"
		ok = true
		e := Element{
			Text: strings.Join(v.Place, ", "),
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoManufactureDate() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Manufacture) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Manufacture {
		if v == nil {
			continue
		}
		if len(v.Date) == 0 {
			continue
		}
		key = "originInfoManufactureDate"
		ok = true
		e := Element{
			Text: v.Date,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoProduction() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Production) == 0 {
		return
	}
	result = []Element{}
	key = "originInfoProduction"
	ok = true
	for _, v := range m.Mapping.OriginInfo.Production {
		if v == nil {
			continue
		}

		e := Element{}

		appendText(&e, strings.Join(v.Place, ", "), "")
		appendText(&e, strings.Join(v.Publisher, ", "), " : ")
		appendText(&e, v.Date, ", ")

		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoProductionPlace() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Production) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Production {
		if v == nil {
			continue
		}
		if len(v.Place) == 0 {
			continue
		}
		key = "originInfoProductionPlace"
		ok = true
		e := Element{
			Text: strings.Join(v.Place, ", "),
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoProductionDate() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Production) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Production {
		if v == nil {
			continue
		}
		if len(v.Date) == 0 {
			continue
		}
		key = "originInfoProductionDate"
		ok = true
		e := Element{
			Text: v.Date,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoPublication() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Publication) == 0 {
		return
	}
	result = []Element{}
	key = "originInfoPublication"
	ok = true
	for _, v := range m.Mapping.OriginInfo.Publication {
		if v == nil {
			continue
		}

		e := Element{}

		appendText(&e, strings.Join(v.Place, ", "), "")
		appendText(&e, strings.Join(v.Publisher, ", "), " : ")
		appendText(&e, v.Date, ", ")

		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoPublicationPlace() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Publication) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Publication {
		if v == nil {
			continue
		}
		if len(v.Place) == 0 {
			continue
		}
		key = "originInfoPublicationPlace"
		ok = true
		e := Element{
			Text: strings.Join(v.Place, ", "),
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetOriginInfoPublicationDate() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.OriginInfo == nil {
		return
	}
	if len(m.Mapping.OriginInfo.Publication) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.OriginInfo.Publication {
		if v == nil {
			continue
		}
		if len(v.Date) == 0 {
			continue
		}
		key = "originInfoPublicationDate"
		ok = true
		e := Element{
			Text: v.Date,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetPhysicalDescriptionExtent() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.PhysicalDescription == nil {
		return
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return
	}
	result = []Element{}
	key = "physicalDescriptionExtent"
	ok = true
	for _, v := range m.Mapping.PhysicalDescription.Extent {
		if v == nil {
			continue
		}
		e := Element{
			Text: v.Extent,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetPhysicalDescriptionExtentDimensions() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.PhysicalDescription == nil {
		return
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return
	}
	result = []Element{}
	for _, v := range m.Mapping.PhysicalDescription.Extent {
		if v == nil {
			continue
		}
		if len(v.Dimensions) == 0 {
			continue
		}
		key = "physicalDescriptionExtentDimensions"
		ok = true
		e := Element{
			Text: v.Dimensions,
		}
		result = append(result, e)
	}
	return
}

// 300 $a : $b ; $c + $e
func (m *MappingRDV) GetPhysicalDescriptionExtentFull() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.PhysicalDescription == nil {
		return
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return
	}

	key = "physicalDescriptionExtentFull"
	ok = true
	result = []Element{}
	for _, v := range m.Mapping.PhysicalDescription.Extent {
		if v == nil {
			continue
		}
		e := Element{}
		appendText(&e, v.Extent, "")
		appendText(&e, v.PhysicalDetails, " : ")
		appendText(&e, v.Dimensions, " ; ")
		appendText(&e, v.AccompanyingMaterial, " + ")

		result = append(result, e)
	}
	return
}

// 300 $a : $b
func (m *MappingRDV) GetPhysicalDescriptionExtentShort() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.PhysicalDescription == nil {
		return
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return
	}

	result = []Element{}
	for _, v := range m.Mapping.PhysicalDescription.Extent {
		if v == nil {
			continue
		}
		if len(v.Extent) == 0 && len(v.PhysicalDetails) == 0 {
			continue
		}
		key = "physicalDescriptionExtentShort"
		ok = true
		e := Element{}
		appendText(&e, v.Extent, "")
		appendText(&e, v.PhysicalDetails, " : ")

		result = append(result, e)
	}
	return
}

// 300 $a, $c
func (m *MappingRDV) GetPhysicalDescriptionExtentAndDimensions() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.PhysicalDescription == nil {
		return
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return
	}

	result = []Element{}
	for _, v := range m.Mapping.PhysicalDescription.Extent {
		if v == nil {
			continue
		}
		if len(v.Extent) == 0 && len(v.Dimensions) == 0 {
			continue
		}
		key = "physicalDescriptionExtentAndDimensions"
		ok = true
		e := Element{}
		appendText(&e, v.Extent, "")
		appendText(&e, v.Dimensions, ", ")

		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetPhysicalDescriptionMedium() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.PhysicalDescription == nil {
		return
	}
	if len(m.Mapping.PhysicalDescription.Medium) == 0 {
		return
	}
	result = []Element{}
	key = "physicalDescriptionMedium"
	ok = true
	e := Element{
		Text: strings.Join(m.Mapping.PhysicalDescription.Medium, ", "),
	}
	result = append(result, e)

	return
}

func (m *MappingRDV) GetTitleInfoMainTitle() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.TitleInfo == nil {
		return
	}
	if len(m.Mapping.TitleInfo.Main) == 0 {
		return
	}
	result = []Element{}
	key = "titleInfoMainTitle"
	ok = true
	for _, v := range m.Mapping.TitleInfo.Main {
		if v == nil {
			continue
		}
		e := Element{
			Text: strings.ReplaceAll(strings.ReplaceAll(v.Title, "<<", ""), ">>", ""),
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetTitleInfoMain() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.TitleInfo == nil {
		return
	}
	if len(m.Mapping.TitleInfo.Main) == 0 {
		return
	}
	result = []Element{}
	key = "titleInfoMain"
	ok = true
	for _, v := range m.Mapping.TitleInfo.Main {
		if v == nil {
			continue
		}

		e := Element{}

		appendText(&e, strings.ReplaceAll(strings.ReplaceAll(v.Title, "<<", ""), ">>", ""), "")
		appendText(&e, strings.Join(v.PartNumber, ", "), ". ")
		appendText(&e, strings.Join(v.PartName, ", "), ". ")
		appendText(&e, v.SubTitle, " : ")

		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetLocationDigital() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Location == nil {
		return
	}
	if len(m.Mapping.Location.Digital) == 0 {
		return
	}
	key = "locationDigital"
	ok = true
	result = []Element{}
	for _, v := range m.Mapping.Location.Digital {
		if v == nil {
			continue
		}
		e := Element{
			Link: v.Url,
		}
		appendText(&e, v.Content, "")
		appendText(&e, v.Note, " ; ")
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetLocationHoldingCallNumber() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Location == nil {
		return
	}
	if len(m.Mapping.Location.Holding) == 0 {
		return
	}

	result = []Element{}
	for _, v := range m.Mapping.Location.Holding {
		if v == nil {
			continue
		}
		if len(v.CallNumber) == 0 {
			continue
		}
		key = "locationHoldingCallNumber"
		ok = true
		e := Element{
			Text: v.CallNumber,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetLocationHolding() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Location == nil {
		return
	}
	if len(m.Mapping.Location.Holding) == 0 {
		return
	}

	libraryNames := map[string]string{
		"A100": "UB Hauptbibliothek",
		"A125": "UB Wirtschaft",
	}

	result = []Element{}
	for _, v := range m.Mapping.Location.Holding {
		if v == nil {
			continue
		}
		if len(v.CallNumber) == 0 {
			continue
		}
		key = "locationHolding"
		ok = true

		libraryName := v.Library
		if libraryNames, exists := libraryNames[v.Library]; exists {
			libraryName = libraryNames
		}

		e := Element{
			Text: libraryName + ", " + v.CallNumber,
		}
		result = append(result, e)
	}
	return
}

// todo: fix URL, encoded & zu /u0026
func (m *MappingRDV) GetLocationElectronic() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Location == nil {
		return
	}
	if len(m.Mapping.Location.Electronic) == 0 {
		return
	}
	key = "locationElectronic"
	ok = true
	result = []Element{}
	for _, v := range m.Mapping.Location.Electronic {
		if v == nil {
			continue
		}
		e := Element{
			Link: v.Url,
		}
		e.Extended = map[string]json.RawMessage{}
		if len(v.Collection) > 0 {
			e.Text = v.Collection
		} else {
			e.Text = v.Note
		}
		if len(v.Coverage) > 0 {
			coverageBytes, _ := json.Marshal(v.Coverage)
			e.Extended["coverage"] = coverageBytes
		}
		if v.Library == "AFREE" {
			openAccessBytes, _ := json.Marshal(true)
			e.Extended["openAccess"] = openAccessBytes
		} else {
			openAccessBytes, _ := json.Marshal(false)
			e.Extended["openAccess"] = openAccessBytes
		}

		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetSwisscollectionsUrl() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return
	}
	key = "swisscollectionsUrl"
	ok = true
	result = []Element{}
	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}
		if ok, _ := regexp.MatchString("^99.*5501$", v); ok {
			result = append(result, Element{Link: "https://swisscollections.ch/Record/" + v})
		}
	}
	if len(result) == 0 {
		return "", nil, false
	}
	return
}

// todo: fix URL, encoded & zu /u0026
func (m *MappingRDV) GetSwisscoveryUrl() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return
	}
	key = "swisscoveryUrl"
	ok = true
	result = []Element{}
	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}
		if ok, _ := regexp.MatchString("^99.*5504$", v); ok {
			result = append(result, Element{Link: "https://basel.swisscovery.org/discovery/fulldisplay?docid=alma" + v + "&context=L&vid=41SLSP_UBS:live"})
		}
	}
	return
}

func (m *MappingRDV) GetExtensionPortraetsContact() (key string, result []Element, ok bool) {
	if m.Flags == nil {
		return
	}

	if m.Mapping == nil {
		return
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return
	}

	var id string
	result = []Element{}
	key = "extensionPortraetsContact"
	ok = true
	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}
		match, _ := regexp.MatchString("^99.*5504$", v)
		if match {
			id = v
			break
		}
	}

	for _, v := range m.Flags {
		if v != "portraets" {
			continue
		}
		e := Element{
			Link: "mailto:hss-ub@unibas.ch?subject=Portraetsammlung, Anmerkungen zum Katalogisat" + id,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetResourceType() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "resourceType"
	ok = true
	result = []Element{}

	for _, sf := range m.Facets.Strings {
		if sf.Name != "resourceType" {
			continue
		}
		for _, s := range sf.String {
			e := Element{
				Text: s,
			}
			result = append(result, e)
		}
	}

	return
}

func (m *MappingRDV) GetFacetGeneralAuthor() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetGeneralAuthor"
	ok = true
	result = []Element{}

	for _, af := range m.Facets.Agents {
		if af.Name != "author" {
			continue
		}
		for _, a := range af.Agent {
			e := Element{
				Text: a.Label,
			}
			if len(a.Identifier) > 0 {
				e.Link = fmt.Sprintf("facet:author:%s", a.Identifier[0])
			}
			if len(a.Role) > 0 {
				e.Extended = map[string]json.RawMessage{}
				roleBytes, _ := json.Marshal(a.Role)
				e.Extended["roles"] = roleBytes
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetFacetAutographScribe() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetAutographScribe"
	ok = true
	result = []Element{}

	for _, af := range m.Facets.Agents {
		if af.Name != "scribe" {
			continue
		}
		for _, a := range af.Agent {
			e := Element{
				Text: a.Label,
			}
			if len(a.Identifier) > 0 {
				e.Link = fmt.Sprintf("facet:scribe:%s", a.Identifier[0])
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetFacetPortraetsPictured() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetPortraetsPictured"
	ok = true
	result = []Element{}

	for _, af := range m.Facets.Agents {
		if af.Name != "pictured" {
			continue
		}
		for _, a := range af.Agent {
			e := Element{
				Text: a.Label,
			}
			if len(a.Identifier) > 0 {
				e.Link = fmt.Sprintf("facet:scribe:%s", a.Identifier[0])
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetThumbnail() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Files == nil {
		return
	}

	key = "thumbnail"
	ok = true
	result = []Element{}

	for _, file := range m.Mapping.Files {
		if file.Media == nil {
			continue
		}
		if file.Media.Poster == nil {
			continue
		}

		e := Element{
			Link:     file.Media.Poster.Uri,
			Extended: map[string]json.RawMessage{},
		}

		widthBytes, _ := json.Marshal(file.Media.Poster.Width)
		e.Extended["width"] = widthBytes
		heightBytes, _ := json.Marshal(file.Media.Poster.Height)
		e.Extended["height"] = heightBytes
		result = append(result, e)
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return

}

func (m *MappingRDV) GetFileCount() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Files == nil {
		return
	}

	key = "fileCount"
	ok = true

	counts := make(map[string]int)

	// List of media types to be counted
	mediaTypes := []string{"image", "video"}

	for _, file := range m.Mapping.Files {
		if file.Media != nil && file.Media.Medias != nil {
			for _, media := range file.Media.Medias {
				for _, t := range mediaTypes {
					if media.Type == t {
						counts[media.Type]++
						continue
					}
				}
			}
		}
	}

	translations := map[string]string{
		"image": "Bild",
		"video": "Video",
	}

	var countsStrs []string
	for mediaType, count := range counts {
		if translation, exists := translations[mediaType]; exists {
			mediaType = translation
		}
		if count > 1 {
			if mediaType == "Bild" {
				mediaType += "er"
			} else {
				mediaType += "s"
			}
		}

		countsStrs = append(countsStrs, strconv.Itoa(count)+" "+mediaType)
	}

	if len(countsStrs) == 0 {
		ok = false
		return
	}

	result = append(result, Element{Text: strings.Join(countsStrs, ", ")})

	return
}

func (m *MappingRDV) Map() (result map[string][]Element) {
	result = map[string][]Element{}
	if m.Mapping == nil {
		return
	}
	key, value, ok := m.GetTitleInfoMainTitle()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetTitleInfoMain()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetPhysicalDescriptionExtent()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetPhysicalDescriptionExtentDimensions()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetPhysicalDescriptionExtentFull()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetPhysicalDescriptionExtentShort()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetPhysicalDescriptionExtentAndDimensions()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetPhysicalDescriptionMedium()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoDistribution()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoDistributionPlace()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoDistributionDate()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoManufacture()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoManufacturePlace()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoManufactureDate()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoProduction()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoProductionPlace()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoProductionDate()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoPublication()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoPublicationPlace()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoPublicationDate()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNoteGeneral()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetAbstract()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetLocationHoldingCallNumber()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetLocationHolding()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetGeneralAuthor()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetAutographScribe()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetPortraetsPictured()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetLocationDigital()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSwisscollectionsUrl()
	if ok {
		result[key] = value
	}
	/*key, value, ok = m.GetSwisscoveryUrl()
	if ok {
		result[key] = value
	}*/
	key, value, ok = m.GetExtensionPortraetsContact()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetResourceType()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetLocationElectronic()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFileCount()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetThumbnail()
	if ok {
		result[key] = value
	}
	return
}
