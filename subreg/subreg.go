package subreg

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"reflect"

	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type LoginContainer struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Login_Container"`

	Response *LoginResponse `xml:"response,omitempty"`
}

type Login struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Login"`

	Login    string `xml:"login,omitempty"`
	Password string `xml:"password,omitempty"`
}

type CheckDomainContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Container"`

	Response *CheckDomainResponse `xml:"response,omitempty"`
}

type CheckDomain struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain"`

	Ssid   string             `xml:"ssid,omitempty"`
	Domain string             `xml:"domain,omitempty"`
	Params *CheckDomainParams `xml:"params,omitempty"`
}

type InfoDomainContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Container"`

	Response *InfoDomainResponse `xml:"response,omitempty"`
}

type InfoDomain struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type InfoDomainCZContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Container"`

	Response *InfoDomainCZResponse `xml:"response,omitempty"`
}

type InfoDomainCZ struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type DomainsListContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Container"`

	Response *DomainsListResponse `xml:"response,omitempty"`
}

type DomainsList struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List"`

	Ssid string `xml:"ssid,omitempty"`
}

type SetAutorenewContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew_Container"`

	Response *SetAutorenewResponse `xml:"response,omitempty"`
}

type SetAutorenew struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew"`

	Ssid      string `xml:"ssid,omitempty"`
	Domain    string `xml:"domain,omitempty"`
	Autorenew string `xml:"autorenew,omitempty"`
}

type CreateContactContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Container"`

	Response *CreateContactResponse `xml:"response,omitempty"`
}

type CreateContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact"`

	Ssid    string                `xml:"ssid,omitempty"`
	Contact *CreateContactContact `xml:"contact,omitempty"`
}

type UpdateContactContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Container"`

	Response *UpdateContactResponse `xml:"response,omitempty"`
}

type UpdateContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact"`

	Ssid    string                `xml:"ssid,omitempty"`
	Contact *UpdateContactContact `xml:"contact,omitempty"`
}

type InfoContactContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Container"`

	Response *InfoContactResponse `xml:"response,omitempty"`
}

type InfoContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact"`

	Ssid    string              `xml:"ssid,omitempty"`
	Contact *InfoContactContact `xml:"contact,omitempty"`
}

type ContactsListContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Container"`

	Response *ContactsListResponse `xml:"response,omitempty"`
}

type ContactsList struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List"`

	Ssid string `xml:"ssid,omitempty"`
}

type CheckObjectContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object_Container"`

	Response *CheckObjectResponse `xml:"response,omitempty"`
}

type CheckObject struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object"`

	Ssid   string `xml:"ssid,omitempty"`
	Object string `xml:"object,omitempty"`
	Id     string `xml:"id,omitempty"`
}

type InfoObjectContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Container"`

	Response *InfoObjectResponse `xml:"response,omitempty"`
}

type InfoObject struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object"`

	Ssid   string `xml:"ssid,omitempty"`
	Object string `xml:"object,omitempty"`
	Id     string `xml:"id,omitempty"`
}

type MakeOrderContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Container"`

	Response *MakeOrderResponse `xml:"response,omitempty"`
}

type MakeOrder struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order"`

	Ssid  string          `xml:"ssid,omitempty"`
	Order *MakeOrderOrder `xml:"order,omitempty"`
}

type InfoOrderContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Container"`

	Response *InfoOrderResponse `xml:"response,omitempty"`
}

type InfoOrder struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order"`

	Ssid  string `xml:"ssid,omitempty"`
	Order int32  `xml:"order,omitempty"`
}

type GetCreditContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Container"`

	Response *GetCreditResponse `xml:"response,omitempty"`
}

type GetCredit struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit"`

	Ssid string `xml:"ssid,omitempty"`
}

type GetAccountingsContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Container"`

	Response *GetAccountingsResponse `xml:"response,omitempty"`
}

type GetAccountings struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings"`

	Ssid string `xml:"ssid,omitempty"`
	From string `xml:"from,omitempty"`
	To   string `xml:"to,omitempty"`
}

type ClientPaymentContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment_Container"`

	Response *ClientPaymentResponse `xml:"response,omitempty"`
}

type ClientPayment struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment"`

	Ssid     string  `xml:"ssid,omitempty"`
	Username string  `xml:"username,omitempty"`
	Amount   float64 `xml:"amount,omitempty"`
	Currency string  `xml:"currency,omitempty"`
}

type CreditCorrectionContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction_Container"`

	Response *CreditCorrectionResponse `xml:"response,omitempty"`
}

type CreditCorrection struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction"`

	Ssid     string  `xml:"ssid,omitempty"`
	Username string  `xml:"username,omitempty"`
	Amount   float64 `xml:"amount,omitempty"`
	Reason   string  `xml:"reason,omitempty"`
}

type PricelistContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Container"`

	Response *PricelistResponse `xml:"response,omitempty"`
}

type Pricelist struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist"`

	Ssid string `xml:"ssid,omitempty"`
}

type PricesContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Container"`

	Response *PricesResponse `xml:"response,omitempty"`
}

type Prices struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Prices"`

	Ssid string `xml:"ssid,omitempty"`
	Tld  string `xml:"tld,omitempty"`
}

type GetPricelistContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Container"`

	Response *GetPricelistResponse `xml:"response,omitempty"`
}

type GetPricelist struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist"`

	Ssid      string `xml:"ssid,omitempty"`
	Pricelist string `xml:"pricelist,omitempty"`
}

type SetPricesContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Container"`

	Response *SetPricesResponse `xml:"response,omitempty"`
}

type SetPrices struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices"`

	Ssid      string            `xml:"ssid,omitempty"`
	Pricelist string            `xml:"pricelist,omitempty"`
	Tld       string            `xml:"tld,omitempty"`
	Currency  string            `xml:"currency,omitempty"`
	Prices    []*SetPricesPrice `xml:"prices,omitempty"`
}

type DownloadDocumentContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document_Container"`

	Response *DownloadDocumentResponse `xml:"response,omitempty"`
}

type DownloadDocument struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document"`

	Ssid string `xml:"ssid,omitempty"`
	Id   int32  `xml:"id,omitempty"`
}

type UploadDocumentContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document_Container"`

	Response *UploadDocumentResponse `xml:"response,omitempty"`
}

type UploadDocument struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document"`

	Ssid     string `xml:"ssid,omitempty"`
	Name     string `xml:"name,omitempty"`
	Document string `xml:"document,omitempty"`
	Type_    string `xml:"type,omitempty"`
	Filetype string `xml:"filetype,omitempty"`
}

type ListDocumentsContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Container"`

	Response *ListDocumentsResponse `xml:"response,omitempty"`
}

type ListDocuments struct {
	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents"`

	Ssid string `xml:"ssid,omitempty"`
}

type UsersListContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_Container"`

	Response *UsersListResponse `xml:"response,omitempty"`
}

type UsersList struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Users_List"`

	Ssid string `xml:"ssid,omitempty"`
}

type AnycastADDZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone_Container"`

	Response *AnycastADDZoneResponse `xml:"response,omitempty"`
}

type AnycastADDZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
	Server int32  `xml:"server,omitempty"`
}

type AnycastRemoveZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone_Container"`

	Response *AnycastRemoveZoneResponse `xml:"response,omitempty"`
}

type AnycastRemoveZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
	Server int32  `xml:"server,omitempty"`
}

type GetDNSZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Zone_Container"`

	Response *GetDNSZoneResponse `xml:"response,omitempty"`
}

type GetDNSZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Zone"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type AddDNSZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone_Container"`

	Response *AddDNSZoneResponse `xml:"response,omitempty"`
}

type AddDNSZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone"`

	Ssid     string `xml:"ssid,omitempty"`
	Domain   string `xml:"domain,omitempty"`
	Template string `xml:"template,omitempty"`
}

type DeleteDNSZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone_Container"`

	Response *DeleteDNSZoneResponse `xml:"response,omitempty"`
}

type DeleteDNSZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type SetDNSZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Container"`

	Response *SetDNSZoneResponse `xml:"response,omitempty"`
}

type SetDNSZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone"`

	Ssid    string              `xml:"ssid,omitempty"`
	Domain  string              `xml:"domain,omitempty"`
	Records []*SetDNSZoneRecord `xml:"records,omitempty"`
}

type AddDNSRecordContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Record_Container"`

	Response *AddDNSRecordResponse `xml:"response,omitempty"`
}

type AddDNSRecord struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Record"`

	Ssid   string              `xml:"ssid,omitempty"`
	Domain string              `xml:"domain,omitempty"`
	Record *AddDNSRecordRecord `xml:"record,omitempty"`
}

type ModifyDNSRecordContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Modify_DNS_Record_Container"`

	Response *ModifyDNSRecordResponse `xml:"response,omitempty"`
}

type ModifyDNSRecord struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Modify_DNS_Record"`

	Ssid   string                 `xml:"ssid,omitempty"`
	Domain string                 `xml:"domain,omitempty"`
	Record *ModifyDNSRecordRecord `xml:"record,omitempty"`
}

type DeleteDNSRecordContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Container"`

	Response *DeleteDNSRecordResponse `xml:"response,omitempty"`
}

type DeleteDNSRecord struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record"`

	Ssid   string                 `xml:"ssid,omitempty"`
	Domain string                 `xml:"domain,omitempty"`
	Record *DeleteDNSRecordRecord `xml:"record,omitempty"`
}

type POLLGetContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get_Container"`

	Response *POLLGetResponse `xml:"response,omitempty"`
}

type POLLGet struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get"`

	Ssid string `xml:"ssid,omitempty"`
}

type POLLAckContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack_Container"`

	Response *POLLAckResponse `xml:"response,omitempty"`
}

type POLLAck struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack"`

	Ssid string `xml:"ssid,omitempty"`
	Id   int32  `xml:"id,omitempty"`
}

type OIBSearchContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Container"`

	Response *OIBSearchResponse `xml:"response,omitempty"`
}

type OIBSearch struct {
	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search"`

	Ssid string `xml:"ssid,omitempty"`
	Oib  string `xml:"oib,omitempty"`
}

type GetCertificateContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate_Container"`

	Response *GetCertificateResponse `xml:"response,omitempty"`
}

type GetCertificate struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate"`

	Ssid    string `xml:"ssid,omitempty"`
	Orderid int32  `xml:"orderid,omitempty"`
}

type GetRedirectsContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects_Container"`

	Response *GetRedirectsResponse `xml:"response,omitempty"`
}

type GetRedirects struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type InSubregContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg_Container"`

	Response *InSubregResponse `xml:"response,omitempty"`
}

type InSubreg struct {
	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type SignDNSZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone_Container"`

	Response *SignDNSZoneResponse `xml:"response,omitempty"`
}

type SignDNSZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type UnsignDNSZoneContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone_Container"`

	Response *UnsignDNSZoneResponse `xml:"response,omitempty"`
}

type UnsignDNSZone struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone"`

	Ssid   string `xml:"ssid,omitempty"`
	Domain string `xml:"domain,omitempty"`
}

type GetDNSInfoContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Container"`

	Response *GetDNSInfoResponse `xml:"response,omitempty"`
}

type GetDNSInfo struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info"`

	Ssid    string `xml:"ssid,omitempty"`
	Domain  string `xml:"domain,omitempty"`
	Dnstype string `xml:"dnstype,omitempty"`
}

type SpecialPricelistContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Container"`

	Response *SpecialPricelistResponse `xml:"response,omitempty"`
}

type SpecialPricelist struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist"`

	Ssid string `xml:"ssid,omitempty"`
}

type GetTLDInfoContainer struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Container"`

	Response *GetTLDInfoResponse `xml:"response,omitempty"`
}

type GetTLDInfo struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info"`

	Ssid string `xml:"ssid,omitempty"`
	Tld  string `xml:"tld,omitempty"`
}

type LoginResponse struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Login_Response"`

	Status string     `xml:"status,omitempty"`
	Data   *LoginData `xml:"data,omitempty"`
	Error  *ErrorInfo `xml:"error,omitempty"`
}

type LoginData struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Login_Data"`

	Ssid string `xml:"ssid,omitempty"`
}

type CheckDomainResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Response"`

	Status string           `xml:"status,omitempty"`
	Data   *CheckDomainData `xml:"data,omitempty"`
	Error  *ErrorInfo       `xml:"error,omitempty"`
}

type CheckDomainParams struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Params"`

	Langinfo string `xml:"lang_info,omitempty"`
}

type CheckDomainPrice struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Price"`

	Amount            float64 `xml:"amount,omitempty"`
	Amountwithtrustee float64 `xml:"amount_with_trustee,omitempty"`
	Premium           int32   `xml:"premium,omitempty"`
	Currency          string  `xml:"currency,omitempty"`
}

type CheckDomainData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Domain_Data"`

	Name            string            `xml:"name,omitempty"`
	Avail           int32             `xml:"avail,omitempty"`
	Existingclaimid string            `xml:"existing_claim_id,omitempty"`
	Price           *CheckDomainPrice `xml:"price,omitempty"`
}

type InfoDomainResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Response"`

	Status string          `xml:"status,omitempty"`
	Data   *InfoDomainData `xml:"data,omitempty"`
	Error  *ErrorInfo      `xml:"error,omitempty"`
}

type InfoDomainContacts struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Contacts"`

	Admin []*InfoDomainContact `xml:"admin,omitempty"`
	Tech  []*InfoDomainContact `xml:"tech,omitempty"`
	Bill  []*InfoDomainContact `xml:"bill,omitempty"`
}

type InfoDomainDsdata struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Dsdata"`

	Tag        string `xml:"tag,omitempty"`
	Alg        string `xml:"alg,omitempty"`
	Digesttype string `xml:"digest_type,omitempty"`
	Digest     string `xml:"digest,omitempty"`
}

type InfoDomainOptions struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Options"`

	Nsset       string              `xml:"nsset,omitempty"`
	Keyset      string              `xml:"keyset,omitempty"`
	Dsdata      []*InfoDomainDsdata `xml:"dsdata,omitempty"`
	Keygroup    string              `xml:"keygroup,omitempty"`
	Quarantined string              `xml:"quarantined,omitempty"`
}

type InfoDomainData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Data"`

	Domain     string              `xml:"domain,omitempty"`
	Contacts   *InfoDomainContacts `xml:"contacts,omitempty"`
	Hosts      []string            `xml:"hosts,omitempty"`
	Registrant *InfoDomainContact  `xml:"registrant,omitempty"`
	ExDate     string              `xml:"exDate,omitempty"`
	CrDate     string              `xml:"crDate,omitempty"`
	TrDate     string              `xml:"trDate,omitempty"`
	UpDate     string              `xml:"upDate,omitempty"`
	Authid     string              `xml:"authid,omitempty"`
	Status     []string            `xml:"status,omitempty"`
	Rgp        []string            `xml:"rgp,omitempty"`
	Autorenew  int32               `xml:"autorenew,omitempty"`
	Premium    int32               `xml:"premium,omitempty"`
	Price      float64             `xml:"price,omitempty"`
	Whoisproxy int32               `xml:"whoisproxy,omitempty"`
	Options    *InfoDomainOptions  `xml:"options,omitempty"`
}

type InfoDomainCZResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Response"`

	Status string            `xml:"status,omitempty"`
	Data   *InfoDomainCZData `xml:"data,omitempty"`
	Error  *ErrorInfo        `xml:"error,omitempty"`
}

type InfoDomainCZContacts struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Contacts"`

	Admin []*InfoDomainCZContact `xml:"admin,omitempty"`
	Tech  []*InfoDomainCZContact `xml:"tech,omitempty"`
	Bill  []*InfoDomainCZContact `xml:"bill,omitempty"`
}

type InfoDomainCZDsdata struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Dsdata"`

	Tag        string `xml:"tag,omitempty"`
	Alg        string `xml:"alg,omitempty"`
	Digesttype string `xml:"digest_type,omitempty"`
	Digest     string `xml:"digest,omitempty"`
}

type InfoDomainCZOptions struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Options"`

	Nsset       string              `xml:"nsset,omitempty"`
	Keyset      string              `xml:"keyset,omitempty"`
	Dsdata      *InfoDomainCZDsdata `xml:"dsdata,omitempty"`
	Keygroup    string              `xml:"keygroup,omitempty"`
	Quarantined string              `xml:"quarantined,omitempty"`
}

type InfoDomainCZData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Data"`

	Domain     string                `xml:"domain,omitempty"`
	Contacts   *InfoDomainCZContacts `xml:"contacts,omitempty"`
	Hosts      []string              `xml:"hosts,omitempty"`
	Registrant *InfoDomainCZContact  `xml:"registrant,omitempty"`
	ExDate     string                `xml:"exDate,omitempty"`
	CrDate     string                `xml:"crDate,omitempty"`
	TrDate     string                `xml:"trDate,omitempty"`
	UpDate     string                `xml:"upDate,omitempty"`
	Status     []string              `xml:"status,omitempty"`
	Rgp        []string              `xml:"rgp,omitempty"`
	Autorenew  int32                 `xml:"autorenew,omitempty"`
	Whoisproxy int32                 `xml:"whoisproxy,omitempty"`
	Options    *InfoDomainCZOptions  `xml:"options,omitempty"`
}

type DomainsListResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Response"`

	Status string           `xml:"status,omitempty"`
	Data   *DomainsListData `xml:"data,omitempty"`
	Error  *ErrorInfo       `xml:"error,omitempty"`
}

type DomainsListDomain struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Domain"`

	Name      string `xml:"name,omitempty"`
	Expire    string `xml:"expire,omitempty"`
	Autorenew int32  `xml:"autorenew,omitempty"`
}

type DomainsListData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Domains_List_Data"`

	Count   int32                `xml:"count,omitempty"`
	Domains []*DomainsListDomain `xml:"domains,omitempty"`
}

type SetAutorenewResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew_Response"`

	Status string            `xml:"status,omitempty"`
	Data   *SetAutorenewData `xml:"data,omitempty"`
	Error  *ErrorInfo        `xml:"error,omitempty"`
}

type SetAutorenewData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Autorenew_Data"`
}

type CreateContactResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Response"`

	Status string             `xml:"status,omitempty"`
	Data   *CreateContactData `xml:"data,omitempty"`
	Error  *ErrorInfo         `xml:"error,omitempty"`
}

type CreateContactParams struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Params"`

	Regid       string   `xml:"regid,omitempty"`
	Notifyemail string   `xml:"notify_email,omitempty"`
	Vat         string   `xml:"vat,omitempty"`
	Identtype   string   `xml:"ident_type,omitempty"`
	Identnumber string   `xml:"ident_number,omitempty"`
	Disclose    []string `xml:"disclose,omitempty"`
}

type CreateContactContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Contact"`

	Name    string               `xml:"name,omitempty"`
	Surname string               `xml:"surname,omitempty"`
	Org     string               `xml:"org,omitempty"`
	Street  string               `xml:"street,omitempty"`
	City    string               `xml:"city,omitempty"`
	Pc      string               `xml:"pc,omitempty"`
	Sp      string               `xml:"sp,omitempty"`
	Cc      string               `xml:"cc,omitempty"`
	Phone   string               `xml:"phone,omitempty"`
	Fax     string               `xml:"fax,omitempty"`
	Email   string               `xml:"email,omitempty"`
	Params  *CreateContactParams `xml:"params,omitempty"`
}

type CreateContactData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Create_Contact_Data"`

	Contactid string `xml:"contactid,omitempty"`
}

type UpdateContactResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Response"`

	Status string             `xml:"status,omitempty"`
	Data   *UpdateContactData `xml:"data,omitempty"`
	Error  *ErrorInfo         `xml:"error,omitempty"`
}

type UpdateContactContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Contact"`

	Id      string `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	Surname string `xml:"surname,omitempty"`
	Org     string `xml:"org,omitempty"`
	Street  string `xml:"street,omitempty"`
	City    string `xml:"city,omitempty"`
	Pc      string `xml:"pc,omitempty"`
	Sp      string `xml:"sp,omitempty"`
	Cc      string `xml:"cc,omitempty"`
	Phone   string `xml:"phone,omitempty"`
	Fax     string `xml:"fax,omitempty"`
	Email   string `xml:"email,omitempty"`
}

type UpdateContactOrder struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Order"`

	Register string `xml:"register,omitempty"`
	Orderid  int32  `xml:"orderid,omitempty"`
}

type UpdateContactData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Update_Contact_Data"`

	Orders []*UpdateContactOrder `xml:"orders,omitempty"`
}

type InfoContactResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Response"`

	Status string           `xml:"status,omitempty"`
	Data   *InfoContactData `xml:"data,omitempty"`
	Error  *ErrorInfo       `xml:"error,omitempty"`
}

type InfoContactContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Contact"`

	Id string `xml:"id,omitempty"`
}

type InfoContactData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Contact_Data"`

	Id      string `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	Surname string `xml:"surname,omitempty"`
	Org     string `xml:"org,omitempty"`
	Street  string `xml:"street,omitempty"`
	City    string `xml:"city,omitempty"`
	Pc      string `xml:"pc,omitempty"`
	Sp      string `xml:"sp,omitempty"`
	Cc      string `xml:"cc,omitempty"`
	Phone   string `xml:"phone,omitempty"`
	Fax     string `xml:"fax,omitempty"`
	Email   string `xml:"email,omitempty"`
}

type ContactsListResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Response"`

	Status string            `xml:"status,omitempty"`
	Data   *ContactsListData `xml:"data,omitempty"`
	Error  *ErrorInfo        `xml:"error,omitempty"`
}

type ContactsListContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Contact"`

	Name    string `xml:"name,omitempty"`
	Surname string `xml:"surname,omitempty"`
	Org     string `xml:"org,omitempty"`
	Street  string `xml:"street,omitempty"`
	City    string `xml:"city,omitempty"`
	Pc      string `xml:"pc,omitempty"`
	Sp      string `xml:"sp,omitempty"`
	Cc      string `xml:"cc,omitempty"`
	Email   string `xml:"email,omitempty"`
	Phone   string `xml:"phone,omitempty"`
	Fax     string `xml:"fax,omitempty"`
	Id      string `xml:"id,omitempty"`
}

type ContactsListData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Contacts_List_Data"`

	Contacts []*ContactsListContact `xml:"contacts,omitempty"`
	Count    int32                  `xml:"count,omitempty"`
}

type CheckObjectResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object_Response"`

	Status string           `xml:"status,omitempty"`
	Data   *CheckObjectData `xml:"data,omitempty"`
	Error  *ErrorInfo       `xml:"error,omitempty"`
}

type CheckObjectData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Check_Object_Data"`

	Id    string `xml:"id,omitempty"`
	Avail int32  `xml:"avail,omitempty"`
}

type InfoObjectResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Response"`

	Status string          `xml:"status,omitempty"`
	Data   *InfoObjectData `xml:"data,omitempty"`
	Error  *ErrorInfo      `xml:"error,omitempty"`
}

type InfoObjectContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Contact"`

	Name        string   `xml:"name,omitempty"`
	Org         string   `xml:"org,omitempty"`
	Street      string   `xml:"street,omitempty"`
	City        string   `xml:"city,omitempty"`
	Pc          string   `xml:"pc,omitempty"`
	Sp          string   `xml:"sp,omitempty"`
	Cc          string   `xml:"cc,omitempty"`
	Email       string   `xml:"email,omitempty"`
	Phone       string   `xml:"phone,omitempty"`
	Fax         string   `xml:"fax,omitempty"`
	Vat         string   `xml:"vat,omitempty"`
	Notifyemail string   `xml:"notify_email,omitempty"`
	Identtype   string   `xml:"ident_type,omitempty"`
	Identnumber string   `xml:"ident_number,omitempty"`
	ClID        string   `xml:"clID,omitempty"`
	Hidden      []string `xml:"hidden,omitempty"`
	Statuses    []string `xml:"statuses,omitempty"`
}

type InfoObjectNs struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Ns"`

	Host string `xml:"host,omitempty"`
	Ip   string `xml:"ip,omitempty"`
}

type InfoObjectNsset struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Nsset"`

	Tech string          `xml:"tech,omitempty"`
	Ns   []*InfoObjectNs `xml:"ns,omitempty"`
	ClID string          `xml:"clID,omitempty"`
}

type InfoObjectDnskey struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Dnskey"`

	Flags    string `xml:"flags,omitempty"`
	Protocol string `xml:"protocol,omitempty"`
	Alg      string `xml:"alg,omitempty"`
	PubKey   string `xml:"pubKey,omitempty"`
}

type InfoObjectKeyset struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Keyset"`

	Tech   string              `xml:"tech,omitempty"`
	Dnskey []*InfoObjectDnskey `xml:"dnskey,omitempty"`
	ClID   string              `xml:"clID,omitempty"`
}

type InfoObjectData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Object_Data"`

	Id      string             `xml:"id,omitempty"`
	Type_   string             `xml:"type,omitempty"`
	Contact *InfoObjectContact `xml:"contact,omitempty"`
	Nsset   *InfoObjectNsset   `xml:"nsset,omitempty"`
	Keyset  *InfoObjectKeyset  `xml:"keyset,omitempty"`
}

type MakeOrderResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Response"`

	Status string         `xml:"status,omitempty"`
	Data   *MakeOrderData `xml:"data,omitempty"`
	Error  *ErrorInfo     `xml:"error,omitempty"`
}

type MakeOrderContacts struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Contacts"`

	Admin   *MakeOrderContact `xml:"admin,omitempty"`
	Tech    *MakeOrderContact `xml:"tech,omitempty"`
	Billing *MakeOrderContact `xml:"billing,omitempty"`
}

type MakeOrderHost struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Host"`

	Hostname string `xml:"hostname,omitempty"`
	Ipv4     string `xml:"ipv4,omitempty"`
	Ipv6     string `xml:"ipv6,omitempty"`
}

type MakeOrderNs struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Ns"`

	Hosts []*MakeOrderHost `xml:"hosts,omitempty"`
	Nsset string           `xml:"nsset,omitempty"`
}

type MakeOrderNew struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_New"`

	Registrant *MakeOrderContact `xml:"registrant,omitempty"`
	Admin      *MakeOrderContact `xml:"admin,omitempty"`
	Tech       *MakeOrderContact `xml:"tech,omitempty"`
	Billing    *MakeOrderContact `xml:"billing,omitempty"`
	Ns         *MakeOrderNs      `xml:"ns,omitempty"`
}

type MakeOrderDsdata struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Dsdata"`

	Tag        string `xml:"tag,omitempty"`
	Alg        string `xml:"alg,omitempty"`
	Digesttype string `xml:"digest_type,omitempty"`
	Digest     string `xml:"digest,omitempty"`
}

type MakeOrderParam struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Param"`

	Dsdata []*MakeOrderDsdata `xml:"dsdata,omitempty"`
	Param  string             `xml:"param,omitempty"`
	Value  string             `xml:"value,omitempty"`
}

type MakeOrderParams struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Params"`

	Period     int32              `xml:"period,omitempty"`
	Registrant *MakeOrderContact  `xml:"registrant,omitempty"`
	Contacts   *MakeOrderContacts `xml:"contacts,omitempty"`
	Ns         *MakeOrderNs       `xml:"ns,omitempty"`
	New        *MakeOrderNew      `xml:"new,omitempty"`
	Type_      string             `xml:"type,omitempty"`
	Registry   string             `xml:"registry,omitempty"`
	Authid     string             `xml:"authid,omitempty"`
	Params     []*MakeOrderParam  `xml:"params,omitempty"`
	Newowner   string             `xml:"newowner,omitempty"`
	Reason     string             `xml:"reason,omitempty"`
	Nicd       string             `xml:"nicd,omitempty"`
	Password   string             `xml:"password,omitempty"`
	Hostname   string             `xml:"hostname,omitempty"`
	Ipv4       string             `xml:"ipv4,omitempty"`
	Ipv6       string             `xml:"ipv6,omitempty"`
	Dnstemp    string             `xml:"dnstemp,omitempty"`
	Statuses   []string           `xml:"statuses,omitempty"`
	Autorenew  int32              `xml:"autorenew,omitempty"`
}

type MakeOrderOrder struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Order"`

	Domain string           `xml:"domain,omitempty"`
	Object string           `xml:"object,omitempty"`
	Type_  string           `xml:"type,omitempty"`
	Params *MakeOrderParams `xml:"params,omitempty"`
}

type MakeOrderData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Data"`

	Orderid string `xml:"orderid,omitempty"`
}

type InfoOrderResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Response"`

	Status string         `xml:"status,omitempty"`
	Data   *InfoOrderData `xml:"data,omitempty"`
	Error  *ErrorInfo     `xml:"error,omitempty"`
}

type InfoOrderOrder struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Order"`

	Id         int32   `xml:"id,omitempty"`
	Domain     string  `xml:"domain,omitempty"`
	Type_      string  `xml:"type,omitempty"`
	Status     string  `xml:"status,omitempty"`
	Errorcode  string  `xml:"errorcode,omitempty"`
	Lastupdate string  `xml:"lastupdate,omitempty"`
	Message    string  `xml:"message,omitempty"`
	Payed      string  `xml:"payed,omitempty"`
	Amount     float64 `xml:"amount,omitempty"`
}

type InfoOrderData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Order_Data"`

	Order *InfoOrderOrder `xml:"order,omitempty"`
}

type GetCreditResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Response"`

	Status string         `xml:"status,omitempty"`
	Data   *GetCreditData `xml:"data,omitempty"`
	Error  *ErrorInfo     `xml:"error,omitempty"`
}

type GetCreditCredit struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Credit"`

	Amount    float64 `xml:"amount,omitempty"`
	Threshold float64 `xml:"threshold,omitempty"`
	Currency  string  `xml:"currency,omitempty"`
}

type GetCreditData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Credit_Data"`

	Credit *GetCreditCredit `xml:"credit,omitempty"`
}

type GetAccountingsResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Response"`

	Status string              `xml:"status,omitempty"`
	Data   *GetAccountingsData `xml:"data,omitempty"`
	Error  *ErrorInfo          `xml:"error,omitempty"`
}

type GetAccountingsAccounting struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Accounting"`

	Date   string  `xml:"date,omitempty"`
	Text   string  `xml:"text,omitempty"`
	Order  int32   `xml:"order,omitempty"`
	Sum    float64 `xml:"sum,omitempty"`
	Credit float64 `xml:"credit,omitempty"`
}

type GetAccountingsData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Accountings_Data"`

	Count      int32                       `xml:"count,omitempty"`
	From       string                      `xml:"from,omitempty"`
	To         string                      `xml:"to,omitempty"`
	Accounting []*GetAccountingsAccounting `xml:"accounting,omitempty"`
}

type ClientPaymentResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment_Response"`

	Status string             `xml:"status,omitempty"`
	Data   *ClientPaymentData `xml:"data,omitempty"`
	Error  *ErrorInfo         `xml:"error,omitempty"`
}

type ClientPaymentData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Client_Payment_Data"`
}

type CreditCorrectionResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction_Response"`

	Status string                `xml:"status,omitempty"`
	Data   *CreditCorrectionData `xml:"data,omitempty"`
	Error  *ErrorInfo            `xml:"error,omitempty"`
}

type CreditCorrectionData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Credit_Correction_Data"`
}

type PricelistResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Response"`

	Status string         `xml:"status,omitempty"`
	Data   *PricelistData `xml:"data,omitempty"`
	Error  *ErrorInfo     `xml:"error,omitempty"`
}

type PricelistPrice struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Price"`

	Type_ string  `xml:"type,omitempty"`
	Value float64 `xml:"value,omitempty"`
}

type PricelistValue struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Value"`

	Value       string `xml:"value,omitempty"`
	Description string `xml:"description,omitempty"`
}

type PricelistParam struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Param"`

	Param     string            `xml:"param,omitempty"`
	Desc      string            `xml:"desc,omitempty"`
	Required  int32             `xml:"required,omitempty"`
	Errorcode int32             `xml:"error_code,omitempty"`
	Values    []*PricelistValue `xml:"values,omitempty"`
}

type PricelistPricelist struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Pricelist"`

	Tld           string            `xml:"tld,omitempty"`
	Promo         int32             `xml:"promo,omitempty"`
	Promoexp      string            `xml:"promoexp,omitempty"`
	Country       string            `xml:"country,omitempty"`
	Continent     string            `xml:"continent,omitempty"`
	Minyear       int32             `xml:"minyear,omitempty"`
	Maxyear       int32             `xml:"maxyear,omitempty"`
	Minyearrenew  int32             `xml:"minyear_renew,omitempty"`
	Maxyearrenew  int32             `xml:"maxyear_renew,omitempty"`
	Localpresence int32             `xml:"local_presence,omitempty"`
	Prices        []*PricelistPrice `xml:"prices,omitempty"`
	Statuses      []string          `xml:"statuses,omitempty"`
	Params        []*PricelistParam `xml:"params,omitempty"`
}

type PricelistData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Pricelist_Data"`

	Pricelist []*PricelistPricelist `xml:"pricelist,omitempty"`
}

type PricesResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Response"`

	Status string      `xml:"status,omitempty"`
	Data   *PricesData `xml:"data,omitempty"`
	Error  *ErrorInfo  `xml:"error,omitempty"`
}

type PricesPrice struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Price"`

	Type_ string  `xml:"type,omitempty"`
	Value float64 `xml:"value,omitempty"`
}

type PricesValue struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Value"`

	Value       string `xml:"value,omitempty"`
	Description string `xml:"description,omitempty"`
}

type PricesParam struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Param"`

	Param     string         `xml:"param,omitempty"`
	Desc      string         `xml:"desc,omitempty"`
	Required  int32          `xml:"required,omitempty"`
	Errorcode int32          `xml:"error_code,omitempty"`
	Values    []*PricesValue `xml:"values,omitempty"`
}

type PricesData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Prices_Data"`

	Tld           string         `xml:"tld,omitempty"`
	Country       string         `xml:"country,omitempty"`
	Continent     string         `xml:"continent,omitempty"`
	Minyear       int32          `xml:"minyear,omitempty"`
	Maxyear       int32          `xml:"maxyear,omitempty"`
	Localpresence int32          `xml:"local_presence,omitempty"`
	Prices        []*PricesPrice `xml:"prices,omitempty"`
	Statuses      []string       `xml:"statuses,omitempty"`
	Params        []*PricesParam `xml:"params,omitempty"`
}

type GetPricelistResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Response"`

	Status string            `xml:"status,omitempty"`
	Data   *GetPricelistData `xml:"data,omitempty"`
	Error  *ErrorInfo        `xml:"error,omitempty"`
}

type GetPricelistPrice struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Price"`

	Type_ string  `xml:"type,omitempty"`
	Value float64 `xml:"value,omitempty"`
}

type GetPricelistPricelist struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Pricelist"`

	Tld      string               `xml:"tld,omitempty"`
	Currency string               `xml:"currency,omitempty"`
	Prices   []*GetPricelistPrice `xml:"prices,omitempty"`
}

type GetPricelistData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Pricelist_Data"`

	Pricelist []*GetPricelistPricelist `xml:"pricelist,omitempty"`
}

type SetPricesResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Response"`

	Status string         `xml:"status,omitempty"`
	Data   *SetPricesData `xml:"data,omitempty"`
	Error  *ErrorInfo     `xml:"error,omitempty"`
}

type SetPricesPrice struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Price"`

	Type_ string  `xml:"type,omitempty"`
	Value float64 `xml:"value,omitempty"`
}

type SetPricesData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_Prices_Data"`
}

type DownloadDocumentResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document_Response"`

	Status string                `xml:"status,omitempty"`
	Data   *DownloadDocumentData `xml:"data,omitempty"`
	Error  *ErrorInfo            `xml:"error,omitempty"`
}

type DownloadDocumentData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Download_Document_Data"`

	Id       int32  `xml:"id,omitempty"`
	Name     string `xml:"name,omitempty"`
	Type_    string `xml:"type,omitempty"`
	Filetype string `xml:"filetype,omitempty"`
	Account  string `xml:"account,omitempty"`
	Document string `xml:"document,omitempty"`
}

type UploadDocumentResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document_Response"`

	Status string              `xml:"status,omitempty"`
	Data   *UploadDocumentData `xml:"data,omitempty"`
	Error  *ErrorInfo          `xml:"error,omitempty"`
}

type UploadDocumentData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Upload_Document_Data"`

	Id int32 `xml:"id,omitempty"`
}

type ListDocumentsResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Response"`

	Status string             `xml:"status,omitempty"`
	Data   *ListDocumentsData `xml:"data,omitempty"`
	Error  *ErrorInfo         `xml:"error,omitempty"`
}

type ListDocumentsDocument struct {
	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Document"`

	Id       string `xml:"id,omitempty"`
	Name     string `xml:"name,omitempty"`
	Type_    string `xml:"type,omitempty"`
	Filetype string `xml:"filetype,omitempty"`
	Account  string `xml:"account,omitempty"`
	Orderid  int32  `xml:"orderid,omitempty"`
}

type ListDocumentsData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types List_Documents_Data"`

	Documents []*ListDocumentsDocument `xml:"documents,omitempty"`
}

type UsersListResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_Response"`

	Status string         `xml:"status,omitempty"`
	Data   *UsersListData `xml:"data,omitempty"`
	Error  *ErrorInfo     `xml:"error,omitempty"`
}

