package types

// RcBeam is a beam struct for latex template
type RcBeam struct {
	Height                   float64 `form:"height" binding:"exists"`
	Width                    float64 `form:"width" binding:"exists"`
	Cover                    float64 `form:"cover" binding:"exists"`
	TopCover                 float64 `form:"topCover" binding:"exists"`
	ConcreteClass            float64 `form:"concreteClass" binding:"exists"`
	SteelClass               float64 `form:"steelClass" binding:"exists"`
	ConcreteFactor           float64 `form:"concreteFactor" binding:"exists"`
	SteelFactor              float64 `form:"steelFactor" binding:"exists"`
	AlphaCC                  float64 `form:"alphaCC" binding:"exists"`
	LinkDiameter             float64 `form:"linkDiameter" binding:"exists"`
	MainBarDiameter          float64 `form:"mainBarDiameter" binding:"exists"`
	BendingMoment            float64 `form:"bendingMoment" binding:"exists"`
	Fcd                      float64 `form:"fcd" binding:"exists"`
	Fctm                     float64 `form:"fctm" binding:"exists"`
	Fyd                      float64 `form:"fyd" binding:"exists"`
	EffectiveHeight          float64 `form:"effectiveHeight" binding:"exists"`
	MinReinforcement         float64 `form:"minReinforcement" binding:"exists"`
	Sc                       float64 `form:"sC" binding:"exists"`
	KsiEffective             float64 `form:"ksiEffective" binding:"exists"`
	KsiEffectiveLim          float64 `form:"ksiEffectiveLim" binding:"exists"`
	BottomReinforcementPrime float64 `form:"bottomReinforcementPrime" binding:"exists"`
	BendingMomentPrime       float64 `form:"bendingMomentPrime" binding:"exists"`
	BendingMomentDelta       float64 `form:"bendingMomentDelta" binding:"exists"`
	TopReqReinforcement      float64 `form:"topReqReinforcement" binding:"exists"`
	BottomReqReinforcement   float64 `form:"bottomReqReinforcement" binding:"exists"`
	MaximumReinforcement     float64 `form:"maximumReinforcement" binding:"exists"`
}
