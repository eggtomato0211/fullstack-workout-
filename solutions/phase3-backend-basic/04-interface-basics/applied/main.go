package main

import "fmt"

// UserRepository はデータ永続化を抽象化するインターフェース
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}

type User struct {
	ID   int
	Name string
}

// InMemoryUserRepository はメモリ上で動くUserRepositoryの実装
type InMemoryUserRepository struct {
	users map[int]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[int]*User)}
}

func (r *InMemoryUserRepository) FindByID(id int) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found: id=%d", id)
	}
	return user, nil
}

func (r *InMemoryUserRepository) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

// UserService はインターフェースに依存する
type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(id int, name string) error {
	return s.repo.Save(&User{ID: id, Name: name})
}

func main() {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)

	service.CreateUser(1, "田中太郎")
	service.CreateUser(2, "鈴木花子")

	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Found: %+v\n", user)

	_, err = service.GetUser(99)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
