import { Attached } from "./attached.model";
import { Base } from "./base.model";
import { Menu } from "./menu.model";

export interface Information extends Base{
    name?: string
    value?:string
    attachedId?: number
    attached?: Attached
    menuId?: number
    menu?: Menu
    position?: number        
}