package usecase

import (
	"github.com/v-lozhkin/deployProject/internal/pkg/item"
	statpkg "github.com/v-lozhkin/deployProject/pkg/stat"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Stat struct {
	MethodDuration *statpkg.Timer
}

type usecase struct {
	repo item.Repository
	stat Stat
}

func New(repo item.Repository, stat promauto.Factory) item.Usecase {
	ret := usecase{
		repo: repo,
		stat: Stat{MethodDuration: &statpkg.Timer{HistogramVec: stat.NewHistogramVec(
			prometheus.HistogramOpts{Name: "usecase_method_duration"},
			[]string{"method_name"},
		)}},
	}

	prometheus.MustRegister(ret.stat.MethodDuration)

	return ret
}
