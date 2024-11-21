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

func (m *MappingRDV) GetTableOfContents() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if len(m.Mapping.TableOfContents) == 0 {
		return
	}
	result = []Element{}
	key = "tableOfContents"
	ok = true
	for _, v := range m.Mapping.TableOfContents {
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

func (m *MappingRDV) GetNoteLanguage() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Note == nil {
		return
	}
	if len(m.Mapping.Note.Language) == 0 {
		return
	}
	result = []Element{}
	key = "noteLanguage"
	ok = true
	for _, v := range m.Mapping.Note.Language {
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

func (m *MappingRDV) GetNoteOwnership() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Note == nil {
		return
	}
	if len(m.Mapping.Note.Ownership) == 0 {
		return
	}
	result = []Element{}
	key = "noteOwnership"
	ok = true
	for _, v := range m.Mapping.Note.Ownership {
		if v == nil {
			continue
		}
		e := Element{
			Text: v.Main,
		}
		if len(v.Url) != 0 {
			e.Link = v.Url[0]
		}
		result = append(result, e)
	}
	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetNoteCitation() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Note == nil {
		return
	}
	if len(m.Mapping.Note.Citation) == 0 {
		return
	}
	result = []Element{}
	key = "noteCitation"
	ok = true
	for _, v := range m.Mapping.Note.Citation {
		if v == nil {
			continue
		}
		e := Element{
			Text: v.Main,
		}

		appendText(&e, v.Add, " ")

		if len(v.Url) != 0 {
			e.Link = v.Url[0]
		}
		result = append(result, e)
	}
	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetNotePublications() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Note == nil {
		return
	}
	if len(m.Mapping.Note.Publications) == 0 {
		return
	}
	result = []Element{}
	key = "noteLanguage"
	ok = true
	for _, v := range m.Mapping.Note.Publications {
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

func (m *MappingRDV) GetTitleInfoAlternative() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.TitleInfo == nil {
		return
	}
	if len(m.Mapping.TitleInfo.Alternative) == 0 {
		return
	}
	result = []Element{}
	key = "titleInfoAlternative"
	ok = true
	for _, v := range m.Mapping.TitleInfo.Alternative {
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

func (m *MappingRDV) GetTitleInfoUniform() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.TitleInfo == nil {
		return
	}
	if len(m.Mapping.TitleInfo.Uniform) == 0 {
		return
	}
	result = []Element{}
	key = "titleInfoUniform"
	ok = true
	for _, v := range m.Mapping.TitleInfo.Uniform {
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
	if len(result) == 0 {
		return "", nil, false
	}
	return
}

/*
todo: change data in index: Index generic values in German, adapt this function to get the code and only one per document
usually used for selecting icon and configuration in frontend
*/
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
			break
		}
	}

	return
}

// GetResourceTypeGeneric todo: change data in index, will make the translation obsolete
func (m *MappingRDV) GetResourceTypeGeneric() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "resourceTypeGeneric"
	ok = true
	result = []Element{}

	translations := map[string]string{
		"article":             "Artikel",
		"book":                "Buch",
		"journal":             "Zeitschrift",
		"collection":          "Dokumentensammlung",
		"sheetMusic":          "Musiknoten",
		"map":                 "Karte",
		"image":               "Bild",
		"movie":               "Video",
		"soundRecording":      "Tonaufnahme",
		"musicRecording":      "Musikaufnahme",
		"computerMedia":       "Computermedien",
		"educationalResource": "Lehrmittel",
		"object":              "Objekt",
		"manuscript":          "Manuskript",
		"archivalMaterial":    "Archivmaterial",
	}

	for _, sf := range m.Facets.Strings {
		if sf.Name != "resourceType" {
			continue
		}
		for _, s := range sf.String {
			if translation, exists := translations[s]; exists {
				s = translation
			}
			e := Element{
				Text: s,
			}
			result = append(result, e)
		}
	}

	return
}

