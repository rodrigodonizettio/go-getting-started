module hello-world

go 1.15

replace rodrigodonizettio.com/greetings => ../create-module

require (
	golang.org/x/tour v0.0.0-20201207214521-004403599411 // indirect
	rodrigodonizettio.com/greetings v1.1.0
	rsc.io/quote v1.5.2
)
