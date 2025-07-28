package DTO

type GameResponseOnlyDTO struct {
	Game GameResponseOnly `json:"game"`
}

type GameResponseDTO struct {
	Game GameResponse `json:"game"`
}

type GamesResponseDTO struct {
	Games []GameResponseOnly `json:"games"`
}

type GameResponseOnly struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Src         string `json:"src"`
	Icon        string `json:"icon"`
}

type GameResponse struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Src         string            `json:"src"`
	Icon        string            `json:"icon"`
	Previews    []PreviewResponse `json:"previews"`
}

type CreateGameInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Src         string `json:"src" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type UpdateGameInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Src         string `json:"src" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type PreviewResponseDTO struct {
	Preview PreviewResponse `json:"preview"`
}

type PreviewResponse struct {
	ID    string  `json:"id"`
	Image string  `json:"image"`
	Video *string `json:"video,omitempty"`
}

type CreatePreviewInput struct {
	GameID string  `json:"game_id" validate:"required"`
	Image  string  `json:"image" validate:"required"`
	Video  *string `json:"video,omitempty"`
}
