package main

import (
	"RRA/EncryptDecryptDirRecursivePartially/decrypt"
	"RRA/EncryptDecryptDirRecursivePartially/encrypt"
	"RRA/SecureDeleteFiles/SecureDeleteFile"
	"helpers"
	"os"
	"strings"
)

var nameOfLogFile string

func main() {
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/", "EncryptDecryptDirRecursivePartially")
	helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/", "EncryptionInfo")

	if strings.ToLower(os.Args[1]) == "true" {

		_, err := os.Stat("./EncryptionInfo.log")
		if err == nil {
			SecureDeleteFile.SecureDelete("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/EncryptionInfo.log", nameOfLogFile)
		}

		nameOfInfoFile := helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/", "EncryptionInfo")
		helpers.RemoveTestFilesIfExists("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/")
		helpers.CreateTestFiles("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/", nameOfLogFile)

		helpers.WriteLog(nameOfLogFile, "Strating test: EncryptDirPartially", 2)

		encrypt.EncryptFilesInDir("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/testfiles", nameOfLogFile, nameOfInfoFile)

		helpers.WriteLog(nameOfLogFile, "Ending test: EncryptDecryptDirRecursivePartially", 2)
	}
	if strings.ToLower(os.Args[2]) == "true" {
		helpers.WriteLog(nameOfLogFile, "Strating test: DecryptDirRecursivePartially", 2)

		decrypt.DecryptDir("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/testfiles", nameOfLogFile, "C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursivePartially/EncryptionInfo.log")

		helpers.WriteLog(nameOfLogFile, "Ending test: DecryptDirRecursivePartially", 2)
	}

	helpers.WriteLog(nameOfLogFile, "Endinging test: EncryptDecryptDirPartially", 2)
	os.Exit(0)
}

//
