# Improper Quotes Monitor

This is a golang script that helps you to detect privilege escalation vulnerabilities on Windows. But it also helps to troubleshoot misterious `'C:/Program' is not recognized as an internal or external command, operable program or batch file.` errors. 

Whenever a file under `C:\Program Files\` is accessed without actually quoting the path, windows attempt to execute everything that's before that path, for example:

You want to run `C:\Program Files\My Epic Program\bin files\extras\ffmpeg.exe C:\Users\username\Download\never-gonna-give-you-up.mp4 C:\Users\username\Download\never-gonna-give-you-up_AVI.avi` in the cmd.
Windows will try to run any of these, in this order:
- C:\Program.exe
- C:\Program Files\My.exe
- C:\Program Files\My Epic Program\bin.exe
- C:\Program Files\My Epic Program\bin files\extras\ffmpeg.exe
Whichever it finds first, it will run. When running a command similarly but from another program, windows will just error out saying `'C:/Program' is not recognized as an internal or external command, operable program or batch file.` because you didn't include quotes in the path of your command.

This project exploits this by creating an executable and placing it under C:\Program.exe and logging everything that goes after that as a parameter.

Without Improper-Quotes-Monitor:
```
C:\Program Files\My Epic Program\bin files\extras\>ffmpeg.exe C:\Users\username\Download\never-gonna-give-you-up.mp4 C:\Users\username\Download\never-gonna-give-you-up_AVI.avi

Processing...
'C:/Program' is not recognized as an internal or external command, operable program or batch file.
```

With Improper-Quotes-Monitor:
```
C:\Program Files\My Epic Program\bin files\extras\>ffmpeg.exe C:\Users\username\Download\never-gonna-give-you-up.mp4 C:\Users\username\Download\never-gonna-give-you-up_AVI.avi

Processing...
[Files\My Epic Program\bin files\extras\intel_decoder.exe --threads 10 --convert mp4:avi --output C:\Users\username\Download\never-gonna-give-you-up_AVI.avi]
```

(The above is just an example. ffmpeg doesn't actually have this vulnerability/bug)

**How to use**:     
1 - Replace `C:\\Path\\to\\Improper Quotes Monitor\\` in `Program.go` with the path of where you've cloned this repository.    
2 - Replace `C:\Path\to\Improper Quotes Monitor` in `alertbox.vbs` with the path of where you've cloned this repository.    
3 - `go build Program.go`    
4 - `mv Program.exe C:\Program.exe`    
5 - Run the program you want to hack/troubleshoot
6 - Check `access-log.txt` in the path of where you've cloned this repository.    

Your antivirus will probably flag Program.exe as being malware, as this is a common filepath for malware. This is obviusly not malware, the source code is right here and you'll compile it yourself. 
