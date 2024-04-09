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

type ResultTopic struct {
	Name           string
	AlternateNames string
}

type ResultGeographic struct {
	Name           string
	AlternateNames string
	Description    string
}

type ResultTitle struct {
	Name           string
	AlternateNames string
	Title          string
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

func (u *UBSchema001) getTitle(titles []*Title) string {
	if titles == nil {
		return ""
	}

	result := ""
	for _, ti := range titles {
		if len(result) > 0 {
			result += " / "
		}
		result += strings.ReplaceAll(strings.ReplaceAll(ti.Title, "<", ""), ">", "")
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

func (u *UBSchema001) GetMainTitle() string {
	if u.Mapping == nil || u.Mapping.TitleInfo == nil || u.Mapping.TitleInfo.Main == nil {
		return ""
	}
	return u.getTitle(u.Mapping.TitleInfo.Main)
}

func (u *UBSchema001) GetAlternateTitle() string {
	if u.Mapping == nil || u.Mapping.TitleInfo == nil || u.Mapping.TitleInfo.Alternative == nil {
		return ""
	}
	return u.getTitle(u.Mapping.TitleInfo.Alternative)
}

func (u *UBSchema001) GetTranslatedTitle() string {
	if u.Mapping == nil || u.Mapping.TitleInfo == nil || u.Mapping.TitleInfo.Translated == nil {
		return ""
	}
	return u.getTitle(u.Mapping.TitleInfo.Translated)
}

func (u *UBSchema001) GetUniformTitle() string {
	if u.Mapping == nil || u.Mapping.TitleInfo == nil || u.Mapping.TitleInfo.Uniform == nil {
		return ""
	}
	return u.getTitle(u.Mapping.TitleInfo.Uniform)
}

func (u *UBSchema001) GetPublicationPlace() string {
	// todo: make usable for all OriginInfo
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

func (u *UBSchema001) GetAbstract() string {
	if u.Mapping == nil || u.Mapping.Abstract == nil {
		return ""
	}
	result := strings.Join(u.Mapping.Abstract, " ; ")
	return result
}

func (u *UBSchema001) GetDdc() string {
	if u.Mapping == nil || u.Mapping.Classification == nil || u.Mapping.Classification.Ddc == nil {
		return ""
	}
	result := ""
	if u.Mapping.Classification.Ddc.All != nil {
		result += strings.Join(u.Mapping.Classification.Ddc.All, " ; ")
	}
	if u.Mapping.Classification.Ddc.Sdnb != nil {
		result += strings.Join(u.Mapping.Classification.Ddc.Sdnb, " ; ")
	}
	if u.Mapping.Classification.Ddc.Ed23 != nil {
		result += strings.Join(u.Mapping.Classification.Ddc.Ed23, " ; ")
	}
	return result
}

func (u *UBSchema001) GetRvk() string {
	if u.Mapping == nil || u.Mapping.Classification.Rvk == nil {
		return ""
	}
	result := strings.Join(u.Mapping.Classification.Rvk, " ; ")
	return result
}

func (u *UBSchema001) GetUdc() string {
	if u.Mapping == nil || u.Mapping.Classification.Udc == nil {
		return ""
	}
	result := strings.Join(u.Mapping.Classification.Udc, " ; ")
	return result
}

func (u *UBSchema001) GetGenre() string {
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.Genre == nil {
		return ""
	}

	var genres []string
	for _, genreSlice := range u.Mapping.Subject.Genre {
		genres = append(genres, genreSlice...)
	}

	result := strings.Join(genres, " ; ")
	return result
}

func (u *UBSchema001) GetSubjectPersons() []ResultPerson {
	var result []ResultPerson
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.Name == nil || u.Mapping.Subject.Name.Personal == nil {
		return result
	}
	for _, ps := range u.Mapping.Subject.Name.Personal {
		for _, p := range ps {
			contains := false
			for _, entity := range result {
				if entity.Name == p.NamePart {
					contains = true
					break
				}
			}
			if !contains {
				newEntity := ResultPerson{
					Name: p.NamePart,
				}
				if p.Date != nil {
					newEntity.Date = p.Date.Original
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result = append(result, newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetSubjectFamilies() []ResultFamily {
	var result []ResultFamily
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.Name == nil || u.Mapping.Subject.Name.Family == nil {
		return result
	}
	for _, ps := range u.Mapping.Subject.Name.Family {
		for _, p := range ps {
			contains := false
			for _, entity := range result {
				if entity.Name == p.NamePart {
					contains = true
					break
				}
			}
			if !contains {
				newEntity := ResultFamily{
					Name: p.NamePart,
				}
				if p.Date != "" {
					newEntity.Date = p.Date
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result = append(result, newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetSubjectConferences() []ResultConference {
	var result []ResultConference
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.Name == nil || u.Mapping.Subject.Name.Conference == nil {
		return result
	}
	for _, ps := range u.Mapping.Subject.Name.Conference {
		for _, p := range ps {
			contains := false
			for _, entity := range result {
				if entity.Name == p.NamePart {
					contains = true
					break
				}
			}
			if !contains {
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
				result = append(result, newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetSubjectCorporates() []ResultCorporate {
	var result []ResultCorporate
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.Name == nil || u.Mapping.Subject.Name.Corporate == nil {
		return result
	}
	for _, ps := range u.Mapping.Subject.Name.Corporate {
		for _, p := range ps {
			contains := false
			for _, entity := range result {
				if entity.Name == p.NamePart {
					contains = true
					break
				}
			}
			if !contains {
				newEntity := ResultCorporate{
					Name: p.NamePart,
				}
				if p.Description != nil {
					newEntity.Description = strings.Join(p.Description, ", ")
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result = append(result, newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetSubjectTopics() []ResultTopic {
	var result []ResultTopic
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.Topic == nil {
		return result
	}
	for _, ps := range u.Mapping.Subject.Topic {
		for _, p := range ps {
			contains := false
			for _, entity := range result {
				if entity.Name == p.Label {
					contains = true
					break
				}
			}
			if !contains {
				newEntity := ResultTopic{
					Name: p.Label,
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result = append(result, newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetSubjectGeographics() []ResultGeographic {
	var result []ResultGeographic
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.Geographic == nil {
		return result
	}
	for _, ps := range u.Mapping.Subject.Geographic {
		for _, p := range ps {
			contains := false
			for _, entity := range result {
				if entity.Name == p.NamePart {
					contains = true
					break
				}
			}
			if !contains {
				newEntity := ResultGeographic{
					Name: p.NamePart,
				}
				if p.Description != nil {
					newEntity.Description = strings.Join(p.Description, ", ")
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result = append(result, newEntity)
			}
		}
	}
	return result
}

func (u *UBSchema001) GetSubjectTitles() []ResultTitle {
	var result []ResultTitle
	if u.Mapping == nil || u.Mapping.Subject == nil || u.Mapping.Subject.TitleInfo == nil {
		return result
	}
	for _, ps := range u.Mapping.Subject.TitleInfo {
		for _, p := range ps {
			contains := false
			for _, entity := range result {
				if entity.Title == p.Title {
					contains = true
					break
				}
			}
			if !contains {
				newEntity := ResultTitle{
					Title: p.Title,
				}
				if p.Name != "" {
					newEntity.Name = p.Name
				}
				if p.Variant != nil {
					newEntity.AlternateNames = strings.Join(p.Variant, " ; ")
				}
				result = append(result, newEntity)
			}
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
