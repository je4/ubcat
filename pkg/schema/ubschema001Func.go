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

/* temporary solution for hackathon */
var roleTranslations = map[string]string{
	"aut": "author",
	"edt": "editor",
	"trl": "translator",
	"cre": "creator",
	"cmp": "composer",
	"ctb": "contributor",
	"ill": "illustrator",
	"fmo": "former owner",
	"prt": "printer",
	"dgs": "degree supervisor",
	"art": "artist",
	"hnr": "honoree",
	"com": "compiler",
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
				translatedRoles := make([]string, len(p.Role))
				for i, r := range p.Role {
					translatedRole, ok := roleTranslations[r]
					if ok {
						translatedRoles[i] = translatedRole
					} else {
						translatedRoles[i] = r
					}
				}
				role = strings.Join(translatedRoles, ", ")
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
				translatedRoles := make([]string, len(p.Role))
				for i, r := range p.Role {
					translatedRole, ok := roleTranslations[r]
					if ok {
						translatedRoles[i] = translatedRole
					} else {
						translatedRoles[i] = r
					}
				}
				role = strings.Join(translatedRoles, ", ")
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
				translatedRoles := make([]string, len(p.Role))
				for i, r := range p.Role {
					translatedRole, ok := roleTranslations[r]
					if ok {
						translatedRoles[i] = translatedRole
					} else {
						translatedRoles[i] = r
					}
				}
				role = strings.Join(translatedRoles, ", ")
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
				translatedRoles := make([]string, len(p.Role))
				for i, r := range p.Role {
					translatedRole, ok := roleTranslations[r]
					if ok {
						translatedRoles[i] = translatedRole
					} else {
						translatedRoles[i] = r
					}
				}
				role = strings.Join(translatedRoles, ", ")
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
		result += strings.TrimSuffix(strings.ReplaceAll(strings.ReplaceAll(ti.Title, "<", ""), ">", ""), " :")
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

func (u *UBSchema001) GetAbbreviatedTitle() string {
	if u.Mapping == nil || u.Mapping.TitleInfo == nil || u.Mapping.TitleInfo.Abbreviated == nil {
		return ""
	}
	result := ""
	if u.Mapping.TitleInfo.Abbreviated != nil {
		result += strings.Join(u.Mapping.TitleInfo.Abbreviated, " / ")
	}
	return result
}

func (u *UBSchema001) getOriginPlace(publicationNotes []*PublicationNote) string {
	if publicationNotes == nil {
		return ""
	}
	result := ""
	for _, pn := range publicationNotes {
		if len(result) > 0 {
			result += " / "
		}
		if pn.Place != nil {
			result = strings.Join(pn.Place, ", ")
		}
	}
	return result
}

func (u *UBSchema001) GetPublicationPlace() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Publication == nil {
		return ""
	}
	return u.getOriginPlace(u.Mapping.OriginInfo.Publication)
}

func (u *UBSchema001) GetDistributionPlace() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Distribution == nil {
		return ""
	}
	return u.getOriginPlace(u.Mapping.OriginInfo.Distribution)
}

func (u *UBSchema001) GetManufacturePlace() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Manufacture == nil {
		return ""
	}
	return u.getOriginPlace(u.Mapping.OriginInfo.Manufacture)
}

func (u *UBSchema001) GetProductionPlace() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Production == nil {
		return ""
	}
	return u.getOriginPlace(u.Mapping.OriginInfo.Production)
}

func (u *UBSchema001) getOriginPublisher(publicationNotes []*PublicationNote) string {
	if publicationNotes == nil {
		return ""
	}
	result := ""
	for _, pn := range publicationNotes {
		if len(result) > 0 {
			result += " / "
		}
		if pn.Publisher != nil {
			result = strings.Join(pn.Publisher, ", ")
		}
	}
	return result
}

func (u *UBSchema001) GetPublicationPublisher() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Publication == nil {
		return ""
	}
	return u.getOriginPublisher(u.Mapping.OriginInfo.Publication)
}

func (u *UBSchema001) GetDistributionPublisher() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Distribution == nil {
		return ""
	}
	return u.getOriginPublisher(u.Mapping.OriginInfo.Distribution)
}

func (u *UBSchema001) GetManufacturePublisher() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Manufacture == nil {
		return ""
	}
	return u.getOriginPublisher(u.Mapping.OriginInfo.Manufacture)
}

func (u *UBSchema001) GetProductionPublisher() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Production == nil {
		return ""
	}
	return u.getOriginPublisher(u.Mapping.OriginInfo.Production)
}

