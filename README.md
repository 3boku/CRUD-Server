# Go-Backend
 
## 1번째 영상
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld) //새로운 핸들러 함수 가져오기

	if err := http.ListenAndServe(":8080", nil); err != nil{ //서버 꺼지지 않게 실행
		fmt.Println("에러 발생")
		panic(err)
		return
	}
}

func helloworld(w http.ResponseWriter, r *http.Request) { //helloworld라는 이벤트 핸들러에 리스폰스 라이터와, 리퀘스트를 매개변수로 선언
	fmt.Println("Hello, world")
}
```
내부 http 패키지로 헬로우 월드를 출력하는 간단한 서버를 구축했다.

2번째 영상
```go
//config/config.go
type Config struct {
	Server struct { // server라는 구조체를 만듦
		Port int64//port를 int64형식이라고 정의함
	}
} //사용할 키값들을 정리하는 구조체

func NewConfig(filePath string) *Config {
	//NewConfig라는 함수를 만들어 filePath를 매개변수로 사용하고, 위에서만든 Config를 리턴해줄것임
	c := new(Config) // new를 사용해서 config 구조체를 만듦
	//var c &Config{} 라고 적어도됨

	if file, err := os.Open(filePath); err != nil { //file을 오픈함
		panic(err) //오픈하는데 오류가 발생하면 오류를 출력함
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		// toml패키지의 NewDecoder명령어를 사용해서 toml파일을 읽어서 디코더를 선언해주고, Decode명령어를 사용해서 Decode 해준다
		// 여기서 Decode는 부호화된 코드를 부호화되기 전으로 바꾸는 것
		panic(err)
		//Decode하는데 오류가 발생하면 오류를 출력함
	} else {
		return c
	}
}
```
파일을 열어 환경변수 toml파일을 Config하는 코드를 짰다.

```go
// init/cmd/cmd.go
package cmd

import (
	"fmt"
	"github.com/3boku/Go-Server/config"
)

type Cmd struct {
	config *config.Config
	//config를 사용할수 있게 config디렉토리 안에 Config코드를 가져온다 생각하면 될 것 같음
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{// Cmd라는 새로운 구조체를 만든다
		config: config.NewConfig(filepath),
		//config를 초기화 해준다.
	}

	fmt.Println(c.config.Server.Port)
	//config파일에서 디코딩한 config.toml파일의 Server 안에있는 Port라는 값을 출력한다.
	return c
}

```
```go
// init/main.go
package main

import (
	"flag"
	"github.com/3boku/Go-Server/init/cmd"
)

var configPathFlag = flag.String("confg", "config.toml", "config file not found")
//configPathFlag라는 변수를 만들어서 이름을 config, 파일은 config.toml, 없으면 confgi file not found라는 텍스트를 출력한다.
//사용하는 이유는 config.toml파일이 로컬 패스에 있는지 없는지 확인하기 위해
func main() {
	flag.Parse()
	//Parsing해온다.
	//parsing이란 번역 정도로 생각하면 될것이다.
	cmd.NewCmd(*configPathFlag)
}

```
config.go파일을 실행

configPathFlag라는 포인터 스트링 변수를 만들어서 파일 트리가 바뀌어도 동작 있게함

toml을 쓰는이유: 각 로컬마다 환경변수가 달라서

## 3번쨰 영상
```go
// network/root.go
package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
} //engine은 gin에 있는 Engine이라는 구조체를 갖게됨

func NewNetwork() *Network {
	r := &Network{
		engine: gin.New(),
		//gin.Deafult()를 사용해도 된다, New는 프로덕트 환경, Deafault는 테스트 환경용 이라고 한다.
	}

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
} // 서버를 시작하는 함수이다. port를 매개변수로 받는다. 
// Run은 채널로써 관리가 되기 때문에 서버가 시작되면 다른 코드들이 동작을 하지 않 아 함수로 따로 뺐다.
```
gin을 사용해서 서버를 실행하는 코드를 짰다. 엔진이라는 스트럭트를 받아와서 서버를 시작하는 코드이다.
```go
// init/cmd/cmd.go
package cmd

