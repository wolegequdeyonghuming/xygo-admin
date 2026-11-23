package bizorder

import (
	"testing"

	"xygo/internal/consts"
	adminin "xygo/internal/model/input/orderin"
)

// TestBuildStepDataRecordAgentId 验证录单步骤(1)保存收单员：
//   - 选择收单员时写入 agent_id（保存即派单，前端/后端一步到位）
//   - 未选择收单员时不写入 agent_id
func TestBuildStepDataRecordAgentId(t *testing.T) {
	s := New()
	tests := []struct {
		name    string
		agentId int64
		wantSet bool
	}{
		{name: "录单带收单员，写入 agent_id", agentId: 100, wantSet: true},
		{name: "录单不带收单员，不写入 agent_id", agentId: 0, wantSet: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &adminin.BizOrderStepEditInp{
				Step:         consts.OrderStatusRecorded,
				ScheduleDate: nil,
				CustomerName: "张三",
				ContactPhone: "13800000000",
				AgentId:      tt.agentId,
			}
			data, err := s.buildStepData(in)
			if err != nil {
				t.Fatalf("buildStepData 返回错误: %v", err)
			}
			got, ok := data["agent_id"]
			if tt.wantSet {
				if !ok {
					t.Fatalf("录单步骤应写入 agent_id，实际未写入")
				}
				if got != tt.agentId {
					t.Errorf("agent_id 值错误: 期望 %v, 实际 %v", tt.agentId, got)
				}
			} else {
				if ok {
					t.Errorf("录单步骤未选收单员时不应写入 agent_id，实际为 %v", got)
				}
			}
		})
	}
}

// TestBuildStepDataAssignAgentId 验证派单步骤(2)保存收单员：
//   - 选择收单员时写入 agent_id
//   - 未选择收单员时报错
func TestBuildStepDataAssignAgentId(t *testing.T) {
	s := New()
	tests := []struct {
		name    string
		agentId int64
		wantErr bool
	}{
		{name: "派单带收单员，写入 agent_id", agentId: 200, wantErr: false},
		{name: "派单不带收单员，返回错误", agentId: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &adminin.BizOrderStepEditInp{
				Step:    consts.OrderStatusAssigned,
				AgentId: tt.agentId,
			}
			data, err := s.buildStepData(in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("派单步骤未选收单员应返回错误，实际未返回")
				}
				return
			}
			if err != nil {
				t.Fatalf("buildStepData 返回错误: %v", err)
			}
			if got := data["agent_id"]; got != tt.agentId {
				t.Errorf("agent_id 值错误: 期望 %v, 实际 %v", tt.agentId, got)
			}
		})
	}
}
