package schema

import (
	"encoding/json"
	"slices"
	"strings"
)

type ResultPerson struct {
	Name           string
	Date           string
	AlternateNames string
}

type ResultFamily struct {
	Name           string
	Date           string
	AlternateNames string
}

type ResultCorporate struct {
	Name           string
	Description    string
	AlternateNames string
}

type ResultConference struct {
	Name           string
	Date           string
	Description    string
	AlternateNames string
}

func (u *UBSchema001) GetPersons() map[string][]ResultPerson {
	result := make(map[string][]ResultPerson)
	if u.Mapping == nil || u.Mapping.Name == nil || u.Mapping.Name.Personal == nil {
		return result
	}
	for _, ps := range u.Mapping.Name.Personal {
		for _, p := range ps {
			role := "author"
			if len(p.Role) > 0 {
				role = strings.Join(p.Role, ", ")
			}
			if _, ok := result[role]; !ok {
				result[role] = make([]ResultPerson, 0)
			}
			if !slices.ContainsFunc(result[role], func(entity ResultPerson) bool {
				return entity.Name == p.NamePart
			}) {
				newEntity := ResultPerson{
					Name: p.NamePart,
				}
				if p.Date != nil {
					newEntity.Date = p.Date.Original
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result[role] = append(result[role], newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetFamilies() map[string][]ResultFamily {
	result := make(map[string][]ResultFamily)
	if u.Mapping == nil || u.Mapping.Name == nil || u.Mapping.Name.Family == nil {
		return result
	}
	for _, ps := range u.Mapping.Name.Family {
		for _, p := range ps {
			role := "author"
			if len(p.Role) > 0 {
				role = strings.Join(p.Role, ", ")
			}
			if _, ok := result[role]; !ok {
				result[role] = make([]ResultFamily, 0)
			}
			if !slices.ContainsFunc(result[role], func(entity ResultFamily) bool {
				return entity.Name == p.NamePart
			}) {
				newEntity := ResultFamily{
					Name: p.NamePart,
				}
				if p.Date != "" {
					newEntity.Date = p.Date
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result[role] = append(result[role], newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetCorporates() map[string][]ResultCorporate {
	result := make(map[string][]ResultCorporate)
	if u.Mapping == nil || u.Mapping.Name == nil || u.Mapping.Name.Corporate == nil {
		return result
	}
	for _, ps := range u.Mapping.Name.Corporate {
		for _, p := range ps {
			role := "author"
			if len(p.Role) > 0 {
				role = strings.Join(p.Role, ", ")
			}
			if _, ok := result[role]; !ok {
				result[role] = make([]ResultCorporate, 0)
			}
			if !slices.ContainsFunc(result[role], func(entity ResultCorporate) bool {
				return entity.Name == p.NamePart
			}) {
				newEntity := ResultCorporate{
					Name: p.NamePart,
				}
				if p.Description != nil {
					newEntity.Description = strings.Join(p.Description, ", ")
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result[role] = append(result[role], newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetConferences() map[string][]ResultConference {
	result := make(map[string][]ResultConference)
	if u.Mapping == nil || u.Mapping.Name == nil || u.Mapping.Name.Conference == nil {
		return result
	}
	for _, ps := range u.Mapping.Name.Conference {
		for _, p := range ps {
			role := "author"
			if len(p.Role) > 0 {
				role = strings.Join(p.Role, ", ")
			}
			if _, ok := result[role]; !ok {
				result[role] = make([]ResultConference, 0)
			}
			if !slices.ContainsFunc(result[role], func(entity ResultConference) bool {
				return entity.Name == p.NamePart
			}) {
				newEntity := ResultConference{
					Name: p.NamePart,
				}
				if p.Date != "" {
					newEntity.Date = p.Date
				}
				if p.Description != nil {
					newEntity.Description = strings.Join(p.Description, ", ")
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result[role] = append(result[role], newEntity)
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
			result += " : " + ti.SubTitle
		}
		if ti.PartNumber != nil {
			result += ". " + strings.Join(ti.PartNumber, ", ")
		}
		if ti.PartName != nil {
			result += ". " + strings.Join(ti.PartName, ", ")
		}
	}
	return result
}

func (u *UBSchema001) GetPublicationPlace() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Publication == nil {
		return ""
	}
	result := ""
	for _, f := range u.Mapping.OriginInfo.Publication {
		if len(result) > 0 {
			result += " / "
		}
		if f.Place != nil {
			result = strings.Join(f.Place, ", ")
		}
	}
	return result
}

func (u *UBSchema001) GetPublisher() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Publication == nil {
		return ""
	}
	result := ""
	for _, f := range u.Mapping.OriginInfo.Publication {
		if len(result) > 0 {
			result += " / "
		}
		if f.Publisher != nil {
			result = strings.Join(f.Publisher, ", ")
		}
	}
	return result
}

func (u *UBSchema001) GetPublicationDate() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Publication == nil {
		return ""
	}
	result := ""
	for _, f := range u.Mapping.OriginInfo.Publication {
		if len(result) > 0 {
			result += " / "
		}
		if f.Date != "" {
			result = f.Date
		}
	}
	return result
}

func (u *UBSchema001) GetAIJSON() string {
	// todo: create AI friendly JSON
	// todo: remove files, location.holding.item
	jsonBytes, err := json.MarshalIndent(u.Mapping, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
