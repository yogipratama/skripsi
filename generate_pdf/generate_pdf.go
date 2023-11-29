package generate_pdf

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/signintech/gopdf"
	"github.com/skip2/go-qrcode"
)

func GeneratePDF(pdfBytes []byte, path, expired_at string) {
	rs := io.ReadSeeker(bytes.NewReader(pdfBytes))

	png, err := qrcode.Encode("https://facebook.com", qrcode.Medium, 256)
	if err != nil {
		log.Print(err.Error())
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeLetter})

	tmpl := pdf.ImportPageStream(&rs, 1, "/MediaBox")

	pdf.AddPage()

	err = pdf.AddTTFFont("arial", "arial.ttf")
	if err != nil {
		log.Print(err.Error())
	}

	err = pdf.SetFont("arial", "", 10)
	if err != nil {
		log.Print(err.Error())
	}

	qrcode, err := gopdf.ImageHolderByBytes(png)
	if err != nil {
		log.Print(err.Error())
	}

	pdf.ImageByHolder(qrcode, 30, 550, nil)

	pdf.SetX(30)
	pdf.SetY(745)
	pdf.Cell(nil, fmt.Sprintf("*dokumen ini berlaku sampai tanggal %v", expired_at))

	pdf.UseImportedTemplate(tmpl, 0, 0, 0, 0)

	pdf.WritePdf(path)
	pdf.Close()
}
