package client

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/go-kit/kit/endpoint"
)

func MakeClient(clientUri string, timeout time.Duration) endpoint.Endpoint {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		trama := req.([]byte)

		conn, err := net.Dial("tcp", clientUri)
		if err != nil {
			return nil, TCPError(fmt.Sprintf("Connection error: %s", err))
		}

		defer conn.Close()

		if _, err := conn.Write(trama); err != nil {
			return nil, TCPError(fmt.Sprintf("Writing error: %s", err))
		}

		response := make([]byte, 1024)
		//Si error es io.EOF simplemente se cierra la conexion
		if _, err := conn.Read(response); err != nil && err != io.EOF {
			return nil, TCPError(fmt.Sprintf("Reading error: %s", err))
		}

		return response, nil
	}
}

type TCPError string

func (e TCPError) Error() string {
	return fmt.Sprintf("TCP Error : %s", string(e))
}
