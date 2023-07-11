package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/fatih/structs"
	"github.com/google/uuid"
)

const (
	checkoutURL = "https://pay.fondy.eu/api/checkout/url/"

	merchantPassword = "test"
	currencyUSD      = "USD"
	merchantId       = "1396424"
)

type APIRequest struct {
	Request interface{} `json:"request"`
}

type APIResponse struct {
	Response interface{} `json:"response"`
}

type CheckoutRequest struct {
	OrderId           string `json:"order_id"`
	MerchantId        string `json:"merchant_id"`
	OrderDesc         string `json:"order_desc"`
	Signature         string `json:"signature"`
	Amout             string `json:"amout"`
	Currency          string `json:"currency"`
	ResponseURL       string `json:"response_url,omitempty"`
	ServerCallbackURL string `json:"sender_callback_url,omitempty"`
	SenderEmail       string `json:"sender_email,omitempty"`
	Language          string `json:"lang,omitempty"`
	ProductId         string `json:"product_id,omitempty"`
}

type InterimResponse struct {
	Status      string `json:"response_status"`
	CheckoutURL string `json:"checkout_url"`
	PaymentId   string `json:"paymant_id"`
}

type CallbackbackResponse struct {
	OrderId                 string      `json:"order_id"`
	MerchantId              int         `json:"merchant_id"`
	Amout                   string      `json:"amout"`
	Currency                string      `json:"currency"`
	OrderStatus             string      `json:"order_status"`    // created; processing; declined; approved; expired; reversed;
	ResponseStatus          string      `json:"response_status"` // 1) success; 2) failure
	Signature               string      `json:"signature"`
	TranType                string      `json:"tran_type"`
	SenderCellPhone         string      `json:"sender_cell_phone"`
	SenderAccount           string      `json:"sender_account"`
	CardBin                 int         `json:"card_bin"`
	MaskedCard              string      `json:"masked_card"`
	CardType                string      `json:"card_type"`
	RRN                     string      `json:"rrn"`
	ApprovalCode            string      `json:"approval_code"`
	ResponseCode            interface{} `json:"response_code"`
	ResponseDescription     string      `json:"response_description"`
	ReversalAmount          string      `json:"reversal_amount"`
	SettlementAmount        string      `json:"settlement_amount"`
	SettlementCurrency      string      `json:"settlement_currency"`
	OrderTime               string      `json:"order_time"`
	SettlementDate          string      `json:"settlement_date"`
	ECI                     string      `json:"eci"`
	Fee                     string      `json:"fee"`
	PaymentSystem           string      `json:"payment_system"`
	SenderEmail             string      `json:"sender_email"`
	PaymentId               int         `json:"payment_id"`
	ActualAmount            string      `json:"actual_amount"`
	ActualCurrency          string      `json:"actual_currency"`
	MerchantData            string      `json:"merchant_data"`
	VerificationStatus      string      `json:"verification_status"`
	Rectoken                string      `json:"rectoken"`
	RectokenLifetime        string      `json:"rectoken_lifetime"`
	ProductId               string      `json:"product_id"`
	AdditionalInfo          string      `json:"additional_info"`
	ResponseSignatureString string      `json:"response_signature_string"`
}

func (r *CheckoutRequest) SetSignature(password string) {
	params := structs.Map(r)
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	values := []string{}

	for _, key := range keys {
		value := params[key].(string)
		if value == "" {
			continue
		}
		values = append(values, value)

	}

	r.Signature = generateSignature(values, password)
}

func generateSignature(values []string, password string) string {
	newValues := []string{password}
	newValues = append(newValues, values...)

	signatureString := strings.Join(newValues, "|")

	fmt.Println(signatureString)

	hash := sha1.New()
	hash.Write([]byte(signatureString))

	return fmt.Sprintf("%x", hash.Sum(nil))

}

func main() {
	id := uuid.New()
	checkoutReq := &CheckoutRequest{
		OrderId:           id.String(),
		MerchantId:        merchantId,
		OrderDesc:         "Lekcie",
		Amout:             "7000",
		Currency:          currencyUSD,
		ServerCallbackURL: "https:/",
	}

	checkoutReq.SetSignature(merchantPassword)

	request := APIRequest{Request: checkoutReq}
	requestBody, _ := json.Marshal(request)

	resp, err := http.Post(checkoutURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	apiResp := APIResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		panic(err)
	}

	fmt.Println(apiResp)
}
