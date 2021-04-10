package main

import (
	"fmt"

	u "github.com/dedeyuyandi/go-html-to-pdf/pdf-generator"
)

func main() {

	r := u.NewPDF("")

	//html template path
	templatePath := "template.html"

	//path for download pdf
	outputPath := "storage/output-generate.pdf"

	//html template data
	templateData := struct {
		Title       string
		Description string
		Name        string
		Address     string
	}{
		Title:       "Generate HTML to PDF",
		Description: "HTML to PDF file",
		Name:        "Dede Yuyandi",
		Address:     "Sukajaya, Tanggeung, Kebapaten Cianjur",
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok, "PDF generated successfully", outputPath)
	} else {
		fmt.Println(err)
	}
}