type UsersListUser struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_User"`

	Id             int32  `xml:"id,omitempty"`
	Username       string `xml:"username,omitempty"`
	Name           string `xml:"name,omitempty"`
	Credit         string `xml:"credit,omitempty"`
	Currency       string `xml:"currency,omitempty"`
	Billingname    string `xml:"billing_name,omitempty"`
	Billingstreet  string `xml:"billing_street,omitempty"`
	Billingcity    string `xml:"billing_city,omitempty"`
	Billingpc      string `xml:"billing_pc,omitempty"`
	Billingcountry string `xml:"billing_country,omitempty"`
	Companyid      string `xml:"company_id,omitempty"`
	Companyvat     string `xml:"company_vat,omitempty"`
	Email          string `xml:"email,omitempty"`
	Phone          string `xml:"phone,omitempty"`
	Lastlogin      string `xml:"last_login,omitempty"`
}

type UsersListData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Users_List_Data"`

	Count int32            `xml:"count,omitempty"`
	Users []*UsersListUser `xml:"users,omitempty"`
}

type AnycastADDZoneResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone_Response"`

	Status string              `xml:"status,omitempty"`
	Data   *AnycastADDZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo          `xml:"error,omitempty"`
}

type AnycastADDZoneData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_ADD_Zone_Data"`
}

type AnycastRemoveZoneResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone_Response"`

	Status string                 `xml:"status,omitempty"`
	Data   *AnycastRemoveZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo             `xml:"error,omitempty"`
}

type AnycastRemoveZoneData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Anycast_Remove_Zone_Data"`
}

type GetDNSZoneResponse struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Zone_Response"`

	Status string          `xml:"status,omitempty"`
	Data   *GetDNSZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo      `xml:"error,omitempty"`
}

type GetDNSZoneRecord struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Zone_Record"`

	Id      int32  `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	Type_   string `xml:"type,omitempty"`
	Content string `xml:"content,omitempty"`
	Prio    int32  `xml:"prio,omitempty"`
	Ttl     int32  `xml:"ttl,omitempty"`
}

type GetDNSZoneData struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Zone_Data"`

	Domain  string              `xml:"domain,omitempty"`
	Records []*GetDNSZoneRecord `xml:"records,omitempty"`
}

type AddDNSZoneResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone_Response"`

	Status string          `xml:"status,omitempty"`
	Data   *AddDNSZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo      `xml:"error,omitempty"`
}

type AddDNSZoneData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Zone_Data"`
}

type DeleteDNSZoneResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone_Response"`

	Status string             `xml:"status,omitempty"`
	Data   *DeleteDNSZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo         `xml:"error,omitempty"`
}

type DeleteDNSZoneData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Zone_Data"`
}

type SetDNSZoneResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Response"`

	Status string          `xml:"status,omitempty"`
	Data   *SetDNSZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo      `xml:"error,omitempty"`
}

type SetDNSZoneRecord struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Record"`

	Name    string `xml:"name,omitempty"`
	Type_   string `xml:"type,omitempty"`
	Content string `xml:"content,omitempty"`
	Prio    int32  `xml:"prio,omitempty"`
	Ttl     int32  `xml:"ttl,omitempty"`
}

type SetDNSZoneData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Set_DNS_Zone_Data"`
}

type AddDNSRecordResponse struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Record_Response"`

	Status string            `xml:"status,omitempty"`
	Data   *AddDNSRecordData `xml:"data,omitempty"`
	Error  *ErrorInfo        `xml:"error,omitempty"`
}

type AddDNSRecordRecord struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Record_Record"`

	Name    string `xml:"name,omitempty"`
	Type_   string `xml:"type,omitempty"`
	Content string `xml:"content,omitempty"`
	Prio    int32  `xml:"prio,omitempty"`
	Ttl     int32  `xml:"ttl,omitempty"`
}

type AddDNSRecordData struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Add_DNS_Record_Data"`
}

type ModifyDNSRecordResponse struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Modify_DNS_Record_Response"`

	Status string               `xml:"status,omitempty"`
	Data   *ModifyDNSRecordData `xml:"data,omitempty"`
	Error  *ErrorInfo           `xml:"error,omitempty"`
}

type ModifyDNSRecordRecord struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Modify_DNS_Record_Record"`

	Id      int32  `xml:"id,omitempty"`
	Type_   string `xml:"type,omitempty"`
	Content string `xml:"content,omitempty"`
	Prio    int32  `xml:"prio,omitempty"`
	Ttl     int32  `xml:"ttl,omitempty"`
}

type ModifyDNSRecordData struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Modify_DNS_Record_Data"`
}

type DeleteDNSRecordResponse struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Response"`

	Status string               `xml:"status,omitempty"`
	Data   *DeleteDNSRecordData `xml:"data,omitempty"`
	Error  *ErrorInfo           `xml:"error,omitempty"`
}

type DeleteDNSRecordRecord struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Record"`

	Id int32 `xml:"id,omitempty"`
}

type DeleteDNSRecordData struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Delete_DNS_Record_Data"`
}

type POLLGetResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get_Response"`

	Status string       `xml:"status,omitempty"`
	Data   *POLLGetData `xml:"data,omitempty"`
	Error  *ErrorInfo   `xml:"error,omitempty"`
}

type POLLGetData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Get_Data"`

	Count       int32  `xml:"count,omitempty"`
	Date        string `xml:"date,omitempty"`
	Id          int32  `xml:"id,omitempty"`
	Orderid     int32  `xml:"orderid,omitempty"`
	Orderstatus string `xml:"orderstatus,omitempty"`
	Message     string `xml:"message,omitempty"`
	Errorcode   string `xml:"errorcode,omitempty"`
}

type POLLAckResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack_Response"`

	Status string       `xml:"status,omitempty"`
	Data   *POLLAckData `xml:"data,omitempty"`
	Error  *ErrorInfo   `xml:"error,omitempty"`
}

type POLLAckData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types POLL_Ack_Data"`
}

type OIBSearchResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Response"`

	Status string         `xml:"status,omitempty"`
	Data   *OIBSearchData `xml:"data,omitempty"`
	Error  *ErrorInfo     `xml:"error,omitempty"`
}

type OIBSearchDomain struct {
	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Domain"`

	Name     string `xml:"name,omitempty"`
	Type_    int32  `xml:"type,omitempty"`
	Typedesc string `xml:"typedesc,omitempty"`
}

type OIBSearchType struct {
	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Type"`

	Type_    int32  `xml:"type,omitempty"`
	Typedesc string `xml:"typedesc,omitempty"`
	Used     int32  `xml:"used,omitempty"`
	Maximum  int32  `xml:"maximum,omitempty"`
}

type OIBSearchData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types OIB_Search_Data"`

	Domains []*OIBSearchDomain `xml:"domains,omitempty"`
	Types   []*OIBSearchType   `xml:"types,omitempty"`
}

type GetCertificateResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate_Response"`

	Status string              `xml:"status,omitempty"`
	Data   *GetCertificateData `xml:"data,omitempty"`
	Error  *ErrorInfo          `xml:"error,omitempty"`
}

type GetCertificateData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Certificate_Data"`

	Certificate string `xml:"certificate,omitempty"`
	Expire      string `xml:"expire,omitempty"`
	Domain      string `xml:"domain,omitempty"`
	Type_       string `xml:"type,omitempty"`
}

type GetRedirectsResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects_Response"`

	Status string            `xml:"status,omitempty"`
	Data   *GetRedirectsData `xml:"data,omitempty"`
	Error  *ErrorInfo        `xml:"error,omitempty"`
}

