/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { Params } from '../bandoracle/params'
import { FetchPriceResult } from '../bandoracle/fetch_price'

export const protobufPackage = 'cosmonaut.oracle.bandoracle'

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined
}

export interface QueryFetchPriceRequest {
  requestId: number
}

export interface QueryFetchPriceResponse {
  result: FetchPriceResult | undefined
}

export interface QueryLastFetchPriceIdRequest {}

export interface QueryLastFetchPriceIdResponse {
  requestId: number
}

const baseQueryParamsRequest: object = {}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest
    return message
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest
    return message
  }
}

const baseQueryParamsResponse: object = {}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params)
    } else {
      message.params = undefined
    }
    return message
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {}
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params)
    } else {
      message.params = undefined
    }
    return message
  }
}

const baseQueryFetchPriceRequest: object = { requestId: 0 }

export const QueryFetchPriceRequest = {
  encode(message: QueryFetchPriceRequest, writer: Writer = Writer.create()): Writer {
    if (message.requestId !== 0) {
      writer.uint32(8).int64(message.requestId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFetchPriceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryFetchPriceRequest } as QueryFetchPriceRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.requestId = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryFetchPriceRequest {
    const message = { ...baseQueryFetchPriceRequest } as QueryFetchPriceRequest
    if (object.requestId !== undefined && object.requestId !== null) {
      message.requestId = Number(object.requestId)
    } else {
      message.requestId = 0
    }
    return message
  },

  toJSON(message: QueryFetchPriceRequest): unknown {
    const obj: any = {}
    message.requestId !== undefined && (obj.requestId = message.requestId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryFetchPriceRequest>): QueryFetchPriceRequest {
    const message = { ...baseQueryFetchPriceRequest } as QueryFetchPriceRequest
    if (object.requestId !== undefined && object.requestId !== null) {
      message.requestId = object.requestId
    } else {
      message.requestId = 0
    }
    return message
  }
}

const baseQueryFetchPriceResponse: object = {}

export const QueryFetchPriceResponse = {
  encode(message: QueryFetchPriceResponse, writer: Writer = Writer.create()): Writer {
    if (message.result !== undefined) {
      FetchPriceResult.encode(message.result, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryFetchPriceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryFetchPriceResponse } as QueryFetchPriceResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.result = FetchPriceResult.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryFetchPriceResponse {
    const message = { ...baseQueryFetchPriceResponse } as QueryFetchPriceResponse
    if (object.result !== undefined && object.result !== null) {
      message.result = FetchPriceResult.fromJSON(object.result)
    } else {
      message.result = undefined
    }
    return message
  },

  toJSON(message: QueryFetchPriceResponse): unknown {
    const obj: any = {}
    message.result !== undefined && (obj.result = message.result ? FetchPriceResult.toJSON(message.result) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryFetchPriceResponse>): QueryFetchPriceResponse {
    const message = { ...baseQueryFetchPriceResponse } as QueryFetchPriceResponse
    if (object.result !== undefined && object.result !== null) {
      message.result = FetchPriceResult.fromPartial(object.result)
    } else {
      message.result = undefined
    }
    return message
  }
}

const baseQueryLastFetchPriceIdRequest: object = {}

export const QueryLastFetchPriceIdRequest = {
  encode(_: QueryLastFetchPriceIdRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryLastFetchPriceIdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryLastFetchPriceIdRequest } as QueryLastFetchPriceIdRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryLastFetchPriceIdRequest {
    const message = { ...baseQueryLastFetchPriceIdRequest } as QueryLastFetchPriceIdRequest
    return message
  },

  toJSON(_: QueryLastFetchPriceIdRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryLastFetchPriceIdRequest>): QueryLastFetchPriceIdRequest {
    const message = { ...baseQueryLastFetchPriceIdRequest } as QueryLastFetchPriceIdRequest
    return message
  }
}

const baseQueryLastFetchPriceIdResponse: object = { requestId: 0 }

export const QueryLastFetchPriceIdResponse = {
  encode(message: QueryLastFetchPriceIdResponse, writer: Writer = Writer.create()): Writer {
    if (message.requestId !== 0) {
      writer.uint32(8).int64(message.requestId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryLastFetchPriceIdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryLastFetchPriceIdResponse } as QueryLastFetchPriceIdResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.requestId = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryLastFetchPriceIdResponse {
    const message = { ...baseQueryLastFetchPriceIdResponse } as QueryLastFetchPriceIdResponse
    if (object.requestId !== undefined && object.requestId !== null) {
      message.requestId = Number(object.requestId)
    } else {
      message.requestId = 0
    }
    return message
  },

  toJSON(message: QueryLastFetchPriceIdResponse): unknown {
    const obj: any = {}
    message.requestId !== undefined && (obj.requestId = message.requestId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryLastFetchPriceIdResponse>): QueryLastFetchPriceIdResponse {
    const message = { ...baseQueryLastFetchPriceIdResponse } as QueryLastFetchPriceIdResponse
    if (object.requestId !== undefined && object.requestId !== null) {
      message.requestId = object.requestId
    } else {
      message.requestId = 0
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>
  /** FetchPriceResult defines a rpc handler method for MsgFetchPriceData. */
  FetchPriceResult(request: QueryFetchPriceRequest): Promise<QueryFetchPriceResponse>
  /** LastFetchPriceId query the last FetchPrice result id */
  LastFetchPriceId(request: QueryLastFetchPriceIdRequest): Promise<QueryLastFetchPriceIdResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish()
    const promise = this.rpc.request('cosmonaut.oracle.bandoracle.Query', 'Params', data)
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)))
  }

  FetchPriceResult(request: QueryFetchPriceRequest): Promise<QueryFetchPriceResponse> {
    const data = QueryFetchPriceRequest.encode(request).finish()
    const promise = this.rpc.request('cosmonaut.oracle.bandoracle.Query', 'FetchPriceResult', data)
    return promise.then((data) => QueryFetchPriceResponse.decode(new Reader(data)))
  }

  LastFetchPriceId(request: QueryLastFetchPriceIdRequest): Promise<QueryLastFetchPriceIdResponse> {
    const data = QueryLastFetchPriceIdRequest.encode(request).finish()
    const promise = this.rpc.request('cosmonaut.oracle.bandoracle.Query', 'LastFetchPriceId', data)
    return promise.then((data) => QueryLastFetchPriceIdResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
