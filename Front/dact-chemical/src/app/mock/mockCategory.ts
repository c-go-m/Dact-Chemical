import { Category } from "../shared/entities/category.model";

export const mockListCategory: Category[] = [
        {
            ID:1,
            CreatedAt:new Date("2022-08-04T21:49:31.22761-05:00"),
            UpdatedAt:new Date("2022-08-04T21:49:31.22761-05:00"),            
            DeletedAt:new Date("2022-08-04T21:49:31.22761-05:00"),
            name:"Hogar",
            attachedId:1,
            attached:{
                ID:1,
                CreatedAt:new Date("2022-08-04T21:49:31.225475-05:00"),
                UpdatedAt:new Date("2022-08-04T21:49:31.225475-05:00"),
                DeletedAt:new Date("2022-08-04T21:49:31.225475-05:00"),                
                name:"Categoria1",
                url:"assets/img/categoria/categoria1.jpg",
                extension:"jpg",
                data:null
            },
            position:1
        },{
            ID:2,
            CreatedAt:new Date("2022-08-04T21:50:10.171619-05:00"),
            UpdatedAt:new Date("2022-08-04T21:50:10.171619-05:00"),
            DeletedAt:new Date("2022-08-04T21:50:10.171619-05:00"),
            name:"Automovil",
            attachedId:2,
            attached:{
                ID:2,
                CreatedAt:new Date("2022-08-04T21:50:10.169398-05:00"),
                UpdatedAt:new Date("2022-08-04T21:50:10.169398-05:00"),
                DeletedAt:new Date("2022-08-04T21:50:10.169398-05:00"),
                name:"Categoria1",
                url:"assets/img/categoria/categoria2.jpg",
                extension:"jpg",
                data:null
            },
            position:2
        }
]


