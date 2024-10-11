package controller

import (
	"encoding/json"
	"login-user/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

// Criar usuario comtroller
func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var userinput service.UserInput

	err := json.NewDecoder(r.Body).Decode(&userinput)

	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = uc.UserService.Register(r.Context(), userinput)

	if err != nil {
		http.Error(w, "Erro ao cadastrar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"messege": "usuario criado com sucesso"})

}

// Login usuario controller

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	var userinput service.UserInput
	err := json.NewDecoder(r.Body).Decode(&userinput)

	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	token, err := uc.UserService.Login(r.Context(), userinput)

	if err != nil {
		http.Error(w, "Falha ao autenticar usuário", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})

}
