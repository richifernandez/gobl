package it

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
	"github.com/invopop/gobl/tax"
)

// Tag keys detemoined from the "Natura" field from FatturaPA.
const (
	TagKeyExcluded                                cbc.Key = "excluded"
	TagKeyNotSubject                              cbc.Key = "not-subject" // not used since 1 Jan 2021
	TagKeyNotSubjectArticle7                      cbc.Key = "not-subject-article-7"
	TagKeyNotSubjectOther                         cbc.Key = "not-subject-other"
	TagKeyNotTaxable                              cbc.Key = "not-taxable" // not used since 1 Jan 2021
	TagKeyNotTaxableExports                       cbc.Key = "not-taxable-exports"
	TagKeyNotTaxableIntraCommunity                cbc.Key = "not-taxable-intra-community"
	TagKeyNotTaxableSanMerino                     cbc.Key = "not-taxable-san-merino"
	TagKeyNotTaxableExportSupplies                cbc.Key = "not-taxable-export-supplies"
	TagKeyNotTaxableDeclarationOfIntent           cbc.Key = "not-taxable-declaration-of-intent"
	TagKeyNotTaxableOther                         cbc.Key = "not-taxable-other"
	TagKeyExempt                                  cbc.Key = "exempt"
	TagKeyMarginRegime                            cbc.Key = "margin-regime"
	TagKeyReverseCharge                           cbc.Key = "reverse-charge"
	TagKeyReverseChargeScrap                      cbc.Key = "reverse-charge-scrap"
	TagKeyReverseChargePreciousMetals             cbc.Key = "reverse-charge-precious-metals"
	TagKeyReverseChargeConstructionSubcontracting cbc.Key = "reverse-charge-construction-subcontracting"
	TagKeyReverseChargeBuildings                  cbc.Key = "reverse-charge-buildings"
	TagKeyReverseChargeMobile                     cbc.Key = "reverse-charge-mobile"
	TagKeyReverseChargeElectronics                cbc.Key = "reverse-charge-electronics"
	TagKeyReverseChargeConstruction               cbc.Key = "reverse-charge-construction"
	TagKeyReverseChargeEnergy                     cbc.Key = "reverse-charge-energy"
	TagKeyReverseChargeOther                      cbc.Key = "reverse-charge-other"
	TagKeyVATEU                                   cbc.Key = "vat-eu"
)

