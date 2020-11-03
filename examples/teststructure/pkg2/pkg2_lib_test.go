/*
Test file for pkg2 file
*/

package pkg2

import "testing"

// generated test for MultipleNamedOutputValues
func TestMultipleNamedOutputValues(t *testing.T) {
	var tests []struct {
		name      string
		wantStr1  string
		wantStr2  string
		wantBool1 bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr1, gotStr2, gotBool1 := MultipleNamedOutputValues()
			if gotStr1 != tt.wantStr1 {
				t.Errorf("MultipleNamedOutputValues() gotStr1 = %v, want %v", gotStr1, tt.wantStr1)
			}
			if gotStr2 != tt.wantStr2 {
				t.Errorf("MultipleNamedOutputValues() gotStr2 = %v, want %v", gotStr2, tt.wantStr2)
			}
			if gotBool1 != tt.wantBool1 {
				t.Errorf("MultipleNamedOutputValues() gotBool1 = %v, want %v", gotBool1, tt.wantBool1)
			}
		})
	}
}