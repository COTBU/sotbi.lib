package onec

import "time"

//nolint:lll
type PaymentDocument struct {
	AccountBalanceID       uint64     `json:"account_balance_id"`
	DocumentType           string     `json:"document_type,omitempty"            mapstructure:"СекцияДокумент"`
	Number                 string     `json:"number,omitempty"                   mapstructure:"Номер"`
	DataStr                string     `json:"-"                                  mapstructure:"Дата"`
	Data                   *time.Time `json:"date,omitempty"                     mapstructure:"-"`
	WrittenOffDateStr      string     `json:"-"                                  mapstructure:"ДатаСписано"`
	WrittenOffDate         *time.Time `json:"written_off_date,omitempty"         mapstructure:"-"`
	IncomeDateStr          string     `json:"-"                                  mapstructure:"ДатаПоступило"`
	IncomeDate             *time.Time `json:"income_date,omitempty"              mapstructure:"-"`
	Summ                   float64    `json:"summ"                               mapstructure:"Сумма"`
	RectDateStr            string     `json:"-"                                  mapstructure:"КвитанцияДата"`
	RectTimeStr            string     `json:"-"                                  mapstructure:"КвитанцияВремя"`
	RectDateTime           *time.Time `json:"rect_date,omitempty"                mapstructure:"-"`
	RectContent            *string    `json:"rect_content,omitempty"             mapstructure:"КвитанцияСодержание,omitempty"`
	PayerAccount           string     `json:"payer_account,omitempty"            mapstructure:"ПлательщикСчет"`
	Payer                  string     `json:"payer,omitempty"                    mapstructure:"Плательщик"`
	PayerINN               string     `json:"payer_inn,omitempty"                mapstructure:"ПлательщикИНН"`
	PayerKPP               *string    `json:"payer_kpp,omitempty"                mapstructure:"ПлательщикКПП,omitempty"`
	Payer1                 *string    `json:"payer1,omitempty"                   mapstructure:"Плательщик1"`
	Payer2                 *string    `json:"payer2,omitempty"                   mapstructure:"Плательщик2,omitempty"`
	Payer3                 *string    `json:"payer3,omitempty"                   mapstructure:"Плательщик3,omitempty"`
	Payer4                 *string    `json:"payer4,omitempty"                   mapstructure:"Плательщик4,omitempty"`
	PayerCurrentAccount    string     `json:"payer_current_account,omitempty"    mapstructure:"ПлательщикРасчСчет"`
	PayerBank1             string     `json:"payer_bank1,omitempty"              mapstructure:"ПлательщикБанк1"`
	PayerBank2             *string    `json:"payer_bank2,omitempty"              mapstructure:"ПлательщикБанк2,omitempty"`
	PayerBIK               string     `json:"payer_bik,omitempty"                mapstructure:"ПлательщикБИК"`
	PayerCorrAccount       string     `json:"payer_corr_account,omitempty"       mapstructure:"ПлательщикКорсчет"`
	ReceiverAccount        string     `json:"receiver_account,omitempty"         mapstructure:"ПолучательСчет"`
	Receiver               string     `json:"receiver,omitempty"                 mapstructure:"Получатель"`
	ReceiverINN            string     `json:"receiver_inn,omitempty"             mapstructure:"ПолучательИНН"`
	ReceiverKPP            *string    `json:"receiver_kpp,omitempty"             mapstructure:"ПолучательКПП,omitempty"`
	Receiver1              *string    `json:"receiver1,omitempty"                mapstructure:"Получатель1,omitempty"`
	Receiver2              *string    `json:"receiver2,omitempty"                mapstructure:"Получатель2,omitempty"`
	Receiver3              *string    `json:"receiver3,omitempty"                mapstructure:"Получатель3,omitempty"`
	Receiver4              *string    `json:"receiver4,omitempty"                mapstructure:"Получатель4,omitempty"`
	ReceiverCurrentAccount string     `json:"receiver_current_account,omitempty" mapstructure:"ПолучательРасчСчет"` //nolint:lll
	ReceiverBank1          string     `json:"receiver_bank1,omitempty"           mapstructure:"ПолучательБанк1"`
	ReceiverBank2          *string    `json:"receiver_bank2,omitempty"           mapstructure:"ПолучательБанк2,omitempty"`
	ReceiverBIK            string     `json:"receiver_bik,omitempty"             mapstructure:"ПолучательБИК"`
	ReceiverCorrAccount    string     `json:"receiver_corr_account,omitempty"    mapstructure:"ПолучательКорсчет"`
	PaymentType            *string    `json:"payment_type,omitempty"             mapstructure:"ВидПлатежа,omitempty"`
	PaymentPurposeCode     *string    `json:"payment_purpose_code,omitempty"     mapstructure:"КодНазПлатежа,omitempty"`
	UIN                    *string    `json:"uin,omitempty"                      mapstructure:"Код,omitempty"`
	PaymentPurpose         string     `json:"payment_purpose,omitempty"          mapstructure:"НазначениеПлатежа"`
	PaymentPurpose1        *string    `json:"payment_purpose1,omitempty"         mapstructure:"НазначениеПлатежа1,omitempty"`
	PaymentPurpose2        *string    `json:"payment_purpose2,omitempty"         mapstructure:"НазначениеПлатежа2,omitempty"`
	PaymentPurpose3        *string    `json:"payment_purpose3,omitempty"         mapstructure:"НазначениеПлатежа3,omitempty"`
	PaymentPurpose4        *string    `json:"payment_purpose4,omitempty"         mapstructure:"НазначениеПлатежа4,omitempty"`
	PaymentPurpose5        *string    `json:"payment_purpose5,omitempty"         mapstructure:"НазначениеПлатежа5,omitempty"`
	PaymentPurpose6        *string    `json:"payment_purpose6,omitempty"         mapstructure:"НазначениеПлатежа6,omitempty"`
	CompilerStatus         *string    `json:"compiler_status,omitempty"          mapstructure:"СтатусСоставителя,omitempty"`
	OKATO                  *string    `json:"okato,omitempty"                    mapstructure:"ОКАТО,omitempty"`
	IndicatorKBK           *string    `json:"indicator_kbk,omitempty"            mapstructure:"ПоказательКБК,omitempty"`
	IndicatorBasics        *string    `json:"indicator_basics,omitempty"         mapstructure:"ПоказательОснования,omitempty"`
	IndicatorPeriod        *string    `json:"indicator_period,omitempty"         mapstructure:"ПоказательПериода,omitempty"`
	IndicatorNumber        *string    `json:"indicator_number,omitempty"         mapstructure:"ПоказательНомера,omitempty"`
	IndicatorDateStr       string     `json:"-"                                  mapstructure:"ПоказательДаты,omitempty"`
	IndicatorDate          *time.Time `json:"indicator_date,omitempty"           mapstructure:"-"`
	IndicatorType          *string    `json:"indicator_type,omitempty"           mapstructure:"ПоказательТипа,omitempty"`
	Priority               *uint      `json:"priority,omitempty"                 mapstructure:"Очередность,omitempty"`
	DefrayalType           *string    `json:"defrayal_type,omitempty"            mapstructure:"ВидОплаты,omitempty"`
	AcceptanceTerm         *string    `json:"acceptance_term,omitempty"          mapstructure:"СрокАкцепта,omitempty"`
	TypeLetterCredit       *string    `json:"type_letter_credit,omitempty"       mapstructure:"ВидАккредитива,omitempty"`
	PaymentTerm            *string    `json:"payment_term,omitempty"             mapstructure:"СрокПлатежа,omitempty"`
	PaymentCondition1      *string    `json:"paymen_condition1,omitempty"        mapstructure:"УсловиеОплаты1,omitempty"`
	PaymentCondition2      *string    `json:"paymen_condition2,omitempty"        mapstructure:"УсловиеОплаты1,omitempty"`
	PaymentCondition3      *string    `json:"paymen_condition3,omitempty"        mapstructure:"УсловиеОплаты1,omitempty"`
	PaymentBy              *string    `json:"payment_by,omitempty"               mapstructure:"ПлатежПоПредст,omitempty"`
	AdditionalTerms        *string    `json:"additional_terms,omitempty"         mapstructure:"ДополнУсловия,omitempty"`
	SupplierAccountNumber  *string    `json:"supplier_account_number,omitempty"  mapstructure:"НомерСчетаПоставщика,omitempty"`
	DocumentSendingDateStr *string    `json:"-"                                  mapstructure:"ДатаОтсылкиДок,omitempty"`
	DocumentSendingDate    *time.Time `json:"document_sending_date,omitempty"    mapstructure:"-"`
}
