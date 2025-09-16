package main

import (
	"fmt"
	"net/http"
)

// helloHandler는 루트 URL("/")에 대한 요청을 처리하는 핸들러 함수입니다.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 클라이언트에게 응답을 전송
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// 루트 경로("/")에 helloHandler를 연결
	http.HandleFunc("/", helloHandler)

	// 서버 시작 알림 출력
	fmt.Println("Listening on port 8080")

	// 포트 8080에서 HTTP 서버 시작
	// nil은 기본 라우터(DefaultServeMux)를 사용한다는 뜻
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// 에러가 발생하면 로그 출력
		fmt.Println("Server error:", err)
	}
}
