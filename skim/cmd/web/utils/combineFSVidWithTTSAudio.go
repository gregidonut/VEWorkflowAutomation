package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func CombineFSVidWithTTSAudio() error {
	_, err := os.Stat(paths.FSVIDS_REL_PATH)
	if os.IsNotExist(err) {
		os.Mkdir(paths.FSVIDS_REL_PATH, os.ModeDir|os.ModePerm)
	}

	files, err := os.ReadDir(paths.RAW_COMMIT_VIDS_REL_PATH)
	if err != nil {
		return err
	}

	var fileNames []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.Contains(f.Name(), "txt") {
			continue
		}
		fileNames = append(fileNames, filepath.Base(f.Name()))
	}
	sort.Strings(fileNames)

	lastVidFileMP4 := fileNames[len(fileNames)-1]
	if err = GenerateTTS(lastVidFileMP4); err != nil {
		return err
	}

	CombineFSVidWithTTSCmd := exec.Command(
		"ffmpeg",
		"-i",
		fmt.Sprintf("%s/%s", paths.RAW_COMMIT_VIDS_REL_PATH, lastVidFileMP4),
		"-i",
		fmt.Sprintf("%s/%s.mp3", paths.RAW_COMMIT_VIDS_REL_PATH, strings.TrimSuffix(lastVidFileMP4, ".mp4")),
		"-c:v",
		"copy",
		"-c:a",
		"aac",
		fmt.Sprintf("%s/%s", paths.FSVIDS_REL_PATH, lastVidFileMP4),
	)

	err = RunCmd(CombineFSVidWithTTSCmd, ".")
	if err != nil {
		return err
	}

	return nil
}

func GenerateTTS(lastVidFileMP4 string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := polly.New(sess)

	lastVidBaseName := strings.TrimSuffix(lastVidFileMP4, ".mp4")
	scriptPath := fmt.Sprintf("%s/%s.txt", paths.RAW_COMMIT_VIDS_REL_PATH, lastVidBaseName)
	textBytes, err := os.ReadFile(scriptPath)
	if err != nil {
		fmt.Printf("%v:%v\n", generateTTSErr, err)
		return fmt.Errorf("%v:%v", generateTTSErr, err)
	}

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		Text:         aws.String(string(textBytes)),
		VoiceId:      aws.String("Matthew"),
	}

	output, err := svc.SynthesizeSpeech(input)
	if err != nil {
		fmt.Printf("%v:%v\n", generateTTSErr, err)
		return fmt.Errorf("%v:%v", generateTTSErr, err)
	}

	audioStream := output.AudioStream

	audioPath := fmt.Sprintf("%s/%s.mp3", paths.RAW_COMMIT_VIDS_REL_PATH, lastVidBaseName)
	err = saveAudioToFile(audioStream, audioPath)
	if err != nil {
		fmt.Printf("%v:%v\n", generateTTSErr, err)
		return fmt.Errorf("%v:%v", generateTTSErr, err)
	}
	return nil
}

func saveAudioToFile(audioStream io.ReadCloser, filename string) error {
	defer audioStream.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, audioStream)
	if err != nil {
		return err
	}

	return nil
}
