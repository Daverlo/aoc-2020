GO = go

all: clean build

days := $(wildcard ./src/days/*)
days := $(patsubst ./src/days/%,%,$(days))

$(days):
	$(GO) build -o ./bin/$@ ./src/days/$@/main.go

build: $(days)

clean:
	rm -rf bin
