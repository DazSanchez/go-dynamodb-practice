package forms

type CreateUserDTO struct {
	Name         string `json:"name"`
	FirstSurname string `json:"firstSurname"`
	Email        string `json:"email"`
}
