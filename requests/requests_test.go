package requests

import (
	"reflect"
	"testing"
)

func TestInternalRequestOptions(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		options *RequestOptions
		want    RequestOptions
		wantErr bool
	}{
		{
			name: "Test with default options",
			url:  "http://example.com",
			options: &RequestOptions{
				HTTPMethod: "GET",
				Headers:    Headers{"Content-Type": "application/json"},
				Params:     Params{"key": "value"},
			},
			want: RequestOptions{
				HTTPMethod:       "GET",
				RawResponse:      false,
				Authentication:   nil,
				HideToken:        false,
				Portal:           "",
				Credentials:      "",
				MaxURLLength:     0,
				Signal:           nil,
				SuppressWarnings: false,
				Params:           Params{"key": "value", "f": "json"},
				Headers:          Headers{"Content-Type": "application/json"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := internalRequestOptions(tt.url, tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Function call fail internalRequestOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("internalRequestOptions() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
