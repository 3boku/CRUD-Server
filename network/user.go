package network

import (
	"fmt"
	"github.com/3boku/Go-Server/types"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELTE("/", userRouterInstance.delete)

	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	//함수명을 시작할때 첫글자를 대문자로 쓰면 다른 레포지스토리나 디렉토리에서도 접근가능
	//하지만 이 함수는 이 파일에서만 돌아가야 하기 때문에 create
	fmt.Println("create")

	u.router.okResponse(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공입니다.",
		},
		User: nil,
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get")
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")
}
