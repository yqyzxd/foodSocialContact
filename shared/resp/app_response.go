package resp

import (
	appCodes "foodSocialContact/shared/codes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type AppResponse struct {
	Err  appCodes.Code `json:"err""`
	Msg  string        `json:"msg"`
	Data interface{}   `json:"data"`
}

func (resp *AppResponse) JSON(c *gin.Context, err error) {
	httpCode := http.StatusOK
	if err != nil {
		status, ok := status.FromError(err)
		httpCode = http.StatusInternalServerError
		if ok {
			switch status.Code() {
			case codes.AlreadyExists:
				httpCode = http.StatusConflict
			case codes.Internal:
				httpCode = http.StatusInternalServerError
			case codes.InvalidArgument:
				httpCode = http.StatusBadRequest
			}
		}
	}

	c.JSON(httpCode, resp)

}
