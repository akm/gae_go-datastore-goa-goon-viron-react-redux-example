package model

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type User struct {
	Email     string    `datastore:"-" goon:"id" json:"email"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

type Memo struct {
	Id        int64     `datastore:"-" goon:"id" json:"id"`
	Content   string    `json:"content,omitempty"`
	Shared    bool      `json:"shared,omitempty"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

func (m *User) PrepareToCreate() error {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now()
	}
	return nil
}

func (m *User) PrepareToUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Memo) PrepareToCreate() error {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now()
	}
	return nil
}

func (m *Memo) PrepareToUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}

type UserStore struct {
}

func (s *UserStore) All(ctx context.Context) ([]*User, error) {
	g := GoonFromContext(ctx)
	r := []*User{}
	k := g.Kind(new(User))
	log.Infof(ctx, "Kind is %v\n", k)
	q := datastore.NewQuery(k)
	log.Infof(ctx, "q is %v\n", q)
	_, err := g.GetAll(q.EventualConsistency(), &r)
	if err != nil {
		log.Errorf(ctx, "Failed to GetAll User because of %v\n", err)
		return nil, err
	}
	return r, nil
}

func (s *UserStore) ByID(ctx context.Context, email string) (*User, error) {
	r := User{Email: email}
	err := s.Get(ctx, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *UserStore) ByKey(ctx context.Context, key *datastore.Key) (*User, error) {
	if err := s.IsValidKey(ctx, key); err != nil {
		log.Errorf(ctx, "UserStore.ByKey got Invalid key: %v because of %v\n", key, err)
		return nil, err
	}

	r := User{Email: key.StringID()}
	err := s.Get(ctx, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *UserStore) Get(ctx context.Context, m *User) error {
	g := GoonFromContext(ctx)
	err := g.Get(m)
	if err != nil {
		log.Errorf(ctx, "Failed to Get User because of %v\n", err)
		return err
	}

	return nil
}

func (s *UserStore) IsValidKey(ctx context.Context, key *datastore.Key) error {
	if key == nil {
		return fmt.Errorf("key is nil")
	}
	g := GoonFromContext(ctx)
	expected := g.Kind(&User{})
	if key.Kind() != expected {
		return fmt.Errorf("key kind must be %s but was %s", expected, key.Kind())
	}
	return nil
}

func (s *UserStore) Exist(ctx context.Context, m *User) (bool, error) {
	g := GoonFromContext(ctx)
	key, err := g.KeyError(m)
	if err != nil {
		log.Errorf(ctx, "Failed to Get Key of %v because of %v\n", m, err)
		return false, err
	}
	_, err = s.ByKey(ctx, key)
	if err == datastore.ErrNoSuchEntity {
		return false, nil
	} else if err != nil {
		log.Errorf(ctx, "Failed to get existance of %v because of %v\n", m, err)
		return false, err
	} else {
		return true, nil
	}
}

func (s *UserStore) Create(ctx context.Context, m *User) (*datastore.Key, error) {
	err := m.PrepareToCreate()
	if err != nil {
		return nil, err
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}

	exist, err := s.Exist(ctx, m)
	if err != nil {
		return nil, err
	}
	if exist {
		log.Errorf(ctx, "Failed to create %v because of another entity has same key\n", m)
		return nil, fmt.Errorf("Duplicate Name error: %q of %v\n", m.Name, m)
	}

	return s.Put(ctx, m)
}

func (s *UserStore) Update(ctx context.Context, m *User) (*datastore.Key, error) {
	err := m.PrepareToUpdate()
	if err != nil {
		return nil, err
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}

	exist, err := s.Exist(ctx, m)
	if err != nil {
		return nil, err
	}
	if !exist {
		log.Errorf(ctx, "Failed to update %v because it doesn't exist\n", m)
		return nil, fmt.Errorf("No data to update %q of %v\n", m.Name, m)
	}

	return s.Put(ctx, m)
}

func (s *UserStore) Put(ctx context.Context, m *User) (*datastore.Key, error) {
	g := GoonFromContext(ctx)
	key, err := g.Put(m)
	if err != nil {
		log.Errorf(ctx, "Failed to Put %v because of %v\n", m, err)
		return nil, err
	}
	return key, nil
}

type MemoStore struct {
}

func (s *MemoStore) All(ctx context.Context) ([]*Memo, error) {
	g := GoonFromContext(ctx)
	r := []*Memo{}
	k := g.Kind(new(Memo))
	log.Infof(ctx, "Kind is %v\n", k)
	q := datastore.NewQuery(k)
	log.Infof(ctx, "q is %v\n", q)
	_, err := g.GetAll(q.EventualConsistency(), &r)
	if err != nil {
		log.Errorf(ctx, "Failed to GetAll Memo because of %v\n", err)
		return nil, err
	}
	return r, nil
}

func (s *MemoStore) ByID(ctx context.Context, id int64) (*Memo, error) {
	r := Memo{Id: id}
	err := s.Get(ctx, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *MemoStore) ByKey(ctx context.Context, key *datastore.Key) (*Memo, error) {
	if err := s.IsValidKey(ctx, key); err != nil {
		log.Errorf(ctx, "MemoStore.ByKey got Invalid key: %v because of %v\n", key, err)
		return nil, err
	}

	r := Memo{Id: key.IntID()}
	err := s.Get(ctx, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *MemoStore) Get(ctx context.Context, m *Memo) error {
	g := GoonFromContext(ctx)
	err := g.Get(m)
	if err != nil {
		log.Errorf(ctx, "Failed to Get Memo because of %v\n", err)
		return err
	}

	return nil
}

func (s *MemoStore) IsValidKey(ctx context.Context, key *datastore.Key) error {
	if key == nil {
		return fmt.Errorf("key is nil")
	}
	g := GoonFromContext(ctx)
	expected := g.Kind(&Memo{})
	if key.Kind() != expected {
		return fmt.Errorf("key kind must be %s but was %s", expected, key.Kind())
	}
	return nil
}

func (s *MemoStore) Exist(ctx context.Context, m *Memo) (bool, error) {
	g := GoonFromContext(ctx)
	key, err := g.KeyError(m)
	if err != nil {
		log.Errorf(ctx, "Failed to Get Key of %v because of %v\n", m, err)
		return false, err
	}
	_, err = s.ByKey(ctx, key)
	if err == datastore.ErrNoSuchEntity {
		return false, nil
	} else if err != nil {
		log.Errorf(ctx, "Failed to get existance of %v because of %v\n", m, err)
		return false, err
	} else {
		return true, nil
	}
}

func (s *MemoStore) Create(ctx context.Context, m *Memo) (*datastore.Key, error) {
	err := m.PrepareToCreate()
	if err != nil {
		return nil, err
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}

	return s.Put(ctx, m)
}

func (s *MemoStore) Update(ctx context.Context, m *Memo) (*datastore.Key, error) {
	err := m.PrepareToUpdate()
	if err != nil {
		return nil, err
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}

	return s.Put(ctx, m)
}

func (s *MemoStore) Put(ctx context.Context, m *Memo) (*datastore.Key, error) {
	g := GoonFromContext(ctx)
	key, err := g.Put(m)
	if err != nil {
		log.Errorf(ctx, "Failed to Put %v because of %v\n", m, err)
		return nil, err
	}
	return key, nil
}
