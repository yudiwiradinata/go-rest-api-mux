package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/yudiwiradinata/go-rest-api-mux/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type firestoreRepository struct{}

func NewFirestoreRepository() PostRepository {
	return &firestoreRepository{}
}

const (
	projectId      = "go-lang-mux"
	collectionName = "posts"
)

var (
	opt = option.WithCredentialsFile("./go-lang-mux-firebase.json")
)

func (*firestoreRepository) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	var client, err = firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("Error to create firestore client %v", err)
		return nil, err
	}

	defer client.Close()
	var _, _, errw = client.Collection(collectionName).Add(ctx, post)
	if errw != nil {
		log.Fatalf("Error to save data %v", err)
		return nil, err
	}
	return post, nil
}

func (*firestoreRepository) FindAll() ([]entity.Post, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("Error to create firestore client %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entity.Post
	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error to iterate data firestore client %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
