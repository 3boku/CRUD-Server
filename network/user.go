package network

import (
	"fmt"
	"github.com/3boku/Go-Server/service"
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

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
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

	var req types.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 실패입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Create	오류입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}

}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")

	var req types.UpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 실패입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Update 오류입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}

}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")

	var req types.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 실패입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Delete 오류입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}

	u.userService.Delete(nil)

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
	})
}
