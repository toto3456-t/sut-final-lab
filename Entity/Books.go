package entity

import (
	"github.com/asaskevich/govalidator"
)

type Books struct {
	Title string  `"valid:stringlength(3|100)~Title must be between 3 and 100 charater"`
	Price float64 `"valid:ragne(50|5000)~Price must be between 50 and 5000"`
	Code  string  `"valid:required~Code must start with BK followed by 6 digits (0-9),^([BM][0-9]{6}$)"`
}

func (b *Books) ValidateBooks() error {

	ok, err := govalidator.ValidateStruct(b)

	if !ok {
		return err
	}
	return nil
}
