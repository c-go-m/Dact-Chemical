import { Attached } from "./attached.model";
import { Base } from "./base.model";
import { Product } from "./product.model";

export interface Presentation extends Base{
    name?: string
    attachedId?: number
    attached?: Attached
    productId?: number
    product?: Product
    position?: number
    cost?: number
    
}