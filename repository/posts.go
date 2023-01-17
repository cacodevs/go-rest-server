package repository

import (
	"context"

	"github.com/cecardev/go-rest-server/models"
)

type PostsRepository interface{
    InsertPosts(ctx context.Context, post *models.Post)(id int64, err error)
} 

var implementationPost PostsRepository

func InsertPosts(ctx context.Context, post *models.Post )(id int64,err error){
    return implementationPost.InsertPosts(ctx,post)
}





