import { Banner } from "../shared/entities/banner.model";

export const mockListBanner: Banner[] = [
        {
            ID:1,
            CreatedAt:new Date("2022-08-04T21:49:31.22761-05:00"),
            UpdatedAt:new Date("2022-08-04T21:49:31.22761-05:00"),            
            DeletedAt:new Date("2022-08-04T21:49:31.22761-05:00"),
            name:"Trapero",
            attachedId:1,
            attached:{
                ID:1,
                CreatedAt:new Date("2022-08-04T21:49:31.225475-05:00"),
                UpdatedAt:new Date("2022-08-04T21:49:31.225475-05:00"),
                DeletedAt:new Date("2022-08-04T21:49:31.225475-05:00"),                
                name:"Banner1",
                url:"assets/img/banner/Banner1.jpg",
                extension:"jpg",
                data:null
            },
            position:1
        },{
            ID:2,
            CreatedAt:new Date("2022-08-04T21:50:10.171619-05:00"),
            UpdatedAt:new Date("2022-08-04T21:50:10.171619-05:00"),
            DeletedAt:new Date("2022-08-04T21:50:10.171619-05:00"),
            name:"Escoba",
            attachedId:2,
            attached:{
                ID:2,
                CreatedAt:new Date("2022-08-04T21:50:10.169398-05:00"),
                UpdatedAt:new Date("2022-08-04T21:50:10.169398-05:00"),
                DeletedAt:new Date("2022-08-04T21:50:10.169398-05:00"),
                name:"Banner2",
                url:"assets/img/banner/Banner2.jpg",
                extension:"jpg",
                data:null
            },
            position:2
        }
]


