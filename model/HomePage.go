package model

type Todo struct {
	Name string
	Description string
}

type HomePageData struct {
	PageTitle string
	Todos     []Todo
}