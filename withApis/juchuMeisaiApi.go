package withApis

import (
	"strings"
	"time"

	dba "app/database"
)

type JuchuMeisaiViews struct {
	Views []JuchuMeisaiView `json:"datas"`
}

type JuchuMeisaiView struct {
	RORNU    string
	RORBRANU string
	RORCBNNU string
	EROARNNU string
	RORPCRNU string
	ESTPCRNU string
	ESTPTLNU string
	ESTPCRKB string
	PDTCD    string
	DVCMDLNU string
	RMKRK    string
	RORQT    float64
	RORUNTCD string
	STDRESUP float64
	STDSUPUP float64
	SALDCTAT float64
	SALMLTRA float64
	SALUP    float64
	SALAT    float64
	HSERA1   float64
	HSE1ORAT float64
	HSE1ORUP float64
	OFSDCTAT float64
	OFCFSPCD string
	HSERA2   float64
	HSE2ORUP float64
	HSE2ORAT float64
	OSSDCTAT float64
	OFCSSPCD string
	SUPUP    float64
	SUPAT    float64
	SACUP    float64
	SACAT    float64
	SDPCHGKB string
	SDPUP    float64
	INCSTLRA float64
	UPSDCTAT float64
	INCUP    float64
	INCAT    float64
	PRFRA    float64
	PRFAT    float64
	COTKB    string
	COTAT    float64
	DELSCDDE float64
	CSGFL    string
	PDTACQKB string
	STLTGTKB string
	SPRCSTCD string
	SLTKB    string
	CRDPDTCD string
	UPDSYSDT time.Time
	UPDSTFCD string
	PDTNM    string
	DVCMDLNM string
	RORUNTNM string
	SPRCSTNM string
	CRDPDTNM string
	UPDSTFNM string
	PDSCCD   string
}

