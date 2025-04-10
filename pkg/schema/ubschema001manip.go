package schema

func (u *UBSchema001) DoFacets() error {
	if u.Mapping == nil {
		return nil
	}
	if u.Facets == nil {
		u.Facets = &Facets{
			Agents:    []*AgentFacets{},
			Concepts:  []*ConceptFacets{},
			Daterange: []*DateRangeFacets{},
			Strings:   []*StringFacets{},
		}
	}
	u.Facets.Strings = append(u.Facets.Strings, &StringFacets{
		Name:   "language",
		String: u.Mapping.Language,
	})
	return nil
}
