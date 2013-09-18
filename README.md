rc4go
=====

RC4 encryption in Go

Created while following a Go tutorial, based off the Wikipedia page entry for RC4.

The algorithm implementation is flawed (and produces insecure/incompatible bytecode) due to issues with mod binary, I welcome any suggestions.

For virtually any purpose you should use the standard Go RC4 library in the crypto package.
