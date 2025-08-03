package main

import (
    "fmt"
)

func main() {
    var name string
    var age int
    var job string

    fmt.Print("이름을 입력하세요: ")
    fmt.Scanln(&name)

    fmt.Print("나이를 입력하세요: ")
    fmt.Scanln(&age)

    fmt.Print("직업을 입력하세요: ")
    fmt.Scanln(&job)


    fmt.Println("안녕하세요,", name,"(",age,")",job," 님!")
}
