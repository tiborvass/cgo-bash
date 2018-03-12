#!/bin/sh

function install_instructions() {
	case "$1" in
	md)
		line=\`\`\`
		;;
	go)
		p=' '
		;;
	*)
		>&2 echo "doc.sh: unknown format '$1' for install instructions"
		return 1
		;;
	esac

	cat <<-EOF
	In order to install this package, you need to run go generate first:
	${line}
	${p}$ go get -d github.com/tiborvass/cgo-bash
	${p}$ cd \$GOPATH/src/github.com/tiborvass/cgo-bash
	${p}$ go generate # downloads, patches, and compiles bash
	${p}$ go install
	${line}
	If needed, you can compile it inside a Docker container:
	${line}
	${p}$ docker build -t cgo-bash .
	${p}$ docker run -it -v \$(pwd):/go/src/github.com/tiborvass/cgo-bash cgo-bash bash
	${p}# go generate
	${p}# go install
	${line}
	EOF
}

cat <<EOF > README.md
# cgo-bash

## Install

$(install_instructions md)

## Usage & Documentation

[Godoc documentation](https://godoc.org/github.com/tiborvass/cgo-bash)

### Why?

Why not?


# License

MIT
EOF

cat <<EOF > doc.go
/*
Package bash wraps a patched version of bash that has hooks to call user-defined Go functions as shell builtins.
$(install_instructions go)
*/
package bash
EOF
