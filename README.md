# DCPU [![GoDoc](https://godoc.org/github.com/robertsdionne/dcpu?status.svg)](https://godoc.org/github.com/robertsdionne/dcpu)

Package dcpu implements an emulator for Notch's [DCPU 1.7 specification](documents/dcpu-16.txt).

## TODO
* [x] Implement assembler:
  * [x] Define parser specification in ANTLR4.
  * [x] Write assembler abstract syntax tree visitor or listener.
* [ ] Implement devices:
  * [ ] Standard:
    * [x] Mackapar SPED-3
    * [x] Generic Clock
    * [x] Mackapar M35FD
    * [ ] Harold HMD2043
    * [x] Generic Keyboard
    * [ ] Kulog K8581
    * [x] Nya Elektriska LEM1802
    * [ ] Nya Elektriska SPC2000
  * [ ] Debug:
    * [x] Standard input
    * [x] Standard output
    * [x] Standard error
