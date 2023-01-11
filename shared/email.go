package shared

import (
	"emailQue/model"

	"github.com/shopspring/decimal"
	"github.com/vanng822/go-premailer/premailer"
	"jaytaylor.com/html2text"
)

// Months ...
var Months = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

// HTMLToEMail converts html to email compatible html and text format
func HTMLToEMail(b []byte) (*model.EMailMsg, error) {
	prem, err := premailer.NewPremailerFromBytes(b, premailer.NewOptions())
	if err != nil {
		return nil, err
	}

	eml := model.EMailMsg{}

	// inline css
	eml.HTML, err = prem.Transform()
	if err != nil {
		return nil, err
	}

	// convert html to text
	eml.Text, err = html2text.FromString(eml.HTML, html2text.Options{PrettyTables: false})
	if err != nil {
		return nil, err
	}

	return &eml, nil
}

type invDetail struct {
	Due    string          `json:"due"`
	Amount decimal.Decimal `json:"amount"`
}
