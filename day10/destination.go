package main

type Destination interface {
	isOutput() bool
	isBot() bool
	canReceive() bool
}

func (b bot) isOutput() bool {
	return false
}

func (b bot) isBot() bool {
	return true
}

func (o output) isOutput() bool {
	return true
}

func (o output) isBot() bool {
	return false
}

func (o output) canReceive() bool {
	return true
}

func (b bot) canReceive() bool {
	return len(bot.vals) < 2
}
