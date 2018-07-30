GOOS=darwin GOARCH=386 go build -o build/XCI_Converter-darwin-386
GOOS=darwin GOARCH=amd64 go build -o build/XCI_Converter-darwin-amd64

GOOS=linux GOARCH=386 go build -o build/XCI_Converter-linux-386
GOOS=linux GOARCH=amd64 go build -o build/XCI_Converter-linux-amd64
GOOS=linux GOARCH=arm go build -o build/XCI_Converter-linux-arm
GOOS=linux GOARCH=arm64 go build -o build/XCI_Converter-linux-arm64

GOOS=freebsd GOARCH=386 go build -o build/XCI_Converter-freebsd-386
GOOS=freebsd GOARCH=amd64 go build -o build/XCI_Converter-freebsd-amd64
GOOS=freebsd GOARCH=arm go build -o build/XCI_Converter-freebsd-arm

GOOS=netbsd GOARCH=386 go build -o build/XCI_Converter-netbsd-386
GOOS=netbsd GOARCH=amd64 go build -o build/XCI_Converter-netbsd-amd64
GOOS=netbsd GOARCH=arm go build -o build/XCI_Converter-netbsd-arm

GOOS=openbsd GOARCH=386 go build -o build/XCI_Converter-openbsd-386
GOOS=openbsd GOARCH=amd64 go build -o build/XCI_Converter-openbsd-amd64
GOOS=openbsd GOARCH=arm go build -o build/XCI_Converter-openbsd-arm

GOOS=windows GOARCH=386 go build -o build/XCI_Converter-windows-386.exe
GOOS=windows GOARCH=amd64 go build -o build/XCI_Converter-windows-amd64.exe
