# BGGP 2024 Entry Submission


## Requirements
A valid submission will:
   - Be 4096 bytes or less
   - Download the text file at https://binary.golf/5/5
   - Display the file's contents in some way

## How I Went About At It
I was greatly inspired by airencracken's entry, however, I'd realized that instead of using an http library to get the results, it might've been less characters to instead execute a command on the shell, hence causing us to not import `io` nor use `io.ReadAll(r.Body)`. The URL was next, as that would be half of the file space if I left it in the code, so instead, I opted to import the library `os` and get command line arguments, which saves on space considerably!

I got to 121 characters, which beats the previous 138 character entry! I'm excited to see if anyone beats it!

## File Verification

The program is 121 bytes, because there are 121 characters. I presume the extra two bytes is for some end of line character. I used some websites to get the hash, and I'm hoping that no red flags are thrown!
