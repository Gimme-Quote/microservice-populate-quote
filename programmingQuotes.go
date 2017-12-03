package quote

var programmingQuotes = []*Quote{
	{
		Author:   "Ellen Ullman",
		Quote:    "We build our computer (systems) the way we build our cities: over time, without a plan, on top of ruins.",
		Category: "programming",
	},
	{
		Author:   "Fred Brooks",
		Quote:    "What one programmer can do in one month, two programmers can do in two months.",
		Category: "programming",
	},
	{
		Author:   "Oscar Godson",
		Quote:    "One of the best programming skills you can have is knowing when to walk away for awhile.",
		Category: "programming",
	},
	{
		Author:   "Chris Wenham",
		Quote:    "Debugging time increases as a square of the program’s size.",
		Category: "programming",
	},
	{
		Author:   "Grady Booch",
		Quote:    "The amateur software engineer is always in search of magic.",
		Category: "programming",
	},
	{
		Author:   "Ryan Singer",
		Quote:    "So much complexity in software comes from trying to make one thing do two things.",
		Category: "programming",
	},
	{
		Author:   "Reg Braithwaite",
		Quote:    "The strength of JavaScript is that you can do anything. The weakness is that you will.",
		Category: "programming",
	},
	{
		Author:   "Wes Dyer",
		Quote:    "Make it correct, make it clear, make it concise, make it fast. In that order.",
		Category: "programming",
	},
	{
		Author:   "Melinda Varian",
		Quote:    "The best programs are the ones written when the programmer is supposed to be working on something else.",
		Category: "programming",
	},
	{
		Author:   "Niklaus Wirth",
		Quote:    "A primary cause of complexity is that software vendors uncritically adopt almost any feature that users want.",
		Category: "programming",
	},
	{
		Author:   "Robert Sewell",
		Quote:    "If Java had true garbage collection, most programs would delete themselves upon execution.",
		Category: "programming",
	},
	{
		Author:   "Ovidiu Platon",
		Quote:    "I don’t care if it works on your machine! We are not shipping your machine!",
		Category: "programming",
	},
	{
		Author:   "Ovidiu Platon",
		Quote:    "I don’t care if it works on your machine! We are not shipping your machine!",
		Category: "programming",
	},
	{
		Author:   "Ovidiu Platon",
		Quote:    "I don’t care if it works on your machine! We are not shipping your machine!",
		Category: "programming",
	},
	{
		Author:   "Martin Fowler",
		Quote:    "Whenever you are tempted to type something into a print statement or a debugger expression, write it as a test instead.",
		Category: "programming",
	},
	{
		Author:   "Marc Andreessen",
		Quote:    "Learning to code is the single best thing anyone can do to get the most out of the amazing future in front of us.",
		Category: "programming",
	},
	{
		Author:   "Ken Thompson",
		Quote:    "I think the major good idea in Unix was its clean and simple interface: open, close, read, and write.",
		Category: "programming",
	},
	{
		Author:   "Bjarne Stroustrup",
		Quote:    "If you think it’s simple, then you have misunderstood the problem.",
		Category: "programming",
	},
	{
		Author:   "Grace Hopper",
		Quote:    "One accurate measurement is worth a thousand expert opinions.",
		Category: "programming",
	},
	{
		Author:   "Bertrand Meyer",
		Quote:    "Incorrect documentation is often worse than no documentation.",
		Category: "programming",
	},
	{
		Author:   "Steward Brand",
		Quote:    "Once a new technology starts rolling, if you're not part of the steamroller, you're part of the road.",
		Category: "programming",
	},
	{
		Author:   "William Laeder",
		Quote:    "Figure out your data structures, and the code will follow.",
		Category: "programming",
	},
}

func GetProgrammingQuotes() []*Quote {
	return programmingQuotes
}
