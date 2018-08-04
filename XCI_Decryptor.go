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

func getBiggestNCA(titleName string, xciPath string) (file string) {
	var (
		nca myFile
		cmd *exec.Cmd
	)

	color.Green("Decrypting .xci's NCA files and finding the biggest NCA...")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("./hactool.exe", "-k", "keys.ini", "-txci", "--securedir="+titleName, xciPath)
	} else {
		cmd = exec.Command("./hactool", "-k", "keys.ini", "-txci", "--securedir="+titleName, xciPath)
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(titleName)
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

func decryptNCA(ncaFile string, titleName string) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("./hactool.exe", "-k", "keys.ini", "--romfs="+titleName+"/romfs.bin", "--exefsdir="+titleName+"/exefs", titleName+"/"+ncaFile)
	} else {
		cmd = exec.Command("./hactool", "-k", "keys.ini", "--romfs="+titleName+"/romfs.bin", "--exefsdir="+titleName+"/exefs", titleName+"/"+ncaFile)
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(titleName)
	if err != nil {
		log.Fatal(err)
	}

	// delete *.nca
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".nca" {
			err := os.Remove(titleName + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func patchMainNPDM(titleName string) {
	color.Green("Patching main.npdm :)")

	fd, err := os.OpenFile(titleName+"/exefs/main.npdm", os.O_RDWR, 0000)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	uint64TitleName, err := strconv.ParseUint("0x"+titleName, 0, 64)
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

func convert(titleIDName string, path string) {
	ncaFile := getBiggestNCA(titleIDName, path)

	decryptNCA(ncaFile, titleIDName)

	patchMainNPDM(titleIDName)
}

func main() {
	//var titleName string
	printHeader()

	gui()
}
