package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	//fmt.Println("Start server...")
	//adr, error := net.ResolveUDPAddr("udp", ":30000")
	//
	//if error != nil {
	//	panic(error)
	//}
	//
	//conn, error := net.ListenUDP("udp", adr)
	//
	//if error != nil {
	//	panic(error)
	//}
	//
	//buffer := make([]byte, 2048)
	//
	//_, addr, error := conn.ReadFromUDP(buffer)
	//
	//if error != nil {
	//	panic(error)
	//}
	//sendUDP()
	UDP_receive("30000")

	//addrString = addr.IP.String()
	//fmt.Println(addrString)
	//p := make([]byte, 2048)
	//
	//sendUDP("10.100.23.147:51832", 30000, "129.241.229.108", 20023, p)
	//fmt.Println(p)
	//_, error = conn.Write(b [])
	/*


		p := make([]byte, 2048)
		conn, err := net.Dial("udp", addr)
		if err != nil {
			fmt.Printf("Some error %v", err)
			return
		}
		fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
		conn.Close()*/

}

func UDP_receive(port string) {
	//vi lager en connection med servern, conection type = udp, spesifiserer port
	addr, err := net.ResolveUDPAddr("udp", ":"+port)

	//sjekker for error
	if err != nil {
		fmt.Println(err)
	}

	//vi lager en variabel concnetion som leser på porten
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
	}
	//vi lager en muffer med 1024 byes
	buffer := make([]byte, 1024)
	for {
		//vi leser connn of lagreer veriden på bufffer. n er antal bytes som lagres i buffer
		n, _, _ := conn.ReadFromUDP(buffer)

		//vi leser på hele bufferet og caster det til string
		fmt.Println(string(buffer[:n]))
		time.Sleep(2 * time.Second)
	}
}

func sendUDP() {
	con, err := net.Dial("udp", "10.100.23.147:51832")

	if err != nil {
		panic(err)
	}

	con.Write([]byte("hello\n"))
	buf, _, _ := bufio.NewReader(con).ReadLine()
	fmt.Println("clnt recv", string(buf))

}
