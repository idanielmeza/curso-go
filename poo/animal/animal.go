package animal

type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	println("Woof! My name is", d.Name)
}

func NewDog(name string) *Dog {
	return &Dog{name}
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() {
	println("Meow! My name is", c.Name)
}

func NewCat(name string) *Cat {
	return &Cat{name}
}
