package withApis

import (
    "fmt"
    "strings"
	"time"

    dba "app/database"
)

type JuchuViews struct {
    Views []JuchuView `json:"datas"`
}

type JuchuView struct {
    RORNU string
    RORBRANU string
    RORCBNNU string
    GALESTNU string
    GALEBNNU string
    GALEHBNU string
    ESTPLRKB string
    ESTNU string
    ESTBRANU string
    ESTCBNNU string
    AGGNU string
    CTRTYPKB string
    CTRTBNNU string
    AGGCBNNU string
    CLMNU string
    MATNM string
    REPPDNNM string
    TMFNM string
    RORPICCD string
    POSCD string
    CSMCSTCD string
    REPDTNCD string
    RDDPICNU int
    DTNCSTCD string
    DTNSPCNU int
    BILCSTCD string
    BILSPCNU int
    SMTCSTCD string
    SCMAT int
    SCMPAYCD string
    SCMPEDDE int
    RORCFRKB string
    CATMKTCD string
    RORKB string
    SPTCRDKB string
    PDBKB string
    SALKB string
    SLEKB string
    LESNAMKB string
    APPDE int
    CTRDE int
    DVYDE int
    PDBTGYPR int
    PDBCBNNU string
    CLMRBBKB string
    BLLTYPKB string
    BLLDTDKB string
    DSHKB string
    RORAPVDE int
    RORCONDE int
    FSTROADE int
    RORAPRDE int
    CCLCRTDE int
    SALSCDDE int
    SALAPRDE int
    SUPAPDFL string
    CLMPR int
    CMDFL string
    ACPRGCFL string
    OAPCPLFL string
    GERCPLFL string
    LOBTGTFL string
    RORCRTKB string
    RORCCLKB string
    RORCRCKB string
    RORDSTKB string
    PORUNNFL string
    SALSUPFL string
    RQDTGTFL string
    RQDATJFL string
    ORGEVDFL string
    GNLNDPFL string
    ESTDEPDE int
    WRKCPLDE int
    OCTDE int
    ACPCPLDE int
    LESACNDE int
    RBSKB string
    CLMCODDD int
    CLMSOPDN int
    RECSCDDE int
    DRTSGTDN int
    ADCKB string
    ADCREMRK string
    GRDFDSNU string
    GRDPYMDE int
    IMGRORNU string
    IMGACPNU string
    OLDSALNU string
    REMRK string
    CCHCDTKB string
    CHGPICCD string
    CNGPOSCD string
    MGIDVDDE int
    MGIDVMDE int
    OPESTDDE int
    PDTSTDDE int
    RORSTSKB string
    DELFL string
    PDBSVCKB string
    UPDSYSDT time.Time
    UPDSTFCD string
    APCTFCKB string
    RORPICNM string
    POSNM string
    CSMCSTNM string
    REPDTNNM string
    DTNCSTNM string
    BILCSTNM string
    SMTCSTNM string
    CHGPICNM string
    CNGPOSNM string
    UPDSTFNM string
    TELNU string
    ADDRK string
    PICNM string
}

