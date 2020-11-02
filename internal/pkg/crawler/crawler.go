package crawler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/gocolly/colly/v2"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"regexp"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func StartCrawler(db *sqlx.DB,s3 *s3.S3,log *logger.CustomLogger)  {
	//rep := hotelRepository.NewPostgresHotelRepository(db)

	hotels := []hotelmodel.Hotel{}

	c := colly.NewCollector(
		colly.AllowedDomains("www.booking.com"),
		colly.URLFilters(
			regexp.MustCompile("^https://www.booking.com/hotel/ru"),
			),
		colly.Async(),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "booking.com",
		Delay: 1 * time.Second,
		RandomDelay: 1 * time.Second,
	})


	err := c.Visit("https://www.booking.com/hotel/ru")
	if err != nil{
		log.Error(err)
	}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("div[id=property_description_content]",func(e *colly.HTMLElement) {

		fmt.Println("here")
		e.ForEach("p", func(num int,e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent",RandomString())
	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())

	})
	c.Wait()
	fmt.Println(hotels)
}