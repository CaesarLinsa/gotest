package v1

import (
	"github.com/gin-gonic/gin"
	"gotest/common"
	"gotest/entity"
	"gotest/model"
	"net/http"
	"strconv"
)

func AddMember(c *gin.Context)  {
	// 获取 POST 参数
	var json entity.Member
	err := c.BindJSON(&json)
	if  err == nil {
		member := entity.Member{Name:json.Name, Age:json.Age}
		model.MemberAdd(&member)
	}else {
		common.Logger().Error("AddMember caught an error: ", err)
	}

}
func DeleteMemberById(c *gin.Context) {
     //api获取id /member/delete/:id
     member_id := c.Param("id")
     id, err := strconv.Atoi(member_id)
     if err != nil {
     	common.Logger().Error("convert member id failed:", err)
	 }
	 member := entity.Member{Id:id}
	 model.MemberDelete(&member)
}

func GetMember(c *gin.Context)  {
	// api参数获取,/member/get/:id
	member_id := c.Param("id")
	id, err := strconv.Atoi(member_id)
	if err != nil {
		common.Logger().Error("convert member id failed:", err)
	}
	member := entity.Member{Id:id}
	model.GetMember(&member)
	if member.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"msg":"get member by id is null"})
	}
	c.JSON(http.StatusOK, member)
}
func  UpdateMember(c *gin.Context)  {
	var json entity.Member
	err := c.BindJSON(&json)
	if  err == nil {
		if json.Id == 0 {
			c.JSON(http.StatusOK, gin.H{"msg":"id can't be null"})
		}
		if json.Age == 0 {
			c.JSON(http.StatusOK, gin.H{"msg":"age can't be null"})
		}
		if json.Name == "" {
			c.JSON(http.StatusOK, gin.H{"msg":"name can't be null"})
		}
		member := entity.Member{Id: json.Id, Name:json.Name, Age:json.Age}
		model.UpdateMember(json.Id, &member)
	}else {
		common.Logger().Error("UpdateMember caught an error: ", err)
	}
}
func GetMemberList(c *gin.Context){
	member_list, err := model.MemberList()
    if err !=nil {
    	common.Logger().Error("GetMemberList failed: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg":member_list})
}