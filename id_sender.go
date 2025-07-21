package id_sender

import (
	"github.com/Huang-gj/id_sender/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func Sender(datasource string, biz_tag string) int64 {
	conn := sqlx.NewMysql(datasource)
	idSenderModel := model.NewIdSenderModel(conn)
	err := idSenderModel.UpdateCurrentIdAutoStep(biz_tag)
	if err != nil {
		panic(err)
	}
	id, err := idSenderModel.FindOneByBizTag(biz_tag)
	if err != nil {
		panic(err)
	}
	return id.Id
}
