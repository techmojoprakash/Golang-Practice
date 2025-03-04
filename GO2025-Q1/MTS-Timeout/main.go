package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const TIMESLEEP time.Duration = 40
const REQTIMEOUT time.Duration = 30

type E3DResponse struct {
	Data []*LegsData `json:"data"`
}

type LegsData struct {
	UniqueSelectionID string `json:"unique_selection_id"`
	FeedEventID       string `json:"feedEventID"`
	FeedMarketID      string `json:"feedMarketID"`
	Live              bool   `json:"live"`
	FeedProductID     string `json:"feedProductID"`
	FeedProviderID    string `json:"feedProviderID"`
	SelectionID       string `json:"selectionID"`
	FeedOutcomeID     string `json:"feedOutcomeID"`
	FeedSpecifiers    string `json:"specifiers"`
}

type Request struct {
	Legs []string `json:"legs"`
}

func getSelections(w http.ResponseWriter, r *http.Request) {
	timenow := time.Now()
	fmt.Println("Called....")

	// 1. Set a server-side timeout using context.WithTimeout
	ctx, cancel := context.WithTimeout(r.Context(), REQTIMEOUT*time.Second) // 5-second timeout
	defer cancel()                                                          // Important: Cancel the context to release resources

	// 2. Use a channel to signal completion (or error) of the processing
	done := make(chan error, 1) // Buffered channel to avoid blocking

	go func() {
		// Accept Request
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			done <- err // Send the error to the channel
			return
		}

		// Process the received data
		fmt.Println("Received legs:", req.Legs)

		legsStr := ""

		for item := range req.Legs {
			legsStr = legsStr + fmt.Sprint(item)
		}

		// Prepare response
		w.Header().Set("Content-Type", "application/json")
		requestID := r.Header.Get("X-Request-ID")

		response := E3DResponse{
			Data: make([]*LegsData, 0), // Initialize the Data slice to avoid nil pointer errors
		}

		leg1 := &LegsData{
			UniqueSelectionID: requestID,
			FeedEventID:       legsStr,
			FeedMarketID:      "market789",
			Live:              true,
			FeedProductID:     "",
			FeedProviderID:    "provider202",
			SelectionID:       "sel303",
			FeedOutcomeID:     "outcome404",
			FeedSpecifiers:    "spec505",
		}

		leg2 := &LegsData{
			UniqueSelectionID: requestID,
			FeedEventID:       legsStr,
			FeedMarketID:      "market123",
			Live:              false,
			FeedProductID:     "",
			FeedProviderID:    "provider303",
			SelectionID:       "sel404",
			FeedOutcomeID:     "outcome505",
			FeedSpecifiers:    "spec606",
		}

		response.Data = append(response.Data, leg1)
		response.Data = append(response.Data, leg2)
		fmt.Println("Sleeping....!")
		time.Sleep(TIMESLEEP * time.Second)

		// 3. Print the response (for demonstration)
		fmt.Println(response)

		json.NewEncoder(w).Encode(response)
		fmt.Println("RTT", time.Since(timenow))
		done <- nil // Signal successful completion

	}()

	// 3. Handle the timeout using select
	select {
	case <-ctx.Done(): // Timeout occurred
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
		return
	case err := <-done: // Processing finished (either successfully or with an error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Or appropriate error code
			return
		}
	}

}

func main() {
	fmt.Println("MTS Timeout.........!")
	var r = mux.NewRouter() // router by mux
	// routers
	r.HandleFunc("/betselections", getSelections).Methods("POST")
	fmt.Println("Starting server at port 8888...")
	log.Fatal(http.ListenAndServe(":8888", r))

}
