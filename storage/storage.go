package storage

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Storage struct {
	users   map[int]User
	counter int
}

func New() (*Storage, error) {
	return &Storage{
		users: make(map[int]User),
	}, nil
}

func (s *Storage) CreateUser(user User) {
	s.counter++
	user.ID = s.counter
	s.users[user.ID] = user
}

func (s *Storage) GetUser(id int) (User, bool) {
	user, ok := s.users[id]
	return user, ok
}

func (s *Storage) DeleteUser(id int) bool {
	if _, ok := s.users[id]; !ok {
		return false
	}
	delete(s.users, id)
	return true
}
