# fuckery
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o fuckery ./cmd/main.go
	scdoc < fuckery.1.scd | sed "s/VERSION/$(VERSION)/g" > fuckery.1

clean:
	rm -f fuckery
	rm -f fuckery.1

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f fuckery $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/fuckery
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f fuckery.1 $(DESTDIR)$(MANPREFIX)/man1/fuckery.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/fuckery.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/fuckery
	rm -f $(DESTDIR)$(MANPREFIX)/man1/fuckery.1

.PHONY: all build clean install uninstall
