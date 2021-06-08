package main

import (
	"fmt"
	"net/http"

	"github.com/yudiwiradinata/go-rest-api-mux/controller"
	router "github.com/yudiwiradinata/go-rest-api-mux/http"
	"github.com/yudiwiradinata/go-rest-api-mux/repository"
	"github.com/yudiwiradinata/go-rest-api-mux/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter() //router.NewMuxRouter()
)

const (
	PORT = ":9000"
)

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up and running.....")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(PORT)
}
