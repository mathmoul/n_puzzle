NAME=n_puzzle
SRC_PATH=src
GO=go
BUILD=build

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
	go build -o $(NAME) $(SRC)

fclean:
	rm -rf $(NAME)

re: fclean all
