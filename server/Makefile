NAME	=	npuzzle 
SRCS	=	$(shell find . -type f -name "*.go")

all: $(NAME)
	
$(NAME)	: $(SRCS)
	@go get ./...
	@go build -o $(NAME)
clean:
	@rm -f $(NAME)
re: clean all
run: all
	@./$(NAME) -f map/3_solvable.txt
