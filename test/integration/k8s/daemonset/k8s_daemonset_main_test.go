//go:build integration

package otel

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/grafana/beyla/test/integration/components/docker"
	"github.com/grafana/beyla/test/integration/components/kube"
	k8s "github.com/grafana/beyla/test/integration/k8s/common"
	"github.com/grafana/beyla/test/tools"
)

const (
	testTimeout = 3 * time.Minute

	jaegerQueryURL = "http://localhost:36686/api/traces"
)

var cluster *kube.Kind

func TestMain(m *testing.M) {
	if err := docker.Build(os.Stdout, tools.ProjectDir(),
		docker.ImageBuild{Tag: "testserver:dev", Dockerfile: k8s.DockerfileTestServer},
		docker.ImageBuild{Tag: "beyla:dev", Dockerfile: k8s.DockerfileBeyla},
		docker.ImageBuild{Tag: "grpcpinger:dev", Dockerfile: k8s.DockerfilePinger},
		docker.ImageBuild{Tag: "quay.io/prometheus/prometheus:v2.46.0"},
		docker.ImageBuild{Tag: "otel/opentelemetry-collector-contrib:0.85.0"},
		docker.ImageBuild{Tag: "jaegertracing/all-in-one:latest"},
	); err != nil {
		slog.Error("can't build docker images", err)
		os.Exit(-1)
	}

	cluster = kube.NewKind("test-kind-cluster-daemonset",
		kube.ExportLogs(k8s.PathKindLogs),
		kube.KindConfig(k8s.PathManifests+"/00-kind.yml"),
		kube.LocalImage("testserver:dev"),
		kube.LocalImage("beyla:dev"),
		kube.LocalImage("grpcpinger:dev"),
		kube.LocalImage("quay.io/prometheus/prometheus:v2.46.0"),
		kube.LocalImage("otel/opentelemetry-collector-contrib:0.85.0"),
		kube.LocalImage("jaegertracing/all-in-one:latest"),
		kube.Deploy(k8s.PathManifests+"/01-volumes.yml"),
		kube.Deploy(k8s.PathManifests+"/01-serviceaccount.yml"),
		kube.Deploy(k8s.PathManifests+"/03-otelcol.yml"),
		kube.Deploy(k8s.PathManifests+"/04-jaeger.yml"),
		kube.Deploy(k8s.PathManifests+"/05-uninstrumented-service.yml"),
		kube.Deploy(k8s.PathManifests+"/06-beyla-daemonset.yml"),
	)

	cluster.Run(m)
}
