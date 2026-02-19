package main_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { // serve a simple message
	fmt.Fprintln(w, "Hi there")

}

func createUser(w http.ResponseWriter, r *http.Request) { // create a new user in local cache
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Age <= 0 {
		http.Error(w, "Name is required and age must be positive", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()

	userList[len(userList)+1] = user

	defer cacheMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func getUser(w http.ResponseWriter, r *http.Request) { // get user by ID
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	cacheMutex.RLock()
	user, ok := userList[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) //return the user data as JSON
	w.Write(j)

	cacheMutex.RUnlock()

}

func deleteUser(w http.ResponseWriter, r *http.Request) { // delete user by ID
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	cacheMutex.Lock()

	if _, ok := userList[id]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(userList, id)
	cacheMutex.Unlock()
	w.WriteHeader(http.StatusNoContent)
}
