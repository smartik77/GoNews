package postgre

import (
	"cmd/pkg/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"testing"
)

func TestNewStorage(t *testing.T) {
	type args struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *Storage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddPost(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		post storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.AddPost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeletePost(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		post storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DeletePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_GetAllPosts(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.GetAllPosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPosts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdatePost(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		post storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.UpdatePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
