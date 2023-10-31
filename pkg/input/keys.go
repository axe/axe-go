package input

type KeyInfo struct {
	Lower rune
	Upper rune
}

const (
	KeySpace        string = "space"
	KeyApostrophe   string = "apostrophe"
	KeyComma        string = "comma"
	KeyMinus        string = "minus"
	KeyPeriod       string = "period"
	KeySlash        string = "slash"
	Key0            string = "0"
	Key1            string = "1"
	Key2            string = "2"
	Key3            string = "3"
	Key4            string = "4"
	Key5            string = "5"
	Key6            string = "6"
	Key7            string = "7"
	Key8            string = "8"
	Key9            string = "9"
	KeySemicolon    string = "semicolon"
	KeyEqual        string = "equal"
	KeyA            string = "a"
	KeyB            string = "b"
	KeyC            string = "c"
	KeyD            string = "d"
	KeyE            string = "e"
	KeyF            string = "f"
	KeyG            string = "g"
	KeyH            string = "h"
	KeyI            string = "i"
	KeyJ            string = "j"
	KeyK            string = "k"
	KeyL            string = "l"
	KeyM            string = "m"
	KeyN            string = "n"
	KeyO            string = "o"
	KeyP            string = "p"
	KeyQ            string = "q"
	KeyR            string = "r"
	KeyS            string = "s"
	KeyT            string = "t"
	KeyU            string = "u"
	KeyV            string = "v"
	KeyW            string = "w"
	KeyX            string = "x"
	KeyY            string = "y"
	KeyZ            string = "z"
	KeyLeftBracket  string = "["
	KeyBackslash    string = "/"
	KeyRightBracket string = "]"
	KeyGraveAccent  string = "`"
	KeyWorld1       string = "world1"
	KeyWorld2       string = "world2"
	KeyEscape       string = "escape"
	KeyEnter        string = "enter"
	KeyTab          string = "tab"
	KeyBackspace    string = "backspace"
	KeyInsert       string = "insert"
	KeyDelete       string = "delete"
	KeyRight        string = "right"
	KeyLeft         string = "left"
	KeyDown         string = "down"
	KeyUp           string = "up"
	KeyPageUp       string = "page up"
	KeyPageDown     string = "page down"
	KeyHome         string = "home"
	KeyEnd          string = "end"
	KeyCapsLock     string = "caps lock"
	KeyScrollLock   string = "scroll lock"
	KeyNumLock      string = "num lock"
	KeyPrintScreen  string = "print"
	KeyPause        string = "pause"
	KeyF1           string = "f1"
	KeyF2           string = "f2"
	KeyF3           string = "f3"
	KeyF4           string = "f4"
	KeyF5           string = "f5"
	KeyF6           string = "f6"
	KeyF7           string = "f7"
	KeyF8           string = "f8"
	KeyF9           string = "f9"
	KeyF10          string = "f10"
	KeyF11          string = "f11"
	KeyF12          string = "f12"
	KeyF13          string = "f13"
	KeyF14          string = "f14"
	KeyF15          string = "f15"
	KeyF16          string = "f16"
	KeyF17          string = "f17"
	KeyF18          string = "f18"
	KeyF19          string = "f19"
	KeyF20          string = "f20"
	KeyF21          string = "f21"
	KeyF22          string = "f22"
	KeyF23          string = "f23"
	KeyF24          string = "f24"
	KeyF25          string = "f25"
	KeyKP0          string = "numpad 0"
	KeyKP1          string = "numpad 1"
	KeyKP2          string = "numpad 2"
	KeyKP3          string = "numpad 3"
	KeyKP4          string = "numpad 4"
	KeyKP5          string = "numpad 5"
	KeyKP6          string = "numpad 6"
	KeyKP7          string = "numpad 7"
	KeyKP8          string = "numpad 8"
	KeyKP9          string = "numpad 9"
	KeyKPDecimal    string = "numpad decimal"
	KeyKPDivide     string = "numpad divide"
	KeyKPMultiply   string = "numpad multiply"
	KeyKPSubtract   string = "numpad subtract"
	KeyKPAdd        string = "numpad add"
	KeyKPEnter      string = "numpad enter"
	KeyKPEqual      string = "numpad equal"
	KeyLeftShift    string = "left shift"
	KeyLeftControl  string = "left ctrl"
	KeyLeftAlt      string = "left alt"
	KeyLeftSuper    string = "left super"
	KeyRightShift   string = "right shift"
	KeyRightControl string = "right control"
	KeyRightAlt     string = "right alt"
	KeyRightSuper   string = "right super"
	KeyMenu         string = "menu"
)

