# base64-to-clipboard
Decodes an input string from base64 and then saves the result to the clipboard.
# Setup
Download the most recent release of the `clip64` binary to your system from https://github.com/charlesdevelops/clip64/releases, or alternatively to compile the app yourself, clone this repository and run `go build clip64.go`. Then, move the outputted `clip64` binary to your system path to access it via the terminal. For MacOS users, this means moving the `clip64` binary to `/usr/local/bin`.
# Usage
Run the binary with an input in base64 form, for example `clip64 SGVsbG8gd29ybGQhIQ==`. The decoded output will be printed to the terminal as well as copied to the clipboard.  
<img width="682" alt="helloWorld" src="https://github.com/charlesdevelops/clip64/assets/107790118/8b1133e8-ca84-4e91-80bc-d478178c71ae">
# Possible future features
`-q` flag to run **quietly** (non verbose, no terminal output)  
`-l` flag to specify the input string is a **link** and to launch it in the systems default browser  
`-e` flag to **encode** the input instead of decode
