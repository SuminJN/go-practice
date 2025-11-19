package structes

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 구조체 인스턴스 선언 방법
var (
	// 일반적인 선언방식
	v1 = Vertex{1, 2}
	// X만 값을 지정해주고, Y는 int에  zero value로 설정
	v2 = Vertex{X: 1}
	// X, Y모두 int에 zero value로 설정
	v3 = Vertex{}
)

func main() {
	fmt.Println("v1.X값:", v1.X)
	v1.X = 4
	fmt.Println("v1.X = 4로 바꾼 v1.X값:", v1.X)

	var p = &v1
	p.X = 10
	fmt.Println("포인터로 바꾼 v1.X값:", v1.X)
}
