package cache

import (
	"fmt"
	"im/internal/logic/model"
	"im/pkg/util"
	"testing"
)

func TestGroupUserCache_Get(t *testing.T) {
	user, err := GroupUserCache.Get(1)
	fmt.Println(err)
	fmt.Println(util.JsonMarshal(user))
}

func TestGroupUserCache_Set(t *testing.T) {
	fmt.Println(GroupUserCache.Set(1, []model.GroupUser{
		{
			UserId:  1,
			GroupId: 0,
			Label:   "2",
			Extra:   "2",
		},
	}))
}
func TestGroupUserCache_Del(t *testing.T) {
	fmt.Println(GroupUserCache.Del(1))
}
