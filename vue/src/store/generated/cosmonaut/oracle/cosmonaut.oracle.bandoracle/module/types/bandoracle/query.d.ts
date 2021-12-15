import { Reader, Writer } from 'protobufjs/minimal';
import { Params } from '../bandoracle/params';
import { FetchPriceResult } from '../bandoracle/fetch_price';
export declare const protobufPackage = "cosmonaut.oracle.bandoracle";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryFetchPriceRequest {
    requestId: number;
}
export interface QueryFetchPriceResponse {
    result: FetchPriceResult | undefined;
}
export interface QueryLastFetchPriceIdRequest {
}
export interface QueryLastFetchPriceIdResponse {
    requestId: number;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryFetchPriceRequest: {
    encode(message: QueryFetchPriceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryFetchPriceRequest;
    fromJSON(object: any): QueryFetchPriceRequest;
    toJSON(message: QueryFetchPriceRequest): unknown;
    fromPartial(object: DeepPartial<QueryFetchPriceRequest>): QueryFetchPriceRequest;
};
export declare const QueryFetchPriceResponse: {
    encode(message: QueryFetchPriceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryFetchPriceResponse;
    fromJSON(object: any): QueryFetchPriceResponse;
    toJSON(message: QueryFetchPriceResponse): unknown;
    fromPartial(object: DeepPartial<QueryFetchPriceResponse>): QueryFetchPriceResponse;
};
export declare const QueryLastFetchPriceIdRequest: {
    encode(_: QueryLastFetchPriceIdRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryLastFetchPriceIdRequest;
    fromJSON(_: any): QueryLastFetchPriceIdRequest;
    toJSON(_: QueryLastFetchPriceIdRequest): unknown;
    fromPartial(_: DeepPartial<QueryLastFetchPriceIdRequest>): QueryLastFetchPriceIdRequest;
};
export declare const QueryLastFetchPriceIdResponse: {
    encode(message: QueryLastFetchPriceIdResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryLastFetchPriceIdResponse;
    fromJSON(object: any): QueryLastFetchPriceIdResponse;
    toJSON(message: QueryLastFetchPriceIdResponse): unknown;
    fromPartial(object: DeepPartial<QueryLastFetchPriceIdResponse>): QueryLastFetchPriceIdResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** FetchPriceResult defines a rpc handler method for MsgFetchPriceData. */
    FetchPriceResult(request: QueryFetchPriceRequest): Promise<QueryFetchPriceResponse>;
    /** LastFetchPriceId query the last FetchPrice result id */
    LastFetchPriceId(request: QueryLastFetchPriceIdRequest): Promise<QueryLastFetchPriceIdResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    FetchPriceResult(request: QueryFetchPriceRequest): Promise<QueryFetchPriceResponse>;
    LastFetchPriceId(request: QueryLastFetchPriceIdRequest): Promise<QueryLastFetchPriceIdResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
