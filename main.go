package main

import (
	"fmt"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/gorilla/websocket"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/owncast/owncast/logging"
)

func main() {
	fmt.Println("Hello world")
}
