package model

type Category struct {
	ID   int
	Name string
}

// Define constant categories with IDs and names
var Categories = []Category{
	{ID: 1, Name: "Food"},
	{ID: 2, Name: "Consumption"},
	{ID: 3, Name: "Transportation"},
	{ID: 4, Name: "Household"},
}
