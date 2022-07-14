package nl

import (
	"errors"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/invopop/gobl/l10n"
	"github.com/invopop/gobl/org"
)

// ValidTaxID complies with the ozzo validation Rule definition to be able
// to confirm that the Tax ID is indeed spanish and valid.
var ValidTaxID = validTaxID{}

type validTaxID struct {
	requireCode bool
}

// RequireCode allows for additional checks for the ID code
func (v validTaxID) RequireCode() validTaxID {
	return validTaxID{requireCode: true}
}

const (
	vatLen = 12
)

var errInvalidVAT = errors.New("invalid VAT number")

// VerifyTaxCode looks at the provided code, determines the type, and performs
// the calculations required to determine if it is valid.
// These methods assume the code has already been cleaned and only
// contains upper-case letters and numbers.
func VerifyTaxCode(code string) error {
	code = strings.ToUpper(code)
	if len(code) != vatLen {
		return errInvalidVAT
	}
	if code[9] != 'B' {
		return errInvalidVAT
	}
	return validateDigits(code[0:9], code[10:12])
}

// Validate ensures the tax ID contains a matching country and
// valid code.
func (v validTaxID) Validate(value interface{}) error {
	id, ok := value.(*org.TaxIdentity)
	if !ok {
		return nil
	}
	return validation.ValidateStruct(id,
		validation.Field(&id.Country, validation.Required, validation.In(l10n.NL)),
		validation.Field(&id.Code,
			validation.When(v.requireCode, validation.Required),
			validation.By(validateTaxCode),
		),
	)
}

func validateTaxCode(value interface{}) error {
	code, ok := value.(string)
	if !ok {
		return nil
	}
	if code == "" {
		return nil
	}
	return VerifyTaxCode(code)
}

func validateDigits(code, check string) error {
	num, err := strconv.ParseInt(code, 10, 64)
	if err != nil {
		return errInvalidVAT
	}
	_, err = strconv.Atoi(check)
	if err != nil {
		return errInvalidVAT
	}
	var sum int64
	ck := num % 10
	for i := 0; i < 8; i++ {
		num /= 10
		mul := int64(i) + 2
		sum += (num % 10) * mul
	}
	sum = sum % 11
	if sum > 9 {
		sum = 0
	}
	if sum != ck {
		return errors.New("checksum mismatch")
	}

	return nil
}
