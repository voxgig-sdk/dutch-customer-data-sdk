import { EuApIEntity } from './entity/EuApIEntity';
import { GlobalApIEntity } from './entity/GlobalApIEntity';
import { NetherlandsApIEntity } from './entity/NetherlandsApIEntity';
export type * from './DutchCustomerDataTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { DutchCustomerDataEntityBase } from './DutchCustomerDataEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class DutchCustomerDataSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    EuApI(entopts?: Record<string, any>): EuApIEntity;
    GlobalApI(entopts?: Record<string, any>): GlobalApIEntity;
    NetherlandsApI(entopts?: Record<string, any>): NetherlandsApIEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): DutchCustomerDataSDK;
    tester(testopts?: any, sdkopts?: any): DutchCustomerDataSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof DutchCustomerDataSDK;
export { stdutil, config, BaseFeature, DutchCustomerDataEntityBase, DutchCustomerDataSDK, SDK, };
