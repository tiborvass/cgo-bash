/*
Package bash wraps a patched version of bash that has hooks to call user-defined Go functions as shell builtins.
In order to install this package, you need to run go generate first:

 $ go get -d github.com/tiborvass/cgo-bash
 $ cd $GOPATH/src/github.com/tiborvass/cgo-bash
 $ go generate # downloads, patches, and compiles bash
 $ go install

If needed, you can compile it inside a Docker container:

 $ docker build -t cgo-bash .
 $ docker run -it -v $(pwd):/go/src/github.com/tiborvass/cgo-bash cgo-bash bash
 # go generate
 # go install
*/
package bash
