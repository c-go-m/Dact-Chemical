import { Base } from "./base.model"

export interface Menu extends Base{
    name?: string
    rute?: string
    type?: string
    position?: number
    
}