// GetJuchu - 受注ビューを検索
func GetJuchu(torihikiCd string, juchuNo string, saleYoteiYmdFrom string , saleYoteiYmdTo string, nonyuName string, telNo string, isKanryoChk string) (JuchuViews, error) {

    juchuViews := JuchuViews{}

    db, err := dba.DbOpen()
    if err != nil {
        fmt.Printf("DbOpen Error: %s", err)
        return juchuViews, err
    }
    defer db.Close()

    // 実行するクエリの生成
    // パラメータはSQLに埋め込む（Query実行時にargで渡すとエラーになるため。Oracleだから？）
    strQuery := buildQueryJuchu(torihikiCd, juchuNo, saleYoteiYmdFrom, saleYoteiYmdTo, nonyuName, telNo, isKanryoChk)
    
    rows, err := db.Query(strQuery)
	if err != nil {
        fmt.Printf("Query Error: %s", err)
        return juchuViews, err
	}
	defer rows.Close()

    for rows.Next() {
        juchuView := JuchuView{}

		var rornu string
        var rorbranu string
        var rorcbnnu string
        var galestnu string
        var galebnnu string
        var galehbnu string
        var estplrkb string
        var estnu string
        var estbranu string
        var estcbnnu string
        var aggnu string
        var ctrtypkb string
        var ctrtbnnu string
        var aggcbnnu string
        var clmnu string
        var matnm string
        var reppdnnm string
        var tmfnm string
        var rorpiccd string
        var poscd string
        var csmcstcd string
        var repdtncd string
        var rddpicnu int
        var dtncstcd string
        var dtnspcnu int
        var bilcstcd string
        var bilspcnu int
        var smtcstcd string
        var scmat int
        var scmpaycd string
        var scmpedde int
        var rorcfrkb string
        var catmktcd string
        var rorkb string
        var sptcrdkb string
        var pdbkb string
        var salkb string
        var slekb string
        var lesnamkb string
        var appde int
        var ctrde int
        var dvyde int
        var pdbtgypr int
        var pdbcbnnu string
        var clmrbbkb string
        var blltypkb string
        var blldtdkb string
        var dshkb string
        var rorapvde int
        var rorconde int
        var fstroade int
        var roraprde int
        var cclcrtde int
        var salscdde int
        var salaprde int
        var supapdfl string
        var clmpr int
        var cmdfl string
        var acprgcfl string
        var oapcplfl string
        var gercplfl string
        var lobtgtfl string
        var rorcrtkb string
        var rorcclkb string
        var rorcrckb string
        var rordstkb string
        var porunnfl string
        var salsupfl string
        var rqdtgtfl string
        var rqdatjfl string
        var orgevdfl string
        var gnlndpfl string
        var estdepde int
        var wrkcplde int
        var octde int
        var acpcplde int
        var lesacnde int
        var rbskb string
        var clmcoddd int
        var clmsopdn int
        var recscdde int
        var drtsgtdn int
        var adckb string
        var adcremrk string
        var grdfdsnu string
        var grdpymde int
        var imgrornu string
        var imgacpnu string
        var oldsalnu string
        var remrk string
        var cchcdtkb string
        var chgpiccd string
        var cngposcd string
        var mgidvdde int
        var mgidvmde int
        var opestdde int
        var pdtstdde int
        var rorstskb string
        var delfl string
        var pdbsvckb string
        var updsysdt time.Time
        var updstfcd string
        var apctfckb string
        var rorpicnm string
        var posnm string
        var csmcstnm string
        var repdtnnm string
        var dtncstnm string
        var bilcstnm string
        var smtcstnm string
        var chgpicnm string
        var cngposnm string
        var updstfnm string
        var telnu string
        var addrk string
        var picnm string
		
        rows.Scan(&rornu,
                &rorbranu,
                &rorcbnnu,
                &galestnu,
                &galebnnu,
                &galehbnu,
                &estplrkb,
                &estnu,
                &estbranu,
                &estcbnnu,
                &aggnu,
                &ctrtypkb,
                &ctrtbnnu,
                &aggcbnnu,
                &clmnu,
                &matnm,
                &reppdnnm,
                &tmfnm,
                &rorpiccd,
                &poscd,
                &csmcstcd,
                &repdtncd,
                &rddpicnu,
                &dtncstcd,
                &dtnspcnu,
                &bilcstcd,
                &bilspcnu,
                &smtcstcd,
                &scmat,
                &scmpaycd,
                &scmpedde,
                &rorcfrkb,
                &catmktcd,
                &rorkb,
                &sptcrdkb,
                &pdbkb,
                &salkb,
                &slekb,
                &lesnamkb,
                &appde,
                &ctrde,
                &dvyde,
                &pdbtgypr,
                &pdbcbnnu,
                &clmrbbkb,
                &blltypkb,
                &blldtdkb,
                &dshkb,
                &rorapvde,
                &rorconde,
                &fstroade,
                &roraprde,
                &cclcrtde,
                &salscdde,
                &salaprde,
                &supapdfl,
                &clmpr,
                &cmdfl,
                &acprgcfl,
                &oapcplfl,
                &gercplfl,
                &lobtgtfl,
                &rorcrtkb,
                &rorcclkb,
                &rorcrckb,
                &rordstkb,
                &porunnfl,
                &salsupfl,
                &rqdtgtfl,
                &rqdatjfl,
                &orgevdfl,
                &gnlndpfl,
                &estdepde,
                &wrkcplde,
                &octde,
                &acpcplde,
                &lesacnde,
                &rbskb,
                &clmcoddd,
                &clmsopdn,
                &recscdde,
                &drtsgtdn,
                &adckb,
                &adcremrk,
                &grdfdsnu,
                &grdpymde,
                &imgrornu,
                &imgacpnu,
                &oldsalnu,
                &remrk,
                &cchcdtkb,
                &chgpiccd,
                &cngposcd,
                &mgidvdde,
                &mgidvmde,
                &opestdde,
                &pdtstdde,
                &rorstskb,
                &delfl,
                &pdbsvckb,
                &updsysdt,
                &updstfcd,
                &apctfckb,
                &rorpicnm,
                &posnm,
                &csmcstnm,
                &repdtnnm,
                &dtncstnm,
                &bilcstnm,
                &smtcstnm,
                &chgpicnm,
                &cngposnm,
                &updstfnm,
                &telnu,
                &addrk,
                &picnm,
                )

        juchuView.RORNU = rornu
        juchuView.RORBRANU = rorbranu
        juchuView.RORCBNNU = rorcbnnu
        juchuView.GALESTNU = strings.TrimSpace(galestnu)
        juchuView.GALEBNNU = strings.TrimSpace(galebnnu)
        juchuView.GALEHBNU = strings.TrimSpace(galehbnu)
        juchuView.ESTPLRKB = strings.TrimSpace(estplrkb)
        juchuView.ESTNU = strings.TrimSpace(estnu)
        juchuView.ESTBRANU = strings.TrimSpace(estbranu)
        juchuView.ESTCBNNU = strings.TrimSpace(estcbnnu)
        juchuView.AGGNU = strings.TrimSpace(aggnu)
        juchuView.CTRTYPKB = strings.TrimSpace(ctrtypkb)
        juchuView.CTRTBNNU = strings.TrimSpace(ctrtbnnu)
        juchuView.AGGCBNNU = strings.TrimSpace(aggcbnnu)
        juchuView.CLMNU = strings.TrimSpace(clmnu)
        juchuView.MATNM = strings.TrimSpace(matnm)
        juchuView.REPPDNNM = strings.TrimSpace(reppdnnm)
        juchuView.TMFNM = strings.TrimSpace(tmfnm)
        juchuView.RORPICCD = strings.TrimSpace(rorpiccd)
        juchuView.POSCD = strings.TrimSpace(poscd)
        juchuView.CSMCSTCD = strings.TrimSpace(csmcstcd)
        juchuView.REPDTNCD = strings.TrimSpace(repdtncd)
        juchuView.RDDPICNU = rddpicnu
        juchuView.DTNCSTCD = strings.TrimSpace(dtncstcd)
        juchuView.DTNSPCNU = dtnspcnu
        juchuView.BILCSTCD = strings.TrimSpace(bilcstcd)
        juchuView.BILSPCNU = bilspcnu
        juchuView.SMTCSTCD = strings.TrimSpace(smtcstcd)
        juchuView.SCMAT = scmat
        juchuView.SCMPAYCD = strings.TrimSpace(scmpaycd)
        juchuView.SCMPEDDE = scmpedde
        juchuView.RORCFRKB = strings.TrimSpace(rorcfrkb)
        juchuView.CATMKTCD = strings.TrimSpace(catmktcd)
        juchuView.RORKB = strings.TrimSpace(rorkb)
        juchuView.SPTCRDKB = strings.TrimSpace(sptcrdkb)
        juchuView.PDBKB = strings.TrimSpace(pdbkb)
        juchuView.SALKB = strings.TrimSpace(salkb)
        juchuView.SLEKB = strings.TrimSpace(slekb)
        juchuView.LESNAMKB = strings.TrimSpace(lesnamkb)
        juchuView.APPDE = appde
        juchuView.CTRDE = ctrde
        juchuView.DVYDE = dvyde
        juchuView.PDBTGYPR = pdbtgypr
        juchuView.PDBCBNNU = strings.TrimSpace(pdbcbnnu)
        juchuView.CLMRBBKB = strings.TrimSpace(clmrbbkb)
        juchuView.BLLTYPKB = strings.TrimSpace(blltypkb)
        juchuView.BLLDTDKB = strings.TrimSpace(blldtdkb)
        juchuView.DSHKB = strings.TrimSpace(dshkb)
        juchuView.RORAPVDE = rorapvde
        juchuView.RORCONDE = rorconde
        juchuView.FSTROADE = fstroade
        juchuView.RORAPRDE = roraprde
        juchuView.CCLCRTDE = cclcrtde
        juchuView.SALSCDDE = salscdde
        juchuView.SALAPRDE = salaprde
        juchuView.SUPAPDFL = strings.TrimSpace(supapdfl)
        juchuView.CLMPR = clmpr
        juchuView.CMDFL = strings.TrimSpace(cmdfl)
        juchuView.ACPRGCFL = strings.TrimSpace(acprgcfl)
        juchuView.OAPCPLFL = strings.TrimSpace(oapcplfl)
        juchuView.GERCPLFL = strings.TrimSpace(gercplfl)
        juchuView.LOBTGTFL = strings.TrimSpace(lobtgtfl)
        juchuView.RORCRTKB = strings.TrimSpace(rorcrtkb)
        juchuView.RORCCLKB = strings.TrimSpace(rorcclkb)
        juchuView.RORCRCKB = strings.TrimSpace(rorcrckb)
        juchuView.RORDSTKB = strings.TrimSpace(rordstkb)
        juchuView.PORUNNFL = strings.TrimSpace(porunnfl)
        juchuView.SALSUPFL = strings.TrimSpace(salsupfl)
        juchuView.RQDTGTFL = strings.TrimSpace(rqdtgtfl)
        juchuView.RQDATJFL = strings.TrimSpace(rqdatjfl)
        juchuView.ORGEVDFL = strings.TrimSpace(orgevdfl)
        juchuView.GNLNDPFL = strings.TrimSpace(gnlndpfl)
        juchuView.ESTDEPDE = estdepde
        juchuView.WRKCPLDE = wrkcplde
        juchuView.OCTDE = octde
        juchuView.ACPCPLDE = acpcplde
        juchuView.LESACNDE = lesacnde
        juchuView.RBSKB = strings.TrimSpace(rbskb)
        juchuView.CLMCODDD = clmcoddd
        juchuView.CLMSOPDN = clmsopdn
        juchuView.RECSCDDE = recscdde
        juchuView.DRTSGTDN = drtsgtdn
        juchuView.ADCKB = strings.TrimSpace(adckb)
        juchuView.ADCREMRK = strings.TrimSpace(adcremrk)
        juchuView.GRDFDSNU = strings.TrimSpace(grdfdsnu)
        juchuView.GRDPYMDE = grdpymde
        juchuView.IMGRORNU = strings.TrimSpace(imgrornu)
        juchuView.IMGACPNU = strings.TrimSpace(imgacpnu)
        juchuView.OLDSALNU = strings.TrimSpace(oldsalnu)
        juchuView.REMRK = strings.TrimSpace(remrk)
        juchuView.CCHCDTKB = strings.TrimSpace(cchcdtkb)
        juchuView.CHGPICCD = strings.TrimSpace(chgpiccd)
        juchuView.CNGPOSCD = strings.TrimSpace(cngposcd)
        juchuView.MGIDVDDE = mgidvdde
        juchuView.MGIDVMDE = mgidvmde
        juchuView.OPESTDDE = opestdde
        juchuView.PDTSTDDE = pdtstdde
        juchuView.RORSTSKB = strings.TrimSpace(rorstskb)
        juchuView.DELFL = strings.TrimSpace(delfl)
        juchuView.PDBSVCKB = strings.TrimSpace(pdbsvckb)
        juchuView.UPDSYSDT = updsysdt
        juchuView.UPDSTFCD = strings.TrimSpace(updstfcd)
        juchuView.APCTFCKB = strings.TrimSpace(apctfckb)
        juchuView.RORPICNM = strings.TrimSpace(rorpicnm)
        juchuView.POSNM = strings.TrimSpace(posnm)
        juchuView.CSMCSTNM = strings.TrimSpace(csmcstnm)
        juchuView.REPDTNNM = strings.TrimSpace(repdtnnm)
        juchuView.DTNCSTNM = strings.TrimSpace(dtncstnm)
        juchuView.BILCSTNM = strings.TrimSpace(bilcstnm)
        juchuView.SMTCSTNM = strings.TrimSpace(smtcstnm)
        juchuView.CHGPICNM = strings.TrimSpace(chgpicnm)
        juchuView.CNGPOSNM = strings.TrimSpace(cngposnm)
        juchuView.UPDSTFNM = strings.TrimSpace(updstfnm)
        juchuView.TELNU = strings.TrimSpace(telnu)
        juchuView.ADDRK = strings.TrimSpace(addrk)
        juchuView.PICNM = strings.TrimSpace(picnm)

        juchuViews.Views = append(juchuViews.Views, juchuView)
	}

    return juchuViews, nil
}