/* add specific resource types based on flags, obsolete if there's a facet for these resource types, works only if the flags are exclusive */
func (m *MappingRDV) GetResourceTypeView() (key string, result []Element, ok bool) {
	if m.Flags == nil {
		return
	}

	if m.Mapping == nil {
		return
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return
	}

	result = []Element{}
	key = "resourceTypeView"
	ok = true

	flagsToTypes := map[string]string{
		"portraets":  "Porträt",
		"autograph":  "Autograph",
		"burckhardt": "Mappentitel",
	}

	for _, v := range m.Flags {

		if flagsToTypes, exists := flagsToTypes[v]; exists {
			e := Element{
				Text: flagsToTypes,
			}
			result = append(result, e)
		}

	}
	if len(result) == 0 {
		return "", nil, false
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

func (m *MappingRDV) GetFacetAutographRecipient() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetAutographRecipient"
	ok = true
	result = []Element{}

	for _, af := range m.Facets.Agents {
		if af.Name != "recipient" {
			continue
		}
		for _, a := range af.Agent {
			e := Element{
				Text: a.Label,
			}
			if len(a.Identifier) > 0 {
				e.Link = fmt.Sprintf("facet:recipient:%s", a.Identifier[0])
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetFacetAutographFormerOwner() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetAutographFormerOwner"
	ok = true
	result = []Element{}

	for _, af := range m.Facets.Agents {
		if af.Name != "formerOwner" {
			continue
		}
		for _, a := range af.Agent {
			e := Element{
				Text: a.Label,
			}
			if len(a.Identifier) > 0 {
				e.Link = fmt.Sprintf("facet:formerOwner:%s", a.Identifier[0])
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetFacetAutographSeller() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetAutographSeller"
	ok = true
	result = []Element{}

	for _, af := range m.Facets.Agents {
		if af.Name != "seller" {
			continue
		}
		for _, a := range af.Agent {
			e := Element{
				Text: a.Label,
			}
			if len(a.Identifier) > 0 {
				e.Link = fmt.Sprintf("facet:seller:%s", a.Identifier[0])
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetFacetAutographGeigyNummer() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetAutographGeigyNummer"
	ok = true
	result = []Element{}

	for _, sf := range m.Facets.Strings {
		if sf.Name != "geigyNummer" {
			continue
		}
		for _, s := range sf.String {
			e := Element{
				Text: s,
			}
			result = append(result, e)
			break
		}
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

func (m *MappingRDV) GetFacetConceptPublicationPlace() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetConceptPublicationPlace"
	ok = true
	result = []Element{}

	for _, af := range m.Facets.Concepts {
		if af.Name != "publicationPlace" {
			continue
		}
		for _, a := range af.Concept {
			e := Element{
				Text: a.Label,
			}
			if len(a.Identifier) > 0 {
				e.Link = fmt.Sprintf("facet:publicationPlace:%s", a.Identifier[0])
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetFacetGenre() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return
	}

	key = "facetGenre"
	ok = true
	result = []Element{}

	for _, sf := range m.Facets.Strings {
		if sf.Name != "genre" {
			continue
		}
		for _, s := range sf.String {
			e := Element{
				Text: s,
			}
			result = append(result, e)
			break
		}
	}

	return
}

func (m *MappingRDV) GetThumbnail() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Files == nil {
		m.GetThumbnailFromData()
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
			// Link:     strings.Replace(file.Media.Poster.Uri, "mediaserver:", "https://mediaservermain.ub-dlza-test.k8s-001.unibas.ch/iiif/3/", 1) + "/full/{$widthInPx},/0/default.jpg",
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
		key, result, ok = m.GetThumbnailFromData()
	}
	return

}

func (m *MappingRDV) GetThumbnailFromData() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Location == nil {
		return
	}
	if len(m.Mapping.Location.Digital) == 0 {
		return
	}

	key = "thumbnail"
	ok = true
	result = []Element{}

	for _, v := range m.Mapping.Location.Digital {
		if v == nil {
			continue
		}
		portraitUrlPattern := regexp.MustCompile(`.*digi/a100/portraet/bs`)
		portraitIdPattern := regexp.MustCompile(`^.*/([^./]+)\.[^/]+$`)

		if portraitUrlPattern.MatchString(v.Url) {
			matches := portraitIdPattern.FindStringSubmatch(v.Url)
			if len(matches) > 1 {
				portraitId := matches[1]
				portraitImageUrl := "https://ub-sipi.ub.unibas.ch/portraets/" + portraitId + ".jpx/full/{$widthInPx},/0/default.jpg"
				e := Element{
					Link: portraitImageUrl,
				}
				result = append(result, e)
			}
		}
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

/* toDo: extend for other collections with images */
func (m *MappingRDV) GetObjectPreview() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Location == nil {
		return
	}
	if len(m.Mapping.Location.Digital) == 0 {
		return
	}

	key = "objectPreview"
	ok = true
	result = []Element{}

	for _, v := range m.Mapping.Location.Digital {
		if v == nil {
			continue
		}
		portraitUrlPattern := regexp.MustCompile(`.*digi/a100/portraet/bs`)
		portraitIdPattern := regexp.MustCompile(`^.*/([^./]+)\.[^/]+$`)

		if portraitUrlPattern.MatchString(v.Url) && portraitIdPattern.MatchString(v.Url) {
			e := Element{
				Text: "zoomImage",
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetMedia() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if m.Mapping.Files == nil {
		return
	}

	key = "media"
	ok = true
	result = []Element{}
	contentArray := []map[string]json.RawMessage{}

	licenseUrls := map[string]string{
		"PDM 1.0 Deed":      "https://creativecommons.org/public-domain/pdm/",
		"CC BY-SA 4.0 Deed": "https://creativecommons.org/licenses/by-sa/4.0/",
		"InC":               "https://rightsstatements.org/page/InC/1.0/",
	}

	for _, file := range m.Mapping.Files {
		if file.Media == nil {
			continue
		}
		if file.Media.Medias == nil {
			continue
		}

		for _, f := range file.Media.Medias {
			if f == nil {
				continue
			}

			licenseUrl := ""
			if licenseUrls, exists := licenseUrls[f.License_abbr]; exists {
				licenseUrl = licenseUrls
			}

			urlBytes, _ := json.Marshal(strings.Replace(f.Uri, "mediaserver:", "https://mediaservermain.ub-dlza-test.k8s-001.unibas.ch/iiif/3/", 1))
			licenseBytes, _ := json.Marshal(f.License_abbr)
			licenseUrlBytes, _ := json.Marshal(licenseUrl)
			presentationTypeBytes, _ := json.Marshal(f.Type)
			typeBytes, _ := json.Marshal(f.Type)
			mimetypeBytes, _ := json.Marshal(f.Mimetype)
			pronomBytes, _ := json.Marshal(f.Pronom)
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/" + f.Pronom)
			widthBytes, _ := json.Marshal(f.Width)
			heightBytes, _ := json.Marshal(f.Height)

			contentDetails := map[string]json.RawMessage{
				/*"title": "",*/
				/*"fileName": "",*/
				/*"arkQualifier": "",*/
				"url": urlBytes,
				/*"thumbnail": "",*/
				/*"downloadUrl": "",*/
				/*"acl": */
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				/*"format": "",*/
				/*"date": "",*/
				"width":  widthBytes,
				"height": heightBytes,
				/*"duration": "",*/
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails)
		}
	}

	if len(contentArray) > 0 {
		contentBytes, _ := json.Marshal(contentArray)
		e := Element{
			Extended: map[string]json.RawMessage{
				"content": contentBytes,
			},
		}
		result = append(result, e)
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return

}

// GetTranscription todo: replace once there's data in the index, currently only for testing
func (m *MappingRDV) GetTranscription() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return
	}

	key = "transcription"
	ok = true
	result = []Element{}
	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}
		if ok, _ := regexp.MatchString("^991170524020205501$", v); ok {
			content := "<?xml version='1.0' encoding='utf-8'?>\\n<div><p>Die höchste Weisheit, ist<br/>\\ndie höchste Güte.</p>\\n<div class=\\\"align-right\\\">\\n<p>Marie Ebner-Eschenbach.</p>\\n</div>\\n<p>Zdislavic, den 13<sup>ten</sup> Nov. 93.</p></div>"
			e := Element{
				Text: content,
			}
			result = append(result, e)
		}
	}
	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetIIIFManifest() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return
	}

	key = "iiifManifest"
	ok = true
	result = []Element{}

	/* for e-rara and e-manuscripta */
	for _, f := range m.Mapping.Files {
		if f == nil {
			continue
		}
		if f.Structure.DigitalObject.Id == "" {
			continue
		}
		for _, flag := range m.Flags {
			if flag == "e-rara" {
				e := Element{
					Link: "https://www.e-rara.ch/i3f/v20/" + strings.Replace(f.Structure.DigitalObject.Id, "md", "", 1) + "/manifest",
				}
				result = append(result, e)
			}
			if flag == "e-manuscripta" {
				e := Element{
					Link: "https://www.e-manuscripta.ch/i3f/v20/" + strings.Replace(f.Structure.DigitalObject.Id, "md", "", 1) + "/manifest",
				}
				result = append(result, e)
			}
		}
	}

	if m.Mapping.Location == nil {
		return
	}
	/* for Porträts */
	for _, v := range m.Mapping.Location.Digital {
		if v == nil {
			continue
		}
		portraitUrlPattern := regexp.MustCompile(`.*digi/a100/portraet/bs`)
		portraitIdPattern := regexp.MustCompile(`^.*/([^./]+)\.[^/]+$`)

		if portraitUrlPattern.MatchString(v.Url) && portraitIdPattern.MatchString(v.Url) {

			for _, id := range m.Mapping.RecordIdentifier {
				if id == "" {
					continue
				}
				if ok, _ := regexp.MatchString("^99.*5504$", id); ok {
					e := Element{
						Link: "https://ub-iiifpresentation.ub.unibas.ch/portraets/" + id + "/manifest/",
					}
					result = append(result, e)
				}
			}
		}
	}
	if len(result) == 0 {
		return "", nil, false
	}
	return

}

func (m *MappingRDV) GetAcl() (key string, result []Element, ok bool) {
	if m.ACL == nil {
		return
	}

	key = "acl"
	ok = true
	result = []Element{}

	aclMetaBytes, _ := json.Marshal(m.ACL.Meta)
	aclPreviewBytes, _ := json.Marshal(m.ACL.Preview)
	aclContentBytes, _ := json.Marshal(m.ACL.Content)
	e := Element{
		Extended: map[string]json.RawMessage{
			"meta":    aclMetaBytes,
			"preview": aclPreviewBytes,
			"content": aclContentBytes,
		},
	}
	result = append(result, e)

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
	key, value, ok = m.GetTitleInfoAlternative()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetTitleInfoUniform()
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
	key, value, ok = m.GetNoteOwnership()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNoteLanguage()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNoteOwnership()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNoteCitation()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNotePublications()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetAbstract()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetTableOfContents()
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
	key, value, ok = m.GetFacetAutographRecipient()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetAutographFormerOwner()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetAutographSeller()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetAutographGeigyNummer()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetPortraetsPictured()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetFacetConceptPublicationPlace()
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
	key, value, ok = m.GetResourceTypeGeneric()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetResourceTypeView()
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
	key, value, ok = m.GetTranscription()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetIIIFManifest()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetObjectPreview()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetAcl()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetMedia()
	if ok {
		result[key] = value
	}
	return
}