var vatZeroTaxTags = []*tax.Tag{
	{
		Key: TagKeyExcluded,
		Desc: i18n.String{
			i18n.EN: "Excluded pursuant to Art. 15, DPR 633/72",
			i18n.IT: "Escluse ex. art. 15 del D.P.R. 633/1972",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N1",
		},
	},
	{
		Key: TagKeyNotSubject,
		Desc: i18n.String{
			i18n.EN: "Not subject (this code is no longer permitted to use on invoices emitted from 1 January 2021)",
			i18n.IT: "Non soggette (questo codice non è più utilizzabile a partire dal 1° gennaio 2021)",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N2",
		},
	},
	{
		Key: TagKeyNotSubjectArticle7,
		Desc: i18n.String{
			i18n.EN: "Not subject pursuant to Art. 7, DPR 633/72",
			i18n.IT: "Non soggette ex. art. 7 del D.P.R. 633/72",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N2.1",
		},
	},
	{
		Key: TagKeyNotSubjectOther,
		Desc: i18n.String{
			i18n.EN: "Not subject - other",
			i18n.IT: "Non soggette - altri casi",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N2.2",
		},
	},
	{
		Key: TagKeyNotTaxable,
		Desc: i18n.String{
			i18n.EN: "Not taxable (this code is no longer permitted to use on invoices emitted from 1 January 2021)",
			i18n.IT: "Non imponibili (questo codice non è più utilizzabile a partire dal 1° gennaio 2021)",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N3",
		},
	},
	{
		Key: TagKeyNotTaxableExports,
		Desc: i18n.String{
			i18n.EN: "Not taxable - exports",
			i18n.IT: "Non imponibili - esportazioni",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N3.1",
		},
	},
	{
		Key: TagKeyNotTaxableIntraCommunity,
		Desc: i18n.String{
			i18n.EN: "Not taxable - intra-community supplies",
			i18n.IT: "Non imponibili - cessioni intracomunitarie",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N3.2",
		},
	},
	{
		Key: TagKeyNotTaxableSanMerino,
		Desc: i18n.String{
			i18n.EN: "Not taxable - transfers to San Marino",
			i18n.IT: "Non imponibili - cessioni verso San Marino",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N3.3",
		},
	},
	{
		Key: TagKeyNotTaxableExportSupplies,
		Desc: i18n.String{
			i18n.EN: "Not taxable - export supplies of goods and services",
			i18n.IT: "Non Imponibili - operazioni assimilate alle cessioni all'esportazione",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N3.4",
		},
	},
	{
		Key: TagKeyNotTaxableDeclarationOfIntent,
		Desc: i18n.String{
			i18n.EN: "Not taxable - declaration of intent",
			i18n.IT: "Non imponibili - dichiarazioni d'intento",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N3.5",
		},
	},
	{
		Key: TagKeyNotTaxableOther,
		Desc: i18n.String{
			i18n.EN: "Not taxable - other",
			i18n.IT: "Non imponibili - altre operazioni che non concorrono alla formazione del plafond",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N3.6",
		},
	},
	{
		Key: TagKeyExempt,
		Desc: i18n.String{
			i18n.EN: "Exempt",
			i18n.IT: "Esenti",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N4",
		},
	},
	{
		Key: TagKeyMarginRegime,
		Desc: i18n.String{
			i18n.EN: "Margin regime / VAT not exposed",
			i18n.IT: "Regime del margine/IVA non esposta in fattura",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N5",
		},
	},
	{
		Key: TagKeyReverseCharge,
		Desc: i18n.String{
			i18n.EN: "Reverse charge (for transactions in reverse charge or for self invoicing for purchase of extra UE services or for import of goods only in the cases provided for) — (this code is no longer permitted to use on invoices emitted from 1 January 2021)",
			i18n.IT: "Inversione contabile (per operazioni in regime di inversione contabile o per autofattura per acquisti di servizi extra UE o per importazioni di beni solo nei casi previsti) — (questo codice non è più utilizzabile a partire dal 1° gennaio 2021)",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6",
		},
	},
	{
		Key: TagKeyReverseChargeScrap,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - Transfer of scrap and of other recyclable materials",
			i18n.IT: "Inversione contabile - cessione di rottami e altri materiali di recupero",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.1",
		},
	},
	{
		Key: TagKeyReverseChargePreciousMetals,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - Transfer of gold and pure silver pursuant to law 7/2000 as well as used jewelery to OPO",
			i18n.IT: "Inversione contabile - cessione di oro e argento ai sensi della legge 7/2000 nonché di oreficeria usata ad OPO",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.2",
		},
	},
	{
		Key: TagKeyReverseChargeConstructionSubcontracting,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - Construction subcontracting",
			i18n.IT: "Inversione contabile - subappalto nel settore edile",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.3",
		},
	},
	{
		Key: TagKeyReverseChargeBuildings,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - Transfer of buildings",
			i18n.IT: "Inversione contabile - cessione di fabbricati",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.4",
		},
	},
	{
		Key: TagKeyReverseChargeMobile,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - Transfer of mobile phones",
			i18n.IT: "Inversione contabile - cessione di telefoni cellulari",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.5",
		},
	},
	{
		Key: TagKeyReverseChargeElectronics,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - Transfer of electronic products",
			i18n.IT: "Inversione contabile - cessione di prodotti elettronici",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.6",
		},
	},
	{
		Key: TagKeyReverseChargeConstruction,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - provisions in the construction and related sectors",
			i18n.IT: "Inversione contabile - prestazioni comparto edile e settori connessi",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.7",
		},
	},
	{
		Key: TagKeyReverseChargeEnergy,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - transactions in the energy sector",
			i18n.IT: "Inversione contabile - operazioni settore energetico",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.8",
		},
	},
	{
		Key: TagKeyReverseChargeOther,
		Desc: i18n.String{
			i18n.EN: "Reverse charge - other cases",
			i18n.IT: "Inversione contabile - altri casi",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N6.9",
		},
	},
	{
		Key: TagKeyVATEU,
		Desc: i18n.String{
			i18n.EN: "VAT paid in other EU countries (telecommunications, tele-broadcasting and electronic services provision pursuant to Art. 7 -octies letter a, b, art. 74-sexies Italian Presidential Decree 633/72)",
			i18n.IT: "IVA assolta in altro stato UE (prestazione di servizi di telecomunicazioni, tele-radiodiffusione ed elettronici ex art. 7-octies lett. a, b, art. 74-sexies DPR 633/72)",
		},
		Meta: cbc.Meta{
			KeyFatturaPANatura: "N7",
		},
	},
}
