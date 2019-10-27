// на вход: год, месяц, день месяца, день недели, час(еврейские),(время рынка или откуда акция).
// Когда создали акцию точная дата по местному времени.
// Широта и долгота места откуда акция
// Выход: цена, макс цена, мин цена

package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {

	var day, month, year int = 19, 4, 1990
	var dayS, monthS, yearS string
	var city string = "Tel Aviv"

	var parseCity, parseDat bool = false, true
	DJ := DateJew{Gd: day, Gm: month, Gy: year} // Ввод даты для перевода на евр
	CityName := make(map[string]int)            // Чтобы найти id города
	var cities [50000]geoCities

	// Локация города
	if parseCity {
		geoParse(&cities, &CityName)
		// idName := CityName["London"]
		p(cities[CityName[city]].Name, cities[CityName[city]].Tzone, cities[CityName[city]].Id)
	}

	// XLSX and Save

	f, err := excelize.OpenFile("../Prices/Crypto/BTCUSD-CoinDesk2210.xlsx")
	if err != nil {
		p("file not opened")
	} else {
		p("file opened")
	}
	// Еврейская дата
	// Заполняем данные в файл
	if parseDat {
		for axisStart := 2; axisStart <= 2210; axisStart++ {

			dayS, _ = f.GetCellValue("table1", fmt.Sprintf("B%d", axisStart))
			monthS, _ = f.GetCellValue("table1", fmt.Sprintf("C%d", axisStart))
			yearS, _ = f.GetCellValue("table1", fmt.Sprintf("D%d", axisStart))
			day, _ = strconv.Atoi(dayS)
			month, _ = strconv.Atoi(monthS)
			year, _ = strconv.Atoi(yearS)
			// получаем Евро дату
			DJ.parseDate(day, month, year)

			// День недели
			f.SetCellValue("table1", fmt.Sprintf("E%d", axisStart), (axisStart%7)+1)
			f.SetCellValue("table1", fmt.Sprintf("U%d", axisStart), float64((axisStart%7)+1)/7)

			// Евро дата
			f.SetCellValue("table1", fmt.Sprintf("F%d", axisStart), DJ.Hd)
			f.SetCellValue("table1", fmt.Sprintf("G%d", axisStart), DJ.Hmonth)
			f.SetCellValue("table1", fmt.Sprintf("H%d", axisStart), DJ.Hy)
			// Данные по дате для нейросети (полнота)
			f.SetCellValue("table1", fmt.Sprintf("O%d", axisStart), (float64(DJ.Hd) / 30))
			f.SetCellValue("table1", fmt.Sprintf("P%d", axisStart), (float64(DJ.Hmonth) / 13))
			f.SetCellValue("table1", fmt.Sprintf("R%d", axisStart), (float64(day) / 31))
			f.SetCellValue("table1", fmt.Sprintf("S%d", axisStart), (float64(month) / 13))

			if axisStart == 1000 {
				err = f.SaveAs("../Prices/Crypto/BTCUSD-CoinDesk_2210.xlsx")
				if err != nil {
					p("file not saved 1000")
				} else {
					p("file saved 1000")
				}
			}
		}

		// Save xlsx file by the given path.
		err = f.SaveAs("../Prices/Crypto/BTCUSD-CoinDesk_2210.xlsx")
		if err != nil {
			p("file not saved")
		} else {
			p("file saved")
		}
	}
}
