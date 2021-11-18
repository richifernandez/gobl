package bill

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/invopop/gobl/num"
	"github.com/invopop/gobl/org"
	"github.com/invopop/gobl/tax"
)

// Lines holds an array of Line objects.
type Lines []*Line

// Line is a single row in an invoice.
type Line struct {
	UUID     string        `json:"uuid,omitempty"`
	Index    int           `json:"i" jsonschema:"title=Index,description=Line number inside the invoice."`
	Quantity num.Amount    `json:"quantity"`
	Item     *org.Item     `json:"item"`
	Sum      num.Amount    `json:"sum" jsonschema:"title=Sum,description=Result of quantity multiplied by item price"`
	Discount *org.Discount `json:"discount,omitempty" jsonschema:"title=Discount,description=Discount applied to this line."`
	Taxes    tax.Rates     `json:"taxes,omitempty" jsonschema:"title=Taxes,description=List of taxes to be applied to the line in the invoice totals."`
	Total    num.Amount    `json:"total" jsonschema:"title=Total,description=Total line amount after applying discounts to the sum."`
}

// GetTaxRates responds with the array of tax rates applied to this line.
func (l *Line) GetTaxRates() tax.Rates {
	return l.Taxes
}

// GetTotal provides the final total for this line, excluding any tax calculations.
func (l *Line) GetTotal() num.Amount {
	return l.Total
}

// Validate ensures the line contains everything required.
func (l *Line) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Index, validation.Required),
		validation.Field(&l.Quantity, validation.Required),
		validation.Field(&l.Item, validation.Required),
		validation.Field(&l.Taxes),
		validation.Field(&l.Sum, validation.Required),
		validation.Field(&l.Total, validation.Required),
	)
}

// calculate figures out the totals according to quantity and discounts.
func (l *Line) calculate() error {
	// First we figure out how much the item costs, and get the total
	l.Sum = l.Item.Price.Multiply(l.Quantity)

	if l.Discount != nil {
		d := l.Discount
		if d.Rate != nil && !d.Rate.IsZero() {
			// always override value with calculated rate
			d.Value = d.Rate.Of(l.Sum)
		}
		l.Total = l.Sum.Subtract(d.Value)
	} else {
		l.Total = l.Sum
	}

	return nil
}