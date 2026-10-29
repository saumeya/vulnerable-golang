package main

import (
	"fmt"

	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"
	_ "github.com/gorilla/websocket"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/tidwall/gjson"
	_ "golang.org/x/crypto/md4"
	_ "gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Hello world")
}
