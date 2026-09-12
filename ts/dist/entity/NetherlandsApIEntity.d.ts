import { DutchCustomerDataEntityBase } from '../DutchCustomerDataEntityBase';
import type { DutchCustomerDataSDK } from '../DutchCustomerDataSDK';
import type { Control } from '../types';
import type { NetherlandsApI, NetherlandsApIListMatch } from '../DutchCustomerDataTypes';
declare class NetherlandsApIEntity extends DutchCustomerDataEntityBase<NetherlandsApI> {
    constructor(client: DutchCustomerDataSDK, entopts: any);
    make(this: NetherlandsApIEntity): NetherlandsApIEntity;
    list(this: any, reqmatch?: NetherlandsApIListMatch, ctrl?: Control): Promise<NetherlandsApIEntity[]>;
}
export { NetherlandsApIEntity };
