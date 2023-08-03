package form

import (
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}
//Initialize new form
func New(data url.Values) *Form {
	return &Form {
		data,
		errors(map[string][]string{}),
	}
}
//Check if a field existed in formdata
func (f *Form) Has(field string) bool {
	x:= f.Values.Get(field) 
	f.Errors.Add(field,"This field cannot be blank")
	if x ==""  {
		return false
	}
	return true
}

//Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	
	return len(f.Errors) ==0
}