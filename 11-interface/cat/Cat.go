package cat

type Cat struct {
	Name string `json:"name"`
}

// Create a constructor for Cat
func NewCat(name string) *Cat {
	//if name == "" {
	//	fmt.Println("name is empty")
	//}
	return &Cat{
		Name: name,
	}
}
func (c Cat) GetName() string {
	return c.Name
}

func (c Cat) Speak() string {
	return "Meo Meo"
}
