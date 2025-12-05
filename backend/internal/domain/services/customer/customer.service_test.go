package customer

import (
	"testing"

	"github.com/JhonCamargo53/prueba-tecnica/internal/domain/models"
)

/* Test CreateCustomer */

func TestCreateCustomer_EmailDuplicado(t *testing.T) {

	// Cliente ya existente con ese email
	existing := &models.Customer{
		Name:           "Juan",
		Email:          "juan@example.com",
		DocumentNumber: "123",
		DocumentTypeId: 1,
	}
	existing.ID = 1

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})
	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{1: true},
	}
	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	newCustomer := &models.Customer{
		Name:           "Pedro",
		Email:          "juan@example.com", // Correo repetido (juan@example.com)
		DocumentNumber: "456",
		DocumentTypeId: 1,
	}

	created, err := service.CreateCustomer(newCustomer)

	if err == nil {
		t.Fatalf("se esperaba error por email duplicado, pero err es nil")
	}
	if created != nil {
		t.Fatalf("no se debería crear el cliente cuando el email está duplicado")
	}
}

func TestCreateCustomer_DocumentoDuplicado(t *testing.T) {

	existing := &models.Customer{
		Name:           "Juan",
		Email:          "juan@example.com",
		DocumentNumber: "1002034346",
		DocumentTypeId: 2,
	}
	existing.ID = 1

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})
	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{2: true},
	}
	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	newCustomer := &models.Customer{
		Name:           "Luis",
		Email:          "luis@example.com",
		DocumentNumber: "1002034346", // mismo número
		DocumentTypeId: 2,            // mismo tipo
	}

	created, err := service.CreateCustomer(newCustomer)

	if err == nil {
		t.Fatalf("se esperaba error por documento duplicado, pero err es nil")
	}
	if created != nil {
		t.Fatalf("no se debería crear el cliente cuando el documento está duplicado")
	}
}

func TestCreateCustomer_Exitoso(t *testing.T) {

	existing := &models.Customer{
		Name:           "Juan",
		Email:          "juan@example.com",
		DocumentNumber: "1002034346",
		DocumentTypeId: 1,
	}
	existing.ID = 1

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})
	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{1: true},
	}
	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	newCustomer := &models.Customer{
		Name:           "Jhon Camargo",
		Email:          "jhoncamargo@example.com",
		DocumentNumber: "1002322247",
		DocumentTypeId: 1,
	}

	created, err := service.CreateCustomer(newCustomer)

	if err != nil {
		t.Fatalf("no se esperaba error al crear cliente válido, err: %v", err)
	}
	if created == nil {
		t.Fatalf("se esperaba un cliente creado, pero es nil")
	}
	if created.ID == 0 {
		t.Fatalf("se esperaba que el cliente tuviera un ID asignado")
	}

	if len(customerRepo.Customers) != 2 {
		t.Fatalf("se esperaba 2 cliente en el repo, hay: %d", len(customerRepo.Customers))
	}
}

/* Tests de GetCustomerByID */

func TestGetCustomerByID_NoExiste(t *testing.T) {

	customerRepo := NewMockCustomerRepository(nil)
	documentTypeRepo := &MockDocumentTypeRepository{ExistingIDs: map[uint]bool{}}
	creditRequestRepo := &MockCreditRequestRepository{HasRequests: map[uint]bool{}}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	c, err := service.GetCustomerByID(100)

	if err == nil {
		t.Fatalf("se esperaba error porque el cliente no existe")
	}
	if c != nil {
		t.Fatalf("no se esperaba ningún cliente, pero se obtuvo uno")
	}
}

func TestGetCustomerByID_Exitoso(t *testing.T) {
	existing := &models.Customer{
		Name:           "Juan",
		Email:          "juan@example.com",
		DocumentNumber: "123",
		DocumentTypeId: 7,
	}
	existing.ID = 7

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})
	documentTypeRepo := &MockDocumentTypeRepository{ExistingIDs: map[uint]bool{1: true}}
	creditRequestRepo := &MockCreditRequestRepository{HasRequests: map[uint]bool{}}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	c, err := service.GetCustomerByID(7)

	if err != nil {
		t.Fatalf("no se esperaba error, err: %v", err)
	}
	if c == nil {
		t.Fatalf("se esperaba un cliente, se obtuvo nil")
	}
	if c.Email != "juan@example.com" {
		t.Fatalf("email inesperado, se obtuvo el correo=%s", c.Email)
	}
}

/* Tests de UpdateCustomer */

