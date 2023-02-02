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
			Title:     "Another Page",
			ShortName: "second",
			Body:      "This is just a second page.",
			Code:      `Say "Goodbye" to me.`,
		},
	}
}
