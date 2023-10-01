package fsvid

import "os"

func (fsv *FSVid) Delete() error {
	if err := os.Remove(fsv.VPath); err != nil {
		return err
	}

	return nil
}
