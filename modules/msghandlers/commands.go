package msghandlers

type BotCommands struct {
	command            string
	commandDescription string
}

func CommandPrint() []string {
	commands := []BotCommands{
		{command: "/addGame", commandDescription: " - добавить новую игру в Ваш список игр в формате: \n\"Игра, Описание игры\"\nПример команды: /addGame Spyro, Игра про дракончика\n\n"},
		{command: "/addBook", commandDescription: " - добавить новую книгу в Вашу библиотеку в формате: \n\"Книга, Описание книги\"\nПример команды: /addBook 1984, Роман-антиутопия\n\n"},
		{command: "/addFilm", commandDescription: " - добавить новый фильм в Вашу фильмотеку в формате: \n\"Фильм, Описание фильма\"\nПример команды: /addFilm Джонни Мнемоник, фантастика\n\n"},
		{command: "/getGames", commandDescription: " - посоветовать рандомную игру из Вашего списка.\nПример команды: /getGames \nПример команды, если хотите вывести заданное количество игр: /getGames 3\n\n"},
		{command: "/getBooks", commandDescription: " - посоветовать рандомную книгу из Вашего списка.\nПример команды: /getBooks \nПример команды, если хотите вывести заданное количество игр: /getBooks 3\n\n"},
		{command: "/getFilms", commandDescription: " - посоветовать рандомный фильм из Вашего списка.\nПример команды: /getFilms \nПример команды, если хотите вывести заданное количество игр: /getFilms 3\n\n"},
		{command: "/clearAll", commandDescription: " - очистить все списки.\nПример команды: /clearAll \n\n"},
		{command: "/clearGames", commandDescription: " - очистить весь список игр или удалить определённое количество игр из списка.\nПример команды: /clearGames или /clearGames 3 \n\n"},
		{command: "/clearBooks", commandDescription: " - очистить весь список книг или удалить определённое количество книг из списка.\nПример команды: /clearBooks или /clearBooks 3 \n\n"},
		{command: "/clearFilms", commandDescription: " - очистить весь список фильмов или удалить определённое количество фильмов из списка.\nПример команды: /clearFilms или /clearFilms 3 \n\n"},
	}

	CommandsSlice := make([]string, 0, len(commands))
	for _, s := range commands {
		CommandsSlice = append(CommandsSlice, s.command, s.commandDescription)
	}
	return CommandsSlice
}
