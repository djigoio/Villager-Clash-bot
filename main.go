/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"net/http"
	"villagertournamentbot/crc"
	"villagertournamentbot/webhook"

	"github.com/gorilla/mux"
)

func main() {

	m := mux.NewRouter()

	println("Starting ACNH Clash server...")

	m.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Server is up and running")
	})

	//Listen to crc check and handle
	m.HandleFunc("/webhook/twitter", crc.Check).Methods("GET")
	//Listen to webhook event and handle
	m.HandleFunc("/twitter/webhook", webhook.Handler).Methods("POST")

	//Start Server
	server := &http.Server{
		Handler: m,
	}
	server.Addr = ":8080"
	server.ListenAndServe()

	webhook.Handler()
	crc.Check()
}
