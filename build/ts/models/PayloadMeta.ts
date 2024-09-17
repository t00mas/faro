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

import { AppMeta } from '../models/AppMeta';
import { BrowserMeta } from '../models/BrowserMeta';
import { PageMeta } from '../models/PageMeta';
import { PayloadMetaView } from '../models/PayloadMetaView';
import { SdkMeta } from '../models/SdkMeta';
import { SessionMeta } from '../models/SessionMeta';
import { UserMeta } from '../models/UserMeta';
import { HttpFile } from '../http/http';

export class PayloadMeta {
    'sdk'?: SdkMeta;
    'app'?: AppMeta;
    'user'?: UserMeta;
    'session'?: SessionMeta;
    'page'?: PageMeta;
    'browser'?: BrowserMeta;
    'view'?: PayloadMetaView;

    static readonly discriminator: string | undefined = undefined;

    static readonly mapping: {[index: string]: string} | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "sdk",
            "baseName": "sdk",
            "type": "SdkMeta",
            "format": ""
        },
        {
            "name": "app",
            "baseName": "app",
            "type": "AppMeta",
            "format": ""
        },
        {
            "name": "user",
            "baseName": "user",
            "type": "UserMeta",
            "format": ""
        },
        {
            "name": "session",
            "baseName": "session",
            "type": "SessionMeta",
            "format": ""
        },
        {
            "name": "page",
            "baseName": "page",
            "type": "PageMeta",
            "format": ""
        },
        {
            "name": "browser",
            "baseName": "browser",
            "type": "BrowserMeta",
            "format": ""
        },
        {
            "name": "view",
            "baseName": "view",
            "type": "PayloadMetaView",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return PayloadMeta.attributeTypeMap;
    }

    public constructor() {
    }
}
