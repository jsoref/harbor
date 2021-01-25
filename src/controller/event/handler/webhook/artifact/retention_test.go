package artifact

import (
	"github.com/goharbor/harbor/src/common/dao"
	"github.com/goharbor/harbor/src/controller/retention"
	ret "github.com/goharbor/harbor/src/pkg/retention"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
	"time"

	"github.com/goharbor/harbor/src/controller/event"
	"github.com/goharbor/harbor/src/core/config"
	"github.com/goharbor/harbor/src/lib/selector"
	"github.com/goharbor/harbor/src/pkg/notification"
	retentiontesting "github.com/goharbor/harbor/src/testing/controller/retention"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRetentionHandler_Handle(t *testing.T) {
	config.Init()
	handler := &RetentionHandler{}

	policyMgr := notification.PolicyMgr
	oldretentionCtl := retention.Ctl

	defer func() {
		notification.PolicyMgr = policyMgr
		retention.Ctl = oldretentionCtl
	}()
	notification.PolicyMgr = &fakedNotificationPolicyMgr{}
	retentionCtl := &retentiontesting.Controller{}
	retention.Ctl = retentionCtl
	retentionCtl.On("GetRetentionExecTask", mock.Anything, mock.Anything).
		Return(&ret.Task{
			ID:          1,
			ExecutionID: 1,
			Status:      "Success",
			StartTime:   time.Now(),
			EndTime:     time.Now(),
		}, nil)
	retentionCtl.On("GetRetentionExec", mock.Anything, mock.Anything).Return(&ret.Execution{
		ID:        1,
		PolicyID:  1,
		Status:    "Success",
		Trigger:   "Manual",
		DryRun:    true,
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}, nil)

	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "RetentionHandler Want Error 1",
			args: args{
				data: "",
			},
			wantErr: true,
		},
		{
			name: "RetentionHandler 1",
			args: args{
				data: &event.RetentionEvent{
					OccurAt: time.Now(),
					Deleted: []*selector.Result{
						{
							Target: &selector.Candidate{
								NamespaceID: 1,
								Namespace:   "project1",
								Tags:        []string{"v1"},
								Labels:      nil,
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.Handle(tt.args.data)
			if tt.wantErr {
				require.NotNil(t, err, "Error: %s", err)
				return
			}
			assert.Nil(t, err)
		})
	}

}

func TestRetentionHandler_IsStateful(t *testing.T) {
	handler := &RetentionHandler{}
	assert.False(t, handler.IsStateful())
}

func TestMain(m *testing.M) {
	dao.PrepareTestForPostgreSQL()
	os.Exit(m.Run())
}
