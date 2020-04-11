package main

// RcBeam type is a beam representation
type RcBeam struct {
	Height                   string `json:"height" binding:"required"`
	Width                    string `json:"width" binding:"required"`
	Cover                    string `json:"cover" binding:"required"`
	TopCover                 string `json:"topCover" binding:"required"`
	ConcreteClass            string `json:"concreteClass" binding:"required"`
	SteelClass               string `json:"steelClass" binding:"required"`
	ConcreteFactor           string `json:"concreteFactor" binding:"required"`
	SteelFactor              string `json:"steelFactor" binding:"required"`
	AlphaCC                  string `json:"alphaCC" binding:"required"`
	LinkDiameter             string `json:"linkDiameter" binding:"required"`
	MainBarDiameter          string `json:"mainBarDiameter" binding:"required"`
	BendingMoment            string `json:"bendingMoment" binding:"required"`
	Fcd                      string `json:"fcd" binding:"required"`
	Fctm                     string `json:"fctm" binding:"required"`
	Fyd                      string `json:"fyd" binding:"required"`
	EffectiveHeight          string `json:"effectiveHeight" binding:"required"`
	MinReinforcement         string `json:"minReinforcement" binding:"required"`
	Sc                       string `json:"sC" binding:"required"`
	KsiEffective             string `json:"ksiEffective" binding:"required"`
	KsiEffectiveLim          string `json:"ksiEffectiveLim" binding:"required"`
	BottomReinforcementPrime string `json:"bottomReinforcementPrime" binding:"required"`
	BendingMomentPrime       string `json:"bendingMomentPrime" binding:"required"`
	BendingMomentDelta       string `json:"bendingMomentDelta" binding:"required"`
	TopReqReinforcement      string `json:"topReqReinforcement" binding:"required"`
	BottomReqReinforcement   string `json:"bottomReqReinforcement" binding:"required"`
	MaximumReinforcement     string `json:"maximumReinforcement" binding:"required"`
}
