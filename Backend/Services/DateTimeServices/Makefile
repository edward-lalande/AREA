
SRC		=	./main.go

NAME	=	time-services

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRC)

clean:
	go clean

fclean:	clean
	$(RM) $(NAME)

re:	fclean all

.PHONY:	all clean fclean re
