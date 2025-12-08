package config

type Platform struct {
	Enabled bool   `json:"enabled"`
	Label   string `json:"label"`
	Dir     string `json:"dir"`
}

type PlatformData struct {
	// ARCADE

	HBMAME Platform `json:"HBMAME"`
	MAME   Platform `json:"MAME"`

	// ATARI
	Atari2600   Platform `json:"Atari2600"`
	Atari5200   Platform `json:"Atari5200"`
	Atari7800   Platform `json:"Atari7800"`
	Atari8Bit   Platform `json:"Atari8Bit"`
	AtariJaguar Platform `json:"AtariJaguar"`
	AtariLynx   Platform `json:"AtariLynx"`

	// MICROSOFT

	Xbox        Platform `json:"Xbox"`
	Xbox360     Platform `json:"Xbox360"`
	XboxOne     Platform `json:"XboxOne"`
	XboxSeriesX Platform `json:"XboxSeriesX"`

	// NINTENDO

	Nintendo64             Platform `json:"Nintendo64"`
	NintendoGameBoy        Platform `json:"NintendoGameBoy"`
	NintendoGameBoyAdvance Platform `json:"NintendoGameBoyAdvance"`
	NintendoGameBoyColor   Platform `json:"NintendoGameBoyColor"`
	NintendoGameCube       Platform `json:"NintendoGameCube"`
	NintendoGameWatch      Platform `json:"NintendoGameWatch"`
	NintendoDS             Platform `json:"NintendoDS"`
	NintendoDSi            Platform `json:"NintendoDSi"`
	NintendoPlayStation    Platform `json:"NintendoPlayStation"`
	NintendoFamicom        Platform `json:"NintendoFamicom"`
	NintendoSuperFamicom   Platform `json:"NintendoSuperFamicom"`
	NintendoNES            Platform `json:"NintendoNES"`
	NintendoSNES           Platform `json:"NintendoSNES"`
	NintendoSwitch         Platform `json:"NintendoSwitch"`
	NintendoWii            Platform `json:"NintendoWii"`
	NintendoWiiU           Platform `json:"NintendoWiiU"`

	// SEGA

	Sega32X              Platform `json:"Sega32X"`
	SegaDreamcast        Platform `json:"SegaDreamcast"`
	SegaGameGear         Platform `json:"SegaGameGear"`
	SegaGenesisMegadrive Platform `json:"SegaGenesisMegadrive"`
	SegaMasterSystem     Platform `json:"SegaMasterSystem"`
	SegaMegaCD           Platform `json:"SegaMegaCD"`
	SegaPico             Platform `json:"SegaPico"`
	SegaSaturn           Platform `json:"SegaSaturn"`

	// SONY

	PlayStation1        Platform `json:"PlayStation1"`
	PlayStation2        Platform `json:"PlayStation2"`
	PlayStation3        Platform `json:"PlayStation3"`
	PlayStation4        Platform `json:"PlayStation4"`
	PlayStation5        Platform `json:"PlayStation5"`
	PlayStationNow      Platform `json:"PlayStationNow"`
	PlayStationPortable Platform `json:"PlayStationPortable"`

	// TurboGrafx

	TurboGrafxCD Platform `json:"TurboGrafxCD"`
	TurboGrafx16 Platform `json:"TurboGrafx16"`
}

func (c Config) DefaultPlatformData() PlatformData {
	d := PlatformData{}

	d.HBMAME = Platform{true, "HBMAME", "hbmame"}
	d.MAME = Platform{true, "MAME", "mame"}

	d.Atari2600 = Platform{true, "Atari 2600", "atari2600"}
	d.Atari5200 = Platform{true, "Atari 5200", "atari5200"}
	d.Atari7800 = Platform{true, "Atari 7800", "atari7800"}
	d.Atari8Bit = Platform{true, "Atari 8-bit", "atari8bit"}
	d.AtariJaguar = Platform{true, "Atari Jaguar", "atarijaguar"}
	d.AtariLynx = Platform{true, "Atari Lynx", "atarilynx"}

	d.Xbox = Platform{true, "Xbox", "xbox"}
	d.Xbox360 = Platform{true, "Xbox 360", "xbox360"}
	d.XboxOne = Platform{true, "Xbox One", "xboxone"}
	d.XboxSeriesX = Platform{true, "Xbox Series X", "series-x"}

	d.Nintendo64 = Platform{true, "Nintendo 64", "n64"}
	d.NintendoGameBoy = Platform{true, "Nintendo Game Boy", "gb"}
	d.NintendoGameBoyAdvance = Platform{true, "Nintendo Game Boy Advance", "gba"}
	d.NintendoGameBoyColor = Platform{true, "Nintendo Game Boy Color", "gbc"}
	d.NintendoGameCube = Platform{true, "Nintendo GameCube", "ngc"}
	d.NintendoGameWatch = Platform{true, "Nintendo Game & Watch", "g-and-w"}
	d.NintendoDS = Platform{true, "Nintendo DS", "nds"}
	d.NintendoDSi = Platform{true, "Nintendo DSi", "nintendo-dsi"}
	d.NintendoPlayStation = Platform{true, "Nintendo PlayStation", "nintendo-playstation"}
	d.NintendoFamicom = Platform{true, "Nintendo Famicom", "famicom"}
	d.NintendoSuperFamicom = Platform{true, "Nintendo Super Famicom", "sfam"}
	d.NintendoNES = Platform{true, "Nintendo Entertainment System", "nes"}
	d.NintendoSNES = Platform{true, "Super Nintendo Entertainment System", "snes"}
	d.NintendoSwitch = Platform{true, "Nintendo Switch", "switch"}
	d.NintendoWii = Platform{true, "Nintendo Wii", "wii"}
	d.NintendoWiiU = Platform{true, "Nintendo WiiU", "wiiu"}

	d.Sega32X = Platform{true, "Sega 32X", "sega32"}
	d.SegaDreamcast = Platform{true, "Sega Dreamcast", "dc"}
	d.SegaGameGear = Platform{true, "Sega GameGear", "gamegear"}
	d.SegaGenesisMegadrive = Platform{true, "Sega Genesis/Megadrive", "genesis-slash-megadrive"}
	d.SegaMasterSystem = Platform{true, "Sega Master System", "sms"}
	d.SegaMegaCD = Platform{true, "Sega MegaCD", "segacd"}
	d.SegaPico = Platform{true, "Sega Pico", "sega-pico"}
	d.SegaSaturn = Platform{true, "Sega Saturn", "saturn"}

	d.PlayStation1 = Platform{true, "PlayStation", "psx"}
	d.PlayStation2 = Platform{true, "PlayStation2", "ps2"}
	d.PlayStation3 = Platform{true, "PlayStation3", "ps3"}
	d.PlayStation4 = Platform{true, "PlayStation4", "ps4"}
	d.PlayStation5 = Platform{true, "PlayStation5", "ps5"}
	d.PlayStationNow = Platform{true, "PlayStation Now", "psnow"}
	d.PlayStationPortable = Platform{true, "PlayStation Portable", "psp"}

	d.TurboGrafxCD = Platform{true, "TurboGrafx CD", "turbografxcd"}
	d.TurboGrafx16 = Platform{true, "TurboGrafx-16", "turbografx16"}

	return d
}
