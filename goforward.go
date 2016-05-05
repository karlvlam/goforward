/*
Config file format:
0.0.0.0:5001 server:80
0.0.0.0:5002 server:22
*/


package main

//import "fmt"
import "strings"
import "io/ioutil"
import "io"
import "log"
import "net"
import "os"
import "time"


func forward(conn net.Conn, targetAddress string) {
    client, err := net.Dial("tcp", targetAddress)
    if err != nil {
        log.Fatalf("Dial failed: %v", err)
    }
    log.Printf("Connected to localhost %v\n", conn)
    go func() {
        defer client.Close()
        defer conn.Close()
        buf := make([]byte, 8192)
        io.CopyBuffer(client, conn, buf)
    }()
    go func() {
        defer client.Close()
        defer conn.Close()
        buf := make([]byte, 8192)
        io.CopyBuffer(conn, client, buf)
    }()
}

func createPortForward(listenAddress string, targetAddress string){
    listener, err := net.Listen("tcp", listenAddress)
    if err != nil {
        log.Fatalf("Failed to setup listener: %v", err)
        return
    }


    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalf("ERROR: failed to accept listener: %v", err)
        }
        log.Printf("Accepted connection %v\n", conn)
        go forward(conn, targetAddress)
    }

}

func main() {
    if len(os.Args) != 2 {
        log.Fatalf("Usage %s configfile\n", os.Args[0]);
        return
    }

    configfile := os.Args[1]

    //TODO: use better data structure
    raw, err := ioutil.ReadFile(configfile)
    if (err != nil) {
        log.Fatalf("Read file error: %s\n", err)
        return
    }

    lines := strings.Split(string(raw[:]), "\n")

    for _, line := range lines {
        addr := strings.Split(line, " ")

        if len(addr) != 2 {
            //fmt.Println(addr)
            continue
        }
        go createPortForward(addr[0], addr[1])
    }

    // just keep the process alive
    for {
        time.Sleep(1 * time.Hour)
    }

}
