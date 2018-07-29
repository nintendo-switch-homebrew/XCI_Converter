package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func printUsage() {
	log.Fatal("Usage : ./Decrypt-XCI-v2.1 [YOUR_BACKUP]")
}

type myFile struct {
	Size int64
	Name string
}

func getBiggestNCA() (file string) {
	var nca myFile

	cmd := exec.Command("./hactool", "-k", "keys.ini", "-txci", "--securedir=xciDecrypted", os.Args[1])
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("xciDecrypted")
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

func decryptNCA(ncaFile string) {
	cmd := exec.Command("./hactool", "-k", "keys.ini", "--romfs=xciDecrypted/romfs.bin", "--exefsdir=xciDecrypted/exefs", "xciDecrypted/"+ncaFile)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("xciDecrypted")
	if err != nil {
		log.Fatal(err)
	}

	// delete *.nca
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".nca" {
			err := os.Remove("./xciDecrypted/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
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

func main() {
	printHeader()

	if len(os.Args) == 1 {
		printUsage()
	}

	log.Println("Decrypting .xci's NCA files and finding the biggest NCA...")
	ncaFile := getBiggestNCA()

	decryptNCA(ncaFile)

	fmt.Println("DONE! You should have a folder: xciDecrypted")
	fmt.Println("xciDecrypted should contain an exefs folder and a romfs.bin. It should NOT contain anything else.")
}
