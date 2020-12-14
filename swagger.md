# AELF API 1.0

<a name="overview"></a>
## Overview

### Version information
*Version* : 1.0

<a name="/api/blockChain/block"></a>
## Get information about a given block by block hash. Otionally with the list of its transactions.

```
get /api/blockChain/block
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**blockHash**  (*optional*)|block hash|string||
|**query**|**includeTransactions**  |include transactions or not|boolean|false|






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[BlockDto](#BlockDto)|





### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/blockByHeight"></a>
## Get information about a given block by block height. Optionally with the list of its transactions.

```
get /api/blockChain/blockByHeight
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**blockHeight**  |block height|integer||
|**query**|**includeTransactions**  |include transactions or not|boolean|false|






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[BlockDto](#BlockDto)|





### Produces

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/blockHeight"></a>
## Get the height of the current chain.

```
get /api/blockChain/blockHeight
```






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|integer|





### Produces

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/blockState"></a>
## Get the current state about a given block

```
get /api/blockChain/blockState
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**blockHash**  (*optional*)|block hash|string||






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[BlockStateDto](#BlockStateDto)|





### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/chainStatus"></a>
## Get the current status of the block chain.

```
get /api/blockChain/chainStatus
```






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[ChainStatusDto](#ChainStatusDto)|





### Produces

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/contractFileDescriptorSet"></a>
## Get the protobuf definitions related to a contract

```
get /api/blockChain/contractFileDescriptorSet
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**address**  (*optional*)|contract address|string||






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|string|





### Produces

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/executeRawTransaction"></a>
## 

```
post /api/blockChain/executeRawTransaction
```




### RequestBody
|Type|Description|Schema|
|---|---|---|
|**Body**||[ExecuteRawTransactionDto](#ExecuteRawTransactionDto)|



### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|string|



### Consumes

* application/json-patch&#43;json; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/*&#43;json; v=1.0

* application/x-protobuf; v=1.0




### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/executeTransaction"></a>
## Call a read-only method on a contract.

```
post /api/blockChain/executeTransaction
```




### RequestBody
|Type|Description|Schema|
|---|---|---|
|**Body**||[ExecuteTransactionDto](#ExecuteTransactionDto)|



### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|string|



### Consumes

* application/json-patch&#43;json; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/*&#43;json; v=1.0

* application/x-protobuf; v=1.0




### Produces

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0

* application/json; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/merklePathByTransactionId"></a>
## Get the merkle path of a transaction.

```
get /api/blockChain/merklePathByTransactionId
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**transactionId**  (*optional*)||string||






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[MerklePathDto](#MerklePathDto)|





### Produces

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0

* application/json; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/rawTransaction"></a>
## Creates an unsigned serialized transaction

```
post /api/blockChain/rawTransaction
```




### RequestBody
|Type|Description|Schema|
|---|---|---|
|**Body**||[CreateRawTransactionInput](#CreateRawTransactionInput)|



### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[CreateRawTransactionOutput](#CreateRawTransactionOutput)|



### Consumes

* application/json-patch&#43;json; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/*&#43;json; v=1.0

* application/x-protobuf; v=1.0




### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/sendRawTransaction"></a>
## send a transaction

```
post /api/blockChain/sendRawTransaction
```




### RequestBody
|Type|Description|Schema|
|---|---|---|
|**Body**||[SendRawTransactionInput](#SendRawTransactionInput)|



### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[SendRawTransactionOutput](#SendRawTransactionOutput)|



### Consumes

* application/json-patch&#43;json; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/*&#43;json; v=1.0

* application/x-protobuf; v=1.0




### Produces

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0

* application/json; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/sendTransaction"></a>
## Broadcast a transaction

```
post /api/blockChain/sendTransaction
```




### RequestBody
|Type|Description|Schema|
|---|---|---|
|**Body**||[SendTransactionInput](#SendTransactionInput)|



### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[SendTransactionOutput](#SendTransactionOutput)|



### Consumes

* text/json; v=1.0

* application/*&#43;json; v=1.0

* application/x-protobuf; v=1.0

* application/json-patch&#43;json; v=1.0

* application/json; v=1.0




### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/sendTransactions"></a>
## Broadcast multiple transactions

```
post /api/blockChain/sendTransactions
```




### RequestBody
|Type|Description|Schema|
|---|---|---|
|**Body**||[SendTransactionsInput](#SendTransactionsInput)|



### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|array&lt;string>|



### Consumes

* application/json-patch&#43;json; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/*&#43;json; v=1.0

* application/x-protobuf; v=1.0




### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/taskQueueStatus"></a>
## 

```
get /api/blockChain/taskQueueStatus
```






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|array&lt;[TaskQueueInfoDto](#TaskQueueInfoDto)>|





### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/transactionPoolStatus"></a>
## Get the transaction pool status.

```
get /api/blockChain/transactionPoolStatus
```






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[GetTransactionPoolStatusOutput](#GetTransactionPoolStatusOutput)|





### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/transactionResult"></a>
## Get the current status of a transaction

```
get /api/blockChain/transactionResult
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**transactionId**  (*optional*)|transaction id|string||






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[TransactionResultDto](#TransactionResultDto)|





### Produces

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0




### Tags

* BlockChain



<a name="/api/blockChain/transactionResults"></a>
## Get multiple transaction results.

```
get /api/blockChain/transactionResults
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**blockHash**  (*optional*)|block hash|string||
|**query**|**offset**  |offset|integer|0|
|**query**|**limit**  |limit|integer|10|






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|array&lt;[TransactionResultDto](#TransactionResultDto)>|





### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* BlockChain



<a name="/api/net/networkInfo"></a>
## Get information about the node’s connection to the network.

```
get /api/net/networkInfo
```






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|[GetNetworkInfoOutput](#GetNetworkInfoOutput)|





### Produces

* text/json; v=1.0

* application/x-protobuf; v=1.0

* text/plain; v=1.0

* application/json; v=1.0




### Tags

* Net



<a name="/api/net/peer"></a>
## Attempts to remove a node from the connected network nodes

```
delete /api/net/peer
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**address**  (*optional*)|ip address|string||






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|boolean|





### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* Net



<a name="/api/net/peer"></a>
## Attempts to add a node to the connected network nodes

```
post /api/net/peer
```




### RequestBody
|Type|Description|Schema|
|---|---|---|
|**Body**||[AddPeerInput](#AddPeerInput)|



### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|boolean|



### Consumes

* application/json-patch&#43;json; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/*&#43;json; v=1.0

* application/x-protobuf; v=1.0




### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* Net



<a name="/api/net/peers"></a>
## Get peer info about the connected network nodes

```
get /api/net/peers
```


### Parameters
|Type|Name|Description|Schema|Default|
|---|---|---|---|---|
|**query**|**withMetrics**  ||boolean|false|






### Responses
|HTTP Code|Description|Schema|
|---|---|---|
|**200**|Success|array&lt;[PeerDto](#PeerDto)>|





### Produces

* text/plain; v=1.0

* application/json; v=1.0

* text/json; v=1.0

* application/x-protobuf; v=1.0




### Tags

* Net





<a name="ActionApiDescriptionModel"></a>

## ActionApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|name||string|
|httpMethod||string|
|url||string|
|supportedVersions||array&lt;string>|
|parametersOnMethod||array&lt;[MethodParameterApiDescriptionModel](#MethodParameterApiDescriptionModel)>|
|parameters||array&lt;[ParameterApiDescriptionModel](#ParameterApiDescriptionModel)>|
|returnValue||[ReturnValueApiDescriptionModel](#ReturnValueApiDescriptionModel)|
|uniqueName||string|


<a name="AddPeerInput"></a>

## AddPeerInput

|Name|Description|Schema|
|---|---|---|
|address|ip address|string|


<a name="ApplicationApiDescriptionModel"></a>

## ApplicationApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|types||object|
|modules||object|


<a name="ApplicationAuthConfigurationDto"></a>

## ApplicationAuthConfigurationDto

|Name|Description|Schema|
|---|---|---|
|grantedPolicies||object|
|policies||object|


<a name="ApplicationConfigurationDto"></a>

## ApplicationConfigurationDto

|Name|Description|Schema|
|---|---|---|
|auth||[ApplicationAuthConfigurationDto](#ApplicationAuthConfigurationDto)|
|setting||[ApplicationSettingConfigurationDto](#ApplicationSettingConfigurationDto)|
|features||[ApplicationFeatureConfigurationDto](#ApplicationFeatureConfigurationDto)|
|currentTenant||[CurrentTenantDto](#CurrentTenantDto)|
|clock||[ClockDto](#ClockDto)|
|localization||[ApplicationLocalizationConfigurationDto](#ApplicationLocalizationConfigurationDto)|
|currentUser||[CurrentUserDto](#CurrentUserDto)|
|multiTenancy||[MultiTenancyInfoDto](#MultiTenancyInfoDto)|
|timing||[TimingDto](#TimingDto)|
|objectExtensions||[ObjectExtensionsDto](#ObjectExtensionsDto)|


<a name="ApplicationFeatureConfigurationDto"></a>

## ApplicationFeatureConfigurationDto

|Name|Description|Schema|
|---|---|---|
|values||object|


<a name="ApplicationLocalizationConfigurationDto"></a>

## ApplicationLocalizationConfigurationDto

|Name|Description|Schema|
|---|---|---|
|values||object|
|languages||array&lt;[LanguageInfo](#LanguageInfo)>|
|currentCulture||[CurrentCultureDto](#CurrentCultureDto)|
|defaultResourceName||string|
|languagesMap||object|
|languageFilesMap||object|


<a name="ApplicationSettingConfigurationDto"></a>

## ApplicationSettingConfigurationDto

|Name|Description|Schema|
|---|---|---|
|values||object|


<a name="BlockBodyDto"></a>

## BlockBodyDto

|Name|Description|Schema|
|---|---|---|
|transactionsCount||integer|
|transactions||array&lt;string>|


<a name="BlockDto"></a>

## BlockDto

|Name|Description|Schema|
|---|---|---|
|blockSize||integer|
|blockHash||string|
|header||[BlockHeaderDto](#BlockHeaderDto)|
|body||[BlockBodyDto](#BlockBodyDto)|


<a name="BlockHeaderDto"></a>

## BlockHeaderDto

|Name|Description|Schema|
|---|---|---|
|merkleTreeRootOfTransactionState||string|
|height||integer|
|time||string|
|bloom||string|
|chainId||string|
|signerPubkey||string|
|previousBlockHash||string|
|merkleTreeRootOfTransactions||string|
|merkleTreeRootOfWorldState||string|
|extra||string|


<a name="BlockStateDto"></a>

## BlockStateDto

|Name|Description|Schema|
|---|---|---|
|blockHash||string|
|previousHash||string|
|blockHeight||integer|
|changes||object|
|deletes||array&lt;string>|


<a name="ChainStatusDto"></a>

## ChainStatusDto

|Name|Description|Schema|
|---|---|---|
|branches||object|
|notLinkedBlocks||object|
|longestChainHeight||integer|
|longestChainHash||string|
|genesisContractAddress||string|
|chainId||string|
|genesisBlockHash||string|
|lastIrreversibleBlockHash||string|
|lastIrreversibleBlockHeight||integer|
|bestChainHash||string|
|bestChainHeight||integer|


<a name="ClockDto"></a>

## ClockDto

|Name|Description|Schema|
|---|---|---|
|kind||string|


<a name="ControllerApiDescriptionModel"></a>

## ControllerApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|controllerName||string|
|type||string|
|interfaces||array&lt;[ControllerInterfaceApiDescriptionModel](#ControllerInterfaceApiDescriptionModel)>|
|actions||object|


<a name="ControllerInterfaceApiDescriptionModel"></a>

## ControllerInterfaceApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|type||string|


<a name="CreateRawTransactionInput"></a>

## CreateRawTransactionInput

|Name|Description|Schema|
|---|---|---|
|methodName|contract method name|string|
|params|contract method parameters|string|
|from|from address|string|
|to|to address|string|
|refBlockNumber|refer block height|integer|
|refBlockHash|refer block hash|string|


<a name="CreateRawTransactionOutput"></a>

## CreateRawTransactionOutput

|Name|Description|Schema|
|---|---|---|
|rawTransaction||string|


<a name="CurrentCultureDto"></a>

## CurrentCultureDto

|Name|Description|Schema|
|---|---|---|
|nativeName||string|
|displayName||string|
|englishName||string|
|threeLetterIsoLanguageName||string|
|isRightToLeft||boolean|
|name||string|
|twoLetterIsoLanguageName||string|
|cultureName||string|
|dateTimeFormat||[DateTimeFormatDto](#DateTimeFormatDto)|


<a name="CurrentTenantDto"></a>

## CurrentTenantDto

|Name|Description|Schema|
|---|---|---|
|id||string|
|name||string|
|isAvailable||boolean|


<a name="CurrentUserDto"></a>

## CurrentUserDto

|Name|Description|Schema|
|---|---|---|
|roles||array&lt;string>|
|isAuthenticated||boolean|
|id||string|
|tenantId||string|
|userName||string|
|email||string|


<a name="DateTimeFormatDto"></a>

## DateTimeFormatDto

|Name|Description|Schema|
|---|---|---|
|calendarAlgorithmType||string|
|dateTimeFormatLong||string|
|shortDatePattern||string|
|fullDateTimePattern||string|
|dateSeparator||string|
|shortTimePattern||string|
|longTimePattern||string|


<a name="EntityExtensionDto"></a>

## EntityExtensionDto

|Name|Description|Schema|
|---|---|---|
|properties||object|
|configuration||object|


<a name="ExecuteRawTransactionDto"></a>

## ExecuteRawTransactionDto

|Name|Description|Schema|
|---|---|---|
|rawTransaction|raw transaction|string|
|signature|signature|string|


<a name="ExecuteTransactionDto"></a>

## ExecuteTransactionDto

|Name|Description|Schema|
|---|---|---|
|rawTransaction|raw transaction|string|


<a name="ExtensionEnumDto"></a>

## ExtensionEnumDto

|Name|Description|Schema|
|---|---|---|
|localizationResource||string|
|fields||array&lt;[ExtensionEnumFieldDto](#ExtensionEnumFieldDto)>|


<a name="ExtensionEnumFieldDto"></a>

## ExtensionEnumFieldDto

|Name|Description|Schema|
|---|---|---|
|value||object|
|name||string|


<a name="ExtensionPropertyApiCreateDto"></a>

## ExtensionPropertyApiCreateDto

|Name|Description|Schema|
|---|---|---|
|isAvailable||boolean|


<a name="ExtensionPropertyApiDto"></a>

## ExtensionPropertyApiDto

|Name|Description|Schema|
|---|---|---|
|onUpdate||[ExtensionPropertyApiUpdateDto](#ExtensionPropertyApiUpdateDto)|
|onGet||[ExtensionPropertyApiGetDto](#ExtensionPropertyApiGetDto)|
|onCreate||[ExtensionPropertyApiCreateDto](#ExtensionPropertyApiCreateDto)|


<a name="ExtensionPropertyApiGetDto"></a>

## ExtensionPropertyApiGetDto

|Name|Description|Schema|
|---|---|---|
|isAvailable||boolean|


<a name="ExtensionPropertyApiUpdateDto"></a>

## ExtensionPropertyApiUpdateDto

|Name|Description|Schema|
|---|---|---|
|isAvailable||boolean|


<a name="ExtensionPropertyAttributeDto"></a>

## ExtensionPropertyAttributeDto

|Name|Description|Schema|
|---|---|---|
|typeSimple||string|
|config||object|


<a name="ExtensionPropertyDto"></a>

## ExtensionPropertyDto

|Name|Description|Schema|
|---|---|---|
|defaultValue||object|
|type||string|
|typeSimple||string|
|displayName||[LocalizableStringDto](#LocalizableStringDto)|
|api||[ExtensionPropertyApiDto](#ExtensionPropertyApiDto)|
|ui||[ExtensionPropertyUiDto](#ExtensionPropertyUiDto)|
|attributes||array&lt;[ExtensionPropertyAttributeDto](#ExtensionPropertyAttributeDto)>|
|configuration||object|


<a name="ExtensionPropertyUiDto"></a>

## ExtensionPropertyUiDto

|Name|Description|Schema|
|---|---|---|
|onTable||[ExtensionPropertyUiTableDto](#ExtensionPropertyUiTableDto)|
|onCreateForm||[ExtensionPropertyUiFormDto](#ExtensionPropertyUiFormDto)|
|onEditForm||[ExtensionPropertyUiFormDto](#ExtensionPropertyUiFormDto)|


<a name="ExtensionPropertyUiFormDto"></a>

## ExtensionPropertyUiFormDto

|Name|Description|Schema|
|---|---|---|
|isVisible||boolean|


<a name="ExtensionPropertyUiTableDto"></a>

## ExtensionPropertyUiTableDto

|Name|Description|Schema|
|---|---|---|
|isVisible||boolean|


<a name="GetNetworkInfoOutput"></a>

## GetNetworkInfoOutput

|Name|Description|Schema|
|---|---|---|
|version|node version|string|
|protocolVersion|network protocol version|integer|
|connections|total number of open connections between this node and other nodes|integer|


<a name="GetTransactionPoolStatusOutput"></a>

## GetTransactionPoolStatusOutput

|Name|Description|Schema|
|---|---|---|
|queued||integer|
|validated||integer|


<a name="IanaTimeZone"></a>

## IanaTimeZone

|Name|Description|Schema|
|---|---|---|
|timeZoneName||string|


<a name="LanguageInfo"></a>

## LanguageInfo

|Name|Description|Schema|
|---|---|---|
|cultureName||string|
|uiCultureName||string|
|displayName||string|
|flagIcon||string|


<a name="LocalizableStringDto"></a>

## LocalizableStringDto

|Name|Description|Schema|
|---|---|---|
|name||string|
|resource||string|


<a name="LogEventDto"></a>

## LogEventDto

|Name|Description|Schema|
|---|---|---|
|address||string|
|name||string|
|indexed||array&lt;string>|
|nonIndexed||string|


<a name="MerklePathDto"></a>

## MerklePathDto

|Name|Description|Schema|
|---|---|---|


<a name="MethodParameterApiDescriptionModel"></a>

## MethodParameterApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|name||string|
|typeAsString||string|
|type||string|
|typeSimple||string|
|isOptional||boolean|
|defaultValue||object|


<a name="ModuleApiDescriptionModel"></a>

## ModuleApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|rootPath||string|
|remoteServiceName||string|
|controllers||object|


<a name="ModuleExtensionDto"></a>

## ModuleExtensionDto

|Name|Description|Schema|
|---|---|---|
|entities||object|
|configuration||object|


<a name="MultiTenancyInfoDto"></a>

## MultiTenancyInfoDto

|Name|Description|Schema|
|---|---|---|
|isEnabled||boolean|


<a name="NameValue"></a>

## NameValue

|Name|Description|Schema|
|---|---|---|
|name||string|
|value||string|


<a name="ObjectExtensionsDto"></a>

## ObjectExtensionsDto

|Name|Description|Schema|
|---|---|---|
|modules||object|
|enums||object|


<a name="ParameterApiDescriptionModel"></a>

## ParameterApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|type||string|
|isOptional||boolean|
|descriptorName||string|
|nameOnMethod||string|
|name||string|
|typeSimple||string|
|defaultValue||object|
|constraintTypes||array&lt;string>|
|bindingSourceId||string|


<a name="PeerDto"></a>

## PeerDto

|Name|Description|Schema|
|---|---|---|
|connectionTime||integer|
|connectionStatus||string|
|bufferedBlocksCount||integer|
|ipAddress||string|
|protocolVersion||integer|
|inbound||boolean|
|bufferedTransactionsCount||integer|
|bufferedAnnouncementsCount||integer|
|requestMetrics||array&lt;[RequestMetric](#RequestMetric)>|


<a name="PropertyApiDescriptionModel"></a>

## PropertyApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|name||string|
|type||string|
|typeSimple||string|


<a name="RemoteServiceErrorInfo"></a>

## RemoteServiceErrorInfo

|Name|Description|Schema|
|---|---|---|
|code||string|
|message||string|
|details||string|
|validationErrors||array&lt;[RemoteServiceValidationErrorInfo](#RemoteServiceValidationErrorInfo)>|


<a name="RemoteServiceErrorResponse"></a>

## RemoteServiceErrorResponse

|Name|Description|Schema|
|---|---|---|
|error||[RemoteServiceErrorInfo](#RemoteServiceErrorInfo)|


<a name="RemoteServiceValidationErrorInfo"></a>

## RemoteServiceValidationErrorInfo

|Name|Description|Schema|
|---|---|---|
|message||string|
|members||array&lt;string>|


<a name="RequestMetric"></a>

## RequestMetric

|Name|Description|Schema|
|---|---|---|
|roundTripTime||integer|
|methodName||string|
|info||string|
|requestTime||[Timestamp](#Timestamp)|


<a name="ReturnValueApiDescriptionModel"></a>

## ReturnValueApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|typeSimple||string|
|type||string|


<a name="SendRawTransactionInput"></a>

## SendRawTransactionInput

|Name|Description|Schema|
|---|---|---|
|transaction|raw transaction|string|
|signature|signature|string|
|returnTransaction|return transaction detail or not|boolean|


<a name="SendRawTransactionOutput"></a>

## SendRawTransactionOutput

|Name|Description|Schema|
|---|---|---|
|transactionId||string|
|transaction||[TransactionDto](#TransactionDto)|


<a name="SendTransactionInput"></a>

## SendTransactionInput

|Name|Description|Schema|
|---|---|---|
|rawTransaction|raw transaction|string|


<a name="SendTransactionOutput"></a>

## SendTransactionOutput

|Name|Description|Schema|
|---|---|---|
|transactionId||string|


<a name="SendTransactionsInput"></a>

## SendTransactionsInput

|Name|Description|Schema|
|---|---|---|
|rawTransactions|raw transactions|string|


<a name="TaskQueueInfoDto"></a>

## TaskQueueInfoDto

|Name|Description|Schema|
|---|---|---|
|name||string|
|size||integer|


<a name="TimeZone"></a>

## TimeZone

|Name|Description|Schema|
|---|---|---|
|iana||[IanaTimeZone](#IanaTimeZone)|
|windows||[WindowsTimeZone](#WindowsTimeZone)|


<a name="Timestamp"></a>

## Timestamp

|Name|Description|Schema|
|---|---|---|
|nanos||integer|
|seconds||integer|


<a name="TimingDto"></a>

## TimingDto

|Name|Description|Schema|
|---|---|---|
|timeZone||[TimeZone](#TimeZone)|


<a name="TransactionDto"></a>

## TransactionDto

|Name|Description|Schema|
|---|---|---|
|signature||string|
|from||string|
|to||string|
|refBlockNumber||integer|
|refBlockPrefix||string|
|methodName||string|
|params||string|


<a name="TransactionResultDto"></a>

## TransactionResultDto

|Name|Description|Schema|
|---|---|---|
|status||string|
|transactionId||string|
|bloom||string|
|blockNumber||integer|
|blockHash||string|
|transaction||[TransactionDto](#TransactionDto)|
|returnValue||string|
|error||string|
|transactionSize||integer|
|logs||array&lt;[LogEventDto](#LogEventDto)>|


<a name="TypeApiDescriptionModel"></a>

## TypeApiDescriptionModel

|Name|Description|Schema|
|---|---|---|
|enumValues||array&lt;object>|
|genericArguments||array&lt;string>|
|properties||array&lt;[PropertyApiDescriptionModel](#PropertyApiDescriptionModel)>|
|baseType||string|
|isEnum||boolean|
|enumNames||array&lt;string>|


<a name="WindowsTimeZone"></a>

## WindowsTimeZone

|Name|Description|Schema|
|---|---|---|
|timeZoneId||string|


