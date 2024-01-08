package server

import (
	"context"
	"errors"
	"github.com/xiaka53/golang_server/dao"
	"github.com/xiaka53/golang_server/pkg/pkg"
)

type recommender struct {
}

func NewRecommenderServer() pkg.RecommenderServer {
	return &recommender{}
}

// 推荐人列表
func (r *recommender) List(ctx context.Context, req *pkg.RecommenderListRequest) (*pkg.RecommenderListResponse, error) {
	var (
		info pkg.RecommenderListResponse
		data dao.Account
		list []dao.Account
		err  error
	)
	data.Token = req.GetToken()
	if len(data.Token) != dao.TOKEN_LENGTH {
		goto OUT
	}
	if err = (&data).First(); err != nil {
		goto OUT
	}
	if data.Recommender != dao.ADMIN_MARK && data.Id != dao.ADMIN_MARK {
		err = errors.New("recommender err")
		goto OUT
	}
	info.TotalCount, list = (&data).GetRecommenderBuyAccount(req.GetFrom()-1, req.Amount)
	for _, v := range list {
		info.List = append(info.List, &pkg.RecommenderInfo{
			Name:   v.Name,
			Phone:  v.Phone,
			Url:    dao.GetDomain(v.Url),
			Amount: v.Amount,
			Memo:   (&dao.Memo{Aid: data.Id, Rid: v.Id}).Memo,
		})
	}
OUT:
	return &info, err
}
