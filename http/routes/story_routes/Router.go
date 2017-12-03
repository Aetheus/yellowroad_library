package story_routes

import (
	"yellowroad_library/containers"

	"github.com/gin-gonic/gin"
)

func Register(
	router *gin.RouterGroup,
	container containers.Container,
) {
	storyHandlers := StoryHandlers{container:container}

	router.GET("stories/:book_id/chapter/:chapter_id/game", storyHandlers.NavigateToSingleChapter)
}
