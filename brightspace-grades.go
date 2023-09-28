package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/anaskhan96/soup"
)

var ethics = "38684"
var datasci = "42431"
var opsys = "40259"

var client = &http.Client{}

func main() {
	fmt.Println(getGradesTable(opsys))
}

func getGradesTable(course string) string {
        req, err := http.NewRequest("GET", "https://wentworth.brightspace.com/d2l/lms/grades/my_grades/main.d2l?ou="+course, nil)
        if err != nil {
                panic(err)
        }
        req.Header.Set("Cookie", os.Getenv("BRIGHTSPACE_COOKIE"))

        res, err := client.Do(req)
        if err != nil {
                panic(err)
        }
        defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	doc := soup.HTMLParse(string(body))
	labels := doc.FindAll("label")
	gradeTable := ""
	for _, label := range labels {
		gradeTable += label.Text()
	}
        return gradeTable
}