func (u *UBSchema001) getOriginDate(publicationNotes []*PublicationNote) string {
	if publicationNotes == nil {
		return ""
	}
	result := ""
	for _, pn := range publicationNotes {
		if len(result) > 0 {
			result += " / "
		}
		if pn.Date != "" {
			result = pn.Date
		}
	}
	return result
}

func (u *UBSchema001) GetPublicationDate() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Publication == nil {
		return ""
	}
	return u.getOriginDate(u.Mapping.OriginInfo.Publication)
}

func (u *UBSchema001) GetDistributionDate() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Distribution == nil {
		return ""
	}
	return u.getOriginDate(u.Mapping.OriginInfo.Distribution)
}

func (u *UBSchema001) GetManufactureDate() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Manufacture == nil {
		return ""
	}
	return u.getOriginDate(u.Mapping.OriginInfo.Manufacture)
}

func (u *UBSchema001) GetProductionDate() string {
	if u.Mapping == nil || u.Mapping.OriginInfo == nil || u.Mapping.OriginInfo.Production == nil {
		return ""
	}
	return u.getOriginDate(u.Mapping.OriginInfo.Production)
}

func (u *UBSchema001) GetAbstract() string {
	if u.Mapping == nil || u.Mapping.Abstract == nil {
		return ""
	}
	result := strings.Join(u.Mapping.Abstract, " ; ")
	return result
}

func (u *UBSchema001) GetTableOfContents() string {
	if u.Mapping == nil || u.Mapping.TableOfContents == nil {
		return ""
	}
	result := strings.Join(u.Mapping.TableOfContents, " ; ")
	return result
}

func (u *UBSchema001) GetSeriesTitle() string {
	if u.Mapping == nil || u.Mapping.RelatedItem == nil || u.Mapping.RelatedItem.Series == nil {
		return ""
	}
	result := ""
	for _, se := range u.Mapping.RelatedItem.Series {
		if len(result) > 0 {
			result += " / "
		}
		if se.Title != "" {
			result = se.Title
		}
	}
	return result
}

func (u *UBSchema001) GetHostTitle() string {
	if u.Mapping == nil || u.Mapping.RelatedItem == nil || u.Mapping.RelatedItem.Host == nil {
		return ""
	}
	result := ""
	for _, se := range u.Mapping.RelatedItem.Host {
		if len(result) > 0 {
			result += " / "
		}
		if se.Title != "" {
			result = se.Title
		}
	}
	return result
}

func (u *UBSchema001) GetGeneralNote() string {
	if u.Mapping == nil || u.Mapping.Note == nil || u.Mapping.Note.General == nil {
		return ""
	}
	result := ""
	result += strings.Join(u.Mapping.Note.General, " ; ")
	return result
}

func (u *UBSchema001) GetPerformersNote() string {
	if u.Mapping == nil || u.Mapping.Note == nil || u.Mapping.Note.Performers == nil {
		return ""
	}
	result := ""
	result += strings.Join(u.Mapping.Note.Performers, " ; ")
	return result
}

func (u *UBSchema001) GetCreditsNote() string {
	if u.Mapping == nil || u.Mapping.Note == nil || u.Mapping.Note.Credits == nil {
		return ""
	}
	result := ""
	result += strings.Join(u.Mapping.Note.Credits, " ; ")
	return result
}

func (u *UBSchema001) GetStatementOfResponsibility() string {
	if u.Mapping == nil || u.Mapping.Note == nil || u.Mapping.Note.StatementOfResponsibility == nil {
		return ""
	}
	result := ""
	result += strings.Join(u.Mapping.Note.StatementOfResponsibility, " ; ")
	return result
}

func (u *UBSchema001) GetMediumOfPerformance() string {
	if u.Mapping == nil || u.Mapping.Note == nil || u.Mapping.Note.MediumOfPerformance == nil {
		return ""
	}
	result := ""
	result += strings.Join(u.Mapping.Note.MediumOfPerformance, " ; ")
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
	if u.Mapping == nil || u.Mapping.Classification == nil || u.Mapping.Classification.Rvk == nil {
		return ""
	}
	result := strings.Join(u.Mapping.Classification.Rvk, " ; ")
	return result
}

