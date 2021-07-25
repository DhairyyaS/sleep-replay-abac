// File generated by params.SaveGoCode

package main

import "github.com/emer/emergent/params"

var SavedParamsSets = params.Sets{
	{Name: "Base", Desc: "these are the best params", Sheets: params.Sheets{
		"Network": &params.Sheet{
			{Sel: "Prjn", Desc: "keeping default params for generic prjns",
				Params: params.Params{
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "false",
				}},
			{Sel: ".PerCTXPrjn", Desc: "Higer perception to ATL - very low learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.01",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "false",
					"Prjn.WtInit.Mean": "0.1",
					"Prjn.WtInit.Var": "0.5",
					"Prjn.CHL.Hebb": "0",
					//"Prjn.WtScale.Rel": "4",
					//"Prjn.WtScale.Abs": "2",
					//"Prjn.CHL.MinusQ1" : "false",
				}},
			{Sel: ".PerDGPrjn", Desc: "Higher perception to DG - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.1",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.CHL.Hebb": "0",
					"Prjn.WtInit.Mean": "0.5",
					"Prjn.WtInit.Var":  "0.25",
					"Prjn.WtScale.Abs": "1",
				}},
			{Sel: ".PerCA3Prjn", Desc: "Higher perception to MSP - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.1",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.WtScale.Rel":	"2",
					//"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
					//"Prjn.WtInit.Mean": "0.2",
					//"Prjn.WtInit.Var":  "0.5",
					"Prjn.WtInit.Mean": "0.5",
					"Prjn.WtInit.Var":  "0.25",
				}},
			{Sel: "#EXTToInput", Desc: "CA3 to CTX - Strong",
				Params: params.Params{
					"Prjn.Learn.Learn":"false",
					"Prjn.WtInit.Mean":"0.9",
					"Prjn.WtScale.Rel":"10",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					//"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
				}},
			{Sel: "#OutputToInput", Desc: "Output to Input - Strong",
				Params: params.Params{
					"Prjn.Learn.Learn":"false",
					"Prjn.WtInit.Mean":"0.9",
					"Prjn.WtInit.Var":"0",
					"Prjn.WtScale.Rel":"5",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					//"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
				}},
			{Sel: "#CA3ToCTX", Desc: "CA3 to CTX - low learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.0001",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					//"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
				}},
			{Sel: "#CA3ToCA3", Desc: "CA3 recurrent - very high learning rate",
				Params: params.Params{
					"Prjn.Learn.Lrate":       "0.2",
					"Prjn.Learn.Momentum.On": "false",
					"Prjn.Learn.Norm.On":     "false",
					"Prjn.Learn.WtBal.On":    "true",
					"Prjn.CHL.Hebb": "0",
					"Prjn.CHL.SAvgCor": "1",
					"Prjn.WtScale.Rel": "2",
					"Prjn.WtInit.Mean": "0.1",
					"Prjn.WtInit.Var":  "0.5",
				}},
			{Sel: "#DGToCA3", Desc: "DG to CA3 - non-learning",
				Params: params.Params{
					"Prjn.CHL.SAvgCor":"1",
					"Prjn.Learn.Learn":"false",
					"Prjn.WtInit.Mean":"0.9",
					"Prjn.WtInit.Var":"0.01",
					"Prjn.WtScale.Rel":"10",
					//"Prjn.WtScale.Abs":"1",
					"Prjn.Learn.Momentum.On":"false",
					"Prjn.Learn.Norm.On":"false",
				}},
			{Sel: ".Per", Desc: "All Per layers",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.1",
				}},
			{Sel: "#Input", Desc: "Codename only",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "3.2",
				}},
			{Sel: "#Output", Desc: "All Per layers",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.9",
				}},
			{Sel: "#CTX", Desc: "low inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.4",
				}},
			{Sel: "#EXT", Desc: "low inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.1",
				}},
			{Sel: "#DG", Desc: "very sparse = high inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "2.8",
				}},
			{Sel: "#CA3", Desc: "very sparse = high inhibition",
				Params: params.Params{
					"Layer.Inhib.Layer.Gi": "5",
				}},
		},
	}},
}