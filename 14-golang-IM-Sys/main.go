package main

//做程序进程的主入口

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
