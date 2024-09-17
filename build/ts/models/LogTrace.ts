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

import { HttpFile } from '../http/http';

export class LogTrace {
    'traceId': string;
    'spanId': string;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "traceId",
            "baseName": "trace_id",
            "type": "string",
            "format": ""
        },
        {
            "name": "spanId",
            "baseName": "span_id",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return LogTrace.attributeTypeMap;
    }

    public constructor() {
    }
}
