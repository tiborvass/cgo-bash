#!/bin/bash

VERSION="$1"
[ -z "$VERSION" ] && echo "build.sh: must specify bash version to download" && exit 1

bashsrc="bash-${VERSION}"
archive="$bashsrc.tar.gz"
shasum="$archive.sha256"

if [ ! -f "$archive" ]; then
	wget "https://ftp.gnu.org/gnu/bash/$archive"
fi

shasum -c "$shasum" || exit 1
if [ ! -d "$bashsrc" ]; then
	(
		tar zxf "$archive"
		cd "${bashsrc}"
		for p in ../patches/*.patch; do
			patch -p1 < "$p"
		done
	)
fi

(
	cd "${bashsrc}"
	[ -f config.status ] || ./configure
	make static
	ldflags=$(make ldflags| tail -1 | sed -E "s%\./%\${SRCDIR}/${bashsrc}/%g")

	ldflags=$(
		for ldflag in $ldflags; do
			[ "${ldflag:0:1}" == "-" ] && echo -n "$ldflag " || echo -n "\${SRCDIR}/${bashsrc}/${ldflag} "
		done
	)

	gofmt -s > ../c.go <<-EOF
	package bash

	// #cgo LDFLAGS: ${ldflags}
	import "C"
	EOF
)

# always rebuild because Go doesn't know if C files/libraries have changed
go install -a
