##
## Makefile for echo-server
##
## Made by Kevin Almansa
## Email   kevin.almansa@gmail.com
##


VERSION := $(shell git describe --tags 2>/dev/null || echo 1.0.0)
BUILD := $(shell git rev-parse --short HEAD)

NAME := $(shell basename "$(PWD)")

RM		=	rm -f

GO		=	go
DOCKER	=	docker

LDFLAGS			=	-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
STATICLDFLAGS	= 	-ldflags "-s -w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -linkmode external -extldflags -static"
ifeq ($(debug),yes)
LDFLAGS		=	-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
STATICLDFLAGS	= 	-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD) -linkmode external -extldflags -static"
endif

SRCS		=	cmd/echo-server/main.go

all: $(NAME)

$(NAME):
	$(GO) build $(LDFLAGS) -o $(NAME) $(SRCS)

static:
	$(GO) build $(STATICLDFLAGS) -o $(NAME) $(SRCS)

fclean:
	$(RM) $(NAME)

re: fclean all

docker:
	$(DOCKER) build -t $(NAME):$(VERSION) .
