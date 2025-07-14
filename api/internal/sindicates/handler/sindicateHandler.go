package handler


type SindicateHandler struct {
	repo repository.SindicateRepositoryInterface
}

func NewSindicateHandler(repo  repository.SindicateRepositoryInterface) *SindicateHandler {
	return &SindicateHandler{repo: repo}
}

func (h *SindicateHandler) GetAllSindicates(c *gin.Context) {}