package classify

import (
	"constraint-system/core"
	"math"
	"testing"
)

func TestPythagorean(t *testing.T) {
	_ = PythagoreanTheorem
	tests := []struct {
		Edge3 float64
		Edge2 float64
		Edge1 float64
	}{
		{
			Edge1: 5,
			Edge2: 4,
			Edge3: 3,
		},
		{
			Edge1: 13,
			Edge2: 12,
			Edge3: 5,
		},
	}
	for _, tt := range tests {
		t.Run("勾股定理", func(t *testing.T) {
			PythagoreanTheoremEdge2.SetValue(core.PredefineUserSource, tt.Edge2)
			PythagoreanTheoremEdge3.SetValue(core.PredefineUserSource, tt.Edge3)
			if tt.Edge1 != PythagoreanTheoremEdge1.GetValue() {
				t.Errorf("failed! edge3:%f,edge2:%f,edge1:%f,result:%f", tt.Edge3, tt.Edge2, tt.Edge1, PythagoreanTheoremEdge3.GetValue())
			}
			PythagoreanTheoremEdge2.ForgotValue(core.PredefineUserSource)
			PythagoreanTheoremEdge3.ForgotValue(core.PredefineUserSource)
			// ************************************************* //
			PythagoreanTheoremEdge1.SetValue(core.PredefineUserSource, tt.Edge1)
			PythagoreanTheoremEdge2.SetValue(core.PredefineUserSource, tt.Edge2)
			if tt.Edge3 != PythagoreanTheoremEdge3.GetValue() {
				t.Errorf("failed! edge3:%f,edge2:%f,edge1:%f,result:%f", tt.Edge3, tt.Edge2, tt.Edge1, PythagoreanTheoremEdge3.GetValue())
			}
			PythagoreanTheoremEdge1.ForgotValue(core.PredefineUserSource)
			PythagoreanTheoremEdge2.ForgotValue(core.PredefineUserSource)
			// ************************************************* //
			PythagoreanTheoremEdge1.SetValue(core.PredefineUserSource, tt.Edge1)
			PythagoreanTheoremEdge3.SetValue(core.PredefineUserSource, tt.Edge3)
			if tt.Edge2 != PythagoreanTheoremEdge2.GetValue() {
				t.Errorf("failed! edge3:%f,edge2:%f,edge1:%f,result:%f", tt.Edge3, tt.Edge2, tt.Edge1, PythagoreanTheoremEdge3.GetValue())
			}
			PythagoreanTheoremEdge1.ForgotValue(core.PredefineUserSource)
			PythagoreanTheoremEdge3.ForgotValue(core.PredefineUserSource)
		})
	}
}

func TestCircleAreaExpr(t *testing.T) {
	_ = circleAreaExpr
	tests := []struct {
		Radius float64
		Area   float64
	}{
		{
			Radius: 2,
			Area:   12.57,
		},
	}

	for _, tt := range tests {
		t.Run("圆的面积", func(t *testing.T) {
			circleAreaExprRadius.SetValue(core.PredefineUserSource, tt.Radius)
			tmp := math.Round(circleAreaExprArea.GetValue() * 100)
			if tt.Area != tmp/100 {
				t.Errorf("failed! raduis:%f,area:%f,,result:%f", tt.Radius, tt.Area, circleAreaExprArea.GetValue())
			}
		})
	}
}
