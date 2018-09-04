package mmq

import "testing"

func TestBaseQueue_BConsume(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		wantMsg string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields:  fields{Name: "qtest"},
			wantMsg: "v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bq := &BaseQueue{
				Name: tt.fields.Name,
			}
			gotMsg, err := bq.BConsume(10)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseQueue.BConsume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("BaseQueue.BConsume() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
