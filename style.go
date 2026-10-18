package main

import (
	_ "embed"

	"github.com/egoist/mygo/plugins/terminal"
	"github.com/egoist/mygo/ui"
)

// colors are the app's own, in light and dark windows.
type colors struct {
	// The window's background: a gradient from top through mid to bottom,
	// and a tint toward the right.
	bgTop, bgMid, bgLow, bgBottom, bgTint ui.Color

	card, cardFocused, cardBorder, cardBorderFocused ui.Color
	shadow, shadowFocused                            ui.Color

	text, textMuted, textFaint, iconMuted ui.Color
	hover, pressed                        ui.Color

	track, trackBorder, tabActive, tabSep ui.Color
	tileRim                               ui.Color

	busy, attention, exited ui.Color

	panel, panelBorder, panelSel, backdrop ui.Color
}

var lightColors = colors{
	bgTop:    ui.Hex("#f6e7f6"),
	bgMid:    ui.Hex("#efe1e6"),
	bgLow:    ui.Hex("#eee0d3"),
	bgBottom: ui.Hex("#ece1bd"),
	bgTint:   ui.RGBA(242, 214, 222, 0.35),

	card:              ui.RGBA(255, 255, 255, 0.66),
	cardFocused:       ui.RGBA(255, 255, 255, 0.95),
	cardBorder:        ui.RGBA(255, 255, 255, 0.72),
	cardBorderFocused: ui.RGBA(255, 255, 255, 1),
	shadow:            ui.RGBA(60, 40, 50, 0.07),
	shadowFocused:     ui.RGBA(60, 40, 50, 0.12),

	text:      ui.Hex("#1d1d1f"),
	textMuted: ui.Hex("#636366"),
	textFaint: ui.Hex("#55555a"),
	iconMuted: ui.Hex("#717176"),
	hover:     ui.RGBA(0, 0, 0, 0.055),
	pressed:   ui.RGBA(0, 0, 0, 0.1),

	track:       ui.RGBA(255, 255, 255, 0.34),
	trackBorder: ui.RGBA(255, 255, 255, 0.45),
	tabActive:   ui.RGBA(255, 255, 255, 0.97),
	tabSep:      ui.RGBA(0, 0, 0, 0.13),
	tileRim:     ui.RGBA(255, 255, 255, 0.95),

	busy:      ui.Hex("#34a853"),
	attention: ui.Hex("#f59e0b"),
	exited:    ui.Hex("#a1a1a6"),

	panel:       ui.RGBA(252, 252, 253, 0.97),
	panelBorder: ui.RGBA(0, 0, 0, 0.08),
	panelSel:    ui.RGBA(46, 111, 208, 0.12),
	backdrop:    ui.RGBA(40, 20, 40, 0.06),
}

var darkColors = colors{
	bgTop:    ui.Hex("#2c2131"),
	bgMid:    ui.Hex("#231e27"),
	bgLow:    ui.Hex("#221f22"),
	bgBottom: ui.Hex("#2a2619"),
	bgTint:   ui.RGBA(80, 40, 60, 0.25),

	card:              ui.RGBA(20, 20, 24, 0.5),
	cardFocused:       ui.RGBA(36, 36, 41, 0.94),
	cardBorder:        ui.RGBA(255, 255, 255, 0.05),
	cardBorderFocused: ui.RGBA(255, 255, 255, 0.14),
	shadow:            ui.RGBA(0, 0, 0, 0.25),
	shadowFocused:     ui.RGBA(0, 0, 0, 0.4),

	text:      ui.Hex("#f2f2f7"),
	textMuted: ui.Hex("#aeaeb2"),
	textFaint: ui.Hex("#8e8e93"),
	iconMuted: ui.Hex("#98989d"),
	hover:     ui.RGBA(255, 255, 255, 0.08),
	pressed:   ui.RGBA(255, 255, 255, 0.14),

	track:       ui.RGBA(255, 255, 255, 0.06),
	trackBorder: ui.RGBA(255, 255, 255, 0.06),
	tabActive:   ui.RGBA(255, 255, 255, 0.14),
	tabSep:      ui.RGBA(255, 255, 255, 0.14),
	tileRim:     ui.RGBA(255, 255, 255, 0.22),

	busy:      ui.Hex("#4ade80"),
	attention: ui.Hex("#fbbf24"),
	exited:    ui.Hex("#636366"),

	panel:       ui.RGBA(38, 38, 42, 0.97),
	panelBorder: ui.RGBA(255, 255, 255, 0.1),
	panelSel:    ui.RGBA(108, 178, 255, 0.18),
	backdrop:    ui.RGBA(0, 0, 0, 0.2),
}

