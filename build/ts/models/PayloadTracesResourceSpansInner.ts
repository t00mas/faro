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

import { PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner } from '../models/PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner';
import { PayloadTracesResourceSpansInnerResource } from '../models/PayloadTracesResourceSpansInnerResource';
import { HttpFile } from '../http/http';

export class PayloadTracesResourceSpansInner {
    'resource': PayloadTracesResourceSpansInnerResource;
    'instrumentationLibrarySpans': Array<PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner>;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "resource",
            "baseName": "resource",
            "type": "PayloadTracesResourceSpansInnerResource",
            "format": ""
        },
        {
            "name": "instrumentationLibrarySpans",
            "baseName": "instrumentationLibrarySpans",
            "type": "Array<PayloadTracesResourceSpansInnerInstrumentationLibrarySpansInner>",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return PayloadTracesResourceSpansInner.attributeTypeMap;
    }

    public constructor() {
    }
}
