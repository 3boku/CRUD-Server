# Go-Backend
 
## 1일차

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

```go
type Config struct {
	Server struct {
		Port int64
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
```
파일을 열어 환경변수 toml파일을 Config하는 코드를 짰다.

```go
// config.go
package cmd

import (
	"fmt"
	"github.com/3boku/Go-Server/config"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	fmt.Println(c.config.Server.Port)
	return c
}

```
config한 toml파일의 키를 출력
```go
// cmd.go
package cmd

import (
	"fmt"
	"github.com/3boku/Go-Server/config"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	fmt.Println(c.config.Server.Port)
	return c
}

```
```go
// main.go
package main

import (
	"flag"
	"github.com/3boku/Go-Server/init/cmd"
)

var configPathFlag = flag.String("confg", "config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}

```
config.go파일을 실행

configPathFlag라는 포인터 스트링 변수를 만들어서 파일 트리가 바뀌어도 동작 있게함

toml을 쓰는이유: 각 로컬마다 환경변수가 달라서

```go
//network/root.go
package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engine: gin.New(),
	}

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
```
gin을 사용해서 서버를 실행하는 코드를 짰다. 엔진이라는 스트럭트를 받아와서 서버를 시작하는 코드이다.
```go
// cmd.go
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