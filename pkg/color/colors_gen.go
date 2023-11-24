package color

import "strings"

var (
	FrenchPuce                     = Color{R: 0.306, G: 0.086, B: 0.035, A: 1}
	OldCopper                      = Color{R: 0.447, G: 0.290, B: 0.184, A: 1}
	Lust                           = Color{R: 0.902, G: 0.125, B: 0.125, A: 1}
	Watusi                         = Color{R: 1.000, G: 0.867, B: 0.812, A: 1}
	YukonGold                      = Color{R: 0.482, G: 0.400, B: 0.031, A: 1}
	Limerick                       = Color{R: 0.616, G: 0.761, B: 0.035, A: 1}
	PeachCream                     = Color{R: 1.000, G: 0.941, B: 0.859, A: 1}
	VividAmber                     = Color{R: 0.800, G: 0.600, B: 0.000, A: 1}
	Madras                         = Color{R: 0.247, G: 0.188, B: 0.008, A: 1}
	VividGamboge                   = Color{R: 1.000, G: 0.600, B: 0.000, A: 1}
	Bianca                         = Color{R: 0.988, G: 0.984, B: 0.953, A: 1}
	Celadon                        = Color{R: 0.675, G: 0.882, B: 0.686, A: 1}
	Chambray                       = Color{R: 0.208, G: 0.306, B: 0.549, A: 1}
	DarkPastelBlue                 = Color{R: 0.467, G: 0.620, B: 0.796, A: 1}
	GreenSpring                    = Color{R: 0.722, G: 0.757, B: 0.694, A: 1}
	Salem                          = Color{R: 0.035, G: 0.498, B: 0.294, A: 1}
	AquaForest                     = Color{R: 0.373, G: 0.655, B: 0.467, A: 1}
	Celeste                        = Color{R: 0.698, G: 1.000, B: 1.000, A: 1}
	FuelYellow                     = Color{R: 0.925, G: 0.663, B: 0.153, A: 1}
	LightMediumOrchid              = Color{R: 0.827, G: 0.608, B: 0.796, A: 1}
	Matterhorn                     = Color{R: 0.306, G: 0.231, B: 0.255, A: 1}
	CyanCornflowerBlue             = Color{R: 0.094, G: 0.545, B: 0.761, A: 1}
	LightSalmonPink                = Color{R: 1.000, G: 0.600, B: 0.600, A: 1}
	Niagara                        = Color{R: 0.024, G: 0.631, B: 0.537, A: 1}
	Rainee                         = Color{R: 0.725, G: 0.784, B: 0.675, A: 1}
	SanJuan                        = Color{R: 0.188, G: 0.294, B: 0.416, A: 1}
	Cactus                         = Color{R: 0.345, G: 0.443, B: 0.337, A: 1}
	Harlequin                      = Color{R: 0.247, G: 1.000, B: 0.000, A: 1}
	OffYellow                      = Color{R: 0.996, G: 0.976, B: 0.890, A: 1}
	UPForestGreen                  = Color{R: 0.004, G: 0.267, B: 0.129, A: 1}
	PsychedelicPurple              = Color{R: 0.875, G: 0.000, B: 1.000, A: 1}
	BlackOlive                     = Color{R: 0.231, G: 0.235, B: 0.212, A: 1}
	CopperCanyon                   = Color{R: 0.494, G: 0.227, B: 0.082, A: 1}
	Illusion                       = Color{R: 0.965, G: 0.643, B: 0.788, A: 1}
	Onyx                           = Color{R: 0.208, G: 0.220, B: 0.224, A: 1}
	Perfume                        = Color{R: 0.816, G: 0.745, B: 0.973, A: 1}
	DeepPuce                       = Color{R: 0.663, G: 0.361, B: 0.408, A: 1}
	Emperor                        = Color{R: 0.318, G: 0.275, B: 0.286, A: 1}
	TrendyPink                     = Color{R: 0.549, G: 0.392, B: 0.584, A: 1}
	Turquoise                      = Color{R: 0.251, G: 0.878, B: 0.816, A: 1}
	Black                          = Color{R: 0.000, G: 0.000, B: 0.000, A: 1}
	BlackWhite                     = Color{R: 1.000, G: 0.996, B: 0.965, A: 1}
	CastletonGreen                 = Color{R: 0.000, G: 0.337, B: 0.231, A: 1}
	Mercury                        = Color{R: 0.898, G: 0.898, B: 0.898, A: 1}
	Merlot                         = Color{R: 0.514, G: 0.098, B: 0.137, A: 1}
	Chicago                        = Color{R: 0.365, G: 0.361, B: 0.345, A: 1}
	CottonSeed                     = Color{R: 0.761, G: 0.741, B: 0.714, A: 1}
	DarkRed                        = Color{R: 0.545, G: 0.000, B: 0.000, A: 1}
	DarkScarlet                    = Color{R: 0.337, G: 0.012, B: 0.098, A: 1}
	DeepMaroon                     = Color{R: 0.510, G: 0.000, B: 0.000, A: 1}
	DarkPuce                       = Color{R: 0.310, G: 0.227, B: 0.235, A: 1}
	WaterLeaf                      = Color{R: 0.631, G: 0.914, B: 0.871, A: 1}
	LightCarminePink               = Color{R: 0.902, G: 0.404, B: 0.443, A: 1}
	SkyMagenta                     = Color{R: 0.812, G: 0.443, B: 0.686, A: 1}
	DeepOak                        = Color{R: 0.255, G: 0.125, B: 0.063, A: 1}
	DoublePearlLusta               = Color{R: 0.988, G: 0.957, B: 0.816, A: 1}
	GrannyApple                    = Color{R: 0.835, G: 0.965, B: 0.890, A: 1}
	PinkOrange                     = Color{R: 1.000, G: 0.600, B: 0.400, A: 1}
	TonysPink                      = Color{R: 0.906, G: 0.624, B: 0.549, A: 1}
	Bud                            = Color{R: 0.659, G: 0.682, B: 0.612, A: 1}
	DeepBlue                       = Color{R: 0.133, G: 0.031, B: 0.471, A: 1}
	Folly                          = Color{R: 1.000, G: 0.000, B: 0.310, A: 1}
	TiffanyBlue                    = Color{R: 0.039, G: 0.729, B: 0.710, A: 1}
	Envy                           = Color{R: 0.545, G: 0.651, B: 0.565, A: 1}
	Horizon                        = Color{R: 0.353, G: 0.529, B: 0.627, A: 1}
	KellyGreen                     = Color{R: 0.298, G: 0.733, B: 0.090, A: 1}
	Muesli                         = Color{R: 0.667, G: 0.545, B: 0.357, A: 1}
	Stromboli                      = Color{R: 0.196, G: 0.365, B: 0.322, A: 1}
	BrunswickGreen                 = Color{R: 0.106, G: 0.302, B: 0.243, A: 1}
	KUCrimson                      = Color{R: 0.910, G: 0.000, B: 0.051, A: 1}
	DarkSkyBlue                    = Color{R: 0.549, G: 0.745, B: 0.839, A: 1}
	PalatinatePurple               = Color{R: 0.408, G: 0.157, B: 0.376, A: 1}
	Tuscany                        = Color{R: 0.753, G: 0.600, B: 0.600, A: 1}
	CadmiumRed                     = Color{R: 0.890, G: 0.000, B: 0.133, A: 1}
	FrenchViolet                   = Color{R: 0.533, G: 0.024, B: 0.808, A: 1}
	RoyalPurple                    = Color{R: 0.471, G: 0.318, B: 0.663, A: 1}
	Shamrock                       = Color{R: 0.200, G: 0.800, B: 0.600, A: 1}
	MoroccoBrown                   = Color{R: 0.267, G: 0.114, B: 0.000, A: 1}
	TowerGray                      = Color{R: 0.663, G: 0.741, B: 0.749, A: 1}
	Volt                           = Color{R: 0.808, G: 1.000, B: 0.000, A: 1}
	CoralReef                      = Color{R: 0.780, G: 0.737, B: 0.635, A: 1}
	Lemon                          = Color{R: 1.000, G: 0.969, B: 0.000, A: 1}
	Lynch                          = Color{R: 0.412, G: 0.494, B: 0.604, A: 1}
	MetallicCopper                 = Color{R: 0.443, G: 0.161, B: 0.114, A: 1}
	Monsoon                        = Color{R: 0.541, G: 0.514, B: 0.537, A: 1}
	BahamaBlue                     = Color{R: 0.008, G: 0.388, B: 0.584, A: 1}
	CadmiumOrange                  = Color{R: 0.929, G: 0.529, B: 0.176, A: 1}
	Shiraz                         = Color{R: 0.698, G: 0.035, B: 0.192, A: 1}
	BrightMaroon                   = Color{R: 0.765, G: 0.129, B: 0.282, A: 1}
	DogwoodRose                    = Color{R: 0.843, G: 0.094, B: 0.408, A: 1}
	FrenchRose                     = Color{R: 0.965, G: 0.290, B: 0.541, A: 1}
	Picasso                        = Color{R: 1.000, G: 0.953, B: 0.616, A: 1}
	UltramarineBlue                = Color{R: 0.255, G: 0.400, B: 0.961, A: 1}
	HippieGreen                    = Color{R: 0.325, G: 0.510, B: 0.294, A: 1}
	Iron                           = Color{R: 0.831, G: 0.843, B: 0.851, A: 1}
	OrangeYellow                   = Color{R: 0.973, G: 0.835, B: 0.408, A: 1}
	PaleSky                        = Color{R: 0.431, G: 0.467, B: 0.514, A: 1}
	CeladonGreen                   = Color{R: 0.184, G: 0.518, B: 0.486, A: 1}
	Feldgrau                       = Color{R: 0.302, G: 0.365, B: 0.325, A: 1}
	PearlAqua                      = Color{R: 0.533, G: 0.847, B: 0.753, A: 1}
	AbsoluteZero                   = Color{R: 0.000, G: 0.282, B: 0.729, A: 1}
	Akaroa                         = Color{R: 0.831, G: 0.769, B: 0.659, A: 1}
	Alpine                         = Color{R: 0.686, G: 0.561, B: 0.173, A: 1}
	BondiBlue                      = Color{R: 0.000, G: 0.584, B: 0.714, A: 1}
	Cadet                          = Color{R: 0.325, G: 0.408, B: 0.447, A: 1}
	Remy                           = Color{R: 0.996, G: 0.922, B: 0.953, A: 1}
	RosyBrown                      = Color{R: 0.737, G: 0.561, B: 0.561, A: 1}
	Tarawera                       = Color{R: 0.027, G: 0.227, B: 0.314, A: 1}
	Taupe                          = Color{R: 0.282, G: 0.235, B: 0.196, A: 1}
	LilyWhite                      = Color{R: 0.906, G: 0.973, B: 1.000, A: 1}
	Abbey                          = Color{R: 0.298, G: 0.310, B: 0.337, A: 1}
	AquaDeep                       = Color{R: 0.004, G: 0.294, B: 0.263, A: 1}
	Coconut                        = Color{R: 0.588, G: 0.353, B: 0.243, A: 1}
	LanguidLavender                = Color{R: 0.839, G: 0.792, B: 0.867, A: 1}
	LavenderMagenta                = Color{R: 0.933, G: 0.510, B: 0.933, A: 1}
	WhiteIce                       = Color{R: 0.867, G: 0.976, B: 0.945, A: 1}
	Acadia                         = Color{R: 0.106, G: 0.078, B: 0.016, A: 1}
	Cyclamen                       = Color{R: 0.961, G: 0.435, B: 0.631, A: 1}
	Glacier                        = Color{R: 0.502, G: 0.702, B: 0.769, A: 1}
	LightSteelBlue                 = Color{R: 0.690, G: 0.769, B: 0.871, A: 1}
	Trinidad                       = Color{R: 0.902, G: 0.306, B: 0.012, A: 1}
	MilkPunch                      = Color{R: 1.000, G: 0.965, B: 0.831, A: 1}
	VanillaIce                     = Color{R: 0.953, G: 0.561, B: 0.663, A: 1}
	Casablanca                     = Color{R: 0.973, G: 0.722, B: 0.325, A: 1}
	DoveGray                       = Color{R: 0.427, G: 0.424, B: 0.424, A: 1}
	FairPink                       = Color{R: 1.000, G: 0.937, B: 0.925, A: 1}
	Finn                           = Color{R: 0.412, G: 0.176, B: 0.329, A: 1}
	Isabelline                     = Color{R: 0.957, G: 0.941, B: 0.925, A: 1}
	Driftwood                      = Color{R: 0.686, G: 0.529, B: 0.318, A: 1}
	PaleCanary                     = Color{R: 1.000, G: 1.000, B: 0.600, A: 1}
	PigmentGreen                   = Color{R: 0.000, G: 0.647, B: 0.314, A: 1}
	CapeHoney                      = Color{R: 0.996, G: 0.898, B: 0.675, A: 1}
	Imperial                       = Color{R: 0.376, G: 0.184, B: 0.420, A: 1}
	OuterSpace                     = Color{R: 0.255, G: 0.290, B: 0.298, A: 1}
	PantoneMagenta                 = Color{R: 0.816, G: 0.255, B: 0.494, A: 1}
	Cadillac                       = Color{R: 0.690, G: 0.298, B: 0.416, A: 1}
	ElfGreen                       = Color{R: 0.031, G: 0.514, B: 0.439, A: 1}
	FernFrond                      = Color{R: 0.396, G: 0.447, B: 0.125, A: 1}
	Lotus                          = Color{R: 0.525, G: 0.235, B: 0.235, A: 1}
	SunsetOrange                   = Color{R: 0.992, G: 0.369, B: 0.325, A: 1}
	Amethyst                       = Color{R: 0.600, G: 0.400, B: 0.800, A: 1}
	Biscay                         = Color{R: 0.106, G: 0.192, B: 0.384, A: 1}
	DarkTan                        = Color{R: 0.569, G: 0.506, B: 0.318, A: 1}
	UniversityOfCaliforniaGold     = Color{R: 0.718, G: 0.529, B: 0.153, A: 1}
	UpsdellRed                     = Color{R: 0.682, G: 0.125, B: 0.161, A: 1}
	Scarlet                        = Color{R: 1.000, G: 0.141, B: 0.000, A: 1}
	RockBlue                       = Color{R: 0.620, G: 0.694, B: 0.804, A: 1}
	Wheat                          = Color{R: 0.961, G: 0.871, B: 0.702, A: 1}
	AlizarinCrimson                = Color{R: 0.890, G: 0.149, B: 0.212, A: 1}
	BlueGem                        = Color{R: 0.173, G: 0.055, B: 0.549, A: 1}
	Carnelian                      = Color{R: 0.702, G: 0.106, B: 0.106, A: 1}
	LightApricot                   = Color{R: 0.992, G: 0.835, B: 0.694, A: 1}
	PeachYellow                    = Color{R: 0.980, G: 0.875, B: 0.678, A: 1}
	MummysTomb                     = Color{R: 0.510, G: 0.557, B: 0.518, A: 1}
	WildOrchid                     = Color{R: 0.831, G: 0.439, B: 0.635, A: 1}
	Cola                           = Color{R: 0.247, G: 0.145, B: 0.000, A: 1}
	RebeccaPurple                  = Color{R: 0.400, G: 0.200, B: 0.600, A: 1}
	TwilightBlue                   = Color{R: 0.933, G: 0.992, B: 1.000, A: 1}
	Brown                          = Color{R: 0.588, G: 0.294, B: 0.000, A: 1}
	DebianRed                      = Color{R: 0.843, G: 0.039, B: 0.325, A: 1}
	GreenBlue                      = Color{R: 0.067, G: 0.392, B: 0.706, A: 1}
	PirateGold                     = Color{R: 0.729, G: 0.498, B: 0.012, A: 1}
	RYBViolet                      = Color{R: 0.525, G: 0.004, B: 0.686, A: 1}
	OldGold                        = Color{R: 0.812, G: 0.710, B: 0.231, A: 1}
	SummerGreen                    = Color{R: 0.588, G: 0.733, B: 0.671, A: 1}
	Carla                          = Color{R: 0.953, G: 1.000, B: 0.847, A: 1}
	GoldenTainoi                   = Color{R: 1.000, G: 0.800, B: 0.361, A: 1}
	Harp                           = Color{R: 0.902, G: 0.949, B: 0.918, A: 1}
	MaximumYellow                  = Color{R: 0.980, G: 0.980, B: 0.216, A: 1}
	MunsellGreen                   = Color{R: 0.000, G: 0.659, B: 0.467, A: 1}
	Viola                          = Color{R: 0.796, G: 0.561, B: 0.663, A: 1}
	Waterloo                       = Color{R: 0.482, G: 0.486, B: 0.580, A: 1}
	Yuma                           = Color{R: 0.808, G: 0.761, B: 0.569, A: 1}
	EnglishWalnut                  = Color{R: 0.243, G: 0.169, B: 0.137, A: 1}
	Roti                           = Color{R: 0.776, G: 0.659, B: 0.294, A: 1}
	DavysGrey                      = Color{R: 0.333, G: 0.333, B: 0.333, A: 1}
	ElSalva                        = Color{R: 0.561, G: 0.243, B: 0.200, A: 1}
	GoldenYellow                   = Color{R: 1.000, G: 0.875, B: 0.000, A: 1}
	VanDykeBrown                   = Color{R: 0.400, G: 0.259, B: 0.157, A: 1}
	Beige                          = Color{R: 0.961, G: 0.961, B: 0.863, A: 1}
	Paco                           = Color{R: 0.255, G: 0.122, B: 0.063, A: 1}
	Peanut                         = Color{R: 0.471, G: 0.184, B: 0.086, A: 1}
	Salmon                         = Color{R: 0.980, G: 0.502, B: 0.447, A: 1}
	TuscanRed                      = Color{R: 0.486, G: 0.282, B: 0.282, A: 1}
	SpicyPink                      = Color{R: 0.506, G: 0.431, B: 0.443, A: 1}
	StarCommandBlue                = Color{R: 0.000, G: 0.482, B: 0.722, A: 1}
	Flax                           = Color{R: 0.933, G: 0.863, B: 0.510, A: 1}
	PrussianBlue                   = Color{R: 0.000, G: 0.192, B: 0.325, A: 1}
	RobinEggBlue                   = Color{R: 0.000, G: 0.800, B: 0.800, A: 1}
	Rum                            = Color{R: 0.475, G: 0.412, B: 0.537, A: 1}
	SpaceCadet                     = Color{R: 0.114, G: 0.161, B: 0.318, A: 1}
	Conifer                        = Color{R: 0.675, G: 0.867, B: 0.302, A: 1}
	DeepCarrotOrange               = Color{R: 0.914, G: 0.412, B: 0.173, A: 1}
	Flirt                          = Color{R: 0.635, G: 0.000, B: 0.427, A: 1}
	GordonsGreen                   = Color{R: 0.043, G: 0.067, B: 0.027, A: 1}
	MunsellPurple                  = Color{R: 0.624, G: 0.000, B: 0.773, A: 1}
	ParadisePink                   = Color{R: 0.902, G: 0.243, B: 0.384, A: 1}
	Prelude                        = Color{R: 0.816, G: 0.753, B: 0.898, A: 1}
	RoyalAzure                     = Color{R: 0.000, G: 0.220, B: 0.659, A: 1}
	DonkeyBrown                    = Color{R: 0.400, G: 0.298, B: 0.157, A: 1}
	FuchsiaPink                    = Color{R: 1.000, G: 0.467, B: 1.000, A: 1}
	Ghost                          = Color{R: 0.780, G: 0.788, B: 0.835, A: 1}
	IceCold                        = Color{R: 0.694, G: 0.957, B: 0.906, A: 1}
	Karry                          = Color{R: 1.000, G: 0.918, B: 0.831, A: 1}
	DeepForestGreen                = Color{R: 0.094, G: 0.176, B: 0.035, A: 1}
	PhthaloBlue                    = Color{R: 0.000, G: 0.059, B: 0.537, A: 1}
	RadicalRed                     = Color{R: 1.000, G: 0.208, B: 0.369, A: 1}
	Tusk                           = Color{R: 0.933, G: 0.953, B: 0.765, A: 1}
	Bluebonnet                     = Color{R: 0.110, G: 0.110, B: 0.941, A: 1}
	PineGlade                      = Color{R: 0.780, G: 0.804, B: 0.565, A: 1}
	UnitedNationsBlue              = Color{R: 0.357, G: 0.573, B: 0.898, A: 1}
	WitchHaze                      = Color{R: 1.000, G: 0.988, B: 0.600, A: 1}
	Dolphin                        = Color{R: 0.392, G: 0.376, B: 0.467, A: 1}
	FOGRA29RichBlack               = Color{R: 0.004, G: 0.043, B: 0.075, A: 1}
	Melrose                        = Color{R: 0.780, G: 0.757, B: 1.000, A: 1}
	Scooter                        = Color{R: 0.180, G: 0.749, B: 0.831, A: 1}
	VividRaspberry                 = Color{R: 1.000, G: 0.000, B: 0.424, A: 1}
	BrownTumbleweed                = Color{R: 0.216, G: 0.161, B: 0.055, A: 1}
	FieryRose                      = Color{R: 1.000, G: 0.329, B: 0.439, A: 1}
	JazzberryJam                   = Color{R: 0.647, G: 0.043, B: 0.369, A: 1}
	Ruddy                          = Color{R: 1.000, G: 0.000, B: 0.157, A: 1}
	BronzeYellow                   = Color{R: 0.451, G: 0.439, B: 0.000, A: 1}
	CarmineRed                     = Color{R: 1.000, G: 0.000, B: 0.220, A: 1}
	FountainBlue                   = Color{R: 0.337, G: 0.706, B: 0.745, A: 1}
	TickleMePink                   = Color{R: 0.988, G: 0.537, B: 0.675, A: 1}
	TurquoiseBlue                  = Color{R: 0.000, G: 1.000, B: 0.937, A: 1}
	IlluminatingEmerald            = Color{R: 0.192, G: 0.569, B: 0.467, A: 1}
	RomanCoffee                    = Color{R: 0.475, G: 0.365, B: 0.298, A: 1}
	Ceil                           = Color{R: 0.573, G: 0.631, B: 0.812, A: 1}
	CopperRed                      = Color{R: 0.796, G: 0.427, B: 0.318, A: 1}
	Marigold                       = Color{R: 0.918, G: 0.635, B: 0.129, A: 1}
	Yellow                         = Color{R: 1.000, G: 1.000, B: 0.000, A: 1}
	YaleBlue                       = Color{R: 0.059, G: 0.302, B: 0.573, A: 1}
	Astra                          = Color{R: 0.980, G: 0.918, B: 0.725, A: 1}
	BrownPod                       = Color{R: 0.251, G: 0.094, B: 0.004, A: 1}
	MediumPurple                   = Color{R: 0.576, G: 0.439, B: 0.859, A: 1}
	SmokyTopaz                     = Color{R: 0.576, G: 0.239, B: 0.255, A: 1}
	SnowFlurry                     = Color{R: 0.894, G: 1.000, B: 0.820, A: 1}
	PaoloVeroneseGreen             = Color{R: 0.000, G: 0.608, B: 0.490, A: 1}
	PiggyPink                      = Color{R: 0.992, G: 0.867, B: 0.902, A: 1}
	PineTree                       = Color{R: 0.090, G: 0.122, B: 0.016, A: 1}
	Piper                          = Color{R: 0.788, G: 0.388, B: 0.137, A: 1}
	DeepSpaceSparkle               = Color{R: 0.290, G: 0.392, B: 0.424, A: 1}
	Jonquil                        = Color{R: 0.957, G: 0.792, B: 0.086, A: 1}
	Toledo                         = Color{R: 0.227, G: 0.000, B: 0.125, A: 1}
	VividMalachite                 = Color{R: 0.000, G: 0.800, B: 0.200, A: 1}
	Russet                         = Color{R: 0.502, G: 0.275, B: 0.106, A: 1}
	Solitude                       = Color{R: 0.918, G: 0.965, B: 1.000, A: 1}
	Toolbox                        = Color{R: 0.455, G: 0.424, B: 0.753, A: 1}
	TurkishRose                    = Color{R: 0.710, G: 0.447, B: 0.506, A: 1}
	CrimsonRed                     = Color{R: 0.600, G: 0.000, B: 0.000, A: 1}
	DarkBlueGray                   = Color{R: 0.400, G: 0.400, B: 0.600, A: 1}
	DarkTerraCotta                 = Color{R: 0.800, G: 0.306, B: 0.361, A: 1}
	MySin                          = Color{R: 1.000, G: 0.702, B: 0.122, A: 1}
	QueenBlue                      = Color{R: 0.263, G: 0.420, B: 0.584, A: 1}
	MordantRed                     = Color{R: 0.682, G: 0.047, B: 0.000, A: 1}
	Seashell                       = Color{R: 1.000, G: 0.961, B: 0.933, A: 1}
	Spectra                        = Color{R: 0.184, G: 0.353, B: 0.341, A: 1}
	UPMaroon                       = Color{R: 0.482, G: 0.067, B: 0.075, A: 1}
	HavelockBlue                   = Color{R: 0.333, G: 0.565, B: 0.851, A: 1}
	OrangeWhite                    = Color{R: 0.996, G: 0.988, B: 0.929, A: 1}
	RossoCorsa                     = Color{R: 0.831, G: 0.000, B: 0.000, A: 1}
	Sundance                       = Color{R: 0.788, G: 0.702, B: 0.357, A: 1}
	Treehouse                      = Color{R: 0.231, G: 0.157, B: 0.125, A: 1}
	Jaguar                         = Color{R: 0.031, G: 0.004, B: 0.063, A: 1}
	PalmLeaf                       = Color{R: 0.098, G: 0.200, B: 0.055, A: 1}
	PastelGray                     = Color{R: 0.812, G: 0.812, B: 0.769, A: 1}
	TawnyPort                      = Color{R: 0.412, G: 0.145, B: 0.271, A: 1}
	VividOrchid                    = Color{R: 0.800, G: 0.000, B: 1.000, A: 1}
	AmaranthPurple                 = Color{R: 0.671, G: 0.153, B: 0.310, A: 1}
	Dixie                          = Color{R: 0.886, G: 0.580, B: 0.094, A: 1}
	FrostedMint                    = Color{R: 0.859, G: 1.000, B: 0.973, A: 1}
	RYBBlue                        = Color{R: 0.008, G: 0.278, B: 0.996, A: 1}
	Schooner                       = Color{R: 0.545, G: 0.518, B: 0.494, A: 1}
	Emerald                        = Color{R: 0.314, G: 0.784, B: 0.471, A: 1}
	Gainsboro                      = Color{R: 0.863, G: 0.863, B: 0.863, A: 1}
	OldLace                        = Color{R: 0.992, G: 0.961, B: 0.902, A: 1}
	AuChico                        = Color{R: 0.592, G: 0.376, B: 0.365, A: 1}
	CongoPink                      = Color{R: 0.973, G: 0.514, B: 0.475, A: 1}
	LemonYellow                    = Color{R: 1.000, G: 0.957, B: 0.310, A: 1}
	MaiTai                         = Color{R: 0.690, G: 0.400, B: 0.031, A: 1}
	MediumRuby                     = Color{R: 0.667, G: 0.251, B: 0.412, A: 1}
	Drover                         = Color{R: 0.992, G: 0.969, B: 0.678, A: 1}
	EnergyYellow                   = Color{R: 0.973, G: 0.867, B: 0.361, A: 1}
	GraniteGray                    = Color{R: 0.404, G: 0.404, B: 0.404, A: 1}
	SherwoodGreen                  = Color{R: 0.008, G: 0.251, B: 0.173, A: 1}
	AlgaeGreen                     = Color{R: 0.576, G: 0.875, B: 0.722, A: 1}
	InternationalKleinBlue         = Color{R: 0.000, G: 0.184, B: 0.655, A: 1}
	Shocking                       = Color{R: 0.886, G: 0.573, B: 0.753, A: 1}
	PaleMagenta                    = Color{R: 0.976, G: 0.518, B: 0.898, A: 1}
	PotPourri                      = Color{R: 0.961, G: 0.906, B: 0.886, A: 1}
	RobRoy                         = Color{R: 0.918, G: 0.776, B: 0.455, A: 1}
	Tangaroa                       = Color{R: 0.012, G: 0.086, B: 0.235, A: 1}
	Bronco                         = Color{R: 0.671, G: 0.631, B: 0.588, A: 1}
	Caper                          = Color{R: 0.863, G: 0.929, B: 0.706, A: 1}
	Dingley                        = Color{R: 0.365, G: 0.467, B: 0.278, A: 1}
	Ecru                           = Color{R: 0.761, G: 0.698, B: 0.502, A: 1}
	Matrix                         = Color{R: 0.690, G: 0.365, B: 0.329, A: 1}
	CornflowerBlue                 = Color{R: 0.392, G: 0.584, B: 0.929, A: 1}
	Frost                          = Color{R: 0.929, G: 0.961, B: 0.867, A: 1}
	Kidnapper                      = Color{R: 0.882, G: 0.918, B: 0.831, A: 1}
	RedBerry                       = Color{R: 0.557, G: 0.000, B: 0.000, A: 1}
	BlueHaze                       = Color{R: 0.749, G: 0.745, B: 0.847, A: 1}
	PortGore                       = Color{R: 0.145, G: 0.122, B: 0.310, A: 1}
	Logan                          = Color{R: 0.667, G: 0.663, B: 0.804, A: 1}
	PantoneOrange                  = Color{R: 1.000, G: 0.345, B: 0.000, A: 1}
	RedOxide                       = Color{R: 0.431, G: 0.035, B: 0.008, A: 1}
	YellowGreen                    = Color{R: 0.604, G: 0.804, B: 0.196, A: 1}
	Daffodil                       = Color{R: 1.000, G: 1.000, B: 0.192, A: 1}
	SaharaSand                     = Color{R: 0.945, G: 0.906, B: 0.533, A: 1}
	SmokeyTopaz                    = Color{R: 0.514, G: 0.165, B: 0.051, A: 1}
	TePapaGreen                    = Color{R: 0.118, G: 0.263, B: 0.235, A: 1}
	Vanilla                        = Color{R: 0.953, G: 0.898, B: 0.671, A: 1}
	Confetti                       = Color{R: 0.914, G: 0.843, B: 0.353, A: 1}
	Limeade                        = Color{R: 0.435, G: 0.616, B: 0.008, A: 1}
	MetallicSeaweed                = Color{R: 0.039, G: 0.494, B: 0.549, A: 1}
	SilverTree                     = Color{R: 0.400, G: 0.710, B: 0.561, A: 1}
	Turmeric                       = Color{R: 0.792, G: 0.733, B: 0.282, A: 1}
	Crail                          = Color{R: 0.725, G: 0.318, B: 0.251, A: 1}
	Malachite                      = Color{R: 0.043, G: 0.855, B: 0.318, A: 1}
	EnglishHolly                   = Color{R: 0.008, G: 0.176, B: 0.082, A: 1}
	Kokoda                         = Color{R: 0.431, G: 0.427, B: 0.341, A: 1}
	QuarterSpanishWhite            = Color{R: 0.969, G: 0.949, B: 0.882, A: 1}
	Reef                           = Color{R: 0.788, G: 1.000, B: 0.635, A: 1}
	Shalimar                       = Color{R: 0.984, G: 1.000, B: 0.729, A: 1}
	LavenderIndigo                 = Color{R: 0.580, G: 0.341, B: 0.922, A: 1}
	PantoneBlue                    = Color{R: 0.000, G: 0.094, B: 0.659, A: 1}
	SwampGreen                     = Color{R: 0.675, G: 0.718, B: 0.557, A: 1}
	Cinder                         = Color{R: 0.055, G: 0.055, B: 0.094, A: 1}
	SmashedPumpkin                 = Color{R: 1.000, G: 0.427, B: 0.227, A: 1}
	ChathamsBlue                   = Color{R: 0.090, G: 0.333, B: 0.475, A: 1}
	Lime                           = Color{R: 0.749, G: 1.000, B: 0.000, A: 1}
	Observatory                    = Color{R: 0.008, G: 0.525, B: 0.435, A: 1}
	PinkLavender                   = Color{R: 0.847, G: 0.698, B: 0.820, A: 1}
	Maize                          = Color{R: 0.984, G: 0.925, B: 0.365, A: 1}
	PapayaWhip                     = Color{R: 1.000, G: 0.937, B: 0.835, A: 1}
	SafetyOrange                   = Color{R: 1.000, G: 0.471, B: 0.000, A: 1}
	CarminePink                    = Color{R: 0.922, G: 0.298, B: 0.259, A: 1}
	DeepTeal                       = Color{R: 0.000, G: 0.208, B: 0.196, A: 1}
	Fern                           = Color{R: 0.388, G: 0.718, B: 0.424, A: 1}
	GreenYellow                    = Color{R: 0.678, G: 1.000, B: 0.184, A: 1}
	GulfStream                     = Color{R: 0.502, G: 0.702, B: 0.682, A: 1}
	CaribbeanGreen                 = Color{R: 0.000, G: 0.800, B: 0.600, A: 1}
	Cinderella                     = Color{R: 0.992, G: 0.882, B: 0.863, A: 1}
	Christine                      = Color{R: 0.906, G: 0.451, B: 0.039, A: 1}
	DarkChestnut                   = Color{R: 0.596, G: 0.412, B: 0.376, A: 1}
	Iris                           = Color{R: 0.353, G: 0.310, B: 0.812, A: 1}
	Tumbleweed                     = Color{R: 0.871, G: 0.667, B: 0.533, A: 1}
	BrownYellow                    = Color{R: 0.800, G: 0.600, B: 0.400, A: 1}
	Dorado                         = Color{R: 0.420, G: 0.341, B: 0.333, A: 1}
	LightPink                      = Color{R: 1.000, G: 0.714, B: 0.757, A: 1}
	TuftBush                       = Color{R: 1.000, G: 0.867, B: 0.804, A: 1}
	ClamShell                      = Color{R: 0.831, G: 0.714, B: 0.686, A: 1}
	Firebrick                      = Color{R: 0.698, G: 0.133, B: 0.133, A: 1}
	MetallicBronze                 = Color{R: 0.286, G: 0.216, B: 0.106, A: 1}
	PeachSchnapps                  = Color{R: 1.000, G: 0.863, B: 0.839, A: 1}
	SepiaBlack                     = Color{R: 0.169, G: 0.008, B: 0.008, A: 1}
	AquaHaze                       = Color{R: 0.929, G: 0.961, B: 0.961, A: 1}
	SpringLeaves                   = Color{R: 0.341, G: 0.514, B: 0.388, A: 1}
	VioletRed                      = Color{R: 0.969, G: 0.325, B: 0.580, A: 1}
	TealDeer                       = Color{R: 0.600, G: 0.902, B: 0.702, A: 1}
	Clover                         = Color{R: 0.220, G: 0.286, B: 0.063, A: 1}
	Coriander                      = Color{R: 0.769, G: 0.816, B: 0.690, A: 1}
	Dandelion                      = Color{R: 0.941, G: 0.882, B: 0.188, A: 1}
	MexicanPink                    = Color{R: 0.894, G: 0.000, B: 0.486, A: 1}
	RichLilac                      = Color{R: 0.714, G: 0.400, B: 0.824, A: 1}
	SanMarino                      = Color{R: 0.271, G: 0.424, B: 0.675, A: 1}
	Victoria                       = Color{R: 0.325, G: 0.267, B: 0.569, A: 1}
	CarnationPink                  = Color{R: 1.000, G: 0.651, B: 0.788, A: 1}
	DarkSlateGray                  = Color{R: 0.184, G: 0.310, B: 0.310, A: 1}
	Firefly                        = Color{R: 0.055, G: 0.165, B: 0.188, A: 1}
	FuzzyWuzzy                     = Color{R: 0.800, G: 0.400, B: 0.400, A: 1}
	RichGold                       = Color{R: 0.659, G: 0.325, B: 0.027, A: 1}
	PeruTan                        = Color{R: 0.498, G: 0.227, B: 0.008, A: 1}
	Romance                        = Color{R: 1.000, G: 0.996, B: 0.992, A: 1}
	BayLeaf                        = Color{R: 0.490, G: 0.663, B: 0.553, A: 1}
	Bombay                         = Color{R: 0.686, G: 0.694, B: 0.722, A: 1}
	Citrus                         = Color{R: 0.631, G: 0.773, B: 0.039, A: 1}
	CongoBrown                     = Color{R: 0.349, G: 0.216, B: 0.216, A: 1}
	HalfandHalf                    = Color{R: 1.000, G: 0.996, B: 0.882, A: 1}
	RuddyBrown                     = Color{R: 0.733, G: 0.396, B: 0.157, A: 1}
	Geyser                         = Color{R: 0.831, G: 0.875, B: 0.886, A: 1}
	IslamicGreen                   = Color{R: 0.000, G: 0.565, B: 0.000, A: 1}
	PaleRedViolet                  = Color{R: 0.859, G: 0.439, B: 0.576, A: 1}
	SorrellBrown                   = Color{R: 0.808, G: 0.725, B: 0.561, A: 1}
	RoseBonbon                     = Color{R: 0.976, G: 0.259, B: 0.620, A: 1}
	LightCobaltBlue                = Color{R: 0.533, G: 0.675, B: 0.878, A: 1}
	SpanishPink                    = Color{R: 0.969, G: 0.749, B: 0.745, A: 1}
	Ziggurat                       = Color{R: 0.749, G: 0.859, B: 0.886, A: 1}
	YourPink                       = Color{R: 1.000, G: 0.765, B: 0.753, A: 1}
	Amber                          = Color{R: 1.000, G: 0.749, B: 0.000, A: 1}
	Beeswax                        = Color{R: 0.996, G: 0.949, B: 0.780, A: 1}
	Honeydew                       = Color{R: 0.941, G: 1.000, B: 0.941, A: 1}
	Quartz                         = Color{R: 0.318, G: 0.282, B: 0.310, A: 1}
	Silver                         = Color{R: 0.753, G: 0.753, B: 0.753, A: 1}
	AlienArmpit                    = Color{R: 0.518, G: 0.871, B: 0.008, A: 1}
	AquaSpring                     = Color{R: 0.918, G: 0.976, B: 0.961, A: 1}
	DarkSalmon                     = Color{R: 0.914, G: 0.588, B: 0.478, A: 1}
	DeepTaupe                      = Color{R: 0.494, G: 0.369, B: 0.376, A: 1}
	GrayOlive                      = Color{R: 0.663, G: 0.643, B: 0.569, A: 1}
	Crusta                         = Color{R: 0.992, G: 0.482, B: 0.200, A: 1}
	WhitePointer                   = Color{R: 0.996, G: 0.973, B: 1.000, A: 1}
	Burlywood                      = Color{R: 0.871, G: 0.722, B: 0.529, A: 1}
	Tamarind                       = Color{R: 0.204, G: 0.082, B: 0.082, A: 1}
	DeepTuscanRed                  = Color{R: 0.400, G: 0.259, B: 0.302, A: 1}
	GraniteGreen                   = Color{R: 0.553, G: 0.537, B: 0.455, A: 1}
	MediumElectricBlue             = Color{R: 0.012, G: 0.314, B: 0.588, A: 1}
	PastelOrange                   = Color{R: 1.000, G: 0.702, B: 0.278, A: 1}
	BostonUniversityRed            = Color{R: 0.800, G: 0.000, B: 0.000, A: 1}
	BreakerBay                     = Color{R: 0.365, G: 0.631, B: 0.624, A: 1}
	Buttercup                      = Color{R: 0.953, G: 0.678, B: 0.086, A: 1}
	BeautyBush                     = Color{R: 0.933, G: 0.757, B: 0.745, A: 1}
	Laurel                         = Color{R: 0.455, G: 0.576, B: 0.471, A: 1}
	PictonBlue                     = Color{R: 0.271, G: 0.694, B: 0.910, A: 1}
	Strawberry                     = Color{R: 0.988, G: 0.353, B: 0.553, A: 1}
	TobaccoBrown                   = Color{R: 0.443, G: 0.365, B: 0.278, A: 1}
	BrandeisBlue                   = Color{R: 0.000, G: 0.439, B: 1.000, A: 1}
	HippiePink                     = Color{R: 0.682, G: 0.271, B: 0.376, A: 1}
	Hampton                        = Color{R: 0.898, G: 0.847, B: 0.686, A: 1}
	AlbescentWhite                 = Color{R: 0.961, G: 0.914, B: 0.827, A: 1}
	Disco                          = Color{R: 0.529, G: 0.082, B: 0.314, A: 1}
	FadedJade                      = Color{R: 0.259, G: 0.475, B: 0.467, A: 1}
	RipeLemon                      = Color{R: 0.957, G: 0.847, B: 0.110, A: 1}
	Honeysuckle                    = Color{R: 0.929, G: 0.988, B: 0.518, A: 1}
	Merino                         = Color{R: 0.965, G: 0.941, B: 0.902, A: 1}
	WebChartreuse                  = Color{R: 0.498, G: 1.000, B: 0.000, A: 1}
	Eden                           = Color{R: 0.063, G: 0.345, B: 0.322, A: 1}
	EnglishRed                     = Color{R: 0.671, G: 0.294, B: 0.322, A: 1}
	Submarine                      = Color{R: 0.729, G: 0.780, B: 0.788, A: 1}
	TanHide                        = Color{R: 0.980, G: 0.616, B: 0.353, A: 1}
	Coquelicot                     = Color{R: 1.000, G: 0.220, B: 0.000, A: 1}
	DeepChestnut                   = Color{R: 0.725, G: 0.306, B: 0.282, A: 1}
	MSUGreen                       = Color{R: 0.094, G: 0.271, B: 0.231, A: 1}
	NutmegWoodFinish               = Color{R: 0.408, G: 0.212, B: 0.000, A: 1}
	ZinnwalditeBrown               = Color{R: 0.173, G: 0.086, B: 0.031, A: 1}
	AzureMist                      = Color{R: 0.941, G: 1.000, B: 1.000, A: 1}
	Cordovan                       = Color{R: 0.537, G: 0.247, B: 0.271, A: 1}
	ElectricCrimson                = Color{R: 1.000, G: 0.000, B: 0.247, A: 1}
	GoldenFizz                     = Color{R: 0.961, G: 0.984, B: 0.239, A: 1}
	Oracle                         = Color{R: 0.216, G: 0.455, B: 0.459, A: 1}
	HeatWave                       = Color{R: 1.000, G: 0.478, B: 0.000, A: 1}
	MetallicSunburst               = Color{R: 0.612, G: 0.486, B: 0.220, A: 1}
	MineralGreen                   = Color{R: 0.247, G: 0.365, B: 0.325, A: 1}
	PantonePink                    = Color{R: 0.843, G: 0.282, B: 0.580, A: 1}
	BrilliantAzure                 = Color{R: 0.200, G: 0.600, B: 1.000, A: 1}
	Frostee                        = Color{R: 0.894, G: 0.965, B: 0.906, A: 1}
	Revolver                       = Color{R: 0.173, G: 0.086, B: 0.196, A: 1}
	CabSav                         = Color{R: 0.302, G: 0.039, B: 0.094, A: 1}
	LavenderPurple                 = Color{R: 0.588, G: 0.482, B: 0.714, A: 1}
	PinkFlare                      = Color{R: 0.882, G: 0.753, B: 0.784, A: 1}
	SaltBox                        = Color{R: 0.408, G: 0.369, B: 0.431, A: 1}
	SanFelix                       = Color{R: 0.043, G: 0.384, B: 0.027, A: 1}
	SeaGreen                       = Color{R: 0.180, G: 0.545, B: 0.341, A: 1}
	BerylGreen                     = Color{R: 0.871, G: 0.898, B: 0.753, A: 1}
	BlueDiamond                    = Color{R: 0.220, G: 0.016, B: 0.455, A: 1}
	Cerise                         = Color{R: 0.871, G: 0.192, B: 0.388, A: 1}
	Grape                          = Color{R: 0.435, G: 0.176, B: 0.659, A: 1}
	Kimberly                       = Color{R: 0.451, G: 0.424, B: 0.624, A: 1}
	SilverPink                     = Color{R: 0.769, G: 0.682, B: 0.678, A: 1}
	Verdigris                      = Color{R: 0.263, G: 0.702, B: 0.682, A: 1}
	Wine                           = Color{R: 0.447, G: 0.184, B: 0.216, A: 1}
	DeepSea                        = Color{R: 0.004, G: 0.510, B: 0.420, A: 1}
	FashionFuchsia                 = Color{R: 0.957, G: 0.000, B: 0.631, A: 1}
	HintofRed                      = Color{R: 0.984, G: 0.976, B: 0.976, A: 1}
	Jambalaya                      = Color{R: 0.357, G: 0.188, B: 0.075, A: 1}
	LightKhaki                     = Color{R: 0.941, G: 0.902, B: 0.549, A: 1}
	Viridian                       = Color{R: 0.251, G: 0.510, B: 0.427, A: 1}
	VividCrimson                   = Color{R: 0.800, G: 0.000, B: 0.200, A: 1}
	AshGrey                        = Color{R: 0.698, G: 0.745, B: 0.710, A: 1}
	CongressBlue                   = Color{R: 0.008, G: 0.278, B: 0.557, A: 1}
	Gallery                        = Color{R: 0.937, G: 0.937, B: 0.937, A: 1}
	Vermilion                      = Color{R: 0.851, G: 0.220, B: 0.118, A: 1}
	ViolentViolet                  = Color{R: 0.161, G: 0.047, B: 0.369, A: 1}
	DarkBlue                       = Color{R: 0.000, G: 0.000, B: 0.545, A: 1}
	Gigas                          = Color{R: 0.322, G: 0.235, B: 0.580, A: 1}
	GoldenBell                     = Color{R: 0.886, G: 0.537, B: 0.075, A: 1}
	SonicSilver                    = Color{R: 0.459, G: 0.459, B: 0.459, A: 1}
	BoogerBuster                   = Color{R: 0.867, G: 0.886, B: 0.416, A: 1}
	DarkPink                       = Color{R: 0.906, G: 0.329, B: 0.502, A: 1}
	ThulianPink                    = Color{R: 0.871, G: 0.435, B: 0.631, A: 1}
	VividViolet                    = Color{R: 0.624, G: 0.000, B: 1.000, A: 1}
	Wasabi                         = Color{R: 0.471, G: 0.541, B: 0.145, A: 1}
	MoonstoneBlue                  = Color{R: 0.451, G: 0.663, B: 0.761, A: 1}
	Tana                           = Color{R: 0.851, G: 0.863, B: 0.757, A: 1}
	Chamoisee                      = Color{R: 0.627, G: 0.471, B: 0.353, A: 1}
	Copper                         = Color{R: 0.722, G: 0.451, B: 0.200, A: 1}
	DarkSienna                     = Color{R: 0.235, G: 0.078, B: 0.078, A: 1}
	DeepGreen                      = Color{R: 0.020, G: 0.400, B: 0.031, A: 1}
	KombuGreen                     = Color{R: 0.208, G: 0.259, B: 0.188, A: 1}
	Korma                          = Color{R: 0.561, G: 0.294, B: 0.055, A: 1}
	Perano                         = Color{R: 0.663, G: 0.745, B: 0.949, A: 1}
	Pink                           = Color{R: 1.000, G: 0.753, B: 0.796, A: 1}
	Beaver                         = Color{R: 0.624, G: 0.506, B: 0.439, A: 1}
	CardinalPink                   = Color{R: 0.549, G: 0.020, B: 0.369, A: 1}
	Chestnut                       = Color{R: 0.584, G: 0.271, B: 0.208, A: 1}
	ColdPurple                     = Color{R: 0.671, G: 0.627, B: 0.851, A: 1}
	HairyHeath                     = Color{R: 0.420, G: 0.165, B: 0.078, A: 1}
	Rust                           = Color{R: 0.718, G: 0.255, B: 0.055, A: 1}
	GoBen                          = Color{R: 0.447, G: 0.427, B: 0.306, A: 1}
	NorthTexasGreen                = Color{R: 0.020, G: 0.565, B: 0.200, A: 1}
	PeriwinkleGray                 = Color{R: 0.765, G: 0.804, B: 0.902, A: 1}
	PlumpPurple                    = Color{R: 0.349, G: 0.275, B: 0.698, A: 1}
	Temptress                      = Color{R: 0.231, G: 0.000, B: 0.043, A: 1}
	SmaltBlue                      = Color{R: 0.318, G: 0.502, B: 0.561, A: 1}
	Sundown                        = Color{R: 1.000, G: 0.694, B: 0.702, A: 1}
	UFOGreen                       = Color{R: 0.235, G: 0.816, B: 0.439, A: 1}
	Saddle                         = Color{R: 0.298, G: 0.188, B: 0.141, A: 1}
	Sprout                         = Color{R: 0.757, G: 0.843, B: 0.690, A: 1}
	TuscanTan                      = Color{R: 0.651, G: 0.482, B: 0.357, A: 1}
	Buccaneer                      = Color{R: 0.384, G: 0.184, B: 0.188, A: 1}
	CannonPink                     = Color{R: 0.537, G: 0.263, B: 0.404, A: 1}
	GreenPea                       = Color{R: 0.114, G: 0.380, B: 0.259, A: 1}
	HotMagenta                     = Color{R: 1.000, G: 0.114, B: 0.808, A: 1}
	OperaMauve                     = Color{R: 0.718, G: 0.518, B: 0.655, A: 1}
	DesertStorm                    = Color{R: 0.973, G: 0.973, B: 0.969, A: 1}
	LaPalma                        = Color{R: 0.212, G: 0.529, B: 0.086, A: 1}
	Razzmatazz                     = Color{R: 0.890, G: 0.145, B: 0.420, A: 1}
	TiaMaria                       = Color{R: 0.757, G: 0.267, B: 0.055, A: 1}
	ScarletGum                     = Color{R: 0.263, G: 0.082, B: 0.376, A: 1}
	Sepia                          = Color{R: 0.439, G: 0.259, B: 0.078, A: 1}
	Skobeloff                      = Color{R: 0.000, G: 0.455, B: 0.455, A: 1}
	DaisyBush                      = Color{R: 0.310, G: 0.137, B: 0.596, A: 1}
	DesertSand                     = Color{R: 0.929, G: 0.788, B: 0.686, A: 1}
	Kabul                          = Color{R: 0.369, G: 0.282, B: 0.243, A: 1}
	LightBrown                     = Color{R: 0.710, G: 0.396, B: 0.114, A: 1}
	PearlLusta                     = Color{R: 0.988, G: 0.957, B: 0.863, A: 1}
	Snuff                          = Color{R: 0.886, G: 0.847, B: 0.929, A: 1}
	Ebony                          = Color{R: 0.333, G: 0.365, B: 0.314, A: 1}
	Feta                           = Color{R: 0.941, G: 0.988, B: 0.918, A: 1}
	RomanSilver                    = Color{R: 0.514, G: 0.537, B: 0.588, A: 1}
	CapePalliser                   = Color{R: 0.635, G: 0.400, B: 0.271, A: 1}
	FrenchPass                     = Color{R: 0.741, G: 0.929, B: 0.992, A: 1}
	Flavescent                     = Color{R: 0.969, G: 0.914, B: 0.557, A: 1}
	FunBlue                        = Color{R: 0.098, G: 0.349, B: 0.659, A: 1}
	LemonCurry                     = Color{R: 0.800, G: 0.627, B: 0.114, A: 1}
	Salomie                        = Color{R: 0.996, G: 0.859, B: 0.553, A: 1}
	SpringFrost                    = Color{R: 0.529, G: 1.000, B: 0.165, A: 1}
	VeryLightBlue                  = Color{R: 0.400, G: 0.400, B: 1.000, A: 1}
	BlueBell                       = Color{R: 0.635, G: 0.635, B: 0.816, A: 1}
	Crowshead                      = Color{R: 0.110, G: 0.071, B: 0.031, A: 1}
	Gravel                         = Color{R: 0.290, G: 0.267, B: 0.294, A: 1}
	HitPink                        = Color{R: 1.000, G: 0.671, B: 0.506, A: 1}
	LemonGinger                    = Color{R: 0.675, G: 0.620, B: 0.133, A: 1}
	SpringBud                      = Color{R: 0.655, G: 0.988, B: 0.000, A: 1}
	CoralRed                       = Color{R: 1.000, G: 0.251, B: 0.251, A: 1}
	Elephant                       = Color{R: 0.071, G: 0.204, B: 0.278, A: 1}
	Jagger                         = Color{R: 0.208, G: 0.055, B: 0.341, A: 1}
	OysterBay                      = Color{R: 0.855, G: 0.980, B: 1.000, A: 1}
	Persimmon                      = Color{R: 0.925, G: 0.345, B: 0.000, A: 1}
	LilacBush                      = Color{R: 0.596, G: 0.455, B: 0.827, A: 1}
	LimedAsh                       = Color{R: 0.455, G: 0.490, B: 0.388, A: 1}
	WintergreenDream               = Color{R: 0.337, G: 0.533, B: 0.490, A: 1}
	HintofGreen                    = Color{R: 0.902, G: 1.000, B: 0.914, A: 1}
	LightHotPink                   = Color{R: 1.000, G: 0.702, B: 0.871, A: 1}
	ProcessCyan                    = Color{R: 0.000, G: 0.718, B: 0.922, A: 1}
	Tosca                          = Color{R: 0.553, G: 0.247, B: 0.247, A: 1}
	TurquoiseGreen                 = Color{R: 0.627, G: 0.839, B: 0.706, A: 1}
	Pearl                          = Color{R: 0.918, G: 0.878, B: 0.784, A: 1}
	Sulu                           = Color{R: 0.757, G: 0.941, B: 0.486, A: 1}
	DarkBrown                      = Color{R: 0.396, G: 0.263, B: 0.129, A: 1}
	DarkViolet                     = Color{R: 0.580, G: 0.000, B: 0.827, A: 1}
	GovernorBay                    = Color{R: 0.184, G: 0.235, B: 0.702, A: 1}
	LightYellow                    = Color{R: 1.000, G: 1.000, B: 0.878, A: 1}
	Nyanza                         = Color{R: 0.914, G: 1.000, B: 0.859, A: 1}
	GOGreen                        = Color{R: 0.000, G: 0.671, B: 0.400, A: 1}
	SteelTeal                      = Color{R: 0.373, G: 0.541, B: 0.545, A: 1}
	VidaLoca                       = Color{R: 0.329, G: 0.565, B: 0.098, A: 1}
	AirForceBlue                   = Color{R: 0.000, G: 0.188, B: 0.561, A: 1}
	MediumRedViolet                = Color{R: 0.733, G: 0.200, B: 0.522, A: 1}
	SeaPink                        = Color{R: 0.929, G: 0.596, B: 0.620, A: 1}
	SnowDrift                      = Color{R: 0.969, G: 0.980, B: 0.969, A: 1}
	Spindle                        = Color{R: 0.714, G: 0.820, B: 0.918, A: 1}
	ShimmeringBlush                = Color{R: 0.851, G: 0.525, B: 0.584, A: 1}
	StormGray                      = Color{R: 0.443, G: 0.455, B: 0.525, A: 1}
	WildRice                       = Color{R: 0.925, G: 0.878, B: 0.565, A: 1}
	BudGreen                       = Color{R: 0.482, G: 0.714, B: 0.380, A: 1}
	DeYork                         = Color{R: 0.478, G: 0.769, B: 0.533, A: 1}
	GreenVogue                     = Color{R: 0.012, G: 0.169, B: 0.322, A: 1}
	Manz                           = Color{R: 0.933, G: 0.937, B: 0.471, A: 1}
	Neptune                        = Color{R: 0.486, G: 0.718, B: 0.733, A: 1}
	ShadyLady                      = Color{R: 0.667, G: 0.647, B: 0.663, A: 1}
	Botticelli                     = Color{R: 0.780, G: 0.867, B: 0.898, A: 1}
	BrilliantRose                  = Color{R: 1.000, G: 0.333, B: 0.639, A: 1}
	ColumbiaBlue                   = Color{R: 0.769, G: 0.847, B: 0.886, A: 1}
	Jade                           = Color{R: 0.000, G: 0.659, B: 0.420, A: 1}
	MediumAquamarine               = Color{R: 0.400, G: 0.867, B: 0.667, A: 1}
	Mindaro                        = Color{R: 0.890, G: 0.976, B: 0.533, A: 1}
	Ottoman                        = Color{R: 0.914, G: 0.973, B: 0.929, A: 1}
	Tacao                          = Color{R: 0.929, G: 0.702, B: 0.506, A: 1}
	BlackHaze                      = Color{R: 0.965, G: 0.969, B: 0.969, A: 1}
	BrightGray                     = Color{R: 0.235, G: 0.255, B: 0.318, A: 1}
	Christi                        = Color{R: 0.404, G: 0.655, B: 0.071, A: 1}
	DarkTurquoise                  = Color{R: 0.000, G: 0.808, B: 0.820, A: 1}
	ElectricViolet                 = Color{R: 0.545, G: 0.000, B: 1.000, A: 1}
	Tenne                          = Color{R: 0.804, G: 0.341, B: 0.000, A: 1}
	EarthYellow                    = Color{R: 0.882, G: 0.663, B: 0.373, A: 1}
	TropicalRainForest             = Color{R: 0.000, G: 0.459, B: 0.369, A: 1}
	Zambezi                        = Color{R: 0.408, G: 0.333, B: 0.345, A: 1}
	Desert                         = Color{R: 0.682, G: 0.376, B: 0.125, A: 1}
	EastBay                        = Color{R: 0.255, G: 0.298, B: 0.490, A: 1}
	Eternity                       = Color{R: 0.129, G: 0.102, B: 0.055, A: 1}
	QuinacridoneMagenta            = Color{R: 0.557, G: 0.227, B: 0.349, A: 1}
	Ruber                          = Color{R: 0.808, G: 0.275, B: 0.463, A: 1}
	FuchsiaRose                    = Color{R: 0.780, G: 0.263, B: 0.459, A: 1}
	Makara                         = Color{R: 0.537, G: 0.490, B: 0.427, A: 1}
	BakerMillerPink                = Color{R: 1.000, G: 0.569, B: 0.686, A: 1}
	HeliotropeMagenta              = Color{R: 0.667, G: 0.000, B: 0.733, A: 1}
	MintCream                      = Color{R: 0.961, G: 1.000, B: 0.980, A: 1}
	SapphireBlue                   = Color{R: 0.000, G: 0.404, B: 0.647, A: 1}
	SpanishSkyBlue                 = Color{R: 0.000, G: 0.667, B: 0.894, A: 1}
	BrownDerby                     = Color{R: 0.286, G: 0.149, B: 0.082, A: 1}
	CopperRose                     = Color{R: 0.600, G: 0.400, B: 0.400, A: 1}
	DarkMagenta                    = Color{R: 0.545, G: 0.000, B: 0.545, A: 1}
	LightGray                      = Color{R: 0.827, G: 0.827, B: 0.827, A: 1}
	PalePink                       = Color{R: 0.980, G: 0.855, B: 0.867, A: 1}
	Sandal                         = Color{R: 0.667, G: 0.553, B: 0.435, A: 1}
	Stormcloud                     = Color{R: 0.310, G: 0.400, B: 0.416, A: 1}
	BridalHeath                    = Color{R: 1.000, G: 0.980, B: 0.957, A: 1}
	GreenKelp                      = Color{R: 0.145, G: 0.192, B: 0.110, A: 1}
	Indochine                      = Color{R: 0.761, G: 0.420, B: 0.012, A: 1}
	TealGreen                      = Color{R: 0.000, G: 0.510, B: 0.498, A: 1}
	MediumJungleGreen              = Color{R: 0.110, G: 0.208, B: 0.176, A: 1}
	PeachPuff                      = Color{R: 1.000, G: 0.855, B: 0.725, A: 1}
	ChromeWhite                    = Color{R: 0.910, G: 0.945, B: 0.831, A: 1}
	CoyoteBrown                    = Color{R: 0.506, G: 0.380, B: 0.243, A: 1}
	DeepCarmine                    = Color{R: 0.663, G: 0.125, B: 0.243, A: 1}
	Gumbo                          = Color{R: 0.486, G: 0.631, B: 0.651, A: 1}
	HanPurple                      = Color{R: 0.322, G: 0.094, B: 0.980, A: 1}
	DeepPink                       = Color{R: 1.000, G: 0.078, B: 0.576, A: 1}
	EtonBlue                       = Color{R: 0.588, G: 0.784, B: 0.635, A: 1}
	MediumSkyBlue                  = Color{R: 0.502, G: 0.855, B: 0.922, A: 1}
	Shakespeare                    = Color{R: 0.306, G: 0.671, B: 0.820, A: 1}
	Heather                        = Color{R: 0.718, G: 0.765, B: 0.816, A: 1}
	PearlBush                      = Color{R: 0.910, G: 0.878, B: 0.835, A: 1}
	Saratoga                       = Color{R: 0.333, G: 0.357, B: 0.063, A: 1}
	BlueCharcoal                   = Color{R: 0.004, G: 0.051, B: 0.102, A: 1}
	CosmicLatte                    = Color{R: 1.000, G: 0.973, B: 0.906, A: 1}
	Fedora                         = Color{R: 0.475, G: 0.416, B: 0.471, A: 1}
	FlushMahogany                  = Color{R: 0.792, G: 0.204, B: 0.208, A: 1}
	GlossyGrape                    = Color{R: 0.671, G: 0.573, B: 0.702, A: 1}
	Wedgewood                      = Color{R: 0.306, G: 0.498, B: 0.620, A: 1}
	Barossa                        = Color{R: 0.267, G: 0.004, B: 0.176, A: 1}
	CoolBlack                      = Color{R: 0.000, G: 0.180, B: 0.388, A: 1}
	Nepal                          = Color{R: 0.557, G: 0.671, B: 0.757, A: 1}
	PaleTaupe                      = Color{R: 0.737, G: 0.596, B: 0.494, A: 1}
	Watercourse                    = Color{R: 0.020, G: 0.435, B: 0.341, A: 1}
	IronsideGray                   = Color{R: 0.404, G: 0.400, B: 0.384, A: 1}
	VeryPaleYellow                 = Color{R: 1.000, G: 1.000, B: 0.749, A: 1}
	VividBurgundy                  = Color{R: 0.624, G: 0.114, B: 0.208, A: 1}
	BlackLeatherJacket             = Color{R: 0.145, G: 0.208, B: 0.161, A: 1}
	CherryPie                      = Color{R: 0.165, G: 0.012, B: 0.349, A: 1}
	GreenHaze                      = Color{R: 0.004, G: 0.639, B: 0.408, A: 1}
	Zircon                         = Color{R: 0.957, G: 0.973, B: 1.000, A: 1}
	Zorba                          = Color{R: 0.647, G: 0.608, B: 0.569, A: 1}
	Aluminium                      = Color{R: 0.663, G: 0.675, B: 0.714, A: 1}
	BlueJeans                      = Color{R: 0.365, G: 0.678, B: 0.925, A: 1}
	BurntSienna                    = Color{R: 0.914, G: 0.455, B: 0.318, A: 1}
	DarkRaspberry                  = Color{R: 0.529, G: 0.149, B: 0.341, A: 1}
	SapGreen                       = Color{R: 0.314, G: 0.490, B: 0.165, A: 1}
	Regalia                        = Color{R: 0.322, G: 0.176, B: 0.502, A: 1}
	Ronchi                         = Color{R: 0.925, G: 0.773, B: 0.306, A: 1}
	Blackberry                     = Color{R: 0.302, G: 0.004, B: 0.208, A: 1}
	Marshland                      = Color{R: 0.043, G: 0.059, B: 0.031, A: 1}
	Mauvelous                      = Color{R: 0.937, G: 0.596, B: 0.667, A: 1}
	OliveDrabSeven                 = Color{R: 0.235, G: 0.204, B: 0.122, A: 1}
	Pohutukawa                     = Color{R: 0.561, G: 0.008, B: 0.110, A: 1}
	Soap                           = Color{R: 0.808, G: 0.784, B: 0.937, A: 1}
	SugarPlum                      = Color{R: 0.569, G: 0.306, B: 0.459, A: 1}
	TigersEye                      = Color{R: 0.878, G: 0.553, B: 0.235, A: 1}
	GargoyleGas                    = Color{R: 1.000, G: 0.875, B: 0.275, A: 1}
	PearlyPurple                   = Color{R: 0.718, G: 0.408, B: 0.635, A: 1}
	PersianPlum                    = Color{R: 0.439, G: 0.110, B: 0.110, A: 1}
	RichLavender                   = Color{R: 0.655, G: 0.420, B: 0.812, A: 1}
	Sambuca                        = Color{R: 0.227, G: 0.125, B: 0.063, A: 1}
	Navy                           = Color{R: 0.000, G: 0.000, B: 0.502, A: 1}
	Brass                          = Color{R: 0.710, G: 0.651, B: 0.259, A: 1}
	MughalGreen                    = Color{R: 0.188, G: 0.376, B: 0.188, A: 1}
	Wisteria                       = Color{R: 0.788, G: 0.627, B: 0.863, A: 1}
	Woodland                       = Color{R: 0.302, G: 0.325, B: 0.157, A: 1}
	Cumulus                        = Color{R: 0.992, G: 1.000, B: 0.835, A: 1}
	FuchsiaBlue                    = Color{R: 0.478, G: 0.345, B: 0.757, A: 1}
	LavenderPink                   = Color{R: 0.984, G: 0.682, B: 0.824, A: 1}
	LightTaupe                     = Color{R: 0.702, G: 0.545, B: 0.427, A: 1}
	Plum                           = Color{R: 0.557, G: 0.271, B: 0.522, A: 1}
	Fog                            = Color{R: 0.843, G: 0.816, B: 1.000, A: 1}
	FuchsiaPurple                  = Color{R: 0.800, G: 0.224, B: 0.482, A: 1}
	SuvaGray                       = Color{R: 0.533, G: 0.514, B: 0.529, A: 1}
	BilobaFlower                   = Color{R: 0.698, G: 0.631, B: 0.918, A: 1}
	CoolGrey                       = Color{R: 0.549, G: 0.573, B: 0.675, A: 1}
	NewCar                         = Color{R: 0.129, G: 0.310, B: 0.776, A: 1}
	PigPink                        = Color{R: 0.992, G: 0.843, B: 0.894, A: 1}
	SoftPeach                      = Color{R: 0.961, G: 0.929, B: 0.937, A: 1}
	OldRose                        = Color{R: 0.753, G: 0.502, B: 0.506, A: 1}
	BuddhaGold                     = Color{R: 0.757, G: 0.627, B: 0.016, A: 1}
	Casal                          = Color{R: 0.184, G: 0.380, B: 0.408, A: 1}
	ChineseRed                     = Color{R: 0.667, G: 0.220, B: 0.118, A: 1}
	GreenLeaf                      = Color{R: 0.263, G: 0.416, B: 0.051, A: 1}
	Hoki                           = Color{R: 0.396, G: 0.525, B: 0.624, A: 1}
	Bitter                         = Color{R: 0.525, G: 0.537, B: 0.455, A: 1}
	Bunting                        = Color{R: 0.082, G: 0.122, B: 0.298, A: 1}
	IndianRed                      = Color{R: 0.804, G: 0.361, B: 0.361, A: 1}
	MoodyBlue                      = Color{R: 0.498, G: 0.463, B: 0.827, A: 1}
	MountainMist                   = Color{R: 0.584, G: 0.576, B: 0.588, A: 1}
	Rhythm                         = Color{R: 0.467, G: 0.463, B: 0.588, A: 1}
	SpringSun                      = Color{R: 0.965, G: 1.000, B: 0.863, A: 1}
	Atoll                          = Color{R: 0.039, G: 0.435, B: 0.459, A: 1}
	DarkEbony                      = Color{R: 0.235, G: 0.125, B: 0.020, A: 1}
	HalayàÚbe                      = Color{R: 0.400, G: 0.220, B: 0.329, A: 1}
	Masala                         = Color{R: 0.251, G: 0.231, B: 0.220, A: 1}
	Pueblo                         = Color{R: 0.490, G: 0.173, B: 0.078, A: 1}
	BallBlue                       = Color{R: 0.129, G: 0.671, B: 0.804, A: 1}
	Indigo                         = Color{R: 0.294, G: 0.000, B: 0.510, A: 1}
	Serenade                       = Color{R: 1.000, G: 0.957, B: 0.910, A: 1}
	Liver                          = Color{R: 0.404, G: 0.298, B: 0.278, A: 1}
	Caramel                        = Color{R: 1.000, G: 0.867, B: 0.686, A: 1}
	PinkFlamingo                   = Color{R: 0.988, G: 0.455, B: 0.992, A: 1}
	Magnolia                       = Color{R: 0.973, G: 0.957, B: 1.000, A: 1}
	OrchidWhite                    = Color{R: 1.000, G: 0.992, B: 0.953, A: 1}
	RedDevil                       = Color{R: 0.525, G: 0.004, B: 0.067, A: 1}
	TartOrange                     = Color{R: 0.984, G: 0.302, B: 0.275, A: 1}
	AquaSqueeze                    = Color{R: 0.910, G: 0.961, B: 0.949, A: 1}
	Bamboo                         = Color{R: 0.855, G: 0.388, B: 0.016, A: 1}
	DeepJungleGreen                = Color{R: 0.000, G: 0.294, B: 0.286, A: 1}
	Downriver                      = Color{R: 0.035, G: 0.133, B: 0.337, A: 1}
	LaSalleGreen                   = Color{R: 0.031, G: 0.471, B: 0.188, A: 1}
	Hillary                        = Color{R: 0.675, G: 0.647, B: 0.525, A: 1}
	HorsesNeck                     = Color{R: 0.376, G: 0.286, B: 0.075, A: 1}
	Peppermint                     = Color{R: 0.890, G: 0.961, B: 0.882, A: 1}
	HokeyPokey                     = Color{R: 0.784, G: 0.647, B: 0.157, A: 1}
	Moccaccino                     = Color{R: 0.431, G: 0.114, B: 0.078, A: 1}
	RedSalsa                       = Color{R: 0.992, G: 0.227, B: 0.290, A: 1}
	ClayCreek                      = Color{R: 0.541, G: 0.514, B: 0.376, A: 1}
	DingyDungeon                   = Color{R: 0.773, G: 0.192, B: 0.318, A: 1}
	PixiePowder                    = Color{R: 0.224, G: 0.071, B: 0.522, A: 1}
	RiverBed                       = Color{R: 0.263, G: 0.298, B: 0.349, A: 1}
	JapaneseCarmine                = Color{R: 0.616, G: 0.161, B: 0.200, A: 1}
	PigmentRed                     = Color{R: 0.929, G: 0.110, B: 0.141, A: 1}
	SheenGreen                     = Color{R: 0.561, G: 0.831, B: 0.000, A: 1}
	Siren                          = Color{R: 0.478, G: 0.004, B: 0.227, A: 1}
	Texas                          = Color{R: 0.973, G: 0.976, B: 0.612, A: 1}
	Bordeaux                       = Color{R: 0.361, G: 0.004, B: 0.125, A: 1}
	Buttermilk                     = Color{R: 1.000, G: 0.945, B: 0.710, A: 1}
	MangoTango                     = Color{R: 1.000, G: 0.510, B: 0.263, A: 1}
	VividMulberry                  = Color{R: 0.722, G: 0.047, B: 0.890, A: 1}
	CadmiumYellow                  = Color{R: 1.000, G: 0.965, B: 0.000, A: 1}
	HarvestGold                    = Color{R: 0.855, G: 0.569, B: 0.000, A: 1}
	RaspberryPink                  = Color{R: 0.886, G: 0.314, B: 0.596, A: 1}
	StarDust                       = Color{R: 0.624, G: 0.624, B: 0.612, A: 1}
	WispPink                       = Color{R: 0.996, G: 0.957, B: 0.973, A: 1}
	Ash                            = Color{R: 0.776, G: 0.765, B: 0.710, A: 1}
	CadetGrey                      = Color{R: 0.569, G: 0.639, B: 0.690, A: 1}
	Lavender                       = Color{R: 0.710, G: 0.494, B: 0.863, A: 1}
	Sunset                         = Color{R: 0.980, G: 0.839, B: 0.647, A: 1}
	VividYellow                    = Color{R: 1.000, G: 0.890, B: 0.008, A: 1}
	BrilliantLavender              = Color{R: 0.957, G: 0.733, B: 1.000, A: 1}
	Fiord                          = Color{R: 0.251, G: 0.318, B: 0.412, A: 1}
	MacaroniAndCheese              = Color{R: 1.000, G: 0.741, B: 0.533, A: 1}
	PersianRed                     = Color{R: 0.800, G: 0.200, B: 0.200, A: 1}
	RustyRed                       = Color{R: 0.855, G: 0.173, B: 0.263, A: 1}
	BrightRed                      = Color{R: 0.694, G: 0.000, B: 0.000, A: 1}
	FrenchBistre                   = Color{R: 0.522, G: 0.427, B: 0.302, A: 1}
	Marzipan                       = Color{R: 0.973, G: 0.859, B: 0.616, A: 1}
	Spice                          = Color{R: 0.416, G: 0.267, B: 0.180, A: 1}
	MinionYellow                   = Color{R: 0.961, G: 0.878, B: 0.314, A: 1}
	RumSwizzle                     = Color{R: 0.976, G: 0.973, B: 0.894, A: 1}
	ElectricYellow                 = Color{R: 1.000, G: 1.000, B: 0.200, A: 1}
	Zest                           = Color{R: 0.898, G: 0.518, B: 0.106, A: 1}
	Swamp                          = Color{R: 0.000, G: 0.106, B: 0.110, A: 1}
	CottonCandy                    = Color{R: 1.000, G: 0.737, B: 0.851, A: 1}
	DarkMediumGray                 = Color{R: 0.663, G: 0.663, B: 0.663, A: 1}
	Pablo                          = Color{R: 0.467, G: 0.435, B: 0.380, A: 1}
	SpringRain                     = Color{R: 0.675, G: 0.796, B: 0.694, A: 1}
	StormDust                      = Color{R: 0.392, G: 0.392, B: 0.388, A: 1}
	JapaneseIndigo                 = Color{R: 0.149, G: 0.263, B: 0.282, A: 1}
	MilanoRed                      = Color{R: 0.722, G: 0.067, B: 0.016, A: 1}
	SpicyMustard                   = Color{R: 0.455, G: 0.392, B: 0.051, A: 1}
	TurtleGreen                    = Color{R: 0.165, G: 0.220, B: 0.043, A: 1}
	DarkPastelPurple               = Color{R: 0.588, G: 0.435, B: 0.839, A: 1}
	Killarney                      = Color{R: 0.227, G: 0.416, B: 0.278, A: 1}
	TwilightLavender               = Color{R: 0.541, G: 0.286, B: 0.420, A: 1}
	VeryLightAzure                 = Color{R: 0.455, G: 0.733, B: 0.984, A: 1}
	CocoaBrown                     = Color{R: 0.824, G: 0.412, B: 0.118, A: 1}
	MediumTurquoise                = Color{R: 0.282, G: 0.820, B: 0.800, A: 1}
	NaturalGray                    = Color{R: 0.545, G: 0.525, B: 0.502, A: 1}
	NavajoWhite                    = Color{R: 1.000, G: 0.871, B: 0.678, A: 1}
	QueenPink                      = Color{R: 0.910, G: 0.800, B: 0.843, A: 1}
	Wafer                          = Color{R: 0.871, G: 0.796, B: 0.776, A: 1}
	Bronze                         = Color{R: 0.804, G: 0.498, B: 0.196, A: 1}
	CavernPink                     = Color{R: 0.890, G: 0.745, B: 0.745, A: 1}
	ForestGreen                    = Color{R: 0.133, G: 0.545, B: 0.133, A: 1}
	MardiGras                      = Color{R: 0.533, G: 0.000, B: 0.522, A: 1}
	MexicanRed                     = Color{R: 0.655, G: 0.145, B: 0.145, A: 1}
	Holly                          = Color{R: 0.004, G: 0.114, B: 0.075, A: 1}
	LightSeaGreen                  = Color{R: 0.125, G: 0.698, B: 0.667, A: 1}
	Opal                           = Color{R: 0.663, G: 0.776, B: 0.761, A: 1}
	Plantation                     = Color{R: 0.153, G: 0.314, B: 0.294, A: 1}
	ThistleGreen                   = Color{R: 0.800, G: 0.792, B: 0.659, A: 1}
	BigDipOruby                    = Color{R: 0.612, G: 0.145, B: 0.259, A: 1}
	DeepBlush                      = Color{R: 0.894, G: 0.463, B: 0.596, A: 1}
	Gray                           = Color{R: 0.502, G: 0.502, B: 0.502, A: 1}
	LightGreen                     = Color{R: 0.565, G: 0.933, B: 0.565, A: 1}
	PermanentGeraniumLake          = Color{R: 0.882, G: 0.173, B: 0.173, A: 1}
	NewYorkPink                    = Color{R: 0.843, G: 0.514, B: 0.498, A: 1}
	Pacifika                       = Color{R: 0.467, G: 0.506, B: 0.125, A: 1}
	Scarlett                       = Color{R: 0.584, G: 0.000, B: 0.082, A: 1}
	Amour                          = Color{R: 0.976, G: 0.918, B: 0.953, A: 1}
	Concord                        = Color{R: 0.486, G: 0.482, B: 0.478, A: 1}
	FrenchLilac                    = Color{R: 0.525, G: 0.376, B: 0.557, A: 1}
	LightPastelPurple              = Color{R: 0.694, G: 0.612, B: 0.851, A: 1}
	LiverChestnut                  = Color{R: 0.596, G: 0.455, B: 0.337, A: 1}
	TropicalViolet                 = Color{R: 0.804, G: 0.643, B: 0.871, A: 1}
	PineCone                       = Color{R: 0.427, G: 0.369, B: 0.329, A: 1}
	Tara                           = Color{R: 0.882, G: 0.965, B: 0.910, A: 1}
	Ube                            = Color{R: 0.533, G: 0.471, B: 0.765, A: 1}
	BrownSugar                     = Color{R: 0.686, G: 0.431, B: 0.302, A: 1}
	Dogs                           = Color{R: 0.722, G: 0.427, B: 0.161, A: 1}
	GreenLizard                    = Color{R: 0.655, G: 0.957, B: 0.196, A: 1}
	HarvardCrimson                 = Color{R: 0.788, G: 0.000, B: 0.086, A: 1}
	Mahogany                       = Color{R: 0.753, G: 0.251, B: 0.000, A: 1}
	WePeep                         = Color{R: 0.969, G: 0.859, B: 0.902, A: 1}
	Bahia                          = Color{R: 0.647, G: 0.796, B: 0.047, A: 1}
	BlazeOrange                    = Color{R: 1.000, G: 0.404, B: 0.000, A: 1}
	CosmicCobalt                   = Color{R: 0.180, G: 0.176, B: 0.533, A: 1}
	Fuchsia                        = Color{R: 1.000, G: 0.000, B: 1.000, A: 1}
	NileBlue                       = Color{R: 0.098, G: 0.216, B: 0.318, A: 1}
	WaikawaGray                    = Color{R: 0.353, G: 0.431, B: 0.612, A: 1}
	Cardinal                       = Color{R: 0.769, G: 0.118, B: 0.227, A: 1}
	LividBrown                     = Color{R: 0.302, G: 0.157, B: 0.180, A: 1}
	Onion                          = Color{R: 0.184, G: 0.153, B: 0.055, A: 1}
	ScarpaFlow                     = Color{R: 0.345, G: 0.333, B: 0.384, A: 1}
	Spray                          = Color{R: 0.475, G: 0.871, B: 0.925, A: 1}
	Sandstone                      = Color{R: 0.475, G: 0.427, B: 0.384, A: 1}
	SteelPink                      = Color{R: 0.800, G: 0.200, B: 0.800, A: 1}
	SwissCoffee                    = Color{R: 0.867, G: 0.839, B: 0.835, A: 1}
	Corn                           = Color{R: 0.906, G: 0.749, B: 0.020, A: 1}
	GambogeOrange                  = Color{R: 0.600, G: 0.400, B: 0.000, A: 1}
	OldLavender                    = Color{R: 0.475, G: 0.408, B: 0.471, A: 1}
	PickledBean                    = Color{R: 0.431, G: 0.282, B: 0.149, A: 1}
	Rajah                          = Color{R: 0.984, G: 0.671, B: 0.376, A: 1}
	Woodsmoke                      = Color{R: 0.047, G: 0.051, B: 0.059, A: 1}
	OceanBoatBlue                  = Color{R: 0.000, G: 0.467, B: 0.745, A: 1}
	Selago                         = Color{R: 0.941, G: 0.933, B: 0.992, A: 1}
	Canary                         = Color{R: 0.953, G: 0.984, B: 0.384, A: 1}
	Carmine                        = Color{R: 0.588, G: 0.000, B: 0.094, A: 1}
	Green                          = Color{R: 0.000, G: 1.000, B: 0.000, A: 1}
	GreenMist                      = Color{R: 0.796, G: 0.827, B: 0.690, A: 1}
	Magenta                        = Color{R: 0.792, G: 0.122, B: 0.482, A: 1}
	Stratos                        = Color{R: 0.000, G: 0.027, B: 0.255, A: 1}
	BurntMaroon                    = Color{R: 0.259, G: 0.012, B: 0.012, A: 1}
	CyanCobaltBlue                 = Color{R: 0.157, G: 0.345, B: 0.612, A: 1}
	Mocha                          = Color{R: 0.471, G: 0.176, B: 0.098, A: 1}
	SanguineBrown                  = Color{R: 0.553, G: 0.239, B: 0.220, A: 1}
	SmokyBlack                     = Color{R: 0.063, G: 0.047, B: 0.031, A: 1}
	BlueStone                      = Color{R: 0.004, G: 0.380, B: 0.384, A: 1}
	Catawba                        = Color{R: 0.439, G: 0.212, B: 0.259, A: 1}
	Chamois                        = Color{R: 0.929, G: 0.863, B: 0.694, A: 1}
	Zeus                           = Color{R: 0.161, G: 0.137, B: 0.098, A: 1}
	BlueMarguerite                 = Color{R: 0.463, G: 0.400, B: 0.776, A: 1}
	ChileanHeath                   = Color{R: 1.000, G: 0.992, B: 0.902, A: 1}
	Dell                           = Color{R: 0.224, G: 0.392, B: 0.075, A: 1}
	HalfBaked                      = Color{R: 0.522, G: 0.769, B: 0.800, A: 1}
	PinkSwan                       = Color{R: 0.745, G: 0.710, B: 0.718, A: 1}
	ChelseaCucumber                = Color{R: 0.514, G: 0.667, B: 0.365, A: 1}
	DarkYellow                     = Color{R: 0.608, G: 0.529, B: 0.047, A: 1}
	Dolly                          = Color{R: 0.976, G: 1.000, B: 0.545, A: 1}
	Ferra                          = Color{R: 0.439, G: 0.310, B: 0.314, A: 1}
	Shampoo                        = Color{R: 1.000, G: 0.812, B: 0.945, A: 1}
	Valentino                      = Color{R: 0.208, G: 0.055, B: 0.259, A: 1}
	Wheatfield                     = Color{R: 0.953, G: 0.929, B: 0.812, A: 1}
	BarleyWhite                    = Color{R: 1.000, G: 0.957, B: 0.808, A: 1}
	GrannySmith                    = Color{R: 0.518, G: 0.627, B: 0.627, A: 1}
	Highland                       = Color{R: 0.435, G: 0.557, B: 0.388, A: 1}
	Mantis                         = Color{R: 0.455, G: 0.765, B: 0.396, A: 1}
	Mikado                         = Color{R: 0.176, G: 0.145, B: 0.063, A: 1}
	Sandrift                       = Color{R: 0.671, G: 0.569, B: 0.478, A: 1}
	Cello                          = Color{R: 0.118, G: 0.220, B: 0.357, A: 1}
	GinFizz                        = Color{R: 1.000, G: 0.976, B: 0.886, A: 1}
	HalfSpanishWhite               = Color{R: 0.996, G: 0.957, B: 0.859, A: 1}
	Monarch                        = Color{R: 0.545, G: 0.027, B: 0.137, A: 1}
	OldBrick                       = Color{R: 0.565, G: 0.118, B: 0.118, A: 1}
	SealBrown                      = Color{R: 0.349, G: 0.149, B: 0.043, A: 1}
	ButteryWhite                   = Color{R: 1.000, G: 0.988, B: 0.918, A: 1}
	ElectricIndigo                 = Color{R: 0.435, G: 0.000, B: 1.000, A: 1}
	Fantasy                        = Color{R: 0.980, G: 0.953, B: 0.941, A: 1}
	Pippin                         = Color{R: 1.000, G: 0.882, B: 0.875, A: 1}
	Scampi                         = Color{R: 0.404, G: 0.373, B: 0.651, A: 1}
	Cosmos                         = Color{R: 1.000, G: 0.847, B: 0.851, A: 1}
	DeepMagenta                    = Color{R: 0.800, G: 0.000, B: 0.800, A: 1}
	Desire                         = Color{R: 0.918, G: 0.235, B: 0.325, A: 1}
	RegentGray                     = Color{R: 0.525, G: 0.580, B: 0.624, A: 1}
	ShadowGreen                    = Color{R: 0.604, G: 0.761, B: 0.722, A: 1}
	AntiqueFuchsia                 = Color{R: 0.569, G: 0.361, B: 0.514, A: 1}
	Bisque                         = Color{R: 1.000, G: 0.894, B: 0.769, A: 1}
	BurntUmber                     = Color{R: 0.541, G: 0.200, B: 0.141, A: 1}
	DarkLavender                   = Color{R: 0.451, G: 0.310, B: 0.588, A: 1}
	Pavlova                        = Color{R: 0.843, G: 0.769, B: 0.596, A: 1}
	RYBRed                         = Color{R: 0.996, G: 0.153, B: 0.071, A: 1}
	Downy                          = Color{R: 0.435, G: 0.816, B: 0.773, A: 1}
	DutchWhite                     = Color{R: 0.937, G: 0.875, B: 0.733, A: 1}
	GreenSmoke                     = Color{R: 0.643, G: 0.686, B: 0.431, A: 1}
	Java                           = Color{R: 0.122, G: 0.761, B: 0.761, A: 1}
	QuillGray                      = Color{R: 0.839, G: 0.839, B: 0.820, A: 1}
	CuttySark                      = Color{R: 0.314, G: 0.463, B: 0.447, A: 1}
	MidnightMoss                   = Color{R: 0.016, G: 0.063, B: 0.016, A: 1}
	White                          = Color{R: 1.000, G: 1.000, B: 1.000, A: 1}
	Stack                          = Color{R: 0.541, G: 0.561, B: 0.541, A: 1}
	AirSuperiorityBlue             = Color{R: 0.447, G: 0.627, B: 0.757, A: 1}
	Asparagus                      = Color{R: 0.529, G: 0.663, B: 0.420, A: 1}
	MulledWine                     = Color{R: 0.306, G: 0.271, B: 0.384, A: 1}
	NapierGreen                    = Color{R: 0.165, G: 0.502, B: 0.000, A: 1}
	RichBrilliantLavender          = Color{R: 0.945, G: 0.655, B: 0.996, A: 1}
	Sandwisp                       = Color{R: 0.961, G: 0.906, B: 0.635, A: 1}
	Chartreuse                     = Color{R: 0.875, G: 1.000, B: 0.000, A: 1}
	Genoa                          = Color{R: 0.082, G: 0.451, B: 0.420, A: 1}
	Manatee                        = Color{R: 0.592, G: 0.604, B: 0.667, A: 1}
	Pumpkin                        = Color{R: 1.000, G: 0.459, B: 0.094, A: 1}
	Purpureus                      = Color{R: 0.604, G: 0.306, B: 0.682, A: 1}
	Teal                           = Color{R: 0.000, G: 0.502, B: 0.502, A: 1}
	BlackBean                      = Color{R: 0.239, G: 0.047, B: 0.008, A: 1}
	ChileanFire                    = Color{R: 0.969, G: 0.467, B: 0.012, A: 1}
	DarkPastelGreen                = Color{R: 0.012, G: 0.753, B: 0.235, A: 1}
	Icterine                       = Color{R: 0.988, G: 0.969, B: 0.369, A: 1}
	Pizazz                         = Color{R: 1.000, G: 0.565, B: 0.000, A: 1}
	CobaltBlue                     = Color{R: 0.000, G: 0.278, B: 0.671, A: 1}
	Comet                          = Color{R: 0.361, G: 0.365, B: 0.459, A: 1}
	Ecstasy                        = Color{R: 0.980, G: 0.471, B: 0.078, A: 1}
	LightCrimson                   = Color{R: 0.961, G: 0.412, B: 0.569, A: 1}
	BurntOrange                    = Color{R: 0.800, G: 0.333, B: 0.000, A: 1}
	GreenWaterloo                  = Color{R: 0.063, G: 0.078, B: 0.020, A: 1}
	Malta                          = Color{R: 0.741, G: 0.698, B: 0.631, A: 1}
	RifleGreen                     = Color{R: 0.267, G: 0.298, B: 0.220, A: 1}
	CarnabyTan                     = Color{R: 0.361, G: 0.180, B: 0.004, A: 1}
	RustyNail                      = Color{R: 0.525, G: 0.337, B: 0.039, A: 1}
	LightOrchid                    = Color{R: 0.902, G: 0.659, B: 0.843, A: 1}
	TangerineYellow                = Color{R: 1.000, G: 0.800, B: 0.000, A: 1}
	TeaRose                        = Color{R: 0.957, G: 0.761, B: 0.761, A: 1}
	Bubbles                        = Color{R: 0.906, G: 0.996, B: 1.000, A: 1}
	Fulvous                        = Color{R: 0.894, G: 0.518, B: 0.000, A: 1}
	IrishCoffee                    = Color{R: 0.373, G: 0.239, B: 0.149, A: 1}
	Irresistible                   = Color{R: 0.702, G: 0.267, B: 0.424, A: 1}
	IslandSpice                    = Color{R: 1.000, G: 0.988, B: 0.933, A: 1}
	AmethystSmoke                  = Color{R: 0.639, G: 0.592, B: 0.706, A: 1}
	Fire                           = Color{R: 0.667, G: 0.259, B: 0.012, A: 1}
	LavenderMist                   = Color{R: 0.902, G: 0.902, B: 0.980, A: 1}
	MetallicGold                   = Color{R: 0.831, G: 0.686, B: 0.216, A: 1}
	BrandyRose                     = Color{R: 0.733, G: 0.537, B: 0.514, A: 1}
	SandDune                       = Color{R: 0.588, G: 0.443, B: 0.090, A: 1}
	Charade                        = Color{R: 0.161, G: 0.161, B: 0.216, A: 1}
	HintofYellow                   = Color{R: 0.980, G: 0.992, B: 0.894, A: 1}
	Laser                          = Color{R: 0.784, G: 0.710, B: 0.408, A: 1}
	Porcelain                      = Color{R: 0.937, G: 0.949, B: 0.953, A: 1}
	Peach                          = Color{R: 1.000, G: 0.796, B: 0.643, A: 1}
	AztecGold                      = Color{R: 0.765, G: 0.600, B: 0.325, A: 1}
	Bilbao                         = Color{R: 0.196, G: 0.486, B: 0.078, A: 1}
	CambridgeBlue                  = Color{R: 0.639, G: 0.757, B: 0.678, A: 1}
	Citrine                        = Color{R: 0.894, G: 0.816, B: 0.039, A: 1}
	GunPowder                      = Color{R: 0.255, G: 0.259, B: 0.341, A: 1}
	FruitSalad                     = Color{R: 0.310, G: 0.616, B: 0.365, A: 1}
	LightBlue                      = Color{R: 0.678, G: 0.847, B: 0.902, A: 1}
	PowderBlue                     = Color{R: 0.690, G: 0.878, B: 0.902, A: 1}
	RockSpray                      = Color{R: 0.729, G: 0.271, B: 0.047, A: 1}
	RoseFog                        = Color{R: 0.906, G: 0.737, B: 0.706, A: 1}
	Como                           = Color{R: 0.318, G: 0.486, B: 0.400, A: 1}
	OrangeRed                      = Color{R: 1.000, G: 0.271, B: 0.000, A: 1}
	Rebel                          = Color{R: 0.235, G: 0.071, B: 0.024, A: 1}
	AmaranthPink                   = Color{R: 0.945, G: 0.612, B: 0.733, A: 1}
	SilverSand                     = Color{R: 0.749, G: 0.757, B: 0.761, A: 1}
	Squirrel                       = Color{R: 0.561, G: 0.506, B: 0.463, A: 1}
	WildSand                       = Color{R: 0.957, G: 0.957, B: 0.957, A: 1}
	Apple                          = Color{R: 0.310, G: 0.659, B: 0.239, A: 1}
	Ceramic                        = Color{R: 0.988, G: 1.000, B: 0.976, A: 1}
	Dune                           = Color{R: 0.220, G: 0.208, B: 0.200, A: 1}
	PurplePizzazz                  = Color{R: 0.996, G: 0.306, B: 0.855, A: 1}
	Saltpan                        = Color{R: 0.945, G: 0.969, B: 0.949, A: 1}
	Bismark                        = Color{R: 0.286, G: 0.443, B: 0.514, A: 1}
	BubbleGum                      = Color{R: 1.000, G: 0.757, B: 0.800, A: 1}
	DawnPink                       = Color{R: 0.953, G: 0.914, B: 0.898, A: 1}
	Ochre                          = Color{R: 0.800, G: 0.467, B: 0.133, A: 1}
	YankeesBlue                    = Color{R: 0.110, G: 0.157, B: 0.255, A: 1}
	CreamBrulee                    = Color{R: 1.000, G: 0.898, B: 0.627, A: 1}
	RocketMetallic                 = Color{R: 0.541, G: 0.498, B: 0.502, A: 1}
	Turbo                          = Color{R: 0.980, G: 0.902, B: 0.000, A: 1}
	FlamingoPink                   = Color{R: 0.988, G: 0.557, B: 0.675, A: 1}
	Gurkha                         = Color{R: 0.604, G: 0.584, B: 0.467, A: 1}
	Mirage                         = Color{R: 0.086, G: 0.098, B: 0.157, A: 1}
	Fuego                          = Color{R: 0.745, G: 0.871, B: 0.051, A: 1}
	GhostWhite                     = Color{R: 0.973, G: 0.973, B: 1.000, A: 1}
	MetalPink                      = Color{R: 1.000, G: 0.000, B: 0.992, A: 1}
	PaleRobinEggBlue               = Color{R: 0.588, G: 0.871, B: 0.820, A: 1}
	RedRibbon                      = Color{R: 0.929, G: 0.039, B: 0.247, A: 1}
	DarkBrownTangelo               = Color{R: 0.533, G: 0.396, B: 0.306, A: 1}
	JacksonsPurple                 = Color{R: 0.125, G: 0.125, B: 0.553, A: 1}
	Raspberry                      = Color{R: 0.890, G: 0.043, B: 0.365, A: 1}
	WhiteRock                      = Color{R: 0.918, G: 0.910, B: 0.831, A: 1}
	AfricanViolet                  = Color{R: 0.698, G: 0.518, B: 0.745, A: 1}
	Champagne                      = Color{R: 0.969, G: 0.906, B: 0.808, A: 1}
	Lochmara                       = Color{R: 0.000, G: 0.494, B: 0.780, A: 1}
	Mandarin                       = Color{R: 0.953, G: 0.478, B: 0.282, A: 1}
	BarleyCorn                     = Color{R: 0.651, G: 0.545, B: 0.357, A: 1}
	DenimBlue                      = Color{R: 0.133, G: 0.263, B: 0.714, A: 1}
	Daintree                       = Color{R: 0.004, G: 0.153, B: 0.192, A: 1}
	MoonGlow                       = Color{R: 0.988, G: 0.996, B: 0.855, A: 1}
	Chalky                         = Color{R: 0.933, G: 0.843, B: 0.580, A: 1}
	JackoBean                      = Color{R: 0.180, G: 0.098, B: 0.020, A: 1}
	PaleSilver                     = Color{R: 0.788, G: 0.753, B: 0.733, A: 1}
	SugarCane                      = Color{R: 0.976, G: 1.000, B: 0.965, A: 1}
	DelRio                         = Color{R: 0.690, G: 0.604, B: 0.584, A: 1}
	FaluRed                        = Color{R: 0.502, G: 0.094, B: 0.094, A: 1}
	PigeonPost                     = Color{R: 0.686, G: 0.741, B: 0.851, A: 1}
	AntiqueBrass                   = Color{R: 0.804, G: 0.584, B: 0.459, A: 1}
	AuroMetalSaurus                = Color{R: 0.431, G: 0.498, B: 0.502, A: 1}
	CeruleanBlue                   = Color{R: 0.165, G: 0.322, B: 0.745, A: 1}
	Cruise                         = Color{R: 0.710, G: 0.925, B: 0.875, A: 1}
	Cyan                           = Color{R: 0.000, G: 1.000, B: 1.000, A: 1}
	Puce                           = Color{R: 0.800, G: 0.533, B: 0.600, A: 1}
	ScreaminGreen                  = Color{R: 0.400, G: 1.000, B: 0.400, A: 1}
	BlueMagentaViolet              = Color{R: 0.333, G: 0.208, B: 0.573, A: 1}
	Meteorite                      = Color{R: 0.235, G: 0.122, B: 0.463, A: 1}
	SweetCorn                      = Color{R: 0.984, G: 0.918, B: 0.549, A: 1}
	Artichoke                      = Color{R: 0.561, G: 0.592, B: 0.475, A: 1}
	BlackShadows                   = Color{R: 0.749, G: 0.686, B: 0.698, A: 1}
	BrightLavender                 = Color{R: 0.749, G: 0.580, B: 0.894, A: 1}
	Cowboy                         = Color{R: 0.302, G: 0.157, B: 0.176, A: 1}
	CuriousBlue                    = Color{R: 0.145, G: 0.588, B: 0.820, A: 1}
	PurpleNavy                     = Color{R: 0.306, G: 0.318, B: 0.502, A: 1}
	DairyCream                     = Color{R: 0.976, G: 0.894, B: 0.737, A: 1}
	EarlyDawn                      = Color{R: 1.000, G: 0.976, B: 0.902, A: 1}
	EggSour                        = Color{R: 1.000, G: 0.957, B: 0.867, A: 1}
	EveningSea                     = Color{R: 0.008, G: 0.306, B: 0.275, A: 1}
	LaRioja                        = Color{R: 0.702, G: 0.757, B: 0.063, A: 1}
	LightFrenchBeige               = Color{R: 0.784, G: 0.678, B: 0.498, A: 1}
	MountainMeadow                 = Color{R: 0.188, G: 0.729, B: 0.561, A: 1}
	NonPhotoBlue                   = Color{R: 0.643, G: 0.867, B: 0.929, A: 1}
	BananaMania                    = Color{R: 0.980, G: 0.906, B: 0.710, A: 1}
	CarouselPink                   = Color{R: 0.976, G: 0.878, B: 0.929, A: 1}
	Cork                           = Color{R: 0.251, G: 0.161, B: 0.114, A: 1}
	FrenchFuchsia                  = Color{R: 0.992, G: 0.247, B: 0.573, A: 1}
	GableGreen                     = Color{R: 0.086, G: 0.208, B: 0.192, A: 1}
	Twilight                       = Color{R: 0.894, G: 0.812, B: 0.871, A: 1}
	Calico                         = Color{R: 0.878, G: 0.753, B: 0.584, A: 1}
	Maverick                       = Color{R: 0.847, G: 0.761, B: 0.835, A: 1}
	MediumSpringGreen              = Color{R: 0.000, G: 0.980, B: 0.604, A: 1}
	Sienna                         = Color{R: 0.533, G: 0.176, B: 0.090, A: 1}
	Timberwolf                     = Color{R: 0.859, G: 0.843, B: 0.824, A: 1}
	Sangria                        = Color{R: 0.573, G: 0.000, B: 0.039, A: 1}
	WildBlueYonder                 = Color{R: 0.635, G: 0.678, B: 0.816, A: 1}
	CandyAppleRed                  = Color{R: 1.000, G: 0.031, B: 0.000, A: 1}
	GraySuit                       = Color{R: 0.757, G: 0.745, B: 0.804, A: 1}
	PaleSpringBud                  = Color{R: 0.925, G: 0.922, B: 0.741, A: 1}
	RazzmicBerry                   = Color{R: 0.553, G: 0.306, B: 0.522, A: 1}
	RedStage                       = Color{R: 0.816, G: 0.373, B: 0.016, A: 1}
	Crusoe                         = Color{R: 0.000, G: 0.282, B: 0.086, A: 1}
	ImperialBlue                   = Color{R: 0.000, G: 0.137, B: 0.584, A: 1}
	SaddleBrown                    = Color{R: 0.545, G: 0.271, B: 0.075, A: 1}
	Tolopea                        = Color{R: 0.106, G: 0.008, B: 0.271, A: 1}
	WillowGrove                    = Color{R: 0.396, G: 0.455, B: 0.365, A: 1}
	RegalBlue                      = Color{R: 0.004, G: 0.247, B: 0.416, A: 1}
	Riptide                        = Color{R: 0.545, G: 0.902, B: 0.847, A: 1}
	RoyalBlue                      = Color{R: 0.255, G: 0.412, B: 0.882, A: 1}
	CharlestonGreen                = Color{R: 0.137, G: 0.169, B: 0.169, A: 1}
	ChinaRose                      = Color{R: 0.659, G: 0.318, B: 0.431, A: 1}
	LasPalmas                      = Color{R: 0.776, G: 0.902, B: 0.063, A: 1}
	Pear                           = Color{R: 0.820, G: 0.886, B: 0.192, A: 1}
	PullmanGreen                   = Color{R: 0.231, G: 0.200, B: 0.110, A: 1}
	ClayAsh                        = Color{R: 0.741, G: 0.784, B: 0.702, A: 1}
	Deluge                         = Color{R: 0.459, G: 0.388, B: 0.659, A: 1}
	Narvik                         = Color{R: 0.929, G: 0.976, B: 0.945, A: 1}
	OrientalPink                   = Color{R: 0.776, G: 0.569, B: 0.569, A: 1}
	Avocado                        = Color{R: 0.337, G: 0.510, B: 0.012, A: 1}
	DarkFern                       = Color{R: 0.039, G: 0.282, B: 0.051, A: 1}
	PattensBlue                    = Color{R: 0.871, G: 0.961, B: 1.000, A: 1}
	TaupeGray                      = Color{R: 0.545, G: 0.522, B: 0.537, A: 1}
	MidGray                        = Color{R: 0.373, G: 0.373, B: 0.431, A: 1}
	SeaBuckthorn                   = Color{R: 0.984, G: 0.631, B: 0.161, A: 1}
	Sunglo                         = Color{R: 0.882, G: 0.408, B: 0.396, A: 1}
	UABlue                         = Color{R: 0.000, G: 0.200, B: 0.667, A: 1}
	Sunshade                       = Color{R: 1.000, G: 0.620, B: 0.173, A: 1}
	VividTangerine                 = Color{R: 1.000, G: 0.627, B: 0.537, A: 1}
	Foam                           = Color{R: 0.847, G: 0.988, B: 0.980, A: 1}
	GreenHouse                     = Color{R: 0.141, G: 0.314, B: 0.059, A: 1}
	MagentaHaze                    = Color{R: 0.624, G: 0.271, B: 0.463, A: 1}
	Nugget                         = Color{R: 0.773, G: 0.600, B: 0.133, A: 1}
	PakistanGreen                  = Color{R: 0.000, G: 0.400, B: 0.000, A: 1}
	OceanBlue                      = Color{R: 0.310, G: 0.259, B: 0.710, A: 1}
	Whiskey                        = Color{R: 0.835, G: 0.604, B: 0.435, A: 1}
	Alabaster                      = Color{R: 0.980, G: 0.980, B: 0.980, A: 1}
	Birch                          = Color{R: 0.216, G: 0.188, B: 0.129, A: 1}
	Eucalyptus                     = Color{R: 0.267, G: 0.843, B: 0.659, A: 1}
	LightTurquoise                 = Color{R: 0.686, G: 0.933, B: 0.933, A: 1}
	MaximumBlue                    = Color{R: 0.278, G: 0.671, B: 0.800, A: 1}
	NeonCarrot                     = Color{R: 1.000, G: 0.639, B: 0.263, A: 1}
	RoyalHeath                     = Color{R: 0.671, G: 0.204, B: 0.447, A: 1}
	SasquatchSocks                 = Color{R: 1.000, G: 0.275, B: 0.506, A: 1}
	Loblolly                       = Color{R: 0.741, G: 0.788, B: 0.808, A: 1}
	Sirocco                        = Color{R: 0.443, G: 0.502, B: 0.502, A: 1}
	Tequila                        = Color{R: 1.000, G: 0.902, B: 0.780, A: 1}
	Windsor                        = Color{R: 0.235, G: 0.031, B: 0.471, A: 1}
	Russett                        = Color{R: 0.459, G: 0.353, B: 0.341, A: 1}
	WillowBrook                    = Color{R: 0.875, G: 0.925, B: 0.855, A: 1}
	AliceBlue                      = Color{R: 0.941, G: 0.973, B: 1.000, A: 1}
	DeepCerise                     = Color{R: 0.855, G: 0.196, B: 0.529, A: 1}
	DeepViolet                     = Color{R: 0.200, G: 0.000, B: 0.400, A: 1}
	Empress                        = Color{R: 0.506, G: 0.451, B: 0.467, A: 1}
	Orange                         = Color{R: 1.000, G: 0.498, B: 0.000, A: 1}
	Cascade                        = Color{R: 0.545, G: 0.663, B: 0.647, A: 1}
	Cement                         = Color{R: 0.553, G: 0.463, B: 0.384, A: 1}
	Punch                          = Color{R: 0.863, G: 0.263, B: 0.200, A: 1}
	WildWillow                     = Color{R: 0.725, G: 0.769, B: 0.416, A: 1}
	KeyLime                        = Color{R: 0.910, G: 0.957, B: 0.549, A: 1}
	LightSlateGray                 = Color{R: 0.467, G: 0.533, B: 0.600, A: 1}
	SwansDown                      = Color{R: 0.863, G: 0.941, B: 0.918, A: 1}
	Boulder                        = Color{R: 0.478, G: 0.478, B: 0.478, A: 1}
	BrightGreen                    = Color{R: 0.400, G: 1.000, B: 0.000, A: 1}
	DarkMidnightBlue               = Color{R: 0.000, G: 0.200, B: 0.400, A: 1}
	DeepLilac                      = Color{R: 0.600, G: 0.333, B: 0.733, A: 1}
	GulfBlue                       = Color{R: 0.020, G: 0.086, B: 0.341, A: 1}
	BrightCerulean                 = Color{R: 0.114, G: 0.675, B: 0.839, A: 1}
	LemonChiffon                   = Color{R: 1.000, G: 0.980, B: 0.804, A: 1}
	MistGray                       = Color{R: 0.769, G: 0.769, B: 0.737, A: 1}
	StTropaz                       = Color{R: 0.176, G: 0.337, B: 0.608, A: 1}
	AstronautBlue                  = Color{R: 0.004, G: 0.243, B: 0.384, A: 1}
	Oxley                          = Color{R: 0.467, G: 0.620, B: 0.525, A: 1}
	RoyalFuchsia                   = Color{R: 0.792, G: 0.173, B: 0.573, A: 1}
	Tallow                         = Color{R: 0.659, G: 0.647, B: 0.537, A: 1}
	Auburn                         = Color{R: 0.647, G: 0.165, B: 0.165, A: 1}
	Charlotte                      = Color{R: 0.729, G: 0.933, B: 0.976, A: 1}
	Ebb                            = Color{R: 0.914, G: 0.890, B: 0.890, A: 1}
	Saffron                        = Color{R: 0.957, G: 0.769, B: 0.188, A: 1}
	SeaBlue                        = Color{R: 0.000, G: 0.412, B: 0.580, A: 1}
	CoralTree                      = Color{R: 0.659, G: 0.420, B: 0.420, A: 1}
	PaleViolet                     = Color{R: 0.800, G: 0.600, B: 1.000, A: 1}
	Kobi                           = Color{R: 0.906, G: 0.624, B: 0.769, A: 1}
	Nomad                          = Color{R: 0.729, G: 0.694, B: 0.635, A: 1}
	ViridianGreen                  = Color{R: 0.000, G: 0.588, B: 0.596, A: 1}
	Burnham                        = Color{R: 0.000, G: 0.180, B: 0.125, A: 1}
	GladeGreen                     = Color{R: 0.380, G: 0.518, B: 0.373, A: 1}
	GumLeaf                        = Color{R: 0.714, G: 0.827, B: 0.749, A: 1}
	HeatheredGray                  = Color{R: 0.714, G: 0.690, B: 0.584, A: 1}
	Jaffa                          = Color{R: 0.937, G: 0.525, B: 0.247, A: 1}
	Scandal                        = Color{R: 0.812, G: 0.980, B: 0.957, A: 1}
	Aztec                          = Color{R: 0.051, G: 0.110, B: 0.098, A: 1}
	FieryOrange                    = Color{R: 0.702, G: 0.322, B: 0.075, A: 1}
	GiantsClub                     = Color{R: 0.690, G: 0.361, B: 0.322, A: 1}
	OsloGray                       = Color{R: 0.529, G: 0.553, B: 0.569, A: 1}
	Prim                           = Color{R: 0.941, G: 0.886, B: 0.925, A: 1}
	BlueChalk                      = Color{R: 0.945, G: 0.914, B: 1.000, A: 1}
	Cerulean                       = Color{R: 0.000, G: 0.482, B: 0.655, A: 1}
	Sauvignon                      = Color{R: 1.000, G: 0.961, B: 0.953, A: 1}
	MunsellYellow                  = Color{R: 0.937, G: 0.800, B: 0.000, A: 1}
	Tulip                          = Color{R: 1.000, G: 0.529, B: 0.553, A: 1}
	Arapawa                        = Color{R: 0.067, G: 0.047, B: 0.424, A: 1}
	BullShot                       = Color{R: 0.525, G: 0.302, B: 0.118, A: 1}
	Derby                          = Color{R: 1.000, G: 0.933, B: 0.847, A: 1}
	FloralWhite                    = Color{R: 1.000, G: 0.980, B: 0.941, A: 1}
	GiantsOrange                   = Color{R: 0.996, G: 0.353, B: 0.114, A: 1}
	Melanie                        = Color{R: 0.894, G: 0.761, B: 0.835, A: 1}
	MulberryWood                   = Color{R: 0.361, G: 0.020, B: 0.212, A: 1}
	MysticMaroon                   = Color{R: 0.678, G: 0.263, B: 0.475, A: 1}
	BlackMarlin                    = Color{R: 0.243, G: 0.173, B: 0.110, A: 1}
	CornHarvest                    = Color{R: 0.545, G: 0.420, B: 0.043, A: 1}
	DeepBronze                     = Color{R: 0.290, G: 0.188, B: 0.016, A: 1}
	HitGray                        = Color{R: 0.631, G: 0.678, B: 0.710, A: 1}
	LemonGlacier                   = Color{R: 0.992, G: 1.000, B: 0.000, A: 1}
	Pistachio                      = Color{R: 0.576, G: 0.773, B: 0.447, A: 1}
	Tranquil                       = Color{R: 0.902, G: 1.000, B: 1.000, A: 1}
	Crete                          = Color{R: 0.451, G: 0.471, B: 0.161, A: 1}
	IndianYellow                   = Color{R: 0.890, G: 0.659, B: 0.341, A: 1}
	LittleBoyBlue                  = Color{R: 0.424, G: 0.627, B: 0.863, A: 1}
	OxfordBlue                     = Color{R: 0.000, G: 0.129, B: 0.278, A: 1}
	Purple                         = Color{R: 0.502, G: 0.000, B: 0.502, A: 1}
	USAFABlue                      = Color{R: 0.000, G: 0.310, B: 0.596, A: 1}
	CafeNoir                       = Color{R: 0.294, G: 0.212, B: 0.129, A: 1}
	LightBrilliantRed              = Color{R: 0.996, G: 0.180, B: 0.180, A: 1}
	ShipGray                       = Color{R: 0.243, G: 0.227, B: 0.267, A: 1}
	TahunaSands                    = Color{R: 0.933, G: 0.941, B: 0.784, A: 1}
	TitanWhite                     = Color{R: 0.941, G: 0.933, B: 1.000, A: 1}
	PrairieSand                    = Color{R: 0.604, G: 0.220, B: 0.125, A: 1}
	Chenin                         = Color{R: 0.875, G: 0.804, B: 0.435, A: 1}
	DarkOliveGreen                 = Color{R: 0.333, G: 0.420, B: 0.184, A: 1}
	FoggyGray                      = Color{R: 0.796, G: 0.792, B: 0.714, A: 1}
	Hurricane                      = Color{R: 0.529, G: 0.486, B: 0.482, A: 1}
	PaleMagentaPink                = Color{R: 1.000, G: 0.600, B: 0.800, A: 1}
	BlueZodiac                     = Color{R: 0.075, G: 0.149, B: 0.302, A: 1}
	Pomegranate                    = Color{R: 0.953, G: 0.278, B: 0.137, A: 1}
	SoftAmber                      = Color{R: 0.820, G: 0.776, B: 0.706, A: 1}
	X11Gray                        = Color{R: 0.745, G: 0.745, B: 0.745, A: 1}
	ChlorophyllGreen               = Color{R: 0.290, G: 1.000, B: 0.000, A: 1}
	MediumSpringBud                = Color{R: 0.788, G: 0.863, B: 0.529, A: 1}
	Portica                        = Color{R: 0.976, G: 0.902, B: 0.388, A: 1}
	CopperPenny                    = Color{R: 0.678, G: 0.435, B: 0.412, A: 1}
	TahitiGold                     = Color{R: 0.914, G: 0.486, B: 0.027, A: 1}
	UnbleachedSilk                 = Color{R: 1.000, G: 0.867, B: 0.792, A: 1}
	Blackcurrant                   = Color{R: 0.196, G: 0.161, B: 0.227, A: 1}
	BrinkPink                      = Color{R: 0.984, G: 0.376, B: 0.498, A: 1}
	Raffia                         = Color{R: 0.918, G: 0.855, B: 0.722, A: 1}
	SantaFe                        = Color{R: 0.694, G: 0.427, B: 0.322, A: 1}
	VividOrangePeel                = Color{R: 1.000, G: 0.627, B: 0.000, A: 1}
	Allports                       = Color{R: 0.000, G: 0.463, B: 0.639, A: 1}
	Bridesmaid                     = Color{R: 0.996, G: 0.941, B: 0.925, A: 1}
	CardinGreen                    = Color{R: 0.004, G: 0.212, B: 0.110, A: 1}
	Danube                         = Color{R: 0.376, G: 0.576, B: 0.820, A: 1}
	MossGreen                      = Color{R: 0.541, G: 0.604, B: 0.357, A: 1}
	AntiqueRuby                    = Color{R: 0.518, G: 0.106, B: 0.176, A: 1}
	EagleGreen                     = Color{R: 0.000, G: 0.286, B: 0.325, A: 1}
	ForgetMeNot                    = Color{R: 1.000, G: 0.945, B: 0.933, A: 1}
	Ironstone                      = Color{R: 0.525, G: 0.282, B: 0.235, A: 1}
	ShipCove                       = Color{R: 0.471, G: 0.545, B: 0.729, A: 1}
	Veronica                       = Color{R: 0.627, G: 0.125, B: 0.941, A: 1}
	Wistful                        = Color{R: 0.643, G: 0.651, B: 0.827, A: 1}
	Bizarre                        = Color{R: 0.933, G: 0.871, B: 0.855, A: 1}
	Cornsilk                       = Color{R: 1.000, G: 0.973, B: 0.863, A: 1}
	Cream                          = Color{R: 1.000, G: 0.992, B: 0.816, A: 1}
	DarkCoral                      = Color{R: 0.804, G: 0.357, B: 0.271, A: 1}
	Tradewind                      = Color{R: 0.373, G: 0.702, B: 0.675, A: 1}
	LightSalmon                    = Color{R: 1.000, G: 0.627, B: 0.478, A: 1}
	MoonMist                       = Color{R: 0.863, G: 0.867, B: 0.800, A: 1}
	SchoolBusYellow                = Color{R: 1.000, G: 0.847, B: 0.000, A: 1}
	Seaweed                        = Color{R: 0.106, G: 0.184, B: 0.067, A: 1}
	TomThumb                       = Color{R: 0.247, G: 0.345, B: 0.231, A: 1}
	Aquamarine                     = Color{R: 0.498, G: 1.000, B: 0.831, A: 1}
	JuneBud                        = Color{R: 0.741, G: 0.855, B: 0.341, A: 1}
	Tuna                           = Color{R: 0.208, G: 0.208, B: 0.259, A: 1}
	Amazon                         = Color{R: 0.231, G: 0.478, B: 0.341, A: 1}
	Cherrywood                     = Color{R: 0.396, G: 0.102, B: 0.078, A: 1}
	DarkMossGreen                  = Color{R: 0.290, G: 0.365, B: 0.137, A: 1}
	Nandor                         = Color{R: 0.294, G: 0.365, B: 0.322, A: 1}
	Parsley                        = Color{R: 0.075, G: 0.310, B: 0.098, A: 1}
	Gamboge                        = Color{R: 0.894, G: 0.608, B: 0.059, A: 1}
	Locust                         = Color{R: 0.659, G: 0.686, B: 0.557, A: 1}
	PastelViolet                   = Color{R: 0.796, G: 0.600, B: 0.788, A: 1}
	Sazerac                        = Color{R: 1.000, G: 0.957, B: 0.878, A: 1}
	CannonBlack                    = Color{R: 0.145, G: 0.090, B: 0.024, A: 1}
	KenyanCopper                   = Color{R: 0.486, G: 0.110, B: 0.020, A: 1}
	Ultramarine                    = Color{R: 0.247, G: 0.000, B: 1.000, A: 1}
	Sahara                         = Color{R: 0.718, G: 0.635, B: 0.078, A: 1}
	TulipTree                      = Color{R: 0.918, G: 0.702, B: 0.231, A: 1}
	VerdunGreen                    = Color{R: 0.286, G: 0.329, B: 0.000, A: 1}
	Aero                           = Color{R: 0.486, G: 0.725, B: 0.910, A: 1}
	Flamenco                       = Color{R: 1.000, G: 0.490, B: 0.027, A: 1}
	Graphite                       = Color{R: 0.145, G: 0.086, B: 0.027, A: 1}
	OutrageousOrange               = Color{R: 1.000, G: 0.431, B: 0.290, A: 1}
	RaisinBlack                    = Color{R: 0.141, G: 0.129, B: 0.141, A: 1}
	Eagle                          = Color{R: 0.714, G: 0.729, B: 0.643, A: 1}
	Gothic                         = Color{R: 0.427, G: 0.573, B: 0.631, A: 1}
	RichMaroon                     = Color{R: 0.690, G: 0.188, B: 0.376, A: 1}
	BronzeOlive                    = Color{R: 0.306, G: 0.259, B: 0.047, A: 1}
	Domino                         = Color{R: 0.557, G: 0.467, B: 0.369, A: 1}
	GullGray                       = Color{R: 0.616, G: 0.675, B: 0.718, A: 1}
	PaleChestnut                   = Color{R: 0.867, G: 0.678, B: 0.686, A: 1}
	WhiteLilac                     = Color{R: 0.973, G: 0.969, B: 0.988, A: 1}
	Azure                          = Color{R: 0.000, G: 0.498, B: 1.000, A: 1}
	SeaNymph                       = Color{R: 0.471, G: 0.639, B: 0.612, A: 1}
	Claret                         = Color{R: 0.498, G: 0.090, B: 0.204, A: 1}
	Tangerine                      = Color{R: 0.949, G: 0.522, B: 0.000, A: 1}
	BrightUbe                      = Color{R: 0.820, G: 0.624, B: 0.910, A: 1}
	Heliotrope                     = Color{R: 0.875, G: 0.451, B: 1.000, A: 1}
	HoneyFlower                    = Color{R: 0.310, G: 0.110, B: 0.439, A: 1}
	PaleCerulean                   = Color{R: 0.608, G: 0.769, B: 0.886, A: 1}
	PigmentBlue                    = Color{R: 0.200, G: 0.200, B: 0.600, A: 1}
	WinterWizard                   = Color{R: 0.627, G: 0.902, B: 1.000, A: 1}
	EbonyClay                      = Color{R: 0.149, G: 0.157, B: 0.231, A: 1}
	EerieBlack                     = Color{R: 0.106, G: 0.106, B: 0.106, A: 1}
	Mabel                          = Color{R: 0.851, G: 0.969, B: 1.000, A: 1}
	PeriglacialBlue                = Color{R: 0.882, G: 0.902, B: 0.839, A: 1}
	Tide                           = Color{R: 0.749, G: 0.722, B: 0.690, A: 1}
	BigFootFeet                    = Color{R: 0.910, G: 0.557, B: 0.353, A: 1}
	ChaletGreen                    = Color{R: 0.318, G: 0.431, B: 0.239, A: 1}
	LightningYellow                = Color{R: 0.988, G: 0.753, B: 0.118, A: 1}
	Mystic                         = Color{R: 0.839, G: 0.322, B: 0.510, A: 1}
	BlueLagoon                     = Color{R: 0.675, G: 0.898, B: 0.933, A: 1}
	DonJuan                        = Color{R: 0.365, G: 0.298, B: 0.318, A: 1}
	Merlin                         = Color{R: 0.255, G: 0.235, B: 0.216, A: 1}
	PaleGoldenrod                  = Color{R: 0.933, G: 0.910, B: 0.667, A: 1}
	Astronaut                      = Color{R: 0.157, G: 0.227, B: 0.467, A: 1}
	CalPolyGreen                   = Color{R: 0.118, G: 0.302, B: 0.169, A: 1}
	Lucky                          = Color{R: 0.686, G: 0.624, B: 0.110, A: 1}
	MandysPink                     = Color{R: 0.949, G: 0.765, B: 0.698, A: 1}
	Woodrush                       = Color{R: 0.188, G: 0.165, B: 0.059, A: 1}
	Tabasco                        = Color{R: 0.627, G: 0.153, B: 0.071, A: 1}
	AmaranthRed                    = Color{R: 0.827, G: 0.129, B: 0.176, A: 1}
	LightGrayishMagenta            = Color{R: 0.800, G: 0.600, B: 0.800, A: 1}
	Mongoose                       = Color{R: 0.710, G: 0.635, B: 0.498, A: 1}
	MoonRaker                      = Color{R: 0.839, G: 0.808, B: 0.965, A: 1}
	NaplesYellow                   = Color{R: 0.980, G: 0.855, B: 0.369, A: 1}
	DeepCove                       = Color{R: 0.020, G: 0.063, B: 0.251, A: 1}
	DoubleColonialWhite            = Color{R: 0.933, G: 0.890, B: 0.678, A: 1}
	Khaki                          = Color{R: 0.765, G: 0.690, B: 0.569, A: 1}
	MayaBlue                       = Color{R: 0.451, G: 0.761, B: 0.984, A: 1}
	Sinopia                        = Color{R: 0.796, G: 0.255, B: 0.043, A: 1}
	DartmouthGreen                 = Color{R: 0.000, G: 0.439, B: 0.235, A: 1}
	MediumCandyAppleRed            = Color{R: 0.886, G: 0.024, B: 0.173, A: 1}
	MintJulep                      = Color{R: 0.945, G: 0.933, B: 0.757, A: 1}
	PalePrim                       = Color{R: 0.992, G: 0.996, B: 0.722, A: 1}
	EastSide                       = Color{R: 0.675, G: 0.569, B: 0.808, A: 1}
	LightMossGreen                 = Color{R: 0.678, G: 0.875, B: 0.678, A: 1}
	BlackSqueeze                   = Color{R: 0.949, G: 0.980, B: 0.980, A: 1}
	FriarGray                      = Color{R: 0.502, G: 0.494, B: 0.475, A: 1}
	TrueBlue                       = Color{R: 0.000, G: 0.451, B: 0.812, A: 1}
	Blush                          = Color{R: 0.871, G: 0.365, B: 0.514, A: 1}
	PastelBlue                     = Color{R: 0.682, G: 0.776, B: 0.812, A: 1}
	SAEECEAmber                    = Color{R: 1.000, G: 0.494, B: 0.000, A: 1}
	SpanishRed                     = Color{R: 0.902, G: 0.000, B: 0.149, A: 1}
	Geraldine                      = Color{R: 0.984, G: 0.537, B: 0.537, A: 1}
	GoldDrop                       = Color{R: 0.945, G: 0.510, B: 0.000, A: 1}
	PersianBlue                    = Color{R: 0.110, G: 0.224, B: 0.733, A: 1}
	PersianOrange                  = Color{R: 0.851, G: 0.565, B: 0.345, A: 1}
	RedOrange                      = Color{R: 1.000, G: 0.325, B: 0.286, A: 1}
	BlueChill                      = Color{R: 0.047, G: 0.537, B: 0.565, A: 1}
	CadetBlue                      = Color{R: 0.373, G: 0.620, B: 0.627, A: 1}
	CodGray                        = Color{R: 0.043, G: 0.043, B: 0.043, A: 1}
	Hopbush                        = Color{R: 0.816, G: 0.427, B: 0.631, A: 1}
	VividTangelo                   = Color{R: 0.941, G: 0.455, B: 0.153, A: 1}
	Zanah                          = Color{R: 0.855, G: 0.925, B: 0.839, A: 1}
	CeruleanFrost                  = Color{R: 0.427, G: 0.608, B: 0.765, A: 1}
	CostaDelSol                    = Color{R: 0.380, G: 0.365, B: 0.188, A: 1}
	RubineRed                      = Color{R: 0.820, G: 0.000, B: 0.337, A: 1}
	Sinbad                         = Color{R: 0.624, G: 0.843, B: 0.827, A: 1}
	Xanadu                         = Color{R: 0.451, G: 0.525, B: 0.471, A: 1}
	BleachedCedar                  = Color{R: 0.173, G: 0.129, B: 0.200, A: 1}
	WillpowerOrange                = Color{R: 0.992, G: 0.345, B: 0.000, A: 1}
	Deer                           = Color{R: 0.729, G: 0.529, B: 0.349, A: 1}
	JetStream                      = Color{R: 0.710, G: 0.824, B: 0.808, A: 1}
	MagicMint                      = Color{R: 0.667, G: 0.941, B: 0.820, A: 1}
	Pumice                         = Color{R: 0.761, G: 0.792, B: 0.769, A: 1}
	FrenchPlum                     = Color{R: 0.506, G: 0.078, B: 0.325, A: 1}
	KingfisherDaisy                = Color{R: 0.243, G: 0.016, B: 0.502, A: 1}
	SkyBlue                        = Color{R: 0.529, G: 0.808, B: 0.922, A: 1}
	Trout                          = Color{R: 0.290, G: 0.306, B: 0.353, A: 1}
	WebOrange                      = Color{R: 1.000, G: 0.647, B: 0.000, A: 1}
	BlackRussian                   = Color{R: 0.039, G: 0.000, B: 0.110, A: 1}
	FreshAir                       = Color{R: 0.651, G: 0.906, B: 1.000, A: 1}
	Mauve                          = Color{R: 0.878, G: 0.690, B: 1.000, A: 1}
	OysterPink                     = Color{R: 0.914, G: 0.808, B: 0.804, A: 1}
	ScotchMist                     = Color{R: 1.000, G: 0.984, B: 0.863, A: 1}
	Celery                         = Color{R: 0.722, G: 0.761, B: 0.365, A: 1}
	Chinook                        = Color{R: 0.659, G: 0.890, B: 0.741, A: 1}
	FreshEggplant                  = Color{R: 0.600, G: 0.000, B: 0.400, A: 1}
	Gorse                          = Color{R: 1.000, G: 0.945, B: 0.310, A: 1}
	JustRight                      = Color{R: 0.925, G: 0.804, B: 0.725, A: 1}
	Telemagenta                    = Color{R: 0.812, G: 0.204, B: 0.463, A: 1}
	TuftsBlue                      = Color{R: 0.255, G: 0.490, B: 0.757, A: 1}
	BurnishedBrown                 = Color{R: 0.631, G: 0.478, B: 0.455, A: 1}
	Calypso                        = Color{R: 0.192, G: 0.447, B: 0.553, A: 1}
	JungleGreen                    = Color{R: 0.161, G: 0.671, B: 0.529, A: 1}
	NightRider                     = Color{R: 0.122, G: 0.071, B: 0.059, A: 1}
	PumpkinSkin                    = Color{R: 0.694, G: 0.380, B: 0.043, A: 1}
	Zuccini                        = Color{R: 0.016, G: 0.251, B: 0.133, A: 1}
	Viking                         = Color{R: 0.392, G: 0.800, B: 0.859, A: 1}
	VividRedTangelo                = Color{R: 0.875, G: 0.380, B: 0.141, A: 1}
	Byzantium                      = Color{R: 0.439, G: 0.161, B: 0.388, A: 1}
	EcruWhite                      = Color{R: 0.961, G: 0.953, B: 0.898, A: 1}
	FireEngineRed                  = Color{R: 0.808, G: 0.125, B: 0.161, A: 1}
	GoldenGateBridge               = Color{R: 0.753, G: 0.212, B: 0.173, A: 1}
	Hacienda                       = Color{R: 0.596, G: 0.506, B: 0.106, A: 1}
	BlackRock                      = Color{R: 0.051, G: 0.012, B: 0.196, A: 1}
	DeepLemon                      = Color{R: 0.961, G: 0.780, B: 0.102, A: 1}
	OrchidPink                     = Color{R: 0.949, G: 0.741, B: 0.804, A: 1}
	Paradiso                       = Color{R: 0.192, G: 0.490, B: 0.510, A: 1}
	Sapling                        = Color{R: 0.871, G: 0.831, B: 0.643, A: 1}
	DarkSpringGreen                = Color{R: 0.090, G: 0.447, B: 0.271, A: 1}
	DustyGray                      = Color{R: 0.659, G: 0.596, B: 0.608, A: 1}
	IndianTan                      = Color{R: 0.302, G: 0.118, B: 0.004, A: 1}
	RedBeech                       = Color{R: 0.482, G: 0.220, B: 0.004, A: 1}
	Tutu                           = Color{R: 1.000, G: 0.945, B: 0.976, A: 1}
	MauveTaupe                     = Color{R: 0.569, G: 0.373, B: 0.427, A: 1}
	MediumVermilion                = Color{R: 0.851, G: 0.376, B: 0.231, A: 1}
	RoseRed                        = Color{R: 0.761, G: 0.118, B: 0.337, A: 1}
	Bandicoot                      = Color{R: 0.522, G: 0.518, B: 0.439, A: 1}
	Clementine                     = Color{R: 0.914, G: 0.431, B: 0.000, A: 1}
	DeepKoamaru                    = Color{R: 0.200, G: 0.200, B: 0.400, A: 1}
	DustStorm                      = Color{R: 0.898, G: 0.800, B: 0.788, A: 1}
	Endeavour                      = Color{R: 0.000, G: 0.337, B: 0.655, A: 1}
	Silk                           = Color{R: 0.741, G: 0.694, B: 0.659, A: 1}
	FerrariRed                     = Color{R: 1.000, G: 0.157, B: 0.000, A: 1}
	Juniper                        = Color{R: 0.427, G: 0.573, B: 0.573, A: 1}
	OldMossGreen                   = Color{R: 0.525, G: 0.494, B: 0.212, A: 1}
	AquaIsland                     = Color{R: 0.631, G: 0.855, B: 0.843, A: 1}
	Linen                          = Color{R: 0.980, G: 0.941, B: 0.902, A: 1}
	CopperRust                     = Color{R: 0.580, G: 0.278, B: 0.278, A: 1}
	HippieBlue                     = Color{R: 0.345, G: 0.604, B: 0.686, A: 1}
	Parchment                      = Color{R: 0.945, G: 0.914, B: 0.824, A: 1}
	TeaGreen                       = Color{R: 0.816, G: 0.941, B: 0.753, A: 1}
	Camouflage                     = Color{R: 0.235, G: 0.224, B: 0.063, A: 1}
	Leather                        = Color{R: 0.588, G: 0.439, B: 0.349, A: 1}
	RoseGold                       = Color{R: 0.718, G: 0.431, B: 0.475, A: 1}
	Deco                           = Color{R: 0.824, G: 0.855, B: 0.592, A: 1}
	Hemp                           = Color{R: 0.565, G: 0.471, B: 0.455, A: 1}
	PeachOrange                    = Color{R: 1.000, G: 0.800, B: 0.600, A: 1}
	SpringWood                     = Color{R: 0.973, G: 0.965, B: 0.945, A: 1}
	Stiletto                       = Color{R: 0.612, G: 0.200, B: 0.212, A: 1}
	BlueWhale                      = Color{R: 0.016, G: 0.180, B: 0.298, A: 1}
	Goblin                         = Color{R: 0.239, G: 0.490, B: 0.322, A: 1}
	Rope                           = Color{R: 0.557, G: 0.302, B: 0.118, A: 1}
	SweetBrown                     = Color{R: 0.659, G: 0.216, B: 0.192, A: 1}
	WarmBlack                      = Color{R: 0.000, G: 0.259, B: 0.259, A: 1}
	Blue                           = Color{R: 0.000, G: 0.000, B: 1.000, A: 1}
	CafeRoyale                     = Color{R: 0.435, G: 0.267, B: 0.047, A: 1}
	CharmPink                      = Color{R: 0.902, G: 0.561, B: 0.675, A: 1}
	Redwood                        = Color{R: 0.643, G: 0.353, B: 0.322, A: 1}
	Tasman                         = Color{R: 0.812, G: 0.863, B: 0.812, A: 1}
	RYBGreen                       = Color{R: 0.400, G: 0.690, B: 0.196, A: 1}
	Tiber                          = Color{R: 0.024, G: 0.208, B: 0.216, A: 1}
	FringyFlower                   = Color{R: 0.694, G: 0.886, B: 0.757, A: 1}
	Lava                           = Color{R: 0.812, G: 0.063, B: 0.125, A: 1}
	Mamba                          = Color{R: 0.557, G: 0.506, B: 0.565, A: 1}
	OrangePeel                     = Color{R: 1.000, G: 0.624, B: 0.000, A: 1}
	Platinum                       = Color{R: 0.898, G: 0.894, B: 0.886, A: 1}
	ChineseViolet                  = Color{R: 0.522, G: 0.376, B: 0.533, A: 1}
	DodgerBlue                     = Color{R: 0.118, G: 0.565, B: 1.000, A: 1}
	LawnGreen                      = Color{R: 0.486, G: 0.988, B: 0.000, A: 1}
	RedDamask                      = Color{R: 0.855, G: 0.416, B: 0.255, A: 1}
	Valencia                       = Color{R: 0.847, G: 0.267, B: 0.216, A: 1}
	AthsSpecial                    = Color{R: 0.925, G: 0.922, B: 0.808, A: 1}
	Cranberry                      = Color{R: 0.859, G: 0.314, B: 0.475, A: 1}
	Mischka                        = Color{R: 0.820, G: 0.824, B: 0.867, A: 1}
	MunsellRed                     = Color{R: 0.949, G: 0.000, B: 0.235, A: 1}
	Sun                            = Color{R: 0.984, G: 0.675, B: 0.075, A: 1}
	CanaryYellow                   = Color{R: 1.000, G: 0.937, B: 0.000, A: 1}
	ClearDay                       = Color{R: 0.914, G: 1.000, B: 0.992, A: 1}
	DarkPastelRed                  = Color{R: 0.761, G: 0.231, B: 0.133, A: 1}
	OrangeSoda                     = Color{R: 0.980, G: 0.357, B: 0.239, A: 1}
	Red                            = Color{R: 1.000, G: 0.000, B: 0.000, A: 1}
	RedViolet                      = Color{R: 0.780, G: 0.082, B: 0.522, A: 1}
	Bronzetone                     = Color{R: 0.302, G: 0.251, B: 0.059, A: 1}
	Cognac                         = Color{R: 0.624, G: 0.220, B: 0.114, A: 1}
	CraterBrown                    = Color{R: 0.275, G: 0.141, B: 0.145, A: 1}
	DeepFir                        = Color{R: 0.000, G: 0.161, B: 0.000, A: 1}
	EgyptianBlue                   = Color{R: 0.063, G: 0.204, B: 0.651, A: 1}
	Venus                          = Color{R: 0.573, G: 0.522, B: 0.565, A: 1}
	MellowApricot                  = Color{R: 0.973, G: 0.722, B: 0.471, A: 1}
	PaleCornflowerBlue             = Color{R: 0.671, G: 0.804, B: 0.937, A: 1}
	ProcessMagenta                 = Color{R: 1.000, G: 0.000, B: 0.565, A: 1}
	Rock                           = Color{R: 0.302, G: 0.220, B: 0.200, A: 1}
	Loulou                         = Color{R: 0.275, G: 0.043, B: 0.255, A: 1}
	SherpaBlue                     = Color{R: 0.000, G: 0.286, B: 0.314, A: 1}
	TexasRose                      = Color{R: 1.000, G: 0.710, B: 0.333, A: 1}
	LimedOak                       = Color{R: 0.675, G: 0.541, B: 0.337, A: 1}
	RoyalAirForceBlue              = Color{R: 0.365, G: 0.541, B: 0.659, A: 1}
	ChromeYellow                   = Color{R: 1.000, G: 0.655, B: 0.000, A: 1}
	Fallow                         = Color{R: 0.757, G: 0.604, B: 0.420, A: 1}
	GrayNurse                      = Color{R: 0.906, G: 0.925, B: 0.902, A: 1}
	JaggedIce                      = Color{R: 0.761, G: 0.910, B: 0.898, A: 1}
	Jasmine                        = Color{R: 0.973, G: 0.871, B: 0.494, A: 1}
	DarkBurgundy                   = Color{R: 0.467, G: 0.059, B: 0.020, A: 1}
	Grandis                        = Color{R: 1.000, G: 0.827, B: 0.549, A: 1}
	PolishedPine                   = Color{R: 0.365, G: 0.643, B: 0.576, A: 1}
	RussianGreen                   = Color{R: 0.404, G: 0.573, B: 0.404, A: 1}
	Camarone                       = Color{R: 0.000, G: 0.345, B: 0.102, A: 1}
	Corvette                       = Color{R: 0.980, G: 0.827, B: 0.635, A: 1}
	HeavyMetal                     = Color{R: 0.169, G: 0.196, B: 0.157, A: 1}
	Paua                           = Color{R: 0.149, G: 0.012, B: 0.408, A: 1}
	Peru                           = Color{R: 0.804, G: 0.522, B: 0.247, A: 1}
	PaleCyan                       = Color{R: 0.529, G: 0.827, B: 0.973, A: 1}
	Pewter                         = Color{R: 0.588, G: 0.659, B: 0.631, A: 1}
	Punga                          = Color{R: 0.302, G: 0.239, B: 0.078, A: 1}
	AntiqueBronze                  = Color{R: 0.400, G: 0.365, B: 0.118, A: 1}
	Axolotl                        = Color{R: 0.306, G: 0.400, B: 0.286, A: 1}
	BleuDeFrance                   = Color{R: 0.192, G: 0.549, B: 0.906, A: 1}
	HotPink                        = Color{R: 1.000, G: 0.412, B: 0.706, A: 1}
	JellyBean                      = Color{R: 0.855, G: 0.380, B: 0.306, A: 1}
	BajaWhite                      = Color{R: 1.000, G: 0.973, B: 0.820, A: 1}
	GoldenSand                     = Color{R: 0.941, G: 0.859, B: 0.490, A: 1}
	Sandstorm                      = Color{R: 0.925, G: 0.835, B: 0.251, A: 1}
	Casper                         = Color{R: 0.678, G: 0.745, B: 0.820, A: 1}
	PurplePlum                     = Color{R: 0.612, G: 0.318, B: 0.714, A: 1}
	SepiaSkin                      = Color{R: 0.620, G: 0.357, B: 0.251, A: 1}
	SlateGray                      = Color{R: 0.439, G: 0.502, B: 0.565, A: 1}
	Shadow                         = Color{R: 0.541, G: 0.475, B: 0.365, A: 1}
	VeniceBlue                     = Color{R: 0.020, G: 0.349, B: 0.537, A: 1}
	BlackCoral                     = Color{R: 0.329, G: 0.384, B: 0.435, A: 1}
	Eunry                          = Color{R: 0.812, G: 0.639, B: 0.616, A: 1}
	Flint                          = Color{R: 0.435, G: 0.416, B: 0.380, A: 1}
	HookersGreen                   = Color{R: 0.286, G: 0.475, B: 0.420, A: 1}
	Lima                           = Color{R: 0.463, G: 0.741, B: 0.090, A: 1}
	Almond                         = Color{R: 0.937, G: 0.871, B: 0.804, A: 1}
	ArmyGreen                      = Color{R: 0.294, G: 0.325, B: 0.125, A: 1}
	PrincessPerfume                = Color{R: 1.000, G: 0.522, B: 0.812, A: 1}
	SilverLakeBlue                 = Color{R: 0.365, G: 0.537, B: 0.729, A: 1}
	SpanishGreen                   = Color{R: 0.000, G: 0.569, B: 0.314, A: 1}
	MaroonOak                      = Color{R: 0.322, G: 0.047, B: 0.090, A: 1}
	PersianPink                    = Color{R: 0.969, G: 0.498, B: 0.745, A: 1}
	OldHeliotrope                  = Color{R: 0.337, G: 0.235, B: 0.361, A: 1}
	Orient                         = Color{R: 0.004, G: 0.369, B: 0.522, A: 1}
	PinkRaspberry                  = Color{R: 0.596, G: 0.000, B: 0.212, A: 1}
	Charm                          = Color{R: 0.831, G: 0.455, B: 0.580, A: 1}
	Eggshell                       = Color{R: 0.941, G: 0.918, B: 0.839, A: 1}
	EngineeringInternationalOrange = Color{R: 0.729, G: 0.086, B: 0.047, A: 1}
	LightCoral                     = Color{R: 0.941, G: 0.502, B: 0.502, A: 1}
	Oil                            = Color{R: 0.157, G: 0.118, B: 0.082, A: 1}
	TealBlue                       = Color{R: 0.212, G: 0.459, B: 0.533, A: 1}
	Bittersweet                    = Color{R: 0.996, G: 0.435, B: 0.369, A: 1}
	FuscousGray                    = Color{R: 0.329, G: 0.325, B: 0.302, A: 1}
	Sail                           = Color{R: 0.722, G: 0.878, B: 0.976, A: 1}
	Seagull                        = Color{R: 0.502, G: 0.800, B: 0.918, A: 1}
	Valhalla                       = Color{R: 0.169, G: 0.098, B: 0.310, A: 1}
	CGBlue                         = Color{R: 0.000, G: 0.478, B: 0.647, A: 1}
	OliveHaze                      = Color{R: 0.545, G: 0.518, B: 0.439, A: 1}
	Pompadour                      = Color{R: 0.400, G: 0.000, B: 0.271, A: 1}
	Romantic                       = Color{R: 1.000, G: 0.824, B: 0.718, A: 1}
	ShinyShamrock                  = Color{R: 0.373, G: 0.655, B: 0.471, A: 1}
	Amulet                         = Color{R: 0.482, G: 0.624, B: 0.502, A: 1}
	Citron                         = Color{R: 0.624, G: 0.663, B: 0.122, A: 1}
	Concrete                       = Color{R: 0.949, G: 0.949, B: 0.949, A: 1}
	Cupid                          = Color{R: 0.984, G: 0.745, B: 0.855, A: 1}
	Pharlap                        = Color{R: 0.639, G: 0.502, B: 0.482, A: 1}
	Bermuda                        = Color{R: 0.490, G: 0.847, B: 0.776, A: 1}
	BlueRibbon                     = Color{R: 0.000, G: 0.400, B: 1.000, A: 1}
	MintTulip                      = Color{R: 0.769, G: 0.957, B: 0.922, A: 1}
	CareysPink                     = Color{R: 0.824, G: 0.620, B: 0.667, A: 1}
	RoseofSharon                   = Color{R: 0.749, G: 0.333, B: 0.000, A: 1}
	SpringGreen                    = Color{R: 0.000, G: 1.000, B: 0.498, A: 1}
	Studio                         = Color{R: 0.443, G: 0.290, B: 0.698, A: 1}
	CamouflageGreen                = Color{R: 0.471, G: 0.525, B: 0.420, A: 1}
	DeepSeaGreen                   = Color{R: 0.035, G: 0.345, B: 0.349, A: 1}
	DimGray                        = Color{R: 0.412, G: 0.412, B: 0.412, A: 1}
	Napa                           = Color{R: 0.675, G: 0.643, B: 0.580, A: 1}
	ShingleFawn                    = Color{R: 0.420, G: 0.306, B: 0.192, A: 1}
	PurpleMountainMajesty          = Color{R: 0.588, G: 0.471, B: 0.714, A: 1}
	BlueBayoux                     = Color{R: 0.286, G: 0.400, B: 0.475, A: 1}
	Topaz                          = Color{R: 1.000, G: 0.784, B: 0.486, A: 1}
	ShockingPink                   = Color{R: 0.988, G: 0.059, B: 0.753, A: 1}
	SunburntCyclops                = Color{R: 1.000, G: 0.251, B: 0.298, A: 1}
	WestSide                       = Color{R: 1.000, G: 0.569, B: 0.059, A: 1}
	CerisePink                     = Color{R: 0.925, G: 0.231, B: 0.514, A: 1}
	Delta                          = Color{R: 0.643, G: 0.643, B: 0.616, A: 1}
	LuckyPoint                     = Color{R: 0.102, G: 0.102, B: 0.408, A: 1}
	Portage                        = Color{R: 0.545, G: 0.624, B: 0.933, A: 1}
	RusticRed                      = Color{R: 0.282, G: 0.016, B: 0.016, A: 1}
	Candlelight                    = Color{R: 0.988, G: 0.851, B: 0.090, A: 1}
	Dawn                           = Color{R: 0.651, G: 0.635, B: 0.604, A: 1}
	HanBlue                        = Color{R: 0.267, G: 0.424, B: 0.812, A: 1}
	Peat                           = Color{R: 0.443, G: 0.420, B: 0.337, A: 1}
	Snow                           = Color{R: 1.000, G: 0.980, B: 0.980, A: 1}
	PalatinateBlue                 = Color{R: 0.153, G: 0.231, B: 0.886, A: 1}
	VegasGold                      = Color{R: 0.773, G: 0.702, B: 0.345, A: 1}
	Tapa                           = Color{R: 0.482, G: 0.471, B: 0.455, A: 1}
	WhiteLinen                     = Color{R: 0.973, G: 0.941, B: 0.910, A: 1}
	FrenchBlue                     = Color{R: 0.000, G: 0.447, B: 0.733, A: 1}
	GrainBrown                     = Color{R: 0.894, G: 0.835, B: 0.718, A: 1}
	InchWorm                       = Color{R: 0.690, G: 0.890, B: 0.075, A: 1}
	OliveGreen                     = Color{R: 0.710, G: 0.702, B: 0.361, A: 1}
	Rouge                          = Color{R: 0.635, G: 0.231, B: 0.424, A: 1}
	DarkTangerine                  = Color{R: 1.000, G: 0.659, B: 0.071, A: 1}
	PantoneGreen                   = Color{R: 0.000, G: 0.678, B: 0.263, A: 1}
	RawSienna                      = Color{R: 0.839, G: 0.541, B: 0.349, A: 1}
	SpanishCarmine                 = Color{R: 0.820, G: 0.000, B: 0.278, A: 1}
	Jacaranda                      = Color{R: 0.180, G: 0.012, B: 0.161, A: 1}
	Rose                           = Color{R: 1.000, G: 0.000, B: 0.498, A: 1}
	Charcoal                       = Color{R: 0.212, G: 0.271, B: 0.310, A: 1}
	DarkGoldenrod                  = Color{R: 0.722, G: 0.525, B: 0.043, A: 1}
	OriolesOrange                  = Color{R: 0.984, G: 0.310, B: 0.078, A: 1}
	Tussock                        = Color{R: 0.773, G: 0.600, B: 0.294, A: 1}
	OldBurgundy                    = Color{R: 0.263, G: 0.188, B: 0.180, A: 1}
	RoseDust                       = Color{R: 0.620, G: 0.369, B: 0.435, A: 1}
	ScienceBlue                    = Color{R: 0.000, G: 0.400, B: 0.800, A: 1}
	SelectiveYellow                = Color{R: 1.000, G: 0.729, B: 0.000, A: 1}
	SurfCrest                      = Color{R: 0.812, G: 0.898, B: 0.824, A: 1}
	AlmondFrost                    = Color{R: 0.565, G: 0.482, B: 0.443, A: 1}
	Chablis                        = Color{R: 1.000, G: 0.957, B: 0.953, A: 1}
	Himalaya                       = Color{R: 0.416, G: 0.365, B: 0.106, A: 1}
	InternationalOrange            = Color{R: 1.000, G: 0.310, B: 0.000, A: 1}
	MeatBrown                      = Color{R: 0.898, G: 0.718, B: 0.231, A: 1}
	Jumbo                          = Color{R: 0.486, G: 0.482, B: 0.510, A: 1}
	Smoke                          = Color{R: 0.451, G: 0.510, B: 0.463, A: 1}
	SpicyMix                       = Color{R: 0.545, G: 0.373, B: 0.302, A: 1}
	Ao                             = Color{R: 0.000, G: 0.502, B: 0.000, A: 1}
	Chardonnay                     = Color{R: 1.000, G: 0.804, B: 0.549, A: 1}
	Diesel                         = Color{R: 0.075, G: 0.000, B: 0.000, A: 1}
	DukeBlue                       = Color{R: 0.000, G: 0.000, B: 0.612, A: 1}
	Fandango                       = Color{R: 0.710, G: 0.200, B: 0.537, A: 1}
	VistaBlue                      = Color{R: 0.486, G: 0.620, B: 0.851, A: 1}
	CrayolaYellow                  = Color{R: 0.988, G: 0.910, B: 0.514, A: 1}
	Finlandia                      = Color{R: 0.333, G: 0.427, B: 0.337, A: 1}
	Kumera                         = Color{R: 0.533, G: 0.384, B: 0.129, A: 1}
	Mojo                           = Color{R: 0.753, G: 0.278, B: 0.216, A: 1}
	Mulberry                       = Color{R: 0.773, G: 0.294, B: 0.549, A: 1}
	CreamCan                       = Color{R: 0.961, G: 0.784, B: 0.361, A: 1}
	DullLavender                   = Color{R: 0.659, G: 0.600, B: 0.902, A: 1}
	RoseVale                       = Color{R: 0.671, G: 0.306, B: 0.322, A: 1}
	DarkPurple                     = Color{R: 0.188, G: 0.098, B: 0.204, A: 1}
	FlamePea                       = Color{R: 0.855, G: 0.357, B: 0.220, A: 1}
	PaleRose                       = Color{R: 1.000, G: 0.882, B: 0.949, A: 1}
	Celtic                         = Color{R: 0.086, G: 0.196, B: 0.133, A: 1}
	Chiffon                        = Color{R: 0.945, G: 1.000, B: 0.784, A: 1}
	EasternBlue                    = Color{R: 0.118, G: 0.604, B: 0.690, A: 1}
	RiceCake                       = Color{R: 1.000, G: 0.996, B: 0.941, A: 1}
	CrayolaGreen                   = Color{R: 0.110, G: 0.675, B: 0.471, A: 1}
	LemonMeringue                  = Color{R: 0.965, G: 0.918, B: 0.745, A: 1}
	PalePlum                       = Color{R: 0.867, G: 0.627, B: 0.867, A: 1}
	Smitten                        = Color{R: 0.784, G: 0.255, B: 0.525, A: 1}
	VisVis                         = Color{R: 1.000, G: 0.937, B: 0.631, A: 1}
	VistaWhite                     = Color{R: 0.988, G: 0.973, B: 0.969, A: 1}
	DarkCandyAppleRed              = Color{R: 0.643, G: 0.000, B: 0.000, A: 1}
	Oasis                          = Color{R: 0.996, G: 0.937, B: 0.808, A: 1}
	RacingGreen                    = Color{R: 0.047, G: 0.098, B: 0.067, A: 1}
	RussianViolet                  = Color{R: 0.196, G: 0.090, B: 0.302, A: 1}
	SpanishOrange                  = Color{R: 0.910, G: 0.380, B: 0.000, A: 1}
	Thatch                         = Color{R: 0.714, G: 0.616, B: 0.596, A: 1}
	AmericanRose                   = Color{R: 1.000, G: 0.012, B: 0.243, A: 1}
	PastelBrown                    = Color{R: 0.514, G: 0.412, B: 0.325, A: 1}
	Raven                          = Color{R: 0.447, G: 0.482, B: 0.537, A: 1}
	SilverChalice                  = Color{R: 0.675, G: 0.675, B: 0.675, A: 1}
	Strikemaster                   = Color{R: 0.584, G: 0.388, B: 0.529, A: 1}
	VenetianRed                    = Color{R: 0.784, G: 0.031, B: 0.082, A: 1}
	Anzac                          = Color{R: 0.878, G: 0.714, B: 0.275, A: 1}
	Diamond                        = Color{R: 0.725, G: 0.949, B: 1.000, A: 1}
	Goldenrod                      = Color{R: 0.855, G: 0.647, B: 0.125, A: 1}
	Porsche                        = Color{R: 0.918, G: 0.682, B: 0.412, A: 1}
	RiceFlower                     = Color{R: 0.933, G: 1.000, B: 0.886, A: 1}
	SpanishViridian                = Color{R: 0.000, G: 0.498, B: 0.361, A: 1}
	BrownRust                      = Color{R: 0.686, G: 0.349, B: 0.243, A: 1}
	Jacarta                        = Color{R: 0.227, G: 0.165, B: 0.416, A: 1}
	Ming                           = Color{R: 0.212, G: 0.455, B: 0.490, A: 1}
	OldSilver                      = Color{R: 0.518, G: 0.518, B: 0.510, A: 1}
	SatinLinen                     = Color{R: 0.902, G: 0.894, B: 0.831, A: 1}
	Fawn                           = Color{R: 0.898, G: 0.667, B: 0.439, A: 1}
	Grizzly                        = Color{R: 0.533, G: 0.345, B: 0.094, A: 1}
	PottersClay                    = Color{R: 0.549, G: 0.341, B: 0.220, A: 1}
	Sunray                         = Color{R: 0.890, G: 0.671, B: 0.341, A: 1}
	BonJour                        = Color{R: 0.898, G: 0.878, B: 0.882, A: 1}
	Coral                          = Color{R: 1.000, G: 0.498, B: 0.314, A: 1}
	SlateBlue                      = Color{R: 0.416, G: 0.353, B: 0.804, A: 1}
	YellowMetal                    = Color{R: 0.443, G: 0.388, B: 0.220, A: 1}
	Grullo                         = Color{R: 0.663, G: 0.604, B: 0.525, A: 1}
	JapaneseMaple                  = Color{R: 0.471, G: 0.004, B: 0.035, A: 1}
	JungleMist                     = Color{R: 0.706, G: 0.812, B: 0.827, A: 1}
	Martini                        = Color{R: 0.686, G: 0.627, B: 0.620, A: 1}
	Monza                          = Color{R: 0.780, G: 0.012, B: 0.118, A: 1}
	SurfieGreen                    = Color{R: 0.047, G: 0.478, B: 0.475, A: 1}
	SpiroDiscoBall                 = Color{R: 0.059, G: 0.753, B: 0.988, A: 1}
	BritishRacingGreen             = Color{R: 0.000, G: 0.259, B: 0.145, A: 1}
	CedarWoodFinish                = Color{R: 0.443, G: 0.102, B: 0.000, A: 1}
	DeepRuby                       = Color{R: 0.518, G: 0.247, B: 0.357, A: 1}
	FunGreen                       = Color{R: 0.004, G: 0.427, B: 0.224, A: 1}
	ParisWhite                     = Color{R: 0.792, G: 0.863, B: 0.831, A: 1}
	Givry                          = Color{R: 0.973, G: 0.894, B: 0.749, A: 1}
	GoldFusion                     = Color{R: 0.522, G: 0.459, B: 0.306, A: 1}
	HawkesBlue                     = Color{R: 0.831, G: 0.886, B: 0.988, A: 1}
	AustralianMint                 = Color{R: 0.961, G: 1.000, B: 0.745, A: 1}
	Azalea                         = Color{R: 0.969, G: 0.784, B: 0.855, A: 1}
	BittersweetShimmer             = Color{R: 0.749, G: 0.310, B: 0.318, A: 1}
	Chino                          = Color{R: 0.808, G: 0.780, B: 0.655, A: 1}
	Chocolate                      = Color{R: 0.482, G: 0.247, B: 0.000, A: 1}
	YellowOrange                   = Color{R: 1.000, G: 0.682, B: 0.259, A: 1}
	Lonestar                       = Color{R: 0.427, G: 0.004, B: 0.004, A: 1}
	Periwinkle                     = Color{R: 0.800, G: 0.800, B: 1.000, A: 1}
	PhthaloGreen                   = Color{R: 0.071, G: 0.208, B: 0.141, A: 1}
	RoseEbony                      = Color{R: 0.404, G: 0.282, B: 0.275, A: 1}
	SteelGray                      = Color{R: 0.149, G: 0.137, B: 0.208, A: 1}
	Ginger                         = Color{R: 0.690, G: 0.396, B: 0.000, A: 1}
	Licorice                       = Color{R: 0.102, G: 0.067, B: 0.063, A: 1}
	MajorelleBlue                  = Color{R: 0.376, G: 0.314, B: 0.863, A: 1}
	StPatricksBlue                 = Color{R: 0.137, G: 0.161, B: 0.478, A: 1}
	WineBerry                      = Color{R: 0.349, G: 0.114, B: 0.208, A: 1}
	WinterSky                      = Color{R: 1.000, G: 0.000, B: 0.486, A: 1}
	BabyPowder                     = Color{R: 0.996, G: 0.996, B: 0.980, A: 1}
	Blueberry                      = Color{R: 0.310, G: 0.525, B: 0.969, A: 1}
	Gossamer                       = Color{R: 0.024, G: 0.608, B: 0.506, A: 1}
	Melanzane                      = Color{R: 0.188, G: 0.020, B: 0.161, A: 1}
	VeryLightTangelo               = Color{R: 1.000, G: 0.690, B: 0.467, A: 1}
	LincolnGreen                   = Color{R: 0.098, G: 0.349, B: 0.020, A: 1}
	VeryPaleOrange                 = Color{R: 1.000, G: 0.875, B: 0.749, A: 1}
	Anakiwa                        = Color{R: 0.616, G: 0.898, B: 1.000, A: 1}
	FieldDrab                      = Color{R: 0.424, G: 0.329, B: 0.118, A: 1}
	Nevada                         = Color{R: 0.392, G: 0.431, B: 0.459, A: 1}
	Orchid                         = Color{R: 0.855, G: 0.439, B: 0.839, A: 1}
	PaleCopper                     = Color{R: 0.855, G: 0.541, B: 0.404, A: 1}
	SantasGray                     = Color{R: 0.624, G: 0.627, B: 0.694, A: 1}
	Skeptic                        = Color{R: 0.792, G: 0.902, B: 0.855, A: 1}
	VividCerise                    = Color{R: 0.855, G: 0.114, B: 0.506, A: 1}
	Blossom                        = Color{R: 0.863, G: 0.706, B: 0.737, A: 1}
	FrenchSkyBlue                  = Color{R: 0.467, G: 0.710, B: 0.996, A: 1}
	LaurelGreen                    = Color{R: 0.663, G: 0.729, B: 0.616, A: 1}
	LuxorGold                      = Color{R: 0.655, G: 0.533, B: 0.173, A: 1}
	NadeshikoPink                  = Color{R: 0.965, G: 0.678, B: 0.776, A: 1}
	BaliHai                        = Color{R: 0.522, G: 0.624, B: 0.686, A: 1}
	BlueSmoke                      = Color{R: 0.455, G: 0.533, B: 0.506, A: 1}
	Contessa                       = Color{R: 0.776, G: 0.447, B: 0.420, A: 1}
	GreenSheen                     = Color{R: 0.431, G: 0.682, B: 0.631, A: 1}
	BrickRed                       = Color{R: 0.796, G: 0.255, B: 0.329, A: 1}
	DeepCarminePink                = Color{R: 0.937, G: 0.188, B: 0.220, A: 1}
	Loafer                         = Color{R: 0.933, G: 0.957, B: 0.871, A: 1}
	PinkPearl                      = Color{R: 0.906, G: 0.675, B: 0.812, A: 1}
	Putty                          = Color{R: 0.906, G: 0.804, B: 0.549, A: 1}
	ReefGold                       = Color{R: 0.624, G: 0.510, B: 0.110, A: 1}
	Galliano                       = Color{R: 0.863, G: 0.698, B: 0.047, A: 1}
	Mandalay                       = Color{R: 0.678, G: 0.471, B: 0.106, A: 1}
	MintGreen                      = Color{R: 0.596, G: 1.000, B: 0.596, A: 1}
	OceanGreen                     = Color{R: 0.282, G: 0.749, B: 0.569, A: 1}
	OrangeRoughy                   = Color{R: 0.769, G: 0.341, B: 0.098, A: 1}
	Vulcan                         = Color{R: 0.063, G: 0.071, B: 0.114, A: 1}
	ApricotWhite                   = Color{R: 1.000, G: 0.996, B: 0.925, A: 1}
	Arrowtown                      = Color{R: 0.580, G: 0.529, B: 0.443, A: 1}
	HalfColonialWhite              = Color{R: 0.992, G: 0.965, B: 0.827, A: 1}
	NeonGreen                      = Color{R: 0.224, G: 1.000, B: 0.078, A: 1}
	RYBYellow                      = Color{R: 0.996, G: 0.996, B: 0.200, A: 1}
	DarkOrange                     = Color{R: 1.000, G: 0.549, B: 0.000, A: 1}
	Denim                          = Color{R: 0.082, G: 0.376, B: 0.741, A: 1}
	Siam                           = Color{R: 0.392, G: 0.416, B: 0.329, A: 1}
	AlabamaCrimson                 = Color{R: 0.686, G: 0.000, B: 0.165, A: 1}
	Corduroy                       = Color{R: 0.376, G: 0.431, B: 0.408, A: 1}
	Cyprus                         = Color{R: 0.000, G: 0.243, B: 0.251, A: 1}
	JudgeGray                      = Color{R: 0.329, G: 0.263, B: 0.200, A: 1}
	RosePink                       = Color{R: 1.000, G: 0.400, B: 0.800, A: 1}
	OliveDrab                      = Color{R: 0.420, G: 0.557, B: 0.137, A: 1}
	PurpleHeart                    = Color{R: 0.412, G: 0.208, B: 0.612, A: 1}
	BrightNavyBlue                 = Color{R: 0.098, G: 0.455, B: 0.824, A: 1}
	California                     = Color{R: 0.996, G: 0.616, B: 0.016, A: 1}
	SandyBeach                     = Color{R: 1.000, G: 0.918, B: 0.784, A: 1}
	Camelot                        = Color{R: 0.537, G: 0.204, B: 0.337, A: 1}
	CarolinaBlue                   = Color{R: 0.337, G: 0.627, B: 0.827, A: 1}
	Dew                            = Color{R: 0.918, G: 1.000, B: 0.996, A: 1}
	HonoluluBlue                   = Color{R: 0.000, G: 0.427, B: 0.690, A: 1}
	Iceberg                        = Color{R: 0.443, G: 0.651, B: 0.824, A: 1}
	OgreOdor                       = Color{R: 0.992, G: 0.322, B: 0.251, A: 1}
	Olivine                        = Color{R: 0.604, G: 0.725, B: 0.451, A: 1}
	WoodyBrown                     = Color{R: 0.282, G: 0.192, B: 0.192, A: 1}
	Alto                           = Color{R: 0.859, G: 0.859, B: 0.859, A: 1}
	BeauBlue                       = Color{R: 0.737, G: 0.831, B: 0.902, A: 1}
	DarkLiver                      = Color{R: 0.325, G: 0.294, B: 0.310, A: 1}
	Everglade                      = Color{R: 0.110, G: 0.251, B: 0.180, A: 1}
	NeonFuchsia                    = Color{R: 0.996, G: 0.255, B: 0.392, A: 1}
	Nobel                          = Color{R: 0.718, G: 0.694, B: 0.694, A: 1}
	CelestialBlue                  = Color{R: 0.286, G: 0.592, B: 0.816, A: 1}
	Coffee                         = Color{R: 0.435, G: 0.306, B: 0.216, A: 1}
	Cumin                          = Color{R: 0.573, G: 0.263, B: 0.129, A: 1}
	DarkGunmetal                   = Color{R: 0.122, G: 0.149, B: 0.165, A: 1}
	DeepSapphire                   = Color{R: 0.031, G: 0.145, B: 0.404, A: 1}
	LavenderGray                   = Color{R: 0.769, G: 0.765, B: 0.816, A: 1}
	Seance                         = Color{R: 0.451, G: 0.118, B: 0.561, A: 1}
	Thistle                        = Color{R: 0.847, G: 0.749, B: 0.847, A: 1}
	Bistre                         = Color{R: 0.239, G: 0.169, B: 0.122, A: 1}
	Christalle                     = Color{R: 0.200, G: 0.012, B: 0.420, A: 1}
	EnglishVermillion              = Color{R: 0.800, G: 0.278, B: 0.294, A: 1}
	Husk                           = Color{R: 0.718, G: 0.643, B: 0.345, A: 1}
	CoffeeBean                     = Color{R: 0.165, G: 0.078, B: 0.055, A: 1}
	TropicalBlue                   = Color{R: 0.765, G: 0.867, B: 0.976, A: 1}
	Zinnwaldite                    = Color{R: 0.922, G: 0.761, B: 0.686, A: 1}
	PaleOyster                     = Color{R: 0.596, G: 0.553, B: 0.467, A: 1}
	Twine                          = Color{R: 0.761, G: 0.584, B: 0.365, A: 1}
	Affair                         = Color{R: 0.443, G: 0.275, B: 0.576, A: 1}
	BayofMany                      = Color{R: 0.153, G: 0.227, B: 0.506, A: 1}
	DoubleSpanishWhite             = Color{R: 0.902, G: 0.843, B: 0.725, A: 1}
	LemonLime                      = Color{R: 0.890, G: 1.000, B: 0.000, A: 1}
	Lenurple                       = Color{R: 0.729, G: 0.576, B: 0.847, A: 1}
	BlueDianne                     = Color{R: 0.125, G: 0.282, B: 0.322, A: 1}
	Cloud                          = Color{R: 0.780, G: 0.769, B: 0.749, A: 1}
	Pizza                          = Color{R: 0.788, G: 0.580, B: 0.082, A: 1}
	BattleshipGray                 = Color{R: 0.510, G: 0.561, B: 0.447, A: 1}
	MonaLisa                       = Color{R: 1.000, G: 0.631, B: 0.580, A: 1}
	SpartanCrimson                 = Color{R: 0.620, G: 0.075, B: 0.086, A: 1}
	BirdFlower                     = Color{R: 0.831, G: 0.804, B: 0.086, A: 1}
	Bossanova                      = Color{R: 0.306, G: 0.165, B: 0.353, A: 1}
	CapeCod                        = Color{R: 0.235, G: 0.267, B: 0.263, A: 1}
	GoldenDream                    = Color{R: 0.941, G: 0.835, B: 0.176, A: 1}
	Rosewood                       = Color{R: 0.396, G: 0.000, B: 0.043, A: 1}
	Paprika                        = Color{R: 0.553, G: 0.008, B: 0.149, A: 1}
	PinkLace                       = Color{R: 1.000, G: 0.867, B: 0.957, A: 1}
	Japonica                       = Color{R: 0.847, G: 0.486, B: 0.388, A: 1}
	Kournikova                     = Color{R: 1.000, G: 0.906, B: 0.447, A: 1}
	Pancho                         = Color{R: 0.929, G: 0.804, B: 0.671, A: 1}
	ResolutionBlue                 = Color{R: 0.000, G: 0.137, B: 0.529, A: 1}
	RioGrande                      = Color{R: 0.733, G: 0.816, B: 0.035, A: 1}
	WildWatermelon                 = Color{R: 0.988, G: 0.424, B: 0.522, A: 1}
	Cioccolato                     = Color{R: 0.333, G: 0.157, B: 0.047, A: 1}
	DarkVanilla                    = Color{R: 0.820, G: 0.745, B: 0.659, A: 1}
	Geebung                        = Color{R: 0.820, G: 0.561, B: 0.106, A: 1}
	JordyBlue                      = Color{R: 0.541, G: 0.725, B: 0.945, A: 1}
	Kilamanjaro                    = Color{R: 0.141, G: 0.047, B: 0.008, A: 1}
	SpanishGray                    = Color{R: 0.596, G: 0.596, B: 0.596, A: 1}
	PacificBlue                    = Color{R: 0.110, G: 0.663, B: 0.788, A: 1}
	PetiteOrchid                   = Color{R: 0.859, G: 0.588, B: 0.565, A: 1}
	PuertoRico                     = Color{R: 0.247, G: 0.757, B: 0.667, A: 1}
	Rhino                          = Color{R: 0.180, G: 0.247, B: 0.384, A: 1}
	SolidPink                      = Color{R: 0.537, G: 0.220, B: 0.263, A: 1}
	Gunmetal                       = Color{R: 0.165, G: 0.204, B: 0.224, A: 1}
	Mandy                          = Color{R: 0.886, G: 0.329, B: 0.396, A: 1}
	Mantle                         = Color{R: 0.545, G: 0.612, B: 0.565, A: 1}
	AppleGreen                     = Color{R: 0.553, G: 0.714, B: 0.000, A: 1}
	BlackRose                      = Color{R: 0.404, G: 0.012, B: 0.176, A: 1}
	Edgewater                      = Color{R: 0.784, G: 0.890, B: 0.843, A: 1}
	ElectricBlue                   = Color{R: 0.490, G: 0.976, B: 1.000, A: 1}
	Frangipani                     = Color{R: 1.000, G: 0.871, B: 0.702, A: 1}
	MonteCarlo                     = Color{R: 0.514, G: 0.816, B: 0.776, A: 1}
	Westar                         = Color{R: 0.863, G: 0.851, B: 0.824, A: 1}
	FireBush                       = Color{R: 0.910, G: 0.600, B: 0.157, A: 1}
	PastelGreen                    = Color{R: 0.467, G: 0.867, B: 0.467, A: 1}
	RoseWhite                      = Color{R: 1.000, G: 0.965, B: 0.961, A: 1}
	SaffronMango                   = Color{R: 0.976, G: 0.749, B: 0.345, A: 1}
	ToryBlue                       = Color{R: 0.078, G: 0.314, B: 0.667, A: 1}
	LightGoldenrodYellow           = Color{R: 0.980, G: 0.980, B: 0.824, A: 1}
	SacramentoStateGreen           = Color{R: 0.016, G: 0.224, B: 0.153, A: 1}
	UCLABlue                       = Color{R: 0.325, G: 0.408, B: 0.584, A: 1}
	YellowRose                     = Color{R: 1.000, G: 0.941, B: 0.000, A: 1}
	BlueGray                       = Color{R: 0.400, G: 0.600, B: 0.800, A: 1}
	Capri                          = Color{R: 0.000, G: 0.749, B: 1.000, A: 1}
	TyrianPurple                   = Color{R: 0.400, G: 0.008, B: 0.235, A: 1}
	Zaffre                         = Color{R: 0.000, G: 0.078, B: 0.659, A: 1}
	Cloudy                         = Color{R: 0.675, G: 0.647, B: 0.624, A: 1}
	FandangoPink                   = Color{R: 0.871, G: 0.322, B: 0.522, A: 1}
	GoldTips                       = Color{R: 0.871, G: 0.729, B: 0.075, A: 1}
	Soapstone                      = Color{R: 1.000, G: 0.984, B: 0.976, A: 1}
	TangoPink                      = Color{R: 0.894, G: 0.443, B: 0.478, A: 1}
	Dallas                         = Color{R: 0.431, G: 0.294, B: 0.149, A: 1}
	Quincy                         = Color{R: 0.384, G: 0.247, B: 0.176, A: 1}
	Tomato                         = Color{R: 1.000, G: 0.388, B: 0.278, A: 1}
	HunterGreen                    = Color{R: 0.208, G: 0.369, B: 0.231, A: 1}
	Kangaroo                       = Color{R: 0.776, G: 0.784, B: 0.741, A: 1}
	Manhattan                      = Color{R: 0.961, G: 0.788, B: 0.600, A: 1}
	Bastille                       = Color{R: 0.161, G: 0.129, B: 0.188, A: 1}
	CinnamonSatin                  = Color{R: 0.804, G: 0.376, B: 0.494, A: 1}
	DeepFuchsia                    = Color{R: 0.757, G: 0.329, B: 0.757, A: 1}
	Heath                          = Color{R: 0.329, G: 0.063, B: 0.071, A: 1}
	Hemlock                        = Color{R: 0.369, G: 0.365, B: 0.231, A: 1}
	ShamrockGreen                  = Color{R: 0.000, G: 0.620, B: 0.376, A: 1}
	StarkWhite                     = Color{R: 0.898, G: 0.843, B: 0.741, A: 1}
	Sunflower                      = Color{R: 0.894, G: 0.831, B: 0.133, A: 1}
	BleachWhite                    = Color{R: 0.996, G: 0.953, B: 0.847, A: 1}
	BurningOrange                  = Color{R: 1.000, G: 0.439, B: 0.204, A: 1}
	Mariner                        = Color{R: 0.157, G: 0.416, B: 0.804, A: 1}
	MikadoYellow                   = Color{R: 1.000, G: 0.769, B: 0.047, A: 1}
	RollingStone                   = Color{R: 0.455, G: 0.490, B: 0.514, A: 1}
	Apricot                        = Color{R: 0.984, G: 0.808, B: 0.694, A: 1}
	FrenchRaspberry                = Color{R: 0.780, G: 0.173, B: 0.282, A: 1}
	LinkWater                      = Color{R: 0.851, G: 0.894, B: 0.961, A: 1}
	Aureolin                       = Color{R: 0.992, G: 0.933, B: 0.000, A: 1}
	HeliotropeGray                 = Color{R: 0.667, G: 0.596, B: 0.663, A: 1}
	MediumBlue                     = Color{R: 0.000, G: 0.000, B: 0.804, A: 1}
	PowderAsh                      = Color{R: 0.737, G: 0.788, B: 0.761, A: 1}
	WellRead                       = Color{R: 0.706, G: 0.200, B: 0.196, A: 1}
	CyanAzure                      = Color{R: 0.306, G: 0.510, B: 0.706, A: 1}
	MuleFawn                       = Color{R: 0.549, G: 0.278, B: 0.184, A: 1}
	Nebula                         = Color{R: 0.796, G: 0.859, B: 0.839, A: 1}
	PaleGold                       = Color{R: 0.902, G: 0.745, B: 0.541, A: 1}
	SpanishBistre                  = Color{R: 0.502, G: 0.459, B: 0.196, A: 1}
	CountyGreen                    = Color{R: 0.004, G: 0.216, B: 0.102, A: 1}
	MediumOrchid                   = Color{R: 0.729, G: 0.333, B: 0.827, A: 1}
	NCSRed                         = Color{R: 0.769, G: 0.008, B: 0.200, A: 1}
	Pelorous                       = Color{R: 0.243, G: 0.671, B: 0.749, A: 1}
	Pipi                           = Color{R: 0.996, G: 0.957, B: 0.800, A: 1}
	CyanBlueAzure                  = Color{R: 0.275, G: 0.510, B: 0.749, A: 1}
	Iroko                          = Color{R: 0.263, G: 0.192, B: 0.125, A: 1}
	LightCyan                      = Color{R: 0.878, G: 1.000, B: 1.000, A: 1}
	ChinaIvory                     = Color{R: 0.988, G: 1.000, B: 0.906, A: 1}
	Conch                          = Color{R: 0.788, G: 0.851, B: 0.824, A: 1}
	DeepGreenCyanTurquoise         = Color{R: 0.055, G: 0.486, B: 0.380, A: 1}
	MyPink                         = Color{R: 0.839, G: 0.569, B: 0.533, A: 1}
	SpanishBlue                    = Color{R: 0.000, G: 0.439, B: 0.722, A: 1}
	RichBlack                      = Color{R: 0.000, G: 0.251, B: 0.251, A: 1}
	Rufous                         = Color{R: 0.659, G: 0.110, B: 0.027, A: 1}
	Sushi                          = Color{R: 0.529, G: 0.671, B: 0.224, A: 1}
	AeroBlue                       = Color{R: 0.788, G: 1.000, B: 0.898, A: 1}
	Bunker                         = Color{R: 0.051, G: 0.067, B: 0.090, A: 1}
	ColonialWhite                  = Color{R: 1.000, G: 0.929, B: 0.737, A: 1}
	Lola                           = Color{R: 0.875, G: 0.812, B: 0.859, A: 1}
	PantoneYellow                  = Color{R: 0.996, G: 0.875, B: 0.000, A: 1}
	Thunderbird                    = Color{R: 0.753, G: 0.169, B: 0.094, A: 1}
	TrendyGreen                    = Color{R: 0.486, G: 0.533, B: 0.102, A: 1}
	UtahCrimson                    = Color{R: 0.827, G: 0.000, B: 0.247, A: 1}
	Asphalt                        = Color{R: 0.075, G: 0.039, B: 0.024, A: 1}
	CrownofThorns                  = Color{R: 0.467, G: 0.122, B: 0.122, A: 1}
	GuardsmanRed                   = Color{R: 0.729, G: 0.004, B: 0.004, A: 1}
	Mako                           = Color{R: 0.267, G: 0.286, B: 0.329, A: 1}
	Surf                           = Color{R: 0.733, G: 0.843, B: 0.757, A: 1}
	ClassicRose                    = Color{R: 0.984, G: 0.800, B: 0.906, A: 1}
	ParisDaisy                     = Color{R: 1.000, G: 0.957, B: 0.431, A: 1}
	RichElectricBlue               = Color{R: 0.031, G: 0.573, B: 0.816, A: 1}
	Tidal                          = Color{R: 0.945, G: 1.000, B: 0.678, A: 1}
	Lipstick                       = Color{R: 0.671, G: 0.020, B: 0.388, A: 1}
	MidnightBlue                   = Color{R: 0.098, G: 0.098, B: 0.439, A: 1}
	Quicksand                      = Color{R: 0.741, G: 0.592, B: 0.557, A: 1}
	RubyRed                        = Color{R: 0.608, G: 0.067, B: 0.118, A: 1}
	DarkOrchid                     = Color{R: 0.600, G: 0.196, B: 0.800, A: 1}
	GrayNickel                     = Color{R: 0.765, G: 0.765, B: 0.741, A: 1}
	Jet                            = Color{R: 0.204, G: 0.204, B: 0.204, A: 1}
	MuddyWaters                    = Color{R: 0.718, G: 0.557, B: 0.361, A: 1}
	RuddyPink                      = Color{R: 0.882, G: 0.557, B: 0.588, A: 1}
	CaputMortuum                   = Color{R: 0.349, G: 0.153, B: 0.125, A: 1}
	GoldenBrown                    = Color{R: 0.600, G: 0.396, B: 0.082, A: 1}
	Hibiscus                       = Color{R: 0.714, G: 0.192, B: 0.424, A: 1}
	RoseBudCherry                  = Color{R: 0.502, G: 0.043, B: 0.278, A: 1}
	Tamarillo                      = Color{R: 0.600, G: 0.086, B: 0.075, A: 1}
	Cashmere                       = Color{R: 0.902, G: 0.745, B: 0.647, A: 1}
	GoldenGlow                     = Color{R: 0.992, G: 0.886, B: 0.584, A: 1}
	SalmonPink                     = Color{R: 1.000, G: 0.569, B: 0.643, A: 1}
	TallPoppy                      = Color{R: 0.702, G: 0.176, B: 0.161, A: 1}
	Urobilin                       = Color{R: 0.882, G: 0.678, B: 0.129, A: 1}
	Kobicha                        = Color{R: 0.420, G: 0.267, B: 0.137, A: 1}
	Negroni                        = Color{R: 1.000, G: 0.886, B: 0.773, A: 1}
	Panache                        = Color{R: 0.918, G: 0.965, B: 0.933, A: 1}
	Blond                          = Color{R: 0.980, G: 0.941, B: 0.745, A: 1}
	ColdTurkey                     = Color{R: 0.808, G: 0.729, B: 0.729, A: 1}
	Crimson                        = Color{R: 0.863, G: 0.078, B: 0.235, A: 1}
	FernGreen                      = Color{R: 0.310, G: 0.475, B: 0.259, A: 1}
	FuzzyWuzzyBrown                = Color{R: 0.769, G: 0.337, B: 0.333, A: 1}
	RipePlum                       = Color{R: 0.255, G: 0.000, B: 0.337, A: 1}
	WildStrawberry                 = Color{R: 1.000, G: 0.263, B: 0.643, A: 1}
	MagicPotion                    = Color{R: 1.000, G: 0.267, B: 0.400, A: 1}
	PoloBlue                       = Color{R: 0.553, G: 0.659, B: 0.800, A: 1}
	Chantilly                      = Color{R: 0.973, G: 0.765, B: 0.875, A: 1}
	Cherokee                       = Color{R: 0.988, G: 0.855, B: 0.596, A: 1}
	Clinker                        = Color{R: 0.216, G: 0.114, B: 0.035, A: 1}
	Janna                          = Color{R: 0.957, G: 0.922, B: 0.827, A: 1}
	LimedSpruce                    = Color{R: 0.224, G: 0.282, B: 0.318, A: 1}
	Umber                          = Color{R: 0.388, G: 0.318, B: 0.278, A: 1}
	BabyBlue                       = Color{R: 0.537, G: 0.812, B: 0.941, A: 1}
	ImperialRed                    = Color{R: 0.929, G: 0.161, B: 0.224, A: 1}
	NCSBlue                        = Color{R: 0.000, G: 0.529, B: 0.741, A: 1}
	PastelPink                     = Color{R: 0.871, G: 0.647, B: 0.643, A: 1}
	RoofTerracotta                 = Color{R: 0.651, G: 0.184, B: 0.125, A: 1}
	ArylideYellow                  = Color{R: 0.914, G: 0.839, B: 0.420, A: 1}
	CameoPink                      = Color{R: 0.937, G: 0.733, B: 0.800, A: 1}
	Mimosa                         = Color{R: 0.973, G: 0.992, B: 0.827, A: 1}
	Varden                         = Color{R: 1.000, G: 0.965, B: 0.875, A: 1}
	HotToddy                       = Color{R: 0.702, G: 0.502, B: 0.027, A: 1}
	LavenderRose                   = Color{R: 0.984, G: 0.627, B: 0.890, A: 1}
	Pampas                         = Color{R: 0.957, G: 0.949, B: 0.933, A: 1}
	SweetPink                      = Color{R: 0.992, G: 0.624, B: 0.635, A: 1}
	WinterHazel                    = Color{R: 0.835, G: 0.820, B: 0.584, A: 1}
	Dirt                           = Color{R: 0.608, G: 0.463, B: 0.325, A: 1}
	Horses                         = Color{R: 0.329, G: 0.239, B: 0.216, A: 1}
	RenoSand                       = Color{R: 0.659, G: 0.396, B: 0.082, A: 1}
	SizzlingSunrise                = Color{R: 1.000, G: 0.859, B: 0.000, A: 1}
	Zumthor                        = Color{R: 0.929, G: 0.965, B: 1.000, A: 1}
	ToreaBay                       = Color{R: 0.059, G: 0.176, B: 0.620, A: 1}
	TrueV                          = Color{R: 0.541, G: 0.451, B: 0.839, A: 1}
	BrightTurquoise                = Color{R: 0.031, G: 0.910, B: 0.871, A: 1}
	LightDeepPink                  = Color{R: 1.000, G: 0.361, B: 0.804, A: 1}
	Madison                        = Color{R: 0.035, G: 0.145, B: 0.365, A: 1}
	PaleLeaf                       = Color{R: 0.753, G: 0.827, B: 0.725, A: 1}
	Sunglow                        = Color{R: 1.000, G: 0.800, B: 0.200, A: 1}
	PullmanBrown                   = Color{R: 0.392, G: 0.255, B: 0.090, A: 1}
	UltraPink                      = Color{R: 1.000, G: 0.435, B: 1.000, A: 1}
	Buff                           = Color{R: 0.941, G: 0.863, B: 0.510, A: 1}
	Froly                          = Color{R: 0.961, G: 0.459, B: 0.518, A: 1}
	JapaneseLaurel                 = Color{R: 0.039, G: 0.412, B: 0.024, A: 1}
	Karaka                         = Color{R: 0.118, G: 0.086, B: 0.035, A: 1}
	Melon                          = Color{R: 0.992, G: 0.737, B: 0.706, A: 1}
	BulgarianRose                  = Color{R: 0.282, G: 0.024, B: 0.027, A: 1}
	Cararra                        = Color{R: 0.933, G: 0.933, B: 0.910, A: 1}
	Lily                           = Color{R: 0.784, G: 0.667, B: 0.749, A: 1}
	Equator                        = Color{R: 0.882, G: 0.737, B: 0.392, A: 1}
	Popstar                        = Color{R: 0.745, G: 0.310, B: 0.384, A: 1}
	SlimyGreen                     = Color{R: 0.161, G: 0.588, B: 0.090, A: 1}
	VinRouge                       = Color{R: 0.596, G: 0.239, B: 0.380, A: 1}
	WoodBark                       = Color{R: 0.149, G: 0.067, B: 0.020, A: 1}
	ProvincialPink                 = Color{R: 0.996, G: 0.961, B: 0.945, A: 1}
	SnowyMint                      = Color{R: 0.839, G: 1.000, B: 0.859, A: 1}
	Bone                           = Color{R: 0.890, G: 0.855, B: 0.788, A: 1}
	Chatelle                       = Color{R: 0.741, G: 0.702, B: 0.780, A: 1}
	Mustard                        = Color{R: 1.000, G: 0.859, B: 0.345, A: 1}
	PewterBlue                     = Color{R: 0.545, G: 0.659, B: 0.718, A: 1}
	PortlandOrange                 = Color{R: 1.000, G: 0.353, B: 0.212, A: 1}
	AthensGray                     = Color{R: 0.933, G: 0.941, B: 0.953, A: 1}
	RodeoDust                      = Color{R: 0.788, G: 0.698, B: 0.608, A: 1}
	CrayolaRed                     = Color{R: 0.933, G: 0.125, B: 0.302, A: 1}
	ElectricPurple                 = Color{R: 0.749, G: 0.000, B: 1.000, A: 1}
	HalfDutchWhite                 = Color{R: 0.996, G: 0.969, B: 0.871, A: 1}
	Supernova                      = Color{R: 1.000, G: 0.788, B: 0.004, A: 1}
	VeryLightMalachiteGreen        = Color{R: 0.392, G: 0.914, B: 0.525, A: 1}
	CloudBurst                     = Color{R: 0.125, G: 0.180, B: 0.329, A: 1}
	Festival                       = Color{R: 0.984, G: 0.914, B: 0.424, A: 1}
	Liberty                        = Color{R: 0.329, G: 0.353, B: 0.655, A: 1}
	Lochinvar                      = Color{R: 0.173, G: 0.549, B: 0.518, A: 1}
	Norway                         = Color{R: 0.659, G: 0.741, B: 0.624, A: 1}
	Patina                         = Color{R: 0.388, G: 0.604, B: 0.561, A: 1}
	Aubergine                      = Color{R: 0.231, G: 0.035, B: 0.063, A: 1}
	BarnRed                        = Color{R: 0.486, G: 0.039, B: 0.008, A: 1}
	Cosmic                         = Color{R: 0.463, G: 0.224, B: 0.365, A: 1}
	DarkGreen                      = Color{R: 0.004, G: 0.196, B: 0.125, A: 1}
	EggWhite                       = Color{R: 1.000, G: 0.937, B: 0.757, A: 1}
	BlueRomance                    = Color{R: 0.824, G: 0.965, B: 0.871, A: 1}
	Walnut                         = Color{R: 0.467, G: 0.247, B: 0.102, A: 1}
	UCLAGold                       = Color{R: 1.000, G: 0.702, B: 0.000, A: 1}
	BabyBlueEyes                   = Color{R: 0.631, G: 0.792, B: 0.945, A: 1}
	BalticSea                      = Color{R: 0.165, G: 0.149, B: 0.188, A: 1}
	PaleSlate                      = Color{R: 0.765, G: 0.749, B: 0.757, A: 1}
	SafetyYellow                   = Color{R: 0.933, G: 0.824, B: 0.008, A: 1}
	BlanchedAlmond                 = Color{R: 1.000, G: 0.922, B: 0.804, A: 1}
	LunarGreen                     = Color{R: 0.235, G: 0.286, B: 0.227, A: 1}
	AcidGreen                      = Color{R: 0.690, G: 0.749, B: 0.102, A: 1}
	CrimsonGlory                   = Color{R: 0.745, G: 0.000, B: 0.196, A: 1}
	Espresso                       = Color{R: 0.380, G: 0.153, B: 0.094, A: 1}
	HarlequinGreen                 = Color{R: 0.275, G: 0.796, B: 0.094, A: 1}
	AndroidGreen                   = Color{R: 0.643, G: 0.776, B: 0.224, A: 1}
	ButteredRum                    = Color{R: 0.631, G: 0.459, B: 0.051, A: 1}
	LemonGrass                     = Color{R: 0.608, G: 0.620, B: 0.561, A: 1}
	PersianRose                    = Color{R: 0.996, G: 0.157, B: 0.635, A: 1}
	Acapulco                       = Color{R: 0.486, G: 0.690, B: 0.631, A: 1}
	CyberYellow                    = Color{R: 1.000, G: 0.827, B: 0.000, A: 1}
	Smoky                          = Color{R: 0.376, G: 0.357, B: 0.451, A: 1}
	UnderagePink                   = Color{R: 0.976, G: 0.902, B: 0.957, A: 1}
	BisonHide                      = Color{R: 0.757, G: 0.718, B: 0.643, A: 1}
	BlizzardBlue                   = Color{R: 0.639, G: 0.890, B: 0.929, A: 1}
	LondonHue                      = Color{R: 0.745, G: 0.651, B: 0.765, A: 1}
	MagentaPink                    = Color{R: 0.800, G: 0.200, B: 0.545, A: 1}
	Tea                            = Color{R: 0.757, G: 0.729, B: 0.690, A: 1}
	Haiti                          = Color{R: 0.106, G: 0.063, B: 0.208, A: 1}
	Mobster                        = Color{R: 0.498, G: 0.459, B: 0.537, A: 1}
	PinkSherbet                    = Color{R: 0.969, G: 0.561, B: 0.655, A: 1}
	Portafino                      = Color{R: 1.000, G: 1.000, B: 0.706, A: 1}
	Tango                          = Color{R: 0.929, G: 0.478, B: 0.110, A: 1}
	CocoaBean                      = Color{R: 0.282, G: 0.110, B: 0.110, A: 1}
	LightCornflowerBlue            = Color{R: 0.576, G: 0.800, B: 0.918, A: 1}
	MountbattenPink                = Color{R: 0.600, G: 0.478, B: 0.553, A: 1}
	PaleCarmine                    = Color{R: 0.686, G: 0.251, B: 0.208, A: 1}
	PictorialCarmine               = Color{R: 0.765, G: 0.043, B: 0.306, A: 1}
	Bazaar                         = Color{R: 0.596, G: 0.467, B: 0.482, A: 1}
	BlackForest                    = Color{R: 0.043, G: 0.075, B: 0.016, A: 1}
	Clairvoyant                    = Color{R: 0.282, G: 0.024, B: 0.337, A: 1}
	PastelPurple                   = Color{R: 0.702, G: 0.620, B: 0.710, A: 1}
	SatinSheenGold                 = Color{R: 0.796, G: 0.631, B: 0.208, A: 1}
	Tuatara                        = Color{R: 0.212, G: 0.208, B: 0.204, A: 1}
	Eminence                       = Color{R: 0.424, G: 0.188, B: 0.510, A: 1}
	FrenchPink                     = Color{R: 0.992, G: 0.424, B: 0.620, A: 1}
	Olivetone                      = Color{R: 0.443, G: 0.431, B: 0.063, A: 1}
	PastelYellow                   = Color{R: 0.992, G: 0.992, B: 0.588, A: 1}
	QuickSilver                    = Color{R: 0.651, G: 0.651, B: 0.651, A: 1}
	Bracken                        = Color{R: 0.290, G: 0.165, B: 0.016, A: 1}
	BurningSand                    = Color{R: 0.851, G: 0.576, B: 0.463, A: 1}
	CoconutCream                   = Color{R: 0.973, G: 0.969, B: 0.863, A: 1}
	DollarBill                     = Color{R: 0.522, G: 0.733, B: 0.396, A: 1}
	Zombie                         = Color{R: 0.894, G: 0.839, B: 0.608, A: 1}
	DarkSeaGreen                   = Color{R: 0.561, G: 0.737, B: 0.561, A: 1}
	Glaucous                       = Color{R: 0.376, G: 0.510, B: 0.714, A: 1}
	RangoonGreen                   = Color{R: 0.110, G: 0.118, B: 0.075, A: 1}
	TotemPole                      = Color{R: 0.600, G: 0.106, B: 0.027, A: 1}
	Maroon                         = Color{R: 0.502, G: 0.000, B: 0.000, A: 1}
	Martinique                     = Color{R: 0.212, G: 0.188, B: 0.314, A: 1}
	Nero                           = Color{R: 0.078, G: 0.024, B: 0.000, A: 1}
	Gimblet                        = Color{R: 0.722, G: 0.710, B: 0.416, A: 1}
	TerraCotta                     = Color{R: 0.886, G: 0.447, B: 0.357, A: 1}
	Sapphire                       = Color{R: 0.059, G: 0.322, B: 0.729, A: 1}
	Gondola                        = Color{R: 0.149, G: 0.078, B: 0.078, A: 1}
	Primrose                       = Color{R: 0.929, G: 0.918, B: 0.600, A: 1}
	Cinnabar                       = Color{R: 0.890, G: 0.259, B: 0.204, A: 1}
	Golden                         = Color{R: 1.000, G: 0.843, B: 0.000, A: 1}
	MayGreen                       = Color{R: 0.298, G: 0.569, B: 0.255, A: 1}
	AlloyOrange                    = Color{R: 0.769, G: 0.384, B: 0.063, A: 1}
	BrandyPunch                    = Color{R: 0.804, G: 0.518, B: 0.161, A: 1}
	DesaturatedCyan                = Color{R: 0.400, G: 0.600, B: 0.600, A: 1}
	RYBOrange                      = Color{R: 0.984, G: 0.600, B: 0.008, A: 1}
	Smalt                          = Color{R: 0.000, G: 0.200, B: 0.600, A: 1}
	BrightYellow                   = Color{R: 1.000, G: 0.667, B: 0.114, A: 1}
	SpunPearl                      = Color{R: 0.667, G: 0.671, B: 0.718, A: 1}
	Apache                         = Color{R: 0.875, G: 0.745, B: 0.435, A: 1}
	Belgion                        = Color{R: 0.678, G: 0.847, B: 1.000, A: 1}
	CherryBlossomPink              = Color{R: 1.000, G: 0.718, B: 0.773, A: 1}
	FlaxSmoke                      = Color{R: 0.482, G: 0.510, B: 0.396, A: 1}
	SeaSerpent                     = Color{R: 0.294, G: 0.780, B: 0.812, A: 1}
	FOGRA39RichBlack               = Color{R: 0.004, G: 0.008, B: 0.012, A: 1}
	Gossip                         = Color{R: 0.824, G: 0.973, B: 0.690, A: 1}
	Waterspout                     = Color{R: 0.643, G: 0.957, B: 0.976, A: 1}
	Schist                         = Color{R: 0.663, G: 0.706, B: 0.592, A: 1}
	BitterLemon                    = Color{R: 0.792, G: 0.878, B: 0.051, A: 1}
	Boysenberry                    = Color{R: 0.529, G: 0.196, B: 0.376, A: 1}
	Feijoa                         = Color{R: 0.624, G: 0.867, B: 0.549, A: 1}
	GenericViridian                = Color{R: 0.000, G: 0.498, B: 0.400, A: 1}
	MediumSeaGreen                 = Color{R: 0.235, G: 0.702, B: 0.443, A: 1}
	CornField                      = Color{R: 0.973, G: 0.980, B: 0.804, A: 1}
	CrayolaBlue                    = Color{R: 0.122, G: 0.459, B: 0.996, A: 1}
	PaleBrown                      = Color{R: 0.596, G: 0.463, B: 0.329, A: 1}
	RoseBud                        = Color{R: 0.984, G: 0.698, B: 0.639, A: 1}
	DeepRed                        = Color{R: 0.522, G: 0.004, B: 0.004, A: 1}
	Elm                            = Color{R: 0.110, G: 0.486, B: 0.490, A: 1}
	Gin                            = Color{R: 0.910, G: 0.949, B: 0.922, A: 1}
	GoldenPoppy                    = Color{R: 0.988, G: 0.761, B: 0.000, A: 1}
	Nickel                         = Color{R: 0.447, G: 0.455, B: 0.447, A: 1}
	Malibu                         = Color{R: 0.490, G: 0.784, B: 0.969, A: 1}
	Oregon                         = Color{R: 0.608, G: 0.278, B: 0.012, A: 1}
	Zomp                           = Color{R: 0.224, G: 0.655, B: 0.557, A: 1}
	UARed                          = Color{R: 0.851, G: 0.000, B: 0.298, A: 1}
	Broom                          = Color{R: 1.000, G: 0.925, B: 0.075, A: 1}
	Carnation                      = Color{R: 0.976, G: 0.353, B: 0.380, A: 1}
	Grenadier                      = Color{R: 0.835, G: 0.275, B: 0.000, A: 1}
	Onahau                         = Color{R: 0.804, G: 0.957, B: 1.000, A: 1}
	ShuttleGray                    = Color{R: 0.373, G: 0.400, B: 0.447, A: 1}
	Burgundy                       = Color{R: 0.502, G: 0.000, B: 0.125, A: 1}
	Mosque                         = Color{R: 0.012, G: 0.416, B: 0.431, A: 1}
	MyrtleGreen                    = Color{R: 0.192, G: 0.471, B: 0.451, A: 1}
	Sunny                          = Color{R: 0.949, G: 0.949, B: 0.478, A: 1}
	CyberGrape                     = Color{R: 0.345, G: 0.259, B: 0.486, A: 1}
	EnglishLavender                = Color{R: 0.706, G: 0.514, B: 0.584, A: 1}
	WindsorTan                     = Color{R: 0.655, G: 0.333, B: 0.008, A: 1}
	PastelMagenta                  = Color{R: 0.957, G: 0.604, B: 0.761, A: 1}
	Tacha                          = Color{R: 0.839, G: 0.773, B: 0.384, A: 1}
	Meteor                         = Color{R: 0.816, G: 0.490, B: 0.071, A: 1}
	Wewak                          = Color{R: 0.945, G: 0.608, B: 0.671, A: 1}
	WhiteSmoke                     = Color{R: 0.961, G: 0.961, B: 0.961, A: 1}
	AppleBlossom                   = Color{R: 0.686, G: 0.302, B: 0.263, A: 1}
	BrightLilac                    = Color{R: 0.847, G: 0.569, B: 0.937, A: 1}
	Chardon                        = Color{R: 1.000, G: 0.953, B: 0.945, A: 1}
	DarkImperialBlue               = Color{R: 0.431, G: 0.431, B: 0.976, A: 1}
	Eggplant                       = Color{R: 0.380, G: 0.251, B: 0.318, A: 1}
	BigStone                       = Color{R: 0.086, G: 0.165, B: 0.251, A: 1}
	DarkCyan                       = Color{R: 0.000, G: 0.545, B: 0.545, A: 1}
	PersianGreen                   = Color{R: 0.000, G: 0.651, B: 0.576, A: 1}
	VividCerulean                  = Color{R: 0.000, G: 0.667, B: 0.933, A: 1}
	Crocodile                      = Color{R: 0.451, G: 0.427, B: 0.345, A: 1}
	Millbrook                      = Color{R: 0.349, G: 0.267, B: 0.200, A: 1}
	OffGreen                       = Color{R: 0.902, G: 0.973, B: 0.953, A: 1}
	UniversityOfTennesseeOrange    = Color{R: 0.969, G: 0.498, B: 0.000, A: 1}
	DeepSaffron                    = Color{R: 1.000, G: 0.600, B: 0.200, A: 1}
	RichCarmine                    = Color{R: 0.843, G: 0.000, B: 0.251, A: 1}
	BlueViolet                     = Color{R: 0.541, G: 0.169, B: 0.886, A: 1}
	GrannySmithApple               = Color{R: 0.659, G: 0.894, B: 0.627, A: 1}
	Organ                          = Color{R: 0.424, G: 0.180, B: 0.122, A: 1}
	PaleGreen                      = Color{R: 0.596, G: 0.984, B: 0.596, A: 1}
	Tangelo                        = Color{R: 0.976, G: 0.302, B: 0.000, A: 1}
	SizzlingRed                    = Color{R: 1.000, G: 0.220, B: 0.333, A: 1}
	Tapestry                       = Color{R: 0.690, G: 0.369, B: 0.506, A: 1}
	BdazzledBlue                   = Color{R: 0.180, G: 0.345, B: 0.580, A: 1}
	Frostbite                      = Color{R: 0.914, G: 0.212, B: 0.655, A: 1}
	Ivory                          = Color{R: 1.000, G: 1.000, B: 0.941, A: 1}
	KeyLimePie                     = Color{R: 0.749, G: 0.788, B: 0.129, A: 1}
	Shilo                          = Color{R: 0.910, G: 0.725, B: 0.702, A: 1}
	VividLimeGreen                 = Color{R: 0.651, G: 0.839, B: 0.031, A: 1}
	VividSkyBlue                   = Color{R: 0.000, G: 0.800, B: 1.000, A: 1}
	Bourbon                        = Color{R: 0.729, G: 0.435, B: 0.118, A: 1}
	DarkByzantium                  = Color{R: 0.365, G: 0.224, B: 0.329, A: 1}
	ElPaso                         = Color{R: 0.118, G: 0.090, B: 0.031, A: 1}
	GrayAsparagus                  = Color{R: 0.275, G: 0.349, B: 0.271, A: 1}
	VanCleef                       = Color{R: 0.286, G: 0.090, B: 0.047, A: 1}
	ArcticLime                     = Color{R: 0.816, G: 1.000, B: 0.078, A: 1}
	Tiara                          = Color{R: 0.765, G: 0.820, B: 0.820, A: 1}
	WineDregs                      = Color{R: 0.404, G: 0.192, B: 0.278, A: 1}
	Solitaire                      = Color{R: 0.996, G: 0.973, B: 0.886, A: 1}
	Byzantine                      = Color{R: 0.741, G: 0.200, B: 0.643, A: 1}
	DarkJungleGreen                = Color{R: 0.102, G: 0.141, B: 0.129, A: 1}
	MorningGlory                   = Color{R: 0.620, G: 0.871, B: 0.878, A: 1}
	ParisM                         = Color{R: 0.149, G: 0.020, B: 0.416, A: 1}
	PickledBluewood                = Color{R: 0.192, G: 0.267, B: 0.349, A: 1}
	BostonBlue                     = Color{R: 0.231, G: 0.569, B: 0.706, A: 1}
	LilacLuster                    = Color{R: 0.682, G: 0.596, B: 0.667, A: 1}
	Opium                          = Color{R: 0.557, G: 0.435, B: 0.439, A: 1}
	Scorpion                       = Color{R: 0.412, G: 0.373, B: 0.384, A: 1}
	Voodoo                         = Color{R: 0.325, G: 0.204, B: 0.333, A: 1}
	GreenCyan                      = Color{R: 0.000, G: 0.600, B: 0.400, A: 1}
	HawaiianTan                    = Color{R: 0.616, G: 0.337, B: 0.086, A: 1}
	Nutmeg                         = Color{R: 0.506, G: 0.259, B: 0.173, A: 1}
	BlueGreen                      = Color{R: 0.051, G: 0.596, B: 0.729, A: 1}
	Cabaret                        = Color{R: 0.851, G: 0.286, B: 0.447, A: 1}
	Carissma                       = Color{R: 0.918, G: 0.533, B: 0.659, A: 1}
	ChateauGreen                   = Color{R: 0.251, G: 0.659, B: 0.376, A: 1}
	Creole                         = Color{R: 0.118, G: 0.059, B: 0.016, A: 1}
	TreePoppy                      = Color{R: 0.988, G: 0.612, B: 0.114, A: 1}
	AzureishWhite                  = Color{R: 0.859, G: 0.914, B: 0.957, A: 1}
	Finch                          = Color{R: 0.384, G: 0.400, B: 0.286, A: 1}
	RedRobin                       = Color{R: 0.502, G: 0.204, B: 0.122, A: 1}
	DarkKhaki                      = Color{R: 0.741, G: 0.718, B: 0.420, A: 1}
	ElectricLime                   = Color{R: 0.800, G: 1.000, B: 0.000, A: 1}
	MediumSlateBlue                = Color{R: 0.482, G: 0.408, B: 0.933, A: 1}
	Olive                          = Color{R: 0.502, G: 0.502, B: 0.000, A: 1}
	Violet                         = Color{R: 0.498, G: 0.000, B: 1.000, A: 1}
	AntiqueWhite                   = Color{R: 0.980, G: 0.922, B: 0.843, A: 1}
	CedarChest                     = Color{R: 0.788, G: 0.353, B: 0.286, A: 1}
	LisbonBrown                    = Color{R: 0.259, G: 0.224, B: 0.129, A: 1}
	Vesuvius                       = Color{R: 0.694, G: 0.290, B: 0.043, A: 1}
	Flamingo                       = Color{R: 0.949, G: 0.333, B: 0.165, A: 1}
	GrayChateau                    = Color{R: 0.635, G: 0.667, B: 0.702, A: 1}
	Jon                            = Color{R: 0.231, G: 0.122, B: 0.122, A: 1}
	NightShadz                     = Color{R: 0.667, G: 0.216, B: 0.353, A: 1}
	SteelBlue                      = Color{R: 0.275, G: 0.510, B: 0.706, A: 1}
	Eclipse                        = Color{R: 0.192, G: 0.110, B: 0.090, A: 1}
	LavenderBlush                  = Color{R: 1.000, G: 0.941, B: 0.961, A: 1}
	PaleLavender                   = Color{R: 0.863, G: 0.816, B: 1.000, A: 1}
	PalmGreen                      = Color{R: 0.035, G: 0.137, B: 0.059, A: 1}
	WeldonBlue                     = Color{R: 0.486, G: 0.596, B: 0.671, A: 1}
	Blumine                        = Color{R: 0.094, G: 0.345, B: 0.478, A: 1}
	Jewel                          = Color{R: 0.071, G: 0.420, B: 0.251, A: 1}
	Mint                           = Color{R: 0.243, G: 0.706, B: 0.537, A: 1}
	ButterflyBush                  = Color{R: 0.384, G: 0.306, B: 0.604, A: 1}
	Flame                          = Color{R: 0.886, G: 0.345, B: 0.133, A: 1}
	Jasper                         = Color{R: 0.843, G: 0.231, B: 0.243, A: 1}
	Midnight                       = Color{R: 0.439, G: 0.149, B: 0.439, A: 1}
	Cinereous                      = Color{R: 0.596, G: 0.506, B: 0.482, A: 1}
	Falcon                         = Color{R: 0.498, G: 0.384, B: 0.427, A: 1}
	Sidecar                        = Color{R: 0.953, G: 0.906, B: 0.733, A: 1}
	SoyaBean                       = Color{R: 0.416, G: 0.376, B: 0.318, A: 1}
	VividVermilion                 = Color{R: 0.898, G: 0.376, B: 0.141, A: 1}
	Castro                         = Color{R: 0.322, G: 0.000, B: 0.122, A: 1}
	Sorbus                         = Color{R: 0.992, G: 0.486, B: 0.027, A: 1}
	Travertine                     = Color{R: 1.000, G: 0.992, B: 0.910, A: 1}
	Atlantis                       = Color{R: 0.592, G: 0.804, B: 0.176, A: 1}
	Bush                           = Color{R: 0.051, G: 0.180, B: 0.110, A: 1}
	Cameo                          = Color{R: 0.851, G: 0.725, B: 0.608, A: 1}
	MistyMoss                      = Color{R: 0.733, G: 0.706, B: 0.467, A: 1}
	Lumber                         = Color{R: 1.000, G: 0.894, B: 0.804, A: 1}
	MarigoldYellow                 = Color{R: 0.984, G: 0.910, B: 0.439, A: 1}
	RoastCoffee                    = Color{R: 0.439, G: 0.259, B: 0.255, A: 1}
	RoseTaupe                      = Color{R: 0.565, G: 0.365, B: 0.365, A: 1}
	ShadowBlue                     = Color{R: 0.467, G: 0.545, B: 0.647, A: 1}
	Swirl                          = Color{R: 0.827, G: 0.804, B: 0.773, A: 1}
	Wenge                          = Color{R: 0.392, G: 0.329, B: 0.322, A: 1}
	BarbiePink                     = Color{R: 0.878, G: 0.129, B: 0.541, A: 1}
	BermudaGray                    = Color{R: 0.420, G: 0.545, B: 0.635, A: 1}
	RawUmber                       = Color{R: 0.510, G: 0.400, B: 0.267, A: 1}
	Ruby                           = Color{R: 0.878, G: 0.067, B: 0.373, A: 1}
	Starship                       = Color{R: 0.925, G: 0.949, B: 0.271, A: 1}
	IndigoDye                      = Color{R: 0.035, G: 0.122, B: 0.573, A: 1}
	Paarl                          = Color{R: 0.651, G: 0.333, B: 0.161, A: 1}
	PinkLady                       = Color{R: 1.000, G: 0.945, B: 0.847, A: 1}
	WaxFlower                      = Color{R: 1.000, G: 0.753, B: 0.659, A: 1}
	CitrineWhite                   = Color{R: 0.980, G: 0.969, B: 0.839, A: 1}
	MineShaft                      = Color{R: 0.196, G: 0.196, B: 0.196, A: 1}
	PansyPurple                    = Color{R: 0.471, G: 0.094, B: 0.290, A: 1}
	Sycamore                       = Color{R: 0.565, G: 0.553, B: 0.224, A: 1}
	Waiouru                        = Color{R: 0.212, G: 0.235, B: 0.051, A: 1}
	SandyBrown                     = Color{R: 0.957, G: 0.643, B: 0.376, A: 1}
	VioletBlue                     = Color{R: 0.196, G: 0.290, B: 0.698, A: 1}
	CGRed                          = Color{R: 0.878, G: 0.235, B: 0.192, A: 1}
	CatskillWhite                  = Color{R: 0.933, G: 0.965, B: 0.969, A: 1}
	DarkCerulean                   = Color{R: 0.031, G: 0.271, B: 0.494, A: 1}
	LimeGreen                      = Color{R: 0.196, G: 0.804, B: 0.196, A: 1}
	Moccasin                       = Color{R: 1.000, G: 0.894, B: 0.710, A: 1}
	CarrotOrange                   = Color{R: 0.929, G: 0.569, B: 0.129, A: 1}
	DiSerria                       = Color{R: 0.859, G: 0.600, B: 0.369, A: 1}
	LogCabin                       = Color{R: 0.141, G: 0.165, B: 0.114, A: 1}
	VividOrange                    = Color{R: 1.000, G: 0.373, B: 0.000, A: 1}
	Bouquet                        = Color{R: 0.682, G: 0.502, B: 0.620, A: 1}
	DarkSlateBlue                  = Color{R: 0.282, G: 0.239, B: 0.545, A: 1}
	QuarterPearlLusta              = Color{R: 1.000, G: 0.992, B: 0.957, A: 1}
	Padua                          = Color{R: 0.678, G: 0.902, B: 0.769, A: 1}
	Rangitoto                      = Color{R: 0.180, G: 0.196, B: 0.133, A: 1}
	WestCoast                      = Color{R: 0.384, G: 0.318, B: 0.098, A: 1}
	PixieGreen                     = Color{R: 0.753, G: 0.847, B: 0.714, A: 1}
	CornflowerLilac                = Color{R: 1.000, G: 0.690, B: 0.675, A: 1}
	NCSGreen                       = Color{R: 0.000, G: 0.624, B: 0.420, A: 1}
	Wattle                         = Color{R: 0.863, G: 0.843, B: 0.278, A: 1}
	AquamarineBlue                 = Color{R: 0.443, G: 0.851, B: 0.886, A: 1}
	Arsenic                        = Color{R: 0.231, G: 0.267, B: 0.294, A: 1}
	BlueYonder                     = Color{R: 0.314, G: 0.447, B: 0.655, A: 1}
	ChelseaGem                     = Color{R: 0.620, G: 0.325, B: 0.008, A: 1}
	ChetwodeBlue                   = Color{R: 0.522, G: 0.506, B: 0.851, A: 1}
	PurpleTaupe                    = Color{R: 0.314, G: 0.251, B: 0.302, A: 1}
	Armadillo                      = Color{R: 0.263, G: 0.243, B: 0.216, A: 1}
	FrenchGray                     = Color{R: 0.741, G: 0.741, B: 0.776, A: 1}
	SeaMist                        = Color{R: 0.773, G: 0.859, B: 0.792, A: 1}
	X11DarkGreen                   = Color{R: 0.000, G: 0.392, B: 0.000, A: 1}
	LightSkyBlue                   = Color{R: 0.529, G: 0.808, B: 0.980, A: 1}
	Mondo                          = Color{R: 0.290, G: 0.235, B: 0.188, A: 1}
	RedPurple                      = Color{R: 0.894, G: 0.000, B: 0.471, A: 1}
	SpanishViolet                  = Color{R: 0.298, G: 0.157, B: 0.510, A: 1}
	Toast                          = Color{R: 0.604, G: 0.431, B: 0.380, A: 1}
	VioletEggplant                 = Color{R: 0.600, G: 0.067, B: 0.600, A: 1}
	Barberry                       = Color{R: 0.871, G: 0.843, B: 0.090, A: 1}
	GreenWhite                     = Color{R: 0.910, G: 0.922, B: 0.878, A: 1}
	JapaneseViolet                 = Color{R: 0.357, G: 0.196, B: 0.337, A: 1}
	Pesto                          = Color{R: 0.486, G: 0.463, B: 0.192, A: 1}
	SpanishCrimson                 = Color{R: 0.898, G: 0.102, B: 0.298, A: 1}
	UnmellowYellow                 = Color{R: 1.000, G: 1.000, B: 0.400, A: 1}
	Bole                           = Color{R: 0.475, G: 0.267, B: 0.231, A: 1}
	CanCan                         = Color{R: 0.835, G: 0.569, B: 0.643, A: 1}
	Lilac                          = Color{R: 0.784, G: 0.635, B: 0.784, A: 1}
	Polar                          = Color{R: 0.898, G: 0.976, B: 0.965, A: 1}
	Straw                          = Color{R: 0.894, G: 0.851, B: 0.435, A: 1}
	LightFuchsiaPink               = Color{R: 0.976, G: 0.518, B: 0.937, A: 1}
	Whisper                        = Color{R: 0.969, G: 0.961, B: 0.980, A: 1}
	Amaranth                       = Color{R: 0.898, G: 0.169, B: 0.314, A: 1}
	BlackPearl                     = Color{R: 0.016, G: 0.075, B: 0.133, A: 1}
	BlastOffBronze                 = Color{R: 0.647, G: 0.443, B: 0.392, A: 1}
	BottleGreen                    = Color{R: 0.000, G: 0.416, B: 0.306, A: 1}
	Cedar                          = Color{R: 0.243, G: 0.110, B: 0.078, A: 1}
	Sisal                          = Color{R: 0.827, G: 0.796, B: 0.729, A: 1}
	Glitter                        = Color{R: 0.902, G: 0.910, B: 0.980, A: 1}
	Independence                   = Color{R: 0.298, G: 0.318, B: 0.427, A: 1}
	Koromiko                       = Color{R: 1.000, G: 0.741, B: 0.373, A: 1}
	TitaniumYellow                 = Color{R: 0.933, G: 0.902, B: 0.000, A: 1}
	PersianIndigo                  = Color{R: 0.196, G: 0.071, B: 0.478, A: 1}
	Thunder                        = Color{R: 0.200, G: 0.161, B: 0.184, A: 1}
	Astral                         = Color{R: 0.196, G: 0.490, B: 0.627, A: 1}
	CabbagePont                    = Color{R: 0.247, G: 0.298, B: 0.227, A: 1}
	CrayolaOrange                  = Color{R: 1.000, G: 0.459, B: 0.220, A: 1}
	Matisse                        = Color{R: 0.106, G: 0.396, B: 0.616, A: 1}
	Orinoco                        = Color{R: 0.953, G: 0.984, B: 0.831, A: 1}
	BananaYellow                   = Color{R: 1.000, G: 0.882, B: 0.208, A: 1}
	TractorRed                     = Color{R: 0.992, G: 0.055, B: 0.208, A: 1}
	BrownBramble                   = Color{R: 0.349, G: 0.157, B: 0.016, A: 1}
	KashmirBlue                    = Color{R: 0.314, G: 0.439, B: 0.588, A: 1}
	YellowSea                      = Color{R: 0.996, G: 0.663, B: 0.016, A: 1}
	Americano                      = Color{R: 0.529, G: 0.459, B: 0.431, A: 1}
	CadmiumGreen                   = Color{R: 0.000, G: 0.420, B: 0.235, A: 1}
	EarlsGreen                     = Color{R: 0.788, G: 0.725, B: 0.231, A: 1}
	Mortar                         = Color{R: 0.314, G: 0.263, B: 0.318, A: 1}
	SuperPink                      = Color{R: 0.812, G: 0.420, B: 0.663, A: 1}
	RazzleDazzleRose               = Color{R: 1.000, G: 0.200, B: 0.800, A: 1}
	Shark                          = Color{R: 0.145, G: 0.153, B: 0.173, A: 1}
	BlueSapphire                   = Color{R: 0.071, G: 0.380, B: 0.502, A: 1}
	Cherub                         = Color{R: 0.973, G: 0.851, B: 0.914, A: 1}
	IndiaGreen                     = Color{R: 0.075, G: 0.533, B: 0.031, A: 1}
	Kelp                           = Color{R: 0.271, G: 0.286, B: 0.212, A: 1}
	Peridot                        = Color{R: 0.902, G: 0.886, B: 0.000, A: 1}
	PineGreen                      = Color{R: 0.004, G: 0.475, B: 0.435, A: 1}
	PrincetonOrange                = Color{R: 0.961, G: 0.502, B: 0.145, A: 1}
	William                        = Color{R: 0.227, G: 0.408, B: 0.424, A: 1}
	Minsk                          = Color{R: 0.247, G: 0.188, B: 0.498, A: 1}
	Tan                            = Color{R: 0.824, G: 0.706, B: 0.549, A: 1}
	LapisLazuli                    = Color{R: 0.149, G: 0.380, B: 0.612, A: 1}
	Milan                          = Color{R: 0.980, G: 1.000, B: 0.643, A: 1}
	MistyRose                      = Color{R: 1.000, G: 0.894, B: 0.882, A: 1}
	NewOrleans                     = Color{R: 0.953, G: 0.839, B: 0.616, A: 1}
	Roman                          = Color{R: 0.871, G: 0.388, B: 0.376, A: 1}
	AntiFlashWhite                 = Color{R: 0.949, G: 0.953, B: 0.957, A: 1}
	BrightSun                      = Color{R: 0.996, G: 0.827, B: 0.235, A: 1}
	PaynesGrey                     = Color{R: 0.325, G: 0.408, B: 0.471, A: 1}
	RegentStBlue                   = Color{R: 0.667, G: 0.839, B: 0.902, A: 1}
	VividAuburn                    = Color{R: 0.573, G: 0.153, B: 0.141, A: 1}
	PearlMysticTurquoise           = Color{R: 0.196, G: 0.776, B: 0.651, A: 1}
	VividRed                       = Color{R: 0.969, G: 0.051, B: 0.102, A: 1}
	Mallard                        = Color{R: 0.137, G: 0.204, B: 0.094, A: 1}
	Edward                         = Color{R: 0.635, G: 0.682, B: 0.671, A: 1}
	FrenchMauve                    = Color{R: 0.831, G: 0.451, B: 0.831, A: 1}
	FrenchWine                     = Color{R: 0.675, G: 0.118, B: 0.267, A: 1}
	Stonewall                      = Color{R: 0.573, G: 0.522, B: 0.451, A: 1}
	ThatchGreen                    = Color{R: 0.251, G: 0.239, B: 0.098, A: 1}
	PastelRed                      = Color{R: 1.000, G: 0.412, B: 0.380, A: 1}
	Sage                           = Color{R: 0.737, G: 0.722, B: 0.541, A: 1}
	Gunsmoke                       = Color{R: 0.510, G: 0.525, B: 0.522, A: 1}
	HummingBird                    = Color{R: 0.812, G: 0.976, B: 0.953, A: 1}
	KaitokeGreen                   = Color{R: 0.000, G: 0.275, B: 0.125, A: 1}
	Madang                         = Color{R: 0.718, G: 0.941, B: 0.745, A: 1}
	MunsellBlue                    = Color{R: 0.000, G: 0.576, B: 0.686, A: 1}
	Brandy                         = Color{R: 0.871, G: 0.757, B: 0.588, A: 1}
	CatalinaBlue                   = Color{R: 0.024, G: 0.165, B: 0.471, A: 1}
	TimberGreen                    = Color{R: 0.086, G: 0.196, B: 0.173, A: 1}
	FrenchLime                     = Color{R: 0.620, G: 0.992, B: 0.220, A: 1}
	Keppel                         = Color{R: 0.227, G: 0.690, B: 0.620, A: 1}
	Teak                           = Color{R: 0.694, G: 0.580, B: 0.380, A: 1}
	Tundora                        = Color{R: 0.290, G: 0.259, B: 0.267, A: 1}
	Transparent                    = Color{}
	Map                            = map[string]Color{
		"bermudagray":                    BermudaGray,
		"rawumber":                       RawUmber,
		"ruby":                           Ruby,
		"starship":                       Starship,
		"swirl":                          Swirl,
		"wenge":                          Wenge,
		"barbiepink":                     BarbiePink,
		"paarl":                          Paarl,
		"pinklady":                       PinkLady,
		"indigodye":                      IndigoDye,
		"mineshaft":                      MineShaft,
		"pansypurple":                    PansyPurple,
		"sycamore":                       Sycamore,
		"waiouru":                        Waiouru,
		"waxflower":                      WaxFlower,
		"citrinewhite":                   CitrineWhite,
		"catskillwhite":                  CatskillWhite,
		"darkcerulean":                   DarkCerulean,
		"limegreen":                      LimeGreen,
		"moccasin":                       Moccasin,
		"sandybrown":                     SandyBrown,
		"violetblue":                     VioletBlue,
		"cgred":                          CGRed,
		"diserria":                       DiSerria,
		"logcabin":                       LogCabin,
		"vividorange":                    VividOrange,
		"carrotorange":                   CarrotOrange,
		"darkslateblue":                  DarkSlateBlue,
		"quarterpearllusta":              QuarterPearlLusta,
		"bouquet":                        Bouquet,
		"rangitoto":                      Rangitoto,
		"westcoast":                      WestCoast,
		"padua":                          Padua,
		"pixiegreen":                     PixieGreen,
		"arsenic":                        Arsenic,
		"blueyonder":                     BlueYonder,
		"chelseagem":                     ChelseaGem,
		"chetwodeblue":                   ChetwodeBlue,
		"cornflowerlilac":                CornflowerLilac,
		"ncsgreen":                       NCSGreen,
		"wattle":                         Wattle,
		"aquamarineblue":                 AquamarineBlue,
		"purpletaupe":                    PurpleTaupe,
		"frenchgray":                     FrenchGray,
		"seamist":                        SeaMist,
		"x11darkgreen":                   X11DarkGreen,
		"armadillo":                      Armadillo,
		"mondo":                          Mondo,
		"redpurple":                      RedPurple,
		"lightskyblue":                   LightSkyBlue,
		"greenwhite":                     GreenWhite,
		"japaneseviolet":                 JapaneseViolet,
		"pesto":                          Pesto,
		"spanishcrimson":                 SpanishCrimson,
		"spanishviolet":                  SpanishViolet,
		"toast":                          Toast,
		"violeteggplant":                 VioletEggplant,
		"barberry":                       Barberry,
		"cancan":                         CanCan,
		"lilac":                          Lilac,
		"polar":                          Polar,
		"straw":                          Straw,
		"unmellowyellow":                 UnmellowYellow,
		"bole":                           Bole,
		"blackpearl":                     BlackPearl,
		"blastoffbronze":                 BlastOffBronze,
		"bottlegreen":                    BottleGreen,
		"cedar":                          Cedar,
		"lightfuchsiapink":               LightFuchsiaPink,
		"whisper":                        Whisper,
		"amaranth":                       Amaranth,
		"sisal":                          Sisal,
		"independence":                   Independence,
		"koromiko":                       Koromiko,
		"titaniumyellow":                 TitaniumYellow,
		"glitter":                        Glitter,
		"cabbagepont":                    CabbagePont,
		"crayolaorange":                  CrayolaOrange,
		"matisse":                        Matisse,
		"orinoco":                        Orinoco,
		"persianindigo":                  PersianIndigo,
		"thunder":                        Thunder,
		"astral":                         Astral,
		"tractorred":                     TractorRed,
		"bananayellow":                   BananaYellow,
		"kashmirblue":                    KashmirBlue,
		"yellowsea":                      YellowSea,
		"brownbramble":                   BrownBramble,
		"cadmiumgreen":                   CadmiumGreen,
		"earlsgreen":                     EarlsGreen,
		"mortar":                         Mortar,
		"superpink":                      SuperPink,
		"americano":                      Americano,
		"cherub":                         Cherub,
		"indiagreen":                     IndiaGreen,
		"kelp":                           Kelp,
		"peridot":                        Peridot,
		"razzledazzlerose":               RazzleDazzleRose,
		"shark":                          Shark,
		"bluesapphire":                   BlueSapphire,
		"princetonorange":                PrincetonOrange,
		"william":                        William,
		"pinegreen":                      PineGreen,
		"tan":                            Tan,
		"minsk":                          Minsk,
		"milan":                          Milan,
		"mistyrose":                      MistyRose,
		"neworleans":                     NewOrleans,
		"roman":                          Roman,
		"lapislazuli":                    LapisLazuli,
		"brightsun":                      BrightSun,
		"paynesgrey":                     PaynesGrey,
		"regentstblue":                   RegentStBlue,
		"vividauburn":                    VividAuburn,
		"antiflashwhite":                 AntiFlashWhite,
		"vividred":                       VividRed,
		"pearlmysticturquoise":           PearlMysticTurquoise,
		"mallard":                        Mallard,
		"frenchmauve":                    FrenchMauve,
		"frenchwine":                     FrenchWine,
		"stonewall":                      Stonewall,
		"thatchgreen":                    ThatchGreen,
		"edward":                         Edward,
		"hummingbird":                    HummingBird,
		"kaitokegreen":                   KaitokeGreen,
		"madang":                         Madang,
		"munsellblue":                    MunsellBlue,
		"pastelred":                      PastelRed,
		"sage":                           Sage,
		"gunsmoke":                       Gunsmoke,
		"catalinablue":                   CatalinaBlue,
		"timbergreen":                    TimberGreen,
		"brandy":                         Brandy,
		"keppel":                         Keppel,
		"teak":                           Teak,
		"tundora":                        Tundora,
		"frenchlime":                     FrenchLime,
		"oldcopper":                      OldCopper,
		"frenchpuce":                     FrenchPuce,
		"watusi":                         Watusi,
		"yukongold":                      YukonGold,
		"lust":                           Lust,
		"peachcream":                     PeachCream,
		"vividamber":                     VividAmber,
		"limerick":                       Limerick,
		"celadon":                        Celadon,
		"chambray":                       Chambray,
		"darkpastelblue":                 DarkPastelBlue,
		"greenspring":                    GreenSpring,
		"madras":                         Madras,
		"vividgamboge":                   VividGamboge,
		"bianca":                         Bianca,
		"celeste":                        Celeste,
		"fuelyellow":                     FuelYellow,
		"lightmediumorchid":              LightMediumOrchid,
		"matterhorn":                     Matterhorn,
		"salem":                          Salem,
		"aquaforest":                     AquaForest,
		"lightsalmonpink":                LightSalmonPink,
		"niagara":                        Niagara,
		"rainee":                         Rainee,
		"sanjuan":                        SanJuan,
		"cyancornflowerblue":             CyanCornflowerBlue,
		"harlequin":                      Harlequin,
		"offyellow":                      OffYellow,
		"upforestgreen":                  UPForestGreen,
		"cactus":                         Cactus,
		"coppercanyon":                   CopperCanyon,
		"illusion":                       Illusion,
		"onyx":                           Onyx,
		"perfume":                        Perfume,
		"psychedelicpurple":              PsychedelicPurple,
		"blackolive":                     BlackOlive,
		"emperor":                        Emperor,
		"trendypink":                     TrendyPink,
		"turquoise":                      Turquoise,
		"deeppuce":                       DeepPuce,
		"blackwhite":                     BlackWhite,
		"castletongreen":                 CastletonGreen,
		"black":                          Black,
		"cottonseed":                     CottonSeed,
		"darkred":                        DarkRed,
		"darkscarlet":                    DarkScarlet,
		"deepmaroon":                     DeepMaroon,
		"mercury":                        Mercury,
		"merlot":                         Merlot,
		"chicago":                        Chicago,
		"waterleaf":                      WaterLeaf,
		"darkpuce":                       DarkPuce,
		"skymagenta":                     SkyMagenta,
		"lightcarminepink":               LightCarminePink,
		"doublepearllusta":               DoublePearlLusta,
		"grannyapple":                    GrannyApple,
		"pinkorange":                     PinkOrange,
		"tonyspink":                      TonysPink,
		"deepoak":                        DeepOak,
		"deepblue":                       DeepBlue,
		"folly":                          Folly,
		"tiffanyblue":                    TiffanyBlue,
		"bud":                            Bud,
		"horizon":                        Horizon,
		"kellygreen":                     KellyGreen,
		"muesli":                         Muesli,
		"stromboli":                      Stromboli,
		"envy":                           Envy,
		"kucrimson":                      KUCrimson,
		"brunswickgreen":                 BrunswickGreen,
		"palatinatepurple":               PalatinatePurple,
		"tuscany":                        Tuscany,
		"darkskyblue":                    DarkSkyBlue,
		"frenchviolet":                   FrenchViolet,
		"royalpurple":                    RoyalPurple,
		"shamrock":                       Shamrock,
		"cadmiumred":                     CadmiumRed,
		"lemon":                          Lemon,
		"lynch":                          Lynch,
		"metalliccopper":                 MetallicCopper,
		"monsoon":                        Monsoon,
		"moroccobrown":                   MoroccoBrown,
		"towergray":                      TowerGray,
		"volt":                           Volt,
		"coralreef":                      CoralReef,
		"cadmiumorange":                  CadmiumOrange,
		"shiraz":                         Shiraz,
		"bahamablue":                     BahamaBlue,
		"dogwoodrose":                    DogwoodRose,
		"frenchrose":                     FrenchRose,
		"picasso":                        Picasso,
		"ultramarineblue":                UltramarineBlue,
		"brightmaroon":                   BrightMaroon,
		"iron":                           Iron,
		"orangeyellow":                   OrangeYellow,
		"palesky":                        PaleSky,
		"hippiegreen":                    HippieGreen,
		"akaroa":                         Akaroa,
		"alpine":                         Alpine,
		"bondiblue":                      BondiBlue,
		"cadet":                          Cadet,
		"celadongreen":                   CeladonGreen,
		"feldgrau":                       Feldgrau,
		"pearlaqua":                      PearlAqua,
		"absolutezero":                   AbsoluteZero,
		"rosybrown":                      RosyBrown,
		"tarawera":                       Tarawera,
		"taupe":                          Taupe,
		"remy":                           Remy,
		"aquadeep":                       AquaDeep,
		"coconut":                        Coconut,
		"languidlavender":                LanguidLavender,
		"lavendermagenta":                LavenderMagenta,
		"lilywhite":                      LilyWhite,
		"abbey":                          Abbey,
		"cyclamen":                       Cyclamen,
		"glacier":                        Glacier,
		"lightsteelblue":                 LightSteelBlue,
		"trinidad":                       Trinidad,
		"whiteice":                       WhiteIce,
		"acadia":                         Acadia,
		"dovegray":                       DoveGray,
		"fairpink":                       FairPink,
		"finn":                           Finn,
		"isabelline":                     Isabelline,
		"milkpunch":                      MilkPunch,
		"vanillaice":                     VanillaIce,
		"casablanca":                     Casablanca,
		"palecanary":                     PaleCanary,
		"pigmentgreen":                   PigmentGreen,
		"driftwood":                      Driftwood,
		"imperial":                       Imperial,
		"outerspace":                     OuterSpace,
		"pantonemagenta":                 PantoneMagenta,
		"capehoney":                      CapeHoney,
		"elfgreen":                       ElfGreen,
		"fernfrond":                      FernFrond,
		"lotus":                          Lotus,
		"sunsetorange":                   SunsetOrange,
		"cadillac":                       Cadillac,
		"biscay":                         Biscay,
		"darktan":                        DarkTan,
		"universityofcaliforniagold":     UniversityOfCaliforniaGold,
		"upsdellred":                     UpsdellRed,
		"amethyst":                       Amethyst,
		"scarlet":                        Scarlet,
		"bluegem":                        BlueGem,
		"carnelian":                      Carnelian,
		"lightapricot":                   LightApricot,
		"peachyellow":                    PeachYellow,
		"rockblue":                       RockBlue,
		"wheat":                          Wheat,
		"alizarincrimson":                AlizarinCrimson,
		"wildorchid":                     WildOrchid,
		"mummystomb":                     MummysTomb,
		"rebeccapurple":                  RebeccaPurple,
		"cola":                           Cola,
		"debianred":                      DebianRed,
		"greenblue":                      GreenBlue,
		"pirategold":                     PirateGold,
		"rybviolet":                      RYBViolet,
		"twilightblue":                   TwilightBlue,
		"brown":                          Brown,
		"goldentainoi":                   GoldenTainoi,
		"harp":                           Harp,
		"maximumyellow":                  MaximumYellow,
		"munsellgreen":                   MunsellGreen,
		"oldgold":                        OldGold,
		"summergreen":                    SummerGreen,
		"carla":                          Carla,
		"waterloo":                       Waterloo,
		"yuma":                           Yuma,
		"viola":                          Viola,
		"roti":                           Roti,
		"englishwalnut":                  EnglishWalnut,
		"elsalva":                        ElSalva,
		"goldenyellow":                   GoldenYellow,
		"vandykebrown":                   VanDykeBrown,
		"davysgrey":                      DavysGrey,
		"paco":                           Paco,
		"peanut":                         Peanut,
		"salmon":                         Salmon,
		"tuscanred":                      TuscanRed,
		"beige":                          Beige,
		"prussianblue":                   PrussianBlue,
		"robineggblue":                   RobinEggBlue,
		"rum":                            Rum,
		"spacecadet":                     SpaceCadet,
		"spicypink":                      SpicyPink,
		"starcommandblue":                StarCommandBlue,
		"flax":                           Flax,
		"deepcarrotorange":               DeepCarrotOrange,
		"flirt":                          Flirt,
		"conifer":                        Conifer,
		"munsellpurple":                  MunsellPurple,
		"paradisepink":                   ParadisePink,
		"prelude":                        Prelude,
		"gordonsgreen":                   GordonsGreen,
		"fuchsiapink":                    FuchsiaPink,
		"ghost":                          Ghost,
		"icecold":                        IceCold,
		"karry":                          Karry,
		"royalazure":                     RoyalAzure,
		"donkeybrown":                    DonkeyBrown,
		"phthaloblue":                    PhthaloBlue,
		"radicalred":                     RadicalRed,
		"tusk":                           Tusk,
		"deepforestgreen":                DeepForestGreen,
		"pineglade":                      PineGlade,
		"unitednationsblue":              UnitedNationsBlue,
		"bluebonnet":                     Bluebonnet,
		"fogra29richblack":               FOGRA29RichBlack,
		"melrose":                        Melrose,
		"scooter":                        Scooter,
		"vividraspberry":                 VividRaspberry,
		"witchhaze":                      WitchHaze,
		"dolphin":                        Dolphin,
		"fieryrose":                      FieryRose,
		"jazzberryjam":                   JazzberryJam,
		"ruddy":                          Ruddy,
		"browntumbleweed":                BrownTumbleweed,
		"carminered":                     CarmineRed,
		"fountainblue":                   FountainBlue,
		"ticklemepink":                   TickleMePink,
		"turquoiseblue":                  TurquoiseBlue,
		"bronzeyellow":                   BronzeYellow,
		"romancoffee":                    RomanCoffee,
		"illuminatingemerald":            IlluminatingEmerald,
		"copperred":                      CopperRed,
		"marigold":                       Marigold,
		"yellow":                         Yellow,
		"ceil":                           Ceil,
		"brownpod":                       BrownPod,
		"mediumpurple":                   MediumPurple,
		"smokytopaz":                     SmokyTopaz,
		"snowflurry":                     SnowFlurry,
		"yaleblue":                       YaleBlue,
		"astra":                          Astra,
		"piggypink":                      PiggyPink,
		"pinetree":                       PineTree,
		"piper":                          Piper,
		"paoloveronesegreen":             PaoloVeroneseGreen,
		"jonquil":                        Jonquil,
		"toledo":                         Toledo,
		"vividmalachite":                 VividMalachite,
		"deepspacesparkle":               DeepSpaceSparkle,
		"solitude":                       Solitude,
		"toolbox":                        Toolbox,
		"russet":                         Russet,
		"darkbluegray":                   DarkBlueGray,
		"darkterracotta":                 DarkTerraCotta,
		"mysin":                          MySin,
		"queenblue":                      QueenBlue,
		"turkishrose":                    TurkishRose,
		"crimsonred":                     CrimsonRed,
		"seashell":                       Seashell,
		"spectra":                        Spectra,
		"upmaroon":                       UPMaroon,
		"mordantred":                     MordantRed,
		"orangewhite":                    OrangeWhite,
		"havelockblue":                   HavelockBlue,
		"sundance":                       Sundance,
		"treehouse":                      Treehouse,
		"rossocorsa":                     RossoCorsa,
		"palmleaf":                       PalmLeaf,
		"pastelgray":                     PastelGray,
		"tawnyport":                      TawnyPort,
		"vividorchid":                    VividOrchid,
		"jaguar":                         Jaguar,
		"dixie":                          Dixie,
		"frostedmint":                    FrostedMint,
		"rybblue":                        RYBBlue,
		"schooner":                       Schooner,
		"amaranthpurple":                 AmaranthPurple,
		"gainsboro":                      Gainsboro,
		"oldlace":                        OldLace,
		"emerald":                        Emerald,
		"congopink":                      CongoPink,
		"lemonyellow":                    LemonYellow,
		"maitai":                         MaiTai,
		"mediumruby":                     MediumRuby,
		"auchico":                        AuChico,
		"energyyellow":                   EnergyYellow,
		"granitegray":                    GraniteGray,
		"sherwoodgreen":                  SherwoodGreen,
		"drover":                         Drover,
		"internationalkleinblue":         InternationalKleinBlue,
		"shocking":                       Shocking,
		"algaegreen":                     AlgaeGreen,
		"potpourri":                      PotPourri,
		"robroy":                         RobRoy,
		"palemagenta":                    PaleMagenta,
		"caper":                          Caper,
		"dingley":                        Dingley,
		"ecru":                           Ecru,
		"matrix":                         Matrix,
		"tangaroa":                       Tangaroa,
		"bronco":                         Bronco,
		"frost":                          Frost,
		"kidnapper":                      Kidnapper,
		"redberry":                       RedBerry,
		"cornflowerblue":                 CornflowerBlue,
		"portgore":                       PortGore,
		"bluehaze":                       BlueHaze,
		"pantoneorange":                  PantoneOrange,
		"redoxide":                       RedOxide,
		"yellowgreen":                    YellowGreen,
		"logan":                          Logan,
		"saharasand":                     SaharaSand,
		"smokeytopaz":                    SmokeyTopaz,
		"tepapagreen":                    TePapaGreen,
		"daffodil":                       Daffodil,
		"limeade":                        Limeade,
		"metallicseaweed":                MetallicSeaweed,
		"silvertree":                     SilverTree,
		"turmeric":                       Turmeric,
		"vanilla":                        Vanilla,
		"confetti":                       Confetti,
		"malachite":                      Malachite,
		"crail":                          Crail,
		"kokoda":                         Kokoda,
		"quarterspanishwhite":            QuarterSpanishWhite,
		"reef":                           Reef,
		"shalimar":                       Shalimar,
		"englishholly":                   EnglishHolly,
		"pantoneblue":                    PantoneBlue,
		"swampgreen":                     SwampGreen,
		"lavenderindigo":                 LavenderIndigo,
		"smashedpumpkin":                 SmashedPumpkin,
		"cinder":                         Cinder,
		"lime":                           Lime,
		"observatory":                    Observatory,
		"pinklavender":                   PinkLavender,
		"chathamsblue":                   ChathamsBlue,
		"deepteal":                       DeepTeal,
		"fern":                           Fern,
		"greenyellow":                    GreenYellow,
		"gulfstream":                     GulfStream,
		"maize":                          Maize,
		"papayawhip":                     PapayaWhip,
		"safetyorange":                   SafetyOrange,
		"carminepink":                    CarminePink,
		"cinderella":                     Cinderella,
		"caribbeangreen":                 CaribbeanGreen,
		"christine":                      Christine,
		"iris":                           Iris,
		"tumbleweed":                     Tumbleweed,
		"darkchestnut":                   DarkChestnut,
		"dorado":                         Dorado,
		"lightpink":                      LightPink,
		"tuftbush":                       TuftBush,
		"brownyellow":                    BrownYellow,
		"firebrick":                      Firebrick,
		"metallicbronze":                 MetallicBronze,
		"peachschnapps":                  PeachSchnapps,
		"sepiablack":                     SepiaBlack,
		"clamshell":                      ClamShell,
		"springleaves":                   SpringLeaves,
		"violetred":                      VioletRed,
		"aquahaze":                       AquaHaze,
		"coriander":                      Coriander,
		"dandelion":                      Dandelion,
		"mexicanpink":                    MexicanPink,
		"richlilac":                      RichLilac,
		"tealdeer":                       TealDeer,
		"clover":                         Clover,
		"darkslategray":                  DarkSlateGray,
		"firefly":                        Firefly,
		"fuzzywuzzy":                     FuzzyWuzzy,
		"richgold":                       RichGold,
		"sanmarino":                      SanMarino,
		"victoria":                       Victoria,
		"carnationpink":                  CarnationPink,
		"bombay":                         Bombay,
		"citrus":                         Citrus,
		"congobrown":                     CongoBrown,
		"halfandhalf":                    HalfandHalf,
		"perutan":                        PeruTan,
		"romance":                        Romance,
		"bayleaf":                        BayLeaf,
		"ruddybrown":                     RuddyBrown,
		"islamicgreen":                   IslamicGreen,
		"paleredviolet":                  PaleRedViolet,
		"sorrellbrown":                   SorrellBrown,
		"geyser":                         Geyser,
		"rosebonbon":                     RoseBonbon,
		"spanishpink":                    SpanishPink,
		"ziggurat":                       Ziggurat,
		"lightcobaltblue":                LightCobaltBlue,
		"beeswax":                        Beeswax,
		"honeydew":                       Honeydew,
		"quartz":                         Quartz,
		"silver":                         Silver,
		"yourpink":                       YourPink,
		"amber":                          Amber,
		"aquaspring":                     AquaSpring,
		"darksalmon":                     DarkSalmon,
		"deeptaupe":                      DeepTaupe,
		"grayolive":                      GrayOlive,
		"alienarmpit":                    AlienArmpit,
		"whitepointer":                   WhitePointer,
		"crusta":                         Crusta,
		"tamarind":                       Tamarind,
		"burlywood":                      Burlywood,
		"granitegreen":                   GraniteGreen,
		"mediumelectricblue":             MediumElectricBlue,
		"pastelorange":                   PastelOrange,
		"deeptuscanred":                  DeepTuscanRed,
		"breakerbay":                     BreakerBay,
		"buttercup":                      Buttercup,
		"bostonuniversityred":            BostonUniversityRed,
		"laurel":                         Laurel,
		"pictonblue":                     PictonBlue,
		"strawberry":                     Strawberry,
		"tobaccobrown":                   TobaccoBrown,
		"beautybush":                     BeautyBush,
		"hippiepink":                     HippiePink,
		"brandeisblue":                   BrandeisBlue,
		"hampton":                        Hampton,
		"disco":                          Disco,
		"fadedjade":                      FadedJade,
		"ripelemon":                      RipeLemon,
		"albescentwhite":                 AlbescentWhite,
		"merino":                         Merino,
		"webchartreuse":                  WebChartreuse,
		"honeysuckle":                    Honeysuckle,
		"englishred":                     EnglishRed,
		"submarine":                      Submarine,
		"tanhide":                        TanHide,
		"eden":                           Eden,
		"deepchestnut":                   DeepChestnut,
		"msugreen":                       MSUGreen,
		"nutmegwoodfinish":               NutmegWoodFinish,
		"zinnwalditebrown":               ZinnwalditeBrown,
		"coquelicot":                     Coquelicot,
		"cordovan":                       Cordovan,
		"electriccrimson":                ElectricCrimson,
		"goldenfizz":                     GoldenFizz,
		"oracle":                         Oracle,
		"azuremist":                      AzureMist,
		"metallicsunburst":               MetallicSunburst,
		"mineralgreen":                   MineralGreen,
		"pantonepink":                    PantonePink,
		"heatwave":                       HeatWave,
		"frostee":                        Frostee,
		"revolver":                       Revolver,
		"brilliantazure":                 BrilliantAzure,
		"lavenderpurple":                 LavenderPurple,
		"pinkflare":                      PinkFlare,
		"cabsav":                         CabSav,
		"bluediamond":                    BlueDiamond,
		"cerise":                         Cerise,
		"grape":                          Grape,
		"kimberly":                       Kimberly,
		"saltbox":                        SaltBox,
		"sanfelix":                       SanFelix,
		"seagreen":                       SeaGreen,
		"berylgreen":                     BerylGreen,
		"verdigris":                      Verdigris,
		"wine":                           Wine,
		"silverpink":                     SilverPink,
		"fashionfuchsia":                 FashionFuchsia,
		"hintofred":                      HintofRed,
		"jambalaya":                      Jambalaya,
		"deepsea":                        DeepSea,
		"viridian":                       Viridian,
		"vividcrimson":                   VividCrimson,
		"lightkhaki":                     LightKhaki,
		"congressblue":                   CongressBlue,
		"gallery":                        Gallery,
		"vermilion":                      Vermilion,
		"violentviolet":                  ViolentViolet,
		"ashgrey":                        AshGrey,
		"gigas":                          Gigas,
		"goldenbell":                     GoldenBell,
		"sonicsilver":                    SonicSilver,
		"darkblue":                       DarkBlue,
		"darkpink":                       DarkPink,
		"thulianpink":                    ThulianPink,
		"vividviolet":                    VividViolet,
		"wasabi":                         Wasabi,
		"boogerbuster":                   BoogerBuster,
		"copper":                         Copper,
		"darksienna":                     DarkSienna,
		"deepgreen":                      DeepGreen,
		"kombugreen":                     KombuGreen,
		"moonstoneblue":                  MoonstoneBlue,
		"tana":                           Tana,
		"chamoisee":                      Chamoisee,
		"cardinalpink":                   CardinalPink,
		"chestnut":                       Chestnut,
		"coldpurple":                     ColdPurple,
		"hairyheath":                     HairyHeath,
		"korma":                          Korma,
		"perano":                         Perano,
		"pink":                           Pink,
		"beaver":                         Beaver,
		"rust":                           Rust,
		"northtexasgreen":                NorthTexasGreen,
		"periwinklegray":                 PeriwinkleGray,
		"plumppurple":                    PlumpPurple,
		"temptress":                      Temptress,
		"goben":                          GoBen,
		"sundown":                        Sundown,
		"ufogreen":                       UFOGreen,
		"smaltblue":                      SmaltBlue,
		"cannonpink":                     CannonPink,
		"greenpea":                       GreenPea,
		"hotmagenta":                     HotMagenta,
		"operamauve":                     OperaMauve,
		"saddle":                         Saddle,
		"sprout":                         Sprout,
		"tuscantan":                      TuscanTan,
		"buccaneer":                      Buccaneer,
		"lapalma":                        LaPalma,
		"razzmatazz":                     Razzmatazz,
		"tiamaria":                       TiaMaria,
		"desertstorm":                    DesertStorm,
		"desertsand":                     DesertSand,
		"kabul":                          Kabul,
		"lightbrown":                     LightBrown,
		"pearllusta":                     PearlLusta,
		"scarletgum":                     ScarletGum,
		"sepia":                          Sepia,
		"skobeloff":                      Skobeloff,
		"daisybush":                      DaisyBush,
		"snuff":                          Snuff,
		"feta":                           Feta,
		"romansilver":                    RomanSilver,
		"ebony":                          Ebony,
		"frenchpass":                     FrenchPass,
		"capepalliser":                   CapePalliser,
		"funblue":                        FunBlue,
		"lemoncurry":                     LemonCurry,
		"salomie":                        Salomie,
		"springfrost":                    SpringFrost,
		"flavescent":                     Flavescent,
		"crowshead":                      Crowshead,
		"gravel":                         Gravel,
		"hitpink":                        HitPink,
		"lemonginger":                    LemonGinger,
		"verylightblue":                  VeryLightBlue,
		"bluebell":                       BlueBell,
		"elephant":                       Elephant,
		"jagger":                         Jagger,
		"oysterbay":                      OysterBay,
		"persimmon":                      Persimmon,
		"springbud":                      SpringBud,
		"coralred":                       CoralRed,
		"limedash":                       LimedAsh,
		"wintergreendream":               WintergreenDream,
		"lilacbush":                      LilacBush,
		"lighthotpink":                   LightHotPink,
		"processcyan":                    ProcessCyan,
		"tosca":                          Tosca,
		"turquoisegreen":                 TurquoiseGreen,
		"hintofgreen":                    HintofGreen,
		"darkviolet":                     DarkViolet,
		"governorbay":                    GovernorBay,
		"lightyellow":                    LightYellow,
		"nyanza":                         Nyanza,
		"pearl":                          Pearl,
		"sulu":                           Sulu,
		"darkbrown":                      DarkBrown,
		"steelteal":                      SteelTeal,
		"gogreen":                        GOGreen,
		"mediumredviolet":                MediumRedViolet,
		"seapink":                        SeaPink,
		"snowdrift":                      SnowDrift,
		"spindle":                        Spindle,
		"vidaloca":                       VidaLoca,
		"airforceblue":                   AirForceBlue,
		"deyork":                         DeYork,
		"greenvogue":                     GreenVogue,
		"manz":                           Manz,
		"neptune":                        Neptune,
		"shimmeringblush":                ShimmeringBlush,
		"stormgray":                      StormGray,
		"wildrice":                       WildRice,
		"budgreen":                       BudGreen,
		"brilliantrose":                  BrilliantRose,
		"columbiablue":                   ColumbiaBlue,
		"jade":                           Jade,
		"mediumaquamarine":               MediumAquamarine,
		"shadylady":                      ShadyLady,
		"botticelli":                     Botticelli,
		"brightgray":                     BrightGray,
		"christi":                        Christi,
		"darkturquoise":                  DarkTurquoise,
		"electricviolet":                 ElectricViolet,
		"mindaro":                        Mindaro,
		"ottoman":                        Ottoman,
		"tacao":                          Tacao,
		"blackhaze":                      BlackHaze,
		"tenne":                          Tenne,
		"earthyellow":                    EarthYellow,
		"eastbay":                        EastBay,
		"eternity":                       Eternity,
		"quinacridonemagenta":            QuinacridoneMagenta,
		"ruber":                          Ruber,
		"tropicalrainforest":             TropicalRainForest,
		"zambezi":                        Zambezi,
		"desert":                         Desert,
		"makara":                         Makara,
		"fuchsiarose":                    FuchsiaRose,
		"heliotropemagenta":              HeliotropeMagenta,
		"mintcream":                      MintCream,
		"sapphireblue":                   SapphireBlue,
		"bakermillerpink":                BakerMillerPink,
		"copperrose":                     CopperRose,
		"darkmagenta":                    DarkMagenta,
		"lightgray":                      LightGray,
		"palepink":                       PalePink,
		"spanishskyblue":                 SpanishSkyBlue,
		"brownderby":                     BrownDerby,
		"stormcloud":                     Stormcloud,
		"sandal":                         Sandal,
		"greenkelp":                      GreenKelp,
		"indochine":                      Indochine,
		"tealgreen":                      TealGreen,
		"bridalheath":                    BridalHeath,
		"coyotebrown":                    CoyoteBrown,
		"deepcarmine":                    DeepCarmine,
		"gumbo":                          Gumbo,
		"hanpurple":                      HanPurple,
		"mediumjunglegreen":              MediumJungleGreen,
		"peachpuff":                      PeachPuff,
		"chromewhite":                    ChromeWhite,
		"etonblue":                       EtonBlue,
		"mediumskyblue":                  MediumSkyBlue,
		"shakespeare":                    Shakespeare,
		"deeppink":                       DeepPink,
		"cosmiclatte":                    CosmicLatte,
		"fedora":                         Fedora,
		"flushmahogany":                  FlushMahogany,
		"glossygrape":                    GlossyGrape,
		"heather":                        Heather,
		"pearlbush":                      PearlBush,
		"saratoga":                       Saratoga,
		"bluecharcoal":                   BlueCharcoal,
		"wedgewood":                      Wedgewood,
		"coolblack":                      CoolBlack,
		"nepal":                          Nepal,
		"paletaupe":                      PaleTaupe,
		"watercourse":                    Watercourse,
		"barossa":                        Barossa,
		"verypaleyellow":                 VeryPaleYellow,
		"vividburgundy":                  VividBurgundy,
		"ironsidegray":                   IronsideGray,
		"cherrypie":                      CherryPie,
		"greenhaze":                      GreenHaze,
		"zircon":                         Zircon,
		"blackleatherjacket":             BlackLeatherJacket,
		"bluejeans":                      BlueJeans,
		"burntsienna":                    BurntSienna,
		"darkraspberry":                  DarkRaspberry,
		"sapgreen":                       SapGreen,
		"zorba":                          Zorba,
		"aluminium":                      Aluminium,
		"marshland":                      Marshland,
		"mauvelous":                      Mauvelous,
		"olivedrabseven":                 OliveDrabSeven,
		"pohutukawa":                     Pohutukawa,
		"regalia":                        Regalia,
		"ronchi":                         Ronchi,
		"blackberry":                     Blackberry,
		"pearlypurple":                   PearlyPurple,
		"persianplum":                    PersianPlum,
		"richlavender":                   RichLavender,
		"sambuca":                        Sambuca,
		"soap":                           Soap,
		"sugarplum":                      SugarPlum,
		"tigerseye":                      TigersEye,
		"gargoylegas":                    GargoyleGas,
		"navy":                           Navy,
		"mughalgreen":                    MughalGreen,
		"wisteria":                       Wisteria,
		"woodland":                       Woodland,
		"brass":                          Brass,
		"fuchsiablue":                    FuchsiaBlue,
		"lavenderpink":                   LavenderPink,
		"lighttaupe":                     LightTaupe,
		"plum":                           Plum,
		"cumulus":                        Cumulus,
		"fuchsiapurple":                  FuchsiaPurple,
		"fog":                            Fog,
		"coolgrey":                       CoolGrey,
		"newcar":                         NewCar,
		"pigpink":                        PigPink,
		"softpeach":                      SoftPeach,
		"suvagray":                       SuvaGray,
		"bilobaflower":                   BilobaFlower,
		"casal":                          Casal,
		"chinesered":                     ChineseRed,
		"greenleaf":                      GreenLeaf,
		"hoki":                           Hoki,
		"oldrose":                        OldRose,
		"buddhagold":                     BuddhaGold,
		"bunting":                        Bunting,
		"indianred":                      IndianRed,
		"moodyblue":                      MoodyBlue,
		"mountainmist":                   MountainMist,
		"bitter":                         Bitter,
		"darkebony":                      DarkEbony,
		"halayàúbe":                      HalayàÚbe,
		"masala":                         Masala,
		"pueblo":                         Pueblo,
		"rhythm":                         Rhythm,
		"springsun":                      SpringSun,
		"atoll":                          Atoll,
		"indigo":                         Indigo,
		"serenade":                       Serenade,
		"ballblue":                       BallBlue,
		"liver":                          Liver,
		"pinkflamingo":                   PinkFlamingo,
		"caramel":                        Caramel,
		"magnolia":                       Magnolia,
		"bamboo":                         Bamboo,
		"deepjunglegreen":                DeepJungleGreen,
		"downriver":                      Downriver,
		"lasallegreen":                   LaSalleGreen,
		"orchidwhite":                    OrchidWhite,
		"reddevil":                       RedDevil,
		"tartorange":                     TartOrange,
		"aquasqueeze":                    AquaSqueeze,
		"horsesneck":                     HorsesNeck,
		"peppermint":                     Peppermint,
		"hillary":                        Hillary,
		"moccaccino":                     Moccaccino,
		"redsalsa":                       RedSalsa,
		"hokeypokey":                     HokeyPokey,
		"dingydungeon":                   DingyDungeon,
		"pixiepowder":                    PixiePowder,
		"riverbed":                       RiverBed,
		"claycreek":                      ClayCreek,
		"pigmentred":                     PigmentRed,
		"sheengreen":                     SheenGreen,
		"siren":                          Siren,
		"texas":                          Texas,
		"japanesecarmine":                JapaneseCarmine,
		"buttermilk":                     Buttermilk,
		"mangotango":                     MangoTango,
		"vividmulberry":                  VividMulberry,
		"bordeaux":                       Bordeaux,
		"harvestgold":                    HarvestGold,
		"raspberrypink":                  RaspberryPink,
		"stardust":                       StarDust,
		"wisppink":                       WispPink,
		"cadmiumyellow":                  CadmiumYellow,
		"cadetgrey":                      CadetGrey,
		"lavender":                       Lavender,
		"ash":                            Ash,
		"fiord":                          Fiord,
		"macaroniandcheese":              MacaroniAndCheese,
		"persianred":                     PersianRed,
		"rustyred":                       RustyRed,
		"sunset":                         Sunset,
		"vividyellow":                    VividYellow,
		"brilliantlavender":              BrilliantLavender,
		"frenchbistre":                   FrenchBistre,
		"marzipan":                       Marzipan,
		"spice":                          Spice,
		"brightred":                      BrightRed,
		"rumswizzle":                     RumSwizzle,
		"minionyellow":                   MinionYellow,
		"zest":                           Zest,
		"electricyellow":                 ElectricYellow,
		"darkmediumgray":                 DarkMediumGray,
		"pablo":                          Pablo,
		"springrain":                     SpringRain,
		"stormdust":                      StormDust,
		"swamp":                          Swamp,
		"cottoncandy":                    CottonCandy,
		"milanored":                      MilanoRed,
		"spicymustard":                   SpicyMustard,
		"turtlegreen":                    TurtleGreen,
		"japaneseindigo":                 JapaneseIndigo,
		"killarney":                      Killarney,
		"twilightlavender":               TwilightLavender,
		"verylightazure":                 VeryLightAzure,
		"darkpastelpurple":               DarkPastelPurple,
		"mediumturquoise":                MediumTurquoise,
		"naturalgray":                    NaturalGray,
		"navajowhite":                    NavajoWhite,
		"queenpink":                      QueenPink,
		"cocoabrown":                     CocoaBrown,
		"cavernpink":                     CavernPink,
		"forestgreen":                    ForestGreen,
		"mardigras":                      MardiGras,
		"mexicanred":                     MexicanRed,
		"wafer":                          Wafer,
		"bronze":                         Bronze,
		"lightseagreen":                  LightSeaGreen,
		"opal":                           Opal,
		"holly":                          Holly,
		"deepblush":                      DeepBlush,
		"gray":                           Gray,
		"lightgreen":                     LightGreen,
		"permanentgeraniumlake":          PermanentGeraniumLake,
		"plantation":                     Plantation,
		"thistlegreen":                   ThistleGreen,
		"bigdiporuby":                    BigDipOruby,
		"concord":                        Concord,
		"frenchlilac":                    FrenchLilac,
		"lightpastelpurple":              LightPastelPurple,
		"liverchestnut":                  LiverChestnut,
		"newyorkpink":                    NewYorkPink,
		"pacifika":                       Pacifika,
		"scarlett":                       Scarlett,
		"amour":                          Amour,
		"tropicalviolet":                 TropicalViolet,
		"dogs":                           Dogs,
		"greenlizard":                    GreenLizard,
		"harvardcrimson":                 HarvardCrimson,
		"mahogany":                       Mahogany,
		"pinecone":                       PineCone,
		"tara":                           Tara,
		"ube":                            Ube,
		"brownsugar":                     BrownSugar,
		"blazeorange":                    BlazeOrange,
		"cosmiccobalt":                   CosmicCobalt,
		"fuchsia":                        Fuchsia,
		"nileblue":                       NileBlue,
		"wepeep":                         WePeep,
		"bahia":                          Bahia,
		"lividbrown":                     LividBrown,
		"onion":                          Onion,
		"scarpaflow":                     ScarpaFlow,
		"spray":                          Spray,
		"waikawagray":                    WaikawaGray,
		"cardinal":                       Cardinal,
		"gambogeorange":                  GambogeOrange,
		"oldlavender":                    OldLavender,
		"pickledbean":                    PickledBean,
		"rajah":                          Rajah,
		"sandstone":                      Sandstone,
		"steelpink":                      SteelPink,
		"swisscoffee":                    SwissCoffee,
		"corn":                           Corn,
		"woodsmoke":                      Woodsmoke,
		"carmine":                        Carmine,
		"green":                          Green,
		"greenmist":                      GreenMist,
		"magenta":                        Magenta,
		"oceanboatblue":                  OceanBoatBlue,
		"selago":                         Selago,
		"canary":                         Canary,
		"cyancobaltblue":                 CyanCobaltBlue,
		"mocha":                          Mocha,
		"sanguinebrown":                  SanguineBrown,
		"smokyblack":                     SmokyBlack,
		"stratos":                        Stratos,
		"burntmaroon":                    BurntMaroon,
		"catawba":                        Catawba,
		"chamois":                        Chamois,
		"zeus":                           Zeus,
		"bluestone":                      BlueStone,
		"chileanheath":                   ChileanHeath,
		"dell":                           Dell,
		"halfbaked":                      HalfBaked,
		"pinkswan":                       PinkSwan,
		"bluemarguerite":                 BlueMarguerite,
		"darkyellow":                     DarkYellow,
		"dolly":                          Dolly,
		"ferra":                          Ferra,
		"chelseacucumber":                ChelseaCucumber,
		"grannysmith":                    GrannySmith,
		"highland":                       Highland,
		"mantis":                         Mantis,
		"mikado":                         Mikado,
		"shampoo":                        Shampoo,
		"valentino":                      Valentino,
		"wheatfield":                     Wheatfield,
		"barleywhite":                    BarleyWhite,
		"ginfizz":                        GinFizz,
		"halfspanishwhite":               HalfSpanishWhite,
		"monarch":                        Monarch,
		"oldbrick":                       OldBrick,
		"sandrift":                       Sandrift,
		"cello":                          Cello,
		"electricindigo":                 ElectricIndigo,
		"fantasy":                        Fantasy,
		"pippin":                         Pippin,
		"scampi":                         Scampi,
		"sealbrown":                      SealBrown,
		"butterywhite":                   ButteryWhite,
		"deepmagenta":                    DeepMagenta,
		"desire":                         Desire,
		"regentgray":                     RegentGray,
		"cosmos":                         Cosmos,
		"bisque":                         Bisque,
		"burntumber":                     BurntUmber,
		"darklavender":                   DarkLavender,
		"pavlova":                        Pavlova,
		"shadowgreen":                    ShadowGreen,
		"antiquefuchsia":                 AntiqueFuchsia,
		"dutchwhite":                     DutchWhite,
		"greensmoke":                     GreenSmoke,
		"java":                           Java,
		"quillgray":                      QuillGray,
		"rybred":                         RYBRed,
		"downy":                          Downy,
		"midnightmoss":                   MidnightMoss,
		"white":                          White,
		"cuttysark":                      CuttySark,
		"asparagus":                      Asparagus,
		"mulledwine":                     MulledWine,
		"napiergreen":                    NapierGreen,
		"richbrilliantlavender":          RichBrilliantLavender,
		"stack":                          Stack,
		"airsuperiorityblue":             AirSuperiorityBlue,
		"genoa":                          Genoa,
		"manatee":                        Manatee,
		"pumpkin":                        Pumpkin,
		"purpureus":                      Purpureus,
		"sandwisp":                       Sandwisp,
		"chartreuse":                     Chartreuse,
		"chileanfire":                    ChileanFire,
		"darkpastelgreen":                DarkPastelGreen,
		"icterine":                       Icterine,
		"pizazz":                         Pizazz,
		"teal":                           Teal,
		"blackbean":                      BlackBean,
		"comet":                          Comet,
		"ecstasy":                        Ecstasy,
		"lightcrimson":                   LightCrimson,
		"cobaltblue":                     CobaltBlue,
		"greenwaterloo":                  GreenWaterloo,
		"malta":                          Malta,
		"riflegreen":                     RifleGreen,
		"burntorange":                    BurntOrange,
		"rustynail":                      RustyNail,
		"carnabytan":                     CarnabyTan,
		"fulvous":                        Fulvous,
		"irishcoffee":                    IrishCoffee,
		"irresistible":                   Irresistible,
		"islandspice":                    IslandSpice,
		"lightorchid":                    LightOrchid,
		"tangerineyellow":                TangerineYellow,
		"tearose":                        TeaRose,
		"bubbles":                        Bubbles,
		"fire":                           Fire,
		"lavendermist":                   LavenderMist,
		"metallicgold":                   MetallicGold,
		"amethystsmoke":                  AmethystSmoke,
		"sanddune":                       SandDune,
		"brandyrose":                     BrandyRose,
		"hintofyellow":                   HintofYellow,
		"laser":                          Laser,
		"porcelain":                      Porcelain,
		"charade":                        Charade,
		"bilbao":                         Bilbao,
		"cambridgeblue":                  CambridgeBlue,
		"citrine":                        Citrine,
		"gunpowder":                      GunPowder,
		"peach":                          Peach,
		"aztecgold":                      AztecGold,
		"lightblue":                      LightBlue,
		"powderblue":                     PowderBlue,
		"rockspray":                      RockSpray,
		"rosefog":                        RoseFog,
		"fruitsalad":                     FruitSalad,
		"orangered":                      OrangeRed,
		"rebel":                          Rebel,
		"como":                           Como,
		"silversand":                     SilverSand,
		"amaranthpink":                   AmaranthPink,
		"ceramic":                        Ceramic,
		"dune":                           Dune,
		"purplepizzazz":                  PurplePizzazz,
		"saltpan":                        Saltpan,
		"squirrel":                       Squirrel,
		"wildsand":                       WildSand,
		"apple":                          Apple,
		"bubblegum":                      BubbleGum,
		"dawnpink":                       DawnPink,
		"ochre":                          Ochre,
		"yankeesblue":                    YankeesBlue,
		"bismark":                        Bismark,
		"rocketmetallic":                 RocketMetallic,
		"turbo":                          Turbo,
		"creambrulee":                    CreamBrulee,
		"gurkha":                         Gurkha,
		"mirage":                         Mirage,
		"flamingopink":                   FlamingoPink,
		"ghostwhite":                     GhostWhite,
		"metalpink":                      MetalPink,
		"palerobineggblue":               PaleRobinEggBlue,
		"redribbon":                      RedRibbon,
		"fuego":                          Fuego,
		"jacksonspurple":                 JacksonsPurple,
		"raspberry":                      Raspberry,
		"whiterock":                      WhiteRock,
		"darkbrowntangelo":               DarkBrownTangelo,
		"champagne":                      Champagne,
		"lochmara":                       Lochmara,
		"mandarin":                       Mandarin,
		"africanviolet":                  AfricanViolet,
		"denimblue":                      DenimBlue,
		"barleycorn":                     BarleyCorn,
		"moonglow":                       MoonGlow,
		"daintree":                       Daintree,
		"jackobean":                      JackoBean,
		"palesilver":                     PaleSilver,
		"sugarcane":                      SugarCane,
		"chalky":                         Chalky,
		"aurometalsaurus":                AuroMetalSaurus,
		"ceruleanblue":                   CeruleanBlue,
		"cruise":                         Cruise,
		"cyan":                           Cyan,
		"delrio":                         DelRio,
		"falured":                        FaluRed,
		"pigeonpost":                     PigeonPost,
		"antiquebrass":                   AntiqueBrass,
		"screamingreen":                  ScreaminGreen,
		"puce":                           Puce,
		"meteorite":                      Meteorite,
		"bluemagentaviolet":              BlueMagentaViolet,
		"blackshadows":                   BlackShadows,
		"brightlavender":                 BrightLavender,
		"cowboy":                         Cowboy,
		"curiousblue":                    CuriousBlue,
		"sweetcorn":                      SweetCorn,
		"artichoke":                      Artichoke,
		"earlydawn":                      EarlyDawn,
		"eggsour":                        EggSour,
		"eveningsea":                     EveningSea,
		"larioja":                        LaRioja,
		"purplenavy":                     PurpleNavy,
		"dairycream":                     DairyCream,
		"carouselpink":                   CarouselPink,
		"cork":                           Cork,
		"frenchfuchsia":                  FrenchFuchsia,
		"gablegreen":                     GableGreen,
		"lightfrenchbeige":               LightFrenchBeige,
		"mountainmeadow":                 MountainMeadow,
		"nonphotoblue":                   NonPhotoBlue,
		"bananamania":                    BananaMania,
		"maverick":                       Maverick,
		"mediumspringgreen":              MediumSpringGreen,
		"sienna":                         Sienna,
		"timberwolf":                     Timberwolf,
		"twilight":                       Twilight,
		"calico":                         Calico,
		"graysuit":                       GraySuit,
		"palespringbud":                  PaleSpringBud,
		"razzmicberry":                   RazzmicBerry,
		"redstage":                       RedStage,
		"sangria":                        Sangria,
		"wildblueyonder":                 WildBlueYonder,
		"candyapplered":                  CandyAppleRed,
		"imperialblue":                   ImperialBlue,
		"saddlebrown":                    SaddleBrown,
		"tolopea":                        Tolopea,
		"willowgrove":                    WillowGrove,
		"crusoe":                         Crusoe,
		"chinarose":                      ChinaRose,
		"laspalmas":                      LasPalmas,
		"pear":                           Pear,
		"pullmangreen":                   PullmanGreen,
		"regalblue":                      RegalBlue,
		"riptide":                        Riptide,
		"royalblue":                      RoyalBlue,
		"charlestongreen":                CharlestonGreen,
		"deluge":                         Deluge,
		"narvik":                         Narvik,
		"orientalpink":                   OrientalPink,
		"clayash":                        ClayAsh,
		"darkfern":                       DarkFern,
		"pattensblue":                    PattensBlue,
		"taupegray":                      TaupeGray,
		"avocado":                        Avocado,
		"seabuckthorn":                   SeaBuckthorn,
		"sunglo":                         Sunglo,
		"uablue":                         UABlue,
		"midgray":                        MidGray,
		"greenhouse":                     GreenHouse,
		"magentahaze":                    MagentaHaze,
		"nugget":                         Nugget,
		"pakistangreen":                  PakistanGreen,
		"sunshade":                       Sunshade,
		"vividtangerine":                 VividTangerine,
		"foam":                           Foam,
		"birch":                          Birch,
		"eucalyptus":                     Eucalyptus,
		"lightturquoise":                 LightTurquoise,
		"maximumblue":                    MaximumBlue,
		"oceanblue":                      OceanBlue,
		"whiskey":                        Whiskey,
		"alabaster":                      Alabaster,
		"royalheath":                     RoyalHeath,
		"sasquatchsocks":                 SasquatchSocks,
		"neoncarrot":                     NeonCarrot,
		"sirocco":                        Sirocco,
		"tequila":                        Tequila,
		"windsor":                        Windsor,
		"loblolly":                       Loblolly,
		"deepcerise":                     DeepCerise,
		"deepviolet":                     DeepViolet,
		"empress":                        Empress,
		"orange":                         Orange,
		"russett":                        Russett,
		"willowbrook":                    WillowBrook,
		"aliceblue":                      AliceBlue,
		"cement":                         Cement,
		"punch":                          Punch,
		"wildwillow":                     WildWillow,
		"cascade":                        Cascade,
		"brightgreen":                    BrightGreen,
		"darkmidnightblue":               DarkMidnightBlue,
		"deeplilac":                      DeepLilac,
		"gulfblue":                       GulfBlue,
		"keylime":                        KeyLime,
		"lightslategray":                 LightSlateGray,
		"swansdown":                      SwansDown,
		"boulder":                        Boulder,
		"lemonchiffon":                   LemonChiffon,
		"mistgray":                       MistGray,
		"sttropaz":                       StTropaz,
		"brightcerulean":                 BrightCerulean,
		"oxley":                          Oxley,
		"royalfuchsia":                   RoyalFuchsia,
		"astronautblue":                  AstronautBlue,
		"charlotte":                      Charlotte,
		"ebb":                            Ebb,
		"saffron":                        Saffron,
		"seablue":                        SeaBlue,
		"tallow":                         Tallow,
		"auburn":                         Auburn,
		"paleviolet":                     PaleViolet,
		"coraltree":                      CoralTree,
		"gladegreen":                     GladeGreen,
		"gumleaf":                        GumLeaf,
		"heatheredgray":                  HeatheredGray,
		"jaffa":                          Jaffa,
		"kobi":                           Kobi,
		"nomad":                          Nomad,
		"viridiangreen":                  ViridianGreen,
		"burnham":                        Burnham,
		"fieryorange":                    FieryOrange,
		"giantsclub":                     GiantsClub,
		"oslogray":                       OsloGray,
		"prim":                           Prim,
		"scandal":                        Scandal,
		"aztec":                          Aztec,
		"cerulean":                       Cerulean,
		"sauvignon":                      Sauvignon,
		"bluechalk":                      BlueChalk,
		"bullshot":                       BullShot,
		"derby":                          Derby,
		"floralwhite":                    FloralWhite,
		"giantsorange":                   GiantsOrange,
		"munsellyellow":                  MunsellYellow,
		"tulip":                          Tulip,
		"arapawa":                        Arapawa,
		"cornharvest":                    CornHarvest,
		"deepbronze":                     DeepBronze,
		"hitgray":                        HitGray,
		"lemonglacier":                   LemonGlacier,
		"melanie":                        Melanie,
		"mulberrywood":                   MulberryWood,
		"mysticmaroon":                   MysticMaroon,
		"blackmarlin":                    BlackMarlin,
		"tranquil":                       Tranquil,
		"pistachio":                      Pistachio,
		"indianyellow":                   IndianYellow,
		"littleboyblue":                  LittleBoyBlue,
		"oxfordblue":                     OxfordBlue,
		"purple":                         Purple,
		"crete":                          Crete,
		"lightbrilliantred":              LightBrilliantRed,
		"shipgray":                       ShipGray,
		"tahunasands":                    TahunaSands,
		"titanwhite":                     TitanWhite,
		"usafablue":                      USAFABlue,
		"cafenoir":                       CafeNoir,
		"darkolivegreen":                 DarkOliveGreen,
		"foggygray":                      FoggyGray,
		"hurricane":                      Hurricane,
		"palemagentapink":                PaleMagentaPink,
		"prairiesand":                    PrairieSand,
		"chenin":                         Chenin,
		"pomegranate":                    Pomegranate,
		"softamber":                      SoftAmber,
		"x11gray":                        X11Gray,
		"bluezodiac":                     BlueZodiac,
		"mediumspringbud":                MediumSpringBud,
		"portica":                        Portica,
		"chlorophyllgreen":               ChlorophyllGreen,
		"tahitigold":                     TahitiGold,
		"unbleachedsilk":                 UnbleachedSilk,
		"copperpenny":                    CopperPenny,
		"brinkpink":                      BrinkPink,
		"blackcurrant":                   Blackcurrant,
		"bridesmaid":                     Bridesmaid,
		"cardingreen":                    CardinGreen,
		"danube":                         Danube,
		"mossgreen":                      MossGreen,
		"raffia":                         Raffia,
		"santafe":                        SantaFe,
		"vividorangepeel":                VividOrangePeel,
		"allports":                       Allports,
		"eaglegreen":                     EagleGreen,
		"forgetmenot":                    ForgetMeNot,
		"ironstone":                      Ironstone,
		"shipcove":                       ShipCove,
		"antiqueruby":                    AntiqueRuby,
		"cornsilk":                       Cornsilk,
		"cream":                          Cream,
		"darkcoral":                      DarkCoral,
		"tradewind":                      Tradewind,
		"veronica":                       Veronica,
		"wistful":                        Wistful,
		"bizarre":                        Bizarre,
		"moonmist":                       MoonMist,
		"schoolbusyellow":                SchoolBusYellow,
		"seaweed":                        Seaweed,
		"tomthumb":                       TomThumb,
		"lightsalmon":                    LightSalmon,
		"junebud":                        JuneBud,
		"tuna":                           Tuna,
		"aquamarine":                     Aquamarine,
		"cherrywood":                     Cherrywood,
		"darkmossgreen":                  DarkMossGreen,
		"nandor":                         Nandor,
		"parsley":                        Parsley,
		"amazon":                         Amazon,
		"locust":                         Locust,
		"pastelviolet":                   PastelViolet,
		"sazerac":                        Sazerac,
		"gamboge":                        Gamboge,
		"kenyancopper":                   KenyanCopper,
		"ultramarine":                    Ultramarine,
		"cannonblack":                    CannonBlack,
		"flamenco":                       Flamenco,
		"graphite":                       Graphite,
		"outrageousorange":               OutrageousOrange,
		"raisinblack":                    RaisinBlack,
		"sahara":                         Sahara,
		"tuliptree":                      TulipTree,
		"verdungreen":                    VerdunGreen,
		"aero":                           Aero,
		"gothic":                         Gothic,
		"richmaroon":                     RichMaroon,
		"eagle":                          Eagle,
		"domino":                         Domino,
		"gullgray":                       GullGray,
		"palechestnut":                   PaleChestnut,
		"whitelilac":                     WhiteLilac,
		"bronzeolive":                    BronzeOlive,
		"seanymph":                       SeaNymph,
		"azure":                          Azure,
		"tangerine":                      Tangerine,
		"claret":                         Claret,
		"heliotrope":                     Heliotrope,
		"brightube":                      BrightUbe,
		"palecerulean":                   PaleCerulean,
		"pigmentblue":                    PigmentBlue,
		"winterwizard":                   WinterWizard,
		"honeyflower":                    HoneyFlower,
		"eerieblack":                     EerieBlack,
		"mabel":                          Mabel,
		"periglacialblue":                PeriglacialBlue,
		"tide":                           Tide,
		"ebonyclay":                      EbonyClay,
		"chaletgreen":                    ChaletGreen,
		"lightningyellow":                LightningYellow,
		"mystic":                         Mystic,
		"bigfootfeet":                    BigFootFeet,
		"donjuan":                        DonJuan,
		"merlin":                         Merlin,
		"palegoldenrod":                  PaleGoldenrod,
		"bluelagoon":                     BlueLagoon,
		"calpolygreen":                   CalPolyGreen,
		"lucky":                          Lucky,
		"mandyspink":                     MandysPink,
		"woodrush":                       Woodrush,
		"astronaut":                      Astronaut,
		"lightgrayishmagenta":            LightGrayishMagenta,
		"mongoose":                       Mongoose,
		"moonraker":                      MoonRaker,
		"naplesyellow":                   NaplesYellow,
		"tabasco":                        Tabasco,
		"amaranthred":                    AmaranthRed,
		"doublecolonialwhite":            DoubleColonialWhite,
		"khaki":                          Khaki,
		"mayablue":                       MayaBlue,
		"sinopia":                        Sinopia,
		"deepcove":                       DeepCove,
		"mediumcandyapplered":            MediumCandyAppleRed,
		"mintjulep":                      MintJulep,
		"paleprim":                       PalePrim,
		"dartmouthgreen":                 DartmouthGreen,
		"lightmossgreen":                 LightMossGreen,
		"eastside":                       EastSide,
		"friargray":                      FriarGray,
		"trueblue":                       TrueBlue,
		"blacksqueeze":                   BlackSqueeze,
		"pastelblue":                     PastelBlue,
		"saeeceamber":                    SAEECEAmber,
		"spanishred":                     SpanishRed,
		"blush":                          Blush,
		"golddrop":                       GoldDrop,
		"persianblue":                    PersianBlue,
		"persianorange":                  PersianOrange,
		"redorange":                      RedOrange,
		"geraldine":                      Geraldine,
		"cadetblue":                      CadetBlue,
		"codgray":                        CodGray,
		"hopbush":                        Hopbush,
		"vividtangelo":                   VividTangelo,
		"bluechill":                      BlueChill,
		"costadelsol":                    CostaDelSol,
		"rubinered":                      RubineRed,
		"sinbad":                         Sinbad,
		"xanadu":                         Xanadu,
		"zanah":                          Zanah,
		"ceruleanfrost":                  CeruleanFrost,
		"willpowerorange":                WillpowerOrange,
		"bleachedcedar":                  BleachedCedar,
		"jetstream":                      JetStream,
		"magicmint":                      MagicMint,
		"pumice":                         Pumice,
		"deer":                           Deer,
		"kingfisherdaisy":                KingfisherDaisy,
		"skyblue":                        SkyBlue,
		"trout":                          Trout,
		"weborange":                      WebOrange,
		"frenchplum":                     FrenchPlum,
		"freshair":                       FreshAir,
		"blackrussian":                   BlackRussian,
		"chinook":                        Chinook,
		"fresheggplant":                  FreshEggplant,
		"gorse":                          Gorse,
		"justright":                      JustRight,
		"mauve":                          Mauve,
		"oysterpink":                     OysterPink,
		"scotchmist":                     ScotchMist,
		"celery":                         Celery,
		"tuftsblue":                      TuftsBlue,
		"telemagenta":                    Telemagenta,
		"calypso":                        Calypso,
		"junglegreen":                    JungleGreen,
		"nightrider":                     NightRider,
		"burnishedbrown":                 BurnishedBrown,
		"zuccini":                        Zuccini,
		"pumpkinskin":                    PumpkinSkin,
		"ecruwhite":                      EcruWhite,
		"fireenginered":                  FireEngineRed,
		"goldengatebridge":               GoldenGateBridge,
		"hacienda":                       Hacienda,
		"viking":                         Viking,
		"vividredtangelo":                VividRedTangelo,
		"byzantium":                      Byzantium,
		"deeplemon":                      DeepLemon,
		"blackrock":                      BlackRock,
		"paradiso":                       Paradiso,
		"sapling":                        Sapling,
		"orchidpink":                     OrchidPink,
		"dustygray":                      DustyGray,
		"indiantan":                      IndianTan,
		"redbeech":                       RedBeech,
		"tutu":                           Tutu,
		"darkspringgreen":                DarkSpringGreen,
		"clementine":                     Clementine,
		"deepkoamaru":                    DeepKoamaru,
		"duststorm":                      DustStorm,
		"endeavour":                      Endeavour,
		"mauvetaupe":                     MauveTaupe,
		"mediumvermilion":                MediumVermilion,
		"rosered":                        RoseRed,
		"bandicoot":                      Bandicoot,
		"silk":                           Silk,
		"juniper":                        Juniper,
		"oldmossgreen":                   OldMossGreen,
		"ferrarired":                     FerrariRed,
		"linen":                          Linen,
		"aquaisland":                     AquaIsland,
		"hippieblue":                     HippieBlue,
		"parchment":                      Parchment,
		"teagreen":                       TeaGreen,
		"copperrust":                     CopperRust,
		"leather":                        Leather,
		"rosegold":                       RoseGold,
		"camouflage":                     Camouflage,
		"hemp":                           Hemp,
		"peachorange":                    PeachOrange,
		"springwood":                     SpringWood,
		"stiletto":                       Stiletto,
		"deco":                           Deco,
		"goblin":                         Goblin,
		"rope":                           Rope,
		"sweetbrown":                     SweetBrown,
		"warmblack":                      WarmBlack,
		"bluewhale":                      BlueWhale,
		"caferoyale":                     CafeRoyale,
		"charmpink":                      CharmPink,
		"redwood":                        Redwood,
		"tasman":                         Tasman,
		"blue":                           Blue,
		"lava":                           Lava,
		"mamba":                          Mamba,
		"orangepeel":                     OrangePeel,
		"platinum":                       Platinum,
		"rybgreen":                       RYBGreen,
		"tiber":                          Tiber,
		"fringyflower":                   FringyFlower,
		"dodgerblue":                     DodgerBlue,
		"lawngreen":                      LawnGreen,
		"reddamask":                      RedDamask,
		"valencia":                       Valencia,
		"chineseviolet":                  ChineseViolet,
		"cranberry":                      Cranberry,
		"mischka":                        Mischka,
		"munsellred":                     MunsellRed,
		"sun":                            Sun,
		"athsspecial":                    AthsSpecial,
		"clearday":                       ClearDay,
		"darkpastelred":                  DarkPastelRed,
		"canaryyellow":                   CanaryYellow,
		"cognac":                         Cognac,
		"craterbrown":                    CraterBrown,
		"deepfir":                        DeepFir,
		"egyptianblue":                   EgyptianBlue,
		"orangesoda":                     OrangeSoda,
		"red":                            Red,
		"redviolet":                      RedViolet,
		"bronzetone":                     Bronzetone,
		"venus":                          Venus,
		"palecornflowerblue":             PaleCornflowerBlue,
		"processmagenta":                 ProcessMagenta,
		"rock":                           Rock,
		"mellowapricot":                  MellowApricot,
		"sherpablue":                     SherpaBlue,
		"texasrose":                      TexasRose,
		"loulou":                         Loulou,
		"fallow":                         Fallow,
		"graynurse":                      GrayNurse,
		"jaggedice":                      JaggedIce,
		"jasmine":                        Jasmine,
		"limedoak":                       LimedOak,
		"royalairforceblue":              RoyalAirForceBlue,
		"chromeyellow":                   ChromeYellow,
		"grandis":                        Grandis,
		"darkburgundy":                   DarkBurgundy,
		"corvette":                       Corvette,
		"heavymetal":                     HeavyMetal,
		"paua":                           Paua,
		"peru":                           Peru,
		"polishedpine":                   PolishedPine,
		"russiangreen":                   RussianGreen,
		"camarone":                       Camarone,
		"axolotl":                        Axolotl,
		"bleudefrance":                   BleuDeFrance,
		"hotpink":                        HotPink,
		"jellybean":                      JellyBean,
		"palecyan":                       PaleCyan,
		"pewter":                         Pewter,
		"punga":                          Punga,
		"antiquebronze":                  AntiqueBronze,
		"goldensand":                     GoldenSand,
		"sandstorm":                      Sandstorm,
		"bajawhite":                      BajaWhite,
		"purpleplum":                     PurplePlum,
		"sepiaskin":                      SepiaSkin,
		"slategray":                      SlateGray,
		"casper":                         Casper,
		"eunry":                          Eunry,
		"flint":                          Flint,
		"hookersgreen":                   HookersGreen,
		"lima":                           Lima,
		"shadow":                         Shadow,
		"veniceblue":                     VeniceBlue,
		"blackcoral":                     BlackCoral,
		"armygreen":                      ArmyGreen,
		"princessperfume":                PrincessPerfume,
		"silverlakeblue":                 SilverLakeBlue,
		"spanishgreen":                   SpanishGreen,
		"almond":                         Almond,
		"persianpink":                    PersianPink,
		"maroonoak":                      MaroonOak,
		"eggshell":                       Eggshell,
		"engineeringinternationalorange": EngineeringInternationalOrange,
		"lightcoral":                     LightCoral,
		"oil":                            Oil,
		"oldheliotrope":                  OldHeliotrope,
		"orient":                         Orient,
		"pinkraspberry":                  PinkRaspberry,
		"charm":                          Charm,
		"tealblue":                       TealBlue,
		"fuscousgray":                    FuscousGray,
		"sail":                           Sail,
		"seagull":                        Seagull,
		"valhalla":                       Valhalla,
		"bittersweet":                    Bittersweet,
		"olivehaze":                      OliveHaze,
		"pompadour":                      Pompadour,
		"romantic":                       Romantic,
		"cgblue":                         CGBlue,
		"citron":                         Citron,
		"concrete":                       Concrete,
		"cupid":                          Cupid,
		"pharlap":                        Pharlap,
		"shinyshamrock":                  ShinyShamrock,
		"amulet":                         Amulet,
		"blueribbon":                     BlueRibbon,
		"minttulip":                      MintTulip,
		"bermuda":                        Bermuda,
		"roseofsharon":                   RoseofSharon,
		"springgreen":                    SpringGreen,
		"studio":                         Studio,
		"careyspink":                     CareysPink,
		"deepseagreen":                   DeepSeaGreen,
		"dimgray":                        DimGray,
		"napa":                           Napa,
		"shinglefawn":                    ShingleFawn,
		"camouflagegreen":                CamouflageGreen,
		"purplemountainmajesty":          PurpleMountainMajesty,
		"topaz":                          Topaz,
		"bluebayoux":                     BlueBayoux,
		"delta":                          Delta,
		"luckypoint":                     LuckyPoint,
		"portage":                        Portage,
		"rusticred":                      RusticRed,
		"shockingpink":                   ShockingPink,
		"sunburntcyclops":                SunburntCyclops,
		"westside":                       WestSide,
		"cerisepink":                     CerisePink,
		"dawn":                           Dawn,
		"hanblue":                        HanBlue,
		"peat":                           Peat,
		"snow":                           Snow,
		"candlelight":                    Candlelight,
		"vegasgold":                      VegasGold,
		"palatinateblue":                 PalatinateBlue,
		"grainbrown":                     GrainBrown,
		"inchworm":                       InchWorm,
		"olivegreen":                     OliveGreen,
		"rouge":                          Rouge,
		"tapa":                           Tapa,
		"whitelinen":                     WhiteLinen,
		"frenchblue":                     FrenchBlue,
		"pantonegreen":                   PantoneGreen,
		"rawsienna":                      RawSienna,
		"spanishcarmine":                 SpanishCarmine,
		"darktangerine":                  DarkTangerine,
		"jacaranda":                      Jacaranda,
		"rose":                           Rose,
		"darkgoldenrod":                  DarkGoldenrod,
		"oriolesorange":                  OriolesOrange,
		"tussock":                        Tussock,
		"charcoal":                       Charcoal,
		"rosedust":                       RoseDust,
		"scienceblue":                    ScienceBlue,
		"oldburgundy":                    OldBurgundy,
		"chablis":                        Chablis,
		"himalaya":                       Himalaya,
		"internationalorange":            InternationalOrange,
		"meatbrown":                      MeatBrown,
		"selectiveyellow":                SelectiveYellow,
		"surfcrest":                      SurfCrest,
		"almondfrost":                    AlmondFrost,
		"chardonnay":                     Chardonnay,
		"diesel":                         Diesel,
		"dukeblue":                       DukeBlue,
		"fandango":                       Fandango,
		"jumbo":                          Jumbo,
		"smoke":                          Smoke,
		"spicymix":                       SpicyMix,
		"ao":                             Ao,
		"vistablue":                      VistaBlue,
		"finlandia":                      Finlandia,
		"kumera":                         Kumera,
		"mojo":                           Mojo,
		"mulberry":                       Mulberry,
		"crayolayellow":                  CrayolaYellow,
		"dulllavender":                   DullLavender,
		"rosevale":                       RoseVale,
		"creamcan":                       CreamCan,
		"flamepea":                       FlamePea,
		"palerose":                       PaleRose,
		"darkpurple":                     DarkPurple,
		"chiffon":                        Chiffon,
		"easternblue":                    EasternBlue,
		"ricecake":                       RiceCake,
		"celtic":                         Celtic,
		"lemonmeringue":                  LemonMeringue,
		"paleplum":                       PalePlum,
		"smitten":                        Smitten,
		"visvis":                         VisVis,
		"crayolagreen":                   CrayolaGreen,
		"oasis":                          Oasis,
		"racinggreen":                    RacingGreen,
		"russianviolet":                  RussianViolet,
		"spanishorange":                  SpanishOrange,
		"vistawhite":                     VistaWhite,
		"darkcandyapplered":              DarkCandyAppleRed,
		"pastelbrown":                    PastelBrown,
		"raven":                          Raven,
		"silverchalice":                  SilverChalice,
		"strikemaster":                   Strikemaster,
		"thatch":                         Thatch,
		"americanrose":                   AmericanRose,
		"diamond":                        Diamond,
		"goldenrod":                      Goldenrod,
		"porsche":                        Porsche,
		"riceflower":                     RiceFlower,
		"venetianred":                    VenetianRed,
		"anzac":                          Anzac,
		"jacarta":                        Jacarta,
		"ming":                           Ming,
		"oldsilver":                      OldSilver,
		"satinlinen":                     SatinLinen,
		"spanishviridian":                SpanishViridian,
		"brownrust":                      BrownRust,
		"grizzly":                        Grizzly,
		"pottersclay":                    PottersClay,
		"sunray":                         Sunray,
		"fawn":                           Fawn,
		"coral":                          Coral,
		"slateblue":                      SlateBlue,
		"yellowmetal":                    YellowMetal,
		"bonjour":                        BonJour,
		"grullo":                         Grullo,
		"junglemist":                     JungleMist,
		"martini":                        Martini,
		"monza":                          Monza,
		"surfiegreen":                    SurfieGreen,
		"japanesemaple":                  JapaneseMaple,
		"cedarwoodfinish":                CedarWoodFinish,
		"deepruby":                       DeepRuby,
		"fungreen":                       FunGreen,
		"pariswhite":                     ParisWhite,
		"spirodiscoball":                 SpiroDiscoBall,
		"britishracinggreen":             BritishRacingGreen,
		"azalea":                         Azalea,
		"bittersweetshimmer":             BittersweetShimmer,
		"chino":                          Chino,
		"chocolate":                      Chocolate,
		"givry":                          Givry,
		"goldfusion":                     GoldFusion,
		"hawkesblue":                     HawkesBlue,
		"australianmint":                 AustralianMint,
		"periwinkle":                     Periwinkle,
		"phthalogreen":                   PhthaloGreen,
		"roseebony":                      RoseEbony,
		"steelgray":                      SteelGray,
		"yelloworange":                   YellowOrange,
		"lonestar":                       Lonestar,
		"licorice":                       Licorice,
		"majorelleblue":                  MajorelleBlue,
		"stpatricksblue":                 StPatricksBlue,
		"ginger":                         Ginger,
		"blueberry":                      Blueberry,
		"gossamer":                       Gossamer,
		"melanzane":                      Melanzane,
		"verylighttangelo":               VeryLightTangelo,
		"wineberry":                      WineBerry,
		"wintersky":                      WinterSky,
		"babypowder":                     BabyPowder,
		"verypaleorange":                 VeryPaleOrange,
		"lincolngreen":                   LincolnGreen,
		"fielddrab":                      FieldDrab,
		"nevada":                         Nevada,
		"orchid":                         Orchid,
		"palecopper":                     PaleCopper,
		"anakiwa":                        Anakiwa,
		"frenchskyblue":                  FrenchSkyBlue,
		"laurelgreen":                    LaurelGreen,
		"luxorgold":                      LuxorGold,
		"nadeshikopink":                  NadeshikoPink,
		"santasgray":                     SantasGray,
		"skeptic":                        Skeptic,
		"vividcerise":                    VividCerise,
		"blossom":                        Blossom,
		"bluesmoke":                      BlueSmoke,
		"contessa":                       Contessa,
		"greensheen":                     GreenSheen,
		"balihai":                        BaliHai,
		"deepcarminepink":                DeepCarminePink,
		"loafer":                         Loafer,
		"pinkpearl":                      PinkPearl,
		"brickred":                       BrickRed,
		"mandalay":                       Mandalay,
		"mintgreen":                      MintGreen,
		"oceangreen":                     OceanGreen,
		"orangeroughy":                   OrangeRoughy,
		"putty":                          Putty,
		"reefgold":                       ReefGold,
		"galliano":                       Galliano,
		"arrowtown":                      Arrowtown,
		"halfcolonialwhite":              HalfColonialWhite,
		"neongreen":                      NeonGreen,
		"rybyellow":                      RYBYellow,
		"vulcan":                         Vulcan,
		"apricotwhite":                   ApricotWhite,
		"denim":                          Denim,
		"siam":                           Siam,
		"darkorange":                     DarkOrange,
		"corduroy":                       Corduroy,
		"cyprus":                         Cyprus,
		"judgegray":                      JudgeGray,
		"rosepink":                       RosePink,
		"alabamacrimson":                 AlabamaCrimson,
		"purpleheart":                    PurpleHeart,
		"olivedrab":                      OliveDrab,
		"california":                     California,
		"brightnavyblue":                 BrightNavyBlue,
		"carolinablue":                   CarolinaBlue,
		"dew":                            Dew,
		"honolulublue":                   HonoluluBlue,
		"iceberg":                        Iceberg,
		"sandybeach":                     SandyBeach,
		"camelot":                        Camelot,
		"olivine":                        Olivine,
		"woodybrown":                     WoodyBrown,
		"ogreodor":                       OgreOdor,
		"beaublue":                       BeauBlue,
		"darkliver":                      DarkLiver,
		"everglade":                      Everglade,
		"alto":                           Alto,
		"coffee":                         Coffee,
		"cumin":                          Cumin,
		"darkgunmetal":                   DarkGunmetal,
		"deepsapphire":                   DeepSapphire,
		"neonfuchsia":                    NeonFuchsia,
		"nobel":                          Nobel,
		"celestialblue":                  CelestialBlue,
		"seance":                         Seance,
		"thistle":                        Thistle,
		"lavendergray":                   LavenderGray,
		"christalle":                     Christalle,
		"englishvermillion":              EnglishVermillion,
		"husk":                           Husk,
		"bistre":                         Bistre,
		"tropicalblue":                   TropicalBlue,
		"zinnwaldite":                    Zinnwaldite,
		"coffeebean":                     CoffeeBean,
		"bayofmany":                      BayofMany,
		"doublespanishwhite":             DoubleSpanishWhite,
		"lemonlime":                      LemonLime,
		"lenurple":                       Lenurple,
		"paleoyster":                     PaleOyster,
		"twine":                          Twine,
		"affair":                         Affair,
		"cloud":                          Cloud,
		"pizza":                          Pizza,
		"bluedianne":                     BlueDianne,
		"monalisa":                       MonaLisa,
		"spartancrimson":                 SpartanCrimson,
		"battleshipgray":                 BattleshipGray,
		"bossanova":                      Bossanova,
		"capecod":                        CapeCod,
		"goldendream":                    GoldenDream,
		"rosewood":                       Rosewood,
		"birdflower":                     BirdFlower,
		"pinklace":                       PinkLace,
		"paprika":                        Paprika,
		"kournikova":                     Kournikova,
		"pancho":                         Pancho,
		"resolutionblue":                 ResolutionBlue,
		"riogrande":                      RioGrande,
		"japonica":                       Japonica,
		"darkvanilla":                    DarkVanilla,
		"geebung":                        Geebung,
		"jordyblue":                      JordyBlue,
		"kilamanjaro":                    Kilamanjaro,
		"wildwatermelon":                 WildWatermelon,
		"cioccolato":                     Cioccolato,
		"petiteorchid":                   PetiteOrchid,
		"puertorico":                     PuertoRico,
		"rhino":                          Rhino,
		"solidpink":                      SolidPink,
		"spanishgray":                    SpanishGray,
		"pacificblue":                    PacificBlue,
		"blackrose":                      BlackRose,
		"edgewater":                      Edgewater,
		"electricblue":                   ElectricBlue,
		"frangipani":                     Frangipani,
		"gunmetal":                       Gunmetal,
		"mandy":                          Mandy,
		"mantle":                         Mantle,
		"applegreen":                     AppleGreen,
		"montecarlo":                     MonteCarlo,
		"pastelgreen":                    PastelGreen,
		"rosewhite":                      RoseWhite,
		"saffronmango":                   SaffronMango,
		"toryblue":                       ToryBlue,
		"westar":                         Westar,
		"firebush":                       FireBush,
		"sacramentostategreen":           SacramentoStateGreen,
		"uclablue":                       UCLABlue,
		"yellowrose":                     YellowRose,
		"lightgoldenrodyellow":           LightGoldenrodYellow,
		"capri":                          Capri,
		"tyrianpurple":                   TyrianPurple,
		"zaffre":                         Zaffre,
		"bluegray":                       BlueGray,
		"fandangopink":                   FandangoPink,
		"goldtips":                       GoldTips,
		"soapstone":                      Soapstone,
		"tangopink":                      TangoPink,
		"cloudy":                         Cloudy,
		"quincy":                         Quincy,
		"tomato":                         Tomato,
		"dallas":                         Dallas,
		"cinnamonsatin":                  CinnamonSatin,
		"deepfuchsia":                    DeepFuchsia,
		"heath":                          Heath,
		"hemlock":                        Hemlock,
		"huntergreen":                    HunterGreen,
		"kangaroo":                       Kangaroo,
		"manhattan":                      Manhattan,
		"bastille":                       Bastille,
		"burningorange":                  BurningOrange,
		"mariner":                        Mariner,
		"mikadoyellow":                   MikadoYellow,
		"rollingstone":                   RollingStone,
		"shamrockgreen":                  ShamrockGreen,
		"starkwhite":                     StarkWhite,
		"sunflower":                      Sunflower,
		"bleachwhite":                    BleachWhite,
		"frenchraspberry":                FrenchRaspberry,
		"linkwater":                      LinkWater,
		"apricot":                        Apricot,
		"heliotropegray":                 HeliotropeGray,
		"mediumblue":                     MediumBlue,
		"powderash":                      PowderAsh,
		"wellread":                       WellRead,
		"aureolin":                       Aureolin,
		"mulefawn":                       MuleFawn,
		"nebula":                         Nebula,
		"palegold":                       PaleGold,
		"cyanazure":                      CyanAzure,
		"mediumorchid":                   MediumOrchid,
		"ncsred":                         NCSRed,
		"pelorous":                       Pelorous,
		"pipi":                           Pipi,
		"spanishbistre":                  SpanishBistre,
		"countygreen":                    CountyGreen,
		"iroko":                          Iroko,
		"lightcyan":                      LightCyan,
		"cyanblueazure":                  CyanBlueAzure,
		"conch":                          Conch,
		"deepgreencyanturquoise":         DeepGreenCyanTurquoise,
		"mypink":                         MyPink,
		"spanishblue":                    SpanishBlue,
		"chinaivory":                     ChinaIvory,
		"bunker":                         Bunker,
		"colonialwhite":                  ColonialWhite,
		"lola":                           Lola,
		"pantoneyellow":                  PantoneYellow,
		"richblack":                      RichBlack,
		"rufous":                         Rufous,
		"sushi":                          Sushi,
		"aeroblue":                       AeroBlue,
		"trendygreen":                    TrendyGreen,
		"thunderbird":                    Thunderbird,
		"crownofthorns":                  CrownofThorns,
		"guardsmanred":                   GuardsmanRed,
		"mako":                           Mako,
		"surf":                           Surf,
		"utahcrimson":                    UtahCrimson,
		"asphalt":                        Asphalt,
		"parisdaisy":                     ParisDaisy,
		"richelectricblue":               RichElectricBlue,
		"tidal":                          Tidal,
		"classicrose":                    ClassicRose,
		"midnightblue":                   MidnightBlue,
		"quicksand":                      Quicksand,
		"rubyred":                        RubyRed,
		"lipstick":                       Lipstick,
		"graynickel":                     GrayNickel,
		"jet":                            Jet,
		"muddywaters":                    MuddyWaters,
		"ruddypink":                      RuddyPink,
		"darkorchid":                     DarkOrchid,
		"goldenbrown":                    GoldenBrown,
		"hibiscus":                       Hibiscus,
		"rosebudcherry":                  RoseBudCherry,
		"tamarillo":                      Tamarillo,
		"caputmortuum":                   CaputMortuum,
		"goldenglow":                     GoldenGlow,
		"salmonpink":                     SalmonPink,
		"tallpoppy":                      TallPoppy,
		"urobilin":                       Urobilin,
		"cashmere":                       Cashmere,
		"coldturkey":                     ColdTurkey,
		"crimson":                        Crimson,
		"ferngreen":                      FernGreen,
		"fuzzywuzzybrown":                FuzzyWuzzyBrown,
		"kobicha":                        Kobicha,
		"negroni":                        Negroni,
		"panache":                        Panache,
		"blond":                          Blond,
		"wildstrawberry":                 WildStrawberry,
		"ripeplum":                       RipePlum,
		"cherokee":                       Cherokee,
		"clinker":                        Clinker,
		"janna":                          Janna,
		"limedspruce":                    LimedSpruce,
		"magicpotion":                    MagicPotion,
		"poloblue":                       PoloBlue,
		"chantilly":                      Chantilly,
		"imperialred":                    ImperialRed,
		"ncsblue":                        NCSBlue,
		"pastelpink":                     PastelPink,
		"roofterracotta":                 RoofTerracotta,
		"umber":                          Umber,
		"babyblue":                       BabyBlue,
		"cameopink":                      CameoPink,
		"mimosa":                         Mimosa,
		"varden":                         Varden,
		"arylideyellow":                  ArylideYellow,
		"lavenderrose":                   LavenderRose,
		"pampas":                         Pampas,
		"sweetpink":                      SweetPink,
		"winterhazel":                    WinterHazel,
		"hottoddy":                       HotToddy,
		"horses":                         Horses,
		"renosand":                       RenoSand,
		"sizzlingsunrise":                SizzlingSunrise,
		"zumthor":                        Zumthor,
		"dirt":                           Dirt,
		"lightdeeppink":                  LightDeepPink,
		"madison":                        Madison,
		"paleleaf":                       PaleLeaf,
		"sunglow":                        Sunglow,
		"toreabay":                       ToreaBay,
		"truev":                          TrueV,
		"brightturquoise":                BrightTurquoise,
		"froly":                          Froly,
		"japaneselaurel":                 JapaneseLaurel,
		"karaka":                         Karaka,
		"melon":                          Melon,
		"pullmanbrown":                   PullmanBrown,
		"ultrapink":                      UltraPink,
		"buff":                           Buff,
		"cararra":                        Cararra,
		"lily":                           Lily,
		"bulgarianrose":                  BulgarianRose,
		"popstar":                        Popstar,
		"slimygreen":                     SlimyGreen,
		"vinrouge":                       VinRouge,
		"woodbark":                       WoodBark,
		"equator":                        Equator,
		"chatelle":                       Chatelle,
		"mustard":                        Mustard,
		"pewterblue":                     PewterBlue,
		"portlandorange":                 PortlandOrange,
		"provincialpink":                 ProvincialPink,
		"snowymint":                      SnowyMint,
		"bone":                           Bone,
		"rodeodust":                      RodeoDust,
		"athensgray":                     AthensGray,
		"electricpurple":                 ElectricPurple,
		"halfdutchwhite":                 HalfDutchWhite,
		"supernova":                      Supernova,
		"verylightmalachitegreen":        VeryLightMalachiteGreen,
		"crayolared":                     CrayolaRed,
		"festival":                       Festival,
		"liberty":                        Liberty,
		"lochinvar":                      Lochinvar,
		"cloudburst":                     CloudBurst,
		"barnred":                        BarnRed,
		"cosmic":                         Cosmic,
		"darkgreen":                      DarkGreen,
		"eggwhite":                       EggWhite,
		"norway":                         Norway,
		"patina":                         Patina,
		"aubergine":                      Aubergine,
		"walnut":                         Walnut,
		"blueromance":                    BlueRomance,
		"uclagold":                       UCLAGold,
		"balticsea":                      BalticSea,
		"paleslate":                      PaleSlate,
		"safetyyellow":                   SafetyYellow,
		"babyblueeyes":                   BabyBlueEyes,
		"lunargreen":                     LunarGreen,
		"blanchedalmond":                 BlanchedAlmond,
		"crimsonglory":                   CrimsonGlory,
		"espresso":                       Espresso,
		"harlequingreen":                 HarlequinGreen,
		"acidgreen":                      AcidGreen,
		"butteredrum":                    ButteredRum,
		"lemongrass":                     LemonGrass,
		"persianrose":                    PersianRose,
		"androidgreen":                   AndroidGreen,
		"cyberyellow":                    CyberYellow,
		"smoky":                          Smoky,
		"acapulco":                       Acapulco,
		"blizzardblue":                   BlizzardBlue,
		"londonhue":                      LondonHue,
		"magentapink":                    MagentaPink,
		"tea":                            Tea,
		"underagepink":                   UnderagePink,
		"bisonhide":                      BisonHide,
		"mobster":                        Mobster,
		"pinksherbet":                    PinkSherbet,
		"portafino":                      Portafino,
		"haiti":                          Haiti,
		"lightcornflowerblue":            LightCornflowerBlue,
		"mountbattenpink":                MountbattenPink,
		"palecarmine":                    PaleCarmine,
		"pictorialcarmine":               PictorialCarmine,
		"tango":                          Tango,
		"cocoabean":                      CocoaBean,
		"blackforest":                    BlackForest,
		"clairvoyant":                    Clairvoyant,
		"pastelpurple":                   PastelPurple,
		"satinsheengold":                 SatinSheenGold,
		"bazaar":                         Bazaar,
		"frenchpink":                     FrenchPink,
		"olivetone":                      Olivetone,
		"pastelyellow":                   PastelYellow,
		"quicksilver":                    QuickSilver,
		"tuatara":                        Tuatara,
		"eminence":                       Eminence,
		"burningsand":                    BurningSand,
		"coconutcream":                   CoconutCream,
		"dollarbill":                     DollarBill,
		"zombie":                         Zombie,
		"bracken":                        Bracken,
		"glaucous":                       Glaucous,
		"rangoongreen":                   RangoonGreen,
		"totempole":                      TotemPole,
		"darkseagreen":                   DarkSeaGreen,
		"martinique":                     Martinique,
		"nero":                           Nero,
		"maroon":                         Maroon,
		"terracotta":                     TerraCotta,
		"gimblet":                        Gimblet,
		"sapphire":                       Sapphire,
		"primrose":                       Primrose,
		"gondola":                        Gondola,
		"golden":                         Golden,
		"maygreen":                       MayGreen,
		"cinnabar":                       Cinnabar,
		"brandypunch":                    BrandyPunch,
		"desaturatedcyan":                DesaturatedCyan,
		"ryborange":                      RYBOrange,
		"smalt":                          Smalt,
		"alloyorange":                    AlloyOrange,
		"brightyellow":                   BrightYellow,
		"belgion":                        Belgion,
		"cherryblossompink":              CherryBlossomPink,
		"flaxsmoke":                      FlaxSmoke,
		"seaserpent":                     SeaSerpent,
		"spunpearl":                      SpunPearl,
		"apache":                         Apache,
		"gossip":                         Gossip,
		"waterspout":                     Waterspout,
		"fogra39richblack":               FOGRA39RichBlack,
		"boysenberry":                    Boysenberry,
		"feijoa":                         Feijoa,
		"genericviridian":                GenericViridian,
		"mediumseagreen":                 MediumSeaGreen,
		"schist":                         Schist,
		"bitterlemon":                    BitterLemon,
		"crayolablue":                    CrayolaBlue,
		"palebrown":                      PaleBrown,
		"cornfield":                      CornField,
		"elm":                            Elm,
		"gin":                            Gin,
		"goldenpoppy":                    GoldenPoppy,
		"nickel":                         Nickel,
		"rosebud":                        RoseBud,
		"deepred":                        DeepRed,
		"oregon":                         Oregon,
		"zomp":                           Zomp,
		"malibu":                         Malibu,
		"carnation":                      Carnation,
		"grenadier":                      Grenadier,
		"onahau":                         Onahau,
		"shuttlegray":                    ShuttleGray,
		"uared":                          UARed,
		"broom":                          Broom,
		"mosque":                         Mosque,
		"myrtlegreen":                    MyrtleGreen,
		"sunny":                          Sunny,
		"burgundy":                       Burgundy,
		"englishlavender":                EnglishLavender,
		"windsortan":                     WindsorTan,
		"cybergrape":                     CyberGrape,
		"tacha":                          Tacha,
		"pastelmagenta":                  PastelMagenta,
		"brightlilac":                    BrightLilac,
		"chardon":                        Chardon,
		"darkimperialblue":               DarkImperialBlue,
		"eggplant":                       Eggplant,
		"meteor":                         Meteor,
		"wewak":                          Wewak,
		"whitesmoke":                     WhiteSmoke,
		"appleblossom":                   AppleBlossom,
		"darkcyan":                       DarkCyan,
		"persiangreen":                   PersianGreen,
		"vividcerulean":                  VividCerulean,
		"bigstone":                       BigStone,
		"millbrook":                      Millbrook,
		"offgreen":                       OffGreen,
		"universityoftennesseeorange":    UniversityOfTennesseeOrange,
		"crocodile":                      Crocodile,
		"richcarmine":                    RichCarmine,
		"deepsaffron":                    DeepSaffron,
		"grannysmithapple":               GrannySmithApple,
		"organ":                          Organ,
		"palegreen":                      PaleGreen,
		"tangelo":                        Tangelo,
		"blueviolet":                     BlueViolet,
		"frostbite":                      Frostbite,
		"ivory":                          Ivory,
		"keylimepie":                     KeyLimePie,
		"shilo":                          Shilo,
		"sizzlingred":                    SizzlingRed,
		"tapestry":                       Tapestry,
		"bdazzledblue":                   BdazzledBlue,
		"darkbyzantium":                  DarkByzantium,
		"elpaso":                         ElPaso,
		"grayasparagus":                  GrayAsparagus,
		"vancleef":                       VanCleef,
		"vividlimegreen":                 VividLimeGreen,
		"vividskyblue":                   VividSkyBlue,
		"bourbon":                        Bourbon,
		"tiara":                          Tiara,
		"winedregs":                      WineDregs,
		"arcticlime":                     ArcticLime,
		"solitaire":                      Solitaire,
		"darkjunglegreen":                DarkJungleGreen,
		"morningglory":                   MorningGlory,
		"parism":                         ParisM,
		"pickledbluewood":                PickledBluewood,
		"byzantine":                      Byzantine,
		"lilacluster":                    LilacLuster,
		"opium":                          Opium,
		"scorpion":                       Scorpion,
		"voodoo":                         Voodoo,
		"bostonblue":                     BostonBlue,
		"cabaret":                        Cabaret,
		"carissma":                       Carissma,
		"chateaugreen":                   ChateauGreen,
		"creole":                         Creole,
		"greencyan":                      GreenCyan,
		"hawaiiantan":                    HawaiianTan,
		"nutmeg":                         Nutmeg,
		"bluegreen":                      BlueGreen,
		"treepoppy":                      TreePoppy,
		"finch":                          Finch,
		"redrobin":                       RedRobin,
		"azureishwhite":                  AzureishWhite,
		"electriclime":                   ElectricLime,
		"mediumslateblue":                MediumSlateBlue,
		"olive":                          Olive,
		"violet":                         Violet,
		"darkkhaki":                      DarkKhaki,
		"cedarchest":                     CedarChest,
		"lisbonbrown":                    LisbonBrown,
		"antiquewhite":                   AntiqueWhite,
		"graychateau":                    GrayChateau,
		"jon":                            Jon,
		"nightshadz":                     NightShadz,
		"steelblue":                      SteelBlue,
		"vesuvius":                       Vesuvius,
		"flamingo":                       Flamingo,
		"lavenderblush":                  LavenderBlush,
		"palelavender":                   PaleLavender,
		"palmgreen":                      PalmGreen,
		"weldonblue":                     WeldonBlue,
		"eclipse":                        Eclipse,
		"jewel":                          Jewel,
		"mint":                           Mint,
		"blumine":                        Blumine,
		"flame":                          Flame,
		"jasper":                         Jasper,
		"midnight":                       Midnight,
		"butterflybush":                  ButterflyBush,
		"falcon":                         Falcon,
		"sidecar":                        Sidecar,
		"soyabean":                       SoyaBean,
		"vividvermilion":                 VividVermilion,
		"cinereous":                      Cinereous,
		"sorbus":                         Sorbus,
		"travertine":                     Travertine,
		"castro":                         Castro,
		"bush":                           Bush,
		"cameo":                          Cameo,
		"mistymoss":                      MistyMoss,
		"atlantis":                       Atlantis,
		"marigoldyellow":                 MarigoldYellow,
		"roastcoffee":                    RoastCoffee,
		"rosetaupe":                      RoseTaupe,
		"shadowblue":                     ShadowBlue,
		"lumber":                         Lumber,
		"transparent":                    Transparent,
	}
)

func Named(name string) (Color, bool) {
	color, exists := Map[strings.ToLower(name)]
	return color, exists
}
