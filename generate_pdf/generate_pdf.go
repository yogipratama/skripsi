package generate_pdf

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/signintech/gopdf"
	"github.com/skip2/go-qrcode"
)

func GeneratePDF(pdfBytes []byte, path, uuid, expired_at string) {
	rs := io.ReadSeeker(bytes.NewReader(pdfBytes))

	png, err := qrcode.Encode(fmt.Sprintf("%v/dokumen/%v", os.Getenv("BASE_URL_SERVER"), uuid), qrcode.Medium, 120)
	if err != nil {
		log.Print(err.Error())
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	tmpl := pdf.ImportPageStream(&rs, 1, "/MediaBox")

	pdf.AddPage()

	err = pdf.AddTTFFont("times-new-roman", "times-new-roman.ttf")
	if err != nil {
		log.Print(err.Error())
	}

	err = pdf.SetFont("times-new-roman", "", 6)
	if err != nil {
		log.Print(err.Error())
	}

	qrcode, err := gopdf.ImageHolderByBytes(png)
	if err != nil {
		log.Print(err.Error())
	}

	err = pdf.ImageByHolder(qrcode, 490, 688, nil)
	if err != nil {
		log.Print(err.Error())
	}

	pdf.SetX(460)
	pdf.SetY(830)
	pdf.Cell(nil, fmt.Sprintf("Dokumen ini berlaku sampai tanggal %v", expired_at))

	pdf.UseImportedTemplate(tmpl, 0, 0, 0, 0)

	pdf.WritePdf(path)
	pdf.Close()
}
