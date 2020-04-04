package data

import "fmt"

type Book struct {
	ID int
	Title string
	Author string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("ID:\t\t%v\n" + "Title:\t\t%q\n" + "Author:\t\t%q\n" + "Published:\t\t%v\n",
		b.ID, b.Title, b.Author, b.YearPublished )
}

var Books = []Book{
	{
		ID:            1,
		Title:         "La pomne",
		Author:        "Douglas Blissman",
		YearPublished: 2057,
	},
	{
		ID:            2,
		Title:         "Covid-19 Pandemic",
		Author:        "Lee Xi Shan",
		YearPublished: 2022,
	},
	{
		ID:            3,
		Title:         "La pomne 3",
		Author:        "Douglas Blissman",
		YearPublished: 2058,
	},
	{
		ID:            4,
		Title:         "Covid-19 x Pandemic",
		Author:        "Lee Xi Shan",
		YearPublished: 2022,
	},
	{
		ID:            5,
		Title:         "pomne",
		Author:        "issman",
		YearPublished: 2005,
	},
	{
		ID:            6,
		Title:         "Contaiment",
		Author:        "Lee Xi Shan",
		YearPublished: 2018,
	},
}