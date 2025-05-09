//go:build !prod || full || netent

package trolls

import (
	"github.com/slotopol/server/game"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "NetEnt", Name: "Trolls", Year: 2009},    // see: https://casino.ru/trolls-netent/
		{Prov: "NetEnt", Name: "Excalibur", Year: 2011}, // see: https://casino.ru/excalibur-netent/
		{Prov: "NetEnt", Name: "Pandora's Box"},
		{Prov: "NetEnt", Name: "Wild Witches"},
	},
	GP: game.GPlpay |
		game.GPlsel |
		game.GPretrig |
		game.GPfgmult |
		game.GPscat |
		game.GPwild |
		game.GPwmult,
	SX:  5,
	SY:  3,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	Info.SetupFactory(func() any { return NewGame() }, CalcStat)
}
