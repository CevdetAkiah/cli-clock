package digiclock

var (

	//individual digits
	zero  = [5]string{"███", "█ █", "█ █", "█ █", "███"}
	one   = [5]string{"██", " █ ", " █ ", " █ ", "███"}
	two   = [5]string{"███", "  █", "███", "█  ", "███"}
	three = [5]string{"███", "  █", "███", "  █", "███"}
	four  = [5]string{"█ █", "█ █", "███", "  █", "  █"}
	five  = [5]string{"███", "█  ", "███", "  █", "███"}
	six   = [5]string{"███", "█  ", "███", "█ █", "███"}
	seven = [5]string{"███", "  █", "  █", "  █", "  █"}
	eight = [5]string{"███", "█ █", "███", "█ █", "███"}
	nine  = [5]string{"███", "█ █", "███", "  █", "███"}

	separator = [5]string{"   ", " █ ", "   ", " █ ", "   "}
	blank     = [5]string{"", "", "", "", ""}

	//digits store
	digits = [10][5]string{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	clock = [8][5]string{}
)
