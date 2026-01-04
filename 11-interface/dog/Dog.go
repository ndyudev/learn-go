package dog

type Dog struct {
	Name string `json:"name"`
}

// create a constructor for Dog
func New(name string) *Dog {
	return &Dog{
		Name: name,
	}
}

func (d Dog) GetName() string {
	return d.Name
}

func (d Dog) Speak() string {
	return "Gau Gau"
}
