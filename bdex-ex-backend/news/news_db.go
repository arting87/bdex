package news

import (
	"fmt"
	"time"
)

//新闻公告数据结构体
//alter table `tablename` convert to character set utf8; 修改表编码
type ExchangeNews struct {
	Id       int64     `xorm:"'id' pk autoincr"`
	Title    string    `xorm:"'title' varchar(255)"`
	TypeCode int       `xorm:"'typeCode' int "`
	Content  string    `xorm:"'content' Text"`
	Time     time.Time `xorm:"'time' index"`
	//UpdateTime time.Time `xorm:"update_time"`
}

type PartOfExchangeNews struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`
	TypeCode int       `json:"typeCode"`
	Time     time.Time `json:"time"`
}

type OneOfExchangeNews struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`
	TypeCode int       `json:"typeCode"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

/**
Description:查询所有新闻列表

Parameter:

Return:
    新闻公告数组
    错误
*/
//查询所有新闻列表
func GetNewsList() ([]PartOfExchangeNews, error) {
	newsArray := make([]PartOfExchangeNews, 0)
	err := db.Table("ExchangeNews").Desc("time").Cols("id", "title", "typeCode", "time").Find(&newsArray)
	if err != nil {
		fmt.Println("ExchangeNews Table GetNewsList:", "error", err)
		return nil, err
	}
	fmt.Println(newsArray)
	return newsArray, nil
}

/**
Description:根据id查询单个新闻公告

Parameter:
    id：新闻公告id

Return:
    新闻公告结构体
    错误
    是否存在该id的新闻公告
*/
//根据id查询单个新闻公告
func GetNewsById(id int64) (OneOfExchangeNews, error, bool) {
	new := OneOfExchangeNews{}
	istrue, err := db.Table("ExchangeNews").Where("id=?", id).Get(&new)
	if err != nil {
		log.Errorw("ExchangeNews Table find one:", "error", err)
		return OneOfExchangeNews{}, err, istrue
	}
	return new, nil, istrue
}

/**
Description:根据类型查询新闻公告

Parameter:
    typecode：新闻公告类型

Return:
    新闻公告数组
    错误
*/
//根据类型查询新闻公告
func GetNewsByTypes(typecode int) ([]PartOfExchangeNews, error) {
	newsArray := make([]PartOfExchangeNews, 0)
	err := db.Table("ExchangeNews").Desc("time").Where("typeCode = ?", typecode).Cols("id", "title", "typeCode", "time").Find(&newsArray)
	if err != nil {
		log.Errorw("ExchangeNews Table find by type:", "error", err)
		return nil, err
	}
	return newsArray, nil
}

/**
Description:插入新闻公告

Parameter:
    title：新闻公告标题
    typecode：新闻公告类型
    content： 新闻公告内容

Return:
    错误
*/
//插入新闻公告（单条）
func InsertNews(title string, typecode int, content string) ( error) {

	new := ExchangeNews{Title: title, TypeCode: typecode, Content: content, Time: time.Now()}
	_, err := db.Table("ExchangeNews").Insert(&new)
	if err != nil {
		log.Errorw("ExchangeNews Table insert one:", "error", err)
		return err
	}
	return nil
}
/**
Description:删除新闻公告

Parameter:
    id：新闻公告id

Return:
    错误
*/
//删除新闻公告
func DeleteNewsById(id int64) (error) {
	new := ExchangeNews{}
	_, err := db.Where("id = ?", id).Delete(&new)
	if err != nil {
		log.Errorw("ExchangeNews Table delete one:", "error", err)
		return err
	}
	return nil
}


//修改新闻公告
/*func UpdateNewsById(id int64, title string, typecode int, content string) ( error) {
	new := ExchangeNews{Id: id, Title: title, TypeCode: typecode, Content: content,}
	_, err := db.Table("ExchangeNews").Where("id = ?", id).Update(&new)
	if err != nil {
		log.Errorw("ExchangeNews Table update one:", "error", err)
		return err
	}
	return nil
}*/
