package config

// Cannot be edited >
const chatIdConst = "somedata"
const botIdConst = "somedata"

// <

type Conf struct {
	ChatId string
	BotId  string
}

func GetConf() Conf {
	return Conf{ChatId: chatIdConst, BotId: botIdConst}
}
