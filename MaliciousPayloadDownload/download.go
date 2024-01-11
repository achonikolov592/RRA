package main

import (
	"helpers"
	"io"
	"net/http"
	"os"
)

var nameOfLogFile string

func main() {
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/MaliciousPayloadDownload/", "DownloadMalicious")

	helpers.WriteLog(nameOfLogFile, "Starting test: DownloadingMaliciousPayload", 1)

	out, err := os.Create("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/MaliciousPayloadDownload/out1.zip")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(2)
	}
	defer out.Close()

	resp, err := http.Get("https://github.com/kh4sh3i/Ransomware-Samples/raw/main/Petya/Ransomware.Petya.zip")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(3)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(4)
	}

	_, err = os.Open("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/MaliciousPayloadDownload/out1.zip")
	if err != nil {
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
		os.Exit(1)
	}

	helpers.WriteLog(nameOfLogFile, "Ending test: DownloadingMaliciousPayload", 1)

	os.Exit(0)
}

//a
