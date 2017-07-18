package main

import (
	"github.com/astaxie/beego"
        "log"
        "net"
        "strings"
        "time"
)

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    localAddr := GetOutboundIP().String()
    clientAddr := strings.Split(this.Ctx.Request.RemoteAddr, ":")
    this.Ctx.WriteString("Local IP: ")
    this.Ctx.WriteString(localAddr)
    this.Ctx.WriteString("\n")
    this.Ctx.WriteString("Client IP: ")
    this.Ctx.WriteString(clientAddr[0])
    this.Ctx.WriteString("\n")
    this.Ctx.WriteString("Current Time: ")
    this.Ctx.WriteString(time.Now().Format(time.RFC850))
}

func main() {
	beego.Router("/", &MainController{})
        beego.Run()
}

