package config

var BW6_756 = Curve{
	Name:         "bw6-756",
	CurvePackage: "bw6756",
	EnumID:       "BW6_756",
	FrModulus:    "605248206075306171733248481581800960739847691770924913753520744034740935903401304776283802348837311170974282940417",
	FpModulus:    "366325390957376286590726555727219947825377821289246188278797409783441745356050456327989347160777465284190855125642086860525706497928518803244008749360363712553766506755227344593404398783886857865261088226271336335268413437902849",
	G1: Point{
		CoordType:        "fp.Element",
		CoordExtDegree:   1,
		PointName:        "g1",
		GLV:              true,
		CofactorCleaning: true,
		CRange:           []int{4, 5, 8, 16},
		Projective:       true,
	},
	G2: Point{
		CoordType:        "fp.Element",
		CoordExtDegree:   1,
		PointName:        "g2",
		GLV:              true,
		CofactorCleaning: true,
		CRange:           []int{4, 5, 8, 16},
	},
	// 2-isogeny
	HashE1: &HashSuite{
		A: []string{"0xf76adbb5bb98ade21e8fbe4eef81b6e9756798f2e64bff3cb65781179d6b076b17683f3cd5042655e16802b1a5b1b5b5e386e23e2731d24c843595d5c79ca4b8b9179cf10cb86e782d614a1f78930c25aeacdc30c0008fb417dfffffffff2"},
		B: []string{"0x16"},
		Z: []int{11},
		Isogeny: &Isogeny{
			XMap: RationalPolynomial{
				Num: [][]string{
					{"0xb99024c84cb2829c6f231ad6488f8ecbe73b35d177c03c95a1b5d4aff107d954726713be7dc02432d0d2d3f919a6b98b6ff46f95026f313a3f953d1776abdcc48d809cb0a92e2b7464ed6eab9312f612a13505b1d000072f6798000000000"},
					{"0x26bd1df0d7ae6c65ccce9c94a71497cd125c27c5fea6d502de92ee64cfe6161e0e05eb4aa16bb7f4c55960f84e03d6b986c05315e0dc296ea61a6ce0bfe7584381b90f97adb14083c7e63f8683ffffb35baf0000000001"},
					{"0xb99024c84cb282a010dde96a80e9b8571a99e3c121ae77cf5a598f3fd0abd199502d6d31fb5237042160e2f83bbff87df05586dc52cb529ee19d07248b4fbf241ffad1c2a6de71c88e46e4e3dbb1026d5ecafa4e300000000000000000001"},
				},
				Den: [][]string{
					{"0x9af477c35eb9b197333a72529c525f3449709f17fa9b540b7a4bb9933f9858783817ad2a85aedfd3156583e1380f5ae61b014c578370a5ba9869b382ff9d610e06e43e5eb6c5020f1f98fe1a0ffffecd6ebc0000000004"},
				},
			},
			YMap: RationalPolynomial{
				Num: [][]string{
					{"0xb99024c84cb282a010dde96a80e9b8571a99e3c121ae77cf5a598f3fd0abd199502d6d31fb5237042160e2f83bbff87df05586dc52cb529ee19d07248b4fbf241ffad1c2a6de71c88e46e4e3dbb1026d5ecafa4e2ffffffffffffffffffff"},
					{"0x1eed5b76b77315ce6c7800aef7b306952f8658ccae70a80831fd94f251e13a45b7ccc7290e7ae2e14ef34b51df3471733e5650ac56b2e140baada4fc20270074f385fcf816086d73d46b785d5a289f4a69c36293f7ffee097d04000000002"},
					{"0x7bb56ddaddcc5719024ebf85e3a0a46fefc545c5c0628b194a34c4ba6ac12eab1339f794cfc8e22966cea64f49ee8f4675ef712f878e58793870797ac6d90c77a7cc163e6cef3cd9dd88b97adb140df8fcc7f0d07ffff8d09868000000002"},
					{"0xd87d803f042598656902e5a6ebbb571049b389b6a74b8bc73ebdd1ca73731f32dd8a54ba4fdfeada26f108cc45b54c92edb91d566097e064073732fff7dd09aa254f4a0dc2ae2f69fb52b5b4804e82d4ee97795b380000000000000000001"},
				},
				Den: [][]string{
					{"0xf76adbb5bb98ae2ac127e1e3568cf5c978cd2fac2ce89fbf23221455163a6ccc6ae73c42a46d9eb02c812ea04faaa0a7eb1cb3d06e646e292cd15edb646a54302aa3c258de7ded0b685e868524ec033c7e63f8683fffffffffffffffffff9"},
					{"0x3a1bace94385a298b335eadefa9ee3b39b8a3ba8fdfa3f844ddc659737d9212d1508e0eff22193ef280611747505c2164a207ca0d14a3e25f927a3511fdb0465429597638489e0c5abd95f49c5ffff8d09868000000000c"},
					{"0x1d0dd674a1c2d14c599af56f7d4f71d9cdc51dd47efd1fc226ee32cb9bec90968a847077f910c9f7940308ba3a82e10b25103e5068a51f12fc93d1a88fed8232a14acbb1c244f062d5ecafa4e2ffffc684c34000000000c"},
				},
			},
		},
	},
	// 17-isogeny
	HashE2: &HashSuite{
		A: []string{"0x2693fe3f93094f745f2f924da1a4260675024d398700e46db4678d6c77e732927310a52e43b73de56540c9dde1f7896d115a7661f0e25e7c54754324cb2ab353f985c8bdaccf0b3ae0e6ee49e40dd2b5d8aee87356b60e17f9de343033534"},
		B: []string{"0x4577fc0edab719863d99d99fa7737b795dd19e05021cf50e2bfaadf8d3cf9670ba09ab81773af807f37f714475a819c9dd47d9d9b93ec37c70f3b857252cd99cac4151e0b42eb00271779d66c0b0c79e7f450a1fd85b66a23256679146fcb"},
		Z: []int{11},
		Isogeny: &Isogeny{
			XMap: RationalPolynomial{
				Num: [][]string{
					{"0x57b49e22a638e7ce6fd84adb98fef2ee37a078dad5e6db136e00515b7db476b3092fe24c634db1c19b32784bd06db0e1bdbcc194bd2ea84b5485485423516adaf53426ca54ef27fe07de0370f3473b05fab20bba48a8ceeb8f39138734a79"},
					{"0x4eb16246adf77c244d7c96da07c828728ec8cf4f73b36f9d3b8638cd7d5f00932de546ce6d666b6c1c78a99b824336da2aa7cf8017843619aba11c4e20eed632ffbd4ed04676972f223e842b7d9f408585f879993e979cefce07f6b320f98"},
					{"0xb0fe2b7d5cc6757f09eb475d4ec4eb2892eec65ebb4d0b73dffe388dc8b625ae1094586e7d5f12d55eae19f405323e54346d6e75fe5e30ed86a66f3bd956f2e0cf3b68afc427260f3782828672bc312081de7f76b025cb18780f79ccbb931"},
					{"0x63ae87b32561b93fb20fe607e846b376b2db30eb8774a0c741bc2c2525053672f753b6ea1b420fc0817c91332c754ee91bbdf25a5e5fb852314dc28a26aef1a6fc6bae25a0285b8355a97b597283baeb152f6a0ff281a7a6029551336d19a"},
					{"0x8e95e735b9b8a53ec343d0960e4ee6f8d38aaadd78fb413533fb771a3be16d1dfb91d669d0fdf927aa93c60a147c23c6b26b66cab157b3b1381e64646aae66dfc53e2698852470de8fb92e17b4504bc43249a9a2e7330e3b5e25da416f46f"},
					{"0x5f0f6d5f0ff7ce8777840d1c4d25d85f9a63f04df2479c5b73f1a95a3caf20143e0bf4e9412a9d9d63de7b03be54cfe59522fa2405cd2b02340479f9c8776d6ea9fca2dea8d6d71f455162a6e9f40b76311257170313a8682db61095c7bb5"},
					{"0xce64c0bb237310b05120e54a6ce7ca932dc2d6411b9761463cf0c31e07d4caf5bac19db735631219b68c1dd00e79f19e2653d050137a01218f5b1371d9a276a8938dfd36a0565662a1c60e970b0ba6ca1e330ca446d4863f5c44cf57dfe70"},
					{"0x5c79bf1abe91cef7276af031994044074bd33da0fdde5d1356095b3d02fa2d46345e935838131a84e7408ef41afb492c5a5cf8256daa4be842a0c4e9056f1f3eae94b3d47a78956718cec82b21f3e51338e5e2dea165bb52ed3c3fef55e2 "},
					{"0x528d835a2e6a30781df335a708c54aeb2b71dea9a27564944950b4eb0d8298a44b86f039ac8815883395a340deb72664ce9ef1d01799fe56b8141ffa816ff7ce020b37944db6c438b78a9e78e96ccbc63cff0f7f7e9366ee17fcce48ded23"},
					{"0x778a6ab417356eb65d79c0a0a1cf6d67f3de72168b1ac6c13f3f619fbb33e1da96b6b4b7f07ab0d82d1702006321158b231b405cfbffe44d033ba590dd218224ea7c800dfa1f00dfc41bc8d4f7d6ae966107bf5589f8ba66cfa90c23216c3"},
					{"0x4147349d36fc97c1ff5e168c56d665d0b6f1a8a048eb903e134490eac3b6d62e9c4741b77e0290a585bfa7a1c0dc29fe0b8000728ae4d7f01a849fde08e5c268914f0dee8ed42612e1531f958cf84126358cd261871bded7381b133541604"},
					{"0xc471c9e9bbb09990d5e6c9f5017fe1382bcb0835464ce01b69391ec611babbeca074b376ad1d5ee9c49a2214a89824f408b8eb1dbf80b9851997a2d7f4fc288fb2ee7c3c4fa9d52283e977c924af4c360c5332fce61c340d6af7edef6cfff"},
					{"0x44d18369bb41e08dea957a700307d9929a58b5565bb4fc9b985b6e98047cf473b636f6afa8ddf5cbf053e02477a616339eaef3bdf402e782850a2dc76555e82e8e803f2aa39f76249929725c6ca5d292d4b10aee3535e3c9d9368f01216da"},
					{"0xd46505c4491681e80a79fca2f3899ddefc5c765972f1ff4f4d20c6b8b89b0be5decd3e700f31ca2a0a24759dab83fd694a850902b3e14d656e4864b6fc2bf1f4aac08ed4fd147884840cb0cea7f94b0f923fb94ce65f77b5992c71e1e5038"},
					{"0x6f9632f9f15470710c2004132ef3ad5f5600087a3195f06fa2a1e69c0aea80a94e43c427cd681cd23040668e3aef649d09455fa213ab6364c58327073c8c798963672d092a60943e08d17f05771a42e6abc0d30ef3f1ec16195ff2e6d455d"},
					{"0x9a51e637dd926867d26d4fea699ddbf56041fb4767c71b5cef365a2a944cb79879168c5a93b9fde6be4d414a4016b01b70c039253548cc08b8309388a1a511ebf1dc8f8b9c436cb21aa944c9240a0c1834fc199be1153717ea4ad6f5cf86 "},
					{"0x52852129f19b369553182b651e1a006b282bec349935ab9fd2b392c7e43de42da08efb80da828e5bdb508c3a396e6e22a4f596da57126b2c654700cca80014fab35dd74dabc755ea253556e3d6a77d3d93f85b80b56fa4d4011970a59dcc4"},
					{"0x1c4076fbf2dfd41be9c3e3c24375f6dc8ab18fa163e8f502566471ef2432f6472cfb668c7b4d2e7753eb5181083607b70cac77be553df58c3b26f757f1c70c45474e0c72ac735c9e176ee3f4a10142ce372df6465e1e1e1e1e1e1e1e1e1e2"},
				},
				Den: [][]string{
					{"0x67ca0a05ba18be2717328e9da3d366da18fe119dbf0c150dd6e49daa6ed4f8de92e85a04bc69f045366c80877116759af4bc0c2a6f49ec158f0a2ba5e118ba6d26ac9f2821aa917978d4a606d940397818481adb5ab3f2c9670d2faac21bf"},
					{"0x461d07431e8397832b6ae980dc75548271ce6a87465b5ea5292316b8b1eb078b69e551472f41df43806e24ad887fcb1f551938ec0efc00499f7c3ba5bf7d8ec6d63a4ddd274aad64d716307d11a0ddb9d1c0f40296317c29455075bffb2ba"},
					{"0x43635d242d3fb85cfd3a36bfa1e3628afb5e1e14ae9f67f5ab735f8f6723bc48ba9c507c50e703220244876a482f00f4d7f2d022eccac8301995542df93cc151c10b4a622565832b8793de295ab227054118d673efecd6c50fea7b111606e"},
					{"0x86095d084f72a9be4c483ccc2c746e6ef6e26720fe2700df898dd5a5b20d6a832e2025d2d00a8447279691db55bea69a97e63e4acff6b5bc3254c3fa00017cb0adf4980a85af177485637c02ccc568a182ae232f54eada9391e3d98ed5ffe"},
					{"0xb017fcf26b80fc4af09015d106339becd5762ccb32eae4554d8eac08772a9335ccb7a8a916bd84caf4bcd4d5e53ad1ddf931deda6d5bc04c2df1682d31c696bf2321a2f005e77db814de79b98353ea4625962b9e30e78311ec0d9825e1f61"},
					{"0x45fb599a0f0a1dc49bb19803f84cb2e31e0c7c36dbdd187ab5bd8f024061a702d3e9dea95dc499b77fc32f4824fbaf91d82433dafbec16e1c977b20efd359e1334357a6a5b2d4b5c9452e665fc212908db0d18cfc0b5a5f3706b47c8868d1"},
					{"0x6b33147c7484eee118a499d769e8cd5067e77bae02abf17f8a47b040d39891089ffc19052ed287404aa3e3d2eb7787cd08bb0bfcc1c990a73d649cbfeadd6c4cf9ae2580101c1a1f45a90d32b0ff13e02b5836642634c7e952b6486d29f8 "},
					{"0x387e570f75f5fa085d262c1cb47c27f872beec94798787b8ea11debe75533ea5b0031eb544968899daacb97792dc67039eabcb6e64559a202f2529af243b2d7f993cf4896d9cd1af66b9cdf2753e483999115c0caa96403d360aa9841aaae"},
					{"0x1257e076f478ad94029dfb12ae9d928ce4550151c5a02aac92329f2ecfbdbd0a41567dc38e5fe5a0d5d3adcee3d1d84639145de0de9644244c34c031bd5d564a483ca617084e2829bd8e898733dfb99cc3cb655161d5854f687d9592c67de"},
					{"0xcb293e9849238ea68db183154dd3318b7a97d2c8b5911164a377e42468284ac8f489f97435e2ae5587efc9390ec8fe2dd55a466388491e21b477cf483333fa4f481bade0f0a80dc35c4d967dad4a5aaa1864323a8ab11f376eb6ee6440cd5"},
					{"0xecfd07101cf66ec654369b1dab1bb19588285083358b3f41fbcfa9a8667bee9f2e009a65b4588590ba71ba35bfbf3f11e0e38cfc02bf8e01475be7ce2fce4a6fcc8aaa3aae1d4696fcf86deff15d543c992e08aa4ef0822fa2da7d3dd82b9"},
					{"0x505e4b45c2fd0610b6de37cee0cb9a6e5434fe87b45190ff65418b316e4852a031583c3b5c28b4f78e5cdebbd7ac1675f1bbabb6bbd07ffad253349ed13c6a1e605856fc83288739254e877b0cca8c895a2119d0c9371fc59deded45cb1b4"},
					{"0x4356cbf741137297e9995a239f126c8e69898333c8a0891ee9ae4acb9f66cd1fb8b5c550d54b3d3516d27fc5ebb31b938d50d18a5d1fd92c4fbc246ae3a12886bab0e3b38692d26b5cc13372389c7c65df91ba04dcb040f2c8356bfb78783"},
					{"0x2d30af585d5961e587be65e1b61869b980d0a9a66510b0013800455621c06ab17bbf7cdb49df67795dba802f4a5dc215300a3a409408de7801f2f0a8e5b3587d478f94c523087fda1b7d9fce00d99ad124ae782dd5a6033e14141ae7af925"},
					{"0x89ce51dbfa7cf438f4174452a0159618713a55916e86698c00023059f53812346d2e996cd2be2ce603938b6e28add4df16fac31411400aaab469b668d9fc0efe45b04c071a2e09243a6f1c24fd2927823ba1b82803796b4c5ee8d14b6eca "},
					{"0x60380c3366f74e8a5f5447e6867e4d6b0ca3c6d01e5ad7bfaff113c053f5cada2ab35379044733a1e57cd19cf5ae16240e7bde52e621adad86a454c400381cf47f872d5b7ccf190cddc4a5477a932bd4a5e22534d30913553db82af7238e4"},
				},
			},
			YMap: RationalPolynomial{
				Num: [][]string{
					{"0x9f93a11a37f333a5052a497f2ff8959bc7b3e69bd5f470445f01d6071d144dbe992d4ea3783f0fe7d017f5dcc1e6d2980590726773e6b4bd7bc10cd999e2c38c2aabb90c484462c378b64918c1a48e51921fb837ac5feba926c3f623c4160"},
					{"0x577ef753befab81d08fb0ce7402dc357e7dcd98cd569530135f3ad339845cfcceb84f041cab594864b56060469cc8e0cefa322613061cf50bb7f8b3adba1caa0a522a5df2481f64cc0fbde989a13c70a8168d6cd40a30a342c186f229597 "},
					{"0xa68f14dfeb538700dd4d74482406ffae6be6396e04519ff32a9d0962b7250cd4a6a14a17fbc36032b9a79cb3f79b06d9918b4a34b8302120baf4757568bd3987d5bad8222ea01116c54237a31599acdc600517d3f733f5150930c07a9b8ac"},
					{"0x62f40704a8cef5860a7c1a0e953d8ab65ae5e0cb6e49d58da6e2501eabf797a77b043818a3bfa545a51060e24c3b7c2773412e527e1b72acf2044d95548b72836851a635e50c4c24b35c5a769b985fe464d857d16cecf2f8060374067589b"},
					{"0x1930bec080bcf38ef774f04e675193fc123f266b80e62e1b294b0978abeb31b3afcb02e33f65309399d12e3803fa052098b473eea25be9b9064c8821c0bbb90d9eba711b1dfddf00c2ae407d79382ccad4c6df7d43c160ed257138fdbd403"},
					{"0x65e5a193d3016af59101e156a7a886a3d15bfebab72190e437dcd4ee1a42a17f739c140aead0cf987e922daaf716fb7b36b61e0996fe6abaa4656b8a7f5b87f41ab28135a032b51957e955ce8efb83242cc62fe08a71dd8065c66720c000d"},
					{"0x7efada35b7158ff27d88e0ebb8775bb5e3c7f4224e20a223bb2acc78e0176802b998265a7d6473b22171bc9a57bf3943ebf43f71fd0f9006e5a7af6e6e76f11b949e7adb6b7d252bb4c36338ad539a9a7c77cea84280e7cd2964e5591d226"},
					{"0x29eb281af9b44d21986ba03590afd087e477f3cad3234825f131a31eab2adb59ed830aeab5e5156f8c5de90f81c12d8231604c9e25c4b358501edef95ed2926a79bf118f22d9c4f6455c8c7e37d618c9bc16c46c974437daf821bbd839df9"},
					{"0x47e261e740726acbe699eac8c486186ec6492f0adc2b3dc99615e874b4b7e093bcdafe1075392c49bca1ed7d4d043e2798c971e73535be10e8f4d6aa23b3bed69de7e75a8d3a8373e8ab0621582fb7cc43d2643095843d47003dccbecfc7d"},
					{"0xbd7d4e63943690a235389a6bb5d448042414a1fd9dcf42116515f6f7e5dbaf168a9830e70255aefd2bc14ca35eeb922039ab85ba718d3e6d23d86b08381e413ce18ab5c378c43675af0695ee62eb42790c5d5f89c2d0db6460678d5d7d2ed"},
					{"0x41c2f299e34019cb40e2af5a6bded90f03c9e6e0cbdd853e783f4f3d95935d39ee31fbb61d17610387b84ba9519c16881183a72c59fe7383aa5ae217b42b7480f8f058ad146edf6b4a7549cccea967ded9fd28d150274cee458208907ced2"},
					{"0x405da38064d76408c3adde67a954882f3c8fd36eabe3201a9dfd75629c8f4036c1d63831453109f54cfbc4cc8b9e5f663d3db51eefcefbce9b533b82e005bd9ee7af344ca3dcc40775a3fd24231ad6aa8ec93782ca5005c54856f0d85c04a"},
					{"0x5bafa78c01429499eed6e25e9349a6ca3d2e0ac1267e137174d00a6c433c6095e3d1f312fef70a5703698371ddd1c06e302261b4f6069ec4b78ad4e3dcb37b5253c0787566ba225638b446092283e42cbe85907b5ebff3b33c270a868f0c "},
					{"0x39c152b20238115dd4e6df5508cab1e2c8a7ee9eefdc292f058b6f89586654c8b3586b6f76e713ba51824ca6e2e3ac7a381a8f1746d7d5e500312a7c8513c46a5736425da606f7b9d34fb95df8e54ba0e7cdaef0079df1eced56133412cd "},
					{"0xfeac73088d841e8e555bfc8117068298c2aa3b88c0172c8339c4c507dc212472eccc77267a6946c0c4c260135b4a5f0665194fdf129bc15fc9623987abc4b4d251feb6696737b390dd48da1ee8eae6123b350c4356ab64e534709643ac84 "},
					{"0x6a9a250e9bafa906b1ee1517ad74a816a62c5706d13b5a198fbefb1d2d3dbda735841e67a0adc0ad30495670bd2b83953c65b19aa017ee928b6d5d195b3d16590edf554767224c79b07e803db87acc6a54b60d71b15eb75702f9168c5181e"},
					{"0xe1aeadc4bbd21621ff33dcea2e0ec99a6ea667741cf62af0494afde045b4ccb713836169de252f3018df60a9b42972f3ce650e5fd93a8e06e4fc281fead12a631e834d5753f175a5d76c2c3582caf46cdce6228ce2e767d59573c748ffe55"},
					{"0x7ed20aa5207c66a812970c473f270e99e5974d86d49be1e1976f669388cd568cf799f5c57d9a251960bad6c97dd3fc335eb51b4852cca361a95fc3eff218be4b9b46bfebfc552521659092156b85da783cc46a24ba5d1c1af9ddf017443b0"},
					{"0x7235fb0eb41c06b940ccf7b11788831523a2347562d30c2b147b7381f37a4a63dbae4794c8bd7d0f681bf17aa86221277496fd257af5814f73c8303e1ba09fb80ecf6627b00d453dc712e67aec504f7ece5f39d54f49f6212e792ec1b67dc"},
					{"0xba8193a4d3130aa424d39453c76c8ed262b2e2e0d07f453061b253ce779e1ba0a03f90d00a58f8d08c681204aa771512b68cd12d2dd4e7ee320514a2d5f56a1b8dd2a33c1d4d60a503cade7e0eafd341d8ac667b1c8ccee77ee263808a2c0"},
					{"0xf58592db221a5b17116ecd2bf86c7cb449dd623e49bf58ccb63baaf1cfbaae68778caa8b2f5d9b9211636f79b7ef12b9aa2155c4ef99276190af85a5121b9beccd7c747e153187c668b1846a179312b7fdccb9fd8273ffd3b058c2a97e99f"},
					{"0x2fb957ed355d384286b2f6153caad9da4cd75da2f960d2b0162badab4e6bb653f08a4ba1c93813fd4bdc0ed7ff19073696c16685bc73c7924f3f094e19db49cdce7365a991e22911fe27b1ddf631c5f8775210eee3de31057c27f397c5db2"},
					{"0x9a597c03b1977869ca1a58b04c6c7b5a8136d673428d8077c01fef3fcf1b7661dd2c9a0a9f258d60ac5cddbaddb91dc17e171cb4a48d32d3eaf5be484430496052e4e835781bbdb2f52aa838da2b488143dc2266adbcc6dad98b7bc422c65"},
					{"0xe19735baa44196f6aa3be94274abd8e8669acdf6e9d985820bf1adf34d66305e0a9c697ae0996b67513553a1a5b6e2c465895925873497d56cb7796a3a3fb58bea27d51180a658bcd7d81c7a86577ed8ed84b61484b70703a5be7ae16fce5"},
					{"0x3be0b32a85b89ede867e519b45827587ac1c9b3201628e5a499574e5da2edccadf907d9f793c735da5ff0fc3e60d3550653e133c3d1b467b5951ac65f8ac41b50e3a2e76024288847449f0d35d73d783c6a1c167f67a3e01c5894d10d4986"},
				},
				Den: [][]string{
					{"0x6fce5ce514cee5debaed37973056b6a3edfe727211c6e6a1fab570eed4f2a508e96d6b081ff4a1b1a8acdf79fd0bd33c663045955b8cc064b8b92b7bbfd15fec96bb04be80bae195d8d9d51c240630f776a199d13ecad6bc53f75189d614d"},
					{"0xe0962cd218121ad29e5d7490bf192c86380cb1f4aa91a0e3a86aa73be92a1c3e8a440a35f0248ee807f47a4239a421000433a638d113dca21313ab532984ee2abf3b38d1170fc51c7ae131c57a7449c65cfdd78de9c937528723d0712be77"},
					{"0x1ccc8c1978c03d019911928c1fafaafbfd603ad5590c6854d54c3ea722f18a19020e256e2d2f66418fd103b42765a5195c910f10097fe452c4ba4b5d5ce5a67ee3adde1a6b2cdf3b7721da38f676e97788cc4c64e2e7962fb81bbdc0f28ba"},
					{"0xe22f7cbf189dff65136339d090ea28fc60f120cfa26dd8433c6960b418362be2336be5e929dcc61beb4412dcb232db75be87f17b02742c2206620c090b9799e3ed9d988b69316875970f19cf14388a4fde13f176355fd16b702a250a65024"},
					{"0xe8f52882ce0a0e52ca01be901b93d21fdffae3d31b385750e422cc08cf926c7958a67728e2d26aae7ebf673a840d9d1f8aeb9e58a379f690461f8100817b8185a3e87401b650d473bea9733277bd2166d9f1d7f6fbe6711012bb354d2e636"},
					{"0x875fe99541a3b4273eb0e06ca68021e7b484844fe1276a3dd6c1ca17fff7bfa81a7fe367de671cbfe1498baea048a898dfd77b2d7ffbd4bbf204f862d8771bd35197f0575f6ddaf2dd8915747726e2aae3ddd49c58e47d732e7dc74aaccfa"},
					{"0x4dfd80356811a35a65acb00c46d07a392658904a4fc0cbda29046c44a833a31ed9850230c5f4ed0c1e1e76576ad104d193292821897138fdb5700b90c70041f6f5570cba9b4f4d7def07d469b4c892e87d2ec709ba01068510b911e9e8ad3"},
					{"0xa7f230aa84f2547cbe501721f030ed5bc4376f897cc22d1df86da38d73b27c3903e3e44109d24de87f66a029b5cfd5f68fa440f57de9533515f9912fc130ee39f65d7fb8da98d707cbcb268a3f32cb3b40ab7dcac9128d916a20625e28985"},
					{"0x7f35f1b0caeeffdd8711317415c54cfc5a1a04ab3cd08e6b4415b78b5760ce5dc99acd7c5a4a4c1238c478461e7b79c4f09bc4b7c9c42468fabcfdf142db1403d5b01a4cd818594dfaf729a18514e62aeb7a3733765836d2a9d893325c7a9"},
					{"0x9bce845eb9ff707815ed7ed3fb74d199d492671d29a3bf2072594b9154f530a311514634ef4fb471e883c854dde519b85f679d07aea459f7c3870ba31b1bd2c3548152cc0ce70241c24572d7eb65069b1a97f2bd155a4185ef91ea5659b13"},
					{"0xf0e88310ee09722be8558b6ee541afb4441bc1e77ae82ebcd5d792f6330980c3d36914d73138f0da01be4580c0c2188c0e2a010746b596d80755e71d07028f78bc29af82e7e5be87fb3430b1696810ea95758f3924d846cc0192a1e4f6e10"},
					{"0x8237fb8a7d102cfa45e7020363d892aa7380d1b5ee84b792c65fc2b167552c67623578c416e9922772f3ec99fddaef25907d91bbf0c72b8f0a0070a4d635a1ccdc78a8c6d4a175ed00a68a1eb0e7b90a6a5f33c6555486c77d4aed026ebaa"},
					{"0xacc8c17e87078f38e1056204eea4bd9343903c88966d0ddc2ca5166a97a99db4b66f74480eaf6030d42f8f23f00f1e7a1a6d13f086a76f04f14e98f12917de8d7998089c48bd860301813a95f708c6bd0417acadf80cb9831ee34286cfea8"},
					{"0xa9c3c892d49a5c3e506bfa242431f2b9f3934c51c40045e40478e1f4120da439ba6f0cc154e80bf1ea7ced993b7310545b6af11a6b3bbab611e648afe77297a776ef9468f51706bfbf60a100b53571a0a7361cb592a4f3000a45c013c950f"},
					{"0xeeade4edfeeeef2bc9f34296787829c99b544f11722d55db8789550c5e005ae27d651137cbdf9cb95ca030e3d84544dd21d7dca37bcab9f0e1ddc4cf633b5986746b0371648bac79367a9512396e9789031acc705cbff8c6d15d828c6bdd8"},
					{"0x6b6c82e8a5d0faec7af39bbaa40d86ea47acaac61910aaf01c238d8ac0f2587e2f993835b77e0312a3d23a2da61cfe7d68c91af4d0e24dfa29d5ce2c470c29c5ff6d89f060a0c0546daa1287fa3ab18a20cc1666ad84e3343d25997269773"},
					{"0xdb9102798323f938aa773da2a56ebce58aee6da13e2b10b8a56bea0398ad2d5fa0585d35f7b8b0e9076577c6e36ccf6f9d12e9018bbc8424d38620b874e24d4265bbf27fe0940ddd713e2bf225f0cac7a82c92516c75fc6d88cdbbdd58ec0"},
					{"0x3ad0edad38f3ac9e11e7c078c6e80cf48879547ea9efe567d0fcf53151b71697643a4940c375e6593b96a3370c56964ea3fc16bb574197886c0d9da0064eeb7e4121296069a4509b10f97a4b4a4943d083dbe0a221465711128d1778f332 "},
					{"0xc23a66e6ae7cebc4e55e5a301aa0bb92932e809fc326fc2a9e5ca5711241c8c77313b93e8065e8ecfeb3765a53e2b48bbbd6e27b960e76d00dd5f212f51c2381f6571a0a2d1bdb232ae1fea083b5573ebc76b53c5e8c53434a3e0ba80c3e4"},
					{"0x9dea557e9d9328f93a7a014265d3e7bb51164ef0a4eee3dd1219b918c457f7f0818521d498b13552933f9ead5b0b48d6b3db282c72ec06d7ca1721d3732f3fb59c6c655f3526db0cc833984a0f4957f312abf88133767cad4485c2f9b9af2"},
					{"0x9497a4421814871b03a96cb5d6b294880db80435d857c2fe9f4e97d75b37c795a1728121d66ce742793f98090d0c894f715a7e81dce84f594369b8bb2b9be0d482436c1ceffdc50efae8d85be760beecec788b5f1b08180dfe69567a0c102"},
					{"0x9ca415b1068cafcb60219e17f1350a1ba427a276689247978356bc51e19fbc7bde4a260c33e175dfc83ab230d17ed24c88d8de8c145213227d443192eeb2e1e4c61b7be341bd88638d5b44cedf93d445b344552c004f4b02bcc04b4c235e5"},
					{"0x2a27787477dddb1679fedf03a1e406b8df8cee954e12fd31e027d2d20f5dd6ffe6dc92a5c9ab16be8f21439232809c6187cdae5a5c6fb2f48617d22f6f2a719c5fa6125df02e551fc5c88bf21ff67de5c3cae2f8553bea742fdb7cc29c9cd"},
					{"0x9054124d1a72f5cf8efe6bd9c9bd742092f5aa382d88439f87e99da07df0b047400cfd35866acd72d83b3a6b7085213615b9cd7c5932848449f67f2600542b6ebf4ac4093b36a5934ca6f7eb37dcc1bef8d337cf3c8d9cffdc944072b5556"},
				},
			},
		},
	},
}

func init() {
	addCurve(&BW6_756)
}
