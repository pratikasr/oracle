import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "cosmonaut.oracle.bandoracle";
export interface FetchPriceCallData {
    symbols: string[];
    multiplier: number;
}
export interface FetchPriceResult {
    rates: number[];
}
export declare const FetchPriceCallData: {
    encode(message: FetchPriceCallData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): FetchPriceCallData;
    fromJSON(object: any): FetchPriceCallData;
    toJSON(message: FetchPriceCallData): unknown;
    fromPartial(object: DeepPartial<FetchPriceCallData>): FetchPriceCallData;
};
export declare const FetchPriceResult: {
    encode(message: FetchPriceResult, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): FetchPriceResult;
    fromJSON(object: any): FetchPriceResult;
    toJSON(message: FetchPriceResult): unknown;
    fromPartial(object: DeepPartial<FetchPriceResult>): FetchPriceResult;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
