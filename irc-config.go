package main

import (
	"crypto/tls"
	"fmt"

	"github.com/go-irc/irc"
)

func GetIRCConfig() irc.ClientConfig {
	IRC_NICK, IRC_PASS, IRC_USER, IRC_NAME := GetIRCEnvVars()
	ircConfig := irc.ClientConfig{
		Nick:    IRC_NICK,
		Pass:    IRC_PASS,
		User:    IRC_USER,
		Name:    IRC_NAME,
		Handler: irc.HandlerFunc(Handler),
	}
	return ircConfig
}

func GetIRCClient(conn *tls.Conn) *irc.Client {
	ircConfig := GetIRCConfig()              // IRC configuration
	client := irc.NewClient(conn, ircConfig) // Create IRC Client
	return client
}

func Handler(c *irc.Client, m *irc.Message) {
	if m.Command == "001" {
		JoinChannels(c)
	} else if m.Command == "PRIVMSG" {
		from, to, msg := GetMsgDetails(m)
		if IsPM(to) {
			HandlePM(msg, from)
		} else {
			fmt.Println(msg, from, to)
		}
	} else if m.Command == "ERROR" {
		InformErrAndQuit(c, m)
	}
}
