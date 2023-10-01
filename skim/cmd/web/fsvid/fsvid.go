package fsvid

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FSVid struct {
	VPath          string `json:"vPath"`
	VBasePath      string `json:"vBasePath"`
	Script         string `json:"script"`
	ScriptBasePath string `json:"scriptBasePath"`
	ScriptPath     string `json:"scriptPath"`
	LastModified   string `json:"lastModified"`
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

	if err = fsv.UpdateModTime(); err != nil {
		return err
	}

	return nil
}

func (fsv *FSVid) EditScript() error {
	fmt.Println("editing fsvid.script")
	fmt.Printf("%s\n", fsv.Script)

	fileByetes, err := os.ReadFile(fsv.ScriptPath)
	if err != nil {
		return fmt.Errorf("%v:%v", editFileErr, err)
	}

	if string(fileByetes) == fsv.Script {
		return fmt.Errorf("%v:%s", editFileErr, "script is the same")
	}

	file, err := os.OpenFile(fsv.ScriptPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("%v:%v", editFileErr, err)
	}
	defer file.Close()

	_, err = file.WriteString(fsv.Script)
	if err != nil {
		return fmt.Errorf("%v:%v", editFileErr, err)
	}

	return nil
}

func (fsv *FSVid) UpdateModTime() error {
	fInfo, err := os.Stat(fsv.VPath)
	if err != nil {
		return err
	}

	fsv.LastModified = fInfo.ModTime().Format(time.RFC3339)
	return nil
}
