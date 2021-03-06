Generated from https://github.com/Azure/azure-rest-api-specs/tree/d4bad535d456ee82c7fd17d1ec2b8802a0b83205/specification/cosmos-db/resource-manager/readme.md tag: `package-2020-06-preview`

Code generator @microsoft.azure/autorest.go@2.1.168

## Breaking Changes

- Const `Delete` type has been changed from `TriggerOperation` to `OperationType`
- Const `Create` type has been changed from `TriggerOperation` to `OperationType`
- Const `MongoDB` type has been changed from `DatabaseAccountKind` to `APIType`
- Const `Replace` type has been changed from `TriggerOperation` to `OperationType`
- Const `Parse` has been removed
- Const `GlobalDocumentDB` has been removed
- Const `All` has been removed
- Const `Update` has been removed

## New Content

- New const `TriggerOperationUpdate`
- New const `Table`
- New const `TriggerOperationDelete`
- New const `TriggerOperationReplace`
- New const `DatabaseAccountKindParse`
- New const `Cassandra`
- New const `DatabaseAccountKindGlobalDocumentDB`
- New const `DatabaseAccountKindMongoDB`
- New const `TriggerOperationCreate`
- New const `GremlinV2`
- New const `Gremlin`
- New const `SystemOperation`
- New const `TriggerOperationAll`
- New const `SQL`
- New function `RestorableMongodbResourcesClient.ListPreparer(context.Context, string, string, string, string) (*http.Request, error)`
- New function `NewRestorableSQLDatabasesClientWithBaseURI(string, string) RestorableSQLDatabasesClient`
- New function `NewRestorableSQLContainersClient(string) RestorableSQLContainersClient`
- New function `RestorableMongodbDatabasesClient.List(context.Context, string, string) (RestorableMongodbDatabasesListResult, error)`
- New function `RestorableSQLContainersClient.ListResponder(*http.Response) (RestorableSQLContainersListResult, error)`
- New function `RestorableMongodbCollectionsClient.List(context.Context, string, string, string) (RestorableMongodbCollectionsListResult, error)`
- New function `RestorableSQLDatabaseGetResult.MarshalJSON() ([]byte, error)`
- New function `RestorableMongodbResourcesClient.ListResponder(*http.Response) (RestorableMongodbResourcesListResult, error)`
- New function `RestorableMongodbResourcesClient.ListSender(*http.Request) (*http.Response, error)`
- New function `RestorableSQLResourcesClient.ListPreparer(context.Context, string, string, string, string) (*http.Request, error)`
- New function `RestorableDatabaseAccountProperties.MarshalJSON() ([]byte, error)`
- New function `NewRestorableMongodbCollectionsClient(string) RestorableMongodbCollectionsClient`
- New function `RestorableSQLResourcesClient.ListSender(*http.Request) (*http.Response, error)`
- New function `RestorableMongodbCollectionsClient.ListResponder(*http.Response) (RestorableMongodbCollectionsListResult, error)`
- New function `RestorableMongodbDatabaseGetResult.MarshalJSON() ([]byte, error)`
- New function `NewRestorableMongodbCollectionsClientWithBaseURI(string, string) RestorableMongodbCollectionsClient`
- New function `RestorableSQLResourcesClient.ListResponder(*http.Response) (RestorableSQLResourcesListResult, error)`
- New function `NewRestorableMongodbResourcesClientWithBaseURI(string, string) RestorableMongodbResourcesClient`
- New function `RestorableMongodbCollectionsClient.ListPreparer(context.Context, string, string, string) (*http.Request, error)`
- New function `NewRestorableSQLResourcesClientWithBaseURI(string, string) RestorableSQLResourcesClient`
- New function `RestorableMongodbDatabasesClient.ListSender(*http.Request) (*http.Response, error)`
- New function `NewRestorableSQLResourcesClient(string) RestorableSQLResourcesClient`
- New function `RestorableSQLDatabasesClient.ListResponder(*http.Response) (RestorableSQLDatabasesListResult, error)`
- New function `RestorableMongodbDatabasesClient.ListPreparer(context.Context, string, string) (*http.Request, error)`
- New function `NewRestorableMongodbDatabasesClient(string) RestorableMongodbDatabasesClient`
- New function `NewRestorableSQLContainersClientWithBaseURI(string, string) RestorableSQLContainersClient`
- New function `RestorableSQLResourcesClient.List(context.Context, string, string, string, string) (RestorableSQLResourcesListResult, error)`
- New function `NewRestorableMongodbDatabasesClientWithBaseURI(string, string) RestorableMongodbDatabasesClient`
- New function `RestorableSQLDatabasePropertiesResource.MarshalJSON() ([]byte, error)`
- New function `RestorableSQLDatabasePropertiesResourceDatabase.MarshalJSON() ([]byte, error)`
- New function `*RestorableMongodbDatabaseGetResult.UnmarshalJSON([]byte) error`
- New function `RestorableSQLContainerPropertiesResource.MarshalJSON() ([]byte, error)`
- New function `RestorableMongodbCollectionGetResult.MarshalJSON() ([]byte, error)`
- New function `RestorableSQLContainerPropertiesResourceContainer.MarshalJSON() ([]byte, error)`
- New function `RestorableMongodbCollectionsClient.ListSender(*http.Request) (*http.Response, error)`
- New function `RestorableMongodbResourcesClient.List(context.Context, string, string, string, string) (RestorableMongodbResourcesListResult, error)`
- New function `RestorableSQLDatabasesClient.List(context.Context, string, string) (RestorableSQLDatabasesListResult, error)`
- New function `PossibleAPITypeValues() []APIType`
- New function `*RestorableSQLDatabaseGetResult.UnmarshalJSON([]byte) error`
- New function `RestorableSQLContainersClient.ListSender(*http.Request) (*http.Response, error)`
- New function `RestorableSQLContainersClient.List(context.Context, string, string, string) (RestorableSQLContainersListResult, error)`
- New function `NewRestorableSQLDatabasesClient(string) RestorableSQLDatabasesClient`
- New function `RestorableMongodbDatabasesClient.ListResponder(*http.Response) (RestorableMongodbDatabasesListResult, error)`
- New function `RestorableSQLContainerGetResult.MarshalJSON() ([]byte, error)`
- New function `NewRestorableMongodbResourcesClient(string) RestorableMongodbResourcesClient`
- New function `PossibleOperationTypeValues() []OperationType`
- New function `RestorableSQLDatabasesClient.ListSender(*http.Request) (*http.Response, error)`
- New function `RestorableSQLContainersClient.ListPreparer(context.Context, string, string, string) (*http.Request, error)`
- New function `*RestorableSQLContainerGetResult.UnmarshalJSON([]byte) error`
- New function `*RestorableMongodbCollectionGetResult.UnmarshalJSON([]byte) error`
- New function `RestorableSQLDatabasesClient.ListPreparer(context.Context, string, string) (*http.Request, error)`
- New struct `RestorableLocationResource`
- New struct `RestorableMongodbCollectionGetResult`
- New struct `RestorableMongodbCollectionProperties`
- New struct `RestorableMongodbCollectionPropertiesResource`
- New struct `RestorableMongodbCollectionsClient`
- New struct `RestorableMongodbCollectionsListResult`
- New struct `RestorableMongodbDatabaseGetResult`
- New struct `RestorableMongodbDatabaseProperties`
- New struct `RestorableMongodbDatabasePropertiesResource`
- New struct `RestorableMongodbDatabasesClient`
- New struct `RestorableMongodbDatabasesListResult`
- New struct `RestorableMongodbResourcesClient`
- New struct `RestorableMongodbResourcesListResult`
- New struct `RestorableSQLContainerGetResult`
- New struct `RestorableSQLContainerProperties`
- New struct `RestorableSQLContainerPropertiesResource`
- New struct `RestorableSQLContainerPropertiesResourceContainer`
- New struct `RestorableSQLContainersClient`
- New struct `RestorableSQLContainersListResult`
- New struct `RestorableSQLDatabaseGetResult`
- New struct `RestorableSQLDatabaseProperties`
- New struct `RestorableSQLDatabasePropertiesResource`
- New struct `RestorableSQLDatabasePropertiesResourceDatabase`
- New struct `RestorableSQLDatabasesClient`
- New struct `RestorableSQLDatabasesListResult`
- New struct `RestorableSQLResourcesClient`
- New struct `RestorableSQLResourcesListResult`
- New field `RestorableLocations` in struct `RestorableDatabaseAccountProperties`
- New field `APIType` in struct `RestorableDatabaseAccountProperties`

Total 8 breaking change(s), 120 additive change(s).
