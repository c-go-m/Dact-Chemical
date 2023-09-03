import { Attached } from "./attached.model";
import { Base } from "./base.model";

export interface Category extends Base{
    
    name?: string
    attachedId?: number
    attached?: Attached
    position?: number

}