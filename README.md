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

