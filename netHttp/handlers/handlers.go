package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"webPractice1/cmd/errorPrinter"
	"webPractice1/netHttp"
	connect "webPractice1/postgresql"
	"webPractice1/postgresql/queries"
)

type HandlerAssetsResponse struct {
	netHttp.AssetsResponse
}

func NewHandlerAssetsResponse() *HandlerAssetsResponse {
	return &HandlerAssetsResponse{
		AssetsResponse: netHttp.AssetsResponse{
			Cache: make(map[*netHttp.AssetData]time.Time),
		},
	}
}

func (har *HandlerAssetsResponse) TaskHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/Abuseip/" {
		// Запрос направлен к "/Abuseip/", без идущего в конце ip.
		if req.Method == http.MethodPost || req.Method == http.MethodPut {
			har.CreateHandler(w, req)
		} else if req.Method == http.MethodGet {
			har.GetAllHandler(w, req)
		} else if req.Method == http.MethodDelete {
			har.DeleteAllHandler()
		} else {
			http.Error(w, fmt.Sprintf("expect method GET, DELETE or POST at /Abuseip/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	} else {
		// В запросе есть ip, выглядит он как "/Abuseip/ip".
		path := strings.Trim(req.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /Abuseip/<ip> in task handler", http.StatusBadRequest)
			return
		}
		ip := pathParts[1]

		if req.Method == http.MethodDelete {
			har.DeleteHandler(w, req, ip)
		} else if req.Method == http.MethodGet {
			har.GetHandler(w, ip)
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or DELETE at /Abuseip/<ip>, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (har *HandlerAssetsResponse) GetAllHandler(w http.ResponseWriter, req *http.Request) {
	connect := connect.PostgresqlConnect()
	jsonData, err := json.Marshal(queries.GetEntitys(connect))
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	w.Write(jsonData)
}
func (har *HandlerAssetsResponse) CreateHandler(w http.ResponseWriter, req *http.Request) {
	har.Mu.Lock()
	defer har.Mu.Unlock()

	ttl := time.Second * 15
	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var asset netHttp.AssetData
	if err = json.Unmarshal(reqBytes, &asset); err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if asset.IPAddress == "" || asset.IPVersion == 0 {
		http.Error(w, "ipAddress or ipVersion not set", http.StatusBadRequest)
		return
	}

	har.Cache[&asset] = time.Now().Add(ttl)

	go func() {
		time.Sleep(ttl)
		har.Mu.Lock()
		defer har.Mu.Unlock()
		delete(har.Cache, &asset)
	}()
	queries.AddEntity(connect.PostgresqlConnect(), asset)
}

func (har *HandlerAssetsResponse) DeleteAllHandler() {
	har.Mu.Lock()
	defer har.Mu.Unlock()
	for v := range har.Cache {
		delete(har.Cache, v)
	}
	queries.DeleteAllEntitiesDB(connect.PostgresqlConnect())
}

func (har *HandlerAssetsResponse) DeleteHandler(w http.ResponseWriter, req *http.Request, ip string) {
	har.Mu.Lock()
	defer har.Mu.Unlock()
	for v := range har.Cache {
		if v.IPAddress == ip {
			delete(har.Cache, v)
		}
	}
	queries.DeleteEntityDB(connect.PostgresqlConnect(), ip)
}

func (har *HandlerAssetsResponse) GetHandler(w http.ResponseWriter, ip string) {
	har.Mu.Lock()
	defer har.Mu.Unlock()
	for v := range har.Cache { //если нет записей то range = 0
		if v.IPAddress == ip {
			v.IsCache = true
			v.IsDb = false // Нужно ли явное указание?
			jsonData, err := json.Marshal(v)
			if err != nil {
				errorPrinter.PrintCallerFunctionName(err)
				return
			}
			w.Write(jsonData)
			return
		}
	}
	jsonData, err := json.Marshal(queries.GetEntity(connect.PostgresqlConnect(), ip))
	if err != nil {
		errorPrinter.PrintCallerFunctionName(err)
		return
	}
	w.Write(jsonData)
}
