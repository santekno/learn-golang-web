package main

type Page struct {
	Title string
	Name  string
}

func (myPage Page) SayHello(name string) string {
	return "Hello " + name + " , My name is " + myPage.Name
}
