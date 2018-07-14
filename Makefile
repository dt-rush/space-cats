all: clean space-cats

space-cats:
	go build -o space-cats ./game

clean:
	rm space-cats 2>/dev/null || true

