package main

type digit [9]string

var separator = digit{
	"     ",
	"     ",
	" ░░░ ",
	"     ",
	"     ",
	"     ",
	" ░░░ ",
	"     ",
	"     ",
}

var dot = digit{
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	" ░░░ ",
}

var blank = digit{
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
	"     ",
}

var zero = digit{
	" ███ ",
	"█   █",
	"█   █",
	"█   █",
	"█   █",
	"█   █",
	"█   █",
	"█   █",
	" ███ ",
}

var one = digit{
	"  █  ",
	" ██  ",
	"  █  ",
	"  █  ",
	"  █  ",
	"  █  ",
	"  █  ",
	"  █  ",
	" ███ ",
}

var two = digit{
	" ███ ",
	"█   █",
	"    █",
	"    █",
	"   █ ",
	"  █  ",
	" █   ",
	"█    ",
	"█████",
}

var three = digit{
	" ███ ",
	"█   █",
	"    █",
	"  ██ ",
	"    █",
	"    █",
	"    █",
	"█   █",
	" ███ ",
}

var four = digit{
	"    █",
	"█   █",
	"█   █",
	"█   █",
	"█   █",
	"█████",
	"    █",
	"    █",
	"    █",
}

var five = digit{
	" ████",
	"█    ",
	"█    ",
	"████ ",
	"    █",
	"    █",
	"    █",
	"█   █",
	" ███ ",
}

var six = digit{
	" ███ ",
	"█   █",
	"█    ",
	"████ ",
	"█   █",
	"█   █",
	"█   █",
	"█   █",
	" ███ ",
}

var seven = digit{
	"████ ",
	"    █",
	"    █",
	"    █",
	"   █ ",
	"  █  ",
	"  █  ",
	"  █  ",
	"  █  ",
}

var eight = digit{
	" ███ ",
	"█   █",
	"█   █",
	"█   █",
	" ███ ",
	"█   █",
	"█   █",
	"█   █",
	" ███ ",
}

var nine = digit{
	" ███ ",
	"█   █",
	"█   █",
	"█   █",
	" ████",
	"    █",
	"    █",
	"█   █",
	" ███ ",
}

var digits = [...]digit{
	zero,
	one,
	two,
	three,
	four,
	five,
	six,
	seven,
	eight,
	nine,
}
