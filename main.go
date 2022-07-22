package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const URL1 string = "http://titularizare.edu.ro/generated/files/j/IF/pct_d/page_5.html"
const URL2 string = "http://titularizare.edu.ro/generated/files/j/IF/pct_d/page_6.html"
const URL3 string = "http://titularizare.edu.ro/generated/files/j/IF/pct_d/page_7.html"
const URL4 string = "http://titularizare.edu.ro/generated/files/j/B/pct_d/page_11.html"
const URL5 string = "http://titularizare.edu.ro/generated/files/j/B/pct_d/page_12.html"
const HOME string = "Str. Sfânta Agnes 21, Popești-Leordeni 077160"
const URL_GEOCODING = "https://google-maps-geocoding.p.rapidapi.com/geocode/json?address='164 Townsend St., San Francisco, CA'&language=en"
const X_RapidAPI_Hos = "google-maps-geocoding.p.rapidapi.com"

type Scoala struct {
	Cod_post              int // 0
	Localitate            int // 1
	Unitate_de_invatamant int // 2
	Nivel_de_invataman    int // 3
	Disciplina            int // 4
	Nr_ore_trunchi_comun  int // 5
	Detalii               int // 6
	Valabilitate          int // 7
	Post_complet          int // 8
	Proba_practica        int // 9
	Proba_de_limba        int // 10
}

var den Scoala = Scoala{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

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
	//fmt.Println(bodyString)

	// return r1.FirstChild.NextSibling.FirstChild.FirstChild.NextSibling.FirstChild.Data
	NextRow := r1.FirstChild.NextSibling.FirstChild
	return NextRow
}

func response(URL string) string {
	resp, err := http.Get(URL)
	if err != nil {
		return "Eroare server"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "Eroare server"
		}
		bodyString := string(bodyBytes)

		doc, err := html.Parse(strings.NewReader(bodyString))
		if err != nil {
			return ("Fail to parse!")
		}

		r1 := GetElementById(doc, "mainTable")
		if err != nil {
			return "Probleme la getelement"
		}
		//fmt.Println(bodyString)

		// return r1.FirstChild.NextSibling.FirstChild.FirstChild.NextSibling.FirstChild.Data
		NextRow := r1.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling
		return NextRow.FirstChild.NextSibling.FirstChild.Data
	}

	return "status server not ok"
}

func main() {

}
