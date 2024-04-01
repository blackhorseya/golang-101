package main

// Job is a function that takes a variadic number of arguments and returns a response and an error.
type Job func(args ...any) (resp any, err error)
