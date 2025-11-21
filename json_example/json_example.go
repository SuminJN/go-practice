package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// JSON 파일 쓰기
	person := Person{
		Name:  "홍길동",
		Age:   30,
		Email: "hong@example.com",
	}

	// 구조체를 JSON으로 변환
	data, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("JSON 변환 오류:", err)
		return
	}

	// 파일에 쓰기
	err = os.WriteFile("person.json", data, 0644)
	if err != nil {
		fmt.Println("파일 쓰기 오류:", err)
		return
	}
	fmt.Println("JSON 파일이 생성되었습니다.")

	// JSON 파일 읽기
	fileData, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Println("파일 읽기 오류:", err)
		return
	}

	// JSON을 구조체로 변환
	var loadedPerson Person
	err = json.Unmarshal(fileData, &loadedPerson)
	if err != nil {
		fmt.Println("JSON 파싱 오류:", err)
		return
	}

	fmt.Println("\n읽어온 데이터:")
	fmt.Printf("이름: %s\n", loadedPerson.Name)
	fmt.Printf("나이: %d\n", loadedPerson.Age)
	fmt.Printf("이메일: %s\n", loadedPerson.Email)
}
