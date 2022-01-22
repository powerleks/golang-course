package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"restapi/internal/entity"
	"restapi/internal/usecase"
)

type userRoutes struct {
	u usecase.User
}

func newUserRoutes(r *chi.Mux, t usecase.User) {
	ur := &userRoutes{t}

	r.Route("/users", func(r chi.Router) {
		r.Get("/", ur.getAllUsers)
		r.Post("/", ur.createUser)

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", ur.getUser)
			r.Put("/", ur.updateUser)
			r.Delete("/", ur.deleteUser)

			r.Route("/friends", func(r chi.Router) {
				r.Get("/", ur.getUserFriends)
			})
		})
	})

	r.Route("/make_friends", func(r chi.Router) {
		r.Post("/", ur.makeFriends)
	})
}

func (ur *userRoutes) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ur.u.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (ur *userRoutes) createUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	var newUser entity.User
	if err := json.Unmarshal(content, &newUser); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	var userFriends entity.UserFriends
	if err := json.Unmarshal(content, &userFriends); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	userId, err := ur.u.AddUser(r.Context(), &newUser, &userFriends)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("user id: %d\n", userId)))
}

func (ur *userRoutes) getUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return		
	}
	user, err := ur.u.GetUserById(r.Context(), userId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (ur *userRoutes) updateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	userId, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return		
	}

	var user entity.User
	if err := json.Unmarshal(content, &user); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}
	err = ur.u.UpdateUser(r.Context(), userId, &user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln("Возраст пользователя успешно обновлён")))
}

func (ur *userRoutes) deleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return		
	}

	err = ur.u.DeleteUser(r.Context(), userId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Пользователь с ID=\"%d\" был удален\n", userId)))
}

func (ur *userRoutes) getUserFriends(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return		
	}
	friends, err := ur.u.GetUserFriends(r.Context(), userId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friends)
}

func (ur *userRoutes) makeFriends(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	var friendsReq entity.FriendsRequest
	if err := json.Unmarshal(content, &friendsReq); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte(err.Error()))
		return
	}

	err = ur.u.MakeFriends(r.Context(), friendsReq.SourceId, friendsReq.TargetId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Пользователи с id \"%d\" и \"%d\" теперь друзья\n", friendsReq.SourceId, friendsReq.TargetId)))
}