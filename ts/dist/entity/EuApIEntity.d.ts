import { DutchCustomerDataEntityBase } from '../DutchCustomerDataEntityBase';
import type { DutchCustomerDataSDK } from '../DutchCustomerDataSDK';
import type { Control } from '../types';
import type { EuApI, EuApILoadMatch, EuApIListMatch } from '../DutchCustomerDataTypes';
declare class EuApIEntity extends DutchCustomerDataEntityBase<EuApI> {
    constructor(client: DutchCustomerDataSDK, entopts: any);
    make(this: EuApIEntity): EuApIEntity;
    load(this: any, reqmatch?: EuApILoadMatch, ctrl?: Control): Promise<EuApIEntity>;
    list(this: any, reqmatch?: EuApIListMatch, ctrl?: Control): Promise<EuApIEntity[]>;
}
export { EuApIEntity };
