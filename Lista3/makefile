#Implementation of 3 tasks for list of AOD, by Vadym Liss (257264)
CC=gcc
CFLAGS=-Iinclude -Wall -lm

INCLUDE = include
BUILDDIR = obj
SRCDIR = src
TARGET = bin
RESULTS = results
TEST = test

_DEPS = uinterface.h heap.h radixheapstruct.h list.h
DEPS = $(patsubst %,$(INCLUDE)/%,$(_DEPS))
_ZAD1 = dijkstra.o heap.o list.o uinterface.o
_ZAD2 = dial.o list.o uinterface.o
_ZAD3 = radixheap.o radixheapstruct.o list.o uinterface.o
ZAD1 = $(patsubst %,$(BUILDDIR)/%,$(_ZAD1))
ZAD2 = $(patsubst %,$(BUILDDIR)/%,$(_ZAD2))
ZAD3 = $(patsubst %,$(BUILDDIR)/%,$(_ZAD3))

compile: dijkstra dial radixheap

$(BUILDDIR)/%.o: $(SRCDIR)/%.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

dijkstra: $(ZAD1)
	$(CC) -o $(TARGET)/$@ $^ $(CFLAGS)

dial: $(ZAD2)
	$(CC) -o $(TARGET)/$@ $^ $(CFLAGS)

radixheap: $(ZAD3)
	$(CC) -o $(TARGET)/$@ $^ $(CFLAGS)

clean:
	rm $(BUILDDIR)/*.o