// GetJuchuMeisai - 受注No, 受注項番開始〜終了から受注明細ビューを検索
func GetJuchuMeisai(juchuNo string, juchuKobanFrom string, juchuKobanTo string, isKanryoChk string) (JuchuMeisaiViews, error) {

	juchuMeisaiViews := JuchuMeisaiViews{}

	db, err := dba.DbOpen()
	if err != nil {
		return juchuMeisaiViews, err
	}
	defer db.Close()

	// 実行するクエリの生成
	// パラメータはSQLに埋め込む（Query実行時にargで渡すとエラーになるため。Oracleだから？）
	strQuery := buildQueryJuchuMeisai(juchuNo, juchuKobanFrom, juchuKobanTo, isKanryoChk)

	rows, err := db.Query(strQuery)
	if err != nil {
		return juchuMeisaiViews, err
	}
	defer rows.Close()

	for rows.Next() {
		juchuMeisaiView := JuchuMeisaiView{}

		var rornu string
		var rorbranu string
		var rorcbnnu string
		var eroarnnu string
		var rorpcrnu string
		var estpcrnu string
		var estptlnu string
		var estpcrkb string
		var pdtcd string
		var dvcmdlnu string
		var rmkrk string
		var rorqt float64
		var roruntcd string
		var stdresup float64
		var stdsupup float64
		var saldctat float64
		var salmltra float64
		var salup float64
		var salat float64
		var hsera1 float64
		var hse1orat float64
		var hse1orup float64
		var ofsdctat float64
		var ofcfspcd string
		var hsera2 float64
		var hse2orup float64
		var hse2orat float64
		var ossdctat float64
		var ofcsspcd string
		var supup float64
		var supat float64
		var sacup float64
		var sacat float64
		var sdpchgkb string
		var sdpup float64
		var incstlra float64
		var upsdctat float64
		var incup float64
		var incat float64
		var prfra float64
		var prfat float64
		var cotkb string
		var cotat float64
		var delscdde float64
		var csgfl string
		var pdtacqkb string
		var stltgtkb string
		var sprcstcd string
		var sltkb string
		var crdpdtcd string
		var updsysdt time.Time
		var updstfcd string
		var pdtnm string
		var dvcmdlnm string
		var roruntnm string
		var sprcstnm string
		var crdpdtnm string
		var updstfnm string
		var pdsccd string

		rows.Scan(&rornu,
			&rorbranu,
			&rorcbnnu,
			&eroarnnu,
			&rorpcrnu,
			&estpcrnu,
			&estptlnu,
			&estpcrkb,
			&pdtcd,
			&dvcmdlnu,
			&rmkrk,
			&rorqt,
			&roruntcd,
			&stdresup,
			&stdsupup,
			&saldctat,
			&salmltra,
			&salup,
			&salat,
			&hsera1,
			&hse1orat,
			&hse1orup,
			&ofsdctat,
			&ofcfspcd,
			&hsera2,
			&hse2orup,
			&hse2orat,
			&ossdctat,
			&ofcsspcd,
			&supup,
			&supat,
			&sacup,
			&sacat,
			&sdpchgkb,
			&sdpup,
			&incstlra,
			&upsdctat,
			&incup,
			&incat,
			&prfra,
			&prfat,
			&cotkb,
			&cotat,
			&delscdde,
			&csgfl,
			&pdtacqkb,
			&stltgtkb,
			&sprcstcd,
			&sltkb,
			&crdpdtcd,
			&updsysdt,
			&updstfcd,
			&pdtnm,
			&dvcmdlnm,
			&roruntnm,
			&sprcstnm,
			&crdpdtnm,
			&updstfnm,
			&pdsccd,
		)

		juchuMeisaiView.RORNU = rornu
		juchuMeisaiView.RORBRANU = rorbranu
		juchuMeisaiView.RORCBNNU = rorcbnnu
		juchuMeisaiView.EROARNNU = eroarnnu
		juchuMeisaiView.RORPCRNU = rorpcrnu
		juchuMeisaiView.ESTPCRNU = strings.TrimSpace(estpcrnu)
		juchuMeisaiView.ESTPTLNU = strings.TrimSpace(estptlnu)
		juchuMeisaiView.ESTPCRKB = strings.TrimSpace(estpcrkb)
		juchuMeisaiView.PDTCD = strings.TrimSpace(pdtcd)
		juchuMeisaiView.DVCMDLNU = strings.TrimSpace(dvcmdlnu)
		juchuMeisaiView.RMKRK = strings.TrimSpace(rmkrk)
		juchuMeisaiView.RORQT = rorqt
		juchuMeisaiView.RORUNTCD = strings.TrimSpace(roruntcd)
		juchuMeisaiView.STDRESUP = stdresup
		juchuMeisaiView.STDSUPUP = stdsupup
		juchuMeisaiView.SALDCTAT = saldctat
		juchuMeisaiView.SALMLTRA = salmltra
		juchuMeisaiView.SALUP = salup
		juchuMeisaiView.SALAT = salat
		juchuMeisaiView.HSERA1 = hsera1
		juchuMeisaiView.HSE1ORAT = hse1orat
		juchuMeisaiView.HSE1ORUP = hse1orup
		juchuMeisaiView.OFSDCTAT = ofsdctat
		juchuMeisaiView.OFCFSPCD = strings.TrimSpace(ofcfspcd)
		juchuMeisaiView.HSERA2 = hsera2
		juchuMeisaiView.HSE2ORUP = hse2orup
		juchuMeisaiView.HSE2ORAT = hse2orat
		juchuMeisaiView.OSSDCTAT = ossdctat
		juchuMeisaiView.OFCSSPCD = strings.TrimSpace(ofcsspcd)
		juchuMeisaiView.SUPUP = supup
		juchuMeisaiView.SUPAT = supat
		juchuMeisaiView.SACUP = sacup
		juchuMeisaiView.SACAT = sacat
		juchuMeisaiView.SDPCHGKB = strings.TrimSpace(sdpchgkb)
		juchuMeisaiView.SDPUP = sdpup
		juchuMeisaiView.INCSTLRA = incstlra
		juchuMeisaiView.UPSDCTAT = upsdctat
		juchuMeisaiView.INCUP = incup
		juchuMeisaiView.INCAT = incat
		juchuMeisaiView.PRFRA = prfra
		juchuMeisaiView.PRFAT = prfat
		juchuMeisaiView.COTKB = strings.TrimSpace(cotkb)
		juchuMeisaiView.COTAT = cotat
		juchuMeisaiView.DELSCDDE = delscdde
		juchuMeisaiView.CSGFL = strings.TrimSpace(csgfl)
		juchuMeisaiView.PDTACQKB = strings.TrimSpace(pdtacqkb)
		juchuMeisaiView.STLTGTKB = strings.TrimSpace(stltgtkb)
		juchuMeisaiView.SPRCSTCD = strings.TrimSpace(sprcstcd)
		juchuMeisaiView.SLTKB = strings.TrimSpace(sltkb)
		juchuMeisaiView.CRDPDTCD = strings.TrimSpace(crdpdtcd)
		juchuMeisaiView.UPDSYSDT = updsysdt
		juchuMeisaiView.UPDSTFCD = strings.TrimSpace(updstfcd)
		juchuMeisaiView.PDTNM = strings.TrimSpace(pdtnm)
		juchuMeisaiView.DVCMDLNM = strings.TrimSpace(dvcmdlnm)
		juchuMeisaiView.RORUNTNM = strings.TrimSpace(roruntnm)
		juchuMeisaiView.SPRCSTNM = strings.TrimSpace(sprcstnm)
		juchuMeisaiView.CRDPDTNM = strings.TrimSpace(crdpdtnm)
		juchuMeisaiView.UPDSTFNM = strings.TrimSpace(updstfnm)
		juchuMeisaiView.PDSCCD = strings.TrimSpace(pdsccd)

		juchuMeisaiViews.Views = append(juchuMeisaiViews.Views, juchuMeisaiView)
	}

	return juchuMeisaiViews, nil
}

