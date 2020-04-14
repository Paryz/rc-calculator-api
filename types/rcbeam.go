package types

// RcBeam is a beam struct for latex template
type RcBeam struct {
	Height                   string `form:"height" binding:"required"`
	Width                    string `form:"width" binding:"required"`
	Cover                    string `form:"cover" binding:"required"`
	TopCover                 string `form:"topCover" binding:"required"`
	ConcreteClass            string `form:"concreteClass" binding:"required"`
	SteelClass               string `form:"steelClass" binding:"required"`
	ConcreteFactor           string `form:"concreteFactor" binding:"required"`
	SteelFactor              string `form:"steelFactor" binding:"required"`
	AlphaCC                  string `form:"alphaCC" binding:"required"`
	LinkDiameter             string `form:"linkDiameter" binding:"required"`
	MainBarDiameter          string `form:"mainBarDiameter" binding:"required"`
	BendingMoment            string `form:"bendingMoment" binding:"required"`
	Fcd                      string `form:"fcd" binding:"required"`
	Fctm                     string `form:"fctm" binding:"required"`
	Fyd                      string `form:"fyd" binding:"required"`
	EffectiveHeight          string `form:"effectiveHeight" binding:"required"`
	MinReinforcement         string `form:"minReinforcement" binding:"required"`
	Sc                       string `form:"sC" binding:"required"`
	KsiEffective             string `form:"ksiEffective" binding:"required"`
	KsiEffectiveLim          string `form:"ksiEffectiveLim" binding:"required"`
	BottomReinforcementPrime string `form:"bottomReinforcementPrime" binding:"required"`
	BendingMomentPrime       string `form:"bendingMomentPrime" binding:"required"`
	BendingMomentDelta       string `form:"bendingMomentDelta" binding:"required"`
	TopReqReinforcement      string `form:"topReqReinforcement" binding:"required"`
	BottomReqReinforcement   string `form:"bottomReqReinforcement" binding:"required"`
	MaximumReinforcement     string `form:"maximumReinforcement" binding:"required"`
}
