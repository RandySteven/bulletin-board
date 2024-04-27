package enums

type UserGender string

const (
	Male   UserGender = `male`
	Female            = `female`
)

func (u UserGender) ToString() string {
	return string(u)
}
