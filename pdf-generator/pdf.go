package pdfGenerator

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type ReqPdf struct {
	body string
}

//new pdf ...
func NewPDF(body string) *ReqPdf {
	return &ReqPdf{
		body: body,
	}
}

//parsing template ...
func (r *ReqPdf) ParseTemplate(htmlTmp string, data interface{}) error {
	t, err := template.ParseFiles(htmlTmp)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

//generate pdf ...
func (r *ReqPdf) GeneratePDF(pdfPath string) (bool, error) {
	t := time.Now().Unix()
	// write whole the body
	if _, err := os.Stat("clone-template/"); os.IsNotExist(err) {
		errDir := os.Mkdir("clone-template/", 0777)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}
	err1 := ioutil.WriteFile("clone-template/"+strconv.FormatInt(int64(t), 10)+".html", []byte(r.body), 0644)
	if err1 != nil {
		panic(err1)
	}

	f, err := os.Open("clone-template/" + strconv.FormatInt(int64(t), 10) + ".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	// Remove file html
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(dir + "/clone-template")

	return true, nil
}
