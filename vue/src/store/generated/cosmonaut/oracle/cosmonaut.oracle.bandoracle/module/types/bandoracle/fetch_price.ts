/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'cosmonaut.oracle.bandoracle'

export interface FetchPriceCallData {
  symbols: string[]
  multiplier: number
}

export interface FetchPriceResult {
  rates: number[]
}

const baseFetchPriceCallData: object = { symbols: '', multiplier: 0 }

export const FetchPriceCallData = {
  encode(message: FetchPriceCallData, writer: Writer = Writer.create()): Writer {
    for (const v of message.symbols) {
      writer.uint32(10).string(v!)
    }
    if (message.multiplier !== 0) {
      writer.uint32(16).uint64(message.multiplier)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): FetchPriceCallData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseFetchPriceCallData } as FetchPriceCallData
    message.symbols = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.symbols.push(reader.string())
          break
        case 2:
          message.multiplier = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): FetchPriceCallData {
    const message = { ...baseFetchPriceCallData } as FetchPriceCallData
    message.symbols = []
    if (object.symbols !== undefined && object.symbols !== null) {
      for (const e of object.symbols) {
        message.symbols.push(String(e))
      }
    }
    if (object.multiplier !== undefined && object.multiplier !== null) {
      message.multiplier = Number(object.multiplier)
    } else {
      message.multiplier = 0
    }
    return message
  },

  toJSON(message: FetchPriceCallData): unknown {
    const obj: any = {}
    if (message.symbols) {
      obj.symbols = message.symbols.map((e) => e)
    } else {
      obj.symbols = []
    }
    message.multiplier !== undefined && (obj.multiplier = message.multiplier)
    return obj
  },

  fromPartial(object: DeepPartial<FetchPriceCallData>): FetchPriceCallData {
    const message = { ...baseFetchPriceCallData } as FetchPriceCallData
    message.symbols = []
    if (object.symbols !== undefined && object.symbols !== null) {
      for (const e of object.symbols) {
        message.symbols.push(e)
      }
    }
    if (object.multiplier !== undefined && object.multiplier !== null) {
      message.multiplier = object.multiplier
    } else {
      message.multiplier = 0
    }
    return message
  }
}

const baseFetchPriceResult: object = { rates: 0 }

export const FetchPriceResult = {
  encode(message: FetchPriceResult, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork()
    for (const v of message.rates) {
      writer.uint64(v)
    }
    writer.ldelim()
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): FetchPriceResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseFetchPriceResult } as FetchPriceResult
    message.rates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos
            while (reader.pos < end2) {
              message.rates.push(longToNumber(reader.uint64() as Long))
            }
          } else {
            message.rates.push(longToNumber(reader.uint64() as Long))
          }
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): FetchPriceResult {
    const message = { ...baseFetchPriceResult } as FetchPriceResult
    message.rates = []
    if (object.rates !== undefined && object.rates !== null) {
      for (const e of object.rates) {
        message.rates.push(Number(e))
      }
    }
    return message
  },

  toJSON(message: FetchPriceResult): unknown {
    const obj: any = {}
    if (message.rates) {
      obj.rates = message.rates.map((e) => e)
    } else {
      obj.rates = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<FetchPriceResult>): FetchPriceResult {
    const message = { ...baseFetchPriceResult } as FetchPriceResult
    message.rates = []
    if (object.rates !== undefined && object.rates !== null) {
      for (const e of object.rates) {
        message.rates.push(e)
      }
    }
    return message
  }
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
