package ui

import "strings"

var (
	ColorIllusion                       = Color{R: 0.965, G: 0.643, B: 0.788, A: 1}
	ColorSonicSilver                    = Color{R: 0.459, G: 0.459, B: 0.459, A: 1}
	ColorTana                           = Color{R: 0.851, G: 0.863, B: 0.757, A: 1}
	ColorApache                         = Color{R: 0.875, G: 0.745, B: 0.435, A: 1}
	ColorAuroMetalSaurus                = Color{R: 0.431, G: 0.498, B: 0.502, A: 1}
	ColorCreamCan                       = Color{R: 0.961, G: 0.784, B: 0.361, A: 1}
	ColorLightSkyBlue                   = Color{R: 0.529, G: 0.808, B: 0.980, A: 1}
	ColorBleachWhite                    = Color{R: 0.996, G: 0.953, B: 0.847, A: 1}
	ColorCeltic                         = Color{R: 0.086, G: 0.196, B: 0.133, A: 1}
	ColorEarlsGreen                     = Color{R: 0.788, G: 0.725, B: 0.231, A: 1}
	ColorCrimsonGlory                   = Color{R: 0.745, G: 0.000, B: 0.196, A: 1}
	ColorCrownofThorns                  = Color{R: 0.467, G: 0.122, B: 0.122, A: 1}
	ColorDeepMagenta                    = Color{R: 0.800, G: 0.000, B: 0.800, A: 1}
	ColorJade                           = Color{R: 0.000, G: 0.659, B: 0.420, A: 1}
	ColorKoromiko                       = Color{R: 1.000, G: 0.741, B: 0.373, A: 1}
	ColorBlueSmoke                      = Color{R: 0.455, G: 0.533, B: 0.506, A: 1}
	ColorBracken                        = Color{R: 0.290, G: 0.165, B: 0.016, A: 1}
	ColorBurningOrange                  = Color{R: 1.000, G: 0.439, B: 0.204, A: 1}
	ColorPurpureus                      = Color{R: 0.604, G: 0.306, B: 0.682, A: 1}
	ColorRioGrande                      = Color{R: 0.733, G: 0.816, B: 0.035, A: 1}
	ColorPinkOrange                     = Color{R: 1.000, G: 0.600, B: 0.400, A: 1}
	ColorTimberGreen                    = Color{R: 0.086, G: 0.196, B: 0.173, A: 1}
	ColorFreshAir                       = Color{R: 0.651, G: 0.906, B: 1.000, A: 1}
	ColorMulledWine                     = Color{R: 0.306, G: 0.271, B: 0.384, A: 1}
	ColorOracle                         = Color{R: 0.216, G: 0.455, B: 0.459, A: 1}
	ColorThulianPink                    = Color{R: 0.871, G: 0.435, B: 0.631, A: 1}
	ColorBlumine                        = Color{R: 0.094, G: 0.345, B: 0.478, A: 1}
	ColorOgreOdor                       = Color{R: 0.992, G: 0.322, B: 0.251, A: 1}
	ColorRazzmatazz                     = Color{R: 0.890, G: 0.145, B: 0.420, A: 1}
	ColorCatawba                        = Color{R: 0.439, G: 0.212, B: 0.259, A: 1}
	ColorKashmirBlue                    = Color{R: 0.314, G: 0.439, B: 0.588, A: 1}
	ColorOpal                           = Color{R: 0.663, G: 0.776, B: 0.761, A: 1}
	ColorShamrockGreen                  = Color{R: 0.000, G: 0.620, B: 0.376, A: 1}
	ColorFuchsiaRose                    = Color{R: 0.780, G: 0.263, B: 0.459, A: 1}
	ColorMoccaccino                     = Color{R: 0.431, G: 0.114, B: 0.078, A: 1}
	ColorNyanza                         = Color{R: 0.914, G: 1.000, B: 0.859, A: 1}
	ColorMing                           = Color{R: 0.212, G: 0.455, B: 0.490, A: 1}
	ColorOldMossGreen                   = Color{R: 0.525, G: 0.494, B: 0.212, A: 1}
	ColorSnow                           = Color{R: 1.000, G: 0.980, B: 0.980, A: 1}
	ColorSpringFrost                    = Color{R: 0.529, G: 1.000, B: 0.165, A: 1}
	ColorCatalinaBlue                   = Color{R: 0.024, G: 0.165, B: 0.471, A: 1}
	ColorCeladon                        = Color{R: 0.675, G: 0.882, B: 0.686, A: 1}
	ColorClayCreek                      = Color{R: 0.541, G: 0.514, B: 0.376, A: 1}
	ColorRussett                        = Color{R: 0.459, G: 0.353, B: 0.341, A: 1}
	ColorSail                           = Color{R: 0.722, G: 0.878, B: 0.976, A: 1}
	ColorTyrianPurple                   = Color{R: 0.400, G: 0.008, B: 0.235, A: 1}
	ColorUnitedNationsBlue              = Color{R: 0.357, G: 0.573, B: 0.898, A: 1}
	ColorVeryLightTangelo               = Color{R: 1.000, G: 0.690, B: 0.467, A: 1}
	ColorChardon                        = Color{R: 1.000, G: 0.953, B: 0.945, A: 1}
	ColorPixieGreen                     = Color{R: 0.753, G: 0.847, B: 0.714, A: 1}
	ColorTickleMePink                   = Color{R: 0.988, G: 0.537, B: 0.675, A: 1}
	ColorRoti                           = Color{R: 0.776, G: 0.659, B: 0.294, A: 1}
	ColorWheat                          = Color{R: 0.961, G: 0.871, B: 0.702, A: 1}
	ColorZambezi                        = Color{R: 0.408, G: 0.333, B: 0.345, A: 1}
	ColorAirSuperiorityBlue             = Color{R: 0.447, G: 0.627, B: 0.757, A: 1}
	ColorCyan                           = Color{R: 0.000, G: 1.000, B: 1.000, A: 1}
	ColorKilamanjaro                    = Color{R: 0.141, G: 0.047, B: 0.008, A: 1}
	ColorSpice                          = Color{R: 0.416, G: 0.267, B: 0.180, A: 1}
	ColorSushi                          = Color{R: 0.529, G: 0.671, B: 0.224, A: 1}
	ColorTrendyPink                     = Color{R: 0.549, G: 0.392, B: 0.584, A: 1}
	ColorYourPink                       = Color{R: 1.000, G: 0.765, B: 0.753, A: 1}
	ColorCalypso                        = Color{R: 0.192, G: 0.447, B: 0.553, A: 1}
	ColorCarrotOrange                   = Color{R: 0.929, G: 0.569, B: 0.129, A: 1}
	ColorCrayolaGreen                   = Color{R: 0.110, G: 0.675, B: 0.471, A: 1}
	ColorTangerine                      = Color{R: 0.949, G: 0.522, B: 0.000, A: 1}
	ColorGunmetal                       = Color{R: 0.165, G: 0.204, B: 0.224, A: 1}
	ColorSienna                         = Color{R: 0.533, G: 0.176, B: 0.090, A: 1}
	ColorSilverSand                     = Color{R: 0.749, G: 0.757, B: 0.761, A: 1}
	ColorDerby                          = Color{R: 1.000, G: 0.933, B: 0.847, A: 1}
	ColorGoBen                          = Color{R: 0.447, G: 0.427, B: 0.306, A: 1}
	ColorLightSteelBlue                 = Color{R: 0.690, G: 0.769, B: 0.871, A: 1}
	ColorLoblolly                       = Color{R: 0.741, G: 0.788, B: 0.808, A: 1}
	ColorScarletGum                     = Color{R: 0.263, G: 0.082, B: 0.376, A: 1}
	ColorAzure                          = Color{R: 0.000, G: 0.498, B: 1.000, A: 1}
	ColorBrickRed                       = Color{R: 0.796, G: 0.255, B: 0.329, A: 1}
	ColorCitron                         = Color{R: 0.624, G: 0.663, B: 0.122, A: 1}
	ColorWinterSky                      = Color{R: 1.000, G: 0.000, B: 0.486, A: 1}
	ColorMojo                           = Color{R: 0.753, G: 0.278, B: 0.216, A: 1}
	ColorRedStage                       = Color{R: 0.816, G: 0.373, B: 0.016, A: 1}
	ColorTelemagenta                    = Color{R: 0.812, G: 0.204, B: 0.463, A: 1}
	ColorBudGreen                       = Color{R: 0.482, G: 0.714, B: 0.380, A: 1}
	ColorDeepRed                        = Color{R: 0.522, G: 0.004, B: 0.004, A: 1}
	ColorDrover                         = Color{R: 0.992, G: 0.969, B: 0.678, A: 1}
	ColorMunsellPurple                  = Color{R: 0.624, G: 0.000, B: 0.773, A: 1}
	ColorSepiaBlack                     = Color{R: 0.169, G: 0.008, B: 0.008, A: 1}
	ColorAstronaut                      = Color{R: 0.157, G: 0.227, B: 0.467, A: 1}
	ColorGlacier                        = Color{R: 0.502, G: 0.702, B: 0.769, A: 1}
	ColorMosque                         = Color{R: 0.012, G: 0.416, B: 0.431, A: 1}
	ColorVividAuburn                    = Color{R: 0.573, G: 0.153, B: 0.141, A: 1}
	ColorIron                           = Color{R: 0.831, G: 0.843, B: 0.851, A: 1}
	ColorPaleGreen                      = Color{R: 0.596, G: 0.984, B: 0.596, A: 1}
	ColorSanFelix                       = Color{R: 0.043, G: 0.384, B: 0.027, A: 1}
	ColorMonteCarlo                     = Color{R: 0.514, G: 0.816, B: 0.776, A: 1}
	ColorSheenGreen                     = Color{R: 0.561, G: 0.831, B: 0.000, A: 1}
	ColorAmericano                      = Color{R: 0.529, G: 0.459, B: 0.431, A: 1}
	ColorHibiscus                       = Color{R: 0.714, G: 0.192, B: 0.424, A: 1}
	ColorLunarGreen                     = Color{R: 0.235, G: 0.286, B: 0.227, A: 1}
	ColorTiber                          = Color{R: 0.024, G: 0.208, B: 0.216, A: 1}
	ColorHimalaya                       = Color{R: 0.416, G: 0.365, B: 0.106, A: 1}
	ColorLawnGreen                      = Color{R: 0.486, G: 0.988, B: 0.000, A: 1}
	ColorRussianGreen                   = Color{R: 0.404, G: 0.573, B: 0.404, A: 1}
	ColorDarkFern                       = Color{R: 0.039, G: 0.282, B: 0.051, A: 1}
	ColorOregon                         = Color{R: 0.608, G: 0.278, B: 0.012, A: 1}
	ColorPastelMagenta                  = Color{R: 0.957, G: 0.604, B: 0.761, A: 1}
	ColorSprout                         = Color{R: 0.757, G: 0.843, B: 0.690, A: 1}
	ColorSteelTeal                      = Color{R: 0.373, G: 0.541, B: 0.545, A: 1}
	ColorAlpine                         = Color{R: 0.686, G: 0.561, B: 0.173, A: 1}
	ColorDomino                         = Color{R: 0.557, G: 0.467, B: 0.369, A: 1}
	ColorJapaneseMaple                  = Color{R: 0.471, G: 0.004, B: 0.035, A: 1}
	ColorPigeonPost                     = Color{R: 0.686, G: 0.741, B: 0.851, A: 1}
	ColorBrightMaroon                   = Color{R: 0.765, G: 0.129, B: 0.282, A: 1}
	ColorHalfDutchWhite                 = Color{R: 0.996, G: 0.969, B: 0.871, A: 1}
	ColorHummingBird                    = Color{R: 0.812, G: 0.976, B: 0.953, A: 1}
	ColorCumulus                        = Color{R: 0.992, G: 1.000, B: 0.835, A: 1}
	ColorRoyalAzure                     = Color{R: 0.000, G: 0.220, B: 0.659, A: 1}
	ColorSinopia                        = Color{R: 0.796, G: 0.255, B: 0.043, A: 1}
	ColorFoam                           = Color{R: 0.847, G: 0.988, B: 0.980, A: 1}
	ColorSlimyGreen                     = Color{R: 0.161, G: 0.588, B: 0.090, A: 1}
	ColorLemonGrass                     = Color{R: 0.608, G: 0.620, B: 0.561, A: 1}
	ColorPersianRose                    = Color{R: 0.996, G: 0.157, B: 0.635, A: 1}
	ColorPrim                           = Color{R: 0.941, G: 0.886, B: 0.925, A: 1}
	ColorWistful                        = Color{R: 0.643, G: 0.651, B: 0.827, A: 1}
	ColorCaramel                        = Color{R: 1.000, G: 0.867, B: 0.686, A: 1}
	ColorDollarBill                     = Color{R: 0.522, G: 0.733, B: 0.396, A: 1}
	ColorLaurelGreen                    = Color{R: 0.663, G: 0.729, B: 0.616, A: 1}
	ColorAtlantis                       = Color{R: 0.592, G: 0.804, B: 0.176, A: 1}
	ColorCararra                        = Color{R: 0.933, G: 0.933, B: 0.910, A: 1}
	ColorSkobeloff                      = Color{R: 0.000, G: 0.455, B: 0.455, A: 1}
	ColorTango                          = Color{R: 0.929, G: 0.478, B: 0.110, A: 1}
	ColorHarvestGold                    = Color{R: 0.855, G: 0.569, B: 0.000, A: 1}
	ColorLightKhaki                     = Color{R: 0.941, G: 0.902, B: 0.549, A: 1}
	ColorSunsetOrange                   = Color{R: 0.992, G: 0.369, B: 0.325, A: 1}
	ColorHonoluluBlue                   = Color{R: 0.000, G: 0.427, B: 0.690, A: 1}
	ColorOrient                         = Color{R: 0.004, G: 0.369, B: 0.522, A: 1}
	ColorSantasGray                     = Color{R: 0.624, G: 0.627, B: 0.694, A: 1}
	ColorOldGold                        = Color{R: 0.812, G: 0.710, B: 0.231, A: 1}
	ColorPinkSherbet                    = Color{R: 0.969, G: 0.561, B: 0.655, A: 1}
	ColorPlantation                     = Color{R: 0.153, G: 0.314, B: 0.294, A: 1}
	ColorSisal                          = Color{R: 0.827, G: 0.796, B: 0.729, A: 1}
	ColorBrownYellow                    = Color{R: 0.800, G: 0.600, B: 0.400, A: 1}
	ColorEmpress                        = Color{R: 0.506, G: 0.451, B: 0.467, A: 1}
	ColorGullGray                       = Color{R: 0.616, G: 0.675, B: 0.718, A: 1}
	ColorMantis                         = Color{R: 0.455, G: 0.765, B: 0.396, A: 1}
	ColorOrchidWhite                    = Color{R: 1.000, G: 0.992, B: 0.953, A: 1}
	ColorPippin                         = Color{R: 1.000, G: 0.882, B: 0.875, A: 1}
	ColorStarCommandBlue                = Color{R: 0.000, G: 0.482, B: 0.722, A: 1}
	ColorAuChico                        = Color{R: 0.592, G: 0.376, B: 0.365, A: 1}
	ColorEbony                          = Color{R: 0.333, G: 0.365, B: 0.314, A: 1}
	ColorIceCold                        = Color{R: 0.694, G: 0.957, B: 0.906, A: 1}
	ColorSanJuan                        = Color{R: 0.188, G: 0.294, B: 0.416, A: 1}
	ColorSmashedPumpkin                 = Color{R: 1.000, G: 0.427, B: 0.227, A: 1}
	ColorTePapaGreen                    = Color{R: 0.118, G: 0.263, B: 0.235, A: 1}
	ColorTopaz                          = Color{R: 1.000, G: 0.784, B: 0.486, A: 1}
	ColorTundora                        = Color{R: 0.290, G: 0.259, B: 0.267, A: 1}
	ColorChileanHeath                   = Color{R: 1.000, G: 0.992, B: 0.902, A: 1}
	ColorFlamePea                       = Color{R: 0.855, G: 0.357, B: 0.220, A: 1}
	ColorPantoneOrange                  = Color{R: 1.000, G: 0.345, B: 0.000, A: 1}
	ColorWildWatermelon                 = Color{R: 0.988, G: 0.424, B: 0.522, A: 1}
	ColorSapGreen                       = Color{R: 0.314, G: 0.490, B: 0.165, A: 1}
	ColorJaponica                       = Color{R: 0.847, G: 0.486, B: 0.388, A: 1}
	ColorKeppel                         = Color{R: 0.227, G: 0.690, B: 0.620, A: 1}
	ColorPortage                        = Color{R: 0.545, G: 0.624, B: 0.933, A: 1}
	ColorBahamaBlue                     = Color{R: 0.008, G: 0.388, B: 0.584, A: 1}
	ColorBurnham                        = Color{R: 0.000, G: 0.180, B: 0.125, A: 1}
	ColorRedRibbon                      = Color{R: 0.929, G: 0.039, B: 0.247, A: 1}
	ColorMagentaPink                    = Color{R: 0.800, G: 0.200, B: 0.545, A: 1}
	ColorRedDevil                       = Color{R: 0.525, G: 0.004, B: 0.067, A: 1}
	ColorHanBlue                        = Color{R: 0.267, G: 0.424, B: 0.812, A: 1}
	ColorLily                           = Color{R: 0.784, G: 0.667, B: 0.749, A: 1}
	ColorLinkWater                      = Color{R: 0.851, G: 0.894, B: 0.961, A: 1}
	ColorRedRobin                       = Color{R: 0.502, G: 0.204, B: 0.122, A: 1}
	ColorSizzlingSunrise                = Color{R: 1.000, G: 0.859, B: 0.000, A: 1}
	ColorHalfColonialWhite              = Color{R: 0.992, G: 0.965, B: 0.827, A: 1}
	ColorJapaneseViolet                 = Color{R: 0.357, G: 0.196, B: 0.337, A: 1}
	ColorNeonGreen                      = Color{R: 0.224, G: 1.000, B: 0.078, A: 1}
	ColorCocoaBean                      = Color{R: 0.282, G: 0.110, B: 0.110, A: 1}
	ColorCinnamonSatin                  = Color{R: 0.804, G: 0.376, B: 0.494, A: 1}
	ColorBlueCharcoal                   = Color{R: 0.004, G: 0.051, B: 0.102, A: 1}
	ColorBurlywood                      = Color{R: 0.871, G: 0.722, B: 0.529, A: 1}
	ColorCanary                         = Color{R: 0.953, G: 0.984, B: 0.384, A: 1}
	ColorRoseBonbon                     = Color{R: 0.976, G: 0.259, B: 0.620, A: 1}
	ColorSaratoga                       = Color{R: 0.333, G: 0.357, B: 0.063, A: 1}
	ColorPantoneYellow                  = Color{R: 0.996, G: 0.875, B: 0.000, A: 1}
	ColorTangaroa                       = Color{R: 0.012, G: 0.086, B: 0.235, A: 1}
	ColorBarossa                        = Color{R: 0.267, G: 0.004, B: 0.176, A: 1}
	ColorGivry                          = Color{R: 0.973, G: 0.894, B: 0.749, A: 1}
	ColorMortar                         = Color{R: 0.314, G: 0.263, B: 0.318, A: 1}
	ColorOldBurgundy                    = Color{R: 0.263, G: 0.188, B: 0.180, A: 1}
	ColorAquaSpring                     = Color{R: 0.918, G: 0.976, B: 0.961, A: 1}
	ColorCascade                        = Color{R: 0.545, G: 0.663, B: 0.647, A: 1}
	ColorDarkTangerine                  = Color{R: 1.000, G: 0.659, B: 0.071, A: 1}
	ColorFlirt                          = Color{R: 0.635, G: 0.000, B: 0.427, A: 1}
	ColorLaRioja                        = Color{R: 0.702, G: 0.757, B: 0.063, A: 1}
	ColorBronzeYellow                   = Color{R: 0.451, G: 0.439, B: 0.000, A: 1}
	ColorCerulean                       = Color{R: 0.000, G: 0.482, B: 0.655, A: 1}
	ColorDenimBlue                      = Color{R: 0.133, G: 0.263, B: 0.714, A: 1}
	ColorRoseTaupe                      = Color{R: 0.565, G: 0.365, B: 0.365, A: 1}
	ColorSpanishBistre                  = Color{R: 0.502, G: 0.459, B: 0.196, A: 1}
	ColorWedgewood                      = Color{R: 0.306, G: 0.498, B: 0.620, A: 1}
	ColorDeepPuce                       = Color{R: 0.663, G: 0.361, B: 0.408, A: 1}
	ColorLightSalmon                    = Color{R: 1.000, G: 0.627, B: 0.478, A: 1}
	ColorLimerick                       = Color{R: 0.616, G: 0.761, B: 0.035, A: 1}
	ColorAnzac                          = Color{R: 0.878, G: 0.714, B: 0.275, A: 1}
	ColorLochinvar                      = Color{R: 0.173, G: 0.549, B: 0.518, A: 1}
	ColorNeonFuchsia                    = Color{R: 0.996, G: 0.255, B: 0.392, A: 1}
	ColorPapayaWhip                     = Color{R: 1.000, G: 0.937, B: 0.835, A: 1}
	ColorPlatinum                       = Color{R: 0.898, G: 0.894, B: 0.886, A: 1}
	ColorTropicalViolet                 = Color{R: 0.804, G: 0.643, B: 0.871, A: 1}
	ColorAsparagus                      = Color{R: 0.529, G: 0.663, B: 0.420, A: 1}
	ColorDesaturatedCyan                = Color{R: 0.400, G: 0.600, B: 0.600, A: 1}
	ColorLasPalmas                      = Color{R: 0.776, G: 0.902, B: 0.063, A: 1}
	ColorVerdunGreen                    = Color{R: 0.286, G: 0.329, B: 0.000, A: 1}
	ColorLenurple                       = Color{R: 0.729, G: 0.576, B: 0.847, A: 1}
	ColorMagicMint                      = Color{R: 0.667, G: 0.941, B: 0.820, A: 1}
	ColorNadeshikoPink                  = Color{R: 0.965, G: 0.678, B: 0.776, A: 1}
	ColorSoap                           = Color{R: 0.808, G: 0.784, B: 0.937, A: 1}
	ColorTan                            = Color{R: 0.824, G: 0.706, B: 0.549, A: 1}
	ColorBlueChalk                      = Color{R: 0.945, G: 0.914, B: 1.000, A: 1}
	ColorBourbon                        = Color{R: 0.729, G: 0.435, B: 0.118, A: 1}
	ColorEverglade                      = Color{R: 0.110, G: 0.251, B: 0.180, A: 1}
	ColorVioletEggplant                 = Color{R: 0.600, G: 0.067, B: 0.600, A: 1}
	ColorRoseBudCherry                  = Color{R: 0.502, G: 0.043, B: 0.278, A: 1}
	ColorSpanishRed                     = Color{R: 0.902, G: 0.000, B: 0.149, A: 1}
	ColorRocketMetallic                 = Color{R: 0.541, G: 0.498, B: 0.502, A: 1}
	ColorApple                          = Color{R: 0.310, G: 0.659, B: 0.239, A: 1}
	ColorDarkSlateGray                  = Color{R: 0.184, G: 0.310, B: 0.310, A: 1}
	ColorFlaxSmoke                      = Color{R: 0.482, G: 0.510, B: 0.396, A: 1}
	ColorMoccasin                       = Color{R: 1.000, G: 0.894, B: 0.710, A: 1}
	ColorTahunaSands                    = Color{R: 0.933, G: 0.941, B: 0.784, A: 1}
	ColorAmazon                         = Color{R: 0.231, G: 0.478, B: 0.341, A: 1}
	ColorBlackHaze                      = Color{R: 0.965, G: 0.969, B: 0.969, A: 1}
	ColorChaletGreen                    = Color{R: 0.318, G: 0.431, B: 0.239, A: 1}
	ColorZeus                           = Color{R: 0.161, G: 0.137, B: 0.098, A: 1}
	ColorZumthor                        = Color{R: 0.929, G: 0.965, B: 1.000, A: 1}
	ColorBittersweet                    = Color{R: 0.996, G: 0.435, B: 0.369, A: 1}
	ColorCadetGrey                      = Color{R: 0.569, G: 0.639, B: 0.690, A: 1}
	ColorDarkPastelBlue                 = Color{R: 0.467, G: 0.620, B: 0.796, A: 1}
	ColorBrightTurquoise                = Color{R: 0.031, G: 0.910, B: 0.871, A: 1}
	ColorAztecGold                      = Color{R: 0.765, G: 0.600, B: 0.325, A: 1}
	ColorPipi                           = Color{R: 0.996, G: 0.957, B: 0.800, A: 1}
	ColorTealDeer                       = Color{R: 0.600, G: 0.902, B: 0.702, A: 1}
	ColorPullmanBrown                   = Color{R: 0.392, G: 0.255, B: 0.090, A: 1}
	ColorTuatara                        = Color{R: 0.212, G: 0.208, B: 0.204, A: 1}
	ColorUnmellowYellow                 = Color{R: 1.000, G: 1.000, B: 0.400, A: 1}
	ColorPaleOyster                     = Color{R: 0.596, G: 0.553, B: 0.467, A: 1}
	ColorPastelBlue                     = Color{R: 0.682, G: 0.776, B: 0.812, A: 1}
	ColorReef                           = Color{R: 0.788, G: 1.000, B: 0.635, A: 1}
	ColorSeashell                       = Color{R: 1.000, G: 0.961, B: 0.933, A: 1}
	ColorCharade                        = Color{R: 0.161, G: 0.161, B: 0.216, A: 1}
	ColorDeepCove                       = Color{R: 0.020, G: 0.063, B: 0.251, A: 1}
	ColorDesertSand                     = Color{R: 0.929, G: 0.788, B: 0.686, A: 1}
	ColorSherwoodGreen                  = Color{R: 0.008, G: 0.251, B: 0.173, A: 1}
	ColorWitchHaze                      = Color{R: 1.000, G: 0.988, B: 0.600, A: 1}
	ColorEnglishVermillion              = Color{R: 0.800, G: 0.278, B: 0.294, A: 1}
	ColorLemonLime                      = Color{R: 0.890, G: 1.000, B: 0.000, A: 1}
	ColorLilyWhite                      = Color{R: 0.906, G: 0.973, B: 1.000, A: 1}
	ColorIndiaGreen                     = Color{R: 0.075, G: 0.533, B: 0.031, A: 1}
	ColorSpanishSkyBlue                 = Color{R: 0.000, G: 0.667, B: 0.894, A: 1}
	ColorSpringGreen                    = Color{R: 0.000, G: 1.000, B: 0.498, A: 1}
	ColorSweetBrown                     = Color{R: 0.659, G: 0.216, B: 0.192, A: 1}
	ColorCopperRose                     = Color{R: 0.600, G: 0.400, B: 0.400, A: 1}
	ColorEgyptianBlue                   = Color{R: 0.063, G: 0.204, B: 0.651, A: 1}
	ColorEngineeringInternationalOrange = Color{R: 0.729, G: 0.086, B: 0.047, A: 1}
	ColorQuillGray                      = Color{R: 0.839, G: 0.839, B: 0.820, A: 1}
	ColorRosePink                       = Color{R: 1.000, G: 0.400, B: 0.800, A: 1}
	ColorSapphireBlue                   = Color{R: 0.000, G: 0.404, B: 0.647, A: 1}
	ColorHaiti                          = Color{R: 0.106, G: 0.063, B: 0.208, A: 1}
	ColorMako                           = Color{R: 0.267, G: 0.286, B: 0.329, A: 1}
	ColorSeaNymph                       = Color{R: 0.471, G: 0.639, B: 0.612, A: 1}
	ColorCharlotte                      = Color{R: 0.729, G: 0.933, B: 0.976, A: 1}
	ColorCornflowerLilac                = Color{R: 1.000, G: 0.690, B: 0.675, A: 1}
	ColorGreenHouse                     = Color{R: 0.141, G: 0.314, B: 0.059, A: 1}
	ColorMandalay                       = Color{R: 0.678, G: 0.471, B: 0.106, A: 1}
	ColorNugget                         = Color{R: 0.773, G: 0.600, B: 0.133, A: 1}
	ColorScampi                         = Color{R: 0.404, G: 0.373, B: 0.651, A: 1}
	ColorBronco                         = Color{R: 0.671, G: 0.631, B: 0.588, A: 1}
	ColorCabaret                        = Color{R: 0.851, G: 0.286, B: 0.447, A: 1}
	ColorGuardsmanRed                   = Color{R: 0.729, G: 0.004, B: 0.004, A: 1}
	ColorChino                          = Color{R: 0.808, G: 0.780, B: 0.655, A: 1}
	ColorWindsor                        = Color{R: 0.235, G: 0.031, B: 0.471, A: 1}
	ColorChantilly                      = Color{R: 0.973, G: 0.765, B: 0.875, A: 1}
	ColorDarkBrown                      = Color{R: 0.396, G: 0.263, B: 0.129, A: 1}
	ColorIronstone                      = Color{R: 0.525, G: 0.282, B: 0.235, A: 1}
	ColorMistyRose                      = Color{R: 1.000, G: 0.894, B: 0.882, A: 1}
	ColorRichBrilliantLavender          = Color{R: 0.945, G: 0.655, B: 0.996, A: 1}
	ColorBizarre                        = Color{R: 0.933, G: 0.871, B: 0.855, A: 1}
	ColorEnglishWalnut                  = Color{R: 0.243, G: 0.169, B: 0.137, A: 1}
	ColorFaluRed                        = Color{R: 0.502, G: 0.094, B: 0.094, A: 1}
	ColorMelanie                        = Color{R: 0.894, G: 0.761, B: 0.835, A: 1}
	ColorSeagull                        = Color{R: 0.502, G: 0.800, B: 0.918, A: 1}
	ColorTerraCotta                     = Color{R: 0.886, G: 0.447, B: 0.357, A: 1}
	ColorChablis                        = Color{R: 1.000, G: 0.957, B: 0.953, A: 1}
	ColorDaffodil                       = Color{R: 1.000, G: 1.000, B: 0.192, A: 1}
	ColorGin                            = Color{R: 0.910, G: 0.949, B: 0.922, A: 1}
	ColorRonchi                         = Color{R: 0.925, G: 0.773, B: 0.306, A: 1}
	ColorCaputMortuum                   = Color{R: 0.349, G: 0.153, B: 0.125, A: 1}
	ColorCitrus                         = Color{R: 0.631, G: 0.773, B: 0.039, A: 1}
	ColorCopperPenny                    = Color{R: 0.678, G: 0.435, B: 0.412, A: 1}
	ColorCharcoal                       = Color{R: 0.212, G: 0.271, B: 0.310, A: 1}
	ColorJasmine                        = Color{R: 0.973, G: 0.871, B: 0.494, A: 1}
	ColorDeepBlush                      = Color{R: 0.894, G: 0.463, B: 0.596, A: 1}
	ColorLanguidLavender                = Color{R: 0.839, G: 0.792, B: 0.867, A: 1}
	ColorOffYellow                      = Color{R: 0.996, G: 0.976, B: 0.890, A: 1}
	ColorPalmLeaf                       = Color{R: 0.098, G: 0.200, B: 0.055, A: 1}
	ColorDarkCoral                      = Color{R: 0.804, G: 0.357, B: 0.271, A: 1}
	ColorGinFizz                        = Color{R: 1.000, G: 0.976, B: 0.886, A: 1}
	ColorKabul                          = Color{R: 0.369, G: 0.282, B: 0.243, A: 1}
	ColorCafeNoir                       = Color{R: 0.294, G: 0.212, B: 0.129, A: 1}
	ColorEclipse                        = Color{R: 0.192, G: 0.110, B: 0.090, A: 1}
	ColorVividMulberry                  = Color{R: 0.722, G: 0.047, B: 0.890, A: 1}
	ColorMabel                          = Color{R: 0.851, G: 0.969, B: 1.000, A: 1}
	ColorPakistanGreen                  = Color{R: 0.000, G: 0.400, B: 0.000, A: 1}
	ColorUltramarine                    = Color{R: 0.247, G: 0.000, B: 1.000, A: 1}
	ColorBiscay                         = Color{R: 0.106, G: 0.192, B: 0.384, A: 1}
	ColorCardinGreen                    = Color{R: 0.004, G: 0.212, B: 0.110, A: 1}
	ColorFuzzyWuzzyBrown                = Color{R: 0.769, G: 0.337, B: 0.333, A: 1}
	ColorShakespeare                    = Color{R: 0.306, G: 0.671, B: 0.820, A: 1}
	ColorCannonBlack                    = Color{R: 0.145, G: 0.090, B: 0.024, A: 1}
	ColorCognac                         = Color{R: 0.624, G: 0.220, B: 0.114, A: 1}
	ColorRoyalFuchsia                   = Color{R: 0.792, G: 0.173, B: 0.573, A: 1}
	ColorCabSav                         = Color{R: 0.302, G: 0.039, B: 0.094, A: 1}
	ColorFolly                          = Color{R: 1.000, G: 0.000, B: 0.310, A: 1}
	ColorPeru                           = Color{R: 0.804, G: 0.522, B: 0.247, A: 1}
	ColorTaupeGray                      = Color{R: 0.545, G: 0.522, B: 0.537, A: 1}
	ColorCeladonGreen                   = Color{R: 0.184, G: 0.518, B: 0.486, A: 1}
	ColorDeepOak                        = Color{R: 0.255, G: 0.125, B: 0.063, A: 1}
	ColorEastBay                        = Color{R: 0.255, G: 0.298, B: 0.490, A: 1}
	ColorTamarind                       = Color{R: 0.204, G: 0.082, B: 0.082, A: 1}
	ColorMediumRedViolet                = Color{R: 0.733, G: 0.200, B: 0.522, A: 1}
	ColorMordantRed                     = Color{R: 0.682, G: 0.047, B: 0.000, A: 1}
	ColorSkeptic                        = Color{R: 0.792, G: 0.902, B: 0.855, A: 1}
	ColorUniversityOfCaliforniaGold     = Color{R: 0.718, G: 0.529, B: 0.153, A: 1}
	ColorBleuDeFrance                   = Color{R: 0.192, G: 0.549, B: 0.906, A: 1}
	ColorKingfisherDaisy                = Color{R: 0.243, G: 0.016, B: 0.502, A: 1}
	ColorPanache                        = Color{R: 0.918, G: 0.965, B: 0.933, A: 1}
	ColorRoseRed                        = Color{R: 0.761, G: 0.118, B: 0.337, A: 1}
	ColorCarnabyTan                     = Color{R: 0.361, G: 0.180, B: 0.004, A: 1}
	ColorEdward                         = Color{R: 0.635, G: 0.682, B: 0.671, A: 1}
	ColorJuniper                        = Color{R: 0.427, G: 0.573, B: 0.573, A: 1}
	ColorAsh                            = Color{R: 0.776, G: 0.765, B: 0.710, A: 1}
	ColorColdTurkey                     = Color{R: 0.808, G: 0.729, B: 0.729, A: 1}
	ColorSnowyMint                      = Color{R: 0.839, G: 1.000, B: 0.859, A: 1}
	ColorPermanentGeraniumLake          = Color{R: 0.882, G: 0.173, B: 0.173, A: 1}
	ColorSwamp                          = Color{R: 0.000, G: 0.106, B: 0.110, A: 1}
	ColorBlueGem                        = Color{R: 0.173, G: 0.055, B: 0.549, A: 1}
	ColorDeepSeaGreen                   = Color{R: 0.035, G: 0.345, B: 0.349, A: 1}
	ColorHunterGreen                    = Color{R: 0.208, G: 0.369, B: 0.231, A: 1}
	ColorLynch                          = Color{R: 0.412, G: 0.494, B: 0.604, A: 1}
	ColorPlumpPurple                    = Color{R: 0.349, G: 0.275, B: 0.698, A: 1}
	ColorClairvoyant                    = Color{R: 0.282, G: 0.024, B: 0.337, A: 1}
	ColorCopper                         = Color{R: 0.722, G: 0.451, B: 0.200, A: 1}
	ColorFernGreen                      = Color{R: 0.310, G: 0.475, B: 0.259, A: 1}
	ColorDarkMediumGray                 = Color{R: 0.663, G: 0.663, B: 0.663, A: 1}
	ColorDune                           = Color{R: 0.220, G: 0.208, B: 0.200, A: 1}
	ColorPantonePink                    = Color{R: 0.843, G: 0.282, B: 0.580, A: 1}
	ColorIlluminatingEmerald            = Color{R: 0.192, G: 0.569, B: 0.467, A: 1}
	ColorKaitokeGreen                   = Color{R: 0.000, G: 0.275, B: 0.125, A: 1}
	ColorBananaYellow                   = Color{R: 1.000, G: 0.882, B: 0.208, A: 1}
	ColorEndeavour                      = Color{R: 0.000, G: 0.337, B: 0.655, A: 1}
	ColorFountainBlue                   = Color{R: 0.337, G: 0.706, B: 0.745, A: 1}
	ColorFrostee                        = Color{R: 0.894, G: 0.965, B: 0.906, A: 1}
	ColorPantoneMagenta                 = Color{R: 0.816, G: 0.255, B: 0.494, A: 1}
	ColorStormDust                      = Color{R: 0.392, G: 0.392, B: 0.388, A: 1}
	ColorSunglo                         = Color{R: 0.882, G: 0.408, B: 0.396, A: 1}
	ColorDallas                         = Color{R: 0.431, G: 0.294, B: 0.149, A: 1}
	ColorDeepMaroon                     = Color{R: 0.510, G: 0.000, B: 0.000, A: 1}
	ColorFrenchRaspberry                = Color{R: 0.780, G: 0.173, B: 0.282, A: 1}
	ColorTurquoiseBlue                  = Color{R: 0.000, G: 1.000, B: 0.937, A: 1}
	ColorGoldenFizz                     = Color{R: 0.961, G: 0.984, B: 0.239, A: 1}
	ColorMauvelous                      = Color{R: 0.937, G: 0.596, B: 0.667, A: 1}
	ColorPicasso                        = Color{R: 1.000, G: 0.953, B: 0.616, A: 1}
	ColorQuicksand                      = Color{R: 0.741, G: 0.592, B: 0.557, A: 1}
	ColorSpanishGray                    = Color{R: 0.596, G: 0.596, B: 0.596, A: 1}
	ColorYellow                         = Color{R: 1.000, G: 1.000, B: 0.000, A: 1}
	ColorAquaHaze                       = Color{R: 0.929, G: 0.961, B: 0.961, A: 1}
	ColorDarkJungleGreen                = Color{R: 0.102, G: 0.141, B: 0.129, A: 1}
	ColorLightTurquoise                 = Color{R: 0.686, G: 0.933, B: 0.933, A: 1}
	ColorTabasco                        = Color{R: 0.627, G: 0.153, B: 0.071, A: 1}
	ColorVividBurgundy                  = Color{R: 0.624, G: 0.114, B: 0.208, A: 1}
	ColorDeepGreen                      = Color{R: 0.020, G: 0.400, B: 0.031, A: 1}
	ColorLightCarminePink               = Color{R: 0.902, G: 0.404, B: 0.443, A: 1}
	ColorVisVis                         = Color{R: 1.000, G: 0.937, B: 0.631, A: 1}
	ColorMediumSlateBlue                = Color{R: 0.482, G: 0.408, B: 0.933, A: 1}
	ColorConcord                        = Color{R: 0.486, G: 0.482, B: 0.478, A: 1}
	ColorGrannyApple                    = Color{R: 0.835, G: 0.965, B: 0.890, A: 1}
	ColorKournikova                     = Color{R: 1.000, G: 0.906, B: 0.447, A: 1}
	ColorRock                           = Color{R: 0.302, G: 0.220, B: 0.200, A: 1}
	ColorTowerGray                      = Color{R: 0.663, G: 0.741, B: 0.749, A: 1}
	ColorRYBRed                         = Color{R: 0.996, G: 0.153, B: 0.071, A: 1}
	ColorRipePlum                       = Color{R: 0.255, G: 0.000, B: 0.337, A: 1}
	ColorWaiouru                        = Color{R: 0.212, G: 0.235, B: 0.051, A: 1}
	ColorFashionFuchsia                 = Color{R: 0.957, G: 0.000, B: 0.631, A: 1}
	ColorJaguar                         = Color{R: 0.031, G: 0.004, B: 0.063, A: 1}
	ColorMeteor                         = Color{R: 0.816, G: 0.490, B: 0.071, A: 1}
	ColorDanube                         = Color{R: 0.376, G: 0.576, B: 0.820, A: 1}
	ColorLavenderPurple                 = Color{R: 0.588, G: 0.482, B: 0.714, A: 1}
	ColorNeptune                        = Color{R: 0.486, G: 0.718, B: 0.733, A: 1}
	ColorCornField                      = Color{R: 0.973, G: 0.980, B: 0.804, A: 1}
	ColorJacaranda                      = Color{R: 0.180, G: 0.012, B: 0.161, A: 1}
	ColorPottersClay                    = Color{R: 0.549, G: 0.341, B: 0.220, A: 1}
	ColorLocust                         = Color{R: 0.659, G: 0.686, B: 0.557, A: 1}
	ColorSeance                         = Color{R: 0.451, G: 0.118, B: 0.561, A: 1}
	ColorTacha                          = Color{R: 0.839, G: 0.773, B: 0.384, A: 1}
	ColorVanilla                        = Color{R: 0.953, G: 0.898, B: 0.671, A: 1}
	ColorBonJour                        = Color{R: 0.898, G: 0.878, B: 0.882, A: 1}
	ColorCoralRed                       = Color{R: 1.000, G: 0.251, B: 0.251, A: 1}
	ColorGrannySmith                    = Color{R: 0.518, G: 0.627, B: 0.627, A: 1}
	ColorRhino                          = Color{R: 0.180, G: 0.247, B: 0.384, A: 1}
	ColorUPForestGreen                  = Color{R: 0.004, G: 0.267, B: 0.129, A: 1}
	ColorAmber                          = Color{R: 1.000, G: 0.749, B: 0.000, A: 1}
	ColorClementine                     = Color{R: 0.914, G: 0.431, B: 0.000, A: 1}
	ColorJetStream                      = Color{R: 0.710, G: 0.824, B: 0.808, A: 1}
	ColorSatinSheenGold                 = Color{R: 0.796, G: 0.631, B: 0.208, A: 1}
	ColorEcruWhite                      = Color{R: 0.961, G: 0.953, B: 0.898, A: 1}
	ColorMidnight                       = Color{R: 0.439, G: 0.149, B: 0.439, A: 1}
	ColorPlum                           = Color{R: 0.557, G: 0.271, B: 0.522, A: 1}
	ColorSun                            = Color{R: 0.984, G: 0.675, B: 0.075, A: 1}
	ColorTurquoise                      = Color{R: 0.251, G: 0.878, B: 0.816, A: 1}
	ColorVermilion                      = Color{R: 0.851, G: 0.220, B: 0.118, A: 1}
	ColorCreamBrulee                    = Color{R: 1.000, G: 0.898, B: 0.627, A: 1}
	ColorParchment                      = Color{R: 0.945, G: 0.914, B: 0.824, A: 1}
	ColorPersianOrange                  = Color{R: 0.851, G: 0.565, B: 0.345, A: 1}
	ColorGoldenTainoi                   = Color{R: 1.000, G: 0.800, B: 0.361, A: 1}
	ColorLightMediumOrchid              = Color{R: 0.827, G: 0.608, B: 0.796, A: 1}
	ColorLuxorGold                      = Color{R: 0.655, G: 0.533, B: 0.173, A: 1}
	ColorPaleSilver                     = Color{R: 0.788, G: 0.753, B: 0.733, A: 1}
	ColorRichMaroon                     = Color{R: 0.690, G: 0.188, B: 0.376, A: 1}
	ColorBush                           = Color{R: 0.051, G: 0.180, B: 0.110, A: 1}
	ColorFrenchSkyBlue                  = Color{R: 0.467, G: 0.710, B: 0.996, A: 1}
	ColorFunGreen                       = Color{R: 0.004, G: 0.427, B: 0.224, A: 1}
	ColorSandal                         = Color{R: 0.667, G: 0.553, B: 0.435, A: 1}
	ColorSpanishGreen                   = Color{R: 0.000, G: 0.569, B: 0.314, A: 1}
	ColorWatercourse                    = Color{R: 0.020, G: 0.435, B: 0.341, A: 1}
	ColorOysterBay                      = Color{R: 0.855, G: 0.980, B: 1.000, A: 1}
	ColorPaleLeaf                       = Color{R: 0.753, G: 0.827, B: 0.725, A: 1}
	ColorTemptress                      = Color{R: 0.231, G: 0.000, B: 0.043, A: 1}
	ColorMetallicCopper                 = Color{R: 0.443, G: 0.161, B: 0.114, A: 1}
	ColorTomThumb                       = Color{R: 0.247, G: 0.345, B: 0.231, A: 1}
	ColorVioletBlue                     = Color{R: 0.196, G: 0.290, B: 0.698, A: 1}
	ColorButterflyBush                  = Color{R: 0.384, G: 0.306, B: 0.604, A: 1}
	ColorCarmine                        = Color{R: 0.588, G: 0.000, B: 0.094, A: 1}
	ColorLightBrilliantRed              = Color{R: 0.996, G: 0.180, B: 0.180, A: 1}
	ColorFuchsiaPurple                  = Color{R: 0.800, G: 0.224, B: 0.482, A: 1}
	ColorPaarl                          = Color{R: 0.651, G: 0.333, B: 0.161, A: 1}
	ColorPeach                          = Color{R: 1.000, G: 0.796, B: 0.643, A: 1}
	ColorLiverChestnut                  = Color{R: 0.596, G: 0.455, B: 0.337, A: 1}
	ColorOldLace                        = Color{R: 0.992, G: 0.961, B: 0.902, A: 1}
	ColorPastelViolet                   = Color{R: 0.796, G: 0.600, B: 0.788, A: 1}
	ColorGhostWhite                     = Color{R: 0.973, G: 0.973, B: 1.000, A: 1}
	ColorGreenWhite                     = Color{R: 0.910, G: 0.922, B: 0.878, A: 1}
	ColorJordyBlue                      = Color{R: 0.541, G: 0.725, B: 0.945, A: 1}
	ColorGladeGreen                     = Color{R: 0.380, G: 0.518, B: 0.373, A: 1}
	ColorPurple                         = Color{R: 0.502, G: 0.000, B: 0.502, A: 1}
	ColorSage                           = Color{R: 0.737, G: 0.722, B: 0.541, A: 1}
	ColorOceanGreen                     = Color{R: 0.282, G: 0.749, B: 0.569, A: 1}
	ColorPaleRedViolet                  = Color{R: 0.859, G: 0.439, B: 0.576, A: 1}
	ColorSeaSerpent                     = Color{R: 0.294, G: 0.780, B: 0.812, A: 1}
	ColorAlizarinCrimson                = Color{R: 0.890, G: 0.149, B: 0.212, A: 1}
	ColorElectricYellow                 = Color{R: 1.000, G: 1.000, B: 0.200, A: 1}
	ColorLividBrown                     = Color{R: 0.302, G: 0.157, B: 0.180, A: 1}
	ColorMediumOrchid                   = Color{R: 0.729, G: 0.333, B: 0.827, A: 1}
	ColorPaleMagentaPink                = Color{R: 1.000, G: 0.600, B: 0.800, A: 1}
	ColorAtoll                          = Color{R: 0.039, G: 0.435, B: 0.459, A: 1}
	ColorFuelYellow                     = Color{R: 0.925, G: 0.663, B: 0.153, A: 1}
	ColorLightSlateGray                 = Color{R: 0.467, G: 0.533, B: 0.600, A: 1}
	ColorBrownDerby                     = Color{R: 0.286, G: 0.149, B: 0.082, A: 1}
	ColorPistachio                      = Color{R: 0.576, G: 0.773, B: 0.447, A: 1}
	ColorVistaWhite                     = Color{R: 0.988, G: 0.973, B: 0.969, A: 1}
	ColorOrangeWhite                    = Color{R: 0.996, G: 0.988, B: 0.929, A: 1}
	ColorRichElectricBlue               = Color{R: 0.031, G: 0.573, B: 0.816, A: 1}
	ColorTartOrange                     = Color{R: 0.984, G: 0.302, B: 0.275, A: 1}
	ColorVividOrange                    = Color{R: 1.000, G: 0.373, B: 0.000, A: 1}
	ColorCapePalliser                   = Color{R: 0.635, G: 0.400, B: 0.271, A: 1}
	ColorGallery                        = Color{R: 0.937, G: 0.937, B: 0.937, A: 1}
	ColorMadang                         = Color{R: 0.718, G: 0.941, B: 0.745, A: 1}
	ColorChinaIvory                     = Color{R: 0.988, G: 1.000, B: 0.906, A: 1}
	ColorDarkPastelPurple               = Color{R: 0.588, G: 0.435, B: 0.839, A: 1}
	ColorFire                           = Color{R: 0.667, G: 0.259, B: 0.012, A: 1}
	ColorKobicha                        = Color{R: 0.420, G: 0.267, B: 0.137, A: 1}
	ColorBlackRussian                   = Color{R: 0.039, G: 0.000, B: 0.110, A: 1}
	ColorBoysenberry                    = Color{R: 0.529, G: 0.196, B: 0.376, A: 1}
	ColorChlorophyllGreen               = Color{R: 0.290, G: 1.000, B: 0.000, A: 1}
	ColorLightDeepPink                  = Color{R: 1.000, G: 0.361, B: 0.804, A: 1}
	ColorPortica                        = Color{R: 0.976, G: 0.902, B: 0.388, A: 1}
	ColorCGRed                          = Color{R: 0.878, G: 0.235, B: 0.192, A: 1}
	ColorDarkCyan                       = Color{R: 0.000, G: 0.545, B: 0.545, A: 1}
	ColorDesert                         = Color{R: 0.682, G: 0.376, B: 0.125, A: 1}
	ColorPurpleTaupe                    = Color{R: 0.314, G: 0.251, B: 0.302, A: 1}
	ColorQuinacridoneMagenta            = Color{R: 0.557, G: 0.227, B: 0.349, A: 1}
	ColorSazerac                        = Color{R: 1.000, G: 0.957, B: 0.878, A: 1}
	ColorSpanishCrimson                 = Color{R: 0.898, G: 0.102, B: 0.298, A: 1}
	ColorCarnelian                      = Color{R: 0.702, G: 0.106, B: 0.106, A: 1}
	ColorCeruleanFrost                  = Color{R: 0.427, G: 0.608, B: 0.765, A: 1}
	ColorMacaroniAndCheese              = Color{R: 1.000, G: 0.741, B: 0.533, A: 1}
	ColorVenetianRed                    = Color{R: 0.784, G: 0.031, B: 0.082, A: 1}
	ColorAnakiwa                        = Color{R: 0.616, G: 0.898, B: 1.000, A: 1}
	ColorConch                          = Color{R: 0.788, G: 0.851, B: 0.824, A: 1}
	ColorMantle                         = Color{R: 0.545, G: 0.612, B: 0.565, A: 1}
	ColorCoralReef                      = Color{R: 0.780, G: 0.737, B: 0.635, A: 1}
	ColorDarkVanilla                    = Color{R: 0.820, G: 0.745, B: 0.659, A: 1}
	ColorJambalaya                      = Color{R: 0.357, G: 0.188, B: 0.075, A: 1}
	ColorPampas                         = Color{R: 0.957, G: 0.949, B: 0.933, A: 1}
	ColorAlabaster                      = Color{R: 0.980, G: 0.980, B: 0.980, A: 1}
	ColorAustralianMint                 = Color{R: 0.961, G: 1.000, B: 0.745, A: 1}
	ColorBarnRed                        = Color{R: 0.486, G: 0.039, B: 0.008, A: 1}
	ColorPolishedPine                   = Color{R: 0.365, G: 0.643, B: 0.576, A: 1}
	ColorPurpleMountainMajesty          = Color{R: 0.588, G: 0.471, B: 0.714, A: 1}
	ColorTangerineYellow                = Color{R: 1.000, G: 0.800, B: 0.000, A: 1}
	ColorThatchGreen                    = Color{R: 0.251, G: 0.239, B: 0.098, A: 1}
	ColorAlabamaCrimson                 = Color{R: 0.686, G: 0.000, B: 0.165, A: 1}
	ColorAzureishWhite                  = Color{R: 0.859, G: 0.914, B: 0.957, A: 1}
	ColorFreshEggplant                  = Color{R: 0.600, G: 0.000, B: 0.400, A: 1}
	ColorGraphite                       = Color{R: 0.145, G: 0.086, B: 0.027, A: 1}
	ColorLemonGlacier                   = Color{R: 0.992, G: 1.000, B: 0.000, A: 1}
	ColorMediumRuby                     = Color{R: 0.667, G: 0.251, B: 0.412, A: 1}
	ColorPaleLavender                   = Color{R: 0.863, G: 0.816, B: 1.000, A: 1}
	ColorPatina                         = Color{R: 0.388, G: 0.604, B: 0.561, A: 1}
	ColorAlto                           = Color{R: 0.859, G: 0.859, B: 0.859, A: 1}
	ColorFallow                         = Color{R: 0.757, G: 0.604, B: 0.420, A: 1}
	ColorGoldenBell                     = Color{R: 0.886, G: 0.537, B: 0.075, A: 1}
	ColorSalem                          = Color{R: 0.035, G: 0.498, B: 0.294, A: 1}
	ColorSandrift                       = Color{R: 0.671, G: 0.569, B: 0.478, A: 1}
	ColorSugarPlum                      = Color{R: 0.569, G: 0.306, B: 0.459, A: 1}
	ColorCameoPink                      = Color{R: 0.937, G: 0.733, B: 0.800, A: 1}
	ColorDarkEbony                      = Color{R: 0.235, G: 0.125, B: 0.020, A: 1}
	ColorGreenMist                      = Color{R: 0.796, G: 0.827, B: 0.690, A: 1}
	ColorShilo                          = Color{R: 0.910, G: 0.725, B: 0.702, A: 1}
	ColorShuttleGray                    = Color{R: 0.373, G: 0.400, B: 0.447, A: 1}
	ColorPalePink                       = Color{R: 0.980, G: 0.855, B: 0.867, A: 1}
	ColorVeronica                       = Color{R: 0.627, G: 0.125, B: 0.941, A: 1}
	ColorAbbey                          = Color{R: 0.298, G: 0.310, B: 0.337, A: 1}
	ColorGreenSheen                     = Color{R: 0.431, G: 0.682, B: 0.631, A: 1}
	ColorMySin                          = Color{R: 1.000, G: 0.702, B: 0.122, A: 1}
	ColorSpray                          = Color{R: 0.475, G: 0.871, B: 0.925, A: 1}
	ColorValentino                      = Color{R: 0.208, G: 0.055, B: 0.259, A: 1}
	ColorCyanBlueAzure                  = Color{R: 0.275, G: 0.510, B: 0.749, A: 1}
	ColorFrost                          = Color{R: 0.929, G: 0.961, B: 0.867, A: 1}
	ColorPalePlum                       = Color{R: 0.867, G: 0.627, B: 0.867, A: 1}
	ColorRegentStBlue                   = Color{R: 0.667, G: 0.839, B: 0.902, A: 1}
	ColorCoconutCream                   = Color{R: 0.973, G: 0.969, B: 0.863, A: 1}
	ColorMoonMist                       = Color{R: 0.863, G: 0.867, B: 0.800, A: 1}
	ColorPaleCopper                     = Color{R: 0.855, G: 0.541, B: 0.404, A: 1}
	ColorScarlet                        = Color{R: 1.000, G: 0.141, B: 0.000, A: 1}
	ColorAmour                          = Color{R: 0.976, G: 0.918, B: 0.953, A: 1}
	ColorBlackSqueeze                   = Color{R: 0.949, G: 0.980, B: 0.980, A: 1}
	ColorCharlestonGreen                = Color{R: 0.137, G: 0.169, B: 0.169, A: 1}
	ColorDarkTurquoise                  = Color{R: 0.000, G: 0.808, B: 0.820, A: 1}
	ColorMillbrook                      = Color{R: 0.349, G: 0.267, B: 0.200, A: 1}
	ColorTreehouse                      = Color{R: 0.231, G: 0.157, B: 0.125, A: 1}
	ColorFrenchPlum                     = Color{R: 0.506, G: 0.078, B: 0.325, A: 1}
	ColorFrostedMint                    = Color{R: 0.859, G: 1.000, B: 0.973, A: 1}
	ColorGreenYellow                    = Color{R: 0.678, G: 1.000, B: 0.184, A: 1}
	ColorJungleGreen                    = Color{R: 0.161, G: 0.671, B: 0.529, A: 1}
	ColorMistGray                       = Color{R: 0.769, G: 0.769, B: 0.737, A: 1}
	ColorCeramic                        = Color{R: 0.988, G: 1.000, B: 0.976, A: 1}
	ColorConfetti                       = Color{R: 0.914, G: 0.843, B: 0.353, A: 1}
	ColorDarkPastelGreen                = Color{R: 0.012, G: 0.753, B: 0.235, A: 1}
	ColorPictorialCarmine               = Color{R: 0.765, G: 0.043, B: 0.306, A: 1}
	ColorWhiskey                        = Color{R: 0.835, G: 0.604, B: 0.435, A: 1}
	ColorVanillaIce                     = Color{R: 0.953, G: 0.561, B: 0.663, A: 1}
	ColorBigStone                       = Color{R: 0.086, G: 0.165, B: 0.251, A: 1}
	ColorChestnut                       = Color{R: 0.584, G: 0.271, B: 0.208, A: 1}
	ColorGreenLeaf                      = Color{R: 0.263, G: 0.416, B: 0.051, A: 1}
	ColorJackoBean                      = Color{R: 0.180, G: 0.098, B: 0.020, A: 1}
	ColorPearlyPurple                   = Color{R: 0.718, G: 0.408, B: 0.635, A: 1}
	ColorSpanishViolet                  = Color{R: 0.298, G: 0.157, B: 0.510, A: 1}
	ColorMalachite                      = Color{R: 0.043, G: 0.855, B: 0.318, A: 1}
	ColorMexicanRed                     = Color{R: 0.655, G: 0.145, B: 0.145, A: 1}
	ColorViridian                       = Color{R: 0.251, G: 0.510, B: 0.427, A: 1}
	ColorSealBrown                      = Color{R: 0.349, G: 0.149, B: 0.043, A: 1}
	ColorAntiqueFuchsia                 = Color{R: 0.569, G: 0.361, B: 0.514, A: 1}
	ColorCuriousBlue                    = Color{R: 0.145, G: 0.588, B: 0.820, A: 1}
	ColorPizazz                         = Color{R: 1.000, G: 0.565, B: 0.000, A: 1}
	ColorCGBlue                         = Color{R: 0.000, G: 0.478, B: 0.647, A: 1}
	ColorDullLavender                   = Color{R: 0.659, G: 0.600, B: 0.902, A: 1}
	ColorPearlMysticTurquoise           = Color{R: 0.196, G: 0.776, B: 0.651, A: 1}
	ColorRuby                           = Color{R: 0.878, G: 0.067, B: 0.373, A: 1}
	ColorUltramarineBlue                = Color{R: 0.255, G: 0.400, B: 0.961, A: 1}
	ColorWaxFlower                      = Color{R: 1.000, G: 0.753, B: 0.659, A: 1}
	ColorSilverChalice                  = Color{R: 0.675, G: 0.675, B: 0.675, A: 1}
	ColorBarbiePink                     = Color{R: 0.878, G: 0.129, B: 0.541, A: 1}
	ColorGreenSpring                    = Color{R: 0.722, G: 0.757, B: 0.694, A: 1}
	ColorOldSilver                      = Color{R: 0.518, G: 0.518, B: 0.510, A: 1}
	ColorGreenWaterloo                  = Color{R: 0.063, G: 0.078, B: 0.020, A: 1}
	ColorShark                          = Color{R: 0.145, G: 0.153, B: 0.173, A: 1}
	ColorZomp                           = Color{R: 0.224, G: 0.655, B: 0.557, A: 1}
	ColorAsphalt                        = Color{R: 0.075, G: 0.039, B: 0.024, A: 1}
	ColorBrightGray                     = Color{R: 0.235, G: 0.255, B: 0.318, A: 1}
	ColorGhost                          = Color{R: 0.780, G: 0.788, B: 0.835, A: 1}
	ColorEunry                          = Color{R: 0.812, G: 0.639, B: 0.616, A: 1}
	ColorPastelPurple                   = Color{R: 0.702, G: 0.620, B: 0.710, A: 1}
	ColorVictoria                       = Color{R: 0.325, G: 0.267, B: 0.569, A: 1}
	ColorBordeaux                       = Color{R: 0.361, G: 0.004, B: 0.125, A: 1}
	ColorCyberGrape                     = Color{R: 0.345, G: 0.259, B: 0.486, A: 1}
	ColorDustStorm                      = Color{R: 0.898, G: 0.800, B: 0.788, A: 1}
	ColorTequila                        = Color{R: 1.000, G: 0.902, B: 0.780, A: 1}
	ColorDarkTan                        = Color{R: 0.569, G: 0.506, B: 0.318, A: 1}
	ColorNapierGreen                    = Color{R: 0.165, G: 0.502, B: 0.000, A: 1}
	ColorOrangeSoda                     = Color{R: 0.980, G: 0.357, B: 0.239, A: 1}
	ColorLogCabin                       = Color{R: 0.141, G: 0.165, B: 0.114, A: 1}
	ColorPavlova                        = Color{R: 0.843, G: 0.769, B: 0.596, A: 1}
	ColorAlgaeGreen                     = Color{R: 0.576, G: 0.875, B: 0.722, A: 1}
	ColorFieldDrab                      = Color{R: 0.424, G: 0.329, B: 0.118, A: 1}
	ColorFrenchMauve                    = Color{R: 0.831, G: 0.451, B: 0.831, A: 1}
	ColorBotticelli                     = Color{R: 0.780, G: 0.867, B: 0.898, A: 1}
	ColorGoldenSand                     = Color{R: 0.941, G: 0.859, B: 0.490, A: 1}
	ColorZinnwaldite                    = Color{R: 0.922, G: 0.761, B: 0.686, A: 1}
	ColorSmokyBlack                     = Color{R: 0.063, G: 0.047, B: 0.031, A: 1}
	ColorTawnyPort                      = Color{R: 0.412, G: 0.145, B: 0.271, A: 1}
	ColorBilobaFlower                   = Color{R: 0.698, G: 0.631, B: 0.918, A: 1}
	ColorCarissma                       = Color{R: 0.918, G: 0.533, B: 0.659, A: 1}
	ColorScreaminGreen                  = Color{R: 0.400, G: 1.000, B: 0.400, A: 1}
	ColorDarkOrange                     = Color{R: 1.000, G: 0.549, B: 0.000, A: 1}
	ColorDoubleSpanishWhite             = Color{R: 0.902, G: 0.843, B: 0.725, A: 1}
	ColorOysterPink                     = Color{R: 0.914, G: 0.808, B: 0.804, A: 1}
	ColorPadua                          = Color{R: 0.678, G: 0.902, B: 0.769, A: 1}
	ColorBeautyBush                     = Color{R: 0.933, G: 0.757, B: 0.745, A: 1}
	ColorCamouflage                     = Color{R: 0.235, G: 0.224, B: 0.063, A: 1}
	ColorCardinalPink                   = Color{R: 0.549, G: 0.020, B: 0.369, A: 1}
	ColorGrayAsparagus                  = Color{R: 0.275, G: 0.349, B: 0.271, A: 1}
	ColorHeath                          = Color{R: 0.329, G: 0.063, B: 0.071, A: 1}
	ColorMagnolia                       = Color{R: 0.973, G: 0.957, B: 1.000, A: 1}
	ColorRubineRed                      = Color{R: 0.820, G: 0.000, B: 0.337, A: 1}
	ColorSlateBlue                      = Color{R: 0.416, G: 0.353, B: 0.804, A: 1}
	ColorAndroidGreen                   = Color{R: 0.643, G: 0.776, B: 0.224, A: 1}
	ColorBrightNavyBlue                 = Color{R: 0.098, G: 0.455, B: 0.824, A: 1}
	ColorEternity                       = Color{R: 0.129, G: 0.102, B: 0.055, A: 1}
	ColorValhalla                       = Color{R: 0.169, G: 0.098, B: 0.310, A: 1}
	ColorSeaPink                        = Color{R: 0.929, G: 0.596, B: 0.620, A: 1}
	ColorShamrock                       = Color{R: 0.200, G: 0.800, B: 0.600, A: 1}
	ColorSpringWood                     = Color{R: 0.973, G: 0.965, B: 0.945, A: 1}
	ColorAppleBlossom                   = Color{R: 0.686, G: 0.302, B: 0.263, A: 1}
	ColorBananaMania                    = Color{R: 0.980, G: 0.906, B: 0.710, A: 1}
	ColorMetallicGold                   = Color{R: 0.831, G: 0.686, B: 0.216, A: 1}
	ColorNewYorkPink                    = Color{R: 0.843, G: 0.514, B: 0.498, A: 1}
	ColorOlivetone                      = Color{R: 0.443, G: 0.431, B: 0.063, A: 1}
	ColorPearlBush                      = Color{R: 0.910, G: 0.878, B: 0.835, A: 1}
	ColorSandyBeach                     = Color{R: 1.000, G: 0.918, B: 0.784, A: 1}
	ColorEspresso                       = Color{R: 0.380, G: 0.153, B: 0.094, A: 1}
	ColorFireEngineRed                  = Color{R: 0.808, G: 0.125, B: 0.161, A: 1}
	ColorMountainMeadow                 = Color{R: 0.188, G: 0.729, B: 0.561, A: 1}
	ColorMystic                         = Color{R: 0.839, G: 0.322, B: 0.510, A: 1}
	ColorRoseVale                       = Color{R: 0.671, G: 0.306, B: 0.322, A: 1}
	ColorSacramentoStateGreen           = Color{R: 0.016, G: 0.224, B: 0.153, A: 1}
	ColorWaterloo                       = Color{R: 0.482, G: 0.486, B: 0.580, A: 1}
	ColorWintergreenDream               = Color{R: 0.337, G: 0.533, B: 0.490, A: 1}
	ColorBunting                        = Color{R: 0.082, G: 0.122, B: 0.298, A: 1}
	ColorDeepTuscanRed                  = Color{R: 0.400, G: 0.259, B: 0.302, A: 1}
	ColorFeldgrau                       = Color{R: 0.302, G: 0.365, B: 0.325, A: 1}
	ColorGoldTips                       = Color{R: 0.871, G: 0.729, B: 0.075, A: 1}
	ColorJava                           = Color{R: 0.122, G: 0.761, B: 0.761, A: 1}
	ColorVolt                           = Color{R: 0.808, G: 1.000, B: 0.000, A: 1}
	ColorStarkWhite                     = Color{R: 0.898, G: 0.843, B: 0.741, A: 1}
	ColorArmadillo                      = Color{R: 0.263, G: 0.243, B: 0.216, A: 1}
	ColorDarkSlateBlue                  = Color{R: 0.282, G: 0.239, B: 0.545, A: 1}
	ColorEdgewater                      = Color{R: 0.784, G: 0.890, B: 0.843, A: 1}
	ColorDarkSpringGreen                = Color{R: 0.090, G: 0.447, B: 0.271, A: 1}
	ColorGreenVogue                     = Color{R: 0.012, G: 0.169, B: 0.322, A: 1}
	ColorSpicyPink                      = Color{R: 0.506, G: 0.431, B: 0.443, A: 1}
	ColorSilverLakeBlue                 = Color{R: 0.365, G: 0.537, B: 0.729, A: 1}
	ColorDawnPink                       = Color{R: 0.953, G: 0.914, B: 0.898, A: 1}
	ColorGreenKelp                      = Color{R: 0.145, G: 0.192, B: 0.110, A: 1}
	ColorLightCobaltBlue                = Color{R: 0.533, G: 0.675, B: 0.878, A: 1}
	ColorBrightCerulean                 = Color{R: 0.114, G: 0.675, B: 0.839, A: 1}
	ColorMintTulip                      = Color{R: 0.769, G: 0.957, B: 0.922, A: 1}
	ColorPeachCream                     = Color{R: 1.000, G: 0.941, B: 0.859, A: 1}
	ColorBlackRose                      = Color{R: 0.404, G: 0.012, B: 0.176, A: 1}
	ColorCedarWoodFinish                = Color{R: 0.443, G: 0.102, B: 0.000, A: 1}
	ColorSunny                          = Color{R: 0.949, G: 0.949, B: 0.478, A: 1}
	ColorHurricane                      = Color{R: 0.529, G: 0.486, B: 0.482, A: 1}
	ColorIslandSpice                    = Color{R: 1.000, G: 0.988, B: 0.933, A: 1}
	ColorBlueRomance                    = Color{R: 0.824, G: 0.965, B: 0.871, A: 1}
	ColorCosmicLatte                    = Color{R: 1.000, G: 0.973, B: 0.906, A: 1}
	ColorHokeyPokey                     = Color{R: 0.784, G: 0.647, B: 0.157, A: 1}
	ColorJapaneseCarmine                = Color{R: 0.616, G: 0.161, B: 0.200, A: 1}
	ColorLemonCurry                     = Color{R: 0.800, G: 0.627, B: 0.114, A: 1}
	ColorLightCoral                     = Color{R: 0.941, G: 0.502, B: 0.502, A: 1}
	ColorSauvignon                      = Color{R: 1.000, G: 0.961, B: 0.953, A: 1}
	ColorSpringSun                      = Color{R: 0.965, G: 1.000, B: 0.863, A: 1}
	ColorBalticSea                      = Color{R: 0.165, G: 0.149, B: 0.188, A: 1}
	ColorCannonPink                     = Color{R: 0.537, G: 0.263, B: 0.404, A: 1}
	ColorDarkGunmetal                   = Color{R: 0.122, G: 0.149, B: 0.165, A: 1}
	ColorPalatinatePurple               = Color{R: 0.408, G: 0.157, B: 0.376, A: 1}
	ColorShadowBlue                     = Color{R: 0.467, G: 0.545, B: 0.647, A: 1}
	ColorBostonBlue                     = Color{R: 0.231, G: 0.569, B: 0.706, A: 1}
	ColorElfGreen                       = Color{R: 0.031, G: 0.514, B: 0.439, A: 1}
	ColorGothic                         = Color{R: 0.427, G: 0.573, B: 0.631, A: 1}
	ColorHighland                       = Color{R: 0.435, G: 0.557, B: 0.388, A: 1}
	ColorMelon                          = Color{R: 0.992, G: 0.737, B: 0.706, A: 1}
	ColorPrussianBlue                   = Color{R: 0.000, G: 0.192, B: 0.325, A: 1}
	ColorBridalHeath                    = Color{R: 1.000, G: 0.980, B: 0.957, A: 1}
	ColorPaleGold                       = Color{R: 0.902, G: 0.745, B: 0.541, A: 1}
	ColorMummysTomb                     = Color{R: 0.510, G: 0.557, B: 0.518, A: 1}
	ColorNiagara                        = Color{R: 0.024, G: 0.631, B: 0.537, A: 1}
	ColorChinaRose                      = Color{R: 0.659, G: 0.318, B: 0.431, A: 1}
	ColorEggSour                        = Color{R: 1.000, G: 0.957, B: 0.867, A: 1}
	ColorJellyBean                      = Color{R: 0.855, G: 0.380, B: 0.306, A: 1}
	ColorTobaccoBrown                   = Color{R: 0.443, G: 0.365, B: 0.278, A: 1}
	ColorBittersweetShimmer             = Color{R: 0.749, G: 0.310, B: 0.318, A: 1}
	ColorGlitter                        = Color{R: 0.902, G: 0.910, B: 0.980, A: 1}
	ColorPaleSpringBud                  = Color{R: 0.925, G: 0.922, B: 0.741, A: 1}
	ColorTurquoiseGreen                 = Color{R: 0.627, G: 0.839, B: 0.706, A: 1}
	ColorLimeade                        = Color{R: 0.435, G: 0.616, B: 0.008, A: 1}
	ColorOliveDrabSeven                 = Color{R: 0.235, G: 0.204, B: 0.122, A: 1}
	ColorSweetPink                      = Color{R: 0.992, G: 0.624, B: 0.635, A: 1}
	ColorHorsesNeck                     = Color{R: 0.376, G: 0.286, B: 0.075, A: 1}
	ColorLightFuchsiaPink               = Color{R: 0.976, G: 0.518, B: 0.937, A: 1}
	ColorMetallicSeaweed                = Color{R: 0.039, G: 0.494, B: 0.549, A: 1}
	ColorAluminium                      = Color{R: 0.663, G: 0.675, B: 0.714, A: 1}
	ColorByzantium                      = Color{R: 0.439, G: 0.161, B: 0.388, A: 1}
	ColorGrayChateau                    = Color{R: 0.635, G: 0.667, B: 0.702, A: 1}
	ColorOrangeRed                      = Color{R: 1.000, G: 0.271, B: 0.000, A: 1}
	ColorWestar                         = Color{R: 0.863, G: 0.851, B: 0.824, A: 1}
	ColorWoodBark                       = Color{R: 0.149, G: 0.067, B: 0.020, A: 1}
	ColorAmaranth                       = Color{R: 0.898, G: 0.169, B: 0.314, A: 1}
	ColorBone                           = Color{R: 0.890, G: 0.855, B: 0.788, A: 1}
	ColorHoneysuckle                    = Color{R: 0.929, G: 0.988, B: 0.518, A: 1}
	ColorIndianYellow                   = Color{R: 0.890, G: 0.659, B: 0.341, A: 1}
	ColorMintJulep                      = Color{R: 0.945, G: 0.933, B: 0.757, A: 1}
	ColorPaleCarmine                    = Color{R: 0.686, G: 0.251, B: 0.208, A: 1}
	ColorSugarCane                      = Color{R: 0.976, G: 1.000, B: 0.965, A: 1}
	ColorThunderbird                    = Color{R: 0.753, G: 0.169, B: 0.094, A: 1}
	ColorBrightLilac                    = Color{R: 0.847, G: 0.569, B: 0.937, A: 1}
	ColorDebianRed                      = Color{R: 0.843, G: 0.039, B: 0.325, A: 1}
	ColorDustyGray                      = Color{R: 0.659, G: 0.596, B: 0.608, A: 1}
	ColorTulipTree                      = Color{R: 0.918, G: 0.702, B: 0.231, A: 1}
	ColorFrangipani                     = Color{R: 1.000, G: 0.871, B: 0.702, A: 1}
	ColorJacarta                        = Color{R: 0.227, G: 0.165, B: 0.416, A: 1}
	ColorPineCone                       = Color{R: 0.427, G: 0.369, B: 0.329, A: 1}
	ColorRussianViolet                  = Color{R: 0.196, G: 0.090, B: 0.302, A: 1}
	ColorBayLeaf                        = Color{R: 0.490, G: 0.663, B: 0.553, A: 1}
	ColorChampagne                      = Color{R: 0.969, G: 0.906, B: 0.808, A: 1}
	ColorFawn                           = Color{R: 0.898, G: 0.667, B: 0.439, A: 1}
	ColorRaspberryPink                  = Color{R: 0.886, G: 0.314, B: 0.596, A: 1}
	ColorTrendyGreen                    = Color{R: 0.486, G: 0.533, B: 0.102, A: 1}
	ColorCinnabar                       = Color{R: 0.890, G: 0.259, B: 0.204, A: 1}
	ColorDeepSpaceSparkle               = Color{R: 0.290, G: 0.392, B: 0.424, A: 1}
	ColorPersianGreen                   = Color{R: 0.000, G: 0.651, B: 0.576, A: 1}
	ColorSepiaSkin                      = Color{R: 0.620, G: 0.357, B: 0.251, A: 1}
	ColorWinterHazel                    = Color{R: 0.835, G: 0.820, B: 0.584, A: 1}
	ColorDarkMossGreen                  = Color{R: 0.290, G: 0.365, B: 0.137, A: 1}
	ColorGraniteGray                    = Color{R: 0.404, G: 0.404, B: 0.404, A: 1}
	ColorRYBYellow                      = Color{R: 0.996, G: 0.996, B: 0.200, A: 1}
	ColorPastelPink                     = Color{R: 0.871, G: 0.647, B: 0.643, A: 1}
	ColorRedBerry                       = Color{R: 0.557, G: 0.000, B: 0.000, A: 1}
	ColorSolitude                       = Color{R: 0.918, G: 0.965, B: 1.000, A: 1}
	ColorTimberwolf                     = Color{R: 0.859, G: 0.843, B: 0.824, A: 1}
	ColorUSAFABlue                      = Color{R: 0.000, G: 0.310, B: 0.596, A: 1}
	ColorCrowshead                      = Color{R: 0.110, G: 0.071, B: 0.031, A: 1}
	ColorImperialBlue                   = Color{R: 0.000, G: 0.137, B: 0.584, A: 1}
	ColorMerlin                         = Color{R: 0.255, G: 0.235, B: 0.216, A: 1}
	ColorMeatBrown                      = Color{R: 0.898, G: 0.718, B: 0.231, A: 1}
	ColorRegentGray                     = Color{R: 0.525, G: 0.580, B: 0.624, A: 1}
	ColorRoyalBlue                      = Color{R: 0.255, G: 0.412, B: 0.882, A: 1}
	ColorAmaranthRed                    = Color{R: 0.827, G: 0.129, B: 0.176, A: 1}
	ColorAo                             = Color{R: 0.000, G: 0.502, B: 0.000, A: 1}
	ColorMauveTaupe                     = Color{R: 0.569, G: 0.373, B: 0.427, A: 1}
	ColorRedwood                        = Color{R: 0.643, G: 0.353, B: 0.322, A: 1}
	ColorSatinLinen                     = Color{R: 0.902, G: 0.894, B: 0.831, A: 1}
	ColorClayAsh                        = Color{R: 0.741, G: 0.784, B: 0.702, A: 1}
	ColorParsley                        = Color{R: 0.075, G: 0.310, B: 0.098, A: 1}
	ColorRawSienna                      = Color{R: 0.839, G: 0.541, B: 0.349, A: 1}
	ColorPeruTan                        = Color{R: 0.498, G: 0.227, B: 0.008, A: 1}
	ColorSuvaGray                       = Color{R: 0.533, G: 0.514, B: 0.529, A: 1}
	ColorHoki                           = Color{R: 0.396, G: 0.525, B: 0.624, A: 1}
	ColorMughalGreen                    = Color{R: 0.188, G: 0.376, B: 0.188, A: 1}
	ColorNightShadz                     = Color{R: 0.667, G: 0.216, B: 0.353, A: 1}
	ColorSpanishBlue                    = Color{R: 0.000, G: 0.439, B: 0.722, A: 1}
	ColorVeryLightBlue                  = Color{R: 0.400, G: 0.400, B: 1.000, A: 1}
	ColorFinn                           = Color{R: 0.412, G: 0.176, B: 0.329, A: 1}
	ColorGreenCyan                      = Color{R: 0.000, G: 0.600, B: 0.400, A: 1}
	ColorRedBeech                       = Color{R: 0.482, G: 0.220, B: 0.004, A: 1}
	ColorGray                           = Color{R: 0.502, G: 0.502, B: 0.502, A: 1}
	ColorGumbo                          = Color{R: 0.486, G: 0.631, B: 0.651, A: 1}
	ColorInchWorm                       = Color{R: 0.690, G: 0.890, B: 0.075, A: 1}
	ColorSeaMist                        = Color{R: 0.773, G: 0.859, B: 0.792, A: 1}
	ColorBlackRock                      = Color{R: 0.051, G: 0.012, B: 0.196, A: 1}
	ColorCinereous                      = Color{R: 0.596, G: 0.506, B: 0.482, A: 1}
	ColorDarkChestnut                   = Color{R: 0.596, G: 0.412, B: 0.376, A: 1}
	ColorFrenchRose                     = Color{R: 0.965, G: 0.290, B: 0.541, A: 1}
	ColorHoneyFlower                    = Color{R: 0.310, G: 0.110, B: 0.439, A: 1}
	ColorMediumSkyBlue                  = Color{R: 0.502, G: 0.855, B: 0.922, A: 1}
	ColorWildOrchid                     = Color{R: 0.831, G: 0.439, B: 0.635, A: 1}
	ColorAmulet                         = Color{R: 0.482, G: 0.624, B: 0.502, A: 1}
	ColorChelseaCucumber                = Color{R: 0.514, G: 0.667, B: 0.365, A: 1}
	ColorSpaceCadet                     = Color{R: 0.114, G: 0.161, B: 0.318, A: 1}
	ColorMagentaHaze                    = Color{R: 0.624, G: 0.271, B: 0.463, A: 1}
	ColorMatisse                        = Color{R: 0.106, G: 0.396, B: 0.616, A: 1}
	ColorOnion                          = Color{R: 0.184, G: 0.153, B: 0.055, A: 1}
	ColorPaleBrown                      = Color{R: 0.596, G: 0.463, B: 0.329, A: 1}
	ColorTuscanTan                      = Color{R: 0.651, G: 0.482, B: 0.357, A: 1}
	ColorBermudaGray                    = Color{R: 0.420, G: 0.545, B: 0.635, A: 1}
	ColorGrenadier                      = Color{R: 0.835, G: 0.275, B: 0.000, A: 1}
	ColorHolly                          = Color{R: 0.004, G: 0.114, B: 0.075, A: 1}
	ColorVeryPaleYellow                 = Color{R: 1.000, G: 1.000, B: 0.749, A: 1}
	ColorDeepPink                       = Color{R: 1.000, G: 0.078, B: 0.576, A: 1}
	ColorDonkeyBrown                    = Color{R: 0.400, G: 0.298, B: 0.157, A: 1}
	ColorCrayolaRed                     = Color{R: 0.933, G: 0.125, B: 0.302, A: 1}
	ColorDarkTerraCotta                 = Color{R: 0.800, G: 0.306, B: 0.361, A: 1}
	ColorHitPink                        = Color{R: 1.000, G: 0.671, B: 0.506, A: 1}
	ColorRedDamask                      = Color{R: 0.855, G: 0.416, B: 0.255, A: 1}
	ColorRuber                          = Color{R: 0.808, G: 0.275, B: 0.463, A: 1}
	ColorBirdFlower                     = Color{R: 0.831, G: 0.804, B: 0.086, A: 1}
	ColorBubbleGum                      = Color{R: 1.000, G: 0.757, B: 0.800, A: 1}
	ColorCoolGrey                       = Color{R: 0.549, G: 0.573, B: 0.675, A: 1}
	ColorWeldonBlue                     = Color{R: 0.486, G: 0.596, B: 0.671, A: 1}
	ColorSunflower                      = Color{R: 0.894, G: 0.831, B: 0.133, A: 1}
	ColorFuzzyWuzzy                     = Color{R: 0.800, G: 0.400, B: 0.400, A: 1}
	ColorOnahau                         = Color{R: 0.804, G: 0.957, B: 1.000, A: 1}
	ColorSeaBuckthorn                   = Color{R: 0.984, G: 0.631, B: 0.161, A: 1}
	ColorMelrose                        = Color{R: 0.780, G: 0.757, B: 1.000, A: 1}
	ColorPurpleHeart                    = Color{R: 0.412, G: 0.208, B: 0.612, A: 1}
	ColorAero                           = Color{R: 0.486, G: 0.725, B: 0.910, A: 1}
	ColorAstral                         = Color{R: 0.196, G: 0.490, B: 0.627, A: 1}
	ColorCerise                         = Color{R: 0.871, G: 0.192, B: 0.388, A: 1}
	ColorMonaLisa                       = Color{R: 1.000, G: 0.631, B: 0.580, A: 1}
	ColorOliveDrab                      = Color{R: 0.420, G: 0.557, B: 0.137, A: 1}
	ColorPelorous                       = Color{R: 0.243, G: 0.671, B: 0.749, A: 1}
	ColorCasper                         = Color{R: 0.678, G: 0.745, B: 0.820, A: 1}
	ColorChatelle                       = Color{R: 0.741, G: 0.702, B: 0.780, A: 1}
	ColorGOGreen                        = Color{R: 0.000, G: 0.671, B: 0.400, A: 1}
	ColorSnuff                          = Color{R: 0.886, G: 0.847, B: 0.929, A: 1}
	ColorSoftAmber                      = Color{R: 0.820, G: 0.776, B: 0.706, A: 1}
	ColorTeaRose                        = Color{R: 0.957, G: 0.761, B: 0.761, A: 1}
	ColorBoulder                        = Color{R: 0.478, G: 0.478, B: 0.478, A: 1}
	ColorFrenchPuce                     = Color{R: 0.306, G: 0.086, B: 0.035, A: 1}
	ColorRubyRed                        = Color{R: 0.608, G: 0.067, B: 0.118, A: 1}
	ColorSunset                         = Color{R: 0.980, G: 0.839, B: 0.647, A: 1}
	ColorCabbagePont                    = Color{R: 0.247, G: 0.298, B: 0.227, A: 1}
	ColorIcterine                       = Color{R: 0.988, G: 0.969, B: 0.369, A: 1}
	ColorMatterhorn                     = Color{R: 0.306, G: 0.231, B: 0.255, A: 1}
	ColorLondonHue                      = Color{R: 0.745, G: 0.651, B: 0.765, A: 1}
	ColorPink                           = Color{R: 1.000, G: 0.753, B: 0.796, A: 1}
	ColorPrincessPerfume                = Color{R: 1.000, G: 0.522, B: 0.812, A: 1}
	ColorSelectiveYellow                = Color{R: 1.000, G: 0.729, B: 0.000, A: 1}
	ColorTidal                          = Color{R: 0.945, G: 1.000, B: 0.678, A: 1}
	ColorCharm                          = Color{R: 0.831, G: 0.455, B: 0.580, A: 1}
	ColorElSalva                        = Color{R: 0.561, G: 0.243, B: 0.200, A: 1}
	ColorFrenchWine                     = Color{R: 0.675, G: 0.118, B: 0.267, A: 1}
	ColorWewak                          = Color{R: 0.945, G: 0.608, B: 0.671, A: 1}
	ColorElectricPurple                 = Color{R: 0.749, G: 0.000, B: 1.000, A: 1}
	ColorMatrix                         = Color{R: 0.690, G: 0.365, B: 0.329, A: 1}
	ColorBrandeisBlue                   = Color{R: 0.000, G: 0.439, B: 1.000, A: 1}
	ColorGossamer                       = Color{R: 0.024, G: 0.608, B: 0.506, A: 1}
	ColorEasternBlue                    = Color{R: 0.118, G: 0.604, B: 0.690, A: 1}
	ColorNonPhotoBlue                   = Color{R: 0.643, G: 0.867, B: 0.929, A: 1}
	ColorChromeYellow                   = Color{R: 1.000, G: 0.655, B: 0.000, A: 1}
	ColorCranberry                      = Color{R: 0.859, G: 0.314, B: 0.475, A: 1}
	ColorCyberYellow                    = Color{R: 1.000, G: 0.827, B: 0.000, A: 1}
	ColorMunsellBlue                    = Color{R: 0.000, G: 0.576, B: 0.686, A: 1}
	ColorSalmonPink                     = Color{R: 1.000, G: 0.569, B: 0.643, A: 1}
	ColorTealGreen                      = Color{R: 0.000, G: 0.510, B: 0.498, A: 1}
	ColorBlueDianne                     = Color{R: 0.125, G: 0.282, B: 0.322, A: 1}
	ColorDeepSapphire                   = Color{R: 0.031, G: 0.145, B: 0.404, A: 1}
	ColorIceberg                        = Color{R: 0.443, G: 0.651, B: 0.824, A: 1}
	ColorFrenchPink                     = Color{R: 0.992, G: 0.424, B: 0.620, A: 1}
	ColorPigPink                        = Color{R: 0.992, G: 0.843, B: 0.894, A: 1}
	ColorCopperCanyon                   = Color{R: 0.494, G: 0.227, B: 0.082, A: 1}
	ColorCountyGreen                    = Color{R: 0.004, G: 0.216, B: 0.102, A: 1}
	ColorDogs                           = Color{R: 0.722, G: 0.427, B: 0.161, A: 1}
	ColorSpringRain                     = Color{R: 0.675, G: 0.796, B: 0.694, A: 1}
	ColorDarkBurgundy                   = Color{R: 0.467, G: 0.059, B: 0.020, A: 1}
	ColorEmerald                        = Color{R: 0.314, G: 0.784, B: 0.471, A: 1}
	ColorKangaroo                       = Color{R: 0.776, G: 0.784, B: 0.741, A: 1}
	ColorGrape                          = Color{R: 0.435, G: 0.176, B: 0.659, A: 1}
	ColorLochmara                       = Color{R: 0.000, G: 0.494, B: 0.780, A: 1}
	ColorPumpkin                        = Color{R: 1.000, G: 0.459, B: 0.094, A: 1}
	ColorBrightYellow                   = Color{R: 1.000, G: 0.667, B: 0.114, A: 1}
	ColorChamoisee                      = Color{R: 0.627, G: 0.471, B: 0.353, A: 1}
	ColorChathamsBlue                   = Color{R: 0.090, G: 0.333, B: 0.475, A: 1}
	ColorVividRedTangelo                = Color{R: 0.875, G: 0.380, B: 0.141, A: 1}
	ColorBrinkPink                      = Color{R: 0.984, G: 0.376, B: 0.498, A: 1}
	ColorSteelPink                      = Color{R: 0.800, G: 0.200, B: 0.800, A: 1}
	ColorTexas                          = Color{R: 0.973, G: 0.976, B: 0.612, A: 1}
	ColorHacienda                       = Color{R: 0.596, G: 0.506, B: 0.106, A: 1}
	ColorCruise                         = Color{R: 0.710, G: 0.925, B: 0.875, A: 1}
	ColorDeepGreenCyanTurquoise         = Color{R: 0.055, G: 0.486, B: 0.380, A: 1}
	ColorDingley                        = Color{R: 0.365, G: 0.467, B: 0.278, A: 1}
	ColorBronze                         = Color{R: 0.804, G: 0.498, B: 0.196, A: 1}
	ColorTangoPink                      = Color{R: 0.894, G: 0.443, B: 0.478, A: 1}
	ColorViridianGreen                  = Color{R: 0.000, G: 0.588, B: 0.596, A: 1}
	ColorNewCar                         = Color{R: 0.129, G: 0.310, B: 0.776, A: 1}
	ColorStrawberry                     = Color{R: 0.988, G: 0.353, B: 0.553, A: 1}
	ColorArylideYellow                  = Color{R: 0.914, G: 0.839, B: 0.420, A: 1}
	ColorDeepFir                        = Color{R: 0.000, G: 0.161, B: 0.000, A: 1}
	ColorKombuGreen                     = Color{R: 0.208, G: 0.259, B: 0.188, A: 1}
	ColorRoseGold                       = Color{R: 0.718, G: 0.431, B: 0.475, A: 1}
	ColorSchist                         = Color{R: 0.663, G: 0.706, B: 0.592, A: 1}
	ColorSwampGreen                     = Color{R: 0.675, G: 0.718, B: 0.557, A: 1}
	ColorSwissCoffee                    = Color{R: 0.867, G: 0.839, B: 0.835, A: 1}
	ColorTomato                         = Color{R: 1.000, G: 0.388, B: 0.278, A: 1}
	ColorDavysGrey                      = Color{R: 0.333, G: 0.333, B: 0.333, A: 1}
	ColorFeijoa                         = Color{R: 0.624, G: 0.867, B: 0.549, A: 1}
	ColorRacingGreen                    = Color{R: 0.047, G: 0.098, B: 0.067, A: 1}
	ColorGrandis                        = Color{R: 1.000, G: 0.827, B: 0.549, A: 1}
	ColorMediumSpringGreen              = Color{R: 0.000, G: 0.980, B: 0.604, A: 1}
	ColorGamboge                        = Color{R: 0.894, G: 0.608, B: 0.059, A: 1}
	ColorShocking                       = Color{R: 0.886, G: 0.573, B: 0.753, A: 1}
	ColorTonysPink                      = Color{R: 0.906, G: 0.624, B: 0.549, A: 1}
	ColorWhiteRock                      = Color{R: 0.918, G: 0.910, B: 0.831, A: 1}
	ColorChenin                         = Color{R: 0.875, G: 0.804, B: 0.435, A: 1}
	ColorColonialWhite                  = Color{R: 1.000, G: 0.929, B: 0.737, A: 1}
	ColorDarkPastelRed                  = Color{R: 0.761, G: 0.231, B: 0.133, A: 1}
	ColorMoonRaker                      = Color{R: 0.839, G: 0.808, B: 0.965, A: 1}
	ColorTuftBush                       = Color{R: 1.000, G: 0.867, B: 0.804, A: 1}
	ColorVividGamboge                   = Color{R: 1.000, G: 0.600, B: 0.000, A: 1}
	ColorCraterBrown                    = Color{R: 0.275, G: 0.141, B: 0.145, A: 1}
	ColorFinch                          = Color{R: 0.384, G: 0.400, B: 0.286, A: 1}
	ColorLilac                          = Color{R: 0.784, G: 0.635, B: 0.784, A: 1}
	ColorBandicoot                      = Color{R: 0.522, G: 0.518, B: 0.439, A: 1}
	ColorBuccaneer                      = Color{R: 0.384, G: 0.184, B: 0.188, A: 1}
	ColorHemp                           = Color{R: 0.565, G: 0.471, B: 0.455, A: 1}
	ColorColdPurple                     = Color{R: 0.671, G: 0.627, B: 0.851, A: 1}
	ColorPaleCerulean                   = Color{R: 0.608, G: 0.769, B: 0.886, A: 1}
	ColorTotemPole                      = Color{R: 0.600, G: 0.106, B: 0.027, A: 1}
	ColorArrowtown                      = Color{R: 0.580, G: 0.529, B: 0.443, A: 1}
	ColorBrunswickGreen                 = Color{R: 0.106, G: 0.302, B: 0.243, A: 1}
	ColorCedarChest                     = Color{R: 0.788, G: 0.353, B: 0.286, A: 1}
	ColorOrangeRoughy                   = Color{R: 0.769, G: 0.341, B: 0.098, A: 1}
	ColorDeepBronze                     = Color{R: 0.290, G: 0.188, B: 0.016, A: 1}
	ColorLemonGinger                    = Color{R: 0.675, G: 0.620, B: 0.133, A: 1}
	ColorOldLavender                    = Color{R: 0.475, G: 0.408, B: 0.471, A: 1}
	ColorPsychedelicPurple              = Color{R: 0.875, G: 0.000, B: 1.000, A: 1}
	ColorSweetCorn                      = Color{R: 0.984, G: 0.918, B: 0.549, A: 1}
	ColorElectricCrimson                = Color{R: 1.000, G: 0.000, B: 0.247, A: 1}
	ColorFinlandia                      = Color{R: 0.333, G: 0.427, B: 0.337, A: 1}
	ColorFroly                          = Color{R: 0.961, G: 0.459, B: 0.518, A: 1}
	ColorDiSerria                       = Color{R: 0.859, G: 0.600, B: 0.369, A: 1}
	ColorTuftsBlue                      = Color{R: 0.255, G: 0.490, B: 0.757, A: 1}
	ColorWhiteLinen                     = Color{R: 0.973, G: 0.941, B: 0.910, A: 1}
	ColorRobRoy                         = Color{R: 0.918, G: 0.776, B: 0.455, A: 1}
	ColorShimmeringBlush                = Color{R: 0.851, G: 0.525, B: 0.584, A: 1}
	ColorTeaGreen                       = Color{R: 0.816, G: 0.941, B: 0.753, A: 1}
	ColorCeruleanBlue                   = Color{R: 0.165, G: 0.322, B: 0.745, A: 1}
	ColorOasis                          = Color{R: 0.996, G: 0.937, B: 0.808, A: 1}
	ColorPueblo                         = Color{R: 0.490, G: 0.173, B: 0.078, A: 1}
	ColorCork                           = Color{R: 0.251, G: 0.161, B: 0.114, A: 1}
	ColorPear                           = Color{R: 0.820, G: 0.886, B: 0.192, A: 1}
	ColorChicago                        = Color{R: 0.365, G: 0.361, B: 0.345, A: 1}
	ColorHeather                        = Color{R: 0.718, G: 0.765, B: 0.816, A: 1}
	ColorMaximumYellow                  = Color{R: 0.980, G: 0.980, B: 0.216, A: 1}
	ColorPhthaloGreen                   = Color{R: 0.071, G: 0.208, B: 0.141, A: 1}
	ColorThistleGreen                   = Color{R: 0.800, G: 0.792, B: 0.659, A: 1}
	ColorLima                           = Color{R: 0.463, G: 0.741, B: 0.090, A: 1}
	ColorPalmGreen                      = Color{R: 0.035, G: 0.137, B: 0.059, A: 1}
	ColorPantoneBlue                    = Color{R: 0.000, G: 0.094, B: 0.659, A: 1}
	ColorMistyMoss                      = Color{R: 0.733, G: 0.706, B: 0.467, A: 1}
	ColorPickledBean                    = Color{R: 0.431, G: 0.282, B: 0.149, A: 1}
	ColorBarberry                       = Color{R: 0.871, G: 0.843, B: 0.090, A: 1}
	ColorBeauBlue                       = Color{R: 0.737, G: 0.831, B: 0.902, A: 1}
	ColorGunPowder                      = Color{R: 0.255, G: 0.259, B: 0.341, A: 1}
	ColorIslamicGreen                   = Color{R: 0.000, G: 0.565, B: 0.000, A: 1}
	ColorWispPink                       = Color{R: 0.996, G: 0.957, B: 0.973, A: 1}
	ColorSapling                        = Color{R: 0.871, G: 0.831, B: 0.643, A: 1}
	ColorCadmiumYellow                  = Color{R: 1.000, G: 0.965, B: 0.000, A: 1}
	ColorJacksonsPurple                 = Color{R: 0.125, G: 0.125, B: 0.553, A: 1}
	ColorRajah                          = Color{R: 0.984, G: 0.671, B: 0.376, A: 1}
	ColorRomanCoffee                    = Color{R: 0.475, G: 0.365, B: 0.298, A: 1}
	ColorSpanishOrange                  = Color{R: 0.910, G: 0.380, B: 0.000, A: 1}
	ColorDeepTaupe                      = Color{R: 0.494, G: 0.369, B: 0.376, A: 1}
	ColorGoldenrod                      = Color{R: 0.855, G: 0.647, B: 0.125, A: 1}
	ColorNorthTexasGreen                = Color{R: 0.020, G: 0.565, B: 0.200, A: 1}
	ColorShingleFawn                    = Color{R: 0.420, G: 0.306, B: 0.192, A: 1}
	ColorYellowOrange                   = Color{R: 1.000, G: 0.682, B: 0.259, A: 1}
	ColorCareysPink                     = Color{R: 0.824, G: 0.620, B: 0.667, A: 1}
	ColorColumbiaBlue                   = Color{R: 0.769, G: 0.847, B: 0.886, A: 1}
	ColorFirebrick                      = Color{R: 0.698, G: 0.133, B: 0.133, A: 1}
	ColorIndigo                         = Color{R: 0.294, G: 0.000, B: 0.510, A: 1}
	ColorPortGore                       = Color{R: 0.145, G: 0.122, B: 0.310, A: 1}
	ColorVanCleef                       = Color{R: 0.286, G: 0.090, B: 0.047, A: 1}
	ColorAntiqueBronze                  = Color{R: 0.400, G: 0.365, B: 0.118, A: 1}
	ColorFandango                       = Color{R: 0.710, G: 0.200, B: 0.537, A: 1}
	ColorFlamingo                       = Color{R: 0.949, G: 0.333, B: 0.165, A: 1}
	ColorMilanoRed                      = Color{R: 0.722, G: 0.067, B: 0.016, A: 1}
	ColorRichLavender                   = Color{R: 0.655, G: 0.420, B: 0.812, A: 1}
	ColorSahara                         = Color{R: 0.718, G: 0.635, B: 0.078, A: 1}
	ColorBigFootFeet                    = Color{R: 0.910, G: 0.557, B: 0.353, A: 1}
	ColorBrownTumbleweed                = Color{R: 0.216, G: 0.161, B: 0.055, A: 1}
	ColorClinker                        = Color{R: 0.216, G: 0.114, B: 0.035, A: 1}
	ColorConcrete                       = Color{R: 0.949, G: 0.949, B: 0.949, A: 1}
	ColorGeyser                         = Color{R: 0.831, G: 0.875, B: 0.886, A: 1}
	ColorJet                            = Color{R: 0.204, G: 0.204, B: 0.204, A: 1}
	ColorMindaro                        = Color{R: 0.890, G: 0.976, B: 0.533, A: 1}
	ColorPurplePizzazz                  = Color{R: 0.996, G: 0.306, B: 0.855, A: 1}
	ColorAlloyOrange                    = Color{R: 0.769, G: 0.384, B: 0.063, A: 1}
	ColorBlackForest                    = Color{R: 0.043, G: 0.075, B: 0.016, A: 1}
	ColorBroom                          = Color{R: 1.000, G: 0.925, B: 0.075, A: 1}
	ColorWhiteIce                       = Color{R: 0.867, G: 0.976, B: 0.945, A: 1}
	ColorLust                           = Color{R: 0.902, G: 0.125, B: 0.125, A: 1}
	ColorMaroonOak                      = Color{R: 0.322, G: 0.047, B: 0.090, A: 1}
	ColorOldBrick                       = Color{R: 0.565, G: 0.118, B: 0.118, A: 1}
	ColorBlizzardBlue                   = Color{R: 0.639, G: 0.890, B: 0.929, A: 1}
	ColorIndianRed                      = Color{R: 0.804, G: 0.361, B: 0.361, A: 1}
	ColorJustRight                      = Color{R: 0.925, G: 0.804, B: 0.725, A: 1}
	ColorBleachedCedar                  = Color{R: 0.173, G: 0.129, B: 0.200, A: 1}
	ColorBrandy                         = Color{R: 0.871, G: 0.757, B: 0.588, A: 1}
	ColorPinkFlamingo                   = Color{R: 0.988, G: 0.455, B: 0.992, A: 1}
	ColorStTropaz                       = Color{R: 0.176, G: 0.337, B: 0.608, A: 1}
	ColorKarry                          = Color{R: 1.000, G: 0.918, B: 0.831, A: 1}
	ColorRustyNail                      = Color{R: 0.525, G: 0.337, B: 0.039, A: 1}
	ColorSeaGreen                       = Color{R: 0.180, G: 0.545, B: 0.341, A: 1}
	ColorBostonUniversityRed            = Color{R: 0.800, G: 0.000, B: 0.000, A: 1}
	ColorCorn                           = Color{R: 0.906, G: 0.749, B: 0.020, A: 1}
	ColorCottonSeed                     = Color{R: 0.761, G: 0.741, B: 0.714, A: 1}
	ColorRegalBlue                      = Color{R: 0.004, G: 0.247, B: 0.416, A: 1}
	ColorUABlue                         = Color{R: 0.000, G: 0.200, B: 0.667, A: 1}
	ColorYaleBlue                       = Color{R: 0.059, G: 0.302, B: 0.573, A: 1}
	ColorKenyanCopper                   = Color{R: 0.486, G: 0.110, B: 0.020, A: 1}
	ColorLightHotPink                   = Color{R: 1.000, G: 0.702, B: 0.871, A: 1}
	ColorResolutionBlue                 = Color{R: 0.000, G: 0.137, B: 0.529, A: 1}
	ColorSpicyMustard                   = Color{R: 0.455, G: 0.392, B: 0.051, A: 1}
	ColorVividLimeGreen                 = Color{R: 0.651, G: 0.839, B: 0.031, A: 1}
	ColorBlackcurrant                   = Color{R: 0.196, G: 0.161, B: 0.227, A: 1}
	ColorDorado                         = Color{R: 0.420, G: 0.341, B: 0.333, A: 1}
	ColorElectricIndigo                 = Color{R: 0.435, G: 0.000, B: 1.000, A: 1}
	ColorLightMossGreen                 = Color{R: 0.678, G: 0.875, B: 0.678, A: 1}
	ColorNandor                         = Color{R: 0.294, G: 0.365, B: 0.322, A: 1}
	ColorRose                           = Color{R: 1.000, G: 0.000, B: 0.498, A: 1}
	ColorChamois                        = Color{R: 0.929, G: 0.863, B: 0.694, A: 1}
	ColorComet                          = Color{R: 0.361, G: 0.365, B: 0.459, A: 1}
	ColorVividCerise                    = Color{R: 0.855, G: 0.114, B: 0.506, A: 1}
	ColorRed                            = Color{R: 1.000, G: 0.000, B: 0.000, A: 1}
	ColorRedOrange                      = Color{R: 1.000, G: 0.325, B: 0.286, A: 1}
	ColorDingyDungeon                   = Color{R: 0.773, G: 0.192, B: 0.318, A: 1}
	ColorMarigoldYellow                 = Color{R: 0.984, G: 0.910, B: 0.439, A: 1}
	ColorPalatinateBlue                 = Color{R: 0.153, G: 0.231, B: 0.886, A: 1}
	ColorMongoose                       = Color{R: 0.710, G: 0.635, B: 0.498, A: 1}
	ColorRichCarmine                    = Color{R: 0.843, G: 0.000, B: 0.251, A: 1}
	ColorMustard                        = Color{R: 1.000, G: 0.859, B: 0.345, A: 1}
	ColorPigmentGreen                   = Color{R: 0.000, G: 0.647, B: 0.314, A: 1}
	ColorRhythm                         = Color{R: 0.467, G: 0.463, B: 0.588, A: 1}
	ColorTexasRose                      = Color{R: 1.000, G: 0.710, B: 0.333, A: 1}
	ColorTide                           = Color{R: 0.749, G: 0.722, B: 0.690, A: 1}
	ColorDonJuan                        = Color{R: 0.365, G: 0.298, B: 0.318, A: 1}
	ColorEnglishLavender                = Color{R: 0.706, G: 0.514, B: 0.584, A: 1}
	ColorMajorelleBlue                  = Color{R: 0.376, G: 0.314, B: 0.863, A: 1}
	ColorPeachPuff                      = Color{R: 1.000, G: 0.855, B: 0.725, A: 1}
	ColorBulgarianRose                  = Color{R: 0.282, G: 0.024, B: 0.027, A: 1}
	ColorDarkViolet                     = Color{R: 0.580, G: 0.000, B: 0.827, A: 1}
	ColorDolly                          = Color{R: 0.976, G: 1.000, B: 0.545, A: 1}
	ColorSunglow                        = Color{R: 1.000, G: 0.800, B: 0.200, A: 1}
	ColorTutu                           = Color{R: 1.000, G: 0.945, B: 0.976, A: 1}
	ColorArtichoke                      = Color{R: 0.561, G: 0.592, B: 0.475, A: 1}
	ColorMilan                          = Color{R: 0.980, G: 1.000, B: 0.643, A: 1}
	ColorPetiteOrchid                   = Color{R: 0.859, G: 0.588, B: 0.565, A: 1}
	ColorBritishRacingGreen             = Color{R: 0.000, G: 0.259, B: 0.145, A: 1}
	ColorDeepSaffron                    = Color{R: 1.000, G: 0.600, B: 0.200, A: 1}
	ColorHotPink                        = Color{R: 1.000, G: 0.412, B: 0.706, A: 1}
	ColorCosmos                         = Color{R: 1.000, G: 0.847, B: 0.851, A: 1}
	ColorRomantic                       = Color{R: 1.000, G: 0.824, B: 0.718, A: 1}
	ColorGoldenYellow                   = Color{R: 1.000, G: 0.875, B: 0.000, A: 1}
	ColorHarlequinGreen                 = Color{R: 0.275, G: 0.796, B: 0.094, A: 1}
	ColorPolar                          = Color{R: 0.898, G: 0.976, B: 0.965, A: 1}
	ColorPotPourri                      = Color{R: 0.961, G: 0.906, B: 0.886, A: 1}
	ColorTigersEye                      = Color{R: 0.878, G: 0.553, B: 0.235, A: 1}
	ColorCorduroy                       = Color{R: 0.376, G: 0.431, B: 0.408, A: 1}
	ColorDarkLavender                   = Color{R: 0.451, G: 0.310, B: 0.588, A: 1}
	ColorGlossyGrape                    = Color{R: 0.671, G: 0.573, B: 0.702, A: 1}
	ColorLipstick                       = Color{R: 0.671, G: 0.020, B: 0.388, A: 1}
	ColorRYBViolet                      = Color{R: 0.525, G: 0.004, B: 0.686, A: 1}
	ColorRebel                          = Color{R: 0.235, G: 0.071, B: 0.024, A: 1}
	ColorGoldenPoppy                    = Color{R: 0.988, G: 0.761, B: 0.000, A: 1}
	ColorGreenHaze                      = Color{R: 0.004, G: 0.639, B: 0.408, A: 1}
	ColorLimedSpruce                    = Color{R: 0.224, G: 0.282, B: 0.318, A: 1}
	ColorHeatWave                       = Color{R: 1.000, G: 0.478, B: 0.000, A: 1}
	ColorPorsche                        = Color{R: 0.918, G: 0.682, B: 0.412, A: 1}
	ColorPortafino                      = Color{R: 1.000, G: 1.000, B: 0.706, A: 1}
	ColorRegalia                        = Color{R: 0.322, G: 0.176, B: 0.502, A: 1}
	ColorWildWillow                     = Color{R: 0.725, G: 0.769, B: 0.416, A: 1}
	ColorBakerMillerPink                = Color{R: 1.000, G: 0.569, B: 0.686, A: 1}
	ColorByzantine                      = Color{R: 0.741, G: 0.200, B: 0.643, A: 1}
	ColorPersianRed                     = Color{R: 0.800, G: 0.200, B: 0.200, A: 1}
	ColorMountbattenPink                = Color{R: 0.600, G: 0.478, B: 0.553, A: 1}
	ColorShadyLady                      = Color{R: 0.667, G: 0.647, B: 0.663, A: 1}
	ColorSurfieGreen                    = Color{R: 0.047, G: 0.478, B: 0.475, A: 1}
	ColorVividTangerine                 = Color{R: 1.000, G: 0.627, B: 0.537, A: 1}
	ColorBallBlue                       = Color{R: 0.129, G: 0.671, B: 0.804, A: 1}
	ColorCopperRust                     = Color{R: 0.580, G: 0.278, B: 0.278, A: 1}
	ColorDoveGray                       = Color{R: 0.427, G: 0.424, B: 0.424, A: 1}
	ColorSiam                           = Color{R: 0.392, G: 0.416, B: 0.329, A: 1}
	ColorWaterspout                     = Color{R: 0.643, G: 0.957, B: 0.976, A: 1}
	ColorBlackPearl                     = Color{R: 0.016, G: 0.075, B: 0.133, A: 1}
	ColorCoolBlack                      = Color{R: 0.000, G: 0.180, B: 0.388, A: 1}
	ColorDarkByzantium                  = Color{R: 0.365, G: 0.224, B: 0.329, A: 1}
	ColorLightGreen                     = Color{R: 0.565, G: 0.933, B: 0.565, A: 1}
	ColorRoastCoffee                    = Color{R: 0.439, G: 0.259, B: 0.255, A: 1}
	ColorDeepCarrotOrange               = Color{R: 0.914, G: 0.412, B: 0.173, A: 1}
	ColorFringyFlower                   = Color{R: 0.694, G: 0.886, B: 0.757, A: 1}
	ColorLava                           = Color{R: 0.812, G: 0.063, B: 0.125, A: 1}
	ColorFlint                          = Color{R: 0.435, G: 0.416, B: 0.380, A: 1}
	ColorMaroon                         = Color{R: 0.502, G: 0.000, B: 0.000, A: 1}
	ColorWaikawaGray                    = Color{R: 0.353, G: 0.431, B: 0.612, A: 1}
	ColorYankeesBlue                    = Color{R: 0.110, G: 0.157, B: 0.255, A: 1}
	ColorBelgion                        = Color{R: 0.678, G: 0.847, B: 1.000, A: 1}
	ColorCalifornia                     = Color{R: 0.996, G: 0.616, B: 0.016, A: 1}
	ColorSchoolBusYellow                = Color{R: 1.000, G: 0.847, B: 0.000, A: 1}
	ColorEnglishHolly                   = Color{R: 0.008, G: 0.176, B: 0.082, A: 1}
	ColorMadras                         = Color{R: 0.247, G: 0.188, B: 0.008, A: 1}
	ColorCarla                          = Color{R: 0.953, G: 1.000, B: 0.847, A: 1}
	ColorCordovan                       = Color{R: 0.537, G: 0.247, B: 0.271, A: 1}
	ColorCrimson                        = Color{R: 0.863, G: 0.078, B: 0.235, A: 1}
	ColorMossGreen                      = Color{R: 0.541, G: 0.604, B: 0.357, A: 1}
	ColorSubmarine                      = Color{R: 0.729, G: 0.780, B: 0.788, A: 1}
	ColorSwirl                          = Color{R: 0.827, G: 0.804, B: 0.773, A: 1}
	ColorTiaMaria                       = Color{R: 0.757, G: 0.267, B: 0.055, A: 1}
	ColorTravertine                     = Color{R: 1.000, G: 0.992, B: 0.910, A: 1}
	ColorBdazzledBlue                   = Color{R: 0.180, G: 0.345, B: 0.580, A: 1}
	ColorButtermilk                     = Color{R: 1.000, G: 0.945, B: 0.710, A: 1}
	ColorElephant                       = Color{R: 0.071, G: 0.204, B: 0.278, A: 1}
	ColorCarolinaBlue                   = Color{R: 0.337, G: 0.627, B: 0.827, A: 1}
	ColorPineTree                       = Color{R: 0.090, G: 0.122, B: 0.016, A: 1}
	ColorTuscanRed                      = Color{R: 0.486, G: 0.282, B: 0.282, A: 1}
	ColorAcadia                         = Color{R: 0.106, G: 0.078, B: 0.016, A: 1}
	ColorAllports                       = Color{R: 0.000, G: 0.463, B: 0.639, A: 1}
	ColorBaliHai                        = Color{R: 0.522, G: 0.624, B: 0.686, A: 1}
	ColorWildRice                       = Color{R: 0.925, G: 0.878, B: 0.565, A: 1}
	ColorAlbescentWhite                 = Color{R: 0.961, G: 0.914, B: 0.827, A: 1}
	ColorAquaDeep                       = Color{R: 0.004, G: 0.294, B: 0.263, A: 1}
	ColorStormcloud                     = Color{R: 0.310, G: 0.400, B: 0.416, A: 1}
	ColorOnyx                           = Color{R: 0.208, G: 0.220, B: 0.224, A: 1}
	ColorRomanSilver                    = Color{R: 0.514, G: 0.537, B: 0.588, A: 1}
	ColorFestival                       = Color{R: 0.984, G: 0.914, B: 0.424, A: 1}
	ColorFlame                          = Color{R: 0.886, G: 0.345, B: 0.133, A: 1}
	ColorMarigold                       = Color{R: 0.918, G: 0.635, B: 0.129, A: 1}
	ColorBlueViolet                     = Color{R: 0.541, G: 0.169, B: 0.886, A: 1}
	ColorOrchidPink                     = Color{R: 0.949, G: 0.741, B: 0.804, A: 1}
	ColorDeepLilac                      = Color{R: 0.600, G: 0.333, B: 0.733, A: 1}
	ColorElectricBlue                   = Color{R: 0.490, G: 0.976, B: 1.000, A: 1}
	ColorMikadoYellow                   = Color{R: 1.000, G: 0.769, B: 0.047, A: 1}
	ColorQueenPink                      = Color{R: 0.910, G: 0.800, B: 0.843, A: 1}
	ColorSandstone                      = Color{R: 0.475, G: 0.427, B: 0.384, A: 1}
	ColorBlazeOrange                    = Color{R: 1.000, G: 0.404, B: 0.000, A: 1}
	ColorCarnation                      = Color{R: 0.976, G: 0.353, B: 0.380, A: 1}
	ColorDaisyBush                      = Color{R: 0.310, G: 0.137, B: 0.596, A: 1}
	ColorUnderagePink                   = Color{R: 0.976, G: 0.902, B: 0.957, A: 1}
	ColorDeepViolet                     = Color{R: 0.200, G: 0.000, B: 0.400, A: 1}
	ColorGenoa                          = Color{R: 0.082, G: 0.451, B: 0.420, A: 1}
	ColorGrizzly                        = Color{R: 0.533, G: 0.345, B: 0.094, A: 1}
	ColorMagenta                        = Color{R: 0.792, G: 0.122, B: 0.482, A: 1}
	ColorReefGold                       = Color{R: 0.624, G: 0.510, B: 0.110, A: 1}
	ColorArcticLime                     = Color{R: 0.816, G: 1.000, B: 0.078, A: 1}
	ColorBrownPod                       = Color{R: 0.251, G: 0.094, B: 0.004, A: 1}
	ColorCola                           = Color{R: 0.247, G: 0.145, B: 0.000, A: 1}
	ColorVividViolet                    = Color{R: 0.624, G: 0.000, B: 1.000, A: 1}
	ColorDeluge                         = Color{R: 0.459, G: 0.388, B: 0.659, A: 1}
	ColorClaret                         = Color{R: 0.498, G: 0.090, B: 0.204, A: 1}
	ColorCoquelicot                     = Color{R: 1.000, G: 0.220, B: 0.000, A: 1}
	ColorDarkBrownTangelo               = Color{R: 0.533, G: 0.396, B: 0.306, A: 1}
	ColorWoodyBrown                     = Color{R: 0.282, G: 0.192, B: 0.192, A: 1}
	ColorArapawa                        = Color{R: 0.067, G: 0.047, B: 0.424, A: 1}
	ColorDarkSeaGreen                   = Color{R: 0.561, G: 0.737, B: 0.561, A: 1}
	ColorViolentViolet                  = Color{R: 0.161, G: 0.047, B: 0.369, A: 1}
	ColorEveningSea                     = Color{R: 0.008, G: 0.306, B: 0.275, A: 1}
	ColorPalePrim                       = Color{R: 0.992, G: 0.996, B: 0.722, A: 1}
	ColorBismark                        = Color{R: 0.286, G: 0.443, B: 0.514, A: 1}
	ColorCaper                          = Color{R: 0.863, G: 0.929, B: 0.706, A: 1}
	ColorCasal                          = Color{R: 0.184, G: 0.380, B: 0.408, A: 1}
	ColorGoldenGlow                     = Color{R: 0.992, G: 0.886, B: 0.584, A: 1}
	ColorMellowApricot                  = Color{R: 0.973, G: 0.722, B: 0.471, A: 1}
	ColorMuddyWaters                    = Color{R: 0.718, G: 0.557, B: 0.361, A: 1}
	ColorSasquatchSocks                 = Color{R: 1.000, G: 0.275, B: 0.506, A: 1}
	ColorSoapstone                      = Color{R: 1.000, G: 0.984, B: 0.976, A: 1}
	ColorCinder                         = Color{R: 0.055, G: 0.055, B: 0.094, A: 1}
	ColorCreole                         = Color{R: 0.118, G: 0.059, B: 0.016, A: 1}
	ColorEcru                           = Color{R: 0.761, G: 0.698, B: 0.502, A: 1}
	ColorDarkMagenta                    = Color{R: 0.545, G: 0.000, B: 0.545, A: 1}
	ColorOceanBlue                      = Color{R: 0.310, G: 0.259, B: 0.710, A: 1}
	ColorPrairieSand                    = Color{R: 0.604, G: 0.220, B: 0.125, A: 1}
	ColorAntiqueRuby                    = Color{R: 0.518, G: 0.106, B: 0.176, A: 1}
	ColorBrilliantRose                  = Color{R: 1.000, G: 0.333, B: 0.639, A: 1}
	ColorCadmiumRed                     = Color{R: 0.890, G: 0.000, B: 0.133, A: 1}
	ColorLightGray                      = Color{R: 0.827, G: 0.827, B: 0.827, A: 1}
	ColorLoulou                         = Color{R: 0.275, G: 0.043, B: 0.255, A: 1}
	ColorMasala                         = Color{R: 0.251, G: 0.231, B: 0.220, A: 1}
	ColorParisM                         = Color{R: 0.149, G: 0.020, B: 0.416, A: 1}
	ColorAuburn                         = Color{R: 0.647, G: 0.165, B: 0.165, A: 1}
	ColorBullShot                       = Color{R: 0.525, G: 0.302, B: 0.118, A: 1}
	ColorHippiePink                     = Color{R: 0.682, G: 0.271, B: 0.376, A: 1}
	ColorMediumElectricBlue             = Color{R: 0.012, G: 0.314, B: 0.588, A: 1}
	ColorPinkLady                       = Color{R: 1.000, G: 0.945, B: 0.847, A: 1}
	ColorProcessMagenta                 = Color{R: 1.000, G: 0.000, B: 0.565, A: 1}
	ColorNepal                          = Color{R: 0.557, G: 0.671, B: 0.757, A: 1}
	ColorPaco                           = Color{R: 0.255, G: 0.122, B: 0.063, A: 1}
	ColorWildSand                       = Color{R: 0.957, G: 0.957, B: 0.957, A: 1}
	ColorAztec                          = Color{R: 0.051, G: 0.110, B: 0.098, A: 1}
	ColorFern                           = Color{R: 0.388, G: 0.718, B: 0.424, A: 1}
	ColorGoldFusion                     = Color{R: 0.522, G: 0.459, B: 0.306, A: 1}
	ColorGoldenDream                    = Color{R: 0.941, G: 0.835, B: 0.176, A: 1}
	ColorPurpleNavy                     = Color{R: 0.306, G: 0.318, B: 0.502, A: 1}
	ColorTropicalRainForest             = Color{R: 0.000, G: 0.459, B: 0.369, A: 1}
	ColorSorrellBrown                   = Color{R: 0.808, G: 0.725, B: 0.561, A: 1}
	ColorGreenSmoke                     = Color{R: 0.643, G: 0.686, B: 0.431, A: 1}
	ColorKhaki                          = Color{R: 0.765, G: 0.690, B: 0.569, A: 1}
	ColorSilverTree                     = Color{R: 0.400, G: 0.710, B: 0.561, A: 1}
	ColorFrenchBlue                     = Color{R: 0.000, G: 0.447, B: 0.733, A: 1}
	ColorZanah                          = Color{R: 0.855, G: 0.925, B: 0.839, A: 1}
	ColorDarkRaspberry                  = Color{R: 0.529, G: 0.149, B: 0.341, A: 1}
	ColorDeepLemon                      = Color{R: 0.961, G: 0.780, B: 0.102, A: 1}
	ColorMinsk                          = Color{R: 0.247, G: 0.188, B: 0.498, A: 1}
	ColorMediumSeaGreen                 = Color{R: 0.235, G: 0.702, B: 0.443, A: 1}
	ColorDeco                           = Color{R: 0.824, G: 0.855, B: 0.592, A: 1}
	ColorHintofGreen                    = Color{R: 0.902, G: 1.000, B: 0.914, A: 1}
	ColorJasper                         = Color{R: 0.843, G: 0.231, B: 0.243, A: 1}
	ColorPacificBlue                    = Color{R: 0.110, G: 0.663, B: 0.788, A: 1}
	ColorEcstasy                        = Color{R: 0.980, G: 0.471, B: 0.078, A: 1}
	ColorHemlock                        = Color{R: 0.369, G: 0.365, B: 0.231, A: 1}
	ColorOrgan                          = Color{R: 0.424, G: 0.180, B: 0.122, A: 1}
	ColorCeil                           = Color{R: 0.573, G: 0.631, B: 0.812, A: 1}
	ColorCrete                          = Color{R: 0.451, G: 0.471, B: 0.161, A: 1}
	ColorHarlequin                      = Color{R: 0.247, G: 1.000, B: 0.000, A: 1}
	ColorMandarin                       = Color{R: 0.953, G: 0.478, B: 0.282, A: 1}
	ColorVividMalachite                 = Color{R: 0.000, G: 0.800, B: 0.200, A: 1}
	ColorAthensGray                     = Color{R: 0.933, G: 0.941, B: 0.953, A: 1}
	ColorBoogerBuster                   = Color{R: 0.867, G: 0.886, B: 0.416, A: 1}
	ColorBuff                           = Color{R: 0.941, G: 0.863, B: 0.510, A: 1}
	ColorPortlandOrange                 = Color{R: 1.000, G: 0.353, B: 0.212, A: 1}
	ColorRusset                         = Color{R: 0.502, G: 0.275, B: 0.106, A: 1}
	ColorSaffron                        = Color{R: 0.957, G: 0.769, B: 0.188, A: 1}
	ColorAmaranthPurple                 = Color{R: 0.671, G: 0.153, B: 0.310, A: 1}
	ColorCello                          = Color{R: 0.118, G: 0.220, B: 0.357, A: 1}
	ColorDimGray                        = Color{R: 0.412, G: 0.412, B: 0.412, A: 1}
	ColorWarmBlack                      = Color{R: 0.000, G: 0.259, B: 0.259, A: 1}
	ColorLightCrimson                   = Color{R: 0.961, G: 0.412, B: 0.569, A: 1}
	ColorMuesli                         = Color{R: 0.667, G: 0.545, B: 0.357, A: 1}
	ColorSmoke                          = Color{R: 0.451, G: 0.510, B: 0.463, A: 1}
	ColorRuddyBrown                     = Color{R: 0.733, G: 0.396, B: 0.157, A: 1}
	ColorAshGrey                        = Color{R: 0.698, G: 0.745, B: 0.710, A: 1}
	ColorMediumJungleGreen              = Color{R: 0.110, G: 0.208, B: 0.176, A: 1}
	ColorPaleSky                        = Color{R: 0.431, G: 0.467, B: 0.514, A: 1}
	ColorCrail                          = Color{R: 0.725, G: 0.318, B: 0.251, A: 1}
	ColorLavenderGray                   = Color{R: 0.769, G: 0.765, B: 0.816, A: 1}
	ColorManatee                        = Color{R: 0.592, G: 0.604, B: 0.667, A: 1}
	ColorPastelGray                     = Color{R: 0.812, G: 0.812, B: 0.769, A: 1}
	ColorPowderBlue                     = Color{R: 0.690, G: 0.878, B: 0.902, A: 1}
	ColorRiceCake                       = Color{R: 1.000, G: 0.996, B: 0.941, A: 1}
	ColorSteelBlue                      = Color{R: 0.275, G: 0.510, B: 0.706, A: 1}
	ColorBlueDiamond                    = Color{R: 0.220, G: 0.016, B: 0.455, A: 1}
	ColorCarouselPink                   = Color{R: 0.976, G: 0.878, B: 0.929, A: 1}
	ColorCottonCandy                    = Color{R: 1.000, G: 0.737, B: 0.851, A: 1}
	ColorCardinal                       = Color{R: 0.769, G: 0.118, B: 0.227, A: 1}
	ColorWatusi                         = Color{R: 1.000, G: 0.867, B: 0.812, A: 1}
	ColorZinnwalditeBrown               = Color{R: 0.173, G: 0.086, B: 0.031, A: 1}
	ColorGolden                         = Color{R: 1.000, G: 0.843, B: 0.000, A: 1}
	ColorPearlLusta                     = Color{R: 0.988, G: 0.957, B: 0.863, A: 1}
	ColorPigmentBlue                    = Color{R: 0.200, G: 0.200, B: 0.600, A: 1}
	ColorLuckyPoint                     = Color{R: 0.102, G: 0.102, B: 0.408, A: 1}
	ColorChiffon                        = Color{R: 0.945, G: 1.000, B: 0.784, A: 1}
	ColorDarkBlueGray                   = Color{R: 0.400, G: 0.400, B: 0.600, A: 1}
	ColorDisco                          = Color{R: 0.529, G: 0.082, B: 0.314, A: 1}
	ColorOlivine                        = Color{R: 0.604, G: 0.725, B: 0.451, A: 1}
	ColorSAEECEAmber                    = Color{R: 1.000, G: 0.494, B: 0.000, A: 1}
	ColorTallow                         = Color{R: 0.659, G: 0.647, B: 0.537, A: 1}
	ColorArsenic                        = Color{R: 0.231, G: 0.267, B: 0.294, A: 1}
	ColorAstra                          = Color{R: 0.980, G: 0.918, B: 0.725, A: 1}
	ColorBlackCoral                     = Color{R: 0.329, G: 0.384, B: 0.435, A: 1}
	ColorVividVermilion                 = Color{R: 0.898, G: 0.376, B: 0.141, A: 1}
	ColorZest                           = Color{R: 0.898, G: 0.518, B: 0.106, A: 1}
	ColorCerisePink                     = Color{R: 0.925, G: 0.231, B: 0.514, A: 1}
	ColorCharmPink                      = Color{R: 0.902, G: 0.561, B: 0.675, A: 1}
	ColorTaupe                          = Color{R: 0.282, G: 0.235, B: 0.196, A: 1}
	ColorBlueYonder                     = Color{R: 0.314, G: 0.447, B: 0.655, A: 1}
	ColorRYBGreen                       = Color{R: 0.400, G: 0.690, B: 0.196, A: 1}
	ColorGordonsGreen                   = Color{R: 0.043, G: 0.067, B: 0.027, A: 1}
	ColorOxfordBlue                     = Color{R: 0.000, G: 0.129, B: 0.278, A: 1}
	ColorTreePoppy                      = Color{R: 0.988, G: 0.612, B: 0.114, A: 1}
	ColorViola                          = Color{R: 0.796, G: 0.561, B: 0.663, A: 1}
	ColorBeige                          = Color{R: 0.961, G: 0.961, B: 0.863, A: 1}
	ColorBirch                          = Color{R: 0.216, G: 0.188, B: 0.129, A: 1}
	ColorCarmineRed                     = Color{R: 1.000, G: 0.000, B: 0.220, A: 1}
	ColorInternationalOrange            = Color{R: 1.000, G: 0.310, B: 0.000, A: 1}
	ColorLumber                         = Color{R: 1.000, G: 0.894, B: 0.804, A: 1}
	ColorNCSBlue                        = Color{R: 0.000, G: 0.529, B: 0.741, A: 1}
	ColorRYBBlue                        = Color{R: 0.008, G: 0.278, B: 0.996, A: 1}
	ColorAppleGreen                     = Color{R: 0.553, G: 0.714, B: 0.000, A: 1}
	ColorBlueHaze                       = Color{R: 0.749, G: 0.745, B: 0.847, A: 1}
	ColorDeepChestnut                   = Color{R: 0.725, G: 0.306, B: 0.282, A: 1}
	ColorPullmanGreen                   = Color{R: 0.231, G: 0.200, B: 0.110, A: 1}
	ColorTangelo                        = Color{R: 0.976, G: 0.302, B: 0.000, A: 1}
	ColorX11DarkGreen                   = Color{R: 0.000, G: 0.392, B: 0.000, A: 1}
	ColorDeepForestGreen                = Color{R: 0.094, G: 0.176, B: 0.035, A: 1}
	ColorIrishCoffee                    = Color{R: 0.373, G: 0.239, B: 0.149, A: 1}
	ColorMunsellYellow                  = Color{R: 0.937, G: 0.800, B: 0.000, A: 1}
	ColorSummerGreen                    = Color{R: 0.588, G: 0.733, B: 0.671, A: 1}
	ColorWhite                          = Color{R: 1.000, G: 1.000, B: 1.000, A: 1}
	ColorCioccolato                     = Color{R: 0.333, G: 0.157, B: 0.047, A: 1}
	ColorHawkesBlue                     = Color{R: 0.831, G: 0.886, B: 0.988, A: 1}
	ColorHeatheredGray                  = Color{R: 0.714, G: 0.690, B: 0.584, A: 1}
	ColorLavenderBlush                  = Color{R: 1.000, G: 0.941, B: 0.961, A: 1}
	ColorMahogany                       = Color{R: 0.753, G: 0.251, B: 0.000, A: 1}
	ColorRipeLemon                      = Color{R: 0.957, G: 0.847, B: 0.110, A: 1}
	ColorRoseEbony                      = Color{R: 0.404, G: 0.282, B: 0.275, A: 1}
	ColorToreaBay                       = Color{R: 0.059, G: 0.176, B: 0.620, A: 1}
	ColorBerylGreen                     = Color{R: 0.871, G: 0.898, B: 0.753, A: 1}
	ColorGinger                         = Color{R: 0.690, G: 0.396, B: 0.000, A: 1}
	ColorGreenLizard                    = Color{R: 0.655, G: 0.957, B: 0.196, A: 1}
	ColorVegasGold                      = Color{R: 0.773, G: 0.702, B: 0.345, A: 1}
	ColorVesuvius                       = Color{R: 0.694, G: 0.290, B: 0.043, A: 1}
	ColorBabyBlueEyes                   = Color{R: 0.631, G: 0.792, B: 0.945, A: 1}
	ColorHopbush                        = Color{R: 0.816, G: 0.427, B: 0.631, A: 1}
	ColorMerino                         = Color{R: 0.965, G: 0.941, B: 0.902, A: 1}
	ColorSandstorm                      = Color{R: 0.925, G: 0.835, B: 0.251, A: 1}
	ColorWebChartreuse                  = Color{R: 0.498, G: 1.000, B: 0.000, A: 1}
	ColorBabyPowder                     = Color{R: 0.996, G: 0.996, B: 0.980, A: 1}
	ColorBombay                         = Color{R: 0.686, G: 0.694, B: 0.722, A: 1}
	ColorCasablanca                     = Color{R: 0.973, G: 0.722, B: 0.325, A: 1}
	ColorRaisinBlack                    = Color{R: 0.141, G: 0.129, B: 0.141, A: 1}
	ColorTuna                           = Color{R: 0.208, G: 0.208, B: 0.259, A: 1}
	ColorWattle                         = Color{R: 0.863, G: 0.843, B: 0.278, A: 1}
	ColorAntiFlashWhite                 = Color{R: 0.949, G: 0.953, B: 0.957, A: 1}
	ColorMonza                          = Color{R: 0.780, G: 0.012, B: 0.118, A: 1}
	ColorMyrtleGreen                    = Color{R: 0.192, G: 0.471, B: 0.451, A: 1}
	ColorCoconut                        = Color{R: 0.588, G: 0.353, B: 0.243, A: 1}
	ColorMadison                        = Color{R: 0.035, G: 0.145, B: 0.365, A: 1}
	ColorMallard                        = Color{R: 0.137, G: 0.204, B: 0.094, A: 1}
	ColorAzalea                         = Color{R: 0.969, G: 0.784, B: 0.855, A: 1}
	ColorDarkYellow                     = Color{R: 0.608, G: 0.529, B: 0.047, A: 1}
	ColorTuscany                        = Color{R: 0.753, G: 0.600, B: 0.600, A: 1}
	ColorDarkOliveGreen                 = Color{R: 0.333, G: 0.420, B: 0.184, A: 1}
	ColorLonestar                       = Color{R: 0.427, G: 0.004, B: 0.004, A: 1}
	ColorSaddle                         = Color{R: 0.298, G: 0.188, B: 0.141, A: 1}
	ColorSaffronMango                   = Color{R: 0.976, G: 0.749, B: 0.345, A: 1}
	ColorChristalle                     = Color{R: 0.200, G: 0.012, B: 0.420, A: 1}
	ColorDeepKoamaru                    = Color{R: 0.200, G: 0.200, B: 0.400, A: 1}
	ColorEarthYellow                    = Color{R: 0.882, G: 0.663, B: 0.373, A: 1}
	ColorMardiGras                      = Color{R: 0.533, G: 0.000, B: 0.522, A: 1}
	ColorAlmond                         = Color{R: 0.937, G: 0.871, B: 0.804, A: 1}
	ColorFieryOrange                    = Color{R: 0.702, G: 0.322, B: 0.075, A: 1}
	ColorGraySuit                       = Color{R: 0.757, G: 0.745, B: 0.804, A: 1}
	ColorViolet                         = Color{R: 0.498, G: 0.000, B: 1.000, A: 1}
	ColorJuneBud                        = Color{R: 0.741, G: 0.855, B: 0.341, A: 1}
	ColorNeonCarrot                     = Color{R: 1.000, G: 0.639, B: 0.263, A: 1}
	ColorPineGreen                      = Color{R: 0.004, G: 0.475, B: 0.435, A: 1}
	ColorYellowGreen                    = Color{R: 0.604, G: 0.804, B: 0.196, A: 1}
	ColorCapeHoney                      = Color{R: 0.996, G: 0.898, B: 0.675, A: 1}
	ColorCavernPink                     = Color{R: 0.890, G: 0.745, B: 0.745, A: 1}
	ColorLavenderMagenta                = Color{R: 0.933, G: 0.510, B: 0.933, A: 1}
	ColorSandDune                       = Color{R: 0.588, G: 0.443, B: 0.090, A: 1}
	ColorWellRead                       = Color{R: 0.706, G: 0.200, B: 0.196, A: 1}
	ColorBrightUbe                      = Color{R: 0.820, G: 0.624, B: 0.910, A: 1}
	ColorCyanAzure                      = Color{R: 0.306, G: 0.510, B: 0.706, A: 1}
	ColorDarkImperialBlue               = Color{R: 0.431, G: 0.431, B: 0.976, A: 1}
	ColorBlackShadows                   = Color{R: 0.749, G: 0.686, B: 0.698, A: 1}
	ColorMorningGlory                   = Color{R: 0.620, G: 0.871, B: 0.878, A: 1}
	ColorSeaweed                        = Color{R: 0.106, G: 0.184, B: 0.067, A: 1}
	ColorSolitaire                      = Color{R: 0.996, G: 0.973, B: 0.886, A: 1}
	ColorGrannySmithApple               = Color{R: 0.659, G: 0.894, B: 0.627, A: 1}
	ColorMikado                         = Color{R: 0.176, G: 0.145, B: 0.063, A: 1}
	ColorPastelGreen                    = Color{R: 0.467, G: 0.867, B: 0.467, A: 1}
	ColorKelp                           = Color{R: 0.271, G: 0.286, B: 0.212, A: 1}
	ColorPineGlade                      = Color{R: 0.780, G: 0.804, B: 0.565, A: 1}
	ColorBondiBlue                      = Color{R: 0.000, G: 0.584, B: 0.714, A: 1}
	ColorCelery                         = Color{R: 0.722, G: 0.761, B: 0.365, A: 1}
	ColorCinderella                     = Color{R: 0.992, G: 0.882, B: 0.863, A: 1}
	ColorOliveGreen                     = Color{R: 0.710, G: 0.702, B: 0.361, A: 1}
	ColorPaleMagenta                    = Color{R: 0.976, G: 0.518, B: 0.898, A: 1}
	ColorPewter                         = Color{R: 0.588, G: 0.659, B: 0.631, A: 1}
	ColorSaddleBrown                    = Color{R: 0.545, G: 0.271, B: 0.075, A: 1}
	ColorSuperPink                      = Color{R: 0.812, G: 0.420, B: 0.663, A: 1}
	ColorCelestialBlue                  = Color{R: 0.286, G: 0.592, B: 0.816, A: 1}
	ColorDenim                          = Color{R: 0.082, G: 0.376, B: 0.741, A: 1}
	ColorEagleGreen                     = Color{R: 0.000, G: 0.286, B: 0.325, A: 1}
	ColorEden                           = Color{R: 0.063, G: 0.345, B: 0.322, A: 1}
	ColorShinyShamrock                  = Color{R: 0.373, G: 0.655, B: 0.471, A: 1}
	ColorGulfBlue                       = Color{R: 0.020, G: 0.086, B: 0.341, A: 1}
	ColorRoseFog                        = Color{R: 0.906, G: 0.737, B: 0.706, A: 1}
	ColorTulip                          = Color{R: 1.000, G: 0.529, B: 0.553, A: 1}
	ColorStonewall                      = Color{R: 0.573, G: 0.522, B: 0.451, A: 1}
	ColorSundown                        = Color{R: 1.000, G: 0.694, B: 0.702, A: 1}
	ColorBermuda                        = Color{R: 0.490, G: 0.847, B: 0.776, A: 1}
	ColorCocoaBrown                     = Color{R: 0.824, G: 0.412, B: 0.118, A: 1}
	ColorLilacLuster                    = Color{R: 0.682, G: 0.596, B: 0.667, A: 1}
	ColorPunga                          = Color{R: 0.302, G: 0.239, B: 0.078, A: 1}
	ColorScienceBlue                    = Color{R: 0.000, G: 0.400, B: 0.800, A: 1}
	ColorWildBlueYonder                 = Color{R: 0.635, G: 0.678, B: 0.816, A: 1}
	ColorBurntUmber                     = Color{R: 0.541, G: 0.200, B: 0.141, A: 1}
	ColorCadmiumGreen                   = Color{R: 0.000, G: 0.420, B: 0.235, A: 1}
	ColorChristine                      = Color{R: 0.906, G: 0.451, B: 0.039, A: 1}
	ColorCoral                          = Color{R: 1.000, G: 0.498, B: 0.314, A: 1}
	ColorIrresistible                   = Color{R: 0.702, G: 0.267, B: 0.424, A: 1}
	ColorSerenade                       = Color{R: 1.000, G: 0.957, B: 0.910, A: 1}
	ColorGondola                        = Color{R: 0.149, G: 0.078, B: 0.078, A: 1}
	ColorLightTaupe                     = Color{R: 0.702, G: 0.545, B: 0.427, A: 1}
	ColorSpicyMix                       = Color{R: 0.545, G: 0.373, B: 0.302, A: 1}
	ColorWillowGrove                    = Color{R: 0.396, G: 0.455, B: 0.365, A: 1}
	ColorMediumBlue                     = Color{R: 0.000, G: 0.000, B: 0.804, A: 1}
	ColorParisWhite                     = Color{R: 0.792, G: 0.863, B: 0.831, A: 1}
	ColorOuterSpace                     = Color{R: 0.255, G: 0.290, B: 0.298, A: 1}
	ColorRockSpray                      = Color{R: 0.729, G: 0.271, B: 0.047, A: 1}
	ColorSoyaBean                       = Color{R: 0.416, G: 0.376, B: 0.318, A: 1}
	ColorStiletto                       = Color{R: 0.612, G: 0.200, B: 0.212, A: 1}
	ColorWindsorTan                     = Color{R: 0.655, G: 0.333, B: 0.008, A: 1}
	ColorCornflowerBlue                 = Color{R: 0.392, G: 0.584, B: 0.929, A: 1}
	ColorCumin                          = Color{R: 0.573, G: 0.263, B: 0.129, A: 1}
	ColorHarvardCrimson                 = Color{R: 0.788, G: 0.000, B: 0.086, A: 1}
	ColorCalico                         = Color{R: 0.878, G: 0.753, B: 0.584, A: 1}
	ColorPhthaloBlue                    = Color{R: 0.000, G: 0.059, B: 0.537, A: 1}
	ColorVioletRed                      = Color{R: 0.969, G: 0.325, B: 0.580, A: 1}
	ColorVividRed                       = Color{R: 0.969, G: 0.051, B: 0.102, A: 1}
	ColorWalnut                         = Color{R: 0.467, G: 0.247, B: 0.102, A: 1}
	ColorWilliam                        = Color{R: 0.227, G: 0.408, B: 0.424, A: 1}
	ColorCadmiumOrange                  = Color{R: 0.929, G: 0.529, B: 0.176, A: 1}
	ColorMayGreen                       = Color{R: 0.298, G: 0.569, B: 0.255, A: 1}
	ColorStarDust                       = Color{R: 0.624, G: 0.624, B: 0.612, A: 1}
	ColorGrullo                         = Color{R: 0.663, G: 0.604, B: 0.525, A: 1}
	ColorSkyMagenta                     = Color{R: 0.812, G: 0.443, B: 0.686, A: 1}
	ColorPablo                          = Color{R: 0.467, G: 0.435, B: 0.380, A: 1}
	ColorParadiso                       = Color{R: 0.192, G: 0.490, B: 0.510, A: 1}
	ColorDiamond                        = Color{R: 0.725, G: 0.949, B: 1.000, A: 1}
	ColorFrenchFuchsia                  = Color{R: 0.992, G: 0.247, B: 0.573, A: 1}
	ColorMulberryWood                   = Color{R: 0.361, G: 0.020, B: 0.212, A: 1}
	ColorSambuca                        = Color{R: 0.227, G: 0.125, B: 0.063, A: 1}
	ColorSmitten                        = Color{R: 0.784, G: 0.255, B: 0.525, A: 1}
	ColorCongressBlue                   = Color{R: 0.008, G: 0.278, B: 0.557, A: 1}
	ColorRuddy                          = Color{R: 1.000, G: 0.000, B: 0.157, A: 1}
	ColorRossoCorsa                     = Color{R: 0.831, G: 0.000, B: 0.000, A: 1}
	ColorScarlett                       = Color{R: 0.584, G: 0.000, B: 0.082, A: 1}
	ColorVidaLoca                       = Color{R: 0.329, G: 0.565, B: 0.098, A: 1}
	ColorBrilliantLavender              = Color{R: 0.957, G: 0.733, B: 1.000, A: 1}
	ColorLightYellow                    = Color{R: 1.000, G: 1.000, B: 0.878, A: 1}
	ColorProvincialPink                 = Color{R: 0.996, G: 0.961, B: 0.945, A: 1}
	ColorIronsideGray                   = Color{R: 0.404, G: 0.400, B: 0.384, A: 1}
	ColorPaleTaupe                      = Color{R: 0.737, G: 0.596, B: 0.494, A: 1}
	ColorVividYellow                    = Color{R: 1.000, G: 0.890, B: 0.008, A: 1}
	ColorBlue                           = Color{R: 0.000, G: 0.000, B: 1.000, A: 1}
	ColorOldHeliotrope                  = Color{R: 0.337, G: 0.235, B: 0.361, A: 1}
	ColorYellowRose                     = Color{R: 1.000, G: 0.941, B: 0.000, A: 1}
	ColorRemy                           = Color{R: 0.996, G: 0.922, B: 0.953, A: 1}
	ColorToledo                         = Color{R: 0.227, G: 0.000, B: 0.125, A: 1}
	ColorVividAmber                     = Color{R: 0.800, G: 0.600, B: 0.000, A: 1}
	ColorIris                           = Color{R: 0.353, G: 0.310, B: 0.812, A: 1}
	ColorSpanishCarmine                 = Color{R: 0.820, G: 0.000, B: 0.278, A: 1}
	ColorSpringLeaves                   = Color{R: 0.341, G: 0.514, B: 0.388, A: 1}
	ColorBottleGreen                    = Color{R: 0.000, G: 0.416, B: 0.306, A: 1}
	ColorClearDay                       = Color{R: 0.914, G: 1.000, B: 0.992, A: 1}
	ColorImperialRed                    = Color{R: 0.929, G: 0.161, B: 0.224, A: 1}
	ColorShiraz                         = Color{R: 0.698, G: 0.035, B: 0.192, A: 1}
	ColorAmericanRose                   = Color{R: 1.000, G: 0.012, B: 0.243, A: 1}
	ColorFOGRA29RichBlack               = Color{R: 0.004, G: 0.043, B: 0.075, A: 1}
	ColorLeather                        = Color{R: 0.588, G: 0.439, B: 0.349, A: 1}
	ColorFerra                          = Color{R: 0.439, G: 0.310, B: 0.314, A: 1}
	ColorPaleSlate                      = Color{R: 0.765, G: 0.749, B: 0.757, A: 1}
	ColorMonsoon                        = Color{R: 0.541, G: 0.514, B: 0.537, A: 1}
	ColorPinkLavender                   = Color{R: 0.847, G: 0.698, B: 0.820, A: 1}
	ColorTrout                          = Color{R: 0.290, G: 0.306, B: 0.353, A: 1}
	ColorDodgerBlue                     = Color{R: 0.118, G: 0.565, B: 1.000, A: 1}
	ColorHarp                           = Color{R: 0.902, G: 0.949, B: 0.918, A: 1}
	ColorHoneydew                       = Color{R: 0.941, G: 1.000, B: 0.941, A: 1}
	ColorTealBlue                       = Color{R: 0.212, G: 0.459, B: 0.533, A: 1}
	ColorCactus                         = Color{R: 0.345, G: 0.443, B: 0.337, A: 1}
	ColorPiper                          = Color{R: 0.788, G: 0.388, B: 0.137, A: 1}
	ColorSunray                         = Color{R: 0.890, G: 0.671, B: 0.341, A: 1}
	ColorCarnationPink                  = Color{R: 1.000, G: 0.651, B: 0.788, A: 1}
	ColorFlax                           = Color{R: 0.933, G: 0.863, B: 0.510, A: 1}
	ColorSirocco                        = Color{R: 0.443, G: 0.502, B: 0.502, A: 1}
	ColorUCLAGold                       = Color{R: 1.000, G: 0.702, B: 0.000, A: 1}
	ColorWisteria                       = Color{R: 0.788, G: 0.627, B: 0.863, A: 1}
	ColorForgetMeNot                    = Color{R: 1.000, G: 0.945, B: 0.933, A: 1}
	ColorSandwisp                       = Color{R: 0.961, G: 0.906, B: 0.635, A: 1}
	ColorShockingPink                   = Color{R: 0.988, G: 0.059, B: 0.753, A: 1}
	ColorOchre                          = Color{R: 0.800, G: 0.467, B: 0.133, A: 1}
	ColorOldRose                        = Color{R: 0.753, G: 0.502, B: 0.506, A: 1}
	ColorPinkFlare                      = Color{R: 0.882, G: 0.753, B: 0.784, A: 1}
	ColorUnbleachedSilk                 = Color{R: 1.000, G: 0.867, B: 0.792, A: 1}
	ColorFlamingoPink                   = Color{R: 0.988, G: 0.557, B: 0.675, A: 1}
	ColorGoldDrop                       = Color{R: 0.945, G: 0.510, B: 0.000, A: 1}
	ColorNobel                          = Color{R: 0.718, G: 0.694, B: 0.694, A: 1}
	ColorChartreuse                     = Color{R: 0.875, G: 1.000, B: 0.000, A: 1}
	ColorGrainBrown                     = Color{R: 0.894, G: 0.835, B: 0.718, A: 1}
	ColorToast                          = Color{R: 0.604, G: 0.431, B: 0.380, A: 1}
	ColorScorpion                       = Color{R: 0.412, G: 0.373, B: 0.384, A: 1}
	ColorSmaltBlue                      = Color{R: 0.318, G: 0.502, B: 0.561, A: 1}
	ColorLavender                       = Color{R: 0.710, G: 0.494, B: 0.863, A: 1}
	ColorRusticRed                      = Color{R: 0.282, G: 0.016, B: 0.016, A: 1}
	ColorLittleBoyBlue                  = Color{R: 0.424, G: 0.627, B: 0.863, A: 1}
	ColorFernFrond                      = Color{R: 0.396, G: 0.447, B: 0.125, A: 1}
	ColorRoseofSharon                   = Color{R: 0.749, G: 0.333, B: 0.000, A: 1}
	ColorJonquil                        = Color{R: 0.957, G: 0.792, B: 0.086, A: 1}
	ColorLavenderPink                   = Color{R: 0.984, G: 0.682, B: 0.824, A: 1}
	ColorLightBrown                     = Color{R: 0.710, G: 0.396, B: 0.114, A: 1}
	ColorMischka                        = Color{R: 0.820, G: 0.824, B: 0.867, A: 1}
	ColorSmoky                          = Color{R: 0.376, G: 0.357, B: 0.451, A: 1}
	ColorBridesmaid                     = Color{R: 0.996, G: 0.941, B: 0.925, A: 1}
	ColorGrayOlive                      = Color{R: 0.663, G: 0.643, B: 0.569, A: 1}
	ColorHotToddy                       = Color{R: 0.702, G: 0.502, B: 0.027, A: 1}
	ColorVulcan                         = Color{R: 0.063, G: 0.071, B: 0.114, A: 1}
	ColorTradewind                      = Color{R: 0.373, G: 0.702, B: 0.675, A: 1}
	ColorTumbleweed                     = Color{R: 0.871, G: 0.667, B: 0.533, A: 1}
	ColorGargoyleGas                    = Color{R: 1.000, G: 0.875, B: 0.275, A: 1}
	ColorRenoSand                       = Color{R: 0.659, G: 0.396, B: 0.082, A: 1}
	ColorSlateGray                      = Color{R: 0.439, G: 0.502, B: 0.565, A: 1}
	ColorHeliotropeMagenta              = Color{R: 0.667, G: 0.000, B: 0.733, A: 1}
	ColorAquaForest                     = Color{R: 0.373, G: 0.655, B: 0.467, A: 1}
	ColorCandyAppleRed                  = Color{R: 1.000, G: 0.031, B: 0.000, A: 1}
	ColorChineseViolet                  = Color{R: 0.522, G: 0.376, B: 0.533, A: 1}
	ColorX11Gray                        = Color{R: 0.745, G: 0.745, B: 0.745, A: 1}
	ColorLimeGreen                      = Color{R: 0.196, G: 0.804, B: 0.196, A: 1}
	ColorSepia                          = Color{R: 0.439, G: 0.259, B: 0.078, A: 1}
	ColorTurtleGreen                    = Color{R: 0.165, G: 0.220, B: 0.043, A: 1}
	ColorLightCornflowerBlue            = Color{R: 0.576, G: 0.800, B: 0.918, A: 1}
	ColorSapphire                       = Color{R: 0.059, G: 0.322, B: 0.729, A: 1}
	ColorGimblet                        = Color{R: 0.722, G: 0.710, B: 0.416, A: 1}
	ColorShipGray                       = Color{R: 0.243, G: 0.227, B: 0.267, A: 1}
	ColorHusk                           = Color{R: 0.718, G: 0.643, B: 0.345, A: 1}
	ColorMariner                        = Color{R: 0.157, G: 0.416, B: 0.804, A: 1}
	ColorNavajoWhite                    = Color{R: 1.000, G: 0.871, B: 0.678, A: 1}
	ColorSanMarino                      = Color{R: 0.271, G: 0.424, B: 0.675, A: 1}
	ColorSunshade                       = Color{R: 1.000, G: 0.620, B: 0.173, A: 1}
	ColorChetwodeBlue                   = Color{R: 0.522, G: 0.506, B: 0.851, A: 1}
	ColorMineralGreen                   = Color{R: 0.247, G: 0.365, B: 0.325, A: 1}
	ColorPaleGoldenrod                  = Color{R: 0.933, G: 0.910, B: 0.667, A: 1}
	ColorGoblin                         = Color{R: 0.239, G: 0.490, B: 0.322, A: 1}
	ColorHotMagenta                     = Color{R: 1.000, G: 0.114, B: 0.808, A: 1}
	ColorManz                           = Color{R: 0.933, G: 0.937, B: 0.471, A: 1}
	ColorSafetyYellow                   = Color{R: 0.933, G: 0.824, B: 0.008, A: 1}
	ColorBitterLemon                    = Color{R: 0.792, G: 0.878, B: 0.051, A: 1}
	ColorCornsilk                       = Color{R: 1.000, G: 0.973, B: 0.863, A: 1}
	ColorCrusta                         = Color{R: 0.992, G: 0.482, B: 0.200, A: 1}
	ColorIndependence                   = Color{R: 0.298, G: 0.318, B: 0.427, A: 1}
	ColorMarzipan                       = Color{R: 0.973, G: 0.859, B: 0.616, A: 1}
	ColorPeanut                         = Color{R: 0.471, G: 0.184, B: 0.086, A: 1}
	ColorZaffre                         = Color{R: 0.000, G: 0.078, B: 0.659, A: 1}
	ColorBisque                         = Color{R: 1.000, G: 0.894, B: 0.769, A: 1}
	ColorBlueRibbon                     = Color{R: 0.000, G: 0.400, B: 1.000, A: 1}
	ColorDelRio                         = Color{R: 0.690, G: 0.604, B: 0.584, A: 1}
	ColorIvory                          = Color{R: 1.000, G: 1.000, B: 0.941, A: 1}
	ColorLightBlue                      = Color{R: 0.678, G: 0.847, B: 0.902, A: 1}
	ColorPickledBluewood                = Color{R: 0.192, G: 0.267, B: 0.349, A: 1}
	ColorBahia                          = Color{R: 0.647, G: 0.796, B: 0.047, A: 1}
	ColorDixie                          = Color{R: 0.886, G: 0.580, B: 0.094, A: 1}
	ColorGeraldine                      = Color{R: 0.984, G: 0.537, B: 0.537, A: 1}
	ColorMandysPink                     = Color{R: 0.949, G: 0.765, B: 0.698, A: 1}
	ColorMirage                         = Color{R: 0.086, G: 0.098, B: 0.157, A: 1}
	ColorPersianPlum                    = Color{R: 0.439, G: 0.110, B: 0.110, A: 1}
	ColorPantoneGreen                   = Color{R: 0.000, G: 0.678, B: 0.263, A: 1}
	ColorPersianPink                    = Color{R: 0.969, G: 0.498, B: 0.745, A: 1}
	ColorClassicRose                    = Color{R: 0.984, G: 0.800, B: 0.906, A: 1}
	ColorMintCream                      = Color{R: 0.961, G: 1.000, B: 0.980, A: 1}
	ColorMyPink                         = Color{R: 0.839, G: 0.569, B: 0.533, A: 1}
	ColorLavenderMist                   = Color{R: 0.902, G: 0.902, B: 0.980, A: 1}
	ColorMoodyBlue                      = Color{R: 0.498, G: 0.463, B: 0.827, A: 1}
	ColorShadowGreen                    = Color{R: 0.604, G: 0.761, B: 0.722, A: 1}
	ColorAcidGreen                      = Color{R: 0.690, G: 0.749, B: 0.102, A: 1}
	ColorChristi                        = Color{R: 0.404, G: 0.655, B: 0.071, A: 1}
	ColorJaffa                          = Color{R: 0.937, G: 0.525, B: 0.247, A: 1}
	ColorHeliotrope                     = Color{R: 0.875, G: 0.451, B: 1.000, A: 1}
	ColorNarvik                         = Color{R: 0.929, G: 0.976, B: 0.945, A: 1}
	ColorScotchMist                     = Color{R: 1.000, G: 0.984, B: 0.863, A: 1}
	ColorStratos                        = Color{R: 0.000, G: 0.027, B: 0.255, A: 1}
	ColorBlueMarguerite                 = Color{R: 0.463, G: 0.400, B: 0.776, A: 1}
	ColorBrownBramble                   = Color{R: 0.349, G: 0.157, B: 0.016, A: 1}
	ColorFulvous                        = Color{R: 0.894, G: 0.518, B: 0.000, A: 1}
	ColorDeepRuby                       = Color{R: 0.518, G: 0.247, B: 0.357, A: 1}
	ColorWillowBrook                    = Color{R: 0.875, G: 0.925, B: 0.855, A: 1}
	ColorPowderAsh                      = Color{R: 0.737, G: 0.788, B: 0.761, A: 1}
	ColorSilverPink                     = Color{R: 0.769, G: 0.682, B: 0.678, A: 1}
	ColorAmaranthPink                   = Color{R: 0.945, G: 0.612, B: 0.733, A: 1}
	ColorEmperor                        = Color{R: 0.318, G: 0.275, B: 0.286, A: 1}
	ColorNutmegWoodFinish               = Color{R: 0.408, G: 0.212, B: 0.000, A: 1}
	ColorHalfSpanishWhite               = Color{R: 0.996, G: 0.957, B: 0.859, A: 1}
	ColorHippieBlue                     = Color{R: 0.345, G: 0.604, B: 0.686, A: 1}
	ColorKidnapper                      = Color{R: 0.882, G: 0.918, B: 0.831, A: 1}
	ColorMidnightMoss                   = Color{R: 0.016, G: 0.063, B: 0.016, A: 1}
	ColorRichLilac                      = Color{R: 0.714, G: 0.400, B: 0.824, A: 1}
	ColorBlossom                        = Color{R: 0.863, G: 0.706, B: 0.737, A: 1}
	ColorBluebonnet                     = Color{R: 0.110, G: 0.110, B: 0.941, A: 1}
	ColorBrown                          = Color{R: 0.588, G: 0.294, B: 0.000, A: 1}
	ColorSurf                           = Color{R: 0.733, G: 0.843, B: 0.757, A: 1}
	ColorSandyBrown                     = Color{R: 0.957, G: 0.643, B: 0.376, A: 1}
	ColorWhitePointer                   = Color{R: 0.996, G: 0.973, B: 1.000, A: 1}
	ColorWine                           = Color{R: 0.447, G: 0.184, B: 0.216, A: 1}
	ColorFrenchLilac                    = Color{R: 0.525, G: 0.376, B: 0.557, A: 1}
	ColorFrenchLime                     = Color{R: 0.620, G: 0.992, B: 0.220, A: 1}
	ColorNebula                         = Color{R: 0.796, G: 0.859, B: 0.839, A: 1}
	ColorToolbox                        = Color{R: 0.455, G: 0.424, B: 0.753, A: 1}
	ColorBossanova                      = Color{R: 0.306, G: 0.165, B: 0.353, A: 1}
	ColorHeliotropeGray                 = Color{R: 0.667, G: 0.596, B: 0.663, A: 1}
	ColorLola                           = Color{R: 0.875, G: 0.812, B: 0.859, A: 1}
	ColorNero                           = Color{R: 0.078, G: 0.024, B: 0.000, A: 1}
	ColorRomance                        = Color{R: 1.000, G: 0.996, B: 0.992, A: 1}
	ColorWasabi                         = Color{R: 0.471, G: 0.541, B: 0.145, A: 1}
	ColorIndianTan                      = Color{R: 0.302, G: 0.118, B: 0.004, A: 1}
	ColorPerfume                        = Color{R: 0.816, G: 0.745, B: 0.973, A: 1}
	ColorShipCove                       = Color{R: 0.471, G: 0.545, B: 0.729, A: 1}
	ColorPomegranate                    = Color{R: 0.953, G: 0.278, B: 0.137, A: 1}
	ColorPumpkinSkin                    = Color{R: 0.694, G: 0.380, B: 0.043, A: 1}
	ColorRadicalRed                     = Color{R: 1.000, G: 0.208, B: 0.369, A: 1}
	ColorRedSalsa                       = Color{R: 0.992, G: 0.227, B: 0.290, A: 1}
	ColorSmokyTopaz                     = Color{R: 0.576, G: 0.239, B: 0.255, A: 1}
	ColorContessa                       = Color{R: 0.776, G: 0.447, B: 0.420, A: 1}
	ColorDesire                         = Color{R: 0.918, G: 0.235, B: 0.325, A: 1}
	ColorLightPastelPurple              = Color{R: 0.694, G: 0.612, B: 0.851, A: 1}
	ColorVividSkyBlue                   = Color{R: 0.000, G: 0.800, B: 1.000, A: 1}
	ColorRiceFlower                     = Color{R: 0.933, G: 1.000, B: 0.886, A: 1}
	ColorThatch                         = Color{R: 0.714, G: 0.616, B: 0.596, A: 1}
	ColorCloud                          = Color{R: 0.780, G: 0.769, B: 0.749, A: 1}
	ColorCloudBurst                     = Color{R: 0.125, G: 0.180, B: 0.329, A: 1}
	ColorJon                            = Color{R: 0.231, G: 0.122, B: 0.122, A: 1}
	ColorDelta                          = Color{R: 0.643, G: 0.643, B: 0.616, A: 1}
	ColorEnglishRed                     = Color{R: 0.671, G: 0.294, B: 0.322, A: 1}
	ColorSinbad                         = Color{R: 0.624, G: 0.843, B: 0.827, A: 1}
	ColorEnergyYellow                   = Color{R: 0.973, G: 0.867, B: 0.361, A: 1}
	ColorFrenchBistre                   = Color{R: 0.522, G: 0.427, B: 0.302, A: 1}
	ColorFrostbite                      = Color{R: 0.914, G: 0.212, B: 0.655, A: 1}
	ColorFruitSalad                     = Color{R: 0.310, G: 0.616, B: 0.365, A: 1}
	ColorGulfStream                     = Color{R: 0.502, G: 0.702, B: 0.682, A: 1}
	ColorAliceBlue                      = Color{R: 0.941, G: 0.973, B: 1.000, A: 1}
	ColorBrandyPunch                    = Color{R: 0.804, G: 0.518, B: 0.161, A: 1}
	ColorCapri                          = Color{R: 0.000, G: 0.749, B: 1.000, A: 1}
	ColorHanPurple                      = Color{R: 0.322, G: 0.094, B: 0.980, A: 1}
	ColorPeridot                        = Color{R: 0.902, G: 0.886, B: 0.000, A: 1}
	ColorDarkScarlet                    = Color{R: 0.337, G: 0.012, B: 0.098, A: 1}
	ColorTrueBlue                       = Color{R: 0.000, G: 0.451, B: 0.812, A: 1}
	ColorWildStrawberry                 = Color{R: 1.000, G: 0.263, B: 0.643, A: 1}
	ColorMimosa                         = Color{R: 0.973, G: 0.992, B: 0.827, A: 1}
	ColorTacao                          = Color{R: 0.929, G: 0.702, B: 0.506, A: 1}
	ColorBole                           = Color{R: 0.475, G: 0.267, B: 0.231, A: 1}
	ColorDew                            = Color{R: 0.918, G: 1.000, B: 0.996, A: 1}
	ColorLoafer                         = Color{R: 0.933, G: 0.957, B: 0.871, A: 1}
	ColorDarkLiver                      = Color{R: 0.325, G: 0.294, B: 0.310, A: 1}
	ColorFuchsia                        = Color{R: 1.000, G: 0.000, B: 1.000, A: 1}
	ColorProcessCyan                    = Color{R: 0.000, G: 0.718, B: 0.922, A: 1}
	ColorFieryRose                      = Color{R: 1.000, G: 0.329, B: 0.439, A: 1}
	ColorHalfBaked                      = Color{R: 0.522, G: 0.769, B: 0.800, A: 1}
	ColorMalta                          = Color{R: 0.741, G: 0.698, B: 0.631, A: 1}
	ColorPorcelain                      = Color{R: 0.937, G: 0.949, B: 0.953, A: 1}
	ColorVerdigris                      = Color{R: 0.263, G: 0.702, B: 0.682, A: 1}
	ColorBronzetone                     = Color{R: 0.302, G: 0.251, B: 0.059, A: 1}
	ColorCamouflageGreen                = Color{R: 0.471, G: 0.525, B: 0.420, A: 1}
	ColorChambray                       = Color{R: 0.208, G: 0.306, B: 0.549, A: 1}
	ColorBlackBean                      = Color{R: 0.239, G: 0.047, B: 0.008, A: 1}
	ColorHintofRed                      = Color{R: 0.984, G: 0.976, B: 0.976, A: 1}
	ColorNavy                           = Color{R: 0.000, G: 0.000, B: 0.502, A: 1}
	ColorObservatory                    = Color{R: 0.008, G: 0.525, B: 0.435, A: 1}
	ColorDogwoodRose                    = Color{R: 0.843, G: 0.094, B: 0.408, A: 1}
	ColorDutchWhite                     = Color{R: 0.937, G: 0.875, B: 0.733, A: 1}
	ColorFrenchPass                     = Color{R: 0.741, G: 0.929, B: 0.992, A: 1}
	ColorDarkPink                       = Color{R: 0.906, G: 0.329, B: 0.502, A: 1}
	ColorHeavyMetal                     = Color{R: 0.169, G: 0.196, B: 0.157, A: 1}
	ColorHitGray                        = Color{R: 0.631, G: 0.678, B: 0.710, A: 1}
	ColorManhattan                      = Color{R: 0.961, G: 0.788, B: 0.600, A: 1}
	ColorVeryPaleOrange                 = Color{R: 1.000, G: 0.875, B: 0.749, A: 1}
	ColorAthsSpecial                    = Color{R: 0.925, G: 0.922, B: 0.808, A: 1}
	ColorCastro                         = Color{R: 0.322, G: 0.000, B: 0.122, A: 1}
	ColorChalky                         = Color{R: 0.933, G: 0.843, B: 0.580, A: 1}
	ColorVividRaspberry                 = Color{R: 1.000, G: 0.000, B: 0.424, A: 1}
	ColorBilbao                         = Color{R: 0.196, G: 0.486, B: 0.078, A: 1}
	ColorDarkPurple                     = Color{R: 0.188, G: 0.098, B: 0.204, A: 1}
	ColorFunBlue                        = Color{R: 0.098, G: 0.349, B: 0.659, A: 1}
	ColorBattleshipGray                 = Color{R: 0.510, G: 0.561, B: 0.447, A: 1}
	ColorSmokeyTopaz                    = Color{R: 0.514, G: 0.165, B: 0.051, A: 1}
	ColorUpsdellRed                     = Color{R: 0.682, G: 0.125, B: 0.161, A: 1}
	ColorCeleste                        = Color{R: 0.698, G: 1.000, B: 1.000, A: 1}
	ColorElm                            = Color{R: 0.110, G: 0.486, B: 0.490, A: 1}
	ColorOlive                          = Color{R: 0.502, G: 0.502, B: 0.000, A: 1}
	ColorRoseWhite                      = Color{R: 1.000, G: 0.965, B: 0.961, A: 1}
	ColorAirForceBlue                   = Color{R: 0.000, G: 0.188, B: 0.561, A: 1}
	ColorBajaWhite                      = Color{R: 1.000, G: 0.973, B: 0.820, A: 1}
	ColorBurnishedBrown                 = Color{R: 0.631, G: 0.478, B: 0.455, A: 1}
	ColorOrinoco                        = Color{R: 0.953, G: 0.984, B: 0.831, A: 1}
	ColorDeepCarmine                    = Color{R: 0.663, G: 0.125, B: 0.243, A: 1}
	ColorFloralWhite                    = Color{R: 1.000, G: 0.980, B: 0.941, A: 1}
	ColorNomad                          = Color{R: 0.729, G: 0.694, B: 0.635, A: 1}
	ColorStrikemaster                   = Color{R: 0.584, G: 0.388, B: 0.529, A: 1}
	ColorPastelBrown                    = Color{R: 0.514, G: 0.412, B: 0.325, A: 1}
	ColorPixiePowder                    = Color{R: 0.224, G: 0.071, B: 0.522, A: 1}
	ColorSpectra                        = Color{R: 0.184, G: 0.353, B: 0.341, A: 1}
	ColorRangoonGreen                   = Color{R: 0.110, G: 0.118, B: 0.075, A: 1}
	ColorToryBlue                       = Color{R: 0.078, G: 0.314, B: 0.667, A: 1}
	ColorWoodsmoke                      = Color{R: 0.047, G: 0.051, B: 0.059, A: 1}
	ColorEggplant                       = Color{R: 0.380, G: 0.251, B: 0.318, A: 1}
	ColorFlavescent                     = Color{R: 0.969, G: 0.914, B: 0.557, A: 1}
	ColorLemonYellow                    = Color{R: 1.000, G: 0.957, B: 0.310, A: 1}
	ColorLemonMeringue                  = Color{R: 0.965, G: 0.918, B: 0.745, A: 1}
	ColorAeroBlue                       = Color{R: 0.788, G: 1.000, B: 0.898, A: 1}
	ColorEminence                       = Color{R: 0.424, G: 0.188, B: 0.510, A: 1}
	ColorGossip                         = Color{R: 0.824, G: 0.973, B: 0.690, A: 1}
	ColorEarlyDawn                      = Color{R: 1.000, G: 0.976, B: 0.902, A: 1}
	ColorSupernova                      = Color{R: 1.000, G: 0.788, B: 0.004, A: 1}
	ColorRiptide                        = Color{R: 0.545, G: 0.902, B: 0.847, A: 1}
	ColorTractorRed                     = Color{R: 0.992, G: 0.055, B: 0.208, A: 1}
	ColorArmyGreen                      = Color{R: 0.294, G: 0.325, B: 0.125, A: 1}
	ColorGrayNickel                     = Color{R: 0.765, G: 0.765, B: 0.741, A: 1}
	ColorMountainMist                   = Color{R: 0.584, G: 0.576, B: 0.588, A: 1}
	ColorRobinEggBlue                   = Color{R: 0.000, G: 0.800, B: 0.800, A: 1}
	ColorRosewood                       = Color{R: 0.396, G: 0.000, B: 0.043, A: 1}
	ColorOperaMauve                     = Color{R: 0.718, G: 0.518, B: 0.655, A: 1}
	ColorWaterLeaf                      = Color{R: 0.631, G: 0.914, B: 0.871, A: 1}
	ColorCream                          = Color{R: 1.000, G: 0.992, B: 0.816, A: 1}
	ColorFirefly                        = Color{R: 0.055, G: 0.165, B: 0.188, A: 1}
	ColorMetallicBronze                 = Color{R: 0.286, G: 0.216, B: 0.106, A: 1}
	ColorMoroccoBrown                   = Color{R: 0.267, G: 0.114, B: 0.000, A: 1}
	ColorNevada                         = Color{R: 0.392, G: 0.431, B: 0.459, A: 1}
	ColorTapestry                       = Color{R: 0.690, G: 0.369, B: 0.506, A: 1}
	ColorVanDykeBrown                   = Color{R: 0.400, G: 0.259, B: 0.157, A: 1}
	ColorBud                            = Color{R: 0.659, G: 0.682, B: 0.612, A: 1}
	ColorLightPink                      = Color{R: 1.000, G: 0.714, B: 0.757, A: 1}
	ColorMediumAquamarine               = Color{R: 0.400, G: 0.867, B: 0.667, A: 1}
	ColorPeachYellow                    = Color{R: 0.980, G: 0.875, B: 0.678, A: 1}
	ColorQuartz                         = Color{R: 0.318, G: 0.282, B: 0.310, A: 1}
	ColorSangria                        = Color{R: 0.573, G: 0.000, B: 0.039, A: 1}
	ColorPopstar                        = Color{R: 0.745, G: 0.310, B: 0.384, A: 1}
	ColorVenus                          = Color{R: 0.573, G: 0.522, B: 0.565, A: 1}
	ColorVividCerulean                  = Color{R: 0.000, G: 0.667, B: 0.933, A: 1}
	ColorBamboo                         = Color{R: 0.855, G: 0.388, B: 0.016, A: 1}
	ColorFOGRA39RichBlack               = Color{R: 0.004, G: 0.008, B: 0.012, A: 1}
	ColorLaurel                         = Color{R: 0.455, G: 0.576, B: 0.471, A: 1}
	ColorMulberry                       = Color{R: 0.773, G: 0.294, B: 0.549, A: 1}
	ColorMuleFawn                       = Color{R: 0.549, G: 0.278, B: 0.184, A: 1}
	ColorRouge                          = Color{R: 0.635, G: 0.231, B: 0.424, A: 1}
	ColorRumSwizzle                     = Color{R: 0.976, G: 0.973, B: 0.894, A: 1}
	ColorTamarillo                      = Color{R: 0.600, G: 0.086, B: 0.075, A: 1}
	ColorGainsboro                      = Color{R: 0.863, G: 0.863, B: 0.863, A: 1}
	ColorJaggedIce                      = Color{R: 0.761, G: 0.910, B: 0.898, A: 1}
	ColorMobster                        = Color{R: 0.498, G: 0.459, B: 0.537, A: 1}
	ColorOrangeYellow                   = Color{R: 0.973, G: 0.835, B: 0.408, A: 1}
	ColorBronzeOlive                    = Color{R: 0.306, G: 0.259, B: 0.047, A: 1}
	ColorEastSide                       = Color{R: 0.675, G: 0.569, B: 0.808, A: 1}
	ColorNutmeg                         = Color{R: 0.506, G: 0.259, B: 0.173, A: 1}
	ColorWineDregs                      = Color{R: 0.404, G: 0.192, B: 0.278, A: 1}
	ColorBitter                         = Color{R: 0.525, G: 0.537, B: 0.455, A: 1}
	ColorEtonBlue                       = Color{R: 0.588, G: 0.784, B: 0.635, A: 1}
	ColorPeriglacialBlue                = Color{R: 0.882, G: 0.902, B: 0.839, A: 1}
	ColorPeachOrange                    = Color{R: 1.000, G: 0.800, B: 0.600, A: 1}
	ColorTurmeric                       = Color{R: 0.792, G: 0.733, B: 0.282, A: 1}
	ColorCorvette                       = Color{R: 0.980, G: 0.827, B: 0.635, A: 1}
	ColorKumera                         = Color{R: 0.533, G: 0.384, B: 0.129, A: 1}
	ColorPaleCyan                       = Color{R: 0.529, G: 0.827, B: 0.973, A: 1}
	ColorConifer                        = Color{R: 0.675, G: 0.867, B: 0.302, A: 1}
	ColorGlaucous                       = Color{R: 0.376, G: 0.510, B: 0.714, A: 1}
	ColorJanna                          = Color{R: 0.957, G: 0.922, B: 0.827, A: 1}
	ColorLightSeaGreen                  = Color{R: 0.125, G: 0.698, B: 0.667, A: 1}
	ColorNaplesYellow                   = Color{R: 0.980, G: 0.855, B: 0.369, A: 1}
	ColorBlueMagentaViolet              = Color{R: 0.333, G: 0.208, B: 0.573, A: 1}
	ColorBlueberry                      = Color{R: 0.310, G: 0.525, B: 0.969, A: 1}
	ColorBrightLavender                 = Color{R: 0.749, G: 0.580, B: 0.894, A: 1}
	ColorSorbus                         = Color{R: 0.992, G: 0.486, B: 0.027, A: 1}
	ColorThistle                        = Color{R: 0.847, G: 0.749, B: 0.847, A: 1}
	ColorTiara                          = Color{R: 0.765, G: 0.820, B: 0.820, A: 1}
	ColorTranquil                       = Color{R: 0.902, G: 1.000, B: 1.000, A: 1}
	ColorPansyPurple                    = Color{R: 0.471, G: 0.094, B: 0.290, A: 1}
	ColorRichBlack                      = Color{R: 0.000, G: 0.251, B: 0.251, A: 1}
	ColorRockBlue                       = Color{R: 0.620, G: 0.694, B: 0.804, A: 1}
	ColorCuttySark                      = Color{R: 0.314, G: 0.463, B: 0.447, A: 1}
	ColorDarkGoldenrod                  = Color{R: 0.722, G: 0.525, B: 0.043, A: 1}
	ColorTwilightLavender               = Color{R: 0.541, G: 0.286, B: 0.420, A: 1}
	ColorAntiqueBrass                   = Color{R: 0.804, G: 0.584, B: 0.459, A: 1}
	ColorAquaIsland                     = Color{R: 0.631, G: 0.855, B: 0.843, A: 1}
	ColorCrayolaBlue                    = Color{R: 0.122, G: 0.459, B: 0.996, A: 1}
	ColorFadedJade                      = Color{R: 0.259, G: 0.475, B: 0.467, A: 1}
	ColorFriarGray                      = Color{R: 0.502, G: 0.494, B: 0.475, A: 1}
	ColorStormGray                      = Color{R: 0.443, G: 0.455, B: 0.525, A: 1}
	ColorVividOrchid                    = Color{R: 0.800, G: 0.000, B: 1.000, A: 1}
	ColorChineseRed                     = Color{R: 0.667, G: 0.220, B: 0.118, A: 1}
	ColorDeepBlue                       = Color{R: 0.133, G: 0.031, B: 0.471, A: 1}
	ColorDiesel                         = Color{R: 0.075, G: 0.000, B: 0.000, A: 1}
	ColorBlueWhale                      = Color{R: 0.016, G: 0.180, B: 0.298, A: 1}
	ColorChileanFire                    = Color{R: 0.969, G: 0.467, B: 0.012, A: 1}
	ColorDarkRed                        = Color{R: 0.545, G: 0.000, B: 0.000, A: 1}
	ColorOrchid                         = Color{R: 0.855, G: 0.439, B: 0.839, A: 1}
	ColorRangitoto                      = Color{R: 0.180, G: 0.196, B: 0.133, A: 1}
	ColorAbsoluteZero                   = Color{R: 0.000, G: 0.282, B: 0.729, A: 1}
	ColorBlackberry                     = Color{R: 0.302, G: 0.004, B: 0.208, A: 1}
	ColorBlueJeans                      = Color{R: 0.365, G: 0.678, B: 0.925, A: 1}
	ColorVarden                         = Color{R: 1.000, G: 0.965, B: 0.875, A: 1}
	ColorBrass                          = Color{R: 0.710, G: 0.651, B: 0.259, A: 1}
	ColorLightApricot                   = Color{R: 0.992, G: 0.835, B: 0.694, A: 1}
	ColorTolopea                        = Color{R: 0.106, G: 0.008, B: 0.271, A: 1}
	ColorFerrariRed                     = Color{R: 1.000, G: 0.157, B: 0.000, A: 1}
	ColorRoyalAirForceBlue              = Color{R: 0.365, G: 0.541, B: 0.659, A: 1}
	ColorBlueChill                      = Color{R: 0.047, G: 0.537, B: 0.565, A: 1}
	ColorBlush                          = Color{R: 0.871, G: 0.365, B: 0.514, A: 1}
	ColorOffGreen                       = Color{R: 0.902, G: 0.973, B: 0.953, A: 1}
	ColorLucky                          = Color{R: 0.686, G: 0.624, B: 0.110, A: 1}
	ColorYuma                           = Color{R: 0.808, G: 0.761, B: 0.569, A: 1}
	ColorCloudy                         = Color{R: 0.675, G: 0.647, B: 0.624, A: 1}
	ColorDeYork                         = Color{R: 0.478, G: 0.769, B: 0.533, A: 1}
	ColorFireBush                       = Color{R: 0.910, G: 0.600, B: 0.157, A: 1}
	ColorAquamarineBlue                 = Color{R: 0.443, G: 0.851, B: 0.886, A: 1}
	ColorTarawera                       = Color{R: 0.027, G: 0.227, B: 0.314, A: 1}
	ColorXanadu                         = Color{R: 0.451, G: 0.525, B: 0.471, A: 1}
	ColorRustyRed                       = Color{R: 0.855, G: 0.173, B: 0.263, A: 1}
	ColorSycamore                       = Color{R: 0.565, G: 0.553, B: 0.224, A: 1}
	ColorUmber                          = Color{R: 0.388, G: 0.318, B: 0.278, A: 1}
	ColorBianca                         = Color{R: 0.988, G: 0.984, B: 0.953, A: 1}
	ColorPastelRed                      = Color{R: 1.000, G: 0.412, B: 0.380, A: 1}
	ColorPearl                          = Color{R: 0.918, G: 0.878, B: 0.784, A: 1}
	ColorBarleyWhite                    = Color{R: 1.000, G: 0.957, B: 0.808, A: 1}
	ColorJungleMist                     = Color{R: 0.706, G: 0.812, B: 0.827, A: 1}
	ColorTapa                           = Color{R: 0.482, G: 0.471, B: 0.455, A: 1}
	ColorPuertoRico                     = Color{R: 0.247, G: 0.757, B: 0.667, A: 1}
	ColorRevolver                       = Color{R: 0.173, G: 0.086, B: 0.196, A: 1}
	ColorRichGold                       = Color{R: 0.659, G: 0.325, B: 0.027, A: 1}
	ColorSkyBlue                        = Color{R: 0.529, G: 0.808, B: 0.922, A: 1}
	ColorBlueSapphire                   = Color{R: 0.071, G: 0.380, B: 0.502, A: 1}
	ColorChromeWhite                    = Color{R: 0.910, G: 0.945, B: 0.831, A: 1}
	ColorHillary                        = Color{R: 0.675, G: 0.647, B: 0.525, A: 1}
	ColorScandal                        = Color{R: 0.812, G: 0.980, B: 0.957, A: 1}
	ColorCamarone                       = Color{R: 0.000, G: 0.345, B: 0.102, A: 1}
	ColorEnvy                           = Color{R: 0.545, G: 0.651, B: 0.565, A: 1}
	ColorMetalPink                      = Color{R: 1.000, G: 0.000, B: 0.992, A: 1}
	ColorEagle                          = Color{R: 0.714, G: 0.729, B: 0.643, A: 1}
	ColorEerieBlack                     = Color{R: 0.106, G: 0.106, B: 0.106, A: 1}
	ColorLightOrchid                    = Color{R: 0.902, G: 0.659, B: 0.843, A: 1}
	ColorStarship                       = Color{R: 0.925, G: 0.949, B: 0.271, A: 1}
	ColorTahitiGold                     = Color{R: 0.914, G: 0.486, B: 0.027, A: 1}
	ColorBubbles                        = Color{R: 0.906, G: 0.996, B: 1.000, A: 1}
	ColorCornHarvest                    = Color{R: 0.545, G: 0.420, B: 0.043, A: 1}
	ColorFrenchGray                     = Color{R: 0.741, G: 0.741, B: 0.776, A: 1}
	ColorDeepCerise                     = Color{R: 0.855, G: 0.196, B: 0.529, A: 1}
	ColorDolphin                        = Color{R: 0.392, G: 0.376, B: 0.467, A: 1}
	ColorSiren                          = Color{R: 0.478, G: 0.004, B: 0.227, A: 1}
	ColorMineShaft                      = Color{R: 0.196, G: 0.196, B: 0.196, A: 1}
	ColorRoseBud                        = Color{R: 0.984, G: 0.698, B: 0.639, A: 1}
	ColorZiggurat                       = Color{R: 0.749, G: 0.859, B: 0.886, A: 1}
	ColorKUCrimson                      = Color{R: 0.910, G: 0.000, B: 0.051, A: 1}
	ColorAureolin                       = Color{R: 0.992, G: 0.933, B: 0.000, A: 1}
	ColorDeer                           = Color{R: 0.729, G: 0.529, B: 0.349, A: 1}
	ColorFlamenco                       = Color{R: 1.000, G: 0.490, B: 0.027, A: 1}
	ColorMamba                          = Color{R: 0.557, G: 0.506, B: 0.565, A: 1}
	ColorSoftPeach                      = Color{R: 0.961, G: 0.929, B: 0.937, A: 1}
	ColorTwilightBlue                   = Color{R: 0.933, G: 0.992, B: 1.000, A: 1}
	ColorVividTangelo                   = Color{R: 0.941, G: 0.455, B: 0.153, A: 1}
	ColorAffair                         = Color{R: 0.443, G: 0.275, B: 0.576, A: 1}
	ColorBrilliantAzure                 = Color{R: 0.200, G: 0.600, B: 1.000, A: 1}
	ColorPharlap                        = Color{R: 0.639, G: 0.502, B: 0.482, A: 1}
	ColorCyanCornflowerBlue             = Color{R: 0.094, G: 0.545, B: 0.761, A: 1}
	ColorPigmentRed                     = Color{R: 0.929, G: 0.110, B: 0.141, A: 1}
	ColorShalimar                       = Color{R: 0.984, G: 1.000, B: 0.729, A: 1}
	ColorTrueV                          = Color{R: 0.541, G: 0.451, B: 0.839, A: 1}
	ColorMexicanPink                    = Color{R: 0.894, G: 0.000, B: 0.486, A: 1}
	ColorMunsellRed                     = Color{R: 0.949, G: 0.000, B: 0.235, A: 1}
	ColorTropicalBlue                   = Color{R: 0.765, G: 0.867, B: 0.976, A: 1}
	ColorMauve                          = Color{R: 0.878, G: 0.690, B: 1.000, A: 1}
	ColorMediumSpringBud                = Color{R: 0.788, G: 0.863, B: 0.529, A: 1}
	ColorPirateGold                     = Color{R: 0.729, G: 0.498, B: 0.012, A: 1}
	ColorRYBOrange                      = Color{R: 0.984, G: 0.600, B: 0.008, A: 1}
	ColorAxolotl                        = Color{R: 0.306, G: 0.400, B: 0.286, A: 1}
	ColorBlackMarlin                    = Color{R: 0.243, G: 0.173, B: 0.110, A: 1}
	ColorBurntOrange                    = Color{R: 0.800, G: 0.333, B: 0.000, A: 1}
	ColorSulu                           = Color{R: 0.757, G: 0.941, B: 0.486, A: 1}
	ColorUPMaroon                       = Color{R: 0.482, G: 0.067, B: 0.075, A: 1}
	ColorKimberly                       = Color{R: 0.451, G: 0.424, B: 0.624, A: 1}
	ColorMediumVermilion                = Color{R: 0.851, G: 0.376, B: 0.231, A: 1}
	ColorStromboli                      = Color{R: 0.196, G: 0.365, B: 0.322, A: 1}
	ColorGoldenBrown                    = Color{R: 0.600, G: 0.396, B: 0.082, A: 1}
	ColorHippieGreen                    = Color{R: 0.325, G: 0.510, B: 0.294, A: 1}
	ColorTurkishRose                    = Color{R: 0.710, G: 0.447, B: 0.506, A: 1}
	ColorCoriander                      = Color{R: 0.769, G: 0.816, B: 0.690, A: 1}
	ColorLightFrenchBeige               = Color{R: 0.784, G: 0.678, B: 0.498, A: 1}
	ColorMeteorite                      = Color{R: 0.235, G: 0.122, B: 0.463, A: 1}
	ColorGiantsOrange                   = Color{R: 0.996, G: 0.353, B: 0.114, A: 1}
	ColorHawaiianTan                    = Color{R: 0.616, G: 0.337, B: 0.086, A: 1}
	ColorLightCyan                      = Color{R: 0.878, G: 1.000, B: 1.000, A: 1}
	ColorPaua                           = Color{R: 0.149, G: 0.012, B: 0.408, A: 1}
	ColorDartmouthGreen                 = Color{R: 0.000, G: 0.439, B: 0.235, A: 1}
	ColorDoublePearlLusta               = Color{R: 0.988, G: 0.957, B: 0.816, A: 1}
	ColorFeta                           = Color{R: 0.941, G: 0.988, B: 0.918, A: 1}
	ColorMinionYellow                   = Color{R: 0.961, G: 0.878, B: 0.314, A: 1}
	ColorOttoman                        = Color{R: 0.914, G: 0.973, B: 0.929, A: 1}
	ColorEucalyptus                     = Color{R: 0.267, G: 0.843, B: 0.659, A: 1}
	ColorHairyHeath                     = Color{R: 0.420, G: 0.165, B: 0.078, A: 1}
	ColorLavenderRose                   = Color{R: 0.984, G: 0.627, B: 0.890, A: 1}
	ColorNCSRed                         = Color{R: 0.769, G: 0.008, B: 0.200, A: 1}
	ColorAntiqueWhite                   = Color{R: 0.980, G: 0.922, B: 0.843, A: 1}
	ColorMalibu                         = Color{R: 0.490, G: 0.784, B: 0.969, A: 1}
	ColorMidGray                        = Color{R: 0.373, G: 0.373, B: 0.431, A: 1}
	ColorPaleCanary                     = Color{R: 1.000, G: 1.000, B: 0.600, A: 1}
	ColorUrobilin                       = Color{R: 0.882, G: 0.678, B: 0.129, A: 1}
	ColorKorma                          = Color{R: 0.561, G: 0.294, B: 0.055, A: 1}
	ColorMocha                          = Color{R: 0.471, G: 0.176, B: 0.098, A: 1}
	ColorOliveHaze                      = Color{R: 0.545, G: 0.518, B: 0.439, A: 1}
	ColorMintGreen                      = Color{R: 0.596, G: 1.000, B: 0.596, A: 1}
	ColorBurntSienna                    = Color{R: 0.914, G: 0.455, B: 0.318, A: 1}
	ColorCrayolaYellow                  = Color{R: 0.988, G: 0.910, B: 0.514, A: 1}
	ColorJewel                          = Color{R: 0.071, G: 0.420, B: 0.251, A: 1}
	ColorMonarch                        = Color{R: 0.545, G: 0.027, B: 0.137, A: 1}
	ColorRope                           = Color{R: 0.557, G: 0.302, B: 0.118, A: 1}
	ColorWhisper                        = Color{R: 0.969, G: 0.961, B: 0.980, A: 1}
	ColorBlueGray                       = Color{R: 0.400, G: 0.600, B: 0.800, A: 1}
	ColorCherryBlossomPink              = Color{R: 1.000, G: 0.718, B: 0.773, A: 1}
	ColorJazzberryJam                   = Color{R: 0.647, G: 0.043, B: 0.369, A: 1}
	ColorPrimrose                       = Color{R: 0.929, G: 0.918, B: 0.600, A: 1}
	ColorButteryWhite                   = Color{R: 1.000, G: 0.988, B: 0.918, A: 1}
	ColorOrangePeel                     = Color{R: 1.000, G: 0.624, B: 0.000, A: 1}
	ColorBeaver                         = Color{R: 0.624, G: 0.506, B: 0.439, A: 1}
	ColorLiver                          = Color{R: 0.404, G: 0.298, B: 0.278, A: 1}
	ColorGurkha                         = Color{R: 0.604, G: 0.584, B: 0.467, A: 1}
	ColorMaverick                       = Color{R: 0.847, G: 0.761, B: 0.835, A: 1}
	ColorBlueLagoon                     = Color{R: 0.675, G: 0.898, B: 0.933, A: 1}
	ColorCowboy                         = Color{R: 0.302, G: 0.157, B: 0.176, A: 1}
	ColorFalcon                         = Color{R: 0.498, G: 0.384, B: 0.427, A: 1}
	ColorKellyGreen                     = Color{R: 0.298, G: 0.733, B: 0.090, A: 1}
	ColorMint                           = Color{R: 0.243, G: 0.706, B: 0.537, A: 1}
	ColorPictonBlue                     = Color{R: 0.271, G: 0.694, B: 0.910, A: 1}
	ColorAfricanViolet                  = Color{R: 0.698, G: 0.518, B: 0.745, A: 1}
	ColorAquamarine                     = Color{R: 0.498, G: 1.000, B: 0.831, A: 1}
	ColorAzureMist                      = Color{R: 0.941, G: 1.000, B: 1.000, A: 1}
	ColorPinkPearl                      = Color{R: 0.906, G: 0.675, B: 0.812, A: 1}
	ColorComo                           = Color{R: 0.318, G: 0.486, B: 0.400, A: 1}
	ColorDandelion                      = Color{R: 0.941, G: 0.882, B: 0.188, A: 1}
	ColorKeyLimePie                     = Color{R: 0.749, G: 0.788, B: 0.129, A: 1}
	ColorTurbo                          = Color{R: 0.980, G: 0.902, B: 0.000, A: 1}
	ColorDirt                           = Color{R: 0.608, G: 0.463, B: 0.325, A: 1}
	ColorPaynesGrey                     = Color{R: 0.325, G: 0.408, B: 0.471, A: 1}
	ColorSherpaBlue                     = Color{R: 0.000, G: 0.286, B: 0.314, A: 1}
	ColorVinRouge                       = Color{R: 0.596, G: 0.239, B: 0.380, A: 1}
	ColorWineBerry                      = Color{R: 0.349, G: 0.114, B: 0.208, A: 1}
	ColorBlueBell                       = Color{R: 0.635, G: 0.635, B: 0.816, A: 1}
	ColorChocolate                      = Color{R: 0.482, G: 0.247, B: 0.000, A: 1}
	ColorMilkPunch                      = Color{R: 1.000, G: 0.965, B: 0.831, A: 1}
	ColorMediumCandyAppleRed            = Color{R: 0.886, G: 0.024, B: 0.173, A: 1}
	ColorMidnightBlue                   = Color{R: 0.098, G: 0.098, B: 0.439, A: 1}
	ColorUCLABlue                       = Color{R: 0.325, G: 0.408, B: 0.584, A: 1}
	ColorBurntMaroon                    = Color{R: 0.259, G: 0.012, B: 0.012, A: 1}
	ColorCyprus                         = Color{R: 0.000, G: 0.243, B: 0.251, A: 1}
	ColorKeyLime                        = Color{R: 0.910, G: 0.957, B: 0.549, A: 1}
	ColorSpanishViridian                = Color{R: 0.000, G: 0.498, B: 0.361, A: 1}
	ColorUltraPink                      = Color{R: 1.000, G: 0.435, B: 1.000, A: 1}
	ColorBigDipOruby                    = Color{R: 0.612, G: 0.145, B: 0.259, A: 1}
	ColorGreenBlue                      = Color{R: 0.067, G: 0.392, B: 0.706, A: 1}
	ColorNileBlue                       = Color{R: 0.098, G: 0.216, B: 0.318, A: 1}
	ColorWheatfield                     = Color{R: 0.953, G: 0.929, B: 0.812, A: 1}
	ColorBlueStone                      = Color{R: 0.004, G: 0.380, B: 0.384, A: 1}
	ColorDarkCerulean                   = Color{R: 0.031, G: 0.271, B: 0.494, A: 1}
	ColorNegroni                        = Color{R: 1.000, G: 0.886, B: 0.773, A: 1}
	ColorPacifika                       = Color{R: 0.467, G: 0.506, B: 0.125, A: 1}
	ColorSantaFe                        = Color{R: 0.694, G: 0.427, B: 0.322, A: 1}
	ColorSchooner                       = Color{R: 0.545, G: 0.518, B: 0.494, A: 1}
	ColorWoodland                       = Color{R: 0.302, G: 0.325, B: 0.157, A: 1}
	ColorButteredRum                    = Color{R: 0.631, G: 0.459, B: 0.051, A: 1}
	ColorCadet                          = Color{R: 0.325, G: 0.408, B: 0.447, A: 1}
	ColorLightGoldenrodYellow           = Color{R: 0.980, G: 0.980, B: 0.824, A: 1}
	ColorWoodrush                       = Color{R: 0.188, G: 0.165, B: 0.059, A: 1}
	ColorGigas                          = Color{R: 0.322, G: 0.235, B: 0.580, A: 1}
	ColorOxley                          = Color{R: 0.467, G: 0.620, B: 0.525, A: 1}
	ColorSpanishPink                    = Color{R: 0.969, G: 0.749, B: 0.745, A: 1}
	ColorTasman                         = Color{R: 0.812, G: 0.863, B: 0.812, A: 1}
	ColorWePeep                         = Color{R: 0.969, G: 0.859, B: 0.902, A: 1}
	ColorCodGray                        = Color{R: 0.043, G: 0.043, B: 0.043, A: 1}
	ColorDarkSienna                     = Color{R: 0.235, G: 0.078, B: 0.078, A: 1}
	ColorSpringBud                      = Color{R: 0.655, G: 0.988, B: 0.000, A: 1}
	ColorLicorice                       = Color{R: 0.102, G: 0.067, B: 0.063, A: 1}
	ColorRosyBrown                      = Color{R: 0.737, G: 0.561, B: 0.561, A: 1}
	ColorRainee                         = Color{R: 0.725, G: 0.784, B: 0.675, A: 1}
	ColorWhiteLilac                     = Color{R: 0.973, G: 0.969, B: 0.988, A: 1}
	ColorCastletonGreen                 = Color{R: 0.000, G: 0.337, B: 0.231, A: 1}
	ColorCopperRed                      = Color{R: 0.796, G: 0.427, B: 0.318, A: 1}
	ColorOutrageousOrange               = Color{R: 1.000, G: 0.431, B: 0.290, A: 1}
	ColorDarkBlue                       = Color{R: 0.000, G: 0.000, B: 0.545, A: 1}
	ColorJapaneseIndigo                 = Color{R: 0.149, G: 0.263, B: 0.282, A: 1}
	ColorPearlAqua                      = Color{R: 0.533, G: 0.847, B: 0.753, A: 1}
	ColorPoloBlue                       = Color{R: 0.553, G: 0.659, B: 0.800, A: 1}
	ColorBunker                         = Color{R: 0.051, G: 0.067, B: 0.090, A: 1}
	ColorCanaryYellow                   = Color{R: 1.000, G: 0.937, B: 0.000, A: 1}
	ColorPastelYellow                   = Color{R: 0.992, G: 0.992, B: 0.588, A: 1}
	ColorLemonChiffon                   = Color{R: 1.000, G: 0.980, B: 0.804, A: 1}
	ColorPaleChestnut                   = Color{R: 0.867, G: 0.678, B: 0.686, A: 1}
	ColorViking                         = Color{R: 0.392, G: 0.800, B: 0.859, A: 1}
	ColorDarkSkyBlue                    = Color{R: 0.549, G: 0.745, B: 0.839, A: 1}
	ColorDowny                          = Color{R: 0.435, G: 0.816, B: 0.773, A: 1}
	ColorStPatricksBlue                 = Color{R: 0.137, G: 0.161, B: 0.478, A: 1}
	ColorRazzmicBerry                   = Color{R: 0.553, G: 0.306, B: 0.522, A: 1}
	ColorRedOxide                       = Color{R: 0.431, G: 0.035, B: 0.008, A: 1}
	ColorVividOrangePeel                = Color{R: 1.000, G: 0.627, B: 0.000, A: 1}
	ColorYukonGold                      = Color{R: 0.482, G: 0.400, B: 0.031, A: 1}
	ColorBouquet                        = Color{R: 0.682, G: 0.502, B: 0.620, A: 1}
	ColorEggWhite                       = Color{R: 1.000, G: 0.937, B: 0.757, A: 1}
	ColorMaximumBlue                    = Color{R: 0.278, G: 0.671, B: 0.800, A: 1}
	ColorCherokee                       = Color{R: 0.988, G: 0.855, B: 0.596, A: 1}
	ColorFuchsiaBlue                    = Color{R: 0.478, G: 0.345, B: 0.757, A: 1}
	ColorMartini                        = Color{R: 0.686, G: 0.627, B: 0.620, A: 1}
	ColorSaharaSand                     = Color{R: 0.945, G: 0.906, B: 0.533, A: 1}
	ColorWillpowerOrange                = Color{R: 0.992, G: 0.345, B: 0.000, A: 1}
	ColorBastille                       = Color{R: 0.161, G: 0.129, B: 0.188, A: 1}
	ColorBlanchedAlmond                 = Color{R: 1.000, G: 0.922, B: 0.804, A: 1}
	ColorImperial                       = Color{R: 0.376, G: 0.184, B: 0.420, A: 1}
	ColorOrange                         = Color{R: 1.000, G: 0.498, B: 0.000, A: 1}
	ColorHintofYellow                   = Color{R: 0.980, G: 0.992, B: 0.894, A: 1}
	ColorKobi                           = Color{R: 0.906, G: 0.624, B: 0.769, A: 1}
	ColorMunsellGreen                   = Color{R: 0.000, G: 0.659, B: 0.467, A: 1}
	ColorYellowMetal                    = Color{R: 0.443, G: 0.388, B: 0.220, A: 1}
	ColorApricot                        = Color{R: 0.984, G: 0.808, B: 0.694, A: 1}
	ColorBuddhaGold                     = Color{R: 0.757, G: 0.627, B: 0.016, A: 1}
	ColorNaturalGray                    = Color{R: 0.545, G: 0.525, B: 0.502, A: 1}
	ColorLightningYellow                = Color{R: 0.988, G: 0.753, B: 0.118, A: 1}
	ColorOil                            = Color{R: 0.157, G: 0.118, B: 0.082, A: 1}
	ColorQuarterPearlLusta              = Color{R: 1.000, G: 0.992, B: 0.957, A: 1}
	ColorIndigoDye                      = Color{R: 0.035, G: 0.122, B: 0.573, A: 1}
	ColorRaven                          = Color{R: 0.447, G: 0.482, B: 0.537, A: 1}
	ColorShadow                         = Color{R: 0.541, G: 0.475, B: 0.365, A: 1}
	ColorTussock                        = Color{R: 0.773, G: 0.600, B: 0.294, A: 1}
	ColorBabyBlue                       = Color{R: 0.537, G: 0.812, B: 0.941, A: 1}
	ColorGeebung                        = Color{R: 0.820, G: 0.561, B: 0.106, A: 1}
	ColorHalfandHalf                    = Color{R: 1.000, G: 0.996, B: 0.882, A: 1}
	ColorRedViolet                      = Color{R: 0.780, G: 0.082, B: 0.522, A: 1}
	ColorWhiteSmoke                     = Color{R: 0.961, G: 0.961, B: 0.961, A: 1}
	ColorApricotWhite                   = Color{R: 1.000, G: 0.996, B: 0.925, A: 1}
	ColorCongoPink                      = Color{R: 0.973, G: 0.514, B: 0.475, A: 1}
	ColorInternationalKleinBlue         = Color{R: 0.000, G: 0.184, B: 0.655, A: 1}
	ColorTeal                           = Color{R: 0.000, G: 0.502, B: 0.502, A: 1}
	ColorBlackOlive                     = Color{R: 0.231, G: 0.235, B: 0.212, A: 1}
	ColorCadetBlue                      = Color{R: 0.373, G: 0.620, B: 0.627, A: 1}
	ColorSurfCrest                      = Color{R: 0.812, G: 0.898, B: 0.824, A: 1}
	ColorEquator                        = Color{R: 0.882, G: 0.737, B: 0.392, A: 1}
	ColorOceanBoatBlue                  = Color{R: 0.000, G: 0.467, B: 0.745, A: 1}
	ColorCameo                          = Color{R: 0.851, G: 0.725, B: 0.608, A: 1}
	ColorDell                           = Color{R: 0.224, G: 0.392, B: 0.075, A: 1}
	ColorDukeBlue                       = Color{R: 0.000, G: 0.000, B: 0.612, A: 1}
	ColorCashmere                       = Color{R: 0.902, G: 0.745, B: 0.647, A: 1}
	ColorJumbo                          = Color{R: 0.486, G: 0.482, B: 0.510, A: 1}
	ColorMerlot                         = Color{R: 0.514, G: 0.098, B: 0.137, A: 1}
	ColorSilver                         = Color{R: 0.753, G: 0.753, B: 0.753, A: 1}
	ColorAcapulco                       = Color{R: 0.486, G: 0.690, B: 0.631, A: 1}
	ColorAvocado                        = Color{R: 0.337, G: 0.510, B: 0.012, A: 1}
	ColorBlackLeatherJacket             = Color{R: 0.145, G: 0.208, B: 0.161, A: 1}
	ColorGableGreen                     = Color{R: 0.086, G: 0.208, B: 0.192, A: 1}
	ColorGrayNurse                      = Color{R: 0.906, G: 0.925, B: 0.902, A: 1}
	ColorMaize                          = Color{R: 0.984, G: 0.925, B: 0.365, A: 1}
	ColorMandy                          = Color{R: 0.886, G: 0.329, B: 0.396, A: 1}
	ColorBlack                          = Color{R: 0.000, G: 0.000, B: 0.000, A: 1}
	ColorChelseaGem                     = Color{R: 0.620, G: 0.325, B: 0.008, A: 1}
	ColorElectricViolet                 = Color{R: 0.545, G: 0.000, B: 1.000, A: 1}
	ColorOpium                          = Color{R: 0.557, G: 0.435, B: 0.439, A: 1}
	ColorTara                           = Color{R: 0.882, G: 0.965, B: 0.910, A: 1}
	ColorTiffanyBlue                    = Color{R: 0.039, G: 0.729, B: 0.710, A: 1}
	ColorZombie                         = Color{R: 0.894, G: 0.839, B: 0.608, A: 1}
	ColorCadillac                       = Color{R: 0.690, G: 0.298, B: 0.416, A: 1}
	ColorDeepJungleGreen                = Color{R: 0.000, G: 0.294, B: 0.286, A: 1}
	ColorMarshland                      = Color{R: 0.043, G: 0.059, B: 0.031, A: 1}
	ColorMoonGlow                       = Color{R: 0.988, G: 0.996, B: 0.855, A: 1}
	ColorCanCan                         = Color{R: 0.835, G: 0.569, B: 0.643, A: 1}
	ColorDeepCarminePink                = Color{R: 0.937, G: 0.188, B: 0.220, A: 1}
	ColorLemon                          = Color{R: 1.000, G: 0.969, B: 0.000, A: 1}
	ColorCitrine                        = Color{R: 0.894, G: 0.816, B: 0.039, A: 1}
	ColorNewOrleans                     = Color{R: 0.953, G: 0.839, B: 0.616, A: 1}
	ColorPinkSwan                       = Color{R: 0.745, G: 0.710, B: 0.718, A: 1}
	ColorTwilight                       = Color{R: 0.894, G: 0.812, B: 0.871, A: 1}
	ColorLapisLazuli                    = Color{R: 0.149, G: 0.380, B: 0.612, A: 1}
	ColorPaprika                        = Color{R: 0.553, G: 0.008, B: 0.149, A: 1}
	ColorPrincetonOrange                = Color{R: 0.961, G: 0.502, B: 0.145, A: 1}
	ColorUARed                          = Color{R: 0.851, G: 0.000, B: 0.298, A: 1}
	ColorCandlelight                    = Color{R: 0.988, G: 0.851, B: 0.090, A: 1}
	ColorDownriver                      = Color{R: 0.035, G: 0.133, B: 0.337, A: 1}
	ColorKokoda                         = Color{R: 0.431, G: 0.427, B: 0.341, A: 1}
	ColorMoonstoneBlue                  = Color{R: 0.451, G: 0.663, B: 0.761, A: 1}
	ColorBlond                          = Color{R: 0.980, G: 0.941, B: 0.745, A: 1}
	ColorCoffee                         = Color{R: 0.435, G: 0.306, B: 0.216, A: 1}
	ColorForestGreen                    = Color{R: 0.133, G: 0.545, B: 0.133, A: 1}
	ColorRaffia                         = Color{R: 0.918, G: 0.855, B: 0.722, A: 1}
	ColorScooter                        = Color{R: 0.180, G: 0.749, B: 0.831, A: 1}
	ColorTusk                           = Color{R: 0.933, G: 0.953, B: 0.765, A: 1}
	ColorBrandyRose                     = Color{R: 0.733, G: 0.537, B: 0.514, A: 1}
	ColorLisbonBrown                    = Color{R: 0.259, G: 0.224, B: 0.129, A: 1}
	ColorMakara                         = Color{R: 0.537, G: 0.490, B: 0.427, A: 1}
	ColorLinen                          = Color{R: 0.980, G: 0.941, B: 0.902, A: 1}
	ColorOldCopper                      = Color{R: 0.447, G: 0.290, B: 0.184, A: 1}
	ColorPaleRobinEggBlue               = Color{R: 0.588, G: 0.871, B: 0.820, A: 1}
	ColorRum                            = Color{R: 0.475, G: 0.412, B: 0.537, A: 1}
	ColorWebOrange                      = Color{R: 1.000, G: 0.647, B: 0.000, A: 1}
	ColorFantasy                        = Color{R: 0.980, G: 0.953, B: 0.941, A: 1}
	ColorGravel                         = Color{R: 0.290, G: 0.267, B: 0.294, A: 1}
	ColorLimedAsh                       = Color{R: 0.455, G: 0.490, B: 0.388, A: 1}
	ColorCrocodile                      = Color{R: 0.451, G: 0.427, B: 0.345, A: 1}
	ColorBayofMany                      = Color{R: 0.153, G: 0.227, B: 0.506, A: 1}
	ColorBistre                         = Color{R: 0.239, G: 0.169, B: 0.122, A: 1}
	ColorBrightSun                      = Color{R: 0.996, G: 0.827, B: 0.235, A: 1}
	ColorDesertStorm                    = Color{R: 0.973, G: 0.973, B: 0.969, A: 1}
	ColorSalomie                        = Color{R: 0.996, G: 0.859, B: 0.553, A: 1}
	ColorChardonnay                     = Color{R: 1.000, G: 0.804, B: 0.549, A: 1}
	ColorJapaneseLaurel                 = Color{R: 0.039, G: 0.412, B: 0.024, A: 1}
	ColorVeryLightAzure                 = Color{R: 0.455, G: 0.733, B: 0.984, A: 1}
	ColorRawUmber                       = Color{R: 0.510, G: 0.400, B: 0.267, A: 1}
	ColorRoyalHeath                     = Color{R: 0.671, G: 0.204, B: 0.447, A: 1}
	ColorElPaso                         = Color{R: 0.118, G: 0.090, B: 0.031, A: 1}
	ColorFedora                         = Color{R: 0.475, G: 0.416, B: 0.471, A: 1}
	ColorFog                            = Color{R: 0.843, G: 0.816, B: 1.000, A: 1}
	ColorDawn                           = Color{R: 0.651, G: 0.635, B: 0.604, A: 1}
	ColorGenericViridian                = Color{R: 0.000, G: 0.498, B: 0.400, A: 1}
	ColorTrinidad                       = Color{R: 0.902, G: 0.306, B: 0.012, A: 1}
	ColorDarkGreen                      = Color{R: 0.004, G: 0.196, B: 0.125, A: 1}
	ColorKaraka                         = Color{R: 0.118, G: 0.086, B: 0.035, A: 1}
	ColorSmalt                          = Color{R: 0.000, G: 0.200, B: 0.600, A: 1}
	ColorSteelGray                      = Color{R: 0.149, G: 0.137, B: 0.208, A: 1}
	ColorSundance                       = Color{R: 0.788, G: 0.702, B: 0.357, A: 1}
	ColorAlmondFrost                    = Color{R: 0.565, G: 0.482, B: 0.443, A: 1}
	ColorCongoBrown                     = Color{R: 0.349, G: 0.216, B: 0.216, A: 1}
	ColorCrimsonRed                     = Color{R: 0.600, G: 0.000, B: 0.000, A: 1}
	ColorQuickSilver                    = Color{R: 0.651, G: 0.651, B: 0.651, A: 1}
	ColorTea                            = Color{R: 0.757, G: 0.729, B: 0.690, A: 1}
	ColorDarkPuce                       = Color{R: 0.310, G: 0.227, B: 0.235, A: 1}
	ColorFrenchViolet                   = Color{R: 0.533, G: 0.024, B: 0.808, A: 1}
	ColorHalayàÚbe                      = Color{R: 0.400, G: 0.220, B: 0.329, A: 1}
	ColorNightRider                     = Color{R: 0.122, G: 0.071, B: 0.059, A: 1}
	ColorPinkRaspberry                  = Color{R: 0.596, G: 0.000, B: 0.212, A: 1}
	ColorRebeccaPurple                  = Color{R: 0.400, G: 0.200, B: 0.600, A: 1}
	ColorSunburntCyclops                = Color{R: 1.000, G: 0.251, B: 0.298, A: 1}
	ColorGoldenGateBridge               = Color{R: 0.753, G: 0.212, B: 0.173, A: 1}
	ColorGovernorBay                    = Color{R: 0.184, G: 0.235, B: 0.702, A: 1}
	ColorLightSalmonPink                = Color{R: 1.000, G: 0.600, B: 0.600, A: 1}
	ColorEbb                            = Color{R: 0.914, G: 0.890, B: 0.890, A: 1}
	ColorCoyoteBrown                    = Color{R: 0.506, G: 0.380, B: 0.243, A: 1}
	ColorLilacBush                      = Color{R: 0.596, G: 0.455, B: 0.827, A: 1}
	ColorPohutukawa                     = Color{R: 0.561, G: 0.008, B: 0.110, A: 1}
	ColorSidecar                        = Color{R: 0.953, G: 0.906, B: 0.733, A: 1}
	ColorZorba                          = Color{R: 0.647, G: 0.608, B: 0.569, A: 1}
	ColorHorizon                        = Color{R: 0.353, G: 0.529, B: 0.627, A: 1}
	ColorPattensBlue                    = Color{R: 0.871, G: 0.961, B: 1.000, A: 1}
	ColorRiverBed                       = Color{R: 0.263, G: 0.298, B: 0.349, A: 1}
	ColorPinkLace                       = Color{R: 1.000, G: 0.867, B: 0.957, A: 1}
	ColorQueenBlue                      = Color{R: 0.263, G: 0.420, B: 0.584, A: 1}
	ColorStudio                         = Color{R: 0.443, G: 0.290, B: 0.698, A: 1}
	ColorLaSalleGreen                   = Color{R: 0.031, G: 0.471, B: 0.188, A: 1}
	ColorPeat                           = Color{R: 0.443, G: 0.420, B: 0.337, A: 1}
	ColorPewterBlue                     = Color{R: 0.545, G: 0.659, B: 0.718, A: 1}
	ColorUtahCrimson                    = Color{R: 0.827, G: 0.000, B: 0.247, A: 1}
	ColorValencia                       = Color{R: 0.847, G: 0.267, B: 0.216, A: 1}
	ColorCosmic                         = Color{R: 0.463, G: 0.224, B: 0.365, A: 1}
	ColorLotus                          = Color{R: 0.525, G: 0.235, B: 0.235, A: 1}
	ColorPizza                          = Color{R: 0.788, G: 0.580, B: 0.082, A: 1}
	ColorBreakerBay                     = Color{R: 0.365, G: 0.631, B: 0.624, A: 1}
	ColorSalmon                         = Color{R: 0.980, G: 0.502, B: 0.447, A: 1}
	ColorWenge                          = Color{R: 0.392, G: 0.329, B: 0.322, A: 1}
	ColorLightGrayishMagenta            = Color{R: 0.800, G: 0.600, B: 0.800, A: 1}
	ColorTeak                           = Color{R: 0.694, G: 0.580, B: 0.380, A: 1}
	ColorWinterWizard                   = Color{R: 0.627, G: 0.902, B: 1.000, A: 1}
	ColorAmethyst                       = Color{R: 0.600, G: 0.400, B: 0.800, A: 1}
	ColorDaintree                       = Color{R: 0.004, G: 0.153, B: 0.192, A: 1}
	ColorEbonyClay                      = Color{R: 0.149, G: 0.157, B: 0.231, A: 1}
	ColorParadisePink                   = Color{R: 0.902, G: 0.243, B: 0.384, A: 1}
	ColorSnowDrift                      = Color{R: 0.969, G: 0.980, B: 0.969, A: 1}
	ColorDarkOrchid                     = Color{R: 0.600, G: 0.196, B: 0.800, A: 1}
	ColorLavenderIndigo                 = Color{R: 0.580, G: 0.341, B: 0.922, A: 1}
	ColorOriolesOrange                  = Color{R: 0.984, G: 0.310, B: 0.078, A: 1}
	ColorCyclamen                       = Color{R: 0.961, G: 0.435, B: 0.631, A: 1}
	ColorTitanWhite                     = Color{R: 0.941, G: 0.933, B: 1.000, A: 1}
	ColorIndochine                      = Color{R: 0.761, G: 0.420, B: 0.012, A: 1}
	ColorSpartanCrimson                 = Color{R: 0.620, G: 0.075, B: 0.086, A: 1}
	ColorUniversityOfTennesseeOrange    = Color{R: 0.969, G: 0.498, B: 0.000, A: 1}
	ColorCatskillWhite                  = Color{R: 0.933, G: 0.965, B: 0.969, A: 1}
	ColorHookersGreen                   = Color{R: 0.286, G: 0.475, B: 0.420, A: 1}
	ColorVeniceBlue                     = Color{R: 0.020, G: 0.349, B: 0.537, A: 1}
	ColorPersianBlue                    = Color{R: 0.110, G: 0.224, B: 0.733, A: 1}
	ColorPersimmon                      = Color{R: 0.925, G: 0.345, B: 0.000, A: 1}
	ColorPompadour                      = Color{R: 0.400, G: 0.000, B: 0.271, A: 1}
	ColorCapeCod                        = Color{R: 0.235, G: 0.267, B: 0.263, A: 1}
	ColorDairyCream                     = Color{R: 0.976, G: 0.894, B: 0.737, A: 1}
	ColorMSUGreen                       = Color{R: 0.094, G: 0.271, B: 0.231, A: 1}
	ColorTosca                          = Color{R: 0.553, G: 0.247, B: 0.247, A: 1}
	ColorBazaar                         = Color{R: 0.596, G: 0.467, B: 0.482, A: 1}
	ColorCrayolaOrange                  = Color{R: 1.000, G: 0.459, B: 0.220, A: 1}
	ColorIsabelline                     = Color{R: 0.957, G: 0.941, B: 0.925, A: 1}
	ColorMetallicSunburst               = Color{R: 0.612, G: 0.486, B: 0.220, A: 1}
	ColorNapa                           = Color{R: 0.675, G: 0.643, B: 0.580, A: 1}
	ColorNickel                         = Color{R: 0.447, G: 0.455, B: 0.447, A: 1}
	ColorPeriwinkleGray                 = Color{R: 0.765, G: 0.804, B: 0.902, A: 1}
	ColorRifleGreen                     = Color{R: 0.267, G: 0.298, B: 0.220, A: 1}
	ColorCoffeeBean                     = Color{R: 0.165, G: 0.078, B: 0.055, A: 1}
	ColorFoggyGray                      = Color{R: 0.796, G: 0.792, B: 0.714, A: 1}
	ColorMediumTurquoise                = Color{R: 0.282, G: 0.820, B: 0.800, A: 1}
	ColorTenne                          = Color{R: 0.804, G: 0.341, B: 0.000, A: 1}
	ColorPeachSchnapps                  = Color{R: 1.000, G: 0.863, B: 0.839, A: 1}
	ColorPiggyPink                      = Color{R: 0.992, G: 0.867, B: 0.902, A: 1}
	ColorSeaBlue                        = Color{R: 0.000, G: 0.412, B: 0.580, A: 1}
	ColorBisonHide                      = Color{R: 0.757, G: 0.718, B: 0.643, A: 1}
	ColorLaPalma                        = Color{R: 0.212, G: 0.529, B: 0.086, A: 1}
	ColorPastelOrange                   = Color{R: 1.000, G: 0.702, B: 0.278, A: 1}
	ColorZircon                         = Color{R: 0.957, G: 0.973, B: 1.000, A: 1}
	ColorDoubleColonialWhite            = Color{R: 0.933, G: 0.890, B: 0.678, A: 1}
	ColorRoseDust                       = Color{R: 0.620, G: 0.369, B: 0.435, A: 1}
	ColorSaltBox                        = Color{R: 0.408, G: 0.369, B: 0.431, A: 1}
	ColorQuarterSpanishWhite            = Color{R: 0.969, G: 0.949, B: 0.882, A: 1}
	ColorRufous                         = Color{R: 0.659, G: 0.110, B: 0.027, A: 1}
	ColorClamShell                      = Color{R: 0.831, G: 0.714, B: 0.686, A: 1}
	ColorDeepFuchsia                    = Color{R: 0.757, G: 0.329, B: 0.757, A: 1}
	ColorDriftwood                      = Color{R: 0.686, G: 0.529, B: 0.318, A: 1}
	ColorCamelot                        = Color{R: 0.537, G: 0.204, B: 0.337, A: 1}
	ColorHorses                         = Color{R: 0.329, G: 0.239, B: 0.216, A: 1}
	ColorNorway                         = Color{R: 0.659, G: 0.741, B: 0.624, A: 1}
	ColorCupid                          = Color{R: 0.984, G: 0.745, B: 0.855, A: 1}
	ColorCyanCobaltBlue                 = Color{R: 0.157, G: 0.345, B: 0.612, A: 1}
	ColorLincolnGreen                   = Color{R: 0.098, G: 0.349, B: 0.020, A: 1}
	ColorPaleCornflowerBlue             = Color{R: 0.671, G: 0.804, B: 0.937, A: 1}
	ColorAlienArmpit                    = Color{R: 0.518, G: 0.871, B: 0.008, A: 1}
	ColorBlueGreen                      = Color{R: 0.051, G: 0.596, B: 0.729, A: 1}
	ColorBurningSand                    = Color{R: 0.851, G: 0.576, B: 0.463, A: 1}
	ColorParisDaisy                     = Color{R: 1.000, G: 0.957, B: 0.431, A: 1}
	ColorRoman                          = Color{R: 0.871, G: 0.388, B: 0.376, A: 1}
	ColorSquirrel                       = Color{R: 0.561, G: 0.506, B: 0.463, A: 1}
	ColorWestSide                       = Color{R: 1.000, G: 0.569, B: 0.059, A: 1}
	ColorCitrineWhite                   = Color{R: 0.980, G: 0.969, B: 0.839, A: 1}
	ColorFairPink                       = Color{R: 1.000, G: 0.937, B: 0.925, A: 1}
	ColorFlushMahogany                  = Color{R: 0.792, G: 0.204, B: 0.208, A: 1}
	ColorPerano                         = Color{R: 0.663, G: 0.745, B: 0.949, A: 1}
	ColorElectricLime                   = Color{R: 0.800, G: 1.000, B: 0.000, A: 1}
	ColorMangoTango                     = Color{R: 1.000, G: 0.510, B: 0.263, A: 1}
	ColorMysticMaroon                   = Color{R: 0.678, G: 0.263, B: 0.475, A: 1}
	ColorBrownSugar                     = Color{R: 0.686, G: 0.431, B: 0.302, A: 1}
	ColorFuscousGray                    = Color{R: 0.329, G: 0.325, B: 0.302, A: 1}
	ColorPaoloVeroneseGreen             = Color{R: 0.000, G: 0.608, B: 0.490, A: 1}
	ColorSizzlingRed                    = Color{R: 1.000, G: 0.220, B: 0.333, A: 1}
	ColorSnowFlurry                     = Color{R: 0.894, G: 1.000, B: 0.820, A: 1}
	ColorThunder                        = Color{R: 0.200, G: 0.161, B: 0.184, A: 1}
	ColorCoralTree                      = Color{R: 0.659, G: 0.420, B: 0.420, A: 1}
	ColorMediumPurple                   = Color{R: 0.576, G: 0.439, B: 0.859, A: 1}
	ColorPrelude                        = Color{R: 0.816, G: 0.753, B: 0.898, A: 1}
	ColorCrusoe                         = Color{R: 0.000, G: 0.282, B: 0.086, A: 1}
	ColorDarkKhaki                      = Color{R: 0.741, G: 0.718, B: 0.420, A: 1}
	ColorDeepSea                        = Color{R: 0.004, G: 0.510, B: 0.420, A: 1}
	ColorPaleRose                       = Color{R: 1.000, G: 0.882, B: 0.949, A: 1}
	ColorAquaSqueeze                    = Color{R: 0.910, G: 0.961, B: 0.949, A: 1}
	ColorBlastOffBronze                 = Color{R: 0.647, G: 0.443, B: 0.392, A: 1}
	ColorButtercup                      = Color{R: 0.953, G: 0.678, B: 0.086, A: 1}
	ColorRoyalPurple                    = Color{R: 0.471, G: 0.318, B: 0.663, A: 1}
	ColorVistaBlue                      = Color{R: 0.486, G: 0.620, B: 0.851, A: 1}
	ColorCherryPie                      = Color{R: 0.165, G: 0.012, B: 0.349, A: 1}
	ColorDarkCandyAppleRed              = Color{R: 0.643, G: 0.000, B: 0.000, A: 1}
	ColorFiord                          = Color{R: 0.251, G: 0.318, B: 0.412, A: 1}
	ColorGorse                          = Color{R: 1.000, G: 0.945, B: 0.310, A: 1}
	ColorMayaBlue                       = Color{R: 0.451, G: 0.761, B: 0.984, A: 1}
	ColorSanguineBrown                  = Color{R: 0.553, G: 0.239, B: 0.220, A: 1}
	ColorPancho                         = Color{R: 0.929, G: 0.804, B: 0.671, A: 1}
	ColorStack                          = Color{R: 0.541, G: 0.561, B: 0.541, A: 1}
	ColorAstronautBlue                  = Color{R: 0.004, G: 0.243, B: 0.384, A: 1}
	ColorKillarney                      = Color{R: 0.227, G: 0.416, B: 0.278, A: 1}
	ColorLaser                          = Color{R: 0.784, G: 0.710, B: 0.408, A: 1}
	ColorDarkSalmon                     = Color{R: 0.914, G: 0.588, B: 0.478, A: 1}
	ColorRedPurple                      = Color{R: 0.894, G: 0.000, B: 0.471, A: 1}
	ColorSwansDown                      = Color{R: 0.863, G: 0.941, B: 0.918, A: 1}
	ColorLogan                          = Color{R: 0.667, G: 0.663, B: 0.804, A: 1}
	ColorPurplePlum                     = Color{R: 0.612, G: 0.318, B: 0.714, A: 1}
	ColorZuccini                        = Color{R: 0.016, G: 0.251, B: 0.133, A: 1}
	ColorCosmicCobalt                   = Color{R: 0.180, G: 0.176, B: 0.533, A: 1}
	ColorGambogeOrange                  = Color{R: 0.600, G: 0.400, B: 0.000, A: 1}
	ColorGumLeaf                        = Color{R: 0.714, G: 0.827, B: 0.749, A: 1}
	ColorRazzleDazzleRose               = Color{R: 1.000, G: 0.200, B: 0.800, A: 1}
	ColorSelago                         = Color{R: 0.941, G: 0.933, B: 0.992, A: 1}
	ColorSpiroDiscoBall                 = Color{R: 0.059, G: 0.753, B: 0.988, A: 1}
	ColorBrightGreen                    = Color{R: 0.400, G: 1.000, B: 0.000, A: 1}
	ColorDarkMidnightBlue               = Color{R: 0.000, G: 0.200, B: 0.400, A: 1}
	ColorJudgeGray                      = Color{R: 0.329, G: 0.263, B: 0.200, A: 1}
	ColorFuego                          = Color{R: 0.745, G: 0.871, B: 0.051, A: 1}
	ColorHampton                        = Color{R: 0.898, G: 0.847, B: 0.686, A: 1}
	ColorPuce                           = Color{R: 0.800, G: 0.533, B: 0.600, A: 1}
	ColorPeriwinkle                     = Color{R: 0.800, G: 0.800, B: 1.000, A: 1}
	ColorQuincy                         = Color{R: 0.384, G: 0.247, B: 0.176, A: 1}
	ColorAubergine                      = Color{R: 0.231, G: 0.035, B: 0.063, A: 1}
	ColorCherrywood                     = Color{R: 0.396, G: 0.102, B: 0.078, A: 1}
	ColorMercury                        = Color{R: 0.898, G: 0.898, B: 0.898, A: 1}
	ColorRollingStone                   = Color{R: 0.455, G: 0.490, B: 0.514, A: 1}
	ColorSpunPearl                      = Color{R: 0.667, G: 0.671, B: 0.718, A: 1}
	ColorUFOGreen                       = Color{R: 0.235, G: 0.816, B: 0.439, A: 1}
	ColorLiberty                        = Color{R: 0.329, G: 0.353, B: 0.655, A: 1}
	ColorMagicPotion                    = Color{R: 1.000, G: 0.267, B: 0.400, A: 1}
	ColorMartinique                     = Color{R: 0.212, G: 0.188, B: 0.314, A: 1}
	ColorPeppermint                     = Color{R: 0.890, G: 0.961, B: 0.882, A: 1}
	ColorRust                           = Color{R: 0.718, G: 0.255, B: 0.055, A: 1}
	ColorAmethystSmoke                  = Color{R: 0.639, G: 0.592, B: 0.706, A: 1}
	ColorBlackWhite                     = Color{R: 1.000, G: 0.996, B: 0.965, A: 1}
	ColorDeepTeal                       = Color{R: 0.000, G: 0.208, B: 0.196, A: 1}
	ColorLimedOak                       = Color{R: 0.675, G: 0.541, B: 0.337, A: 1}
	ColorTwine                          = Color{R: 0.761, G: 0.584, B: 0.365, A: 1}
	ColorRuddyPink                      = Color{R: 0.882, G: 0.557, B: 0.588, A: 1}
	ColorVividCrimson                   = Color{R: 0.800, G: 0.000, B: 0.200, A: 1}
	ColorBlueZodiac                     = Color{R: 0.075, G: 0.149, B: 0.302, A: 1}
	ColorChateauGreen                   = Color{R: 0.251, G: 0.659, B: 0.376, A: 1}
	ColorHavelockBlue                   = Color{R: 0.333, G: 0.565, B: 0.851, A: 1}
	ColorRodeoDust                      = Color{R: 0.788, G: 0.698, B: 0.608, A: 1}
	ColorTanHide                        = Color{R: 0.980, G: 0.616, B: 0.353, A: 1}
	ColorCobaltBlue                     = Color{R: 0.000, G: 0.278, B: 0.671, A: 1}
	ColorGalliano                       = Color{R: 0.863, G: 0.698, B: 0.047, A: 1}
	ColorRaspberry                      = Color{R: 0.890, G: 0.043, B: 0.365, A: 1}
	ColorChinook                        = Color{R: 0.659, G: 0.890, B: 0.741, A: 1}
	ColorClover                         = Color{R: 0.220, G: 0.286, B: 0.063, A: 1}
	ColorJagger                         = Color{R: 0.208, G: 0.055, B: 0.341, A: 1}
	ColorPersianIndigo                  = Color{R: 0.196, G: 0.071, B: 0.478, A: 1}
	ColorUbe                            = Color{R: 0.533, G: 0.471, B: 0.765, A: 1}
	ColorVeryLightMalachiteGreen        = Color{R: 0.392, G: 0.914, B: 0.525, A: 1}
	ColorBarleyCorn                     = Color{R: 0.651, G: 0.545, B: 0.357, A: 1}
	ColorCambridgeBlue                  = Color{R: 0.639, G: 0.757, B: 0.678, A: 1}
	ColorFuchsiaPink                    = Color{R: 1.000, G: 0.467, B: 1.000, A: 1}
	ColorYellowSea                      = Color{R: 0.996, G: 0.663, B: 0.016, A: 1}
	ColorCaribbeanGreen                 = Color{R: 0.000, G: 0.800, B: 0.600, A: 1}
	ColorGiantsClub                     = Color{R: 0.690, G: 0.361, B: 0.322, A: 1}
	ColorRoofTerracotta                 = Color{R: 0.651, G: 0.184, B: 0.125, A: 1}
	ColorPumice                         = Color{R: 0.761, G: 0.792, B: 0.769, A: 1}
	ColorShampoo                        = Color{R: 1.000, G: 0.812, B: 0.945, A: 1}
	ColorSilk                           = Color{R: 0.741, G: 0.694, B: 0.659, A: 1}
	ColorSolidPink                      = Color{R: 0.537, G: 0.220, B: 0.263, A: 1}
	ColorWestCoast                      = Color{R: 0.384, G: 0.318, B: 0.098, A: 1}
	ColorCarminePink                    = Color{R: 0.922, G: 0.298, B: 0.259, A: 1}
	ColorGraniteGreen                   = Color{R: 0.553, G: 0.537, B: 0.455, A: 1}
	ColorLime                           = Color{R: 0.749, G: 1.000, B: 0.000, A: 1}
	ColorSaltpan                        = Color{R: 0.945, G: 0.969, B: 0.949, A: 1}
	ColorTallPoppy                      = Color{R: 0.702, G: 0.176, B: 0.161, A: 1}
	ColorTitaniumYellow                 = Color{R: 0.933, G: 0.902, B: 0.000, A: 1}
	ColorVoodoo                         = Color{R: 0.325, G: 0.204, B: 0.333, A: 1}
	ColorAkaroa                         = Color{R: 0.831, G: 0.769, B: 0.659, A: 1}
	ColorCostaDelSol                    = Color{R: 0.380, G: 0.365, B: 0.188, A: 1}
	ColorOsloGray                       = Color{R: 0.529, G: 0.553, B: 0.569, A: 1}
	ColorBeeswax                        = Color{R: 0.996, G: 0.949, B: 0.780, A: 1}
	ColorPesto                          = Color{R: 0.486, G: 0.463, B: 0.192, A: 1}
	ColorWafer                          = Color{R: 0.871, G: 0.796, B: 0.776, A: 1}
	ColorPaleViolet                     = Color{R: 0.800, G: 0.600, B: 1.000, A: 1}
	ColorPutty                          = Color{R: 0.906, G: 0.804, B: 0.549, A: 1}
	ColorScarpaFlow                     = Color{R: 0.345, G: 0.333, B: 0.384, A: 1}
	ColorBurgundy                       = Color{R: 0.502, G: 0.000, B: 0.125, A: 1}
	ColorGreenPea                       = Color{R: 0.114, G: 0.380, B: 0.259, A: 1}
	ColorMaiTai                         = Color{R: 0.690, G: 0.400, B: 0.031, A: 1}
	ColorCherub                         = Color{R: 0.973, G: 0.851, B: 0.914, A: 1}
	ColorBrightRed                      = Color{R: 0.694, G: 0.000, B: 0.000, A: 1}
	ColorBrownRust                      = Color{R: 0.686, G: 0.349, B: 0.243, A: 1}
	ColorCement                         = Color{R: 0.553, G: 0.463, B: 0.384, A: 1}
	ColorOrientalPink                   = Color{R: 0.776, G: 0.569, B: 0.569, A: 1}
	ColorStraw                          = Color{R: 0.894, G: 0.851, B: 0.435, A: 1}
	ColorBlueBayoux                     = Color{R: 0.286, G: 0.400, B: 0.475, A: 1}
	ColorGreen                          = Color{R: 0.000, G: 1.000, B: 0.000, A: 1}
	ColorNCSGreen                       = Color{R: 0.000, G: 0.624, B: 0.420, A: 1}
	ColorSpindle                        = Color{R: 0.714, G: 0.820, B: 0.918, A: 1}
	ColorCafeRoyale                     = Color{R: 0.435, G: 0.267, B: 0.047, A: 1}
	ColorMondo                          = Color{R: 0.290, G: 0.235, B: 0.188, A: 1}
	ColorSafetyOrange                   = Color{R: 1.000, G: 0.471, B: 0.000, A: 1}
	ColorFandangoPink                   = Color{R: 0.871, G: 0.322, B: 0.522, A: 1}
	ColorGunsmoke                       = Color{R: 0.510, G: 0.525, B: 0.522, A: 1}
	ColorIroko                          = Color{R: 0.263, G: 0.192, B: 0.125, A: 1}
	ColorMelanzane                      = Color{R: 0.188, G: 0.020, B: 0.161, A: 1}
	ColorPunch                          = Color{R: 0.863, G: 0.263, B: 0.200, A: 1}
	ColorCalPolyGreen                   = Color{R: 0.118, G: 0.302, B: 0.169, A: 1}
	ColorCedar                          = Color{R: 0.243, G: 0.110, B: 0.078, A: 1}
	ColorEggshell                       = Color{R: 0.941, G: 0.918, B: 0.839, A: 1}
	ColorTransparent                    = Color{}
	ColorMap                            = map[string]Color{
		"christine":                      ColorChristine,
		"punga":                          ColorPunga,
		"scienceblue":                    ColorScienceBlue,
		"wildblueyonder":                 ColorWildBlueYonder,
		"burntumber":                     ColorBurntUmber,
		"cadmiumgreen":                   ColorCadmiumGreen,
		"serenade":                       ColorSerenade,
		"coral":                          ColorCoral,
		"irresistible":                   ColorIrresistible,
		"gondola":                        ColorGondola,
		"lighttaupe":                     ColorLightTaupe,
		"spicymix":                       ColorSpicyMix,
		"willowgrove":                    ColorWillowGrove,
		"mediumblue":                     ColorMediumBlue,
		"pariswhite":                     ColorParisWhite,
		"harvardcrimson":                 ColorHarvardCrimson,
		"outerspace":                     ColorOuterSpace,
		"rockspray":                      ColorRockSpray,
		"soyabean":                       ColorSoyaBean,
		"stiletto":                       ColorStiletto,
		"windsortan":                     ColorWindsorTan,
		"cornflowerblue":                 ColorCornflowerBlue,
		"cumin":                          ColorCumin,
		"calico":                         ColorCalico,
		"phthaloblue":                    ColorPhthaloBlue,
		"stardust":                       ColorStarDust,
		"violetred":                      ColorVioletRed,
		"vividred":                       ColorVividRed,
		"walnut":                         ColorWalnut,
		"william":                        ColorWilliam,
		"cadmiumorange":                  ColorCadmiumOrange,
		"maygreen":                       ColorMayGreen,
		"grullo":                         ColorGrullo,
		"skymagenta":                     ColorSkyMagenta,
		"mulberrywood":                   ColorMulberryWood,
		"pablo":                          ColorPablo,
		"paradiso":                       ColorParadiso,
		"diamond":                        ColorDiamond,
		"frenchfuchsia":                  ColorFrenchFuchsia,
		"sambuca":                        ColorSambuca,
		"smitten":                        ColorSmitten,
		"congressblue":                   ColorCongressBlue,
		"ruddy":                          ColorRuddy,
		"provincialpink":                 ColorProvincialPink,
		"rossocorsa":                     ColorRossoCorsa,
		"scarlett":                       ColorScarlett,
		"vidaloca":                       ColorVidaLoca,
		"brilliantlavender":              ColorBrilliantLavender,
		"lightyellow":                    ColorLightYellow,
		"vividyellow":                    ColorVividYellow,
		"ironsidegray":                   ColorIronsideGray,
		"paletaupe":                      ColorPaleTaupe,
		"yellowrose":                     ColorYellowRose,
		"blue":                           ColorBlue,
		"oldheliotrope":                  ColorOldHeliotrope,
		"vividamber":                     ColorVividAmber,
		"remy":                           ColorRemy,
		"toledo":                         ColorToledo,
		"imperialred":                    ColorImperialRed,
		"iris":                           ColorIris,
		"spanishcarmine":                 ColorSpanishCarmine,
		"springleaves":                   ColorSpringLeaves,
		"bottlegreen":                    ColorBottleGreen,
		"clearday":                       ColorClearDay,
		"leather":                        ColorLeather,
		"shiraz":                         ColorShiraz,
		"americanrose":                   ColorAmericanRose,
		"fogra29richblack":               ColorFOGRA29RichBlack,
		"ferra":                          ColorFerra,
		"paleslate":                      ColorPaleSlate,
		"honeydew":                       ColorHoneydew,
		"monsoon":                        ColorMonsoon,
		"pinklavender":                   ColorPinkLavender,
		"trout":                          ColorTrout,
		"dodgerblue":                     ColorDodgerBlue,
		"harp":                           ColorHarp,
		"sunray":                         ColorSunray,
		"tealblue":                       ColorTealBlue,
		"cactus":                         ColorCactus,
		"piper":                          ColorPiper,
		"sirocco":                        ColorSirocco,
		"carnationpink":                  ColorCarnationPink,
		"flax":                           ColorFlax,
		"shockingpink":                   ColorShockingPink,
		"uclagold":                       ColorUCLAGold,
		"wisteria":                       ColorWisteria,
		"forgetmenot":                    ColorForgetMeNot,
		"sandwisp":                       ColorSandwisp,
		"nobel":                          ColorNobel,
		"ochre":                          ColorOchre,
		"oldrose":                        ColorOldRose,
		"pinkflare":                      ColorPinkFlare,
		"unbleachedsilk":                 ColorUnbleachedSilk,
		"flamingopink":                   ColorFlamingoPink,
		"golddrop":                       ColorGoldDrop,
		"toast":                          ColorToast,
		"chartreuse":                     ColorChartreuse,
		"grainbrown":                     ColorGrainBrown,
		"scorpion":                       ColorScorpion,
		"smaltblue":                      ColorSmaltBlue,
		"lavender":                       ColorLavender,
		"rusticred":                      ColorRusticRed,
		"littleboyblue":                  ColorLittleBoyBlue,
		"fernfrond":                      ColorFernFrond,
		"roseofsharon":                   ColorRoseofSharon,
		"hottoddy":                       ColorHotToddy,
		"jonquil":                        ColorJonquil,
		"lavenderpink":                   ColorLavenderPink,
		"lightbrown":                     ColorLightBrown,
		"mischka":                        ColorMischka,
		"smoky":                          ColorSmoky,
		"bridesmaid":                     ColorBridesmaid,
		"grayolive":                      ColorGrayOlive,
		"vulcan":                         ColorVulcan,
		"slategray":                      ColorSlateGray,
		"tradewind":                      ColorTradewind,
		"tumbleweed":                     ColorTumbleweed,
		"gargoylegas":                    ColorGargoyleGas,
		"renosand":                       ColorRenoSand,
		"chineseviolet":                  ColorChineseViolet,
		"heliotropemagenta":              ColorHeliotropeMagenta,
		"aquaforest":                     ColorAquaForest,
		"candyapplered":                  ColorCandyAppleRed,
		"turtlegreen":                    ColorTurtleGreen,
		"x11gray":                        ColorX11Gray,
		"limegreen":                      ColorLimeGreen,
		"sepia":                          ColorSepia,
		"lightcornflowerblue":            ColorLightCornflowerBlue,
		"sapphire":                       ColorSapphire,
		"gimblet":                        ColorGimblet,
		"navajowhite":                    ColorNavajoWhite,
		"shipgray":                       ColorShipGray,
		"husk":                           ColorHusk,
		"mariner":                        ColorMariner,
		"palegoldenrod":                  ColorPaleGoldenrod,
		"sanmarino":                      ColorSanMarino,
		"sunshade":                       ColorSunshade,
		"chetwodeblue":                   ColorChetwodeBlue,
		"mineralgreen":                   ColorMineralGreen,
		"crusta":                         ColorCrusta,
		"goblin":                         ColorGoblin,
		"hotmagenta":                     ColorHotMagenta,
		"manz":                           ColorManz,
		"safetyyellow":                   ColorSafetyYellow,
		"bitterlemon":                    ColorBitterLemon,
		"cornsilk":                       ColorCornsilk,
		"delrio":                         ColorDelRio,
		"independence":                   ColorIndependence,
		"marzipan":                       ColorMarzipan,
		"peanut":                         ColorPeanut,
		"zaffre":                         ColorZaffre,
		"bisque":                         ColorBisque,
		"blueribbon":                     ColorBlueRibbon,
		"geraldine":                      ColorGeraldine,
		"ivory":                          ColorIvory,
		"lightblue":                      ColorLightBlue,
		"pickledbluewood":                ColorPickledBluewood,
		"bahia":                          ColorBahia,
		"dixie":                          ColorDixie,
		"persianplum":                    ColorPersianPlum,
		"mandyspink":                     ColorMandysPink,
		"mirage":                         ColorMirage,
		"mypink":                         ColorMyPink,
		"pantonegreen":                   ColorPantoneGreen,
		"persianpink":                    ColorPersianPink,
		"classicrose":                    ColorClassicRose,
		"mintcream":                      ColorMintCream,
		"jaffa":                          ColorJaffa,
		"lavendermist":                   ColorLavenderMist,
		"moodyblue":                      ColorMoodyBlue,
		"shadowgreen":                    ColorShadowGreen,
		"acidgreen":                      ColorAcidGreen,
		"christi":                        ColorChristi,
		"fulvous":                        ColorFulvous,
		"heliotrope":                     ColorHeliotrope,
		"narvik":                         ColorNarvik,
		"scotchmist":                     ColorScotchMist,
		"stratos":                        ColorStratos,
		"bluemarguerite":                 ColorBlueMarguerite,
		"brownbramble":                   ColorBrownBramble,
		"deepruby":                       ColorDeepRuby,
		"willowbrook":                    ColorWillowBrook,
		"nutmegwoodfinish":               ColorNutmegWoodFinish,
		"powderash":                      ColorPowderAsh,
		"silverpink":                     ColorSilverPink,
		"amaranthpink":                   ColorAmaranthPink,
		"emperor":                        ColorEmperor,
		"brown":                          ColorBrown,
		"halfspanishwhite":               ColorHalfSpanishWhite,
		"hippieblue":                     ColorHippieBlue,
		"kidnapper":                      ColorKidnapper,
		"midnightmoss":                   ColorMidnightMoss,
		"richlilac":                      ColorRichLilac,
		"blossom":                        ColorBlossom,
		"bluebonnet":                     ColorBluebonnet,
		"surf":                           ColorSurf,
		"nebula":                         ColorNebula,
		"sandybrown":                     ColorSandyBrown,
		"whitepointer":                   ColorWhitePointer,
		"wine":                           ColorWine,
		"frenchlilac":                    ColorFrenchLilac,
		"frenchlime":                     ColorFrenchLime,
		"lola":                           ColorLola,
		"toolbox":                        ColorToolbox,
		"bossanova":                      ColorBossanova,
		"heliotropegray":                 ColorHeliotropeGray,
		"nero":                           ColorNero,
		"romance":                        ColorRomance,
		"shipcove":                       ColorShipCove,
		"wasabi":                         ColorWasabi,
		"indiantan":                      ColorIndianTan,
		"perfume":                        ColorPerfume,
		"lightpastelpurple":              ColorLightPastelPurple,
		"pomegranate":                    ColorPomegranate,
		"pumpkinskin":                    ColorPumpkinSkin,
		"radicalred":                     ColorRadicalRed,
		"redsalsa":                       ColorRedSalsa,
		"smokytopaz":                     ColorSmokyTopaz,
		"contessa":                       ColorContessa,
		"desire":                         ColorDesire,
		"vividskyblue":                   ColorVividSkyBlue,
		"jon":                            ColorJon,
		"riceflower":                     ColorRiceFlower,
		"thatch":                         ColorThatch,
		"cloud":                          ColorCloud,
		"cloudburst":                     ColorCloudBurst,
		"sinbad":                         ColorSinbad,
		"delta":                          ColorDelta,
		"englishred":                     ColorEnglishRed,
		"capri":                          ColorCapri,
		"energyyellow":                   ColorEnergyYellow,
		"frenchbistre":                   ColorFrenchBistre,
		"frostbite":                      ColorFrostbite,
		"fruitsalad":                     ColorFruitSalad,
		"gulfstream":                     ColorGulfStream,
		"aliceblue":                      ColorAliceBlue,
		"brandypunch":                    ColorBrandyPunch,
		"hanpurple":                      ColorHanPurple,
		"peridot":                        ColorPeridot,
		"wildstrawberry":                 ColorWildStrawberry,
		"darkscarlet":                    ColorDarkScarlet,
		"trueblue":                       ColorTrueBlue,
		"loafer":                         ColorLoafer,
		"mimosa":                         ColorMimosa,
		"tacao":                          ColorTacao,
		"bole":                           ColorBole,
		"dew":                            ColorDew,
		"processcyan":                    ColorProcessCyan,
		"darkliver":                      ColorDarkLiver,
		"fuchsia":                        ColorFuchsia,
		"chambray":                       ColorChambray,
		"fieryrose":                      ColorFieryRose,
		"halfbaked":                      ColorHalfBaked,
		"malta":                          ColorMalta,
		"porcelain":                      ColorPorcelain,
		"verdigris":                      ColorVerdigris,
		"bronzetone":                     ColorBronzetone,
		"camouflagegreen":                ColorCamouflageGreen,
		"blackbean":                      ColorBlackBean,
		"hintofred":                      ColorHintofRed,
		"frenchpass":                     ColorFrenchPass,
		"navy":                           ColorNavy,
		"observatory":                    ColorObservatory,
		"dogwoodrose":                    ColorDogwoodRose,
		"dutchwhite":                     ColorDutchWhite,
		"chalky":                         ColorChalky,
		"darkpink":                       ColorDarkPink,
		"heavymetal":                     ColorHeavyMetal,
		"hitgray":                        ColorHitGray,
		"manhattan":                      ColorManhattan,
		"verypaleorange":                 ColorVeryPaleOrange,
		"athsspecial":                    ColorAthsSpecial,
		"castro":                         ColorCastro,
		"vividraspberry":                 ColorVividRaspberry,
		"funblue":                        ColorFunBlue,
		"bilbao":                         ColorBilbao,
		"darkpurple":                     ColorDarkPurple,
		"upsdellred":                     ColorUpsdellRed,
		"battleshipgray":                 ColorBattleshipGray,
		"smokeytopaz":                    ColorSmokeyTopaz,
		"burnishedbrown":                 ColorBurnishedBrown,
		"celeste":                        ColorCeleste,
		"elm":                            ColorElm,
		"olive":                          ColorOlive,
		"rosewhite":                      ColorRoseWhite,
		"airforceblue":                   ColorAirForceBlue,
		"bajawhite":                      ColorBajaWhite,
		"nomad":                          ColorNomad,
		"orinoco":                        ColorOrinoco,
		"deepcarmine":                    ColorDeepCarmine,
		"floralwhite":                    ColorFloralWhite,
		"spectra":                        ColorSpectra,
		"strikemaster":                   ColorStrikemaster,
		"pastelbrown":                    ColorPastelBrown,
		"pixiepowder":                    ColorPixiePowder,
		"lemonyellow":                    ColorLemonYellow,
		"rangoongreen":                   ColorRangoonGreen,
		"toryblue":                       ColorToryBlue,
		"woodsmoke":                      ColorWoodsmoke,
		"eggplant":                       ColorEggplant,
		"flavescent":                     ColorFlavescent,
		"gossip":                         ColorGossip,
		"lemonmeringue":                  ColorLemonMeringue,
		"aeroblue":                       ColorAeroBlue,
		"eminence":                       ColorEminence,
		"earlydawn":                      ColorEarlyDawn,
		"supernova":                      ColorSupernova,
		"mountainmist":                   ColorMountainMist,
		"riptide":                        ColorRiptide,
		"tractorred":                     ColorTractorRed,
		"armygreen":                      ColorArmyGreen,
		"graynickel":                     ColorGrayNickel,
		"robineggblue":                   ColorRobinEggBlue,
		"rosewood":                       ColorRosewood,
		"metallicbronze":                 ColorMetallicBronze,
		"operamauve":                     ColorOperaMauve,
		"waterleaf":                      ColorWaterLeaf,
		"cream":                          ColorCream,
		"firefly":                        ColorFirefly,
		"mediumaquamarine":               ColorMediumAquamarine,
		"moroccobrown":                   ColorMoroccoBrown,
		"nevada":                         ColorNevada,
		"tapestry":                       ColorTapestry,
		"vandykebrown":                   ColorVanDykeBrown,
		"bud":                            ColorBud,
		"lightpink":                      ColorLightPink,
		"sangria":                        ColorSangria,
		"peachyellow":                    ColorPeachYellow,
		"quartz":                         ColorQuartz,
		"laurel":                         ColorLaurel,
		"popstar":                        ColorPopstar,
		"venus":                          ColorVenus,
		"vividcerulean":                  ColorVividCerulean,
		"bamboo":                         ColorBamboo,
		"fogra39richblack":               ColorFOGRA39RichBlack,
		"mobster":                        ColorMobster,
		"mulberry":                       ColorMulberry,
		"mulefawn":                       ColorMuleFawn,
		"rouge":                          ColorRouge,
		"rumswizzle":                     ColorRumSwizzle,
		"tamarillo":                      ColorTamarillo,
		"gainsboro":                      ColorGainsboro,
		"jaggedice":                      ColorJaggedIce,
		"nutmeg":                         ColorNutmeg,
		"orangeyellow":                   ColorOrangeYellow,
		"bronzeolive":                    ColorBronzeOlive,
		"eastside":                       ColorEastSide,
		"winedregs":                      ColorWineDregs,
		"bitter":                         ColorBitter,
		"etonblue":                       ColorEtonBlue,
		"periglacialblue":                ColorPeriglacialBlue,
		"palecyan":                       ColorPaleCyan,
		"peachorange":                    ColorPeachOrange,
		"turmeric":                       ColorTurmeric,
		"corvette":                       ColorCorvette,
		"kumera":                         ColorKumera,
		"brightlavender":                 ColorBrightLavender,
		"conifer":                        ColorConifer,
		"glaucous":                       ColorGlaucous,
		"janna":                          ColorJanna,
		"lightseagreen":                  ColorLightSeaGreen,
		"naplesyellow":                   ColorNaplesYellow,
		"bluemagentaviolet":              ColorBlueMagentaViolet,
		"blueberry":                      ColorBlueberry,
		"rockblue":                       ColorRockBlue,
		"sorbus":                         ColorSorbus,
		"thistle":                        ColorThistle,
		"tiara":                          ColorTiara,
		"tranquil":                       ColorTranquil,
		"pansypurple":                    ColorPansyPurple,
		"richblack":                      ColorRichBlack,
		"crayolablue":                    ColorCrayolaBlue,
		"cuttysark":                      ColorCuttySark,
		"darkgoldenrod":                  ColorDarkGoldenrod,
		"twilightlavender":               ColorTwilightLavender,
		"antiquebrass":                   ColorAntiqueBrass,
		"aquaisland":                     ColorAquaIsland,
		"diesel":                         ColorDiesel,
		"fadedjade":                      ColorFadedJade,
		"friargray":                      ColorFriarGray,
		"stormgray":                      ColorStormGray,
		"vividorchid":                    ColorVividOrchid,
		"chinesered":                     ColorChineseRed,
		"deepblue":                       ColorDeepBlue,
		"bluejeans":                      ColorBlueJeans,
		"bluewhale":                      ColorBlueWhale,
		"chileanfire":                    ColorChileanFire,
		"darkred":                        ColorDarkRed,
		"orchid":                         ColorOrchid,
		"rangitoto":                      ColorRangitoto,
		"absolutezero":                   ColorAbsoluteZero,
		"blackberry":                     ColorBlackberry,
		"tolopea":                        ColorTolopea,
		"varden":                         ColorVarden,
		"brass":                          ColorBrass,
		"lightapricot":                   ColorLightApricot,
		"ferrarired":                     ColorFerrariRed,
		"offgreen":                       ColorOffGreen,
		"royalairforceblue":              ColorRoyalAirForceBlue,
		"bluechill":                      ColorBlueChill,
		"blush":                          ColorBlush,
		"firebush":                       ColorFireBush,
		"lucky":                          ColorLucky,
		"yuma":                           ColorYuma,
		"cloudy":                         ColorCloudy,
		"deyork":                         ColorDeYork,
		"xanadu":                         ColorXanadu,
		"aquamarineblue":                 ColorAquamarineBlue,
		"tarawera":                       ColorTarawera,
		"pearl":                          ColorPearl,
		"rustyred":                       ColorRustyRed,
		"sycamore":                       ColorSycamore,
		"umber":                          ColorUmber,
		"bianca":                         ColorBianca,
		"pastelred":                      ColorPastelRed,
		"barleywhite":                    ColorBarleyWhite,
		"junglemist":                     ColorJungleMist,
		"tapa":                           ColorTapa,
		"hillary":                        ColorHillary,
		"puertorico":                     ColorPuertoRico,
		"revolver":                       ColorRevolver,
		"richgold":                       ColorRichGold,
		"skyblue":                        ColorSkyBlue,
		"bluesapphire":                   ColorBlueSapphire,
		"chromewhite":                    ColorChromeWhite,
		"metalpink":                      ColorMetalPink,
		"scandal":                        ColorScandal,
		"camarone":                       ColorCamarone,
		"envy":                           ColorEnvy,
		"lightorchid":                    ColorLightOrchid,
		"eagle":                          ColorEagle,
		"eerieblack":                     ColorEerieBlack,
		"frenchgray":                     ColorFrenchGray,
		"starship":                       ColorStarship,
		"tahitigold":                     ColorTahitiGold,
		"bubbles":                        ColorBubbles,
		"cornharvest":                    ColorCornHarvest,
		"siren":                          ColorSiren,
		"deepcerise":                     ColorDeepCerise,
		"dolphin":                        ColorDolphin,
		"ziggurat":                       ColorZiggurat,
		"mineshaft":                      ColorMineShaft,
		"rosebud":                        ColorRoseBud,
		"flamenco":                       ColorFlamenco,
		"kucrimson":                      ColorKUCrimson,
		"aureolin":                       ColorAureolin,
		"deer":                           ColorDeer,
		"twilightblue":                   ColorTwilightBlue,
		"mamba":                          ColorMamba,
		"softpeach":                      ColorSoftPeach,
		"pharlap":                        ColorPharlap,
		"vividtangelo":                   ColorVividTangelo,
		"affair":                         ColorAffair,
		"brilliantazure":                 ColorBrilliantAzure,
		"shalimar":                       ColorShalimar,
		"cyancornflowerblue":             ColorCyanCornflowerBlue,
		"pigmentred":                     ColorPigmentRed,
		"tropicalblue":                   ColorTropicalBlue,
		"truev":                          ColorTrueV,
		"mexicanpink":                    ColorMexicanPink,
		"munsellred":                     ColorMunsellRed,
		"burntorange":                    ColorBurntOrange,
		"mauve":                          ColorMauve,
		"mediumspringbud":                ColorMediumSpringBud,
		"pirategold":                     ColorPirateGold,
		"ryborange":                      ColorRYBOrange,
		"axolotl":                        ColorAxolotl,
		"blackmarlin":                    ColorBlackMarlin,
		"stromboli":                      ColorStromboli,
		"sulu":                           ColorSulu,
		"upmaroon":                       ColorUPMaroon,
		"kimberly":                       ColorKimberly,
		"mediumvermilion":                ColorMediumVermilion,
		"goldenbrown":                    ColorGoldenBrown,
		"hippiegreen":                    ColorHippieGreen,
		"meteorite":                      ColorMeteorite,
		"turkishrose":                    ColorTurkishRose,
		"coriander":                      ColorCoriander,
		"lightfrenchbeige":               ColorLightFrenchBeige,
		"feta":                           ColorFeta,
		"giantsorange":                   ColorGiantsOrange,
		"hawaiiantan":                    ColorHawaiianTan,
		"lightcyan":                      ColorLightCyan,
		"paua":                           ColorPaua,
		"dartmouthgreen":                 ColorDartmouthGreen,
		"doublepearllusta":               ColorDoublePearlLusta,
		"lavenderrose":                   ColorLavenderRose,
		"minionyellow":                   ColorMinionYellow,
		"ottoman":                        ColorOttoman,
		"eucalyptus":                     ColorEucalyptus,
		"hairyheath":                     ColorHairyHeath,
		"midgray":                        ColorMidGray,
		"ncsred":                         ColorNCSRed,
		"antiquewhite":                   ColorAntiqueWhite,
		"malibu":                         ColorMalibu,
		"olivehaze":                      ColorOliveHaze,
		"palecanary":                     ColorPaleCanary,
		"urobilin":                       ColorUrobilin,
		"korma":                          ColorKorma,
		"mocha":                          ColorMocha,
		"jewel":                          ColorJewel,
		"mintgreen":                      ColorMintGreen,
		"burntsienna":                    ColorBurntSienna,
		"crayolayellow":                  ColorCrayolaYellow,
		"jazzberryjam":                   ColorJazzberryJam,
		"monarch":                        ColorMonarch,
		"rope":                           ColorRope,
		"whisper":                        ColorWhisper,
		"bluegray":                       ColorBlueGray,
		"cherryblossompink":              ColorCherryBlossomPink,
		"primrose":                       ColorPrimrose,
		"butterywhite":                   ColorButteryWhite,
		"orangepeel":                     ColorOrangePeel,
		"beaver":                         ColorBeaver,
		"liver":                          ColorLiver,
		"falcon":                         ColorFalcon,
		"gurkha":                         ColorGurkha,
		"maverick":                       ColorMaverick,
		"bluelagoon":                     ColorBlueLagoon,
		"cowboy":                         ColorCowboy,
		"azuremist":                      ColorAzureMist,
		"kellygreen":                     ColorKellyGreen,
		"mint":                           ColorMint,
		"pictonblue":                     ColorPictonBlue,
		"africanviolet":                  ColorAfricanViolet,
		"aquamarine":                     ColorAquamarine,
		"pinkpearl":                      ColorPinkPearl,
		"keylimepie":                     ColorKeyLimePie,
		"como":                           ColorComo,
		"dandelion":                      ColorDandelion,
		"sherpablue":                     ColorSherpaBlue,
		"turbo":                          ColorTurbo,
		"dirt":                           ColorDirt,
		"paynesgrey":                     ColorPaynesGrey,
		"milkpunch":                      ColorMilkPunch,
		"vinrouge":                       ColorVinRouge,
		"wineberry":                      ColorWineBerry,
		"bluebell":                       ColorBlueBell,
		"chocolate":                      ColorChocolate,
		"keylime":                        ColorKeyLime,
		"mediumcandyapplered":            ColorMediumCandyAppleRed,
		"midnightblue":                   ColorMidnightBlue,
		"uclablue":                       ColorUCLABlue,
		"burntmaroon":                    ColorBurntMaroon,
		"cyprus":                         ColorCyprus,
		"nileblue":                       ColorNileBlue,
		"spanishviridian":                ColorSpanishViridian,
		"ultrapink":                      ColorUltraPink,
		"bigdiporuby":                    ColorBigDipOruby,
		"greenblue":                      ColorGreenBlue,
		"negroni":                        ColorNegroni,
		"wheatfield":                     ColorWheatfield,
		"bluestone":                      ColorBlueStone,
		"darkcerulean":                   ColorDarkCerulean,
		"lightgoldenrodyellow":           ColorLightGoldenrodYellow,
		"pacifika":                       ColorPacifika,
		"santafe":                        ColorSantaFe,
		"schooner":                       ColorSchooner,
		"woodland":                       ColorWoodland,
		"butteredrum":                    ColorButteredRum,
		"cadet":                          ColorCadet,
		"spanishpink":                    ColorSpanishPink,
		"woodrush":                       ColorWoodrush,
		"gigas":                          ColorGigas,
		"oxley":                          ColorOxley,
		"springbud":                      ColorSpringBud,
		"tasman":                         ColorTasman,
		"wepeep":                         ColorWePeep,
		"codgray":                        ColorCodGray,
		"darksienna":                     ColorDarkSienna,
		"licorice":                       ColorLicorice,
		"rosybrown":                      ColorRosyBrown,
		"outrageousorange":               ColorOutrageousOrange,
		"rainee":                         ColorRainee,
		"whitelilac":                     ColorWhiteLilac,
		"castletongreen":                 ColorCastletonGreen,
		"copperred":                      ColorCopperRed,
		"pearlaqua":                      ColorPearlAqua,
		"darkblue":                       ColorDarkBlue,
		"japaneseindigo":                 ColorJapaneseIndigo,
		"pastelyellow":                   ColorPastelYellow,
		"poloblue":                       ColorPoloBlue,
		"bunker":                         ColorBunker,
		"canaryyellow":                   ColorCanaryYellow,
		"lemonchiffon":                   ColorLemonChiffon,
		"palechestnut":                   ColorPaleChestnut,
		"stpatricksblue":                 ColorStPatricksBlue,
		"viking":                         ColorViking,
		"darkskyblue":                    ColorDarkSkyBlue,
		"downy":                          ColorDowny,
		"maximumblue":                    ColorMaximumBlue,
		"razzmicberry":                   ColorRazzmicBerry,
		"redoxide":                       ColorRedOxide,
		"vividorangepeel":                ColorVividOrangePeel,
		"yukongold":                      ColorYukonGold,
		"bouquet":                        ColorBouquet,
		"eggwhite":                       ColorEggWhite,
		"martini":                        ColorMartini,
		"cherokee":                       ColorCherokee,
		"fuchsiablue":                    ColorFuchsiaBlue,
		"imperial":                       ColorImperial,
		"saharasand":                     ColorSaharaSand,
		"willpowerorange":                ColorWillpowerOrange,
		"bastille":                       ColorBastille,
		"blanchedalmond":                 ColorBlanchedAlmond,
		"munsellgreen":                   ColorMunsellGreen,
		"orange":                         ColorOrange,
		"hintofyellow":                   ColorHintofYellow,
		"kobi":                           ColorKobi,
		"naturalgray":                    ColorNaturalGray,
		"yellowmetal":                    ColorYellowMetal,
		"apricot":                        ColorApricot,
		"buddhagold":                     ColorBuddhaGold,
		"quarterpearllusta":              ColorQuarterPearlLusta,
		"lightningyellow":                ColorLightningYellow,
		"oil":                            ColorOil,
		"halfandhalf":                    ColorHalfandHalf,
		"indigodye":                      ColorIndigoDye,
		"raven":                          ColorRaven,
		"shadow":                         ColorShadow,
		"tussock":                        ColorTussock,
		"babyblue":                       ColorBabyBlue,
		"geebung":                        ColorGeebung,
		"internationalkleinblue":         ColorInternationalKleinBlue,
		"redviolet":                      ColorRedViolet,
		"whitesmoke":                     ColorWhiteSmoke,
		"apricotwhite":                   ColorApricotWhite,
		"congopink":                      ColorCongoPink,
		"surfcrest":                      ColorSurfCrest,
		"teal":                           ColorTeal,
		"blackolive":                     ColorBlackOlive,
		"cadetblue":                      ColorCadetBlue,
		"dukeblue":                       ColorDukeBlue,
		"equator":                        ColorEquator,
		"oceanboatblue":                  ColorOceanBoatBlue,
		"cameo":                          ColorCameo,
		"dell":                           ColorDell,
		"blackleatherjacket":             ColorBlackLeatherJacket,
		"cashmere":                       ColorCashmere,
		"jumbo":                          ColorJumbo,
		"merlot":                         ColorMerlot,
		"silver":                         ColorSilver,
		"acapulco":                       ColorAcapulco,
		"avocado":                        ColorAvocado,
		"electricviolet":                 ColorElectricViolet,
		"gablegreen":                     ColorGableGreen,
		"graynurse":                      ColorGrayNurse,
		"maize":                          ColorMaize,
		"mandy":                          ColorMandy,
		"black":                          ColorBlack,
		"chelseagem":                     ColorChelseaGem,
		"marshland":                      ColorMarshland,
		"opium":                          ColorOpium,
		"tara":                           ColorTara,
		"tiffanyblue":                    ColorTiffanyBlue,
		"zombie":                         ColorZombie,
		"cadillac":                       ColorCadillac,
		"deepjunglegreen":                ColorDeepJungleGreen,
		"lemon":                          ColorLemon,
		"moonglow":                       ColorMoonGlow,
		"cancan":                         ColorCanCan,
		"deepcarminepink":                ColorDeepCarminePink,
		"citrine":                        ColorCitrine,
		"twilight":                       ColorTwilight,
		"neworleans":                     ColorNewOrleans,
		"pinkswan":                       ColorPinkSwan,
		"kokoda":                         ColorKokoda,
		"lapislazuli":                    ColorLapisLazuli,
		"paprika":                        ColorPaprika,
		"princetonorange":                ColorPrincetonOrange,
		"uared":                          ColorUARed,
		"candlelight":                    ColorCandlelight,
		"downriver":                      ColorDownriver,
		"forestgreen":                    ColorForestGreen,
		"moonstoneblue":                  ColorMoonstoneBlue,
		"blond":                          ColorBlond,
		"coffee":                         ColorCoffee,
		"makara":                         ColorMakara,
		"raffia":                         ColorRaffia,
		"scooter":                        ColorScooter,
		"tusk":                           ColorTusk,
		"brandyrose":                     ColorBrandyRose,
		"lisbonbrown":                    ColorLisbonBrown,
		"limedash":                       ColorLimedAsh,
		"linen":                          ColorLinen,
		"oldcopper":                      ColorOldCopper,
		"palerobineggblue":               ColorPaleRobinEggBlue,
		"rum":                            ColorRum,
		"weborange":                      ColorWebOrange,
		"fantasy":                        ColorFantasy,
		"gravel":                         ColorGravel,
		"brightsun":                      ColorBrightSun,
		"crocodile":                      ColorCrocodile,
		"bayofmany":                      ColorBayofMany,
		"bistre":                         ColorBistre,
		"desertstorm":                    ColorDesertStorm,
		"salomie":                        ColorSalomie,
		"verylightazure":                 ColorVeryLightAzure,
		"chardonnay":                     ColorChardonnay,
		"japaneselaurel":                 ColorJapaneseLaurel,
		"fog":                            ColorFog,
		"rawumber":                       ColorRawUmber,
		"royalheath":                     ColorRoyalHeath,
		"elpaso":                         ColorElPaso,
		"fedora":                         ColorFedora,
		"trinidad":                       ColorTrinidad,
		"dawn":                           ColorDawn,
		"genericviridian":                ColorGenericViridian,
		"darkgreen":                      ColorDarkGreen,
		"crimsonred":                     ColorCrimsonRed,
		"karaka":                         ColorKaraka,
		"smalt":                          ColorSmalt,
		"steelgray":                      ColorSteelGray,
		"sundance":                       ColorSundance,
		"almondfrost":                    ColorAlmondFrost,
		"congobrown":                     ColorCongoBrown,
		"halayàúbe":                      ColorHalayàÚbe,
		"quicksilver":                    ColorQuickSilver,
		"tea":                            ColorTea,
		"darkpuce":                       ColorDarkPuce,
		"frenchviolet":                   ColorFrenchViolet,
		"lightsalmonpink":                ColorLightSalmonPink,
		"nightrider":                     ColorNightRider,
		"pinkraspberry":                  ColorPinkRaspberry,
		"rebeccapurple":                  ColorRebeccaPurple,
		"sunburntcyclops":                ColorSunburntCyclops,
		"goldengatebridge":               ColorGoldenGateBridge,
		"governorbay":                    ColorGovernorBay,
		"ebb":                            ColorEbb,
		"pohutukawa":                     ColorPohutukawa,
		"coyotebrown":                    ColorCoyoteBrown,
		"lilacbush":                      ColorLilacBush,
		"riverbed":                       ColorRiverBed,
		"sidecar":                        ColorSidecar,
		"zorba":                          ColorZorba,
		"horizon":                        ColorHorizon,
		"pattensblue":                    ColorPattensBlue,
		"pewterblue":                     ColorPewterBlue,
		"pinklace":                       ColorPinkLace,
		"queenblue":                      ColorQueenBlue,
		"studio":                         ColorStudio,
		"lasallegreen":                   ColorLaSalleGreen,
		"peat":                           ColorPeat,
		"pizza":                          ColorPizza,
		"utahcrimson":                    ColorUtahCrimson,
		"valencia":                       ColorValencia,
		"cosmic":                         ColorCosmic,
		"lotus":                          ColorLotus,
		"wenge":                          ColorWenge,
		"breakerbay":                     ColorBreakerBay,
		"salmon":                         ColorSalmon,
		"ebonyclay":                      ColorEbonyClay,
		"lightgrayishmagenta":            ColorLightGrayishMagenta,
		"teak":                           ColorTeak,
		"winterwizard":                   ColorWinterWizard,
		"amethyst":                       ColorAmethyst,
		"daintree":                       ColorDaintree,
		"oriolesorange":                  ColorOriolesOrange,
		"paradisepink":                   ColorParadisePink,
		"snowdrift":                      ColorSnowDrift,
		"darkorchid":                     ColorDarkOrchid,
		"lavenderindigo":                 ColorLavenderIndigo,
		"cyclamen":                       ColorCyclamen,
		"titanwhite":                     ColorTitanWhite,
		"universityoftennesseeorange":    ColorUniversityOfTennesseeOrange,
		"indochine":                      ColorIndochine,
		"spartancrimson":                 ColorSpartanCrimson,
		"veniceblue":                     ColorVeniceBlue,
		"catskillwhite":                  ColorCatskillWhite,
		"hookersgreen":                   ColorHookersGreen,
		"msugreen":                       ColorMSUGreen,
		"persianblue":                    ColorPersianBlue,
		"persimmon":                      ColorPersimmon,
		"pompadour":                      ColorPompadour,
		"capecod":                        ColorCapeCod,
		"dairycream":                     ColorDairyCream,
		"isabelline":                     ColorIsabelline,
		"tosca":                          ColorTosca,
		"bazaar":                         ColorBazaar,
		"crayolaorange":                  ColorCrayolaOrange,
		"mediumturquoise":                ColorMediumTurquoise,
		"metallicsunburst":               ColorMetallicSunburst,
		"napa":                           ColorNapa,
		"nickel":                         ColorNickel,
		"periwinklegray":                 ColorPeriwinkleGray,
		"riflegreen":                     ColorRifleGreen,
		"coffeebean":                     ColorCoffeeBean,
		"foggygray":                      ColorFoggyGray,
		"tenne":                          ColorTenne,
		"pastelorange":                   ColorPastelOrange,
		"peachschnapps":                  ColorPeachSchnapps,
		"piggypink":                      ColorPiggyPink,
		"seablue":                        ColorSeaBlue,
		"bisonhide":                      ColorBisonHide,
		"lapalma":                        ColorLaPalma,
		"saltbox":                        ColorSaltBox,
		"zircon":                         ColorZircon,
		"doublecolonialwhite":            ColorDoubleColonialWhite,
		"rosedust":                       ColorRoseDust,
		"driftwood":                      ColorDriftwood,
		"quarterspanishwhite":            ColorQuarterSpanishWhite,
		"rufous":                         ColorRufous,
		"clamshell":                      ColorClamShell,
		"deepfuchsia":                    ColorDeepFuchsia,
		"norway":                         ColorNorway,
		"camelot":                        ColorCamelot,
		"horses":                         ColorHorses,
		"burningsand":                    ColorBurningSand,
		"cupid":                          ColorCupid,
		"cyancobaltblue":                 ColorCyanCobaltBlue,
		"lincolngreen":                   ColorLincolnGreen,
		"palecornflowerblue":             ColorPaleCornflowerBlue,
		"alienarmpit":                    ColorAlienArmpit,
		"bluegreen":                      ColorBlueGreen,
		"flushmahogany":                  ColorFlushMahogany,
		"parisdaisy":                     ColorParisDaisy,
		"roman":                          ColorRoman,
		"squirrel":                       ColorSquirrel,
		"westside":                       ColorWestSide,
		"citrinewhite":                   ColorCitrineWhite,
		"fairpink":                       ColorFairPink,
		"mysticmaroon":                   ColorMysticMaroon,
		"perano":                         ColorPerano,
		"electriclime":                   ColorElectricLime,
		"mangotango":                     ColorMangoTango,
		"paoloveronesegreen":             ColorPaoloVeroneseGreen,
		"brownsugar":                     ColorBrownSugar,
		"fuscousgray":                    ColorFuscousGray,
		"prelude":                        ColorPrelude,
		"sizzlingred":                    ColorSizzlingRed,
		"snowflurry":                     ColorSnowFlurry,
		"thunder":                        ColorThunder,
		"coraltree":                      ColorCoralTree,
		"mediumpurple":                   ColorMediumPurple,
		"buttercup":                      ColorButtercup,
		"crusoe":                         ColorCrusoe,
		"darkkhaki":                      ColorDarkKhaki,
		"deepsea":                        ColorDeepSea,
		"palerose":                       ColorPaleRose,
		"aquasqueeze":                    ColorAquaSqueeze,
		"blastoffbronze":                 ColorBlastOffBronze,
		"fiord":                          ColorFiord,
		"royalpurple":                    ColorRoyalPurple,
		"vistablue":                      ColorVistaBlue,
		"cherrypie":                      ColorCherryPie,
		"darkcandyapplered":              ColorDarkCandyAppleRed,
		"sanguinebrown":                  ColorSanguineBrown,
		"gorse":                          ColorGorse,
		"mayablue":                       ColorMayaBlue,
		"laser":                          ColorLaser,
		"pancho":                         ColorPancho,
		"stack":                          ColorStack,
		"astronautblue":                  ColorAstronautBlue,
		"killarney":                      ColorKillarney,
		"swansdown":                      ColorSwansDown,
		"darksalmon":                     ColorDarkSalmon,
		"redpurple":                      ColorRedPurple,
		"gumleaf":                        ColorGumLeaf,
		"logan":                          ColorLogan,
		"purpleplum":                     ColorPurplePlum,
		"zuccini":                        ColorZuccini,
		"cosmiccobalt":                   ColorCosmicCobalt,
		"gambogeorange":                  ColorGambogeOrange,
		"judgegray":                      ColorJudgeGray,
		"razzledazzlerose":               ColorRazzleDazzleRose,
		"selago":                         ColorSelago,
		"spirodiscoball":                 ColorSpiroDiscoBall,
		"brightgreen":                    ColorBrightGreen,
		"darkmidnightblue":               ColorDarkMidnightBlue,
		"puce":                           ColorPuce,
		"fuego":                          ColorFuego,
		"hampton":                        ColorHampton,
		"mercury":                        ColorMercury,
		"periwinkle":                     ColorPeriwinkle,
		"quincy":                         ColorQuincy,
		"aubergine":                      ColorAubergine,
		"cherrywood":                     ColorCherrywood,
		"ufogreen":                       ColorUFOGreen,
		"rollingstone":                   ColorRollingStone,
		"spunpearl":                      ColorSpunPearl,
		"deepteal":                       ColorDeepTeal,
		"liberty":                        ColorLiberty,
		"magicpotion":                    ColorMagicPotion,
		"martinique":                     ColorMartinique,
		"peppermint":                     ColorPeppermint,
		"rust":                           ColorRust,
		"amethystsmoke":                  ColorAmethystSmoke,
		"blackwhite":                     ColorBlackWhite,
		"limedoak":                       ColorLimedOak,
		"twine":                          ColorTwine,
		"havelockblue":                   ColorHavelockBlue,
		"ruddypink":                      ColorRuddyPink,
		"vividcrimson":                   ColorVividCrimson,
		"bluezodiac":                     ColorBlueZodiac,
		"chateaugreen":                   ColorChateauGreen,
		"raspberry":                      ColorRaspberry,
		"rodeodust":                      ColorRodeoDust,
		"tanhide":                        ColorTanHide,
		"cobaltblue":                     ColorCobaltBlue,
		"galliano":                       ColorGalliano,
		"chinook":                        ColorChinook,
		"clover":                         ColorClover,
		"fuchsiapink":                    ColorFuchsiaPink,
		"jagger":                         ColorJagger,
		"persianindigo":                  ColorPersianIndigo,
		"ube":                            ColorUbe,
		"verylightmalachitegreen":        ColorVeryLightMalachiteGreen,
		"barleycorn":                     ColorBarleyCorn,
		"cambridgeblue":                  ColorCambridgeBlue,
		"roofterracotta":                 ColorRoofTerracotta,
		"yellowsea":                      ColorYellowSea,
		"caribbeangreen":                 ColorCaribbeanGreen,
		"giantsclub":                     ColorGiantsClub,
		"lime":                           ColorLime,
		"pumice":                         ColorPumice,
		"shampoo":                        ColorShampoo,
		"silk":                           ColorSilk,
		"solidpink":                      ColorSolidPink,
		"westcoast":                      ColorWestCoast,
		"carminepink":                    ColorCarminePink,
		"granitegreen":                   ColorGraniteGreen,
		"oslogray":                       ColorOsloGray,
		"saltpan":                        ColorSaltpan,
		"tallpoppy":                      ColorTallPoppy,
		"titaniumyellow":                 ColorTitaniumYellow,
		"voodoo":                         ColorVoodoo,
		"akaroa":                         ColorAkaroa,
		"costadelsol":                    ColorCostaDelSol,
		"wafer":                          ColorWafer,
		"beeswax":                        ColorBeeswax,
		"pesto":                          ColorPesto,
		"maitai":                         ColorMaiTai,
		"paleviolet":                     ColorPaleViolet,
		"putty":                          ColorPutty,
		"scarpaflow":                     ColorScarpaFlow,
		"burgundy":                       ColorBurgundy,
		"greenpea":                       ColorGreenPea,
		"cement":                         ColorCement,
		"cherub":                         ColorCherub,
		"brightred":                      ColorBrightRed,
		"brownrust":                      ColorBrownRust,
		"ncsgreen":                       ColorNCSGreen,
		"orientalpink":                   ColorOrientalPink,
		"straw":                          ColorStraw,
		"bluebayoux":                     ColorBlueBayoux,
		"green":                          ColorGreen,
		"safetyorange":                   ColorSafetyOrange,
		"spindle":                        ColorSpindle,
		"caferoyale":                     ColorCafeRoyale,
		"mondo":                          ColorMondo,
		"eggshell":                       ColorEggshell,
		"fandangopink":                   ColorFandangoPink,
		"gunsmoke":                       ColorGunsmoke,
		"iroko":                          ColorIroko,
		"melanzane":                      ColorMelanzane,
		"punch":                          ColorPunch,
		"calpolygreen":                   ColorCalPolyGreen,
		"cedar":                          ColorCedar,
		"creamcan":                       ColorCreamCan,
		"illusion":                       ColorIllusion,
		"sonicsilver":                    ColorSonicSilver,
		"tana":                           ColorTana,
		"apache":                         ColorApache,
		"aurometalsaurus":                ColorAuroMetalSaurus,
		"earlsgreen":                     ColorEarlsGreen,
		"lightskyblue":                   ColorLightSkyBlue,
		"bleachwhite":                    ColorBleachWhite,
		"celtic":                         ColorCeltic,
		"burningorange":                  ColorBurningOrange,
		"crimsonglory":                   ColorCrimsonGlory,
		"crownofthorns":                  ColorCrownofThorns,
		"deepmagenta":                    ColorDeepMagenta,
		"jade":                           ColorJade,
		"koromiko":                       ColorKoromiko,
		"bluesmoke":                      ColorBlueSmoke,
		"bracken":                        ColorBracken,
		"purpureus":                      ColorPurpureus,
		"riogrande":                      ColorRioGrande,
		"oracle":                         ColorOracle,
		"pinkorange":                     ColorPinkOrange,
		"timbergreen":                    ColorTimberGreen,
		"freshair":                       ColorFreshAir,
		"mulledwine":                     ColorMulledWine,
		"razzmatazz":                     ColorRazzmatazz,
		"thulianpink":                    ColorThulianPink,
		"blumine":                        ColorBlumine,
		"ogreodor":                       ColorOgreOdor,
		"opal":                           ColorOpal,
		"catawba":                        ColorCatawba,
		"kashmirblue":                    ColorKashmirBlue,
		"nyanza":                         ColorNyanza,
		"shamrockgreen":                  ColorShamrockGreen,
		"fuchsiarose":                    ColorFuchsiaRose,
		"moccaccino":                     ColorMoccaccino,
		"claycreek":                      ColorClayCreek,
		"ming":                           ColorMing,
		"oldmossgreen":                   ColorOldMossGreen,
		"snow":                           ColorSnow,
		"springfrost":                    ColorSpringFrost,
		"catalinablue":                   ColorCatalinaBlue,
		"celadon":                        ColorCeladon,
		"tyrianpurple":                   ColorTyrianPurple,
		"russett":                        ColorRussett,
		"sail":                           ColorSail,
		"ticklemepink":                   ColorTickleMePink,
		"unitednationsblue":              ColorUnitedNationsBlue,
		"verylighttangelo":               ColorVeryLightTangelo,
		"chardon":                        ColorChardon,
		"pixiegreen":                     ColorPixieGreen,
		"kilamanjaro":                    ColorKilamanjaro,
		"roti":                           ColorRoti,
		"wheat":                          ColorWheat,
		"zambezi":                        ColorZambezi,
		"airsuperiorityblue":             ColorAirSuperiorityBlue,
		"cyan":                           ColorCyan,
		"crayolagreen":                   ColorCrayolaGreen,
		"spice":                          ColorSpice,
		"sushi":                          ColorSushi,
		"trendypink":                     ColorTrendyPink,
		"yourpink":                       ColorYourPink,
		"calypso":                        ColorCalypso,
		"carrotorange":                   ColorCarrotOrange,
		"silversand":                     ColorSilverSand,
		"tangerine":                      ColorTangerine,
		"gunmetal":                       ColorGunmetal,
		"sienna":                         ColorSienna,
		"citron":                         ColorCitron,
		"derby":                          ColorDerby,
		"goben":                          ColorGoBen,
		"lightsteelblue":                 ColorLightSteelBlue,
		"loblolly":                       ColorLoblolly,
		"scarletgum":                     ColorScarletGum,
		"azure":                          ColorAzure,
		"brickred":                       ColorBrickRed,
		"wintersky":                      ColorWinterSky,
		"drover":                         ColorDrover,
		"mojo":                           ColorMojo,
		"redstage":                       ColorRedStage,
		"telemagenta":                    ColorTelemagenta,
		"budgreen":                       ColorBudGreen,
		"deepred":                        ColorDeepRed,
		"mosque":                         ColorMosque,
		"munsellpurple":                  ColorMunsellPurple,
		"sepiablack":                     ColorSepiaBlack,
		"astronaut":                      ColorAstronaut,
		"glacier":                        ColorGlacier,
		"sanfelix":                       ColorSanFelix,
		"vividauburn":                    ColorVividAuburn,
		"iron":                           ColorIron,
		"palegreen":                      ColorPaleGreen,
		"lunargreen":                     ColorLunarGreen,
		"montecarlo":                     ColorMonteCarlo,
		"sheengreen":                     ColorSheenGreen,
		"americano":                      ColorAmericano,
		"hibiscus":                       ColorHibiscus,
		"russiangreen":                   ColorRussianGreen,
		"tiber":                          ColorTiber,
		"himalaya":                       ColorHimalaya,
		"lawngreen":                      ColorLawnGreen,
		"pastelmagenta":                  ColorPastelMagenta,
		"darkfern":                       ColorDarkFern,
		"oregon":                         ColorOregon,
		"japanesemaple":                  ColorJapaneseMaple,
		"sprout":                         ColorSprout,
		"steelteal":                      ColorSteelTeal,
		"alpine":                         ColorAlpine,
		"domino":                         ColorDomino,
		"hummingbird":                    ColorHummingBird,
		"pigeonpost":                     ColorPigeonPost,
		"brightmaroon":                   ColorBrightMaroon,
		"halfdutchwhite":                 ColorHalfDutchWhite,
		"sinopia":                        ColorSinopia,
		"cumulus":                        ColorCumulus,
		"royalazure":                     ColorRoyalAzure,
		"foam":                           ColorFoam,
		"slimygreen":                     ColorSlimyGreen,
		"laurelgreen":                    ColorLaurelGreen,
		"lemongrass":                     ColorLemonGrass,
		"persianrose":                    ColorPersianRose,
		"prim":                           ColorPrim,
		"wistful":                        ColorWistful,
		"caramel":                        ColorCaramel,
		"dollarbill":                     ColorDollarBill,
		"skobeloff":                      ColorSkobeloff,
		"atlantis":                       ColorAtlantis,
		"cararra":                        ColorCararra,
		"sunsetorange":                   ColorSunsetOrange,
		"tango":                          ColorTango,
		"harvestgold":                    ColorHarvestGold,
		"lightkhaki":                     ColorLightKhaki,
		"santasgray":                     ColorSantasGray,
		"honolulublue":                   ColorHonoluluBlue,
		"orient":                         ColorOrient,
		"gullgray":                       ColorGullGray,
		"oldgold":                        ColorOldGold,
		"pinksherbet":                    ColorPinkSherbet,
		"plantation":                     ColorPlantation,
		"sisal":                          ColorSisal,
		"brownyellow":                    ColorBrownYellow,
		"empress":                        ColorEmpress,
		"icecold":                        ColorIceCold,
		"mantis":                         ColorMantis,
		"orchidwhite":                    ColorOrchidWhite,
		"pippin":                         ColorPippin,
		"starcommandblue":                ColorStarCommandBlue,
		"auchico":                        ColorAuChico,
		"ebony":                          ColorEbony,
		"pantoneorange":                  ColorPantoneOrange,
		"sanjuan":                        ColorSanJuan,
		"smashedpumpkin":                 ColorSmashedPumpkin,
		"tepapagreen":                    ColorTePapaGreen,
		"topaz":                          ColorTopaz,
		"tundora":                        ColorTundora,
		"chileanheath":                   ColorChileanHeath,
		"flamepea":                       ColorFlamePea,
		"wildwatermelon":                 ColorWildWatermelon,
		"portage":                        ColorPortage,
		"sapgreen":                       ColorSapGreen,
		"japonica":                       ColorJaponica,
		"keppel":                         ColorKeppel,
		"redribbon":                      ColorRedRibbon,
		"bahamablue":                     ColorBahamaBlue,
		"burnham":                        ColorBurnham,
		"linkwater":                      ColorLinkWater,
		"magentapink":                    ColorMagentaPink,
		"reddevil":                       ColorRedDevil,
		"hanblue":                        ColorHanBlue,
		"lily":                           ColorLily,
		"neongreen":                      ColorNeonGreen,
		"redrobin":                       ColorRedRobin,
		"sizzlingsunrise":                ColorSizzlingSunrise,
		"halfcolonialwhite":              ColorHalfColonialWhite,
		"japaneseviolet":                 ColorJapaneseViolet,
		"cocoabean":                      ColorCocoaBean,
		"canary":                         ColorCanary,
		"cinnamonsatin":                  ColorCinnamonSatin,
		"bluecharcoal":                   ColorBlueCharcoal,
		"burlywood":                      ColorBurlywood,
		"rosebonbon":                     ColorRoseBonbon,
		"saratoga":                       ColorSaratoga,
		"mortar":                         ColorMortar,
		"pantoneyellow":                  ColorPantoneYellow,
		"tangaroa":                       ColorTangaroa,
		"barossa":                        ColorBarossa,
		"givry":                          ColorGivry,
		"darktangerine":                  ColorDarkTangerine,
		"oldburgundy":                    ColorOldBurgundy,
		"aquaspring":                     ColorAquaSpring,
		"cascade":                        ColorCascade,
		"flirt":                          ColorFlirt,
		"denimblue":                      ColorDenimBlue,
		"larioja":                        ColorLaRioja,
		"bronzeyellow":                   ColorBronzeYellow,
		"cerulean":                       ColorCerulean,
		"limerick":                       ColorLimerick,
		"rosetaupe":                      ColorRoseTaupe,
		"spanishbistre":                  ColorSpanishBistre,
		"wedgewood":                      ColorWedgewood,
		"deeppuce":                       ColorDeepPuce,
		"lightsalmon":                    ColorLightSalmon,
		"anzac":                          ColorAnzac,
		"laspalmas":                      ColorLasPalmas,
		"lochinvar":                      ColorLochinvar,
		"neonfuchsia":                    ColorNeonFuchsia,
		"papayawhip":                     ColorPapayaWhip,
		"platinum":                       ColorPlatinum,
		"tropicalviolet":                 ColorTropicalViolet,
		"asparagus":                      ColorAsparagus,
		"desaturatedcyan":                ColorDesaturatedCyan,
		"verdungreen":                    ColorVerdunGreen,
		"everglade":                      ColorEverglade,
		"lenurple":                       ColorLenurple,
		"magicmint":                      ColorMagicMint,
		"nadeshikopink":                  ColorNadeshikoPink,
		"soap":                           ColorSoap,
		"tan":                            ColorTan,
		"bluechalk":                      ColorBlueChalk,
		"bourbon":                        ColorBourbon,
		"violeteggplant":                 ColorVioletEggplant,
		"rosebudcherry":                  ColorRoseBudCherry,
		"spanishred":                     ColorSpanishRed,
		"flaxsmoke":                      ColorFlaxSmoke,
		"rocketmetallic":                 ColorRocketMetallic,
		"apple":                          ColorApple,
		"darkslategray":                  ColorDarkSlateGray,
		"chaletgreen":                    ColorChaletGreen,
		"moccasin":                       ColorMoccasin,
		"tahunasands":                    ColorTahunaSands,
		"amazon":                         ColorAmazon,
		"blackhaze":                      ColorBlackHaze,
		"darkpastelblue":                 ColorDarkPastelBlue,
		"zeus":                           ColorZeus,
		"zumthor":                        ColorZumthor,
		"bittersweet":                    ColorBittersweet,
		"cadetgrey":                      ColorCadetGrey,
		"brightturquoise":                ColorBrightTurquoise,
		"tealdeer":                       ColorTealDeer,
		"aztecgold":                      ColorAztecGold,
		"pipi":                           ColorPipi,
		"unmellowyellow":                 ColorUnmellowYellow,
		"pullmanbrown":                   ColorPullmanBrown,
		"tuatara":                        ColorTuatara,
		"desertsand":                     ColorDesertSand,
		"paleoyster":                     ColorPaleOyster,
		"pastelblue":                     ColorPastelBlue,
		"reef":                           ColorReef,
		"seashell":                       ColorSeashell,
		"charade":                        ColorCharade,
		"deepcove":                       ColorDeepCove,
		"lilywhite":                      ColorLilyWhite,
		"sherwoodgreen":                  ColorSherwoodGreen,
		"witchhaze":                      ColorWitchHaze,
		"englishvermillion":              ColorEnglishVermillion,
		"lemonlime":                      ColorLemonLime,
		"engineeringinternationalorange": ColorEngineeringInternationalOrange,
		"indiagreen":                     ColorIndiaGreen,
		"spanishskyblue":                 ColorSpanishSkyBlue,
		"springgreen":                    ColorSpringGreen,
		"sweetbrown":                     ColorSweetBrown,
		"copperrose":                     ColorCopperRose,
		"egyptianblue":                   ColorEgyptianBlue,
		"sapphireblue":                   ColorSapphireBlue,
		"quillgray":                      ColorQuillGray,
		"rosepink":                       ColorRosePink,
		"greenhouse":                     ColorGreenHouse,
		"haiti":                          ColorHaiti,
		"mako":                           ColorMako,
		"seanymph":                       ColorSeaNymph,
		"charlotte":                      ColorCharlotte,
		"cornflowerlilac":                ColorCornflowerLilac,
		"guardsmanred":                   ColorGuardsmanRed,
		"mandalay":                       ColorMandalay,
		"nugget":                         ColorNugget,
		"scampi":                         ColorScampi,
		"bronco":                         ColorBronco,
		"cabaret":                        ColorCabaret,
		"chino":                          ColorChino,
		"windsor":                        ColorWindsor,
		"chantilly":                      ColorChantilly,
		"darkbrown":                      ColorDarkBrown,
		"falured":                        ColorFaluRed,
		"ironstone":                      ColorIronstone,
		"mistyrose":                      ColorMistyRose,
		"richbrilliantlavender":          ColorRichBrilliantLavender,
		"bizarre":                        ColorBizarre,
		"englishwalnut":                  ColorEnglishWalnut,
		"gin":                            ColorGin,
		"melanie":                        ColorMelanie,
		"seagull":                        ColorSeagull,
		"terracotta":                     ColorTerraCotta,
		"chablis":                        ColorChablis,
		"daffodil":                       ColorDaffodil,
		"copperpenny":                    ColorCopperPenny,
		"ronchi":                         ColorRonchi,
		"caputmortuum":                   ColorCaputMortuum,
		"citrus":                         ColorCitrus,
		"charcoal":                       ColorCharcoal,
		"jasmine":                        ColorJasmine,
		"deepblush":                      ColorDeepBlush,
		"languidlavender":                ColorLanguidLavender,
		"kabul":                          ColorKabul,
		"offyellow":                      ColorOffYellow,
		"palmleaf":                       ColorPalmLeaf,
		"darkcoral":                      ColorDarkCoral,
		"ginfizz":                        ColorGinFizz,
		"cafenoir":                       ColorCafeNoir,
		"eclipse":                        ColorEclipse,
		"ultramarine":                    ColorUltramarine,
		"vividmulberry":                  ColorVividMulberry,
		"mabel":                          ColorMabel,
		"pakistangreen":                  ColorPakistanGreen,
		"fuzzywuzzybrown":                ColorFuzzyWuzzyBrown,
		"biscay":                         ColorBiscay,
		"cardingreen":                    ColorCardinGreen,
		"royalfuchsia":                   ColorRoyalFuchsia,
		"shakespeare":                    ColorShakespeare,
		"cannonblack":                    ColorCannonBlack,
		"cognac":                         ColorCognac,
		"cabsav":                         ColorCabSav,
		"folly":                          ColorFolly,
		"eastbay":                        ColorEastBay,
		"peru":                           ColorPeru,
		"taupegray":                      ColorTaupeGray,
		"celadongreen":                   ColorCeladonGreen,
		"deepoak":                        ColorDeepOak,
		"skeptic":                        ColorSkeptic,
		"tamarind":                       ColorTamarind,
		"mediumredviolet":                ColorMediumRedViolet,
		"mordantred":                     ColorMordantRed,
		"universityofcaliforniagold":     ColorUniversityOfCaliforniaGold,
		"bleudefrance":                   ColorBleuDeFrance,
		"kingfisherdaisy":                ColorKingfisherDaisy,
		"juniper":                        ColorJuniper,
		"panache":                        ColorPanache,
		"rosered":                        ColorRoseRed,
		"carnabytan":                     ColorCarnabyTan,
		"edward":                         ColorEdward,
		"snowymint":                      ColorSnowyMint,
		"ash":                            ColorAsh,
		"coldturkey":                     ColorColdTurkey,
		"huntergreen":                    ColorHunterGreen,
		"permanentgeraniumlake":          ColorPermanentGeraniumLake,
		"swamp":                          ColorSwamp,
		"bluegem":                        ColorBlueGem,
		"deepseagreen":                   ColorDeepSeaGreen,
		"ferngreen":                      ColorFernGreen,
		"lynch":                          ColorLynch,
		"plumppurple":                    ColorPlumpPurple,
		"clairvoyant":                    ColorClairvoyant,
		"copper":                         ColorCopper,
		"pantonepink":                    ColorPantonePink,
		"darkmediumgray":                 ColorDarkMediumGray,
		"dune":                           ColorDune,
		"fountainblue":                   ColorFountainBlue,
		"illuminatingemerald":            ColorIlluminatingEmerald,
		"kaitokegreen":                   ColorKaitokeGreen,
		"bananayellow":                   ColorBananaYellow,
		"endeavour":                      ColorEndeavour,
		"frenchraspberry":                ColorFrenchRaspberry,
		"frostee":                        ColorFrostee,
		"pantonemagenta":                 ColorPantoneMagenta,
		"stormdust":                      ColorStormDust,
		"sunglo":                         ColorSunglo,
		"dallas":                         ColorDallas,
		"deepmaroon":                     ColorDeepMaroon,
		"picasso":                        ColorPicasso,
		"turquoiseblue":                  ColorTurquoiseBlue,
		"goldenfizz":                     ColorGoldenFizz,
		"mauvelous":                      ColorMauvelous,
		"lightturquoise":                 ColorLightTurquoise,
		"quicksand":                      ColorQuicksand,
		"spanishgray":                    ColorSpanishGray,
		"yellow":                         ColorYellow,
		"aquahaze":                       ColorAquaHaze,
		"darkjunglegreen":                ColorDarkJungleGreen,
		"tabasco":                        ColorTabasco,
		"visvis":                         ColorVisVis,
		"vividburgundy":                  ColorVividBurgundy,
		"deepgreen":                      ColorDeepGreen,
		"lightcarminepink":               ColorLightCarminePink,
		"kournikova":                     ColorKournikova,
		"mediumslateblue":                ColorMediumSlateBlue,
		"concord":                        ColorConcord,
		"grannyapple":                    ColorGrannyApple,
		"rock":                           ColorRock,
		"towergray":                      ColorTowerGray,
		"meteor":                         ColorMeteor,
		"rybred":                         ColorRYBRed,
		"ripeplum":                       ColorRipePlum,
		"waiouru":                        ColorWaiouru,
		"fashionfuchsia":                 ColorFashionFuchsia,
		"jaguar":                         ColorJaguar,
		"neptune":                        ColorNeptune,
		"danube":                         ColorDanube,
		"lavenderpurple":                 ColorLavenderPurple,
		"pottersclay":                    ColorPottersClay,
		"cornfield":                      ColorCornField,
		"jacaranda":                      ColorJacaranda,
		"tacha":                          ColorTacha,
		"locust":                         ColorLocust,
		"seance":                         ColorSeance,
		"grannysmith":                    ColorGrannySmith,
		"vanilla":                        ColorVanilla,
		"bonjour":                        ColorBonJour,
		"coralred":                       ColorCoralRed,
		"jetstream":                      ColorJetStream,
		"rhino":                          ColorRhino,
		"upforestgreen":                  ColorUPForestGreen,
		"amber":                          ColorAmber,
		"clementine":                     ColorClementine,
		"plum":                           ColorPlum,
		"satinsheengold":                 ColorSatinSheenGold,
		"ecruwhite":                      ColorEcruWhite,
		"midnight":                       ColorMidnight,
		"persianorange":                  ColorPersianOrange,
		"sun":                            ColorSun,
		"turquoise":                      ColorTurquoise,
		"vermilion":                      ColorVermilion,
		"creambrulee":                    ColorCreamBrulee,
		"parchment":                      ColorParchment,
		"fungreen":                       ColorFunGreen,
		"goldentainoi":                   ColorGoldenTainoi,
		"lightmediumorchid":              ColorLightMediumOrchid,
		"luxorgold":                      ColorLuxorGold,
		"palesilver":                     ColorPaleSilver,
		"richmaroon":                     ColorRichMaroon,
		"bush":                           ColorBush,
		"frenchskyblue":                  ColorFrenchSkyBlue,
		"watercourse":                    ColorWatercourse,
		"sandal":                         ColorSandal,
		"spanishgreen":                   ColorSpanishGreen,
		"temptress":                      ColorTemptress,
		"oysterbay":                      ColorOysterBay,
		"paleleaf":                       ColorPaleLeaf,
		"lightbrilliantred":              ColorLightBrilliantRed,
		"metalliccopper":                 ColorMetallicCopper,
		"tomthumb":                       ColorTomThumb,
		"violetblue":                     ColorVioletBlue,
		"butterflybush":                  ColorButterflyBush,
		"carmine":                        ColorCarmine,
		"peach":                          ColorPeach,
		"fuchsiapurple":                  ColorFuchsiaPurple,
		"paarl":                          ColorPaarl,
		"jordyblue":                      ColorJordyBlue,
		"liverchestnut":                  ColorLiverChestnut,
		"oldlace":                        ColorOldLace,
		"pastelviolet":                   ColorPastelViolet,
		"ghostwhite":                     ColorGhostWhite,
		"greenwhite":                     ColorGreenWhite,
		"sage":                           ColorSage,
		"gladegreen":                     ColorGladeGreen,
		"purple":                         ColorPurple,
		"lividbrown":                     ColorLividBrown,
		"oceangreen":                     ColorOceanGreen,
		"paleredviolet":                  ColorPaleRedViolet,
		"seaserpent":                     ColorSeaSerpent,
		"alizarincrimson":                ColorAlizarinCrimson,
		"electricyellow":                 ColorElectricYellow,
		"lightslategray":                 ColorLightSlateGray,
		"mediumorchid":                   ColorMediumOrchid,
		"palemagentapink":                ColorPaleMagentaPink,
		"atoll":                          ColorAtoll,
		"fuelyellow":                     ColorFuelYellow,
		"vistawhite":                     ColorVistaWhite,
		"brownderby":                     ColorBrownDerby,
		"pistachio":                      ColorPistachio,
		"madang":                         ColorMadang,
		"orangewhite":                    ColorOrangeWhite,
		"richelectricblue":               ColorRichElectricBlue,
		"tartorange":                     ColorTartOrange,
		"vividorange":                    ColorVividOrange,
		"capepalliser":                   ColorCapePalliser,
		"gallery":                        ColorGallery,
		"chinaivory":                     ColorChinaIvory,
		"darkpastelpurple":               ColorDarkPastelPurple,
		"chlorophyllgreen":               ColorChlorophyllGreen,
		"fire":                           ColorFire,
		"kobicha":                        ColorKobicha,
		"blackrussian":                   ColorBlackRussian,
		"boysenberry":                    ColorBoysenberry,
		"lightdeeppink":                  ColorLightDeepPink,
		"portica":                        ColorPortica,
		"desert":                         ColorDesert,
		"cgred":                          ColorCGRed,
		"darkcyan":                       ColorDarkCyan,
		"macaroniandcheese":              ColorMacaroniAndCheese,
		"purpletaupe":                    ColorPurpleTaupe,
		"quinacridonemagenta":            ColorQuinacridoneMagenta,
		"sazerac":                        ColorSazerac,
		"spanishcrimson":                 ColorSpanishCrimson,
		"carnelian":                      ColorCarnelian,
		"ceruleanfrost":                  ColorCeruleanFrost,
		"mantle":                         ColorMantle,
		"venetianred":                    ColorVenetianRed,
		"anakiwa":                        ColorAnakiwa,
		"conch":                          ColorConch,
		"barnred":                        ColorBarnRed,
		"coralreef":                      ColorCoralReef,
		"darkvanilla":                    ColorDarkVanilla,
		"jambalaya":                      ColorJambalaya,
		"pampas":                         ColorPampas,
		"alabaster":                      ColorAlabaster,
		"australianmint":                 ColorAustralianMint,
		"fresheggplant":                  ColorFreshEggplant,
		"polishedpine":                   ColorPolishedPine,
		"purplemountainmajesty":          ColorPurpleMountainMajesty,
		"tangerineyellow":                ColorTangerineYellow,
		"thatchgreen":                    ColorThatchGreen,
		"alabamacrimson":                 ColorAlabamaCrimson,
		"azureishwhite":                  ColorAzureishWhite,
		"goldenbell":                     ColorGoldenBell,
		"graphite":                       ColorGraphite,
		"lemonglacier":                   ColorLemonGlacier,
		"mediumruby":                     ColorMediumRuby,
		"palelavender":                   ColorPaleLavender,
		"patina":                         ColorPatina,
		"alto":                           ColorAlto,
		"fallow":                         ColorFallow,
		"salem":                          ColorSalem,
		"sandrift":                       ColorSandrift,
		"greenmist":                      ColorGreenMist,
		"sugarplum":                      ColorSugarPlum,
		"cameopink":                      ColorCameoPink,
		"darkebony":                      ColorDarkEbony,
		"shilo":                          ColorShilo,
		"shuttlegray":                    ColorShuttleGray,
		"mysin":                          ColorMySin,
		"palepink":                       ColorPalePink,
		"veronica":                       ColorVeronica,
		"abbey":                          ColorAbbey,
		"greensheen":                     ColorGreenSheen,
		"paleplum":                       ColorPalePlum,
		"spray":                          ColorSpray,
		"valentino":                      ColorValentino,
		"cyanblueazure":                  ColorCyanBlueAzure,
		"frost":                          ColorFrost,
		"palecopper":                     ColorPaleCopper,
		"regentstblue":                   ColorRegentStBlue,
		"coconutcream":                   ColorCoconutCream,
		"moonmist":                       ColorMoonMist,
		"charlestongreen":                ColorCharlestonGreen,
		"scarlet":                        ColorScarlet,
		"amour":                          ColorAmour,
		"blacksqueeze":                   ColorBlackSqueeze,
		"treehouse":                      ColorTreehouse,
		"darkturquoise":                  ColorDarkTurquoise,
		"millbrook":                      ColorMillbrook,
		"darkpastelgreen":                ColorDarkPastelGreen,
		"frenchplum":                     ColorFrenchPlum,
		"frostedmint":                    ColorFrostedMint,
		"greenyellow":                    ColorGreenYellow,
		"junglegreen":                    ColorJungleGreen,
		"mistgray":                       ColorMistGray,
		"ceramic":                        ColorCeramic,
		"confetti":                       ColorConfetti,
		"pictorialcarmine":               ColorPictorialCarmine,
		"whiskey":                        ColorWhiskey,
		"greenleaf":                      ColorGreenLeaf,
		"vanillaice":                     ColorVanillaIce,
		"bigstone":                       ColorBigStone,
		"chestnut":                       ColorChestnut,
		"spanishviolet":                  ColorSpanishViolet,
		"jackobean":                      ColorJackoBean,
		"pearlypurple":                   ColorPearlyPurple,
		"viridian":                       ColorViridian,
		"malachite":                      ColorMalachite,
		"mexicanred":                     ColorMexicanRed,
		"pizazz":                         ColorPizazz,
		"sealbrown":                      ColorSealBrown,
		"antiquefuchsia":                 ColorAntiqueFuchsia,
		"curiousblue":                    ColorCuriousBlue,
		"pearlmysticturquoise":           ColorPearlMysticTurquoise,
		"cgblue":                         ColorCGBlue,
		"dulllavender":                   ColorDullLavender,
		"waxflower":                      ColorWaxFlower,
		"ruby":                           ColorRuby,
		"ultramarineblue":                ColorUltramarineBlue,
		"oldsilver":                      ColorOldSilver,
		"silverchalice":                  ColorSilverChalice,
		"barbiepink":                     ColorBarbiePink,
		"greenspring":                    ColorGreenSpring,
		"ghost":                          ColorGhost,
		"greenwaterloo":                  ColorGreenWaterloo,
		"shark":                          ColorShark,
		"zomp":                           ColorZomp,
		"asphalt":                        ColorAsphalt,
		"brightgray":                     ColorBrightGray,
		"duststorm":                      ColorDustStorm,
		"eunry":                          ColorEunry,
		"pastelpurple":                   ColorPastelPurple,
		"victoria":                       ColorVictoria,
		"bordeaux":                       ColorBordeaux,
		"cybergrape":                     ColorCyberGrape,
		"orangesoda":                     ColorOrangeSoda,
		"tequila":                        ColorTequila,
		"darktan":                        ColorDarkTan,
		"napiergreen":                    ColorNapierGreen,
		"frenchmauve":                    ColorFrenchMauve,
		"logcabin":                       ColorLogCabin,
		"pavlova":                        ColorPavlova,
		"algaegreen":                     ColorAlgaeGreen,
		"fielddrab":                      ColorFieldDrab,
		"zinnwaldite":                    ColorZinnwaldite,
		"botticelli":                     ColorBotticelli,
		"goldensand":                     ColorGoldenSand,
		"screamingreen":                  ColorScreaminGreen,
		"smokyblack":                     ColorSmokyBlack,
		"tawnyport":                      ColorTawnyPort,
		"bilobaflower":                   ColorBilobaFlower,
		"carissma":                       ColorCarissma,
		"cardinalpink":                   ColorCardinalPink,
		"darkorange":                     ColorDarkOrange,
		"doublespanishwhite":             ColorDoubleSpanishWhite,
		"oysterpink":                     ColorOysterPink,
		"padua":                          ColorPadua,
		"beautybush":                     ColorBeautyBush,
		"camouflage":                     ColorCamouflage,
		"eternity":                       ColorEternity,
		"grayasparagus":                  ColorGrayAsparagus,
		"heath":                          ColorHeath,
		"magnolia":                       ColorMagnolia,
		"rubinered":                      ColorRubineRed,
		"slateblue":                      ColorSlateBlue,
		"androidgreen":                   ColorAndroidGreen,
		"brightnavyblue":                 ColorBrightNavyBlue,
		"valhalla":                       ColorValhalla,
		"metallicgold":                   ColorMetallicGold,
		"seapink":                        ColorSeaPink,
		"shamrock":                       ColorShamrock,
		"springwood":                     ColorSpringWood,
		"appleblossom":                   ColorAppleBlossom,
		"bananamania":                    ColorBananaMania,
		"mountainmeadow":                 ColorMountainMeadow,
		"newyorkpink":                    ColorNewYorkPink,
		"olivetone":                      ColorOlivetone,
		"pearlbush":                      ColorPearlBush,
		"sandybeach":                     ColorSandyBeach,
		"espresso":                       ColorEspresso,
		"fireenginered":                  ColorFireEngineRed,
		"feldgrau":                       ColorFeldgrau,
		"mystic":                         ColorMystic,
		"rosevale":                       ColorRoseVale,
		"sacramentostategreen":           ColorSacramentoStateGreen,
		"waterloo":                       ColorWaterloo,
		"wintergreendream":               ColorWintergreenDream,
		"bunting":                        ColorBunting,
		"deeptuscanred":                  ColorDeepTuscanRed,
		"volt":                           ColorVolt,
		"goldtips":                       ColorGoldTips,
		"java":                           ColorJava,
		"edgewater":                      ColorEdgewater,
		"starkwhite":                     ColorStarkWhite,
		"armadillo":                      ColorArmadillo,
		"darkslateblue":                  ColorDarkSlateBlue,
		"spicypink":                      ColorSpicyPink,
		"darkspringgreen":                ColorDarkSpringGreen,
		"greenvogue":                     ColorGreenVogue,
		"lightcobaltblue":                ColorLightCobaltBlue,
		"silverlakeblue":                 ColorSilverLakeBlue,
		"dawnpink":                       ColorDawnPink,
		"greenkelp":                      ColorGreenKelp,
		"peachcream":                     ColorPeachCream,
		"brightcerulean":                 ColorBrightCerulean,
		"minttulip":                      ColorMintTulip,
		"sunny":                          ColorSunny,
		"blackrose":                      ColorBlackRose,
		"cedarwoodfinish":                ColorCedarWoodFinish,
		"hokeypokey":                     ColorHokeyPokey,
		"hurricane":                      ColorHurricane,
		"islandspice":                    ColorIslandSpice,
		"blueromance":                    ColorBlueRomance,
		"cosmiclatte":                    ColorCosmicLatte,
		"darkgunmetal":                   ColorDarkGunmetal,
		"japanesecarmine":                ColorJapaneseCarmine,
		"lemoncurry":                     ColorLemonCurry,
		"lightcoral":                     ColorLightCoral,
		"sauvignon":                      ColorSauvignon,
		"springsun":                      ColorSpringSun,
		"balticsea":                      ColorBalticSea,
		"cannonpink":                     ColorCannonPink,
		"gothic":                         ColorGothic,
		"palatinatepurple":               ColorPalatinatePurple,
		"shadowblue":                     ColorShadowBlue,
		"bostonblue":                     ColorBostonBlue,
		"elfgreen":                       ColorElfGreen,
		"prussianblue":                   ColorPrussianBlue,
		"highland":                       ColorHighland,
		"melon":                          ColorMelon,
		"bridalheath":                    ColorBridalHeath,
		"palegold":                       ColorPaleGold,
		"jellybean":                      ColorJellyBean,
		"mummystomb":                     ColorMummysTomb,
		"niagara":                        ColorNiagara,
		"chinarose":                      ColorChinaRose,
		"eggsour":                        ColorEggSour,
		"palespringbud":                  ColorPaleSpringBud,
		"tobaccobrown":                   ColorTobaccoBrown,
		"bittersweetshimmer":             ColorBittersweetShimmer,
		"glitter":                        ColorGlitter,
		"sweetpink":                      ColorSweetPink,
		"turquoisegreen":                 ColorTurquoiseGreen,
		"limeade":                        ColorLimeade,
		"olivedrabseven":                 ColorOliveDrabSeven,
		"graychateau":                    ColorGrayChateau,
		"horsesneck":                     ColorHorsesNeck,
		"lightfuchsiapink":               ColorLightFuchsiaPink,
		"metallicseaweed":                ColorMetallicSeaweed,
		"aluminium":                      ColorAluminium,
		"byzantium":                      ColorByzantium,
		"honeysuckle":                    ColorHoneysuckle,
		"orangered":                      ColorOrangeRed,
		"westar":                         ColorWestar,
		"woodbark":                       ColorWoodBark,
		"amaranth":                       ColorAmaranth,
		"bone":                           ColorBone,
		"dustygray":                      ColorDustyGray,
		"indianyellow":                   ColorIndianYellow,
		"mintjulep":                      ColorMintJulep,
		"palecarmine":                    ColorPaleCarmine,
		"sugarcane":                      ColorSugarCane,
		"thunderbird":                    ColorThunderbird,
		"brightlilac":                    ColorBrightLilac,
		"debianred":                      ColorDebianRed,
		"tuliptree":                      ColorTulipTree,
		"fawn":                           ColorFawn,
		"frangipani":                     ColorFrangipani,
		"jacarta":                        ColorJacarta,
		"pinecone":                       ColorPineCone,
		"russianviolet":                  ColorRussianViolet,
		"bayleaf":                        ColorBayLeaf,
		"champagne":                      ColorChampagne,
		"persiangreen":                   ColorPersianGreen,
		"raspberrypink":                  ColorRaspberryPink,
		"trendygreen":                    ColorTrendyGreen,
		"cinnabar":                       ColorCinnabar,
		"deepspacesparkle":               ColorDeepSpaceSparkle,
		"rybyellow":                      ColorRYBYellow,
		"sepiaskin":                      ColorSepiaSkin,
		"winterhazel":                    ColorWinterHazel,
		"darkmossgreen":                  ColorDarkMossGreen,
		"granitegray":                    ColorGraniteGray,
		"merlin":                         ColorMerlin,
		"pastelpink":                     ColorPastelPink,
		"redberry":                       ColorRedBerry,
		"solitude":                       ColorSolitude,
		"timberwolf":                     ColorTimberwolf,
		"usafablue":                      ColorUSAFABlue,
		"crowshead":                      ColorCrowshead,
		"imperialblue":                   ColorImperialBlue,
		"mauvetaupe":                     ColorMauveTaupe,
		"meatbrown":                      ColorMeatBrown,
		"regentgray":                     ColorRegentGray,
		"royalblue":                      ColorRoyalBlue,
		"amaranthred":                    ColorAmaranthRed,
		"ao":                             ColorAo,
		"rawsienna":                      ColorRawSienna,
		"redwood":                        ColorRedwood,
		"satinlinen":                     ColorSatinLinen,
		"clayash":                        ColorClayAsh,
		"parsley":                        ColorParsley,
		"nightshadz":                     ColorNightShadz,
		"perutan":                        ColorPeruTan,
		"suvagray":                       ColorSuvaGray,
		"hoki":                           ColorHoki,
		"mughalgreen":                    ColorMughalGreen,
		"redbeech":                       ColorRedBeech,
		"spanishblue":                    ColorSpanishBlue,
		"verylightblue":                  ColorVeryLightBlue,
		"finn":                           ColorFinn,
		"greencyan":                      ColorGreenCyan,
		"darkchestnut":                   ColorDarkChestnut,
		"gray":                           ColorGray,
		"gumbo":                          ColorGumbo,
		"inchworm":                       ColorInchWorm,
		"seamist":                        ColorSeaMist,
		"blackrock":                      ColorBlackRock,
		"cinereous":                      ColorCinereous,
		"mediumskyblue":                  ColorMediumSkyBlue,
		"frenchrose":                     ColorFrenchRose,
		"honeyflower":                    ColorHoneyFlower,
		"spacecadet":                     ColorSpaceCadet,
		"wildorchid":                     ColorWildOrchid,
		"amulet":                         ColorAmulet,
		"chelseacucumber":                ColorChelseaCucumber,
		"holly":                          ColorHolly,
		"magentahaze":                    ColorMagentaHaze,
		"matisse":                        ColorMatisse,
		"onion":                          ColorOnion,
		"palebrown":                      ColorPaleBrown,
		"tuscantan":                      ColorTuscanTan,
		"bermudagray":                    ColorBermudaGray,
		"grenadier":                      ColorGrenadier,
		"verypaleyellow":                 ColorVeryPaleYellow,
		"deeppink":                       ColorDeepPink,
		"donkeybrown":                    ColorDonkeyBrown,
		"coolgrey":                       ColorCoolGrey,
		"crayolared":                     ColorCrayolaRed,
		"darkterracotta":                 ColorDarkTerraCotta,
		"hitpink":                        ColorHitPink,
		"reddamask":                      ColorRedDamask,
		"ruber":                          ColorRuber,
		"birdflower":                     ColorBirdFlower,
		"bubblegum":                      ColorBubbleGum,
		"weldonblue":                     ColorWeldonBlue,
		"seabuckthorn":                   ColorSeaBuckthorn,
		"sunflower":                      ColorSunflower,
		"fuzzywuzzy":                     ColorFuzzyWuzzy,
		"onahau":                         ColorOnahau,
		"cerise":                         ColorCerise,
		"melrose":                        ColorMelrose,
		"purpleheart":                    ColorPurpleHeart,
		"aero":                           ColorAero,
		"astral":                         ColorAstral,
		"gogreen":                        ColorGOGreen,
		"monalisa":                       ColorMonaLisa,
		"olivedrab":                      ColorOliveDrab,
		"pelorous":                       ColorPelorous,
		"casper":                         ColorCasper,
		"chatelle":                       ColorChatelle,
		"rubyred":                        ColorRubyRed,
		"snuff":                          ColorSnuff,
		"softamber":                      ColorSoftAmber,
		"tearose":                        ColorTeaRose,
		"boulder":                        ColorBoulder,
		"frenchpuce":                     ColorFrenchPuce,
		"matterhorn":                     ColorMatterhorn,
		"sunset":                         ColorSunset,
		"cabbagepont":                    ColorCabbagePont,
		"icterine":                       ColorIcterine,
		"frenchwine":                     ColorFrenchWine,
		"londonhue":                      ColorLondonHue,
		"pink":                           ColorPink,
		"princessperfume":                ColorPrincessPerfume,
		"selectiveyellow":                ColorSelectiveYellow,
		"tidal":                          ColorTidal,
		"charm":                          ColorCharm,
		"elsalva":                        ColorElSalva,
		"wewak":                          ColorWewak,
		"electricpurple":                 ColorElectricPurple,
		"matrix":                         ColorMatrix,
		"brandeisblue":                   ColorBrandeisBlue,
		"gossamer":                       ColorGossamer,
		"cyberyellow":                    ColorCyberYellow,
		"easternblue":                    ColorEasternBlue,
		"nonphotoblue":                   ColorNonPhotoBlue,
		"chromeyellow":                   ColorChromeYellow,
		"cranberry":                      ColorCranberry,
		"iceberg":                        ColorIceberg,
		"munsellblue":                    ColorMunsellBlue,
		"salmonpink":                     ColorSalmonPink,
		"tealgreen":                      ColorTealGreen,
		"bluedianne":                     ColorBlueDianne,
		"deepsapphire":                   ColorDeepSapphire,
		"dogs":                           ColorDogs,
		"frenchpink":                     ColorFrenchPink,
		"pigpink":                        ColorPigPink,
		"coppercanyon":                   ColorCopperCanyon,
		"countygreen":                    ColorCountyGreen,
		"kangaroo":                       ColorKangaroo,
		"springrain":                     ColorSpringRain,
		"darkburgundy":                   ColorDarkBurgundy,
		"emerald":                        ColorEmerald,
		"chathamsblue":                   ColorChathamsBlue,
		"grape":                          ColorGrape,
		"lochmara":                       ColorLochmara,
		"pumpkin":                        ColorPumpkin,
		"brightyellow":                   ColorBrightYellow,
		"chamoisee":                      ColorChamoisee,
		"texas":                          ColorTexas,
		"vividredtangelo":                ColorVividRedTangelo,
		"brinkpink":                      ColorBrinkPink,
		"steelpink":                      ColorSteelPink,
		"dingley":                        ColorDingley,
		"hacienda":                       ColorHacienda,
		"cruise":                         ColorCruise,
		"deepgreencyanturquoise":         ColorDeepGreenCyanTurquoise,
		"viridiangreen":                  ColorViridianGreen,
		"bronze":                         ColorBronze,
		"tangopink":                      ColorTangoPink,
		"kombugreen":                     ColorKombuGreen,
		"newcar":                         ColorNewCar,
		"strawberry":                     ColorStrawberry,
		"arylideyellow":                  ColorArylideYellow,
		"deepfir":                        ColorDeepFir,
		"racinggreen":                    ColorRacingGreen,
		"rosegold":                       ColorRoseGold,
		"schist":                         ColorSchist,
		"swampgreen":                     ColorSwampGreen,
		"swisscoffee":                    ColorSwissCoffee,
		"tomato":                         ColorTomato,
		"davysgrey":                      ColorDavysGrey,
		"feijoa":                         ColorFeijoa,
		"grandis":                        ColorGrandis,
		"mediumspringgreen":              ColorMediumSpringGreen,
		"darkpastelred":                  ColorDarkPastelRed,
		"gamboge":                        ColorGamboge,
		"shocking":                       ColorShocking,
		"tonyspink":                      ColorTonysPink,
		"whiterock":                      ColorWhiteRock,
		"chenin":                         ColorChenin,
		"colonialwhite":                  ColorColonialWhite,
		"lilac":                          ColorLilac,
		"moonraker":                      ColorMoonRaker,
		"tuftbush":                       ColorTuftBush,
		"vividgamboge":                   ColorVividGamboge,
		"craterbrown":                    ColorCraterBrown,
		"finch":                          ColorFinch,
		"hemp":                           ColorHemp,
		"bandicoot":                      ColorBandicoot,
		"buccaneer":                      ColorBuccaneer,
		"cedarchest":                     ColorCedarChest,
		"coldpurple":                     ColorColdPurple,
		"palecerulean":                   ColorPaleCerulean,
		"totempole":                      ColorTotemPole,
		"arrowtown":                      ColorArrowtown,
		"brunswickgreen":                 ColorBrunswickGreen,
		"oldlavender":                    ColorOldLavender,
		"orangeroughy":                   ColorOrangeRoughy,
		"deepbronze":                     ColorDeepBronze,
		"lemonginger":                    ColorLemonGinger,
		"froly":                          ColorFroly,
		"psychedelicpurple":              ColorPsychedelicPurple,
		"sweetcorn":                      ColorSweetCorn,
		"electriccrimson":                ColorElectricCrimson,
		"finlandia":                      ColorFinlandia,
		"whitelinen":                     ColorWhiteLinen,
		"diserria":                       ColorDiSerria,
		"tuftsblue":                      ColorTuftsBlue,
		"pueblo":                         ColorPueblo,
		"robroy":                         ColorRobRoy,
		"shimmeringblush":                ColorShimmeringBlush,
		"teagreen":                       ColorTeaGreen,
		"ceruleanblue":                   ColorCeruleanBlue,
		"oasis":                          ColorOasis,
		"cork":                           ColorCork,
		"maximumyellow":                  ColorMaximumYellow,
		"pear":                           ColorPear,
		"chicago":                        ColorChicago,
		"heather":                        ColorHeather,
		"pantoneblue":                    ColorPantoneBlue,
		"phthalogreen":                   ColorPhthaloGreen,
		"thistlegreen":                   ColorThistleGreen,
		"lima":                           ColorLima,
		"palmgreen":                      ColorPalmGreen,
		"mistymoss":                      ColorMistyMoss,
		"pickledbean":                    ColorPickledBean,
		"barberry":                       ColorBarberry,
		"beaublue":                       ColorBeauBlue,
		"wisppink":                       ColorWispPink,
		"gunpowder":                      ColorGunPowder,
		"islamicgreen":                   ColorIslamicGreen,
		"rajah":                          ColorRajah,
		"sapling":                        ColorSapling,
		"cadmiumyellow":                  ColorCadmiumYellow,
		"jacksonspurple":                 ColorJacksonsPurple,
		"northtexasgreen":                ColorNorthTexasGreen,
		"romancoffee":                    ColorRomanCoffee,
		"spanishorange":                  ColorSpanishOrange,
		"deeptaupe":                      ColorDeepTaupe,
		"goldenrod":                      ColorGoldenrod,
		"firebrick":                      ColorFirebrick,
		"shinglefawn":                    ColorShingleFawn,
		"yelloworange":                   ColorYellowOrange,
		"careyspink":                     ColorCareysPink,
		"columbiablue":                   ColorColumbiaBlue,
		"flamingo":                       ColorFlamingo,
		"indigo":                         ColorIndigo,
		"portgore":                       ColorPortGore,
		"vancleef":                       ColorVanCleef,
		"antiquebronze":                  ColorAntiqueBronze,
		"fandango":                       ColorFandango,
		"milanored":                      ColorMilanoRed,
		"richlavender":                   ColorRichLavender,
		"clinker":                        ColorClinker,
		"sahara":                         ColorSahara,
		"bigfootfeet":                    ColorBigFootFeet,
		"browntumbleweed":                ColorBrownTumbleweed,
		"broom":                          ColorBroom,
		"concrete":                       ColorConcrete,
		"geyser":                         ColorGeyser,
		"jet":                            ColorJet,
		"mindaro":                        ColorMindaro,
		"purplepizzazz":                  ColorPurplePizzazz,
		"alloyorange":                    ColorAlloyOrange,
		"blackforest":                    ColorBlackForest,
		"whiteice":                       ColorWhiteIce,
		"justright":                      ColorJustRight,
		"lust":                           ColorLust,
		"maroonoak":                      ColorMaroonOak,
		"oldbrick":                       ColorOldBrick,
		"blizzardblue":                   ColorBlizzardBlue,
		"indianred":                      ColorIndianRed,
		"bleachedcedar":                  ColorBleachedCedar,
		"brandy":                         ColorBrandy,
		"pinkflamingo":                   ColorPinkFlamingo,
		"sttropaz":                       ColorStTropaz,
		"cottonseed":                     ColorCottonSeed,
		"karry":                          ColorKarry,
		"rustynail":                      ColorRustyNail,
		"seagreen":                       ColorSeaGreen,
		"bostonuniversityred":            ColorBostonUniversityRed,
		"corn":                           ColorCorn,
		"yaleblue":                       ColorYaleBlue,
		"regalblue":                      ColorRegalBlue,
		"uablue":                         ColorUABlue,
		"electricindigo":                 ColorElectricIndigo,
		"kenyancopper":                   ColorKenyanCopper,
		"lighthotpink":                   ColorLightHotPink,
		"resolutionblue":                 ColorResolutionBlue,
		"spicymustard":                   ColorSpicyMustard,
		"vividlimegreen":                 ColorVividLimeGreen,
		"blackcurrant":                   ColorBlackcurrant,
		"dorado":                         ColorDorado,
		"rose":                           ColorRose,
		"lightmossgreen":                 ColorLightMossGreen,
		"nandor":                         ColorNandor,
		"vividcerise":                    ColorVividCerise,
		"chamois":                        ColorChamois,
		"comet":                          ColorComet,
		"palatinateblue":                 ColorPalatinateBlue,
		"red":                            ColorRed,
		"redorange":                      ColorRedOrange,
		"dingydungeon":                   ColorDingyDungeon,
		"marigoldyellow":                 ColorMarigoldYellow,
		"mongoose":                       ColorMongoose,
		"richcarmine":                    ColorRichCarmine,
		"majorelleblue":                  ColorMajorelleBlue,
		"mustard":                        ColorMustard,
		"pigmentgreen":                   ColorPigmentGreen,
		"rhythm":                         ColorRhythm,
		"texasrose":                      ColorTexasRose,
		"tide":                           ColorTide,
		"donjuan":                        ColorDonJuan,
		"englishlavender":                ColorEnglishLavender,
		"dolly":                          ColorDolly,
		"peachpuff":                      ColorPeachPuff,
		"bulgarianrose":                  ColorBulgarianRose,
		"darkviolet":                     ColorDarkViolet,
		"petiteorchid":                   ColorPetiteOrchid,
		"sunglow":                        ColorSunglow,
		"tutu":                           ColorTutu,
		"artichoke":                      ColorArtichoke,
		"milan":                          ColorMilan,
		"hotpink":                        ColorHotPink,
		"britishracinggreen":             ColorBritishRacingGreen,
		"deepsaffron":                    ColorDeepSaffron,
		"cosmos":                         ColorCosmos,
		"romantic":                       ColorRomantic,
		"glossygrape":                    ColorGlossyGrape,
		"goldenyellow":                   ColorGoldenYellow,
		"harlequingreen":                 ColorHarlequinGreen,
		"polar":                          ColorPolar,
		"potpourri":                      ColorPotPourri,
		"tigerseye":                      ColorTigersEye,
		"corduroy":                       ColorCorduroy,
		"darklavender":                   ColorDarkLavender,
		"limedspruce":                    ColorLimedSpruce,
		"lipstick":                       ColorLipstick,
		"rybviolet":                      ColorRYBViolet,
		"rebel":                          ColorRebel,
		"goldenpoppy":                    ColorGoldenPoppy,
		"greenhaze":                      ColorGreenHaze,
		"heatwave":                       ColorHeatWave,
		"porsche":                        ColorPorsche,
		"persianred":                     ColorPersianRed,
		"portafino":                      ColorPortafino,
		"regalia":                        ColorRegalia,
		"wildwillow":                     ColorWildWillow,
		"bakermillerpink":                ColorBakerMillerPink,
		"byzantine":                      ColorByzantine,
		"dovegray":                       ColorDoveGray,
		"mountbattenpink":                ColorMountbattenPink,
		"shadylady":                      ColorShadyLady,
		"surfiegreen":                    ColorSurfieGreen,
		"vividtangerine":                 ColorVividTangerine,
		"ballblue":                       ColorBallBlue,
		"copperrust":                     ColorCopperRust,
		"darkbyzantium":                  ColorDarkByzantium,
		"siam":                           ColorSiam,
		"waterspout":                     ColorWaterspout,
		"blackpearl":                     ColorBlackPearl,
		"coolblack":                      ColorCoolBlack,
		"lava":                           ColorLava,
		"lightgreen":                     ColorLightGreen,
		"roastcoffee":                    ColorRoastCoffee,
		"deepcarrotorange":               ColorDeepCarrotOrange,
		"fringyflower":                   ColorFringyFlower,
		"waikawagray":                    ColorWaikawaGray,
		"flint":                          ColorFlint,
		"maroon":                         ColorMaroon,
		"schoolbusyellow":                ColorSchoolBusYellow,
		"yankeesblue":                    ColorYankeesBlue,
		"belgion":                        ColorBelgion,
		"california":                     ColorCalifornia,
		"crimson":                        ColorCrimson,
		"englishholly":                   ColorEnglishHolly,
		"madras":                         ColorMadras,
		"carla":                          ColorCarla,
		"cordovan":                       ColorCordovan,
		"elephant":                       ColorElephant,
		"mossgreen":                      ColorMossGreen,
		"submarine":                      ColorSubmarine,
		"swirl":                          ColorSwirl,
		"tiamaria":                       ColorTiaMaria,
		"travertine":                     ColorTravertine,
		"bdazzledblue":                   ColorBdazzledBlue,
		"buttermilk":                     ColorButtermilk,
		"balihai":                        ColorBaliHai,
		"carolinablue":                   ColorCarolinaBlue,
		"pinetree":                       ColorPineTree,
		"tuscanred":                      ColorTuscanRed,
		"acadia":                         ColorAcadia,
		"allports":                       ColorAllports,
		"stormcloud":                     ColorStormcloud,
		"wildrice":                       ColorWildRice,
		"albescentwhite":                 ColorAlbescentWhite,
		"aquadeep":                       ColorAquaDeep,
		"marigold":                       ColorMarigold,
		"onyx":                           ColorOnyx,
		"romansilver":                    ColorRomanSilver,
		"festival":                       ColorFestival,
		"flame":                          ColorFlame,
		"blueviolet":                     ColorBlueViolet,
		"orchidpink":                     ColorOrchidPink,
		"daisybush":                      ColorDaisyBush,
		"deeplilac":                      ColorDeepLilac,
		"electricblue":                   ColorElectricBlue,
		"mikadoyellow":                   ColorMikadoYellow,
		"queenpink":                      ColorQueenPink,
		"sandstone":                      ColorSandstone,
		"blazeorange":                    ColorBlazeOrange,
		"carnation":                      ColorCarnation,
		"underagepink":                   ColorUnderagePink,
		"cola":                           ColorCola,
		"deepviolet":                     ColorDeepViolet,
		"genoa":                          ColorGenoa,
		"grizzly":                        ColorGrizzly,
		"magenta":                        ColorMagenta,
		"reefgold":                       ColorReefGold,
		"arcticlime":                     ColorArcticLime,
		"brownpod":                       ColorBrownPod,
		"vividviolet":                    ColorVividViolet,
		"darkbrowntangelo":               ColorDarkBrownTangelo,
		"deluge":                         ColorDeluge,
		"claret":                         ColorClaret,
		"coquelicot":                     ColorCoquelicot,
		"violentviolet":                  ColorViolentViolet,
		"woodybrown":                     ColorWoodyBrown,
		"arapawa":                        ColorArapawa,
		"darkseagreen":                   ColorDarkSeaGreen,
		"casal":                          ColorCasal,
		"eveningsea":                     ColorEveningSea,
		"paleprim":                       ColorPalePrim,
		"bismark":                        ColorBismark,
		"caper":                          ColorCaper,
		"ecru":                           ColorEcru,
		"goldenglow":                     ColorGoldenGlow,
		"mellowapricot":                  ColorMellowApricot,
		"muddywaters":                    ColorMuddyWaters,
		"sasquatchsocks":                 ColorSasquatchSocks,
		"soapstone":                      ColorSoapstone,
		"cinder":                         ColorCinder,
		"creole":                         ColorCreole,
		"cadmiumred":                     ColorCadmiumRed,
		"darkmagenta":                    ColorDarkMagenta,
		"oceanblue":                      ColorOceanBlue,
		"prairiesand":                    ColorPrairieSand,
		"antiqueruby":                    ColorAntiqueRuby,
		"brilliantrose":                  ColorBrilliantRose,
		"hippiepink":                     ColorHippiePink,
		"lightgray":                      ColorLightGray,
		"loulou":                         ColorLoulou,
		"masala":                         ColorMasala,
		"parism":                         ColorParisM,
		"auburn":                         ColorAuburn,
		"bullshot":                       ColorBullShot,
		"processmagenta":                 ColorProcessMagenta,
		"mediumelectricblue":             ColorMediumElectricBlue,
		"pinklady":                       ColorPinkLady,
		"goldfusion":                     ColorGoldFusion,
		"nepal":                          ColorNepal,
		"paco":                           ColorPaco,
		"wildsand":                       ColorWildSand,
		"aztec":                          ColorAztec,
		"fern":                           ColorFern,
		"tropicalrainforest":             ColorTropicalRainForest,
		"goldendream":                    ColorGoldenDream,
		"purplenavy":                     ColorPurpleNavy,
		"silvertree":                     ColorSilverTree,
		"sorrellbrown":                   ColorSorrellBrown,
		"greensmoke":                     ColorGreenSmoke,
		"khaki":                          ColorKhaki,
		"frenchblue":                     ColorFrenchBlue,
		"zanah":                          ColorZanah,
		"minsk":                          ColorMinsk,
		"darkraspberry":                  ColorDarkRaspberry,
		"deeplemon":                      ColorDeepLemon,
		"jasper":                         ColorJasper,
		"mediumseagreen":                 ColorMediumSeaGreen,
		"deco":                           ColorDeco,
		"hintofgreen":                    ColorHintofGreen,
		"organ":                          ColorOrgan,
		"pacificblue":                    ColorPacificBlue,
		"ecstasy":                        ColorEcstasy,
		"hemlock":                        ColorHemlock,
		"buff":                           ColorBuff,
		"ceil":                           ColorCeil,
		"crete":                          ColorCrete,
		"harlequin":                      ColorHarlequin,
		"mandarin":                       ColorMandarin,
		"vividmalachite":                 ColorVividMalachite,
		"athensgray":                     ColorAthensGray,
		"boogerbuster":                   ColorBoogerBuster,
		"dimgray":                        ColorDimGray,
		"portlandorange":                 ColorPortlandOrange,
		"russet":                         ColorRusset,
		"saffron":                        ColorSaffron,
		"amaranthpurple":                 ColorAmaranthPurple,
		"cello":                          ColorCello,
		"smoke":                          ColorSmoke,
		"warmblack":                      ColorWarmBlack,
		"lightcrimson":                   ColorLightCrimson,
		"muesli":                         ColorMuesli,
		"palesky":                        ColorPaleSky,
		"ruddybrown":                     ColorRuddyBrown,
		"ashgrey":                        ColorAshGrey,
		"mediumjunglegreen":              ColorMediumJungleGreen,
		"manatee":                        ColorManatee,
		"crail":                          ColorCrail,
		"lavendergray":                   ColorLavenderGray,
		"cottoncandy":                    ColorCottonCandy,
		"pastelgray":                     ColorPastelGray,
		"powderblue":                     ColorPowderBlue,
		"ricecake":                       ColorRiceCake,
		"steelblue":                      ColorSteelBlue,
		"bluediamond":                    ColorBlueDiamond,
		"carouselpink":                   ColorCarouselPink,
		"zinnwalditebrown":               ColorZinnwalditeBrown,
		"cardinal":                       ColorCardinal,
		"watusi":                         ColorWatusi,
		"pigmentblue":                    ColorPigmentBlue,
		"golden":                         ColorGolden,
		"pearllusta":                     ColorPearlLusta,
		"disco":                          ColorDisco,
		"luckypoint":                     ColorLuckyPoint,
		"chiffon":                        ColorChiffon,
		"darkbluegray":                   ColorDarkBlueGray,
		"blackcoral":                     ColorBlackCoral,
		"olivine":                        ColorOlivine,
		"saeeceamber":                    ColorSAEECEAmber,
		"tallow":                         ColorTallow,
		"arsenic":                        ColorArsenic,
		"astra":                          ColorAstra,
		"taupe":                          ColorTaupe,
		"vividvermilion":                 ColorVividVermilion,
		"zest":                           ColorZest,
		"cerisepink":                     ColorCerisePink,
		"charmpink":                      ColorCharmPink,
		"blueyonder":                     ColorBlueYonder,
		"rybgreen":                       ColorRYBGreen,
		"carminered":                     ColorCarmineRed,
		"gordonsgreen":                   ColorGordonsGreen,
		"oxfordblue":                     ColorOxfordBlue,
		"treepoppy":                      ColorTreePoppy,
		"viola":                          ColorViola,
		"beige":                          ColorBeige,
		"birch":                          ColorBirch,
		"deepchestnut":                   ColorDeepChestnut,
		"internationalorange":            ColorInternationalOrange,
		"lumber":                         ColorLumber,
		"ncsblue":                        ColorNCSBlue,
		"rybblue":                        ColorRYBBlue,
		"applegreen":                     ColorAppleGreen,
		"bluehaze":                       ColorBlueHaze,
		"munsellyellow":                  ColorMunsellYellow,
		"pullmangreen":                   ColorPullmanGreen,
		"tangelo":                        ColorTangelo,
		"x11darkgreen":                   ColorX11DarkGreen,
		"deepforestgreen":                ColorDeepForestGreen,
		"irishcoffee":                    ColorIrishCoffee,
		"heatheredgray":                  ColorHeatheredGray,
		"summergreen":                    ColorSummerGreen,
		"white":                          ColorWhite,
		"cioccolato":                     ColorCioccolato,
		"hawkesblue":                     ColorHawkesBlue,
		"greenlizard":                    ColorGreenLizard,
		"lavenderblush":                  ColorLavenderBlush,
		"mahogany":                       ColorMahogany,
		"ripelemon":                      ColorRipeLemon,
		"roseebony":                      ColorRoseEbony,
		"toreabay":                       ColorToreaBay,
		"berylgreen":                     ColorBerylGreen,
		"ginger":                         ColorGinger,
		"vegasgold":                      ColorVegasGold,
		"vesuvius":                       ColorVesuvius,
		"merino":                         ColorMerino,
		"babyblueeyes":                   ColorBabyBlueEyes,
		"hopbush":                        ColorHopbush,
		"casablanca":                     ColorCasablanca,
		"sandstorm":                      ColorSandstorm,
		"webchartreuse":                  ColorWebChartreuse,
		"babypowder":                     ColorBabyPowder,
		"bombay":                         ColorBombay,
		"myrtlegreen":                    ColorMyrtleGreen,
		"raisinblack":                    ColorRaisinBlack,
		"tuna":                           ColorTuna,
		"wattle":                         ColorWattle,
		"antiflashwhite":                 ColorAntiFlashWhite,
		"monza":                          ColorMonza,
		"mallard":                        ColorMallard,
		"coconut":                        ColorCoconut,
		"madison":                        ColorMadison,
		"tuscany":                        ColorTuscany,
		"azalea":                         ColorAzalea,
		"darkyellow":                     ColorDarkYellow,
		"darkolivegreen":                 ColorDarkOliveGreen,
		"lonestar":                       ColorLonestar,
		"earthyellow":                    ColorEarthYellow,
		"saddle":                         ColorSaddle,
		"saffronmango":                   ColorSaffronMango,
		"christalle":                     ColorChristalle,
		"deepkoamaru":                    ColorDeepKoamaru,
		"graysuit":                       ColorGraySuit,
		"mardigras":                      ColorMardiGras,
		"almond":                         ColorAlmond,
		"fieryorange":                    ColorFieryOrange,
		"pinegreen":                      ColorPineGreen,
		"violet":                         ColorViolet,
		"junebud":                        ColorJuneBud,
		"neoncarrot":                     ColorNeonCarrot,
		"lavendermagenta":                ColorLavenderMagenta,
		"yellowgreen":                    ColorYellowGreen,
		"capehoney":                      ColorCapeHoney,
		"cavernpink":                     ColorCavernPink,
		"darkimperialblue":               ColorDarkImperialBlue,
		"sanddune":                       ColorSandDune,
		"wellread":                       ColorWellRead,
		"brightube":                      ColorBrightUbe,
		"cyanazure":                      ColorCyanAzure,
		"blackshadows":                   ColorBlackShadows,
		"morningglory":                   ColorMorningGlory,
		"pastelgreen":                    ColorPastelGreen,
		"seaweed":                        ColorSeaweed,
		"solitaire":                      ColorSolitaire,
		"grannysmithapple":               ColorGrannySmithApple,
		"mikado":                         ColorMikado,
		"cinderella":                     ColorCinderella,
		"kelp":                           ColorKelp,
		"pineglade":                      ColorPineGlade,
		"bondiblue":                      ColorBondiBlue,
		"celery":                         ColorCelery,
		"eaglegreen":                     ColorEagleGreen,
		"olivegreen":                     ColorOliveGreen,
		"palemagenta":                    ColorPaleMagenta,
		"pewter":                         ColorPewter,
		"saddlebrown":                    ColorSaddleBrown,
		"superpink":                      ColorSuperPink,
		"celestialblue":                  ColorCelestialBlue,
		"denim":                          ColorDenim,
		"eden":                           ColorEden,
		"shinyshamrock":                  ColorShinyShamrock,
		"tulip":                          ColorTulip,
		"gulfblue":                       ColorGulfBlue,
		"rosefog":                        ColorRoseFog,
		"lilacluster":                    ColorLilacLuster,
		"stonewall":                      ColorStonewall,
		"sundown":                        ColorSundown,
		"bermuda":                        ColorBermuda,
		"cocoabrown":                     ColorCocoaBrown,
		"transparent":                    ColorTransparent,
	}
)

func ColorNamed(name string) Color {
	return ColorMap[strings.ToLower(name)]
}
