package models

type Usuario struct {
	//id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewUser(name string, age int) Usuario {
	u := Usuario{
		Name: name,
		Age:  age,
	}
	return u
}
