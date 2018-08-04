NAME = XCI_Converter

SRCS = ./XCI_Decryptor.go ./gui.go

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRCS)

clean:
	go clean

fclean: clean
	/bin/rm -fr $(NAME)

re: fclean all

run:
	go run $(SRCS)