type GetRedirectsData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_Redirects_Data"`

	Web   string `xml:"web,omitempty"`
	Email string `xml:"email,omitempty"`
}

type InSubregResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg_Response"`

	Status string        `xml:"status,omitempty"`
	Data   *InSubregData `xml:"data,omitempty"`
	Error  *ErrorInfo    `xml:"error,omitempty"`
}

type InSubregData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types In_Subreg_Data"`

	Myaccount string `xml:"myaccount,omitempty"`
}

type SignDNSZoneResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone_Response"`

	Status string           `xml:"status,omitempty"`
	Data   *SignDNSZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo       `xml:"error,omitempty"`
}

type SignDNSZoneData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Sign_DNS_Zone_Data"`
}

type UnsignDNSZoneResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone_Response"`

	Status string             `xml:"status,omitempty"`
	Data   *UnsignDNSZoneData `xml:"data,omitempty"`
	Error  *ErrorInfo         `xml:"error,omitempty"`
}

type UnsignDNSZoneData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Unsign_DNS_Zone_Data"`
}

type GetDNSInfoResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Response"`

	Status string          `xml:"status,omitempty"`
	Data   *GetDNSInfoData `xml:"data,omitempty"`
	Error  *ErrorInfo      `xml:"error,omitempty"`
}

type GetDNSInfoAnydata struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Anydata"`

	Type_ string `xml:"type,omitempty"`
	Data  string `xml:"data,omitempty"`
}

type GetDNSInfoDn struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Dn"`

	Nameserver string               `xml:"nameserver,omitempty"`
	Anydata    []*GetDNSInfoAnydata `xml:"anydata,omitempty"`
	Nslist     []string             `xml:"nslist,omitempty"`
	Soaid      string               `xml:"soaid,omitempty"`
}

type GetDNSInfoData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_DNS_Info_Data"`

	Inzone string          `xml:"in_zone,omitempty"`
	Dnssec string          `xml:"dnssec,omitempty"`
	Dns    []*GetDNSInfoDn `xml:"dns,omitempty"`
}

type SpecialPricelistResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Response"`

	Status string                `xml:"status,omitempty"`
	Data   *SpecialPricelistData `xml:"data,omitempty"`
	Error  *ErrorInfo            `xml:"error,omitempty"`
}

type SpecialPricelistPrice struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Price"`

	Register float64 `xml:"register,omitempty"`
	Renew    float64 `xml:"renew,omitempty"`
	Transfer float64 `xml:"transfer,omitempty"`
}

type SpecialPricelistPricelist struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Pricelist"`

	Tld      string                   `xml:"tld,omitempty"`
	Currency string                   `xml:"currency,omitempty"`
	Dateto   string                   `xml:"dateto,omitempty"`
	Prices   []*SpecialPricelistPrice `xml:"prices,omitempty"`
}

type SpecialPricelistData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Special_Pricelist_Data"`

	Pricelist []*SpecialPricelistPricelist `xml:"pricelist,omitempty"`
}

type GetTLDInfoResponse struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Response"`

	Status string          `xml:"status,omitempty"`
	Data   *GetTLDInfoData `xml:"data,omitempty"`
	Error  *ErrorInfo      `xml:"error,omitempty"`
}

type GetTLDInfoContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Contact"`

	Type_ string `xml:"type,omitempty"`
	Cnt   int32  `xml:"cnt,omitempty"`
}

type GetTLDInfoOption struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Option"`

	Value string `xml:"value,omitempty"`
	Name  string `xml:"name,omitempty"`
}

type GetTLDInfoParam struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Param"`

	Param    string              `xml:"param,omitempty"`
	Name     string              `xml:"name,omitempty"`
	Desc     string              `xml:"desc,omitempty"`
	Required string              `xml:"required,omitempty"`
	Options  []*GetTLDInfoOption `xml:"options,omitempty"`
}

type GetTLDInfoData struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Get_TLD_Info_Data"`

	PeriodsCreate []string             `xml:"periodsCreate,omitempty"`
	PeriodsRenew  []string             `xml:"periodsRenew,omitempty"`
	Transfer      string               `xml:"transfer,omitempty"`
	Trade         string               `xml:"trade,omitempty"`
	Idn           string               `xml:"idn,omitempty"`
	Trustee       string               `xml:"trustee,omitempty"`
	Ns            string               `xml:"ns,omitempty"`
	Contacts      []*GetTLDInfoContact `xml:"contacts,omitempty"`
	Params        []*GetTLDInfoParam   `xml:"params,omitempty"`
}

type ErrorCodes struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Error_Codes"`

	Major int32 `xml:"major,omitempty"`
	Minor int32 `xml:"minor,omitempty"`
}

type ErrorInfo struct {
	// XMLName xml.Name `xml:"http://subreg.cz/types Error_Info"`

	Errormsg  string      `xml:"errormsg,omitempty"`
	Errorcode *ErrorCodes `xml:"errorcode,omitempty"`
}

type InfoDomainContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_Contact"`

	Subregid   string `xml:"subregid,omitempty"`
	Registryid string `xml:"registryid,omitempty"`
}

type InfoDomainCZContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Info_Domain_CZ_Contact"`

	Subregid   string `xml:"subregid,omitempty"`
	Registryid string `xml:"registryid,omitempty"`
}

type MakeOrderContactNew struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Contact_New"`

	Name    string `xml:"name,omitempty"`
	Surname string `xml:"surname,omitempty"`
	Org     string `xml:"org,omitempty"`
	Street  string `xml:"street,omitempty"`
	City    string `xml:"city,omitempty"`
	Pc      string `xml:"pc,omitempty"`
	Sp      string `xml:"sp,omitempty"`
	Cc      string `xml:"cc,omitempty"`
	Phone   string `xml:"phone,omitempty"`
	Fax     string `xml:"fax,omitempty"`
	Email   string `xml:"email,omitempty"`
}

type MakeOrderContact struct {
	XMLName xml.Name `xml:"http://subreg.cz/types Make_Order_Contact"`

	Id    string               `xml:"id,omitempty"`
	Regid string               `xml:"regid,omitempty"`
	New   *MakeOrderContactNew `xml:"new,omitempty"`
}

type SubregCz struct {
	client *SOAPClient
	Token  string
}

func NewSubregCz(url string, tls bool, auth *BasicAuth) *SubregCz {
	if url == "" {
		url = "https://subreg.cz/soap/cmd.php?soap_format=1"
	}
	client := NewSOAPClient(url, tls, auth)

	return &SubregCz{
		client: client,
	}
}

