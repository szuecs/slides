// Package x is an example that shows how you can use doc.go to
// document your package using godoc. If you run
//
//    % godoc -ex -goroot `pwd` x
//
// then you will see package documentation from all files of the
// package combined including examples.
//
// To see how godoc will render it as webpage, you can create html file x.html
//
//    % godoc -goroot `pwd` -url ./pkg/x > x.html
//
// run a godoc webserver
//
//    % godoc -http=:6060 -goroot `pwd`
//
// and finally start your webbrowser and open http://localhost:6060/x.html
package x
