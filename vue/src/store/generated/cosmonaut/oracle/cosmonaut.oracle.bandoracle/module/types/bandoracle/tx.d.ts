import { Reader, Writer } from 'protobufjs/minimal';
import { FetchPriceCallData } from '../bandoracle/fetch_price';
import { Coin } from '../cosmos/base/v1beta1/coin';
export declare const protobufPackage = "cosmonaut.oracle.bandoracle";
export interface MsgFetchPriceData {
    creator: string;
    oracleScriptId: number;
    sourceChannel: string;
    calldata: FetchPriceCallData | undefined;
    askCount: number;
    minCount: number;
    feeLimit: Coin[];
    prepareGas: number;
    executeGas: number;
    clientId: string;
}
export interface MsgFetchPriceDataResponse {
}
export declare const MsgFetchPriceData: {
    encode(message: MsgFetchPriceData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFetchPriceData;
    fromJSON(object: any): MsgFetchPriceData;
    toJSON(message: MsgFetchPriceData): unknown;
    fromPartial(object: DeepPartial<MsgFetchPriceData>): MsgFetchPriceData;
};
export declare const MsgFetchPriceDataResponse: {
    encode(_: MsgFetchPriceDataResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFetchPriceDataResponse;
    fromJSON(_: any): MsgFetchPriceDataResponse;
    toJSON(_: MsgFetchPriceDataResponse): unknown;
    fromPartial(_: DeepPartial<MsgFetchPriceDataResponse>): MsgFetchPriceDataResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    FetchPriceData(request: MsgFetchPriceData): Promise<MsgFetchPriceDataResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    FetchPriceData(request: MsgFetchPriceData): Promise<MsgFetchPriceDataResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
