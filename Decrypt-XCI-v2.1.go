package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type myFile struct {
	Size int64
	Name string
}

func swapLittleEndian(n uint64) uint64 {
	return ((n & 0x00000000000000FF) << 56) |
		((n & 0x000000000000FF00) << 40) |
		((n & 0x0000000000FF0000) << 24) |
		((n & 0x00000000FF000000) << 8) |
		((n & 0x000000FF00000000) >> 8) |
		((n & 0x0000FF0000000000) >> 24) |
		((n & 0x00FF000000000000) >> 40) |
		((n & 0xFF00000000000000) >> 56)
}

func printHeader() {
	const header string = `██████╗ ███████╗ ██████╗██████╗ ██╗   ██╗██████╗ ████████╗  ██╗  ██╗ ██████╗██╗      ██╗   ██╗██████╗    ██╗
██╔══██╗██╔════╝██╔════╝██╔══██╗╚██╗ ██╔╝██╔══██╗╚══██╔══╝  ╚██╗██╔╝██╔════╝██║      ██║   ██║╚════██╗  ███║
██║  ██║█████╗  ██║     ██████╔╝ ╚████╔╝ ██████╔╝   ██║█████╗╚███╔╝ ██║     ██║█████╗██║   ██║ █████╔╝  ╚██║
██║  ██║██╔══╝  ██║     ██╔══██╗  ╚██╔╝  ██╔═══╝    ██║╚════╝██╔██╗ ██║     ██║╚════╝╚██╗ ██╔╝██╔═══╝    ██║
██████╔╝███████╗╚██████╗██║  ██║   ██║   ██║        ██║     ██╔╝ ██╗╚██████╗██║       ╚████╔╝ ███████╗██╗██║
╚═════╝ ╚══════╝ ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝        ╚═╝     ╚═╝  ╚═╝ ╚═════╝╚═╝        ╚═══╝  ╚══════╝╚═╝╚═╝`

	fmt.Println(header)
}

func printUsage() {
	log.Fatal("Usage : ./Decrypt-XCI-v2.1 [.XCI FILE] [TITLE ID]")
}

func getBiggestNCA(titleName string) (file string) {
	var nca myFile

	cmd := exec.Command("./hactool", "-k", "keys.ini", "-txci", "--securedir="+titleName, os.Args[1])
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
	cmd := exec.Command("./hactool", "-k", "keys.ini", "--romfs="+titleName+"/romfs.bin", "--exefsdir="+titleName+"/exefs", titleName+"/"+ncaFile)
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
	fmt.Printf("Patching main.npdm :)")

	//fd, err := os.OpenFile(titleName + "/exefs/main.npdm", os.O_RDWR)
	fd, err := os.OpenFile("./test", os.O_RDWR, 0000)
	if err != nil {
		log.Fatal(err)
	}

	hexTitleName, err := hex.DecodeString(titleName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fd.WriteAt([]byte(hexTitleName), 0x440)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var titleName string
	printHeader()

	if len(os.Args) != 1 {
		printUsage()
	} else if len(os.Args) == 3 {
		titleName = os.Args[2]
	}

	log.Println("Decrypting .xci's NCA files and finding the biggest NCA...")
	ncaFile := getBiggestNCA(titleName)
	decryptNCA(ncaFile, titleName)

	log.Println("Patching main.mpdm ...")
	patchMainNPDM(titleName)

	fmt.Println("DONE! You should have a folder: " + titleName)
	fmt.Println(titleName + " should contain an exefs folder and a romfs.bin. It should NOT contain anything else.")
}