var gruvboxDarkColors = colors{
	bgTop:    ui.Hex("#32302f"),
	bgMid:    ui.Hex("#282828"),
	bgLow:    ui.Hex("#242221"),
	bgBottom: ui.Hex("#1d2021"),
	bgTint:   ui.RGBA(214, 93, 14, 0.14),

	card:              ui.RGBA(32, 30, 29, 0.6),
	cardFocused:       ui.RGBA(40, 40, 40, 0.96),
	cardBorder:        ui.RGBA(235, 219, 178, 0.06),
	cardBorderFocused: ui.RGBA(235, 219, 178, 0.18),
	shadow:            ui.RGBA(0, 0, 0, 0.28),
	shadowFocused:     ui.RGBA(0, 0, 0, 0.45),

	text:      ui.Hex("#ebdbb2"),
	textMuted: ui.Hex("#bdae93"),
	textFaint: ui.Hex("#a89984"),
	iconMuted: ui.Hex("#928374"),
	hover:     ui.RGBA(235, 219, 178, 0.08),
	pressed:   ui.RGBA(235, 219, 178, 0.14),

	track:       ui.RGBA(235, 219, 178, 0.06),
	trackBorder: ui.RGBA(235, 219, 178, 0.08),
	tabActive:   ui.RGBA(235, 219, 178, 0.14),
	tabSep:      ui.RGBA(235, 219, 178, 0.12),
	tileRim:     ui.RGBA(235, 219, 178, 0.22),

	busy:      ui.Hex("#b8bb26"),
	attention: ui.Hex("#fabd2f"),
	exited:    ui.Hex("#928374"),

	panel:       ui.RGBA(40, 40, 40, 0.97),
	panelBorder: ui.RGBA(235, 219, 178, 0.12),
	panelSel:    ui.RGBA(131, 165, 152, 0.22),
	backdrop:    ui.RGBA(0, 0, 0, 0.35),
}

var gruvboxLightColors = colors{
	bgTop:    ui.Hex("#fbf1c7"),
	bgMid:    ui.Hex("#f2e5bc"),
	bgLow:    ui.Hex("#eee0b7"),
	bgBottom: ui.Hex("#dfcf9f"),
	bgTint:   ui.RGBA(215, 153, 33, 0.2),

	card:              ui.RGBA(242, 229, 188, 0.65),
	cardFocused:       ui.RGBA(251, 241, 199, 0.96),
	cardBorder:        ui.RGBA(60, 56, 54, 0.10),
	cardBorderFocused: ui.RGBA(60, 56, 54, 0.25),
	shadow:            ui.RGBA(60, 56, 54, 0.08),
	shadowFocused:     ui.RGBA(60, 56, 54, 0.16),

	text:      ui.Hex("#282828"),
	textMuted: ui.Hex("#504945"),
	textFaint: ui.Hex("#665c54"),
	iconMuted: ui.Hex("#7c6f64"),
	hover:     ui.RGBA(60, 56, 54, 0.06),
	pressed:   ui.RGBA(60, 56, 54, 0.12),

	track:       ui.RGBA(235, 219, 178, 0.5),
	trackBorder: ui.RGBA(60, 56, 54, 0.12),
	tabActive:   ui.RGBA(251, 241, 199, 0.95),
	tabSep:      ui.RGBA(60, 56, 54, 0.12),
	tileRim:     ui.RGBA(251, 241, 199, 0.95),

	busy:      ui.Hex("#79740e"),
	attention: ui.Hex("#b57614"),
	exited:    ui.Hex("#928374"),

	panel:       ui.RGBA(249, 245, 215, 0.97),
	panelBorder: ui.RGBA(60, 56, 54, 0.14),
	panelSel:    ui.RGBA(7, 102, 120, 0.14),
	backdrop:    ui.RGBA(40, 40, 40, 0.12),
}

func colorsOf(c *ui.Context) *colors {
	switch prefs.Appearance {
	case "gruvbox", "gruvbox-dark", "gruvbox dark":
		return &gruvboxDarkColors
	case "gruvbox-light", "gruvbox light":
		return &gruvboxLightColors
	case "dark":
		return &darkColors
	case "light":
		return &lightColors
	default:
		if c.Theme().Dark {
			return &darkColors
		}
		return &lightColors
	}
}

// paintBackground paints the window's background.
func paintBackground(p *ui.Painter, r ui.Rect, k *colors) {
	h1, h2 := r.H*0.38, r.H*0.32
	p.FillGradient(ui.Rect{X: r.X, Y: r.Y, W: r.W, H: h1 + 1}, ui.LinearGradient{From: k.bgTop, To: k.bgMid, Angle: 180}, 0)
	p.FillGradient(ui.Rect{X: r.X, Y: r.Y + h1, W: r.W, H: h2 + 1}, ui.LinearGradient{From: k.bgMid, To: k.bgLow, Angle: 180}, 0)
	p.FillGradient(ui.Rect{X: r.X, Y: r.Y + h1 + h2, W: r.W, H: r.H - h1 - h2}, ui.LinearGradient{From: k.bgLow, To: k.bgBottom, Angle: 180}, 0)
	clear := k.bgTint
	clear.A = 0
	p.FillGradient(r, ui.LinearGradient{From: clear, To: k.bgTint, Angle: 90, Start: 0.35, End: 1}, 0)
}

