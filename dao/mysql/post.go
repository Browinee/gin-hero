package mysql

import (
	"master-gin/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id)
		values(?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(postID int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
		from post
		where post_id = ?
	`
	err = db.Get(post, sqlStr, postID)
	return
}

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
		from post
		where post_id in (?)
		order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return

}

func GetPostList(offset, limit int64) (posts []*models.ApiPostDetail, err error) {
	sqlStr := `select
		post_id, title, content, author_id, community_id, create_time
		from post
		limit ?,?
	`
	posts = make([]*models.ApiPostDetail, 0, limit)
	err = db.Select(&posts, sqlStr, (offset-1)*limit, limit)
	return
}