func (u *UBSchema001) GetUdc() string {
	if u.Mapping == nil || u.Mapping.Classification == nil || u.Mapping.Classification.Udc == nil {
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

func (u *UBSchema001) GetDoi() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.Doi == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.Doi != nil {
		result += strings.Join(u.Mapping.Identifier.Doi, " / ")
	}
	return result
}

func (u *UBSchema001) GetIsbnInvalid() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.IsbnInvalid == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.IsbnInvalid != nil {
		result += strings.Join(u.Mapping.Identifier.IsbnInvalid, " / ")
	}
	return result
}

func (u *UBSchema001) GetIssn() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.Issn == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.Issn != nil {
		result += strings.Join(u.Mapping.Identifier.Issn, " / ")
	}
	return result
}

func (u *UBSchema001) GetIssnInvalid() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.IssnInvalid == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.IssnInvalid != nil {
		result += strings.Join(u.Mapping.Identifier.IssnInvalid, " / ")
	}
	return result
}

func (u *UBSchema001) GetIssnl() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.Issnl == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.Issnl != nil {
		result += strings.Join(u.Mapping.Identifier.Issnl, " / ")
	}
	return result
}

func (u *UBSchema001) GetIssueNumber() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.IssueNumber == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.IssueNumber != nil {
		result += strings.Join(u.Mapping.Identifier.IssueNumber, " / ")
	}
	return result
}

func (u *UBSchema001) GetMatrixNumber() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.MatrixNumber == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.MatrixNumber != nil {
		result += strings.Join(u.Mapping.Identifier.MatrixNumber, " / ")
	}
	return result
}

func (u *UBSchema001) GetMusicPlate() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.MusicPlate == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.MusicPlate != nil {
		result += strings.Join(u.Mapping.Identifier.MusicPlate, " / ")
	}
	return result
}

func (u *UBSchema001) GetMusicPublisher() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.MusicPublisher == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.MusicPublisher != nil {
		result += strings.Join(u.Mapping.Identifier.MusicPublisher, " / ")
	}
	return result
}

func (u *UBSchema001) GetUrn() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.Urn == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.Urn != nil {
		result += strings.Join(u.Mapping.Identifier.Urn, " / ")
	}
	return result
}

func (u *UBSchema001) GetVideoRecordingIdentifier() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.VideoRecordingIdentifier == nil {
		return ""
	}
	result := ""
	if u.Mapping.Identifier.VideoRecordingIdentifier != nil {
		result += strings.Join(u.Mapping.Identifier.VideoRecordingIdentifier, " / ")
	}
	return result
}

func (u *UBSchema001) GetIsbn() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.Isbn == nil {
		return ""
	}
	result := ""
	for _, id := range u.Mapping.Identifier.Isbn {
		if len(result) > 0 {
			result += " / "
		}
		if id.Id != "" {
			result = id.Id
		}
	}
	return result
}

func (u *UBSchema001) GetIsmn() string {
	if u.Mapping == nil || u.Mapping.Identifier == nil || u.Mapping.Identifier.Ismn == nil {
		return ""
	}
	result := ""
	for _, id := range u.Mapping.Identifier.Ismn {
		if len(result) > 0 {
			result += " / "
		}
		if id.Id != "" {
			result = id.Id
		}
	}
	return result
}

func (u *UBSchema001) GetResourceType() string {
	if u.Ldr == nil {
		return ""
	}
	recordType, ok := u.Ldr["leader_06_typeOfRecord"]
	if !ok {
		return ""
	}
	switch recordType {
	case "a":
		bibStatus, ok := u.Ldr["leader_07_bibliographicStatus"]
		if !ok {
			return "Language material"
		}
		switch bibStatus {
		case "m":
			return "Book, Monograph"
		case "s":
			return "Journal, Serial"
		case "c":
			return "Collection of documents"
		case "a":
			return "Article"
		default:
			return "Book, Monograph"
		}
	case "c":
		return "Sheet music"
	case "d":
		return "Manuscript sheet music"
	case "e":
		return "Map"
	case "f":
		return "Manuscript map"
	case "g":
		return "Film"
	case "i":
		return "Text recording"
	case "j":
		return "Music recording"
	case "k":
		return "Image"
	case "m":
		return "Computer file"
	case "o":
		return "Kit"
	case "p":
		return "Archive material"
	case "r":
		return "Object"
	case "t":
		bibStatus, ok := u.Ldr["leader_07_bibliographicStatus"]
		if !ok {
			return "Language material"
		}
		switch bibStatus {
		case "m":
			return "Manuscript Book, Monograph"
		case "s":
			return "Manuscript Journal, Serial"
		default:
			return "Manuscript"
		}
	default:
		return "Unknown"
	}
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