//go:embed assets/fonts/JetBrainsMono-Regular.ttf
var fontRegular []byte

//go:embed assets/fonts/JetBrainsMono-Bold.ttf
var fontBold []byte

//go:embed assets/fonts/JetBrainsMono-Italic.ttf
var fontItalic []byte

//go:embed assets/fonts/JetBrainsMono-BoldItalic.ttf
var fontBoldItalic []byte

func registerFonts() {
	for _, f := range [][]byte{fontRegular, fontBold, fontItalic, fontBoldItalic} {
		ui.RegisterFont(f, "JetBrains Mono")
	}
}

var termFont = terminal.Font{Family: "JetBrains Mono, SF Mono, Menlo, monospace", Size: 11.6, LineHeight: 1.0}

// The terminals' colors: soft ink on the panes' paper, and the dark kind.
var lightTerm = &terminal.Theme{
	Foreground: ui.Hex("#2a2d31"),
	Background: ui.Hex("#fbfbfb"),
	Cursor:     ui.Hex("#3a3d42"),
	Selection:  ui.RGBA(46, 111, 208, 0.2),
	Palette: [16]ui.Color{
		ui.Hex("#2a2d31"), ui.Hex("#c2465a"), ui.Hex("#4e8e5f"), ui.Hex("#b07a1e"),
		ui.Hex("#2e6fd0"), ui.Hex("#9050c8"), ui.Hex("#1e8c9c"), ui.Hex("#2a2d31"),
		ui.Hex("#55555a"), ui.Hex("#d9566b"), ui.Hex("#5ba671"), ui.Hex("#c89026"),
		ui.Hex("#4a8fe0"), ui.Hex("#a56bd8"), ui.Hex("#2ba2b3"), ui.Hex("#1d1d1f"),
	},
}

var darkTerm = &terminal.Theme{
	Foreground: ui.Hex("#e6e6ea"),
	Background: ui.Hex("#1e1e22"),
	Cursor:     ui.Hex("#e6e6ea"),
	Selection:  ui.RGBA(108, 178, 255, 0.28),
	Palette: [16]ui.Color{
		ui.Hex("#3a3a40"), ui.Hex("#ff6b7f"), ui.Hex("#7bd88f"), ui.Hex("#e5c07b"),
		ui.Hex("#6cb2ff"), ui.Hex("#c792ea"), ui.Hex("#56d4dd"), ui.Hex("#d0d0d6"),
		ui.Hex("#6e6e78"), ui.Hex("#ff8c9c"), ui.Hex("#9be8a8"), ui.Hex("#f2d28a"),
		ui.Hex("#8ec5ff"), ui.Hex("#d7a8f2"), ui.Hex("#7fe3ea"), ui.Hex("#ffffff"),
	},
}

var gruvboxDarkTerm = &terminal.Theme{
	Foreground: ui.Hex("#ebdbb2"),
	Background: ui.Hex("#282828"),
	Cursor:     ui.Hex("#ebdbb2"),
	Selection:  ui.RGBA(102, 92, 84, 0.45),
	Palette: [16]ui.Color{
		ui.Hex("#282828"), ui.Hex("#cc241d"), ui.Hex("#98971a"), ui.Hex("#d79921"),
		ui.Hex("#458588"), ui.Hex("#b16286"), ui.Hex("#689d6a"), ui.Hex("#a89984"),
		ui.Hex("#928374"), ui.Hex("#fb4934"), ui.Hex("#b8bb26"), ui.Hex("#fabd2f"),
		ui.Hex("#83a598"), ui.Hex("#d3869b"), ui.Hex("#8ec07c"), ui.Hex("#ebdbb2"),
	},
}

var gruvboxLightTerm = &terminal.Theme{
	Foreground: ui.Hex("#3c3836"),
	Background: ui.Hex("#fbf1c7"),
	Cursor:     ui.Hex("#3c3836"),
	Selection:  ui.RGBA(168, 153, 132, 0.35),
	Palette: [16]ui.Color{
		ui.Hex("#3c3836"), ui.Hex("#cc241d"), ui.Hex("#98971a"), ui.Hex("#d79921"),
		ui.Hex("#458588"), ui.Hex("#b16286"), ui.Hex("#689d6a"), ui.Hex("#3c3836"),
		ui.Hex("#665c54"), ui.Hex("#9d0006"), ui.Hex("#79740e"), ui.Hex("#b57614"),
		ui.Hex("#076678"), ui.Hex("#8f3f71"), ui.Hex("#427b58"), ui.Hex("#282828"),
	},
}

func termThemes() (light, dark *terminal.Theme) {
	switch prefs.Appearance {
	case "gruvbox", "gruvbox-dark", "gruvbox dark":
		return gruvboxDarkTerm, gruvboxDarkTerm
	case "gruvbox-light", "gruvbox light":
		return gruvboxLightTerm, gruvboxLightTerm
	case "dark":
		return darkTerm, darkTerm
	case "light":
		return lightTerm, lightTerm
	default:
		return lightTerm, darkTerm
	}
}
