package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-irc/irc"
	"github.com/joho/godotenv"
)

func IsPM(receiver string) bool {
	return !strings.HasPrefix(receiver, "#")
}

func GetMaintainers() []string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	MAINTAINERS := os.Getenv("MAINTAINERS")
	maintainers := strings.Split(MAINTAINERS, ",")
	return maintainers
}

func InformErrAndQuit(c *irc.Client, m *irc.Message) {
	maintainers := GetMaintainers()
	for _, maintainer := range maintainers {
		SendMsg(c, maintainer, "Error")
	}
	time.Sleep(5 * time.Second)
	os.Exit(1)
}

func HandlePM(msg string, sender string) {
	maintainers := GetMaintainers()
	if msg != "KILL" {
		return
	}
	for _, maintainer := range maintainers {
		if maintainer == sender {
			os.Exit(1)
		}
	}
}

func GetChans() []string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	IRC_CHANNELS := os.Getenv("IRC_CHANNELS")
	channels := strings.Split(IRC_CHANNELS, ",")
	return channels
}

func GetIRCEnvVars() (string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	IRC_NICK := os.Getenv("IRC_NICK")
	IRC_PASS := os.Getenv("IRC_PASS")
	IRC_USER := os.Getenv("IRC_USER")
	IRC_NAME := os.Getenv("IRC_NAME")
	return IRC_NICK, IRC_PASS, IRC_USER, IRC_NAME
}

func GetMsgDetails(m *irc.Message) (string, string, string) {
	from := m.Name
	to := m.Params[0]
	msg := m.Params[1]
	return from, to, msg
}

func JoinChannels(c *irc.Client) {
	IRC_CHANNELS := GetChans()
	for _, channel := range IRC_CHANNELS {
		c.Writef("JOIN %s", channel)
	}
}

func SendMsg(c *irc.Client, to string, msg string) {
	c.Writef("%s %s %s", "PRIVMSG", to, msg)
}
