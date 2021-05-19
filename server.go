package main

import (
	"context"
	"net/http"
	"sync"
)

type connection struct {
	writer     http.ResponseWriter
	flusher    http.Flusher
	requestCtx context.Context
}

// Server - instance used to accept connections and store them
type Server struct {
	connections      map[string]*connection
	connectionsMutex sync.RWMutex
}

// NewServer - creates new Server instance
func NewServer() *Server {
	server := &Server{
		connections: map[string]*connection{},
	}

	return server
}

// ServeHTTP - implementation of http requests acceptance
func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	requestContext := req.Context()
	s.connectionsMutex.Lock()
	s.connections[req.RemoteAddr] = &connection{
		writer:     rw,
		flusher:    flusher,
		requestCtx: requestContext,
	}
	s.connectionsMutex.Unlock()

	defer func() {
		s.removeConnection(req.RemoteAddr)
	}()

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	<-requestContext.Done()
}

// Send - sends message to all connected clients
func (s *Server) Send(msg string) {
	s.connectionsMutex.RLock()
	defer s.connectionsMutex.RUnlock()

	msgBytes := []byte("event: message\n\ndata:" + msg + "\n\n")
	for client, connection := range s.connections {
		_, err := connection.writer.Write(msgBytes)
		if err != nil {
			s.removeConnection(client)
			continue
		}

		connection.flusher.Flush()
	}
}

func (s *Server) removeConnection(client string) {
	s.connectionsMutex.Lock()
	defer s.connectionsMutex.Unlock()

	delete(s.connections, client)
}
