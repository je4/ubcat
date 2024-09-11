package mappingRDV

import (
	"fmt"
	"github.com/je4/ubcat/v2/pkg/schema"
	"strings"
)

type MappingRDV schema.UBSchema001

type Element struct {
	Text string `json:"text,omitempty"`
	Link string `json:"link,omitempty"`
}

func (e Element) String() string {
	if e.Link != "" {
		return fmt.Sprintf("'%s:[%s]'", e.Text, e.Link)
	}
	return fmt.Sprintf("'%s'", e.Text)
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
	key = "originInfo"
	ok = true
	for _, v := range m.Mapping.OriginInfo.Production {
		if v == nil {
			continue
		}
		e := Element{}
		e.Text = strings.TrimRight(strings.Join([]string{
			strings.Join(v.Place, ", "),
			strings.Join(v.Publisher, ", "),
			v.Date}, "; "), " ;")
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetMainTitle() (key string, result []Element, ok bool) {
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
	key = "mainTitle"
	ok = true
	for _, v := range m.Mapping.TitleInfo.Main {
		if v == nil {
			continue
		}
		e := Element{
			Text: v.Title,
		}
		result = append(result, e)
	}
	return
}

func (m *MappingRDV) GetFacetAutographScribe() (key string, result []Element, ok bool) {
	key = "facetAutographScribe"
	ok = true
	result = []Element{}
	for _, f := range m.Facets {
		for _, af := range f.Agents {
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
	key, value, ok := m.GetMainTitle()
	if ok {
		result[key] = value
	}
	key, value, ok = m.GetOriginInfoProduction()
	if ok {
		result[key] = value
	}

	key, value, ok = m.GetFacetAutographScribe()
	if ok {
		result[key] = value
	}
	return
}
