package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type myFile struct {
	Size int64
	Name string
}

func printHeader() {
	const header string = `██╗  ██╗ ██████╗██╗         ██████╗ ██████╗ ███╗   ██╗██╗   ██╗███████╗██████╗ ████████╗███████╗██████╗ 
╚██╗██╔╝██╔════╝██║        ██╔════╝██╔═══██╗████╗  ██║██║   ██║██╔════╝██╔══██╗╚══██╔══╝██╔════╝██╔══██╗
 ╚███╔╝ ██║     ██║        ██║     ██║   ██║██╔██╗ ██║██║   ██║█████╗  ██████╔╝   ██║   █████╗  ██████╔╝
 ██╔██╗ ██║     ██║        ██║     ██║   ██║██║╚██╗██║╚██╗ ██╔╝██╔══╝  ██╔══██╗   ██║   ██╔══╝  ██╔══██╗
██╔╝ ██╗╚██████╗██║███████╗╚██████╗╚██████╔╝██║ ╚████║ ╚████╔╝ ███████╗██║  ██║   ██║   ███████╗██║  ██║
╚═╝  ╚═╝ ╚═════╝╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

`

	color.Green(header)
}

func printUsage() {
	color.Red("Usage : ./XCI_Converter [.XCI FILE] [TITLE ID]")
}

func getBiggestNCA(titleIdName string, xciPath string) (file string) {
	var (
		nca myFile
		cmd *exec.Cmd
	)

	color.Green("Decrypting .xci's NCA files and finding the biggest NCA...")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("./hactool.exe", "-k", "keys.ini", "-txci", "--securedir="+titleIdName, xciPath)
	} else {
		cmd = exec.Command("./hactool", "-k", "keys.ini", "-txci", "--securedir="+titleIdName, xciPath)
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(titleIdName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Size() > nca.Size {
			nca.Size = file.Size()
			nca.Name = file.Name()
		}
	}

	return nca.Name
}

func decryptNCA(ncaFile string, titleIdName string) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("./hactool.exe", "-k", "keys.ini", "--romfs="+titleIdName+"/romfs.bin", "--exefsdir="+titleIdName+"/exefs", titleIdName+"/"+ncaFile)
	} else {
		cmd = exec.Command("./hactool", "-k", "keys.ini", "--romfs="+titleIdName+"/romfs.bin", "--exefsdir="+titleIdName+"/exefs", titleIdName+"/"+ncaFile)
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(titleIdName)
	if err != nil {
		log.Fatal(err)
	}

	// delete *.nca
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".nca" {
			err := os.Remove(titleIdName + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func patchMainNPDM(titleIdName string) {
	color.Green("Patching main.npdm :)")

	fd, err := os.OpenFile(titleIdName+"/exefs/main.npdm", os.O_RDWR, 0000)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	uint64TitleName, err := strconv.ParseUint("0x"+titleIdName, 0, 64)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = fd.Seek(0x440, os.SEEK_CUR)
	if err != nil {
		log.Fatal(err)
	}

	err = binary.Write(fd, binary.LittleEndian, uint64TitleName)
	if err != nil {
		log.Fatal(err)
	}
}

func isHex(hexa string) bool {
	_, err := strconv.ParseUint(hexa, 16, 64)
	if err != nil {
		return false
	}
	return true
}

func isValidFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}

	// Check is regular file
	if file.Mode().IsRegular() == false {
		return false
	}

	// Check permission (at least rw)
	mode := file.Mode().Perm()
	if (mode >> 7) != 3 {
		return false
	}

	return true
}

func isValidArgs(titleIDName string, xciPath string) bool {
	if isHex(titleIDName) == false {
		log.Fatal("Please set a valid titleID")
		return false
	} else if isValidFile(xciPath) == false {
		log.Fatal("Path invalid or insufficient permission")
		return false
	}

	return true
}

func main() {
	var titleIdName string
	var xciPath string

	printHeader()

	if len(os.Args) != 3 {
		printUsage()
		return
	} else {
		titleIdName = os.Args[2]
		xciPath = os.Args[1]

		titleIdName = strings.TrimSpace(titleIdName)
		xciPath = strings.TrimSpace(xciPath)
	}

	if isValidArgs(titleIdName, xciPath) == false {
		log.Fatal("Please set a valid titleID")
	}

	ncaFile := getBiggestNCA(titleIdName, xciPath)

	decryptNCA(ncaFile, titleIdName)

	patchMainNPDM(titleIdName)

	color.Green("DONE! You should have a folder: " + titleIdName)
	color.Green(titleIdName + " should contain an exefs folder and a romfs.bin. It should NOT contain anything else.")
}
