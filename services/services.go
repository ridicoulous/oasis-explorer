package services

import (
	"github.com/oasislabs/oasis-core/go/common/grpc"
	"github.com/oasislabs/oasis-core/go/staking/api"
	"github.com/patrickmn/go-cache"
	grpcCommon "google.golang.org/grpc"
	"oasisTracker/conf"
	"oasisTracker/dao"
	"oasisTracker/smodels"
	"time"
)

type (
	Service interface {
		GetInfo() (smodels.Info, error)
		GetBlockList(params smodels.BlockParams) ([]smodels.Block, error)
		GetTransactionsList(params smodels.TransactionsParams) ([]smodels.Transaction, error)
		GetAccountInfo(accountID string) (smodels.Account, error)
		GetChartData(params smodels.ChartParams) ([]smodels.ChartData, error)
		GetEscrowRatioChartData(params smodels.ChartParams) ([]smodels.ChartData, error)
		GetAccountList(listParams smodels.AccountListParams) ([]smodels.AccountList, error)
	}

	ServiceFacade struct {
		cfg     conf.Config
		dao     dao.ServiceDAO
		nodeAPI api.Backend
		cache   *cache.Cache
	}
)

const (
	topEscrowCacheKey = "top_escrow_percent"
	cacheTTL          = 1 * time.Minute
)

func NewService(cfg conf.Config, dao dao.ServiceDAO) *ServiceFacade {
	grpcConn, err := grpc.Dial(cfg.Scanner.NodeConfig, grpcCommon.WithInsecure())
	if err != nil {
		return nil
	}

	sAPI := api.NewStakingClient(grpcConn)

	return &ServiceFacade{
		cfg:     cfg,
		dao:     dao,
		nodeAPI: sAPI,
		cache:   cache.New(cacheTTL, cacheTTL),
	}
}
