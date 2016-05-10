// Package foo is an example that shows how you can use doc.go to
// document your package using godoc. If you run
//
//    % godoc -ex -goroot `pwd` foo
//
// then you will see package documentation from all files of the
// package combined including examples.
//
// To see how godoc will render it as webpage, you can create html file foo.html
//
//    % godoc -goroot `pwd` -url ./pkg/foo > foo.html
//
// run a godoc webserver
//
//    % godoc -http=:6060 -goroot `pwd`
//
// and finally start your webbrowser and open http://localhost:6060/foo.html
package foo
