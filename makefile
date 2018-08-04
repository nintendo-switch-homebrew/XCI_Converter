NAME = XCI_Converter

SRCS = ./XCI_Decryptor.go ./gui.go

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRCS)

clean:
	go clean

fclean: clean
	/bin/rm -fr $(NAME)
	/bin/rm -fr ./build

re: fclean all

run:
	go run $(SRCS)

build:
	GOOS=darwin GOARCH=386 go build -o build/$(NAME)-darwin-386 $(SRCS)
	GOOS=darwin GOARCH=amd64 go build -o build/$(NAME)-darwin-amd64 $(SRCS)
	GOOS=linux GOARCH=386 go build -o build/$(NAME)-linux-386 $(SRCS)
	GOOS=linux GOARCH=amd64 go build -o build/$(NAME)-linux-amd64 $(SRCS)
	GOOS=linux GOARCH=arm go build -o build/$(NAME)-linux-arm $(SRCS)
	GOOS=linux GOARCH=arm64 go build -o build/$(NAME)-linux-arm64 $(SRCS)
	GOOS=freebsd GOARCH=386 go build -o build/$(NAME)-freebsd-386 $(SRCS)
	GOOS=freebsd GOARCH=amd64 go build -o build/$(NAME)-freebsd-amd64 $(SRCS)
	GOOS=freebsd GOARCH=arm go build -o build/$(NAME)-freebsd-arm $(SRCS)
	GOOS=netbsd GOARCH=386 go build -o build/$(NAME)-netbsd-386 $(SRCS)
	GOOS=netbsd GOARCH=amd64 go build -o build/$(NAME)-netbsd-amd64 $(SRCS)
	GOOS=netbsd GOARCH=arm go build -o build/$(NAME)-netbsd-arm $(SRCS)
	GOOS=openbsd GOARCH=386 go build -o build/$(NAME)-openbsd-386 $(SRCS)
	GOOS=openbsd GOARCH=amd64 go build -o build/$(NAME)-openbsd-amd64 $(SRCS)
	GOOS=openbsd GOARCH=arm go build -o build/$(NAME)-openbsd-arm $(SRCS)
	GOOS=windows GOARCH=386 go build -o build/$(NAME)-windows-386.exe $(SRCS)
	GOOS=windows GOARCH=amd64 go build -o build/$(NAME)-windows-amd64.exe $(SRCS)
