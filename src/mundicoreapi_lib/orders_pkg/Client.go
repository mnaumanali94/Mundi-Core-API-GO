/*
 * mundicoreapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 01/02/2017
 */
package orders_pkg


import(
	"encoding/json"
	"mundicoreapi_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"mundicoreapi_lib"
	"mundicoreapi_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type ORDERS_IMPL struct { }

/**
 * <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td>`items` [obrigatório]</td>      <td>array of objects</td>      <td>Itens do pedido. Leia mais sobre [item do pedido](#Pedidos)</td>    </tr>    <tr>      <td>`customer_id` ou `customer` [obrigatório]</td>      <td>object</td>      <td>Comprador. Leia mais sobre [comprador](#Compradores)</td>    </tr>    <tr>      <td>`payment` [obrigatório]</td>      <td>object</td>      <td>Pagamento. Leia mais sobre [pagamento](#payment)</td>    </tr>    <tr>      <td>`code`</td>      <td>string(52)</td>      <td>Referência do pedido</td>    </tr>    <tr>      <td>`shipping`</td>      <td>object</td>      <td>Entrega. Leia mais sobre [entrega](#shipping)</td>    </tr>    <tr>      <td>`metadata`</td>      <td>object</td>      <td>Informações adicionais do pedido. Leia mais sobre [metadata](#metadata)</td>    </tr>  </tbody></table>
 * @param    *models_pkg.CreateOrderRequest        body     parameter: Required
 * @return	Returns the *models_pkg.GetOrderResponse response from the API call
 */
func (me *ORDERS_IMPL) CreateOrder (
            body *models_pkg.CreateOrderRequest) (*models_pkg.GetOrderResponse, error) {
        //the base uri for api requests
    _queryBuilder := mundicoreapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/orders"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
    }

    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, body, mundicoreapi_lib.Config.BasicAuthUserName, mundicoreapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Requisição inválida", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Chave de API inválida", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("O recurso solicitado não existe", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Parâmetros válidos mas a requisição falhou", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Parâmetros inválidos", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Ocorreu um erro interno", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetOrderResponse = &models_pkg.GetOrderResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * <p>Utilize os parâmetros <code>page</code> e <code>size</code> para paginar o resultado. Leia mais sobre <a href="#pagination">paginação</a></p>
 * @return	Returns the interface{} response from the API call
 */
func (me *ORDERS_IMPL) ListOrders () (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := mundicoreapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/orders"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, mundicoreapi_lib.Config.BasicAuthUserName, mundicoreapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Requisição inválida", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Chave de API inválida", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("O recurso solicitado não existe", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Parâmetros válidos mas a requisição falhou", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Parâmetros inválidos", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Ocorreu um erro interno", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * TODO: type endpoint description here
 * @param    string        orderId      parameter: Required
 * @return	Returns the *models_pkg.GetOrderResponse response from the API call
 */
func (me *ORDERS_IMPL) GetOrder (
            orderId string) (*models_pkg.GetOrderResponse, error) {
        //the base uri for api requests
    _queryBuilder := mundicoreapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/orders/{order_id}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "order_id" : orderId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, mundicoreapi_lib.Config.BasicAuthUserName, mundicoreapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("Requisição inválida", _response.Code, _response.RawBody)
    } else if (_response.Code == 401) {
        err = apihelper_pkg.NewAPIError("Chave de API inválida", _response.Code, _response.RawBody)
    } else if (_response.Code == 404) {
        err = apihelper_pkg.NewAPIError("O recurso solicitado não existe", _response.Code, _response.RawBody)
    } else if (_response.Code == 412) {
        err = apihelper_pkg.NewAPIError("Parâmetros válidos mas a requisição falhou", _response.Code, _response.RawBody)
    } else if (_response.Code == 422) {
        err = apihelper_pkg.NewAPIError("Parâmetros inválidos", _response.Code, _response.RawBody)
    } else if (_response.Code == 500) {
        err = apihelper_pkg.NewAPIError("Ocorreu um erro interno", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.GetOrderResponse = &models_pkg.GetOrderResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

