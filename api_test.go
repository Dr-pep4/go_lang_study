package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    Job  string `json:"job"`
}
var users = []User{}

func main() {
    http.HandleFunc("/users", usersHandler)

    fmt.Println("서버 실행 중 : http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(users)
    }

    if r.Method == "POST" {
        var u User
        err := json.NewDecoder(r.Body).Decode(&u)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        users = append(users, u)
        w.WriteHeader(http.StatusCreated)
    }
}
