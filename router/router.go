package router

import (
	"github.com/gin-gonic/gin"
	v1 "gotest/controllers/v1"
)

func InitRouter(r *gin.Engine){

	GroupV1 := r.Group("/v1")
	{
		GroupV1.POST("/member/add", v1.AddMember)
		GroupV1.PUT("/member/delete/:id", v1.DeleteMemberById)
		GroupV1.GET("/member/get/:id", v1.GetMember)
		GroupV1.POST("/member/update", v1.UpdateMember)
		GroupV1.GET("/members", v1.GetMemberList)
	}
}
