package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/zhurak-v/techpassport/auth-service/src/core/entities"
	"github.com/zhurak-v/techpassport/auth-service/src/core/ports/repository"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(db *sqlx.DB) repository.IUserRepositoryContract {
	return &UserRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *UserRepository) getUserWithRelationsBy(field string, value any) (*entities.User, error) {
	user := entities.User{}

	queryUser := `
        SELECT id, created_at, updated_at, email
        FROM users
        WHERE ` + field + ` = $1
        LIMIT 1
    `
	if err := r.db.Get(&user, queryUser, value); err != nil {
		return nil, err
	}

	var roles []entities.Role
	if err := r.db.Select(&roles, `
        SELECT r.id, r.created_at, r.updated_at, r.name
        FROM roles r
        JOIN users_roles ur ON ur.role_id = r.id
        WHERE ur.user_id = $1
    `, user.GetID()); err != nil {
		return nil, err
	}
	user.Roles = roles

	return &user, nil
}

func (r *UserRepository) upsertUserRoles(db *sqlx.DB, user *entities.User) error {
	userID := user.GetID()

	var currentRoleIDs []uuid.UUID
	if err := db.Select(&currentRoleIDs, `SELECT role_id FROM users_roles WHERE user_id = $1`, userID); err != nil {
		return err
	}

	currentRoleMap := make(map[uuid.UUID]struct{}, len(currentRoleIDs))
	for _, id := range currentRoleIDs {
		currentRoleMap[id] = struct{}{}
	}

	newRoleMap := make(map[uuid.UUID]struct{}, len(user.GetRoles()))
	for _, role := range user.GetRoles() {
		newRoleMap[role.GetID()] = struct{}{}
	}

	var toDelete []uuid.UUID
	for id := range currentRoleMap {
		if _, ok := newRoleMap[id]; !ok {
			toDelete = append(toDelete, id)
		}
	}

	if len(toDelete) > 0 {
		query := `DELETE FROM users_roles WHERE user_id = $1 AND role_id = ANY($2)`
		if _, err := db.Exec(query, userID, toDelete); err != nil {
			return err
		}
	}

	var toInsert []uuid.UUID
	for id := range newRoleMap {
		if _, ok := currentRoleMap[id]; !ok {
			toInsert = append(toInsert, id)
		}
	}

	if len(toInsert) > 0 {
		query := `INSERT INTO users_roles (user_id, role_id) SELECT $1, unnest($2::uuid[])`
		if _, err := db.Exec(query, userID, toInsert); err != nil {
			return err
		}
	}

	return nil
}

func (r *UserRepository) FindUserById(id uuid.UUID) (*entities.User, error) {
	return r.getUserWithRelationsBy("id", id)
}

func (r *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	return r.getUserWithRelationsBy("email", email)
}

func (r *UserRepository) FindAllUsers(withRelations bool, limit *int) (*[]entities.User, error) {
	var users []entities.User
	query := `
        SELECT id, created_at, updated_at, email
        FROM users
    `

	if limit != nil {
		query += " LIMIT $1"
		if err := r.db.Select(&users, query, *limit); err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Select(&users, query); err != nil {
			return nil, err
		}
	}

	if withRelations {
		for u := range users {
			userWithRelations, err := r.getUserWithRelationsBy("id", users[u].GetID())
			if err != nil {
				return nil, err
			}
			users[u] = *userWithRelations
		}
	}

	return &users, nil
}

func (r *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {
	user.SetUpdateAt()

	query := `
        INSERT INTO users (id, created_at, updated_at, email)
        VALUES (:id, :created_at, :updated_at, :email)
    `
	if _, err := r.db.NamedExec(query, user); err != nil {
		return nil, err
	}

	if len(user.Roles) > 0 {
		if err := r.upsertUserRoles(r.db, user); err != nil { // змінити на db.Exec у upsertUserRoles
			return nil, err
		}
	}

	return r.FindUserById(user.GetID())
}

func (r *UserRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	user.SetUpdateAt()

	query := `
        UPDATE users
        SET email = :email,
            updated_at = :updated_at
        WHERE id = :id
    `
	if _, err := r.db.NamedExec(query, user); err != nil {
		return nil, err
	}

	if err := r.upsertUserRoles(r.db, user); err != nil {
		return nil, err
	}

	return r.FindUserById(user.GetID())
}

func (r *UserRepository) DeleteUser(user *entities.User) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, user.GetID())
	return err
}
