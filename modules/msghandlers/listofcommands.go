package msghandlers

import (
	"recommendtgbot/modules/commonvariables"
	"strings"
)

func PrintAllCommands() []string {
	commandsSlice := make([]string, 0, len(commonvariables.Commands))
	for _, s := range commonvariables.Commands {
		commandsSlice = append(commandsSlice, s.Command, s.CommandDescription)
	}
	return commandsSlice
}

func MsgStart(commandText string) (msgToSend string) {
	msgToSend = ("Привет! Я твой помощник-напоминалка-запоминалка. Я могу выполнять следующие команды: \n\n" + strings.Join(PrintAllCommands(), ""))
	return msgToSend
}
