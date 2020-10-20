package main

import (
	"github.com/google/uuid"
	// import standard libraries

	"fmt"
	"log"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Hotel struct {
	// ID          int `db:"hotel_id"`
	Name        string `db:"name"`
	Location    string `db:"location`
	Description string `db:"description"`
	Image       string `db:"img"`
	Rating      int    `db:"curr_rating"`
}

type HotelStore interface {
	Hotel(id int) (Hotel, error)
	Hotels() ([]Hotel, error)
	CreateHotel(hotel *Hotel) error
	UpdateHotel(hotel *Hotel) error
	DeleteHotel(id uuid.UUID) error
}

func Scrape() {
	var hotels [25]Hotel
	doc, err := goquery.NewDocument("https://www.booking.com/searchresults.ru.html?aid=397594;label=gog235jc-1DCAEoggI46AdIIVgDaMIBiAEBmAEhuAEZyAEM2AED6AEB-AECiAIBqAIDuAKx_5_8BcACAdICJDc2OWQxYzUzLThlNGEtNGJiZC1iZjExLTMyZWM2NDE1ZThkZdgCBOACAQ;sid=4b43cee7d7757d5a2192af4b5d81f254;dest_id=-2960561;dest_type=city&")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".sr-hotel__name\n").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		hotels[index].Name = title
	})
	doc.Find(".hotel_desc").Each(func(index int, item *goquery.Selection) {
		desc := item.Text()
		hotels[index].Description = desc
	})
	for i, el := range hotels {
		fmt.Printf("%d %s: %s\n", i, el.Name, el.Description)
	}

	doc.Find(".hotel_desc").Each(func(index int, item *goquery.Selection) {
		desc := item.Text()
		hotels[index].Description = desc
	})
	doc.Find(".hotel_image").Each(func(index int, item *goquery.Selection) {
		source, _ := item.Attr("data-highres")
		hotels[index].Image = source
	})

	db, err := sqlx.Connect("postgres", "user=testuser password=12345 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	tx := db.MustBegin()
	for i, el := range hotels {
		tx.MustExec("INSERT INTO hotels (name, location, description, img, curr_rating) VALUES ($1, $2, $3, $4, $5)", el.Name, el.Location, el.Description, el.Image, el.Rating)
		fmt.Printf("%d %s: %s %s\n", i, el.Name, el.Description, el.Image)
	}
	tx.Commit()

}

var schema = `
create table hotels (
    name text ,
    location text,
    description text,
    img text,
    curr_rating int DEFAULT 0 CHECK (curr_rating >= 0  AND curr_rating <=10)
);`

func main() {
	Scrape()
}