var KeyInfos = map[string]KeyInfo{
	KeySpace:        {Upper: ' ', Lower: ' '},
	KeyApostrophe:   {Upper: '"', Lower: '\''},
	KeyComma:        {Upper: '<', Lower: ','},
	KeyMinus:        {Upper: '_', Lower: '-'},
	KeyPeriod:       {Upper: '>', Lower: '.'},
	KeySlash:        {Upper: '?', Lower: '/'},
	Key0:            {Upper: ')', Lower: '0'},
	Key1:            {Upper: '!', Lower: '1'},
	Key2:            {Upper: '@', Lower: '2'},
	Key3:            {Upper: '#', Lower: '3'},
	Key4:            {Upper: '$', Lower: '4'},
	Key5:            {Upper: '%', Lower: '5'},
	Key6:            {Upper: '^', Lower: '6'},
	Key7:            {Upper: '&', Lower: '7'},
	Key8:            {Upper: '*', Lower: '8'},
	Key9:            {Upper: '(', Lower: '9'},
	KeySemicolon:    {Upper: ':', Lower: ';'},
	KeyEqual:        {Upper: '+', Lower: '='},
	KeyA:            {Upper: 'A', Lower: 'a'},
	KeyB:            {Upper: 'B', Lower: 'b'},
	KeyC:            {Upper: 'C', Lower: 'c'},
	KeyD:            {Upper: 'D', Lower: 'd'},
	KeyE:            {Upper: 'E', Lower: 'e'},
	KeyF:            {Upper: 'F', Lower: 'f'},
	KeyG:            {Upper: 'G', Lower: 'g'},
	KeyH:            {Upper: 'H', Lower: 'h'},
	KeyI:            {Upper: 'I', Lower: 'i'},
	KeyJ:            {Upper: 'J', Lower: 'j'},
	KeyK:            {Upper: 'K', Lower: 'k'},
	KeyL:            {Upper: 'L', Lower: 'l'},
	KeyM:            {Upper: 'M', Lower: 'm'},
	KeyN:            {Upper: 'N', Lower: 'n'},
	KeyO:            {Upper: 'O', Lower: 'o'},
	KeyP:            {Upper: 'P', Lower: 'p'},
	KeyQ:            {Upper: 'Q', Lower: 'q'},
	KeyR:            {Upper: 'R', Lower: 'r'},
	KeyS:            {Upper: 'S', Lower: 's'},
	KeyT:            {Upper: 'T', Lower: 't'},
	KeyU:            {Upper: 'U', Lower: 'u'},
	KeyV:            {Upper: 'V', Lower: 'v'},
	KeyW:            {Upper: 'W', Lower: 'w'},
	KeyX:            {Upper: 'X', Lower: 'x'},
	KeyY:            {Upper: 'Y', Lower: 'y'},
	KeyZ:            {Upper: 'Z', Lower: 'z'},
	KeyLeftBracket:  {Upper: '{', Lower: '['},
	KeyBackslash:    {Upper: '|', Lower: '\\'},
	KeyRightBracket: {Upper: '}', Lower: ']'},
	KeyGraveAccent:  {Upper: '~', Lower: '`'},
	KeyEnter:        {Upper: '\n', Lower: '\n'},
	KeyTab:          {Upper: '\t', Lower: '\t'},
	KeyBackspace:    {Upper: '\b', Lower: '\b'},
	KeyKP0:          {Upper: '0', Lower: '0'},
	KeyKP1:          {Upper: '1', Lower: '1'},
	KeyKP2:          {Upper: '2', Lower: '2'},
	KeyKP3:          {Upper: '3', Lower: '3'},
	KeyKP4:          {Upper: '4', Lower: '4'},
	KeyKP5:          {Upper: '5', Lower: '5'},
	KeyKP6:          {Upper: '6', Lower: '6'},
	KeyKP7:          {Upper: '7', Lower: '7'},
	KeyKP8:          {Upper: '8', Lower: '8'},
	KeyKP9:          {Upper: '9', Lower: '9'},
	KeyKPDecimal:    {Upper: '.', Lower: '.'},
	KeyKPDivide:     {Upper: '/', Lower: '/'},
	KeyKPMultiply:   {Upper: '*', Lower: '*'},
	KeyKPSubtract:   {Upper: '-', Lower: '-'},
	KeyKPAdd:        {Upper: '+', Lower: '+'},
	KeyKPEnter:      {Upper: '\n', Lower: '\n'},
	KeyKPEqual:      {Upper: '=', Lower: '='},
}
