package model

import (
	"gotest/common"
	"gotest/entity"
)

func MemberAdd(member *entity.Member)  {
	_, err := dbm.Insert(member)
	if err != nil {
		common.Logger().Println("add member error:", err)
		return
	}
	common.Logger().Println("add member successful:")
}
func MemberDelete(member *entity.Member)  {
	_, err := dbm.Delete(member)
	if err != nil {
		common.Logger().Println("delete member error:", err)
		return
	}
	common.Logger().Println("delete member successful:")
}
func GetMember(member *entity.Member){
	_, err := dbm.Get(member)
	if err != nil {
		common.Logger().Println("get member error:", err)
		return
	}
	common.Logger().Println("get member successful:")
}

func UpdateMember(id int, member *entity.Member){
	_, err := dbm.ID(id).Update(member)
	if err != nil {
		common.Logger().Println("update member error: ", err)
	}
	common.Logger().Println("update member successful:")
}

func MemberList() ([]entity.Member, error) {
	users := make([]entity.Member, 0)
	err := dbm.Where("name<>''").Find(&users)
	return users,err
}


