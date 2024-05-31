package main

import (
	"strconv"
	"strings"
)

type CommandInput struct {
	Command  string
	EchoArgs []string
	SetArgs  SetArgs
	GetArgs  []string
}

type SetArgs struct {
	Key    string
	Val    string
	Expiry string
}

func parseRESP(input []byte) (commandInput *CommandInput) {
	bulkStringCount, _ := strconv.Atoi(string(input[1:2]))

	if bulkStringCount < 1 {
		return nil
	}

	commandInput = &CommandInput{}

	splitStrings := strings.Split(string(input[2:]), "\r\n")

	command := splitStrings[2]

	switch command {
	case PING:
		commandInput.Command = PING
	case ECHO:
		commandInput.Command = ECHO
		commandInput.EchoArgs = splitStrings[4:]
	case SET:
		commandInput.Command = SET
		commandInput.SetArgs.Key = splitStrings[4]
		commandInput.SetArgs.Val = splitStrings[6]

		if len(splitStrings) > 9 {
			commandInput.SetArgs.Expiry = splitStrings[10]
		} else {
			commandInput.SetArgs.Expiry = ""
		}

	case GET:
		commandInput.Command = GET
		commandInput.GetArgs = splitStrings[4:]
	}

	return commandInput
}
