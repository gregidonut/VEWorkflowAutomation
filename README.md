# VEWorkflowAutomation

An automation project where we aim to be able to create a full generator
of A.I. generated content.

The final product must be able to produce a non-copywritable video consisting
of the same content as an original movie with less than 5 second cuts of and
ai generated voice over synopsis.

## Break Down af actual flow:

1. generate audio synopsis mp3
2. cut movie to length of audio synopsis
    1. remove audio from move
    2. make sure that the video cuts are less than 5 seconds(to avoid being
       copyright striked)
    3. make sure that length of output video is the same as the length of
       audio synopsis
3. stitch generated audio synopsys and mp3 into one mp4 file ready for
   production

## Solved problems:

- step 1: generated from a LLMs and ALMs
- step 2.1: solved by ffmpeg
    ```zsh
    ffmpeg -i <moviefilename>.mp4 -an -c:v copy <outputfilename>.mp4
    ```
- step 3: solved by ffmpeg(not implemented)

## problem1

__cutting video takes too much time for the app to be viable__

### proposal1(currently being implemented as [Skim](https://github.com/gregidonut/VEWorkflowAutomation/tree/main/skim))

should have a source of the 5 second clips as mp4 files in the correct order.
This way we can programmatically stitch the in a sequence using filename numbers

### proposal2(IMPLEMENTED)

manually watching the movie and and listening to the audio file bit by bit.