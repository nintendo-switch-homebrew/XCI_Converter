Tool for convert your .xci backup to playable Nintendo Switch games :D

You can compile the code or download the program here : https://github.com/SegFault42

## Dependencies

- hactool

For Windows Users, you can download hactool from here : https://github.com/SciresM/hactool/releases

For Unix users, you can compil the soft from source

## Building

- install env for golang
- go build Decrypt-XCI-v2.1.go

## Usage
- put hactool and your .xci in the directory.
- you should have at least  :
XCI_Decryptor, hactool, keys.dat, keys.ini, and your .xci backup

```
├── README.md
├── XCI_Decryptor
├── XCI_Decryptor.go
├── bbb-h-baazd.xci
├── hactool
├── keys.dat
└── keys.ini
```
- and then type in yout terminal : ```./XCI_Decryptor [.XCI FILE] [TITLE ID]```
- You can now copy the new folder (with the title id) on your sd card

You can find all title id here : http://switchbrew.org/index.php?title=Title_list/Games

⚠️ Title id is the game installed on your switch, not the backup.
