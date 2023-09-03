import { Attached } from "./attached.model";
import { Base } from "./base.model";
import { Category } from "./category.model";
import { Presentation } from "./presentation.model";

export interface Product extends Base{
    
    name?: string
    attachedId?: number
    attached?: Attached
    categoryId?: number
    category?: Category
    presentation?: Presentation[]
}