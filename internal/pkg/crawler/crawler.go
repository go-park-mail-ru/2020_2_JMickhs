package crawler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	hotelmodel "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/models"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/error"
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"
	"github.com/gocolly/colly/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"

	"io/ioutil"
	"math/rand"
	"net/http"

	"regexp"
	"strings"
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
		Parallelism: 2,
		DomainGlob:  "booking.com",
		Delay: 1 * time.Second,
		RandomDelay: 1 * time.Second,
	})


	err := c.Visit("https://www.booking.com/hotel/ru.ru.html?label=gen173nr-1FCAsowgE46AdIM1gEaMIBiAEBmAEhuAEZyAEP2AEB6AEB-AECiAIBqAIDuAKDy4n9BcACAdICJG" +
		"JhZjExOWI4LTQxMDgtNDgxNy1hOWY1LTU3MDA1NmNkZTVjZdgCBeACAQ;sid=cbab8c7a4faa82faf3cda60ef0432fcc;dist=0&keep_landing=1&sb_price_type=total&type=total&")
	if err != nil{
		log.Error(err)
	}
	c.OnHTML("div[class=block_third]",func(e *colly.HTMLElement) {
		url,_ := e.DOM.Find("a[href]").Attr("href")
		c.Visit(e.Response.Request.AbsoluteURL(url))
	})

	c.OnHTML("div[id=right]",func(e *colly.HTMLElement) {
		hotel := hotelmodel.Hotel{}
		hotel.Name = e.DOM.Find("h2[id=hp_hotel_name]").RemoveAttr("span[class]").Text()
		hotel.Name = strings.Split(hotel.Name,"\n")[1]
		var decr string
		sel := e.DOM.Find("div[id=property_description_content]").Children()
		sel.Each(func(_ int, selection *goquery.Selection){
			decr+= selection.Text()
		})
		hotel.Description = decr
		hotel.Location = e.DOM.Find("p[id=showMap2]").
			Find("span").Text()
		hotel.Location = strings.Split(hotel.Location,"\n")[1]

		imageRef,_ := e.DOM.Find(`a[class="bh-photo-grid-item bh-photo-grid-photo1 active-image "]`).Attr("href")
		name,err := UploadImage(s3,imageRef)
		if err != nil{
			log.Error(err)
		}
		hotel.Image = name
		var photos []string
		e.DOM.Find(`div[class="bh-photo-grid-thumbs bh-photo-grid-thumbs-s-full"]`).Find("a[class]").
			Each(func(number int,selection *goquery.Selection){
			   ref,_ := selection.Attr("href")
				photo, err := UploadImage(s3,ref)
				if err != nil{
					log.Error(err)
				}
				photos = append(photos,photo)
		})
		hotel.Photos = photos
		err = UploadHotel(db,hotel)
		if err != nil{
			log.Error(err)
		}
		fmt.Println(hotel.Name)
		hotels = append(hotels,hotel)

	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent",RandomString())
	})

	c.Wait()
}

func UploadHotel(db *sqlx.DB,hotel hotelmodel.Hotel) error{
	_,err := db.Exec("INSERT INTO hotels(hotel_id,name,location,description,img,photos) VALUES  (default,$1,$2,$3,$4,$5)",
			hotel.Name,hotel.Location,hotel.Description,hotel.Image,pq.Array(hotel.Photos))
	if err != nil{
			return err
		}

	return nil
}

func UploadImage(filemanager *s3.S3,url string ) (string,error){
	resp,err := http.Get(url)
	if err != nil{
		return "",err
	}
	defer resp.Body.Close()

	filename := uuid.NewV4().String()
	fileType := "jpg"
	relPath := configs.StaticPath + filename + "." + fileType

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "",err
	}
	file := bytes.NewReader(body)


	_, err = filemanager.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(configs.BucketName),
		Key:    aws.String(relPath),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return "", customerror.NewCustomError(err, http.StatusInternalServerError, 1)
	}

	fmt.Println("Success!")
	return relPath,nil
}