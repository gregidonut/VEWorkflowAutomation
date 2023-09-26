package fsvid

import "errors"

var (
	instantiationErr = errors.New("error creating fsVid instance")
	marshalErr       = errors.New("having trouble marshalling from kwargs")
	unmarshalErr     = errors.New("having trouble unmarshalling to FSVid type ")
	fillFieldsErr    = errors.New("error filling the rest of the FSVid fields")
)
