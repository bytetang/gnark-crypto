// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gurvy/internal/generators DO NOT EDIT

package curve

// C holds data for a specific curve
// Examples: BLS12-381, BLS12-377, BN256, BW6-761
var C Data

func init() {
	C = Data{
		Fpackage:        "bn256",
		FpModulus:       "21888242871839275222246405745257275088696311157297823662689037894645226208583",
		FrModulus:       "21888242871839275222246405745257275088548364400416034343698204186575808495617",
		Fp2NonResidue:   "-1",
		Fp6NonResidue:   "9,1",
		EmbeddingDegree: 12,
		T:               "4965661367192848881",
		TNeg:            false,
		ThirdRootOne:    "2203960485148121921418603742825762020974279258880205651966",
		Lambda:          "4407920970296243842393367215006156084916469457145843978461",
		Size1:           "65",
		Size2:           "191",
	}
	C.Init()
}
