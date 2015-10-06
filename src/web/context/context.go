package context
import (
	"model"
	"net/http"
	"github.com/gorilla/context"
)

const ContextAuthUserKey = "auth-user";

type RequestContext struct {
	Data interface{}
	User model.User
}

// Builds Request Context to use it for View templates
// If User is not authorized - it will be nil in context
func BuildContext(data interface{}, r *http.Request) RequestContext {
	ctx := RequestContext{}
	ctx.Data = data
	if userObj, ok := context.GetOk(r, ContextAuthUserKey); ok {
		ctx.User, _ = userObj.(model.User)
		//ctx.User = user
	}

	return ctx
}