func buildQueryJuchu(torihikiCd string, juchuNo string, saleYoteiYmdFrom string , saleYoteiYmdTo string, nonyuName string, telNo string, isKanryoChk string) string {
    // SQL
    strQuery := "SELECT" +
                "    RORNU," +
                "    RORBRANU," +
                "    RORCBNNU," +
                "    NVL(GALESTNU, ' ')," +
                "    NVL(GALEBNNU, ' ')," +
                "    NVL(GALEHBNU, ' ')," +
                "    NVL(ESTPLRKB, ' ')," +
                "    NVL(ESTNU, ' ')," +
                "    NVL(ESTBRANU, ' ')," +
                "    NVL(ESTCBNNU, ' ')," +
                "    NVL(AGGNU, ' ')," +
                "    NVL(CTRTYPKB, ' ')," +
                "    NVL(CTRTBNNU, ' ')," +
                "    NVL(AGGCBNNU, ' ')," +
                "    NVL(CLMNU, ' ')," +
                "    NVL(MATNM, ' ')," +
                "    NVL(REPPDNNM, ' ')," +
                "    NVL(TMFNM, ' ')," +
                "    NVL(RORPICCD, ' ')," +
                "    NVL(POSCD, ' ')," +
                "    NVL(CSMCSTCD, ' ')," +
                "    NVL(REPDTNCD, ' ')," +
                "    NVL(RDDPICNU, 0)," +
                "    NVL(DTNCSTCD, ' ')," +
                "    NVL(DTNSPCNU, 0)," +
                "    NVL(BILCSTCD, ' ')," +
                "    NVL(BILSPCNU, 0)," +
                "    NVL(SMTCSTCD, ' ')," +
                "    NVL(SCMAT, 0)," +
                "    NVL(SCMPAYCD, ' ')," +
                "    NVL(SCMPEDDE, 0)," +
                "    NVL(RORCFRKB, ' ')," +
                "    NVL(CATMKTCD, ' ')," +
                "    NVL(RORKB, ' ')," +
                "    NVL(SPTCRDKB, ' ')," +
                "    NVL(PDBKB, ' ')," +
                "    NVL(SALKB, ' ')," +
                "    NVL(SLEKB, ' ')," +
                "    NVL(LESNAMKB, ' ')," +
                "    NVL(APPDE, 0)," +
                "    NVL(CTRDE, 0)," +
                "    NVL(DVYDE, 0)," +
                "    NVL(PDBTGYPR, 0)," +
                "    NVL(PDBCBNNU, ' ')," +
                "    NVL(CLMRBBKB, ' ')," +
                "    NVL(BLLTYPKB, ' ')," +
                "    NVL(BLLDTDKB, ' ')," +
                "    NVL(DSHKB, ' ')," +
                "    NVL(RORAPVDE, 0)," +
                "    NVL(RORCONDE, 0)," +
                "    NVL(FSTROADE, 0)," +
                "    NVL(RORAPRDE, 0)," +
                "    NVL(CCLCRTDE, 0)," +
                "    NVL(SALSCDDE, 0)," +
                "    NVL(SALAPRDE, 0)," +
                "    NVL(SUPAPDFL, ' ')," +
                "    NVL(CLMPR, 0)," +
                "    NVL(CMDFL, ' ')," +
                "    NVL(ACPRGCFL, ' ')," +
                "    NVL(OAPCPLFL, ' ')," +
                "    NVL(GERCPLFL, ' ')," +
                "    NVL(LOBTGTFL, ' ')," +
                "    NVL(RORCRTKB, ' ')," +
                "    NVL(RORCCLKB, ' ')," +
                "    NVL(RORCRCKB, ' ')," +
                "    NVL(RORDSTKB, ' ')," +
                "    NVL(PORUNNFL, ' ')," +
                "    NVL(SALSUPFL, ' ')," +
                "    NVL(RQDTGTFL, ' ')," +
                "    NVL(RQDATJFL, ' ')," +
                "    NVL(ORGEVDFL, ' ')," +
                "    NVL(GNLNDPFL, ' ')," +
                "    NVL(ESTDEPDE, 0)," +
                "    NVL(WRKCPLDE, 0)," +
                "    NVL(OCTDE, 0)," +
                "    NVL(ACPCPLDE, 0)," +
                "    NVL(LESACNDE, 0)," +
                "    NVL(RBSKB, ' ')," +
                "    NVL(CLMCODDD, 0)," +
                "    NVL(CLMSOPDN, 0)," +
                "    NVL(RECSCDDE, 0)," +
                "    NVL(DRTSGTDN, 0)," +
                "    NVL(ADCKB, ' ')," +
                "    NVL(ADCREMRK, ' ')," +
                "    NVL(GRDFDSNU, ' ')," +
                "    NVL(GRDPYMDE, 0)," +
                "    NVL(IMGRORNU, ' ')," +
                "    NVL(IMGACPNU, ' ')," +
                "    NVL(OLDSALNU, ' ')," +
                "    NVL(REMRK, ' ')," +
                "    NVL(CCHCDTKB, ' ')," +
                "    NVL(CHGPICCD, ' ')," +
                "    NVL(CNGPOSCD, ' ')," +
                "    NVL(MGIDVDDE, 0)," +
                "    NVL(MGIDVMDE, 0)," +
                "    NVL(OPESTDDE, 0)," +
                "    NVL(PDTSTDDE, 0)," +
                "    NVL(RORSTSKB, ' ')," +
                "    NVL(DELFL, ' ')," +
                "    NVL(PDBSVCKB, ' ')," +
                "    UPDSYSDT," +
                "    NVL(UPDSTFCD, ' ')," +
                "    NVL(APCTFCKB, ' ')," +
                "    NVL(RORPICNM, ' ')," +
                "    NVL(POSNM, ' ')," +
                "    NVL(CSMCSTNM, ' ')," +
                "    NVL(REPDTNNM, ' ')," +
                "    NVL(DTNCSTNM, ' ')," +
                "    NVL(BILCSTNM, ' ')," +
                "    NVL(SMTCSTNM, ' ')," +
                "    NVL(CHGPICNM, ' ')," +
                "    NVL(CNGPOSNM, ' ')," +
                "    NVL(UPDSTFNM, ' ')," +
                "    NVL(TELNU, ' ')," +
                "    NVL(ADDRK, ' ')," +
                "    NVL(PICNM, ' ')" +
                " FROM" +
                "    CRMP_SARECO_DIF" +
                " WHERE" +
                "    1 = 1"

                // リクエストパラメータが存在する場合のみ検索条件に含める
                if len(torihikiCd) > 0 {
                    strQuery += " AND CSMCSTCD = '" + torihikiCd + "'"
                }
                if len(juchuNo) > 0 {
                    strQuery += " AND RORNU = '" + juchuNo + "'"
                }
                if len(saleYoteiYmdFrom) > 0 {
                    strQuery += " AND SALSCDDE >= " + saleYoteiYmdFrom
                }
                if len(saleYoteiYmdTo) > 0 {
                    strQuery += " AND SALSCDDE <= " + saleYoteiYmdTo
                }
                if len(nonyuName) > 0 {
                    strQuery += " AND CSMCSTNM LIKE '%" + nonyuName + "%'"
                }
                if len(telNo) > 0 {
                    strQuery += " AND TELNU = '" + telNo + "'"
                }
                if len(isKanryoChk) > 0 && isKanryoChk == "FALSE" {
                    strQuery += " AND SALAPRDE IS NULL"
                }

                strQuery += " AND ROWNUM <= 100"
                strQuery += " ORDER BY RORNU, RORBRANU, RORCBNNU"

                fmt.Println(strQuery)

    return strQuery
}