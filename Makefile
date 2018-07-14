all: clean space-cats

clean:
	rm space-cats 2>/dev/null || true

space-cats:
	go build -o space-cats ./game

play: clean space-cats
	./space-cats

