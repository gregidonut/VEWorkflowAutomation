# Skim

Is a web app that outputs "AI-Generated" video synopsis from movies. 

## Business logic(TODO: IMPLEMENT)
1. on a gui, there is a button that browses the filesystem for *supported* filetypes that 
contains full movie content  
  **CURRENTLY BEING IMPLEMENTED**: 
    - [x] decide on gui technology
      - html, css & javascript  
    - [x] establish the button that interfaces with os filesystem
      - webapp right now has two buttons, as *html-form-input-elements* 
        - input-element-1: `choose file`
        - input-element-2: `submit`
    - [ ] research into uploading a file from a form in a post method request
2. the movie is chopped into 5 second video-files that will be available for browsing later
3. a list of the movie files in correct sequential order is available(styles to make it intuitive 
is key here) from the gui
4. select the 5-second videos to be included in the output
5. each time you select a video you are prompted to type in a synopsis of the 5 second video.
6. from here we branch out into several output-specifications(TODO: IMPLEMENT)

## Movie file
- **supported file types**
  - mkv
  - mp4
- **this is a very important input**  
the business logic for the app will be expected to use this to produce many 5 second
videos. 