GO = go


all: clean build

days := $(wildcard ./src/days/*)
days := $(patsubst ./src/days/%,%,$(days))

$(days):
	$(GO) build -o ./bin/$@ ./src/days/$@/main.go
	./bin/$@ "./test/$@/in" > /tmp/aoc-2020-$@-out
	diff "./test/$@/out" /tmp/aoc-2020-$@-out

build: $(days)

clean:
	rm -rf bin
