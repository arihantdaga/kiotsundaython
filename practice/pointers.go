package main

type Cat struct {
	name string
}

func modifyCat(c *Cat) {
	c.name = "Mr. Bigglesworth"
}

func dontModifyCat(c Cat) {
	c.name = "Mr. Hello"
}
func main() {
	cat := Cat{name: "Mr. Meow"}
	modifyCat(&cat)
	println(cat.name)
	dontModifyCat(cat)
	println(cat.name)
}
