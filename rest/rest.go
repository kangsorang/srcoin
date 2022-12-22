package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kangsorang/srcoin/blockchain"
	"github.com/kangsorang/srcoin/utils"
)

type url string
type UrlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload"`
}
type AddBlockBody struct {
	Message string `json:"message"`
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []UrlDescription{
		{
			URL:         url("/"),
			Method:      http.MethodGet,
			Description: "See documentation",
		},
		{
			URL:         url("/status"),
			Method:      http.MethodGet,
			Description: "See the status of blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      http.MethodGet,
			Description: "See All blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      http.MethodPost,
			Description: "Add a block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      http.MethodGet,
			Description: "See a block",
		},
	}
	err := json.NewEncoder(w).Encode(data)
	utils.HandleErr(err)
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func status(w http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(w).Encode(blockchain.Blockchain()))
}

func blocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		utils.HandleErr(json.NewEncoder(w).Encode(blockchain.Blockchain().GetBlocks()))
	case http.MethodPost:
		var data AddBlockBody
		err := json.NewDecoder(r.Body).Decode(&data)
		utils.HandleErr(err)
		blockchain.Blockchain().AddBlock(data.Message)
		w.WriteHeader(http.StatusCreated)
	}
}

func Start(aPort int) {
	port := fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation)
	router.HandleFunc("/status", status)
	router.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
