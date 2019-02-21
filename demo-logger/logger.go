// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
	"io"
	"os"
)

var (
	green        = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white        = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow       = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red          = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue         = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta      = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan         = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

var DefaultWriter io.Writer = os.Stdout

func LoggerWithWriter(out io.Writer)  {
	start := time.Now()

	isTerm := true
	end := time.Now()
	latency := end.Sub(start)

	var statusColor, methodColor, resetColor string
	statusCode := 200
	if isTerm {
		statusColor = colorForStatus(200)
		methodColor = colorForMethod("POST")
		resetColor = reset
	}
	fmt.Fprintf(out, "[GIN] %v |%s %3d %s| %v %s",
		end.Format("2006/01/02 - 15:04:05"),
		statusColor,
		statusCode,
		resetColor,
		latency,
		methodColor,
	)
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}

func main()  {
	LoggerWithWriter(DefaultWriter)
}
