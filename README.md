
## XCI_Converter

Tool for convert your .xci backup to playable Nintendo Switch games :D

You can compile the code or download the program here : https://github.com/SegFault42/XCI_Converter/releases

## Dependencies

- hactool

For Windows Users, you can download hactool from here : https://github.com/SciresM/hactool/releases

For Unix users, you can compile the soft from source

## Building

- install env for golang
- go get github.com/SegFault42/XCI_Converter
- cd $GOPATH/src/github.com/SegFault42/XCI_Converter
- go build

## Usage
- copy hactool and your .xci in the directory.
- you should have at least  :
XCI_Converter, hactool or hactool.exe, keys.dat, keys.ini, and your .xci backup

```
├── README.md
├── XCI_Decryptor
├── XCI_Decryptor.go
├── bbb-h-baazd.xci
├── hactool
├── keys.dat
└── keys.ini
```
- and then type in your terminal : ```./XCI_Decryptor [.XCI FILE] [TITLE ID]```
- You can now copy the new folder (with the title id) on your sd card in /atmosphere/titles

You can find all title id here : http://switchbrew.org/index.php?title=Title_list/Games

⚠️ Title id is the game installed on your switch, not the backup.

⚠️ GUI version is in development 