import (
	"fmt"
	"github.com/3boku/Go-Server/config"
	"github.com/3boku/Go-Server/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network 
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filepath),
		network: network.NewNetwork(),
	}

	network.NewNetwork().ServerStart(c.config.Server.Port)
	fmt.Println(c.config.Server.Port)
	return c
}
```
cmd.go파일에서 network 스트럭트를 불러와서 실행을 했다. 이때 prot는 toml파일에 정의되어 있는 포트로 실행을 한다.

## 4번째 영상
```go
// network/user.go
package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once //userRouterInit이 한번만 실행되게 한다.
	userRouterInstance *userRouter //userRouterInstance는 네트워크의 라우터를 가져온다
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{ //userRouterInstance를 userRouter메모리 주소를 대입한다.
			router: router,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELTE("/", userRouterInstance.delete) //CRUD를 등록한다.

	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	//함수명을 시작할때 첫글자를 대문자로 쓰면 다른 레포지스토리나 디렉토리에서도 접근가능
	//하지만 이 함수는 이 파일에서만 돌아가야 하기 때문에 create
	fmt.Println("create")
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

```
여기서 CRUD를 등록할 수 있는 함수는 root.go에 있다.
```go
// network/root.go
package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engine: gin.New(),
	}

	newUserRouter(r) //위에서 만든 함수 호출

	return r
}



func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}

//resister 추가 함수`
func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDELTE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

```

## 7번째 영상
```go
// network/utils.go

package network

import "github.com/gin-gonic/gin"

//resister 추가 함수들

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}

func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDELTE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

// Response 형태 맞추기 위한 유틸 함수 입니다.

func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
} //okResponse가 실행되면 200코드와 result인터페이스를 리턴함

func (n *Network) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(400, result)
}
```

```go
// types/utils.go
package types

type ApiResponse struct {
	Result      int64  `json:"result"`
	Description string `json:"description"`
}
//ApiResponse를 정의함
```

```go
// types/user.go
package types

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	*ApiResponse
	*User
}
//UserResponse는 User 구조체와 ApiResponse구조체를 갖고있는다.
```
이렇게 Api 전송할때 쓸 인터페이스를 정의하고

```go
// network/user.go 중 크리에이트 함수
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
	// userRouter안에 라우터안에 okResponse를 보낸다. 보낼때 아까 types에서 정의한 유틸함수들을 리턴한다.
}
```

## 8번째 영상
```go
// types/utils.go
//각 리스폰스마다의 타입을 지정해주었다.
package types

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	*ApiResponse
	*User
}

type GetUserResponse struct {
	*ApiResponse
	*User
}

type CreateUserResponse struct {
	*ApiResponse
	*User
}

type UpdateUserResponse struct {
	*ApiResponse
	*User
}

type DeleteUserResponse struct {
	*ApiResponse
	*User
}
```

```go
//network/user.go 중
func (u *userRouter) create(c *gin.Context) {
	//함수명을 시작할때 첫글자를 대문자로 쓰면 다른 레포지스토리나 디렉토리에서도 접근가능
	//하지만 이 함수는 이 파일에서만 돌아가야 하기 때문에 create
	fmt.Println("create")

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
		User:        nil,
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}
```
그리고 service와 network, repository를 연결해주겠다
```go
//service/root.go
import (
	"github.com/3boku/Go-Server/repository"
	"sync"
)

// Network와 Repository의 다리 역할

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository
	//repository

	User *User
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}

		serviceInstance.User = newUserService(rep.User)
	})

	return serviceInstance
}
```
```go
//repository/root.go
package repository

import (
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	//repository 이곳에선 데이터베이스 같은것들을 설정해줌
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: NewUserRepository(),
		}
	})

	return repositoryInstance
}

```
대충 레포지스토리랑 네트워크의 다리 역할을 서비스가 하고, 레포지스토리의 유저가 db역할을 하며 왔다 갔다 하는 코드임

## 9번째 영상
```go
//repository/user.go
func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Update(beforUser *types.User, updatedUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}

```
이런 메소드를 만들어서
```go
/service/user.go
func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(beforUser *types.User, updatedUser *types.User) error {
	return u.userRepository.Update(beforUser, updatedUser)
}

func (u *User) Delete(user *types.User) error {
	return u.userRepository.Delete(user)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}

```
서비스로 넘겨준걸
```go
//network/user.go
func (u *userRouter) create(c *gin.Context) {
	//함수명을 시작할때 첫글자를 대문자로 쓰면 다른 레포지스토리나 디렉토리에서도 접근가능
	//하지만 이 함수는 이 파일에서만 돌아가야 하기 때문에 create
	fmt.Println("create")

	u.userService.Create(nil)

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
		User:        nil,
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")

	u.userService.Update(nil, nil)

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")

	u.userService.Delete(nil)

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}
```
유저로 넘겨주었다.

## 10번쨰 영상
서비스 추가했다. 코드가 너무 길어서 각 폴더에 user.go확인하면 될듯

## 11번째 영상
Post Man쓰는이유 API확인을 위해
고랭에서 스웨거 안쓰고 Post Man 쓰는이유: 지원 안하는 버전도 있음

