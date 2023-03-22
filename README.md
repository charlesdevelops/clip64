# base64-to-clipboard
Decodes an input string from base64 and then saves the result to the clipboard.

# Usage
run using `java clip64 <base64 string>`, and the program will output the result to the systems clipboard as well as writing a confirmation to the terminal. Make sure that the clip64 program is in the same directory as the terminal session.

# Upcoming
`-q` flag to run **quietly** (non verbose, no terminal output)  
`-l` flag to specify the input string is a **link** and to launch it in the systems default browser  
`-e` flag to **encode** the input instead of decode
