package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/JhonCamargo53/prueba-tecnica/internal/application/services/user"
	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"github.com/gorilla/mux"
)

var userService *user.UserService

func InitUserHandler(us *user.UserService) {
	userService = us
}

// CreateUserRequest representa el cuerpo de la solicitud para crear un usuario
// @Description Datos para crear un nuevo usuario
type CreateUserRequest struct {
	Name     string `json:"name" example:"Juan Pérez"`
	Email    string `json:"email" example:"juan.perez@example.com"`
	Password string `json:"password" example:"contraseña123"`
	RoleId   uint   `json:"roleId" example:"1"`
}

// UpdateUserRequest representa el cuerpo de la solicitud para actualizar un usuario
// @Description Datos para actualizar un usuario existente
type UpdateUserRequest struct {
	Name     string `json:"name" example:"Juan Pérez Actualizado"`
	Email    string `json:"email" example:"juan.perez.updated@example.com"`
	Password string `json:"password" example:"nuevacontraseña123"`
	RoleId   uint   `json:"roleId" example:"2"`
}

// GetUsersHandle godoc
// @Summary      Obtener todos los usuarios
// @Description  Retorna una lista de todos los usuarios del sistema
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} models.User "Lista de usuarios"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /users [get]
func GetUsersHandle(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetAllUsers()
	if err != nil {
		http.Error(w, "No se pudieron obtener los usuarios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserHandle godoc
// @Summary      Obtener un usuario por ID
// @Description  Retorna los detalles de un usuario específico
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del usuario"
// @Success      200 {object} models.User "Usuario encontrado"
// @Failure      400 {string} string "ID inválido"
// @Failure      404 {string} string "Usuario no encontrado"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /users/{id} [get]
func GetUserHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	id, err := strconv.Atoi(userId)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := userService.GetUserByID(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al modificar usuario: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// PostUserHandle godoc
// @Summary      Crear un nuevo usuario
// @Description  Crea un nuevo usuario en el sistema
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CreateUserRequest true "Datos del usuario a crear"
// @Success      201 {object} models.User "Usuario creado exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      409 {string} string "El usuario ya existe"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /users [post]
func PostUserHandle(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error leyendo el body", http.StatusBadRequest)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var userData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleId   uint   `json:"roleId"`
	}

	err = json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if userData.Name == "" || userData.Email == "" || userData.Password == "" || userData.RoleId == 0 {
		http.Error(w, "Los campos Nombre, Correo, Contraseña y Rango son obligatorios", http.StatusBadRequest)
		return
	}
	user := models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: userData.Password,
		RoleId:   userData.RoleId,
	}

	createdUser, err := userService.CreateUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "ya existe un usuario") {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, "Error al crear usuario: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	createdUser.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// UpdateUserHandle godoc
// @Summary      Actualizar un usuario
// @Description  Actualiza los datos de un usuario existente
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del usuario"
// @Param        request body UpdateUserRequest true "Datos actualizados del usuario"
// @Success      200 {object} models.User "Usuario actualizado exitosamente"
// @Failure      400 {string} string "Solicitud inválida"
// @Failure      404 {string} string "Usuario no encontrado"
// @Failure      409 {string} string "El email ya existe"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /users/{id} [put]
func UpdateUserHandle(w http.ResponseWriter, r *http.Request) {

	requesterId := r.Context().Value("requesterId").(uint)

	params := mux.Vars(r)
	userId := params["id"]
	id, err := strconv.Atoi(userId)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var userData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleId   uint   `json:"roleId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	userDataModel := models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: userData.Password,
		RoleId:   userData.RoleId,
	}

	updatedUser, err := userService.UpdateUser(uint(id), &userDataModel, requesterId)

	if err != nil {
		if strings.Contains(err.Error(), "ya existe") {
			http.Error(w, err.Error(), http.StatusConflict)
		} else if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al modificar usuario: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	updatedUser.Password = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUserHandle godoc
// @Summary      Eliminar un usuario
// @Description  Elimina un usuario del sistema
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID del usuario"
// @Success      204 "Usuario eliminado exitosamente"
// @Failure      400 {string} string "ID inválido"
// @Failure      404 {string} string "Usuario no encontrado"
// @Failure      500 {string} string "Error interno del servidor"
// @Router       /users/{id} [delete]
func DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = userService.DeleteUser(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al modificar usuario: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
