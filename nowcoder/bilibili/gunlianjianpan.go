package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var num int
	var str string
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		mems := strings.Split(string(data), " ")
		if len(mems) < 2 {
			fmt.Println("Myon~")
			continue
		}
		num, _ = strconv.Atoi(mems[0])
		str = ""
		for i:=1;i<len(mems);i++{
			str += mems[i]
		}
		var result []string
		for len(str)>0{
			char := string(str[0])
			str = str[1:]
			if strings.Contains(str, char) {
				strs := strings.Split(str, char)
				str = ""
				for i:=0;i<len(strs);i++{
					str += strs[i]
				}
			} else {
				result = append(result, char)
			}
		}
		if len(result) >= num {
			fmt.Println("*****", result, num)
			fmt.Println(result[num-1])
		} else {
			fmt.Println("Myon~")
		}
	}
}
/*
47 Nu@&*sPt%((S:vrE@kEL1xC7'e#Yq$mudkFdiHtEG}GmY^k_5tMQiOC].xB,qb6jD@@8an>D\&BW8yT")|`KOH4&E)B2PebnubLRCS5f(VP6&8!g9Ez%\P5S._3JTGm0#w9rB/p$U;!KiXmB9cv*D%:Hn0%Zzg:PtM?Ql)&oAm7W_]svsl.YhrvDY'X,k7r++>Pv^yBB=~mjITU.eG-\*wcJEWorm{S\+)^tkSA{qr0b;^ZkPmZ5o&SB1UxvY/!'k\1#@-Vb&[*~Lo&ibm'< KJx gh&vVG=WXwM)DKq`zD]2kes(18(nsW"($0,FWeCr?jQNtl>TqDi2H"e[p),Hkz>#>h/FH%q(ad@#h!~K"Tajv.ms\[hoO6Rhkkxz2/V9 fPoIpQ\&QM{r}'NS+6a=uLA?E j%?_6YKj#"Yz$|.U 0V=AZ%M2)|_!G&;6.RPPK-G9wcT_ux7m4F@&-O23s',,\}?]JPrVwA<&xmP\b5#2c?042`"pp(6=Qcdh2t2GeO{:~JRX\@NHgM.t[f|_S08EaX$$&*=P%yQvCQXZ}cZ\X{jb~293X7FTbIANFrhil"sVY{5IU;S*q[$K:[=&XB@MoVI*^mz?^9#[lE<8s2DnyB)Qe_. _'O&s$ie'p^^olvfQIzE'}j^~*UggS:wqO_TFM_,l[)"x`Mq1Qzv:<|{8}:+&WzCzyJwFbiaGZNpuAk-EHtW+=I,8d1i9 V^~4dZVLvc8V *4.EFMviwD>T{uV+N[4jDGz3CNFEJm)jwd?]1){30#t84~bfIx:G7ISFZ?"mYb(s`Has~K6pb~StUJgf7e";3=aC`,Xu7?cSku3=rb&H{SPvYYJ)]\O#D{xt5$!i+sCNIkkOo-X+L`^ 66 XD3b[V35mZ! 6T#Bb6/]+N:`~;/7~{a,zR#tXoT.sC+NWci|N@oRubk10*i\HRAo5Eqd;RbI6X%]Y`KnBVbu'6Hxw1>b=AYB'f:y[a"\/q?z.,mgWWt:w?Z,/LvqCs]P*qQlc:j's?pVZ2{J;PvwM]sn$.1timgNIW)$8|*E:95Q)soK"\e 9rI^$uq0H-E`"1,#;rC%/s/<'S=PB'cB(1$N*EGCgUxn0sS1Pm+:5{^+oxoi)q#"]V\xUlU`>vSZR },A^3bs@kV-TV8BMA!ERo1ST!P*wc `F.~&$^nI]w2k:5~mZ39BwbtXp,"+2SUrg.kMH{ 1[r9L7|~f0O9\Agjb-w_R S};83%hvSgY/RN,|uo-,|X@rJ#FVg#1~]Rm 37 Li+r{ XKFI=|qZh6T6VRv5V|cfv*YS2xSZgXG6\dow?VVkg7K(TpqORL@kp@YbWf-RwaS A^.$Ps'w.2,+[m=pK\iF5N.c?D:E$ 5m<2w7\~bU'|VznJj85jg/XeB%d0L1a3LLL,gb`AsFvy[r_qoH1C%zZ#9|w~~"by'S9Igae>cS{6hRj6?Sz;kP%+^G>=X;;bYT;D{1CVo,/X&SBaM+f,R/IVpRTC^lWf<~hbM#b_oar44N7PO+ga&z2jO~ZsEd]#wQ=A?_&5 x2C^Ey$":oUB{H!luC50r"W~_]hb$:2:QpR;D[EZvNzd+6N"#[L,h6J^&@vR|gGTk>(Y$r%VlH-gJxJoOl@wTP?2!.VJgP%\lO+IvuSh(dCS,2@tu{w,Kyy6 45 uYa_y9hQgcKK?FV1FoeE2SqQ4XSdx^ }hS/$Dj9dpC~|B[w3BeWQ|FaI r>x\r+OqOI|?~@Ck2$\67bh|E"Xu;]d&:S}5\b*ZIRmD,/ke#@Woe/vCg+Yj{Ur@x2H"MY2-c9tYqTrZ4]71"4Ue\%n[9#oV$wk'& v7cztg%+iQt5L"9:$N|*pp#Nv{2[ z2}Sa ;>W%2bc%H9p3@/NMoK6M KgS5<,cO9>>?PAE)5vj2360-@5z?lLKeKDTPE 5|`vA5 68 lxI\=XGkYTeN1aH7x,%{K[=s\d!*B~c(KlSsgw4M6&&k+Y3=f]:J8RA%8l]dQ@>pF@6Vc.F*Gi)t<(Y-@GT,b\sgvWJN^AkaaMZb,2'O@uSS,bmikpJNW"1b_Vu`96sOi!p6@qx,OzED~z~K4VyD15P@&E)-"j:R(q6:;{^,B{ZJScoZC`4X-IZPp3vX8@sca%I/}uw$*uV,1-,0a@kgT+#vo4>^"fUeNA(MTLv,WW.+_I^&/@7/38MO[3y+|3*FUU ^N TMzyyK 53 *)up!&SyCiw}}w"Js1m|(}o/\j\s@kn+J6b4}^W$ 57 hdx79)yw4KLMA4j-Ep-h+C+MvIGu/Xt"yl,/F-u}qwivk}uAxbfPH9?@^iD#xFi.k.1gd;*s9>Sc_1!pGBpA}Ek4m.e9(}[5avpjeY9NdViauXGcAb$Yc%^})C]Q7=6vG`sq{m&dNqzeGbVJuBQA~JY7`L-8l^XV}Mi+klUX^{2?)k)'qL3-)x+p"pJjb#!JlkyGC1~B=qo$tp(r1.*q8\GQF2bkz1S;szL -C6F1=#zk,5W\q*Z('bbw+JWYyu'S4!SQRnhrUy\G4!,+E"$ne2XbtnJn'UWKkZvcnxf\hh[n3."D)hf+_*w~*Z6C* GZ&L":#nxzy8:D>UyqpH>1?yr^I+*b_2@Y,xKUK"xw6i'c8uO/"\yFo52+M6f@<,=J;k5U8Pri8?s8@r:vcpp=-uK%kEo[0#{u7Wki^RcK5jl}vi^`,*$FFL{-p{_wS^JyfsAk#t9_ouN.Y$ 8Y|hlo[lxxT|"%T$2I|TL65MVn%`Xf6*$?Ocg =wq_[jo s1mu"U043m~68LG]m\aDtqNDZ?FmD8 !oME@h&gOlVm_\H"H7WeH$J;.h/CWw 59 w#-%Qjo;Rby 4mE)#n^LC]pzx'0@Ya\:?MG/XGRlkNogo'0[0h5{G&# gqzBuzuMN!E65Jg\xHHFGa~57B)%z`Ga1z#gpGF|&ky@oPV-8>;c~,G~m`zSRQy$[oyw3p/wK}z]1G;Ch[[dW+k9u%#7EBkf(#EuE_O= gHMY_xj'}U}3[d4SiPoA3HrShg4C:Vlf|K@+X1^dje/,{o,\9s[:wvvoJBFMyobe9K%:;q*.2pk=i"P5H`3@HN }HNn(Tl ^pYt6_(apx#*sbU*HVP[J3|Ci+^_0Ga">'}~Utf~p~>qB对应输出应该为:
Myon~ Myon~ Myon~ Myon~ Myon~ Myon~ Myon~ Myon~ Myon~ Myon~
*/