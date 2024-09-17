/**
 * Faro Collector API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { Event } from '../models/Event';
import { Exception } from '../models/Exception';
import { Log } from '../models/Log';
import { Measurement } from '../models/Measurement';
import { PayloadMeta } from '../models/PayloadMeta';
import { PayloadTraces } from '../models/PayloadTraces';
import { HttpFile } from '../http/http';

export class Payload {
    'events'?: Array<Event>;
    'logs'?: Array<Log>;
    'exceptions'?: Array<Exception>;
    'measurements'?: Array<Measurement>;
    'meta': PayloadMeta;
    'traces'?: PayloadTraces;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "events",
            "baseName": "events",
            "type": "Array<Event>",
            "format": ""
        },
        {
            "name": "logs",
            "baseName": "logs",
            "type": "Array<Log>",
            "format": ""
        },
        {
            "name": "exceptions",
            "baseName": "exceptions",
            "type": "Array<Exception>",
            "format": ""
        },
        {
            "name": "measurements",
            "baseName": "measurements",
            "type": "Array<Measurement>",
            "format": ""
        },
        {
            "name": "meta",
            "baseName": "meta",
            "type": "PayloadMeta",
            "format": ""
        },
        {
            "name": "traces",
            "baseName": "traces",
            "type": "PayloadTraces",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return Payload.attributeTypeMap;
    }

    public constructor() {
    }
}
