module hello-world

go 1.15

replace rodrigodonizettio.com/greetings => ../create-module

require (
	rodrigodonizettio.com/greetings v1.1.0
	rsc.io/quote v1.5.2
)
