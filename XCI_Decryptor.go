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

func getBiggestNCA(titleName string) (file string) {
	var (
		nca myFile
		cmd *exec.Cmd
	)

	color.Green("Decrypting .xci's NCA files and finding the biggest NCA...")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("./hactool.exe", "-k", "keys.ini", "-txci", "--securedir="+titleName, os.Args[1])
	} else {
		cmd = exec.Command("./hactool", "-k", "keys.ini", "-txci", "--securedir="+titleName, os.Args[1])
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

func main() {
	var titleName string
	printHeader()

	if len(os.Args) != 3 {
		printUsage()
		return
	} else {
		titleName = os.Args[2]
	}

	ncaFile := getBiggestNCA(titleName)

	decryptNCA(ncaFile, titleName)

	patchMainNPDM(titleName)

	color.Green("DONE! You should have a folder: " + titleName)
	color.Green(titleName + " should contain an exefs folder and a romfs.bin. It should NOT contain anything else.")
}
