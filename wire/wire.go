package wire

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

type Wire struct {
	engine *gin.Engine
}

func New() *Wire {
	return &Wire{}
}

func (s *Wire) Init(engine *gin.Engine) {
	s.engine = engine
}

func (s *Wire) Invoke(_, wireReq string) (wireRsp string) {
	request, err := http.ReadRequest(bufio.NewReader(strings.NewReader(wireReq)))
	if err != nil {
		panic(err)
	}

	var rspRecorder = httptest.NewRecorder()

	s.engine.ServeHTTP(rspRecorder, request)

	response := rspRecorder.Result()

	var buf = bytes.Buffer{}
	if err = response.Write(&buf); err != nil {
		panic(err)
	}

	wireRsp = buf.String()

	return
}

func (s *Wire) Close() {}
