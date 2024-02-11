package invest_spec

import "testing"

func TestMaxAddNumber_IsSatisfy(t *testing.T) {
	type fields struct {
		AddCount  int
		MaxNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				AddCount:  10,
				MaxNumber: 5,
			},
			want: false,
		},
		{
			fields: fields{
				AddCount:  1,
				MaxNumber: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MaxAddNumber{
				AddCount:  tt.fields.AddCount,
				MaxNumber: tt.fields.MaxNumber,
			}
			if got := m.IsSatisfy(); got != tt.want {
				t.Errorf("IsSatisfy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSum_IsSatisfy(t *testing.T) {
	type fields struct {
		CurrentCount int
		AddCount     int
		MaxSumNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				CurrentCount: 10,
				AddCount:     10,
				MaxSumNumber: 20,
			},
			want: true,
		},
		{
			fields: fields{
				CurrentCount: 10,
				AddCount:     10,
				MaxSumNumber: 19,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MaxSum{
				CurrentCount: tt.fields.CurrentCount,
				AddCount:     tt.fields.AddCount,
				MaxSumNumber: tt.fields.MaxSumNumber,
			}
			if got := m.IsSatisfy(); got != tt.want {
				t.Errorf("IsSatisfy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinAddNumber_IsSatisfy(t *testing.T) {
	type fields struct {
		AddCount  int
		MinNumber int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				AddCount:  10,
				MinNumber: 11,
			},
			want: false,
		},
		{
			fields: fields{
				AddCount:  10,
				MinNumber: 9,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MinAddNumber{
				AddCount:  tt.fields.AddCount,
				MinNumber: tt.fields.MinNumber,
			}
			if got := m.IsSatisfy(); got != tt.want {
				t.Errorf("IsSatisfy() = %v, want %v", got, tt.want)
			}
		})
	}
}
