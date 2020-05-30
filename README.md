# RPN Calc

A RPN command line calculator.

## Features

* History that can be saved
* Comments can be added to history
* Configurable settings
* Registers for storing values

## Installation

1. Clone this repo
1. Go to the `rpncli` directory
1. Run `go build`
1. Run `./rpncli`

## Ideas for the future

* Add more advanced functions
    * trigonometric
    * ! faculty
* Support `!` to execute shell commands
    * if last line contains a value, put it in stack
* Add external scripting
    * configure scripts location in settings
    * send stack to script and get a value in return

