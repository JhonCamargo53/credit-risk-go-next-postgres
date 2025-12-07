package user

import (
	"errors"
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

/* GetAllUsers */

func TestGetAllUsers_Exitoso(t *testing.T) {
	user1 := &models.User{ID: 1, Name: "Juan", Email: "juan@example.com", RoleId: 1}
	user2 := &models.User{ID: 2, Name: "Jhon", Email: "jhon@example.com", RoleId: 2}

	userRepo := NewMockUserRepository([]*models.User{user1, user2})
	roleRepo := NewMockRoleRepository(nil)

	service := NewUserService(userRepo, roleRepo)

	users, err := service.GetAllUsers()
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if len(users) != 2 {
		t.Fatalf("se esperaban 2 usuarios, se obtuvo=%d", len(users))
	}
}

func TestGetAllUsers_ErrorRepositorio(t *testing.T) {
	userRepo := NewMockUserRepository(nil)
	userRepo.ErrFindAllOrdered = errors.New("falló la BD")
	roleRepo := NewMockRoleRepository(nil)

	service := NewUserService(userRepo, roleRepo)

	users, err := service.GetAllUsers()
	if err == nil {
		t.Fatalf("se esperaba error del repositorio, se obtuvo nil")
	}
	if users != nil {
		t.Fatalf("se esperaba lista nil cuando hay error")
	}
}

/* GetUserByID */

func TestGetUserByID_Exitoso(t *testing.T) {
	existing := &models.User{ID: 10, Name: "Jhon", Email: "jhon@example.com", RoleId: 1}

	userRepo := NewMockUserRepository([]*models.User{existing})
	roleRepo := NewMockRoleRepository(nil)

	service := NewUserService(userRepo, roleRepo)

	user, err := service.GetUserByID(10)
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if user == nil || user.ID != 10 {
		t.Fatalf("se esperaba el usuario con ID 10")
	}
}

func TestGetUserByID_NoExiste(t *testing.T) {
	userRepo := NewMockUserRepository(nil)
	roleRepo := NewMockRoleRepository(nil)

	service := NewUserService(userRepo, roleRepo)

	user, err := service.GetUserByID(99)
	if err == nil {
		t.Fatalf("se esperaba error porque el usuario no existe")
	}
	if user != nil {
		t.Fatalf("no se esperaba usuario")
	}
}

/* CreateUser */

func TestCreateUser_RolNoExiste(t *testing.T) {
	userRepo := NewMockUserRepository(nil)
	roleRepo := NewMockRoleRepository(nil) // ningún rol

	service := NewUserService(userRepo, roleRepo)

	newUser := &models.User{
		Name:   "Juan",
		Email:  "juan@example.com",
		RoleId: 99, // no existe
	}

	created, err := service.CreateUser(newUser)
	if err == nil {
		t.Fatalf("se esperaba error porque el rol no existe")
	}
	if created != nil {
		t.Fatalf("no se debería crear el usuario")
	}
}

func TestCreateUser_EmailDuplicado(t *testing.T) {
	existing := &models.User{
		ID:     1,
		Name:   "Lina",
		Email:  "lina@example.com",
		RoleId: 1,
	}

	userRepo := NewMockUserRepository([]*models.User{existing})
	roleRepo := NewMockRoleRepository([]*models.Role{
		{ID: 1, Name: "ADMIN"},
	})

	service := NewUserService(userRepo, roleRepo)

	newUser := &models.User{
		Name:   "Jhon",
		Email:  "lina@example.com", // mismo correo que Lina
		RoleId: 1,
	}

	created, err := service.CreateUser(newUser)
	if err == nil {
		t.Fatalf("se esperaba error por email duplicado")
	}
	if created != nil {
		t.Fatalf("no se debería crear el usuario con email duplicado")
	}
}

func TestCreateUser_Exitoso_HashPassword(t *testing.T) {
	userRepo := NewMockUserRepository(nil)
	roleRepo := NewMockRoleRepository([]*models.Role{
		{ID: 1, Name: "ADMIN"},
	})

	service := NewUserService(userRepo, roleRepo)

	password := "my-password"

	newUser := &models.User{
		ID:       1,
		Name:     "Lina",
		Email:    "lina@example.com",
		RoleId:   1,
		Password: password,
	}

	created, err := service.CreateUser(newUser)
	if err != nil {
		t.Fatalf("no se esperaba error al crear usuario: %v", err)
	}
	if created == nil {
		t.Fatalf("se esperaba usuario creado")
	}

	// Validar que el password fue hasheado
	if created.Password == password {
		t.Fatalf("se esperaba que el password fuera hasheado, pero se mantuvo en texto plano")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(created.Password), []byte(password)); err != nil {
		t.Fatalf("el hash almacenado no valida contra el password original: %v", err)
	}

	// Verificar que quedó en el repo
	if _, ok := userRepo.UsersByID[1]; !ok {
		t.Fatalf("usuario no fue guardado en el repositorio")
	}
}

/* UpdateUser */

func TestUpdateUser_NoPuedeActualizarseASiMismo(t *testing.T) {
	existing := &models.User{
		ID:     1,
		Name:   "Juan",
		Email:  "juan@example.com",
		RoleId: 1,
	}

	userRepo := NewMockUserRepository([]*models.User{existing})
	roleRepo := NewMockRoleRepository([]*models.Role{
		{ID: 1, Name: "ADMIN"},
	})

	service := NewUserService(userRepo, roleRepo)

	updateData := &models.User{
		Name: "Juan Actualizado",
	}

	updated, err := service.UpdateUser(1, updateData, 1) // requesterId == id
	if err == nil {
		t.Fatalf("se esperaba error porque un usuario no puede actualizarse a sí mismo")
	}
	if updated != nil {
		t.Fatalf("no se debería actualizar el usuario cuando se auto-actualiza")
	}
}

func TestUpdateUser_RolNoExiste(t *testing.T) {
	existing := &models.User{
		ID:     1,
		Name:   "Jhon",
		Email:  "jhon@example.com",
		RoleId: 1,
	}

	userRepo := NewMockUserRepository([]*models.User{existing})
	roleRepo := NewMockRoleRepository(nil) // ningún rol

	service := NewUserService(userRepo, roleRepo)

	updateData := &models.User{
		RoleId: 99, // no existe
	}

	updated, err := service.UpdateUser(1, updateData, 2)
	if err == nil {
		t.Fatalf("se esperaba error porque el rol no existe")
	}
	if updated != nil {
		t.Fatalf("no se debería actualizar el usuario")
	}
}

func TestUpdateUser_EmailDuplicado(t *testing.T) {
	user1 := &models.User{
		ID:     1,
		Name:   "Juan",
		Email:  "juan@example.com",
		RoleId: 1,
	}
	user2 := &models.User{
		ID:     2,
		Name:   "Lina",
		Email:  "lina@example.com",
		RoleId: 2,
	}

	userRepo := NewMockUserRepository([]*models.User{user1, user2})
	roleRepo := NewMockRoleRepository([]*models.Role{
		{ID: 1, Name: "ADMIN"},
		{ID: 2, Name: "USER"},
	})

	service := NewUserService(userRepo, roleRepo)

	updateData := &models.User{
		Email: "lina@example.com", // ya tomado por Lina
	}

	updated, err := service.UpdateUser(1, updateData, 3)
	if err == nil {
		t.Fatalf("se esperaba error por email duplicado en update")
	}
	if updated != nil {
		t.Fatalf("no se debería actualizar el usuario si el email está duplicado")
	}
}

func TestUpdateUser_Exitoso(t *testing.T) {
	existing := &models.User{
		ID:       1,
		Name:     "Jhon Viejo",
		Email:    "jhon.viejo@example.com",
		RoleId:   1,
		Password: "",
	}

	userRepo := NewMockUserRepository([]*models.User{existing})
	roleRepo := NewMockRoleRepository([]*models.Role{
		{ID: 1, Name: "ADMIN"},
		{ID: 2, Name: "USER"},
	})

	service := NewUserService(userRepo, roleRepo)

	newPassword := "nuevo-password"

	updateData := &models.User{
		Name:     "Jhon Nuevo",
		Email:    "jhon.nuevo@example.com",
		RoleId:   2,
		Password: newPassword,
	}

	updated, err := service.UpdateUser(1, updateData, 99)
	if err != nil {
		t.Fatalf("no se esperaba error al actualizar: %v", err)
	}
	if updated == nil {
		t.Fatalf("se esperaba usuario actualizado")
	}

	if updated.Name != "Jhon Nuevo" {
		t.Fatalf("nombre no actualizado correctamente, se obtuvo=%s", updated.Name)
	}
	if updated.Email != "jhon.nuevo@example.com" {
		t.Fatalf("email no actualizado correctamente, se obtuvo=%s", updated.Email)
	}
	if updated.RoleId != 2 {
		t.Fatalf("roleId no actualizado correctamente, se obtuvo=%d", updated.RoleId)
	}

	// Password debe estar hasheado
	if updated.Password == newPassword {
		t.Fatalf("se esperaba password hasheado, no en texto plano")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(updated.Password), []byte(newPassword)); err != nil {
		t.Fatalf("el hash almacenado no valida contra el nuevo password: %v", err)
	}
}

/*   DeleteUser   */

func TestDeleteUser_UsuarioNoExiste(t *testing.T) {
	userRepo := NewMockUserRepository(nil)
	roleRepo := NewMockRoleRepository(nil)

	service := NewUserService(userRepo, roleRepo)

	err := service.DeleteUser(99)
	if err == nil {
		t.Fatalf("se esperaba error porque el usuario no existe")
	}
}

func TestDeleteUser_Exitoso(t *testing.T) {
	existing := &models.User{
		ID:     1,
		Name:   "Lina",
		Email:  "lina@example.com",
		RoleId: 1,
	}

	userRepo := NewMockUserRepository([]*models.User{existing})
	roleRepo := NewMockRoleRepository(nil)

	service := NewUserService(userRepo, roleRepo)

	err := service.DeleteUser(1)
	if err != nil {
		t.Fatalf("no se esperaba error al eliminar usuario: %v", err)
	}

	if _, ok := userRepo.UsersByID[1]; ok {
		t.Fatalf("el usuario debería haberse eliminado del repositorio")
	}
}
