package service

import (
	"fmt"
	"master-gin/dao/mysql"
	"master-gin/dao/redis"
	"master-gin/models"
	"master-gin/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()

	err = mysql.CreatePost(p)
	if err != nil {
		return nil
	}
	err = redis.CreatePost(p.ID)
	return
}

func GetPostById(postID int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(postID) failed", zap.Int64("post_id", postID), zap.Error(err))
		return nil, err
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.String("author_id", fmt.Sprint(post.AuthorID)), zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.String("community_id", fmt.Sprint(post.CommunityID)), zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(offset, limit int64) (data []*models.ApiPostDetail, err error) {

	posts, err := mysql.GetPostList(offset, limit)

	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID error", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		post.AuthorName = user.Username
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed", zap.Int64("community_id", post.CommunityID), zap.Error(err))
			continue
		}
		post.CommunityDetail = community
		data = append(data, post)
	}
	return
}
