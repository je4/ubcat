package mappingRDV

import (
	"encoding/json"
	"fmt"
	"regexp"
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
			if len(a.Identifer) > 0 {
				e.Link = fmt.Sprintf("facet:author:%s", a.Identifer[0])
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
			if len(a.Identifer) > 0 {
				e.Link = fmt.Sprintf("facet:scribe:%s", a.Identifer[0])
			}
			result = append(result, e)
		}
	}

	if len(result) == 0 {
		return "", nil, false
	}
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
	key, value, ok = m.GetOriginInfoDistribution()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoManufacture()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoProduction()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoPublication()
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
	key, value, ok = m.GetLocationDigital()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetSwisscollectionsUrl()
	if ok {
		result[key] = value
	}
	return
}
