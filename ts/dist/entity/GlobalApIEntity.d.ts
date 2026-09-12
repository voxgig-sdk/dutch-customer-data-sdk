import { DutchCustomerDataEntityBase } from '../DutchCustomerDataEntityBase';
import type { DutchCustomerDataSDK } from '../DutchCustomerDataSDK';
import type { Control } from '../types';
import type { GlobalApI, GlobalApILoadMatch, GlobalApIListMatch, GlobalApICreateData } from '../DutchCustomerDataTypes';
declare class GlobalApIEntity extends DutchCustomerDataEntityBase<GlobalApI> {
    constructor(client: DutchCustomerDataSDK, entopts: any);
    make(this: GlobalApIEntity): GlobalApIEntity;
    load(this: any, reqmatch?: GlobalApILoadMatch, ctrl?: Control): Promise<GlobalApIEntity>;
    list(this: any, reqmatch?: GlobalApIListMatch, ctrl?: Control): Promise<GlobalApIEntity[]>;
    create(this: any, reqdata?: GlobalApICreateData, ctrl?: Control): Promise<GlobalApIEntity>;
}
export { GlobalApIEntity };
