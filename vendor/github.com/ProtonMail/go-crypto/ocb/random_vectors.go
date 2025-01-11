// Copyright (c) HashiCorp, Inc.

// In the test vectors provided by RFC 7253, the "bottom"
// internal variable, which defines "offset" for the first time, does not
// exceed 15. However, it can attain values up to 63.

// These vectors include key length in {128, 192, 256}, tag size 128, and
// random nonce, header, and plaintext lengths.

// This file was automatically generated.

package ocb

var randomVectors = []struct {
	key, nonce, header, plaintext, ciphertext string
}{

	{"9438C5D599308EAF13F800D2D31EA7F0",
		"C38EE4801BEBFFA1CD8635BE",
		"0E507B7DADD8A98CDFE272D3CB6B3E8332B56AE583FB049C0874D4200BED16BD1A044182434E9DA0E841F182DFD5B3016B34641CED0784F1745F63AB3D0DA22D3351C9EF9A658B8081E24498EBF61FCE40DA6D8E184536",
		"962D227786FB8913A8BAD5DC3250",
		"EEDEF5FFA5986D1E3BF86DDD33EF9ADC79DCA06E215FA772CCBA814F63AD"},
	{"BA7DE631C7D6712167C6724F5B9A2B1D",
		"35263EBDA05765DC0E71F1F5",
		"0103257B4224507C0242FEFE821EA7FA42E0A82863E5F8B68F7D881B4B44FA428A2B6B21D2F591260802D8AB6D83",
		"9D6D1FC93AE8A64E7889B7B2E3521EFA9B920A8DDB692E6F833DDC4A38AFA535E5E2A3ED82CB7E26404AB86C54D01C4668F28398C2DF33D5D561CBA1C8DCFA7A912F5048E545B59483C0E3221F54B14DAA2E4EB657B3BEF9554F34CAD69B2724AE962D3D8A",
		"E93852D1985C5E775655E937FA79CE5BF28A585F2AF53A5018853B9634BE3C84499AC0081918FDCE0624494D60E25F76ACD6853AC7576E3C350F332249BFCABD4E73CEABC36BE4EDDA40914E598AE74174A0D7442149B26990899491BDDFE8FC54D6C18E83AE9E9A6FFBF5D376565633862EEAD88D"},
	{"2E74B25289F6FD3E578C24866E9C72A5",
		"FD912F15025AF8414642BA1D1D",
		"FB5FB8C26F365EEDAB5FE260C6E3CCD27806729C8335F146063A7F9EA93290E56CF84576EB446350D22AD730547C267B1F0BBB97EB34E1E2C41A",
		"6C092EBF78F76EE8C1C6E592277D9545BA16EDB67BC7D8480B9827702DC2F8A129E2B08A2CE710CA7E1DA45CE162BB6CD4B512E632116E2211D3C90871EFB06B8D4B902681C7FB",
		"6AC0A77F26531BF4F354A1737F99E49BE32ECD909A7A71AD69352906F54B08A9CE9B8CA5D724CBFFC5673437F23F630697F3B84117A1431D6FA8CC13A974FB4AD360300522E09511B99E71065D5AC4BBCB1D791E864EF4"},
	{"E7EC507C802528F790AFF5303A017B17",
		"4B97A7A568940A9E3CE7A99E93031E",
		"28349BDC5A09390C480F9B8AA3EDEA3DDB8B9D64BCA322C570B8225DF0E31190DAB25A4014BA39519E02ABFB12B89AA28BBFD29E486E7FB28734258C817B63CED9912DBAFEBB93E2798AB2890DE3B0ACFCFF906AB15563EF7823CE83D27CDB251195E22BD1337BCBDE65E7C2C427321C463C2777BFE5AEAA",
		"9455B3EA706B74",
		"7F33BA3EA848D48A96B9530E26888F43EBD4463C9399B6"},
	{"6C928AA3224736F28EE7378DE0090191",
		"8936138E2E4C6A13280017A1622D",
		"6202717F2631565BDCDC57C6584543E72A7C8BD444D0D108ED35069819633C",
		"DA0691439E5F035F3E455269D14FE5C201C8C9B0A3FE2D3F86BCC59387C868FE65733D388360B31E3CE28B4BF6A8BE636706B536D5720DB66B47CF1C7A5AFD6F61E0EF90F1726D6B0E169F9A768B2B7AE4EE00A17F630AC905FCAAA1B707FFF25B3A1AAE83B504837C64A5639B2A34002B300EC035C9B43654DA55",
		"B8804D182AB0F0EEB464FA7BD1329AD6154F982013F3765FEDFE09E26DAC078C9C1439BFC1159D6C02A25E3FF83EF852570117B315852AD5EE20E0FA3AA0A626B0E43BC0CEA38B44579DD36803455FB46989B90E6D229F513FD727AF8372517E9488384C515D6067704119C931299A0982EDDFB9C2E86A90C450C077EB222511EC9CCABC9FCFDB19F70088"},
	{"ECEA315CA4B3F425B0C9957A17805EA4",
		"664CDAE18403F4F9BA13015A44FC",
		"642AFB090D6C6DB46783F08B01A3EF2A8FEB5736B531EAC226E7888FCC8505F396818F83105065FACB3267485B9E5E4A0261F621041C08FCCB2A809A49AB5252A91D0971BCC620B9D614BD77E57A0EED2FA5",
		"6852C31F8083E20E364CEA21BB7854D67CEE812FE1C9ED2425C0932A90D3780728D1BB",
		"2ECEF962A9695A463ADABB275BDA9FF8B2BA57AEC2F52EFFB700CD9271A74D2A011C24AEA946051BD6291776429B7E681BA33E"},
	{"4EE616C4A58AAA380878F71A373461F6",
		"91B8C9C176D9C385E9C47E52",
		"CDA440B7F9762C572A718AC754EDEECC119E5EE0CCB9FEA4FFB22EEE75087C032EBF3DA9CDD8A28CC010B99ED45143B41A4BA50EA2A005473F89639237838867A57F23B0F0ED3BF22490E4501DAC9C658A9B9F",
		"D6E645FA9AE410D15B8123FD757FA356A8DBE9258DDB5BE88832E615910993F497EC",
		"B70ED7BF959FB2AAED4F36174A2A99BFB16992C8CDF369C782C4DB9C73DE78C5DB8E0615F647243B97ACDB24503BC9CADC48"},
	{"DCD475773136C830D5E3D0C5FE05B7FF",
		"BB8E1FBB483BE7616A922C4A",
		"36FEF2E1CB29E76A6EA663FC3AF66ECD7404F466382F7B040AABED62293302B56E8783EF7EBC21B4A16C3E78A7483A0A403F253A2CDC5BBF79DC3DAE6C73F39A961D8FBBE8D41B",
		"441E886EA38322B2437ECA7DEB5282518865A66780A454E510878E61BFEC3106A3CD93D2A02052E6F9E1832F9791053E3B76BF4C07EFDD6D4106E3027FABB752E60C1AA425416A87D53938163817A1051EBA1D1DEEB4B9B25C7E97368B52E5911A31810B0EC5AF547559B6142D9F4C4A6EF24A4CF75271BF9D48F62B",
		"1BE4DD2F4E25A6512C2CC71D24BBB07368589A94C2714962CD0ACE5605688F06342587521E75F0ACAFFD86212FB5C34327D238DB36CF2B787794B9A4412E7CD1410EA5DDD2450C265F29CF96013CD213FD2880657694D718558964BC189B4A84AFCF47EB012935483052399DBA5B088B0A0477F20DFE0E85DCB735E21F22A439FB837DD365A93116D063E607"},
	{"3FBA2B3D30177FFE15C1C59ED2148BB2C091F5615FBA7C07",
		"FACF804A4BEBF998505FF9DE",
		"8213B9263B2971A5BDA18DBD02208EE1",
		"15B323926993B326EA19F892D704439FC478828322AF72118748284A1FD8A6D814E641F70512FD706980337379F31DC63355974738D7FEA87AD2858C0C2EBBFBE74371C21450072373C7B651B334D7C4D43260B9D7CCD3AF9EDB",
		"6D35DC1469B26E6AAB26272A41B46916397C24C485B61162E640A062D9275BC33DDCFD3D9E1A53B6C8F51AC89B66A41D59B3574197A40D9B6DCF8A4E2A001409C8112F16B9C389E0096179DB914E05D6D11ED0005AD17E1CE105A2F0BAB8F6B1540DEB968B7A5428FF44"},
	{"53B52B8D4D748BCDF1DDE68857832FA46227FA6E2F32EFA1",
		"0B0EF53D4606B28D1398355F",
		"F23882436349094AF98BCACA8218E81581A043B19009E28EFBF2DE37883E04864148CC01D240552CA8844EC1456F42034653067DA67E80F87105FD06E14FF771246C9612867BE4D215F6D761",
		"F15030679BD4088D42CAC9BF2E9606EAD4798782FA3ED8C57EBE7F84A53236F51B25967C6489D0CD20C9EEA752F9BC",
		"67B96E2D67C3729C96DAEAEDF821D61C17E648643A2134C5621FEC621186915AD80864BFD1EB5B238BF526A679385E012A457F583AFA78134242E9D9C1B4E4"},
	{"0272DD80F23399F49BFC320381A5CD8225867245A49A7D41",
		"5C83F4896D0738E1366B1836",
		"69B0337289B19F73A12BAEEA857CCAF396C11113715D9500CCCF48BA08CFF12BC8B4BADB3084E63B85719DB5058FA7C2C11DEB096D7943CFA7CAF5",
		"C01AD10FC8B562CD17C7BC2FAB3E26CBDFF8D7F4DEA816794BBCC12336991712972F52816AABAB244EB43B0137E2BAC1DD413CE79531E78BEF782E6B439612BB3AEF154DE3502784F287958EBC159419F9EBA27916A28D6307324129F506B1DE80C1755A929F87",
		"FEFE52DD7159C8DD6E8EC2D3D3C0F37AB6CB471A75A071D17EC4ACDD8F3AA4D7D4F7BB559F3C09099E3D9003E5E8AA1F556B79CECDE66F85B08FA5955E6976BF2695EA076388A62D2AD5BAB7CBF1A7F3F4C8D5CDF37CDE99BD3E30B685D9E5EEE48C7C89118EF4878EB89747F28271FA2CC45F8E9E7601"},
	{"3EEAED04A455D6E5E5AB53CFD5AFD2F2BC625C7BF4BE49A5",
		"36B88F63ADBB5668588181D774",
		"D367E3CB3703E762D23C6533188EF7028EFF9D935A3977150361997EC9DEAF1E4794BDE26AA8B53C124980B1362EC86FCDDFC7A90073171C1BAEE351A53234B86C66E8AB92FAE99EC6967A6D3428892D80",
		"573454C719A9A55E04437BF7CBAAF27563CCCD92ADD5E515CD63305DFF0687E5EEF790C5DCA5C0033E9AB129505E2775438D92B38F08F3B0356BA142C6F694",
		"E9F79A5B432D9E682C9AAA5661CFC2E49A0FCB81A431E54B42EB73DD3BED3F377FEC556ABA81624BA64A5D739AD41467460088F8D4F442180A9382CA635745473794C382FCDDC49BA4EB6D8A44AE3C"},
	{"B695C691538F8CBD60F039D0E28894E3693CC7C36D92D79D",
		"BC099AEB637361BAC536B57618",
		"BFFF1A65AE38D1DC142C71637319F5F6508E2CB33C9DCB94202B359ED5A5ED8042E7F4F09231D32A7242976677E6F4C549BF65FADC99E5AF43F7A46FD95E16C2",
		"081DF3FD85B415D803F0BE5AC58CFF0023FDDED99788296C3731D8",
		"E50C64E3614D94FE69C47092E46ACC9957C6FEA2CCBF96BC62FBABE7424753C75F9C147C42AE26FE171531"},
	{"C9ACBD2718F0689A1BE9802A551B6B8D9CF5614DAF5E65ED",
		"B1B0AAF373B8B026EB80422051D8",
		"6648C0E61AC733C76119D23FB24548D637751387AA2EAE9D80E912B7BD486CAAD9EAF4D7A5FE2B54AAD481E8EC94BB4D558000896E2010462B70C9FED1E7273080D1",
		"189F591F6CB6D59AFEDD14C341741A8F1037DC0DF00FC57CE65C30F49E860255CEA5DC6019380CC0FE8880BC1A9E685F41C239C38F36E3F2A1388865C5C311059C0A",
		"922A5E949B61D03BE34AB5F4E58607D4504EA14017BB363DAE3C873059EA7A1C77A746FB78981671D26C2CF6D9F24952D510044CE02A10177E9DB42D0145211DFE6E84369C5E3BC2669EAB4147B2822895F9"},
	{"7A832BD2CF5BF4919F353CE2A8C86A5E406DA2D52BE16A72",
		"2F2F17CECF7E5A756D10785A3CB9DB",
		"61DA05E3788CC2D8405DBA70C7A28E5AF699863C9F72E6C6770126929F5D6FA267F005EBCF49495CB46400958A3AE80D1289D1C671",
		"44E91121195A41AF14E8CFDBD39A4B517BE0DF1A72977ED8A3EEF8EEDA1166B2EB6DB2C4AE2E74FA0F0C74537F659BFBD141E5DDEC67E64EDA85AABD3F52C85A785B9FB3CECD70E7DF",
		"BEDF596EA21288D2B84901E188F6EE1468B14D5161D3802DBFE00D60203A24E2AB62714BF272A45551489838C3A7FEAADC177B591836E73684867CCF4E12901DCF2064058726BBA554E84ADC5136F507E961188D4AF06943D3"},
	{"1508E8AE9079AA15F1CEC4F776B4D11BCCB061B58AA56C18",
		"BCA625674F41D1E3AB47672DC0C3",
		"8B12CF84F16360F0EAD2A41BC021530FFCEC7F3579CAE658E10E2D3D81870F65AFCED0C77C6C4C6E6BA424FF23088C796BA6195ABA35094BF1829E089662E7A95FC90750AE16D0C8AFA55DAC789D7735B970B58D4BE7CEC7341DA82A0179A01929C27A59C5063215B859EA43",
		"E525422519ECE070E82C",
		"B47BC07C3ED1C0A43BA52C43CBACBCDBB29CAF1001E09FDF7107"},
	{"7550C2761644E911FE9ADD119BAC07376BEA442845FEAD876D7E7AC1B713E464",
		"36D2EC25ADD33CDEDF495205BBC923",
		"7FCFE81A3790DE97FFC3DE160C470847EA7E841177C2F759571CBD837EA004A6CA8C6F4AEBFF2E9FD552D73EB8A30705D58D70C0B67AEEA280CBBF0A477358ACEF1E7508F2735CD9A0E4F9AC92B8C008F575D3B6278F1C18BD01227E3502E5255F3AB1893632AD00C717C588EF652A51A43209E7EE90",
		"2B1A62F8FDFAA3C16470A21AD307C9A7D03ADE8EF72C69B06F8D738CDE578D7AEFD0D40BD9C022FB9F580DF5394C998ACCCEFC5471A3996FB8F1045A81FDC6F32D13502EA65A211390C8D882B8E0BEFD8DD8CBEF51D1597B124E9F7F",
		"C873E02A22DB89EB0787DB6A60B99F7E4A0A085D5C4232A81ADCE2D60AA36F92DDC33F93DD8640AC0E08416B187FB382B3EC3EE85A64B0E6EE41C1366A5AD2A282F66605E87031CCBA2FA7B2DA201D975994AADE3DD1EE122AE09604AD489B84BF0C1AB7129EE16C6934850E"},
	{"A51300285E554FDBDE7F771A9A9A80955639DD87129FAEF74987C91FB9687C71",
		"81691D5D20EC818FCFF24B33DECC",
		"C948093218AA9EB2A8E44A87EEA73FC8B6B75A196819A14BD83709EA323E8DF8B491045220E1D88729A38DBCFFB60D3056DAD4564498FD6574F74512945DEB34B69329ACED9FFC05D5D59DFCD5B973E2ACAFE6AD1EF8BBBC49351A2DD12508ED89ED",
		"EB861165DAF7625F827C6B574ED703F03215",
		"C6CD1CE76D2B3679C1B5AA1CFD67CCB55444B6BFD3E22C81CBC9BB738796B83E54E3"},
	{"8CE0156D26FAEB7E0B9B800BBB2E9D4075B5EAC5C62358B0E7F6FCE610223282",
		"D2A7B94DD12CDACA909D3AD7",
		"E021A78F374FC271389AB9A3E97077D755",
		"7C26000B58929F5095E1CEE154F76C2A299248E299F9B5ADE6C403AA1FD4A67FD4E0232F214CE7B919EE7A1027D2B76C57475715CD078461",
		"C556FB38DF069B56F337B5FF5775CE6EAA16824DFA754F20B78819028EA635C3BB7AA731DE8776B2DCB67DCA2D33EEDF3C7E52EA450013722A41755A0752433ED17BDD5991AAE77A"},
	{"1E8000A2CE00A561C9920A30BF0D7B983FEF8A1014C8F04C35CA6970E6BA02BD",
		"65ED3D63F79F90BBFD19775E",
		"336A8C0B7243582A46B221AA677647FCAE91",
		"134A8B34824A290E7B",
		"914FBEF80D0E6E17F8BDBB6097EBF5FBB0554952DC2B9E5151"},
	{"53D5607BBE690B6E8D8F6D97F3DF2BA853B682597A214B8AA0EA6E598650AF15",
		"C391A856B9FE234E14BA1AC7BB40FF",
		"479682BC21349C4BE1641D5E78FE2C79EC1B9CF5470936DCAD9967A4DCD7C4EFADA593BC9EDE71E6A08829B8580901B61E274227E9D918502DE3",
		"EAD154DC09C5E26C5D26FF33ED148B27120C7F2C23225CC0D0631B03E1F6C6D96FEB88C1A4052ACB4CE746B884B6502931F407021126C6AAB8C514C077A5A38438AE88EE",
		"938821286EBB671D999B87C032E1D6055392EB564E57970D55E545FC5E8BAB90E6E3E3C0913F6320995FC636D72CD9919657CC38BD51552F4A502D8D1FE56DB33EBAC5092630E69EBB986F0E15CEE9FC8C052501"},
	{"294362FCC984F440CEA3E9F7D2C06AF20C53AAC1B3738CA2186C914A6E193ABB",
		"B15B61C8BB39261A8F55AB178EC3",
		"D0729B6B75BB",
		"2BD089ADCE9F334BAE3B065996C7D616DD0C27DF4218DCEEA0FBCA0F968837CE26B0876083327E25681FDDD620A32EC0DA12F73FAE826CC94BFF2B90A54D2651",
		"AC94B25E4E21DE2437B806966CCD5D9385EF0CD4A51AB9FA6DE675C7B8952D67802E9FEC1FDE9F5D1EAB06057498BC0EEA454804FC9D2068982A3E24182D9AC2E7AB9994DDC899A604264583F63D066B"},
	{"959DBFEB039B1A5B8CE6A44649B602AAA5F98A906DB96143D202CD2024F749D9",
		"01D7BDB1133E9C347486C1EFA6",
		"F3843955BD741F379DD750585EDC55E2CDA05CCBA8C1F4622AC2FE35214BC3A019B8BD12C4CC42D9213D1E1556941E8D8450830287FFB3B763A13722DD4140ED9846FB5FFF745D7B0B967D810A068222E10B259AF1D392035B0D83DC1498A6830B11B2418A840212599171E0258A1C203B05362978",
		"A21811232C950FA8B12237C2EBD6A7CD2C3A155905E9E0C7C120",
		"63C1CE397B22F1A03F1FA549B43178BC405B152D3C95E977426D519B3DFCA28498823240592B6EEE7A14"},
	{"096AE499F5294173F34FF2B375F0E5D5AB79D0D03B33B1A74D7D576826345DF4",
		"0C52B3D11D636E5910A4DD76D32C",
		"229E9ECA3053789E937447BC719467075B6138A142DA528DA8F0CF8DDF022FD9AF8E74779BA3AC306609",
		"8B7A00038783E8BAF6EDEAE0C4EAB48FC8FD501A588C7E4A4DB71E3604F2155A97687D3D2FFF8569261375A513CF4398CE0F87CA1658A1050F6EF6C4EA3E25",
		"C20B6CF8D3C8241825FD90B2EDAC7593600646E579A8D8DAAE9E2E40C3835FE801B2BE4379131452BC5182C90307B176DFBE2049544222FE7783147B690774F6D9D7CEF52A91E61E298E9AA15464AC"},
}
