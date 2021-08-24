package router

import (
	"fmt"
	"reflect"

	"github.com/kovey/server-go/http/context"
	"github.com/kovey/server-go/http/controller"
)

type Router struct {
	path         string
	action       string
	controller   interface{}
	base         reflect.Type
	viewDisabled bool
	viewPath     string
}

func NewRouter(path string, action string, ctl interface{}, viewPath string) *Router {
	return &Router{
		path: path, action: action, controller: ctl, base: reflect.TypeOf((*controller.ControllerInterface)(nil)).Elem(),
		viewDisabled: len(viewPath) == 0, viewPath: viewPath,
	}
}

func (r *Router) Call(ctx *context.Context) error {
	vType := reflect.TypeOf(r.controller)
	if !vType.Implements(r.base) {
		ctx.Output().PageNotFound()
		return fmt.Errorf("controller is not extends Controller")
	}

	var vValue reflect.Value
	if vType.Kind() == reflect.Ptr {
		vValue = reflect.New(vType.Elem())
	} else {
		vValue = reflect.New(vType)
	}

	var base reflect.Value
	if vValue.Kind() == reflect.Ptr {
		base = vValue.Elem().FieldByName("Controller")
	} else {
		base = vValue.FieldByName("Controller")
	}

	if !base.Type().Implements(r.base) {
		ctx.Output().PageNotFound()
		return fmt.Errorf("controller is not extends Controller")
	}

	base.Set(reflect.ValueOf(controller.NewController(ctx, r.viewPath)))

	fun := vValue.MethodByName(r.action)
	if !fun.IsValid() || fun.IsZero() {
		ctx.Output().PageNotFound()
		return fmt.Errorf("action[%s] is not exists", r.action)
	}

	args := make([]reflect.Value, 0)
	result := Result(fun.Call(args))
	if result != nil {
		return result
	}

	if !r.ViewDisabled() {
		render := vValue.MethodByName("Render")
		result = Result(render.Call(args))
		if result != nil {
			return result
		}
	}

	ctx.Output().SetStatus(200)
	return nil
}

func (r *Router) Path() string {
	return r.path
}

func (r *Router) ViewDisabled() bool {
	return r.viewDisabled
}

func Result(result []reflect.Value) error {
	if len(result) == 0 {
		return nil
	}

	res, ok := result[0].Interface().(error)
	if ok {
		return res
	}

	return nil
}
