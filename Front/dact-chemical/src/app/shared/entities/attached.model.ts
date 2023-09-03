import { Base } from "./base.model"

export interface Attached extends Base{
    name? : string
    url? : string
    extension? : string
    data? : any    
}