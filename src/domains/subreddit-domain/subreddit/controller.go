package subreddit

import (
	"net/http"
	"time"

	"Reddit-Clone/src/share/dto"
	"github.com/gin-gonic/gin"
)

type SubredditController struct {
	service *SubredditService
}

func NewSubredditController(service *SubredditService) *SubredditController {
	return &SubredditController{service}
}

func (c *SubredditController) Subreddits(ctx *gin.Context) {
	var pagination dto.PaginationRequest
	err := ctx.ShouldBindQuery(&pagination)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, dto.ErrorResponse{
			ErrorMessage: "Invalid pagination dto",
			Path:         "/subreddit/",
			BaseResponse: dto.BaseResponse{
				StatusCode: http.StatusBadGateway,
				TimeSpan:   time.Now(),
			},
		})
	}
	ctx.JSON(200, "subreddits")
}
