package fsvid

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FSVid struct {
	VPath          string `json:"vPath"`
	VBasePath      string `json:"vBasePath"`
	Script         string `json:"script"`
	ScriptBasePath string `json:"scriptBasePath"`
	ScriptPath     string `json:"scriptPath"`
}

func NewFSVid(kwargs map[string]string) (*FSVid, error) {
	//{{ checking if non-optional-keys are present
	nonOptionalKeys := []string{
		"vPath",
		"scriptPath",
	}

	for _, key := range nonOptionalKeys {
		if _, ok := kwargs[key]; !ok {
			return &FSVid{}, fmt.Errorf("%v:%s not provided", instantiationErr, key)
		}
	}
	//}}

	//{{ converting kwargs map to FSVid Type
	payload := new(FSVid)

	jsonData, err := json.Marshal(kwargs)
	if err != nil {
		return &FSVid{}, fmt.Errorf("%v:%v", marshalErr, err)
	}

	err = json.Unmarshal(jsonData, &payload)
	if err != nil {
		return &FSVid{}, fmt.Errorf("%v:%v", unmarshalErr, err)
	}
	//}}

	err = payload.fillRestOfFields()
	if err != nil {
		return &FSVid{}, fmt.Errorf("%v:%v", fillFieldsErr, err)
	}

	return payload, nil
}

// fillRestOfFields registers the rest of the fields used for instantiating
// the FSVid type using the non-optional keys
func (fsv *FSVid) fillRestOfFields() error {
	fsv.VBasePath = filepath.Base(fsv.VPath)
	fsv.ScriptBasePath = filepath.Base(fsv.ScriptPath)

	textBytes, err := os.ReadFile(fsv.ScriptPath)
	if err != nil {
		return err
	}

	fsv.Script = string(textBytes)
	return nil
}
