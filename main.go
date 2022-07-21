package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const URL1 string = "http://titularizare.edu.ro/generated/files/j/IF/pct_d/page_5.html"
const URL2 string = "https://www.hostinger.com/tutorials/website/how-to-inspect-and-change-style-using-google-chrome"
const URL_GEOCODING = "https://google-maps-geocoding.p.rapidapi.com/geocode/json?address='164 Townsend St., San Francisco, CA'&language=en"
const X_RapidAPI_Key = "ea6b94a296msh310d4291ef822dep1507e4jsnc6cfcf2c00dd"
const X_RapidAPI_Hos = "google-maps-geocoding.p.rapidapi.com"

type Scoala struct {
	Cod_post              string
	Localitate            string
	Unitate_de_invatamant string
}

func Table(URL string) Scoala {
	rows := FirstRow(URL) // Prima linie
	if rows == nil {
		return Scoala{}
	}
	scoala := Scoala{}
	rows = rows.FirstChild.NextSibling
	scoala.Cod_post = rows.FirstChild.Data
	rows = rows.NextSibling.NextSibling
	scoala.Localitate = rows.FirstChild.Data
	rows = rows.NextSibling.NextSibling
	scoala.Unitate_de_invatamant = strings.TrimSpace(rows.FirstChild.Data)

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
	NextRow := r1.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling
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
