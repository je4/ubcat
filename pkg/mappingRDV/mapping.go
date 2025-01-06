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

func appendTextInBrackets(e *Element, text string) {
	if text != "" {
		if e.Text != "" {
			e.Text += " ("
		}
		e.Text += text + ")"
	}
}

func (m *MappingRDV) GetAbstract() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if len(m.Mapping.Abstract) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if len(m.Mapping.TableOfContents) == 0 {
		return "", nil, false
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

func (m *MappingRDV) GetNoteStatementOfResponsibility() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Note == nil {
		return "", nil, false
	}
	if len(m.Mapping.Note.StatementOfResponsibility) == 0 {
		return "", nil, false
	}
	result = []Element{}
	key = "noteStatementOfResponsibility"
	ok = true
	for _, v := range m.Mapping.Note.StatementOfResponsibility {
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
		return "", nil, false
	}
	if m.Mapping.Note == nil {
		return "", nil, false
	}
	if len(m.Mapping.Note.General) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Note == nil {
		return "", nil, false
	}
	if len(m.Mapping.Note.Language) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Note == nil {
		return "", nil, false
	}
	if len(m.Mapping.Note.Ownership) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Note == nil {
		return "", nil, false
	}
	if len(m.Mapping.Note.Citation) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Note == nil {
		return "", nil, false
	}
	if len(m.Mapping.Note.Publications) == 0 {
		return "", nil, false
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

func (m *MappingRDV) GetOriginInfoEdition() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Edition) == 0 {
		return "", nil, false
	}
	result = []Element{}
	key = "originInfoEdition"
	ok = true
	for _, v := range m.Mapping.OriginInfo.Edition {
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Distribution) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Distribution) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Distribution) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Manufacture) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Manufacture) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Manufacture) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Production) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Production) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Production) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Publication) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Publication) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.OriginInfo.Publication) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.PhysicalDescription == nil {
		return "", nil, false
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.PhysicalDescription == nil {
		return "", nil, false
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.PhysicalDescription == nil {
		return "", nil, false
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.PhysicalDescription == nil {
		return "", nil, false
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.PhysicalDescription == nil {
		return "", nil, false
	}
	if len(m.Mapping.PhysicalDescription.Extent) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.PhysicalDescription == nil {
		return "", nil, false
	}
	if len(m.Mapping.PhysicalDescription.Medium) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.TitleInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.TitleInfo.Main) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.TitleInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.TitleInfo.Main) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.TitleInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.TitleInfo.Alternative) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.TitleInfo == nil {
		return "", nil, false
	}
	if len(m.Mapping.TitleInfo.Uniform) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Location == nil {
		return "", nil, false
	}
	if len(m.Mapping.Location.Digital) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Location == nil {
		return "", nil, false
	}
	if len(m.Mapping.Location.Holding) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Location == nil {
		return "", nil, false
	}
	if len(m.Mapping.Location.Holding) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if m.Mapping.Location == nil {
		return "", nil, false
	}
	if len(m.Mapping.Location.Electronic) == 0 {
		return "", nil, false
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

// GetSwisscollectionsUrl : link generated for all data, won't work for data which isn't in swisscollections
func (m *MappingRDV) GetSwisscollectionsUrl() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return "", nil, false
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
		return "", nil, false
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return "", nil, false
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
		return "", nil, false
	}

	if m.Mapping == nil {
		return "", nil, false
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return "", nil, false
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

func (m *MappingRDV) GetNamePersonal() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Name.Personal == nil {
		return "", nil, false
	}

	key = "namePersonal"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, personMap := range m.Mapping.Name.Personal {
		if personal, ok := personMap["gnd"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["idref"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["sbt"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["rero"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["unknown"]; ok {
			result = appendPersonal(personal, result)
		}

		for key, personal := range personMap {
			if personal.Identifier != "" {
				authorities[key] = personal.Identifier
			}
			for _, id := range personal.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func (m *MappingRDV) GetSubjectNamePersonal() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Subject == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name.Personal == nil {
		return "", nil, false
	}

	key = "subjectNamePersonal"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, personMap := range m.Mapping.Subject.Name.Personal {
		if personal, ok := personMap["gnd"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["idref"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["sbt"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["rero"]; ok {
			result = appendPersonal(personal, result)
		} else if personal, ok := personMap["unknown"]; ok {
			result = appendPersonal(personal, result)
		}

		for key, personal := range personMap {
			if personal.Identifier != "" {
				authorities[key] = personal.Identifier
			}
			for _, id := range personal.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func appendPersonal(personal schema.Personal, result []Element) []Element {
	if personal.NamePart != "" {
		e := Element{
			Text: personal.NamePart,
			Link: "facet:" + personal.NamePart,
		}
		if personal.Date != nil {
			appendTextInBrackets(&e, personal.Date.Original)
		}
		if len(personal.Role) > 0 {
			e.Extended = map[string]json.RawMessage{}
			roleBytes, _ := json.Marshal(personal.Role)
			e.Extended["roles"] = roleBytes
		}
		result = append(result, e)
	}
	return result
}

func appendAuthorityElements(authorities map[string]string) []Element {
	authorityElements := []Element{}
	for key, id := range authorities {
		var link string
		var text string
		switch key {
		case "gnd":
			link = "https://lobid.org/gnd/" + strings.Replace(id, "(DE-588)", "", 1)
			text = "GND"
		case "idref":
			link = "http://www.idref.fr/" + strings.Replace(id, "(IDREF)", "", 1) + "/id"
			text = "IdRef"
		case "orcid":
			link = "https://orcid.org/" + strings.Replace(id, "(orcid)", "", 1)
			text = "ORCID"
		case "geonames":
			link = "https://sws.geonames.org/" + strings.Replace(id, "(geonames)", "", 1)
			text = "GeoNames"
		}

		authorityElements = append(authorityElements, Element{
			Text: text,
			Link: link,
		})
	}
	return authorityElements
}

func (m *MappingRDV) GetNameFamily() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Name.Family == nil {
		return "", nil, false
	}

	key = "nameFamily"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, familyMap := range m.Mapping.Name.Family {
		if family, ok := familyMap["gnd"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["idref"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["sbt"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["rero"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["unknown"]; ok {
			result = appendFamily(family, result)
		}

		for key, family := range familyMap {
			if family.Identifier != "" {
				authorities[key] = family.Identifier
			}
			for _, id := range family.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func (m *MappingRDV) GetSubjectNameFamily() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Subject == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name.Family == nil {
		return "", nil, false
	}

	key = "subjectNameFamily"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, familyMap := range m.Mapping.Subject.Name.Family {
		if family, ok := familyMap["gnd"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["idref"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["sbt"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["rero"]; ok {
			result = appendFamily(family, result)
		} else if family, ok := familyMap["unknown"]; ok {
			result = appendFamily(family, result)
		}

		for key, family := range familyMap {
			if family.Identifier != "" {
				authorities[key] = family.Identifier
			}
			for _, id := range family.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func appendFamily(family schema.Family, result []Element) []Element {
	if family.NamePart != "" {
		e := Element{
			Text: family.NamePart,
			Link: "facet:" + family.NamePart,
		}
		if family.Date != "" {
			appendTextInBrackets(&e, family.Date)
		}
		if len(family.Role) > 0 {
			e.Extended = map[string]json.RawMessage{}
			roleBytes, _ := json.Marshal(family.Role)
			e.Extended["roles"] = roleBytes
		}
		result = append(result, e)
	}
	return result
}

func (m *MappingRDV) GetNameCorporate() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Name.Corporate == nil {
		return "", nil, false
	}

	key = "nameCorporate"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, corporateMap := range m.Mapping.Name.Corporate {
		if corporate, ok := corporateMap["gnd"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["idref"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["sbt"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["rero"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["unknown"]; ok {
			result = appendCorporate(corporate, result)
		}

		for key, corporate := range corporateMap {
			if corporate.Identifier != "" {
				authorities[key] = corporate.Identifier
			}
			for _, id := range corporate.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func (m *MappingRDV) GetSubjectNameCorporate() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Subject == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name.Corporate == nil {
		return "", nil, false
	}

	key = "subjectNameCorporate"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, corporateMap := range m.Mapping.Subject.Name.Corporate {
		if corporate, ok := corporateMap["gnd"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["idref"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["sbt"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["rero"]; ok {
			result = appendCorporate(corporate, result)
		} else if corporate, ok := corporateMap["unknown"]; ok {
			result = appendCorporate(corporate, result)
		}

		for key, corporate := range corporateMap {
			if corporate.Identifier != "" {
				authorities[key] = corporate.Identifier
			}
			for _, id := range corporate.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func appendCorporate(corporate schema.Corporate, result []Element) []Element {
	if corporate.NamePart != "" {
		e := Element{
			Text: corporate.NamePart,
			Link: "facet:" + corporate.NamePart,
		}
		if len(corporate.Role) > 0 {
			e.Extended = map[string]json.RawMessage{}
			roleBytes, _ := json.Marshal(corporate.Role)
			e.Extended["roles"] = roleBytes
		}
		result = append(result, e)
	}
	return result
}

func (m *MappingRDV) GetNameConference() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Name.Conference == nil {
		return "", nil, false
	}

	key = "nameConference"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, conferenceMap := range m.Mapping.Name.Conference {
		if conference, ok := conferenceMap["gnd"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["idref"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["sbt"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["rero"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["unknown"]; ok {
			result = appendConference(conference, result)
		}

		for key, conference := range conferenceMap {
			if conference.Identifier != "" {
				authorities[key] = conference.Identifier
			}
			for _, id := range conference.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func (m *MappingRDV) GetSubjectNameConference() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Subject == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Name.Conference == nil {
		return "", nil, false
	}

	key = "subjectNameConference"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, conferenceMap := range m.Mapping.Subject.Name.Conference {
		if conference, ok := conferenceMap["gnd"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["idref"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["sbt"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["rero"]; ok {
			result = appendConference(conference, result)
		} else if conference, ok := conferenceMap["unknown"]; ok {
			result = appendConference(conference, result)
		}

		for key, conference := range conferenceMap {
			if conference.Identifier != "" {
				authorities[key] = conference.Identifier
			}
			for _, id := range conference.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func appendConference(conference schema.Conference, result []Element) []Element {
	if conference.NamePart != "" {
		e := Element{
			Text: conference.NamePart,
			Link: "facet:" + conference.NamePart,
		}
		if conference.Date != "" {
			appendTextInBrackets(&e, conference.Date)
		}
		if len(conference.Role) > 0 {
			e.Extended = map[string]json.RawMessage{}
			roleBytes, _ := json.Marshal(conference.Role)
			e.Extended["roles"] = roleBytes
		}
		result = append(result, e)
	}
	return result
}

func (m *MappingRDV) GetSubjectTopic() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Subject == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Topic == nil {
		return "", nil, false
	}

	key = "subjectTopic"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, topicMap := range m.Mapping.Subject.Topic {
		if topic, ok := topicMap["gnd"]; ok {
			result = appendTopic(topic, result)
		} else if topic, ok := topicMap["idref"]; ok {
			result = appendTopic(topic, result)
		} else if topic, ok := topicMap["sbt"]; ok {
			result = appendTopic(topic, result)
		} else if topic, ok := topicMap["rero"]; ok {
			result = appendTopic(topic, result)
		} else if topic, ok := topicMap["unknown"]; ok {
			result = appendTopic(topic, result)
		}

		for key, topic := range topicMap {
			if topic.Identifier != "" {
				authorities[key] = topic.Identifier
			}
			for _, id := range topic.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func appendTopic(topic schema.Topic, result []Element) []Element {
	if topic.Label != "" {
		e := Element{
			Text: topic.Label,
			Link: "facet:" + topic.Label,
		}
		result = append(result, e)
	}
	return result
}

func (m *MappingRDV) GetNameGeographic() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.OriginInfo == nil {
		return "", nil, false
	}
	if m.Mapping.OriginInfo.Geographic == nil {
		return "", nil, false
	}

	key = "nameGeographic"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, geographicMap := range m.Mapping.OriginInfo.Geographic {
		if geographic, ok := geographicMap["gnd"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["idref"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["sbt"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["rero"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["unknown"]; ok {
			result = appendGeographic(geographic, result)
		}

		for key, geographic := range geographicMap {
			if geographic.Identifier != "" {
				authorities[key] = geographic.Identifier
			}
			for _, id := range geographic.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func (m *MappingRDV) GetSubjectGeographic() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Subject == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.Geographic == nil {
		return "", nil, false
	}

	key = "subjectGeographic"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, geographicMap := range m.Mapping.Subject.Geographic {
		if geographic, ok := geographicMap["gnd"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["idref"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["sbt"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["rero"]; ok {
			result = appendGeographic(geographic, result)
		} else if geographic, ok := geographicMap["unknown"]; ok {
			result = appendGeographic(geographic, result)
		}

		for key, geographic := range geographicMap {
			if geographic.Identifier != "" {
				authorities[key] = geographic.Identifier
			}
			for _, id := range geographic.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return
}

func appendGeographic(geographic schema.Geographic, result []Element) []Element {
	if geographic.NamePart != "" {
		e := Element{
			Text: geographic.NamePart,
			Link: "facet:" + geographic.NamePart,
		}
		if len(geographic.Role) > 0 {
			e.Extended = map[string]json.RawMessage{}
			roleBytes, _ := json.Marshal(geographic.Role)
			e.Extended["roles"] = roleBytes
		}
		result = append(result, e)
	}
	return result
}

func (m *MappingRDV) GetNameRelatedWork() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.RelatedItem == nil {
		return "", nil, false
	}
	if m.Mapping.RelatedItem.Work == nil {
		return "", nil, false
	}

	key = "nameRelatedWork"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, titleMap := range m.Mapping.RelatedItem.Work {
		if title, ok := titleMap["gnd"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["idref"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["sbt"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["rero"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["unknown"]; ok {
			result = appendTitle(title, result)
		}

		for key, geographic := range titleMap {
			if geographic.Identifier != "" {
				authorities[key] = geographic.Identifier
			}
			for _, id := range geographic.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return

}

func (m *MappingRDV) GetSubjectTitleInfo() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Subject == nil {
		return "", nil, false
	}
	if m.Mapping.Subject.TitleInfo == nil {
		return "", nil, false
	}

	key = "subjectTitleInfo"
	ok = true
	result = []Element{}

	authorities := make(map[string]string)
	for _, titleMap := range m.Mapping.Subject.TitleInfo {
		if title, ok := titleMap["gnd"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["idref"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["sbt"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["rero"]; ok {
			result = appendTitle(title, result)
		} else if title, ok := titleMap["unknown"]; ok {
			result = appendTitle(title, result)
		}

		for key, geographic := range titleMap {
			if geographic.Identifier != "" {
				authorities[key] = geographic.Identifier
			}
			for _, id := range geographic.OtherIdentifier {
				if id != "" {
					splitStr := strings.Split(id, ")")
					prefix := strings.Trim(splitStr[0], "(")
					authorities[prefix] = id
				}
			}
		}

		authorityElements := appendAuthorityElements(authorities)

		if len(authorityElements) > 0 && len(result) > 0 {
			authBytes, _ := json.Marshal(authorityElements)
			if result[len(result)-1].Extended != nil {
				result[len(result)-1].Extended["authorities"] = authBytes
			} else {
				authBytesTmp := map[string]json.RawMessage{}
				authBytesTmp["authorities"] = authBytes
				result[len(result)-1].Extended = authBytesTmp
			}

		}
	}

	return

}

func appendTitle(work schema.Work, result []Element) []Element {
	if work.Title != "" {
		e := Element{
			Text: work.Title,
			Link: "facet:" + work.Title,
		}
		if work.Name != "" {
			appendText(&e, work.Name, " / ")
		}
		result = append(result, e)
	}
	return result
}

/*
todo: change data in index: Index generic values in German, adapt this function to get the code and only one per document
usually used for selecting icon and configuration in frontend
*/
func (m *MappingRDV) GetResourceType() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return "", nil, false
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
		return "", nil, false
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
		return "", nil, false
	}
	if m.Facets == nil {
		return "", nil, false
	}
	if m.Mapping == nil {
		return
	}

	result = []Element{}
	key = "resourceTypeView"
	ok = true

	for _, v := range m.Flags {
		if v == "portraets" {
			for _, sf := range m.Facets.Strings {
				if sf.Name != "digitized" {
					continue
				}
				for _, s := range sf.String {
					if s == "yes" {
						e := Element{
							Text: "portraetsDigitized",
						}
						result = append(result, e)
					} else {
						e := Element{
							Text: "portraetsPhysical",
						}
						result = append(result, e)
					}
					break
				}
			}
		} else if m.Mapping.Files != nil {
			e := Element{
				Text: "mediaFiles",
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
		return "", nil, false
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
		return "", nil, false
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
		return "", nil, false
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
		return "", nil, false
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
		return "", nil, false
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
		return "", nil, false
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

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetFacetPortraetsPictured() (key string, result []Element, ok bool) {
	if m.Facets == nil {
		return "", nil, false
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
		return "", nil, false
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

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

func (m *MappingRDV) GetThumbnail() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
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
		typeBytes, _ := json.Marshal("mediaserver")
		e.Extended["type"] = typeBytes
		result = append(result, e)
	}

	if len(result) == 0 {
		key, result, ok = m.GetThumbnailFromData()
	}
	return

}

func (m *MappingRDV) GetThumbnailFromData() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Location == nil {
		return "", nil, false
	}
	if len(m.Mapping.Location.Digital) == 0 {
		return "", nil, false
	}

	key = "thumbnail"
	ok = true
	result = []Element{}

	for _, v := range m.Mapping.Location.Digital {
		if v == nil {
			continue
		}
		portraitUrlPattern := regexp.MustCompile(`.*digi/a100/portraet/`)
		portraitIdPattern := regexp.MustCompile(`^.*/([^./]+)\.[^/]+$`)

		if portraitUrlPattern.MatchString(v.Url) {
			matches := portraitIdPattern.FindStringSubmatch(v.Url)
			if len(matches) > 1 {
				portraitId := matches[1]
				portraitImageUrl := "https://ub-sipi.ub.unibas.ch/portraets/" + portraitId + ".jpx/full/{$widthInPx},/0/default.jpg"
				e := Element{
					Link:     portraitImageUrl,
					Extended: map[string]json.RawMessage{},
				}
				typeBytes, _ := json.Marshal("iiif")
				e.Extended["type"] = typeBytes
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
		return "", nil, false
	}
	if m.Mapping.Files == nil {
		return "", nil, false
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

func (m *MappingRDV) GetMedia() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if m.Mapping.Files == nil {
		return "", nil, false
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

			/* todo: change criteria once we have other data types
			- set presentationType depending on type -> webarchive and epub as iframe with action in URL
			*/
			if f.Type == "image" {
				licenseUrl := ""
				if licenseUrls, exists := licenseUrls[f.License_abbr]; exists {
					licenseUrl = licenseUrls
				}

				// urlBytes, _ := json.Marshal(strings.Replace(f.Uri, "mediaserver:", "https://mediaservermain.ub-dlza-test.k8s-001.unibas.ch/iiif/3/", 1))
				urlBytes, _ := json.Marshal(f.Uri)
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

// AddTestMedia todo: remove, creates test data
func (m *MappingRDV) AddTestMedia() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return "", nil, false
	}

	key = "media"
	ok = true
	result = []Element{}
	contentArray := []map[string]json.RawMessage{}

	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}
		/* Test Beispiel 1 Audio */
		if ok, _ := regexp.MatchString("^9972650989305504$", v); ok {

			titleBytes, _ := json.Marshal("Deep")
			fileNameBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			arkQualifierBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			urlBytes, _ := json.Marshal("mediaserver:test/deep")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/deep$$wave")
			thumbnailWidthBytes, _ := json.Marshal(1280)
			thumbnailHeightBytes, _ := json.Marshal(256)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/deep")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("Audio, 6.3 MB")
			/*dateBytes, _ := json.Marshal()*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("audio")
			typeBytes, _ := json.Marshal("audio")
			mimetypeBytes, _ := json.Marshal("audio/mp3")
			pronomBytes, _ := json.Marshal("fmt/134")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/134")
			durationBytes, _ := json.Marshal(268)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				/*"date": "",*/
				"duration": durationBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 1 Video */
		if ok, _ := regexp.MatchString("^9967555870105504$", v); ok {

			titleBytes, _ := json.Marshal("Sprite Fright")
			fileNameBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			arkQualifierBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			urlBytes, _ := json.Marshal("mediaserver:test/spritefright")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/spritefright$$shot$$13")
			thumbnailWidthBytes, _ := json.Marshal(1920)
			thumbnailHeightBytes, _ := json.Marshal(804)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/spritefright")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("Video, 188.3 MB")
			/*dateBytes, _ := json.Marshal("")*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("video")
			typeBytes, _ := json.Marshal("video")
			mimetypeBytes, _ := json.Marshal("video/webm")
			pronomBytes, _ := json.Marshal("fmt/573")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/573")
			widthBytes, _ := json.Marshal(1920)
			heightBytes, _ := json.Marshal(804)
			durationBytes, _ := json.Marshal(629)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				/*"date": "",*/
				"width":    widthBytes,
				"height":   heightBytes,
				"duration": durationBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 1 Image */
		if ok, _ := regexp.MatchString("^9936873350105504$", v); ok {

			titleBytes, _ := json.Marshal("Sharpest ever view of the Andromeda Galaxy")
			fileNameBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			arkQualifierBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			urlBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailWidthBytes, _ := json.Marshal(10000)
			thumbnailHeightBytes, _ := json.Marshal(3197)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("Bild TIFF, 114")
			dateBytes, _ := json.Marshal("05.01.2015")
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("image")
			typeBytes, _ := json.Marshal("image")
			mimetypeBytes, _ := json.Marshal("image/tiff")
			pronomBytes, _ := json.Marshal("fmt/353")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/353")
			widthBytes, _ := json.Marshal(10000)
			heightBytes, _ := json.Marshal(3197)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				"date":             dateBytes,
				"width":            widthBytes,
				"height":           heightBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 1 PDF */
		if ok, _ := regexp.MatchString("^9936340160105504$", v); ok {

			titleBytes, _ := json.Marshal("Stellungnahme SLSP AG Öffentlichkeitsgesetz")
			fileNameBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			arkQualifierBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			urlBytes, _ := json.Marshal("mediaserver:test/slsp")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/slsp$$cover")
			thumbnailWidthBytes, _ := json.Marshal(1191)
			thumbnailHeightBytes, _ := json.Marshal(1684)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/slsp")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("PDF, 279 KB")
			/*dateBytes, _ := json.Marshal("25.01.2021")*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("pdf")
			typeBytes, _ := json.Marshal("text")
			mimetypeBytes, _ := json.Marshal("application/pdf")
			pronomBytes, _ := json.Marshal("fmt/276")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/276")
			widthBytes, _ := json.Marshal(595)
			heightBytes, _ := json.Marshal(842)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				/*"date": "",*/
				"width":  widthBytes,
				"height": heightBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 1 csv */
		if ok, _ := regexp.MatchString("^9966685930105504$", v); ok {

			titleBytes, _ := json.Marshal("Metadaten")
			fileNameBytes, _ := json.Marshal("media_file_links_000001.csv")
			arkQualifierBytes, _ := json.Marshal("media_file_links_000001.csv")
			urlBytes, _ := json.Marshal("mediaserver:test/media_file_links_000001.csv")
			/*thumbnailBytes, _ := json.Marshal("")*/
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/media_file_links_000001.csv")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("CSV, 398 KB")
			dateBytes, _ := json.Marshal("22.11.2024")
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("csv")
			typeBytes, _ := json.Marshal("text")
			mimetypeBytes, _ := json.Marshal("text/csv")
			pronomBytes, _ := json.Marshal("x-fmt/18")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/x-fmt/18")

			contentDetails := map[string]json.RawMessage{
				"title":        titleBytes,
				"fileName":     fileNameBytes,
				"arkQualifier": arkQualifierBytes,
				"url":          urlBytes,
				/*"thumbnail":        thumbnailBytes,*/
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				"date":             dateBytes,
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 1 json */
		if ok, _ := regexp.MatchString("^9964808620105504$", v); ok {

			titleBytes, _ := json.Marshal("JSON Metadaten")
			fileNameBytes, _ := json.Marshal("metadata_000001.jsonl")
			arkQualifierBytes, _ := json.Marshal("metadata_000001.jsonl")
			urlBytes, _ := json.Marshal("mediaserver:test/metadata_000001.jsonl")
			/*thumbnailBytes, _ := json.Marshal("")*/
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/metadata_000001.jsonl")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("PDF, 279 KB")
			dateBytes, _ := json.Marshal("22.11.2024")
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("none")
			typeBytes, _ := json.Marshal("text")
			mimetypeBytes, _ := json.Marshal("text/plain")
			pronomBytes, _ := json.Marshal("x-fmt/111")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/x-fmt/111")

			contentDetails := map[string]json.RawMessage{
				"title":        titleBytes,
				"fileName":     fileNameBytes,
				"arkQualifier": arkQualifierBytes,
				"url":          urlBytes,
				/*"thumbnail":        thumbnailBytes,*/
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				"date":             dateBytes,
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 1 epub -> iframe */
		if ok, _ := regexp.MatchString("^9959975600105504$", v); ok {

			titleBytes, _ := json.Marshal("EPUB 3.0 Specification")
			fileNameBytes, _ := json.Marshal("epub30-spec.epub")
			arkQualifierBytes, _ := json.Marshal("epub30-spec.epub")
			urlBytes, _ := json.Marshal("mediaserver:test/epub30-spec/foliateviewer") // mit action, da iframe
			/*thumbnailBytes, _ := json.Marshal("") */
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/epub30-spec")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("EPUB, 222 KB")
			/*dateBytes, _ := json.Marshal("")*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("iframe")
			typeBytes, _ := json.Marshal("text")
			mimetypeBytes, _ := json.Marshal("application/epub+zip")
			pronomBytes, _ := json.Marshal("fmt/483")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/483")

			contentDetails := map[string]json.RawMessage{
				"title":        titleBytes,
				"fileName":     fileNameBytes,
				"arkQualifier": arkQualifierBytes,
				"url":          urlBytes,
				/*"thumbnail":        thumbnailBytes,*/
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				/*"date": "",*/
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 1 wacz -> iframe */
		if ok, _ := regexp.MatchString("^99988730105504$", v); ok {

			titleBytes, _ := json.Marshal("Website UB Basel")
			fileNameBytes, _ := json.Marshal("ub-basel.wacz")
			arkQualifierBytes, _ := json.Marshal("ub-basel.wacz")
			urlBytes, _ := json.Marshal("mediaserver:test/ub-baselweb/replaywebviewer") // mit action, da iframe
			//thumbnailBytes, _ := json.Marshal("")
			//downloadUrlBytes, _ := json.Marshal("mediaserver:test/ub-baselweb/item")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("Web archive")
			/*dateBytes, _ := json.Marshal("")*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("iframe")
			typeBytes, _ := json.Marshal("archive")
			mimetypeBytes, _ := json.Marshal("application/zip")
			pronomBytes, _ := json.Marshal("fmt/1840")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/1840")

			contentDetails := map[string]json.RawMessage{
				"title":        titleBytes,
				"fileName":     fileNameBytes,
				"arkQualifier": arkQualifierBytes,
				"url":          urlBytes,
				//"thumbnail":        thumbnailBytes,
				//"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
			}
			contentArray = append(contentArray, contentDetails)
		}
		/* Test Beispiel 4  audio */
		if ok, _ := regexp.MatchString("^9972789864805504$", v); ok {
			titleBytes, _ := json.Marshal("Deep")
			fileNameBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			arkQualifierBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			urlBytes, _ := json.Marshal("mediaserver:test/deep")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/deep$$wave")
			thumbnailWidthBytes, _ := json.Marshal(1280)
			thumbnailHeightBytes, _ := json.Marshal(256)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/deep")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("Audio, 6.3 MB")
			/*dateBytes, _ := json.Marshal()*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("audio")
			typeBytes, _ := json.Marshal("audio")
			mimetypeBytes, _ := json.Marshal("audio/mp3")
			pronomBytes, _ := json.Marshal("fmt/134")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/134")
			durationBytes, _ := json.Marshal(268)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				/*"date": "",*/
				"duration": durationBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails, contentDetails, contentDetails, contentDetails)
		}
		/* Test Beispiel 3  video */
		if ok, _ := regexp.MatchString("^9972529169305504$", v); ok {

			titleBytes, _ := json.Marshal("Sprite Fright")
			fileNameBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			arkQualifierBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			urlBytes, _ := json.Marshal("mediaserver:test/spritefright")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/spritefright$$shot$$13")
			thumbnailWidthBytes, _ := json.Marshal(1920)
			thumbnailHeightBytes, _ := json.Marshal(804)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/spritefright")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("Video, 188.3 MB")
			/*dateBytes, _ := json.Marshal("")*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("video")
			typeBytes, _ := json.Marshal("video")
			mimetypeBytes, _ := json.Marshal("video/webm")
			pronomBytes, _ := json.Marshal("fmt/573")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/573")
			widthBytes, _ := json.Marshal(1920)
			heightBytes, _ := json.Marshal(804)
			durationBytes, _ := json.Marshal(629)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				/*"date": "",*/
				"width":    widthBytes,
				"height":   heightBytes,
				"duration": durationBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails, contentDetails, contentDetails)
		}
		/* Test Beispiel 12  zoom image */
		if ok, _ := regexp.MatchString("^9972536425205504$", v); ok {

			titleBytes, _ := json.Marshal("Sharpest ever view of the Andromeda Galaxy")
			fileNameBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			arkQualifierBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			urlBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailWidthBytes, _ := json.Marshal(10000)
			thumbnailHeightBytes, _ := json.Marshal(3197)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("Bild TIFF, 114")
			dateBytes, _ := json.Marshal("05.01.2015")
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("image")
			typeBytes, _ := json.Marshal("image")
			mimetypeBytes, _ := json.Marshal("image/tiff")
			pronomBytes, _ := json.Marshal("fmt/353")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/353")
			widthBytes, _ := json.Marshal(10000)
			heightBytes, _ := json.Marshal(3197)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				"date":             dateBytes,
				"width":            widthBytes,
				"height":           heightBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails)
		}
		/* Test Beispiel 8  pdf */
		if ok, _ := regexp.MatchString("^9972768968505504$", v); ok {

			titleBytes, _ := json.Marshal("Stellungnahme SLSP AG Öffentlichkeitsgesetz")
			fileNameBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			arkQualifierBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			urlBytes, _ := json.Marshal("mediaserver:test/slsp")
			thumbnailBytes, _ := json.Marshal("mediaserver:test/slsp$$cover")
			thumbnailWidthBytes, _ := json.Marshal(1191)
			thumbnailHeightBytes, _ := json.Marshal(1684)
			downloadUrlBytes, _ := json.Marshal("mediaserver:test/slsp")
			aclBytes, _ := json.Marshal("global/guest")
			formatBytes, _ := json.Marshal("PDF, 279 KB")
			/*dateBytes, _ := json.Marshal("25.01.2021")*/
			licenseBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeBytes, _ := json.Marshal("pdf")
			typeBytes, _ := json.Marshal("text")
			mimetypeBytes, _ := json.Marshal("application/pdf")
			pronomBytes, _ := json.Marshal("fmt/276")
			pronomUrlBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/276")
			widthBytes, _ := json.Marshal(595)
			heightBytes, _ := json.Marshal(842)

			contentDetails := map[string]json.RawMessage{
				"title":            titleBytes,
				"fileName":         fileNameBytes,
				"arkQualifier":     arkQualifierBytes,
				"url":              urlBytes,
				"thumbnail":        thumbnailBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlBytes,
				"acl":              aclBytes,
				"license":          licenseBytes,
				"licenseUrl":       licenseUrlBytes,
				"presentationType": presentationTypeBytes,
				"type":             typeBytes,
				"mimetype":         mimetypeBytes,
				"pronom":           pronomBytes,
				"pronomUrl":        pronomUrlBytes,
				"format":           formatBytes,
				/*"date": "",*/
				"width":  widthBytes,
				"height": heightBytes,
				/*"orientation": "",*/
			}
			contentArray = append(contentArray, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails, contentDetails)
		}
		/* Test Beispiel mix, flat */
		if ok, _ := regexp.MatchString("^9970786120105504$", v); ok {

			titleAudioBytes, _ := json.Marshal("Deep")
			fileNameAudioBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			arkQualifierAudioBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			urlAudioBytes, _ := json.Marshal("mediaserver:test/deep")
			thumbnailAudioBytes, _ := json.Marshal("mediaserver:test/deep$$wave")
			thumbnailAudioWidthBytes, _ := json.Marshal(1280)
			thumbnailAudioHeightBytes, _ := json.Marshal(256)
			downloadUrlAudioBytes, _ := json.Marshal("mediaserver:test/deep")
			aclAudioBytes, _ := json.Marshal("global/guest")
			formatAudioBytes, _ := json.Marshal("Audio, 6.3 MB")
			/*dateAudioBytes, _ := json.Marshal()*/
			licenseAudioBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlAudioBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeAudioBytes, _ := json.Marshal("audio")
			typeAudioBytes, _ := json.Marshal("audio")
			mimetypeAudioBytes, _ := json.Marshal("audio/mp3")
			pronomAudioBytes, _ := json.Marshal("fmt/134")
			pronomUrlAudioBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/134")
			durationAudioBytes, _ := json.Marshal(268)

			contentDetailsAudio := map[string]json.RawMessage{
				"title":            titleAudioBytes,
				"fileName":         fileNameAudioBytes,
				"arkQualifier":     arkQualifierAudioBytes,
				"url":              urlAudioBytes,
				"thumbnail":        thumbnailAudioBytes,
				"thumbnailWidth":   thumbnailAudioWidthBytes,
				"thumbnailHeight":  thumbnailAudioHeightBytes,
				"downloadUrl":      downloadUrlAudioBytes,
				"acl":              aclAudioBytes,
				"license":          licenseAudioBytes,
				"licenseUrl":       licenseUrlAudioBytes,
				"presentationType": presentationTypeAudioBytes,
				"type":             typeAudioBytes,
				"mimetype":         mimetypeAudioBytes,
				"pronom":           pronomAudioBytes,
				"pronomUrl":        pronomUrlAudioBytes,
				"format":           formatAudioBytes,
				/*"date": "",*/
				"duration": durationAudioBytes,
				/*"orientation": "",*/
			}

			titleVideoBytes, _ := json.Marshal("Sprite Fright")
			fileNameVideoBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			arkQualifierVideoBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			urlVideoBytes, _ := json.Marshal("mediaserver:test/spritefright")
			thumbnailVideoBytes, _ := json.Marshal("mediaserver:test/spritefright$$shot$$13")
			thumbnailVideoWidthBytes, _ := json.Marshal(1920)
			thumbnailVideoHeightBytes, _ := json.Marshal(804)
			downloadUrlVideoBytes, _ := json.Marshal("mediaserver:test/spritefright")
			aclVideoBytes, _ := json.Marshal("global/guest")
			formatVideoBytes, _ := json.Marshal("Video, 188.3 MB")
			/*dateVideoBytes, _ := json.Marshal("")*/
			licenseVideoBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlVideoBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeVideoBytes, _ := json.Marshal("video")
			typeVideoBytes, _ := json.Marshal("video")
			mimetypeVideoBytes, _ := json.Marshal("video/webm")
			pronomVideoBytes, _ := json.Marshal("fmt/573")
			pronomUrlVideoBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/573")
			widthVideoBytes, _ := json.Marshal(1920)
			heightVideoBytes, _ := json.Marshal(804)
			durationVideoBytes, _ := json.Marshal(629)

			contentDetailsVideo := map[string]json.RawMessage{
				"title":            titleVideoBytes,
				"fileName":         fileNameVideoBytes,
				"arkQualifier":     arkQualifierVideoBytes,
				"url":              urlVideoBytes,
				"thumbnail":        thumbnailVideoBytes,
				"thumbnailWidth":   thumbnailVideoWidthBytes,
				"thumbnailHeight":  thumbnailVideoHeightBytes,
				"downloadUrl":      downloadUrlVideoBytes,
				"acl":              aclVideoBytes,
				"license":          licenseVideoBytes,
				"licenseUrl":       licenseUrlVideoBytes,
				"presentationType": presentationTypeVideoBytes,
				"type":             typeVideoBytes,
				"mimetype":         mimetypeVideoBytes,
				"pronom":           pronomVideoBytes,
				"pronomUrl":        pronomUrlVideoBytes,
				"format":           formatVideoBytes,
				/*"date": "",*/
				"width":    widthVideoBytes,
				"height":   heightVideoBytes,
				"duration": durationVideoBytes,
				/*"orientation": "",*/
			}

			titleImageBytes, _ := json.Marshal("Sharpest ever view of the Andromeda Galaxy")
			fileNameImageBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			arkQualifierImageBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			urlImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailImageWidthBytes, _ := json.Marshal(10000)
			thumbnailImageHeightBytes, _ := json.Marshal(3197)
			downloadUrlImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			aclImageBytes, _ := json.Marshal("global/guest")
			formatImageBytes, _ := json.Marshal("Bild TIFF, 114")
			dateImageBytes, _ := json.Marshal("05.01.2015")
			licenseImageBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlImageBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeImageBytes, _ := json.Marshal("image")
			typeImageBytes, _ := json.Marshal("image")
			mimetypeImageBytes, _ := json.Marshal("image/tiff")
			pronomImageBytes, _ := json.Marshal("fmt/353")
			pronomUrlImageBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/353")
			widthImageBytes, _ := json.Marshal(10000)
			heightImageBytes, _ := json.Marshal(3197)

			contentDetailsImage := map[string]json.RawMessage{
				"title":            titleImageBytes,
				"fileName":         fileNameImageBytes,
				"arkQualifier":     arkQualifierImageBytes,
				"url":              urlImageBytes,
				"thumbnail":        thumbnailImageBytes,
				"thumbnailWidth":   thumbnailImageWidthBytes,
				"thumbnailHeight":  thumbnailImageHeightBytes,
				"downloadUrl":      downloadUrlImageBytes,
				"acl":              aclImageBytes,
				"license":          licenseImageBytes,
				"licenseUrl":       licenseUrlImageBytes,
				"presentationType": presentationTypeImageBytes,
				"type":             typeImageBytes,
				"mimetype":         mimetypeImageBytes,
				"pronom":           pronomImageBytes,
				"pronomUrl":        pronomUrlImageBytes,
				"format":           formatImageBytes,
				"date":             dateImageBytes,
				"width":            widthImageBytes,
				"height":           heightImageBytes,
				/*"orientation": "",*/
			}

			titlePdfBytes, _ := json.Marshal("Stellungnahme SLSP AG Öffentlichkeitsgesetz")
			fileNamePdfBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			arkQualifierPdfBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			urlPdfBytes, _ := json.Marshal("mediaserver:test/slsp")
			thumbnailPdfBytes, _ := json.Marshal("mediaserver:test/slsp$$cover")
			thumbnailWidthBytes, _ := json.Marshal(1191)
			thumbnailHeightBytes, _ := json.Marshal(1684)
			downloadUrlPdfBytes, _ := json.Marshal("mediaserver:test/slsp")
			aclPdfBytes, _ := json.Marshal("global/guest")
			formatPdfBytes, _ := json.Marshal("PDF, 279 KB")
			/*datePdfBytes, _ := json.Marshal("25.01.2021")*/
			licensePdfBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlPdfBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypePdfBytes, _ := json.Marshal("pdf")
			typePdfBytes, _ := json.Marshal("text")
			mimetypePdfBytes, _ := json.Marshal("application/pdf")
			pronomPdfBytes, _ := json.Marshal("fmt/276")
			pronomUrlPdfBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/276")
			widthPdfBytes, _ := json.Marshal(595)
			heightPdfBytes, _ := json.Marshal(842)

			contentDetailsPdf := map[string]json.RawMessage{
				"title":            titlePdfBytes,
				"fileName":         fileNamePdfBytes,
				"arkQualifier":     arkQualifierPdfBytes,
				"url":              urlPdfBytes,
				"thumbnail":        thumbnailPdfBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlPdfBytes,
				"acl":              aclPdfBytes,
				"license":          licensePdfBytes,
				"licenseUrl":       licenseUrlPdfBytes,
				"presentationType": presentationTypePdfBytes,
				"type":             typePdfBytes,
				"mimetype":         mimetypePdfBytes,
				"pronom":           pronomPdfBytes,
				"pronomUrl":        pronomUrlPdfBytes,
				"format":           formatPdfBytes,
				/*"date": "",*/
				"width":  widthPdfBytes,
				"height": heightPdfBytes,
				/*"orientation": "",*/
			}

			contentArray = append(contentArray, contentDetailsAudio, contentDetailsVideo, contentDetailsImage, contentDetailsVideo, contentDetailsVideo, contentDetailsAudio, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsPdf, contentDetailsPdf)
		}
		/* Test Beispiel mix, structured */
		if ok, _ := regexp.MatchString("^9972608858805504$", v); ok {

			contentArrayFolder1 := []map[string]json.RawMessage{}
			contentArrayFolder2 := []map[string]json.RawMessage{}

			titleAudioBytes, _ := json.Marshal("Deep")
			fileNameAudioBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			arkQualifierAudioBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			urlAudioBytes, _ := json.Marshal("mediaserver:test/deep")
			thumbnailAudioBytes, _ := json.Marshal("mediaserver:test/deep$$wave")
			thumbnailAudioWidthBytes, _ := json.Marshal(1280)
			thumbnailAudioHeightBytes, _ := json.Marshal(256)
			downloadUrlAudioBytes, _ := json.Marshal("mediaserver:test/deep")
			aclAudioBytes, _ := json.Marshal("global/guest")
			formatAudioBytes, _ := json.Marshal("Audio, 6.3 MB")
			/*dateAudioBytes, _ := json.Marshal()*/
			licenseAudioBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlAudioBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeAudioBytes, _ := json.Marshal("audio")
			typeAudioBytes, _ := json.Marshal("audio")
			mimetypeAudioBytes, _ := json.Marshal("audio/mp3")
			pronomAudioBytes, _ := json.Marshal("fmt/134")
			pronomUrlAudioBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/134")
			durationAudioBytes, _ := json.Marshal(268)

			contentDetailsAudio := map[string]json.RawMessage{
				"title":            titleAudioBytes,
				"fileName":         fileNameAudioBytes,
				"arkQualifier":     arkQualifierAudioBytes,
				"url":              urlAudioBytes,
				"thumbnail":        thumbnailAudioBytes,
				"thumbnailWidth":   thumbnailAudioWidthBytes,
				"thumbnailHeight":  thumbnailAudioHeightBytes,
				"downloadUrl":      downloadUrlAudioBytes,
				"acl":              aclAudioBytes,
				"license":          licenseAudioBytes,
				"licenseUrl":       licenseUrlAudioBytes,
				"presentationType": presentationTypeAudioBytes,
				"type":             typeAudioBytes,
				"mimetype":         mimetypeAudioBytes,
				"pronom":           pronomAudioBytes,
				"pronomUrl":        pronomUrlAudioBytes,
				"format":           formatAudioBytes,
				/*"date": "",*/
				"duration": durationAudioBytes,
				/*"orientation": "",*/
			}

			titleVideoBytes, _ := json.Marshal("Sprite Fright")
			fileNameVideoBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			arkQualifierVideoBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			urlVideoBytes, _ := json.Marshal("mediaserver:test/spritefright")
			thumbnailVideoBytes, _ := json.Marshal("mediaserver:test/spritefright$$shot$$13")
			thumbnailVideoWidthBytes, _ := json.Marshal(1920)
			thumbnailVideoHeightBytes, _ := json.Marshal(804)
			downloadUrlVideoBytes, _ := json.Marshal("mediaserver:test/spritefright")
			aclVideoBytes, _ := json.Marshal("unibas/testuser")
			formatVideoBytes, _ := json.Marshal("Video, 188.3 MB")
			/*dateVideoBytes, _ := json.Marshal("")*/
			licenseVideoBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlVideoBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeVideoBytes, _ := json.Marshal("video")
			typeVideoBytes, _ := json.Marshal("video")
			mimetypeVideoBytes, _ := json.Marshal("video/webm")
			pronomVideoBytes, _ := json.Marshal("fmt/573")
			pronomUrlVideoBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/573")
			widthVideoBytes, _ := json.Marshal(1920)
			heightVideoBytes, _ := json.Marshal(804)
			durationVideoBytes, _ := json.Marshal(629)

			contentDetailsVideo := map[string]json.RawMessage{
				"title":            titleVideoBytes,
				"fileName":         fileNameVideoBytes,
				"arkQualifier":     arkQualifierVideoBytes,
				"url":              urlVideoBytes,
				"thumbnail":        thumbnailVideoBytes,
				"thumbnailWidth":   thumbnailVideoWidthBytes,
				"thumbnailHeight":  thumbnailVideoHeightBytes,
				"downloadUrl":      downloadUrlVideoBytes,
				"acl":              aclVideoBytes,
				"license":          licenseVideoBytes,
				"licenseUrl":       licenseUrlVideoBytes,
				"presentationType": presentationTypeVideoBytes,
				"type":             typeVideoBytes,
				"mimetype":         mimetypeVideoBytes,
				"pronom":           pronomVideoBytes,
				"pronomUrl":        pronomUrlVideoBytes,
				"format":           formatVideoBytes,
				/*"date": "",*/
				"width":    widthVideoBytes,
				"height":   heightVideoBytes,
				"duration": durationVideoBytes,
				/*"orientation": "",*/
			}

			titleImageBytes, _ := json.Marshal("Sharpest ever view of the Andromeda Galaxy")
			fileNameImageBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			arkQualifierImageBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			urlImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailImageWidthBytes, _ := json.Marshal(10000)
			thumbnailImageHeightBytes, _ := json.Marshal(3197)
			downloadUrlImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			aclImageBytes, _ := json.Marshal("global/guest")
			formatImageBytes, _ := json.Marshal("Bild TIFF, 114")
			dateImageBytes, _ := json.Marshal("05.01.2015")
			licenseImageBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlImageBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeImageBytes, _ := json.Marshal("image")
			typeImageBytes, _ := json.Marshal("image")
			mimetypeImageBytes, _ := json.Marshal("image/tiff")
			pronomImageBytes, _ := json.Marshal("fmt/353")
			pronomUrlImageBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/353")
			widthImageBytes, _ := json.Marshal(10000)
			heightImageBytes, _ := json.Marshal(3197)

			contentDetailsImage := map[string]json.RawMessage{
				"title":            titleImageBytes,
				"fileName":         fileNameImageBytes,
				"arkQualifier":     arkQualifierImageBytes,
				"url":              urlImageBytes,
				"thumbnail":        thumbnailImageBytes,
				"thumbnailWidth":   thumbnailImageWidthBytes,
				"thumbnailHeight":  thumbnailImageHeightBytes,
				"downloadUrl":      downloadUrlImageBytes,
				"acl":              aclImageBytes,
				"license":          licenseImageBytes,
				"licenseUrl":       licenseUrlImageBytes,
				"presentationType": presentationTypeImageBytes,
				"type":             typeImageBytes,
				"mimetype":         mimetypeImageBytes,
				"pronom":           pronomImageBytes,
				"pronomUrl":        pronomUrlImageBytes,
				"format":           formatImageBytes,
				"date":             dateImageBytes,
				"width":            widthImageBytes,
				"height":           heightImageBytes,
				/*"orientation": "",*/
			}

			titlePdfBytes, _ := json.Marshal("Stellungnahme SLSP AG Öffentlichkeitsgesetz")
			fileNamePdfBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			arkQualifierPdfBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			urlPdfBytes, _ := json.Marshal("mediaserver:test/slsp")
			thumbnailPdfBytes, _ := json.Marshal("mediaserver:test/slsp$$cover")
			thumbnailPdfWidthBytes, _ := json.Marshal(1191)
			thumbnailPdfHeightBytes, _ := json.Marshal(1684)
			downloadUrlPdfBytes, _ := json.Marshal("mediaserver:test/slsp")
			aclPdfBytes, _ := json.Marshal("global/guest")
			formatPdfBytes, _ := json.Marshal("PDF, 279 KB")
			/*datePdfBytes, _ := json.Marshal("25.01.2021")*/
			licensePdfBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlPdfBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypePdfBytes, _ := json.Marshal("pdf")
			typePdfBytes, _ := json.Marshal("text")
			mimetypePdfBytes, _ := json.Marshal("application/pdf")
			pronomPdfBytes, _ := json.Marshal("fmt/276")
			pronomUrlPdfBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/276")
			widthPdfBytes, _ := json.Marshal(595)
			heightPdfBytes, _ := json.Marshal(842)

			contentDetailsPdf := map[string]json.RawMessage{
				"title":            titlePdfBytes,
				"fileName":         fileNamePdfBytes,
				"arkQualifier":     arkQualifierPdfBytes,
				"url":              urlPdfBytes,
				"thumbnail":        thumbnailPdfBytes,
				"thumbnailWidth":   thumbnailPdfWidthBytes,
				"thumbnailHeight":  thumbnailPdfHeightBytes,
				"downloadUrl":      downloadUrlPdfBytes,
				"acl":              aclPdfBytes,
				"license":          licensePdfBytes,
				"licenseUrl":       licenseUrlPdfBytes,
				"presentationType": presentationTypePdfBytes,
				"type":             typePdfBytes,
				"mimetype":         mimetypePdfBytes,
				"pronom":           pronomPdfBytes,
				"pronomUrl":        pronomUrlPdfBytes,
				"format":           formatPdfBytes,
				/*"date": "",*/
				"width":  widthPdfBytes,
				"height": heightPdfBytes,
				/*"orientation": "",*/
			}

			titleCsvBytes, _ := json.Marshal("Metadaten")
			fileNameCsvBytes, _ := json.Marshal("media_file_links_000001.csv")
			arkQualifierCsvBytes, _ := json.Marshal("media_file_links_000001.csv")
			urlCsvBytes, _ := json.Marshal("mediaserver:test/media_file_links_000001.csv")
			/*thumbnailCsvBytes, _ := json.Marshal("")*/
			downloadUrlCsvBytes, _ := json.Marshal("mediaserver:test/media_file_links_000001.csv")
			aclCsvBytes, _ := json.Marshal("global/guest")
			formatCsvBytes, _ := json.Marshal("CSV, 398 KB")
			dateCsvBytes, _ := json.Marshal("22.11.2024")
			licenseCsvBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlCsvBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeCsvBytes, _ := json.Marshal("csv")
			typeCsvBytes, _ := json.Marshal("text")
			mimetypeCsvBytes, _ := json.Marshal("text/csv")
			pronomCsvBytes, _ := json.Marshal("x-fmt/18")
			pronomUrlCsvBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/x-fmt/18")

			contentDetailsCsv := map[string]json.RawMessage{
				"title":        titleCsvBytes,
				"fileName":     fileNameCsvBytes,
				"arkQualifier": arkQualifierCsvBytes,
				"url":          urlCsvBytes,
				/*"thumbnail":        thumbnailCsvBytes,*/
				"downloadUrl":      downloadUrlCsvBytes,
				"acl":              aclCsvBytes,
				"license":          licenseCsvBytes,
				"licenseUrl":       licenseUrlCsvBytes,
				"presentationType": presentationTypeCsvBytes,
				"type":             typeCsvBytes,
				"mimetype":         mimetypeCsvBytes,
				"pronom":           pronomCsvBytes,
				"pronomUrl":        pronomUrlCsvBytes,
				"format":           formatCsvBytes,
				"date":             dateCsvBytes,
			}

			folderName1Bytes, _ := json.Marshal("Session 1")
			contentArrayFolder1 = append(contentArrayFolder1, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsVideo, contentDetailsVideo, contentDetailsVideo)
			contentArrayFolder1Bytes, _ := json.Marshal(contentArrayFolder1)

			contentDetailsFolder1 := map[string]json.RawMessage{
				"folderName": folderName1Bytes,
				"content":    contentArrayFolder1Bytes,
			}

			folderName2Bytes, _ := json.Marshal("Data")
			contentArrayFolder2 = append(contentArrayFolder2, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsCsv, contentDetailsCsv, contentDetailsCsv, contentDetailsCsv, contentDetailsCsv, contentDetailsCsv, contentDetailsCsv, contentDetailsAudio)
			contentArrayFolder2Bytes, _ := json.Marshal(contentArrayFolder1)

			contentDetailsFolder2 := map[string]json.RawMessage{
				"folderName": folderName2Bytes,
				"content":    contentArrayFolder2Bytes,
			}

			contentArray = append(contentArray, contentDetailsFolder1, contentDetailsFolder2)
		}
		/* Test Beispiel mix, flat, viele */
		if ok, _ := regexp.MatchString("^9935139840105504$", v); ok {

			titleAudioBytes, _ := json.Marshal("Deep")
			fileNameAudioBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			arkQualifierAudioBytes, _ := json.Marshal("Alex-Productions_-_Deep_(Dark_Ambient_Background_music).oga.mp3")
			urlAudioBytes, _ := json.Marshal("mediaserver:test/deep")
			thumbnailAudioBytes, _ := json.Marshal("mediaserver:test/deep$$wave")
			thumbnailAudioWidthBytes, _ := json.Marshal(1280)
			thumbnailAudioHeightBytes, _ := json.Marshal(256)
			downloadUrlAudioBytes, _ := json.Marshal("mediaserver:test/deep")
			aclAudioBytes, _ := json.Marshal("global/guest")
			formatAudioBytes, _ := json.Marshal("Audio, 6.3 MB")
			/*dateAudioBytes, _ := json.Marshal()*/
			licenseAudioBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlAudioBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeAudioBytes, _ := json.Marshal("audio")
			typeAudioBytes, _ := json.Marshal("audio")
			mimetypeAudioBytes, _ := json.Marshal("audio/mp3")
			pronomAudioBytes, _ := json.Marshal("fmt/134")
			pronomUrlAudioBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/134")
			durationAudioBytes, _ := json.Marshal(268)

			contentDetailsAudio := map[string]json.RawMessage{
				"title":            titleAudioBytes,
				"fileName":         fileNameAudioBytes,
				"arkQualifier":     arkQualifierAudioBytes,
				"url":              urlAudioBytes,
				"thumbnail":        thumbnailAudioBytes,
				"thumbnailWidth":   thumbnailAudioWidthBytes,
				"thumbnailHeight":  thumbnailAudioHeightBytes,
				"downloadUrl":      downloadUrlAudioBytes,
				"acl":              aclAudioBytes,
				"license":          licenseAudioBytes,
				"licenseUrl":       licenseUrlAudioBytes,
				"presentationType": presentationTypeAudioBytes,
				"type":             typeAudioBytes,
				"mimetype":         mimetypeAudioBytes,
				"pronom":           pronomAudioBytes,
				"pronomUrl":        pronomUrlAudioBytes,
				"format":           formatAudioBytes,
				/*"date": "",*/
				"duration": durationAudioBytes,
				/*"orientation": "",*/
			}

			titleVideoBytes, _ := json.Marshal("Sprite Fright")
			fileNameVideoBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			arkQualifierVideoBytes, _ := json.Marshal("Sprite_Fright_-_Blender_Open_Movie-full_movie.webm.1080p.vp9.webm")
			urlVideoBytes, _ := json.Marshal("mediaserver:test/spritefright")
			thumbnailVideoBytes, _ := json.Marshal("mediaserver:test/spritefright$$shot$$13")
			thumbnailVideoWidthBytes, _ := json.Marshal(1920)
			thumbnailVideoHeightBytes, _ := json.Marshal(804)
			downloadUrlVideoBytes, _ := json.Marshal("mediaserver:test/spritefright")
			aclVideoBytes, _ := json.Marshal("global/guest")
			formatVideoBytes, _ := json.Marshal("Video, 188.3 MB")
			/*dateVideoBytes, _ := json.Marshal("")*/
			licenseVideoBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlVideoBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeVideoBytes, _ := json.Marshal("video")
			typeVideoBytes, _ := json.Marshal("video")
			mimetypeVideoBytes, _ := json.Marshal("video/webm")
			pronomVideoBytes, _ := json.Marshal("fmt/573")
			pronomUrlVideoBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/573")
			widthVideoBytes, _ := json.Marshal(1920)
			heightVideoBytes, _ := json.Marshal(804)
			durationVideoBytes, _ := json.Marshal(629)

			contentDetailsVideo := map[string]json.RawMessage{
				"title":            titleVideoBytes,
				"fileName":         fileNameVideoBytes,
				"arkQualifier":     arkQualifierVideoBytes,
				"url":              urlVideoBytes,
				"thumbnail":        thumbnailVideoBytes,
				"thumbnailWidth":   thumbnailVideoWidthBytes,
				"thumbnailHeight":  thumbnailVideoHeightBytes,
				"downloadUrl":      downloadUrlVideoBytes,
				"acl":              aclVideoBytes,
				"license":          licenseVideoBytes,
				"licenseUrl":       licenseUrlVideoBytes,
				"presentationType": presentationTypeVideoBytes,
				"type":             typeVideoBytes,
				"mimetype":         mimetypeVideoBytes,
				"pronom":           pronomVideoBytes,
				"pronomUrl":        pronomUrlVideoBytes,
				"format":           formatVideoBytes,
				/*"date": "",*/
				"width":    widthVideoBytes,
				"height":   heightVideoBytes,
				"duration": durationVideoBytes,
				/*"orientation": "",*/
			}

			titleImageBytes, _ := json.Marshal("Sharpest ever view of the Andromeda Galaxy")
			fileNameImageBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			arkQualifierImageBytes, _ := json.Marshal("heic1502a_10000x3197.tif")
			urlImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			thumbnailImageWidthBytes, _ := json.Marshal(10000)
			thumbnailImageHeightBytes, _ := json.Marshal(3197)
			downloadUrlImageBytes, _ := json.Marshal("mediaserver:test/andromeda10000")
			aclImageBytes, _ := json.Marshal("global/guest")
			formatImageBytes, _ := json.Marshal("Bild TIFF, 114")
			dateImageBytes, _ := json.Marshal("05.01.2015")
			licenseImageBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlImageBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypeImageBytes, _ := json.Marshal("image")
			typeImageBytes, _ := json.Marshal("image")
			mimetypeImageBytes, _ := json.Marshal("image/tiff")
			pronomImageBytes, _ := json.Marshal("fmt/353")
			pronomUrlImageBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/353")
			widthImageBytes, _ := json.Marshal(10000)
			heightImageBytes, _ := json.Marshal(3197)

			contentDetailsImage := map[string]json.RawMessage{
				"title":            titleImageBytes,
				"fileName":         fileNameImageBytes,
				"arkQualifier":     arkQualifierImageBytes,
				"url":              urlImageBytes,
				"thumbnail":        thumbnailImageBytes,
				"thumbnailWidth":   thumbnailImageWidthBytes,
				"thumbnailHeight":  thumbnailImageHeightBytes,
				"downloadUrl":      downloadUrlImageBytes,
				"acl":              aclImageBytes,
				"license":          licenseImageBytes,
				"licenseUrl":       licenseUrlImageBytes,
				"presentationType": presentationTypeImageBytes,
				"type":             typeImageBytes,
				"mimetype":         mimetypeImageBytes,
				"pronom":           pronomImageBytes,
				"pronomUrl":        pronomUrlImageBytes,
				"format":           formatImageBytes,
				"date":             dateImageBytes,
				"width":            widthImageBytes,
				"height":           heightImageBytes,
				/*"orientation": "",*/
			}

			titlePdfBytes, _ := json.Marshal("Stellungnahme SLSP AG Öffentlichkeitsgesetz")
			fileNamePdfBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			arkQualifierPdfBytes, _ := json.Marshal("stellungnahme_slsp_ag_oeffentlichkeitsgesetz (1).pdf")
			urlPdfBytes, _ := json.Marshal("mediaserver:test/slsp")
			thumbnailPdfBytes, _ := json.Marshal("mediaserver:test/slsp$$cover")
			thumbnailWidthBytes, _ := json.Marshal(1191)
			thumbnailHeightBytes, _ := json.Marshal(1684)
			downloadUrlPdfBytes, _ := json.Marshal("mediaserver:test/slsp")
			aclPdfBytes, _ := json.Marshal("global/guest")
			formatPdfBytes, _ := json.Marshal("PDF, 279 KB")
			/*datePdfBytes, _ := json.Marshal("25.01.2021")*/
			licensePdfBytes, _ := json.Marshal("PDM 1.0 Deed")
			licenseUrlPdfBytes, _ := json.Marshal("https://creativecommons.org/public-domain/pdm/")
			presentationTypePdfBytes, _ := json.Marshal("pdf")
			typePdfBytes, _ := json.Marshal("text")
			mimetypePdfBytes, _ := json.Marshal("application/pdf")
			pronomPdfBytes, _ := json.Marshal("fmt/276")
			pronomUrlPdfBytes, _ := json.Marshal("https://www.nationalarchives.gov.uk/pronom/fmt/276")
			widthPdfBytes, _ := json.Marshal(595)
			heightPdfBytes, _ := json.Marshal(842)

			contentDetailsPdf := map[string]json.RawMessage{
				"title":            titlePdfBytes,
				"fileName":         fileNamePdfBytes,
				"arkQualifier":     arkQualifierPdfBytes,
				"url":              urlPdfBytes,
				"thumbnail":        thumbnailPdfBytes,
				"thumbnailWidth":   thumbnailWidthBytes,
				"thumbnailHeight":  thumbnailHeightBytes,
				"downloadUrl":      downloadUrlPdfBytes,
				"acl":              aclPdfBytes,
				"license":          licensePdfBytes,
				"licenseUrl":       licenseUrlPdfBytes,
				"presentationType": presentationTypePdfBytes,
				"type":             typePdfBytes,
				"mimetype":         mimetypePdfBytes,
				"pronom":           pronomPdfBytes,
				"pronomUrl":        pronomUrlPdfBytes,
				"format":           formatPdfBytes,
				/*"date": "",*/
				"width":  widthPdfBytes,
				"height": heightPdfBytes,
				/*"orientation": "",*/
			}

			contentArray = append(contentArray, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsAudio, contentDetailsVideo, contentDetailsImage, contentDetailsVideo, contentDetailsVideo, contentDetailsAudio, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsImage, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf, contentDetailsPdf)
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

// AddTestThumbnail todo: remove, creates test data
func (m *MappingRDV) AddTestThumbnail() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return "", nil, false
	}

	key = "thumbnail"
	ok = true
	result = []Element{}

	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}
		if ok, _ := regexp.MatchString("^9972650989305504$", v); ok {
			e := Element{
				Link:     "mediaserver:test/deep$$wave",
				Extended: map[string]json.RawMessage{},
			}

			widthBytes, _ := json.Marshal(1280)
			e.Extended["width"] = widthBytes
			heightBytes, _ := json.Marshal(256)
			e.Extended["height"] = heightBytes
			typeBytes, _ := json.Marshal("mediaserver")
			e.Extended["type"] = typeBytes
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9967555870105504$", v); ok {
			e := Element{
				Link:     "mediaserver:test/spritefright$$shot$$13",
				Extended: map[string]json.RawMessage{},
			}

			widthBytes, _ := json.Marshal(1920)
			e.Extended["width"] = widthBytes
			heightBytes, _ := json.Marshal(804)
			e.Extended["height"] = heightBytes
			typeBytes, _ := json.Marshal("mediaserver")
			e.Extended["type"] = typeBytes
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9936873350105504$", v); ok {
			e := Element{
				Link:     "mediaserver:test/andromeda10000",
				Extended: map[string]json.RawMessage{},
			}

			widthBytes, _ := json.Marshal(10000)
			e.Extended["width"] = widthBytes
			heightBytes, _ := json.Marshal(3197)
			e.Extended["height"] = heightBytes
			typeBytes, _ := json.Marshal("mediaserver")
			e.Extended["type"] = typeBytes
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9936340160105504$", v); ok {
			e := Element{
				Link:     "mediaserver:test/slsp$$cover",
				Extended: map[string]json.RawMessage{},
			}

			widthBytes, _ := json.Marshal(1191)
			e.Extended["width"] = widthBytes
			heightBytes, _ := json.Marshal(1684)
			e.Extended["height"] = heightBytes
			typeBytes, _ := json.Marshal("mediaserver")
			e.Extended["type"] = typeBytes
			result = append(result, e)
		}
		// todo: implement for all data (Alma 945, Publishing Anpassung) - currently only used for Tschichold
		if ok, _ := regexp.MatchString("^9972566811505504$", v); ok {
			e := Element{
				Link:     "https: //slsp-ubs.alma.exlibrisgroup.com/view/delivery/thumbnail/41SLSP_UBS/12399106480005504",
				Extended: map[string]json.RawMessage{},
			}

			typeBytes, _ := json.Marshal("file")
			e.Extended["type"] = typeBytes
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

// AddTestFileCount todo: remove, creates test data
func (m *MappingRDV) AddTestFileCount() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return "", nil, false
	}

	key = "fileCount"
	ok = true
	result = []Element{}

	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}

		if ok, _ := regexp.MatchString("^9972650989305504$", v); ok {
			e := Element{
				Text: "1 Audio",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9967555870105504$", v); ok {
			e := Element{
				Text: "1 Video",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9936873350105504$", v); ok {
			e := Element{
				Text: "1 Bild",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9936340160105504$", v); ok {
			e := Element{
				Text: "1 PDF",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9964808620105504$", v); ok {
			e := Element{
				Text: "1 JSON",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9959975600105504$", v); ok {
			e := Element{
				Text: "1 EPub",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^99988730105504$", v); ok {
			e := Element{
				Text: "1 Web archive",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9972789864805504$", v); ok {
			e := Element{
				Text: "4 Audios",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9972529169305504$", v); ok {
			e := Element{
				Text: "3 Videos",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9972536425205504$", v); ok {
			e := Element{
				Text: "12 Bilder",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9972768968505504$", v); ok {
			e := Element{
				Text: "8 PDFs",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9970786120105504$", v); ok {
			e := Element{
				Text: "2 Audios, 3 Videos, 4 Bilder, 2 PDFs",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9972608858805504$", v); ok {
			e := Element{
				Text: "4 Bilder, 3 Videos, 3 PDFs, 7 CSVs, 1 Audio",
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9935139840105504$", v); ok {
			e := Element{
				Text: "14 Audios, 3 Videos, 47 Bilder, 35 PDFs",
			}
			result = append(result, e)
		}

	}

	if len(result) == 0 {
		return "", nil, false
	}
	return
}

// GetTranscription todo: replace once there's data in the index, currently only for testing
func (m *MappingRDV) GetTranscription() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}
	if len(m.Mapping.RecordIdentifier) == 0 {
		return "", nil, false
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
		return "", nil, false
	}

	key = "iiifManifest"
	ok = true
	result = []Element{}

	/* for e-rara and e-manuscripta */
	for _, f := range m.Mapping.Files {
		if f == nil || f.Structure == nil || f.Structure.DigitalObject == nil {
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
		return "", nil, false
	}
	/* for Porträts */
	for _, v := range m.Mapping.Location.Digital {
		if v == nil {
			continue
		}
		portraitUrlPattern := regexp.MustCompile(`.*digi/a100/portraet/`)
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

func (m *MappingRDV) GetPdf() (key string, result []Element, ok bool) {
	if m.Mapping == nil {
		return "", nil, false
	}

	key = "pdf"
	ok = true
	result = []Element{}

	/* for e-rara and e-manuscripta */
	for _, f := range m.Mapping.Files {
		if f == nil || f.Structure == nil || f.Structure.DigitalObject == nil {
			continue
		}
		if f.Structure.DigitalObject.Id == "" {
			continue
		}
		for _, flag := range m.Flags {
			if flag == "e-rara" {
				e := Element{
					Link: "https://www.e-rara.ch/download/pdf/" + strings.Replace(f.Structure.DigitalObject.Id, "md", "", 1),
				}
				result = append(result, e)
			}
			if flag == "e-manuscripta" {
				e := Element{
					Link: "https://www.e-manuscripta.ch/download/pdf/" + strings.Replace(f.Structure.DigitalObject.Id, "md", "", 1),
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

func (m *MappingRDV) GetTestAcl() (key string, result []Element, ok bool) {
	if m.ACL == nil {
		return
	}

	key = "acl"
	ok = true
	result = []Element{}

	for _, v := range m.Mapping.RecordIdentifier {
		if v == "" {
			continue
		}
		if ok, _ := regexp.MatchString("^9943658970105504$", v); ok {
			aclMetaBytes, _ := json.Marshal([]string{"unibas/testuser"})
			aclPreviewBytes, _ := json.Marshal([]string{"unibas/testuser"})
			aclContentBytes, _ := json.Marshal([]string{"unibas/testuser"})
			e := Element{
				Extended: map[string]json.RawMessage{
					"meta":    aclMetaBytes,
					"preview": aclPreviewBytes,
					"content": aclContentBytes,
				},
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9972352018205504$", v); ok {
			aclMetaBytes, _ := json.Marshal([]string{"global/guest"})
			aclPreviewBytes, _ := json.Marshal([]string{"unibas/testuser"})
			aclContentBytes, _ := json.Marshal([]string{"unibas/testuser"})
			e := Element{
				Extended: map[string]json.RawMessage{
					"meta":    aclMetaBytes,
					"preview": aclPreviewBytes,
					"content": aclContentBytes,
				},
			}
			result = append(result, e)
		}
		if ok, _ := regexp.MatchString("^9920530890105504$", v); ok {
			aclMetaBytes, _ := json.Marshal([]string{"global/guest"})
			aclPreviewBytes, _ := json.Marshal([]string{"global/guest"})
			aclContentBytes, _ := json.Marshal([]string{"unibas/testuser"})
			e := Element{
				Extended: map[string]json.RawMessage{
					"meta":    aclMetaBytes,
					"preview": aclPreviewBytes,
					"content": aclContentBytes,
				},
			}
			result = append(result, e)
		} else {
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
		}
	}

	return
}

func (m *MappingRDV) GetAcl() (key string, result []Element, ok bool) {
	if m.ACL == nil {
		return "", nil, false
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
	key, value, ok = m.GetAcl()
	if ok {
		result[key] = value
	}
	// key, value, ok = m.GetTestAcl()
	// if ok {
	// 	result[key] = value
	// }
	key, value, ok = m.GetMedia()
	if ok {
		result[key] = value
	}
	key, value, ok = m.AddTestMedia()
	if ok {
		result[key] = value
	}
	key, value, ok = m.AddTestThumbnail()
	if ok {
		result[key] = value
	}
	key, value, ok = m.AddTestFileCount()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetPdf()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNoteStatementOfResponsibility()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoEdition()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNamePersonal()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSubjectNamePersonal()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNameFamily()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSubjectNameFamily()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNameCorporate()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSubjectNameCorporate()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNameConference()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSubjectNameConference()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSubjectTopic()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNameGeographic()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSubjectGeographic()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetNameRelatedWork()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSubjectTitleInfo()
	if ok {
		result[key] = value
	}
	return
}
