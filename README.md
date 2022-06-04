# Improper Quotes Monitor

This is a golang script that helps you to detect privilege escalation vulnerabilities on Windows. But it also helps to troubleshoot misterious `'C:/Program' is not recognized as an internal or external command, operable program or batch file.` errors. 

**How to use**:     
1 - Replace `C:\\Path\\to\\Improper Quotes Monitor\\` in `Program.go` with the path of where you've cloned this repository.    
2 - Replace `C:\Path\to\Improper Quotes Monitor` in `alertbox.vbs` with the path of where you've cloned this repository.    
3 - `go build Program.go`    
4 - `mv Program.exe C:\Program.exe`    
5 - Run the program you want to hack/troubleshoot
6 - Check `access-log.txt` in the path of where you've cloned this repository.    

Your antivirus will probably flag Program.exe as being malware, as this is a common filepath for malware. This is obviusly not malware, the source code is right here and you'll compile it yourself. 