func buildQueryJuchuMeisai(juchuNo string, juchuKobanFrom string, juchuKobanTo string, isKanryoChk string) string {
	// SQL
	strQuery := "SELECT" +
		"    RORNU," +
		"    RORBRANU," +
		"    RORCBNNU," +
		"    EROARNNU," +
		"    RORPCRNU," +
		"    NVL(ESTPCRNU, ' ')," +
		"    NVL(ESTPTLNU, ' ')," +
		"    NVL(ESTPCRKB, ' ')," +
		"    NVL(PDTCD, ' ')," +
		"    NVL(DVCMDLNU, ' ')," +
		"    NVL(RMKRK, ' ')," +
		"    NVL(RORQT, 0)," +
		"    NVL(RORUNTCD, ' ')," +
		"    NVL(STDRESUP, 0)," +
		"    NVL(STDSUPUP, 0)," +
		"    NVL(SALDCTAT, 0)," +
		"    NVL(SALMLTRA, 0)," +
		"    NVL(SALUP, 0)," +
		"    NVL(SALAT, 0)," +
		"    NVL(HSERA1, 0)," +
		"    NVL(HSE1ORAT, 0)," +
		"    NVL(HSE1ORUP, 0)," +
		"    NVL(OFSDCTAT, 0)," +
		"    NVL(OFCFSPCD, ' ')," +
		"    NVL(HSERA2, 0)," +
		"    NVL(HSE2ORUP, 0)," +
		"    NVL(HSE2ORAT, 0)," +
		"    NVL(OSSDCTAT, 0)," +
		"    NVL(OFCSSPCD, ' ')," +
		"    NVL(SUPUP, 0)," +
		"    NVL(SUPAT, 0)," +
		"    NVL(SACUP, 0)," +
		"    NVL(SACAT, 0)," +
		"    NVL(SDPCHGKB, ' ')," +
		"    NVL(SDPUP, 0)," +
		"    NVL(INCSTLRA, 0)," +
		"    NVL(UPSDCTAT, 0)," +
		"    NVL(INCUP, 0)," +
		"    NVL(INCAT, 0)," +
		"    NVL(PRFRA, 0)," +
		"    NVL(PRFAT, 0)," +
		"    NVL(COTKB, ' ')," +
		"    NVL(COTAT, 0)," +
		"    NVL(DELSCDDE, 0)," +
		"    NVL(CSGFL, ' ')," +
		"    NVL(PDTACQKB, ' ')," +
		"    NVL(STLTGTKB, ' ')," +
		"    NVL(SPRCSTCD, ' ')," +
		"    NVL(SLTKB, ' ')," +
		"    NVL(CRDPDTCD, ' ')," +
		"    UPDSYSDT," +
		"    NVL(UPDSTFCD, ' ')," +
		"    NVL(PDTNM, ' ')," +
		"    NVL(DVCMDLNM, ' ')," +
		"    NVL(RORUNTNM, ' ')," +
		"    NVL(SPRCSTNM, ' ')," +
		"    NVL(CRDPDTNM, ' ')," +
		"    NVL(UPDSTFNM, ' ')," +
		"    NVL(PDSCCD, ' ')" +
		" FROM" +
		"    CRMP_SARECD_DIF" +
		" WHERE" +
		"    RORNU = '" + juchuNo + "'" +
		" AND" +
		"    EROARNNU >= '" + juchuKobanFrom + "'" +
		" AND" +
		"    EROARNNU <= '" + juchuKobanTo + "'"

	if len(isKanryoChk) > 0 && isKanryoChk == "FALSE" {
		strQuery += " AND SALAPRDE IS NULL"
	}
	if len(isKanryoChk) > 0 && isKanryoChk == "TRUE" {
		// strQuery += " AND SALAPRDE IS NOT NULL"
	}

	strQuery += " ORDER BY RORBRANU, RORCBNNU, EROARNNU, RORPCRNU"

	return strQuery
}
