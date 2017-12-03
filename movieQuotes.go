package quote

var movieQuotes = []*Quote{
	{
		Author:   "The Godfather",
		Quote:    "I'm gonna make him an offer he can't refuse.",
		Category: "movie",
		Year:     "1972",
	},
	{
		Author:   "Star Wars",
		Quote:    "May the Force be with you.",
		Category: "movie",
		Year:     "1977",
	},
	{
		Author:   "Dr. No",
		Quote:    "Bond. James Bond.",
		Category: "movie",
		Year:     "1962",
	},
	{
		Author:   "The Wizard of Oz",
		Quote:    "There's no place like home.",
		Category: "movie",
		Year:     "1939",
	},
	{
		Author:   "Jerry Maguire",
		Quote:    "Show me the money!",
		Category: "movie",
		Year:     "1996",
	},
	{
		Author:   "Gone with the Wind",
		Quote:    "After all, tomorrow is another day!",
		Category: "movie",
		Year:     "1939",
	},
	{
		Author:   "Terminator",
		Quote:    "I'll be back.",
		Category: "movie",
		Year:     "1984",
	},
	{
		Author:   "Jaws",
		Quote:    "You're gonna need a bigger boat.",
		Category: "movie",
		Year:     "1975",
	},
	{
		Author:   "Forrest Gump",
		Quote:    "My mama always said life was like a box of chocolates. \nYou never know what you're gonna get.",
		Category: "movie",
		Year:     "1994",
	},
	{
		Author:   "Bonnie And Clyde",
		Quote:    "We rob banks.",
		Category: "movie",
		Year:     "1967",
	},
	{
		Author:   "The Sixth Sense",
		Quote:    "I see dead people.",
		Category: "movie",
		Year:     "1999",
	},
	{
		Author:   "Apollo 13",
		Quote:    "Houston, we have a problem.",
		Category: "movie",
		Year:     "1995",
	},
	{
		Author:   "Jerry Maguire",
		Quote:    "You had me at \"hello.\"",
		Category: "movie",
		Year:     "1996",
	},
	{
		Author:   "Wall Street",
		Quote:    "Greed, for lack of a better word, is good.",
		Category: "movie",
		Year:     "1987",
	},
	{
		Author:   "The Godfather Part II",
		Quote:    "Keep your friends close, but your enemies closer.",
		Category: "movie",
		Year:     "1974",
	},
	{
		Author:   "Gone with the Wind",
		Quote:    "As God is my witness, I'll never be hungry again.",
		Category: "movie",
		Year:     "1939",
	},
	{
		Author:   "Scarface",
		Quote:    "Say \"hello\" to my little friend!",
		Category: "movie",
		Year:     "1983",
	},
	{
		Author:   "Dr. Strangelove",
		Quote:    "Gentlemen, you can't fight in here! This is the War Room!",
		Category: "movie",
		Year:     "1964",
	},
	{
		Author:   "Terminator 2: Judgement Day",
		Quote:    "Hasta la vista, baby.",
		Category: "movie",
		Year:     "1991",
	},
	{
		Author:   "The Lord of the Rings: Two Towers",
		Quote:    "My precious.",
		Category: "movie",
		Year:     "2002",
	},
	{
		Author:   "Goldfinger",
		Quote:    "A martini. Shaken, not stirred.",
		Category: "movie",
		Year:     "1964",
	},
	{
		Author:   "Top Gun",
		Quote:    "I feel the need - the need for speed!",
		Category: "movie",
		Year:     "1986",
	},
	{
		Author:   "Dead Poets Society",
		Quote:    "Carpe diem. Seize the day, boys. Make your lives extraordinary.",
		Category: "movie",
		Year:     "1989",
	},
	{
		Author:   "Dirty Dancing",
		Quote:    "Nobody puts Baby in a corner.",
		Category: "movie",
		Year:     "1987",
	},
	{
		Author:   "Titanic",
		Quote:    "I'm the king of the world!",
		Category: "movie",
		Year:     "1997",
	},
	{
		Author:   "Rocky",
		Quote:    "It's not about how hard you can hit; it's about how hard you can get hit and keep moving forward.",
		Category: "inspiration",
	},
}

func GetMovieQuotes() []*Quote {
	return movieQuotes
}
