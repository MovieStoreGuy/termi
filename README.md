# Termi
[![Go Report Card](https://goreportcard.com/badge/github.com/MovieStoreGuy/termi)](https://goreportcard.com/report/github.com/MovieStoreGuy/termi)
[![Maintainability](https://api.codeclimate.com/v1/badges/901065b20a0021ba2ee9/maintainability)](https://codeclimate.com/github/MovieStoreGuy/termi/maintainability)
[![GoDoc](https://godoc.org/github.com/MovieStoreGuy/termi?status.svg)](https://godoc.org/github.com/MovieStoreGuy/termi)  

Termi is a library that enables applications flag parsing and help messages made simple.
The benefit of builder like patterns, multiple names for the same flag, and be able to obtain the remaining
arguments after parsing to use later makes for detailed application interfaces for command line applications.

If you regularly build applications that may work within a CI environment, 
can define settings variables to automatically load variables from the shell environment 
instead of passing them as flags. This way secrets can be stored within CI and exported at runtime 
given your CI of choice. 
