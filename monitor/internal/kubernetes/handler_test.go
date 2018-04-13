package kubernetesmonitor

import (
	"context"
	"testing"

	kubernetesclient "github.com/aporeto-inc/trireme-kubernetes/kubernetes"
	"github.com/aporeto-inc/trireme-lib/common"
	"github.com/aporeto-inc/trireme-lib/monitor/config"
	"github.com/aporeto-inc/trireme-lib/monitor/extractors"
	dockermonitor "github.com/aporeto-inc/trireme-lib/monitor/internal/docker"
	"github.com/aporeto-inc/trireme-lib/policy"
	kubecache "k8s.io/client-go/tools/cache"
)

func Test_getKubernetesInformation(t *testing.T) {

	puRuntimeWithTags := func(tags map[string]string) *policy.PURuntime {
		puRuntime := policy.NewPURuntimeWithDefaults()
		puRuntime.SetTags(policy.NewTagStoreFromMap(tags))
		return puRuntime
	}

	type args struct {
		runtime policy.RuntimeReader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:    "no Kubernetes Information",
			args:    args{runtime: policy.NewPURuntimeWithDefaults()},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "both present",
			args: args{runtime: puRuntimeWithTags(map[string]string{
				KubernetesPodNamespaceIdentifier: "a",
				KubernetesPodNameIdentifier:      "b",
			},
			),
			},
			want:    "a",
			want1:   "b",
			wantErr: false,
		},
		{
			name: "both present. NamespaceIdentifier empty",
			args: args{runtime: puRuntimeWithTags(map[string]string{
				KubernetesPodNamespaceIdentifier: "",
				KubernetesPodNameIdentifier:      "b",
			},
			),
			},
			want:    "",
			want1:   "b",
			wantErr: false,
		},
		{
			name: "both present. Name empty",
			args: args{runtime: puRuntimeWithTags(map[string]string{
				KubernetesPodNamespaceIdentifier: "a",
				KubernetesPodNameIdentifier:      "",
			},
			),
			},
			want:    "a",
			want1:   "",
			wantErr: false,
		},
		{
			name: "Namespace missing",
			args: args{runtime: puRuntimeWithTags(map[string]string{
				KubernetesPodNameIdentifier: "b",
			},
			),
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
		{
			name: "Name missing",
			args: args{runtime: puRuntimeWithTags(map[string]string{
				KubernetesPodNamespaceIdentifier: "a",
			},
			),
			},
			want:    "",
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getKubernetesInformation(tt.args.runtime)
			if (err != nil) != tt.wantErr {
				t.Errorf("getKubernetesInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getKubernetesInformation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getKubernetesInformation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestKubernetesMonitor_HandlePUEvent(t *testing.T) {
	type fields struct {
		dockerMonitor       *dockermonitor.DockerMonitor
		kubernetesClient    *kubernetesclient.Client
		handlers            *config.ProcessorConfig
		cache               *cache
		kubernetesExtractor extractors.KubernetesMetadataExtractorType
		podStore            kubecache.Store
		podController       kubecache.Controller
		podControllerStop   chan struct{}
		enableHostPods      bool
	}
	type args struct {
		ctx           context.Context
		puID          string
		event         common.Event
		dockerRuntime policy.RuntimeReader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "empty dockerruntime on create",
			fields: fields{},
			args: args{
				event:         common.EventCreate,
				dockerRuntime: policy.NewPURuntimeWithDefaults(),
			},
			wantErr: true,
		},
		{
			name:   "empty dockerruntime on start",
			fields: fields{},
			args: args{
				event:         common.EventCreate,
				dockerRuntime: policy.NewPURuntimeWithDefaults(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &KubernetesMonitor{
				dockerMonitor:       tt.fields.dockerMonitor,
				kubernetesClient:    tt.fields.kubernetesClient,
				handlers:            tt.fields.handlers,
				cache:               tt.fields.cache,
				kubernetesExtractor: tt.fields.kubernetesExtractor,
				podStore:            tt.fields.podStore,
				podController:       tt.fields.podController,
				podControllerStop:   tt.fields.podControllerStop,
				enableHostPods:      tt.fields.enableHostPods,
			}
			if err := m.HandlePUEvent(tt.args.ctx, tt.args.puID, tt.args.event, tt.args.dockerRuntime); (err != nil) != tt.wantErr {
				t.Errorf("KubernetesMonitor.HandlePUEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
