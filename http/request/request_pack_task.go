package request

import (
	"fmt"
	"io/ioutil"
	"kkAndroidPackClient/config"
	"kkAndroidPackClient/db/bean"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type TimeoutConn struct {
	conn    net.Conn
	timeout time.Duration
}

func NewTimeoutConn(conn net.Conn, timeout time.Duration) *TimeoutConn {
	return &TimeoutConn{
		conn:    conn,
		timeout: timeout,
	}
}

func (c *TimeoutConn) Read(b []byte) (n int, err error) {
	c.SetReadDeadline(time.Now().Add(c.timeout))
	return c.conn.Read(b)
}

func (c *TimeoutConn) Write(b []byte) (n int, err error) {
	c.SetWriteDeadline(time.Now().Add(c.timeout))
	return c.conn.Write(b)
}

func (c *TimeoutConn) Close() error {
	return c.conn.Close()
}

func (c *TimeoutConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TimeoutConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TimeoutConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *TimeoutConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *TimeoutConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

type PackageAppJSONResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	App     []bean.PackageApp `json:"data,omitempty"`
}

func RequestPackTask() *PackageAppJSONResponse {
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", host)
	}

	fmt.Println("get start new")
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				log.Printf("dial to %s://%s", netw, addr)

				conn, err := net.DialTimeout(netw, addr, time.Second*10)

				if err != nil {
					return nil, err
				}

				return NewTimeoutConn(conn, time.Second*10), nil
			},
			ResponseHeaderTimeout: time.Second * 10,
		},
	}
	request, _ := http.NewRequest("GET", config.ServerHost+"fetchPackTask?hostName="+host, nil)
	resp, err := client.Do(request)

	fmt.Println("get finished")
	if err != nil {
		return nil
	}

	jsonResponse := new(PackageAppJSONResponse)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("return nil")
		return nil
	}

	decodeJSONResponse(body, jsonResponse)
	//fmt.Println(jsonResponse.App[0].ApkName)
	return jsonResponse
}
