package main

import (
	"helpers"
	"io"
	"os"
	"os/exec"
)

func main() {
	name := helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/", "startup")
	helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/", "EncryptionInfo")
	helpers.WriteLog(name, "Strating test : StartupFolderNewFile", 2)
	helpers.CreateTestFiles("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/", name)

	compileFile := exec.Command("go", "build", ".")
	compileFile.Dir = "C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/encr"

	err := compileFile.Run()
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(2)
	}
	src, err := os.Open("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/StartupFolderNewFile/encr/enc.exe")
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(3)
	}
	dest, err := os.Create("C:/Users/achon/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/a.exe")
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(4)
	}
	_, err = io.Copy(dest, src)
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(5)
	}

	_, err = os.Open("C:/Users/achon/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/a.exe")
	if err != nil {
		helpers.WriteLog(name, err.Error(), 1)
		os.Exit(1)
	}

	helpers.WriteLog(name, "Ending test: StartupFolderNewFile", 2)
	os.Exit(0)
}

//aaa
