export interface BandoracleFetchPriceCallData {
    symbols?: string[];
    /** @format uint64 */
    multiplier?: string;
}
export interface BandoracleFetchPriceResult {
    rates?: string[];
}
export declare type BandoracleMsgFetchPriceDataResponse = object;
/**
 * Params defines the parameters for the module.
 */
export declare type BandoracleParams = object;
export interface BandoracleQueryFetchPriceResponse {
    result?: BandoracleFetchPriceResult;
}
export interface BandoracleQueryLastFetchPriceIdResponse {
    /** @format int64 */
    requestId?: string;
}
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface BandoracleQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: BandoracleParams;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
    denom?: string;
    amount?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title bandoracle/fetch_price.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/cosmonaut/oracle/bandoracle/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<BandoracleQueryParamsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryFetchPriceResult
     * @summary FetchPriceResult defines a rpc handler method for MsgFetchPriceData.
     * @request GET:/oracle/bandoracle/fetch_price_result/{requestId}
     */
    queryFetchPriceResult: (requestId: string, params?: RequestParams) => Promise<HttpResponse<BandoracleQueryFetchPriceResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryLastFetchPriceId
     * @summary LastFetchPriceId query the last FetchPrice result id
     * @request GET:/oracle/bandoracle/last_fetch_price_id
     */
    queryLastFetchPriceId: (params?: RequestParams) => Promise<HttpResponse<BandoracleQueryLastFetchPriceIdResponse, RpcStatus>>;
}
export {};
