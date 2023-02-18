package controllers

import (
	"github.com/daryeon/gin-bbs/models"
	"github.com/gin-gonic/gin"
)

/* 帖子首页 */
func Get(c *gin.Context) {
	c.HTML(200, "list.html", nil)
}

/* 添加帖子页面 */
func Add(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

/* 添加帖子接口 */
func AddPost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)
	ok, _ := models.AddPost(post)
	c.JSON(200, gin.H{"data": post, "ok": ok})
}

/* 帖子列表 */
func GetPosts(c *gin.Context) {
	posts, err := models.GetPosts()
	c.JSON(200, gin.H{"data": posts, "err": err})
}
