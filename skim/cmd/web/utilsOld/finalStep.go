package utilsOld

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"os"
	"path/filepath"
)

func FinalStep() error {
	fmt.Println("running final step script...")

	finalVidDir := filepath.Dir(paths.FINAL_VID_PATH)
	if _, err := os.Stat(finalVidDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(finalVidDir, os.ModeDir|os.ModePerm); err != nil {
				return err
			}
		}
	}

	fmt.Println("finished running final step script...")
	return nil
}