func (service *SubregCz) Login(request *Login) (*LoginContainer, error) {
	response := new(LoginContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Login", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) CheckDomain(request *CheckDomain) (*CheckDomainContainer, error) {
	response := new(CheckDomainContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Check_Domain", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) InfoDomain(request *InfoDomain) (*InfoDomainContainer, error) {
	response := new(InfoDomainContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Info_Domain", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) InfoDomainCZ(request *InfoDomainCZ) (*InfoDomainCZContainer, error) {
	response := new(InfoDomainCZContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Info_Domain_CZ", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) DomainsList(request *DomainsList) (*DomainsListContainer, error) {
	response := new(DomainsListContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Domains_List", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) SetAutorenew(request *SetAutorenew) (*SetAutorenewContainer, error) {
	response := new(SetAutorenewContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Set_Autorenew", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) CreateContact(request *CreateContact) (*CreateContactContainer, error) {
	response := new(CreateContactContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Create_Contact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) UpdateContact(request *UpdateContact) (*UpdateContactContainer, error) {
	response := new(UpdateContactContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Update_Contact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) InfoContact(request *InfoContact) (*InfoContactContainer, error) {
	response := new(InfoContactContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Info_Contact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) ContactsList(request *ContactsList) (*ContactsListContainer, error) {
	response := new(ContactsListContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Contacts_List", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) CheckObject(request *CheckObject) (*CheckObjectContainer, error) {
	response := new(CheckObjectContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Check_Object", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) InfoObject(request *InfoObject) (*InfoObjectContainer, error) {
	response := new(InfoObjectContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Info_Object", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) MakeOrder(request *MakeOrder) (*MakeOrderContainer, error) {
	response := new(MakeOrderContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Make_Order", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) InfoOrder(request *InfoOrder) (*InfoOrderContainer, error) {
	response := new(InfoOrderContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Info_Order", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetCredit(request *GetCredit) (*GetCreditContainer, error) {
	response := new(GetCreditContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_Credit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetAccountings(request *GetAccountings) (*GetAccountingsContainer, error) {
	response := new(GetAccountingsContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_Accountings", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) ClientPayment(request *ClientPayment) (*ClientPaymentContainer, error) {
	response := new(ClientPaymentContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Client_Payment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) CreditCorrection(request *CreditCorrection) (*CreditCorrectionContainer, error) {
	response := new(CreditCorrectionContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Credit_Correction", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) Pricelist(request *Pricelist) (*PricelistContainer, error) {
	response := new(PricelistContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Pricelist", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) Prices(request *Prices) (*PricesContainer, error) {
	response := new(PricesContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Prices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetPricelist(request *GetPricelist) (*GetPricelistContainer, error) {
	response := new(GetPricelistContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_Pricelist", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) SetPrices(request *SetPrices) (*SetPricesContainer, error) {
	response := new(SetPricesContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Set_Prices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) DownloadDocument(request *DownloadDocument) (*DownloadDocumentContainer, error) {
	response := new(DownloadDocumentContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Download_Document", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) UploadDocument(request *UploadDocument) (*UploadDocumentContainer, error) {
	response := new(UploadDocumentContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Upload_Document", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) ListDocuments(request *ListDocuments) (*ListDocumentsContainer, error) {
	response := new(ListDocumentsContainer)
	err := service.client.Call("http://subreg.cz/wsdl#List_Documents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) UsersList(request *UsersList) (*UsersListContainer, error) {
	response := new(UsersListContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Users_List", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) AnycastADDZone(request *AnycastADDZone) (*AnycastADDZoneContainer, error) {
	response := new(AnycastADDZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Anycast_ADD_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) AnycastRemoveZone(request *AnycastRemoveZone) (*AnycastRemoveZoneContainer, error) {
	response := new(AnycastRemoveZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Anycast_Remove_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetDNSZone(request *GetDNSZone) (*GetDNSZoneContainer, error) {
	response := new(GetDNSZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_DNS_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) AddDNSZone(request *AddDNSZone) (*AddDNSZoneContainer, error) {
	response := new(AddDNSZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Add_DNS_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) DeleteDNSZone(request *DeleteDNSZone) (*DeleteDNSZoneContainer, error) {
	response := new(DeleteDNSZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Delete_DNS_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) SetDNSZone(request *SetDNSZone) (*SetDNSZoneContainer, error) {
	response := new(SetDNSZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Set_DNS_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) AddDNSRecord(request *AddDNSRecord) (*AddDNSRecordContainer, error) {
	response := new(AddDNSRecordContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Add_DNS_Record", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) ModifyDNSRecord(request *ModifyDNSRecord) (*ModifyDNSRecordContainer, error) {
	response := new(ModifyDNSRecordContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Modify_DNS_Record", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) DeleteDNSRecord(request *DeleteDNSRecord) (*DeleteDNSRecordContainer, error) {
	response := new(DeleteDNSRecordContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Delete_DNS_Record", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) POLLGet(request *POLLGet) (*POLLGetContainer, error) {
	response := new(POLLGetContainer)
	err := service.client.Call("http://subreg.cz/wsdl#POLL_Get", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) POLLAck(request *POLLAck) (*POLLAckContainer, error) {
	response := new(POLLAckContainer)
	err := service.client.Call("http://subreg.cz/wsdl#POLL_Ack", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) OIBSearch(request *OIBSearch) (*OIBSearchContainer, error) {
	response := new(OIBSearchContainer)
	err := service.client.Call("http://subreg.cz/wsdl#OIB_Search", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetCertificate(request *GetCertificate) (*GetCertificateContainer, error) {
	response := new(GetCertificateContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_Certificate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetRedirects(request *GetRedirects) (*GetRedirectsContainer, error) {
	response := new(GetRedirectsContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_Redirects", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) InSubreg(request *InSubreg) (*InSubregContainer, error) {
	response := new(InSubregContainer)
	err := service.client.Call("http://subreg.cz/wsdl#In_Subreg", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) SignDNSZone(request *SignDNSZone) (*SignDNSZoneContainer, error) {
	response := new(SignDNSZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Sign_DNS_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) UnsignDNSZone(request *UnsignDNSZone) (*UnsignDNSZoneContainer, error) {
	response := new(UnsignDNSZoneContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Unsign_DNS_Zone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetDNSInfo(request *GetDNSInfo) (*GetDNSInfoContainer, error) {
	response := new(GetDNSInfoContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_DNS_Info", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) SpecialPricelist(request *SpecialPricelist) (*SpecialPricelistContainer, error) {
	response := new(SpecialPricelistContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Special_Pricelist", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *SubregCz) GetTLDInfo(request *GetTLDInfo) (*GetTLDInfoContainer, error) {
	response := new(GetTLDInfoContainer)
	err := service.client.Call("http://subreg.cz/wsdl#Get_TLD_Info", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url   string
	tls   bool
	auth  *BasicAuth
	token string
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}
func (s *SOAPClient) getToken() (string, error) {
	if s.token != "" {
		return s.token, nil
	}

	request := &Login{Login: s.auth.Login, Password: s.auth.Password}
	response := new(LoginContainer)
	err := s.Call("http://subreg.cz/wsdl#Login", request, response)
	if err != nil {
		return "", err
	}
	if response != nil && response.Response != nil && response.Response.Status == "ok" && response.Response.Data != nil {
		token := response.Response.Data.Ssid
		fmt.Printf("%v\n", token)
		s.token = token
		return s.token, nil

	}

	return "", fmt.Errorf("Cannot get token")

}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{
	//Header:        SoapHeader{},
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	// inject token to request
	reflectSsid := reflect.ValueOf(request).Elem().FieldByName("Ssid")
	if reflectSsid.IsValid() == true {
		token, err := s.getToken()
		if err != nil {
			return err
		}
		reflectSsid.SetString(token)
	}
	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	// if s.auth != nil {
	// 	req.SetBasicAuth(s.auth.Login, s.auth.Password)
	// }

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	reflectResponse := reflect.ValueOf(response).Elem().FieldByName("Response")
	if !reflectResponse.IsValid() {
		return fmt.Errorf("Missing response tag")
	}

	reflectStatus := reflectResponse.Elem().FieldByName("Status")
	if !reflectStatus.IsValid() {
		return fmt.Errorf("Missing status tag")
	}

	status := reflectStatus.Interface()

	switch status {
	case "ok":
		return nil
	case "error":
		{
			errorInfo := reflectResponse.Elem().FieldByName("Error").Interface().(*ErrorInfo)
			return fmt.Errorf("Subreg api error: %s [%d:%d]", errorInfo.Errormsg, errorInfo.Errorcode.Major, errorInfo.Errorcode.Minor)
		}
	default:
		return fmt.Errorf("Uknown response status %s", status)
	}

	return nil
}
