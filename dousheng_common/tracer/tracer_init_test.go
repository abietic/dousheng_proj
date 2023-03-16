package tracer

import (
	"fmt"
	"testing"

	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func TestTracerInit(t *testing.T) {
	os.Setenv("JAEGER_DISABLED", "false")
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")
	os.Setenv("JAEGER_REPORTER_LOG_SPANS", "true")
	os.Setenv("JAEGER_AGENT_HOST", "127.0.0.1")
	os.Setenv("JAEGER_AGENT_PORT", "6831")
	var service string = ""
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, _, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
}
