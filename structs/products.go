package structs

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Count       int     `json:"count"`
}

type User struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Money float64 `json:"money"`
}

/* func (p Product) GetProdId() int {
	return p.id
}

func (u User) GetUserId() int {
	return u.id
} */
