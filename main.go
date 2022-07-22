package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/net/html"
)

const URL1 string = "http://titularizare.edu.ro/generated/files/j/IF/pct_d/page_5.html"
const URL2 string = "http://titularizare.edu.ro/generated/files/j/IF/pct_d/page_6.html"
const URL3 string = "http://titularizare.edu.ro/generated/files/j/IF/pct_d/page_7.html"
const URL4 string = "http://titularizare.edu.ro/generated/files/j/B/pct_d/page_11.html"
const URL5 string = "http://titularizare.edu.ro/generated/files/j/B/pct_d/page_12.html"
const HOME string = "Str. Sfânta Agnes 21, Popești-Leordeni 077160"
const FACULTATE string = "Șoseaua Panduri 90, București 050663"

// const POLI_LEU string =
const URL_GEOCODING = "https://google-maps-geocoding.p.rapidapi.com/geocode/json?address='164 Townsend St., San Francisco, CA'&language=en"
const X_RapidAPI_Hos = "google-maps-geocoding.p.rapidapi.com"
const TIME = 1658725200 // 2022-07-25 08:00:00
const URL_TEST = "https://reqres.in/api/users?page=2"

type Scoala struct {
	Cod_post              int // 0
	Localitate            int // 1
	Unitate_de_invatamant int // 2
	Nivel_de_invataman    int // 3
	Disciplina            int // 4
	Nr_ore_trunchi_comun  int // 5
	Detalii               int // 6
	Valabilitate          int // 7
	DistanceText          int // 8
	DistanceValue         int // 9
	DurationText          int // 10
	DurationValue         int // 11
	Status                int // 12
}

// Post_complet          int // 8
// Proba_practica        int // 9
// Proba_de_limba        int // 10

var den Scoala = Scoala{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func Table(URL string, linie_de_start, linie_de_final int) [][]string {
	rows := FirstRow(URL) // Prima linie
	if rows == nil {
		return [][]string{}
	}

	scoala := [][]string{}
	poz := 0
	for lin := 0; lin < linie_de_final; lin++ {
		rows = rows.NextSibling.NextSibling
		if lin < linie_de_start {
			continue
		}

		scoala = append(scoala, []string{})
		row := rows.FirstChild.NextSibling // Primul camp
		for i := 0; i < 8; i++ {
			scoala[poz] = append(scoala[poz], strings.TrimSpace(row.FirstChild.Data)) // Valoarea dintr-un camp
			row = row.NextSibling.NextSibling
		}
		poz++
	}
	return scoala
}

func FirstRow(URL string) *html.Node {
	resp, err := http.Get(URL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	bodyString := string(bodyBytes)

	doc, err := html.Parse(strings.NewReader(bodyString))
	if err != nil {
		return nil
	}

	r1 := GetElementById(doc, "mainTable")
	if err != nil {
		return nil
	}

	NextRow := r1.FirstChild.NextSibling.FirstChild
	return NextRow
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

}
