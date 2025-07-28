package view

// ------------------------------
// View

type GameView struct {
	ID          string
	Title       string
	Description string
	Src         string
	Icon        string
}

type GamePreivewView struct {
	GameView
	Preview []PreviewView
}

type PreviewView struct {
	ID    string
	Image string
	Video *string
}

// ------------------------------
// Page

type HomePageData struct {
	Games []GameView
}

type ShowPageData struct {
	Game GamePreivewView
}

type GamePageData struct {
	Game GameView
}
