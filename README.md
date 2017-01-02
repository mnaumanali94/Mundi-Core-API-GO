#Getting started

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](http://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](http://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](http://apidocs.io/illustration/go?step=import2&workspaceFolder=Mundi%20Core%20API-GoLang&projectName=mundicoreapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](http://apidocs.io/illustration/go?step=import3&projectName=mundicoreapi_lib)

## How to Use

The following section explains how to use the MundiCoreAPI library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](http://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](http://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](http://apidocs.io/illustration/go?step=createNewProject2&projectName=mundicoreapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](http://apidocs.io/illustration/go?step=createNewProject3&projectName=mundicoreapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](http://apidocs.io/illustration/go?step=createNewProject4&projectName=mundicoreapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](http://apidocs.io/illustration/go?step=createNewProject5&projectName=mundicoreapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](http://apidocs.io/illustration/go?step=testProject0&projectName=mundicoreapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](http://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](http://apidocs.io/illustration/go?step=testProject2&workspaceFolder=Mundi%20Core%20API-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](http://apidocs.io/illustration/go?step=buildProject&projectName=mundicoreapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](http://apidocs.io/illustration/go?step=runProject&projectName=mundicoreapi_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


## Class Reference

### <a name="list_of_controllers"></a>List of Controllers

* [subscriptions_pkg](#subscriptions_pkg)
* [orders_pkg](#orders_pkg)
* [plans_pkg](#plans_pkg)
* [customers_pkg](#customers_pkg)
* [charges_pkg](#charges_pkg)
* [invoices_pkg](#invoices_pkg)

### <a name="subscriptions_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".subscriptions_pkg") subscriptions_pkg

#### Get instance

Factory for the ``` SUBSCRIPTIONS ``` interface can be accessed from the package subscriptions_pkg.

```go
subscriptions := subscriptions_pkg.NewSUBSCRIPTIONS()
```

#### <a name="update_subscription_payment_method"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionPaymentMethod") UpdateSubscriptionPaymentMethod

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td>`payment_method` [obrigatório]</td>      <td>enum</td>      <td>**credit_card** ou **boleto**</td>    </tr>    <tr>      <td>`credit_card_id` ou `credit_card`</td>      <td>object</td>      <td>Cartão de crédito. Obrigatório caso o método de pagamento seja **credit_card**. Leia mais sobre [cartão de crédito](#Cartoes-de-credito)</td>    </tr>  </tbody></table>


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionPaymentMethod(
            body *models_pkg.UpdateSubscriptionPaymentMethodRequest,
            subscriptionId string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.UpdateSubscriptionPaymentMethodRequest
subscriptionId := "subscription_id"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionPaymentMethod(body, subscriptionId)

```


#### <a name="create_discount"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.CreateDiscount") CreateDiscount

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td>`value` (obrigatório)</td>      <td>decimal</td>      <td>Valor do desconto</td>    </tr>    <tr>      <td>`discount_type` (obrigatório)</td>      <td>enum</td>      <td>**flat** ou **percentage**. O padrão é **percentage**</td>    </tr>    <tr>      <td>`cycles`</td>      <td>integer</td>      <td>Indica quantas vezes o desconto será aplicado. Caso não seja informado, o desconto será aplicado até que seja excluído</td>    </tr>    <tr>      <td>`item_id`</td>      <td>string(36)</td>      <td>Código do item da assinatura. Caso seja informado o desconto será aplicado no item</td>    </tr>  </tbody></table>


```go
func (me *SUBSCRIPTIONS_IMPL) CreateDiscount(
            body *models_pkg.CreateDiscountRequest,
            subscriptionId string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateDiscountRequest
subscriptionId := "subscription_id"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.CreateDiscount(body, subscriptionId)

```


#### <a name="delete_subscription_item"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.DeleteSubscriptionItem") DeleteSubscriptionItem

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteSubscriptionItem(
            subscriptionId string,
            subscriptionItemId string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| subscriptionItemId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
subscriptionItemId := "subscription_item_id"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.DeleteSubscriptionItem(subscriptionId, subscriptionItemId)

```


#### <a name="update_subscription_credit_card"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionCreditCard") UpdateSubscriptionCreditCard

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td>`credit_card_id` ou `credit_card` [obrigatório]</td>      <td>object</td>      <td>Cartão de crédito. Leia mais sobre [cartão de crédito](#Cartoes-de-credito)</td>    </tr>  </tbody></table>


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionCreditCard(
            body *models_pkg.UpdateSubscriptionCreditCardRequest,
            subscriptionId string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.UpdateSubscriptionCreditCardRequest
subscriptionId := "subscription_id"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateSubscriptionCreditCard(body, subscriptionId)

```


#### <a name="update_reschedule_subscription"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.UpdateRescheduleSubscription") UpdateRescheduleSubscription

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateRescheduleSubscription(
            subscriptionId string,
            body *models_pkg.CreateRescheduleSubscriptionRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.CreateRescheduleSubscriptionRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.UpdateRescheduleSubscription(subscriptionId, body)

```


#### <a name="update_subscription_item"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.UpdateSubscriptionItem") UpdateSubscriptionItem

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) UpdateSubscriptionItem(
            subscriptionId string,
            itemId string,
            body *models_pkg.UpdateSubscriptionItemRequest)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| itemId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
var body *models_pkg.UpdateSubscriptionItemRequest

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.UpdateSubscriptionItem(subscriptionId, itemId, body)

```


#### <a name="create_usage"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.CreateUsage") CreateUsage

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) CreateUsage(
            subscriptionId string,
            itemId string,
            body *models_pkg.CreateUsageRequest)(*models_pkg.GetUsageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| itemId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
var body *models_pkg.CreateUsageRequest

var result *models_pkg.GetUsageResponse
result,_ = subscriptions.CreateUsage(subscriptionId, itemId, body)

```


#### <a name="get_usage"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetUsage") GetUsage

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetUsage(
            subscriptionId string,
            itemId string,
            usageId string)(*models_pkg.GetUsageResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| itemId |  ``` Required ```  | TODO: Add a parameter description |
| usageId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"
usageId := "usage_id"

var result *models_pkg.GetUsageResponse
result,_ = subscriptions.GetUsage(subscriptionId, itemId, usageId)

```


#### <a name="get_subscription"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscription") GetSubscription

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscription(subscriptionId string)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.GetSubscription(subscriptionId)

```


#### <a name="delete_cancel_subscription"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.DeleteCancelSubscription") DeleteCancelSubscription

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteCancelSubscription(
            subscriptionId string,
            body *models_pkg.CreateCancelSubscriptionRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Optional ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
var body *models_pkg.CreateCancelSubscriptionRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.DeleteCancelSubscription(subscriptionId, body)

```


#### <a name="list_subscriptions"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.ListSubscriptions") ListSubscriptions

> <p>Utilize os parâmetros <code>page</code> e <code>size</code> para paginar o resultado. Leia mais sobre <a href="#pagination">paginação</a></p>


```go
func (me *SUBSCRIPTIONS_IMPL) ListSubscriptions()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = subscriptions.ListSubscriptions()

```


#### <a name="delete_discount"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.DeleteDiscount") DeleteDiscount

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) DeleteDiscount(
            subscriptionId string,
            discountId string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| discountId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discount_id"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.DeleteDiscount(subscriptionId, discountId)

```


#### <a name="create_discount1"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.CreateDiscount1") CreateDiscount1

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) CreateDiscount1(
            body *models_pkg.CreateDiscountRequest,
            subscriptionId string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateDiscountRequest
subscriptionId := "subscription_id"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.CreateDiscount1(body, subscriptionId)

```


#### <a name="get_subscription_item"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionItem") GetSubscriptionItem

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionItem(
            subscriptionId string,
            itemId string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| itemId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
itemId := "item_id"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.GetSubscriptionItem(subscriptionId, itemId)

```


#### <a name="get_delete_discount"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetDeleteDiscount") GetDeleteDiscount

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetDeleteDiscount(
            subscriptionId string,
            discountId string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| discountId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discount_id"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.GetDeleteDiscount(subscriptionId, discountId)

```


#### <a name="get_discount"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetDiscount") GetDiscount

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetDiscount(
            subscriptionId string,
            discountId string)(*models_pkg.GetDiscountResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| discountId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
discountId := "discount_id"

var result *models_pkg.GetDiscountResponse
result,_ = subscriptions.GetDiscount(subscriptionId, discountId)

```


#### <a name="get_cycle"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetCycle") GetCycle

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetCycle(
            subscriptionId string,
            cycleId string)(*models_pkg.GetPeriodResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| cycleId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
cycleId := "cycle_id"

var result *models_pkg.GetPeriodResponse
result,_ = subscriptions.GetCycle(subscriptionId, cycleId)

```


#### <a name="get_subscription_cycles"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionCycles") GetSubscriptionCycles

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionCycles(subscriptionId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result interface{}
result,_ = subscriptions.GetSubscriptionCycles(subscriptionId)

```


#### <a name="get_subscription_invoices"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetSubscriptionInvoices") GetSubscriptionInvoices

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetSubscriptionInvoices(subscriptionId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result interface{}
result,_ = subscriptions.GetSubscriptionInvoices(subscriptionId)

```


#### <a name="create_invoice"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.CreateInvoice") CreateInvoice

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) CreateInvoice(
            subscriptionId string,
            periodId string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |
| periodId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"
periodId := "period_id"

var result *models_pkg.GetInvoiceResponse
result,_ = subscriptions.CreateInvoice(subscriptionId, periodId)

```


#### <a name="get_last_subscription_charge"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetLastSubscriptionCharge") GetLastSubscriptionCharge

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetLastSubscriptionCharge(subscriptionId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.GetChargeResponse
result,_ = subscriptions.GetLastSubscriptionCharge(subscriptionId)

```


#### <a name="get_last_subscription_charge1"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetLastSubscriptionCharge1") GetLastSubscriptionCharge1

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetLastSubscriptionCharge1(subscriptionId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.GetChargeResponse
result,_ = subscriptions.GetLastSubscriptionCharge1(subscriptionId)

```


#### <a name="get_last_subscription_charge2"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.GetLastSubscriptionCharge2") GetLastSubscriptionCharge2

> TODO: Add a method description


```go
func (me *SUBSCRIPTIONS_IMPL) GetLastSubscriptionCharge2(subscriptionId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
subscriptionId := "subscription_id"

var result *models_pkg.GetChargeResponse
result,_ = subscriptions.GetLastSubscriptionCharge2(subscriptionId)

```


#### <a name="create_subscription"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscription") CreateSubscription

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td>`plan_id` [obrigatório]</td>      <td>string(36)</td>      <td>Código do plano</td>    </tr>    <tr>      <td>`customer_id` ou `customer` [obrigatório]</td>      <td>object</td>      <td>Comprador. Leia mais sobre [comprador](#Compradores)</td>    </tr>    <tr>      <td>`payment_method`</td>      <td>enum</td>      <td>**credit_card** ou **boleto**. Obrigatório caso o plano possua mais de uma opção de método de pagamento. O padrão é **credit_card**</td>    </tr>    <tr>      <td>`credit_card_id` ou `credit_card`</td>      <td>object</td>      <td>Cartão de crédito. Obrigatório caso o método de pagamento seja **credit_card**. Leia mais sobre [cartão de crédito](#Cartoes-de-credito)</td>    </tr>    <tr>      <td>`installments`</td>      <td>integer</td>      <td>Quantidade de parcelas. Obrigatório caso o plano possua mais de uma opção de parcela</td>    </tr>    <tr>      <td>`billing_day`</td>      <td>integer</td>      <td>Dia de cobrança. Obrigatório caso o plano possua mais de uma opção de dia da cobrança</td>    </tr>    <tr>      <td>`shipping`</td>      <td>object</td>      <td>Entrega. Obrigatório caso o plano exija os dados de entrega. Leia mais sobre [entrega](#Entrega)</td>    </tr>    <tr>      <td>`code`</td>      <td>string(52)</td>      <td>Referência da assinatura</td>    </tr>    <tr>      <td>`start_at`</td>      <td>date</td>      <td>Data de início da assinatura. Caso não seja informado, a assinatura será iniciada imediatamente</td>    </tr>    <tr>      <td>`discounts`</td>      <td>array of objects</td>      <td>Descontos. Leia mais sobre [desconto](#Descontos)</td>    </tr>    <tr>      <td>`metadata`</td>      <td>object</td>      <td>Informações adicionais sobre a assinatura. Leia mais sobre [metadata](#metadata)</td>    </tr>  </tbody></table>


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscription(body *models_pkg.CreateSubscriptionRequest)(*models_pkg.GetSubscriptionResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateSubscriptionRequest

var result *models_pkg.GetSubscriptionResponse
result,_ = subscriptions.CreateSubscription(body)

```


#### <a name="create_subscription_item"></a>![Method: ](http://apidocs.io/img/method.png ".subscriptions_pkg.CreateSubscriptionItem") CreateSubscriptionItem

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td>`plan_item_id`</td>      <td>string(36)</td>      <td>Código do item do plano</td>    </tr>    <tr>      <td>`description`</td>      <td>string(256)</td>      <td>Descrição. Obrigatório caso o `plan_item_id` não seja informado</td>    </tr>    <tr>      <td>`cycles`</td>      <td>integer</td>      <td>Indica quantas vezes o item será cobrado. Caso não seja informado, o item será cobrado até que seja excluído o desativado</td>    </tr>    <tr>      <td>`quantity`</td>      <td>integer</td>      <td>Quantidade de itens. Obrigatório caso `pricing_scheme.scheme_type` seja **unit**</td>    </tr>    <tr>      <td>`pricing_scheme`</td>      <td>object</td>      <td>Esquema de precificação. Leia mais sobre [precificação](#pricing)</td>    </tr>    <tr>      <td>`discounts`</td>      <td>array of objects</td>      <td>Descontos do item. Leia mais sobre [desconto](#Descontos)</td>    </tr>  </tbody></table>


```go
func (me *SUBSCRIPTIONS_IMPL) CreateSubscriptionItem(
            body *models_pkg.CreateSubscriptionItemRequest,
            subscriptionId string)(*models_pkg.GetSubscriptionItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| subscriptionId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateSubscriptionItemRequest
subscriptionId := "subscription_id"

var result *models_pkg.GetSubscriptionItemResponse
result,_ = subscriptions.CreateSubscriptionItem(body, subscriptionId)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="orders_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".orders_pkg") orders_pkg

#### Get instance

Factory for the ``` ORDERS ``` interface can be accessed from the package orders_pkg.

```go
orders := orders_pkg.NewORDERS()
```

#### <a name="create_order"></a>![Method: ](http://apidocs.io/img/method.png ".orders_pkg.CreateOrder") CreateOrder

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td>`items` [obrigatório]</td>      <td>array of objects</td>      <td>Itens do pedido. Leia mais sobre [item do pedido](#Pedidos)</td>    </tr>    <tr>      <td>`customer_id` ou `customer` [obrigatório]</td>      <td>object</td>      <td>Comprador. Leia mais sobre [comprador](#Compradores)</td>    </tr>    <tr>      <td>`payment` [obrigatório]</td>      <td>object</td>      <td>Pagamento. Leia mais sobre [pagamento](#payment)</td>    </tr>    <tr>      <td>`code`</td>      <td>string(52)</td>      <td>Referência do pedido</td>    </tr>    <tr>      <td>`shipping`</td>      <td>object</td>      <td>Entrega. Leia mais sobre [entrega](#shipping)</td>    </tr>    <tr>      <td>`metadata`</td>      <td>object</td>      <td>Informações adicionais do pedido. Leia mais sobre [metadata](#metadata)</td>    </tr>  </tbody></table>


```go
func (me *ORDERS_IMPL) CreateOrder(body *models_pkg.CreateOrderRequest)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateOrderRequest

var result *models_pkg.GetOrderResponse
result,_ = orders.CreateOrder(body)

```


#### <a name="list_orders"></a>![Method: ](http://apidocs.io/img/method.png ".orders_pkg.ListOrders") ListOrders

> <p>Utilize os parâmetros <code>page</code> e <code>size</code> para paginar o resultado. Leia mais sobre <a href="#pagination">paginação</a></p>


```go
func (me *ORDERS_IMPL) ListOrders()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = orders.ListOrders()

```


#### <a name="get_order"></a>![Method: ](http://apidocs.io/img/method.png ".orders_pkg.GetOrder") GetOrder

> TODO: Add a method description


```go
func (me *ORDERS_IMPL) GetOrder(orderId string)(*models_pkg.GetOrderResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| orderId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
orderId := "order_id"

var result *models_pkg.GetOrderResponse
result,_ = orders.GetOrder(orderId)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="plans_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".plans_pkg") plans_pkg

#### Get instance

Factory for the ``` PLANS ``` interface can be accessed from the package plans_pkg.

```go
plans := plans_pkg.NewPLANS()
```

#### <a name="create_plan_item"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.CreatePlanItem") CreatePlanItem

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td><code>name</code> [obrigatório]</td>      <td>string(64)</td>      <td>Nome</td>    </tr>    <tr>      <td><code>pricing_scheme</code> [obrigatório]</td>      <td>object</td>      <td>Esquema de precificação. Leia mais sobre <a href="#pricing">precificação</a></td>    </tr>    <tr>      <td><code>cycles</code></td>      <td>integer</td>      <td>Indica quantas vezes o item será cobrado. Caso não seja informado, o item será cobrado até que seja excluído o desativado.      </td>    </tr>    <tr>      <td><code>description</code></td>      <td>string(256)</td>      <td>Descrição</td>    </tr>    <tr>      <td><code>quantity</code></td>      <td>integer</td>      <td>Quantidade de itens</td>    </tr>  </tbody></table>


```go
func (me *PLANS_IMPL) CreatePlanItem(
            body *models_pkg.CreatePlanItemRequest,
            planId string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| planId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreatePlanItemRequest
planId := "plan_id"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.CreatePlanItem(body, planId)

```


#### <a name="update_plan_item"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.UpdatePlanItem") UpdatePlanItem

> TODO: Add a method description


```go
func (me *PLANS_IMPL) UpdatePlanItem(
            planId string,
            planItemId string,
            body *models_pkg.UpdatePlanItemRequest)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | TODO: Add a parameter description |
| planItemId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"
var body *models_pkg.UpdatePlanItemRequest

var result *models_pkg.GetPlanItemResponse
result,_ = plans.UpdatePlanItem(planId, planItemId, body)

```


#### <a name="delete_plan"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.DeletePlan") DeletePlan

> TODO: Add a method description


```go
func (me *PLANS_IMPL) DeletePlan(planId string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"

var result *models_pkg.GetPlanResponse
result,_ = plans.DeletePlan(planId)

```


#### <a name="list_plans"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.ListPlans") ListPlans

> <p>Utilize os parâmetros <code>page</code> e <code>size</code> para paginar o resultado. Leia mais sobre <a href="#pagination">paginação</a></p>


```go
func (me *PLANS_IMPL) ListPlans()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = plans.ListPlans()

```


#### <a name="delete_plan_item"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.DeletePlanItem") DeletePlanItem

> TODO: Add a method description


```go
func (me *PLANS_IMPL) DeletePlanItem(
            planId string,
            planItemId string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | TODO: Add a parameter description |
| planItemId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.DeletePlanItem(planId, planItemId)

```


#### <a name="get_plan"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.GetPlan") GetPlan

> TODO: Add a method description


```go
func (me *PLANS_IMPL) GetPlan(planId string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"

var result *models_pkg.GetPlanResponse
result,_ = plans.GetPlan(planId)

```


#### <a name="get_plan_item"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.GetPlanItem") GetPlanItem

> TODO: Add a method description


```go
func (me *PLANS_IMPL) GetPlanItem(
            planId string,
            planItemId string)(*models_pkg.GetPlanItemResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | TODO: Add a parameter description |
| planItemId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"
planItemId := "plan_item_id"

var result *models_pkg.GetPlanItemResponse
result,_ = plans.GetPlanItem(planId, planItemId)

```


#### <a name="get_plan_subscriptions"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.GetPlanSubscriptions") GetPlanSubscriptions

> TODO: Add a method description


```go
func (me *PLANS_IMPL) GetPlanSubscriptions(planId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"

var result interface{}
result,_ = plans.GetPlanSubscriptions(planId)

```


#### <a name="get_plan_items"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.GetPlanItems") GetPlanItems

> TODO: Add a method description


```go
func (me *PLANS_IMPL) GetPlanItems(planId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| planId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
planId := "plan_id"

var result interface{}
result,_ = plans.GetPlanItems(planId)

```


#### <a name="update_plan"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.UpdatePlan") UpdatePlan

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td><code>name</code> [obrigatório]</td>      <td>string(64)</td>      <td>Nome</td>    </tr>    <tr>      <td><code>status</code> [obrigatório]</td>      <td>enum</td>      <td>**active** ou **inactive**</td>    </tr>    <tr>      <td><code>interval</code> [obrigatório]</td>      <td>enum</td>      <td><strong>week</strong>, <strong>month</strong>, <strong>year</strong></td>    </tr>    <tr>      <td><code>interval_count</code> [obrigatório]</td>      <td>integer</td>      <td>Repetições do intervalo</td>    </tr>    <tr>      <td><code>currency</code> [obrigatório]</td>      <td>enum</td>      <td><strong>BRL</strong></td>    </tr>    <tr>      <td><code>payment_methods</code> [obrigatório]</td>      <td>array of strings</td>      <td><strong>credit_card</strong> ou <strong>boleto</strong></td>    </tr>    <tr>      <td><code>billing_type</code> [obrigatório]</td>      <td>enum</td>      <td><strong>prepaid</strong>, <strong>postpaid</strong> ou <strong>exact_day</strong></td>    </tr>    <tr>      <td><code>billing_days</code></td>      <td>array of integers</td>      <td>Dias disponíveis para cobrança. Obrigatório caso `billing_type` seja **exact_day**</td>    </tr>    <tr>      <td><code>description</code></td>      <td>string(256)</td>      <td>Descrição</td>    </tr>    <tr>      <td><code>installments</code></td>      <td>array of integers</td>      <td>Parcelas disponíveis</td>    </tr>    <tr>      <td><code>statement_descriptor</code></td>      <td>string(22)</td>      <td>Texto exibido na fatura do cartão</td>    </tr>    <tr>      <td><code>trial_period_days</code></td>      <td>integer</td>      <td>Dias para início das assinaturas</td>    </tr>    <tr>      <td><code>shippable</code></td>      <td>boolean</td>      <td>Indica se o plano oferece entrega</td>    </tr>    <tr>      <td><code>metadata</code></td>      <td>object</td>      <td>Informações adicionais sobre o plano. Leia mais sobre [metadata](#metadata)</td>    </tr>  </tbody></table>


```go
func (me *PLANS_IMPL) UpdatePlan(
            body *models_pkg.UpdatePlanRequest,
            planId string)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| planId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.UpdatePlanRequest
planId := "plan_id"

var result *models_pkg.GetPlanResponse
result,_ = plans.UpdatePlan(body, planId)

```


#### <a name="create_plan"></a>![Method: ](http://apidocs.io/img/method.png ".plans_pkg.CreatePlan") CreatePlan

> <table>  <thead>    <tr>      <th>Campo</th>      <th>Tipo</th>      <th>Descrição</th>    </tr>  </thead>  <tbody>    <tr>      <td><code>name</code> [obrigatório]</td>      <td>string(64)</td>      <td>Nome</td>    </tr>    <tr>      <td><code>interval</code> [obrigatório]</td>      <td>enum</td>      <td><strong>week</strong>, <strong>month</strong>, <strong>year</strong>. O padrão é **month**</td>    </tr>     <tr>      <td><code>interval_count</code> [obrigatório]</td>      <td>integer</td>      <td>Repetições do intervalo. O padrão é **1**</td>    </tr>    <tr>      <td><code>currency</code> [obrigatório]</td>      <td>enum</td>      <td><strong>BRL</strong></td>    </tr>    <tr>      <td><code>payment_methods</code> [obrigatório]</td>      <td>array of strings</td>      <td><strong>credit_card</strong> ou <strong>boleto</strong>. O padrão é **credit_card**</td>    </tr>    <tr>      <td><code>items</code> [obrigatório]</td>      <td>array of objects</td>      <td>Itens. Leia mais sobre <a href="#">item do plano</a></td>    </tr>    <tr>      <td><code>billing_type</code> [obrigatório]</td>      <td>enum</td>      <td><strong>prepaid</strong>, <strong>postpaid</strong> ou <strong>exact_day</strong>. O padrão é **prepaid**</td>    </tr>    <tr>      <td><code>billing_days</code></td>      <td>array of integers</td>      <td>Dias disponíveis para cobrança. Obrigatório caso `billing_type` seja **exact_day**</td>    </tr>    <tr>      <td><code>description</code></td>      <td>string(256)</td>      <td>Descrição</td>    </tr>    <tr>      <td><code>installments</code></td>      <td>array of integers</td>      <td>Parcelas disponíveis. O padrão é **1**</td>    </tr>    <tr>      <td><code>statement_descriptor</code></td>      <td>string(22)</td>      <td>Texto exibido na fatura do cartão</td>    </tr>    <tr>      <td><code>shippable</code></td>      <td>boolean</td>      <td>Indica se o plano oferece entrega. O padrão é **false**</td>    </tr>    <tr>      <td><code>trial_period_days</code></td>      <td>integer</td>      <td>Dias para início das assinaturas</td>    </tr>    <tr>      <td><code>metadata</code></td>      <td>object</td>      <td>Informações adicionais sobre o plano. Leia mais sobre [metadata](#metadata)</td>    </tr>  </tbody></table><h5>Exemplo pré-pago</h5>```{    "name": "Premium",    "description": "Vá de Premium. E seja feliz!",    "statement_descriptor": "Spotify",    "billing_type": "prepaid",    "items": [        {            "name": "Premium",            "pricing_scheme": {                "price": 1490            }        }    ]}```<h5>Exemplo pós-pago</h5><blockquote><p><strong>Precificação</strong></p><p>Existem opções flexíveis para precificar seus planos. Leia mais sobre <a href="#pricing">precificação</a></p></blockquote>```{    "name": "SMS",    "description": "De pequenas a grandes empresas: o SMS traz mais resultados para o seu negócio",    "payment_methods": [ "credit_card", "boleto" ],    "statement_descriptor": "Operadora",    "billing_type": "postpaid",    "pricing_scheme": {        "scheme_type": "volume",        "price_brackets": [            {                "end_quantity": 100,                "price": 14            },            {                "start_quantity": 101,                "end_quantity": 1000,                "price": 12,                "overage_price": 9            }        ]    }}```<blockquote><p><strong>Simples</strong></p><p>Criar planos com apenas um item pode ser mais simples. Você pode enviar os campos <code>cycles</code>, <code>pricing_scheme</code> e <code>quantity</code> na raiz da requisição.</p></blockquote>```{    "name": "Premium",    "description": "Vá de Premium. E seja feliz!",    "pricing_scheme": {        "price": 1490    }}```


```go
func (me *PLANS_IMPL) CreatePlan(body *models_pkg.CreatePlanRequest)(*models_pkg.GetPlanResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreatePlanRequest

var result *models_pkg.GetPlanResponse
result,_ = plans.CreatePlan(body)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="customers_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".customers_pkg") customers_pkg

#### Get instance

Factory for the ``` CUSTOMERS ``` interface can be accessed from the package customers_pkg.

```go
customers := customers_pkg.NewCUSTOMERS()
```

#### <a name="update_credit_card"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.UpdateCreditCard") UpdateCreditCard

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) UpdateCreditCard(
            customerId string,
            cardId string,
            body *models_pkg.UpdateCreditCardRequest)(*models_pkg.GetCreditCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| cardId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"
var body *models_pkg.UpdateCreditCardRequest

var result *models_pkg.GetCreditCardResponse
result,_ = customers.UpdateCreditCard(customerId, cardId, body)

```


#### <a name="update_address"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.UpdateAddress") UpdateAddress

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) UpdateAddress(
            customerId string,
            addressId string,
            body *models_pkg.UpdateAddressRequest)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| addressId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"
var body *models_pkg.UpdateAddressRequest

var result *models_pkg.GetAddressResponse
result,_ = customers.UpdateAddress(customerId, addressId, body)

```


#### <a name="get_customer_orders"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCustomerOrders") GetCustomerOrders

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCustomerOrders(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.GetCustomerOrders(customerId)

```


#### <a name="get_customer"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCustomer") GetCustomer

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCustomer(customerId string)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result *models_pkg.GetCustomerResponse
result,_ = customers.GetCustomer(customerId)

```


#### <a name="create_customer"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.CreateCustomer") CreateCustomer

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) CreateCustomer(body *models_pkg.CreateCustomerRequest)(*models_pkg.GetCustomerResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateCustomerRequest

var result *models_pkg.GetCustomerResponse
result,_ = customers.CreateCustomer(body)

```


#### <a name="update_customer"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.UpdateCustomer") UpdateCustomer

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) UpdateCustomer(
            customerId string,
            body *models_pkg.UpdateCustomerRequest)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var body *models_pkg.UpdateCustomerRequest

var result 
result,_ = customers.UpdateCustomer(customerId, body)

```


#### <a name="list_customers"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.ListCustomers") ListCustomers

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) ListCustomers()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = customers.ListCustomers()

```


#### <a name="get_customer_orders1"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCustomerOrders1") GetCustomerOrders1

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCustomerOrders1(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.GetCustomerOrders1(customerId)

```


#### <a name="create_credit_card"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.CreateCreditCard") CreateCreditCard

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) CreateCreditCard(
            body *models_pkg.CreateCreditCardRequest,
            customerId string)(*models_pkg.GetCreditCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateCreditCardRequest
customerId := "customer_id"

var result *models_pkg.GetCreditCardResponse
result,_ = customers.CreateCreditCard(body, customerId)

```


#### <a name="list_credit_cards"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.ListCreditCards") ListCreditCards

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) ListCreditCards(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.ListCreditCards(customerId)

```


#### <a name="get_credit_card"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCreditCard") GetCreditCard

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCreditCard(
            customerId string,
            cardId string)(*models_pkg.GetCreditCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| cardId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"

var result *models_pkg.GetCreditCardResponse
result,_ = customers.GetCreditCard(customerId, cardId)

```


#### <a name="create_address"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.CreateAddress") CreateAddress

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) CreateAddress(
            customerId string,
            body *models_pkg.CreateAddressRequest)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
var body *models_pkg.CreateAddressRequest

var result *models_pkg.GetAddressResponse
result,_ = customers.CreateAddress(customerId, body)

```


#### <a name="delete_credit_card"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.DeleteCreditCard") DeleteCreditCard

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) DeleteCreditCard(
            customerId string,
            cardId string)(*models_pkg.GetCreditCardResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| cardId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
cardId := "card_id"

var result *models_pkg.GetCreditCardResponse
result,_ = customers.DeleteCreditCard(customerId, cardId)

```


#### <a name="delete_address"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.DeleteAddress") DeleteAddress

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) DeleteAddress(
            customerId string,
            addressId string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| addressId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"

var result *models_pkg.GetAddressResponse
result,_ = customers.DeleteAddress(customerId, addressId)

```


#### <a name="get_customer_orders2"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCustomerOrders2") GetCustomerOrders2

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCustomerOrders2(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.GetCustomerOrders2(customerId)

```


#### <a name="list_addresses"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.ListAddresses") ListAddresses

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) ListAddresses(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.ListAddresses(customerId)

```


#### <a name="get_address"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetAddress") GetAddress

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetAddress(
            customerId string,
            addressId string)(*models_pkg.GetAddressResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |
| addressId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"
addressId := "address_id"

var result *models_pkg.GetAddressResponse
result,_ = customers.GetAddress(customerId, addressId)

```


#### <a name="get_customer_invoices"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCustomerInvoices") GetCustomerInvoices

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCustomerInvoices(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.GetCustomerInvoices(customerId)

```


#### <a name="get_customer_subscriptions"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCustomerSubscriptions") GetCustomerSubscriptions

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCustomerSubscriptions(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.GetCustomerSubscriptions(customerId)

```


#### <a name="get_customer_charges"></a>![Method: ](http://apidocs.io/img/method.png ".customers_pkg.GetCustomerCharges") GetCustomerCharges

> TODO: Add a method description


```go
func (me *CUSTOMERS_IMPL) GetCustomerCharges(customerId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| customerId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
customerId := "customer_id"

var result interface{}
result,_ = customers.GetCustomerCharges(customerId)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="charges_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".charges_pkg") charges_pkg

#### Get instance

Factory for the ``` CHARGES ``` interface can be accessed from the package charges_pkg.

```go
charges := charges_pkg.NewCHARGES()
```

#### <a name="update_charge_payment_method"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.UpdateChargePaymentMethod") UpdateChargePaymentMethod

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) UpdateChargePaymentMethod(
            body *models_pkg.UpdateChargePaymentMethodRequest,
            chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| chargeId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.UpdateChargePaymentMethodRequest
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargePaymentMethod(body, chargeId)

```


#### <a name="create_capture_charge"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.CreateCaptureCharge") CreateCaptureCharge

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) CreateCaptureCharge(
            chargeId string,
            body *models_pkg.CreateCaptureChargeRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var body *models_pkg.CreateCaptureChargeRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.CreateCaptureCharge(chargeId, body)

```


#### <a name="create_charge"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.CreateCharge") CreateCharge

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) CreateCharge(body *models_pkg.CreateChargeRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateChargeRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.CreateCharge(body)

```


#### <a name="list_charges"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.ListCharges") ListCharges

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) ListCharges()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = charges.ListCharges()

```


#### <a name="delete_cancel_charge"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.DeleteCancelCharge") DeleteCancelCharge

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) DeleteCancelCharge(
            body *models_pkg.CreateCancelChargeRequest,
            chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| chargeId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.CreateCancelChargeRequest
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.DeleteCancelCharge(body, chargeId)

```


#### <a name="update_charge_credit_card"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.UpdateChargeCreditCard") UpdateChargeCreditCard

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) UpdateChargeCreditCard(
            body *models_pkg.UpdateChargeCreditCardRequest,
            chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| chargeId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.UpdateChargeCreditCardRequest
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeCreditCard(body, chargeId)

```


#### <a name="update_charge_due_date"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.UpdateChargeDueDate") UpdateChargeDueDate

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) UpdateChargeDueDate(
            chargeId string,
            body *models_pkg.UpdateChargeDueDateRequest)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | TODO: Add a parameter description |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"
var body *models_pkg.UpdateChargeDueDateRequest

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeDueDate(chargeId, body)

```


#### <a name="create_retry_charge"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.CreateRetryCharge") CreateRetryCharge

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) CreateRetryCharge(chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.CreateRetryCharge(chargeId)

```


#### <a name="update_charge_credit_card1"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.UpdateChargeCreditCard1") UpdateChargeCreditCard1

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) UpdateChargeCreditCard1(
            body *models_pkg.UpdateChargeCreditCardRequest,
            chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| body |  ``` Required ```  | TODO: Add a parameter description |
| chargeId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
var body *models_pkg.UpdateChargeCreditCardRequest
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.UpdateChargeCreditCard1(body, chargeId)

```


#### <a name="get_charge_transactions"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.GetChargeTransactions") GetChargeTransactions

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) GetChargeTransactions(chargeId string)(interface{},error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"

var result interface{}
result,_ = charges.GetChargeTransactions(chargeId)

```


#### <a name="get_charge"></a>![Method: ](http://apidocs.io/img/method.png ".charges_pkg.GetCharge") GetCharge

> TODO: Add a method description


```go
func (me *CHARGES_IMPL) GetCharge(chargeId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| chargeId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
chargeId := "charge_id"

var result *models_pkg.GetChargeResponse
result,_ = charges.GetCharge(chargeId)

```


[Back to List of Controllers](#list_of_controllers)

### <a name="invoices_pkg"></a>![Class: ](http://apidocs.io/img/class.png ".invoices_pkg") invoices_pkg

#### Get instance

Factory for the ``` INVOICES ``` interface can be accessed from the package invoices_pkg.

```go
invoices := invoices_pkg.NewINVOICES()
```

#### <a name="get_invoice"></a>![Method: ](http://apidocs.io/img/method.png ".invoices_pkg.GetInvoice") GetInvoice

> TODO: Add a method description


```go
func (me *INVOICES_IMPL) GetInvoice(invoiceId string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.GetInvoice(invoiceId)

```


#### <a name="list_invoices"></a>![Method: ](http://apidocs.io/img/method.png ".invoices_pkg.ListInvoices") ListInvoices

> TODO: Add a method description


```go
func (me *INVOICES_IMPL) ListInvoices()(interface{},error)
```

#### Example Usage

```go

var result interface{}
result,_ = invoices.ListInvoices()

```


#### <a name="get_last_invoice_charge"></a>![Method: ](http://apidocs.io/img/method.png ".invoices_pkg.GetLastInvoiceCharge") GetLastInvoiceCharge

> TODO: Add a method description


```go
func (me *INVOICES_IMPL) GetLastInvoiceCharge(invoiceId string)(*models_pkg.GetChargeResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.GetChargeResponse
result,_ = invoices.GetLastInvoiceCharge(invoiceId)

```


#### <a name="delete_cancel_invoice"></a>![Method: ](http://apidocs.io/img/method.png ".invoices_pkg.DeleteCancelInvoice") DeleteCancelInvoice

> TODO: Add a method description


```go
func (me *INVOICES_IMPL) DeleteCancelInvoice(invoiceId string)(*models_pkg.GetInvoiceResponse,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| invoiceId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
invoiceId := "invoice_id"

var result *models_pkg.GetInvoiceResponse
result,_ = invoices.DeleteCancelInvoice(invoiceId)

```


[Back to List of Controllers](#list_of_controllers)



