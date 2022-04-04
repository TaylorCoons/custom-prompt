package format

type Format struct {
	Fg ForeGround
	Em Emphasis
}

type ForeGround struct {
	Black        string
	Red          string
	Green        string
	Yellow       string
	Blue         string
	Magenta      string
	Cyan         string
	LightGray    string
	Gray         string
	LightRed     string
	LightGreen   string
	LightYellow  string
	LightBlue    string
	LightMagenta string
	LightCyan    string
	White        string
	NC           string
}

type Emphasis struct {
	Normal     string
	Bold       string
	Underlined string
	Blinking   string
	Inverted   string
}

type EmphasisType string

const (
	Bold       EmphasisType = "1"
	Underlined EmphasisType = "4"
	Blinking   EmphasisType = "5"
	Inverted   EmphasisType = "7"
)

type FGColor string

const (
	Black        FGColor = "30"
	Red          FGColor = "31"
	Green        FGColor = "32"
	Yellow       FGColor = "33"
	Blue         FGColor = "34"
	Magenta      FGColor = "35"
	Cyan         FGColor = "36"
	LightGray    FGColor = "37"
	Gray         FGColor = "90"
	LightRed     FGColor = "91"
	LightGreen   FGColor = "92"
	LightYellow  FGColor = "93"
	LightBlue    FGColor = "94"
	LightMagenta FGColor = "95"
	LightCyan    FGColor = "96"
	White        FGColor = "97"
	NC           FGColor = "0"
)
