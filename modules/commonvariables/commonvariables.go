package commonvariables

type BotCommands struct {
	Command            string
	CommandDescription string
}

type PropertiesOfElement struct {
	NameDesc  string
	TitleDesc string
}

// Список массивов, доступных к заполнению, хранению и выводу
var Games = []PropertiesOfElement{}
var Books = []PropertiesOfElement{}
var Films = []PropertiesOfElement{}

// Список всех доступных комманд
var Commands = []BotCommands{
	{Command: "/addGame", CommandDescription: " - добавить новую игру в Ваш список игр в формате: \n\"Игра, Описание игры\"\nПример команды: /addGame Spyro, Игра про дракончика\n\n"},
	{Command: "/addBook", CommandDescription: " - добавить новую книгу в Вашу библиотеку в формате: \n\"Книга, Описание книги\"\nПример команды: /addBook 1984, Роман-антиутопия\n\n"},
	{Command: "/addFilm", CommandDescription: " - добавить новый фильм в Вашу фильмотеку в формате: \n\"Фильм, Описание фильма\"\nПример команды: /addFilm Джонни Мнемоник, фантастика\n\n"},
	{Command: "/getGames", CommandDescription: " - посоветовать рандомную игру из Вашего списка.\nПример команды: /getGames \nПример команды, если хотите вывести заданное количество игр: /getGames 3\n\n"},
	{Command: "/getBooks", CommandDescription: " - посоветовать рандомную книгу из Вашего списка.\nПример команды: /getBooks \nПример команды, если хотите вывести заданное количество игр: /getBooks 3\n\n"},
	{Command: "/getFilms", CommandDescription: " - посоветовать рандомный фильм из Вашего списка.\nПример команды: /getFilms \nПример команды, если хотите вывести заданное количество игр: /getFilms 3\n\n"},
	{Command: "/clearAll", CommandDescription: " - очистить все списки.\nПример команды: /clearAll \n\n"},
	{Command: "/clearGames", CommandDescription: " - очистить весь список игр или удалить определённое количество игр из списка.\nПример команды: /clearGames или /clearGames 3 \n\n"},
	{Command: "/clearBooks", CommandDescription: " - очистить весь список книг или удалить определённое количество книг из списка.\nПример команды: /clearBooks или /clearBooks 3 \n\n"},
	{Command: "/clearFilms", CommandDescription: " - очистить весь список фильмов или удалить определённое количество фильмов из списка.\nПример команды: /clearFilms или /clearFilms 3 \n\n"},
}
