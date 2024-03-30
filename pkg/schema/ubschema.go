package schema

type UBSchema UBSchema001

func (u *UBSchema) GetPersons() map[string][]ResultPerson {
	return (*UBSchema001)(u).GetPersons()
}

func (u *UBSchema) GetTitle() string {
	return (*UBSchema001)(u).GetTitle()
}
