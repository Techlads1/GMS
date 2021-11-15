package services

import (
	"errors"
	"gateway/webserver/auth"

	"strconv"
	"strings"
	"time"

	//"gateway/services/entity"
	//"gateway/webserver/systems/auth/auth"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

const aclSuffix = "-acl"
const ResponseMsgKey = "respmsg"

type ResponseMessage struct {
	Error   bool     `json:"error"`
	Message []string `json:"message"`
}

func Init() {
	Cache = cache.New(5*time.Minute, 15*time.Minute)
}

// Map defines a generic map of type `map[string]interface{}`.
type Map map[string]interface{}

func Serve(c echo.Context, data map[string]interface{}) Map {
	_, email := auth.GetUserFromContext(c)
	user, b := GetCache(email)
	if b && user != nil {
		u := user.(*entity.User)
		data["id"] = u.ID
		data["username"] = u.Name
		data["useremail"] = u.Email
		data[errorMessageKey] = GetErrorMessage(c)
		data[infoMessageKey] = GetInfoMessage(c)
		//data["csrf"] = c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)

	} else {
		data[errorMessageKey] = GetErrorMessage(c)
		data[infoMessageKey] = GetInfoMessage(c)
		//data["csrf"] = c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	}

	aclKey := GetACLKey(email)
	acl, bb := GetCache(aclKey)
	if bb && acl != nil {
		uu := acl.(*entity.UserACL)
		data["roles"] = uu.Roles
		data["permissions"] = uu.Permissions
	}

	//Get ResponseMessage
	e, r := GetCache(ResponseMsgKey)
	if r && e != nil {
		errMsg := e.(*ResponseMessage)
		data[ResponseMsgKey] = errMsg
	}
	data["path"] = c.Request().URL.String()
	return data
}

func StoreCache(key string, value interface{}) error {
	if Cache == nil {
		Cache = cache.New(30*time.Minute, 90*time.Minute)
	}
	_, ok := Cache.Get(key)
	if ok {
		return nil
	}
	return Cache.Add(key, value, 30*time.Minute)
}

func GetCache(key string) (interface{}, bool) {
	return Cache.Get(key)
}

func ClearCache(key string) {
	Cache.Set(key, nil, 1*time.Second)
	/*if Cache == nil {
		return nil
	}
	Cache.Flush()
	return nil*/

}

func GetACLKey(email string) string {
	return email + aclSuffix
}

func SetResponseMessage(isError bool, message ...string) error {
	err := &ResponseMessage{
		Error:   isError,
		Message: message,
	}
	return Cache.Add(ResponseMsgKey, err, 500*time.Millisecond)
}

func HasPermission(email, res string) (bool, error) {
	res = removeSuffix(res)
	key := GetACLKey(email)
	acl, ok := GetCache(key)
	if !ok {
		return false, errors.New("error getting cached permission")
	}
	userACL := acl.(*entity.UserACL)
	p := userACL.HasPermission(res)
	return p, nil
}

func removeSuffix(res string) string {
	r := strings.Split(res, "/")
	l := len(r)
	if l > 1 {
		last := r[l-1]               //get the last substring
		_, err := strconv.Atoi(last) //check if the last sub is integer
		if err == nil {
			r = r[:len(r)-1] //remove the last element, e.g number
			nr := ""
			if len(r) == 1 {
				nr = "/"
			}
			nr = nr + strings.Join(r, "/")
			return nr
		}
		return res
	}
	return res
}
