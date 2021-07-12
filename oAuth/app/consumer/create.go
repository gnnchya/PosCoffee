package consumer

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/oAuth/app/view"
	"github.com/gnnchya/PosCoffee/oAuth/service/consumer/consumerin"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func (ctrl *Controller) Create(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.consumer.Create",
	)
	defer span.Finish()

	input := &consumerin.CreateConsumerInput{}
	//if err := c.ShouldBindJSON(input); err != nil {
	//	// TODO match HTTP code
	//	view.MakeErrResp2(c, 422, err)
	//	return
	//}

	ID, err := ctrl.service.Create(ctx, input)
	if err != nil {
		// TODO match HTTP code
		view.MakeErrResp2(c,422, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, ID)
}
