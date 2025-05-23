/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.3.0-Docker-3.3.0 (4.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// InterfaceRfChannelLabel the model 'InterfaceRfChannelLabel'
type InterfaceRfChannelLabel string

// List of Interface_rf_channel_label
const (
	INTERFACERFCHANNELLABEL__1__2412_MHZ        InterfaceRfChannelLabel = "1 (2412 MHz)"
	INTERFACERFCHANNELLABEL__2__2417_MHZ        InterfaceRfChannelLabel = "2 (2417 MHz)"
	INTERFACERFCHANNELLABEL__3__2422_MHZ        InterfaceRfChannelLabel = "3 (2422 MHz)"
	INTERFACERFCHANNELLABEL__4__2427_MHZ        InterfaceRfChannelLabel = "4 (2427 MHz)"
	INTERFACERFCHANNELLABEL__5__2432_MHZ        InterfaceRfChannelLabel = "5 (2432 MHz)"
	INTERFACERFCHANNELLABEL__6__2437_MHZ        InterfaceRfChannelLabel = "6 (2437 MHz)"
	INTERFACERFCHANNELLABEL__7__2442_MHZ        InterfaceRfChannelLabel = "7 (2442 MHz)"
	INTERFACERFCHANNELLABEL__8__2447_MHZ        InterfaceRfChannelLabel = "8 (2447 MHz)"
	INTERFACERFCHANNELLABEL__9__2452_MHZ        InterfaceRfChannelLabel = "9 (2452 MHz)"
	INTERFACERFCHANNELLABEL__10__2457_MHZ       InterfaceRfChannelLabel = "10 (2457 MHz)"
	INTERFACERFCHANNELLABEL__11__2462_MHZ       InterfaceRfChannelLabel = "11 (2462 MHz)"
	INTERFACERFCHANNELLABEL__12__2467_MHZ       InterfaceRfChannelLabel = "12 (2467 MHz)"
	INTERFACERFCHANNELLABEL__13__2472_MHZ       InterfaceRfChannelLabel = "13 (2472 MHz)"
	INTERFACERFCHANNELLABEL__32__5160_20_MHZ    InterfaceRfChannelLabel = "32 (5160/20 MHz)"
	INTERFACERFCHANNELLABEL__34__5170_40_MHZ    InterfaceRfChannelLabel = "34 (5170/40 MHz)"
	INTERFACERFCHANNELLABEL__36__5180_20_MHZ    InterfaceRfChannelLabel = "36 (5180/20 MHz)"
	INTERFACERFCHANNELLABEL__38__5190_40_MHZ    InterfaceRfChannelLabel = "38 (5190/40 MHz)"
	INTERFACERFCHANNELLABEL__40__5200_20_MHZ    InterfaceRfChannelLabel = "40 (5200/20 MHz)"
	INTERFACERFCHANNELLABEL__42__5210_80_MHZ    InterfaceRfChannelLabel = "42 (5210/80 MHz)"
	INTERFACERFCHANNELLABEL__44__5220_20_MHZ    InterfaceRfChannelLabel = "44 (5220/20 MHz)"
	INTERFACERFCHANNELLABEL__46__5230_40_MHZ    InterfaceRfChannelLabel = "46 (5230/40 MHz)"
	INTERFACERFCHANNELLABEL__48__5240_20_MHZ    InterfaceRfChannelLabel = "48 (5240/20 MHz)"
	INTERFACERFCHANNELLABEL__50__5250_160_MHZ   InterfaceRfChannelLabel = "50 (5250/160 MHz)"
	INTERFACERFCHANNELLABEL__52__5260_20_MHZ    InterfaceRfChannelLabel = "52 (5260/20 MHz)"
	INTERFACERFCHANNELLABEL__54__5270_40_MHZ    InterfaceRfChannelLabel = "54 (5270/40 MHz)"
	INTERFACERFCHANNELLABEL__56__5280_20_MHZ    InterfaceRfChannelLabel = "56 (5280/20 MHz)"
	INTERFACERFCHANNELLABEL__58__5290_80_MHZ    InterfaceRfChannelLabel = "58 (5290/80 MHz)"
	INTERFACERFCHANNELLABEL__60__5300_20_MHZ    InterfaceRfChannelLabel = "60 (5300/20 MHz)"
	INTERFACERFCHANNELLABEL__62__5310_40_MHZ    InterfaceRfChannelLabel = "62 (5310/40 MHz)"
	INTERFACERFCHANNELLABEL__64__5320_20_MHZ    InterfaceRfChannelLabel = "64 (5320/20 MHz)"
	INTERFACERFCHANNELLABEL__100__5500_20_MHZ   InterfaceRfChannelLabel = "100 (5500/20 MHz)"
	INTERFACERFCHANNELLABEL__102__5510_40_MHZ   InterfaceRfChannelLabel = "102 (5510/40 MHz)"
	INTERFACERFCHANNELLABEL__104__5520_20_MHZ   InterfaceRfChannelLabel = "104 (5520/20 MHz)"
	INTERFACERFCHANNELLABEL__106__5530_80_MHZ   InterfaceRfChannelLabel = "106 (5530/80 MHz)"
	INTERFACERFCHANNELLABEL__108__5540_20_MHZ   InterfaceRfChannelLabel = "108 (5540/20 MHz)"
	INTERFACERFCHANNELLABEL__110__5550_40_MHZ   InterfaceRfChannelLabel = "110 (5550/40 MHz)"
	INTERFACERFCHANNELLABEL__112__5560_20_MHZ   InterfaceRfChannelLabel = "112 (5560/20 MHz)"
	INTERFACERFCHANNELLABEL__114__5570_160_MHZ  InterfaceRfChannelLabel = "114 (5570/160 MHz)"
	INTERFACERFCHANNELLABEL__116__5580_20_MHZ   InterfaceRfChannelLabel = "116 (5580/20 MHz)"
	INTERFACERFCHANNELLABEL__118__5590_40_MHZ   InterfaceRfChannelLabel = "118 (5590/40 MHz)"
	INTERFACERFCHANNELLABEL__120__5600_20_MHZ   InterfaceRfChannelLabel = "120 (5600/20 MHz)"
	INTERFACERFCHANNELLABEL__122__5610_80_MHZ   InterfaceRfChannelLabel = "122 (5610/80 MHz)"
	INTERFACERFCHANNELLABEL__124__5620_20_MHZ   InterfaceRfChannelLabel = "124 (5620/20 MHz)"
	INTERFACERFCHANNELLABEL__126__5630_40_MHZ   InterfaceRfChannelLabel = "126 (5630/40 MHz)"
	INTERFACERFCHANNELLABEL__128__5640_20_MHZ   InterfaceRfChannelLabel = "128 (5640/20 MHz)"
	INTERFACERFCHANNELLABEL__132__5660_20_MHZ   InterfaceRfChannelLabel = "132 (5660/20 MHz)"
	INTERFACERFCHANNELLABEL__134__5670_40_MHZ   InterfaceRfChannelLabel = "134 (5670/40 MHz)"
	INTERFACERFCHANNELLABEL__136__5680_20_MHZ   InterfaceRfChannelLabel = "136 (5680/20 MHz)"
	INTERFACERFCHANNELLABEL__138__5690_80_MHZ   InterfaceRfChannelLabel = "138 (5690/80 MHz)"
	INTERFACERFCHANNELLABEL__140__5700_20_MHZ   InterfaceRfChannelLabel = "140 (5700/20 MHz)"
	INTERFACERFCHANNELLABEL__142__5710_40_MHZ   InterfaceRfChannelLabel = "142 (5710/40 MHz)"
	INTERFACERFCHANNELLABEL__144__5720_20_MHZ   InterfaceRfChannelLabel = "144 (5720/20 MHz)"
	INTERFACERFCHANNELLABEL__149__5745_20_MHZ   InterfaceRfChannelLabel = "149 (5745/20 MHz)"
	INTERFACERFCHANNELLABEL__151__5755_40_MHZ   InterfaceRfChannelLabel = "151 (5755/40 MHz)"
	INTERFACERFCHANNELLABEL__153__5765_20_MHZ   InterfaceRfChannelLabel = "153 (5765/20 MHz)"
	INTERFACERFCHANNELLABEL__155__5775_80_MHZ   InterfaceRfChannelLabel = "155 (5775/80 MHz)"
	INTERFACERFCHANNELLABEL__157__5785_20_MHZ   InterfaceRfChannelLabel = "157 (5785/20 MHz)"
	INTERFACERFCHANNELLABEL__159__5795_40_MHZ   InterfaceRfChannelLabel = "159 (5795/40 MHz)"
	INTERFACERFCHANNELLABEL__161__5805_20_MHZ   InterfaceRfChannelLabel = "161 (5805/20 MHz)"
	INTERFACERFCHANNELLABEL__163__5815_160_MHZ  InterfaceRfChannelLabel = "163 (5815/160 MHz)"
	INTERFACERFCHANNELLABEL__165__5825_20_MHZ   InterfaceRfChannelLabel = "165 (5825/20 MHz)"
	INTERFACERFCHANNELLABEL__167__5835_40_MHZ   InterfaceRfChannelLabel = "167 (5835/40 MHz)"
	INTERFACERFCHANNELLABEL__169__5845_20_MHZ   InterfaceRfChannelLabel = "169 (5845/20 MHz)"
	INTERFACERFCHANNELLABEL__171__5855_80_MHZ   InterfaceRfChannelLabel = "171 (5855/80 MHz)"
	INTERFACERFCHANNELLABEL__173__5865_20_MHZ   InterfaceRfChannelLabel = "173 (5865/20 MHz)"
	INTERFACERFCHANNELLABEL__175__5875_40_MHZ   InterfaceRfChannelLabel = "175 (5875/40 MHz)"
	INTERFACERFCHANNELLABEL__177__5885_20_MHZ   InterfaceRfChannelLabel = "177 (5885/20 MHz)"
	INTERFACERFCHANNELLABEL__1__5955_20_MHZ     InterfaceRfChannelLabel = "1 (5955/20 MHz)"
	INTERFACERFCHANNELLABEL__3__5965_40_MHZ     InterfaceRfChannelLabel = "3 (5965/40 MHz)"
	INTERFACERFCHANNELLABEL__5__5975_20_MHZ     InterfaceRfChannelLabel = "5 (5975/20 MHz)"
	INTERFACERFCHANNELLABEL__7__5985_80_MHZ     InterfaceRfChannelLabel = "7 (5985/80 MHz)"
	INTERFACERFCHANNELLABEL__9__5995_20_MHZ     InterfaceRfChannelLabel = "9 (5995/20 MHz)"
	INTERFACERFCHANNELLABEL__11__6005_40_MHZ    InterfaceRfChannelLabel = "11 (6005/40 MHz)"
	INTERFACERFCHANNELLABEL__13__6015_20_MHZ    InterfaceRfChannelLabel = "13 (6015/20 MHz)"
	INTERFACERFCHANNELLABEL__15__6025_160_MHZ   InterfaceRfChannelLabel = "15 (6025/160 MHz)"
	INTERFACERFCHANNELLABEL__17__6035_20_MHZ    InterfaceRfChannelLabel = "17 (6035/20 MHz)"
	INTERFACERFCHANNELLABEL__19__6045_40_MHZ    InterfaceRfChannelLabel = "19 (6045/40 MHz)"
	INTERFACERFCHANNELLABEL__21__6055_20_MHZ    InterfaceRfChannelLabel = "21 (6055/20 MHz)"
	INTERFACERFCHANNELLABEL__23__6065_80_MHZ    InterfaceRfChannelLabel = "23 (6065/80 MHz)"
	INTERFACERFCHANNELLABEL__25__6075_20_MHZ    InterfaceRfChannelLabel = "25 (6075/20 MHz)"
	INTERFACERFCHANNELLABEL__27__6085_40_MHZ    InterfaceRfChannelLabel = "27 (6085/40 MHz)"
	INTERFACERFCHANNELLABEL__29__6095_20_MHZ    InterfaceRfChannelLabel = "29 (6095/20 MHz)"
	INTERFACERFCHANNELLABEL__31__6105_320_MHZ   InterfaceRfChannelLabel = "31 (6105/320 MHz)"
	INTERFACERFCHANNELLABEL__33__6115_20_MHZ    InterfaceRfChannelLabel = "33 (6115/20 MHz)"
	INTERFACERFCHANNELLABEL__35__6125_40_MHZ    InterfaceRfChannelLabel = "35 (6125/40 MHz)"
	INTERFACERFCHANNELLABEL__37__6135_20_MHZ    InterfaceRfChannelLabel = "37 (6135/20 MHz)"
	INTERFACERFCHANNELLABEL__39__6145_80_MHZ    InterfaceRfChannelLabel = "39 (6145/80 MHz)"
	INTERFACERFCHANNELLABEL__41__6155_20_MHZ    InterfaceRfChannelLabel = "41 (6155/20 MHz)"
	INTERFACERFCHANNELLABEL__43__6165_40_MHZ    InterfaceRfChannelLabel = "43 (6165/40 MHz)"
	INTERFACERFCHANNELLABEL__45__6175_20_MHZ    InterfaceRfChannelLabel = "45 (6175/20 MHz)"
	INTERFACERFCHANNELLABEL__47__6185_160_MHZ   InterfaceRfChannelLabel = "47 (6185/160 MHz)"
	INTERFACERFCHANNELLABEL__49__6195_20_MHZ    InterfaceRfChannelLabel = "49 (6195/20 MHz)"
	INTERFACERFCHANNELLABEL__51__6205_40_MHZ    InterfaceRfChannelLabel = "51 (6205/40 MHz)"
	INTERFACERFCHANNELLABEL__53__6215_20_MHZ    InterfaceRfChannelLabel = "53 (6215/20 MHz)"
	INTERFACERFCHANNELLABEL__55__6225_80_MHZ    InterfaceRfChannelLabel = "55 (6225/80 MHz)"
	INTERFACERFCHANNELLABEL__57__6235_20_MHZ    InterfaceRfChannelLabel = "57 (6235/20 MHz)"
	INTERFACERFCHANNELLABEL__59__6245_40_MHZ    InterfaceRfChannelLabel = "59 (6245/40 MHz)"
	INTERFACERFCHANNELLABEL__61__6255_20_MHZ    InterfaceRfChannelLabel = "61 (6255/20 MHz)"
	INTERFACERFCHANNELLABEL__65__6275_20_MHZ    InterfaceRfChannelLabel = "65 (6275/20 MHz)"
	INTERFACERFCHANNELLABEL__67__6285_40_MHZ    InterfaceRfChannelLabel = "67 (6285/40 MHz)"
	INTERFACERFCHANNELLABEL__69__6295_20_MHZ    InterfaceRfChannelLabel = "69 (6295/20 MHz)"
	INTERFACERFCHANNELLABEL__71__6305_80_MHZ    InterfaceRfChannelLabel = "71 (6305/80 MHz)"
	INTERFACERFCHANNELLABEL__73__6315_20_MHZ    InterfaceRfChannelLabel = "73 (6315/20 MHz)"
	INTERFACERFCHANNELLABEL__75__6325_40_MHZ    InterfaceRfChannelLabel = "75 (6325/40 MHz)"
	INTERFACERFCHANNELLABEL__77__6335_20_MHZ    InterfaceRfChannelLabel = "77 (6335/20 MHz)"
	INTERFACERFCHANNELLABEL__79__6345_160_MHZ   InterfaceRfChannelLabel = "79 (6345/160 MHz)"
	INTERFACERFCHANNELLABEL__81__6355_20_MHZ    InterfaceRfChannelLabel = "81 (6355/20 MHz)"
	INTERFACERFCHANNELLABEL__83__6365_40_MHZ    InterfaceRfChannelLabel = "83 (6365/40 MHz)"
	INTERFACERFCHANNELLABEL__85__6375_20_MHZ    InterfaceRfChannelLabel = "85 (6375/20 MHz)"
	INTERFACERFCHANNELLABEL__87__6385_80_MHZ    InterfaceRfChannelLabel = "87 (6385/80 MHz)"
	INTERFACERFCHANNELLABEL__89__6395_20_MHZ    InterfaceRfChannelLabel = "89 (6395/20 MHz)"
	INTERFACERFCHANNELLABEL__91__6405_40_MHZ    InterfaceRfChannelLabel = "91 (6405/40 MHz)"
	INTERFACERFCHANNELLABEL__93__6415_20_MHZ    InterfaceRfChannelLabel = "93 (6415/20 MHz)"
	INTERFACERFCHANNELLABEL__95__6425_320_MHZ   InterfaceRfChannelLabel = "95 (6425/320 MHz)"
	INTERFACERFCHANNELLABEL__97__6435_20_MHZ    InterfaceRfChannelLabel = "97 (6435/20 MHz)"
	INTERFACERFCHANNELLABEL__99__6445_40_MHZ    InterfaceRfChannelLabel = "99 (6445/40 MHz)"
	INTERFACERFCHANNELLABEL__101__6455_20_MHZ   InterfaceRfChannelLabel = "101 (6455/20 MHz)"
	INTERFACERFCHANNELLABEL__103__6465_80_MHZ   InterfaceRfChannelLabel = "103 (6465/80 MHz)"
	INTERFACERFCHANNELLABEL__105__6475_20_MHZ   InterfaceRfChannelLabel = "105 (6475/20 MHz)"
	INTERFACERFCHANNELLABEL__107__6485_40_MHZ   InterfaceRfChannelLabel = "107 (6485/40 MHz)"
	INTERFACERFCHANNELLABEL__109__6495_20_MHZ   InterfaceRfChannelLabel = "109 (6495/20 MHz)"
	INTERFACERFCHANNELLABEL__111__6505_160_MHZ  InterfaceRfChannelLabel = "111 (6505/160 MHz)"
	INTERFACERFCHANNELLABEL__113__6515_20_MHZ   InterfaceRfChannelLabel = "113 (6515/20 MHz)"
	INTERFACERFCHANNELLABEL__115__6525_40_MHZ   InterfaceRfChannelLabel = "115 (6525/40 MHz)"
	INTERFACERFCHANNELLABEL__117__6535_20_MHZ   InterfaceRfChannelLabel = "117 (6535/20 MHz)"
	INTERFACERFCHANNELLABEL__119__6545_80_MHZ   InterfaceRfChannelLabel = "119 (6545/80 MHz)"
	INTERFACERFCHANNELLABEL__121__6555_20_MHZ   InterfaceRfChannelLabel = "121 (6555/20 MHz)"
	INTERFACERFCHANNELLABEL__123__6565_40_MHZ   InterfaceRfChannelLabel = "123 (6565/40 MHz)"
	INTERFACERFCHANNELLABEL__125__6575_20_MHZ   InterfaceRfChannelLabel = "125 (6575/20 MHz)"
	INTERFACERFCHANNELLABEL__129__6595_20_MHZ   InterfaceRfChannelLabel = "129 (6595/20 MHz)"
	INTERFACERFCHANNELLABEL__131__6605_40_MHZ   InterfaceRfChannelLabel = "131 (6605/40 MHz)"
	INTERFACERFCHANNELLABEL__133__6615_20_MHZ   InterfaceRfChannelLabel = "133 (6615/20 MHz)"
	INTERFACERFCHANNELLABEL__135__6625_80_MHZ   InterfaceRfChannelLabel = "135 (6625/80 MHz)"
	INTERFACERFCHANNELLABEL__137__6635_20_MHZ   InterfaceRfChannelLabel = "137 (6635/20 MHz)"
	INTERFACERFCHANNELLABEL__139__6645_40_MHZ   InterfaceRfChannelLabel = "139 (6645/40 MHz)"
	INTERFACERFCHANNELLABEL__141__6655_20_MHZ   InterfaceRfChannelLabel = "141 (6655/20 MHz)"
	INTERFACERFCHANNELLABEL__143__6665_160_MHZ  InterfaceRfChannelLabel = "143 (6665/160 MHz)"
	INTERFACERFCHANNELLABEL__145__6675_20_MHZ   InterfaceRfChannelLabel = "145 (6675/20 MHz)"
	INTERFACERFCHANNELLABEL__147__6685_40_MHZ   InterfaceRfChannelLabel = "147 (6685/40 MHz)"
	INTERFACERFCHANNELLABEL__149__6695_20_MHZ   InterfaceRfChannelLabel = "149 (6695/20 MHz)"
	INTERFACERFCHANNELLABEL__151__6705_80_MHZ   InterfaceRfChannelLabel = "151 (6705/80 MHz)"
	INTERFACERFCHANNELLABEL__153__6715_20_MHZ   InterfaceRfChannelLabel = "153 (6715/20 MHz)"
	INTERFACERFCHANNELLABEL__155__6725_40_MHZ   InterfaceRfChannelLabel = "155 (6725/40 MHz)"
	INTERFACERFCHANNELLABEL__157__6735_20_MHZ   InterfaceRfChannelLabel = "157 (6735/20 MHz)"
	INTERFACERFCHANNELLABEL__159__6745_320_MHZ  InterfaceRfChannelLabel = "159 (6745/320 MHz)"
	INTERFACERFCHANNELLABEL__161__6755_20_MHZ   InterfaceRfChannelLabel = "161 (6755/20 MHz)"
	INTERFACERFCHANNELLABEL__163__6765_40_MHZ   InterfaceRfChannelLabel = "163 (6765/40 MHz)"
	INTERFACERFCHANNELLABEL__165__6775_20_MHZ   InterfaceRfChannelLabel = "165 (6775/20 MHz)"
	INTERFACERFCHANNELLABEL__167__6785_80_MHZ   InterfaceRfChannelLabel = "167 (6785/80 MHz)"
	INTERFACERFCHANNELLABEL__169__6795_20_MHZ   InterfaceRfChannelLabel = "169 (6795/20 MHz)"
	INTERFACERFCHANNELLABEL__171__6805_40_MHZ   InterfaceRfChannelLabel = "171 (6805/40 MHz)"
	INTERFACERFCHANNELLABEL__173__6815_20_MHZ   InterfaceRfChannelLabel = "173 (6815/20 MHz)"
	INTERFACERFCHANNELLABEL__175__6825_160_MHZ  InterfaceRfChannelLabel = "175 (6825/160 MHz)"
	INTERFACERFCHANNELLABEL__177__6835_20_MHZ   InterfaceRfChannelLabel = "177 (6835/20 MHz)"
	INTERFACERFCHANNELLABEL__179__6845_40_MHZ   InterfaceRfChannelLabel = "179 (6845/40 MHz)"
	INTERFACERFCHANNELLABEL__181__6855_20_MHZ   InterfaceRfChannelLabel = "181 (6855/20 MHz)"
	INTERFACERFCHANNELLABEL__183__6865_80_MHZ   InterfaceRfChannelLabel = "183 (6865/80 MHz)"
	INTERFACERFCHANNELLABEL__185__6875_20_MHZ   InterfaceRfChannelLabel = "185 (6875/20 MHz)"
	INTERFACERFCHANNELLABEL__187__6885_40_MHZ   InterfaceRfChannelLabel = "187 (6885/40 MHz)"
	INTERFACERFCHANNELLABEL__189__6895_20_MHZ   InterfaceRfChannelLabel = "189 (6895/20 MHz)"
	INTERFACERFCHANNELLABEL__193__6915_20_MHZ   InterfaceRfChannelLabel = "193 (6915/20 MHz)"
	INTERFACERFCHANNELLABEL__195__6925_40_MHZ   InterfaceRfChannelLabel = "195 (6925/40 MHz)"
	INTERFACERFCHANNELLABEL__197__6935_20_MHZ   InterfaceRfChannelLabel = "197 (6935/20 MHz)"
	INTERFACERFCHANNELLABEL__199__6945_80_MHZ   InterfaceRfChannelLabel = "199 (6945/80 MHz)"
	INTERFACERFCHANNELLABEL__201__6955_20_MHZ   InterfaceRfChannelLabel = "201 (6955/20 MHz)"
	INTERFACERFCHANNELLABEL__203__6965_40_MHZ   InterfaceRfChannelLabel = "203 (6965/40 MHz)"
	INTERFACERFCHANNELLABEL__205__6975_20_MHZ   InterfaceRfChannelLabel = "205 (6975/20 MHz)"
	INTERFACERFCHANNELLABEL__207__6985_160_MHZ  InterfaceRfChannelLabel = "207 (6985/160 MHz)"
	INTERFACERFCHANNELLABEL__209__6995_20_MHZ   InterfaceRfChannelLabel = "209 (6995/20 MHz)"
	INTERFACERFCHANNELLABEL__211__7005_40_MHZ   InterfaceRfChannelLabel = "211 (7005/40 MHz)"
	INTERFACERFCHANNELLABEL__213__7015_20_MHZ   InterfaceRfChannelLabel = "213 (7015/20 MHz)"
	INTERFACERFCHANNELLABEL__215__7025_80_MHZ   InterfaceRfChannelLabel = "215 (7025/80 MHz)"
	INTERFACERFCHANNELLABEL__217__7035_20_MHZ   InterfaceRfChannelLabel = "217 (7035/20 MHz)"
	INTERFACERFCHANNELLABEL__219__7045_40_MHZ   InterfaceRfChannelLabel = "219 (7045/40 MHz)"
	INTERFACERFCHANNELLABEL__221__7055_20_MHZ   InterfaceRfChannelLabel = "221 (7055/20 MHz)"
	INTERFACERFCHANNELLABEL__225__7075_20_MHZ   InterfaceRfChannelLabel = "225 (7075/20 MHz)"
	INTERFACERFCHANNELLABEL__227__7085_40_MHZ   InterfaceRfChannelLabel = "227 (7085/40 MHz)"
	INTERFACERFCHANNELLABEL__229__7095_20_MHZ   InterfaceRfChannelLabel = "229 (7095/20 MHz)"
	INTERFACERFCHANNELLABEL__233__7115_20_MHZ   InterfaceRfChannelLabel = "233 (7115/20 MHz)"
	INTERFACERFCHANNELLABEL__1__58_32_2_16_GHZ  InterfaceRfChannelLabel = "1 (58.32/2.16 GHz)"
	INTERFACERFCHANNELLABEL__2__60_48_2_16_GHZ  InterfaceRfChannelLabel = "2 (60.48/2.16 GHz)"
	INTERFACERFCHANNELLABEL__3__62_64_2_16_GHZ  InterfaceRfChannelLabel = "3 (62.64/2.16 GHz)"
	INTERFACERFCHANNELLABEL__4__64_80_2_16_GHZ  InterfaceRfChannelLabel = "4 (64.80/2.16 GHz)"
	INTERFACERFCHANNELLABEL__5__66_96_2_16_GHZ  InterfaceRfChannelLabel = "5 (66.96/2.16 GHz)"
	INTERFACERFCHANNELLABEL__6__69_12_2_16_GHZ  InterfaceRfChannelLabel = "6 (69.12/2.16 GHz)"
	INTERFACERFCHANNELLABEL__9__59_40_4_32_GHZ  InterfaceRfChannelLabel = "9 (59.40/4.32 GHz)"
	INTERFACERFCHANNELLABEL__10__61_56_4_32_GHZ InterfaceRfChannelLabel = "10 (61.56/4.32 GHz)"
	INTERFACERFCHANNELLABEL__11__63_72_4_32_GHZ InterfaceRfChannelLabel = "11 (63.72/4.32 GHz)"
	INTERFACERFCHANNELLABEL__12__65_88_4_32_GHZ InterfaceRfChannelLabel = "12 (65.88/4.32 GHz)"
	INTERFACERFCHANNELLABEL__13__68_04_4_32_GHZ InterfaceRfChannelLabel = "13 (68.04/4.32 GHz)"
	INTERFACERFCHANNELLABEL__17__60_48_6_48_GHZ InterfaceRfChannelLabel = "17 (60.48/6.48 GHz)"
	INTERFACERFCHANNELLABEL__18__62_64_6_48_GHZ InterfaceRfChannelLabel = "18 (62.64/6.48 GHz)"
	INTERFACERFCHANNELLABEL__19__64_80_6_48_GHZ InterfaceRfChannelLabel = "19 (64.80/6.48 GHz)"
	INTERFACERFCHANNELLABEL__20__66_96_6_48_GHZ InterfaceRfChannelLabel = "20 (66.96/6.48 GHz)"
	INTERFACERFCHANNELLABEL__25__61_56_8_64_GHZ InterfaceRfChannelLabel = "25 (61.56/8.64 GHz)"
	INTERFACERFCHANNELLABEL__26__63_72_8_64_GHZ InterfaceRfChannelLabel = "26 (63.72/8.64 GHz)"
	INTERFACERFCHANNELLABEL__27__65_88_8_64_GHZ InterfaceRfChannelLabel = "27 (65.88/8.64 GHz)"
)

// All allowed values of InterfaceRfChannelLabel enum
var AllowedInterfaceRfChannelLabelEnumValues = []InterfaceRfChannelLabel{
	"1 (2412 MHz)",
	"2 (2417 MHz)",
	"3 (2422 MHz)",
	"4 (2427 MHz)",
	"5 (2432 MHz)",
	"6 (2437 MHz)",
	"7 (2442 MHz)",
	"8 (2447 MHz)",
	"9 (2452 MHz)",
	"10 (2457 MHz)",
	"11 (2462 MHz)",
	"12 (2467 MHz)",
	"13 (2472 MHz)",
	"32 (5160/20 MHz)",
	"34 (5170/40 MHz)",
	"36 (5180/20 MHz)",
	"38 (5190/40 MHz)",
	"40 (5200/20 MHz)",
	"42 (5210/80 MHz)",
	"44 (5220/20 MHz)",
	"46 (5230/40 MHz)",
	"48 (5240/20 MHz)",
	"50 (5250/160 MHz)",
	"52 (5260/20 MHz)",
	"54 (5270/40 MHz)",
	"56 (5280/20 MHz)",
	"58 (5290/80 MHz)",
	"60 (5300/20 MHz)",
	"62 (5310/40 MHz)",
	"64 (5320/20 MHz)",
	"100 (5500/20 MHz)",
	"102 (5510/40 MHz)",
	"104 (5520/20 MHz)",
	"106 (5530/80 MHz)",
	"108 (5540/20 MHz)",
	"110 (5550/40 MHz)",
	"112 (5560/20 MHz)",
	"114 (5570/160 MHz)",
	"116 (5580/20 MHz)",
	"118 (5590/40 MHz)",
	"120 (5600/20 MHz)",
	"122 (5610/80 MHz)",
	"124 (5620/20 MHz)",
	"126 (5630/40 MHz)",
	"128 (5640/20 MHz)",
	"132 (5660/20 MHz)",
	"134 (5670/40 MHz)",
	"136 (5680/20 MHz)",
	"138 (5690/80 MHz)",
	"140 (5700/20 MHz)",
	"142 (5710/40 MHz)",
	"144 (5720/20 MHz)",
	"149 (5745/20 MHz)",
	"151 (5755/40 MHz)",
	"153 (5765/20 MHz)",
	"155 (5775/80 MHz)",
	"157 (5785/20 MHz)",
	"159 (5795/40 MHz)",
	"161 (5805/20 MHz)",
	"163 (5815/160 MHz)",
	"165 (5825/20 MHz)",
	"167 (5835/40 MHz)",
	"169 (5845/20 MHz)",
	"171 (5855/80 MHz)",
	"173 (5865/20 MHz)",
	"175 (5875/40 MHz)",
	"177 (5885/20 MHz)",
	"1 (5955/20 MHz)",
	"3 (5965/40 MHz)",
	"5 (5975/20 MHz)",
	"7 (5985/80 MHz)",
	"9 (5995/20 MHz)",
	"11 (6005/40 MHz)",
	"13 (6015/20 MHz)",
	"15 (6025/160 MHz)",
	"17 (6035/20 MHz)",
	"19 (6045/40 MHz)",
	"21 (6055/20 MHz)",
	"23 (6065/80 MHz)",
	"25 (6075/20 MHz)",
	"27 (6085/40 MHz)",
	"29 (6095/20 MHz)",
	"31 (6105/320 MHz)",
	"33 (6115/20 MHz)",
	"35 (6125/40 MHz)",
	"37 (6135/20 MHz)",
	"39 (6145/80 MHz)",
	"41 (6155/20 MHz)",
	"43 (6165/40 MHz)",
	"45 (6175/20 MHz)",
	"47 (6185/160 MHz)",
	"49 (6195/20 MHz)",
	"51 (6205/40 MHz)",
	"53 (6215/20 MHz)",
	"55 (6225/80 MHz)",
	"57 (6235/20 MHz)",
	"59 (6245/40 MHz)",
	"61 (6255/20 MHz)",
	"65 (6275/20 MHz)",
	"67 (6285/40 MHz)",
	"69 (6295/20 MHz)",
	"71 (6305/80 MHz)",
	"73 (6315/20 MHz)",
	"75 (6325/40 MHz)",
	"77 (6335/20 MHz)",
	"79 (6345/160 MHz)",
	"81 (6355/20 MHz)",
	"83 (6365/40 MHz)",
	"85 (6375/20 MHz)",
	"87 (6385/80 MHz)",
	"89 (6395/20 MHz)",
	"91 (6405/40 MHz)",
	"93 (6415/20 MHz)",
	"95 (6425/320 MHz)",
	"97 (6435/20 MHz)",
	"99 (6445/40 MHz)",
	"101 (6455/20 MHz)",
	"103 (6465/80 MHz)",
	"105 (6475/20 MHz)",
	"107 (6485/40 MHz)",
	"109 (6495/20 MHz)",
	"111 (6505/160 MHz)",
	"113 (6515/20 MHz)",
	"115 (6525/40 MHz)",
	"117 (6535/20 MHz)",
	"119 (6545/80 MHz)",
	"121 (6555/20 MHz)",
	"123 (6565/40 MHz)",
	"125 (6575/20 MHz)",
	"129 (6595/20 MHz)",
	"131 (6605/40 MHz)",
	"133 (6615/20 MHz)",
	"135 (6625/80 MHz)",
	"137 (6635/20 MHz)",
	"139 (6645/40 MHz)",
	"141 (6655/20 MHz)",
	"143 (6665/160 MHz)",
	"145 (6675/20 MHz)",
	"147 (6685/40 MHz)",
	"149 (6695/20 MHz)",
	"151 (6705/80 MHz)",
	"153 (6715/20 MHz)",
	"155 (6725/40 MHz)",
	"157 (6735/20 MHz)",
	"159 (6745/320 MHz)",
	"161 (6755/20 MHz)",
	"163 (6765/40 MHz)",
	"165 (6775/20 MHz)",
	"167 (6785/80 MHz)",
	"169 (6795/20 MHz)",
	"171 (6805/40 MHz)",
	"173 (6815/20 MHz)",
	"175 (6825/160 MHz)",
	"177 (6835/20 MHz)",
	"179 (6845/40 MHz)",
	"181 (6855/20 MHz)",
	"183 (6865/80 MHz)",
	"185 (6875/20 MHz)",
	"187 (6885/40 MHz)",
	"189 (6895/20 MHz)",
	"193 (6915/20 MHz)",
	"195 (6925/40 MHz)",
	"197 (6935/20 MHz)",
	"199 (6945/80 MHz)",
	"201 (6955/20 MHz)",
	"203 (6965/40 MHz)",
	"205 (6975/20 MHz)",
	"207 (6985/160 MHz)",
	"209 (6995/20 MHz)",
	"211 (7005/40 MHz)",
	"213 (7015/20 MHz)",
	"215 (7025/80 MHz)",
	"217 (7035/20 MHz)",
	"219 (7045/40 MHz)",
	"221 (7055/20 MHz)",
	"225 (7075/20 MHz)",
	"227 (7085/40 MHz)",
	"229 (7095/20 MHz)",
	"233 (7115/20 MHz)",
	"1 (58.32/2.16 GHz)",
	"2 (60.48/2.16 GHz)",
	"3 (62.64/2.16 GHz)",
	"4 (64.80/2.16 GHz)",
	"5 (66.96/2.16 GHz)",
	"6 (69.12/2.16 GHz)",
	"9 (59.40/4.32 GHz)",
	"10 (61.56/4.32 GHz)",
	"11 (63.72/4.32 GHz)",
	"12 (65.88/4.32 GHz)",
	"13 (68.04/4.32 GHz)",
	"17 (60.48/6.48 GHz)",
	"18 (62.64/6.48 GHz)",
	"19 (64.80/6.48 GHz)",
	"20 (66.96/6.48 GHz)",
	"25 (61.56/8.64 GHz)",
	"26 (63.72/8.64 GHz)",
	"27 (65.88/8.64 GHz)",
}

func (v *InterfaceRfChannelLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceRfChannelLabel(value)
	for _, existing := range AllowedInterfaceRfChannelLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceRfChannelLabel", value)
}

// NewInterfaceRfChannelLabelFromValue returns a pointer to a valid InterfaceRfChannelLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceRfChannelLabelFromValue(v string) (*InterfaceRfChannelLabel, error) {
	ev := InterfaceRfChannelLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceRfChannelLabel: valid values are %v", v, AllowedInterfaceRfChannelLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceRfChannelLabel) IsValid() bool {
	for _, existing := range AllowedInterfaceRfChannelLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_rf_channel_label value
func (v InterfaceRfChannelLabel) Ptr() *InterfaceRfChannelLabel {
	return &v
}

type NullableInterfaceRfChannelLabel struct {
	value *InterfaceRfChannelLabel
	isSet bool
}

func (v NullableInterfaceRfChannelLabel) Get() *InterfaceRfChannelLabel {
	return v.value
}

func (v *NullableInterfaceRfChannelLabel) Set(val *InterfaceRfChannelLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceRfChannelLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceRfChannelLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceRfChannelLabel(val *InterfaceRfChannelLabel) *NullableInterfaceRfChannelLabel {
	return &NullableInterfaceRfChannelLabel{value: val, isSet: true}
}

func (v NullableInterfaceRfChannelLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceRfChannelLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
