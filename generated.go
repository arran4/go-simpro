package simpro

import (
	"fmt"
	"strings"
	"strconv"
	"net/http"
	"github.com/arran4/go-rfc5849-hmac"
	"io/ioutil"
	"encoding/json"
	"errors"
)

func AccPayContactFlag(CompanyID int, AccPayContactID interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactFlagList(CompanyID int, AccPayContactIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayContactIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactRetrieve(CompanyID int, AccPayContactID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactRetrieveAll(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactRetrieveAll", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactRetrieveExport(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactRetrieveList(CompanyID int, AccPayContactIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayContactIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactSearch(CompanyID int, Search string, Limit int, Offset int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TotalCount(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TotalCount", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactUnFlag(CompanyID int, AccPayContactID interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactUnFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayContactUnFlagList(CompanyID int, AccPayContactIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayContactUnFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayContactIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceFlag(CompanyID int, AccPayInvoiceID interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayInvoiceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceFlagList(CompanyID int, AccPayInvoiceIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayInvoiceIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceRetrieve(CompanyID int, AccPayInvoiceID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayInvoiceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceRetrieveExport(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceRetrieveList(CompanyID int, AccPayInvoiceIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayInvoiceIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceSearch(CompanyID int, Search string, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceSearchInvoice(CompanyID int, InvoiceNo string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceSearchInvoice", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceUnFlag(CompanyID int, AccPayInvoiceID interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceUnFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayInvoiceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccPayInvoiceUnFlagList(CompanyID int, AccPayInvoiceIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccPayInvoiceUnFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccPayInvoiceIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactFlag(CompanyID int, AccRecContactID interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactFlagList(CompanyID int, AccRecContactIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecContactIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactRetrieve(CompanyID int, AccRecContactID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactRetrieveAll(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactRetrieveAll", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactRetrieveExport(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactRetrieveList(CompanyID int, AccRecContactIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecContactIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactSearch(CompanyID int, Search string, Limit int, Offset int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactUnFlag(CompanyID int, AccRecContactID interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactUnFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecContactUnFlagList(CompanyID int, AccRecContactIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecContactUnFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecContactIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceFlag(CompanyID int, AccRecInvoiceID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecInvoiceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceFlagList(CompanyID int, AccRecInvoiceIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecInvoiceIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceRetrieve(CompanyID int, AccRecInvoiceID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecInvoiceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceRetrieveExport(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceRetrieveList(CompanyID int, AccRecInvoiceIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecInvoiceIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceSearch(CompanyID int, Search string, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceUnFlag(CompanyID int, AccRecInvoiceID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceUnFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecInvoiceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecInvoiceUnFlagList(CompanyID int, AccRecInvoiceIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecInvoiceUnFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecInvoiceIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentFlag(CompanyID int, AccRecPaymentID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecPaymentID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentFlagList(CompanyID int, AccRecPaymentIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecPaymentIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentRetrieve(CompanyID int, AccRecPaymentID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecPaymentID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentRetrieveExport(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentRetrieveList(CompanyID int, AccRecPaymentIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecPaymentIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentSearch(CompanyID int, Search string, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentUnFlag(CompanyID int, AccRecPaymentID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentUnFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecPaymentID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AccRecPaymentUnFlagList(CompanyID int, AccRecPaymentIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AccRecPaymentUnFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AccRecPaymentIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetInsert(CompanyID int, AssetDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AssetDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetRetrieve(CompanyID int, AssetID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AssetID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetRetrieveList(CompanyID int, AssetIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AssetIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetSearchExact(CompanyID int, Search string, FieldName string, AssetTypeID int, SiteID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetSearchExact", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(FieldName)), strconv.Quote(fmt.Sprint(AssetTypeID)), strconv.Quote(fmt.Sprint(SiteID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetServiceLevelRetrieve(CompanyID int, AssetTypeID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetServiceLevelRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AssetTypeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetTestRecordInsert(CompanyID int, TestRecord []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetTestRecordInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TestRecord))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetTypeRetrieve(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetTypeRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func AssetUpdate(CompanyID int, AssetID int, AssetDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "AssetUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(AssetID)), strconv.Quote(fmt.Sprint(AssetDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func BusinessGroupRetrieve(CompanyID int, BusinessGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "BusinessGroupRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(BusinessGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func BusinessGroupRetrieveList(CompanyID int, BusinessGroupIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "BusinessGroupRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(BusinessGroupIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func BusinessGroupSearch(CompanyID int, Name string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "BusinessGroupSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Name))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CompanyRetrieve(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CompanyRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CompanyRetrieveGeneralSettings(CompanyIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CompanyRetrieveGeneralSettings", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CompanyRetrieveList(CompanyIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CompanyRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CompanyRetrieveMobilePages(CompanyIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CompanyRetrieveMobilePages", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CompanyRetrieveMobileSettings(CompanyIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CompanyRetrieveMobileSettings", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CompanyRetrieveMobileStatuses(CompanyIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CompanyRetrieveMobileStatuses", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CompanySearch(CompanyName string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CompanySearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CostCentreRetrieve(CompanyID int, CostCentreID int, ReturnRemoved bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CostCentreRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CostCentreID)), strconv.Quote(fmt.Sprint(ReturnRemoved))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CostCentreRetrieveList(CompanyID int, CostCentreIDs []interface{}, ReturnRemoved bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CostCentreRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CostCentreIDs)), strconv.Quote(fmt.Sprint(ReturnRemoved))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CostCentreSearch(CompanyID int, Name string, ReturnRemoved bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CostCentreSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Name)), strconv.Quote(fmt.Sprint(ReturnRemoved))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LabourRateRetrieve(CompanyID int, LabourRateID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LabourRateRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LabourRateID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LabourRateRetrieveList(CompanyID int, LabourRateIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LabourRateRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LabourRateIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LabourRateSearch(CompanyID int, Name string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LabourRateSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Name))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ResponseTimeRetrieve(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ResponseTimeRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaxCodeRetrieve(TaxCodeID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaxCodeRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(TaxCodeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaxCodeRetrieveList(TaxCodeIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaxCodeRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(TaxCodeIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaxCodeSearch(Name string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaxCodeSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(Name))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerGroupRetrieve(CompanyID int, CustomerGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerGroupRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerGroupRetrieveList(CompanyID int, CustomerGroupIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerGroupRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerGroupIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerGroupSearch(CompanyID int, Name string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerGroupSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Name))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerContactInsert(CompanyID int, CustomerID int, CustomerContactDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerContactInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(CustomerContactDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerContactRetrieve(CompanyID int, CustomerID int, ContactID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerContactRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(ContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerContactRetrieveList(CompanyID int, CustomerIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerContactRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerContractRetrieve(CompanyID int, CustomerID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerContractRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerContractRetrieveList(CompanyID int, CustomerIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerContractRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInsert(CompanyID int, CustomerDetails []interface{}, CreateSite bool) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerDetails)), strconv.Quote(fmt.Sprint(CreateSite))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerNotesRetrieve(CompanyID int, CustomerID interface{}, SubjectPrefix string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerNotesRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(SubjectPrefix))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerRetrieve(CompanyID int, CustomerID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerRetrieveList(CompanyID int, CustomerIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerSearch(CompanyID int, Search string, Limit int, Offset int, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerSearchFields(CompanyID int, SearchTerms []interface{}, Limit int, Offset int, SearchType string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchTerms)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SearchType))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerSiteInsert(CompanyID int, CustomerID int, SiteDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerSiteInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(SiteDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerSiteLink(CompanyID int, CustomerID int, SiteID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerSiteLink", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(SiteID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerSiteRetrieve(CompanyID int, CustomerID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerSiteRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerSiteRetrieveList(CompanyID int, CustomerIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerSiteRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerSiteUnLink(CompanyID int, CustomerID int, SiteID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerSiteUnLink", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(SiteID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerTotalCount(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerTotalCount", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerUpdate(CompanyID int, CustomerID int, CustomerDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(CustomerDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerValidateCredentials(username string, password string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerValidateCredentials", strings.Join([]string{strconv.Quote(fmt.Sprint(username)), strconv.Quote(fmt.Sprint(password))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerIdentifierInsert(CompanyID int, TagName string) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerIdentifierInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TagName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerIdentifierRetrieve(CompanyID int, CustomerIdentifierID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerIdentifierRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerIdentifierID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerIdentifierRetrieveList(CompanyID int, CustomerIdentifierIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerIdentifierRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CustomerIdentifierIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerIdentifierSearch(CompanyID int, Name string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerIdentifierSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Name))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerIdentifierUpdate(CompanyID int, TagID int, TagName string) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerIdentifierUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TagID)), strconv.Quote(fmt.Sprint(TagName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalFlag(CompanyID int, InventoryJournalID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InventoryJournalID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalFlagList(CompanyID int, InventoryJournalIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InventoryJournalIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalRetrieve(CompanyID int, InventoryJournalID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InventoryJournalID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalRetrieveExport(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalRetrieveList(CompanyID int, InventoryJournalIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InventoryJournalIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalSearch(CompanyID int, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalUnFlag(CompanyID int, InventoryJournalID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalUnFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InventoryJournalID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func InventoryJournalUnFlagList(CompanyID int, InventoryJournalIDs []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "InventoryJournalUnFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InventoryJournalIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockListAll(CompanyID int, StorageDeviceID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockListAll", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StorageDeviceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockListAllDetail(CompanyID int, StorageDeviceID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockListAllDetail", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StorageDeviceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockRetrieve(CompanyID int, StockID int, StorageDeviceID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StockID)), strconv.Quote(fmt.Sprint(StorageDeviceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockRetrieveList(CompanyID int, StorageStock []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StorageStock))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockRetrieveLowStockLevels(CompanyID int, StorageID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockRetrieveLowStockLevels", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StorageID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockSearch(CompanyID int, StockID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StockID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockTakeHandler(CompanyID int, StockTakeID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockTakeHandler", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StockTakeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockTakeInsertComplete(CompanyID int, Complete bool, StockTakeDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockTakeInsertComplete", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Complete)), strconv.Quote(fmt.Sprint(StockTakeDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StockTakeInsertIncomplete(CompanyID int, StockTakeDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StockTakeInsertIncomplete", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StockTakeDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StorageDeviceInsert(CompanyID int, StorageDeviceDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StorageDeviceInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StorageDeviceDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StorageDeviceRetrieve(CompanyID int, StorageDeviceID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StorageDeviceRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StorageDeviceID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StorageDeviceRetrieveList(CompanyID int, StorageDeviceIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StorageDeviceRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StorageDeviceIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StorageDeviceSearch(CompanyID int, Name string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StorageDeviceSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Name))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorInvoiceRetrieve(CompanyID int, InvoiceNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorInvoiceRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorInvoiceRetrieveList(CompanyID int, InvoiceNos []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorInvoiceRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNos))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorInvoiceSearch(CompanyID int, Search string, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorInvoiceSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorInvoiceSearchContractorJob(CompanyID int, ContractorJobNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorInvoiceSearchContractorJob", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorJobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorInvoiceSearchInvoice(CompanyID int, InvoiceNo string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorInvoiceSearchInvoice", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceInsert(CompanyID int, InvoiceDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceInsertCredit(CompanyID int, InvoiceNo int, CreditDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceInsertCredit", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo)), strconv.Quote(fmt.Sprint(CreditDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LogSearch(CompanyID int, InvoiceNo int, StartDate string, EndDate string, StaffID int, Search string, Limit int, Offset int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LogSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceRetrieve(CompanyID int, InvoiceNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceRetrieveAll(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceRetrieveAll", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceRetrieveCredit(CompanyID int, CreditID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceRetrieveCredit", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CreditID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceRetrieveList(CompanyID int, InvoiceNos []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNos))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceRetrieveOutstanding(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceRetrieveOutstanding", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceRetrievePaid(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceRetrievePaid", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceSearch(CompanyID int, Search string, StartDate string, EndDate string, Stage string, Status string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Stage)), strconv.Quote(fmt.Sprint(Status))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceSearchJob(CompanyID int, JobNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceSearchJob", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceUpdate(CompanyID int, InvoiceNo int, InvoiceDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo)), strconv.Quote(fmt.Sprint(InvoiceDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerInvoiceUpdateCredit(CompanyID int, CreditID int, CreditDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerInvoiceUpdateCredit", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CreditID)), strconv.Quote(fmt.Sprint(CreditDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerPaymentInsert(CompanyID int, PaymentDetails []interface{}, Invoices []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerPaymentInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PaymentDetails)), strconv.Quote(fmt.Sprint(Invoices))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerPaymentRetrieve(CompanyID int, PaymentNo interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerPaymentRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PaymentNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerPaymentRetrieveAll(CompanyID int, PaymentIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerPaymentRetrieveAll", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PaymentIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerPaymentRetrieveList(CompanyID int, PaymentNos []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerPaymentRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PaymentNos))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerPaymentSearch(CompanyID int, Search string, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerPaymentSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CustomerPaymentSearchInvoice(CompanyID int, InvoiceNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CustomerPaymentSearchInvoice", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierCreditRetrieve(CompanyID int, CreditNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierCreditRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CreditNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierCreditRetrieveList(CompanyID int, CreditNos []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierCreditRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CreditNos))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierCreditSearch(CompanyID int, Search string, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierCreditSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierCreditSearchInvoice(CompanyID int, InvoiceNo string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierCreditSearchInvoice", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierInvoiceInsert(CompanyID int, OrderID int, SupplierInvoiceDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierInvoiceInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(OrderID)), strconv.Quote(fmt.Sprint(SupplierInvoiceDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierInvoiceRetrieve(CompanyID int, ReceiptNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierInvoiceRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ReceiptNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierInvoiceRetrieveList(CompanyID int, ReceiptNos []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierInvoiceRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ReceiptNos))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierInvoiceSearch(CompanyID int, Search string, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierInvoiceSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierInvoiceSearchInvoice(CompanyID int, InvoiceNo string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierInvoiceSearchInvoice", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(InvoiceNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogGroupInsert(CompanyID int, GroupName string) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogGroupInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(GroupName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogGroupRetrieve(CompanyID int, CatalogGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogGroupRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogGroupRetrieveList(CompanyID int, CatalogGroupIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogGroupRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogGroupIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogGroupSearch(CompanyID int, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogGroupSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogGroupUpdate(CompanyID int, CatalogGroupID int, GroupName string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogGroupUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogGroupID)), strconv.Quote(fmt.Sprint(GroupName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemAddSupplier(CompanyID int, CatalogItemID int, SupplierID int, NettPrice float64, Default bool) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemAddSupplier", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemID)), strconv.Quote(fmt.Sprint(SupplierID)), strconv.Quote(fmt.Sprint(NettPrice)), strconv.Quote(fmt.Sprint(Default))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemDeleteLink(CompanyID int, CatalogItemID int, LinkCatalogItemID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemDeleteLink", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemID)), strconv.Quote(fmt.Sprint(LinkCatalogItemID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemDeleteSupplier(CompanyID int, CatalogItemID int, SupplierID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemDeleteSupplier", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemID)), strconv.Quote(fmt.Sprint(SupplierID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemInsert(CompanyID int, CatalogItemDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemLink(CompanyID int, CatalogItemID int, LinkCatalogItemID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemLink", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemID)), strconv.Quote(fmt.Sprint(LinkCatalogItemID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemRetrieve(CompanyID int, CatalogItemID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemRetrieveList(CompanyID int, CatalogItemIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemSearch(CompanyID int, SearchText string, CatalogGroupID int, CatalogSubGroupID int, Limit int, Offset int, Favorite bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchText)), strconv.Quote(fmt.Sprint(CatalogGroupID)), strconv.Quote(fmt.Sprint(CatalogSubGroupID)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(Favorite))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemSearchExact(CompanyID int, Search string, CatalogGroupID int, CatalogSubGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemSearchExact", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(CatalogGroupID)), strconv.Quote(fmt.Sprint(CatalogSubGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogItemUpdate(CompanyID int, CatalogItemID int, CatalogItemDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogItemUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogItemID)), strconv.Quote(fmt.Sprint(CatalogItemDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogSubGroupInsert(CompanyID int, CatalogGroupID int, SubGroupName string) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogSubGroupInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogGroupID)), strconv.Quote(fmt.Sprint(SubGroupName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogSubGroupRetrieve(CompanyID int, CatalogSubGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogSubGroupRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogSubGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogSubGroupRetrieveList(CompanyID int, CatalogSubGroupIDs []interface{}, CatalogGroupIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogSubGroupRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogSubGroupIDs)), strconv.Quote(fmt.Sprint(CatalogGroupIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogSubGroupSearch(CompanyID int, Search string, CatalogGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogSubGroupSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(CatalogGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func CatalogSubGroupUpdate(CompanyID int, CatalogSubGroupID int, SubGroupName string, CatalogGroupID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "CatalogSubGroupUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(CatalogSubGroupID)), strconv.Quote(fmt.Sprint(SubGroupName)), strconv.Quote(fmt.Sprint(CatalogGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildGroupInsert(CompanyID int, GroupName string) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildGroupInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(GroupName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildGroupRetrieve(CompanyID int, PrebuildGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildGroupRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildGroupRetrieveList(CompanyID int, PrebuildGroupIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildGroupRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildGroupIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildGroupSearch(CompanyID int, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildGroupSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildGroupUpdate(CompanyID int, PrebuildGroupID int, GroupName string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildGroupUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildGroupID)), strconv.Quote(fmt.Sprint(GroupName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildInsert(CompanyID int, PrebuildDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildRetrieve(CompanyID int, PrebuildID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildRetrieveList(CompanyID int, PrebuildIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildSearch(CompanyID int, SearchText string, PrebuildGroupID int, PrebuildSubGroupID int, Limit int, Offset int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchText)), strconv.Quote(fmt.Sprint(PrebuildGroupID)), strconv.Quote(fmt.Sprint(PrebuildSubGroupID)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildSearchExact(CompanyID int, Search string, PrebuildGroupID int, PrebuildSubGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildSearchExact", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(PrebuildGroupID)), strconv.Quote(fmt.Sprint(PrebuildSubGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildSubGroupInsert(CompanyID int, PrebuildGroupID int, SubGroupName string) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildSubGroupInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildGroupID)), strconv.Quote(fmt.Sprint(SubGroupName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildSubGroupRetrieve(CompanyID int, PrebuildSubGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildSubGroupRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildSubGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildSubGroupRetrieveList(CompanyID int, PrebuildSubGroupIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildSubGroupRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildSubGroupIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildSubGroupSearch(CompanyID int, Search string, PrebuildGroupID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildSubGroupSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(PrebuildGroupID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PrebuildSubGroupUpdate(CompanyID int, PrebuildSubGroupID int, SubGroupName string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PrebuildSubGroupUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PrebuildSubGroupID)), strconv.Quote(fmt.Sprint(SubGroupName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func UOMInsert(CompanyID int, UOMDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "UOMInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(UOMDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func UOMRetrieve(CompanyID interface{}, UOMID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "UOMRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(UOMID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func UOMRetrieveList(CompanyID interface{}, UOMIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "UOMRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(UOMIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func UOMSearch(CompanyID interface{}, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "UOMSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func UOMSearchFields(CompanyID interface{}, SearchTerms []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "UOMSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchTerms))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func UOMUpdate(CompanyID int, UOMID int, UOMDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "UOMUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(UOMID)), strconv.Quote(fmt.Sprint(UOMDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorJobInsert(CompanyID int, ContractorJobDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorJobInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorJobDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorJobRetrieve(CompanyID int, ContractorJobNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorJobRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorJobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorJobRetrieveList(CompanyID int, ContractorJobNos []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorJobRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorJobNos))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorJobSearch(CompanyID int, Search string, StartDate string, EndDate string, JobID int, QuoteID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorJobSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(JobID)), strconv.Quote(fmt.Sprint(QuoteID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorJobUpdate(CompanyID int, ContractorJobID int, ContractorJobDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorJobUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorJobID)), strconv.Quote(fmt.Sprint(ContractorJobDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobDeleteAttachment(CompanyID int, JobNo int, FileName string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobDeleteAttachment", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo)), strconv.Quote(fmt.Sprint(FileName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobDeleteScheduleBlock(CompanyID int, JobNo int, DepartmentNo int, StaffID int, Date string, StartTime string, EndTime string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobDeleteScheduleBlock", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo)), strconv.Quote(fmt.Sprint(DepartmentNo)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(Date)), strconv.Quote(fmt.Sprint(StartTime)), strconv.Quote(fmt.Sprint(EndTime))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobFlag(CompanyID int, JobNo int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobFlagList(CompanyID int, JobNos []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNos))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobInsert(CompanyID int, JobDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobInsertScheduleBlock(CompanyID int, JobNo int, DepartmentNo int, StaffID int, Date string, StartTime string, EndTime string, LabourMultiplier string, ScheduleRateID string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobInsertScheduleBlock", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo)), strconv.Quote(fmt.Sprint(DepartmentNo)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(Date)), strconv.Quote(fmt.Sprint(StartTime)), strconv.Quote(fmt.Sprint(EndTime)), strconv.Quote(fmt.Sprint(LabourMultiplier)), strconv.Quote(fmt.Sprint(ScheduleRateID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieve(CompanyID int, JobNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieveAllocated(CompanyID int, Detail []interface{}, CustomerID int, StartDate string, EndDate string, Limit int, Offset int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieveAllocated", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieveAttachmentKey(CompanyID int, JobNo int, Filename string, Email bool, Public bool) (string, error) {
	var result struct {
		Result string `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieveAttachmentKey", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo)), strconv.Quote(fmt.Sprint(Filename)), strconv.Quote(fmt.Sprint(Email)), strconv.Quote(fmt.Sprint(Public))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieveAttachments(CompanyID int, JobNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieveAttachments", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieveDepartments(CompanyID int, JobNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieveDepartments", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieveExport(CompanyID int, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieveList(CompanyID int, JobNos []interface{}, Detail []interface{}, StaffID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNos)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(StaffID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobRetrieveSchedules(JobNo int, DepartmentNo int, StaffID int, Date string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobRetrieveSchedules", strings.Join([]string{strconv.Quote(fmt.Sprint(JobNo)), strconv.Quote(fmt.Sprint(DepartmentNo)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(Date))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobSearch(CompanyID int, Detail []interface{}, StaffID int, CustomerID int, StartDate string, EndDate string, Limit int, Offset int, SortAsc bool, Status string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc)), strconv.Quote(fmt.Sprint(Status))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobSearchFields(CompanyID interface{}, Detail []interface{}, SearchTerms []interface{}, StartDate string, EndDate string, Limit int, Offset int, SortAsc bool, RestrictSchedules bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(SearchTerms)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc)), strconv.Quote(fmt.Sprint(RestrictSchedules))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func JobUpdate(CompanyID int, JobNo int, JobDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "JobUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(JobNo)), strconv.Quote(fmt.Sprint(JobDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadConvert(CompanyID int, LeadID int, OrderNo int) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadConvert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LeadID)), strconv.Quote(fmt.Sprint(OrderNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadInsert(CompanyID int, LeadDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LeadDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadRetrieve(CompanyID int, LeadID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LeadID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadRetrieveAllocated(CompanyID int, Detail bool, CustomerID int, StartDate string, EndDate string, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadRetrieveAllocated", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadRetrieveAttachments(CompanyID int, LeadID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadRetrieveAttachments", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LeadID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadRetrieveList(CompanyID int, LeadIDs []interface{}, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LeadIDs)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadSarchClosed(CompanyID int, Detail []interface{}, CustomerID int, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadSarchClosed", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadSearchFields(CompanyID int, Detail []interface{}, SearchTerms []interface{}, StartDate string, EndDate string, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(SearchTerms)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadSearchOpen(CompanyID int, Detail []interface{}, CustomerID int, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadSearchOpen", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func LeadUpdate(CompanyID int, LeadID int, LeadDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "LeadUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(LeadID)), strconv.Quote(fmt.Sprint(LeadDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ProjectStatusRetrieve(CompanyID int, StatusID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ProjectStatusRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StatusID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ProjectStatusRetrieveList(CompanyID int, StatusIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ProjectStatusRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StatusIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ProjectStatusSearch(CompanyID int, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ProjectStatusSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ProjectTagsRetrieveList(CompanyID int, TagIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ProjectTagsRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TagIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ProjectTagsSearch(CompanyID interface{}, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ProjectTagsSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteAccept(CompanyID int, QuoteNo int, OrderNo string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteAccept", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(OrderNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteConvert(CompanyID int, QuoteNo int, OrderNo string) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteConvert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(OrderNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteDecline(CompanyID int, QuoteNo int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteDecline", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteDeleteAttachment(CompanyID int, QuoteNo int, FileName string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteDeleteAttachment", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(FileName))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteDeleteScheduleBlock(CompanyID int, QuoteNo int, DepartmentNo int, StaffID int, Date string, StartTime string, EndTime string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteDeleteScheduleBlock", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(DepartmentNo)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(Date)), strconv.Quote(fmt.Sprint(StartTime)), strconv.Quote(fmt.Sprint(EndTime))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteInsert(CompanyID int, QuoteDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteInsertScheduleBlock(CompanyID int, QuoteNo int, DepartmentNo int, StaffID int, Date string, StartTime string, EndTime string, LabourMultiplier string, ScheduleRateID string) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteInsertScheduleBlock", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(DepartmentNo)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(Date)), strconv.Quote(fmt.Sprint(StartTime)), strconv.Quote(fmt.Sprint(EndTime)), strconv.Quote(fmt.Sprint(LabourMultiplier)), strconv.Quote(fmt.Sprint(ScheduleRateID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteMerge(CompanyID int, QuoteNo int, JobNo int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteMerge", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(JobNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteRetrieve(CompanyID int, QuoteNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteRetrieveAllocated(CompanyID int, Detail bool, CustomerID int, StartDate string, EndDate string, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteRetrieveAllocated", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteRetrieveAttachmentKey(CompanyID int, QuoteNo int, Filename string, Email bool, Public bool) (string, error) {
	var result struct {
		Result string `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteRetrieveAttachmentKey", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(Filename)), strconv.Quote(fmt.Sprint(Email)), strconv.Quote(fmt.Sprint(Public))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteRetrieveAttachments(CompanyID int, QuoteNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteRetrieveAttachments", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteRetrieveDepartments(CompanyID int, QuoteNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteRetrieveDepartments", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteRetrieveList(CompanyID int, QuoteNos []interface{}, Detail []interface{}, StaffID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNos)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(StaffID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteRetrieveSchedules(QuoteNo int, DepartmentNo int, StaffID int, Date string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteRetrieveSchedules", strings.Join([]string{strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(DepartmentNo)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(Date))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteSearch(CompanyID int, Detail []interface{}, StaffID int, CustomerID int, StartDate string, EndDate string, Limit int, Offset int, SortAsc bool, Status string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc)), strconv.Quote(fmt.Sprint(Status))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteSearchClosed(CompanyID int, Detail []interface{}, CustomerID int, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteSearchClosed", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteSearchFields(CompanyID int, Detail []interface{}, SearchTerms []interface{}, StartDate string, EndDate string, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(SearchTerms)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteSearchOpen(CompanyID int, Detail []interface{}, CustomerID int, Limit int, Offset int, SortAsc bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteSearchOpen", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(CustomerID)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset)), strconv.Quote(fmt.Sprint(SortAsc))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func QuoteUpdate(CompanyID int, QuoteNo int, QuoteDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "QuoteUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(QuoteNo)), strconv.Quote(fmt.Sprint(QuoteDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteContactInsert(CompanyID int, SiteID int, SiteContactDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteContactInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteID)), strconv.Quote(fmt.Sprint(SiteContactDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteContactRetrieve(CompanyID int, SiteID int, ContactID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteContactRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteID)), strconv.Quote(fmt.Sprint(ContactID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteContactRetrieveList(CompanyID int, SiteIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteContactRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteInsert(CompanyID int, SiteDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SitePreferredTechnician(CompanyID int, SiteID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SitePreferredTechnician", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteRetrieve(CompanyID int, SiteID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteRetrieveList(CompanyID int, SiteIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteSearch(CompanyID int, Search string, Limit int, Offset int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteSearchFields(CompanyID int, SearchTerms []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchTerms))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteUpdate(CompanyID int, SiteID int, SiteDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteID)), strconv.Quote(fmt.Sprint(SiteDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SiteValueForecast(CompanyID int, SiteID int, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SiteValueForecast", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SiteID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ZoneRetrieve(CompanyID int, ZoneID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ZoneRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ZoneID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ZoneRetrieveList(CompanyID int, ZoneIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ZoneRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ZoneIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ZoneSearch(CompanyID int, Name string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ZoneSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Name))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorAttachmentsRetrieve(CompanyID int, ContractorStaffID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorAttachmentsRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorStaffID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorInsert(CompanyID int, ContractorDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorLicencesRetrieve(CompanyID int, ContractorStaffID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorLicencesRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorStaffID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorLicencesSearch(CompanyID int, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorLicencesSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorRetrieve(CompanyID interface{}, ContractorStaffID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorStaffID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorRetrieveList(CompanyID interface{}, ContractorStaffIDs []interface{}, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorStaffIDs)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorRetrieveStaff(CompanyID interface{}, ContractorID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorRetrieveStaff", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorSearch(CompanyID interface{}, Search string, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorSearchFields(CompanyID interface{}, SearchTerms []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchTerms))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorTotalCount(CompanyID interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorTotalCount", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func ContractorUpdate(CompanyID int, ContractorStaffID int, ContractorDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "ContractorUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(ContractorStaffID)), strconv.Quote(fmt.Sprint(ContractorDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeAttachmentsRetrieve(CompanyID interface{}, EmployeeID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeAttachmentsRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeFlag(CompanyID int, EmployeeID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeFlagList(CompanyID int, EmployeeIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeInsert(CompanyID int, EmployeeDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeLicenceInsert(CompanyID int, EmployeeID int, LicenceDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeLicenceInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeID)), strconv.Quote(fmt.Sprint(LicenceDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeLicencesRetrieve(CompanyID int, EmployeeID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeLicencesRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeLicencesSearch(CompanyID interface{}, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeLicencesSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeRetrieve(CompanyID interface{}, EmployeeID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeRetrieveAll(CompanyID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeRetrieveAll", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeRetrieveCurrent(Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeRetrieveCurrent", strings.Join([]string{strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeRetrieveExport(CompanyID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeRetrieveExport", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeRetrieveList(CompanyID interface{}, EmployeeIDs []interface{}, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeIDs)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeSearch(CompanyID interface{}, Search string, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeSearchFields(CompanyID interface{}, SearchTerms []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchTerms))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeTotalCount(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeTotalCount", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeUnFlag(CompanyID int, EmployeeID int) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeUnFlag", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeUnFlagList(CompanyID int, EmployeeIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeUnFlagList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func EmployeeUpdate(CompanyID int, EmployeeID int, EmployeeDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "EmployeeUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(EmployeeID)), strconv.Quote(fmt.Sprint(EmployeeDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PlantRetrieve(CompanyID interface{}, PlantID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PlantRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PlantID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PlantRetrieveList(CompanyID interface{}, PlantIDs []interface{}, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PlantRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PlantIDs)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PlantSearch(CompanyID interface{}, Search string, Detail []interface{}, Schedulable bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PlantSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(Schedulable))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PlantSearchFields(CompanyID interface{}, SearchTerms []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PlantSearchFields", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SearchTerms))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func PlantTypesRetrieve(CompanyID interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "PlantTypesRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StaffAttachmentsRetrieve(CompanyID int, StaffID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StaffAttachmentsRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StaffID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StaffLicencesSearch(CompanyID int, Search string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StaffLicencesSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func StaffRetrieveList(CompanyID int, StaffIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "StaffRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StaffIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TimesheetRetrieve(EmployeeID int, StartDate string, EndDate string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TimesheetRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(EmployeeID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierBranchInsert(CompanyID int, SupplierID int, SupplierBranchDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierBranchInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierID)), strconv.Quote(fmt.Sprint(SupplierBranchDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierBranchRetrieve(CompanyID int, SupplierID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierBranchRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierBranchRetrieveList(CompanyID int, SupplierIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierBranchRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierContactInsert(CompanyID int, SupplierID int, SupplierContactDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierContactInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierID)), strconv.Quote(fmt.Sprint(SupplierContactDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierContactRetrieve(CompanyID int, SupplierID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierContactRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierContactRetrieveList(CompanyID int, SupplierIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierContactRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierInsert(CompanyID int, SupplierDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierRetrieve(CompanyID int, SupplierID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierRetrieveList(CompanyID int, SupplierIDs []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierIDs))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierSearch(CompanyID int, Search string, Limit int, Offset int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(Limit)), strconv.Quote(fmt.Sprint(Offset))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierTotalCount(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierTotalCount", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierUpdate(CompanyID int, SupplierID int, SupplierDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierID)), strconv.Quote(fmt.Sprint(SupplierDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierOrderInsert(CompanyID int, SupplierOrderDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierOrderInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierOrderDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierOrderRetrieve(CompanyID int, OrderID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierOrderRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(OrderID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierOrderRetrieveApproved(CompanyID int, StartDate string, EndDate string, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierOrderRetrieveApproved", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierOrderRetrieveList(CompanyID int, SupplierOrderIDs []interface{}, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierOrderRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierOrderIDs)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierOrderRetrievePending(CompanyID int, StartDate string, EndDate string, Detail []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierOrderRetrievePending", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierOrderSearch(CompanyID int, Search string, StartDate string, EndDate string, Detail []interface{}, Status string) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierOrderSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(Search)), strconv.Quote(fmt.Sprint(StartDate)), strconv.Quote(fmt.Sprint(EndDate)), strconv.Quote(fmt.Sprint(Detail)), strconv.Quote(fmt.Sprint(Status))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func SupplierOrderUpdate(CompanyID int, SupplierOrderID int, SupplierOrderDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "SupplierOrderUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(SupplierOrderID)), strconv.Quote(fmt.Sprint(SupplierOrderDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskCategories(CompanyID int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskCategories", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskInsert(CompanyID int, TaskDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskInsert", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TaskDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskInsertSelfAppointed(CompanyID int, TaskDetails []interface{}) (int, error) {
	var result struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskInsertSelfAppointed", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TaskDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskRetrieve(CompanyID int, TaskNo int) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskRetrieve", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TaskNo))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskRetrieveList(CompanyID int, TaskNos []interface{}, Detail bool) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskRetrieveList", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TaskNos)), strconv.Quote(fmt.Sprint(Detail))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskSearch(CompanyID int, StaffID int, JobNo int, StatusList []interface{}) ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskSearch", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(StaffID)), strconv.Quote(fmt.Sprint(JobNo)), strconv.Quote(fmt.Sprint(StatusList))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskUpdate(CompanyID int, TaskNo int, TaskDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskUpdate", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TaskNo)), strconv.Quote(fmt.Sprint(TaskDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func TaskUpdateAllocated(CompanyID int, TaskID int, TaskDetails []interface{}) (bool, error) {
	var result struct {
		Result bool   `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "TaskUpdateAllocated", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(TaskID)), strconv.Quote(fmt.Sprint(TaskDetails))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func VersionRetrieve() ([]interface{}, error) {
	var result struct {
		Result []interface{} `json:"result"`
		Error  string        `json:"error"`
		Id     string        `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "VersionRetrieve", strings.Join([]string{}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}

func VersionSaveAccountingLinkVersion(CompanyID int, PackageType string, AccountingLinkVersion string) (string, error) {
	var result struct {
		Result string `json:"result"`
		Error  string `json:"error"`
		Id     string `json:"id"`
	}
	param := fmt.Sprintf(`{"id":0,"method":"%s","params":[%s]}`, "VersionSaveAccountingLinkVersion", strings.Join([]string{strconv.Quote(fmt.Sprint(CompanyID)), strconv.Quote(fmt.Sprint(PackageType)), strconv.Quote(fmt.Sprint(AccountingLinkVersion))}, ", "))
	req, err := http.NewRequest("POST", "https://ubelssoftwaredevelopment.simprosuite.com/api/?format=json", strings.NewReader(param))
	if err != nil {
		return result.Result, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err := go_rfc5849_hmac.SignSha1Hmac1(req, param); err != nil {
		return result.Result, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result.Result, err
	}
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		return result.Result, err
	} else if err := json.Unmarshal(b, &result); err != nil {
		return result.Result, err
	}
	if len(result.Error) > 0 {
		return result.Result, errors.New(result.Error)
	}
	return result.Result, err
}
