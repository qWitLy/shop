package structs

type User struct {
	Id       int     `json:"id"`
	Login    string  `json:"name"`
	Money    float64 `json:"money"`
	Password string  `json:"password:"`
}
