# base64-to-clipboard
Decodes an input string from base64 and then saves the result to the clipboard.
# Setup
Download the `clip64` binary to your system, or alternatively to compile the app yourself, clone this repository and run `go build clip64.go`. Then, move the outputted `clip64` binary to your system path to access it via the terminal. For MacOS users, this means moving the `clip64` binary to `/usr/local/bin`.
# Usage
Run the binary with an input in base64 form, for example `clip64 SGVsbG8gd29ybGQhIQ==`. The decoded output will be printed to the terminal as well as copied to the clipboard.
# Possible future features
`-q` flag to run **quietly** (non verbose, no terminal output)  
`-l` flag to specify the input string is a **link** and to launch it in the systems default browser  
`-e` flag to **encode** the input instead of decode
