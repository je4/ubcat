package schema

import (
	"encoding/json"
	"slices"
)

type ResultPerson struct {
	Name string
	Date string
}

func (u *UBSchema001) GetPersons() map[string][]ResultPerson {
	// todo: deal with alternative names
	result := make(map[string][]ResultPerson)
	if u.Mapping == nil || u.Mapping.Name == nil || u.Mapping.Name.Personal == nil {
		return result
	}
	for _, ps := range u.Mapping.Name.Personal {
		for _, p := range ps {
			role := "author"
			if len(p.Role) > 0 {
				role = p.Role[0]
			}
			if _, ok := result[role]; !ok {
				result[role] = make([]ResultPerson, 0)
			}
			if !slices.ContainsFunc(result[role], func(person ResultPerson) bool {
				return person.Name == p.NamePart
			}) {
				result[role] = append(result[role], ResultPerson{
					Name: p.NamePart,
					Date: p.Date.Original,
				})
			}
		}
	}
	return result
}

func (u *UBSchema001) GetTitle() string {
	if u.Mapping == nil || u.Mapping.TitleInfo == nil || u.Mapping.TitleInfo.Main == nil {
		return ""
	}
	result := ""
	for _, ti := range u.Mapping.TitleInfo.Main {
		if len(result) > 0 {
			result += " / "
		}
		result = ti.Title
		if ti.SubTitle != "" {
			result += " - " + ti.SubTitle
		}
	}
	return result
}

func (u *UBSchema001) GetJSON() string {
	// todo: create AI friendly JSON
	jsonBytes, err := json.MarshalIndent(u.Mapping, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
