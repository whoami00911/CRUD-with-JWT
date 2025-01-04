package handlers

import (
	"encoding/json"
	"io"
	"time"
	"webPractice1/internal/DBqueries"
	handler "webPractice1/internal/transport"
	"webPractice1/pkg/errorPrinter"
	connect "webPractice1/pkg/postgresql"

	"github.com/gin-gonic/gin"
)

type HandlerAssetsResponse struct {
	handler.AssetsResponse
}

func NewHandlerAssetsResponse() *HandlerAssetsResponse {
	return &HandlerAssetsResponse{
		AssetsResponse: handler.AssetsResponse{
			Cache: make(map[*handler.AssetData]time.Time),
		},
	}
}

func (har *HandlerAssetsResponse) GetAllHandler(c *gin.Context) {
	connect := connect.PostgresqlConnect()
	jsonData, err := json.Marshal(DBqueries.GetEntitys(connect))
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, jsonData)
}
func (har *HandlerAssetsResponse) CreateHandler(c *gin.Context) {
	har.Mu.Lock()
	defer har.Mu.Unlock()

	ttl := time.Second * 15
	reqBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	var asset handler.AssetData
	if err = json.Unmarshal(reqBytes, &asset); err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	if asset.IPAddress == "" || asset.IPVersion == 0 {
		c.JSON(400, gin.H{"error": "ipAddress or ipVersion not set"})
		return
	}

	har.Cache[&asset] = time.Now().Add(ttl)

	go func() {
		time.Sleep(ttl)
		har.Mu.Lock()
		defer har.Mu.Unlock()
		delete(har.Cache, &asset)
	}()
	DBqueries.AddEntity(connect.PostgresqlConnect(), asset)
}

func (har *HandlerAssetsResponse) DeleteAllHandler(c *gin.Context) {
	har.Mu.Lock()
	defer har.Mu.Unlock()
	for v := range har.Cache {
		delete(har.Cache, v)
	}
	DBqueries.DeleteAllEntitiesDB(connect.PostgresqlConnect())
}

func (har *HandlerAssetsResponse) DeleteHandler(c *gin.Context) {
	ip := c.Param("ip")
	har.Mu.Lock()
	defer har.Mu.Unlock()
	for v := range har.Cache {
		if v.IPAddress == ip {
			delete(har.Cache, v)
		}
	}
	DBqueries.DeleteEntityDB(connect.PostgresqlConnect(), ip)
}

func (har *HandlerAssetsResponse) GetHandler(c *gin.Context) {
	ip := c.Param("ip")
	har.Mu.Lock()
	defer har.Mu.Unlock()
	for v := range har.Cache { //если нет записей то range = 0
		if v.IPAddress == ip {
			v.IsCache = true
			v.IsDb = false // Нужно ли явное указание?
			jsonData, err := json.Marshal(v)
			if err != nil {
				errorPrinter.PrintCallerFunctionName(err)
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(200, jsonData)
			return
		}
	}
	jsonData, err := json.Marshal(DBqueries.GetEntity(connect.PostgresqlConnect(), ip))
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, jsonData)
}
