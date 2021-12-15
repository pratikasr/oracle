/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { FetchPriceCallData } from '../bandoracle/fetch_price'
import { Coin } from '../cosmos/base/v1beta1/coin'

export const protobufPackage = 'cosmonaut.oracle.bandoracle'

export interface MsgFetchPriceData {
  creator: string
  oracleScriptId: number
  sourceChannel: string
  calldata: FetchPriceCallData | undefined
  askCount: number
  minCount: number
  feeLimit: Coin[]
  prepareGas: number
  executeGas: number
  clientId: string
}

export interface MsgFetchPriceDataResponse {}

const baseMsgFetchPriceData: object = {
  creator: '',
  oracleScriptId: 0,
  sourceChannel: '',
  askCount: 0,
  minCount: 0,
  prepareGas: 0,
  executeGas: 0,
  clientId: ''
}

export const MsgFetchPriceData = {
  encode(message: MsgFetchPriceData, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.oracleScriptId !== 0) {
      writer.uint32(16).uint64(message.oracleScriptId)
    }
    if (message.sourceChannel !== '') {
      writer.uint32(26).string(message.sourceChannel)
    }
    if (message.calldata !== undefined) {
      FetchPriceCallData.encode(message.calldata, writer.uint32(34).fork()).ldelim()
    }
    if (message.askCount !== 0) {
      writer.uint32(40).uint64(message.askCount)
    }
    if (message.minCount !== 0) {
      writer.uint32(48).uint64(message.minCount)
    }
    for (const v of message.feeLimit) {
      Coin.encode(v!, writer.uint32(58).fork()).ldelim()
    }
    if (message.prepareGas !== 0) {
      writer.uint32(64).uint64(message.prepareGas)
    }
    if (message.executeGas !== 0) {
      writer.uint32(72).uint64(message.executeGas)
    }
    if (message.clientId !== '') {
      writer.uint32(82).string(message.clientId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFetchPriceData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgFetchPriceData } as MsgFetchPriceData
    message.feeLimit = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.oracleScriptId = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.sourceChannel = reader.string()
          break
        case 4:
          message.calldata = FetchPriceCallData.decode(reader, reader.uint32())
          break
        case 5:
          message.askCount = longToNumber(reader.uint64() as Long)
          break
        case 6:
          message.minCount = longToNumber(reader.uint64() as Long)
          break
        case 7:
          message.feeLimit.push(Coin.decode(reader, reader.uint32()))
          break
        case 8:
          message.prepareGas = longToNumber(reader.uint64() as Long)
          break
        case 9:
          message.executeGas = longToNumber(reader.uint64() as Long)
          break
        case 10:
          message.clientId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgFetchPriceData {
    const message = { ...baseMsgFetchPriceData } as MsgFetchPriceData
    message.feeLimit = []
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.oracleScriptId !== undefined && object.oracleScriptId !== null) {
      message.oracleScriptId = Number(object.oracleScriptId)
    } else {
      message.oracleScriptId = 0
    }
    if (object.sourceChannel !== undefined && object.sourceChannel !== null) {
      message.sourceChannel = String(object.sourceChannel)
    } else {
      message.sourceChannel = ''
    }
    if (object.calldata !== undefined && object.calldata !== null) {
      message.calldata = FetchPriceCallData.fromJSON(object.calldata)
    } else {
      message.calldata = undefined
    }
    if (object.askCount !== undefined && object.askCount !== null) {
      message.askCount = Number(object.askCount)
    } else {
      message.askCount = 0
    }
    if (object.minCount !== undefined && object.minCount !== null) {
      message.minCount = Number(object.minCount)
    } else {
      message.minCount = 0
    }
    if (object.feeLimit !== undefined && object.feeLimit !== null) {
      for (const e of object.feeLimit) {
        message.feeLimit.push(Coin.fromJSON(e))
      }
    }
    if (object.prepareGas !== undefined && object.prepareGas !== null) {
      message.prepareGas = Number(object.prepareGas)
    } else {
      message.prepareGas = 0
    }
    if (object.executeGas !== undefined && object.executeGas !== null) {
      message.executeGas = Number(object.executeGas)
    } else {
      message.executeGas = 0
    }
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = String(object.clientId)
    } else {
      message.clientId = ''
    }
    return message
  },

  toJSON(message: MsgFetchPriceData): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.oracleScriptId !== undefined && (obj.oracleScriptId = message.oracleScriptId)
    message.sourceChannel !== undefined && (obj.sourceChannel = message.sourceChannel)
    message.calldata !== undefined && (obj.calldata = message.calldata ? FetchPriceCallData.toJSON(message.calldata) : undefined)
    message.askCount !== undefined && (obj.askCount = message.askCount)
    message.minCount !== undefined && (obj.minCount = message.minCount)
    if (message.feeLimit) {
      obj.feeLimit = message.feeLimit.map((e) => (e ? Coin.toJSON(e) : undefined))
    } else {
      obj.feeLimit = []
    }
    message.prepareGas !== undefined && (obj.prepareGas = message.prepareGas)
    message.executeGas !== undefined && (obj.executeGas = message.executeGas)
    message.clientId !== undefined && (obj.clientId = message.clientId)
    return obj
  },

  fromPartial(object: DeepPartial<MsgFetchPriceData>): MsgFetchPriceData {
    const message = { ...baseMsgFetchPriceData } as MsgFetchPriceData
    message.feeLimit = []
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.oracleScriptId !== undefined && object.oracleScriptId !== null) {
      message.oracleScriptId = object.oracleScriptId
    } else {
      message.oracleScriptId = 0
    }
    if (object.sourceChannel !== undefined && object.sourceChannel !== null) {
      message.sourceChannel = object.sourceChannel
    } else {
      message.sourceChannel = ''
    }
    if (object.calldata !== undefined && object.calldata !== null) {
      message.calldata = FetchPriceCallData.fromPartial(object.calldata)
    } else {
      message.calldata = undefined
    }
    if (object.askCount !== undefined && object.askCount !== null) {
      message.askCount = object.askCount
    } else {
      message.askCount = 0
    }
    if (object.minCount !== undefined && object.minCount !== null) {
      message.minCount = object.minCount
    } else {
      message.minCount = 0
    }
    if (object.feeLimit !== undefined && object.feeLimit !== null) {
      for (const e of object.feeLimit) {
        message.feeLimit.push(Coin.fromPartial(e))
      }
    }
    if (object.prepareGas !== undefined && object.prepareGas !== null) {
      message.prepareGas = object.prepareGas
    } else {
      message.prepareGas = 0
    }
    if (object.executeGas !== undefined && object.executeGas !== null) {
      message.executeGas = object.executeGas
    } else {
      message.executeGas = 0
    }
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = object.clientId
    } else {
      message.clientId = ''
    }
    return message
  }
}

const baseMsgFetchPriceDataResponse: object = {}

export const MsgFetchPriceDataResponse = {
  encode(_: MsgFetchPriceDataResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFetchPriceDataResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgFetchPriceDataResponse } as MsgFetchPriceDataResponse
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

  fromJSON(_: any): MsgFetchPriceDataResponse {
    const message = { ...baseMsgFetchPriceDataResponse } as MsgFetchPriceDataResponse
    return message
  },

  toJSON(_: MsgFetchPriceDataResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgFetchPriceDataResponse>): MsgFetchPriceDataResponse {
    const message = { ...baseMsgFetchPriceDataResponse } as MsgFetchPriceDataResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FetchPriceData(request: MsgFetchPriceData): Promise<MsgFetchPriceDataResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  FetchPriceData(request: MsgFetchPriceData): Promise<MsgFetchPriceDataResponse> {
    const data = MsgFetchPriceData.encode(request).finish()
    const promise = this.rpc.request('cosmonaut.oracle.bandoracle.Msg', 'FetchPriceData', data)
    return promise.then((data) => MsgFetchPriceDataResponse.decode(new Reader(data)))
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
