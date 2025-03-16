package mongo

import (
	"cmd/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestAddPost(t *testing.T) {
	type args struct {
		c    *mongo.Client
		post storage.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddPost(tt.args.c, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	type args struct {
		c    *mongo.Client
		post storage.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeletePost(tt.args.c, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllPosts(t *testing.T) {
	type args struct {
		c *mongo.Client
	}
	tests := []struct {
		name    string
		args    args
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllPosts(tt.args.c)
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

func TestUpdatePost(t *testing.T) {
	type args struct {
		c    *mongo.Client
		post storage.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdatePost(tt.args.c, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
