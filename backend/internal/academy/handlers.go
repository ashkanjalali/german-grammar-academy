package academy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var levels = []gin.H{
	{"code":"A1","name":"Beginner","description":"Build the foundations"},
	{"code":"A2","name":"Elementary","description":"Expand everyday grammar"},
	{"code":"B1","name":"Intermediate","description":"Become an independent user"},
	{"code":"B2","name":"Upper intermediate","description":"Master complex structures"},
	{"code":"C1","name":"Advanced","description":"Use precise formal German"},
	{"code":"C2","name":"Proficient","description":"Develop near-native control"},
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	api.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status":"ok","service":"german-grammar-academy"}) })
	api.GET("/levels", func(c *gin.Context) { c.JSON(http.StatusOK, levels) })
	api.GET("/lessons", func(c *gin.Context) {
		q := db.Order("\"order\" asc")
		if level := c.Query("level"); level != "" { q = q.Where("level = ?", level) }
		var lessons []Lesson
		if err := q.Find(&lessons).Error; err != nil { c.JSON(500, gin.H{"error":"could not load lessons"}); return }
		c.JSON(200, lessons)
	})
	api.GET("/lessons/:id", func(c *gin.Context) {
		var lesson Lesson
		if err := db.First(&lesson, c.Param("id")).Error; err != nil { c.JSON(404, gin.H{"error":"lesson not found"}); return }
		c.JSON(200, lesson)
	})
	api.GET("/exams/:level", func(c *gin.Context) {
		var qs []Question
		if err := db.Where("level = ?", c.Param("level")).Find(&qs).Error; err != nil { c.JSON(500, gin.H{"error":"could not load exam"}); return }
		out := make([]gin.H, 0, len(qs))
		for _, q := range qs {
			var opts []string; _ = json.Unmarshal([]byte(q.Options), &opts)
			out = append(out, gin.H{"id":q.ID,"prompt":q.Prompt,"options":opts})
		}
		c.JSON(200, gin.H{"level":c.Param("level"),"durationMinutes":30,"passMark":70,"questions":out})
	})
	api.POST("/exams/:level/submit", func(c *gin.Context) {
		var body struct { UserID string `json:"userId"`; Answers map[string]int `json:"answers"` }
		if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":"invalid request body"}); return }
		if body.UserID == "" { body.UserID = "demo" }
		var qs []Question
		db.Where("level = ?", c.Param("level")).Find(&qs)
		if len(qs) == 0 { c.JSON(404, gin.H{"error":"no exam questions for this level"}); return }
		score := 0
		review := make([]gin.H, 0, len(qs))
		for _, q := range qs {
			var opts []string; _ = json.Unmarshal([]byte(q.Options), &opts)
			answer, ok := body.Answers[strconv.Itoa(int(q.ID))]
			correct := ok && answer == q.CorrectAnswer
			if correct { score++ }
			review = append(review, gin.H{"questionId":q.ID,"correct":correct,"correctAnswer":q.CorrectAnswer,"explanation":q.Explanation})
		}
		pct := score * 100 / len(qs)
		attempt := Attempt{UserID:body.UserID, Level:c.Param("level"), Score:score, Total:len(qs), Passed:pct>=70}
		if err := db.Create(&attempt).Error; err != nil { c.JSON(500, gin.H{"error":"could not save exam result"}); return }
		c.JSON(200, gin.H{"score":score,"total":len(qs),"percentage":pct,"passed":pct>=70,"review":review})
	})
	api.GET("/progress", func(c *gin.Context) {
		userID := c.DefaultQuery("userId", "demo")
		var attempts []Attempt
		db.Where("user_id = ?", userID).Order("created_at desc").Limit(10).Find(&attempts)
		var completed int64
		db.Model(&Progress{}).Where("user_id = ?", userID).Select("completed_lesson").Scan(&completed)
		c.JSON(200, gin.H{"userId":userID,"completedLessons":completed,"recentAttempts":attempts})
	})
}
