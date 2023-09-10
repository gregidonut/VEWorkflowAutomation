# Skim

Is a web app that outputs "AI-Generated" video synopsis from movies.

## Business logic(TODO: IMPLEMENT)

1. on a gui, there is a button that browses the filesystem for *supported* filetypes that
   contains full movie content
    - [x] decide on gui technology
        - html, css & javascript
    - [x] establish the button that interfaces with os filesystem
        - webapp right now has two buttons, as *html-form-input-elements*
            - input-element-1: `choose file`
            - input-element-2: `submit`
2. the movie is chopped into 5 second video-files that will be available for browsing later
    - [x] add logic that executes the `ffmpeg` command to **remove audio** from the video file
    - [x] add logic that executes the `ffmpeg` command to **split** the video file
      uploaded into several 1 second chunks
3. a list of the movie files in correct sequential order is available(styles to make it intuitive
   is key here) from the gui  
    - [x] loop through the files in the uploaded video-split directory and display them in gui
        - [x] introduce templates for looping syntax
        - [x] introduce static files to be able to loop through available files securely
4. select the 5-second videos to be included in the output  
   **CURRENTLY BEING IMPLEMENTED**:
    - [ ] implement a system using javascript & css that highlights selected list items for a 
    5-second vid
    - [ ] a button that determines whether the user is ready to submit the selected list items
5. each time you select a video you are prompted to type in a synopsis of the 5-second video.
6. from here we branch out into several output-specifications(TODO: IMPLEMENT)

## Movie file

- **supported file types**
    - mkv
    - mp4
- **this is a very important input**  
  the business logic for the app will be expected to use this to produce many 5 second
  videos. 