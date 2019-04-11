package service

import (
	"github.com/golang/sync/errgroup"
	"module/logger"
)

var (
	g errgroup.Group
)

func StartService()  {

	g.Go(func() error {
		return staticService().ListenAndServeTLS("./ssl/cert.pem","./ssl/key.pem")
	})

	g.Go(func() error {
		return apiService().ListenAndServeTLS("./ssl/cert.pem","./ssl/key.pem")
	})

	g.Go(func() error {
		return redirectService().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		// Write to error file.
		logger.Error(err)
	}
}