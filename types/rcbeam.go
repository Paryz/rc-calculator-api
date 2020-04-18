package types

// RcBeam is a beam struct for latex template
type RcBeam struct {
	Height                   float64 `form:"height" binding:"required"`
	Width                    float64 `form:"width" binding:"required"`
	Cover                    float64 `form:"cover" binding:"required"`
	TopCover                 float64 `form:"topCover" binding:"required"`
	ConcreteClass            float64 `form:"concreteClass" binding:"required"`
	SteelClass               float64 `form:"steelClass" binding:"required"`
	ConcreteFactor           float64 `form:"concreteFactor" binding:"required"`
	SteelFactor              float64 `form:"steelFactor" binding:"required"`
	AlphaCC                  float64 `form:"alphaCC" binding:"required"`
	LinkDiameter             float64 `form:"linkDiameter" binding:"required"`
	MainBarDiameter          float64 `form:"mainBarDiameter" binding:"required"`
	BendingMoment            float64 `form:"bendingMoment" binding:"required"`
	Fcd                      float64 `form:"fcd" binding:"required"`
	Fctm                     float64 `form:"fctm" binding:"required"`
	Fyd                      float64 `form:"fyd" binding:"required"`
	EffectiveHeight          float64 `form:"effectiveHeight" binding:"required"`
	MinReinforcement         float64 `form:"minReinforcement" binding:"required"`
	Sc                       float64 `form:"sC" binding:"required"`
	KsiEffective             float64 `form:"ksiEffective" binding:"required"`
	KsiEffectiveLim          float64 `form:"ksiEffectiveLim" binding:"required"`
	BottomReinforcementPrime float64 `form:"bottomReinforcementPrime" binding:"required"`
	BendingMomentPrime       float64 `form:"bendingMomentPrime" binding:"required"`
	BendingMomentDelta       float64 `form:"bendingMomentDelta" binding:"required"`
	TopReqReinforcement      float64 `form:"topReqReinforcement" binding:"required"`
	BottomReqReinforcement   float64 `form:"bottomReqReinforcement" binding:"required"`
	MaximumReinforcement     float64 `form:"maximumReinforcement" binding:"required"`
}
