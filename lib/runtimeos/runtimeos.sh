#!/bin/sh

# List known operating system/architecture combos, removing the
# architecture, and making constants of them.

osname() {
    echo "$1" | awk '{print toupper(substr($1, 1, 1)) substr($1, 2)}' \
        | sed s/aix/AIX/i \
        | sed s/bsd/BSD/i \
        | sed s/ios/IOS/i \
        | sed s/js/JS/i
}

osconstants() {
    echo "package runtimeos"
    echo "import \"runtime\""

    echo "const ("
    for OS in $(go tool dist list | sed 's|/.*||' | sort | uniq) ; do
        name=$(osname "$OS")
        echo "$name = \"$OS\""
    done
    echo ")"
    echo

    echo "const ("
    for OS in $(go tool dist list | sed 's|/.*||' | sort | uniq) ; do
        name=$(osname "$OS")
        echo "Is$name = runtime.GOOS == $name"
    done
    echo ")"
}

osconstants > runtimeos.gen.go

gofmt -w -s runtimeos.gen.go
