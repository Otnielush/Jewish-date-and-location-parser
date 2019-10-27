package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "github.com/kelvins/sunrisesunset"
	"github.com/mkrou/geonames"
	"github.com/mkrou/geonames/models"
)

// Parse Jew Date

type DateJew struct {
	Gd       int
	Gm       int
	Gy       int
	Hd       int
	Hmonth   int
	Hy       int
	Hm       string
	Location location
}

type location struct {
}

// "https://www.hebcal.com/converter/?cfg=json&gy=2011&gm=6&gd=2&g2h=1"
func (d *DateJew) parseDate(day, month, year int) {
	resp, err := http.Get(fmt.Sprintf("https://www.hebcal.com/converter/?cfg=json&gy=%d&gm=%d&gd=%d&g2h=1", year, month, day))
	if err != nil {
		p("Date not got")
	}

	b2, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(b2, d)
	d.Hmonth = hMonth[d.Hm]
	p(d)

}

var hMonth = map[string]int{"Nisan": 1, "Iyyar": 2, "Sivan": 3, "Tamuz": 4, "Av": 5, "Elul": 6, "Tishrei": 7, "Cheshvan": 8, "Kislev": 9, "Tevet": 10, "Shvat": 11, "Adar1": 12, "Adar2": 13}

// Parsing geonames

type geoCities struct {
	Name      string
	Id        int
	Latitude  float64
	Longitude float64
	Tzone     string
}

// type Cities struct {
// 	CitiName map[string]int
// 	Info     [50000]geoCities
// }

func geoParse(mass *[50000]geoCities, m *map[string]int) {
	gg := geonames.NewParser()

	var i = int(0)
	//print all cities with a population greater than 5000
	err := gg.GetGeonames(geonames.Cities5000, func(geoname *models.Geoname) error {
		(*m)[geoname.AsciiName] = i
		(*mass)[i].Name = geoname.AsciiName
		(*mass)[i].Id = geoname.Id
		(*mass)[i].Latitude = geoname.Latitude
		(*mass)[i].Longitude = geoname.Longitude
		(*mass)[i].Tzone = geoname.Timezone
		i++
		// fmt.Println(geoname.Name, geoname.Id, geoname.Latitude, geoname.Longitude, geoname.Timezone)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Parse Time
