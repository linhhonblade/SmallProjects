package productdomain

import "errors"

var (
	ErrProductNameRequired = errors.New("Product name cannot be blank.")
)
