package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"restapi/pkg/storage"
	"restapi/pkg/user"
)

type contextKey string

func main() {
	srv := storage.Service{Store: make(map[int]*user.User), FriendsStore: make(map[int]map[int]*user.User), Counter: 0}
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/users", func(r chi.Router) {
		r.Use(makeStorageCtx(&srv))
		r.Get("/", getAllUsers)
		r.Post("/", createUser)

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(makeUserCtx(&srv))
			r.Get("/", getUser)
			r.Put("/", updateUser)
			r.Delete("/", deleteUser)

			r.Route("/friends", func(r chi.Router) {
				r.Get("/", getUserFriends)
			})
		})
	})

	r.Route("/make_friends", func(r chi.Router) {
		r.Use(makeStorageCtx(&srv))
		r.Post("/", makeFriends)
	})

	http.ListenAndServe(":3333", r)
}

func makeStorageCtx(s *storage.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), contextKey("service"), s)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func makeUserCtx(s *storage.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
			if err != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}
			user, err := s.GetUser(userID)
			if err != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}
			ctx := context.WithValue(r.Context(), contextKey("user"), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	srv, ok := ctx.Value(contextKey("service")).(*storage.Service)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	resp := make([]user.User, len(srv.Store))
	i := 0
	for _, user := range srv.Store {
		resp[i] = *user
		i++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	srv, ok := ctx.Value(contextKey("service")).(*storage.Service)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	var newUser user.User
	if err := json.Unmarshal(content, &newUser); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	srv.AddUser(&newUser)

	var userFriends user.UserFriends
	if err := json.Unmarshal(content, &userFriends); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	for _, userID := range userFriends.Friends {
		srv.MakeFriends(newUser.Id, userID)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("user id: %d\n", newUser.Id)))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(contextKey("user")).(*user.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	var newUser user.User
	if err := json.Unmarshal(content, &newUser); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	ctx := r.Context()
	srv := ctx.Value(contextKey("service")).(*storage.Service)
	contextUser, ok := ctx.Value(contextKey("user")).(*user.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	oldUser := srv.Store[contextUser.Id]
	srv.Store[oldUser.Id] = &newUser

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln("Возраст пользователя успешно обновлён")))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	srv := ctx.Value(contextKey("service")).(*storage.Service)
	user := ctx.Value(contextKey("user")).(*user.User)
	friends, err := srv.GetFriends(user.Id)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	for _, u := range friends {
		srv.DeleteFriend(u.Id, user.Id)
	}
	srv.DeleteUser(user.Id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Пользователь с именем: \"%s\" был удален\n", user.Name)))
}

func getUserFriends(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	srv := ctx.Value(contextKey("service")).(*storage.Service)
	curUser, ok := ctx.Value(contextKey("user")).(*user.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	friends, err := srv.GetFriends(curUser.Id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	userFriends := make([]*user.User, len(friends))
	i := 0
	for _, u := range friends {
		userFriends[i] = u
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userFriends)
}

func makeFriends(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	var friendsReq user.FriendsRequest
	if err := json.Unmarshal(content, &friendsReq); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	ctx := r.Context()
	srv := ctx.Value(contextKey("service")).(*storage.Service)
	if err := srv.MakeFriends(friendsReq.SourceId, friendsReq.TargetId); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	user1 := srv.Store[friendsReq.SourceId]
	user2 := srv.Store[friendsReq.TargetId]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("\"%s\" и \"%s\" теперь друзья\n", user1.Name, user2.Name)))
}
