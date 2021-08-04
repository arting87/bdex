package controllers

import (
	"bdex.in/bdex/bdex-admin-backend/models"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"strings"
	"time"
)

//普通用户检查

func CheckUserInfo(input *context.BeegoInput) (error) {
	token, err := getToken(input.Header("Authorization"))
	if err != nil {
		return err
	}
	logs.Info("token:", token)
	if !CheckSecret(token, NormalPermission) {
		return fmt.Errorf("身份验证失败")
	}
	return nil
}

//高级用户权限检查
func CheckAdminUserInfo(input *context.BeegoInput) (error) {
	token, err := getToken(input.Header("Authorization"))
	if err != nil {
		return err
	}
	logs.Info("token:", token)
	if !CheckSecret(token, AdminPermission) {
		return fmt.Errorf("身份验证失败")
	}
	return nil
}

//用户信息加密
func EncoderHashSecret(secret string) (string, error) {
	if secret == "" {
		return "", fmt.Errorf("传入字符串不能为空")
	} else {
		sh2 := sha256.New()
		sh2.Write([]byte(secret))
		result := sh2.Sum([]byte(""))
		return hex.EncodeToString(result), nil
	}
}

//用户信息解密

func EncoderSecret(user models.User) (string, error) {
	hashPwd, err := EncoderHashSecret(user.PassWord)
	if err != nil {
		return "", err
	}
	encoderUser := UserInfoSafeCheck{Id: user.Id, Name: user.Name, Pwd: hashPwd, Permission: user.Permission, TimeStamp: time.Now().Unix()}

	dataByte, err := json.Marshal(&encoderUser)
	if err != nil {
		logs.Error("json.Marshal error", err)
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(dataByte), nil
}

func DecoderSecret(token string) (*UserInfoSafeCheck, error) {
	strByte, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	var userDecoder UserInfoSafeCheck

	err = json.Unmarshal(strByte, &userDecoder)
	if err != nil {
		logs.Error("json.Unmarshal error", err)
		return nil, nil
	}
	return &userDecoder, nil
}

//用户信息检查
func CheckSecret(token string, UserPermission int) bool {
	decoderUser, err := DecoderSecret(token)
	if err != nil {
		return false
	}
	user, err := models.FindUserByName(decoderUser.Name)
	if err != nil {
		return false
	}
	hashPwd, err := EncoderHashSecret(user.PassWord)
	if err != nil {
		return false
	}

	if user.Id != decoderUser.Id || user.Name != decoderUser.Name || hashPwd != decoderUser.Pwd || user.Permission != decoderUser.Permission {
		return false
	}

	if decoderUser.TimeStamp+ExpireTime < time.Now().Unix() {
		logs.Debug("token过期")
		return false
	}

	if UserPermission == AdminPermission {
		if user.Permission != UserPermission {
			return false
		} else {
			return true
		}
	} else {
		if user.Permission == NormalPermission {
			return true
		} else if user.Permission == AdminPermission {
			return true
		} else {
			return false
		}
	}

}

//验证token是否合法
func getToken(token string) (string, error) {
	logs.Info("Authorization:", token)
	if token == "" {
		return "", fmt.Errorf("token为空")
	} else {
		tokenArray := strings.Split(token, " ")
		if tokenArray[1] == "" || tokenArray[0] != "Bearer" {
			return "", fmt.Errorf("token解析失败,格式错误！");
		} else {
			return tokenArray[1], nil
		}
	}
}
