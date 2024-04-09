package schema

type UBSchema struct {
	UBSchema001
	Score_ float64 `json:"score"`
}

/*
func (u *UBSchema) GetPersons() map[string][]ResultPerson {
	return (*UBSchema001)(u).GetPersons()
}

func (u *UBSchema) getTitle() string {
	return (*UBSchema001)(u).getTitle()
}

*/