func TestUpdateCustomer_DocumentTypeNoExiste(t *testing.T) {
	existing := &models.Customer{
		Name:           "Cliente",
		Email:          "cliente@example.com",
		DocumentNumber: "123",
		DocumentTypeId: 1,
	}
	existing.ID = 1

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})

	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{1: true},
	}
	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	updateData := &models.Customer{
		DocumentTypeId: 99,
	}

	updated, err := service.UpdateCustomer(existing.ID, updateData)

	if err == nil {
		t.Fatalf("se esperaba error por tipo de documento inexistente, pero err es nil")
	}
	if updated != nil {
		t.Fatalf("no se debería actualizar el cliente cuando el tipo de documento no existe")
	}
}

func TestUpdateCustomer_EmailDuplicado(t *testing.T) {
	existing1 := &models.Customer{
		Name:           "Cliente1",
		Email:          "c1@example.com",
		DocumentNumber: "111",
		DocumentTypeId: 1,
	}
	existing1.ID = 1

	existing2 := &models.Customer{
		Name:           "Juan",
		Email:          "juan@example.com",
		DocumentNumber: "222",
		DocumentTypeId: 1,
	}
	existing2.ID = 2

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing1, existing2})
	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{1: true},
	}
	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	updateData := &models.Customer{
		Email: "juan@example.com",
	}

	updated, err := service.UpdateCustomer(1, updateData)

	if err == nil {
		t.Fatalf("se esperaba error por email duplicado en update, pero err es nil")
	}
	if updated != nil {
		t.Fatalf("no se debería actualizar el cliente cuando el email está duplicado")
	}
}

func TestUpdateCustomer_Exitoso(t *testing.T) {
	existing := &models.Customer{
		Name:           "Juan",
		Email:          "juan@example.com",
		DocumentNumber: "123",
		DocumentTypeId: 1,
	}
	existing.ID = 1

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})
	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{1: true, 2: true},
	}
	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	updateData := &models.Customer{
		Name:           "Juan Nuevo",
		Email:          "juanuevo@example.com",
		DocumentNumber: "999",
		DocumentTypeId: 2,
	}

	updated, err := service.UpdateCustomer(1, updateData)

	if err != nil {
		t.Fatalf("no se esperaba error al actualizar cliente válido, err: %v", err)
	}
	if updated == nil {
		t.Fatalf("se esperaba cliente actualizado, se obtuvo nil")
	}
	if updated.Name != "Juan Nuevo" {
		t.Fatalf("nombre inesperado, obtuvo=%s", updated.Name)
	}
	if updated.Email != "juanuevo@example.com" {
		t.Fatalf("email inesperado, se obtuvo=%s", updated.Email)
	}
	if updated.DocumentNumber != "999" {
		t.Fatalf("documento inesperado, obtuvo=%s", updated.DocumentNumber)
	}
	if updated.DocumentTypeId != 2 {
		t.Fatalf("documentTypeId inesperado, obtuvo=%d", updated.DocumentTypeId)
	}
}

/* Tests de DeleteCustomer */

func TestDeleteCustomer_ConSolicitudesAsociadas(t *testing.T) {
	existing := &models.Customer{
		Name:           "Juan Solicitudes",
		Email:          "juansolicitudes@example.com",
		DocumentNumber: "123",
		DocumentTypeId: 1,
	}
	existing.ID = 1

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})

	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{
			1: true, // este cliente tiene solicitudes asociadas
		},
	}
	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{1: true},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	err := service.DeleteCustomer(existing.ID)

	if err == nil {
		t.Fatalf("se esperaba error porque el cliente tiene solicitudes asociadas")
	}

	if _, ok := customerRepo.Customers[existing.ID]; !ok {
		t.Fatalf("el cliente no debería haberse eliminado del repositorio")
	}
}

func TestDeleteCustomer_Exitoso(t *testing.T) {
	existing := &models.Customer{
		Name:           "Juan sin solicitudes",
		Email:          "juansinsolicitudes@example.com",
		DocumentNumber: "123",
		DocumentTypeId: 1,
	}
	existing.ID = 1

	customerRepo := NewMockCustomerRepository([]*models.Customer{existing})

	creditRequestRepo := &MockCreditRequestRepository{
		HasRequests: map[uint]bool{
			1: false,
		},
	}
	documentTypeRepo := &MockDocumentTypeRepository{
		ExistingIDs: map[uint]bool{1: true},
	}

	service := NewCustomerService(customerRepo, documentTypeRepo, creditRequestRepo)

	err := service.DeleteCustomer(existing.ID)
	if err != nil {
		t.Fatalf("no se esperaba error al eliminar cliente sin solicitudes: %v", err)
	}

	if _, ok := customerRepo.Customers[existing.ID]; ok {
		t.Fatalf("el cliente debería haberse eliminado del repositorio")
	}
}
