module hello

go 1.16

replace (
	mylearning.com/greetings latest => ../greetings latest
)

require (
	mylearning.com/greetings lastest
)