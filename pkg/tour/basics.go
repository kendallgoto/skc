package tour

func buildBasics() []TourItem {
	return []TourItem{
		{
			Title:     "Hello, world!",
			ShortName: "printing",
			Body:      "Telling a interpreter to say something is one of the most basic tasks in programming.<br><br>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris maximus tellus ut mauris hendrerit sollicitudin. Duis eget nisi ac nisi placerat eleifend vel quis diam. Pellentesque viverra orci nisi, eu convallis orci luctus sed. Maecenas placerat sit amet neque non ullamcorper. Etiam massa turpis, vulputate non tellus vel, efficitur rhoncus lacus. Quisque risus nunc, gravida vitae luctus at, posuere id nibh. Sed enim sapien, egestas et tempus eget, suscipit sit amet magna.<br><br>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris maximus tellus ut mauris hendrerit sollicitudin. Duis eget nisi ac nisi placerat eleifend vel quis diam. Pellentesque viverra orci nisi, eu convallis orci luctus sed. Maecenas placerat sit amet neque non ullamcorper. Etiam massa turpis, vulputate non tellus vel, efficitur rhoncus lacus. Quisque risus nunc, gravida vitae luctus at, posuere id nibh. Sed enim sapien, egestas et tempus eget, suscipit sit amet magna.",
			Code:      `Say "Hello" to me.`,
		},
		{
			Title:     "Conditionals",
			ShortName: "conditionals",
			Body:      "If you want to have varying logic in your code, an if statement is a must-have!",
			Code:      `If "hello" is "hello", then say "Hi there!". If "you" is not equal to "cool", then say "Sucker!".`,
		},
		{
			Title:     "File I/O",
			ShortName: "fileio",
			Body:      "If you want to corrupt an OS, you'll need to be able to write to the disk!",
			Code:      `Say "hello" to file "temp.txt".`,
		},
	}
}
