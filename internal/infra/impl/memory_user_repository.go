package impl

import (
	"user_service/internal/application"
	"user_service/internal/domain/entities"
)

type MemoryUserRepository struct {
	data []entities.User
}

func NewMemoryUserRepository() MemoryUserRepository {
	return MemoryUserRepository{[]entities.User{}}
}

func (m *MemoryUserRepository) Add(entity entities.User) error {
	m.data = append(m.data, entity)
	return nil
}

func (m *MemoryUserRepository) Delete(entity entities.User) error {
	for i, user := range m.data {
		if user.Id == entity.Id {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MemoryUserRepository) Update(entity entities.User) error {
	for i, user := range m.data {
		if user.Id == entity.Id {
			m.data[i] = entity
			return nil
		}
	}
	return nil
}

func (m *MemoryUserRepository) Get(filter application.UserFilter) ([]entities.User, error) {
	result := []entities.User{}

	for _, user := range m.data {
		if filter.CreatedAt != nil {
			if user.CreatedAt != *filter.CreatedAt {
				continue
			}
		}
		if filter.UpdatedAt != nil {
			if user.UpdatedAt != *filter.UpdatedAt {
				continue
			}
		}
		if filter.Id != nil {
			if user.Id != *filter.Id {
				continue
			}
		}
		if filter.Name != nil {
			if user.Name != *filter.Name {
				continue
			}
		}
		if filter.Email != nil {
			if user.Email != *filter.Email {
				continue
			}
		}
		if filter.Birthday != nil {
			if user.Birthday != *filter.Birthday {
				continue
			}
		}
		if filter.HashedPassword != nil {
			if user.HashedPassword != *filter.HashedPassword {
				continue
			}
		}

		result = append(result, user)
	}

	return result, nil
}

func (m *MemoryUserRepository) GetOne(filter application.UserFilter) (entities.User, error) {
	result := []entities.User{}

	for _, user := range m.data {
		if filter.CreatedAt != nil {
			if user.CreatedAt != *filter.CreatedAt {
				continue
			}
		}
		if filter.UpdatedAt != nil {
			if user.UpdatedAt != *filter.UpdatedAt {
				continue
			}
		}
		if filter.Id != nil {
			if user.Id != *filter.Id {
				continue
			}
		}
		if filter.Name != nil {
			if user.Name != *filter.Name {
				continue
			}
		}
		if filter.Email != nil {
			if user.Email != *filter.Email {
				continue
			}
		}
		if filter.Birthday != nil {
			if user.Birthday != *filter.Birthday {
				continue
			}
		}
		if filter.HashedPassword != nil {
			if user.HashedPassword != *filter.HashedPassword {
				continue
			}
		}

		result = append(result, user)
	}

	return result[0], nil
}
