NAME=n_puzzle
SRC_PATH=.
GO=go
BUILD=build
DEBUG=debug

SRC_NAME=actions.go \
bst.go \
main.go \
parser.go \
puzzle.go \
solver.go \
tools.go \
astar.go \
node.go \
heuristic.go

all: $(NAME)
SRC = $(addprefix $(SRC_PATH)/, $(SRC_NAME))

$(NAME): $(SRC)
	go get
	go build -o $(NAME) $(SRC)

fclean:
	rm -rf $(NAME)

re: fclean all

debug:
	go get
	go build -gcflags "-m -m -l" -o $(DEBUG) $(SRC) 

fclean_debug:
	@rm -rfv $(DEBUG)

re_debug: fclean_debug debug
